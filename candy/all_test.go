// Package all 提供 All 函数的单元测试
package candy

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestAllInt 测试 All 函数对 int 类型的支持
func TestAllInt(t *testing.T) {
	tests := []struct {
		name      string
		input     []int
		predicate func(int) bool
		expected  bool
	}{
		{
			name:      "空切片",
			input:     []int{},
			predicate: func(x int) bool { return x > 0 },
			expected:  true,
		},
		{
			name:      "nil 切片",
			input:     nil,
			predicate: func(x int) bool { return x > 0 },
			expected:  true,
		},
		{
			name:      "所有元素都满足条件",
			input:     []int{2, 4, 6, 8},
			predicate: func(x int) bool { return x%2 == 0 },
			expected:  true,
		},
		{
			name:      "部分元素满足条件",
			input:     []int{1, 2, 3, 4},
			predicate: func(x int) bool { return x%2 == 0 },
			expected:  false,
		},
		{
			name:      "所有元素都不满足条件",
			input:     []int{1, 3, 5, 7},
			predicate: func(x int) bool { return x%2 == 0 },
			expected:  false,
		},
		{
			name:      "单个元素满足条件",
			input:     []int{42},
			predicate: func(x int) bool { return x > 0 },
			expected:  true,
		},
		{
			name:      "单个元素不满足条件",
			input:     []int{-1},
			predicate: func(x int) bool { return x > 0 },
			expected:  false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// 执行 All 函数
			result := All(tt.input, tt.predicate)

			// 验证结果
			assert.Equal(t, tt.expected, result, "All() 的结果应与期望值相等")
		})
	}
}

// TestAllString 测试 All 函数对 string 类型的支持
func TestAllString(t *testing.T) {
	tests := []struct {
		name      string
		input     []string
		predicate func(string) bool
		expected  bool
	}{
		{
			name:      "字符串类型测试",
			input:     []string{"hello", "world", "golang"},
			predicate: func(s string) bool { return len(s) > 0 },
			expected:  true,
		},
		{
			name:      "字符串类型部分满足",
			input:     []string{"hello", "", "world"},
			predicate: func(s string) bool { return len(s) > 0 },
			expected:  false,
		},
		{
			name:      "字符串类型都不满足",
			input:     []string{"", "", ""},
			predicate: func(s string) bool { return len(s) > 0 },
			expected:  false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// 执行 All 函数
			result := All(tt.input, tt.predicate)

			// 验证结果
			assert.Equal(t, tt.expected, result, "All() 的结果应与期望值相等")
		})
	}
}

// TestAllFloat64 测试 All 函数对 float64 类型的支持
func TestAllFloat64(t *testing.T) {
	tests := []struct {
		name      string
		input     []float64
		predicate func(float64) bool
		expected  bool
	}{
		{
			name:      "浮点数类型测试",
			input:     []float64{1.1, 2.2, 3.3},
			predicate: func(x float64) bool { return x > 1.0 },
			expected:  true,
		},
		{
			name:      "浮点数类型部分满足",
			input:     []float64{0.5, 1.5, 2.5},
			predicate: func(x float64) bool { return x > 1.0 },
			expected:  false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// 执行 All 函数
			result := All(tt.input, tt.predicate)

			// 验证结果
			assert.Equal(t, tt.expected, result, "All() 的结果应与期望值相等")
		})
	}
}

// TestAllShortCircuit 测试 All 函数的短路行为
func TestAllShortCircuit(t *testing.T) {
	// 创建一个计数器来验证函数是否被短路调用
	callCount := 0
	predicate := func(x int) bool {
		callCount++
		return x > 0
	}

	// 测试场景：第二个元素不满足条件，应该不会调用第三个元素
	input := []int{1, -1, 2}
	expected := false

	result := All(input, predicate)

	// 验证结果
	assert.Equal(t, expected, result, "All() 应该返回 false")

	// 验证短路行为：只应该调用前两个元素的判断函数
	assert.Equal(t, 2, callCount, "应该只调用前两个元素的判断函数")
}

// TestAllWithStruct 测试 All 函数对结构体类型的支持
func TestAllWithStruct(t *testing.T) {
	type Person struct {
		Name string
		Age  int
	}

	people := []Person{
		{"Alice", 25},
		{"Bob", 30},
		{"Charlie", 35},
	}

	// 测试所有人员都成年
	result := All(people, func(p Person) bool {
		return p.Age >= 18
	})
	assert.True(t, result, "所有人员都应该成年")

	// 测试所有人员都有特定名字长度
	result = All(people, func(p Person) bool {
		return len(p.Name) >= 3
	})
	assert.True(t, result, "所有人员的名字长度都应该大于等于3")
}

// TestAllWithPointers 测试 All 函数对指针类型的支持
func TestAllWithPointers(t *testing.T) {
	values := []*int{new(int), new(int), new(int)}

	// 设置值
	*values[0] = 10
	*values[1] = 20
	*values[2] = 30

	result := All(values, func(p *int) bool {
		return p != nil && *p > 5
	})
	assert.True(t, result, "所有指针都应该指向大于5的值")
}

// BenchmarkAll 性能基准测试
func BenchmarkAll(b *testing.B) {
	data := make([]int, 1000)
	for i := range data {
		data[i] = i + 1
	}

	predicate := func(x int) bool { return x > 0 }

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		All(data, predicate)
	}
}

// BenchmarkAllParallel 并发性能基准测试
func BenchmarkAllParallel(b *testing.B) {
	data := make([]int, 1000)
	for i := range data {
		data[i] = i + 1
	}

	predicate := func(x int) bool { return x > 0 }

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			All(data, predicate)
		}
	})
}
