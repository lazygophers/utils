package anyx

import (
	"errors"
	"math"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

// stringTestStruct 测试用的结构体
type stringTestStruct struct {
	ID   int
	Name string
	Age  int
}

// jsonMarshalError 用于测试 JSON 序列化错误的类型
type jsonMarshalError struct {
	Msg string
}

// Error 实现 error 接口
func (e *jsonMarshalError) Error() string {
	return e.Msg
}

// TestToString 测试 ToString 函数
func TestToString(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		input    interface{}
		expected string
	}{
		{"bool true", true, "1"},
		{"bool false", false, "0"},
		{"int", 42, "42"},
		{"int8", int8(8), "8"},
		{"int16", int16(16), "16"},
		{"int32", int32(32), "32"},
		{"int64", int64(64), "64"},
		{"uint", uint(42), "42"},
		{"uint8", uint8(8), "8"},
		{"uint16", uint16(16), "16"},
		{"uint32", uint32(32), "32"},
		{"uint64", uint64(64), "64"},
		{"float32 integer", float32(42), "42"},
		{"float32 fraction", float32(3.14), "3.140000104904175"},
		{"float64 integer", float64(42), "42"},
		{"float64 fraction", float64(3.14), "3.140000"},
		{"time.Duration", time.Second, "1s"},
		{"string", "hello", "hello"},
		{"[]byte", []byte("hello"), "hello"},
		{"nil", nil, ""},
		{"error", errors.New("test error"), "test error"},
		{"struct", stringTestStruct{ID: 1, Name: "test", Age: 25}, `{"ID":1,"Name":"test","Age":25}`},
		{"slice", []int{1, 2, 3}, "[1,2,3]"},
		{"map", map[string]int{"key": 42}, `{"key":42}`},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			result := ToString(tt.input)
			assert.Equal(t, tt.expected, result, "ToString() 的结果应与期望值相等")
		})
	}
}

