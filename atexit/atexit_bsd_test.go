//go:build freebsd || openbsd || netbsd || dragonfly

package atexit

import (
	"os"
	"syscall"
	"testing"
	"time"
)

func TestBSDRegister(t *testing.T) {
	executed := false
	Register(func() {
		executed = true
	})

	if len(callbacks) != 1 {
		t.Errorf("Expected 1 callback, got %d", len(callbacks))
	}
}

func TestBSDExecuteCallbacks(t *testing.T) {
	// Reset callbacks
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

func TestBSDSignalHandler(t *testing.T) {
	// Reset for test
	callbacks = nil
	signalOnce = *new(func())

	executed := false
	Register(func() {
		executed = true
	})

	// Give signal handler time to initialize
	time.Sleep(100 * time.Millisecond)

	// Note: We can't easily test actual signal handling in unit tests
	// This test just verifies registration works
	if len(callbacks) != 1 {
		t.Errorf("Expected 1 callback registered")
	}
}

func TestBSDNilCallback(t *testing.T) {
	callbacks = nil
	Register(nil)

	if len(callbacks) != 0 {
		t.Errorf("Expected 0 callbacks when registering nil, got %d", len(callbacks))
	}
}

func TestBSDPanicRecovery(t *testing.T) {
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
