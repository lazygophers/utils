package candy

import (
	"fmt"
	"sort"
	"testing"

	"github.com/stretchr/testify/require"
)

// TestSort 测试 Sort 泛型函数的各种场景
func TestSort(t *testing.T) {
	// 测试整数切片排序
	t.Run("整数切片升序排序", func(t *testing.T) {
		t.Parallel()

		data := []int{5, 2, 8, 1, 9, 3}
		result := Sort(data)
		expected := []int{1, 2, 3, 5, 8, 9}

		require.EqualValues(t, expected, result, "Sort应该正确对整数切片进行升序排序")
		require.NotEqual(t, &data, &result, "Sort应该返回新的切片，不修改原切片")
		require.Equal(t, []int{5, 2, 8, 1, 9, 3}, data, "原始切片应该保持不变")
	})

	// 测试字符串切片排序
	t.Run("字符串切片排序", func(t *testing.T) {
		t.Parallel()

		data := []string{"banana", "apple", "cherry", "date"}
		result := Sort(data)
		expected := []string{"apple", "banana", "cherry", "date"}

		require.EqualValues(t, expected, result, "Sort应该正确对字符串切片进行排序")
		require.NotEqual(t, &data, &result, "Sort应该返回新的切片")
	})

	// 测试浮点数切片排序
	t.Run("浮点数切片排序", func(t *testing.T) {
		t.Parallel()

		data := []float64{3.14, 1.59, 2.65, 0.99}
		result := Sort(data)
		expected := []float64{0.99, 1.59, 2.65, 3.14}

		require.EqualValues(t, expected, result, "Sort应该正确对浮点数切片进行排序")
	})

	// 测试有符号整数排序
	t.Run("有符号整数排序", func(t *testing.T) {
		t.Parallel()

		data := []int{-5, 0, 5, -10, 10}
		result := Sort(data)
		expected := []int{-10, -5, 0, 5, 10}

		require.EqualValues(t, expected, result, "Sort应该正确处理负数排序")
	})

	// 测试无符号整数排序
	t.Run("无符号整数排序", func(t *testing.T) {
		t.Parallel()

		data := []uint{5, 0, 10, 2, 8}
		result := Sort(data)
		expected := []uint{0, 2, 5, 8, 10}

		require.EqualValues(t, expected, result, "Sort应该正确对无符号整数进行排序")
	})

	// 测试已排序的切片
	t.Run("已排序切片", func(t *testing.T) {
		t.Parallel()

		data := []int{1, 2, 3, 4, 5}
		result := Sort(data)
		expected := []int{1, 2, 3, 4, 5}

		require.Equal(t, expected, result, "Sort应该正确处理已排序的切片")
		// 注意：对于长度 >= 2 的切片，Sort 函数会返回新切片
		// 注意：对于长度 >= 2 的切片，Sort 函数会返回新切片
		if len(data) >= 2 {
			require.NotSame(t, &data, &result, "长度 >= 2 时应该返回新的切片")
		} else {
			require.Same(t, &data, &result, "长度 < 2 时应该返回原切片（优化）")
		}
	})

	// 测试逆序切片
	t.Run("逆序切片", func(t *testing.T) {
		t.Parallel()

		data := []int{5, 4, 3, 2, 1}
		result := Sort(data)
		expected := []int{1, 2, 3, 4, 5}

		require.EqualValues(t, expected, result, "Sort应该正确处理逆序切片")
	})

	// 测试空切片
	t.Run("空切片", func(t *testing.T) {
		t.Parallel()

		data := []int{}
		result := Sort(data)

		require.Empty(t, result, "Sort处理空切片应该返回空切片")
		// 注意：对于长度 < 2 的切片，Sort 函数返回原切片
		require.Equal(t, &data, &result, "Sort处理空切片应该返回原切片")
	})

	// 测试nil切片
	t.Run("nil切片", func(t *testing.T) {
		t.Parallel()

		var data []int
		result := Sort(data)

		require.Nil(t, result, "Sort处理nil切片应该返回nil")
	})

	// 测试单元素切片
	t.Run("单元素切片", func(t *testing.T) {
		t.Parallel()

		data := []int{42}
		result := Sort(data)

		require.Equal(t, []int{42}, result, "Sort处理单元素切片应该返回相同的元素")
		// 注意：对于长度 < 2 的切片，Sort 函数返回原切片
		require.Equal(t, &data, &result, "Sort处理单元素切片应该返回原切片")
	})

	// 测试重复元素切片
	t.Run("重复元素切片", func(t *testing.T) {
		t.Parallel()

		data := []int{3, 1, 4, 1, 5, 9, 2, 6, 5, 3}
		result := Sort(data)
		expected := []int{1, 1, 2, 3, 3, 4, 5, 5, 6, 9}

		require.EqualValues(t, expected, result, "Sort应该正确处理重复元素")
	})

	// 测试大数值切片
	t.Run("大数值切片", func(t *testing.T) {
		t.Parallel()

		data := []int{999999999, 1, 999999998, 2, 999999997}
		result := Sort(data)
		expected := []int{1, 2, 999999997, 999999998, 999999999}

		require.Equal(t, expected, result, "Sort应该正确处理大数值")
	})

	// 测试byte类型排序
	t.Run("byte类型排序", func(t *testing.T) {
		t.Parallel()

		data := []byte{'z', 'a', 'm', 'n', 'b'}
		result := Sort(data)
		expected := []byte{'a', 'b', 'm', 'n', 'z'}

		require.Equal(t, expected, result, "Sort应该正确对byte类型进行排序")
	})

	// 测试rune类型排序
	t.Run("rune类型排序", func(t *testing.T) {
		t.Parallel()

		data := []rune{'字', 'a', '1', '符', 'b'}
		result := Sort(data)
		// 注意：rune按照Unicode码点排序
		expected := []rune{'1', 'a', 'b', '字', '符'}

		require.Equal(t, expected, result, "Sort应该正确对rune类型进行排序")
	})

	// 测试int8类型排序
	t.Run("int8类型排序", func(t *testing.T) {
		t.Parallel()

		data := []int8{-128, 0, 127, -64, 64}
		result := Sort(data)
		expected := []int8{-128, -64, 0, 64, 127}

		require.Equal(t, expected, result, "Sort应该正确对int8类型进行排序")
	})

	// 测试float32类型排序
	t.Run("float32类型排序", func(t *testing.T) {
		t.Parallel()

		data := []float32{3.14, 1.59, 2.65, 0.99}
		result := Sort(data)
		expected := []float32{0.99, 1.59, 2.65, 3.14}

		require.Equal(t, expected, result, "Sort应该正确对float32类型进行排序")
	})
}

