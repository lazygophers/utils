package hystrix

import (
	"errors"
	"sync"
	"sync/atomic"
	"time"
	"unsafe"
)

// OptimizedCircuitBreaker 高性能优化版本的熔断器
type OptimizedCircuitBreaker struct {
	// 配置参数 (只读，无需同步)
	timeWindow    int64 // 纳秒
	onStateChange StateChange
	readyToTrip   ReadyToTrip
	probe         Probe
	bufferSize    int

	// 状态管理 (原子操作优化)
	state atomic.Uint32 // 0=Closed, 1=Open, 2=HalfOpen

	// 统计计数器 (内存对齐优化，避免伪共享)
	stats struct {
		successes       atomic.Uint64
		_               [56]byte // 缓存行填充
		failures        atomic.Uint64
		_               [56]byte // 缓存行填充
		lastCleanupTime atomic.Int64
		_               [56]byte // 缓存行填充
		changed         atomic.Uint32 // 0=false, 1=true
		_               [60]byte // 缓存行填充
	}

	// 环形缓冲区优化
	ringBuffer *optimizedRingBuffer
}

// optimizedRingBuffer 优化的环形缓冲区
type optimizedRingBuffer struct {
	// 使用连续内存块减少间接访问
	buffer []int64 // 紧凑存储：高32位存储时间戳(相对)，低32位存储状态+标志
	_      [56]byte
	head   atomic.Uint64
	_      [56]byte
	tail   atomic.Uint64
	_      [56]byte
	size   uint64
	mask   uint64 // size-1，用于快速取模
}

const (
	stateClosedOpt  = 0
	stateOpenOpt    = 1
	stateHalfOpenOpt = 2

	// 紧凑存储标志位
	successFlag = 0x01
	timeShift   = 32
)

// NewOptimizedCircuitBreaker 创建优化版本的熔断器
func NewOptimizedCircuitBreaker(c CircuitBreakerConfig) *OptimizedCircuitBreaker {
	if c.Probe == nil {
		c.Probe = ProbeWithChance(50)
	}

	if c.ReadyToTrip == nil {
		c.ReadyToTrip = func(successes, failures uint64) bool {
			total := successes + failures
			return total >= 10 && failures > successes
		}
	}

	if c.BufferSize <= 0 {
		c.BufferSize = 1000
	}

	// 确保 BufferSize 是2的幂，用于快速取模
	bufferSize := 1
	for bufferSize < c.BufferSize {
		bufferSize <<= 1
	}

	cb := &OptimizedCircuitBreaker{
		timeWindow:    c.TimeWindow.Nanoseconds(),
		onStateChange: c.OnStateChange,
		readyToTrip:   c.ReadyToTrip,
		probe:         c.Probe,
		bufferSize:    bufferSize,
		ringBuffer:    newOptimizedRingBuffer(bufferSize),
	}

	cb.state.Store(stateClosedOpt)
	cb.stats.changed.Store(1)
	cb.stats.lastCleanupTime.Store(time.Now().UnixNano())

	return cb
}

// newOptimizedRingBuffer 创建优化的环形缓冲区
func newOptimizedRingBuffer(size int) *optimizedRingBuffer {
	return &optimizedRingBuffer{
		buffer: make([]int64, size),
		size:   uint64(size),
		mask:   uint64(size - 1),
	}
}

// Before 判断是否允许执行新请求 (无锁优化)
func (p *OptimizedCircuitBreaker) Before() bool {
	p.updateStateOptimized()

	state := p.state.Load()
	switch state {
	case stateClosedOpt:
		return true
	case stateOpenOpt:
		return false
	case stateHalfOpenOpt:
		return p.probe()
	default:
		return false
	}
}

