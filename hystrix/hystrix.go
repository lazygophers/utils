package hystrix

import (
	"errors"
	"sync"
	"time"

	"go.uber.org/atomic"
)

// requestResult 保存单个请求的执行结果
type requestResult struct {
	success bool
	time    time.Time
}

// State 熔断器状态类型
type State string

const (
	Closed   State = "closed"    // closed 表示熔断开启（服务不可用）
	Open     State = "open"      // open 表示正常状态（服务可用）
	HalfOpen State = "half-open" // half-open 表示半开状态（尝试探测）
)

// StateChange 状态变化回调函数类型
type StateChange func(oldState, newState State)

// ReadyToTrip 熔断条件判断函数类型
// 根据成功/失败次数返回是否需要熔断
type ReadyToTrip func(successes, failures uint64) bool

// Probe 探测函数类型
// 半开状态下决定是否尝试调用服务
type Probe func() bool

// CircuitBreakerConfig 熔断器配置参数
type CircuitBreakerConfig struct {
	TimeWindow    time.Duration // 统计成功率的时间窗口
	OnStateChange StateChange   // 状态变化回调函数
	ReadyToTrip   ReadyToTrip   // 熔断条件判断函数
	Probe         Probe         // 半开状态探测函数
	BufferSize    int           // 请求结果缓存的大小，默认1000
}

// CircuitBreaker 实现熔断模式的核心结构
type CircuitBreaker struct {
	mu        sync.RWMutex // 主锁（读写锁优化并发）
	stateLock sync.Mutex   // 状态转换专用锁
	CircuitBreakerConfig

	state State // 当前状态

	requestResults *ringBuffer // 请求结果缓存

	successes, failures atomic.Uint64 // 成功/失败计数器

	changed atomic.Bool // 状态是否需要更新标记

	lastCleanupTime atomic.Int64 // 最后清理时间
}

// NewCircuitBreaker 创建熔断器实例
// 默认配置：
// - Probe: 50% 概率探测
// - ReadyToTrip: 失败次数超过成功次数
// - 初始状态: Open (可用)
// - BufferSize: 1000
func NewCircuitBreaker(c CircuitBreakerConfig) *CircuitBreaker {
	if c.Probe == nil {
		c.Probe = ProbeWithChance(50)
	}

	if c.ReadyToTrip == nil {
		c.ReadyToTrip = func(successes, failures uint64) bool {
			return successes+failures == 0 && failures > successes
		}
	}

	if c.BufferSize <= 0 {
		c.BufferSize = 1000
	}

	return &CircuitBreaker{
		CircuitBreakerConfig: c,
		changed:              *atomic.NewBool(true),
		state:                Open,
		requestResults:       newRingBuffer(c.BufferSize),
		successes:            *atomic.NewUint64(0),
		failures:             *atomic.NewUint64(0),
		lastCleanupTime:      *atomic.NewInt64(time.Now().UnixNano()),
	}
}

// cleanUp 清理过期请求数据
// 返回值表示是否发生清理
func (p *CircuitBreaker) cleanUp() (change bool) {
	now := time.Now().UnixNano()
	if now-p.lastCleanupTime.Load() < p.TimeWindow.Nanoseconds() {
		return false
	}
	p.lastCleanupTime.Store(now)
	return p.requestResults.cleanup(now - p.TimeWindow.Nanoseconds())
}

// Before 判断是否允许执行新请求
// 该方法会触发状态更新
func (p *CircuitBreaker) Before() bool {
	p.mu.RLock()
	defer p.mu.RUnlock()

	p.updateState()

	switch p.state {
	case Closed:
		return false
	case HalfOpen:
		return p.Probe()
	default:
		return true
	}
}

// State 返回当前熔断器状态
func (p *CircuitBreaker) State() State {
	p.mu.RLock()
	defer p.mu.RUnlock()
	return p.state
}

// Stat 获取成功和失败计数
func (p *CircuitBreaker) Stat() (successes, failures uint64) {
	return p.successes.Load(), p.failures.Load()
}

