package candy

import (
	"math/rand"
	"sync"
	"testing"

	"github.com/stretchr/testify/require"
)

// TestSortUsing 测试SortUsing函数
func TestSortUsing(t *testing.T) {
	// 测试1: 基本整数升序排序
	t.Run("整数升序排序", func(t *testing.T) {
		t.Parallel()

		data := []int{5, 2, 8, 1, 9, 3}
		less := func(a, b int) bool {
			return a < b
		}
		
		result := SortUsing(data, less)
		expected := []int{1, 2, 3, 5, 8, 9}
		
		require.Equal(t, expected, result, "升序排序结果应该正确")
		require.Equal(t, []int{5, 2, 8, 1, 9, 3}, data, "原始切片应该保持不变")
	})

	// 测试2: 整数降序排序
	t.Run("整数降序排序", func(t *testing.T) {
		t.Parallel()

		data := []int{5, 2, 8, 1, 9, 3}
		less := func(a, b int) bool {
			return a > b
		}
		
		result := SortUsing(data, less)
		expected := []int{9, 8, 5, 3, 2, 1}
		
		require.Equal(t, expected, result, "降序排序结果应该正确")
	})

	// 测试3: 字符串排序
	t.Run("字符串排序", func(t *testing.T) {
		t.Parallel()

		data := []string{"banana", "apple", "cherry", "date"}
		less := func(a, b string) bool {
			return a < b
		}
		
		result := SortUsing(data, less)
		expected := []string{"apple", "banana", "cherry", "date"}
		
		require.Equal(t, expected, result, "字符串排序结果应该正确")
	})

	// 测试4: 结构体排序（按年龄）
	t.Run("结构体按年龄排序", func(t *testing.T) {
		t.Parallel()

		data := []TestPerson{
			{"Alice", 25},
			{"Bob", 30},
			{"Charlie", 20},
			{"David", 35},
		}
		less := func(a, b TestPerson) bool {
			return a.Age < b.Age
		}
		
		result := SortUsing(data, less)
		expected := []TestPerson{
			{"Charlie", 20},
			{"Alice", 25},
			{"Bob", 30},
			{"David", 35},
		}
		
		require.Equal(t, expected, result, "结构体按年龄排序结果应该正确")
	})

	// 测试5: 结构体排序（按姓名）
	t.Run("结构体按姓名排序", func(t *testing.T) {
		t.Parallel()

		data := []TestPerson{
			{"Alice", 25},
			{"Bob", 30},
			{"Charlie", 20},
			{"David", 35},
		}
		less := func(a, b TestPerson) bool {
			return a.Name < b.Name
		}
		
		result := SortUsing(data, less)
		expected := []TestPerson{
			{"Alice", 25},
			{"Bob", 30},
			{"Charlie", 20},
			{"David", 35},
		}
		
		require.Equal(t, expected, result, "结构体按姓名排序结果应该正确")
	})

	// 测试6: 浮点数排序
	t.Run("浮点数排序", func(t *testing.T) {
		t.Parallel()

		data := []float64{3.14, 1.59, 2.65, 1.41}
		less := func(a, b float64) bool {
			return a < b
		}
		
		result := SortUsing(data, less)
		expected := []float64{1.41, 1.59, 2.65, 3.14}
		
		require.Equal(t, expected, result, "浮点数排序结果应该正确")
	})

	// 测试7: 空切片
	t.Run("空切片", func(t *testing.T) {
		t.Parallel()

		data := []int{}
		less := func(a, b int) bool {
			return a < b
		}
		
		result := SortUsing(data, less)
		
		require.Empty(t, result, "空切片排序应该返回空切片")
	})

	// 测试8: nil切片
	t.Run("nil切片", func(t *testing.T) {
		t.Parallel()

		var data []int
		less := func(a, b int) bool {
			return a < b
		}
		
		result := SortUsing(data, less)
		
		require.Nil(t, result, "nil切片排序应该返回nil")
	})

	// 测试9: 单元素切片
	t.Run("单元素切片", func(t *testing.T) {
		t.Parallel()

		data := []int{42}
		less := func(a, b int) bool {
			return a < b
		}
		
		result := SortUsing(data, less)
		expected := []int{42}
		
		require.Equal(t, expected, result, "单元素切片排序应该返回相同元素")
	})

	// 测试10: 已排序切片
	t.Run("已排序切片", func(t *testing.T) {
		t.Parallel()

		data := []int{1, 2, 3, 4, 5}
		less := func(a, b int) bool {
			return a < b
		}
		
		result := SortUsing(data, less)
		expected := []int{1, 2, 3, 4, 5}
		
		require.Equal(t, expected, result, "已排序切片应该保持排序状态")
	})

	// 测试11: 重复元素排序
	t.Run("重复元素排序", func(t *testing.T) {
		t.Parallel()

		data := []int{3, 1, 4, 1, 5, 9, 2, 6, 5, 3}
		less := func(a, b int) bool {
			return a < b
		}
		
		result := SortUsing(data, less)
		expected := []int{1, 1, 2, 3, 3, 4, 5, 5, 6, 9}
		
		require.Equal(t, expected, result, "重复元素排序结果应该正确")
	})

	// 测试12: 复杂比较函数
	t.Run("复杂比较函数", func(t *testing.T) {
		t.Parallel()

		data := []TestPerson{
			{"Alice", 25},
			{"Bob", 30},
			{"Alice", 20}, // 相同姓名，不同年龄
			{"Bob", 25},  // 相同姓名，不同年龄
		}
		// 先按姓名排序，姓名相同则按年龄排序
		less := func(a, b TestPerson) bool {
			if a.Name != b.Name {
				return a.Name < b.Name
			}
			return a.Age < b.Age
		}
		
		result := SortUsing(data, less)
		expected := []TestPerson{
			{"Alice", 20},
			{"Alice", 25},
			{"Bob", 25},
			{"Bob", 30},
		}
		
		require.Equal(t, expected, result, "复杂比较函数排序结果应该正确")
	})

	// 测试13: 并发安全性
	t.Run("并发安全性", func(t *testing.T) {
		t.Parallel()

		data := []int{5, 2, 8, 1, 9, 3}
		less := func(a, b int) bool {
			return a < b
		}
		
		var wg sync.WaitGroup
		results := make([][]int, 10)
		
		for i := 0; i < 10; i++ {
			wg.Add(1)
			go func(index int) {
				defer wg.Done()
				results[index] = SortUsing(data, less)
			}(i)
		}
		
		wg.Wait()
		
		expected := []int{1, 2, 3, 5, 8, 9}
		for i, result := range results {
			require.Equal(t, expected, result, "并发调用第%d次结果应该正确", i)
		}
	})
}

