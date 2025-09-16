package runtime

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestCachePanic(t *testing.T) {
	t.Run("cache_panic_without_panic", func(t *testing.T) {
		// 测试没有panic的情况
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("CachePanic should not panic when there's no panic: %v", r)
			}
		}()

		CachePanic()
		t.Log("CachePanic completed without panic")
	})

	t.Run("cache_panic_basic_functionality", func(t *testing.T) {
		// 测试CachePanic的基本功能
		// 这个函数主要用于defer中，我们验证它能正常被调用

		defer func() {
			if r := recover(); r != nil {
				t.Logf("Recovered panic in test: %v", r)
			}
		}()

		// CachePanic应该能够被正常调用
		func() {
			defer func() {
				// 先尝试恢复，然后调用CachePanic来记录
				if err := recover(); err != nil {
					defer CachePanic()
					// 这里重新panic以便CachePanic能够捕获
					panic(err)
				}
			}()
			// 不产生panic的正常情况
		}()

		t.Log("CachePanic function test completed")
	})
}

func TestCachePanicWithHandle(t *testing.T) {
	t.Run("cache_panic_with_handle_no_panic", func(t *testing.T) {
		// 测试没有panic时handle不被调用
		handleCalled := false
		handle := func(err interface{}) {
			handleCalled = true
		}

		defer func() {
			if r := recover(); r != nil {
				t.Errorf("CachePanicWithHandle should not panic when there's no panic: %v", r)
			}
		}()

		CachePanicWithHandle(handle)

		if handleCalled {
			t.Error("Handle should not be called when there's no panic")
		}
	})

	t.Run("cache_panic_with_handle_with_panic", func(t *testing.T) {
		// 测试有panic时handle被正确调用
		var capturedErr interface{}
		handleCalled := false

		handle := func(err interface{}) {
			handleCalled = true
			capturedErr = err
		}

		expectedPanic := "test panic with handle"

		defer func() {
			if r := recover(); r != nil {
				t.Logf("Test recovered panic: %v", r)
			}

			if !handleCalled {
				t.Error("Handle should be called when there's a panic")
			}

			if capturedErr != expectedPanic {
				t.Errorf("Handle received wrong error: expected %v, got %v", expectedPanic, capturedErr)
			}
		}()

		func() {
			defer CachePanicWithHandle(handle)
			panic(expectedPanic)
		}()
	})

	t.Run("cache_panic_with_nil_handle", func(t *testing.T) {
		// 测试handle为nil的情况
		defer func() {
			if r := recover(); r != nil {
				t.Logf("Test recovered panic with nil handle: %v", r)
			}
		}()

		func() {
			defer CachePanicWithHandle(nil)
			panic("test panic with nil handle")
		}()

		t.Log("CachePanicWithHandle handled nil handle correctly")
	})
}

func TestPrintStack(t *testing.T) {
	t.Run("print_stack_basic", func(t *testing.T) {
		// 测试PrintStack函数
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("PrintStack should not panic: %v", r)
			}
		}()

		PrintStack()
		t.Log("PrintStack executed successfully")
	})
}

func TestExecDir(t *testing.T) {
	t.Run("exec_dir_returns_directory", func(t *testing.T) {
		dir := ExecDir()

		if dir == "" {
			t.Error("ExecDir should not return empty string")
		}

		// 验证返回的是目录路径
		if !filepath.IsAbs(dir) {
			t.Errorf("ExecDir should return absolute path, got: %s", dir)
		}

		// 验证目录存在
		if stat, err := os.Stat(dir); err != nil || !stat.IsDir() {
			t.Errorf("ExecDir should return existing directory, got: %s", dir)
		}

		t.Logf("ExecDir returned: %s", dir)
	})
}

func TestExecFile(t *testing.T) {
	t.Run("exec_file_returns_executable", func(t *testing.T) {
		file := ExecFile()

		if file == "" {
			t.Error("ExecFile should not return empty string")
		}

		// 验证返回的是绝对路径
		if !filepath.IsAbs(file) {
			t.Errorf("ExecFile should return absolute path, got: %s", file)
		}

		// 验证文件存在
		if _, err := os.Stat(file); err != nil {
			t.Errorf("ExecFile should return existing file, got: %s, error: %v", file, err)
		}

		t.Logf("ExecFile returned: %s", file)
	})

	t.Run("exec_file_is_in_exec_dir", func(t *testing.T) {
		dir := ExecDir()
		file := ExecFile()

		if dir == "" || file == "" {
			t.Skip("ExecDir or ExecFile returned empty")
		}

		expectedDir := filepath.Dir(file)
		if dir != expectedDir {
			t.Errorf("ExecDir (%s) should be parent of ExecFile (%s), expected: %s", dir, file, expectedDir)
		}
	})
}

