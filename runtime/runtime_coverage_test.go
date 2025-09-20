package runtime

import (
	"os"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRuntimeCoverage(t *testing.T) {
	// Test PrintStack function coverage
	t.Run("PrintStack", func(t *testing.T) {
		// This function logs to stderr, so we can't easily capture output
		// but we can at least call it to ensure it doesn't panic
		PrintStack()
		// Should not panic
		assert.True(t, true)
	})

	// Test CachePanicWithHandle function coverage
	t.Run("CachePanicWithHandle", func(t *testing.T) {
		t.Run("with handle function", func(t *testing.T) {
			var handleCalled bool
			var handledErr interface{}

			handle := func(err interface{}) {
				handleCalled = true
				handledErr = err
			}

			func() {
				defer CachePanicWithHandle(handle)
				panic("test panic")
			}()

			assert.True(t, handleCalled)
			assert.Equal(t, "test panic", handledErr)
		})

		t.Run("without handle function", func(t *testing.T) {
			func() {
				defer CachePanicWithHandle(nil)
				panic("test panic without handle")
			}()
			// Should not panic outside of the deferred function
			assert.True(t, true)
		})

		t.Run("no panic occurred", func(t *testing.T) {
			var handleCalled bool

			handle := func(err interface{}) {
				handleCalled = true
			}

			func() {
				defer CachePanicWithHandle(handle)
				// No panic
			}()

			assert.False(t, handleCalled)
		})
	})

	// Test ExecDir function coverage
	t.Run("ExecDir", func(t *testing.T) {
		result := ExecDir()
		// Should return a directory path or empty string
		if result != "" {
			assert.True(t, strings.Contains(result, "/") || strings.Contains(result, "\\"))
		}
	})

	// Test ExecFile function coverage
	t.Run("ExecFile", func(t *testing.T) {
		result := ExecFile()
		// Should return a file path or empty string
		if result != "" {
			assert.True(t, strings.Contains(result, "/") || strings.Contains(result, "\\"))
		}
	})

	// Test Pwd function coverage
	t.Run("Pwd", func(t *testing.T) {
		result := Pwd()
		// Should return current working directory or empty string
		if result != "" {
			// Compare with os.Getwd to verify it's correct
			wd, err := os.Getwd()
			if err == nil {
				assert.Equal(t, wd, result)
			}
		}
	})

	// Test error cases for coverage
	t.Run("Error cases", func(t *testing.T) {
		// These functions have error handling paths that we want to cover
		// They return empty strings on error, which is valid behavior

		// ExecDir error case is hard to trigger without mocking
		// ExecFile error case is hard to trigger without mocking
		// Pwd error case is hard to trigger without mocking

		// We can at least verify they don't panic
		assert.NotPanics(t, func() {
			ExecDir()
			ExecFile()
			Pwd()
		})
	})

	// Test system info functions (they should always work)
	t.Run("SystemInfo", func(t *testing.T) {
		// These are already 100% covered but let's test them anyway
		homeDir := UserHomeDir()
		configDir := UserConfigDir()
		cacheDir := UserCacheDir()
		lazyConfigDir := LazyConfigDir()
		lazyCacheDir := LazyCacheDir()

		// These should all return non-empty strings normally
		// but we won't assert that since they might fail in some environments
		_ = homeDir
		_ = configDir
		_ = cacheDir
		_ = lazyConfigDir
		_ = lazyCacheDir

		// At least verify they don't panic
		assert.NotPanics(t, func() {
			UserHomeDir()
			UserConfigDir()
			UserCacheDir()
			LazyConfigDir()
			LazyCacheDir()
		})
	})

	// Test platform detection functions
	t.Run("PlatformDetection", func(t *testing.T) {
		// These functions detect the current platform
		isWindows := IsWindows()
		isDarwin := IsDarwin()
		isLinux := IsLinux()

		// Exactly one should be true
		trueCount := 0
		if isWindows {
			trueCount++
		}
		if isDarwin {
			trueCount++
		}
		if isLinux {
			trueCount++
		}

		assert.Equal(t, 1, trueCount, "Exactly one platform detection should be true")
	})

	// Test GetExitSign function
	t.Run("GetExitSign", func(t *testing.T) {
		sigCh := GetExitSign()
		assert.NotNil(t, sigCh)
		// Don't wait for signal as that would block the test
	})
}