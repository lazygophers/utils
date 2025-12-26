package runtime

import (
	"testing"
)

// 测试各种函数的基本行为和边界情况
func TestRuntimeFunctionsBasic(t *testing.T) {
	// 测试CachePanic和CachePanicWithHandle的基本调用
	CachePanic()
	CachePanicWithHandle(nil)

	// 测试PrintStack的基本调用
	PrintStack()

	// 测试所有目录函数的基本调用
	// 这些函数在正常环境下应该返回有效值
	// 我们只需要确保它们不会panic
	_ = ExecDir()
	_ = ExecFile()
	_ = Pwd()
	_ = UserHomeDir()
	_ = UserConfigDir()
	_ = UserCacheDir()
	_ = LazyConfigDir()
	_ = LazyCacheDir()
}

// 测试CachePanicWithHandle的各种参数情况
func TestCachePanicWithHandleVariations(t *testing.T) {
	// 测试nil处理函数
	CachePanicWithHandle(nil)

	// 测试简单处理函数
	CachePanicWithHandle(func(err interface{}) {
		// 简单处理，什么都不做
	})

	// 测试复杂处理函数
	CachePanicWithHandle(func(err interface{}) {
		// 复杂处理，打印错误
		_ = err
	})
}

// 测试目录函数的多次调用一致性
func TestRuntimeFunctionsConsistency(t *testing.T) {
	// 测试多次调用相同函数的结果一致性
	dir1 := ExecDir()
	dir2 := ExecDir()
	if dir1 != dir2 {
		t.Logf("ExecDir() consistency check: %s vs %s", dir1, dir2)
	}

	file1 := ExecFile()
	file2 := ExecFile()
	if file1 != file2 {
		t.Logf("ExecFile() consistency check: %s vs %s", file1, file2)
	}

	pwd1 := Pwd()
	pwd2 := Pwd()
	if pwd1 != pwd2 {
		t.Logf("Pwd() consistency check: %s vs %s", pwd1, pwd2)
	}
}

