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

	cb.successes = 10
	cb.failures = 5
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

	cb.successes = 10
	cb.failures = 5
	if total := cb.Total(); total != 15 {
		t.Errorf("Expected total to be 15, got %d", total)
	}
}

func TestAfter(t *testing.T) {
	cb := NewCircuitBreaker(CircuitBreakerConfig{
		TimeWindow: time.Second * 2,
	})

	cb.After(true)
	if cb.successes != 1 {
		t.Error("Success count should be 1 after After(true)")
	}

	cb.After(false)
	if cb.failures != 1 {
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
