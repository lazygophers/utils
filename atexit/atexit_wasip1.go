//go:build wasip1

package atexit

import (
	"os"
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
						// WASI error handling
						// WASI错误处理
					}
				}()
				cb()
			}()
		}
	}
}

// Register registers a callback function to be called on exit
// Note: WASI (WebAssembly System Interface) does not support signal handling.
// Callbacks must be explicitly called using executeCallbacks() or Exit()
// 注册退出时的回调函数
// 注意：WASI（WebAssembly系统接口）不支持信号处理
// 必须使用executeCallbacks()或Exit()显式调用回调
func Register(callback func()) {
	if callback == nil {
		return
	}

	callbacksMu.Lock()
	callbacks = append(callbacks, callback)
	callbacksMu.Unlock()
}

// Exit executes all registered callbacks and then exits with the given code
// This should be used instead of os.Exit in WASI environments
// 执行所有注册的回调函数然后以给定的代码退出
// 在WASI环境中应该使用此函数而不是os.Exit
func Exit(code int) {
	executeCallbacks()
	os.Exit(code)
}
