package runtime

import (
	"os"
	"path/filepath"
	"testing"
)

// 测试路径函数的错误处理分支
func TestPathFunctionsErrorHandling(t *testing.T) {
	t.Run("test_exec_dir_comprehensive", func(t *testing.T) {
		// 测试ExecDir函数的完整逻辑
		result := ExecDir()
		
		if result == "" {
			t.Log("ExecDir returned empty string - error path covered")
		} else {
			// 验证返回的路径是有效的目录
			if !filepath.IsAbs(result) {
				t.Errorf("ExecDir should return absolute path, got: %s", result)
			}
			
			if stat, err := os.Stat(result); err != nil || !stat.IsDir() {
				t.Errorf("ExecDir should return existing directory, got: %s", result)
			}
			
			t.Logf("ExecDir success path: %s", result)
		}
	})

	t.Run("test_exec_file_comprehensive", func(t *testing.T) {
		// 测试ExecFile函数的完整逻辑
		result := ExecFile()
		
		if result == "" {
			t.Log("ExecFile returned empty string - error path covered")
		} else {
			// 验证返回的路径是有效的文件
			if !filepath.IsAbs(result) {
				t.Errorf("ExecFile should return absolute path, got: %s", result)
			}
			
			if stat, err := os.Stat(result); err != nil || stat.IsDir() {
				t.Errorf("ExecFile should return existing file, got: %s", result)
			}
			
			t.Logf("ExecFile success path: %s", result)
		}
	})

	t.Run("test_pwd_comprehensive", func(t *testing.T) {
		// 测试Pwd函数的完整逻辑
		result := Pwd()
		
		if result == "" {
			t.Log("Pwd returned empty string - error path covered")
		} else {
			// 验证返回的路径是有效的目录
			if !filepath.IsAbs(result) {
				t.Errorf("Pwd should return absolute path, got: %s", result)
			}
			
			if stat, err := os.Stat(result); err != nil || !stat.IsDir() {
				t.Errorf("Pwd should return existing directory, got: %s", result)
			}
			
			// 与os.Getwd()比较
			expected, err := os.Getwd()
			if err == nil && result != expected {
				t.Errorf("Pwd (%s) should match os.Getwd (%s)", result, expected)
			}
			
			t.Logf("Pwd success path: %s", result)
		}
	})
}

// 测试在各种环境条件下的路径函数行为
func TestPathFunctionsUnderDifferentConditions(t *testing.T) {
	t.Run("test_functions_consistency", func(t *testing.T) {
		// 确保多次调用返回一致的结果
		results := make(map[string][]string)
		
		functions := map[string]func() string{
			"ExecDir":  ExecDir,
			"ExecFile": ExecFile,
			"Pwd":      Pwd,
		}
		
		// 多次调用每个函数
		for name, fn := range functions {
			for i := 0; i < 3; i++ {
				result := fn()
				results[name] = append(results[name], result)
			}
		}
		
		// 验证一致性
		for name, values := range results {
			first := values[0]
			for i, value := range values {
				if value != first {
					t.Errorf("Function %s returned inconsistent result at call %d: expected %s, got %s", name, i, first, value)
				}
			}
			t.Logf("Function %s returned consistent results: %s", name, first)
		}
	})

	t.Run("test_path_relationships", func(t *testing.T) {
		// 测试路径函数之间的关系
		execDir := ExecDir()
		execFile := ExecFile()
		pwd := Pwd()

		t.Logf("Path relationships - ExecDir: %s, ExecFile: %s, Pwd: %s", execDir, execFile, pwd)

		// 如果execDir和execFile都不为空，验证它们的关系
		if execDir != "" && execFile != "" {
			expectedDir := filepath.Dir(execFile)
			if execDir != expectedDir {
				t.Errorf("ExecDir (%s) should match directory of ExecFile (%s), expected: %s", execDir, execFile, expectedDir)
			}
		}

		// 验证所有返回的路径都是绝对路径
		paths := map[string]string{
			"ExecDir":  execDir,
			"ExecFile": execFile,
			"Pwd":      pwd,
		}

		for name, path := range paths {
			if path != "" && !filepath.IsAbs(path) {
				t.Errorf("%s should return absolute path, got: %s", name, path)
			}
		}
	})
}

