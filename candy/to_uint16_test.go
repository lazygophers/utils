// Package candy 提供语法糖工具函数，包含类型转换、数据处理等便捷功能
package candy

import (
	"testing"
)

// TestToUint16 测试 ToUint16 函数的各种输入转换场景
// 该函数验证将各种类型的输入转换为 uint16 的正确性
func TestToUint16(t *testing.T) {
	// 定义测试用例结构体，包含测试名称、输入值和期望输出
	tests := []struct {
		name  string    // 测试用例名称
		input interface{} // 输入值
		want  uint16    // 期望的输出结果
	}{
		// Bool values
		{"bool_true", true, 1},
		{"bool_false", false, 0},
		
		// Integer types
		{"int_positive", 50000, 50000},
		{"int_overflow", 70000, 4464},
		{"int8_value", int8(127), 127},
		{"int16_value", int16(32767), 32767},
		{"int32_value", int32(100), 100},
		{"int64_value", int64(1000), 1000},
		{"uint_value", uint(42), 42},
		{"uint8_value", uint8(255), 255},
		{"uint16_max", uint16(65535), 65535},
		{"uint32_value", uint32(100), 100},
		{"uint64_value", uint64(1000), 1000},
		
		// Float values
		{"float32_positive", float32(3.14), 3},
		{"float64_positive", float64(3.14), 3},
		{"float_negative", -100.5, 65436}, // 负浮点数的转换，测试补码处理
		
		// String values
		{"string_valid", "65535", 65535},
		{"string_zero", "0", 0},
		{"string_small", "42", 42},
		{"string_invalid", "invalid", 0},
		{"string_empty", "", 0},
		{"string_negative", "-42", 0},
		{"string_float", "3.14", 0},
		
		// Byte slice values
		{"byte_slice_valid", []byte("42"), 42},
		{"byte_slice_invalid", []byte("invalid"), 0},
		
		// Unsupported types
		{"nil_value", nil, 0},
		{"struct_value", struct{}{}, 0},
		{"map_value", map[string]int{}, 0},
	}

	// 遍历所有测试用例
	for _, tt := range tests {
		// 为每个测试用例创建子测试，便于定位问题
		t.Run(tt.name, func(t *testing.T) {
			// 调用 ToUint16 函数进行转换
			if got := ToUint16(tt.input); got != tt.want {
				// 如果结果与期望不符，输出错误信息
				t.Errorf("ToUint16() = %v, want %v", got, tt.want)
			}
		})
	}
}