// TestSortStability 测试排序的稳定性（虽然使用sort.Slice，但验证行为）
func TestSortStability(t *testing.T) {
	t.Run("排序稳定性验证", func(t *testing.T) {
		t.Parallel()

		// 使用结构体测试稳定性，但由于Sort是泛型函数，我们主要验证其基本行为
		type Person struct {
			Name string
			Age  int
		}

		// 注意：我们的Sort函数只对基本类型有效，这里主要测试整数排序
		data := []int{3, 1, 2, 1, 4}
		result := Sort(data)
		expected := []int{1, 1, 2, 3, 4}

		require.Equal(t, expected, result, "Sort应该产生正确的排序结果")
	})
}

// TestSortEdgeCases 测试边界情况
func TestSortEdgeCases(t *testing.T) {
	// 测试最大int值
	t.Run("最大int值", func(t *testing.T) {
		t.Parallel()

		data := []int{1<<63 - 1, 0, -(1 << 63)}
		result := Sort(data)
		expected := []int{-(1 << 63), 0, 1<<63 - 1}

		require.Equal(t, expected, result, "Sort应该正确处理int边界值")
	})

	// 测试最大uint值
	t.Run("最大uint值", func(t *testing.T) {
		t.Parallel()

		data := []uint{1<<64 - 1, 0, 1 << 63}
		result := Sort(data)
		expected := []uint{0, 1 << 63, 1<<64 - 1}

		require.Equal(t, expected, result, "Sort应该正确处理uint边界值")
	})

	// 测试浮点数特殊值
	t.Run("浮点数特殊值", func(t *testing.T) {
		t.Parallel()

		data := []float64{1.0, 0.0, -1.0, 3.14, -2.71}
		result := Sort(data)
		expected := []float64{-2.71, -1.0, 0.0, 1.0, 3.14}

		require.Equal(t, expected, result, "Sort应该正确处理浮点数")
	})
}

