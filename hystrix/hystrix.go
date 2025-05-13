package hystrix

import (
	"errors"
	"github.com/lazygophers/utils/candy"
	"sync"
	"time"
)

type requestResult struct {
	success bool

	time time.Time
}

type State string

const (
	Closed   State = "closed"
	Open     State = "open"
	HalfOpen State = "half-open"
)

type StateChange func(oldState, newState State)
type ReadyToTrip func(successes, failures uint64) bool
type Probe func() bool

type CircuitBreaker struct {
	mu sync.Mutex
	CircuitBreakerConfig

	state State

	requestResults []*requestResult

	successes, failures uint64

	changed bool
}

type CircuitBreakerConfig struct {
	// 测试的时间窗口
	TimeWindow time.Duration
	// 状态变化的回调
	OnStateChange StateChange
	// 如果返回了 true，则会发生负面的状态变化
	ReadyToTrip ReadyToTrip
	// 当半开状态下，是否重试的判断
	Probe Probe
}

func NewCircuitBreaker(c CircuitBreakerConfig) *CircuitBreaker {
	if c.Probe == nil {
		c.Probe = ProbeWithChance(50)
	}

	// 新增 ReadyToTrip 默认实现
	if c.ReadyToTrip == nil {
		c.ReadyToTrip = func(successes, failures uint64) bool {
			return successes+failures == 0 && failures > successes
		}
	}

	return &CircuitBreaker{
		CircuitBreakerConfig: c,

		changed:        true,
		state:          Open,
		requestResults: make([]*requestResult, 0),
		successes:      0,
		failures:       0,
	}
}

func (p *CircuitBreaker) cleanUp() (change bool) {
	now := time.Now()
	for len(p.requestResults) > 0 && now.Sub(p.requestResults[0].time) > p.TimeWindow {
		if len(p.requestResults) > 1 {
			p.requestResults = p.requestResults[1:] // Remove expired results
		} else {
			p.requestResults = make([]*requestResult, 0, 1)
		}
		change = true
	}
	return
}

func (p *CircuitBreaker) Before() bool {
	p.mu.Lock()
	defer p.mu.Unlock()

	p.updateState()

	switch p.state {
	case Closed:
		// close，不可用
		return false

	case HalfOpen:
		// 半开状态，按照概率决定是否尝试
		if p.Probe() {
			return true
		}

		return false
	}

	return true
}

func (p *CircuitBreaker) State() State {
	p.mu.Lock()
	defer p.mu.Unlock()

	return p.state
}

func (p *CircuitBreaker) Stat() (successes, failures uint64) {
	p.mu.Lock()
	defer p.mu.Unlock()

	return p.successes, p.failures
}

func (p *CircuitBreaker) Total() uint64 {
	p.mu.Lock()
	defer p.mu.Unlock()

	return p.successes + p.failures
}

func (p *CircuitBreaker) stat() (successes, failures uint64) {
	p.successes = 0
	p.failures = 0

	for _, r := range p.requestResults {
		if r.success {
			p.successes++
		} else {
			p.failures++
		}
	}

	return p.successes, p.failures
}

func (p *CircuitBreaker) updateState() {
	if !p.changed && !p.cleanUp() {
		return
	}

	// 状态变化逻辑
	oldState := p.state
	successes, failures := p.stat()
	if p.ReadyToTrip(successes, failures) {
		switch oldState {
		case Open:
			p.state = HalfOpen
		case HalfOpen:
			if len(p.requestResults) > 0 && candy.Last(p.requestResults).success {
				p.state = Open
				p.requestResults = make([]*requestResult, 0) // 清空历史记录
			} else {
				p.state = Closed
			}
		}
	} else {
		switch oldState {
		case HalfOpen:
			// 增加对 requestResults 长度的判断
			if len(p.requestResults) > 0 && candy.Last(p.requestResults).success {
				p.state = Open
			} else {
				p.state = Closed
			}
		case Closed:
			// 确保从 Closed 转换到 HalfOpen 的逻辑正确
			if failures > 0 {
				p.state = HalfOpen
			}
		}
	}

	p.changed = false

	// 触发状态变化回调
	if oldState != p.state && p.OnStateChange != nil {
		p.OnStateChange(oldState, p.state)
	}
}

func (p *CircuitBreaker) After(success bool) {
	p.mu.Lock()
	defer p.mu.Unlock()

	p.requestResults = append(p.requestResults, &requestResult{success: success, time: time.Now()})
	p.changed = true
	if success {
		p.successes++
	} else {
		p.failures++
	}
}

func (p *CircuitBreaker) Call(fn func() error) error {
	if !p.Before() {
		return errors.New("circuit breaker is open")
	}

	err := fn()
	p.After(err == nil)

	return err
}
