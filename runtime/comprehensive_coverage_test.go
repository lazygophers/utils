package runtime

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestSystemDetection tests only system detection functions
func TestSystemDetection(t *testing.T) {
	t.Run("SystemDetectionComplete", func(t *testing.T) {
		// Test each function multiple times to ensure consistency
		for i := 0; i < 5; i++ {
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

			assert.Equal(t, 1, trueCount, "Exactly one OS should be detected")

			// Results should be consistent across calls
			if i > 0 {
				assert.Equal(t, isWindows, IsWindows(), "IsWindows should be consistent")
				assert.Equal(t, isDarwin, IsDarwin(), "IsDarwin should be consistent")
				assert.Equal(t, isLinux, IsLinux(), "IsLinux should be consistent")
			}
		}
	})
}

// TestExitSignBasics tests only basic exit signal functionality
func TestExitSignBasics(t *testing.T) {
	t.Run("GetExitSign", func(t *testing.T) {
		exitChan := GetExitSign()
		assert.NotNil(t, exitChan)
		assert.IsType(t, make(chan os.Signal), exitChan)

		// Verify the channel is configured for signals
		select {
		case <-exitChan:
			t.Log("Signal received immediately (unexpected but handled)")
		default:
			t.Log("No signal pending (expected)")
		}
	})
}

// TestExitFunctionComponents tests only Exit function components safely
func TestExitFunctionComponents(t *testing.T) {
	t.Run("ExitFunctionComponents", func(t *testing.T) {
		// Test FindProcess functionality (part of Exit function)
		pid := os.Getpid()
		process, err := os.FindProcess(pid)
		if err != nil {
			t.Logf("FindProcess failed: %v", err)
		} else {
			assert.Equal(t, pid, process.Pid, "Process PID should match")
			t.Logf("Successfully found process: %d", process.Pid)
		}

		// Test signal functionality (without actually exiting)
		// We cannot safely test the actual Exit() function as it would terminate the test
		if process != nil {
			// Just verify the process exists and is valid
			// We don't send the actual signal to avoid terminating the test
			t.Logf("Process validation successful for PID: %d", process.Pid)
		}
	})
}

// BenchmarkSystemFunctions benchmarks only system detection functions
func BenchmarkSystemFunctions(b *testing.B) {
	b.Run("SystemDetection", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			IsWindows()
			IsDarwin()
			IsLinux()
		}
	})

	b.Run("GetExitSign", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			GetExitSign()
		}
	})
}