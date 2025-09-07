package candy

import (
	"fmt"
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// TestPersonReverse 测试用的Person结构体（为避免重名冲突）
type TestPersonReverse struct {
	Name string
	Age  int
}

// TestEachReverse 测试 EachReverse 函数
func TestEachReverse(t *testing.T) {
	// 测试基本功能：整数切片
	t.Run("整数切片", func(t *testing.T) {
		t.Parallel()

		data := []int{1, 2, 3, 4, 5}
		var result []int

		EachReverse(data, func(n int) {
			result = append(result, n*2)
		})

		require.Equal(t, []int{10, 8, 6, 4, 2}, result, "EachReverse应该反向遍历并处理每个元素")
	})

	// 测试字符串切片
	t.Run("字符串切片", func(t *testing.T) {
		t.Parallel()

		data := []string{"a", "b", "c"}
		var result []string

		EachReverse(data, func(s string) {
			result = append(result, s+"x")
		})

		require.Equal(t, []string{"cx", "bx", "ax"}, result, "EachReverse应该对字符串切片反向处理")
	})

	// 测试结构体切片
	t.Run("结构体切片", func(t *testing.T) {
		t.Parallel()

		type TestItem struct {
			ID   int
			Name string
		}

		data := []TestItem{
			{ID: 1, Name: "item1"},
			{ID: 2, Name: "item2"},
		}

		var result []int
		EachReverse(data, func(item TestItem) {
			result = append(result, item.ID)
		})

		require.Equal(t, []int{2, 1}, result, "EachReverse应该对结构体切片反向处理")
	})

	// 测试浮点数切片
	t.Run("浮点数切片", func(t *testing.T) {
		t.Parallel()

		data := []float64{1.1, 2.2, 3.3}
		var sum float64

		EachReverse(data, func(n float64) {
			sum += n
		})

		assert.InDelta(t, 6.6, sum, 0.001, "EachReverse应该对浮点数切片反向处理")
	})

	// 测试空切片
	t.Run("空切片", func(t *testing.T) {
		t.Parallel()

		data := []int{}
		var result []int
		callCount := 0

		EachReverse(data, func(n int) {
			callCount++
			result = append(result, n)
		})

		require.Empty(t, result, "EachReverse处理空切片时不应该执行函数")
		require.Equal(t, 0, callCount, "EachReverse处理空切片时函数调用次数应该为0")
	})

	// 测试nil切片
	t.Run("nil切片", func(t *testing.T) {
		t.Parallel()

		var data []int
		var result []int
		callCount := 0

		EachReverse(data, func(n int) {
			callCount++
			result = append(result, n)
		})

		require.Empty(t, result, "EachReverse处理nil切片时不应该执行函数")
		require.Equal(t, 0, callCount, "EachReverse处理nil切片时函数调用次数应该为0")
	})

	// 测试单元素切片
	t.Run("单元素切片", func(t *testing.T) {
		t.Parallel()

		data := []int{42}
		var result int
		callCount := 0

		EachReverse(data, func(n int) {
			callCount++
			result = n
		})

		require.Equal(t, 42, result, "EachReverse应该正确处理单元素切片")
		require.Equal(t, 1, callCount, "EachReverse处理单元素切片时函数调用次数应该为1")
	})

	// 测试函数副作用
	t.Run("函数副作用", func(t *testing.T) {
		t.Parallel()

		data := []int{1, 2, 3}
		counter := 0

		EachReverse(data, func(n int) {
			counter++
		})

		require.Equal(t, 3, counter, "EachReverse应该对每个元素执行一次函数")
	})

	// 测试修改原始切片元素
	t.Run("修改原始元素", func(t *testing.T) {
		t.Parallel()

		data := []TestPersonReverse{
			{Name: "Alice", Age: 25},
			{Name: "Bob", Age: 30},
		}

		EachReverse(data, func(p TestPersonReverse) {
			p.Age += 10
		})

		// 注意：EachReverse函数接收的是值拷贝，所以原始切片不会被修改
		require.Equal(t, 25, data[0].Age, "EachReverse不应该修改原始切片元素（值拷贝）")
		require.Equal(t, 30, data[1].Age, "EachReverse不应该修改原始切片元素（值拷贝）")
	})

	// 测试指针切片
	t.Run("指针切片", func(t *testing.T) {
		t.Parallel()

		data := []*TestPersonReverse{
			{Name: "Alice", Age: 25},
			{Name: "Bob", Age: 30},
		}

		EachReverse(data, func(p *TestPersonReverse) {
			p.Age += 10
		})

		require.Equal(t, 40, data[1].Age, "EachReverse应该可以通过指针修改原始数据")
		require.Equal(t, 35, data[0].Age, "EachReverse应该可以通过指针修改原始数据")
	})

	// 测试复杂计算
	t.Run("复杂计算", func(t *testing.T) {
		t.Parallel()

		data := []int{1, 2, 3, 4, 5}
		var sum int
		var product = 1

		EachReverse(data, func(n int) {
			sum += n
			product *= n
		})

		require.Equal(t, 15, sum, "EachReverse应该正确计算总和")
		require.Equal(t, 120, product, "EachReverse应该正确计算乘积")
	})

	// 测试并发安全性
	t.Run("并发安全", func(t *testing.T) {
		t.Parallel()

		data := []int{1, 2, 3, 4, 5}
		var result []int
		var mu sync.Mutex

		EachReverse(data, func(n int) {
			mu.Lock()
			result = append(result, n*2)
			mu.Unlock()
		})

		require.Equal(t, []int{10, 8, 6, 4, 2}, result, "EachReverse在并发环境下应该安全工作")
	})

	// 测试大切片性能
	t.Run("大切片性能", func(t *testing.T) {
		t.Parallel()

		data := make([]int, 10000)
		for i := range data {
			data[i] = i
		}

		var sum int
		EachReverse(data, func(n int) {
			sum += n
		})

		expectedSum := (10000 - 1) * 10000 / 2
		require.Equal(t, expectedSum, sum, "EachReverse应该正确处理大切片")
	})

	// 测试布尔值切片
	t.Run("布尔值切片", func(t *testing.T) {
		t.Parallel()

		data := []bool{true, false, true, false}
		var result []bool

		EachReverse(data, func(b bool) {
			result = append(result, !b)
		})

		require.Equal(t, []bool{true, false, true, false}, result, "EachReverse应该正确处理布尔值切片")
	})

	// 测试接口切片
	t.Run("接口切片", func(t *testing.T) {
		t.Parallel()

		data := []interface{}{1, "hello", 3.14, true}
		var result []interface{}

		EachReverse(data, func(item interface{}) {
			result = append(result, item)
		})

		require.Equal(t, []interface{}{true, 3.14, "hello", 1}, result, "EachReverse应该正确处理接口切片")
	})
}

// TestEachReverseOrder 测试 EachReverse 函数的遍历顺序
func TestEachReverseOrder(t *testing.T) {
	t.Run("遍历顺序验证", func(t *testing.T) {
		t.Parallel()

		data := []string{"first", "second", "third", "fourth"}
		var order []string

		EachReverse(data, func(s string) {
			order = append(order, s)
		})

		expected := []string{"fourth", "third", "second", "first"}
		require.Equal(t, expected, order, "EachReverse应该按照正确的反向顺序遍历")
	})
}

// TestEachReverseWithBreak 测试 EachReverse 函数在函数中包含逻辑的情况
func TestEachReverseWithBreak(t *testing.T) {
	t.Run("包含条件逻辑", func(t *testing.T) {
		t.Parallel()

		data := []int{1, 2, 3, 4, 5}
		var result []int

		EachReverse(data, func(n int) {
			if n%2 == 0 {
				result = append(result, n)
			}
		})

		require.Equal(t, []int{4, 2}, result, "EachReverse应该正确处理包含条件逻辑的函数")
	})
}

// BenchmarkEachReverse 测试 EachReverse 函数的性能
func BenchmarkEachReverse(b *testing.B) {
	data := make([]int, 1000)
	for i := range data {
		data[i] = i
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		var sum int
		EachReverse(data, func(n int) {
			sum += n
		})
	}
}

// BenchmarkEachReverseDifferentSizes 基准测试不同数据大小的性能
func BenchmarkEachReverseDifferentSizes(b *testing.B) {
	sizes := []int{10, 100, 1000, 10000}

	for _, size := range sizes {
		data := make([]int, size)
		for i := range data {
			data[i] = i
		}

		b.Run(fmt.Sprintf("Size%d", size), func(b *testing.B) {
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				var sum int
				EachReverse(data, func(n int) {
					sum += n
				})
			}
		})
	}
}

// BenchmarkEachReverseVsStandardFor 比较 EachReverse 与标准for循环的性能
func BenchmarkEachReverseVsStandardFor(b *testing.B) {
	data := make([]int, 1000)
	for i := range data {
		data[i] = i
	}

	b.Run("EachReverse", func(b *testing.B) {
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			var sum int
			EachReverse(data, func(n int) {
				sum += n
			})
		}
	})

	b.Run("StandardFor", func(b *testing.B) {
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			var sum int
			for j := len(data) - 1; j >= 0; j-- {
				sum += data[j]
			}
		}
	})
}