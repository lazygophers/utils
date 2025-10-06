//go:build plan9

package atexit

import (
	"os"
	"sync"
)

var (
	callbacks   []func()
	callbacksMu sync.RWMutex
	exitOnce    sync.Once
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
						// Plan9 error handling
						// Plan9错误处理
					}
				}()
				cb()
			}()
		}
	}
}

// Register registers a callback function to be called on exit
// Note: Plan9 does not support traditional Unix signals.
// Use Exit() instead of os.Exit() to ensure callbacks are executed
// 注册退出时的回调函数
// 注意：Plan9不支持传统的Unix信号
// 使用Exit()而不是os.Exit()以确保回调被执行
func Register(callback func()) {
	if callback == nil {
		return
	}

	callbacksMu.Lock()
	callbacks = append(callbacks, callback)
	callbacksMu.Unlock()
}

// Exit executes all registered callbacks and then exits with the given code
// This should be used instead of os.Exit in Plan9 environments
// 执行所有注册的回调函数然后以给定的代码退出
// 在Plan9环境中应该使用此函数而不是os.Exit
func Exit(code int) {
	executeCallbacks()
	os.Exit(code)
}
