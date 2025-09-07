package candy

import (
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// TestPerson 测试用的Person结构体
type TestPerson struct {
	Name string
	Age  int
}

// TestEach 测试Each函数
func TestEach(t *testing.T) {
	// 测试基本功能：整数切片
	t.Run("整数切片", func(t *testing.T) {
		t.Parallel()

		data := []int{1, 2, 3, 4, 5}
		var result []int

		Each(data, func(index int, value interface{}) {
			item := value.(int)
			result = append(result, item*2)
		})

		require.Equal(t, []int{2, 4, 6, 8, 10}, result, "Each应该对每个元素执行函数")
	})

	// 测试字符串切片
	t.Run("字符串切片", func(t *testing.T) {
		t.Parallel()

		data := []string{"a", "b", "c"}
		var result []string

		Each(data, func(index int, value interface{}) {
			item := value.(string)
			result = append(result, item+"x")
		})

		require.Equal(t, []string{"ax", "bx", "cx"}, result, "Each应该对字符串切片正常工作")
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
		Each(data, func(index int, value interface{}) {
			item := value.(TestItem)
			result = append(result, item.ID)
		})

		require.Equal(t, []int{1, 2}, result, "Each应该对结构体切片正常工作")
	})

	// 测试浮点数切片
	t.Run("浮点数切片", func(t *testing.T) {
		t.Parallel()

		data := []float64{1.1, 2.2, 3.3}
		var sum float64

		Each(data, func(index int, value interface{}) {
			item := value.(float64)
			sum += item
		})

		assert.InDelta(t, 6.6, sum, 0.001, "Each应该对浮点数切片正常工作")
	})

	// 测试空切片
	t.Run("空切片", func(t *testing.T) {
		t.Parallel()

		data := []int{}
		var result []int

		Each(data, func(index int, value interface{}) {
			item := value.(int)
			result = append(result, item)
		})

		require.Empty(t, result, "Each处理空切片时不应该执行函数")
	})

	// 测试nil切片
	t.Run("nil切片", func(t *testing.T) {
		t.Parallel()

		var data []int
		var result []int

		Each(data, func(index int, value interface{}) {
			item := value.(int)
			result = append(result, item)
		})

		require.Empty(t, result, "Each处理nil切片时不应该执行函数")
	})

	// 测试单元素切片
	t.Run("单元素切片", func(t *testing.T) {
		t.Parallel()

		data := []int{42}
		var result int

		Each(data, func(index int, value interface{}) {
			item := value.(int)
			result = item
		})

		require.Equal(t, 42, result, "Each应该正确处理单元素切片")
	})

	// 测试函数副作用
	t.Run("函数副作用", func(t *testing.T) {
		t.Parallel()

		data := []int{1, 2, 3}
		counter := 0

		Each(data, func(index int, value interface{}) {
			counter++
		})

		require.Equal(t, 3, counter, "Each应该对每个元素执行一次函数")
	})

	// 测试修改原始切片元素
	t.Run("修改原始元素", func(t *testing.T) {
		t.Parallel()

		data := []TestPerson{
			{Name: "Alice", Age: 25},
			{Name: "Bob", Age: 30},
		}

		Each(data, func(index int, value interface{}) {
			item := value.(TestPerson)
			item.Age += 10
		})

		// 注意：Each函数接收的是值拷贝，所以原始切片不会被修改
		require.Equal(t, 25, data[0].Age, "Each不应该修改原始切片元素（值拷贝）")
		require.Equal(t, 30, data[1].Age, "Each不应该修改原始切片元素（值拷贝）")
	})

	// 测试指针切片
	t.Run("指针切片", func(t *testing.T) {
		t.Parallel()

		data := []*TestPerson{
			{Name: "Alice", Age: 25},
			{Name: "Bob", Age: 30},
		}

		Each(data, func(index int, value interface{}) {
			item := value.(*TestPerson)
			item.Age += 10
		})

		require.Equal(t, 35, data[0].Age, "Each应该可以通过指针修改原始数据")
		require.Equal(t, 40, data[1].Age, "Each应该可以通过指针修改原始数据")
	})

	// 测试复杂计算
	t.Run("复杂计算", func(t *testing.T) {
		t.Parallel()

		data := []int{1, 2, 3, 4, 5}
		var sum int
		var product = 1

		Each(data, func(index int, value interface{}) {
			item := value.(int)
			sum += item
			product *= item
		})

		require.Equal(t, 15, sum, "Each应该正确计算总和")
		require.Equal(t, 120, product, "Each应该正确计算乘积")
	})

	// 测试并发安全性
	t.Run("并发安全", func(t *testing.T) {
		t.Parallel()

		data := []int{1, 2, 3, 4, 5}
		var result []int
		var mu sync.Mutex

		Each(data, func(index int, value interface{}) {
			item := value.(int)
			mu.Lock()
			result = append(result, item*2)
			mu.Unlock()
		})

		require.Equal(t, []int{2, 4, 6, 8, 10}, result, "Each在并发环境下应该安全工作")
	})
}

// BenchmarkEach 测试Each函数的性能
func BenchmarkEach(b *testing.B) {
	data := make([]int, 1000)
	for i := range data {
		data[i] = i
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		var sum int
		Each(data, func(index int, value interface{}) {
			item := value.(int)
			sum += item
		})
	}
}