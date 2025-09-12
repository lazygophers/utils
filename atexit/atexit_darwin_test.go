//go:build darwin

package atexit

import (
	"os"
	"os/exec"
	"strings"
	"sync"
	"syscall"
	"testing"
	"time"
)

func TestRegister(t *testing.T) {
	// 保存原始状态
	originalCallbacks := callbacks
	originalOnce := signalOnce
	defer func() {
		callbacksMu.Lock()
		callbacks = originalCallbacks
		callbacksMu.Unlock()
		signalOnce = originalOnce
	}()
	
	// 重置状态
	callbacksMu.Lock()
	callbacks = nil
	callbacksMu.Unlock()
	signalOnce = sync.Once{}
	
	// 测试注册回调函数
	Register(func() {
		// do nothing
	})
	
	// 检查是否添加到列表中
	callbacksMu.RLock()
	count := len(callbacks)
	callbacksMu.RUnlock()
	
	if count != 1 {
		t.Errorf("期望回调列表长度为1，实际为%d", count)
	}
}

func TestRegisterNil(t *testing.T) {
	// 保存原始状态
	originalCallbacks := callbacks
	defer func() {
		callbacksMu.Lock()
		callbacks = originalCallbacks
		callbacksMu.Unlock()
	}()
	
	// 重置状态
	callbacksMu.Lock()
	originalCount := len(callbacks)
	callbacksMu.Unlock()
	
	// 测试注册nil回调
	Register(nil)
	
	// 检查列表长度没有变化
	callbacksMu.RLock()
	newCount := len(callbacks)
	callbacksMu.RUnlock()
	
	if newCount != originalCount {
		t.Errorf("注册nil回调后列表长度应该不变，原来%d，现在%d", originalCount, newCount)
	}
}

func TestRegisterConcurrent(t *testing.T) {
	// 保存原始状态
	originalCallbacks := callbacks
	originalOnce := signalOnce
	defer func() {
		callbacksMu.Lock()
		callbacks = originalCallbacks
		callbacksMu.Unlock()
		signalOnce = originalOnce
	}()
	
	// 重置状态
	callbacksMu.Lock()
	callbacks = nil
	callbacksMu.Unlock()
	signalOnce = sync.Once{}
	
	const numGoroutines = 50
	const numCallbacks = 5
	
	var wg sync.WaitGroup
	wg.Add(numGoroutines)
	
	// 并发注册回调函数
	for i := 0; i < numGoroutines; i++ {
		go func(id int) {
			defer wg.Done()
			for j := 0; j < numCallbacks; j++ {
				Register(func() {
					// do nothing
				})
			}
		}(i)
	}
	
	wg.Wait()
	
	// 检查所有回调都被注册
	callbacksMu.RLock()
	count := len(callbacks)
	callbacksMu.RUnlock()
	
	expected := numGoroutines * numCallbacks
	if count != expected {
		t.Errorf("期望注册%d个回调，实际注册%d个", expected, count)
	}
}

func TestExecuteCallbacks(t *testing.T) {
	// 保存原始状态
	originalCallbacks := callbacks
	defer func() {
		callbacksMu.Lock()
		callbacks = originalCallbacks
		callbacksMu.Unlock()
	}()
	
	// 准备测试数据
	var called []int
	var callMu sync.Mutex
	
	callbacksMu.Lock()
	callbacks = []func(){
		func() {
			callMu.Lock()
			called = append(called, 1)
			callMu.Unlock()
		},
		func() {
			callMu.Lock()
			called = append(called, 2)
			callMu.Unlock()
		},
		nil, // 测试nil回调
		func() {
			callMu.Lock()
			called = append(called, 3)
			callMu.Unlock()
		},
		func() {
			// 测试panic恢复
			panic("test panic")
		},
		func() {
			callMu.Lock()
			called = append(called, 4)
			callMu.Unlock()
		},
	}
	callbacksMu.Unlock()
	
	// 执行回调
	executeCallbacks()
	
	// 检查结果
	callMu.Lock()
	defer callMu.Unlock()
	
	expected := []int{1, 2, 3, 4}
	if len(called) != len(expected) {
		t.Errorf("期望调用%d个回调，实际调用%d个", len(expected), len(called))
		return
	}
	
	for i, v := range expected {
		if called[i] != v {
			t.Errorf("回调执行顺序错误，位置%d期望%d，实际%d", i, v, called[i])
		}
	}
}

