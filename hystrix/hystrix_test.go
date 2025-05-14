package hystrix

import (
	"errors"
	"testing"
	"time"
)

func TestNewCircuitBreaker(t *testing.T) {
	cb := NewCircuitBreaker(CircuitBreakerConfig{
		TimeWindow: time.Second * 2,
	})
	if cb.State() != Open {
		t.Errorf("Expected initial state to be Open but got %s", cb.State())
	}
}

func TestBefore(t *testing.T) {
	cb := NewCircuitBreaker(CircuitBreakerConfig{
		TimeWindow: time.Second * 2,
	})
	if !cb.Before() {
		t.Error("Before() should return true when state is Open")
	}

	cb.state = Closed
	if cb.Before() {
		t.Error("Before() should return false when state is Closed")
	}

	cb.state = HalfOpen
	cb.CircuitBreakerConfig.Probe = func() bool { return true }
	if !cb.Before() {
		t.Error("Before() should return true when Probe returns true")
	}

	cb.CircuitBreakerConfig.Probe = func() bool { return false }
	if cb.Before() {
		t.Error("Before() should return false when Probe returns false")
	}
}

func TestState(t *testing.T) {
	cb := NewCircuitBreaker(CircuitBreakerConfig{
		TimeWindow: time.Second * 2,
	})
	if cb.State() != Open {
		t.Errorf("Expected state to be Open but got %s", cb.State())
	}

	cb.state = Closed
	if cb.State() != Closed {
		t.Errorf("Expected state to be Closed but got %s", cb.State())
	}
}

func TestStat(t *testing.T) {
	cb := NewCircuitBreaker(CircuitBreakerConfig{
		TimeWindow: time.Second * 2,
	})
	if successes, failures := cb.Stat(); successes != 0 || failures != 0 {
		t.Errorf("Expected initial stats to be 0, got %d/%d", successes, failures)
	}

	cb.successes.Store(10)
	cb.failures.Store(5)
	if successes, failures := cb.Stat(); successes != 10 || failures != 5 {
		t.Errorf("Expected stats 10/5, got %d/%d", successes, failures)
	}
}

func TestTotal(t *testing.T) {
	cb := NewCircuitBreaker(CircuitBreakerConfig{
		TimeWindow: time.Second * 2,
	})
	if total := cb.Total(); total != 0 {
		t.Errorf("Expected total to be 0, got %d", total)
	}

	cb.successes.Store(10)
	cb.failures.Store(5)
	if total := cb.Total(); total != 15 {
		t.Errorf("Expected total to be 15, got %d", total)
	}
}

func TestAfter(t *testing.T) {
	cb := NewCircuitBreaker(CircuitBreakerConfig{
		TimeWindow: time.Second * 2,
	})

	cb.After(true)
	if cb.successes.Load() != 1 {
		t.Error("Success count should be 1 after After(true)")
	}

	cb.After(false)
	if cb.failures.Load() != 1 {
		t.Error("Failure count should be 1 after After(false)")
	}
}

func TestCall(t *testing.T) {
	cb := NewCircuitBreaker(CircuitBreakerConfig{
		TimeWindow: time.Second * 2,
	})

	err := cb.Call(func() error { return nil })
	if err != nil {
		t.Error("Call() should return nil for successful calls")
	}

	cb.state = Closed
	err = cb.Call(func() error { return nil })
	if err == nil {
		t.Error("Call() should return error when state is Closed")
	}
}

func TestFullLifecycle(t *testing.T) {
	cb := NewCircuitBreaker(CircuitBreakerConfig{
		TimeWindow: time.Second * 2,
		ReadyToTrip: func(successes, failures uint64) bool {
			return failures > successes
		},
		Probe: func() bool { return true },
	})

	// Test successful calls
	for i := 0; i < 10; i++ {
		err := cb.Call(func() error { return nil })
		if err != nil {
			t.Errorf("Expected no error on successful call, got %v", err)
		}
	}

	time.Sleep(time.Millisecond * 500) // Ensure results are within TimeWindow

	if cb.State() != Open {
		t.Error("State should remain Open with successful calls")
	}

	// Test failed calls
	for i := 0; i < 15; i++ {
		err := cb.Call(func() error { return errors.New("test error") })
		if err == nil {
			t.Error("Expected error on failed call")
		}
	}

	time.Sleep(time.Millisecond * 500) // Ensure results are within TimeWindow

	if cb.State() != Closed {
		t.Error("State should transition to Closed after failures")
	}

	// Test recovery
	for i := 0; i < 5; i++ {
		err := cb.Call(func() error { return nil })
		if err == nil {
			t.Errorf("Expected no error during recovery phase, got %v", err)
		}
	}

	time.Sleep(time.Millisecond * 500) // Ensure results are within TimeWindow

	if cb.State() != Closed {
		t.Error("State should transition back to Closed after successful recovery")
	}
}

