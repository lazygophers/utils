package candy

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

// ================================
// 布尔转换测试组
// ================================

func TestToBool(t *testing.T) {
	cases := []struct {
		name     string
		input    interface{}
		expected bool
	}{
		// bool
		{"bool true", true, true},
		{"bool false", false, false},

		// int
		{"int 0", 0, false},
		{"int 1", 1, true},
		{"int -1", -1, true},
		{"int8 0", int8(0), false},
		{"int16 10", int16(10), true},
		{"int32 -10", int32(-10), true},
		{"int64 100", int64(100), true},

		// uint
		{"uint 0", uint(0), false},
		{"uint 1", uint(1), true},
		{"uint8 0", uint8(0), false},
		{"uint16 10", uint16(10), true},
		{"uint32 10", uint32(10), true},
		{"uint64 100", uint64(100), true},

		// float
		{"float32 0.0", float32(0.0), false},
		{"float32 -0.0", float32(-0.0), false},
		{"float32 1.23", float32(1.23), true},
		{"float32 -4.56", float32(-4.56), true},
		{"float32 NaN", float32(math.NaN()), false},
		{"float32 +Inf", float32(math.Inf(1)), true},
		{"float32 -Inf", float32(math.Inf(-1)), true},
		{"float64 0.0", 0.0, false},
		{"float64 -0.0", -0.0, false},
		{"float64 1.23", 1.23, true},
		{"float64 -4.56", -4.56, true},
		{"float64 NaN", math.NaN(), false},
		{"float64 +Inf", math.Inf(1), true},
		{"float64 -Inf", math.Inf(-1), true},

		// string (true values)
		{"string true", "true", true},
		{"string TRUE", "TRUE", true},
		{"string True", "True", true},
		{"string tRuE", "tRuE", true},
		{"string 1", "1", true},
		{"string t", "t", true},
		{"string T", "T", true},
		{"string y", "y", true},
		{"string Y", "Y", true},
		{"string yes", "yes", true},
		{"string YES", "YES", true},
		{"string Yes", "Yes", true},
		{"string yEs", "yEs", true},
		{"string on", "on", true},
		{"string ON", "ON", true},
		{"string On", "On", true},
		{"string oN", "oN", true},

		// string (false values)
		{"string false", "false", false},
		{"string FALSE", "FALSE", false},
		{"string False", "False", false},
		{"string fAlSe", "fAlSe", false},
		{"string 0", "0", false},
		{"string f", "f", false},
		{"string F", "F", false},
		{"string n", "n", false},
		{"string N", "N", false},
		{"string no", "no", false},
		{"string NO", "NO", false},
		{"string No", "No", false},
		{"string nO", "nO", false},
		{"string off", "off", false},
		{"string OFF", "OFF", false},
		{"string Off", "Off", false},
		{"string oFf", "oFf", false},

		// string (other non-empty)
		{"string hello", "hello", true},
		{"string with spaces", "  hello  ", true},

		// string (empty)
		{"string empty", "", false},
		{"string space only", "   ", false},
		{"string tab and newline", " \t\n \r\f\v ", false},

		// []byte (true values)
		{"[]byte true", []byte("true"), true},
		{"[]byte TRUE", []byte("TRUE"), true},
		{"[]byte True", []byte("True"), true},
		{"[]byte 1", []byte("1"), true},
		{"[]byte t", []byte("t"), true},
		{"[]byte T", []byte("T"), true},
		{"[]byte y", []byte("y"), true},
		{"[]byte Y", []byte("Y"), true},
		{"[]byte yes", []byte("yes"), true},
		{"[]byte YES", []byte("YES"), true},
		{"[]byte on", []byte("on"), true},
		{"[]byte ON", []byte("ON"), true},

		// []byte (false values)
		{"[]byte false", []byte("false"), false},
		{"[]byte FALSE", []byte("FALSE"), false},
		{"[]byte False", []byte("False"), false},
		{"[]byte 0", []byte("0"), false},
		{"[]byte f", []byte("f"), false},
		{"[]byte F", []byte("F"), false},
		{"[]byte n", []byte("n"), false},
		{"[]byte N", []byte("N"), false},
		{"[]byte no", []byte("no"), false},
		{"[]byte NO", []byte("NO"), false},
		{"[]byte off", []byte("off"), false},
		{"[]byte OFF", []byte("OFF"), false},

		// []byte (other non-empty)
		{"[]byte hello", []byte("hello"), true},
		{"[]byte with spaces", []byte("  hello  "), true},

		// []byte (empty)
		{"[]byte empty", []byte(""), false},
		{"[]byte space only", []byte("   "), false},
		{"[]byte tab and newline", []byte(" \t\n "), false},
		{"[]byte nil", []byte(nil), false},

		// nil
		{"nil", nil, false},

		// unsupported types
		{"unsupported struct", struct{}{}, false},
		{"unsupported map", make(map[int]int), false},
		{"unsupported slice", []int{1, 2}, false},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			if got := ToBool(tc.input); got != tc.expected {
				t.Errorf("ToBool(%v) = %v; want %v", tc.input, got, tc.expected)
			}
		})
	}
}

func BenchmarkToBool(b *testing.B) {
	cases := []struct {
		input interface{}
	}{
		{true},
		{false},
		{1},
		{0},
		{1.23},
		{0.0},
		{float32(4.56)},
		{math.NaN()},
		{"true"},
		{"false"},
		{"1"},
		{"0"},
		{"t"},
		{"f"},
		{"hello"},
		{""},
		{[]byte("true")},
		{[]byte("false")},
		{[]byte("1")},
		{[]byte("0")},
		{[]byte("t")},
		{[]byte("f")},
		{[]byte("hello")},
		{[]byte("")},
		{[]byte("   ")},
		{nil},
		{struct{}{}},
	}

	b.ReportAllocs()
	b.ResetTimer()

	var r bool
	for i := 0; i < b.N; i++ {
		r = ToBool(cases[i%len(cases)].input)
	}
	_ = r
}