// TestToBytesComprehensive 测试 ToBytes 函数的所有分支
func TestToBytesComprehensive(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		input    interface{}
		expected []byte
	}{
		// 布尔类型
		{"bool true", true, []byte("1")},
		{"bool false", false, []byte("0")},

		// 整数类型
		{"int", int(42), []byte("42")},
		{"int min", int(-2147483648), []byte("-2147483648")},
		{"int max", int(2147483647), []byte("2147483647")},

		{"int8", int8(8), []byte("8")},
		{"int8 min", int8(-128), []byte("-128")},
		{"int8 max", int8(127), []byte("127")},

		{"int16", int16(16), []byte("16")},
		{"int16 min", int16(-32768), []byte("-32768")},
		{"int16 max", int16(32767), []byte("32767")},

		{"int32", int32(32), []byte("32")},
		{"int32 min", int32(-2147483648), []byte("-2147483648")},
		{"int32 max", int32(2147483647), []byte("2147483647")},

		{"int64", int64(64), []byte("64")},
		{"int64 min", int64(-9223372036854775808), []byte("-9223372036854775808")},
		{"int64 max", int64(9223372036854775807), []byte("9223372036854775807")},

		{"uint", uint(42), []byte("42")},
		{"uint max", uint(18446744073709551615), []byte("18446744073709551615")},

		{"uint8", uint8(8), []byte("8")},
		{"uint8 max", uint8(255), []byte("255")},

		{"uint16", uint16(16), []byte("16")},
		{"uint16 max", uint16(65535), []byte("65535")},

		{"uint32", uint32(32), []byte("32")},
		{"uint32 max", uint32(4294967295), []byte("4294967295")},

		{"uint64", uint64(64), []byte("64")},
		{"uint64 max", uint64(18446744073709551615), []byte("18446744073709551615")},

		// 浮点数类型 - 整数
		{"float32 integer", float32(42), []byte("42")},
		{"float32 max integer", float32(16777215), []byte("16777215")},
		{"float64 integer", float64(42), []byte("42")},
		{"float64 max integer", float64(9007199254740991), []byte("9007199254740991")},

		// 浮点数类型 - 小数
		{"float32 fraction", float32(3.14), []byte("3.140000104904175")},
		{"float32 negative fraction", float32(-2.71), []byte("-2.710000038146973")},
		{"float32 small fraction", float32(0.000001), []byte("0.000000999999997")},
		{"float64 fraction", float64(3.14), []byte("3.140000")},
		{"float64 negative fraction", float64(-2.71), []byte("-2.710000")},
		{"float64 small fraction", float64(0.000001), []byte("0.000001")},

		// 时间类型
		{"time.Duration", time.Second, []byte("1s")},
		{"time.Duration 0", time.Duration(0), []byte("0s")},
		{"time.Duration negative", time.Duration(-1), []byte("-1ns")},
		{"time.Duration large", 365 * 24 * time.Hour, []byte("8760h0m0s")},

		// 字符串和字节
		{"string", "hello", []byte("hello")},
		{"string empty", "", []byte{}},
		{"string unicode", "你好世界", []byte("你好世界")},
		{"[]byte", []byte("hello"), []byte("hello")},
		{"[]byte empty", []byte{}, []byte{}},

		// nil 值
		{"nil", nil, nil},

		// 错误类型
		{"error", errors.New("test error"), []byte("test error")},
		{"nil error", error(nil), nil},
		{"custom error", &jsonMarshalError{Msg: "custom error"}, []byte("custom error")},

		// 复杂类型 - JSON 序列化
		{"struct", stringTestStruct{ID: 1, Name: "test", Age: 25}, []byte(`{"ID":1,"Name":"test","Age":25}`)},
		{"slice", []int{1, 2, 3}, []byte("[1,2,3]")},
		{"map", map[string]int{"key": 42}, []byte(`{"key":42}`)},
		{"empty slice", []int{}, []byte("[]")},
		{"empty map", map[string]int{}, []byte("{}")},
		{"complex struct", struct {
			ID      int
			Name    string
			Age     int
			Address *struct {
				Street string
				City   string
			}
		}{
			ID:   1,
			Name: "test",
			Age:  25,
			Address: &struct {
				Street string
				City   string
			}{
				Street: "123 Main St",
				City:   "New York",
			},
		}, []byte(`{"ID":1,"Name":"test","Age":25,"Address":{"Street":"123 Main St","City":"New York"}}`)},

		// JSON 序列化失败的情况
		{"json marshal error", make(chan int), nil}, // channel 不能被 JSON 序列化
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			result := ToBytes(tt.input)
			assert.Equal(t, tt.expected, result, "ToBytes() 的结果应与期望值相等")
		})
	}
}