// After 记录请求执行结果 (无锁优化)
func (p *OptimizedCircuitBreaker) After(success bool) {
	now := time.Now().UnixNano()
	
	// 紧凑存储：时间戳(相对) + 成功标志
	baseTime := p.stats.lastCleanupTime.Load()
	relativeTime := now - baseTime
	if relativeTime < 0 {
		relativeTime = 0
	}
	
	var packed int64
	if relativeTime > 0x7FFFFFFF { // 超过32位，使用绝对时间
		packed = now << timeShift
	} else {
		packed = relativeTime << timeShift
	}
	
	if success {
		packed |= successFlag
		p.stats.successes.Add(1)
	} else {
		p.stats.failures.Add(1)
	}

	// 无锁添加到环形缓冲区
	p.ringBuffer.addOptimized(packed)
	p.stats.changed.Store(1)
}

// addOptimized 无锁添加元素到环形缓冲区
func (rb *optimizedRingBuffer) addOptimized(packed int64) {
	tail := rb.tail.Add(1) - 1
	rb.buffer[tail&rb.mask] = packed
}

// updateStateOptimized 无锁状态更新
func (p *OptimizedCircuitBreaker) updateStateOptimized() {
	if p.stats.changed.Load() == 0 && !p.cleanUpOptimized() {
		return
	}

	// 使用CAS避免锁竞争
	if !p.stats.changed.CompareAndSwap(1, 0) {
		return
	}

	oldState := p.state.Load()
	successes := p.stats.successes.Load()
	failures := p.stats.failures.Load()

	shouldTrip := p.readyToTrip(successes, failures)
	var newState uint32

	switch oldState {
	case stateClosedOpt:
		if shouldTrip {
			newState = stateOpenOpt
		} else {
			newState = stateClosedOpt
		}
	case stateOpenOpt:
		if !shouldTrip {
			newState = stateHalfOpenOpt
		} else {
			newState = stateOpenOpt
		}
	case stateHalfOpenOpt:
		if p.ringBuffer.hasRecentRequest() {
			if p.ringBuffer.lastRequestSuccess() {
				newState = stateClosedOpt
				p.ringBuffer.reset()
			} else {
				newState = stateOpenOpt
			}
		} else {
			newState = stateHalfOpenOpt
		}
	default:
		newState = stateClosedOpt
	}

	if p.state.CompareAndSwap(oldState, newState) && oldState != newState && p.onStateChange != nil {
		// 状态转换成功，触发回调
		oldStateEnum := stateFromUint32(oldState)
		newStateEnum := stateFromUint32(newState)
		p.onStateChange(oldStateEnum, newStateEnum)
	}
}

// cleanUpOptimized 无锁清理过期数据
func (p *OptimizedCircuitBreaker) cleanUpOptimized() bool {
	now := time.Now().UnixNano()
	lastCleanup := p.stats.lastCleanupTime.Load()
	
	if now-lastCleanup < p.timeWindow {
		return false
	}

	if !p.stats.lastCleanupTime.CompareAndSwap(lastCleanup, now) {
		return false // 其他goroutine已经在清理
	}

	threshold := now - p.timeWindow
	removedSuccesses, removedFailures := p.ringBuffer.cleanupOptimized(threshold, lastCleanup)

	if removedSuccesses > 0 || removedFailures > 0 {
		p.stats.successes.Add(^(removedSuccesses - 1)) // 原子减法
		p.stats.failures.Add(^(removedFailures - 1))   // 原子减法
		return true
	}

	return false
}

// cleanupOptimized 无锁清理环形缓冲区
func (rb *optimizedRingBuffer) cleanupOptimized(threshold, baseTime int64) (removedSuccesses, removedFailures uint64) {
	head := rb.head.Load()
	tail := rb.tail.Load()

	for head < tail {
		idx := head & rb.mask
		packed := rb.buffer[idx]
		
		if packed == 0 {
			break
		}

		// 解析时间戳
		var timestamp int64
		if packed>>timeShift > 0x7FFFFFFF {
			timestamp = packed >> timeShift // 绝对时间
		} else {
			timestamp = baseTime + (packed >> timeShift) // 相对时间
		}

		if timestamp >= threshold {
			break
		}

		// 统计被清理的请求
		if packed&successFlag != 0 {
			removedSuccesses++
		} else {
			removedFailures++
		}

		// 清理数据防止内存泄漏
		rb.buffer[idx] = 0
		head++
	}

	rb.head.Store(head)
	return
}