// ================================
// 字符串转换测试组
// ================================

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
		{"channel", make(chan int), ""}, // channels cannot be marshaled to JSON, should return empty string
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
		{"nil input", nil, nil},
		{"nil string slice", ([]string)(nil), nil},
		{"int slice", []int{1, 2, 3}, []string{"1", "2", "3"}},
		{"bool slice", []bool{true, false}, []string{"1", "0"}},                             // ToString converts bool to "1"/"0"
		{"float slice", []float64{1.1, 2.2}, []string{"1.100000", "2.200000"}},              // ToString precision
		{"interface slice", []interface{}{"hello", 42, true}, []string{"hello", "42", "1"}}, // bool becomes "1"
		{"non-slice int", 42, []string{"42"}},
		{"non-slice bool", true, []string{"1"}},                                                // ToString converts true to "1"
		{"non-slice float", 3.14, []string{"3.140000"}},                                        // ToString precision
		{"struct input", struct{ Name string }{Name: "test"}, []string{`{"Name":"test"}`}},     // JSON serialization
		{"map input", map[string]int{"key": 42}, []string{`{"key":42}`}},                       // JSON serialization
		{"slice with nil elements", []interface{}{nil, "test", nil}, []string{"", "test", ""}}, // nil becomes empty string
		{"string with multiple commas", "a,b,c,d,e", []string{"a", "b", "c", "d", "e"}},
		{"string with spaces and commas", "a, b, c", []string{"a", " b", " c"}}, // preserves spaces
		{"empty slice input", []int{}, []string{}},                              // empty slice
		{"single element slice", []string{"only"}, []string{"only"}},            // already covered but explicit
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

// ================================
// 数值转换测试组
// ================================

// TestToInt 测试 ToInt 函数
func TestToInt(t *testing.T) {
	t.Run("bool values", func(t *testing.T) {
		assert.Equal(t, 1, ToInt(true))
		assert.Equal(t, 0, ToInt(false))
	})

	t.Run("int values", func(t *testing.T) {
		assert.Equal(t, 0, ToInt(int(0)))
		assert.Equal(t, 42, ToInt(int(42)))
		assert.Equal(t, -42, ToInt(int(-42)))
	})

	t.Run("int8 values", func(t *testing.T) {
		assert.Equal(t, 127, ToInt(int8(127)))
		assert.Equal(t, -128, ToInt(int8(-128)))
	})

	t.Run("int16 values", func(t *testing.T) {
		assert.Equal(t, 32767, ToInt(int16(32767)))
		assert.Equal(t, -32768, ToInt(int16(-32768)))
	})

	t.Run("int32 values", func(t *testing.T) {
		assert.Equal(t, 100, ToInt(int32(100)))
		assert.Equal(t, -100, ToInt(int32(-100)))
	})

	t.Run("int64 values", func(t *testing.T) {
		assert.Equal(t, 1000, ToInt(int64(1000)))
		assert.Equal(t, -1000, ToInt(int64(-1000)))
	})

	t.Run("uint values", func(t *testing.T) {
		assert.Equal(t, 0, ToInt(uint(0)))
		assert.Equal(t, 42, ToInt(uint(42)))
	})

	t.Run("uint8 values", func(t *testing.T) {
		assert.Equal(t, 255, ToInt(uint8(255)))
		assert.Equal(t, 0, ToInt(uint8(0)))
	})

	t.Run("uint16 values", func(t *testing.T) {
		assert.Equal(t, 65535, ToInt(uint16(65535)))
	})

	t.Run("uint32 values", func(t *testing.T) {
		assert.Equal(t, 100, ToInt(uint32(100)))
	})

	t.Run("uint64 values", func(t *testing.T) {
		assert.Equal(t, 1000, ToInt(uint64(1000)))
	})

	t.Run("float32 values", func(t *testing.T) {
		assert.Equal(t, 3, ToInt(float32(3.14)))
		assert.Equal(t, -3, ToInt(float32(-3.14)))
		assert.Equal(t, 0, ToInt(float32(0.0)))
		assert.Equal(t, 0, ToInt(float32(0.9))) // 截断小数部分
	})

	t.Run("float64 values", func(t *testing.T) {
		assert.Equal(t, 3, ToInt(float64(3.14)))
		assert.Equal(t, -3, ToInt(float64(-3.14)))
		assert.Equal(t, 0, ToInt(float64(0.0)))
		assert.Equal(t, 0, ToInt(float64(0.9))) // 截断小数部分
	})

	t.Run("string values", func(t *testing.T) {
		assert.Equal(t, 42, ToInt("42"))
		assert.Equal(t, 0, ToInt("0"))
		assert.Equal(t, 123, ToInt("123"))
	})

	t.Run("string values", func(t *testing.T) {
		assert.Equal(t, -42, ToInt("-42")) // ToInt实际上支持负数字符串
		assert.Equal(t, 0, ToInt("invalid"))
		assert.Equal(t, 0, ToInt(""))
		assert.Equal(t, 0, ToInt("abc"))
		assert.Equal(t, 0, ToInt("3.14"))
	})

	t.Run("valid byte slice values", func(t *testing.T) {
		assert.Equal(t, 42, ToInt([]byte("42")))
		assert.Equal(t, 0, ToInt([]byte("0")))
		assert.Equal(t, 123, ToInt([]byte("123")))
		assert.Equal(t, -42, ToInt([]byte("-42"))) // ToInt支持负数
	})

	t.Run("invalid byte slice values", func(t *testing.T) {
		assert.Equal(t, 0, ToInt([]byte("invalid")))
		assert.Equal(t, 0, ToInt([]byte("")))
		assert.Equal(t, 0, ToInt([]byte("abc")))
	})

	t.Run("unsupported types", func(t *testing.T) {
		assert.Equal(t, 0, ToInt(nil))
		assert.Equal(t, 0, ToInt(struct{}{}))
		assert.Equal(t, 0, ToInt(map[string]int{}))
		assert.Equal(t, 0, ToInt([]int{1, 2, 3}))
		assert.Equal(t, 0, ToInt(func() {}))
	})
}

