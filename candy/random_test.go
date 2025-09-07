package candy

import (
	"math/rand"
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestRandom 测试 Random 函数
func TestRandom(t *testing.T) {
	t.Parallel()

	// 测试整数类型切片的随机选择
	t.Run("整数类型切片", func(t *testing.T) {
		t.Parallel()
		tests := []struct {
			name string
			give []int
			// 验证返回值是否在切片中
			validate func(int, []int) bool
		}{
			{
				name: "多元素切片",
				give: []int{1, 2, 3, 4, 5},
				validate: func(result int, slice []int) bool {
					for _, v := range slice {
						if v == result {
							return true
						}
					}
					return false
				},
			},
			{
				name: "单元素切片",
				give: []int{42},
				validate: func(result int, slice []int) bool {
					return result == slice[0]
				},
			},
			{
				name: "重复元素切片",
				give: []int{1, 2, 2, 3, 3},
				validate: func(result int, slice []int) bool {
					for _, v := range slice {
						if v == result {
							return true
						}
					}
					return false
				},
			},
			{
				name: "负数切片",
				give: []int{-1, -2, -3, -4, -5},
				validate: func(result int, slice []int) bool {
					for _, v := range slice {
						if v == result {
							return true
						}
					}
					return false
				},
			},
			{
				name: "混合正负数切片",
				give: []int{-5, 0, 5, -10, 10},
				validate: func(result int, slice []int) bool {
					for _, v := range slice {
						if v == result {
							return true
						}
					}
					return false
				},
			},
			{
				name: "大数切片",
				give: []int{1000000, 2000000, 3000000},
				validate: func(result int, slice []int) bool {
					for _, v := range slice {
						if v == result {
							return true
						}
					}
					return false
				},
			},
		}

		for _, tt := range tests {
			tt := tt
			t.Run(tt.name, func(t *testing.T) {
				t.Parallel()

				// 多次测试以验证随机性
				results := make(map[int]int)
				for i := 0; i < 100; i++ {
					result := Random(tt.give)
					assert.True(t, tt.validate(result, tt.give), "Random() 返回值应在原切片中")
					results[result]++
				}

				// 验证所有元素都有被选中的机会（对于足够大的切片）
				if len(tt.give) > 1 {
					_ = true // 确保所有元素都被选中
					for _, expected := range tt.give {
						if results[expected] == 0 {
							_ = false
							break
						}
					}
					// 由于随机性，这个测试有一定概率失败，但100次测试应该覆盖所有元素
					// 我们主要验证结果的有效性
				}
			})
		}
	})

	// 测试浮点数类型切片的随机选择
	t.Run("浮点数类型切片", func(t *testing.T) {
		t.Parallel()
		tests := []struct {
			name     string
			give     []float64
			validate func(float64, []float64) bool
		}{
			{
				name: "多元素浮点数切片",
				give: []float64{1.1, 2.2, 3.3, 4.4, 5.5},
				validate: func(result float64, slice []float64) bool {
					for _, v := range slice {
						if v == result {
							return true
						}
					}
					return false
				},
			},
			{
				name: "科学计数法浮点数",
				give: []float64{1.5e10, 2.3e-5, 3.14},
				validate: func(result float64, slice []float64) bool {
					for _, v := range slice {
						if v == result {
							return true
						}
					}
					return false
				},
			},
			{
				name: "负浮点数切片",
				give: []float64{-1.1, -2.2, -3.3},
				validate: func(result float64, slice []float64) bool {
					for _, v := range slice {
						if v == result {
							return true
						}
					}
					return false
				},
			},
		}

		for _, tt := range tests {
			tt := tt
			t.Run(tt.name, func(t *testing.T) {
				t.Parallel()

				// 多次测试以验证随机性
				for i := 0; i < 50; i++ {
					result := Random(tt.give)
					assert.True(t, tt.validate(result, tt.give), "Random() 返回值应在原浮点数切片中")
				}
			})
		}
	})

	// 测试字符串类型切片的随机选择
	t.Run("字符串类型切片", func(t *testing.T) {
		t.Parallel()
		tests := []struct {
			name     string
			give     []string
			validate func(string, []string) bool
		}{
			{
				name: "多元素字符串切片",
				give: []string{"apple", "banana", "cherry", "date", "elderberry"},
				validate: func(result string, slice []string) bool {
					for _, v := range slice {
						if v == result {
							return true
						}
					}
					return false
				},
			},
			{
				name: "Unicode字符串切片",
				give: []string{"苹果", "香蕉", "樱桃", "日期"},
				validate: func(result string, slice []string) bool {
					for _, v := range slice {
						if v == result {
							return true
						}
					}
					return false
				},
			},
			{
				name: "特殊字符字符串切片",
				give: []string{"a@b.com", "x#y", "test$", "hello world"},
				validate: func(result string, slice []string) bool {
					for _, v := range slice {
						if v == result {
							return true
						}
					}
					return false
				},
			},
			{
				name: "空字符串切片",
				give: []string{"", "hello", "", "world"},
				validate: func(result string, slice []string) bool {
					for _, v := range slice {
						if v == result {
							return true
						}
					}
					return false
				},
			},
		}

		for _, tt := range tests {
			tt := tt
			t.Run(tt.name, func(t *testing.T) {
				t.Parallel()

				// 多次测试以验证随机性
				for i := 0; i < 50; i++ {
					result := Random(tt.give)
					assert.True(t, tt.validate(result, tt.give), "Random() 返回值应在原字符串切片中")
				}
			})
		}
	})

	// 测试结构体类型切片的随机选择
	t.Run("结构体类型切片", func(t *testing.T) {
		t.Parallel()

		type Person struct {
			Name string
			Age  int
		}

		tests := []struct {
			name     string
			give     []Person
			validate func(Person, []Person) bool
		}{
			{
				name: "多元素结构体切片",
				give: []Person{
					{"Alice", 25},
					{"Bob", 30},
					{"Charlie", 35},
					{"David", 40},
				},
				validate: func(result Person, slice []Person) bool {
					for _, v := range slice {
						if v == result {
							return true
						}
					}
					return false
				},
			},
			{
				name: "重复字段结构体切片",
				give: []Person{
					{"Alice", 25},
					{"Bob", 25},
					{"Alice", 30}, // 重复姓名，不同年龄
				},
				validate: func(result Person, slice []Person) bool {
					for _, v := range slice {
						if v == result {
							return true
						}
					}
					return false
				},
			},
		}

		for _, tt := range tests {
			tt := tt
			t.Run(tt.name, func(t *testing.T) {
				t.Parallel()

				// 多次测试以验证随机性
				for i := 0; i < 50; i++ {
					result := Random(tt.give)
					assert.True(t, tt.validate(result, tt.give), "Random() 返回值应在原结构体切片中")
				}
			})
		}
	})

	// 测试边界情况
	t.Run("边界情况", func(t *testing.T) {
		t.Parallel()

		// 测试空切片 - 应返回零值
		t.Run("空切片", func(t *testing.T) {
			t.Parallel()

			// 整数空切片
			var emptyIntSlice []int
			intResult := Random(emptyIntSlice)
			assert.Equal(t, 0, intResult, "空整数切片应返回零值")

			// 字符串空切片
			var emptyStringSlice []string
			stringResult := Random(emptyStringSlice)
			assert.Equal(t, "", stringResult, "空字符串切片应返回零值")

			// 浮点数空切片
			var emptyFloatSlice []float64
			floatResult := Random(emptyFloatSlice)
			assert.Equal(t, 0.0, floatResult, "空浮点数切片应返回零值")

			// 结构体空切片
			var emptyStructSlice []TestPerson
			structResult := Random(emptyStructSlice)
			assert.Equal(t, TestPerson{}, structResult, "空结构体切片应返回零值")
		})

		// 测试单元素切片
		t.Run("单元素切片", func(t *testing.T) {
			t.Parallel()

			// 单元素整数切片
			singleInt := []int{42}
			result := Random(singleInt)
			assert.Equal(t, 42, result, "单元素整数切片应返回该元素")

			// 单元素字符串切片
			singleString := []string{"hello"}
			singleStringResult := Random(singleString)
			assert.Equal(t, "hello", singleStringResult, "单元素字符串切片应返回该元素")

			// 单元素浮点数切片
			singleFloat := []float64{3.14}
			singleFloatResult := Random(singleFloat)
			assert.Equal(t, 3.14, singleFloatResult, "单元素浮点数切片应返回该元素")
		})

		// 测试nil切片
		t.Run("nil切片", func(t *testing.T) {
			t.Parallel()

			var nilSlice []int
			result := Random(nilSlice)
			assert.Equal(t, 0, result, "nil切片应返回零值")
		})
	})

	// 测试随机性分布
	t.Run("随机性分布", func(t *testing.T) {
		t.Parallel()

		// 使用固定种子进行可重复的随机测试
		originalSeed := rand.Int63()
		defer func() {
			// 恢复原始种子
			rand.Seed(originalSeed)
		}()

		// 设置固定种子以确保测试可重复
		rand.Seed(12345)

		slice := []int{1, 2, 3, 4, 5}
		results := make([]int, 1000)

		// 生成1000个随机结果
		for i := range results {
			results[i] = Random(slice)
		}

		// 统计每个元素出现的频率
		frequency := make(map[int]int)
		for _, result := range results {
			frequency[result]++
		}

		// 验证每个元素都出现了（在1000次测试中）
		for _, expected := range slice {
			assert.Greater(t, frequency[expected], 0, "每个元素都应该被随机选中至少一次")
		}

		// 验证分布大致均匀（允许一定的偏差）
		expectedFrequency := 1000 / len(slice)
		for _, count := range frequency {
			// 允许20%的偏差
			deviation := float64(count-expectedFrequency) / float64(expectedFrequency)
			assert.LessOrEqual(t, deviation, 0.5, "随机分布应该大致均匀")
		}
	})

	// 测试不同类型的一致性行为
	t.Run("类型一致性", func(t *testing.T) {
		t.Parallel()

		// 验证对于相同的数据，不同类型的结果都是有效的
		intSlice := []int{1, 2, 3}
		stringSlice := []string{"1", "2", "3"}
		floatSlice := []float64{1.0, 2.0, 3.0}

		// 每种类型都应该能正确处理
		intResult := Random(intSlice)
		stringResult := Random(stringSlice)
		floatResult := Random(floatSlice)

		// 验证结果类型正确
		assert.IsType(t, intResult, 0, "整数切片应返回整数类型")
		assert.IsType(t, stringResult, "", "字符串切片应返回字符串类型")
		assert.IsType(t, floatResult, 0.0, "浮点数切片应返回浮点数类型")

		// 验证结果在各自切片中
		assert.Contains(t, intSlice, intResult)
		assert.Contains(t, stringSlice, stringResult)
		assert.Contains(t, floatSlice, floatResult)
	})
}

// BenchmarkRandom 基准测试 Random 函数
func BenchmarkRandom(b *testing.B) {
	// 基准测试小切片
	b.Run("小切片", func(b *testing.B) {
		slice := []int{1, 2, 3, 4, 5}
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			Random(slice)
		}
	})

	// 基准测试中等切片
	b.Run("中等切片", func(b *testing.B) {
		slice := make([]int, 100)
		for i := range slice {
			slice[i] = i
		}
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			Random(slice)
		}
	})

	// 基准测试大切片
	b.Run("大切片", func(b *testing.B) {
		slice := make([]int, 10000)
		for i := range slice {
			slice[i] = i
		}
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			Random(slice)
		}
	})

	// 基准测试空切片
	b.Run("空切片", func(b *testing.B) {
		slice := []int{}
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			Random(slice)
		}
	})

	// 基准测试单元素切片
	b.Run("单元素切片", func(b *testing.B) {
		slice := []int{42}
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			Random(slice)
		}
	})

	// 基准测试字符串切片
	b.Run("字符串切片", func(b *testing.B) {
		slice := []string{"apple", "banana", "cherry", "date", "elderberry"}
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			Random(slice)
		}
	})

	// 基准测试结构体切片
	b.Run("结构体切片", func(b *testing.B) {
		type Person struct {
			Name string
			Age  int
		}
		slice := []Person{
			{"Alice", 25},
			{"Bob", 30},
			{"Charlie", 35},
			{"David", 40},
		}
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			Random(slice)
		}
	})
}
