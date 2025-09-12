package runtime

import (
	"runtime"
	"testing"
)

func TestIsWindows(t *testing.T) {
	t.Run("is_windows_consistency", func(t *testing.T) {
		result := IsWindows()
		expected := runtime.GOOS == "windows"
		
		if result != expected {
			t.Errorf("IsWindows() = %v, expected %v (runtime.GOOS = %s)", result, expected, runtime.GOOS)
		}
		
		t.Logf("IsWindows() = %v (runtime.GOOS = %s)", result, runtime.GOOS)
	})
}

func TestIsDarwin(t *testing.T) {
	t.Run("is_darwin_consistency", func(t *testing.T) {
		result := IsDarwin()
		expected := runtime.GOOS == "darwin"
		
		if result != expected {
			t.Errorf("IsDarwin() = %v, expected %v (runtime.GOOS = %s)", result, expected, runtime.GOOS)
		}
		
		t.Logf("IsDarwin() = %v (runtime.GOOS = %s)", result, runtime.GOOS)
	})
}

func TestIsLinux(t *testing.T) {
	t.Run("is_linux_consistency", func(t *testing.T) {
		result := IsLinux()
		expected := runtime.GOOS == "linux"
		
		if result != expected {
			t.Errorf("IsLinux() = %v, expected %v (runtime.GOOS = %s)", result, expected, runtime.GOOS)
		}
		
		t.Logf("IsLinux() = %v (runtime.GOOS = %s)", result, runtime.GOOS)
	})
}

// 测试系统检测函数的互斥性
func TestSystemDetectionMutualExclusion(t *testing.T) {
	t.Run("only_one_system_should_be_true", func(t *testing.T) {
		windows := IsWindows()
		darwin := IsDarwin()
		linux := IsLinux()
		
		trueCount := 0
		if windows {
			trueCount++
		}
		if darwin {
			trueCount++
		}
		if linux {
			trueCount++
		}
		
		if trueCount != 1 {
			t.Errorf("Exactly one system should be detected as true, but %d were true (Windows: %v, Darwin: %v, Linux: %v)",
				trueCount, windows, darwin, linux)
		}
		
		t.Logf("System detection: Windows=%v, Darwin=%v, Linux=%v", windows, darwin, linux)
	})
}

// 测试系统检测与runtime.GOOS的完全一致性
func TestSystemDetectionCompleteness(t *testing.T) {
	t.Run("system_detection_covers_current_goos", func(t *testing.T) {
		switch runtime.GOOS {
		case "windows":
			if !IsWindows() {
				t.Error("Should detect Windows when runtime.GOOS is 'windows'")
			}
			if IsDarwin() || IsLinux() {
				t.Error("Should not detect Darwin or Linux when running on Windows")
			}
			
		case "darwin":
			if !IsDarwin() {
				t.Error("Should detect Darwin when runtime.GOOS is 'darwin'")
			}
			if IsWindows() || IsLinux() {
				t.Error("Should not detect Windows or Linux when running on Darwin")
			}
			
		case "linux":
			if !IsLinux() {
				t.Error("Should detect Linux when runtime.GOOS is 'linux'")
			}
			if IsWindows() || IsDarwin() {
				t.Error("Should not detect Windows or Darwin when running on Linux")
			}
			
		default:
			// 对于其他系统（如freebsd, openbsd等），所有函数都应该返回false
			if IsWindows() || IsDarwin() || IsLinux() {
				t.Errorf("All system detection functions should return false for unsupported OS: %s", runtime.GOOS)
			}
			t.Logf("Unsupported OS detected: %s, all functions correctly returned false", runtime.GOOS)
		}
	})
}

// 性能测试
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

// 测试函数调用的稳定性
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

// 并发测试
func TestSystemDetectionConcurrency(t *testing.T) {
	t.Run("concurrent_calls_are_safe", func(t *testing.T) {
		const numGoroutines = 100
		const callsPerGoroutine = 100
		
		results := make(chan [3]bool, numGoroutines)
		
		for i := 0; i < numGoroutines; i++ {
			go func() {
				var localResults [3]bool
				for j := 0; j < callsPerGoroutine; j++ {
					localResults[0] = IsWindows()
					localResults[1] = IsDarwin()
					localResults[2] = IsLinux()
				}
				results <- localResults
			}()
		}
		
		// 收集所有结果
		var firstResult [3]bool
		var firstSet = false
		
		for i := 0; i < numGoroutines; i++ {
			result := <-results
			if !firstSet {
				firstResult = result
				firstSet = true
			} else {
				if result != firstResult {
					t.Errorf("Concurrent call returned different result: expected %v, got %v", firstResult, result)
				}
			}
		}
		
		t.Logf("All concurrent calls returned consistent results: %v", firstResult)
	})
}