// TestToStringSliceComprehensive 测试 ToStringSlice 函数的所有分支
func TestToStringSliceComprehensive(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name      string
		input     interface{}
		separator string
		expected  []string
	}{
		// 布尔切片
		{"[]bool", []bool{true, false, true}, ",", []string{"1", "0", "1"}},
		{"[]bool empty", []bool{}, ",", []string{}},
		{"[]bool single", []bool{true}, ",", []string{"1"}},

		// 整数切片
		{"[]int", []int{1, 2, 3}, ",", []string{"1", "2", "3"}},
		{"[]int negative", []int{-1, 0, 1}, ",", []string{"-1", "0", "1"}},
		{"[]int empty", []int{}, ",", []string{}},
		{"[]int single", []int{42}, ",", []string{"42"}},

		// int8 切片
		{"[]int8", []int8{1, -2, 3}, ",", []string{"1", "-2", "3"}},
		{"[]int8 min max", []int8{-128, 0, 127}, ",", []string{"-128", "0", "127"}},

		// int16 切片
		{"[]int16", []int16{1, -2, 3}, ",", []string{"1", "-2", "3"}},
		{"[]int16 min max", []int16{-32768, 0, 32767}, ",", []string{"-32768", "0", "32767"}},

		// int32 切片
		{"[]int32", []int32{1, -2, 3}, ",", []string{"1", "-2", "3"}},
		{"[]int32 min max", []int32{-2147483648, 0, 2147483647}, ",", []string{"-2147483648", "0", "2147483647"}},

		// int64 切片
		{"[]int64", []int64{1, -2, 3}, ",", []string{"1", "-2", "3"}},
		{"[]int64 min max", []int64{-9223372036854775808, 0, 9223372036854775807}, ",", []string{"-9223372036854775808", "0", "9223372036854775807"}},

		// uint 切片
		{"[]uint", []uint{1, 2, 3}, ",", []string{"1", "2", "3"}},
		{"[]uint max", []uint{0, 18446744073709551615}, ",", []string{"0", "18446744073709551615"}},

		// uint16 切片
		{"[]uint16", []uint16{1, 2, 3}, ",", []string{"1", "2", "3"}},
		{"[]uint16 max", []uint16{0, 65535}, ",", []string{"0", "65535"}},

		// uint32 切片
		{"[]uint32", []uint32{1, 2, 3}, ",", []string{"1", "2", "3"}},
		{"[]uint32 max", []uint32{0, 4294967295}, ",", []string{"0", "4294967295"}},

		// uint64 切片
		{"[]uint64", []uint64{1, 2, 3}, ",", []string{"1", "2", "3"}},
		{"[]uint64 max", []uint64{0, 18446744073709551615}, ",", []string{"0", "18446744073709551615"}},

		// float32 切片
		{"[]float32", []float32{1, 2.5, 3}, ",", []string{"1", "2.5", "3"}},
		{"[]float32 negative", []float32{-1, -2.5, -3}, ",", []string{"-1", "-2.5", "-3"}},
		{"[]float32 integer", []float32{1, 2, 3}, ",", []string{"1", "2", "3"}},

		// float64 切片
		{"[]float64", []float64{1, 2.5, 3}, ",", []string{"1", "2.5", "3"}},
		{"[]float64 negative", []float64{-1, -2.5, -3}, ",", []string{"-1", "-2.5", "-3"}},
		{"[]float64 integer", []float64{1, 2, 3}, ",", []string{"1", "2", "3"}},

		// 字符串切片
		{"[]string", []string{"a", "b", "c"}, ",", []string{"a", "b", "c"}},
		{"[]string empty", []string{}, ",", []string{}},
		{"[]string with unicode", []string{"你好", "世界", "Go"}, ",", []string{"你好", "世界", "Go"}},

		// 字节切片 - JSON 数组
		{"[]byte json array", []byte("[1,2,3]"), ",", []string{"1", "2", "3"}},
		{"[]byte json array empty", []byte("[]"), ",", []string{}},
		{"[]byte json array string", []byte(`["a","b","c"]`), ",", []string{"a", "b", "c"}},
		{"[]byte json array mixed", []byte(`[1,"a",3.14]`), ",", []string{"1", "a", "3.140000"}},
		{"[]byte json array nested", []byte(`[1,[2,3],{"a":"b"}]`), ",", []string{"1", "[2,3]", `{"a":"b"}`}},

		// 字节切片 - 非 JSON 数组
		{"[]byte plain", []byte("hello world"), ",", []string{"hello world"}},
		{"[]byte plain with brackets", []byte("[not json]"), ",", []string{"[not json]"}},
		{"[]byte plain with comma", []byte("a,b,c"), ",", []string{"a", "b", "c"}},

		// 字节切片 - 空分隔符
		{"[]byte empty separator", []byte("a,b,c"), "", []string{"a", "b", "c"}},
		{"[]byte multi separator", []byte("a|b|c"), "|", []string{"a", "b", "c"}},

		// 字符串 - JSON 数组
		{"string json array", "[1,2,3]", ",", []string{"1", "2", "3"}},
		{"string json array string", `["a","b","c"]`, ",", []string{"a", "b", "c"}},
		{"string json array mixed", `[1,"a",3.14]`, ",", []string{"1", "a", "3.140000"}},
		{"string json array nested", `[1,[2,3],{"a":"b"}]`, ",", []string{"1", "[2,3]", `{"a":"b"}`}},
		{"string json array invalid", "[1,2,3", ",", []string{"[1", "2", "3"}}, // 无效的 JSON

		// 字符串 - 非 JSON 数组
		{"string plain", "hello world", ",", []string{"hello world"}},
		{"string plain with brackets", "[not json]", ",", []string{"[not json]"}},
		{"string plain with comma", "a,b,c", ",", []string{"a", "b", "c"}},

		// 字符串 - 空分隔符
		{"string empty separator", "a,b,c", "", []string{"a", "b", "c"}},
		{"string multi separator", "a|b|c", "|", []string{"a", "b", "c"}},

		// []interface{} 切片
		{"[]interface{}", []interface{}{1, "a", 3.14, true}, ",", []string{"1", "a", "3.140000", "1"}},
		{"[]interface{} empty", []interface{}{}, ",", []string{}},
		{"[]interface{} with nil", []interface{}{1, nil, 3}, ",", []string{"1", "", "3"}},

		// 默认分隔符测试
		{"default separator", []int{1, 2, 3}, "", []string{"1", "2", "3"}}, // 不传分隔符参数

		// 不支持的类型
		{"int", 42, ",", nil},
		{"nil", nil, ",", nil},
		{"map", map[string]int{"key": 42}, ",", nil},
		{"struct", stringTestStruct{ID: 1}, ",", nil},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			var result []string
			if tt.separator == "" {
				result = ToStringSlice(tt.input)
			} else {
				result = ToStringSlice(tt.input, tt.separator)
			}
			assert.Equal(t, tt.expected, result, "ToStringSlice() 的结果应与期望值相等")
		})
	}
}

