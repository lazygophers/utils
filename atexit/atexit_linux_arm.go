//go:build linux && arm

package atexit

import (
	"os"
	"os/signal"
	"sync"
	"syscall"
)

var (
	callbacks   []func()
	callbacksMu sync.RWMutex
	signalOnce  sync.Once
)

// initSignalHandler initializes signal handler for Linux ARM platforms
// ARM (32-bit) platforms don't support gomonkey, so we use signal-based approach
// 初始化信号处理 - Linux ARM 平台
// ARM（32位）平台不支持 gomonkey，因此使用基于信号的方式
func initSignalHandler() {
	signalOnce.Do(func() {
		c := make(chan os.Signal, 1)
		// Monitor standard termination signals as recommended by Go documentation
		// According to os/signal docs: SIGINT, SIGTERM cause program to exit
		// SIGHUP is sent when program loses controlling terminal
		// SIGQUIT causes exit with stack dump (Ctrl+\)
		// 监听 Go 文档推荐的标准终止信号
		// 根据 os/signal 文档：SIGINT、SIGTERM 会导致程序退出
		// SIGHUP 在程序失去控制终端时发送
		// SIGQUIT 会导致带堆栈转储的退出（Ctrl+\）
		signal.Notify(c,
			os.Interrupt,    // Covers SIGINT on all platforms / 在所有平台上覆盖 SIGINT
			syscall.SIGTERM, // Termination request / 终止请求
			syscall.SIGHUP,  // Hangup - terminal disconnected / 终端断开
			syscall.SIGQUIT, // Quit with core dump (Ctrl+\) / 退出并转储核心（Ctrl+\）
		)

		go func() {
			<-c
			executeCallbacks()
			// Exit gracefully after handling signal
			// Signal-triggered exits are considered normal/graceful shutdowns
			// 信号触发后优雅退出
			// 信号触发的退出被视为正常/优雅的关闭
			os.Exit(0)
		}()
	})
}

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
						// Log panic if needed
						// 需要时记录 panic
					}
				}()
				cb()
			}()
		}
	}
}

// Register registers a callback function to be called on exit
// 注册退出时的回调函数
func Register(callback func()) {
	if callback == nil {
		return
	}

	// Initialize signal handler on first registration
	// 首次注册时初始化信号处理
	initSignalHandler()

	callbacksMu.Lock()
	callbacks = append(callbacks, callback)
	callbacksMu.Unlock()
}
