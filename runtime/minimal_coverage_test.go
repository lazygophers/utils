package runtime

import (
	"testing"
)

// TestMinimalSystemDetection provides minimal system detection testing
func TestMinimalSystemDetection(t *testing.T) {
	t.Run("BasicSystemDetection", func(t *testing.T) {
		// Test system detection functions only
		isWindows := IsWindows()
		isDarwin := IsDarwin()
		isLinux := IsLinux()

		// Verify exactly one is true
		count := 0
		if isWindows {
			count++
		}
		if isDarwin {
			count++
		}
		if isLinux {
			count++
		}

		if count != 1 {
			t.Errorf("Expected exactly one OS to be detected, got %d", count)
		}

		t.Logf("System detection: Windows=%t, Darwin=%t, Linux=%t", isWindows, isDarwin, isLinux)
	})
}