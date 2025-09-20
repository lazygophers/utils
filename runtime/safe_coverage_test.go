package runtime

import (
	"fmt"
	"os"
	"path/filepath"
	"testing"
)

// TestRuntimeFunctionsSafely 安全地测试运行时函数
func TestRuntimeFunctionsSafely(t *testing.T) {
	t.Run("TestSystemFunctions", func(t *testing.T) {
		// 测试系统检测函数
		isWindows := IsWindows()
		isDarwin := IsDarwin()
		isLinux := IsLinux()

		t.Logf("IsWindows: %t", isWindows)
		t.Logf("IsDarwin: %t", isDarwin)
		t.Logf("IsLinux: %t", isLinux)

		// 验证至少有一个为true
		if !isWindows && !isDarwin && !isLinux {
			t.Error("至少应该有一个系统类型为true")
		}
	})

	t.Run("TestFileSystemFunctions", func(t *testing.T) {
		// 测试文件系统相关函数
		execDir := ExecDir()
		execFile := ExecFile()
		pwd := Pwd()

		t.Logf("ExecDir: %s", execDir)
		t.Logf("ExecFile: %s", execFile)
		t.Logf("Pwd: %s", pwd)

		// 验证返回的路径
		if execDir != "" {
			if !filepath.IsAbs(execDir) {
				t.Error("ExecDir应该返回绝对路径")
			}
		}

		if execFile != "" {
			if !filepath.IsAbs(execFile) {
				t.Error("ExecFile应该返回绝对路径")
			}
		}

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

	t.Run("TestUserDirFunctions", func(t *testing.T) {
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

	t.Run("TestPrintStackSafely", func(t *testing.T) {
		// 测试PrintStack函数
		PrintStack()
		t.Log("PrintStack执行完成")
	})
}

// TestPanicRecoverySafely 安全地测试panic恢复功能
func TestPanicRecoverySafely(t *testing.T) {
	t.Run("TestCachePanicBasic", func(t *testing.T) {
		// 测试基本的CachePanic功能
		executed := false
		func() {
			defer CachePanic()
			defer func() {
				executed = true
			}()
			// 不panic，只测试正常执行
		}()

		if !executed {
			t.Error("defer函数应该被执行")
		}
		t.Log("CachePanic正常情况测试通过")
	})

	t.Run("TestCachePanicWithHandleBasic", func(t *testing.T) {
		// 测试基本的CachePanicWithHandle功能
		handleCalled := false
		handle := func(err interface{}) {
			handleCalled = true
		}

		executed := false
		func() {
			defer CachePanicWithHandle(handle)
			defer func() {
				executed = true
			}()
			// 不panic，只测试正常执行
		}()

		if !executed {
			t.Error("defer函数应该被执行")
		}
		if handleCalled {
			t.Error("没有panic时handle不应该被调用")
		}
		t.Log("CachePanicWithHandle正常情况测试通过")
	})

	t.Run("TestCachePanicWithNilHandle", func(t *testing.T) {
		// 测试nil handle的情况
		executed := false
		func() {
			defer CachePanicWithHandle(nil)
			defer func() {
				executed = true
			}()
			// 不panic，只测试正常执行
		}()

		if !executed {
			t.Error("defer函数应该被执行")
		}
		t.Log("CachePanicWithHandle nil handle测试通过")
	})
}

// TestPanicRecoveryWithActualPanic 使用受控方式测试真正的panic恢复
func TestPanicRecoveryWithActualPanic(t *testing.T) {
	t.Run("TestCachePanicActual", func(t *testing.T) {
		// 在这个测试中，我们会在子goroutine中触发panic
		// 并验证CachePanic能正确处理
		done := make(chan bool, 1)
		var panicOccurred bool

		go func() {
			defer func() {
				done <- true
			}()
			defer CachePanic()

			// 设置标志表示我们即将panic
			panicOccurred = true
			panic("test panic for CachePanic")
		}()

		// 等待goroutine完成
		<-done

		if !panicOccurred {
			t.Error("panic应该已经发生")
		}
		t.Log("CachePanic成功处理了panic")
	})

	t.Run("TestCachePanicWithHandleActual", func(t *testing.T) {
		done := make(chan bool, 1)
		handleCalled := false
		var capturedErr interface{}

		handle := func(err interface{}) {
			handleCalled = true
			capturedErr = err
		}

		go func() {
			defer func() {
				done <- true
			}()
			defer CachePanicWithHandle(handle)

			panic("test panic for CachePanicWithHandle")
		}()

		// 等待goroutine完成
		<-done

		if !handleCalled {
			t.Error("handle函数应该被调用")
		}
		if capturedErr == nil {
			t.Error("handle函数应该接收到错误")
		}
		if capturedErr != "test panic for CachePanicWithHandle" {
			t.Errorf("捕获的错误不匹配: %v", capturedErr)
		}
		t.Log("CachePanicWithHandle成功处理了panic并调用了handle")
	})

	t.Run("TestMultiplePanicTypes", func(t *testing.T) {
		// 测试不同类型的panic
		panicValues := []interface{}{
			"string panic",
			42,
			struct{ msg string }{msg: "struct panic"},
			[]string{"slice", "panic"},
		}

		for i, panicValue := range panicValues {
			t.Run(fmt.Sprintf("PanicType_%d", i), func(t *testing.T) {
				done := make(chan bool, 1)
				handleCalled := false
				var capturedErr interface{}

				handle := func(err interface{}) {
					handleCalled = true
					capturedErr = err
				}

				go func() {
					defer func() {
						done <- true
					}()
					defer CachePanicWithHandle(handle)

					panic(panicValue)
				}()

				<-done

				if !handleCalled {
					t.Errorf("Handle函数应该被调用 (panic类型: %T)", panicValue)
				}
				if capturedErr != panicValue {
					t.Errorf("捕获的错误不匹配，期望: %v, 实际: %v", panicValue, capturedErr)
				}
				t.Logf("成功处理panic类型 %T: %v", panicValue, panicValue)
			})
		}
	})
}

// TestEdgeCases 测试边缘情况
func TestEdgeCases(t *testing.T) {
	t.Run("TestEmptyStackScenario", func(t *testing.T) {
		// 这个测试主要是为了覆盖PrintStack中stack为空的情况
		// 虽然在正常情况下这很难发生
		PrintStack()
		t.Log("PrintStack边缘情况测试完成")
	})

	t.Run("TestFileSystemErrorPaths", func(t *testing.T) {
		// 测试各种文件系统函数在正常情况下的行为
		execDir := ExecDir()
		execFile := ExecFile()
		pwd := Pwd()

		// 即使有错误，这些函数也应该返回（可能为空字符串）
		t.Logf("ExecDir结果: '%s'", execDir)
		t.Logf("ExecFile结果: '%s'", execFile)
		t.Logf("Pwd结果: '%s'", pwd)

		// 测试用户目录函数
		homeDir := UserHomeDir()
		configDir := UserConfigDir()
		cacheDir := UserCacheDir()
		lazyConfigDir := LazyConfigDir()
		lazyCacheDir := LazyCacheDir()

		t.Logf("用户目录函数执行完成")
		t.Logf("HomeDir: '%s'", homeDir)
		t.Logf("ConfigDir: '%s'", configDir)
		t.Logf("CacheDir: '%s'", cacheDir)
		t.Logf("LazyConfigDir: '%s'", lazyConfigDir)
		t.Logf("LazyCacheDir: '%s'", lazyCacheDir)
	})

	t.Run("TestSystemDetectionComplete", func(t *testing.T) {
		// 完整测试系统检测
		isWindows := IsWindows()
		isDarwin := IsDarwin()
		isLinux := IsLinux()

		// 验证逻辑一致性
		count := 0
		if isWindows {
			count++
		}
		if isDarwin {
			count++
		}
		if isLinux {
			count++
		}

		if count == 0 {
			t.Error("至少应该检测到一个操作系统")
		}
		if count > 1 {
			t.Error("不应该同时检测到多个操作系统")
		}

		t.Logf("系统检测结果: Windows=%t, Darwin=%t, Linux=%t", isWindows, isDarwin, isLinux)
	})
}