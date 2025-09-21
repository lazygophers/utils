package runtime

import (
	"runtime/debug"
	"testing"
)

// TestMinimalCoverage tests functions to improve coverage to 98%
func TestMinimalCoverage(t *testing.T) {
	// Test CachePanic function (wrapper around CachePanicWithHandle)
	t.Run("CachePanic", func(t *testing.T) {
		// Test CachePanic without panic
		func() {
			defer CachePanic()
			// Normal execution
		}()

		// Test CachePanic with manual recovery to avoid test framework issues
		func() {
			defer func() {
				if r := recover(); r != nil {
					t.Logf("Manual recovery: %v", r)
				}
			}()
			defer CachePanic()
			panic("test for CachePanic")
		}()
	})

	// Test PrintStack with direct call to cover line 12-14
	t.Run("PrintStackDirect", func(t *testing.T) {
		// Save original max stack
		originalMax := debug.SetMaxStack(8192)
		defer debug.SetMaxStack(originalMax)

		// Call PrintStack directly
		PrintStack()

		// Test with small stack to cover small stack branch
		debug.SetMaxStack(1)
		PrintStack()
		debug.SetMaxStack(8192)
	})

	// Test runtime functions error paths
	t.Run("RuntimeErrorPaths", func(t *testing.T) {
		// Call each function multiple times to potentially hit error branches
		for i := 0; i < 5; i++ {
			ExecDir()
			ExecFile()
			Pwd()
			UserHomeDir()
			UserConfigDir()
			UserCacheDir()
			LazyConfigDir()
			LazyCacheDir()
		}
	})

	// Test CachePanicWithHandle error scenarios
	t.Run("CachePanicWithHandleError", func(t *testing.T) {
		// Test with small stack scenario
		func() {
			defer func() {
				if r := recover(); r != nil {
					t.Logf("Manual recovery in small stack test: %v", r)
				}
			}()

			// Set very small stack size
			debug.SetMaxStack(1)
			defer debug.SetMaxStack(8192)

			defer CachePanicWithHandle(func(err interface{}) {
				// This should hit the small stack branch in the function
			})
			panic("small stack panic test")
		}()
	})
}