// TestToFloat32 测试 ToFloat32 函数
func TestToFloat32(t *testing.T) {
	t.Run("bool values", func(t *testing.T) {
		assert.Equal(t, float32(1), ToFloat32(true))
		assert.Equal(t, float32(0), ToFloat32(false))
	})

	t.Run("int values", func(t *testing.T) {
		assert.Equal(t, float32(0), ToFloat32(int(0)))
		assert.Equal(t, float32(42), ToFloat32(int(42)))
		assert.Equal(t, float32(-42), ToFloat32(int(-42)))
	})

	t.Run("int8 values", func(t *testing.T) {
		assert.Equal(t, float32(127), ToFloat32(int8(127)))
		assert.Equal(t, float32(-128), ToFloat32(int8(-128)))
	})

	t.Run("int16 values", func(t *testing.T) {
		assert.Equal(t, float32(32767), ToFloat32(int16(32767)))
		assert.Equal(t, float32(-32768), ToFloat32(int16(-32768)))
	})

	t.Run("int32 values", func(t *testing.T) {
		assert.Equal(t, float32(100), ToFloat32(int32(100)))
		assert.Equal(t, float32(-100), ToFloat32(int32(-100)))
	})

	t.Run("int64 values", func(t *testing.T) {
		assert.Equal(t, float32(1000), ToFloat32(int64(1000)))
		assert.Equal(t, float32(-1000), ToFloat32(int64(-1000)))
	})

	t.Run("uint values", func(t *testing.T) {
		assert.Equal(t, float32(0), ToFloat32(uint(0)))
		assert.Equal(t, float32(42), ToFloat32(uint(42)))
	})

	t.Run("uint8 values", func(t *testing.T) {
		assert.Equal(t, float32(255), ToFloat32(uint8(255)))
		assert.Equal(t, float32(0), ToFloat32(uint8(0)))
	})

	t.Run("uint16 values", func(t *testing.T) {
		assert.Equal(t, float32(65535), ToFloat32(uint16(65535)))
	})

	t.Run("uint32 values", func(t *testing.T) {
		assert.Equal(t, float32(100), ToFloat32(uint32(100)))
	})

	t.Run("uint64 values", func(t *testing.T) {
		assert.Equal(t, float32(1000), ToFloat32(uint64(1000)))
	})

	t.Run("float32 values", func(t *testing.T) {
		assert.Equal(t, float32(3.14), ToFloat32(float32(3.14)))
		assert.Equal(t, float32(-3.14), ToFloat32(float32(-3.14)))
		assert.Equal(t, float32(0.0), ToFloat32(float32(0.0)))
	})

	t.Run("float64 values", func(t *testing.T) {
		assert.Equal(t, float32(3.14), ToFloat32(float64(3.14)))
		assert.Equal(t, float32(-3.14), ToFloat32(float64(-3.14)))
		assert.Equal(t, float32(0.0), ToFloat32(float64(0.0)))
	})

	t.Run("string values", func(t *testing.T) {
		assert.Equal(t, float32(3.14), ToFloat32("3.14"))
		assert.Equal(t, float32(-3.14), ToFloat32("-3.14"))
		assert.Equal(t, float32(42), ToFloat32("42"))
		assert.Equal(t, float32(0), ToFloat32("0"))
		assert.Equal(t, float32(3.14), ToFloat32("  3.14  ")) // 空格处理
	})

	t.Run("invalid string values", func(t *testing.T) {
		assert.Equal(t, float32(0), ToFloat32("invalid"))
		assert.Equal(t, float32(0), ToFloat32(""))
		assert.Equal(t, float32(0), ToFloat32("abc"))
		assert.Equal(t, float32(0), ToFloat32("3.14.15"))
	})

	t.Run("byte slice values", func(t *testing.T) {
		assert.Equal(t, float32(3.14), ToFloat32([]byte("3.14")))
		assert.Equal(t, float32(-3.14), ToFloat32([]byte("-3.14")))
		assert.Equal(t, float32(42), ToFloat32([]byte("42")))
		assert.Equal(t, float32(3.14), ToFloat32([]byte("  3.14  "))) // 空格处理
	})

	t.Run("invalid byte slice values", func(t *testing.T) {
		assert.Equal(t, float32(0), ToFloat32([]byte("invalid")))
		assert.Equal(t, float32(0), ToFloat32([]byte("")))
		assert.Equal(t, float32(0), ToFloat32([]byte("abc")))
	})

	t.Run("unsupported types", func(t *testing.T) {
		assert.Equal(t, float32(0), ToFloat32(nil))
		assert.Equal(t, float32(0), ToFloat32(struct{}{}))
		assert.Equal(t, float32(0), ToFloat32(map[string]int{}))
		assert.Equal(t, float32(0), ToFloat32([]int{1, 2, 3}))
		assert.Equal(t, float32(0), ToFloat32(func() {}))
	})
}

