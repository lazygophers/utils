package hystrix_test

import (
	"errors"
	"github.com/lazygophers/utils/hystrix"
	"github.com/lazygophers/utils/randx"
	"testing"
	"time"
)

func TestHystrix(t *testing.T) {
	breaker := hystrix.NewCircuitBreaker(hystrix.CircuitBreakerConfig{
		TimeWindow: time.Second * 2,
		OnStateChange: func(oldState, newState hystrix.State) {
			t.Logf("state %s -> %s", oldState, newState)
		},
		ReadyToTrip: func(successes, failures uint64) bool {
			if successes == 0 {
				return failures > 0
			}

			if failures == 0 {
				return false
			}

			return failures > successes // 失败比例大于 50%
		},
	})

	for i := 0; i < 20; i++ {
		breaker.Call(func() error {
			if randx.Booln(50) {
				return errors.New("error")
			}

			return nil
		})
		//t.Error(err)
	}

	time.Sleep(time.Second * 5)

	for i := 0; i < 20; i++ {
		breaker.Call(func() error {
			if randx.Booln(50) {
				return errors.New("error")
			}

			return nil
		})
	}
}

func TestCount(t *testing.T) {
	breaker := hystrix.NewCircuitBreaker(hystrix.CircuitBreakerConfig{
		TimeWindow: time.Minute,
		OnStateChange: func(oldState, newState hystrix.State) {
			t.Logf("state %s -> %s", oldState, newState)
		},
		ReadyToTrip: func(successes, failures uint64) bool {
			return failures > successes
		},
	})

	for i := 0; i < 100; i++ {
		breaker.After(true)
	}

	time.Sleep(time.Second * 6)
	t.Log(breaker.Stat())
}
