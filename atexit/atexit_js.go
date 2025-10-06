//go:build js

package atexit

import (
	"sync"
)

var (
	callbacks   []func()
	callbacksMu sync.RWMutex
)

// executeCallbacks executes all registered callback functions
// 执行所有注册的回调函数
func executeCallbacks() {
	callbacksMu.RLock()
	cbList := make([]func(), len(callbacks))
	copy(cbList, callbacks)
	callbacksMu.RUnlock()

	// Execute callbacks in registration order
	// 按注册顺序执行回调
	for _, cb := range cbList {
		if cb != nil {
			func() {
				defer func() {
					// Catch panics from callbacks to prevent affecting other callbacks
					// 捕获回调函数中的panic，避免影响其他回调的执行
					if r := recover(); r != nil {
						// JS/WASM error handling
						// JS/WASM错误处理
					}
				}()
				cb()
			}()
		}
	}
}

// Register registers a callback function to be called on exit
// Note: In JS/WASM environment, signal handling is not available.
// Callbacks must be explicitly called using executeCallbacks() or Exit()
// 注册退出时的回调函数
// 注意：在JS/WASM环境中，信号处理不可用
// 必须使用executeCallbacks()或Exit()显式调用回调
func Register(callback func()) {
	if callback == nil {
		return
	}

	callbacksMu.Lock()
	callbacks = append(callbacks, callback)
	callbacksMu.Unlock()
}

// Exit executes all registered callbacks and then exits
// This should be used instead of os.Exit in JS/WASM environments
// 执行所有注册的回调函数然后退出
// 在JS/WASM环境中应该使用此函数而不是os.Exit
func Exit(code int) {
	executeCallbacks()
	// In JS/WASM, we can't really exit, but we can panic to stop execution
	// 在JS/WASM中，我们不能真正退出，但可以使用panic停止执行
	if code != 0 {
		panic("exit with non-zero code")
	}
}