// TestToFloat64 测试 ToFloat64 函数
func TestToFloat64(t *testing.T) {
	t.Run("bool values", func(t *testing.T) {
		assert.Equal(t, 1.0, ToFloat64(true))
		assert.Equal(t, 0.0, ToFloat64(false))
	})

	t.Run("int values", func(t *testing.T) {
		assert.Equal(t, 0.0, ToFloat64(int(0)))
		assert.Equal(t, 42.0, ToFloat64(int(42)))
		assert.Equal(t, -42.0, ToFloat64(int(-42)))
	})

	t.Run("int8 values", func(t *testing.T) {
		assert.Equal(t, 127.0, ToFloat64(int8(127)))
		assert.Equal(t, -128.0, ToFloat64(int8(-128)))
	})

	t.Run("int16 values", func(t *testing.T) {
		assert.Equal(t, 32767.0, ToFloat64(int16(32767)))
		assert.Equal(t, -32768.0, ToFloat64(int16(-32768)))
	})

	t.Run("int32 values", func(t *testing.T) {
		assert.Equal(t, 100.0, ToFloat64(int32(100)))
		assert.Equal(t, -100.0, ToFloat64(int32(-100)))
	})

	t.Run("int64 values", func(t *testing.T) {
		assert.Equal(t, 1000.0, ToFloat64(int64(1000)))
		assert.Equal(t, -1000.0, ToFloat64(int64(-1000)))
	})

	t.Run("uint values", func(t *testing.T) {
		assert.Equal(t, 0.0, ToFloat64(uint(0)))
		assert.Equal(t, 42.0, ToFloat64(uint(42)))
	})

	t.Run("uint8 values", func(t *testing.T) {
		assert.Equal(t, 255.0, ToFloat64(uint8(255)))
		assert.Equal(t, 0.0, ToFloat64(uint8(0)))
	})

	t.Run("uint16 values", func(t *testing.T) {
		assert.Equal(t, 65535.0, ToFloat64(uint16(65535)))
	})

	t.Run("uint32 values", func(t *testing.T) {
		assert.Equal(t, 100.0, ToFloat64(uint32(100)))
	})

	t.Run("uint64 values", func(t *testing.T) {
		assert.Equal(t, 1000.0, ToFloat64(uint64(1000)))
	})

	t.Run("float32 values", func(t *testing.T) {
		assert.Equal(t, 3.140000104904175, ToFloat64(float32(3.14))) // float32 precision
		assert.Equal(t, -3.140000104904175, ToFloat64(float32(-3.14)))
		assert.Equal(t, 0.0, ToFloat64(float32(0.0)))
	})

	t.Run("float64 values", func(t *testing.T) {
		assert.Equal(t, 3.14, ToFloat64(float64(3.14)))
		assert.Equal(t, -3.14, ToFloat64(float64(-3.14)))
		assert.Equal(t, 0.0, ToFloat64(float64(0.0)))
	})

	t.Run("string float values", func(t *testing.T) {
		assert.Equal(t, 3.14, ToFloat64("3.14"))
		assert.Equal(t, -3.14, ToFloat64("-3.14"))
		assert.Equal(t, 0.0, ToFloat64("0"))
		assert.Equal(t, 3.14, ToFloat64("  3.14  ")) // 空格处理
	})

	t.Run("string int values", func(t *testing.T) {
		assert.Equal(t, 42.0, ToFloat64("42"))
		assert.Equal(t, -42.0, ToFloat64("-42"))
		assert.Equal(t, 0.0, ToFloat64("0"))
	})

	t.Run("string hex values", func(t *testing.T) {
		assert.Equal(t, 255.0, ToFloat64("0xFF"))
		assert.Equal(t, 255.0, ToFloat64("0xff"))
		assert.Equal(t, 8.0, ToFloat64("0o10")) // octal
	})

	t.Run("invalid string values", func(t *testing.T) {
		assert.Equal(t, 0.0, ToFloat64("invalid"))
		assert.Equal(t, 0.0, ToFloat64(""))
		assert.Equal(t, 0.0, ToFloat64("abc"))
		assert.Equal(t, 0.0, ToFloat64("3.14.15"))
	})

	t.Run("byte slice values", func(t *testing.T) {
		assert.Equal(t, 3.14, ToFloat64([]byte("3.14")))
		assert.Equal(t, -3.14, ToFloat64([]byte("-3.14")))
		assert.Equal(t, 42.0, ToFloat64([]byte("42")))
		assert.Equal(t, 255.0, ToFloat64([]byte("0xFF")))
		assert.Equal(t, 3.14, ToFloat64([]byte("  3.14  "))) // 空格处理
	})

	t.Run("invalid byte slice values", func(t *testing.T) {
		assert.Equal(t, 0.0, ToFloat64([]byte("invalid")))
		assert.Equal(t, 0.0, ToFloat64([]byte("")))
		assert.Equal(t, 0.0, ToFloat64([]byte("abc")))
	})

	t.Run("unsupported types", func(t *testing.T) {
		assert.Equal(t, 0.0, ToFloat64(nil))
		assert.Equal(t, 0.0, ToFloat64(struct{}{}))
		assert.Equal(t, 0.0, ToFloat64(map[string]int{}))
		assert.Equal(t, 0.0, ToFloat64([]int{1, 2, 3}))
		assert.Equal(t, 0.0, ToFloat64(func() {}))
	})
}

// TestToInt64 测试 ToInt64 函数
func TestToInt64(t *testing.T) {
	t.Run("bool values", func(t *testing.T) {
		assert.Equal(t, int64(1), ToInt64(true))
		assert.Equal(t, int64(0), ToInt64(false))
	})

	t.Run("integer types", func(t *testing.T) {
		assert.Equal(t, int64(42), ToInt64(int(42)))
		assert.Equal(t, int64(127), ToInt64(int8(127)))
		assert.Equal(t, int64(32767), ToInt64(int16(32767)))
		assert.Equal(t, int64(100), ToInt64(int32(100)))
		assert.Equal(t, int64(1000), ToInt64(int64(1000)))
		assert.Equal(t, int64(42), ToInt64(uint(42)))
		assert.Equal(t, int64(255), ToInt64(uint8(255)))
		assert.Equal(t, int64(65535), ToInt64(uint16(65535)))
		assert.Equal(t, int64(100), ToInt64(uint32(100)))
		assert.Equal(t, int64(1000), ToInt64(uint64(1000)))
	})

	t.Run("duration values", func(t *testing.T) {
		assert.Equal(t, int64(1000000000), ToInt64(time.Second))
		assert.Equal(t, int64(1000000), ToInt64(time.Millisecond))
	})

	t.Run("float values", func(t *testing.T) {
		assert.Equal(t, int64(3), ToInt64(float32(3.14)))
		assert.Equal(t, int64(3), ToInt64(float64(3.14)))
		assert.Equal(t, int64(-3), ToInt64(float64(-3.14)))
	})

	t.Run("string values", func(t *testing.T) {
		assert.Equal(t, int64(42), ToInt64("42"))
		assert.Equal(t, int64(-42), ToInt64("-42"))
		assert.Equal(t, int64(0), ToInt64("0"))
	})

	t.Run("invalid string values", func(t *testing.T) {
		assert.Equal(t, int64(0), ToInt64("invalid"))
		assert.Equal(t, int64(0), ToInt64(""))
		assert.Equal(t, int64(0), ToInt64("3.14"))
	})

	t.Run("byte slice values", func(t *testing.T) {
		assert.Equal(t, int64(42), ToInt64([]byte("42")))
		assert.Equal(t, int64(-42), ToInt64([]byte("-42")))
		assert.Equal(t, int64(0), ToInt64([]byte("invalid")))
		assert.Equal(t, int64(0), ToInt64([]byte("")))
		assert.Equal(t, int64(0), ToInt64([]byte("3.14")))
	})

	t.Run("unsupported types", func(t *testing.T) {
		assert.Equal(t, int64(0), ToInt64(nil))
		assert.Equal(t, int64(0), ToInt64(struct{}{}))
	})
}