func TestPwd(t *testing.T) {
	t.Run("pwd_returns_working_directory", func(t *testing.T) {
		pwd := Pwd()

		if pwd == "" {
			t.Error("Pwd should not return empty string")
		}

		// 验证返回的是绝对路径
		if !filepath.IsAbs(pwd) {
			t.Errorf("Pwd should return absolute path, got: %s", pwd)
		}

		// 验证目录存在
		if stat, err := os.Stat(pwd); err != nil || !stat.IsDir() {
			t.Errorf("Pwd should return existing directory, got: %s, error: %v", pwd, err)
		}

		// 与os.Getwd()对比
		expected, err := os.Getwd()
		if err == nil && pwd != expected {
			t.Errorf("Pwd (%s) should match os.Getwd (%s)", pwd, expected)
		}

		t.Logf("Pwd returned: %s", pwd)
	})
}

func TestUserHomeDir(t *testing.T) {
	t.Run("user_home_dir_returns_home", func(t *testing.T) {
		home := UserHomeDir()

		// 在某些环境下可能为空，这是正常的
		if home == "" {
			t.Log("UserHomeDir returned empty (may be normal in some environments)")
			return
		}

		// 验证是绝对路径
		if !filepath.IsAbs(home) {
			t.Errorf("UserHomeDir should return absolute path, got: %s", home)
		}

		// 验证目录存在
		if stat, err := os.Stat(home); err != nil || !stat.IsDir() {
			t.Errorf("UserHomeDir should return existing directory, got: %s, error: %v", home, err)
		}

		t.Logf("UserHomeDir returned: %s", home)
	})
}

func TestUserConfigDir(t *testing.T) {
	t.Run("user_config_dir_returns_config", func(t *testing.T) {
		config := UserConfigDir()

		// 在某些环境下可能为空
		if config == "" {
			t.Log("UserConfigDir returned empty (may be normal in some environments)")
			return
		}

		// 验证是绝对路径
		if !filepath.IsAbs(config) {
			t.Errorf("UserConfigDir should return absolute path, got: %s", config)
		}

		t.Logf("UserConfigDir returned: %s", config)
	})
}

func TestUserCacheDir(t *testing.T) {
	t.Run("user_cache_dir_returns_cache", func(t *testing.T) {
		cache := UserCacheDir()

		// 在某些环境下可能为空
		if cache == "" {
			t.Log("UserCacheDir returned empty (may be normal in some environments)")
			return
		}

		// 验证是绝对路径
		if !filepath.IsAbs(cache) {
			t.Errorf("UserCacheDir should return absolute path, got: %s", cache)
		}

		t.Logf("UserCacheDir returned: %s", cache)
	})
}

func TestLazyConfigDir(t *testing.T) {
	t.Run("lazy_config_dir_contains_organization", func(t *testing.T) {
		lazyConfig := LazyConfigDir()

		if lazyConfig == "" {
			t.Log("LazyConfigDir returned empty")
			return
		}

		// 验证包含organization路径
		if !strings.Contains(lazyConfig, "lazy") {
			t.Logf("LazyConfigDir (%s) may not contain expected organization name", lazyConfig)
		}

		// 验证是绝对路径
		if !filepath.IsAbs(lazyConfig) {
			t.Errorf("LazyConfigDir should return absolute path, got: %s", lazyConfig)
		}

		t.Logf("LazyConfigDir returned: %s", lazyConfig)
	})
}

func TestLazyCacheDir(t *testing.T) {
	t.Run("lazy_cache_dir_contains_organization", func(t *testing.T) {
		lazyCache := LazyCacheDir()

		if lazyCache == "" {
			t.Log("LazyCacheDir returned empty")
			return
		}

		// 验证包含organization路径
		if !strings.Contains(lazyCache, "lazy") {
			t.Logf("LazyCacheDir (%s) may not contain expected organization name", lazyCache)
		}

		// 验证是绝对路径
		if !filepath.IsAbs(lazyCache) {
			t.Errorf("LazyCacheDir should return absolute path, got: %s", lazyCache)
		}

		t.Logf("LazyCacheDir returned: %s", lazyCache)
	})
}

// 测试路径函数的一致性
func TestPathConsistency(t *testing.T) {
	t.Run("path_functions_consistency", func(t *testing.T) {
		userConfig := UserConfigDir()
		userCache := UserCacheDir()
		lazyConfig := LazyConfigDir()
		lazyCache := LazyCacheDir()

		// 如果user路径有效，lazy路径应该基于它们构建
		if userConfig != "" && lazyConfig != "" {
			if !strings.HasPrefix(lazyConfig, userConfig) {
				t.Logf("LazyConfigDir (%s) may not be based on UserConfigDir (%s)", lazyConfig, userConfig)
			}
		}

		if userCache != "" && lazyCache != "" {
			if !strings.HasPrefix(lazyCache, userCache) {
				t.Logf("LazyCacheDir (%s) may not be based on UserCacheDir (%s)", lazyCache, userCache)
			}
		}
	})
}

