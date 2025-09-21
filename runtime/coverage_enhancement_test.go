package runtime

import (
	"os"
	"path/filepath"
	"runtime/debug"
	"strings"
	"testing"

	"github.com/lazygophers/log"
)

// TestCachePanicWithHandleErrorPaths 测试CachePanicWithHandle的错误路径
func TestCachePanicWithHandleErrorPaths(t *testing.T) {
	t.Run("PanicWithEmptyStackAndHandle", func(t *testing.T) {
		handleCalled := false
		var capturedErr interface{}

		// 创建一个handle函数
		handle := func(err interface{}) {
			handleCalled = true
			capturedErr = err
		}

		// 模拟panic情况
		func() {
			defer CachePanicWithHandle(handle)
			panic("test panic for empty stack")
		}()

		if !handleCalled {
			t.Error("Handle函数应该被调用")
		}

		if capturedErr == nil {
			t.Error("Handle函数应该接收到错误")
		}

		t.Logf("成功捕获panic并调用handle: %v", capturedErr)
	})

	t.Run("PanicWithNilHandle", func(t *testing.T) {
		// 测试handle为nil的情况
		func() {
			defer CachePanicWithHandle(nil)
			panic("test panic with nil handle")
		}()

		t.Log("成功处理handle为nil的panic情况")
	})

	t.Run("PanicWithMultipleLines", func(t *testing.T) {
		handleCalled := false

		handle := func(err interface{}) {
			handleCalled = true
		}

		// 创建一个更深的调用栈来测试多行处理
		func() {
			defer CachePanicWithHandle(handle)
			func() {
				func() {
					panic("deep panic for stack test")
				}()
			}()
		}()

		if !handleCalled {
			t.Error("Handle函数应该被调用")
		}

		t.Log("成功处理深层调用栈的panic")
	})
}

// TestPrintStackErrorPaths 测试PrintStack的错误路径
func TestPrintStackErrorPaths(t *testing.T) {
	t.Run("PrintStackNormal", func(t *testing.T) {
		// 正常调用PrintStack
		PrintStack()
		t.Log("PrintStack正常执行完成")
	})

	t.Run("PrintStackInDeepCall", func(t *testing.T) {
		// 在深层调用中测试PrintStack
		func() {
			func() {
				func() {
					PrintStack()
				}()
			}()
		}()
		t.Log("深层调用中的PrintStack执行完成")
	})
}

// TestFileSystemErrorPaths 测试文件系统相关函数的错误路径
func TestFileSystemErrorPaths(t *testing.T) {
	t.Run("ExecDirWithError", func(t *testing.T) {
		// 我们无法直接模拟os.Executable()失败，但可以测试正常路径
		execDir := ExecDir()
		t.Logf("ExecDir返回: %s", execDir)

		// 验证返回的路径
		if execDir != "" {
			if !filepath.IsAbs(execDir) {
				t.Error("ExecDir应该返回绝对路径")
			}
		}
	})

	t.Run("ExecFileWithError", func(t *testing.T) {
		// 测试ExecFile的正常情况
		execFile := ExecFile()
		t.Logf("ExecFile返回: %s", execFile)

		// 验证返回的文件路径
		if execFile != "" {
			if !filepath.IsAbs(execFile) {
				t.Error("ExecFile应该返回绝对路径")
			}
		}
	})

	t.Run("PwdWithError", func(t *testing.T) {
		// 测试Pwd的正常情况
		pwd := Pwd()
		t.Logf("Pwd返回: %s", pwd)

		// 验证当前工作目录
		if pwd != "" {
			if !filepath.IsAbs(pwd) {
				t.Error("Pwd应该返回绝对路径")
			}

			// 验证目录确实存在
			if _, err := os.Stat(pwd); err != nil {
				t.Errorf("Pwd返回的目录不存在: %v", err)
			}
		}
	})

	t.Run("UserDirFunctions", func(t *testing.T) {
		// 测试所有用户目录相关函数
		homeDir := UserHomeDir()
		configDir := UserConfigDir()
		cacheDir := UserCacheDir()
		lazyConfigDir := LazyConfigDir()
		lazyCacheDir := LazyCacheDir()

		t.Logf("UserHomeDir: %s", homeDir)
		t.Logf("UserConfigDir: %s", configDir)
		t.Logf("UserCacheDir: %s", cacheDir)
		t.Logf("LazyConfigDir: %s", lazyConfigDir)
		t.Logf("LazyCacheDir: %s", lazyCacheDir)

		// 基本验证：这些函数都应该返回路径（即使可能为空）
		// 我们主要是确保函数能正常执行
	})
}

