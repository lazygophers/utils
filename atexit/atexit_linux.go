//go:build (linux || android) && !arm

package atexit

import (
	gomonkey "github.com/agiledragon/gomonkey/v2"
	"os"
	"sync"
)

var exitCallbackList []func()
var exitCallbackListMu sync.Mutex
var exitPatches *gomonkey.Patches
var mu sync.Mutex

// hookExit intercepts os.Exit calls and executes registered callbacks
// hookExit 拦截 os.Exit 调用并执行注册的回调函数
func hookExit(code int) {
	var cbList []func()
	mu.Lock()
	if exitPatches != nil {
		exitPatches.Reset()
		exitCallbackListMu.Lock()
		cbList = exitCallbackList
		exitCallbackList = nil
		exitCallbackListMu.Unlock()
		exitPatches = nil
	}
	mu.Unlock()

	// Execute all registered callback functions
	// 执行所有注册的回调函数
	for _, cb := range cbList {
		if cb != nil {
			cb()
		}
	}
	os.Exit(code)
}

func init() {
	patches := gomonkey.ApplyFunc(os.Exit, hookExit)
	if patches == nil {
		// If patching fails, fallback to using signal handling
		// 如果 patch 失败，fallback 到使用信号处理
		return
	}
	exitPatches = patches
}

// Register registers a callback function to be called on exit
// 注册退出时的回调函数
func Register(callback func()) {
	if callback == nil {
		return
	}
	exitCallbackListMu.Lock()
	exitCallbackList = append(exitCallbackList, callback)
	exitCallbackListMu.Unlock()
}