// TestToStringSliceEdgeCases 测试边界情况
func TestToStringSliceEdgeCases(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name      string
		input     interface{}
		separator string
		expected  []string
	}{
		{"string with empty parts", "a,,c", ",", []string{"a", "", "c"}},
		{"string with trailing separator", "a,b,", ",", []string{"a", "b", ""}},
		{"string with leading separator", ",a,b", ",", []string{"", "a", "b"}},
		{"string only separators", ",,,", ",", []string{"", "", "", ""}},
		{"string empty separator empty string", "", "", []string{""}},
		{"string empty separator with content", "abc", "", []string{"abc"}},
		{"string multi separator", "a::b::c", "::", []string{"a", "b", "c"}},
		{"string separator with spaces", "a, b, c", ", ", []string{"a", "b", "c"}},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			result := ToStringSlice(tt.input, tt.separator)
			assert.Equal(t, tt.expected, result, "ToStringSlice() 处理边界情况应正确")
		})
	}
}

// TestToStringSliceJsonHandling 测试 JSON 处理相关场景
func TestToStringSliceJsonHandling(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name         string
		input        interface{}
		separator    string
		expected     []string
		expectError  bool
	}{
		// 有效的 JSON 字符串
		{"valid json string", `["value1","value2"]`, ",", []string{"value1", "value2"}, false},
		{"valid json string numbers", `[1,2,3]`, ",", []string{"1", "2", "3"}, false},
		{"valid json string mixed", `[1,"value",3.14]`, ",", []string{"1", "value", "3.140000"}, false},
		{"valid json string nested", `[1,[2,3],{"key":"value"}]`, ",", []string{"1", "[2,3]", `{"key":"value"}`}, false},

		// 有效的 JSON 字节切片
		{"valid json bytes", []byte(`["value1","value2"]`), ",", []string{"value1", "value2"}, false},
		{"valid json bytes numbers", []byte(`[1,2,3]`), ",", []string{"1", "2", "3"}, false},
		{"valid json bytes mixed", []byte(`[1,"value",3.14]`), ",", []string{"1", "value", "3.140000"}, false},

		// 无效的 JSON 字符串
		{"invalid json string", "[invalid json", ",", []string{"[invalid json"}, false},
		{"invalid json string2", "[1,2,,3]", ",", []string{"[1", "2", "", "3]"}, false},
		{"invalid json string3", "[1,2,3", ",", []string{"[1", "2", "3"}, false},
		{"invalid json string4", "[1,,3]", ",", []string{"[1", "", "3]"}, false},

		// 无效的 JSON 字节切片
		{"invalid json bytes", []byte("[invalid json"), ",", []string{"[invalid json"}, false},
		{"invalid json bytes2", []byte("[1,2,,3]"), ",", []string{"[1", "2", "", "3]"}, false},

		// JSON 解析错误但返回原字符串
		{"json error fallback string", "[invalid, json", ",", []string{"[invalid", " json"}, false},
		{"json error fallback bytes", []byte("[invalid, json"), ",", []string{"[invalid", " json"}, false},

		// 看起来像 JSON 但实际不是
		{"looks like json but not", "[123]", ",", []string{"123"}, false},
		{"looks like json but not bytes", []byte("[123]"), ",", []string{"123"}, false},

		// 空 JSON 相关情况
		{"empty json string", "[]", ",", []string{}, false},
		{"empty json bytes", []byte("[]"), ",", []string{}, false},
		{"whitespace json string", "   ", ",", []string{"   "}, false},
		{"whitespace json bytes", []byte("   "), ",", []string{"   "}, false},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			var result []string
			if tt.separator == "" {
				result = ToStringSlice(tt.input)
			} else {
				result = ToStringSlice(tt.input, tt.separator)
			}
			assert.Equal(t, tt.expected, result, "ToStringSlice() 的结果应与期望值相等")
		})
	}
}

