// Package min 提供 Min 函数的测试
package candy

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestMin 测试 Min 函数
func TestMin(t *testing.T) {
	// 测试整数类型切片的最小值查找
	t.Run("整数类型切片", func(t *testing.T) {
		tests := []struct {
			name string
			give []int
			want int
		}{
			{"正整数切片", []int{1, 2, 3, 4, 5}, 1},
			{"负整数切片", []int{-1, -2, -3, -4, -5}, -5},
			{"混合正负整数切片", []int{-5, 0, 5, -10, 10}, -10},
			{"单个正整数", []int{42}, 42},
			{"单个负整数", []int{-42}, -42},
			{"重复值切片", []int{3, 3, 3, 3, 3}, 3},
			{"大数切片", []int{1, 999999999, 999999998}, 1},
		}

		for _, tt := range tests {
			tt := tt // 避免竞态
			t.Run(tt.name, func(t *testing.T) {
				t.Parallel()
				got := Min(tt.give)
				assert.Equal(t, tt.want, got, "Min() 的结果应与期望值相等")
			})
		}
	})

	// 测试不同整数类型
	t.Run("不同整数类型", func(t *testing.T) {
		// int8 类型
		int8Slice := []int8{1, 2, 3, 4, 5}
		assert.Equal(t, int8(1), Min(int8Slice))

		// int16 类型
		int16Slice := []int16{1000, 2000, 3000}
		assert.Equal(t, int16(1000), Min(int16Slice))

		// int32 类型
		int32Slice := []int32{100000, 200000, 300000}
		assert.Equal(t, int32(100000), Min(int32Slice))

		// int64 类型
		int64Slice := []int64{10000000000, 20000000000, 30000000000}
		assert.Equal(t, int64(10000000000), Min(int64Slice))

		// uint 类型
		uintSlice := []uint{1, 2, 3, 4, 5}
		assert.Equal(t, uint(1), Min(uintSlice))

		// uint8 类型
		uint8Slice := []uint8{10, 20, 30, 40, 50}
		assert.Equal(t, uint8(10), Min(uint8Slice))
	})

	// 测试浮点数类型切片的最小值查找
	t.Run("浮点数类型切片", func(t *testing.T) {
		tests := []struct {
			name string
			give []float64
			want float64
		}{
			{"正浮点数切片", []float64{1.1, 2.2, 3.3, 4.4, 5.5}, 1.1},
			{"负浮点数切片", []float64{-1.1, -2.2, -3.3, -4.4, -5.5}, -5.5},
			{"混合正负浮点数切片", []float64{-5.5, 0.0, 5.5, -10.1, 10.1}, -10.1},
			{"单个正浮点数", []float64{3.14159}, 3.14159},
			{"单个负浮点数", []float64{-3.14159}, -3.14159},
			{"重复值浮点数切片", []float64{2.718, 2.718, 2.718}, 2.718},
			{"科学计数法", []float64{1.23e-4, 5.67e8, 9.01e2}, 1.23e-4},
		}

		for _, tt := range tests {
			tt := tt // 避免竞态
			t.Run(tt.name, func(t *testing.T) {
				t.Parallel()
				got := Min(tt.give)
				assert.Equal(t, tt.want, got, "Min() 的结果应与期望值相等")
			})
		}
	})

	// 测试不同浮点数类型
	t.Run("不同浮点数类型", func(t *testing.T) {
		// float32 类型
		float32Slice := []float32{1.1, 2.2, 3.3}
		assert.Equal(t, float32(1.1), Min(float32Slice))
	})

	// 测试字符串类型切片的最小值查找
	t.Run("字符串类型切片", func(t *testing.T) {
		tests := []struct {
			name string
			give []string
			want string
		}{
			{"字母字符串切片", []string{"apple", "banana", "cherry", "date"}, "apple"},
			{"混合大小写字符串切片", []string{"Apple", "banana", "Cherry", "date"}, "Apple"},
			{"数字字符串切片", []string{"1", "2", "10", "20"}, "1"},
			{"特殊字符字符串切片", []string{"!", "@", "#", "$"}, "!"},
			{"中文字符串切片", []string{"苹果", "香蕉", "樱桃", "日期"}, "日期"},
			{"单个字符", []string{"a"}, "a"},
			{"重复字符串切片", []string{"hello", "hello", "hello"}, "hello"},
			{"空字符串和有效字符串", []string{"", "hello", "world"}, ""},
			{"Unicode 字符串", []string{"α", "β", "γ", "δ"}, "α"},
		}

		for _, tt := range tests {
			tt := tt // 避免竞态
			t.Run(tt.name, func(t *testing.T) {
				t.Parallel()
				got := Min(tt.give)
				assert.Equal(t, tt.want, got, "Min() 的结果应与期望值相等")
			})
		}
	})

	// 测试空切片情况
	t.Run("空切片情况", func(t *testing.T) {
		// 整数空切片
		assert.Equal(t, 0, Min([]int{}), "空整数切片应返回零值")

		// 浮点数空切片
		assert.Equal(t, 0.0, Min([]float64{}), "空浮点数切片应返回零值")

		// 字符串空切片
		assert.Equal(t, "", Min([]string{}), "空字符串切片应返回零值")

		// 其他类型的空切片
		assert.Equal(t, int8(0), Min([]int8{}))
		assert.Equal(t, uint16(0), Min([]uint16{}))
		assert.Equal(t, float32(0.0), Min([]float32{}))
	})

	// 测试单元素切片情况
	t.Run("单元素切片情况", func(t *testing.T) {
		tests := []struct {
			name string
			give interface{}
			want interface{}
		}{
			{"单元素整数切片", []int{42}, 42},
			{"单元素浮点数切片", []float64{3.14159}, 3.14159},
			{"单元素字符串切片", []string{"hello"}, "hello"},
			{"单元素负整数切片", []int{-42}, -42},
			{"单元素零值切片", []int{0}, 0},
		}

		for _, tt := range tests {
			tt := tt // 避免竞态
			t.Run(tt.name, func(t *testing.T) {
				t.Parallel()
				switch v := tt.give.(type) {
				case []int:
					got := Min(v)
					assert.Equal(t, tt.want, got)
				case []float64:
					got := Min(v)
					assert.Equal(t, tt.want, got)
				case []string:
					got := Min(v)
					assert.Equal(t, tt.want, got)
				}
			})
		}
	})

	// 测试重复值情况
	t.Run("重复值情况", func(t *testing.T) {
		tests := []struct {
			name string
			give interface{}
			want interface{}
		}{
			{"全部相同的整数", []int{5, 5, 5, 5, 5}, 5},
			{"全部相同的浮点数", []float64{2.5, 2.5, 2.5}, 2.5},
			{"全部相同的字符串", []string{"test", "test", "test"}, "test"},
			{"最小值重复出现", []int{1, 5, 2, 1, 3, 1}, 1},
			{"最大值重复出现", []int{10, 10, 10, 2, 3, 4}, 2},
			{"中间值重复出现", []int{5, 2, 3, 2, 3, 5}, 2},
		}

		for _, tt := range tests {
			tt := tt // 避免竞态
			t.Run(tt.name, func(t *testing.T) {
				t.Parallel()
				switch v := tt.give.(type) {
				case []int:
					got := Min(v)
					assert.Equal(t, tt.want, got)
				case []float64:
					got := Min(v)
					assert.Equal(t, tt.want, got)
				case []string:
					got := Min(v)
					assert.Equal(t, tt.want, got)
				}
			})
		}
	})

	// 验证正确返回最小值元素
	t.Run("验证正确返回最小值元素", func(t *testing.T) {
		// 测试是否返回的是原始切片中的元素，而不是计算得到的值
		originalSlice := []int{1, 2, 3, 4, 5}
		minValue := Min(originalSlice)

		// 验证返回值确实是最小值
		assert.Equal(t, 1, minValue)

		// 验证返回值存在于原切片中
		found := false
		for _, v := range originalSlice {
			if v == minValue {
				found = true
				break
			}
		}
		assert.True(t, found, "返回的最小值应该存在于原切片中")

		// 对于浮点数，考虑精度问题
		floatSlice := []float64{1.1, 2.2, 3.3, 4.4, 5.5}
		floatMin := Min(floatSlice)
		assert.InDelta(t, 1.1, floatMin, 1e-9, "浮点数最小值应该在精度范围内")
	})

	// 测试边界情况
	t.Run("边界情况", func(t *testing.T) {
		// 测试极大整数
		largeIntSlice := []int{1, math.MaxInt64 - 1, math.MaxInt64}
		assert.Equal(t, 1, Min(largeIntSlice))

		// 测试极小整数
		smallIntSlice := []int{-1, math.MinInt64 + 1, math.MinInt64}
		assert.Equal(t, math.MinInt64, Min(smallIntSlice))

		// 测试极大浮点数
		largeFloatSlice := []float64{1.0, math.MaxFloat64 - 1.0, math.MaxFloat64}
		assert.Equal(t, 1.0, Min(largeFloatSlice))

		// 测试极小浮点数
		smallFloatSlice := []float64{-1.0, -math.MaxFloat64, -2.0}
		assert.Equal(t, -math.MaxFloat64, Min(smallFloatSlice))

		// 测试无穷大
		infinitySlice := []float64{1.0, math.Inf(1), 2.0}
		assert.Equal(t, 1.0, Min(infinitySlice))

		// 测试负无穷大
		negativeInfinitySlice := []float64{-1.0, math.Inf(-1), -2.0}
		assert.Equal(t, math.Inf(-1), Min(negativeInfinitySlice))

		// 测试 NaN (NaN 在比较中总是返回 false，所以应该返回第一个元素 NaN)
		nanSlice := []float64{math.NaN(), 1.0, 2.0}
		assert.True(t, math.IsNaN(Min(nanSlice)), "NaN 比较应返回第一个元素 NaN")

		allNanSlice := []float64{math.NaN(), math.NaN(), math.NaN()}
		assert.True(t, math.IsNaN(Min(allNanSlice)), "全 NaN 切片应该返回 NaN")
	})
}

// BenchmarkMin 基准测试 Min 函数
func BenchmarkMin(b *testing.B) {
	// 基准测试整数类型
	b.Run("整数类型", func(b *testing.B) {
		slice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			Min(slice)
		}
	})

	// 基准测试浮点数类型
	b.Run("浮点数类型", func(b *testing.B) {
		slice := []float64{1.1, 2.2, 3.3, 4.4, 5.5, 6.6, 7.7, 8.8, 9.9, 10.0}
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			Min(slice)
		}
	})

	// 基准测试字符串类型
	b.Run("字符串类型", func(b *testing.B) {
		slice := []string{"apple", "banana", "cherry", "date", "elderberry", "fig", "grape"}
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			Min(slice)
		}
	})

	// 基准测试大切片
	b.Run("大切片", func(b *testing.B) {
		slice := make([]int, 1000)
		for i := range slice {
			slice[i] = i
		}
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			Min(slice)
		}
	})

	// 基准测试空切片
	b.Run("空切片", func(b *testing.B) {
		slice := []int{}
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			Min(slice)
		}
	})
}