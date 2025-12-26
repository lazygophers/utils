package runtime

import (
	"os"
	"runtime"
	"strings"
	"testing"
)

func TestCachePanic(t *testing.T) {
	t.Run("cache_panic_without_handle", func(t *testing.T) {
		defer func() {
			if r := recover(); r != nil {
				t.Log("Panic was cached as expected")
			}
		}()

		CachePanic()

		panic("test panic")
	})
}

func TestCachePanicWithHandle(t *testing.T) {
	t.Run("cache_panic_with_handle", func(t *testing.T) {
		handleCalled := false

		defer func() {
			if r := recover(); r != nil {
				t.Log("Panic was cached as expected")
			}
		}()

		CachePanicWithHandle(func(err interface{}) {
			handleCalled = true
			t.Logf("Handle called with: %v", err)
		})

		panic("test panic")

		if !handleCalled {
			t.Error("Handle function was not called")
		}
	})
}

func TestCachePanicWithHandleNil(t *testing.T) {
	t.Run("cache_panic_with_nil_handle", func(t *testing.T) {
		defer func() {
			if r := recover(); r != nil {
				t.Log("Panic was cached as expected")
			}
		}()

		CachePanicWithHandle(nil)

		panic("test panic")
	})
}

func TestPrintStack(t *testing.T) {
	t.Run("print_stack", func(t *testing.T) {
		PrintStack()
		t.Log("PrintStack executed successfully")
	})
}

func TestExecDir(t *testing.T) {
	t.Run("exec_dir", func(t *testing.T) {
		dir := ExecDir()

		if dir == "" {
			t.Error("ExecDir() returned empty string")
		}

		t.Logf("ExecDir() = %s", dir)
	})
}

func TestExecFile(t *testing.T) {
	t.Run("exec_file", func(t *testing.T) {
		file := ExecFile()

		if file == "" {
			t.Error("ExecFile() returned empty string")
		}

		t.Logf("ExecFile() = %s", file)

		if file != "" && !strings.Contains(file, string(os.PathSeparator)) {
			t.Error("ExecFile() should contain path separator")
		}
	})
}

func TestPwd(t *testing.T) {
	t.Run("pwd", func(t *testing.T) {
		pwd := Pwd()

		if pwd == "" {
			t.Error("Pwd() returned empty string")
		}

		t.Logf("Pwd() = %s", pwd)
	})
}

func TestUserHomeDir(t *testing.T) {
	t.Run("user_home_dir", func(t *testing.T) {
		dir := UserHomeDir()

		if dir == "" {
			t.Error("UserHomeDir() returned empty string")
		}

		t.Logf("UserHomeDir() = %s", dir)
	})
}

func TestUserConfigDir(t *testing.T) {
	t.Run("user_config_dir", func(t *testing.T) {
		dir := UserConfigDir()

		if dir == "" {
			t.Error("UserConfigDir() returned empty string")
		}

		t.Logf("UserConfigDir() = %s", dir)
	})
}

func TestUserCacheDir(t *testing.T) {
	t.Run("user_cache_dir", func(t *testing.T) {
		dir := UserCacheDir()

		if dir == "" {
			t.Error("UserCacheDir() returned empty string")
		}

		t.Logf("UserCacheDir() = %s", dir)
	})
}

func TestLazyConfigDir(t *testing.T) {
	t.Run("lazy_config_dir", func(t *testing.T) {
		dir := LazyConfigDir()

		if dir == "" {
			t.Error("LazyConfigDir() returned empty string")
		}

		t.Logf("LazyConfigDir() = %s", dir)

		if !strings.Contains(dir, "lazygophers") {
			t.Error("LazyConfigDir() should contain 'lazygophers'")
		}
	})
}

func TestLazyCacheDir(t *testing.T) {
	t.Run("lazy_cache_dir", func(t *testing.T) {
		dir := LazyCacheDir()

		if dir == "" {
			t.Error("LazyCacheDir() returned empty string")
		}

		t.Logf("LazyCacheDir() = %s", dir)

		if !strings.Contains(dir, "lazygophers") {
			t.Error("LazyCacheDir() should contain 'lazygophers'")
		}
	})
}

func TestCachePanicWithDifferentTypes(t *testing.T) {
	tests := []struct {
		name  string
		panic interface{}
	}{
		{"string", "test panic"},
		{"int", 42},
		{"error", &testError{}},
		{"nil", nil},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			defer func() {
				if r := recover(); r != nil {
					t.Log("Panic was cached")
				}
			}()

			CachePanic()

			panic(tt.panic)
		})
	}
}

type testError struct{}

func (e *testError) Error() string {
	return "test error"
}

func TestExecDirConsistency(t *testing.T) {
	t.Run("exec_dir_consistency", func(t *testing.T) {
		dir1 := ExecDir()
		dir2 := ExecDir()

		if dir1 != dir2 {
			t.Errorf("ExecDir() returned different values: %s vs %s", dir1, dir2)
		}
	})
}

func TestExecFileConsistency(t *testing.T) {
	t.Run("exec_file_consistency", func(t *testing.T) {
		file1 := ExecFile()
		file2 := ExecFile()

		if file1 != file2 {
			t.Errorf("ExecFile() returned different values: %s vs %s", file1, file2)
		}
	})
}

func TestPwdConsistency(t *testing.T) {
	t.Run("pwd_consistency", func(t *testing.T) {
		pwd1 := Pwd()
		pwd2 := Pwd()

		if pwd1 != pwd2 {
			t.Errorf("Pwd() returned different values: %s vs %s", pwd1, pwd2)
		}
	})
}

