// Package candy 提供语法糖工具函数，增强 Go 开发的便利性和可读性
package candy

import (
	"testing"
)

// TestToUint8 测试 ToUint8 函数的功能和边界情况
// 该测试验证了各种输入类型转换为 uint8 的正确性，包括：
// - 布尔值转换
// - 整数转换（包括溢出情况）
// - 负数处理
// - 浮点数转换
// - 字符串解析
// - 边界值测试
func TestToUint8(t *testing.T) {
	// 定义测试用例，覆盖各种输入场景
	tests := []struct {
		name  string  // 测试用例名称
		input interface{} // 输入参数
		want  uint8   // 期望结果
	}{
		// 布尔值转换测试
		{"bool true", true, 1},           // true 转换为 1
		{"bool false", false, 0},         // false 转换为 0
		
		// 整数转换测试
		{"int positive", 200, 200},       // 正常范围内的正整数
		{"int overflow", 300, 44},       // 超出 uint8 范围的整数，测试溢出处理
		{"int negative", -1, 255},        // 负数转换为 uint8 的最大值
		
		// 浮点数转换测试
		{"float positive", 100.5, 100},   // 浮点数转换为 uint8，截断小数部分
		
		// 字符串转换测试
		{"string valid", "128", 128},     // 有效的数字字符串
		{"string invalid", "abc", 0},     // 无效的字符串，返回 0
		
		// 边界值测试
		{"max uint8", uint8(255), 255},   // uint8 的最大值
	}

	// 遍历所有测试用例
	for _, tt := range tests {
		// 使用子测试，每个测试用例独立运行
		t.Run(tt.name, func(t *testing.T) {
			// 调用 ToUint8 函数并验证结果
			if got := ToUint8(tt.input); got != tt.want {
				// 如果结果不符合预期，记录错误信息
				t.Errorf("ToUint8() = %v, want %v", got, tt.want)
			}
		})
	}
}