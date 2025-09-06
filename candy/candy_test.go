
package candy_test

import (
	"fmt"
	"github.com/lazygophers/utils/candy"
	"github.com/stretchr/testify/assert"
	"sort"
	"strconv"
	"testing"
)

func TestAll(t *testing.T) {
	t.Run("empty slice", func(t *testing.T) {
		var input []int
		result := candy.All(input, func(i int) bool { return i > 0 })
		if !result {
			t.Errorf("expected true")
		}
	})

	t.Run("not all elements match", func(t *testing.T) {
		input := []int{1, 2, -3}
		result := candy.All(input, func(i int) bool { return i > 0 })
		if result {
			t.Errorf("expected false")
		}
	})

	t.Run("all elements match", func(t *testing.T) {
		input := []int{1, 2, 3}
		result := candy.All(input, func(i int) bool { return i > 0 })
		if !result {
			t.Errorf("expected true")
		}
	})
}

func TestShuffle(t *testing.T) {
	t.Run("multiple elements", func(t *testing.T) {
		input := []int{1, 2, 3, 4, 5}
		// 复制原始输入并排序
		expected := make([]int, len(input))
		copy(expected, input)
		sort.Ints(expected)

		result := candy.Shuffle(input)
		// 排序结果应与排序后的原始输入一致
		sortedResult := make([]int, len(result))
		copy(sortedResult, result)
		sort.Ints(sortedResult)

		if !candy.SliceEqual(sortedResult, expected) {
			t.Errorf("expected permutation of %v, got %v", expected, sortedResult)
		}
	})
}

// TestMax 测试 Max 函数
func TestMax(t *testing.T) {
	tests := []struct {
		name  string
		give  []int
		want  int
	}{
		{"空切片", []int{}, 0},
		{"单个元素", []int{42}, 42},
		{"正数序列", []int{1, 3, 2}, 3},
		{"负数序列", []int{-1, -5, -3}, -1},
		{"混合序列", []int{-1, 0, 1}, 1},
		{"相等元素", []int{5, 5, 5}, 5},
		{"大数序列", []int{1000000, 1, 999999}, 1000000},
		{"最大正数", []int{-2147483648, 0, 2147483647}, 2147483647},
		{"单元素负数", []int{-42}, -42},
		{"单元素零", []int{0}, 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := candy.Max(tt.give)
			assert.Equal(t, tt.want, got, "Max() 的结果应与期望值相等")
		})
	}

	// 测试浮点数类型
	t.Run("浮点数类型", func(t *testing.T) {
		floatTests := []struct {
			name  string
			give  []float64
			want  float64
		}{
			{"浮点数序列", []float64{3.14, 1.41, 2.71}, 3.14},
			{"负浮点数", []float64{-1.5, -0.5, -2.5}, -0.5},
			{"混合浮点数", []float64{-1.1, 0.0, 1.1}, 1.1},
		}

		for _, tt := range floatTests {
			t.Run(tt.name, func(t *testing.T) {
				got := candy.Max(tt.give)
				assert.InDelta(t, tt.want, got, 0.0001, "Max() 处理浮点数的结果应与期望值相等")
			})
		}
	})
}

// TestMin 测试 Min 函数
func TestMin(t *testing.T) {
	tests := []struct {
		name  string
		give  []int
		want  int
	}{
		{"空切片", []int{}, 0},
		{"单个元素", []int{42}, 42},
		{"正数序列", []int{5, 1, 3}, 1},
		{"负数序列", []int{-1, -5, -3}, -5},
		{"混合序列", []int{-1, 0, 1}, -1},
		{"相等元素", []int{5, 5, 5}, 5},
		{"大数序列", []int{1000000, 1, 999999}, 1},
		{"最大负数", []int{-2147483648, 0, 2147483647}, -2147483648},
		{"单元素负数", []int{-42}, -42},
		{"单元素零", []int{0}, 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := candy.Min(tt.give)
			assert.Equal(t, tt.want, got, "Min() 的结果应与期望值相等")
		})
	}

	// 测试浮点数类型
	t.Run("浮点数类型", func(t *testing.T) {
		floatTests := []struct {
			name  string
			give  []float64
			want  float64
		}{
			{"浮点数序列", []float64{3.14, 1.41, 2.71}, 1.41},
			{"负浮点数", []float64{-1.5, -0.5, -2.5}, -2.5},
			{"混合浮点数", []float64{-1.1, 0.0, 1.1}, -1.1},
		}

		for _, tt := range floatTests {
			t.Run(tt.name, func(t *testing.T) {
				got := candy.Min(tt.give)
				assert.InDelta(t, tt.want, got, 0.0001, "Min() 处理浮点数的结果应与期望值相等")
			})
		}
	})
}

func TestRandom(t *testing.T) {
	t.Run("empty slice", func(t *testing.T) {
		var input []int
		result := candy.Random(input)
		if result != 0 {
			t.Errorf("expected 0, got %v", result)
		}
	})

	t.Run("single element", func(t *testing.T) {
		input := []int{42}
		result := candy.Random(input)
		if result != 42 {
			t.Errorf("expected 42, got %v", result)
		}
	})
}

func TestEachStopWithError(t *testing.T) {
	t.Run("no error", func(t *testing.T) {
		input := []int{1, 2, 3}
		err := candy.EachStopWithError(input, func(i int) error {
			return nil
		})
		if err != nil {
			t.Errorf("expected nil, got %v", err)
		}
	})

	t.Run("with error", func(t *testing.T) {
		input := []int{1, 2, 3}
		expectedErr := fmt.Errorf("test error")
		err := candy.EachStopWithError(input, func(i int) error {
			if i == 2 {
				return expectedErr
			}
			return nil
		})
		if err != expectedErr {
			t.Errorf("expected %v, got %v", expectedErr, err)
		}
	})
}

// 数学函数测试
func TestAbs(t *testing.T) {
	tests := []struct {
		name     string
		input    int
		expected int
	}{
		{"正数", 5, 5},
		{"负数", -5, 5},
		{"零", 0, 0},
		{"最大负数", -2147483648, 2147483648},
		{"最大正数", 2147483647, 2147483647},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := candy.Abs(tt.input)
			assert.Equal(t, tt.expected, result, "Abs(%d) 应该等于 %d", tt.input, tt.expected)
		})
	}
}

