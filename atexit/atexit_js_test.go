//go:build js

package atexit

import (
	"testing"
)

func TestJSRegister(t *testing.T) {
	callbacks = nil

	executed := false
	Register(func() {
		executed = true
	})

	if len(callbacks) != 1 {
		t.Errorf("Expected 1 callback, got %d", len(callbacks))
	}
}

func TestJSExecuteCallbacks(t *testing.T) {
	callbacks = nil

	counter := 0
	Register(func() {
		counter++
	})
	Register(func() {
		counter++
	})

	executeCallbacks()

	if counter != 2 {
		t.Errorf("Expected counter to be 2, got %d", counter)
	}
}

func TestJSNilCallback(t *testing.T) {
	callbacks = nil
	Register(nil)

	if len(callbacks) != 0 {
		t.Errorf("Expected 0 callbacks when registering nil, got %d", len(callbacks))
	}
}

func TestJSPanicRecovery(t *testing.T) {
	callbacks = nil

	counter := 0
	Register(func() {
		panic("test panic")
	})
	Register(func() {
		counter++
	})

	executeCallbacks()

	if counter != 1 {
		t.Errorf("Expected second callback to execute despite first panicking, counter=%d", counter)
	}
}

func TestJSExitSuccess(t *testing.T) {
	callbacks = nil

	executed := false
	Register(func() {
		executed = true
	})

	// Exit with code 0 should not panic
	defer func() {
		if r := recover(); r != nil {
			t.Errorf("Exit(0) should not panic, got: %v", r)
		}
		if !executed {
			t.Errorf("Callback should have been executed")
		}
	}()

	Exit(0)
}

func TestJSExitNonZero(t *testing.T) {
	callbacks = nil

	executed := false
	Register(func() {
		executed = true
	})

	// Exit with non-zero code should panic
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("Exit(1) should panic")
		}
		if !executed {
			t.Errorf("Callback should have been executed before panic")
		}
	}()

	Exit(1)
}
