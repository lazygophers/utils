// Package candy 提供语法糖工具函数，包含类型转换、数据处理等便捷功能
package candy

import (
	"testing"
)

// TestToUint64 测试 ToUint64 函数的各种输入转换场景
// 该函数验证将各种类型的输入转换为 uint64 的正确性，包括：
// - 布尔值转换
// - 整数转换（包括溢出情况）
// - 浮点数转换
// - 字符串解析
// - 边界值测试
// - 不支持类型的处理
func TestToUint64(t *testing.T) {
	// 定义测试用例，覆盖各种输入场景
	tests := []struct {
		name  string      // 测试用例名称
		input interface{} // 输入参数
		want  uint64      // 期望结果
	}{
		// 布尔值转换测试
		{"bool true", true, 1},   // true 转换为 1
		{"bool false", false, 0}, // false 转换为 0

		// 整数转换测试
		{"int positive", 123456, 123456},                           // 正常范围内的正整数
		{"int negative", -1, 18446744073709551615},                 // 负数转换为 uint64 的最大值
		{"int8", int8(127), 127},                                   // int8 类型转换
		{"int16", int16(32767), 32767},                             // int16 类型转换
		{"int32", int32(2147483647), 2147483647},                   // int32 类型转换
		{"int64", int64(9223372036854775807), 9223372036854775807}, // int64 类型转换
		{"uint", uint(100000), 100000},                             // uint 类型转换
		{"uint8", uint8(255), 255},                                 // uint8 类型转换
		{"uint16", uint16(65535), 65535},                           // uint16 类型转换
		{"uint32", uint32(4294967295), 4294967295},                 // uint32 类型转换

		// 浮点数转换测试
		{"float positive", 32767.5, 32767}, // 浮点数转换为 uint64，截断小数部分
		{"float negative", -100.5, 0},      // 负浮点数转换为 uint64 应该为 0

		// 字符串转换测试
		{"string valid", "18446744073709551615", 18446744073709551615}, // 有效的数字字符串
		{"string zero", "0", 0},      // 零值字符串
		{"string invalid", "abc", 0}, // 无效的字符串，返回 0

		// 字节切片转换测试
		{"byte slice valid", []byte("18446744073709551615"), 18446744073709551615}, // 有效的数字字节切片
		{"byte slice invalid", []byte("xyz"), 0},                                   // 无效的字节切片，返回 0

		// 边界值测试
		{"max uint64", uint64(18446744073709551615), 18446744073709551615}, // uint64 的最大值
		{"min uint64", uint64(0), 0},                                       // uint64 的最小值

		// 不支持的类型测试（返回 0）
		{"slice", []int{1, 2}, 0},          // 切片类型不支持
		{"map", map[string]int{"a": 1}, 0}, // 映射类型不支持
		{"nil pointer", (*int)(nil), 0},    // nil 指针返回 0
	}

	// 遍历所有测试用例
	for _, tt := range tests {
		// 使用子测试，每个测试用例独立运行
		t.Run(tt.name, func(t *testing.T) {
			// 调用 ToUint64 函数并验证结果
			if got := ToUint64(tt.input); got != tt.want {
				// 如果结果不符合预期，记录错误信息
				t.Errorf("ToUint64() = %v, want %v", got, tt.want)
			}
		})
	}
}

