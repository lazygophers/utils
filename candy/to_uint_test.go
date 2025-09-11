package candy

import (
	"testing"
)

// TestToUint 测试 ToUint 函数的各种类型转换
// 测试用例涵盖了所有支持的输入类型，包括边界情况和错误情况
func TestToUint(t *testing.T) {
	tests := []struct {
		name  string // 测试用例名称，描述具体的测试场景
		input interface{} // 输入参数，支持多种类型
		want  uint   // 期望的输出结果
	}{
		// 布尔值转换测试
		{"bool true", true, 1},                     // true 应该转换为 1
		{"bool false", false, 0},                   // false 应该转换为 0
		
		// 整数类型转换测试
		{"int positive", 42, 42},                   // 正整数直接转换
		{"int negative", -1, 18446744073709551615}, // 负整数转换为对应的无符号值
		{"uint", uint(100), 100},                   // uint 类型直接返回
		
		// 浮点数转换测试（截断小数部分）
		{"float positive", 3.14, 3},               // 正浮点数截断小数
		
		// 字符串转换测试
		{"string valid", "123", 123},               // 有效数字字符串
		{"string invalid", "abc", 0},              // 无效字符串返回 0
		
		// 字节切片转换测试
		{"byte slice valid", []byte("456"), 456},   // 有效数字字节切片
		{"byte slice invalid", []byte("xyz"), 0},  // 无效字节切片返回 0
		
		// 不支持的类型测试（返回 0）
		{"slice", []int{1, 2}, 0},                 // 切片类型不支持
		{"map", map[string]int{"a": 1}, 0},         // 映射类型不支持
		{"nil pointer", (*int)(nil), 0},           // nil 指针返回 0
		
		// 更多类型测试
		{"int8", int8(100), 100},
		{"int16", int16(1000), 1000}, 
		{"int32", int32(50000), 50000},
		{"int64", int64(70000), 70000},
		{"uint8", uint8(255), 255},
		{"uint16", uint16(65535), 65535},
		{"uint32", uint32(100000), 100000},
		{"uint64", uint64(200000), 200000},
		{"float32", float32(2.71), 2},
		{"float64", float64(2.71), 2},
		
		// 边界值测试
		{"max int", 1<<63 - 1, 9223372036854775807}, // 最大 int64 值
		{"min int", -1 << 63, 9223372036854775808}, // 最小 int64 值
		{"max uint", ^uint(0), ^uint(0)},          // 最大 uint 值
		
		// 更多字符串测试
		{"string empty", "", 0},
		{"string negative", "-10", 0}, // negative strings should fail parsing for uint
		{"string large", "18446744073709551615", 18446744073709551615},
		
		// 更多字节切片测试
		{"byte slice empty", []byte(""), 0},
		{"byte slice negative", []byte("-10"), 0},
		
		// nil 测试
		{"nil", nil, 0},
	}

	// 遍历所有测试用例
	for _, tt := range tests {
		// 使用子测试运行每个测试用例，便于定位问题
		t.Run(tt.name, func(t *testing.T) {
			// 调用 ToUint 函数进行转换
			if got := ToUint(tt.input); got != tt.want {
				// 如果结果不符合预期，输出详细的错误信息
				t.Errorf("ToUint() = %v, want %v", got, tt.want)
			}
		})
	}
}