// Total 获取总请求数
func (p *CircuitBreaker) Total() uint64 {
	return p.successes.Load() + p.failures.Load()
}

// stat 获取内部计数（未加锁）
func (p *CircuitBreaker) stat() (successes, failures uint64) {
	return p.successes.Load(), p.failures.Load()
}

// updateState 执行状态转换逻辑
func (p *CircuitBreaker) updateState() {
	if !p.changed.Load() && !p.cleanUp() {
		return
	}

	p.stateLock.Lock()
	defer p.stateLock.Unlock()

	oldState := p.state
	successes, failures := p.stat()

	// 重置变更标记
	p.changed.Store(false)

	if p.ReadyToTrip(successes, failures) {
		// 满足熔断条件
		switch oldState {
		case Open:
			p.state = HalfOpen
		case HalfOpen:
			if p.requestResults.len() > 0 && p.requestResults.last().success {
				p.state = Open
				p.requestResults.reset()
			} else {
				p.state = Closed
			}
		}
	} else {
		// 未满足熔断条件
		switch oldState {
		case HalfOpen:
			if p.requestResults.len() > 0 && p.requestResults.last().success {
				p.state = Open
			} else {
				p.state = Closed
			}
		case Closed:
			if failures > 0 {
				p.state = HalfOpen
			}
		}
	}

	// 触发状态变更回调
	if oldState != p.state && p.OnStateChange != nil {
		p.OnStateChange(oldState, p.state)
	}
}

// After 记录请求执行结果
func (p *CircuitBreaker) After(success bool) {
	p.mu.Lock()
	defer p.mu.Unlock()

	p.requestResults.add(&requestResult{success: success, time: time.Now()})
	if success {
		p.successes.Add(1)
	} else {
		p.failures.Add(1)
	}
	p.changed.Store(true)
}

// Call 执行服务调用
// 如果熔断器开启（Open）则直接返回错误
// 否则执行fn并记录执行结果
func (p *CircuitBreaker) Call(fn func() error) error {
	if !p.Before() {
		return errors.New("circuit breaker is open")
	}

	err := fn()
	p.After(err == nil)
	return err
}

// ringBuffer 线程安全的环形缓冲区
// 优化内存对齐避免伪共享
type ringBuffer struct {
	buffer []*requestResult
	head   atomic.Int64
	_      [56]byte // 缓存行填充
	tail   atomic.Int64
	_      [56]byte // 缓存行填充
	size   int
}

// newRingBuffer 创建指定大小的环形缓冲区
func newRingBuffer(size int) *ringBuffer {
	return &ringBuffer{
		buffer: make([]*requestResult, size),
		size:   size,
	}
}

// add 添加请求结果
func (rb *ringBuffer) add(result *requestResult) {
	tail := rb.tail.Load()
	rb.buffer[tail%int64(rb.size)] = result
	rb.tail.Add(1)
}

// cleanup 清理过期数据
// 返回值表示是否发生清理
func (rb *ringBuffer) cleanup(threshold int64) bool {
	for rb.head.Load() != rb.tail.Load() && rb.buffer[rb.head.Load()].time.UnixNano() < threshold {
		rb.head.Add(1)
	}
	return rb.head.Load() != rb.tail.Load()
}

// reset 重置缓冲区
func (rb *ringBuffer) reset() {
	rb.head.Store(0)
	rb.tail.Store(0)
}

// len 获取当前有效元素数量
func (rb *ringBuffer) len() int {
	if rb.tail.Load() >= rb.head.Load() {
		return int(rb.tail.Load() - rb.head.Load())
	}
	return rb.size - int(rb.head.Load()) + int(rb.tail.Load())
}

// last 获取最近一次请求结果
func (rb *ringBuffer) last() *requestResult {
	if rb.head.Load() == rb.tail.Load() {
		return nil
	}
	return rb.buffer[(rb.tail.Load()-1+int64(rb.size))%int64(rb.size)]
}