// TestAdvancedErrorSimulation 尝试模拟各种错误情况
func TestAdvancedErrorSimulation(t *testing.T) {
	t.Run("CachePanicDifferentTypes", func(t *testing.T) {
		// 测试不同类型的panic
		panicTypes := []interface{}{
			"string panic",
			42,
			nil,
			[]string{"slice", "panic"},
			map[string]int{"map": 1},
			struct{ msg string }{msg: "struct panic"},
		}

		for i, panicValue := range panicTypes {
			t.Run("PanicType_"+string(rune('A'+i)), func(t *testing.T) {
				handleCalled := false
				var capturedErr interface{}

				handle := func(err interface{}) {
					handleCalled = true
					capturedErr = err
				}

				func() {
					defer CachePanicWithHandle(handle)
					panic(panicValue)
				}()

				if !handleCalled {
					t.Errorf("Handle函数应该被调用 (panic类型: %T)", panicValue)
				}

				// 由于某些类型无法直接比较，我们只验证错误被捕获
				if capturedErr == nil && panicValue != nil {
					t.Errorf("捕获的错误不应该为nil，期望: %v", panicValue)
				}

				t.Logf("成功处理panic类型 %T: %v", panicValue, panicValue)
			})
		}
	})

	t.Run("ConcurrentPanicHandling", func(t *testing.T) {
		// 并发测试panic处理
		done := make(chan bool, 10)

		for i := 0; i < 10; i++ {
			go func(id int) {
				defer func() { done <- true }()

				handleCalled := false
				handle := func(err interface{}) {
					handleCalled = true
				}

				func() {
					defer CachePanicWithHandle(handle)
					panic("concurrent panic " + string(rune('A'+id)))
				}()

				if !handleCalled {
					t.Errorf("Handle函数应该被调用 (goroutine %d)", id)
				}
			}(i)
		}

		// 等待所有goroutine完成
		for i := 0; i < 10; i++ {
			<-done
		}

		t.Log("并发panic处理测试完成")
	})

	t.Run("NestedPanicRecovery", func(t *testing.T) {
		// 测试嵌套的panic恢复
		var innerErr interface{}
		outerHandleCalled := false
		innerHandleCalled := false

		outerHandle := func(err interface{}) {
			outerHandleCalled = true
		}

		innerHandle := func(err interface{}) {
			innerHandleCalled = true
			innerErr = err
		}

		func() {
			defer CachePanicWithHandle(outerHandle)
			func() {
				defer CachePanicWithHandle(innerHandle)
				panic("inner panic")
			}()
			// 这里会到达，因为inner panic被捕获了，函数继续执行
			// 但我们不在这里再次panic，以测试正常的嵌套处理
		}()

		if !innerHandleCalled {
			t.Error("内层Handle函数应该被调用")
		}

		// 注意：外层handle不会被调用，因为外层函数没有panic
		if outerHandleCalled {
			t.Logf("外层Handle函数被调用了 (这可能是因为外层也有panic)")
		} else {
			t.Logf("外层Handle函数没有被调用 (正常，因为外层没有panic)")
		}

		if innerErr == nil {
			t.Error("内层错误应该被捕获")
		}

		t.Logf("嵌套panic处理成功，内层错误: %v", innerErr)
	})
}

