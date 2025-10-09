// Package candy provides convenient utility functions
// abs_test.go tests the Abs function comprehensively
package candy

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestAbs tests the Abs function with various numeric types
// 测试 Abs 函数处理各种数值类型
func TestAbs(t *testing.T) {
	t.Run("SignedIntegers", func(t *testing.T) {
		// int type tests
		// int 类型测试
		t.Run("int", func(t *testing.T) {
			tests := []struct {
				name string
				give int
				want int
			}{
				{"positive int", 42, 42},
				{"negative int", -42, 42},
				{"zero", 0, 0},
				{"max int", math.MaxInt, math.MaxInt},
				{"large negative", -999999, 999999},
			}

			for _, tt := range tests {
				t.Run(tt.name, func(t *testing.T) {
					got := Abs(tt.give)
					assert.Equal(t, tt.want, got)
				})
			}
		})

		// int8 type tests
		// int8 类型测试
		t.Run("int8", func(t *testing.T) {
			tests := []struct {
				name string
				give int8
				want int8
			}{
				{"positive int8", 127, 127},
				{"negative int8", -128, -128}, // Note: -128 cannot be represented as positive int8
				{"zero int8", 0, 0},
				{"small negative", -1, 1},
			}

			for _, tt := range tests {
				t.Run(tt.name, func(t *testing.T) {
					got := Abs(tt.give)
					assert.Equal(t, tt.want, got)
				})
			}
		})

		// int16 type tests
		// int16 类型测试
		t.Run("int16", func(t *testing.T) {
			tests := []struct {
				name string
				give int16
				want int16
			}{
				{"positive int16", 32767, 32767},
				{"negative int16", -32768, -32768},
				{"zero int16", 0, 0},
				{"medium negative", -1000, 1000},
			}

			for _, tt := range tests {
				t.Run(tt.name, func(t *testing.T) {
					got := Abs(tt.give)
					assert.Equal(t, tt.want, got)
				})
			}
		})

		// int32 type tests
		// int32 类型测试
		t.Run("int32", func(t *testing.T) {
			tests := []struct {
				name string
				give int32
				want int32
			}{
				{"positive int32", 2147483647, 2147483647},
				{"negative int32", -2147483648, -2147483648},
				{"zero int32", 0, 0},
				{"large positive", 1000000, 1000000},
				{"large negative", -1000000, 1000000},
			}

			for _, tt := range tests {
				t.Run(tt.name, func(t *testing.T) {
					got := Abs(tt.give)
					assert.Equal(t, tt.want, got)
				})
			}
		})

		// int64 type tests
		// int64 类型测试
		t.Run("int64", func(t *testing.T) {
			tests := []struct {
				name string
				give int64
				want int64
			}{
				{"positive int64", 9223372036854775807, 9223372036854775807},
				{"negative int64", -9223372036854775808, -9223372036854775808},
				{"zero int64", 0, 0},
				{"very large positive", 10000000000, 10000000000},
				{"very large negative", -10000000000, 10000000000},
			}

			for _, tt := range tests {
				t.Run(tt.name, func(t *testing.T) {
					got := Abs(tt.give)
					assert.Equal(t, tt.want, got)
				})
			}
		})
	})

	t.Run("UnsignedIntegers", func(t *testing.T) {
		// uint type tests
		// uint 类型测试
		t.Run("uint", func(t *testing.T) {
			tests := []struct {
				name string
				give uint
				want uint
			}{
				{"positive uint", 42, 42},
				{"zero uint", 0, 0},
				{"large uint", 4294967295, 4294967295},
			}

			for _, tt := range tests {
				t.Run(tt.name, func(t *testing.T) {
					got := Abs(tt.give)
					assert.Equal(t, tt.want, got)
				})
			}
		})

		// uint8 type tests
		// uint8 类型测试
		t.Run("uint8", func(t *testing.T) {
			tests := []struct {
				name string
				give uint8
				want uint8
			}{
				{"max uint8", 255, 255},
				{"zero uint8", 0, 0},
				{"medium uint8", 128, 128},
			}

			for _, tt := range tests {
				t.Run(tt.name, func(t *testing.T) {
					got := Abs(tt.give)
					assert.Equal(t, tt.want, got)
				})
			}
		})

		// uint16 type tests
		// uint16 类型测试
		t.Run("uint16", func(t *testing.T) {
			tests := []struct {
				name string
				give uint16
				want uint16
			}{
				{"max uint16", 65535, 65535},
				{"zero uint16", 0, 0},
				{"medium uint16", 32768, 32768},
			}

			for _, tt := range tests {
				t.Run(tt.name, func(t *testing.T) {
					got := Abs(tt.give)
					assert.Equal(t, tt.want, got)
				})
			}
		})

		// uint32 type tests
		// uint32 类型测试
		t.Run("uint32", func(t *testing.T) {
			tests := []struct {
				name string
				give uint32
				want uint32
			}{
				{"max uint32", 4294967295, 4294967295},
				{"zero uint32", 0, 0},
				{"medium uint32", 2147483648, 2147483648},
			}

			for _, tt := range tests {
				t.Run(tt.name, func(t *testing.T) {
					got := Abs(tt.give)
					assert.Equal(t, tt.want, got)
				})
			}
		})

		// uint64 type tests
		// uint64 类型测试
		t.Run("uint64", func(t *testing.T) {
			tests := []struct {
				name string
				give uint64
				want uint64
			}{
				{"max uint64", 18446744073709551615, 18446744073709551615},
				{"zero uint64", 0, 0},
				{"large uint64", 10000000000000000, 10000000000000000},
			}

			for _, tt := range tests {
				t.Run(tt.name, func(t *testing.T) {
					got := Abs(tt.give)
					assert.Equal(t, tt.want, got)
				})
			}
		})
	})

	t.Run("FloatingPoint", func(t *testing.T) {
		// float32 type tests
		// float32 类型测试
		t.Run("float32", func(t *testing.T) {
			tests := []struct {
				name string
				give float32
				want float32
			}{
				{"positive float32", 3.14, 3.14},
				{"negative float32", -3.14, 3.14},
				{"zero float32", 0.0, 0.0},
				{"very small positive", 1.1754944e-38, 1.1754944e-38},
				{"very small negative", -1.1754944e-38, 1.1754944e-38},
				{"max float32", 3.4028235e38, 3.4028235e38},
			}

			for _, tt := range tests {
				t.Run(tt.name, func(t *testing.T) {
					got := Abs(tt.give)
					assert.Equal(t, tt.want, got)
				})
			}
		})

		// float64 type tests
		// float64 类型测试
		t.Run("float64", func(t *testing.T) {
			tests := []struct {
				name string
				give float64
				want float64
			}{
				{"positive float64", 3.141592653589793, 3.141592653589793},
				{"negative float64", -3.141592653589793, 3.141592653589793},
				{"zero float64", 0.0, 0.0},
				{"very small positive", 2.2250738585072014e-308, 2.2250738585072014e-308},
				{"very small negative", -2.2250738585072014e-308, 2.2250738585072014e-308},
				{"max float64", 1.7976931348623157e308, 1.7976931348623157e308},
				{"negative infinity", math.Inf(-1), math.Inf(1)},
				{"positive infinity", math.Inf(1), math.Inf(1)},
			}

			for _, tt := range tests {
				t.Run(tt.name, func(t *testing.T) {
					got := Abs(tt.give)
					assert.Equal(t, tt.want, got)
				})
			}
		})

		// Special float64 cases
		// float64 特殊情况
		t.Run("SpecialCases", func(t *testing.T) {
			t.Run("NaN", func(t *testing.T) {
				result := Abs(math.NaN())
				// NaN < 0 is false, so Abs(NaN) returns NaN
				assert.True(t, math.IsNaN(result), "Abs(NaN) should return NaN")
			})

			t.Run("PositiveInfinity", func(t *testing.T) {
				result := Abs(math.Inf(1))
				assert.Equal(t, math.Inf(1), result)
			})

			t.Run("NegativeInfinity", func(t *testing.T) {
				result := Abs(math.Inf(-1))
				assert.Equal(t, math.Inf(1), result)
			})
		})
	})
}