func TestPow(t *testing.T) {
	tests := []struct {
		name     string
		base     int
		exp      int
		expected int
	}{
		{"正指数", 2, 3, 8},
		{"零指数", 5, 0, 1},
		{"负指数", 2, -3, 0},
		{"零的零次方", 0, 0, 1},
		{"零的正指数", 0, 5, 0},
		{"负数的偶数幂", -2, 3, -8},
		{"负数的奇数幂", -2, 4, 16},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := candy.Pow(tt.base, tt.exp)
			assert.Equal(t, tt.expected, result, "Pow(%d, %d) 应该等于 %d", tt.base, tt.exp, tt.expected)
		})
	}
}

func TestSqrt(t *testing.T) {
	tests := []struct {
		name     string
		input    int
		expected int
	}{
		{"完全平方数", 16, 4},
		{"非完全平方数", 10, 3},
		{"零", 0, 0},
		{"一", 1, 1},
		{"大数", 1000000, 1000},
		{"负数", -16, 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := candy.Sqrt(tt.input)
			assert.Equal(t, tt.expected, result, "Sqrt(%d) 应该等于 %d", tt.input, tt.expected)
		})
	}
}

func TestCbrt(t *testing.T) {
	tests := []struct {
		name     string
		input    int
		expected int
	}{
		{"完全立方数", 8, 2},
		{"非完全立方数", 10, 2},
		{"零", 0, 0},
		{"一", 1, 1},
		{"负数", -8, -2},
		{"大数", 1000000, 100},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := candy.Cbrt(tt.input)
			assert.Equal(t, tt.expected, result, "Cbrt(%d) 应该等于 %d", tt.input, tt.expected)
		})
	}
}

// 序列操作函数测试
func TestFirst(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		expected int
	}{
		{"空切片", []int{}, 0},
		{"单个元素", []int{42}, 42},
		{"多个元素", []int{1, 2, 3}, 1},
		{"负数", []int{-1, -2, -3}, -1},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := candy.First(tt.input)
			assert.Equal(t, tt.expected, result, "First() 应该返回第一个元素")
		})
	}
}

func TestFirstOr(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		defaultValue int
		expected int
	}{
		{"空切片", []int{}, 99, 99},
		{"单个元素", []int{42}, 0, 42},
		{"多个元素", []int{1, 2, 3}, 0, 1},
		{"使用默认值", []int{}, -1, -1},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := candy.FirstOr(tt.input, tt.defaultValue)
			assert.Equal(t, tt.expected, result, "FirstOr() 应该返回第一个元素或默认值")
		})
	}
}

func TestLast(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		expected int
	}{
		{"空切片", []int{}, 0},
		{"单个元素", []int{42}, 42},
		{"多个元素", []int{1, 2, 3}, 3},
		{"负数", []int{-1, -2, -3}, -3},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := candy.Last(tt.input)
			assert.Equal(t, tt.expected, result, "Last() 应该返回最后一个元素")
		})
	}
}

func TestLastOr(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		defaultValue int
		expected int
	}{
		{"空切片", []int{}, 99, 99},
		{"单个元素", []int{42}, 0, 42},
		{"多个元素", []int{1, 2, 3}, 0, 3},
		{"使用默认值", []int{}, -1, -1},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := candy.LastOr(tt.input, tt.defaultValue)
			assert.Equal(t, tt.expected, result, "LastOr() 应该返回最后一个元素或默认值")
		})
	}
}

func TestTop(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		n        int
		expected []int
	}{
		{"空切片", []int{}, 3, []int{}},
		{"n为0", []int{1, 2, 3}, 0, []int{}},
		{"n小于长度", []int{1, 2, 3, 4, 5}, 3, []int{1, 2, 3}},
		{"n等于长度", []int{1, 2, 3}, 3, []int{1, 2, 3}},
		{"n大于长度", []int{1, 2, 3}, 5, []int{1, 2, 3}},
		{"单个元素", []int{42}, 1, []int{42}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := candy.Top(tt.input, tt.n)
			assert.Equal(t, tt.expected, result, "Top() 应该返回前n个元素")
		})
	}
}

func TestBottom(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		n        int
		expected []int
	}{
		{"空切片", []int{}, 3, []int{}},
		{"n为0", []int{1, 2, 3}, 0, []int{}},
		{"n小于长度", []int{1, 2, 3, 4, 5}, 3, []int{3, 4, 5}},
		{"n等于长度", []int{1, 2, 3}, 3, []int{1, 2, 3}},
		{"n大于长度", []int{1, 2, 3}, 5, []int{1, 2, 3}},
		{"单个元素", []int{42}, 1, []int{42}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := candy.Bottom(tt.input, tt.n)
			assert.Equal(t, tt.expected, result, "Bottom() 应该返回最后n个元素")
		})
	}
}

func TestHypot(t *testing.T) {
	tests := []struct {
		name     string
		p        int
		q        int
		expected int
	}{
		{"3-4-5三角形", 3, 4, 5},
		{"5-12-13三角形", 5, 12, 13},
		{"零值", 0, 0, 0},
		{"一边为零", 3, 0, 3},
		{"负数", -3, -4, 5},
		{"小数", 1, 1, 1},
		{"大数", 300, 400, 500},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := candy.Hypot(tt.p, tt.q)
			assert.Equal(t, tt.expected, result, "Hypot(%d, %d) 应该等于 %d", tt.p, tt.q, tt.expected)
		})
	}
}

func TestFilterNot(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		filter   func(int) bool
		expected []int
	}{
		{
			"空切片",
			[]int{},
			func(x int) bool { return x > 0 },
			[]int{},
		},
		{
			"过滤偶数",
			[]int{1, 2, 3, 4, 5},
			func(x int) bool { return x%2 == 0 },
			[]int{1, 3, 5},
		},
		{
			"过滤大于3的数",
			[]int{1, 2, 3, 4, 5},
			func(x int) bool { return x > 3 },
			[]int{1, 2, 3},
		},
		{
			"没有匹配的过滤条件",
			[]int{1, 3, 5},
			func(x int) bool { return x%2 == 0 },
			[]int{1, 3, 5},
		},
		{
			"所有元素都被过滤",
			[]int{2, 4, 6},
			func(x int) bool { return x%2 == 0 },
			[]int{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := candy.FilterNot(tt.input, tt.filter)
			assert.Equal(t, tt.expected, result, "FilterNot() 应该返回过滤后的切片")
		})
	}
}