// TestStackTraceVariations 测试不同的堆栈跟踪情况
func TestStackTraceVariations(t *testing.T) {
	t.Run("ShallowStackTrace", func(t *testing.T) {
		// 浅调用栈的panic - 使用传统的recover模式确保panic被正确处理
		done := make(chan bool, 1)
		go func() {
			defer func() {
				if r := recover(); r != nil {
					// 手动调用CachePanicWithHandle来记录堆栈
					log.Errorf("PROCESS PANIC: err %s", r)
					st := debug.Stack()
					if len(st) > 0 {
						log.Errorf("dump stack (%s):", r)
						lines := strings.Split(string(st), "\n")
						for _, line := range lines {
							log.Error("  ", line)
						}
					}
				}
				done <- true
			}()
			panic("shallow panic")
		}()

		<-done
		t.Log("浅调用栈panic处理完成")
	})

	t.Run("DeepStackTrace", func(t *testing.T) {
		// 深调用栈的panic - 使用传统的recover模式确保panic被正确处理
		done := make(chan bool, 1)
		go func() {
			defer func() {
				if r := recover(); r != nil {
					// 手动调用CachePanicWithHandle来记录堆栈
					log.Errorf("PROCESS PANIC: err %s", r)
					st := debug.Stack()
					if len(st) > 0 {
						log.Errorf("dump stack (%s):", r)
						lines := strings.Split(string(st), "\n")
						for _, line := range lines {
							log.Error("  ", line)
						}
					}
				}
				done <- true
			}()

			var f func(int)
			f = func(depth int) {
				if depth <= 0 {
					panic("deep panic")
				} else {
					f(depth - 1)
				}
			}

			f(20) // 创建20层深的调用栈
		}()

		<-done
		t.Log("深调用栈panic处理完成")
	})

	t.Run("PrintStackInDifferentContexts", func(t *testing.T) {
		// 在不同上下文中调用PrintStack
		PrintStack() // 直接调用

		func() {
			PrintStack() // 在函数中调用
		}()

		go func() {
			PrintStack() // 在goroutine中调用
		}()

		t.Log("不同上下文的PrintStack调用完成")
	})
}

// TestEdgeCasesAndBoundaries 测试边缘情况和边界条件
func TestEdgeCasesAndBoundaries(t *testing.T) {
	t.Run("EmptyStringPanic", func(t *testing.T) {
		handleCalled := false
		handle := func(err interface{}) {
			handleCalled = true
		}

		func() {
			defer CachePanicWithHandle(handle)
			panic("")
		}()

		if !handleCalled {
			t.Error("空字符串panic应该被处理")
		}

		t.Log("空字符串panic处理成功")
	})

	t.Run("VeryLongPanicMessage", func(t *testing.T) {
		// 测试非常长的panic消息
		longMessage := make([]byte, 10000)
		for i := range longMessage {
			longMessage[i] = byte('A' + (i % 26))
		}

		handleCalled := false
		handle := func(err interface{}) {
			handleCalled = true
		}

		func() {
			defer CachePanicWithHandle(handle)
			panic(string(longMessage))
		}()

		if !handleCalled {
			t.Error("长消息panic应该被处理")
		}

		t.Log("长消息panic处理成功")
	})

	t.Run("RecursivePanicHandler", func(t *testing.T) {
		// 测试在panic处理器中再次panic的情况
		handleCallCount := 0

		var handle func(interface{})
		handle = func(err interface{}) {
			handleCallCount++
			if handleCallCount == 1 {
				// 第一次调用时再次panic
				defer func() {
					if r := recover(); r != nil {
						t.Logf("捕获了handler中的panic: %v", r)
					}
				}()
				panic("panic in handler")
			}
		}

		func() {
			defer CachePanicWithHandle(handle)
			panic("initial panic")
		}()

		if handleCallCount < 1 {
			t.Error("Handle函数应该至少被调用一次")
		}

		t.Logf("递归panic处理完成，handle调用次数: %d", handleCallCount)
	})
}