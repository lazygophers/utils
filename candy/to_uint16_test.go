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
		{"int positive", 50000, 50000},     // 正整数的正常转换
		{"int overflow", 70000, 4464},     // 超过 uint16 范围的整数，测试溢出处理
		{"float negative", -100.5, 65436}, // 负浮点数的转换，测试补码处理
		{"string valid", "65535", 65535}, // 有效字符串的转换
		{"max uint16", uint16(65535), 65535}, // uint16 最大值的转换
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