// TestToInt32 测试 ToInt32 函数
func TestToInt32(t *testing.T) {
	t.Run("bool values", func(t *testing.T) {
		assert.Equal(t, int32(1), ToInt32(true))
		assert.Equal(t, int32(0), ToInt32(false))
	})

	t.Run("integer types", func(t *testing.T) {
		assert.Equal(t, int32(42), ToInt32(int(42)))
		assert.Equal(t, int32(127), ToInt32(int8(127)))
		assert.Equal(t, int32(32767), ToInt32(int16(32767)))
		assert.Equal(t, int32(100), ToInt32(int32(100)))
		assert.Equal(t, int32(1000), ToInt32(int64(1000)))
		assert.Equal(t, int32(42), ToInt32(uint(42)))
		assert.Equal(t, int32(255), ToInt32(uint8(255)))
		assert.Equal(t, int32(65535), ToInt32(uint16(65535)))
		assert.Equal(t, int32(100), ToInt32(uint32(100)))
		assert.Equal(t, int32(1000), ToInt32(uint64(1000)))
	})

	t.Run("float values", func(t *testing.T) {
		assert.Equal(t, int32(3), ToInt32(float32(3.14)))
		assert.Equal(t, int32(3), ToInt32(float64(3.14)))
		assert.Equal(t, int32(-3), ToInt32(float64(-3.14)))
	})

	t.Run("string values", func(t *testing.T) {
		assert.Equal(t, int32(42), ToInt32("42"))
		assert.Equal(t, int32(0), ToInt32("-42")) // negative strings not supported
		assert.Equal(t, int32(0), ToInt32("0"))
	})

	t.Run("invalid string values", func(t *testing.T) {
		assert.Equal(t, int32(0), ToInt32("invalid"))
		assert.Equal(t, int32(0), ToInt32(""))
		assert.Equal(t, int32(0), ToInt32("3.14"))
	})

	t.Run("byte slice values", func(t *testing.T) {
		assert.Equal(t, int32(42), ToInt32([]byte("42")))
		assert.Equal(t, int32(0), ToInt32([]byte("-42"))) // negative strings not supported
	})

	t.Run("unsupported types", func(t *testing.T) {
		assert.Equal(t, int32(0), ToInt32(nil))
		assert.Equal(t, int32(0), ToInt32(struct{}{}))
	})
}

// TestToInt16 测试 ToInt16 函数
func TestToInt16(t *testing.T) {
	t.Run("bool values", func(t *testing.T) {
		assert.Equal(t, int16(1), ToInt16(true))
		assert.Equal(t, int16(0), ToInt16(false))
	})

	t.Run("integer types", func(t *testing.T) {
		assert.Equal(t, int16(42), ToInt16(int(42)))
		assert.Equal(t, int16(127), ToInt16(int8(127)))
		assert.Equal(t, int16(1000), ToInt16(int16(1000)))
		assert.Equal(t, int16(100), ToInt16(int32(100)))
		assert.Equal(t, int16(1000), ToInt16(int64(1000)))
		assert.Equal(t, int16(42), ToInt16(uint(42)))
		assert.Equal(t, int16(255), ToInt16(uint8(255)))
		assert.Equal(t, int16(1000), ToInt16(uint16(1000)))
		assert.Equal(t, int16(100), ToInt16(uint32(100)))
		assert.Equal(t, int16(1000), ToInt16(uint64(1000)))
	})

	t.Run("float values", func(t *testing.T) {
		assert.Equal(t, int16(3), ToInt16(float32(3.14)))
		assert.Equal(t, int16(3), ToInt16(float64(3.14)))
	})

	t.Run("string values", func(t *testing.T) {
		assert.Equal(t, int16(42), ToInt16("42"))
		assert.Equal(t, int16(0), ToInt16("-42")) // negative strings not supported due to ParseUint
		assert.Equal(t, int16(0), ToInt16(""))
		assert.Equal(t, int16(0), ToInt16("3.14"))
	})

	t.Run("byte slice values", func(t *testing.T) {
		assert.Equal(t, int16(42), ToInt16([]byte("42")))
		assert.Equal(t, int16(0), ToInt16([]byte("-42"))) // negative not supported
		assert.Equal(t, int16(0), ToInt16([]byte("")))
		assert.Equal(t, int16(0), ToInt16([]byte("invalid")))
	})

	t.Run("invalid values", func(t *testing.T) {
		assert.Equal(t, int16(0), ToInt16("invalid"))
		assert.Equal(t, int16(0), ToInt16(nil))
		assert.Equal(t, int16(0), ToInt16(struct{}{}))
	})
}

