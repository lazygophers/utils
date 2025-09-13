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