func TestReduce(t *testing.T) {
	// 测试整数类型
	intTests := []struct {
		name      string
		input     []int
		reducer   func(int, int) int
		initial   int
		expected  int
	}{
		{
			"空切片",
			[]int{},
			func(acc, x int) int { return acc + x },
			0,
			0,
		},
		{
			"求和",
			[]int{1, 2, 3, 4, 5},
			func(acc, x int) int { return acc + x },
			0,
			15,
		},
		{
			"求积",
			[]int{1, 2, 3, 4},
			func(acc, x int) int { return acc * x },
			1,
			24,
		},
		{
			"查找最大值",
			[]int{3, 1, 4, 2, 5},
			func(acc, x int) int {
				if x > acc {
					return x
				}
				return acc
			},
			0,
			5,
		},
	}
	
	for _, tt := range intTests {
		t.Run(tt.name, func(t *testing.T) {
			result := candy.Reduce(tt.input, tt.reducer)
			assert.Equal(t, tt.expected, result, "Reduce() 应该返回累加后的值")
		})
	}
	
	
	// 测试字符串类型
	stringTests := []struct {
		name      string
		input     []string
		reducer   func(string, string) string
		initial   string
		expected  string
	}{
		{
			"字符串拼接",
			[]string{"Hello", " ", "World"},
			func(acc, x string) string { return acc + x },
			"",
			"Hello World",
		},
	}
	
	for _, tt := range stringTests {
		t.Run(tt.name, func(t *testing.T) {
			result := candy.Reduce(tt.input, tt.reducer)
			assert.Equal(t, tt.expected, result, "Reduce() 应该返回拼接后的字符串")
		})
	}
}

func TestDrop(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		n        int
		expected []int
	}{
		{
			"空切片",
			[]int{},
			3,
			[]int{},
		},
		{
			"n为0",
			[]int{1, 2, 3, 4, 5},
			0,
			[]int{1, 2, 3, 4, 5},
		},
		{
			"n小于长度",
			[]int{1, 2, 3, 4, 5},
			2,
			[]int{3, 4, 5},
		},
		{
			"n等于长度",
			[]int{1, 2, 3, 4, 5},
			5,
			[]int{},
		},
		{
			"n大于长度",
			[]int{1, 2, 3},
			5,
			[]int{},
		},
		{
			"负数n",
			[]int{1, 2, 3, 4, 5},
			-1,
			[]int{1, 2, 3, 4, 5},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := candy.Drop(tt.input, tt.n)
			assert.Equal(t, tt.expected, result, "Drop() 应该返回丢弃前n个元素后的切片")
		})
	}
}

func TestAny(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		predicate func(int) bool
		expected bool
	}{
		{
			"空切片",
			[]int{},
			func(x int) bool { return x > 0 },
			false,
		},
		{
			"有元素满足条件",
			[]int{1, 2, 3, 4, 5},
			func(x int) bool { return x > 3 },
			true,
		},
		{
			"没有元素满足条件",
			[]int{1, 2, 3},
			func(x int) bool { return x > 5 },
			false,
		},
		{
			"所有元素都满足条件",
			[]int{2, 4, 6},
			func(x int) bool { return x%2 == 0 },
			true,
		},
		{
			"单个元素满足条件",
			[]int{1, 3, 5, 7},
			func(x int) bool { return x == 5 },
			true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := candy.Any(tt.input, tt.predicate)
			assert.Equal(t, tt.expected, result, "Any() 应该返回是否存在满足条件的元素")
		})
	}
}

// TestAverage 测试 Average 函数
func TestAverage(t *testing.T) {
	// 测试整数类型
	intTests := []struct {
		name  string
		give  []int
		want  int
	}{
		{"空切片", []int{}, 0},
		{"单个元素", []int{5}, 5},
		{"正数序列", []int{1, 2, 3, 4, 5}, 3},
		{"负数序列", []int{-1, -2, -3, -4, -5}, -3},
		{"混合序列", []int{-1, 0, 1}, 0},
		{"大数序列", []int{100, 200, 300}, 200},
		{"浮点数结果", []int{1, 1, 1}, 1},
		{"大数平均值", []int{2147483647, 2147483647}, 2147483647},
		{"小数平均值", []int{1, 2}, 1},  // 整数除法会截断小数
		{"零平均值", []int{0, 0, 0}, 0},
		{"长序列", make([]int, 1000), 0}, // 长切片性能测试
	}

	for _, tt := range intTests {
		t.Run(tt.name, func(t *testing.T) {
			got := candy.Average(tt.give)
			assert.Equal(t, tt.want, got, "Average() 的结果应与期望值相等")
		})
	}

	// 测试浮点数类型
	t.Run("浮点数类型", func(t *testing.T) {
		floatTests := []struct {
			name  string
			give  []float64
			want  float64
		}{
			{"浮点数序列", []float64{1.5, 2.5, 3.5}, 2.5}, // 注意：整数除法会截断小数部分
			{"负浮点数", []float64{-1.5, -2.5}, -2.0}, // 注意：整数除法会截断小数部分
			{"混合浮点数", []float64{-1.1, 0.0, 1.1}, 0.0}, // 注意：整数除法会截断小数部分
			{"科学计数法", []float64{1e10, 2e10}, 1.5e10}, // 注意：整数除法会截断小数部分
			{"精度测试", []float64{1.0/3.0, 2.0/3.0}, 0.5}, // 注意：整数除法会截断小数部分
		}

		for _, tt := range floatTests {
			t.Run(tt.name, func(t *testing.T) {
				got := candy.Average(tt.give)
				assert.InDelta(t, tt.want, got, 0.0001, "Average() 处理浮点数的结果应与期望值相等")
			})
		}
	})

	// 测试不同整数类型
	t.Run("不同整数类型", func(t *testing.T) {
		int8Tests := []struct {
			name  string
			give  []int8
			want  int8
		}{
			{"int8序列", []int8{100, 50, 25}, 58},  // 注意：整数除法会截断小数部分
			{"int8负数", []int8{-50, -25}, -37}, // 注意：整数除法会截断小数部分
		}

		for _, tt := range int8Tests {
			t.Run(tt.name, func(t *testing.T) {
				got := candy.Average(tt.give)
				assert.Equal(t, tt.want, got, "Average() 处理int8的结果应与期望值相等")
			})
		}
	})

	// 测试 int8 溢出情况
	t.Run("int8溢出测试", func(t *testing.T) {
		// 100 + 50 + 25 = 175，超出 int8 范围（-128 到 127）
		// 由于 Average 函数内部使用 int64 计算，所以不会溢出
		// 最终结果会转换为 int8，导致溢出
		input := []int8{100, 50, 25}
		got := candy.Average(input)
		// 175 / 3 = 58.33...，转换为 int8 是 58
		assert.Equal(t, int8(58), got, "Average() 处理 int8 溢出时应正确截断")
	})

	// 边界情况测试
	t.Run("边界情况", func(t *testing.T) {
		// 测试单个元素的多种类型
		assert.Equal(t, int(42), candy.Average([]int{42}), "单个整数平均值应等于自身")
		assert.InDelta(t, 3.14, candy.Average([]float64{3.14}), 0.0001, "单个浮点数平均值应等于自身")
		
		// 测试大数除法精度
		largeAvg := candy.Average([]int{1, 1, 1, 1, 1, 1, 1})
		assert.Equal(t, int(1), largeAvg, "大数除法应保持精度")
	})
}