// TestToUint64_EdgeCases 测试 ToUint64 函数的边界情况和特殊值
// 该测试专门验证各种边界条件下的转换行为
func TestToUint64_EdgeCases(t *testing.T) {
	// 定义边界测试用例
	tests := []struct {
		name  string      // 测试用例名称
		input interface{} // 输入参数
		want  uint64      // 期望结果
	}{
		// 数值边界测试
		{"max int64", int64(9223372036854775807), 9223372036854775807},          // 最大 int64 值
		{"min int64", int64(-9223372036854775808), 9223372036854775808},         // 最小 int64 值的 uint64 表示
		{"max float64", float64(1.7976931348623157e+308), 18446744073709551615}, // 最大 float64 值应转换为最大 uint64 值
		{"min float64", float64(-1.7976931348623157e+308), 0},                   // 最小 float64 值转换为 uint64 应该为 0

		// 字符串边界测试
		{"string max value", "18446744073709551615", 18446744073709551615}, // 字符串最大值
		{"string min value", "0", 0},                                       // 字符串最小值
		{"string negative", "-1", 0},                                       // 字符串负数
		{"string with spaces", " 123456 ", 0},                              // 带空格的字符串
		{"string with newline", "123456\n", 0},                             // 带换行符的字符串

		// 字节切片边界测试
		{"byte slice empty", []byte(""), 0},                // 空字节切片
		{"byte slice with newline", []byte("123456\n"), 0}, // 带换行符的字节切片

		// 特殊值测试
		{"zero values", 0, 0},                              // 各种类型的零值
		{"negative zero", -0.0, 0},                         // 负零
		{"infinity", float64(1e100), 18446744073709551615}, // 无穷大应转换为最大 uint64 值
		{"negative infinity", float64(-1e100), 0},          // 负无穷大
	}

	// 遍历所有边界测试用例
	for _, tt := range tests {
		// 使用子测试，每个测试用例独立运行
		t.Run(tt.name, func(t *testing.T) {
			// 调用 ToUint64 函数并验证结果
			if got := ToUint64(tt.input); got != tt.want {
				// 如果结果不符合预期，记录错误信息
				t.Errorf("ToUint64() = %v, want %v", got, tt.want)
			}
		})
	}
}

// TestToUint64_TypeConsistency 测试 ToUint64 函数的类型一致性
// 该测试验证相同值的不同类型输入是否能产生一致的输出
func TestToUint64_TypeConsistency(t *testing.T) {
	// 定义类型一致性测试用例
	tests := []struct {
		name     string        // 测试用例名称
		inputs   []interface{} // 不同类型的输入值
		expected uint64        // 期望的统一输出
	}{
		{
			name: "value 100 across types",
			inputs: []interface{}{
				100,            // int
				int8(100),      // int8
				int16(100),     // int16
				int32(100),     // int32
				int64(100),     // int64
				uint(100),      // uint
				uint8(100),     // uint8
				uint16(100),    // uint16
				uint32(100),    // uint32
				uint64(100),    // uint64
				float32(100.0), // float32
				float64(100.0), // float64
				"100",          // string
				[]byte("100"),  // []byte
			},
			expected: 100,
		},
		{
			name: "value 0 across types",
			inputs: []interface{}{
				0,            // int
				false,        // bool
				int8(0),      // int8
				int16(0),     // int16
				int32(0),     // int32
				int64(0),     // int64
				uint(0),      // uint
				uint8(0),     // uint8
				uint16(0),    // uint16
				uint32(0),    // uint32
				uint64(0),    // uint64
				float32(0.0), // float32
				float64(0.0), // float64
				"0",          // string
				[]byte("0"),  // []byte
			},
			expected: 0,
		},
		{
			name: "value 1 across types",
			inputs: []interface{}{
				1,            // int
				true,         // bool
				int8(1),      // int8
				int16(1),     // int16
				int32(1),     // int32
				int64(1),     // int64
				uint(1),      // uint
				uint8(1),     // uint8
				uint16(1),    // uint16
				uint32(1),    // uint32
				uint64(1),    // uint64
				float32(1.0), // float32
				float64(1.0), // float64
				"1",          // string
				[]byte("1"),  // []byte
			},
			expected: 1,
		},
		{
			name: "value 127 across types",
			inputs: []interface{}{
				127,            // int
				int8(127),      // int8 (max value)
				int16(127),     // int16
				int32(127),     // int32
				int64(127),     // int64
				uint(127),      // uint
				uint8(127),     // uint8
				uint16(127),    // uint16
				uint32(127),    // uint32
				uint64(127),    // uint64
				float32(127.0), // float32
				float64(127.0), // float64
				"127",          // string
				[]byte("127"),  // []byte
			},
			expected: 127,
		},
	}

	// 遍历所有类型一致性测试用例
	for _, tt := range tests {
		// 使用子测试，每个测试用例独立运行
		t.Run(tt.name, func(t *testing.T) {
			// 遍历所有输入类型
			for _, input := range tt.inputs {
				// 调用 ToUint64 函数并验证结果
				if got := ToUint64(input); got != tt.expected {
					// 如果结果不符合预期，记录错误信息
					t.Errorf("ToUint64(%v) = %v, want %v", input, got, tt.expected)
				}
			}
		})
	}
}

