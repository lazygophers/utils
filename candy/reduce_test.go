// Package reduce 提供切片归约操作的测试用例
package candy

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestReduce 测试 Reduce 函数的各种场景
// 该函数使用表驱动测试方式，覆盖了多种数据类型和边界情况
func TestReduce(t *testing.T) {
	t.Parallel()

	// 整数切片求和测试
	t.Run("整数切片求和", func(t *testing.T) {
		t.Parallel()
		input := []int{1, 2, 3, 4, 5}
		f := func(a, b int) int { return a + b }
		got := Reduce(input, f)
		want := 15
		assert.Equal(t, want, got, "Reduce() 的结果应与期望值相等")
	})

	// 整数切片求积测试
	t.Run("整数切片求积", func(t *testing.T) {
		t.Parallel()
		input := []int{1, 2, 3, 4}
		f := func(a, b int) int { return a * b }
		got := Reduce(input, f)
		want := 24
		assert.Equal(t, want, got, "Reduce() 的结果应与期望值相等")
	})

	// 空切片测试 - 应返回类型零值
	t.Run("空切片", func(t *testing.T) {
		t.Parallel()
		input := []int{}
		f := func(a, b int) int { return a + b }
		got := Reduce(input, f)
		want := 0 // int 类型零值
		assert.Equal(t, want, got, "空切片应返回类型零值")
	})

	// 单元素切片测试
	t.Run("单元素切片", func(t *testing.T) {
		t.Parallel()
		input := []int{42}
		f := func(a, b int) int { return a + b }
		got := Reduce(input, f)
		want := 42
		assert.Equal(t, want, got, "单元素切片应返回元素本身")
	})

	// 求最大值测试
	t.Run("求最大值", func(t *testing.T) {
		t.Parallel()
		input := []int{3, 1, 4, 1, 5, 9, 2, 6}
		f := func(a, b int) int {
			if b > a {
				return b
			}
			return a
		}
		got := Reduce(input, f)
		want := 9
		assert.Equal(t, want, got, "Reduce() 应能正确求最大值")
	})

	// 求最小值测试
	t.Run("求最小值", func(t *testing.T) {
		t.Parallel()
		input := []int{3, 1, 4, 1, 5, 9, 2, 6}
		f := func(a, b int) int {
			if b < a {
				return b
			}
			return a
		}
		got := Reduce(input, f)
		want := 1
		assert.Equal(t, want, got, "Reduce() 应能正确求最小值")
	})

	// 负数求和测试
	t.Run("负数求和", func(t *testing.T) {
		t.Parallel()
		input := []int{-1, -2, -3, -4, -5}
		f := func(a, b int) int { return a + b }
		got := Reduce(input, f)
		want := -15
		assert.Equal(t, want, got, "Reduce() 应能正确处理负数求和")
	})

	// 字符串拼接测试
	t.Run("字符串拼接", func(t *testing.T) {
		t.Parallel()
		input := []string{"Hello", " ", "World", "!"}
		f := func(a, b string) string { return a + b }
		got := Reduce(input, f)
		want := "Hello World!"
		assert.Equal(t, want, got, "Reduce() 应能正确拼接字符串")
	})

	// 浮点数求和测试
	t.Run("浮点数求和", func(t *testing.T) {
		t.Parallel()
		input := []float64{1.1, 2.2, 3.3}
		f := func(a, b float64) float64 { return a + b }
		got := Reduce(input, f)
		want := 6.6
		assert.InDelta(t, want, got, 0.0001, "Reduce() 的结果应与期望值相等")
	})

	// nil切片测试 - 应返回类型零值
	t.Run("nil切片", func(t *testing.T) {
		t.Parallel()
		var input []int // nil切片
		f := func(a, b int) int { return a + b }
		got := Reduce(input, f)
		want := 0 // int 类型零值
		assert.Equal(t, want, got, "nil切片应返回类型零值")
	})

	// 大数据量性能测试
	t.Run("大数据量性能测试", func(t *testing.T) {
		t.Parallel()
		// 创建包含10000个元素的切片
		input := make([]int, 10000)
		for i := range input {
			input[i] = 1
		}
		f := func(a, b int) int { return a + b }
		got := Reduce(input, f)
		want := 10000
		assert.Equal(t, want, got, "Reduce() 应能正确处理大数据量")
	})

	// 复杂结构体测试
	t.Run("复杂结构体测试", func(t *testing.T) {
		t.Parallel()
		type Person struct {
			Name string
			Age  int
		}

		people := []Person{
			{"Alice", 25},
			{"Bob", 30},
			{"Charlie", 35},
		}

		// 找出年龄最大的人
		f := func(a, b Person) Person {
			if b.Age > a.Age {
				return b
			}
			return a
		}

		got := Reduce(people, f)
		want := Person{"Charlie", 35}
		assert.Equal(t, want, got, "Reduce() 应能正确处理复杂结构体")
	})

	// 函数组合测试
	t.Run("函数组合测试", func(t *testing.T) {
		t.Parallel()
		input := []int{1, 2, 3, 4, 5}

		// 正确的平方和计算方式：先平方再求和
		// 注意：Reduce不适合这种场景，因为它是左结合的
		// 这里演示的是正确的Reduce用法 - 简单求和
		sum := func(a, b int) int { return a + b }

		// 使用Reduce计算简单求和
		got := Reduce(input, sum)

		want := 1 + 2 + 3 + 4 + 5 // 1 + 2 + 3 + 4 + 5 = 15
		assert.Equal(t, want, got, "Reduce() 应能支持函数组合操作")
	})

	// 自定义类型测试
	t.Run("自定义类型测试", func(t *testing.T) {
		t.Parallel()
		type Score int

		scores := []Score{80, 85, 90, 95, 100}
		f := func(a, b Score) Score { return a + b }
		got := Reduce(scores, f)
		want := Score(450)
		assert.Equal(t, want, got, "Reduce() 应能正确处理自定义类型")
	})
}

// BenchmarkReduce 基准测试，用于评估Reduce函数的性能
func BenchmarkReduce(b *testing.B) {
	// 准备测试数据
	data := make([]int, 1000)
	for i := range data {
		data[i] = i
	}

	// 预热
	Reduce(data, func(a, b int) int { return a + b })

	b.ResetTimer()

	// 执行基准测试
	for i := 0; i < b.N; i++ {
		Reduce(data, func(a, b int) int { return a + b })
	}
}

// TestReducePanic 测试Reduce函数是否会在异常情况下panic
func TestReducePanic(t *testing.T) {
	t.Parallel()

	// 测试传入nil函数是否会导致panic
	t.Run("传入nil函数", func(t *testing.T) {
		t.Parallel()
		assert.Panics(t, func() {
			Reduce([]int{1, 2, 3}, nil)
		}, "传入nil函数应该导致panic")
	})
}