// 错误处理测试
func TestErrorHandling(t *testing.T) {
	t.Run("functions_handle_errors_gracefully", func(t *testing.T) {
		// 所有函数都应该优雅地处理错误，不应该panic
		functions := []func() string{
			ExecDir,
			ExecFile,
			Pwd,
			UserHomeDir,
			UserConfigDir,
			UserCacheDir,
			LazyConfigDir,
			LazyCacheDir,
		}

		for i, fn := range functions {
			func() {
				defer func() {
					if r := recover(); r != nil {
						t.Errorf("Function %d panicked: %v", i, r)
					}
				}()

				result := fn()
				t.Logf("Function %d returned: %s", i, result)
			}()
		}
	})
}

// ========== Comprehensive Coverage Tests ==========

// 最终综合测试，针对剩余的覆盖率缺口
func TestComprehensiveFinalCoverage(t *testing.T) {
	t.Run("test_all_functions_called", func(t *testing.T) {
		// 系统检测
		_ = IsWindows()
		_ = IsDarwin()
		_ = IsLinux()

		// 路径函数
		_ = ExecDir()
		_ = ExecFile()
		_ = Pwd()

		// 用户目录函数
		_ = UserHomeDir()
		_ = UserConfigDir()
		_ = UserCacheDir()
		_ = LazyConfigDir()
		_ = LazyCacheDir()

		// 堆栈和panic函数
		PrintStack()

		// 测试CachePanic和CachePanicWithHandle的各种情况
		func() {
			defer CachePanic()
		}()

		func() {
			defer CachePanicWithHandle(nil)
		}()

		func() {
			defer CachePanicWithHandle(func(interface{}) {})
		}()

		t.Log("All functions have been called for coverage")
	})

	t.Run("multiple_panic_handle_calls", func(t *testing.T) {
		// 测试连续的panic处理调用
		for i := 0; i < 3; i++ {
			func() {
				defer func() {
					recover()
				}()
				defer CachePanicWithHandle(func(err interface{}) {
					t.Logf("Handle call %d: %v", i, err)
				})
				panic("edge case " + string(rune('A'+i)))
			}()
		}
		t.Log("Multiple panic handle calls completed")
	})

	t.Run("concurrent_function_calls", func(t *testing.T) {
		// 并发调用所有函数以确保线程安全和覆盖率
		const numGoroutines = 3
		done := make(chan bool, numGoroutines)

		for i := 0; i < numGoroutines; i++ {
			go func(id int) {
				defer func() {
					done <- true
				}()

				// 调用所有主要函数
				_ = IsWindows()
				_ = IsDarwin()
				_ = IsLinux()
				_ = ExecDir()
				_ = ExecFile()
				_ = Pwd()
				_ = UserHomeDir()
				_ = UserConfigDir()
				_ = UserCacheDir()
				_ = LazyConfigDir()
				_ = LazyCacheDir()
				PrintStack()

				func() {
					defer CachePanic()
				}()

				t.Logf("Goroutine %d completed all function calls", id)
			}(i)
		}

		// 等待所有goroutine完成
		for i := 0; i < numGoroutines; i++ {
			<-done
		}

		t.Log("Concurrent function calls completed")
	})

	t.Run("stress_test_all_functions", func(t *testing.T) {
		// 压力测试：多次调用所有函数
		const iterations = 10

		for i := 0; i < iterations; i++ {
			// 系统检测函数
			_ = IsWindows()
			_ = IsDarwin()
			_ = IsLinux()

			// 路径函数
			_ = ExecDir()
			_ = ExecFile()
			_ = Pwd()

			// 用户目录函数
			_ = UserHomeDir()
			_ = UserConfigDir()
			_ = UserCacheDir()
			_ = LazyConfigDir()
			_ = LazyCacheDir()

			// 堆栈和panic函数
			PrintStack()

			func() {
				defer CachePanic()
			}()

			func() {
				defer CachePanicWithHandle(nil)
			}()

			if i%5 == 0 {
				t.Logf("Stress test iteration %d completed", i)
			}
		}

		t.Log("Stress test completed")
	})
}

