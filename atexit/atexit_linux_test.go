//go:build linux

package atexit

import (
	"os"
	"os/exec"
	"strings"
	"sync"
	"testing"
)

func TestRegister(t *testing.T) {
	// 保存原始状态
	originalList := exitCallbackList
	originalPatches := exitPatches
	defer func() {
		exitCallbackList = originalList
		exitPatches = originalPatches
	}()

	// 重置状态
	exitCallbackListMu.Lock()
	exitCallbackList = nil
	exitCallbackListMu.Unlock()

	// 测试注册回调函数
	Register(func() {
		// do nothing - just testing registration
	})

	// 检查是否添加到列表中
	exitCallbackListMu.Lock()
	count := len(exitCallbackList)
	exitCallbackListMu.Unlock()

	if count != 1 {
		t.Errorf("期望回调列表长度为1，实际为%d", count)
	}
}

func TestRegisterNil(t *testing.T) {
	// 保存原始状态
	originalList := exitCallbackList
	defer func() {
		exitCallbackListMu.Lock()
		exitCallbackList = originalList
		exitCallbackListMu.Unlock()
	}()

	// 重置状态
	exitCallbackListMu.Lock()
	originalCount := len(exitCallbackList)
	exitCallbackListMu.Unlock()

	// 测试注册nil回调
	Register(nil)

	// 检查列表长度没有变化
	exitCallbackListMu.Lock()
	newCount := len(exitCallbackList)
	exitCallbackListMu.Unlock()

	if newCount != originalCount {
		t.Errorf("注册nil回调后列表长度应该不变，原来%d，现在%d", originalCount, newCount)
	}
}

func TestRegisterConcurrent(t *testing.T) {
	// 保存原始状态
	originalList := exitCallbackList
	defer func() {
		exitCallbackListMu.Lock()
		exitCallbackList = originalList
		exitCallbackListMu.Unlock()
	}()

	// 重置状态
	exitCallbackListMu.Lock()
	exitCallbackList = nil
	exitCallbackListMu.Unlock()

	const numGoroutines = 100
	const numCallbacks = 10

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
	exitCallbackListMu.Lock()
	count := len(exitCallbackList)
	exitCallbackListMu.Unlock()

	expected := numGoroutines * numCallbacks
	if count != expected {
		t.Errorf("期望注册%d个回调，实际注册%d个", expected, count)
	}
}

func TestHookExit(t *testing.T) {
	// 保存原始状态
	originalList := exitCallbackList
	originalPatches := exitPatches
	defer func() {
		exitCallbackListMu.Lock()
		exitCallbackList = originalList
		exitCallbackListMu.Unlock()
		mu.Lock()
		exitPatches = originalPatches
		mu.Unlock()
	}()

	// 模拟有patch的情况
	mu.Lock()
	exitPatches = originalPatches // 确保不是nil
	mu.Unlock()

	// 准备回调函数
	var called []int
	var callMu sync.Mutex

	exitCallbackListMu.Lock()
	exitCallbackList = []func(){
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
	}
	exitCallbackListMu.Unlock()

	// 由于hookExit会调用os.Exit，我们需要在子进程中测试
	if os.Getenv("TEST_HOOK_EXIT") == "1" {
		hookExit(42)
		return
	}

	// 使用子进程测试
	cmd := exec.Command(os.Args[0], "-test.run=TestHookExit")
	cmd.Env = append(os.Environ(), "TEST_HOOK_EXIT=1")
	err := cmd.Run()

	// 检查退出码
	if exitError, ok := err.(*exec.ExitError); ok {
		if exitError.ExitCode() != 42 {
			t.Errorf("期望退出码42，实际退出码%d", exitError.ExitCode())
		}
	} else if err != nil {
		t.Errorf("子进程执行出错: %v", err)
	}
}

func TestExitBehavior(t *testing.T) {
	// 通过子进程测试完整的退出行为
	if os.Getenv("BE_CRASHER") == "1" {
		Register(func() {
			os.Stdout.Write([]byte("callback1\n"))
		})
		Register(func() {
			os.Stdout.Write([]byte("callback2\n"))
		})
		os.Exit(0)
		return
	}

	cmd := exec.Command(os.Args[0], "-test.run=TestExitBehavior")
	cmd.Env = append(os.Environ(), "BE_CRASHER=1")
	out, err := cmd.Output()
	if err != nil {
		t.Fatal("子进程执行失败:", err)
	}

	output := string(out)
	if !strings.Contains(output, "callback1") {
		t.Errorf("期望输出包含callback1，实际输出: %s", output)
	}
	if !strings.Contains(output, "callback2") {
		t.Errorf("期望输出包含callback2，实际输出: %s", output)
	}
}

// 基准测试
func BenchmarkRegister(b *testing.B) {
	// 保存原始状态
	originalList := exitCallbackList
	defer func() {
		exitCallbackListMu.Lock()
		exitCallbackList = originalList
		exitCallbackListMu.Unlock()
	}()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Register(func() {
			// do nothing
		})
	}
}

func BenchmarkRegisterConcurrent(b *testing.B) {
	// 保存原始状态
	originalList := exitCallbackList
	defer func() {
		exitCallbackListMu.Lock()
		exitCallbackList = originalList
		exitCallbackListMu.Unlock()
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