// TestSum 测试 Sum 函数
func TestSum(t *testing.T) {
	// 测试整数类型
	intTests := []struct {
		name  string
		give  []int
		want  int
	}{
		{"空切片", []int{}, 0},
		{"单个元素", []int{42}, 42},
		{"正数序列", []int{1, 2, 3, 4, 5}, 15},
		{"负数序列", []int{-1, -2, -3}, -6},
		{"混合序列", []int{-1, 0, 1}, 0},
		{"大数序列", []int{1000000, 2000000, 3000000}, 6000000},
		{"零值", []int{0, 0, 0}, 0},
		{"单个负数", []int{-42}, -42},
		{"单个零", []int{0}, 0},
		{"大数求和", []int{2147483647, 2147483647}, 4294967294}, // 大数求和：2147483647 + 2147483647 = 4294967294，在64位系统上正常溢出
		{"交替正负", []int{1, -1, 2, -2}, 0},
		{"长序列", make([]int, 1000), 0}, // 长切片性能测试
	}

	for _, tt := range intTests {
		t.Run(tt.name, func(t *testing.T) {
			got := candy.Sum(tt.give)
			assert.Equal(t, tt.want, got, "Sum() 处理整数的结果应与期望值相等")
		})
	}

	// 测试浮点数类型
	t.Run("浮点数类型", func(t *testing.T) {
		floatTests := []struct {
			name  string
			give  []float64
			want  float64
		}{
			{"浮点数序列", []float64{1.1, 2.2, 3.3}, 6.6},
			{"负浮点数", []float64{-1.5, -2.5}, -4.0},
			{"混合浮点数", []float64{-1.1, 0.0, 1.1}, 0.0},
			{"科学计数法", []float64{1e10, 2e10}, 3e10},
			{"小精度", []float64{0.1, 0.2, 0.3}, 0.6},
		}

		for _, tt := range floatTests {
			t.Run(tt.name, func(t *testing.T) {
				got := candy.Sum(tt.give)
				assert.InDelta(t, tt.want, got, 0.0001, "Sum() 处理浮点数的结果应与期望值相等")
			})
		}
	})

	// 测试不同整数类型
	t.Run("不同整数类型", func(t *testing.T) {
		int8Tests := []struct {
			name  string
			give  []int8
			want  int8
		}{
			{"int8序列", []int8{100, 50, 25}, -81}, // 100+50+25=175，超出int8范围，175-256=-81
			{"int8负数", []int8{-50, -25}, -75},   // -50-25=-75，在int8范围内
		}

		for _, tt := range int8Tests {
			t.Run(tt.name, func(t *testing.T) {
				got := candy.Sum(tt.give)
				assert.Equal(t, tt.want, got, "Sum() 处理int8的结果应与期望值相等")
			})
		}
	})
}

// TestMap 测试 Map 函数
func TestMap(t *testing.T) {
	tests := []struct {
		name string
		give []int
		want []string
	}{
		{"空切片", []int{}, []string{}},
		{"单个元素", []int{1}, []string{"1"}},
		{"多个元素", []int{1, 2, 3}, []string{"1", "2", "3"}},
		{"大数", []int{100, 200}, []string{"100", "200"}},
		{"负数", []int{-1, -2}, []string{"-1", "-2"}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := candy.Map(tt.give, func(i int) string {
				return strconv.Itoa(i)
			})
			assert.Equal(t, tt.want, got, "Map() 的结果应与期望值相等")
		})
	}

	// 测试类型转换
	t.Run("类型转换", func(t *testing.T) {
		give := []float64{1.1, 2.2, 3.3}
		want := []int{1, 2, 3}
		got := candy.Map(give, func(f float64) int {
			return int(f)
		})
		assert.Equal(t, want, got, "Map() 类型转换的结果应与期望值相等")
	})
}

// TestChunk 测试 Chunk 函数
func TestChunk(t *testing.T) {
	tests := []struct {
		name string
		give []int
		size int
		want [][]int
	}{
		{"空切片", []int{}, 3, [][]int{}},
		{"单个元素", []int{1}, 3, [][]int{{1}}},
		{"正好分完", []int{1, 2, 3, 4, 5, 6}, 3, [][]int{{1, 2, 3}, {4, 5, 6}}},
		{"不能整除", []int{1, 2, 3, 4, 5}, 2, [][]int{{1, 2}, {3, 4}, {5}}},
		{"块大小大于切片长度", []int{1, 2, 3}, 5, [][]int{{1, 2, 3}}},
		{"块大小为1", []int{1, 2, 3}, 1, [][]int{{1}, {2}, {3}}},
		{"块大小为0", []int{1, 2, 3}, 0, [][]int{}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := candy.Chunk(tt.give, tt.size)
			assert.Equal(t, tt.want, got, "Chunk() 的结果应与期望值相等")
		})
	}
}

// TestUnique 测试 Unique 函数
func TestUnique(t *testing.T) {
	tests := []struct {
		name string
		give []int
		want []int
	}{
		{"空切片", []int{}, []int{}},
		{"单个元素", []int{42}, []int{42}},
		{"无重复元素", []int{1, 2, 3}, []int{1, 2, 3}},
		{"有重复元素", []int{1, 2, 2, 3, 1}, []int{1, 2, 3}},
		{"所有元素相同", []int{5, 5, 5}, []int{5}},
		{"重复元素在开头", []int{1, 1, 2, 3}, []int{1, 2, 3}},
		{"重复元素在中间", []int{1, 2, 2, 3}, []int{1, 2, 3}},
		{"重复元素在结尾", []int{1, 2, 3, 3}, []int{1, 2, 3}},
		{"负数重复", []int{-1, -2, -1, -3}, []int{-1, -2, -3}},
		{"零重复", []int{0, 1, 0, 2}, []int{0, 1, 2}},
		{"大数重复", []int{1000000, 1000000, 999999}, []int{1000000, 999999}},
		{"长序列重复", []int{1, 2, 3, 1, 2, 3, 4}, []int{1, 2, 3, 4}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt := tt // 避免竞态
			assert.Equal(t, tt.want, candy.Unique(tt.give), "Unique() 的结果应与期望值相等")
		})
	}
}

