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

type CircuitBreaker struct {
	mu sync.Mutex

	state State

	timeWindow     time.Duration
	requestResults []RequestResult

	onStateChange StateChange
	readyToTrip   ReadyToTrip
}

type CircuitBreakerConfig struct {
	TimeWindow    time.Duration
	OnStateChange StateChange
	ReadyToTrip   ReadyToTrip
}

func NewCircuitBreaker(c CircuitBreakerConfig) *CircuitBreaker {
	return &CircuitBreaker{
		state: Closed,

		timeWindow:     c.TimeWindow,
		requestResults: make([]RequestResult, 0),
		onStateChange:  c.OnStateChange,
		readyToTrip:    c.ReadyToTrip,
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

	if p.state == Open {
		return false
	}

	return true
}

func (p *CircuitBreaker) After(success bool) {
	p.mu.Lock()
	defer p.mu.Unlock()

	result := RequestResult{success: success, time: time.Now()}
	p.requestResults = append(p.requestResults, result)

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
			p.state = Open
		} else {
			p.state = HalfOpen
		}
	} else {
		p.state = Closed
	}

	// 触发状态变化回调
	if oldState != p.state && p.onStateChange != nil {
		p.onStateChange(oldState, p.state)
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