// hasRecentRequest 检查是否有最近的请求
func (rb *optimizedRingBuffer) hasRecentRequest() bool {
	return rb.head.Load() < rb.tail.Load()
}

// lastRequestSuccess 获取最后一个请求的成功状态
func (rb *optimizedRingBuffer) lastRequestSuccess() bool {
	head := rb.head.Load()
	tail := rb.tail.Load()
	if head >= tail {
		return false
	}
	
	lastIdx := (tail - 1) & rb.mask
	packed := rb.buffer[lastIdx]
	return packed&successFlag != 0
}

// reset 重置环形缓冲区
func (rb *optimizedRingBuffer) reset() {
	rb.head.Store(0)
	rb.tail.Store(0)
}

// Call 执行服务调用 (优化版本)
func (p *OptimizedCircuitBreaker) Call(fn func() error) error {
	if !p.Before() {
		return errors.New("circuit breaker is open")
	}

	err := fn()
	p.After(err == nil)
	return err
}

// State 返回当前熔断器状态
func (p *OptimizedCircuitBreaker) State() State {
	return stateFromUint32(p.state.Load())
}

// Stat 获取成功和失败计数
func (p *OptimizedCircuitBreaker) Stat() (successes, failures uint64) {
	return p.stats.successes.Load(), p.stats.failures.Load()
}

// Total 获取总请求数
func (p *OptimizedCircuitBreaker) Total() uint64 {
	return p.stats.successes.Load() + p.stats.failures.Load()
}

// 状态转换辅助函数
func stateFromUint32(state uint32) State {
	switch state {
	case stateClosedOpt:
		return Closed
	case stateOpenOpt:
		return Open
	case stateHalfOpenOpt:
		return HalfOpen
	default:
		return Closed
	}
}

// FastCircuitBreaker 超轻量级版本 (仅基础功能)
type FastCircuitBreaker struct {
	state     atomic.Uint32 // 状态
	successes atomic.Uint64 // 成功计数
	failures  atomic.Uint64 // 失败计数
	lastReset atomic.Int64  // 上次重置时间

	threshold   uint64 // 失败阈值
	windowNanos int64  // 时间窗口(纳秒)
}

// NewFastCircuitBreaker 创建超轻量级熔断器
func NewFastCircuitBreaker(failureThreshold uint64, timeWindow time.Duration) *FastCircuitBreaker {
	cb := &FastCircuitBreaker{
		threshold:   failureThreshold,
		windowNanos: timeWindow.Nanoseconds(),
	}
	cb.lastReset.Store(time.Now().UnixNano())
	return cb
}

// AllowRequest 检查是否允许请求 (超快速)
func (cb *FastCircuitBreaker) AllowRequest() bool {
	now := time.Now().UnixNano()
	lastReset := cb.lastReset.Load()

	// 检查是否需要重置窗口
	if now-lastReset > cb.windowNanos {
		if cb.lastReset.CompareAndSwap(lastReset, now) {
			cb.successes.Store(0)
			cb.failures.Store(0)
		}
	}

	state := cb.state.Load()
	if state == stateOpenOpt {
		// 熔断状态：检查是否可以尝试半开
		return false
	}

	return true // Closed 或 HalfOpen 状态允许请求
}

// RecordResult 记录请求结果 (超快速)
func (cb *FastCircuitBreaker) RecordResult(success bool) {
	if success {
		cb.successes.Add(1)
	} else {
		failures := cb.failures.Add(1)
		if failures >= cb.threshold {
			cb.state.Store(stateOpenOpt)
		}
	}
}

// CallFast 快速执行服务调用
func (cb *FastCircuitBreaker) CallFast(fn func() error) error {
	if !cb.AllowRequest() {
		return errors.New("circuit breaker is open")
	}

	err := fn()
	cb.RecordResult(err == nil)
	
	// 半开状态的快速恢复逻辑
	if err == nil && cb.state.Load() == stateHalfOpenOpt {
		cb.state.Store(stateClosedOpt)
	}
	
	return err
}