// BenchmarkSortUsing 测试SortUsing函数的性能
func BenchmarkSortUsing(b *testing.B) {
	// 基准测试1: 小量数据排序
	b.Run("小量数据排序", func(b *testing.B) {
		data := []int{5, 2, 8, 1, 9, 3, 6, 4, 7}
		less := func(a, b int) bool {
			return a < b
		}
		
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			_ = SortUsing(data, less)
		}
	})

	// 基准测试2: 中量数据排序
	b.Run("中量数据排序", func(b *testing.B) {
		data := make([]int, 1000)
		for i := range data {
			data[i] = i
		}
		// 随机打乱数据
		rand.Shuffle(len(data), func(i, j int) {
			data[i], data[j] = data[j], data[i]
		})
		
		less := func(a, b int) bool {
			return a < b
		}
		
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			_ = SortUsing(data, less)
		}
	})

	// 基准测试3: 大量数据排序
	b.Run("大量数据排序", func(b *testing.B) {
		data := make([]int, 10000)
		for i := range data {
			data[i] = i
		}
		// 随机打乱数据
		rand.Shuffle(len(data), func(i, j int) {
			data[i], data[j] = data[j], data[i]
		})
		
		less := func(a, b int) bool {
			return a < b
		}
		
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			_ = SortUsing(data, less)
		}
	})

	// 基准测试4: 结构体排序
	b.Run("结构体排序", func(b *testing.B) {
		data := make([]TestPerson, 1000)
		for i := range data {
			data[i] = TestPerson{
				Name:  "User" + string(rune(i%26+'A')),
				Age:   i % 100,
			}
		}
		
		less := func(a, b TestPerson) bool {
			return a.Age < b.Age
		}
		
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			_ = SortUsing(data, less)
		}
	})
}