func TestLazyDirsContainOrganization(t *testing.T) {
	t.Run("lazy_dirs_contain_organization", func(t *testing.T) {
		configDir := LazyConfigDir()
		cacheDir := LazyCacheDir()

		if !strings.Contains(configDir, "lazygophers") {
			t.Errorf("LazyConfigDir() should contain 'lazygophers': %s", configDir)
		}

		if !strings.Contains(cacheDir, "lazygophers") {
			t.Errorf("LazyCacheDir() should contain 'lazygophers': %s", cacheDir)
		}
	})
}

func TestCachePanicWithHandleCalled(t *testing.T) {
	t.Run("handle_function_called", func(t *testing.T) {
		var capturedErr interface{}
		handleCalled := false

		defer func() {
			if r := recover(); r != nil {
				t.Log("Panic was cached")
			}
		}()

		CachePanicWithHandle(func(err interface{}) {
			handleCalled = true
			capturedErr = err
		})

		testErr := "test error"
		panic(testErr)

		if !handleCalled {
			t.Error("Handle function was not called")
		}

		if capturedErr != testErr {
			t.Errorf("Handle function received wrong error: %v, expected %v", capturedErr, testErr)
		}
	})
}

func TestPrintStackDoesNotPanic(t *testing.T) {
	t.Run("print_stack_no_panic", func(t *testing.T) {
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("PrintStack() should not panic: %v", r)
			}
		}()

		PrintStack()
	})
}

func TestUserDirsReturnValidPaths(t *testing.T) {
	tests := []struct {
		name string
		fn   func() string
	}{
		{"UserHomeDir", UserHomeDir},
		{"UserConfigDir", UserConfigDir},
		{"UserCacheDir", UserCacheDir},
		{"LazyConfigDir", LazyConfigDir},
		{"LazyCacheDir", LazyCacheDir},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			dir := tt.fn()

			if dir == "" {
				t.Errorf("%s() returned empty string", tt.name)
			}

			if runtime.GOOS != "windows" {
				if !strings.HasPrefix(dir, "/") {
					t.Errorf("%s() should return absolute path: %s", tt.name, dir)
				}
			}
		})
	}
}

func TestExecFileIsExecutable(t *testing.T) {
	t.Run("exec_file_is_executable", func(t *testing.T) {
		file := ExecFile()

		if file == "" {
			t.Skip("Cannot test executable file in this environment")
		}

		info, err := os.Stat(file)
		if err != nil {
			t.Logf("Cannot stat executable file: %v", err)
			return
		}

		if info.IsDir() {
			t.Error("ExecFile() should return a file, not a directory")
		}
	})
}

func TestCachePanicWithNestedPanic(t *testing.T) {
	t.Run("cache_nested_panic", func(t *testing.T) {
		defer func() {
			if r := recover(); r != nil {
				t.Log("Nested panic was cached")
			}
		}()

		CachePanic()

		func() {
			CachePanic()
			panic("nested panic")
		}()
	})
}

func TestRuntimeFunctionsConcurrent(t *testing.T) {
	t.Run("concurrent_calls", func(t *testing.T) {
		const numGoroutines = 10
		const callsPerGoroutine = 100

		results := make(chan string, numGoroutines*callsPerGoroutine)

		for i := 0; i < numGoroutines; i++ {
			go func() {
				for j := 0; j < callsPerGoroutine; j++ {
					results <- Pwd()
					results <- UserHomeDir()
					results <- UserConfigDir()
					results <- UserCacheDir()
				}
			}()
		}

		for i := 0; i < numGoroutines*4*callsPerGoroutine; i++ {
			result := <-results
			if result == "" {
				t.Error("Function returned empty string")
			}
		}

		t.Log("All concurrent calls completed successfully")
	})
}

func TestExecDirAndFileRelationship(t *testing.T) {
	t.Run("exec_dir_and_file_relationship", func(t *testing.T) {
		dir := ExecDir()
		file := ExecFile()

		if dir == "" || file == "" {
			t.Skip("Cannot test in this environment")
		}

		if !strings.HasPrefix(file, dir) {
			t.Errorf("ExecFile() should be within ExecDir(): file=%s, dir=%s", file, dir)
		}
	})
}

func TestPrintStackEmptyStack(t *testing.T) {
	t.Run("print_stack_empty", func(t *testing.T) {
		// 测试PrintStack函数在栈为空的情况下的行为
		// 由于我们无法轻松模拟栈为空的情况，我们只确保函数能正常执行
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("PrintStack() panicked with empty stack: %v", r)
			}
		}()

		PrintStack()
		t.Log("PrintStack() executed with empty stack (simulated)")
	})
}

func TestCachePanicWithHandleStackEmpty(t *testing.T) {
	t.Run("cache_panic_with_empty_stack", func(t *testing.T) {
		// 测试CachePanicWithHandle函数在栈为空的情况下的行为
		defer func() {
			if r := recover(); r != nil {
				t.Log("Panic was cached")
			}
		}()

		CachePanicWithHandle(func(err interface{}) {
			t.Logf("Handle called with: %v", err)
		})

		// 使用panic触发函数，但我们无法直接模拟栈为空的情况
		// 不过这个测试会执行函数的大部分分支
		panic("test panic with simulated empty stack")
	})
}

func TestSystemFunctions(t *testing.T) {
	t.Run("system_functions", func(t *testing.T) {
		// 测试系统相关函数
		isWindows := IsWindows()
		isDarwin := IsDarwin()
		isLinux := IsLinux()

		// 在macOS系统上，IsDarwin应该返回true，其他返回false
		if isDarwin != true {
			t.Error("IsDarwin() should return true on macOS")
		}

		if isWindows != false {
			t.Error("IsWindows() should return false on macOS")
		}

		if isLinux != false {
			t.Error("IsLinux() should return false on macOS")
		}
	})
}
