//go:build windows

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

// initSignalHandler initializes signal handler for Windows-specific signals
// 初始化信号处理 - Windows特定信号
func initSignalHandler() {
	signalOnce.Do(func() {
		c := make(chan os.Signal, 1)
		// Monitor Windows-supported signals
		// According to os/signal docs: ^C (Ctrl-C) sends os.Interrupt
		// Windows also sends SIGTERM for CTRL_CLOSE_EVENT, CTRL_LOGOFF_EVENT, CTRL_SHUTDOWN_EVENT
		// 监听 Windows 支持的信号
		// 根据 os/signal 文档：^C（Ctrl-C）发送 os.Interrupt
		// Windows 还会为 CTRL_CLOSE_EVENT、CTRL_LOGOFF_EVENT、CTRL_SHUTDOWN_EVENT 发送 SIGTERM
		signal.Notify(c,
			os.Interrupt,    // Ctrl+C or Ctrl+Break / Ctrl+C 或 Ctrl+Break
			syscall.SIGTERM, // Close, Logoff, or Shutdown events / 关闭、注销或关机事件
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
						// Windows can use event log for logging
						// Windows 可以使用事件日志记录
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
