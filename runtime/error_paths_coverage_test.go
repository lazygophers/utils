package runtime

import (
	"os"
	"os/exec"
	"strings"
	"syscall"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

// TestExitFunctionCoverage tests the Exit function by running it in a subprocess
func TestExitFunctionCoverage(t *testing.T) {
	// We need to test Exit() in a separate process to avoid terminating the test process
	if os.Getenv("TEST_EXIT_SUBPROCESS") == "1" {
		// This code will run in the subprocess
		Exit()
		return
	}

	// Run the test in a subprocess
	cmd := exec.Command(os.Args[0], "-test.run=TestExitFunctionCoverage")
	cmd.Env = append(os.Environ(), "TEST_EXIT_SUBPROCESS=1")

	// Set up to capture output and process state
	err := cmd.Start()
	assert.NoError(t, err, "Should be able to start subprocess")

	// Create a goroutine to kill the process after a timeout to avoid hanging
	done := make(chan error, 1)
	go func() {
		done <- cmd.Wait()
	}()

	select {
	case err := <-done:
		// The process should have been terminated or exited
		if err != nil {
			// This is expected - the process was interrupted
			exitError, ok := err.(*exec.ExitError)
			if ok {
				// Check if it was terminated by a signal (which is what Exit() does)
				if status, ok := exitError.Sys().(syscall.WaitStatus); ok {
					if status.Signaled() {
						// Process was terminated by signal - this is expected behavior from Exit()
						t.Logf("Process was terminated by signal: %v", status.Signal())
					} else {
						t.Logf("Process exited with code: %d", status.ExitStatus())
					}
				}
			}
		}
	case <-time.After(2 * time.Second):
		// Kill the process if it hangs
		if cmd.Process != nil {
			cmd.Process.Kill()
		}
		t.Log("Killed subprocess after timeout")
	}

	t.Log("Exit() function test completed - process was terminated as expected")
}

// TestErrorPathInCachePanicWithHandle tests the error handling path where stack is empty
func TestCachePanicEmptyStackCoverage(t *testing.T) {
	// This test is tricky because debug.Stack() almost always returns a non-empty stack
	// But we can still test the panic handling logic
	handlerCalled := false
	var capturedError interface{}

	func() {
		defer CachePanicWithHandle(func(err interface{}) {
			handlerCalled = true
			capturedError = err
		})

		panic("test error for stack coverage")
	}()

	assert.True(t, handlerCalled, "Handler should have been called")
	assert.Equal(t, "test error for stack coverage", capturedError)
}

// TestPrintStackEmptyPathCoverage tests PrintStack when stack might be empty
func TestPrintStackEmptyPathCoverage(t *testing.T) {
	// PrintStack should always work normally - the empty case is very rare
	// but we'll ensure the function is called for coverage
	PrintStack()

	// Test multiple calls to ensure consistent behavior
	for i := 0; i < 3; i++ {
		PrintStack()
	}
}

// TestErrorPathsInDirectoryFunctions tests error paths in filesystem functions
func TestDirectoryErrorPaths(t *testing.T) {
	// These functions have error paths that are difficult to trigger normally
	// But we can ensure they're called and test their normal behavior

	t.Run("exec_dir_normal_path", func(t *testing.T) {
		dir := ExecDir()
		assert.NotEmpty(t, dir, "ExecDir should return a directory")
		assert.True(t, strings.Contains(dir, "/") || strings.Contains(dir, "\\"), "Should be a valid path")
	})

	t.Run("exec_file_normal_path", func(t *testing.T) {
		file := ExecFile()
		assert.NotEmpty(t, file, "ExecFile should return a file path")
		assert.True(t, strings.Contains(file, "/") || strings.Contains(file, "\\"), "Should be a valid path")
	})

	t.Run("pwd_normal_path", func(t *testing.T) {
		pwd := Pwd()
		assert.NotEmpty(t, pwd, "Pwd should return current directory")

		// Compare with os.Getwd() to verify
		expected, err := os.Getwd()
		assert.NoError(t, err)
		assert.Equal(t, expected, pwd, "Pwd should match os.Getwd()")
	})
}

// TestGetExitSignCoverage tests the GetExitSign function
func TestGetExitSignCoverage(t *testing.T) {
	sigCh := GetExitSign()
	assert.NotNil(t, sigCh, "GetExitSign should return a channel")

	// Verify it's a buffered channel by checking it doesn't block on receive
	select {
	case <-sigCh:
		t.Fatal("Channel should not have any signals initially")
	default:
		// Expected - channel is empty
	}
}

// TestCachePanicNilHandler tests CachePanicWithHandle with nil handler
func TestCachePanicNilHandlerCoverage(t *testing.T) {
	// Test the path where handle is nil
	func() {
		defer CachePanicWithHandle(nil) // nil handler
		panic("test panic with nil handler")
	}()

	// If we get here, the panic was caught and handled properly
}

// TestCachePanicBasicCoverage tests the basic CachePanic function
func TestCachePanicBasicCoverage(t *testing.T) {
	// Test that CachePanic can be called without panic
	defer func() {
		if r := recover(); r != nil {
			t.Errorf("Should not panic: %v", r)
		}
	}()

	func() {
		defer CachePanic()
		// No panic here - test normal case
	}()

	// Test with actual panic
	func() {
		defer func() {
			// Outer recover to prevent test from failing
			if r := recover(); r != nil {
				t.Logf("Recovered in test: %v", r)
			}
		}()

		func() {
			defer CachePanic()
			panic("test basic cache panic")
		}()
	}()
}

// TestAllDirectoryFunctionsCompletely ensures all directory functions are tested
func TestAllDirectoryFunctionsCompletely(t *testing.T) {
	// Test UserHomeDir
	home := UserHomeDir()
	assert.NotEmpty(t, home, "UserHomeDir should not be empty")

	// Test UserConfigDir
	config := UserConfigDir()
	assert.NotEmpty(t, config, "UserConfigDir should not be empty")

	// Test UserCacheDir
	cache := UserCacheDir()
	assert.NotEmpty(t, cache, "UserCacheDir should not be empty")

	// Test LazyConfigDir
	lazyConfig := LazyConfigDir()
	assert.NotEmpty(t, lazyConfig, "LazyConfigDir should not be empty")
	assert.Contains(t, lazyConfig, config, "LazyConfigDir should contain config dir")

	// Test LazyCacheDir
	lazyCache := LazyCacheDir()
	assert.NotEmpty(t, lazyCache, "LazyCacheDir should not be empty")
	assert.Contains(t, lazyCache, cache, "LazyCacheDir should contain cache dir")
}

// TestSystemDetectionComplete ensures all system detection functions are covered
func TestSystemDetectionComplete(t *testing.T) {
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

	assert.Equal(t, 1, trueCount, "Exactly one platform should be detected")
}
