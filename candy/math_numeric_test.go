// Package candy 提供便捷的语法糖函数
package candy

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// TestMathNumericOperations 数学运算操作测试
func TestMathNumericOperations(t *testing.T) {
	// 基础运算组 - 单元素运算
	t.Run("BasicOperations", func(t *testing.T) {
		// 绝对值测试
		t.Run("Abs", func(t *testing.T) {
			tests := []struct {
				name string
				give interface{}
				want interface{}
			}{
				// 整数类型测试
				{
					name: "正整数绝对值等于自身",
					give: int(42),
					want: int(42),
				},
				{
					name: "负整数绝对值等于其相反数",
					give: int(-42),
					want: int(42),
				},
				{
					name: "零的绝对值等于零",
					give: int(0),
					want: int(0),
				},
				{
					name: "最大负整数的绝对值",
					give: int(-2147483648),
					want: int(2147483648),
				},
				{
					name: "最大正整数的绝对值",
					give: int(2147483647),
					want: int(2147483647),
				},
				// int8 类型测试
				{
					name: "int8 正数",
					give: int8(127),
					want: int8(127),
				},
				{
					name: "int8 负数",
					give: int8(-128),
					want: int8(-128),
				},
				// int16 类型测试
				{
					name: "int16 正数",
					give: int16(32767),
					want: int16(32767),
				},
				{
					name: "int16 负数",
					give: int16(-32768),
					want: int16(-32768),
				},
				// int32 类型测试
				{
					name: "int32 正数",
					give: int32(2147483647),
					want: int32(2147483647),
				},
				{
					name: "int32 负数",
					give: int32(-2147483648),
					want: int32(-2147483648),
				},
				// int64 类型测试
				{
					name: "int64 正数",
					give: int64(9223372036854775807),
					want: int64(9223372036854775807),
				},
				{
					name: "int64 负数",
					give: int64(-9223372036854775808),
					want: int64(-9223372036854775808),
				},
				// uint 类型测试
				{
					name: "uint 正数",
					give: uint(42),
					want: uint(42),
				},
				{
					name: "uint 零",
					give: uint(0),
					want: uint(0),
				},
				// uint8 类型测试
				{
					name: "uint8 最大值",
					give: uint8(255),
					want: uint8(255),
				},
				// float32 类型测试
				{
					name: "float32 正数",
					give: float32(3.14),
					want: float32(3.14),
				},
				{
					name: "float32 负数",
					give: float32(-3.14),
					want: float32(3.14),
				},
				{
					name: "float32 零",
					give: float32(0.0),
					want: float32(0.0),
				},
				{
					name: "float32 极小正数",
					give: float32(1.1754944e-38),
					want: float32(1.1754944e-38),
				},
				{
					name: "float32 极大正数",
					give: float32(3.4028235e38),
					want: float32(3.4028235e38),
				},
				// float64 类型测试
				{
					name: "float64 正数",
					give: float64(3.141592653589793),
					want: float64(3.141592653589793),
				},
				{
					name: "float64 负数",
					give: float64(-3.141592653589793),
					want: float64(3.141592653589793),
				},
				{
					name: "float64 零",
					give: float64(0.0),
					want: float64(0.0),
				},
				{
					name: "float64 极小正数",
					give: float64(2.2250738585072014e-308),
					want: float64(2.2250738585072014e-308),
				},
				{
					name: "float64 极大正数",
					give: float64(1.7976931348623157e308),
					want: float64(1.7976931348623157e308),
				},
			}

			for _, tt := range tests {
				tt := tt // 避免竞态
				t.Run(tt.name, func(t *testing.T) {
					switch give := tt.give.(type) {
					case int:
						got := Abs(give)
						assert.Equal(t, tt.want, got, "Abs() 的结果应与期望值相等")
					case int8:
						got := Abs(give)
						assert.Equal(t, tt.want, got, "Abs() 的结果应与期望值相等")
					case int16:
						got := Abs(give)
						assert.Equal(t, tt.want, got, "Abs() 的结果应与期望值相等")
					case int32:
						got := Abs(give)
						assert.Equal(t, tt.want, got, "Abs() 的结果应与期望值相等")
					case int64:
						got := Abs(give)
						assert.Equal(t, tt.want, got, "Abs() 的结果应与期望值相等")
					case uint:
						got := Abs(give)
						assert.Equal(t, tt.want, got, "Abs() 的结果应与期望值相等")
					case uint8:
						got := Abs(give)
						assert.Equal(t, tt.want, got, "Abs() 的结果应与期望值相等")
					case float32:
						got := Abs(give)
						assert.Equal(t, tt.want, got, "Abs() 的结果应与期望值相等")
					case float64:
						got := Abs(give)
						assert.Equal(t, tt.want, got, "Abs() 的结果应与期望值相等")
					}
				})
			}
		})

		// 平方根测试
		t.Run("Sqrt", func(t *testing.T) {
			tests := []struct {
				name string
				give interface{}
				want interface{}
			}{
				// 整数类型测试
				{
					name: "完全平方数",
					give: int(16),
					want: int(4),
				},
				{
					name: "非完全平方数",
					give: int(2),
					want: int(1),
				},
				{
					name: "零的平方根",
					give: int(0),
					want: int(0),
				},
				{
					name: "一的平方根",
					give: int(1),
					want: int(1),
				},
				// float32 类型测试
				{
					name: "float32 完全平方数",
					give: float32(16.0),
					want: float32(4.0),
				},
				{
					name: "float32 非完全平方数",
					give: float32(2.0),
					want: float32(1.4142135),
				},
				{
					name: "float32 小数平方根",
					give: float32(0.25),
					want: float32(0.5),
				},
				{
					name: "float32 零",
					give: float32(0.0),
					want: float32(0.0),
				},
				// float64 类型测试
				{
					name: "float64 完全平方数",
					give: float64(144.0),
					want: float64(12.0),
				},
				{
					name: "float64 非完全平方数",
					give: float64(3.0),
					want: float64(1.7320508075688772),
				},
				{
					name: "float64 小数平方根",
					give: float64(0.81),
					want: float64(0.9),
				},
				{
					name: "float64 零",
					give: float64(0.0),
					want: float64(0.0),
				},
				// 边界情况
				{
					name: "极大数的平方根",
					give: float64(1e20),
					want: float64(1e10),
				},
				{
					name: "极小正数的平方根",
					give: float64(1e-20),
					want: float64(1e-10),
				},
			}

			for _, tt := range tests {
				tt := tt // 避免竞态
				t.Run(tt.name, func(t *testing.T) {
					switch give := tt.give.(type) {
					case int:
						got := Sqrt(give)
						assert.Equal(t, tt.want, got, "Sqrt() 的结果应与期望值相等")
					case float32:
						got := Sqrt(give)
						assert.InDelta(t, tt.want, got, 1e-6, "Sqrt() 的结果应在允许误差范围内")
					case float64:
						got := Sqrt(give)
						assert.InDelta(t, tt.want, got, 1e-12, "Sqrt() 的结果应在允许误差范围内")
					}
				})
			}
		})

		// 立方根测试
		t.Run("Cbrt", func(t *testing.T) {
			tests := []struct {
				name string
				give interface{}
				want interface{}
			}{
				// 整数类型测试
				{
					name: "完全立方数",
					give: int(27),
					want: int(3),
				},
				{
					name: "负的完全立方数",
					give: int(-8),
					want: int(-2),
				},
				{
					name: "零的立方根",
					give: int(0),
					want: int(0),
				},
				{
					name: "一的立方根",
					give: int(1),
					want: int(1),
				},
				// float32 类型测试
				{
					name: "float32 完全立方数",
					give: float32(8.0),
					want: float32(2.0),
				},
				{
					name: "float32 负的完全立方数",
					give: float32(-27.0),
					want: float32(-3.0),
				},
				{
					name: "float32 非完全立方数",
					give: float32(10.0),
					want: float32(2.1544347),
				},
				{
					name: "float32 零",
					give: float32(0.0),
					want: float32(0.0),
				},
				// float64 类型测试
				{
					name: "float64 完全立方数",
					give: float64(125.0),
					want: float64(5.0),
				},
				{
					name: "float64 负的完全立方数",
					give: float64(-64.0),
					want: float64(-4.0),
				},
				{
					name: "float64 非完全立方数",
					give: float64(20.0),
					want: float64(2.7144176),
				},
				{
					name: "float64 零",
					give: float64(0.0),
					want: float64(0.0),
				},
				// 边界情况
				{
					name: "极大数的立方根",
					give: float64(1e30),
					want: float64(1e10),
				},
				{
					name: "极小正数的立方根",
					give: float64(1e-30),
					want: float64(1e-10),
				},
			}

			for _, tt := range tests {
				tt := tt // 避免竞态
				t.Run(tt.name, func(t *testing.T) {
					switch give := tt.give.(type) {
					case int:
						got := Cbrt(give)
						assert.Equal(t, tt.want, got, "Cbrt() 的结果应与期望值相等")
					case float32:
						got := Cbrt(give)
						assert.InDelta(t, tt.want, got, 1e-6, "Cbrt() 的结果应在允许误差范围内")
					case float64:
						got := Cbrt(give)
						assert.InDelta(t, tt.want, got, 1e-7, "Cbrt() 的结果应在允许误差范围内")
					}
				})
			}
		})

		// 幂运算测试
		t.Run("Pow", func(t *testing.T) {
			tests := []struct {
				name string
				x    interface{}
				y    interface{}
				want interface{}
			}{
				// 整数类型测试
				{
					name: "整数正幂",
					x:    int(2),
					y:    int(3),
					want: int(8),
				},
				{
					name: "整数零幂",
					x:    int(5),
					y:    int(0),
					want: int(1),
				},
				{
					name: "整数负幂",
					x:    int(2),
					y:    int(-2),
					want: int(0),
				},
				{
					name: "零的整数幂",
					x:    int(0),
					y:    int(5),
					want: int(0),
				},
				{
					name: "一的整数幂",
					x:    int(1),
					y:    int(100),
					want: int(1),
				},
				// float32 类型测试
				{
					name: "float32 正幂",
					x:    float32(2.0),
					y:    float32(3.0),
					want: float32(8.0),
				},
				{
					name: "float32 小数幂",
					x:    float32(4.0),
					y:    float32(0.5),
					want: float32(2.0),
				},
				{
					name: "float32 负幂",
					x:    float32(2.0),
					y:    float32(-2.0),
					want: float32(0.25),
				},
				{
					name: "float32 零幂",
					x:    float32(3.14),
					y:    float32(0.0),
					want: float32(1.0),
				},
				// float64 类型测试
				{
					name: "float64 正幂",
					x:    float64(2.0),
					y:    float64(10.0),
					want: float64(1024.0),
				},
				{
					name: "float64 小数幂",
					x:    float64(9.0),
					y:    float64(0.5),
					want: float64(3.0),
				},
				{
					name: "float64 负幂",
					x:    float64(2.0),
					y:    float64(-3.0),
					want: float64(0.125),
				},
				{
					name: "float64 零幂",
					x:    float64(2.718),
					y:    float64(0.0),
					want: float64(1.0),
				},
				// 边界情况
				{
					name: "零的零幂",
					x:    float64(0.0),
					y:    float64(0.0),
					want: float64(1.0),
				},
				{
					name: "负数的整数幂",
					x:    float64(-2.0),
					y:    float64(3.0),
					want: float64(-8.0),
				},
				{
					name: "负数的偶数幂",
					x:    float64(-2.0),
					y:    float64(4.0),
					want: float64(16.0),
				},
			}

			for _, tt := range tests {
				tt := tt // 避免竞态
				t.Run(tt.name, func(t *testing.T) {
					switch x := tt.x.(type) {
					case int:
						y := tt.y.(int)
						got := Pow(x, y)
						assert.Equal(t, tt.want, got, "Pow() 的结果应与期望值相等")
					case float32:
						y := tt.y.(float32)
						got := Pow(x, y)
						assert.InDelta(t, tt.want, got, 1e-6, "Pow() 的结果应在允许误差范围内")
					case float64:
						y := tt.y.(float64)
						got := Pow(x, y)
						assert.InDelta(t, tt.want, got, 1e-10, "Pow() 的结果应在允许误差范围内")
					}
				})
			}
		})
	})

	// 聚合运算组 - 切片运算
	t.Run("AggregateOperations", func(t *testing.T) {
		// 求和测试
		t.Run("Sum", func(t *testing.T) {
			t.Run("positive integers", func(t *testing.T) {
				input := []int{1, 2, 3, 4, 5}
				result := Sum(input)
				assert.Equal(t, 15, result)
			})

			t.Run("empty slice", func(t *testing.T) {
				input := []int{}
				result := Sum(input)
				assert.Equal(t, 0, result)
			})

			t.Run("single element", func(t *testing.T) {
				input := []int{42}
				result := Sum(input)
				assert.Equal(t, 42, result)
			})

			t.Run("negative integers", func(t *testing.T) {
				input := []int{-1, -2, -3}
				result := Sum(input)
				assert.Equal(t, -6, result)
			})

			t.Run("mixed positive and negative", func(t *testing.T) {
				input := []int{-5, 10, -3, 8}
				result := Sum(input)
				assert.Equal(t, 10, result)
			})

			t.Run("zeros", func(t *testing.T) {
				input := []int{0, 0, 0}
				result := Sum(input)
				assert.Equal(t, 0, result)
			})

			t.Run("float64 slice", func(t *testing.T) {
				input := []float64{1.5, 2.5, 3.0}
				result := Sum(input)
				assert.Equal(t, 7.0, result)
			})

			t.Run("float32 slice", func(t *testing.T) {
				input := []float32{1.1, 2.2, 3.3}
				result := Sum(input)
				assert.InDelta(t, 6.6, result, 1e-6)
			})

			t.Run("int32 slice", func(t *testing.T) {
				input := []int32{10, 20, 30}
				result := Sum(input)
				assert.Equal(t, int32(60), result)
			})

			t.Run("int64 slice", func(t *testing.T) {
				input := []int64{100, 200, 300}
				result := Sum(input)
				assert.Equal(t, int64(600), result)
			})

			t.Run("uint slice", func(t *testing.T) {
				input := []uint{1, 2, 3}
				result := Sum(input)
				assert.Equal(t, uint(6), result)
			})

			t.Run("large numbers", func(t *testing.T) {
				input := []int64{1000000, 2000000, 3000000}
				result := Sum(input)
				assert.Equal(t, int64(6000000), result)
			})
		})

		// 平均值测试
		t.Run("Average", func(t *testing.T) {
			tests := []struct {
				name   string
				input  interface{}
				output interface{}
			}{
				{
					name:   "empty int slice",
					input:  []int{},
					output: 0,
				},
				{
					name:   "single int",
					input:  []int{5},
					output: 5,
				},
				{
					name:   "multiple ints",
					input:  []int{1, 2, 3, 4, 5},
					output: 3,
				},
				{
					name:   "negative ints",
					input:  []int{-1, -2, -3},
					output: -2,
				},
				{
					name:   "mixed positive negative",
					input:  []int{-5, 0, 5},
					output: 0,
				},
				{
					name:   "empty float64 slice",
					input:  []float64{},
					output: 0.0,
				},
				{
					name:   "single float64",
					input:  []float64{3.5},
					output: 3.5,
				},
				{
					name:   "multiple float64",
					input:  []float64{1.0, 2.0, 3.0},
					output: 2.0,
				},
				{
					name:   "float64 with precision",
					input:  []float64{1.1, 2.2, 3.3},
					output: 2.2,
				},
				{
					name:   "empty int32 slice",
					input:  []int32{},
					output: int32(0),
				},
				{
					name:   "int32 values",
					input:  []int32{10, 20, 30},
					output: int32(20),
				},
				{
					name:   "empty int64 slice",
					input:  []int64{},
					output: int64(0),
				},
				{
					name:   "int64 values",
					input:  []int64{100, 200, 300},
					output: int64(200),
				},
			}

			for _, tt := range tests {
				t.Run(tt.name, func(t *testing.T) {
					switch input := tt.input.(type) {
					case []int:
						result := Average(input)
						assert.Equal(t, tt.output, result)
					case []float64:
						result := Average(input)
						assert.InDelta(t, tt.output, result, 1e-9)
					case []int32:
						result := Average(input)
						assert.Equal(t, tt.output, result)
					case []int64:
						result := Average(input)
						assert.Equal(t, tt.output, result)
					}
				})
			}
		})

		// 最大值测试
		t.Run("Max", func(t *testing.T) {
			// 测试整数切片
			assert.Equal(t, 3, Max([]int{1, 2, 3}))
			assert.Equal(t, 1, Max([]int{1}))
			assert.Equal(t, 0, Max([]int{0, -1, -2}))
			assert.Equal(t, 9, Max([]int{9, 2, 5, 1, 8}))

			// 测试浮点数切片
			assert.Equal(t, 3.14, Max([]float64{1.1, 2.2, 3.14}))
			assert.Equal(t, 0.5, Max([]float64{-0.1, 0.5, -0.3}))
			assert.Equal(t, 1.0, Max([]float64{1.0}))

			// 测试字符串切片
			assert.Equal(t, "c", Max([]string{"a", "b", "c"}))
			assert.Equal(t, "z", Max([]string{"z", "a", "b"}))
			assert.Equal(t, "hello", Max([]string{"hello"}))

			// 测试空切片
			require.Equal(t, 0, Max([]int{}))
			require.Equal(t, 0.0, Max([]float64{}))
			require.Equal(t, "", Max([]string{}))

			// 测试单元素切片
			require.Equal(t, 42, Max([]int{42}))
			require.Equal(t, 3.14, Max([]float64{3.14}))
			require.Equal(t, "test", Max([]string{"test"}))

			// 测试负数
			require.Equal(t, -1, Max([]int{-5, -3, -1, -10}))
			require.Equal(t, -0.5, Max([]float64{-2.5, -1.5, -0.5, -3.0}))

			// 测试重复值
			require.Equal(t, 5, Max([]int{5, 5, 5, 5}))
			require.Equal(t, "b", Max([]string{"a", "b", "b", "a"}))
		})

		// 最大值边界测试
		t.Run("MaxEdgeCases", func(t *testing.T) {
			// 测试极大值
			require.Equal(t, 2147483647, Max([]int{2147483647, 0, -2147483648}))

			// 测试极小值
			require.Equal(t, -2147483647, Max([]int{-2147483648, -2147483647}))

			// 测试浮点数边界
			require.Equal(t, 1.7976931348623157e+308, Max([]float64{1.7976931348623157e+308, 0.0, -1.7976931348623157e+308}))

			// 测试空字符串
			require.Equal(t, "", Max([]string{""}))

			// 测试单字符字符串
			require.Equal(t, "z", Max([]string{"a", "z", "A", "Z"}))
		})

		// 最大值 nil 切片测试
		t.Run("MaxNilSlice", func(t *testing.T) {
			// 测试 nil 切片
			require.Equal(t, 0, Max([]int(nil)))
			require.Equal(t, 0.0, Max([]float64(nil)))
			require.Equal(t, "", Max([]string(nil)))
		})

		// 最小值测试
		t.Run("Min", func(t *testing.T) {
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
		})
	})
}

