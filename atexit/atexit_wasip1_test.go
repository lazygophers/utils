//go:build wasip1

package atexit

import (
	"testing"
)

func TestWASIRegister(t *testing.T) {
	callbacks = nil

	executed := false
	Register(func() {
		executed = true
	})

	if len(callbacks) != 1 {
		t.Errorf("Expected 1 callback, got %d", len(callbacks))
	}
}

func TestWASIExecuteCallbacks(t *testing.T) {
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

func TestWASINilCallback(t *testing.T) {
	callbacks = nil
	Register(nil)

	if len(callbacks) != 0 {
		t.Errorf("Expected 0 callbacks when registering nil, got %d", len(callbacks))
	}
}

func TestWASIPanicRecovery(t *testing.T) {
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