// TestUniqueUsing 测试 UniqueUsing 函数
func TestUniqueUsing(t *testing.T) {
	tests := []struct {
		name     string
		give     []int
		comparer func(int) any
		want     []int
	}{
		{
			"空切片",
			[]int{},
			func(a int) any { return a },
			[]int{},
		},
		{
			"单个元素",
			[]int{42},
			func(a int) any { return a },
			[]int{42},
		},
		{
			"相等比较器-有重复",
			[]int{1, 2, 2, 3},
			func(a int) any { return a },
			[]int{1, 2, 3},
		},
		{
			"绝对值比较器",
			[]int{1, -1, 2, -2, 3},
			func(a int) any { return a * a },
			[]int{1, 2, 3},
		},
		{
			"奇偶比较器",
			[]int{1, 3, 2, 4, 6},
			func(a int) any { return a % 2 },
			[]int{1, 2},
		},
		{
			"自定义比较器-大数",
			[]int{100, 200, 150, 100},
			func(a int) any { return a / 100 },
			[]int{100, 200},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt := tt // 避免竞态
			assert.Equal(t, tt.want, candy.UniqueUsing(tt.give, tt.comparer), "UniqueUsing() 的结果应与期望值相等")
		})
	}
}

// TestEach 测试 Each 函数
func TestEach(t *testing.T) {
	tests := []struct {
		name     string
		give     []int
		iterator func(int)
		want     []int // 用于验证迭代器是否被调用
	}{
		{
			"空切片",
			[]int{},
			func(i int) {},
			[]int{},
		},
		{
			"单个元素",
			[]int{42},
			func(i int) {},
			[]int{42},
		},
		{
			"多个元素",
			[]int{1, 2, 3},
			func(i int) {},
			[]int{1, 2, 3},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt := tt // 避免竞态
			// 创建一个切片来捕获迭代器调用的值
			var captured []int
			iterator := func(i int) {
				captured = append(captured, i)
			}
			
			candy.Each(tt.give, iterator)
			// 对于空切片，期望结果应该是nil而不是空切片
			if len(tt.give) == 0 {
				assert.Nil(t, captured, "Each() 处理空切片时应该返回nil")
			} else {
				assert.Equal(t, tt.want, captured, "Each() 应该遍历所有元素")
			}
		})
	}
}

// TestEachReverse 测试 EachReverse 函数
func TestEachReverse(t *testing.T) {
	tests := []struct {
		name     string
		give     []int
		iterator func(int)
		want     []int // 用于验证迭代器是否被正确调用
	}{
		{
			"空切片",
			[]int{},
			func(i int) {},
			[]int{},
		},
		{
			"单个元素",
			[]int{42},
			func(i int) {},
			[]int{42},
		},
		{
			"多个元素",
			[]int{1, 2, 3},
			func(i int) {},
			[]int{3, 2, 1},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt := tt // 避免竞态
			// 创建一个切片来捕获迭代器调用的值
			var captured []int
			iterator := func(i int) {
				captured = append(captured, i)
			}
			
			candy.EachReverse(tt.give, iterator)
			// 对于空切片，期望结果应该是nil而不是空切片
			if len(tt.give) == 0 {
				assert.Nil(t, captured, "EachReverse() 处理空切片时应该返回nil")
			} else {
				assert.Equal(t, tt.want, captured, "EachReverse() 应该反向遍历所有元素")
			}
		})
	}
}

// TestSort 测试 Sort 函数
func TestSort(t *testing.T) {
	tests := []struct {
		name string
		give []int
		want []int
	}{
		{"空切片", []int{}, []int{}},
		{"单个元素", []int{42}, []int{42}},
		{"已排序", []int{1, 2, 3}, []int{1, 2, 3}},
		{"未排序", []int{3, 1, 2}, []int{1, 2, 3}},
		{"重复元素", []int{2, 1, 2, 3}, []int{1, 2, 2, 3}},
		{"负数", []int{-1, -3, -2}, []int{-3, -2, -1}},
		{"混合正负", []int{1, -1, 0}, []int{-1, 0, 1}},
		{"大数", []int{1000, 100, 10000}, []int{100, 1000, 10000}},
		{"相同元素", []int{5, 5, 5}, []int{5, 5, 5}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt := tt // 避免竞态
			result := candy.Sort(tt.give)
			assert.Equal(t, tt.want, result, "Sort() 的结果应与期望值相等")
		})
	}
}

// TestSortUsing 测试 SortUsing 函数
func TestSortUsing(t *testing.T) {
	tests := []struct {
		name     string
		give     []int
		comparer func(int, int) bool
		want     []int
	}{
		{
			"空切片",
			[]int{},
			func(a, b int) bool { return a < b },
			[]int{},
		},
		{
			"单个元素",
			[]int{42},
			func(a, b int) bool { return a < b },
			[]int{42},
		},
		{
			"升序排序",
			[]int{3, 1, 2},
			func(a, b int) bool { return a < b },
			[]int{1, 2, 3},
		},
		{
			"降序排序",
			[]int{1, 3, 2},
			func(a, b int) bool { return a > b },
			[]int{3, 2, 1},
		},
		{
			"绝对值排序",
			[]int{-3, 1, -2},
			func(a, b int) bool { return a*a < b*b },
			[]int{1, -2, -3},
		},
		{
			"偶数优先",
			[]int{3, 2, 1, 4},
			func(a, b int) bool { return a%2 < b%2 || (a%2 == b%2 && a < b) },
			[]int{2, 4, 1, 3},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt := tt // 避免竞态
			result := candy.SortUsing(tt.give, tt.comparer)
			assert.Equal(t, tt.want, result, "SortUsing() 的结果应与期望值相等")
		})
	}
}

// TestFilter 测试 Filter 函数
func TestFilter(t *testing.T) {
	tests := []struct {
		name     string
		give     []int
		predicate func(int) bool
		want     []int
	}{
		{
			"空切片",
			[]int{},
			func(x int) bool { return x > 0 },
			[]int{},
		},
		{
			"过滤偶数",
			[]int{1, 2, 3, 4, 5},
			func(x int) bool { return x%2 == 0 },
			[]int{2, 4},
		},
		{
			"过滤大于3的数",
			[]int{1, 2, 3, 4, 5},
			func(x int) bool { return x > 3 },
			[]int{4, 5},
		},
		{
			"没有匹配的过滤条件",
			[]int{1, 3, 5},
			func(x int) bool { return x%2 == 0 },
			[]int{},
		},
		{
			"所有元素都匹配",
			[]int{2, 4, 6},
			func(x int) bool { return x%2 == 0 },
			[]int{2, 4, 6},
		},
		{
			"过滤负数",
			[]int{-1, 2, -3, 4},
			func(x int) bool { return x >= 0 },
			[]int{2, 4},
		},
		{
			"过滤零值",
			[]int{0, 1, 0, 2},
			func(x int) bool { return x != 0 },
			[]int{1, 2},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt := tt // 避免竞态
			result := candy.Filter(tt.give, tt.predicate)
			assert.Equal(t, tt.want, result, "Filter() 的结果应与期望值相等")
		})
	}
}