// 测试用户目录函数的完整覆盖率
func TestUserDirectoryFunctionsComplete(t *testing.T) {
	t.Run("test_user_home_dir_complete", func(t *testing.T) {
		// 测试UserHomeDir的完整逻辑
		result := UserHomeDir()
		
		if result == "" {
			t.Log("UserHomeDir returned empty - may be normal in some environments")
		} else {
			if !filepath.IsAbs(result) {
				t.Errorf("UserHomeDir should return absolute path, got: %s", result)
			}
			t.Logf("UserHomeDir: %s", result)
		}
	})

	t.Run("test_user_config_dir_complete", func(t *testing.T) {
		// 测试UserConfigDir的完整逻辑
		result := UserConfigDir()
		
		if result == "" {
			t.Log("UserConfigDir returned empty - may be normal in some environments")
		} else {
			if !filepath.IsAbs(result) {
				t.Errorf("UserConfigDir should return absolute path, got: %s", result)
			}
			t.Logf("UserConfigDir: %s", result)
		}
	})

	t.Run("test_user_cache_dir_complete", func(t *testing.T) {
		// 测试UserCacheDir的完整逻辑
		result := UserCacheDir()
		
		if result == "" {
			t.Log("UserCacheDir returned empty - may be normal in some environments")
		} else {
			if !filepath.IsAbs(result) {
				t.Errorf("UserCacheDir should return absolute path, got: %s", result)
			}
			t.Logf("UserCacheDir: %s", result)
		}
	})

	t.Run("test_lazy_dirs_complete", func(t *testing.T) {
		// 测试LazyConfigDir和LazyCacheDir的完整逻辑
		lazyConfig := LazyConfigDir()
		lazyCache := LazyCacheDir()

		if lazyConfig == "" {
			t.Log("LazyConfigDir returned empty")
		} else {
			if !filepath.IsAbs(lazyConfig) {
				t.Errorf("LazyConfigDir should return absolute path, got: %s", lazyConfig)
			}
			t.Logf("LazyConfigDir: %s", lazyConfig)
		}

		if lazyCache == "" {
			t.Log("LazyCacheDir returned empty")
		} else {
			if !filepath.IsAbs(lazyCache) {
				t.Errorf("LazyCacheDir should return absolute path, got: %s", lazyCache)
			}
			t.Logf("LazyCacheDir: %s", lazyCache)
		}

		// 测试与基础目录的关系
		userConfig := UserConfigDir()
		userCache := UserCacheDir()

		if userConfig != "" && lazyConfig != "" {
			if !filepath.IsAbs(lazyConfig) || len(lazyConfig) <= len(userConfig) {
				t.Logf("LazyConfigDir may not be properly based on UserConfigDir")
			}
		}

		if userCache != "" && lazyCache != "" {
			if !filepath.IsAbs(lazyCache) || len(lazyCache) <= len(userCache) {
				t.Logf("LazyCacheDir may not be properly based on UserCacheDir")
			}
		}
	})
}

// 测试环境变量对路径函数的影响
func TestPathFunctionsWithEnvironmentChanges(t *testing.T) {
	t.Run("test_with_modified_environment", func(t *testing.T) {
		// 备份原始环境变量
		originalVars := map[string]string{
			"HOME":                os.Getenv("HOME"),
			"XDG_CONFIG_HOME":     os.Getenv("XDG_CONFIG_HOME"),
			"XDG_CACHE_HOME":      os.Getenv("XDG_CACHE_HOME"),
			"APPDATA":             os.Getenv("APPDATA"),
			"LocalAppData":        os.Getenv("LocalAppData"),
		}

		defer func() {
			// 恢复原始环境变量
			for key, value := range originalVars {
				if value == "" {
					os.Unsetenv(key)
				} else {
					os.Setenv(key, value)
				}
			}
		}()

		// 测试在清空环境变量后的行为
		for key := range originalVars {
			os.Unsetenv(key)
		}

		// 测试各个函数在受限环境下的行为
		functions := map[string]func() string{
			"UserHomeDir":    UserHomeDir,
			"UserConfigDir":  UserConfigDir,
			"UserCacheDir":   UserCacheDir,
			"LazyConfigDir":  LazyConfigDir,
			"LazyCacheDir":   LazyCacheDir,
		}

		for name, fn := range functions {
			result := fn()
			if result == "" {
				t.Logf("%s returned empty with cleared environment", name)
			} else {
				t.Logf("%s returned: %s (with cleared environment)", name, result)
			}
		}
	})
}

// 并发测试路径函数的线程安全性
func TestPathFunctionsConcurrency(t *testing.T) {
	t.Run("test_concurrent_path_function_calls", func(t *testing.T) {
		const numGoroutines = 10
		const callsPerGoroutine = 5

		functions := map[string]func() string{
			"ExecDir":        ExecDir,
			"ExecFile":       ExecFile,
			"Pwd":            Pwd,
			"UserHomeDir":    UserHomeDir,
			"UserConfigDir":  UserConfigDir,
			"UserCacheDir":   UserCacheDir,
			"LazyConfigDir":  LazyConfigDir,
			"LazyCacheDir":   LazyCacheDir,
		}

		results := make(chan map[string]string, numGoroutines)

		for i := 0; i < numGoroutines; i++ {
			go func(id int) {
				localResults := make(map[string]string)
				for name, fn := range functions {
					// 多次调用以检测竞争条件
					for j := 0; j < callsPerGoroutine; j++ {
						result := fn()
						if j == 0 {
							localResults[name] = result
						} else if localResults[name] != result {
							t.Errorf("Goroutine %d: %s returned inconsistent result: expected %s, got %s", id, name, localResults[name], result)
						}
					}
				}
				results <- localResults
			}(i)
		}

		// 收集所有结果并验证一致性
		var firstResult map[string]string
		for i := 0; i < numGoroutines; i++ {
			result := <-results
			if i == 0 {
				firstResult = result
			} else {
				for name, value := range result {
					if firstResult[name] != value {
						t.Errorf("Concurrent calls produced different results for %s: expected %s, got %s", name, firstResult[name], value)
					}
				}
			}
		}

		t.Log("All path functions passed concurrency test")
	})
}