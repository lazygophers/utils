package candy

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

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

		// Note: []uint8 is the same type as []byte in Go, so it's handled by the []byte case

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
		{"struct", struct{ ID int }{ID: 1}, ",", nil},

		// 边界情况测试
		{"[]byte invalid json", []byte("[invalid}"), ",", []string{"[invalid}"}},
		{"string invalid json", "[invalid}", ",", []string{"[invalid}"}},
		{"[]byte partial json", []byte("[1,2"), ",", []string{"[1", "2"}},

		// 测试空分隔符的更多情况
		{"[]byte empty sep json valid", []byte("[1,2,3]"), "", []string{"1", "2", "3"}},
		{"[]byte empty sep not json", []byte("hello"), "", []string{"hello"}},
		{"string empty sep json valid", "[1,2,3]", "", []string{"1", "2", "3"}},
		{"string empty sep not json", "hello", "", []string{"hello"}},
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

// TestToStringSliceMissingCoverage tests specific uncovered lines
func TestToStringSliceMissingCoverage(t *testing.T) {
	t.Run("uint8_slice_coverage", func(t *testing.T) {
		// Test []uint8 which is the same as []byte and should be handled by []byte case
		// []uint8 is handled as []byte, so it should split by comma
		data := []uint8{49, 44, 50, 44, 51} // ASCII for "1,2,3"
		result := ToStringSlice(data, ",")
		expected := []string{"1", "2", "3"}
		assert.Equal(t, expected, result)
	})

	t.Run("float32_integer_format", func(t *testing.T) {
		// Test float32 that should be formatted as integer (line 101-102)
		data := []float32{1.0, 2.0, 3.0} // These are whole numbers
		result := ToStringSlice(data, ",")
		expected := []string{"1", "2", "3"}
		assert.Equal(t, expected, result)
	})

	t.Run("float64_integer_format", func(t *testing.T) {
		// Test float64 that should be formatted as integer (line 112-113)
		data := []float64{1.0, 2.0, 3.0} // These are whole numbers
		result := ToStringSlice(data, ",")
		expected := []string{"1", "2", "3"}
		assert.Equal(t, expected, result)
	})

	t.Run("byte_slice_empty_separator", func(t *testing.T) {
		// Test []byte with empty separator to trigger line 3297-3299
		data := []byte("hello")
		result := ToStringSlice(data, "")
		expected := []string{"hello"}
		assert.Equal(t, expected, result)
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