// BenchmarkMathNumericOperations 数学运算性能基准测试
func BenchmarkMathNumericOperations(b *testing.B) {
	// 基础运算性能测试
	b.Run("BasicOperations", func(b *testing.B) {
		// Abs 性能测试暂无，可以根据需要添加
		// Sqrt 性能测试暂无，可以根据需要添加
		// Cbrt 性能测试暂无，可以根据需要添加
		// Pow 性能测试暂无，可以根据需要添加
	})

	// 聚合运算性能测试
	b.Run("AggregateOperations", func(b *testing.B) {
		// Sum 性能测试暂无，可以根据需要添加

		// Average 性能测试暂无，可以根据需要添加

		// Max 性能测试
		b.Run("Max", func(b *testing.B) {
			// 测试小整数切片性能
			smallInts := []int{1, 2, 3, 4, 5}
			b.Run("SmallIntSlice", func(b *testing.B) {
				for i := 0; i < b.N; i++ {
					Max(smallInts)
				}
			})

			// 测试大整数切片性能
			largeInts := make([]int, 1000)
			for i := range largeInts {
				largeInts[i] = i
			}
			b.Run("LargeIntSlice", func(b *testing.B) {
				for i := 0; i < b.N; i++ {
					Max(largeInts)
				}
			})

			// 测试字符串切片性能
			strings := make([]string, 1000)
			for i := range strings {
				strings[i] = string(rune(i % 1000))
			}
			b.Run("StringSlice", func(b *testing.B) {
				for i := 0; i < b.N; i++ {
					Max(strings)
				}
			})

			// 测试浮点数切片性能
			floats := make([]float64, 1000)
			for i := range floats {
				floats[i] = float64(i) * 0.1
			}
			b.Run("FloatSlice", func(b *testing.B) {
				for i := 0; i < b.N; i++ {
					Max(floats)
				}
			})
		})

		// Min 性能测试
		b.Run("Min", func(b *testing.B) {
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
		})
	})
}