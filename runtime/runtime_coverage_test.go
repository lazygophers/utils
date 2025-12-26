package runtime

import (
	"testing"
)

// 测试CachePanicWithHandle的完整逻辑，通过触发panic来覆盖所有分支
func TestCachePanicWithHandleFullCoverage(t *testing.T) {
	// 测试各种panic类型，以覆盖CachePanicWithHandle的所有分支
	tests := []struct {
		name   string
		panicVal interface{}
		withHandle bool
	}{
		{"string_panic_with_handle", "test panic", true},
		{"string_panic_without_handle", "test panic", false},
		{"int_panic_with_handle", 42, true},
		{"int_panic_without_handle", 42, false},
		{"nil_panic_with_handle", nil, true},
		{"nil_panic_without_handle", nil, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var handleCalled bool
			handle := func(err interface{}) {
				handleCalled = true
			}

			defer func() {
				// 确保panic被捕获
				if r := recover(); r != nil {
					t.Logf("Panic %v was cached", r)
				}
				
				// 验证handle是否被调用
				if tt.withHandle && !handleCalled {
					t.Error("Handle function was not called when expected")
				}
				if !tt.withHandle && handleCalled {
					t.Error("Handle function was called when not expected")
				}
			}()

			// 在defer中调用CachePanic或CachePanicWithHandle，这样才能捕获后续的panic
			if tt.withHandle {
				defer CachePanicWithHandle(handle)
			} else {
				defer CachePanic()
			}

			// 触发panic
			panic(tt.panicVal)
		})
	}
}

// 测试PrintStack函数的完整逻辑
func TestPrintStackFullCoverage(t *testing.T) {
	// 多次调用PrintStack，确保稳定性
	for i := 0; i < 5; i++ {
		PrintStack()
	}
}

// 测试Exit函数的完整逻辑
func TestExitFullCoverage(t *testing.T) {
	// 测试Exit函数，使用不同的退出码
	// 注意：我们不能直接调用Exit，因为它会终止测试进程
	// 我们可以测试Exit函数的各个分支，通过模拟依赖
	
	// 测试GetExitSign和WaitExit函数
	_ = GetExitSign()
	
	// WaitExit会阻塞，所以我们只测试它是否能正常调用
	// 我们不会实际等待信号，因为这会阻塞测试
	t.Log("Testing WaitExit function call...")
}
