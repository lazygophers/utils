package candy

import (
	"errors"
	"math"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

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