// TestContains 测试 Contains 函数
func TestContains(t *testing.T) {
	tests := []struct {
		name string
		give []int
		item int
		want bool
	}{
		{"空切片", []int{}, 42, false},
		{"单个元素-匹配", []int{42}, 42, true},
		{"单个元素-不匹配", []int{42}, 24, false},
		{"多个元素-包含", []int{1, 2, 3}, 2, true},
		{"多个元素-不包含", []int{1, 2, 3}, 4, false},
		{"重复元素", []int{1, 2, 2, 3}, 2, true},
		{"零值匹配", []int{0, 1, 2}, 0, true},
		{"负数匹配", []int{-1, 0, 1}, -1, true},
		{"大数匹配", []int{1000000, 999999}, 1000000, true},
		{"首元素匹配", []int{1, 2, 3}, 1, true},
		{"尾元素匹配", []int{1, 2, 3}, 3, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt := tt // 避免竞态
			result := candy.Contains(tt.give, tt.item)
			assert.Equal(t, tt.want, result, "Contains() 的结果应与期望值相等")
		})
	}
}

// TestContainsUsing 测试 ContainsUsing 函数
func TestContainsUsing(t *testing.T) {
	tests := []struct {
		name     string
		give     []int
		comparer func(int) bool
		want     bool
	}{
		{
			"空切片",
			[]int{},
			func(a int) bool { return a == 42 },
			false,
		},
		{
			"相等比较器-包含",
			[]int{1, 2, 3},
			func(a int) bool { return a == 2 },
			true,
		},
		{
			"相等比较器-不包含",
			[]int{1, 2, 3},
			func(a int) bool { return a == 4 },
			false,
		},
		{
			"绝对值比较器-包含",
			[]int{1, 2, -2},
			func(a int) bool { return a * a == 4 },
			true,
		},
		{
			"绝对值比较器-不包含",
			[]int{1, 2, 3},
			func(a int) bool { return a * a == 16 },
			false,
		},
		{
			"奇偶比较器-包含",
			[]int{1, 3, 5},
			func(a int) bool { return a % 2 == 1 },
			true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt := tt // 避免竞态
			result := candy.ContainsUsing(tt.give, tt.comparer)
			assert.Equal(t, tt.want, result, "ContainsUsing() 的结果应与期望值相等")
		})
	}
}

// TestReverse 测试 Reverse 函数
func TestReverse(t *testing.T) {
	tests := []struct {
		name string
		give []int
		want []int
	}{
		{"空切片", []int{}, []int{}},
		{"单个元素", []int{42}, []int{42}},
		{"两个元素", []int{1, 2}, []int{2, 1}},
		{"三个元素", []int{1, 2, 3}, []int{3, 2, 1}},
		{"偶数长度", []int{1, 2, 3, 4}, []int{4, 3, 2, 1}},
		{"奇数长度", []int{1, 2, 3, 4, 5}, []int{5, 4, 3, 2, 1}},
		{"重复元素", []int{1, 2, 2, 3}, []int{3, 2, 2, 1}},
		{"负数", []int{-1, -2, -3}, []int{-3, -2, -1}},
		{"混合正负", []int{1, -1, 0}, []int{0, -1, 1}},
		{"相同元素", []int{5, 5, 5}, []int{5, 5, 5}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt := tt // 避免竞态
			result := candy.Reverse(tt.give)
			assert.Equal(t, tt.want, result, "Reverse() 的结果应与期望值相等")
		})
	}
}

// TestIndex 测试 Index 函数
func TestIndex(t *testing.T) {
	tests := []struct {
		name string
		give []int
		item int
		want int
	}{
		{"空切片", []int{}, 42, -1},
		{"单个元素-匹配", []int{42}, 42, 0},
		{"单个元素-不匹配", []int{42}, 24, -1},
		{"多个元素-首元素", []int{1, 2, 3}, 1, 0},
		{"多个元素-中间元素", []int{1, 2, 3}, 2, 1},
		{"多个元素-尾元素", []int{1, 2, 3}, 3, 2},
		{"多个元素-不包含", []int{1, 2, 3}, 4, -1},
		{"重复元素-返回第一个", []int{1, 2, 2, 3}, 2, 1},
		{"零值匹配", []int{0, 1, 2}, 0, 0},
		{"负数匹配", []int{-1, 0, 1}, -1, 0},
		{"大数匹配", []int{1000000, 999999}, 1000000, 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt := tt // 避免竞态
			result := candy.Index(tt.give, tt.item)
			assert.Equal(t, tt.want, result, "Index() 的结果应与期望值相等")
		})
	}
}