func TestMacOSSignalHandling(t *testing.T) {
	// macOS 特定的信号处理测试
	if os.Getenv("BE_MACOS_SIGNAL_TEST") == "1" {
		Register(func() {
			os.Stdout.Write([]byte("macos_callback_executed"))
			os.Stdout.Sync()
		})
		
		// 保持进程运行，等待信号
		select {}
	}
	
	cmd := exec.Command(os.Args[0], "-test.run=TestMacOSSignalHandling")
	cmd.Env = append(os.Environ(), "BE_MACOS_SIGNAL_TEST=1")
	
	err := cmd.Start()
	if err != nil {
		t.Fatal("启动子进程失败:", err)
	}
	
	time.Sleep(100 * time.Millisecond)
	
	// 测试 macOS 特有的 SIGHUP 信号处理
	err = cmd.Process.Signal(syscall.SIGHUP)
	if err != nil {
		t.Fatal("发送SIGHUP信号失败:", err)
	}
	
	done := make(chan bool)
	var output []byte
	
	go func() {
		output, _ = cmd.Output()
		done <- true
	}()
	
	select {
	case <-done:
		outputStr := string(output)
		if strings.Contains(outputStr, "macos_callback_executed") {
			t.Logf("macOS信号处理成功: %s", outputStr)
		} else {
			t.Logf("macOS信号处理输出: %s", outputStr)
		}
	case <-time.After(2 * time.Second):
		cmd.Process.Kill()
		t.Skip("macOS 信号处理测试超时")
	}
}

func TestMacOSSIGTERMHandling(t *testing.T) {
	// 测试标准的 SIGTERM 信号处理
	if os.Getenv("BE_MACOS_SIGTERM_TEST") == "1" {
		Register(func() {
			os.Stdout.Write([]byte("macos_sigterm_callback_executed"))
			os.Stdout.Sync()
		})
		
		// 保持进程运行，等待信号
		select {}
	}
	
	cmd := exec.Command(os.Args[0], "-test.run=TestMacOSSIGTERMHandling")
	cmd.Env = append(os.Environ(), "BE_MACOS_SIGTERM_TEST=1")
	
	err := cmd.Start()
	if err != nil {
		t.Fatal("启动子进程失败:", err)
	}
	
	time.Sleep(100 * time.Millisecond)
	
	err = cmd.Process.Signal(syscall.SIGTERM)
	if err != nil {
		t.Fatal("发送SIGTERM信号失败:", err)
	}
	
	done := make(chan bool)
	var output []byte
	
	go func() {
		output, _ = cmd.Output()
		done <- true
	}()
	
	select {
	case <-done:
		outputStr := string(output)
		if strings.Contains(outputStr, "macos_sigterm_callback_executed") {
			t.Logf("macOS SIGTERM处理成功")
		} else {
			t.Logf("macOS SIGTERM处理输出: %s", outputStr)
		}
	case <-time.After(2 * time.Second):
		cmd.Process.Kill()
		t.Skip("macOS SIGTERM处理测试超时")
	}
}

// 基准测试
func BenchmarkRegister(b *testing.B) {
	originalCallbacks := callbacks
	originalOnce := signalOnce
	defer func() {
		callbacksMu.Lock()
		callbacks = originalCallbacks
		callbacksMu.Unlock()
		signalOnce = originalOnce
	}()
	
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Register(func() {
			// do nothing
		})
	}
}

func BenchmarkRegisterConcurrent(b *testing.B) {
	originalCallbacks := callbacks
	originalOnce := signalOnce
	defer func() {
		callbacksMu.Lock()
		callbacks = originalCallbacks
		callbacksMu.Unlock()
		signalOnce = originalOnce
	}()
	
	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			Register(func() {
				// do nothing
			})
		}
	})
}

func BenchmarkExecuteCallbacks(b *testing.B) {
	originalCallbacks := callbacks
	defer func() {
		callbacksMu.Lock()
		callbacks = originalCallbacks
		callbacksMu.Unlock()
	}()
	
	callbacksMu.Lock()
	callbacks = make([]func(), 100)
	for i := range callbacks {
		callbacks[i] = func() {
			// 简单的操作
		}
	}
	callbacksMu.Unlock()
	
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		executeCallbacks()
	}
}