// TestToInt8 测试 ToInt8 函数
func TestToInt8(t *testing.T) {
	t.Run("bool values", func(t *testing.T) {
		assert.Equal(t, int8(1), ToInt8(true))
		assert.Equal(t, int8(0), ToInt8(false))
	})

	t.Run("integer types", func(t *testing.T) {
		assert.Equal(t, int8(42), ToInt8(int(42)))
		assert.Equal(t, int8(127), ToInt8(int8(127)))
		assert.Equal(t, int8(100), ToInt8(int16(100)))
		assert.Equal(t, int8(100), ToInt8(int32(100)))
		assert.Equal(t, int8(100), ToInt8(int64(100)))
		assert.Equal(t, int8(42), ToInt8(uint(42)))
		assert.Equal(t, int8(100), ToInt8(uint8(100)))
		assert.Equal(t, int8(100), ToInt8(uint16(100)))
		assert.Equal(t, int8(100), ToInt8(uint32(100)))
		assert.Equal(t, int8(100), ToInt8(uint64(100)))
	})

	t.Run("float values", func(t *testing.T) {
		assert.Equal(t, int8(3), ToInt8(float32(3.14)))
		assert.Equal(t, int8(3), ToInt8(float64(3.14)))
	})

	t.Run("string values", func(t *testing.T) {
		assert.Equal(t, int8(42), ToInt8("42"))
		assert.Equal(t, int8(0), ToInt8("-42")) // negative strings not supported due to ParseUint
		assert.Equal(t, int8(0), ToInt8(""))
		assert.Equal(t, int8(0), ToInt8("3.14"))
	})

	t.Run("byte slice values", func(t *testing.T) {
		assert.Equal(t, int8(42), ToInt8([]byte("42")))
		assert.Equal(t, int8(0), ToInt8([]byte("-42"))) // negative not supported
		assert.Equal(t, int8(0), ToInt8([]byte("")))
		assert.Equal(t, int8(0), ToInt8([]byte("invalid")))
	})

	t.Run("invalid values", func(t *testing.T) {
		assert.Equal(t, int8(0), ToInt8("invalid"))
		assert.Equal(t, int8(0), ToInt8(nil))
		assert.Equal(t, int8(0), ToInt8(struct{}{}))
	})
}

