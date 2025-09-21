package runtime

import (
	"testing"
)

// TestAllRuntimeFunctions 测试所有Runtime函数的基本功能
func TestAllRuntimeFunctions(t *testing.T) {
	t.Run("TestAllUserDirFunctions", func(t *testing.T) {
		// 测试所有用户目录函数
		homeDir := UserHomeDir()
		configDir := UserConfigDir()
		cacheDir := UserCacheDir()
		lazyConfigDir := LazyConfigDir()
		lazyCacheDir := LazyCacheDir()

		t.Logf("UserHomeDir: %s", homeDir)
		t.Logf("UserConfigDir: %s", configDir)
		t.Logf("UserCacheDir: %s", cacheDir)
		t.Logf("LazyConfigDir: %s", lazyConfigDir)
		t.Logf("LazyCacheDir: %s", lazyCacheDir)

		// 基本验证：这些函数都应该能够执行
	})

	t.Run("TestSystemFunctions", func(t *testing.T) {
		// 测试系统检测函数
		isWindows := IsWindows()
		isDarwin := IsDarwin()
		isLinux := IsLinux()

		t.Logf("IsWindows: %t", isWindows)
		t.Logf("IsDarwin: %t", isDarwin)
		t.Logf("IsLinux: %t", isLinux)

		// 验证至少有一个为true
		if !isWindows && !isDarwin && !isLinux {
			t.Error("至少应该有一个系统类型为true")
		}
	})

	t.Run("TestFileSystemFunctions", func(t *testing.T) {
		// 测试文件系统相关函数
		execDir := ExecDir()
		execFile := ExecFile()
		pwd := Pwd()

		t.Logf("ExecDir: %s", execDir)
		t.Logf("ExecFile: %s", execFile)
		t.Logf("Pwd: %s", pwd)

		// 这些函数在正常情况下都应该返回有效路径
	})

	t.Run("TestPanicRecovery", func(t *testing.T) {
		// 测试panic恢复功能
		func() {
			defer CachePanic()
			panic("test panic for coverage")
		}()

		t.Log("CachePanic处理完成")
	})

	t.Run("TestPrintStack", func(t *testing.T) {
		// 测试堆栈打印
		PrintStack()
		t.Log("PrintStack执行完成")
	})
}

// TestPanicWithHandleScenarios 测试带处理器的panic场景
func TestPanicWithHandleScenarios(t *testing.T) {
	t.Run("PanicWithHandler", func(t *testing.T) {
		handled := false
		var capturedErr interface{}

		handle := func(err interface{}) {
			handled = true
			capturedErr = err
		}

		func() {
			defer CachePanicWithHandle(handle)
			panic("test error with handler")
		}()

		if !handled {
			t.Error("处理器应该被调用")
		}

		if capturedErr == nil {
			t.Error("应该捕获到错误")
		}

		t.Logf("成功捕获并处理panic: %v", capturedErr)
	})

	t.Run("PanicWithNilHandler", func(t *testing.T) {
		// 测试nil处理器的情况
		func() {
			defer CachePanicWithHandle(nil)
			panic("test error with nil handler")
		}()

		t.Log("nil处理器panic处理完成")
	})
}

// TestErrorConditions 测试错误条件
func TestErrorConditions(t *testing.T) {
	t.Run("CachePanicNormalCase", func(t *testing.T) {
		// 测试正常情况下的CachePanic（没有panic）
		defer CachePanic()
		// 没有panic，函数应该正常返回
		t.Log("正常情况下的CachePanic测试完成")
	})

	t.Run("CachePanicWithHandleNormalCase", func(t *testing.T) {
		// 测试正常情况下的CachePanicWithHandle
		handle := func(err interface{}) {
			t.Errorf("在正常情况下不应该调用处理器")
		}

		defer CachePanicWithHandle(handle)
		// 没有panic，处理器不应该被调用
		t.Log("正常情况下的CachePanicWithHandle测试完成")
	})
}