// BenchmarkAbs benchmarks the Abs function with different numeric types
// 对 Abs 函数进行不同数值类型的性能基准测试
func BenchmarkAbs(b *testing.B) {
	b.Run("int", func(b *testing.B) {
		val := -42
		for i := 0; i < b.N; i++ {
			_ = Abs(val)
		}
	})

	b.Run("int8", func(b *testing.B) {
		val := int8(-42)
		for i := 0; i < b.N; i++ {
			_ = Abs(val)
		}
	})

	b.Run("int16", func(b *testing.B) {
		val := int16(-4200)
		for i := 0; i < b.N; i++ {
			_ = Abs(val)
		}
	})

	b.Run("int32", func(b *testing.B) {
		val := int32(-42000000)
		for i := 0; i < b.N; i++ {
			_ = Abs(val)
		}
	})

	b.Run("int64", func(b *testing.B) {
		val := int64(-4200000000)
		for i := 0; i < b.N; i++ {
			_ = Abs(val)
		}
	})

	b.Run("float32", func(b *testing.B) {
		val := float32(-3.14)
		for i := 0; i < b.N; i++ {
			_ = Abs(val)
		}
	})

	b.Run("float64", func(b *testing.B) {
		val := -3.141592653589793
		for i := 0; i < b.N; i++ {
			_ = Abs(val)
		}
	})

	b.Run("uint", func(b *testing.B) {
		val := uint(42)
		for i := 0; i < b.N; i++ {
			_ = Abs(val)
		}
	})

	b.Run("uint64", func(b *testing.B) {
		val := uint64(4200000000)
		for i := 0; i < b.N; i++ {
			_ = Abs(val)
		}
	})
}