// BenchmarkToString 性能基准测试
func BenchmarkToString(b *testing.B) {
	b.Run("int", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = ToString(42)
		}
	})

	b.Run("string", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = ToString("hello")
		}
	})

	b.Run("[]byte", func(b *testing.B) {
		data := []byte("hello")
		for i := 0; i < b.N; i++ {
			_ = ToString(data)
		}
	})

	b.Run("float64", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = ToString(3.14)
		}
	})

	b.Run("struct", func(b *testing.B) {
		s := stringTestStruct{ID: 1, Name: "test", Age: 25}
		for i := 0; i < b.N; i++ {
			_ = ToString(s)
		}
	})
}

// BenchmarkToBytes 性能基准测试
func BenchmarkToBytes(b *testing.B) {
	b.Run("string", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = ToBytes("hello")
		}
	})

	b.Run("[]byte", func(b *testing.B) {
		data := []byte("hello")
		for i := 0; i < b.N; i++ {
			_ = ToBytes(data)
		}
	})

	b.Run("int", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = ToBytes(42)
		}
	})

	b.Run("float64", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = ToBytes(3.14)
		}
	})

	b.Run("struct", func(b *testing.B) {
		s := stringTestStruct{ID: 1, Name: "test", Age: 25}
		for i := 0; i < b.N; i++ {
			_ = ToBytes(s)
		}
	})

	b.Run("json marshal error", func(b *testing.B) {
		ch := make(chan int)
		for i := 0; i < b.N; i++ {
			_ = ToBytes(ch)
		}
	})
}

// BenchmarkToStringSlice 性能基准测试
func BenchmarkToStringSlice(b *testing.B) {
	b.Run("[]int", func(b *testing.B) {
		data := []int{1, 2, 3}
		for i := 0; i < b.N; i++ {
			_ = ToStringSlice(data, ",")
		}
	})

	b.Run("[]string", func(b *testing.B) {
		data := []string{"a", "b", "c"}
		for i := 0; i < b.N; i++ {
			_ = ToStringSlice(data, ",")
		}
	})

	b.Run("string", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = ToStringSlice("a,b,c", ",")
		}
	})

	b.Run("[]int large", func(b *testing.B) {
		data := make([]int, 1000)
		for i := 0; i < 1000; i++ {
			data[i] = i
		}
		for i := 0; i < b.N; i++ {
			_ = ToStringSlice(data, ",")
		}
	})

	b.Run("json array string", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = ToStringSlice("[1,2,3]", ",")
		}
	})

	b.Run("json array bytes", func(b *testing.B) {
		data := []byte("[1,2,3]")
		for i := 0; i < b.N; i++ {
			_ = ToStringSlice(data, ",")
		}
	})

	b.Run("[]interface{}", func(b *testing.B) {
		data := []interface{}{1, "a", 3.14, true}
		for i := 0; i < b.N; i++ {
			_ = ToStringSlice(data, ",")
		}
	})
}

