// Package candy 提供语法糖工具函数，包含类型转换、数据处理等便捷功能
package candy

import (
	"testing"
)

// TestToUint32 测试 ToUint32 函数的各种输入转换场景
// 该函数验证将各种类型的输入转换为 uint32 的正确性，包括：
// - 布尔值转换
// - 整数转换（包括溢出情况）
// - 浮点数转换
// - 字符串解析
// - 边界值测试
// - 不支持类型的处理
func TestToUint32(t *testing.T) {
	// 定义测试用例，覆盖各种输入场景
	tests := []struct {
		name  string      // 测试用例名称
		input interface{} // 输入参数
		want  uint32      // 期望结果
	}{
		// 布尔值转换测试
		{"bool true", true, 1},   // true 转换为 1
		{"bool false", false, 0}, // false 转换为 0

		// 整数转换测试
		{"int positive", 123456, 123456},         // 正常范围内的正整数
		{"int negative", -1, 4294967295},         // 负数转换为 uint32 的最大值
		{"int8", int8(127), 127},                 // int8 类型转换
		{"int16", int16(32767), 32767},           // int16 类型转换
		{"int32", int32(2147483647), 2147483647}, // int32 类型转换
		{"int64 overflow", int64(4294967296), 0}, // 超出 uint32 范围的 int64
		{"uint", uint(100000), 100000},           // uint 类型转换
		{"uint8", uint8(255), 255},               // uint8 类型转换
		{"uint16", uint16(65535), 65535},         // uint16 类型转换

		// 浮点数转换测试
		{"float positive", 32767.5, 32767},     // 浮点数转换为 uint32，截断小数部分
		{"float negative", -100.5, 4294967196}, // 负浮点数的转换

		// 字符串转换测试
		{"string valid", "4294967295", 4294967295}, // 有效的数字字符串
		{"string zero", "0", 0},                    // 零值字符串
		{"string invalid", "abc", 0},               // 无效的字符串，返回 0
		{"string overflow", "4294967296", 0},       // 超出范围的字符串

		// 字节切片转换测试
		{"byte slice valid", []byte("123456"), 123456}, // 有效的数字字节切片
		{"byte slice invalid", []byte("xyz"), 0},       // 无效的字节切片，返回 0

		// 边界值测试
		{"max uint32", uint32(4294967295), 4294967295}, // uint32 的最大值
		{"min uint32", uint32(0), 0},                   // uint32 的最小值

		// 不支持的类型测试（返回 0）
		{"slice", []int{1, 2}, 0},          // 切片类型不支持
		{"map", map[string]int{"a": 1}, 0}, // 映射类型不支持
		{"nil pointer", (*int)(nil), 0},    // nil 指针返回 0
		{"struct", struct{}{}, 0},          // 结构体类型不支持
	}

	// 遍历所有测试用例
	for _, tt := range tests {
		// 使用子测试，每个测试用例独立运行
		t.Run(tt.name, func(t *testing.T) {
			// 调用 ToUint32 函数并验证结果
			if got := ToUint32(tt.input); got != tt.want {
				// 如果结果不符合预期，记录错误信息
				t.Errorf("ToUint32() = %v, want %v", got, tt.want)
			}
		})
	}
}

// TestToUint32_EdgeCases 测试 ToUint32 函数的边界情况和特殊值
// 该测试专门验证各种边界条件下的转换行为
func TestToUint32_EdgeCases(t *testing.T) {
	// 定义边界测试用例
	tests := []struct {
		name  string      // 测试用例名称
		input interface{} // 输入参数
		want  uint32      // 期望结果
	}{
		// 数值边界测试
		{"max int64 in uint32 range", int64(4294967295), 4294967295}, // 最大可表示的 int64 值
		{"min int64 in uint32 range", int64(-4294967296), 0},         // 最小可表示的 int64 值
		{"max float64", float64(4294967295.9), 4294967295},           // 最大浮点数
		{"min float64", float64(-4294967296.1), 0},                   // 最小浮点数

		// 字符串边界测试
		{"string max value", "4294967295", 4294967295}, // 字符串最大值
		{"string min value", "0", 0},                   // 字符串最小值
		{"string negative", "-1", 0},                   // 字符串负数
		{"string with spaces", " 123456 ", 0},          // 带空格的字符串
		{"string with newline", "123456\n", 0},         // 带换行符的字符串

		// 字节切片边界测试
		{"byte slice empty", []byte(""), 0},                // 空字节切片
		{"byte slice with newline", []byte("123456\n"), 0}, // 带换行符的字节切片

		// 特殊值测试
		{"zero values", 0, 0},                     // 各种类型的零值
		{"negative zero", -0.0, 0},                // 负零
		{"infinity", float64(1e100), 4294967295},  // 无穷大
		{"negative infinity", float64(-1e100), 0}, // 负无穷大
	}

	// 遍历所有边界测试用例
	for _, tt := range tests {
		// 使用子测试，每个测试用例独立运行
		t.Run(tt.name, func(t *testing.T) {
			// 调用 ToUint32 函数并验证结果
			if got := ToUint32(tt.input); got != tt.want {
				// 如果结果不符合预期，记录错误信息
				t.Errorf("ToUint32() = %v, want %v", got, tt.want)
			}
		})
	}
}

