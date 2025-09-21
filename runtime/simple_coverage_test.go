package runtime

import (
	"os"
	"path/filepath"
	"runtime/debug"
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestRuntimeCoverageImprovement tests runtime functions to improve coverage
func TestRuntimeCoverageImprovement(t *testing.T) {
	t.Run("PrintStackComplete", func(t *testing.T) {
		// Test PrintStack in various scenarios
		PrintStack()

		// Test with different stack sizes to cover both stack branches
		debug.SetMaxStack(1)  // Very small stack instead of 0
		PrintStack()

		debug.SetMaxStack(8192)
		PrintStack()

		debug.SetMaxStack(16384)
		PrintStack()

		// Test during panic recovery
		func() {
			defer func() {
				if r := recover(); r != nil {
					PrintStack() // This covers the panic scenario
				}
			}()
			panic("test for stack printing")
		}()
	})

	t.Run("CachePanicWithHandleComplete", func(t *testing.T) {
		// Test CachePanicWithHandle with nil handle - no panic case
		func() {
			defer CachePanicWithHandle(nil)
			// No panic, should complete normally
		}()

		// Test with panic and nil handle
		func() {
			defer CachePanicWithHandle(nil)
			panic("test panic with nil handle")
		}()

		// Test with panic and custom handle
		customHandleCalled := false
		func() {
			defer CachePanicWithHandle(func(err interface{}) {
				customHandleCalled = true
				assert.Equal(t, "test panic", err)
			})
			panic("test panic")
		}()
		assert.True(t, customHandleCalled)

		// Test with different panic types
		intPanicHandled := false
		func() {
			defer CachePanicWithHandle(func(err interface{}) {
				intPanicHandled = true
				assert.Equal(t, 42, err)
			})
			panic(42)
		}()
		assert.True(t, intPanicHandled)

		// Test with struct panic
		structPanicHandled := false
		testStruct := struct{ msg string }{"test message"}
		func() {
			defer CachePanicWithHandle(func(err interface{}) {
				structPanicHandled = true
				assert.Equal(t, testStruct, err)
			})
			panic(testStruct)
		}()
		assert.True(t, structPanicHandled)

		// Test with small stack scenario by setting stack to very small value
		func() {
			debug.SetMaxStack(1)
			defer func() {
				debug.SetMaxStack(8192) // Reset
			}()
			defer CachePanicWithHandle(func(err interface{}) {
				// This should trigger the small stack branch
			})
			panic("small stack test")
		}()
	})

	t.Run("DirectoryFunctionsBranches", func(t *testing.T) {
		// Test ExecDir multiple times to hit potential error paths
		execDir := ExecDir()
		assert.NotEmpty(t, execDir)
		assert.True(t, filepath.IsAbs(execDir))

		// Verify consistency
		assert.Equal(t, execDir, ExecDir())

		// Test ExecFile multiple times
		execFile := ExecFile()
		assert.NotEmpty(t, execFile)
		assert.True(t, filepath.IsAbs(execFile))

		// Verify consistency
		assert.Equal(t, execFile, ExecFile())

		// Test Pwd multiple times
		pwd := Pwd()
		assert.NotEmpty(t, pwd)
		assert.True(t, filepath.IsAbs(pwd))

		// Verify consistency
		assert.Equal(t, pwd, Pwd())

		// Test that directories exist
		_, err := os.Stat(execDir)
		assert.NoError(t, err)

		_, err = os.Stat(execFile)
		assert.NoError(t, err)

		_, err = os.Stat(pwd)
		assert.NoError(t, err)
	})

	t.Run("DirectoryFunctionsWithChdir", func(t *testing.T) {
		// Save current directory
		originalDir, err := os.Getwd()
		assert.NoError(t, err)

		// Restore at the end
		defer func() {
			os.Chdir(originalDir)
		}()

		// Change to temp directory to test Pwd() error handling
		tempDir := os.TempDir()
		err = os.Chdir(tempDir)
		assert.NoError(t, err)

		// Test functions after directory change
		pwd := Pwd()
		assert.NotEmpty(t, pwd)

		execDir := ExecDir()
		assert.NotEmpty(t, execDir)

		execFile := ExecFile()
		assert.NotEmpty(t, execFile)
	})

	t.Run("UserDirectoriesFull", func(t *testing.T) {
		// Test all user directory functions
		homeDir := UserHomeDir()
		assert.NotEmpty(t, homeDir)
		assert.True(t, filepath.IsAbs(homeDir))

		configDir := UserConfigDir()
		assert.NotEmpty(t, configDir)
		assert.True(t, filepath.IsAbs(configDir))

		cacheDir := UserCacheDir()
		assert.NotEmpty(t, cacheDir)
		assert.True(t, filepath.IsAbs(cacheDir))

		lazyConfigDir := LazyConfigDir()
		assert.NotEmpty(t, lazyConfigDir)
		assert.Contains(t, lazyConfigDir, "lazygophers")
		assert.True(t, filepath.IsAbs(lazyConfigDir))

		lazyCacheDir := LazyCacheDir()
		assert.NotEmpty(t, lazyCacheDir)
		assert.Contains(t, lazyCacheDir, "lazygophers")
		assert.True(t, filepath.IsAbs(lazyCacheDir))

		// Test multiple calls for consistency
		assert.Equal(t, homeDir, UserHomeDir())
		assert.Equal(t, configDir, UserConfigDir())
		assert.Equal(t, cacheDir, UserCacheDir())
		assert.Equal(t, lazyConfigDir, LazyConfigDir())
		assert.Equal(t, lazyCacheDir, LazyCacheDir())
	})

	t.Run("SystemDetectionFull", func(t *testing.T) {
		// Test system detection functions multiple times
		for i := 0; i < 10; i++ {
			windows := IsWindows()
			darwin := IsDarwin()
			linux := IsLinux()

			// Exactly one should be true
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

			assert.Equal(t, 1, trueCount, "Exactly one OS should be detected")
		}
	})

	t.Run("ExitSignFunctionsSafe", func(t *testing.T) {
		// Test GetExitSign - just verify we get a channel
		exitSign := GetExitSign()
		assert.NotNil(t, exitSign)

		// Test that the channel is the right type
		assert.IsType(t, make(chan os.Signal), exitSign)

		// We won't test WaitExit as it would block, but we can verify
		// the function exists and is callable
		// WaitExit() // This would block indefinitely
	})

	t.Run("CachePanicFunction", func(t *testing.T) {
		// Test that CachePanic doesn't interfere with normal execution
		// We'll test the panic case through CachePanicWithHandle
		func() {
			defer CachePanic()
			// Normal execution, no panic
		}()

		// Test CachePanic with a simple call (covers the wrapper)
		// Since CachePanic just calls CachePanicWithHandle(nil),
		// and we already test that thoroughly above, this is sufficient
	})
}

// TestEdgeCasesForBetterCoverage tests edge cases to improve coverage
func TestEdgeCasesForBetterCoverage(t *testing.T) {
	t.Run("MultipleStackSizes", func(t *testing.T) {
		// Test with various stack sizes to cover different code paths
		originalMaxStack := 0 // We can't easily get the original value

		stackSizes := []int{1024, 2048, 4096, 8192, 16384, 32768}
		for _, size := range stackSizes {
			debug.SetMaxStack(size)
			PrintStack()

			// Test with panic
			func() {
				defer func() {
					if r := recover(); r != nil {
						PrintStack()
					}
				}()
				panic("stack size test")
			}()
		}

		// Restore a reasonable stack size
		debug.SetMaxStack(8192)
		_ = originalMaxStack // Avoid unused variable
	})

	t.Run("RepeatedPanicHandling", func(t *testing.T) {
		// Test repeated panic handling to ensure robustness
		for i := 0; i < 20; i++ {
			func() {
				defer CachePanicWithHandle(func(err interface{}) {
					// Just handle the panic
				})
				panic(i)
			}()
		}
	})

	t.Run("EmptyStackTest", func(t *testing.T) {
		// Try to trigger empty stack scenario with small stack size instead of 0
		// Using 0 can cause unpredictable runtime behavior including hangs
		debug.SetMaxStack(1) // Very small stack size instead of 0
		func() {
			defer func() {
				debug.SetMaxStack(8192) // Ensure reset happens
			}()
			defer CachePanicWithHandle(func(err interface{}) {
				// This should trigger the small stack branch
			})
			panic("small stack attempt")
		}()

		// Also test PrintStack with small stack
		debug.SetMaxStack(1)
		PrintStack() // This should hit the small stack branch

		// Reset to reasonable size
		debug.SetMaxStack(8192)
	})

	t.Run("ErrorPathTesting", func(t *testing.T) {
		// These tests try to trigger error conditions to improve coverage
		// Most error paths are difficult to trigger in a real environment,
		// but we can try to hit edge cases

		// Test repeated calls to ensure stability
		for i := 0; i < 5; i++ {
			ExecDir()
			ExecFile()
			Pwd()
		}

		// Save current directory to test edge cases
		originalDir, _ := os.Getwd()
		defer os.Chdir(originalDir)

		// Try changing to root directory
		if err := os.Chdir("/"); err == nil {
			pwd := Pwd()
			assert.NotEmpty(t, pwd)

			execDir := ExecDir()
			assert.NotEmpty(t, execDir)

			execFile := ExecFile()
			assert.NotEmpty(t, execFile)
		}
	})

	t.Run("DifferentWorkingDirectories", func(t *testing.T) {
		// Test with different working directories
		original, _ := os.Getwd()
		defer os.Chdir(original)

		// Test with various directories
		testDirs := []string{
			os.TempDir(),
			"/",
		}

		for _, dir := range testDirs {
			if err := os.Chdir(dir); err == nil {
				pwd := Pwd()
				assert.NotEmpty(t, pwd)

				execDir := ExecDir()
				assert.NotEmpty(t, execDir)

				execFile := ExecFile()
				assert.NotEmpty(t, execFile)
			}
		}
	})
}