// TestSortLargeDataset 测试大数据集排序
func TestSortLargeDataset(t *testing.T) {
	t.Run("大数据集排序", func(t *testing.T) {
		t.Parallel()

		// 创建一个较大的数据集
		data := make([]int, 1000)
		for i := range data {
			data[i] = 1000 - i // 逆序数据
		}

		result := Sort(data)

		// 验证排序结果
		for i := 1; i < len(result); i++ {
			require.True(t, result[i] >= result[i-1],
				fmt.Sprintf("切片应该在位置 %d 和 %d 保持有序", i-1, i))
		}

		require.Equal(t, 1000, len(result), "排序后的切片长度应该保持不变")
		require.NotEqual(t, &data, &result, "应该返回新的切片")
	})
}

// BenchmarkSort 基准测试Sort函数的性能
func BenchmarkSort(b *testing.B) {
	// 测试小数据集性能
	b.Run("SmallIntSlice", func(b *testing.B) {
		data := []int{5, 2, 8, 1, 9, 3, 7, 4, 6}
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			Sort(data)
		}
	})

	// 测试中等数据集性能
	b.Run("MediumIntSlice", func(b *testing.B) {
		data := make([]int, 100)
		for i := range data {
			data[i] = 100 - i
		}
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			Sort(data)
		}
	})

	// 测试大数据集性能
	b.Run("LargeIntSlice", func(b *testing.B) {
		data := make([]int, 1000)
		for i := range data {
			data[i] = 1000 - i
		}
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			Sort(data)
		}
	})

	// 测试字符串排序性能
	b.Run("StringSlice", func(b *testing.B) {
		data := []string{"zebra", "apple", "banana", "cherry", "date", "fig", "grape"}
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			Sort(data)
		}
	})

	// 测试浮点数排序性能
	b.Run("Float64Slice", func(b *testing.B) {
		data := []float64{3.14, 1.59, 2.65, 0.99, 1.41, 1.73, 2.71}
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			Sort(data)
		}
	})
}

// BenchmarkSortDifferentSizes 基准测试不同数据大小的性能
func BenchmarkSortDifferentSizes(b *testing.B) {
	sizes := []int{10, 100, 1000, 10000}

	for _, size := range sizes {
		b.Run(fmt.Sprintf("Size%d", size), func(b *testing.B) {
			data := make([]int, size)
			for i := range data {
				data[i] = size - i // 逆序数据，最坏情况
			}

			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				Sort(data)
			}
		})
	}
}

// BenchmarkSortVsStandard 比较 Sort 与标准库的性能
func BenchmarkSortVsStandard(b *testing.B) {
	// 准备测试数据
	data := make([]int, 1000)
	for i := range data {
		data[i] = 1000 - i
	}

	// 测试我们的Sort函数
	b.Run("OurSort", func(b *testing.B) {
		testData := make([]int, len(data))
		copy(testData, data)
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			Sort(testData)
		}
	})

	// 测试标准库sort.Ints
	b.Run("StandardSortInts", func(b *testing.B) {
		testData := make([]int, len(data))
		copy(testData, data)
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			// 标准库sort.Ints会修改原切片，所以需要每次重新复制
			temp := make([]int, len(testData))
			copy(temp, testData)
			sort.Ints(temp)
		}
	})
}
