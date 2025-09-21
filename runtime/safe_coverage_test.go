package runtime

import (
	"testing"
)

// TestSystemFunctionsSafely 安全地测试系统检测函数
func TestSystemFunctionsSafely(t *testing.T) {
	t.Run("TestSystemDetectionFunctions", func(t *testing.T) {
		// 测试系统检测函数
		isWindows := IsWindows()
		isDarwin := IsDarwin()
		isLinux := IsLinux()

		t.Logf("IsWindows: %t", isWindows)
		t.Logf("IsDarwin: %t", isDarwin)
		t.Logf("IsLinux: %t", isLinux)

		// 验证有且仅有一个为true
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
			t.Errorf("应该有且仅有一个系统类型为true，但实际有 %d 个", count)
		}
	})
}