// TestToUint 测试 ToUint 函数的各种类型转换
func TestToUint(t *testing.T) {
	tests := []struct {
		name  string      // 测试用例名称，描述具体的测试场景
		input interface{} // 输入参数，支持多种类型
		want  uint        // 期望的输出结果
	}{
		// 布尔值转换测试
		{"bool true", true, 1},   // true 应该转换为 1
		{"bool false", false, 0}, // false 应该转换为 0

		// 整数类型转换测试
		{"int positive", 42, 42},                   // 正整数直接转换
		{"int negative", -1, 18446744073709551615}, // 负整数转换为对应的无符号值
		{"uint", uint(100), 100},                   // uint 类型直接返回

		// 浮点数转换测试（截断小数部分）
		{"float positive", 3.14, 3}, // 正浮点数截断小数

		// 字符串转换测试
		{"string valid", "123", 123}, // 有效数字字符串
		{"string invalid", "abc", 0}, // 无效字符串返回 0

		// 字节切片转换测试
		{"byte slice valid", []byte("456"), 456}, // 有效数字字节切片
		{"byte slice invalid", []byte("xyz"), 0}, // 无效字节切片返回 0

		// 不支持的类型测试（返回 0）
		{"slice", []int{1, 2}, 0},          // 切片类型不支持
		{"map", map[string]int{"a": 1}, 0}, // 映射类型不支持
		{"nil pointer", (*int)(nil), 0},    // nil 指针返回 0

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
		{"min int", -1 << 63, 9223372036854775808},  // 最小 int64 值
		{"max uint", ^uint(0), ^uint(0)},            // 最大 uint 值

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

// TestToUint16 测试 ToUint16 函数的各种输入转换场景
func TestToUint16(t *testing.T) {
	// 定义测试用例结构体，包含测试名称、输入值和期望输出
	tests := []struct {
		name  string      // 测试用例名称
		input interface{} // 输入值
		want  uint16      // 期望的输出结果
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

// TestToUint8 测试 ToUint8 函数的功能和边界情况
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
		{"string float", "3.14", 0},

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

// TestToUint32 测试 ToUint32 函数
func TestToUint32(t *testing.T) {
	tests := []struct {
		name  string      // 测试用例名称
		input interface{} // 输入参数
		want  uint32      // 期望结果
	}{
		{"bool true", true, 1},
		{"bool false", false, 0},
		{"int positive", 123456, 123456},
		{"int negative", -1, 4294967295},
		{"uint32 max", uint32(4294967295), 4294967295},
		{"string valid", "4294967295", 4294967295},
		{"string invalid", "abc", 0},
		{"byte slice valid", []byte("123456"), 123456},
		{"nil", nil, 0},
		{"struct", struct{}{}, 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ToUint32(tt.input); got != tt.want {
				t.Errorf("ToUint32() = %v, want %v", got, tt.want)
			}
		})
	}
}

// TestToUint64 测试 ToUint64 函数
func TestToUint64(t *testing.T) {
	tests := []struct {
		name  string      // 测试用例名称
		input interface{} // 输入参数
		want  uint64      // 期望结果
	}{
		{"bool true", true, 1},
		{"bool false", false, 0},
		{"int positive", 123456, 123456},
		{"int negative", -1, 18446744073709551615},
		{"uint64 max", uint64(18446744073709551615), 18446744073709551615},
		{"string valid", "18446744073709551615", 18446744073709551615},
		{"string invalid", "abc", 0},
		{"byte slice valid", []byte("18446744073709551615"), 18446744073709551615},
		{"nil", nil, 0},
		{"struct", struct{}{}, 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ToUint64(tt.input); got != tt.want {
				t.Errorf("ToUint64() = %v, want %v", got, tt.want)
			}
		})
	}
}

// ================================
// 字节转换测试组
// ================================

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

// ================================
// 复合类型转换测试组
// ================================

// TestToMap 测试 ToMap 函数
func TestToMap(t *testing.T) {
	t.Run("json byte slice", func(t *testing.T) {
		input := []byte(`{"name":"John","age":30}`)
		result := ToMap(input)
		expected := map[string]interface{}{
			"name": "John",
			"age":  float64(30), // JSON numbers are float64
		}
		assert.Equal(t, expected, result)
	})

	t.Run("json string", func(t *testing.T) {
		input := `{"city":"New York","population":8000000}`
		result := ToMap(input)
		expected := map[string]interface{}{
			"city":       "New York",
			"population": float64(8000000),
		}
		assert.Equal(t, expected, result)
	})

	t.Run("invalid json byte slice", func(t *testing.T) {
		input := []byte(`invalid json`)
		result := ToMap(input)
		expected := map[string]interface{}{} // fallback to ToMapStringAny
		assert.Equal(t, expected, result)
	})

	t.Run("invalid json string", func(t *testing.T) {
		input := "invalid json"
		result := ToMap(input)
		expected := map[string]interface{}{} // fallback to ToMapStringAny
		assert.Equal(t, expected, result)
	})

	t.Run("empty json", func(t *testing.T) {
		input := "{}"
		result := ToMap(input)
		expected := map[string]interface{}{}
		assert.Equal(t, expected, result)
	})

	t.Run("nil input", func(t *testing.T) {
		result := ToMap(nil)
		assert.Nil(t, result)
	})

	t.Run("map input", func(t *testing.T) {
		input := map[int]string{1: "one", 2: "two"}
		result := ToMap(input)
		expected := map[string]interface{}{
			"1": "one",
			"2": "two",
		}
		assert.Equal(t, expected, result)
	})

	t.Run("non-map input", func(t *testing.T) {
		input := 42
		result := ToMap(input)
		expected := map[string]interface{}{}
		assert.Equal(t, expected, result)
	})
}

// TestToMapStringAny 测试 ToMapStringAny 函数
func TestToMapStringAny(t *testing.T) {
	t.Run("nil input", func(t *testing.T) {
		result := ToMapStringAny(nil)
		assert.Nil(t, result)
	})

	t.Run("int key map", func(t *testing.T) {
		input := map[int]string{1: "one", 2: "two", 3: "three"}
		result := ToMapStringAny(input)
		expected := map[string]interface{}{
			"1": "one",
			"2": "two",
			"3": "three",
		}
		assert.Equal(t, expected, result)
	})

	t.Run("string key map", func(t *testing.T) {
		input := map[string]int{"one": 1, "two": 2}
		result := ToMapStringAny(input)
		expected := map[string]interface{}{
			"one": 1,
			"two": 2,
		}
		assert.Equal(t, expected, result)
	})

	t.Run("mixed type map", func(t *testing.T) {
		input := map[interface{}]interface{}{
			"key1": "value1",
			42:     "value2",
			true:   "value3",
		}
		result := ToMapStringAny(input)
		// 检查结果是否包含所有期望的键值对，不依赖遍历顺序
		assert.Len(t, result, 3)
		assert.Equal(t, "value1", result["key1"])
		assert.Equal(t, "value2", result["42"])
		// 需要检查 ToString(true) 的实际返回值
		// 可能是 "1" 而不是 "true"
		hasTrue := result["true"] == "value3"
		hasOne := result["1"] == "value3"
		assert.True(t, hasTrue || hasOne, "Expected either 'true' or '1' key to map to 'value3'")
	})

	t.Run("empty map", func(t *testing.T) {
		input := map[string]string{}
		result := ToMapStringAny(input)
		expected := map[string]interface{}{}
		assert.Equal(t, expected, result)
	})

	t.Run("non-map input", func(t *testing.T) {
		input := "not a map"
		result := ToMapStringAny(input)
		expected := map[string]interface{}{}
		assert.Equal(t, expected, result)
	})

	t.Run("slice input", func(t *testing.T) {
		input := []int{1, 2, 3}
		result := ToMapStringAny(input)
		expected := map[string]interface{}{}
		assert.Equal(t, expected, result)
	})

	t.Run("struct input", func(t *testing.T) {
		input := struct{ Name string }{"test"}
		result := ToMapStringAny(input)
		expected := map[string]interface{}{}
		assert.Equal(t, expected, result)
	})
}

// TestToSlice tests ToSlice functions
func TestToFloat64Slice(t *testing.T) {
	t.Run("nil_input", func(t *testing.T) {
		result := ToFloat64Slice(nil)
		assert.Nil(t, result)
	})

	t.Run("bool_slice", func(t *testing.T) {
		input := []bool{true, false, true}
		result := ToFloat64Slice(input)
		expected := []float64{1.0, 0.0, 1.0}
		assert.Equal(t, expected, result)
	})

	t.Run("int_slice", func(t *testing.T) {
		input := []int{1, 2, 3}
		result := ToFloat64Slice(input)
		expected := []float64{1.0, 2.0, 3.0}
		assert.Equal(t, expected, result)
	})

	t.Run("string_slice", func(t *testing.T) {
		input := []string{"1.5", "2.7", "invalid"}
		result := ToFloat64Slice(input)
		expected := []float64{1.5, 2.7, 0.0}
		assert.Equal(t, expected, result)
	})

	t.Run("interface_slice", func(t *testing.T) {
		input := []interface{}{1, "2.5", true}
		result := ToFloat64Slice(input)
		expected := []float64{1.0, 2.5, 1.0}
		assert.Equal(t, expected, result)
	})

	t.Run("unsupported_type", func(t *testing.T) {
		input := "not a slice"
		result := ToFloat64Slice(input)
		expected := []float64{}
		assert.Equal(t, expected, result)
	})
}

func TestToInt64Slice(t *testing.T) {
	t.Run("bool_slice", func(t *testing.T) {
		input := []bool{true, false, true}
		result := ToInt64Slice(input)
		expected := []int64{1, 0, 1}
		assert.Equal(t, expected, result)
	})

	t.Run("int_slice", func(t *testing.T) {
		input := []int{1, 2, 3}
		result := ToInt64Slice(input)
		expected := []int64{1, 2, 3}
		assert.Equal(t, expected, result)
	})

	t.Run("float64_slice", func(t *testing.T) {
		input := []float64{1.5, 2.7, 3.9}
		result := ToInt64Slice(input)
		expected := []int64{1, 2, 3} // truncated
		assert.Equal(t, expected, result)
	})

	t.Run("string_slice", func(t *testing.T) {
		input := []string{"42", "-10", "invalid"}
		result := ToInt64Slice(input)
		expected := []int64{42, -10, 0}
		assert.Equal(t, expected, result)
	})

	t.Run("interface_slice", func(t *testing.T) {
		input := []interface{}{1, "42", true}
		result := ToInt64Slice(input)
		expected := []int64{1, 42, 1}
		assert.Equal(t, expected, result)
	})

	t.Run("unsupported_type", func(t *testing.T) {
		input := "not a slice"
		result := ToInt64Slice(input)
		expected := []int64{}
		assert.Equal(t, expected, result)
	})
}

// TestToPtr 测试 ToPtr 函数
func TestToPtr(t *testing.T) {
	t.Run("int", func(t *testing.T) {
		v := 42
		p := ToPtr(v)
		assert.NotNil(t, p)
		assert.Equal(t, v, *p)
	})

	t.Run("string", func(t *testing.T) {
		v := "hello"
		p := ToPtr(v)
		assert.NotNil(t, p)
		assert.Equal(t, v, *p)
	})

	t.Run("bool", func(t *testing.T) {
		v := true
		p := ToPtr(v)
		assert.NotNil(t, p)
		assert.Equal(t, v, *p)
	})

	t.Run("zero-value-int", func(t *testing.T) {
		v := 0
		p := ToPtr(v)
		assert.NotNil(t, p)
		assert.Equal(t, v, *p)
	})

	t.Run("struct", func(t *testing.T) {
		type myStruct struct {
			Field int
		}
		v := myStruct{Field: 100}
		p := ToPtr(v)
		assert.NotNil(t, p)
		assert.Equal(t, v, *p)
	})

	t.Run("slice", func(t *testing.T) {
		v := []int{1, 2, 3}
		p := ToPtr(v)
		assert.NotNil(t, p)
		assert.Equal(t, v, *p)
	})
}

// ================================
// Map变种转换测试组 (来自to_map_variants_test.go)
// ================================

func TestToMapInt32String(t *testing.T) {
	t.Run("basic conversion", func(t *testing.T) {
		input := map[int32]string{1: "one", 2: "two", 3: "three"}
		result := ToMapInt32String(input)
		expected := map[int32]string{1: "one", 2: "two", 3: "three"}
		assert.Equal(t, expected, result)
	})

	t.Run("empty map", func(t *testing.T) {
		input := map[int32]string{}
		result := ToMapInt32String(input)
		expected := map[int32]string{}
		assert.Equal(t, expected, result)
	})

	t.Run("non-map input", func(t *testing.T) {
		input := "not a map"
		result := ToMapInt32String(input)
		expected := map[int32]string{}
		assert.Equal(t, expected, result)
	})

	t.Run("nil input", func(t *testing.T) {
		result := ToMapInt32String(nil)
		expected := map[int32]string{}
		assert.Equal(t, expected, result)
	})
}

func TestToMapInt64String(t *testing.T) {
	t.Run("basic conversion", func(t *testing.T) {
		input := map[int64]string{100: "hundred", 200: "two hundred"}
		result := ToMapInt64String(input)
		expected := map[int64]string{100: "hundred", 200: "two hundred"}
		assert.Equal(t, expected, result)
	})

	t.Run("empty map", func(t *testing.T) {
		input := map[int64]string{}
		result := ToMapInt64String(input)
		expected := map[int64]string{}
		assert.Equal(t, expected, result)
	})

	t.Run("non-map input", func(t *testing.T) {
		input := 42
		result := ToMapInt64String(input)
		expected := map[int64]string{}
		assert.Equal(t, expected, result)
	})
}

func TestToMapStringString(t *testing.T) {
	t.Run("basic conversion", func(t *testing.T) {
		input := map[string]string{"key1": "value1", "key2": "value2"}
		result := ToMapStringString(input)
		expected := map[string]string{"key1": "value1", "key2": "value2"}
		assert.Equal(t, expected, result)
	})

	t.Run("empty map", func(t *testing.T) {
		input := map[string]string{}
		result := ToMapStringString(input)
		expected := map[string]string{}
		assert.Equal(t, expected, result)
	})

	t.Run("non-map input", func(t *testing.T) {
		input := []string{"not", "a", "map"}
		result := ToMapStringString(input)
		expected := map[string]string{}
		assert.Equal(t, expected, result)
	})
}

func TestToMapStringInt64(t *testing.T) {
	t.Run("basic conversion", func(t *testing.T) {
		input := map[string]int64{"count": 100, "total": 500}
		result := ToMapStringInt64(input)
		expected := map[string]int64{"count": 100, "total": 500}
		assert.Equal(t, expected, result)
	})

	t.Run("empty map", func(t *testing.T) {
		input := map[string]int64{}
		result := ToMapStringInt64(input)
		expected := map[string]int64{}
		assert.Equal(t, expected, result)
	})

	t.Run("non-map input", func(t *testing.T) {
		input := struct{}{}
		result := ToMapStringInt64(input)
		expected := map[string]int64{}
		assert.Equal(t, expected, result)
	})
}

func TestToMapStringArrayString(t *testing.T) {
	t.Run("basic conversion", func(t *testing.T) {
		input := map[string][]string{
			"colors": {"red", "blue", "green"},
			"fruits": {"apple", "banana"},
		}
		result := ToMapStringArrayString(input)
		expected := map[string][]string{
			"colors": {"red", "blue", "green"},
			"fruits": {"apple", "banana"},
		}
		assert.Equal(t, expected, result)
	})

	t.Run("empty map", func(t *testing.T) {
		input := map[string][]string{}
		result := ToMapStringArrayString(input)
		expected := map[string][]string{}
		assert.Equal(t, expected, result)
	})

	t.Run("nil input", func(t *testing.T) {
		result := ToMapStringArrayString(nil)
		assert.Nil(t, result)
	})

	t.Run("map with different key types", func(t *testing.T) {
		input := map[int]string{
			1: "one,two",
			2: "three",
		}
		result := ToMapStringArrayString(input)
		expected := map[string][]string{
			"1": {"one", "two"},
			"2": {"three"},
		}
		assert.Equal(t, expected, result)
	})

	t.Run("non-map input should panic", func(t *testing.T) {
		input := []string{"not", "a", "map"}
		assert.Panics(t, func() {
			ToMapStringArrayString(input)
		})
	})
}