// TestToBytesFloatSpecialCases 测试浮点数特殊情况
func TestToBytesFloatSpecialCases(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		input    interface{}
		expected []byte
	}{
		{"float32 NaN", float32(math.NaN()), []byte("NaN")},
		{"float32 +Inf", float32(math.Inf(1)), []byte("+Inf")},
		{"float32 -Inf", float32(math.Inf(-1)), []byte("-Inf")},
		{"float64 NaN", math.NaN(), []byte("NaN")},
		{"float64 +Inf", math.Inf(1), []byte("+Inf")},
		{"float64 -Inf", math.Inf(-1), []byte("-Inf")},
		{"float32 smallest positive", float32(1.401298464324817e-45), []byte("0.000000000000000")},
		{"float32 largest positive", float32(math.MaxFloat32), []byte("340282346638528859811704183484516925440")},
		{"float64 smallest positive", float64(4.9406564584124654e-324), []byte("0.000000")},
		{"float64 largest positive", float64(math.MaxFloat64), []byte("179769313486231570814527423731704356798070567525844996598917476803157260780028538760589558632766878171540458953514382464234321326889464182768467546703537516986049910576551282076245490090389328944075868508455133942304583236903222948165808559332123348274797826204144723168738177180919299881250404026184124858368")},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			result := ToBytes(tt.input)
			assert.Equal(t, tt.expected, result, "ToBytes() 处理浮点数特殊情况应正确")
		})
	}
}

// TestToStringSliceDefaultSeparator 测试默认分隔符
func TestToStringSliceDefaultSeparator(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		input    interface{}
		expected []string
	}{
		{"[]bool default separator", []bool{true, false}, []string{"1", "0"}},
		{"[]int default separator", []int{1, 2, 3}, []string{"1", "2", "3"}},
		{"[]string default separator", []string{"a", "b", "c"}, []string{"a", "b", "c"}},
		{"string default separator", "a,b,c", []string{"a", "b", "c"}},
		{"string default separator different sep", "a|b|c", []string{"a|b|c"}},
		{"[]byte default separator json", []byte("[1,2,3]"), []string{"1", "2", "3"}},
		{"[]byte default separator plain", []byte("a,b,c"), []string{"a", "b", "c"}},
		{"[]interface{} default separator", []interface{}{1, "a", true}, []string{"1", "a", "1"}},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			result := ToStringSlice(tt.input) // 不传分隔符参数
			assert.Equal(t, tt.expected, result, "ToStringSlice() 使用默认分隔符应正确")
		})
	}
}

// TestToBytesPrivate 测试 toBytes 私有函数
func TestToBytesPrivate(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		input    string
		expected []byte
	}{
		{"empty string", "", []byte{}},
		{"simple string", "hello", []byte("hello")},
		{"unicode string", "你好世界", []byte("你好世界")},
		{"with special chars", "a\nb\tc", []byte("a\nb\tc")},
		{"with emojis", "hello 😀", []byte("hello 😀")},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			result := toBytes(tt.input)
			assert.Equal(t, tt.expected, result, "toBytes() 的结果应与期望值相等")
		})
	}
}

// TestToArrayString 测试 ToArrayString 函数
func TestToArrayString(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		input    interface{}
		expected []string
	}{
		{"string input", "hello", []string{"hello"}},
		{"[]string input", []string{"a", "b", "c"}, []string{"a", "b", "c"}},
		{"empty string", "", []string{""}},
		{"empty []string", []string{}, []string{}},
		{"string with comma", "a,b,c", []string{"a", "b", "c"}},
		{"string with separator", "a|b|c", []string{"a|b|c"}},
		{"single element []string", []string{"single"}, []string{"single"}},
		{"unicode string", "你好", []string{"你好"}},
		{"unicode []string", []string{"你好", "世界"}, []string{"你好", "世界"}},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			result := ToArrayString(tt.input)
			assert.Equal(t, tt.expected, result, "ToArrayString() 的结果应与期望值相等")
		})
	}
}