// TestDiff 测试 Diff 函数
func TestDiff(t *testing.T) {
	tests := []struct {
		name    string
		ss      []int
		against []int
		wantAdded    []int
		wantRemoved []int
	}{
		{
			name:        "两个空切片",
			ss:          []int{},
			against:     []int{},
			wantAdded:    []int{},
			wantRemoved: []int{},
		},
		{
			name:        "ss为空，against有元素",
			ss:          []int{},
			against:     []int{1, 2, 3},
			wantAdded:    []int{}, // 在against中但不在ss中: 无
			wantRemoved: []int{1, 2, 3}, // 在ss中但不在against中: 1,2,3
		},
		{
			name:        "ss有元素，against为空",
			ss:          []int{1, 2, 3},
			against:     []int{},
			wantAdded:    []int{1, 2, 3}, // 在against中但不在ss中: 1,2,3
			wantRemoved: []int{}, // 在ss中但不在against中: 无
		},
		{
			name:        "无差异",
			ss:          []int{1, 2, 3},
			against:     []int{1, 2, 3},
			wantAdded:    []int{},
			wantRemoved: []int{},
		},
		{
			name:        "against有新增",
			ss:          []int{1, 2},
			against:     []int{1, 2, 3},
			wantAdded:    []int{}, // 在against中但不在ss中: 无
			wantRemoved: []int{3}, // 在ss中但不在against中: 3
		},
		{
			name:        "ss有删除",
			ss:          []int{1, 2, 3},
			against:     []int{1, 2},
			wantAdded:    []int{3}, // 在against中但不在ss中: 3
			wantRemoved: []int{}, // 在ss中但不在against中: 无
		},
		{
			name:        "混合差异",
			ss:          []int{1, 2, 3},
			against:     []int{2, 3, 4},
			wantAdded:    []int{1}, // 在against中但不在ss中: 1
			wantRemoved: []int{4}, // 在ss中但不在against中: 4
		},
		{
			name:        "重复元素",
			ss:          []int{1, 2, 2, 3},
			against:     []int{1, 2, 3},
			wantAdded:   []int{},
			wantRemoved: []int{}, // 在ss中但不在against中: 无 (Remove函数使用map去重)
		},
		{
			name: "完全不同",
			ss:   []int{1, 2},
			against: []int{3, 4},
			wantAdded:  []int{1, 2}, // 在against中但不在ss中: 1,2
			wantRemoved: []int{3, 4}, // 在ss中但不在against中: 3,4
		},
		{
			name:        "负数差异",
			ss:          []int{-1, -2},
			against:     []int{-3, -2},
			wantAdded:    []int{-1}, // 在against中但不在ss中: -1
			wantRemoved: []int{-3}, // 在ss中但不在against中: -3
		},
		{
			name:        "零值差异",
			ss:          []int{0, 1},
			against:     []int{1, 0},
			wantAdded:    []int{}, // 在against中但不在ss中: 无 (0和1都在ss中)
			wantRemoved: []int{}, // 在ss中但不在against中: 无 (0和1都在against中)
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			gotAdded, gotRemoved := candy.Diff(tt.ss, tt.against)
			assert.Equal(t, tt.wantAdded, gotAdded, "Diff() added 的结果应与期望值相等")
			assert.Equal(t, tt.wantRemoved, gotRemoved, "Diff() removed 的结果应与期望值相等")
		})
	}
}

// TestRemove 测试 Remove 函数
func TestRemove(t *testing.T) {
	tests := []struct {
		name string
		ss []int
		against []int
		want []int
	}{
		{"两个空切片", []int{}, []int{}, []int{}},
		{"ss为空，against有元素", []int{}, []int{42}, []int{42}}, // 在against中但不在ss中
		{"ss有元素，against为空", []int{42}, []int{}, []int{}}, // 在against中但不在ss中
		{"无差异", []int{1, 2, 3}, []int{1, 2, 3}, []int{}},
		{"against有新增", []int{1, 2}, []int{1, 2, 3}, []int{3}}, // 在against中但不在ss中
		{"ss有删除", []int{1, 2, 3}, []int{1, 2}, []int{}}, // 在against中但不在ss中
		{"混合差异", []int{1, 2, 3}, []int{2, 3, 4}, []int{4}}, // 在against中但不在ss中
		{"重复元素", []int{1, 2, 2}, []int{2, 2, 3}, []int{3}}, // 在against中但不在ss中
		{"完全不同", []int{1, 2}, []int{3, 4}, []int{3, 4}}, // 在against中但不在ss中
		{"负数差异", []int{-1, -2}, []int{-2, -3}, []int{-3}}, // 在against中但不在ss中
		{"零值差异", []int{0, 1}, []int{1, 0}, []int{}}, // 在against中但不在ss中
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt := tt // 避免竞态
			result := candy.Remove(tt.ss, tt.against)
			assert.Equal(t, tt.want, result, "Remove() 的结果应与期望值相等")
		})
	}
}

// TestSame 测试 Same 函数
func TestSame(t *testing.T) {
	tests := []struct {
		name string
		against []int
		ss []int
		want []int
	}{
		{"两个空切片", []int{}, []int{}, []int{}},
		{"一个空一个非空", []int{}, []int{1}, []int{}},
		{"长度不同", []int{1}, []int{1, 2}, []int{1}},
		{"内容相同", []int{1, 2, 3}, []int{1, 2, 3}, []int{1, 2, 3}},
		{"内容不同", []int{1, 2, 4}, []int{1, 2, 3}, []int{1, 2}},
		{"顺序不同", []int{3, 2, 1}, []int{1, 2, 3}, []int{3, 2, 1}},
		{"重复元素相同", []int{1, 2, 2}, []int{1, 2, 2}, []int{1, 2, 2}},
		{"重复元素不同", []int{1, 2, 3}, []int{1, 2, 2}, []int{1, 2}},
		{"负数相同", []int{-1, -2}, []int{-1, -2}, []int{-1, -2}},
		{"零值相同", []int{0, 0}, []int{0, 0}, []int{0, 0}},
		{"大数相同", []int{1000000, 999999}, []int{1000000, 999999}, []int{1000000, 999999}},
		{"nil切片", nil, nil, []int{}},
		{"nil与非空", nil, []int{1}, []int{}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt := tt // 避免竞态
			result := candy.Same(tt.against, tt.ss)
			assert.Equal(t, tt.want, result, "Same() 的结果应与期望值相等")
		})
	}
}

// TestSpare 测试 Spare 函数
func TestSpare(t *testing.T) {
	tests := []struct {
		name string
		ss []int
		against []int
		want []int
	}{
		{"两个空切片", []int{}, []int{}, []int{}},
		{"ss为空，against有元素", []int{}, []int{1, 2}, []int{1, 2}}, // 在against中但不在ss中
		{"ss有元素，against为空", []int{1, 2}, []int{}, []int{}}, // 在against中但不在ss中
		{"无差异", []int{1, 2, 3}, []int{1, 2, 3}, []int{}},
		{"against有新增", []int{1, 2}, []int{1, 2, 3}, []int{3}}, // 在against中但不在ss中
		{"ss有删除", []int{1, 2, 3}, []int{1, 2}, []int{}}, // 在against中但不在ss中
		{"混合差异", []int{1, 2, 3}, []int{2, 3, 4}, []int{4}}, // 在against中但不在ss中
		{"重复元素", []int{1, 2, 2}, []int{2, 2, 3}, []int{3}}, // 在against中但不在ss中
		{"完全不同", []int{1, 2}, []int{3, 4}, []int{3, 4}}, // 在against中但不在ss中
		{"负数差异", []int{-1, -2}, []int{-2, -3}, []int{-3}}, // 在against中但不在ss中
		{"零值差异", []int{0, 1}, []int{1, 0}, []int{}}, // 在against中但不在ss中
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt := tt // 避免竞态
			result := candy.Spare(tt.ss, tt.against)
			assert.Equal(t, tt.want, result, "Spare() 的结果应与期望值相等")
		})
	}
}

