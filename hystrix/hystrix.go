package hystrix

import (
	"errors"
	"github.com/lazygophers/utils/candy"
	"go.uber.org/atomic"
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
	expiredAt      *atomic.Time

	successes, failures *atomic.Uint64
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
	// 状态过期时间，如果为 0 则始终进行检测
	StateExpiredTime time.Duration
}

func NewCircuitBreaker(c CircuitBreakerConfig) *CircuitBreaker {
	if c.Probe == nil {
		c.Probe = ProbeWithChance(50)
	}

	return &CircuitBreaker{
		CircuitBreakerConfig: c,

		state:          Open,
		requestResults: make([]*requestResult, 0),
		expiredAt:      atomic.NewTime(time.Now()),
		successes:      atomic.NewUint64(0),
		failures:       atomic.NewUint64(0),
	}
}

func (p *CircuitBreaker) cleanUp() (change bool) {
	now := time.Now()
	for len(p.requestResults) > 0 && now.Sub(p.requestResults[0].time).Truncate(time.Second) > p.TimeWindow {
		p.requestResults = p.requestResults[1:] // Remove expired results
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

	return p.successes.Load(), p.failures.Load()
}

func (p *CircuitBreaker) stat() (successes, failures uint64) {
	p.successes.Store(0)
	p.failures.Store(0)

	for _, r := range p.requestResults {
		if r.success {
			p.successes.Inc()
		} else {
			p.failures.Inc()
		}
	}

	return p.successes.Load(), p.failures.Load()
}

func (p *CircuitBreaker) updateState() {
	if time.Now().Before(p.expiredAt.Load()) || p.cleanUp() {
		return
	}

	p.expiredAt.Store(time.Now().Add(p.StateExpiredTime))

	// 状态变化逻辑
	oldState := p.state
	if p.ReadyToTrip(p.stat()) {
		switch oldState {
		case Open:
			p.state = HalfOpen
		case HalfOpen:
			// 对于当前是关闭的状态，如果最后一个是正常的，那么需要回退
			if candy.LastOr(p.requestResults, &requestResult{success: false}).success {
				p.state = Open
				p.requestResults = make([]*requestResult, 1)
			} else {
				p.state = Closed
			}
		}
	} else {
		switch oldState {
		case HalfOpen:
			p.state = Open
		case Closed:
			p.state = HalfOpen
		}
	}

	// 触发状态变化回调
	if oldState != p.state && p.OnStateChange != nil {
		p.OnStateChange(oldState, p.state)
	}
}

func (p *CircuitBreaker) After(success bool) {
	p.mu.Lock()
	defer p.mu.Unlock()

	p.requestResults = append(p.requestResults, &requestResult{success: success, time: time.Now()})
}

func (p *CircuitBreaker) Call(fn func() error) error {
	if !p.Before() {
		return errors.New("circuit breaker is open")
	}

	err := fn()
	p.After(err == nil)

	return err
}
