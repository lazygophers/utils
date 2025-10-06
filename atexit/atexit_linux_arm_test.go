//go:build linux && arm

package atexit

import (
	"sync/atomic"
	"testing"
	"time"
)

func TestRegister_ARM(t *testing.T) {
	var called int32

	Register(func() {
		atomic.AddInt32(&called, 1)
	})

	// Signal handling is initialized
	// We cannot easily test actual signal delivery in unit tests
	// 信号处理已初始化
	// 单元测试中无法轻松测试实际的信号传递
	if atomic.LoadInt32(&called) != 0 {
		t.Error("Callback should not be called during registration")
	}
}

func TestRegisterNil_ARM(t *testing.T) {
	// Should not panic
	// 不应该 panic
	Register(nil)
}

func TestExecuteCallbacks_ARM(t *testing.T) {
	callbacks = nil
	var called int32

	callbacksMu.Lock()
	callbacks = append(callbacks, func() {
		atomic.AddInt32(&called, 1)
	})
	callbacksMu.Unlock()

	executeCallbacks()

	time.Sleep(10 * time.Millisecond)
	if atomic.LoadInt32(&called) != 1 {
		t.Errorf("Expected callback to be called once, got %d", called)
	}
}

func TestExecuteCallbacksPanic_ARM(t *testing.T) {
	callbacks = nil
	var called int32

	callbacksMu.Lock()
	callbacks = append(callbacks, func() {
		panic("test panic")
	})
	callbacks = append(callbacks, func() {
		atomic.AddInt32(&called, 1)
	})
	callbacksMu.Unlock()

	// Should not panic, second callback should still execute
	// 不应该 panic，第二个回调应该仍然执行
	executeCallbacks()

	time.Sleep(10 * time.Millisecond)
	if atomic.LoadInt32(&called) != 1 {
		t.Errorf("Expected second callback to be called, got %d", called)
	}
}
