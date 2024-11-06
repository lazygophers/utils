package hystrix

import (
	"errors"
	"sync"
	"time"
)

type RequestResult struct {
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

	timeWindow     time.Duration
	requestResults []RequestResult

	onStateChange StateChange
	readyToTrip   ReadyToTrip
}

type CircuitBreakerConfig struct {
	// 测试的时间窗口
	TimeWindow time.Duration
	// 状态变化的回调
	OnStateChange StateChange
	// 如果返回了 true，则会发生状态变化
	ReadyToTrip ReadyToTrip
	// 当半开状态下，是否重试的判断
	Probe Probe
}

func NewCircuitBreaker(c CircuitBreakerConfig) *CircuitBreaker {
	if c.Probe == nil {
		c.Probe = ProbeWithChance(50)
	}

	return &CircuitBreaker{
		CircuitBreakerConfig: c,

		state:          Open,
		requestResults: make([]RequestResult, 0),
	}
}

func (p *CircuitBreaker) cleanUp() {
	now := time.Now()
	for len(p.requestResults) > 0 && now.Sub(p.requestResults[0].time) > p.timeWindow {
		p.requestResults = p.requestResults[1:] // Remove expired results
	}
}

func (p *CircuitBreaker) Before() bool {
	p.mu.Lock()
	defer p.mu.Unlock()

	p.cleanUp()
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

func (p *CircuitBreaker) updateState() {
	// 计算状态变化
	var failures, successes uint64
	for _, r := range p.requestResults {
		if r.success {
			successes++
		} else {
			failures++
		}
	}

	// 状态变化逻辑
	oldState := p.state
	if p.readyToTrip(successes, failures) {
		if oldState == HalfOpen {
			p.state = Closed
		} else {
			p.state = HalfOpen
		}
	} else {
		p.state = Open
	}

	// 触发状态变化回调
	if oldState != p.state && p.onStateChange != nil {
		p.onStateChange(oldState, p.state)
	}
}

func (p *CircuitBreaker) After(success bool) {
	p.mu.Lock()
	defer p.mu.Unlock()

	result := RequestResult{success: success, time: time.Now()}
	p.requestResults = append(p.requestResults, result)
}

func (p *CircuitBreaker) Call(fn func() error) error {
	if !p.Before() {
		return errors.New("circuit breaker is open")
	}

	err := fn()
	p.After(err == nil)

	return err
}