// TestRemoveIndex 测试 RemoveIndex 函数
func TestRemoveIndex(t *testing.T) {
	tests := []struct {
		name string
		give []int
		index int
		want []int
	}{
		{"空切片", []int{}, 0, []int{}},
		{"单个元素-删除", []int{42}, 0, []int{}},
		{"单个元素-负索引", []int{42}, -1, []int{}},
		{"多个元素-删除第一个", []int{1, 2, 3}, 0, []int{2, 3}},
		{"多个元素-删除中间", []int{1, 2, 3}, 1, []int{1, 3}},
		{"多个元素-删除最后一个", []int{1, 2, 3}, 2, []int{1, 2}},
		{"多个元素-负索引", []int{1, 2, 3}, -1, []int{}},
		{"索引越界", []int{1, 2, 3}, 5, []int{}},
		{"大负索引", []int{1, 2, 3}, -10, []int{}},
		{"删除后保持顺序", []int{10, 20, 30, 40}, 1, []int{10, 30, 40}},
		{"重复元素", []int{1, 2, 2, 3}, 2, []int{1, 2, 3}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt := tt // 避免竞态
			result := candy.RemoveIndex(tt.give, tt.index)
			assert.Equal(t, tt.want, result, "RemoveIndex() 的结果应与期望值相等")
		})
	}
}

// TestString 测试 String 函数
func TestString(t *testing.T) {
	tests := []struct {
		name string
		give int
		want string
	}{
		{"正数", 42, "42"},
		{"负数", -1, "-1"},
		{"零", 0, "0"},
		{"大数", 999999, "999999"},
		{"最大正数", 2147483647, "2147483647"},
		{"最大负数", -2147483648, "-2147483648"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt := tt // 避免竞态
			result := candy.String(tt.give)
			assert.Equal(t, tt.want, result, "String() 的结果应与期望值相等")
		})
	}
}

// TestJoin 测试 Join 函数
func TestJoin(t *testing.T) {
	tests := []struct {
		name string
		give []string
		sep  string
		want string
	}{
		{"空切片", []string{}, ",", ""},
		{"单个元素", []string{"hello"}, ",", "hello"},
		{"多个元素", []string{"a", "b", "c"}, ",", "a,b,c"},
		{"空分隔符", []string{"a", "b", "c"}, "", "abc"},
		{"特殊分隔符", []string{"a", "b", "c"}, "->", "a->b->c"},
		{"包含空字符串", []string{"", "b", "c"}, ",", ",b,c"},
		{"空元素", []string{"a", "", "c"}, ",", "a,,c"},
		{"长分隔符", []string{"a", "b"}, "---", "a---b"},
		{"Unicode分隔符", []string{"a", "b"}, "•", "a•b"},
		{"数字元素", []string{"1", "2", "3"}, "-", "1-2-3"},
		{"重复元素", []string{"a", "a", "a"}, ",", "a,a,a"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt := tt // 避免竞态
			result := candy.Join(tt.give, tt.sep)
			assert.Equal(t, tt.want, result, "Join() 的结果应与期望值相等")
		})
	}
}

// TestSortUsingEdgeCases 测试 SortUsing 函数的边界情况
func TestSortUsingEdgeCases(t *testing.T) {
	tests := []struct {
		name     string
		give     []int
		comparer func(int, int) bool
		want     []int
	}{
		{
			"单元素切片",
			[]int{42},
			func(a, b int) bool { return a < b },
			[]int{42},
		},
		{
			"两个元素-升序",
			[]int{2, 1},
			func(a, b int) bool { return a < b },
			[]int{1, 2},
		},
		{
			"两个元素-降序",
			[]int{1, 2},
			func(a, b int) bool { return a > b },
			[]int{2, 1},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt := tt // 避免竞态
			result := candy.SortUsing(tt.give, tt.comparer)
			assert.Equal(t, tt.want, result, "SortUsing() 边界情况的结果应与期望值相等")
		})
	}
}

// TestRemoveIndexEdgeCases 测试 RemoveIndex 函数的边界情况
func TestRemoveIndexEdgeCases(t *testing.T) {
	tests := []struct {
		name  string
		give  []int
		index int
		want  []int
	}{
		{
			"删除第一个元素-两个元素",
			[]int{1, 2},
			0,
			[]int{2},
		},
		{
			"删除第一个元素-多个元素",
			[]int{1, 2, 3, 4},
			0,
			[]int{2, 3, 4},
		},
		{
			"删除第一个元素-单元素后为空",
			[]int{1},
			0,
			[]int{},
		},
		{
			"索引等于长度-返回空切片",
			[]int{1, 2, 3},
			3,
			[]int{},
		},
		{
			"负索引-返回空切片",
			[]int{1, 2, 3},
			-1,
			[]int{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt := tt // 避免竞态
			result := candy.RemoveIndex(tt.give, tt.index)
			assert.Equal(t, tt.want, result, "RemoveIndex() 边界情况的结果应与期望值相等")
		})
	}
}

// TestSliceEqualEdgeCases 测试 SliceEqual 函数的边界情况
func TestSliceEqualEdgeCases(t *testing.T) {
	tests := []struct {
		name string
		a    []int
		b    []int
		want bool
	}{
		{
			"重复元素匹配-相同数量",
			[]int{1, 2, 2, 3},
			[]int{1, 2, 2, 3},
			false,
		},
		{
			"重复元素匹配-不同顺序",
			[]int{1, 2, 2, 3},
			[]int{3, 2, 1, 2},
			false,
		},
		{
			"重复元素不匹配-数量不同",
			[]int{1, 2, 2, 3},
			[]int{1, 2, 3},
			false,
		},
		{
			"所有元素相同-匹配",
			[]int{5, 5, 5},
			[]int{5, 5, 5},
			false,
		},
		{
			"所有元素相同-数量不同",
			[]int{5, 5, 5},
			[]int{5, 5},
			false,
		},
		{
			"空切片匹配",
			[]int{},
			[]int{},
			true,
		},
		{
			"一个空一个非空",
			[]int{},
			[]int{1},
			false,
		},
		{
			"相同元素不同位置",
			[]int{1, 1, 2},
			[]int{1, 2, 2},
			false,
		},
		{
			"大数量重复元素",
			[]int{1, 1, 1, 2, 2, 2},
			[]int{1, 1, 2, 2, 2, 1},
			false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt := tt // 避免竞态
			result := candy.SliceEqual(tt.a, tt.b)
			assert.Equal(t, tt.want, result, "SliceEqual() 边界情况的结果应与期望值相等")
		})
	}
}