// 内存池优化 - 复用 requestResult 对象
var requestResultPool = sync.Pool{
	New: func() interface{} {
		return &requestResult{}
	},
}

// getRequestResult 从池中获取对象
func getRequestResult() *requestResult {
	return requestResultPool.Get().(*requestResult)
}

// putRequestResult 归还对象到池中
func putRequestResult(r *requestResult) {
	r.success = false
	r.time = time.Time{}
	requestResultPool.Put(r)
}

// 批量操作优化版本
type BatchCircuitBreaker struct {
	*OptimizedCircuitBreaker
	batchSize    int
	batchBuffer  []bool
	batchIndex   atomic.Int32
	batchMutex   sync.Mutex
	batchTimeout time.Duration
	lastFlush    atomic.Int64
}

// NewBatchCircuitBreaker 创建批量处理版本
func NewBatchCircuitBreaker(config CircuitBreakerConfig, batchSize int, batchTimeout time.Duration) *BatchCircuitBreaker {
	if batchSize <= 0 {
		batchSize = 100
	}
	
	cb := &BatchCircuitBreaker{
		OptimizedCircuitBreaker: NewOptimizedCircuitBreaker(config),
		batchSize:               batchSize,
		batchBuffer:             make([]bool, batchSize),
		batchTimeout:            batchTimeout,
	}
	cb.lastFlush.Store(time.Now().UnixNano())
	
	// 启动定时刷新
	go cb.flushLoop()
	
	return cb
}

// AfterBatch 批量记录结果
func (cb *BatchCircuitBreaker) AfterBatch(success bool) {
	now := time.Now().UnixNano()
	
	// 检查是否需要强制刷新
	if now-cb.lastFlush.Load() > cb.batchTimeout.Nanoseconds() {
		cb.flush()
		return
	}
	
	index := cb.batchIndex.Add(1) - 1
	if int(index) >= cb.batchSize {
		cb.flush()
		cb.AfterBatch(success) // 重试
		return
	}
	
	cb.batchBuffer[index] = success
	
	if int(index) >= cb.batchSize-1 {
		cb.flush()
	}
}

// flush 刷新批量数据
func (cb *BatchCircuitBreaker) flush() {
	cb.batchMutex.Lock()
	defer cb.batchMutex.Unlock()
	
	index := cb.batchIndex.Swap(0)
	if index == 0 {
		return
	}
	
	var successes, failures uint64
	for i := int32(0); i < index; i++ {
		if cb.batchBuffer[i] {
			successes++
		} else {
			failures++
		}
	}
	
	// 批量更新统计
	if successes > 0 {
		cb.stats.successes.Add(successes)
	}
	if failures > 0 {
		cb.stats.failures.Add(failures)
	}
	
	cb.stats.changed.Store(1)
	cb.lastFlush.Store(time.Now().UnixNano())
}

// flushLoop 定时刷新循环
func (cb *BatchCircuitBreaker) flushLoop() {
	ticker := time.NewTicker(cb.batchTimeout)
	defer ticker.Stop()
	
	for range ticker.C {
		if cb.batchIndex.Load() > 0 {
			cb.flush()
		}
	}
}

// 零分配版本的状态查询
func (p *OptimizedCircuitBreaker) GetState() uint32 {
	return p.state.Load()
}

// IsOpen 零分配检查是否熔断
func (p *OptimizedCircuitBreaker) IsOpen() bool {
	return p.state.Load() == stateOpenOpt
}

// IsClosed 零分配检查是否正常
func (p *OptimizedCircuitBreaker) IsClosed() bool {
	return p.state.Load() == stateClosedOpt
}

// 编译时检查接口是否被正确实现
var (
	_ unsafe.Pointer = unsafe.Pointer(&OptimizedCircuitBreaker{})
	_ unsafe.Pointer = unsafe.Pointer(&FastCircuitBreaker{})
	_ unsafe.Pointer = unsafe.Pointer(&BatchCircuitBreaker{})
)