// TestToUint64_OverflowBehavior 测试 ToUint64 函数的溢出行为
// 该测试专门验证数值溢出时的处理方式
func TestToUint64_OverflowBehavior(t *testing.T) {
	// 定义溢出测试用例
	tests := []struct {
		name  string      // 测试用例名称
		input interface{} // 输入参数
		want  uint64      // 期望结果
	}{
		// uint64 类型不会溢出，因为它是 64 位的最大无符号整数
		// 但我们需要测试其他类型转换为 uint64 时的行为
		{"int64 positive max", int64(9223372036854775807), 9223372036854775807},  // 最大 int64 值
		{"int64 negative max", int64(-9223372036854775808), 9223372036854775808}, // 最小 int64 值的 uint64 表示
		{"uint32 max", uint32(4294967295), 4294967295},                           // uint32 最大值
		{"uint16 max", uint16(65535), 65535},                                     // uint16 最大值
		{"uint8 max", uint8(255), 255},                                           // uint8 最大值

		// 浮点数测试
		{"float64 max", float64(1.7976931348623157e+308), 18446744073709551615}, // 最大 float64 值应转换为最大 uint64 值
		{"float64 negative max", float64(-1.7976931348623157e+308), 0},          // 最小 float64 值

		// 字符串溢出测试
		{"string overflow", "18446744073709551616", 0},                        // 超出 uint64 范围的字符串
		{"string underflow", "-1", 0},                                         // 字符串负数
		{"string large number", "999999999999999999999999999999999999999", 0}, // 极大数值字符串
	}

	// 遍历所有溢出测试用例
	for _, tt := range tests {
		// 使用子测试，每个测试用例独立运行
		t.Run(tt.name, func(t *testing.T) {
			// 调用 ToUint64 函数并验证结果
			if got := ToUint64(tt.input); got != tt.want {
				// 如果结果不符合预期，记录错误信息
				t.Errorf("ToUint64() = %v, want %v", got, tt.want)
			}
		})
	}
}

// TestToUint64_ErrorHandling 测试 ToUint64 函数的错误处理
// 该测试验证函数在面对无效输入时的错误处理能力
func TestToUint64_ErrorHandling(t *testing.T) {
	// 定义错误处理测试用例
	tests := []struct {
		name  string      // 测试用例名称
		input interface{} // 输入参数
		want  uint64      // 期望结果
	}{
		// 无效字符串测试
		{"string empty", "", 0},             // 空字符串
		{"string spaces", "   ", 0},         // 只有空格的字符串
		{"string mixed", "123abc", 0},       // 混合字符的字符串
		{"string scientific", "1.23e10", 0}, // 科学计数法字符串
		{"string hexadecimal", "0x123", 0},  // 十六进制字符串
		{"string octal", "0123", 123},       // 八进制字符串：strconv.ParseUint 在 base 10 下会将 "0123" 解析为 123

		// 无效字节切片测试
		{"byte slice empty", []byte{}, 0},         // 空字节切片
		{"byte slice spaces", []byte("   "), 0},   // 只有空格的字节切片
		{"byte slice mixed", []byte("123abc"), 0}, // 混合字符的字节切片

		// 复杂类型测试
		{"channel", make(chan int), 0},  // 通道类型
		{"function", func() {}, 0},      // 函数类型
		{"complex", complex(1, 2), 0},   // 复数类型
		{"array", [3]int{1, 2, 3}, 0},   // 数组类型
		{"pointer to int", new(int), 0}, // 指针类型

		// nil 值测试
		{"nil slice", ([]int)(nil), 0},        // nil 切片
		{"nil map", (map[string]int)(nil), 0}, // nil 映射
		{"nil channel", (chan int)(nil), 0},   // nil 通道
		{"nil function", (func())(nil), 0},    // nil 函数
	}

	// 遍历所有错误处理测试用例
	for _, tt := range tests {
		// 使用子测试，每个测试用例独立运行
		t.Run(tt.name, func(t *testing.T) {
			// 调用 ToUint64 函数并验证结果
			if got := ToUint64(tt.input); got != tt.want {
				// 如果结果不符合预期，记录错误信息
				t.Errorf("ToUint64() = %v, want %v", got, tt.want)
			}
		})
	}
}
