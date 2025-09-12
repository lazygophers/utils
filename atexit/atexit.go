//go:build !linux && !windows && !darwin

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

// 初始化信号处理 - 通用Unix系统
func initSignalHandler() {
	signalOnce.Do(func() {
		c := make(chan os.Signal, 1)
		// 通用Unix信号处理
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
		
		go func() {
			<-c
			executeCallbacks()
			os.Exit(0)
		}()
	})
}

// 执行所有注册的回调函数
func executeCallbacks() {
	callbacksMu.RLock()
	cbList := make([]func(), len(callbacks))
	copy(cbList, callbacks)
	callbacksMu.RUnlock()
	
	// 按注册顺序执行回调
	for _, cb := range cbList {
		if cb != nil {
			func() {
				defer func() {
					// 捕获回调函数中的panic，避免影响其他回调的执行
					if r := recover(); r != nil {
						// 通用系统的错误处理
					}
				}()
				cb()
			}()
		}
	}
}

// Register 注册退出时的回调函数
func Register(callback func()) {
	if callback == nil {
		return
	}
	
	// 首次注册时初始化信号处理
	initSignalHandler()
	
	callbacksMu.Lock()
	callbacks = append(callbacks, callback)
	callbacksMu.Unlock()
}