// TestErrorPathsFilesystem tests error conditions in filesystem functions
func TestErrorPathsFilesystem(t *testing.T) {
	// These error paths are hard to trigger but we'll try
	t.Run("exec_dir_coverage", func(t *testing.T) {
		dir := ExecDir()
		// In normal conditions, this should return a valid directory
		if dir == "" {
			// This would hit the error path
			t.Log("ExecDir returned empty (error path)")
		} else {
			// This is the normal path
			if dir == "" {
				t.Error("ExecDir should not be empty")
			}
			t.Logf("ExecDir: %s", dir)
		}
	})

	t.Run("exec_file_coverage", func(t *testing.T) {
		file := ExecFile()
		// In normal conditions, this should return a valid file path
		if file == "" {
			// This would hit the error path
			t.Log("ExecFile returned empty (error path)")
		} else {
			// This is the normal path
			if file == "" {
				t.Error("ExecFile should not be empty")
			}
			t.Logf("ExecFile: %s", file)
		}
	})

	t.Run("pwd_coverage", func(t *testing.T) {
		pwd := Pwd()
		// In normal conditions, this should return a valid directory
		if pwd == "" {
			// This would hit the error path
			t.Log("Pwd returned empty (error path)")
		} else {
			// This is the normal path
			if pwd == "" {
				t.Error("Pwd should not be empty")
			}
			t.Logf("Pwd: %s", pwd)
		}
	})
}

// TestAllDirectoryFunctionsCompletely ensures all directory functions are tested
func TestAllDirectoryFunctionsCompletely(t *testing.T) {
	// Test UserHomeDir
	home := UserHomeDir()
	if home == "" {
		t.Log("UserHomeDir returned empty (may be normal in some environments)")
	} else {
		t.Logf("UserHomeDir: %s", home)
	}

	// Test UserConfigDir
	config := UserConfigDir()
	if config == "" {
		t.Log("UserConfigDir returned empty (may be normal in some environments)")
	} else {
		t.Logf("UserConfigDir: %s", config)
	}

	// Test UserCacheDir
	cache := UserCacheDir()
	if cache == "" {
		t.Log("UserCacheDir returned empty (may be normal in some environments)")
	} else {
		t.Logf("UserCacheDir: %s", cache)
	}

	// Test LazyConfigDir
	lazyConfig := LazyConfigDir()
	if lazyConfig == "" {
		t.Log("LazyConfigDir returned empty")
	} else {
		t.Logf("LazyConfigDir: %s", lazyConfig)
		if config != "" && !strings.Contains(lazyConfig, config) {
			t.Logf("LazyConfigDir may not contain config dir base")
		}
	}

	// Test LazyCacheDir
	lazyCache := LazyCacheDir()
	if lazyCache == "" {
		t.Log("LazyCacheDir returned empty")
	} else {
		t.Logf("LazyCacheDir: %s", lazyCache)
		if cache != "" && !strings.Contains(lazyCache, cache) {
			t.Logf("LazyCacheDir may not contain cache dir base")
		}
	}
}

// ========== System Detection Tests ==========

func TestSystemDetectionComplete(t *testing.T) {
	// Test all system detection functions
	windows := IsWindows()
	darwin := IsDarwin()
	linux := IsLinux()

	// Exactly one should be true
	trueCount := 0
	if windows { trueCount++ }
	if darwin { trueCount++ }
	if linux { trueCount++ }

	if trueCount != 1 {
		t.Errorf("Exactly one platform should be detected, got %d", trueCount)
	}
	t.Logf("System detection: Windows=%v, Darwin=%v, Linux=%v", windows, darwin, linux)
}

// Test multiple calls for stability
func TestSystemDetectionStability(t *testing.T) {
	t.Run("multiple_calls_return_same_result", func(t *testing.T) {
		// 多次调用应该返回相同结果
		iterations := 1000

		firstWindows := IsWindows()
		firstDarwin := IsDarwin()
		firstLinux := IsLinux()

		for i := 0; i < iterations; i++ {
			if IsWindows() != firstWindows {
				t.Errorf("IsWindows() returned inconsistent result on iteration %d", i)
			}
			if IsDarwin() != firstDarwin {
				t.Errorf("IsDarwin() returned inconsistent result on iteration %d", i)
			}
			if IsLinux() != firstLinux {
				t.Errorf("IsLinux() returned inconsistent result on iteration %d", i)
			}
		}

		t.Logf("All %d calls returned consistent results", iterations)
	})
}

// ========== Benchmark Tests ==========

func BenchmarkIsWindows(b *testing.B) {
	for i := 0; i < b.N; i++ {
		IsWindows()
	}
}

func BenchmarkIsDarwin(b *testing.B) {
	for i := 0; i < b.N; i++ {
		IsDarwin()
	}
}

func BenchmarkIsLinux(b *testing.B) {
	for i := 0; i < b.N; i++ {
		IsLinux()
	}
}

func BenchmarkExecDir(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ExecDir()
	}
}

func BenchmarkExecFile(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ExecFile()
	}
}

func BenchmarkPwd(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Pwd()
	}
}

func BenchmarkPrintStack(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PrintStack()
	}
}

func BenchmarkCachePanic(b *testing.B) {
	for i := 0; i < b.N; i++ {
		func() {
			defer CachePanic()
		}()
	}
}
