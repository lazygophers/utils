//go:build aix

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

// initSignalHandler initializes signal handler for AIX
// 初始化信号处理 - AIX系统
func initSignalHandler() {
	signalOnce.Do(func() {
		c := make(chan os.Signal, 1)
		// Monitor standard Unix termination signals
		// AIX is a Unix system and supports standard signal set
		// 监听标准 Unix 终止信号
		// AIX 是 Unix 系统，支持标准信号集
		signal.Notify(c,
			os.Interrupt,    // Covers SIGINT / 覆盖 SIGINT
			syscall.SIGTERM, // Termination request / 终止请求
			syscall.SIGHUP,  // Hangup - terminal disconnected / 终端断开
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
						// AIX system error handling
						// AIX系统的错误处理
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