// TestToUint32_TypeConsistency 测试 ToUint32 函数的类型一致性
// 该测试验证相同值的不同类型输入是否能产生一致的输出
func TestToUint32_TypeConsistency(t *testing.T) {
	// 定义类型一致性测试用例
	tests := []struct {
		name     string        // 测试用例名称
		inputs   []interface{} // 不同类型的输入值
		expected uint32        // 期望的统一输出
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
	}

	// 遍历所有类型一致性测试用例
	for _, tt := range tests {
		// 使用子测试，每个测试用例独立运行
		t.Run(tt.name, func(t *testing.T) {
			// 遍历所有输入类型
			for _, input := range tt.inputs {
				// 调用 ToUint32 函数并验证结果
				if got := ToUint32(input); got != tt.expected {
					// 如果结果不符合预期，记录错误信息
					t.Errorf("ToUint32(%v) = %v, want %v", input, got, tt.expected)
				}
			}
		})
	}
}

// TestToUint32_OverflowBehavior 测试 ToUint32 函数的溢出行为
// 该测试专门验证数值溢出时的处理方式
func TestToUint32_OverflowBehavior(t *testing.T) {
	// 定义溢出测试用例
	tests := []struct {
		name  string      // 测试用例名称
		input interface{} // 输入参数
		want  uint32      // 期望结果
	}{
		// 正溢出测试
		{"int64 overflow positive", int64(4294967296), 0},                 // 超出 uint32 范围的正数
		{"int64 overflow max", int64(9223372036854775807), 4294967295},    // 最大 int64 值
		{"uint64 overflow positive", uint64(4294967296), 0},               // 超出 uint32 范围的正数
		{"uint64 overflow max", uint64(18446744073709551615), 4294967295}, // 最大 uint64 值

		// 负数溢出测试
		{"int64 negative large", int64(-4294967297), 4294967295}, // 大负数
		{"int64 negative max", int64(-9223372036854775808), 0},   // 最小 int64 值

		// 浮点数溢出测试
		{"float64 overflow", float64(4294967296.0), 0},            // 浮点数正溢出：超出范围的 float64 转换为 uint32 时回绕到 0
		{"float64 underflow", float64(-4294967297.0), 4294967295}, // 浮点数负溢出：负数转换为 uint32 时回绕到最大值

		// 字符串溢出测试
		{"string overflow", "4294967296", 0},                // 字符串数值溢出
		{"string underflow", "-1", 0},                       // 字符串负数
		{"string large number", "999999999999999999999", 0}, // 极大数值字符串
	}

	// 遍历所有溢出测试用例
	for _, tt := range tests {
		// 使用子测试，每个测试用例独立运行
		t.Run(tt.name, func(t *testing.T) {
			// 调用 ToUint32 函数并验证结果
			if got := ToUint32(tt.input); got != tt.want {
				// 如果结果不符合预期，记录错误信息
				t.Errorf("ToUint32() = %v, want %v", got, tt.want)
			}
		})
	}
}

// TestToUint32_ErrorHandling 测试 ToUint32 函数的错误处理
// 该测试验证函数在面对无效输入时的错误处理能力
func TestToUint32_ErrorHandling(t *testing.T) {
	// 定义错误处理测试用例
	tests := []struct {
		name  string      // 测试用例名称
		input interface{} // 输入参数
		want  uint32      // 期望结果
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
			// 调用 ToUint32 函数并验证结果
			if got := ToUint32(tt.input); got != tt.want {
				// 如果结果不符合预期，记录错误信息
				t.Errorf("ToUint32() = %v, want %v", got, tt.want)
			}
		})
	}
}
