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
		name  string      // 测试用例名称
		input interface{} // 输入参数
		want  uint8       // 期望结果
	}{
		// 布尔值转换测试
		{"bool true", true, 1},   // true 转换为 1
		{"bool false", false, 0}, // false 转换为 0

		// 整数转换测试
		{"int positive", 200, 200}, // 正常范围内的正整数
		{"int overflow", 300, 44},  // 超出 uint8 范围的整数，测试溢出处理
		{"int negative", -1, 255},  // 负数转换为 uint8 的最大值

		// 浮点数转换测试
		{"float positive", 100.5, 100}, // 浮点数转换为 uint8，截断小数部分

		// 字符串转换测试
		{"string valid", "128", 128}, // 有效的数字字符串
		{"string invalid", "abc", 0}, // 无效的字符串，返回 0

		// 更多整数类型测试
		{"int8 positive", int8(100), 100},
		{"int8 negative", int8(-10), 246}, // -10 -> 256-10 = 246
		{"int16", int16(200), 200},
		{"int32", int32(150), 150},
		{"int64", int64(180), 180},
		{"uint", uint(220), 220},
		{"uint16", uint16(300), 44},   // 300 & 0xFF = 44
		{"uint32", uint32(500), 244},  // 500 & 0xFF = 244
		{"uint64", uint64(1000), 232}, // 1000 & 0xFF = 232

		// 更多浮点数测试
		{"float32", float32(123.9), 123},
		{"float64", float64(200.7), 200},
		{"float negative", float64(-5.5), 251}, // -5 -> 256-5 = 251

		// 更多字符串测试
		{"string zero", "0", 0},
		{"string max", "255", 255},
		{"string overflow", "256", 0}, // should fail parsing for uint8
		{"string empty", "", 0},
		{"string negative", "-1", 0}, // negative should fail

		// 字节切片测试
		{"byte slice valid", []byte("100"), 100},
		{"byte slice invalid", []byte("xyz"), 0},
		{"byte slice empty", []byte(""), 0},
		{"byte slice overflow", []byte("300"), 0},

		// 不支持的类型
		{"nil", nil, 0},
		{"struct", struct{}{}, 0},
		{"slice", []int{1, 2, 3}, 0},

		// 边界值测试
		{"max uint8", uint8(255), 255}, // uint8 的最大值
		{"min uint8", uint8(0), 0},     // uint8 的最小值
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