/*
go test -bench=. -benchmem -count=3
goos: darwin
goarch: arm64
pkg: github.com/lazygophers/utils/hystrix
cpu: Apple M3
BenchmarkCall_Success          	19843444	        59.79 ns/op	      32 B/op	       1 allocs/op
BenchmarkCall_Success          	19987354	        59.91 ns/op	      32 B/op	       1 allocs/op
BenchmarkCall_Success          	20103996	        60.67 ns/op	      32 B/op	       1 allocs/op
BenchmarkCall_Failure          	27953673	        42.26 ns/op	      16 B/op	       1 allocs/op
BenchmarkCall_Failure          	26144541	        42.42 ns/op	      16 B/op	       1 allocs/op
BenchmarkCall_Failure          	28086462	        42.10 ns/op	      16 B/op	       1 allocs/op
BenchmarkCall_Success_Parallel 	20359053	        60.05 ns/op	      32 B/op	       1 allocs/op
BenchmarkCall_Success_Parallel 	20110200	        59.65 ns/op	      32 B/op	       1 allocs/op
BenchmarkCall_Success_Parallel 	19292913	        59.65 ns/op	      32 B/op	       1 allocs/op
BenchmarkCall_Failure_Parallel 	28006354	        42.13 ns/op	      16 B/op	       1 allocs/op
BenchmarkCall_Failure_Parallel 	28381402	        42.20 ns/op	      16 B/op	       1 allocs/op
BenchmarkCall_Failure_Parallel 	27560493	        42.17 ns/op	      16 B/op	       1 allocs/op
BenchmarkCall_StateTransition  	13347013	        91.49 ns/op	      32 B/op	       1 allocs/op
BenchmarkCall_StateTransition  	13113416	        91.26 ns/op	      32 B/op	       1 allocs/op
BenchmarkCall_StateTransition  	13127863	        92.29 ns/op	      32 B/op	       1 allocs/op
PASS
ok  	github.com/lazygop
*/

func BenchmarkCall_Success(b *testing.B) {
	cb := NewCircuitBreaker(CircuitBreakerConfig{
		TimeWindow: time.Second * 2,
		ReadyToTrip: func(successes, failures uint64) bool {
			return false // 确保不会触发熔断
		},
	})

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = cb.Call(func() error { return nil })
	}
}

func BenchmarkCall_Failure(b *testing.B) {
	cb := NewCircuitBreaker(CircuitBreakerConfig{
		TimeWindow: time.Second * 2,
		ReadyToTrip: func(successes, failures uint64) bool {
			return failures > successes
		},
	})

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = cb.Call(func() error { return errors.New("test error") })
	}
}

func BenchmarkCall_Success_Parallel(b *testing.B) {
	cb := NewCircuitBreaker(CircuitBreakerConfig{
		TimeWindow: time.Second * 2,
		ReadyToTrip: func(successes, failures uint64) bool {
			return false
		},
	})

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			_ = cb.Call(func() error { return nil })
		}
	})
}

func BenchmarkCall_Failure_Parallel(b *testing.B) {
	cb := NewCircuitBreaker(CircuitBreakerConfig{
		TimeWindow: time.Second * 2,
		ReadyToTrip: func(successes, failures uint64) bool {
			return failures > successes
		},
	})

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			_ = cb.Call(func() error { return errors.New("test error") })
		}
	})
}

// 新增极端并发场景测试
func BenchmarkCall_StateTransition(b *testing.B) {
	cb := NewCircuitBreaker(CircuitBreakerConfig{
		TimeWindow: time.Millisecond * 100,
		ReadyToTrip: func(s, f uint64) bool {
			return f > s
		},
	})

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			// 交替成功/失败请求触发状态转换
			if time.Now().UnixNano()%2 == 0 {
				_ = cb.Call(func() error { return nil })
			} else {
				_ = cb.Call(func() error { return errors.New("error") })
			}
		}
	})
}
