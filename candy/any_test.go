package candy

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestAny 测试 Any 函数的功能
func TestAny(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		funcTest func(int) bool
		expected bool
	}{
		{
			name:     "空切片返回 false",
			input:    []int{},
			funcTest: func(i int) bool { return i > 0 },
			expected: false,
		},
		{
			name:     "有一个元素满足条件返回 true",
			input:    []int{1, 2, 3},
			funcTest: func(i int) bool { return i == 2 },
			expected: true,
		},
		{
			name:     "所有元素都满足条件返回 true",
			input:    []int{2, 4, 6},
			funcTest: func(i int) bool { return i%2 == 0 },
			expected: true,
		},
		{
			name:     "没有元素满足条件返回 false",
			input:    []int{1, 3, 5},
			funcTest: func(i int) bool { return i%2 == 0 },
			expected: false,
		},
		{
			name:     "单元素切片满足条件",
			input:    []int{42},
			funcTest: func(i int) bool { return i > 0 },
			expected: true,
		},
		{
			name:     "单元素切片不满足条件",
			input:    []int{42},
			funcTest: func(i int) bool { return i < 0 },
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Any(tt.input, tt.funcTest)
			assert.Equal(t, tt.expected, result)
		})
	}
}

// TestAnyWithString 测试 Any 函数对字符串类型的处理
func TestAnyWithString(t *testing.T) {
	tests := []struct {
		name     string
		input    []string
		funcTest func(string) bool
		expected bool
	}{
		{
			name:     "包含空字符串",
			input:    []string{"hello", "", "world"},
			funcTest: func(s string) bool { return s == "" },
			expected: true,
		},
		{
			name:     "不包含空字符串",
			input:    []string{"hello", "world"},
			funcTest: func(s string) bool { return s == "" },
			expected: false,
		},
		{
			name:     "包含特定前缀",
			input:    []string{"apple", "banana", "orange"},
			funcTest: func(s string) bool { return len(s) > 5 },
			expected: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Any(tt.input, tt.funcTest)
			assert.Equal(t, tt.expected, result)
		})
	}
}

// TestAnyWithStruct 测试 Any 函数对结构体类型的处理
func TestAnyWithStruct(t *testing.T) {
	type Person struct {
		Name string
		Age  int
	}

	people := []Person{
		{"Alice", 25},
		{"Bob", 30},
		{"Charlie", 35},
	}

	// 测试查找年龄大于30的人
	result := Any(people, func(p Person) bool { return p.Age > 30 })
	assert.True(t, result)

	// 测试查找年龄小于20的人
	result = Any(people, func(p Person) bool { return p.Age < 20 })
	assert.False(t, result)
}

// BenchmarkAny 性能基准测试
func BenchmarkAny(b *testing.B) {
	data := make([]int, 1000)
	for i := range data {
		data[i] = i
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Any(data, func(n int) bool { return n > 500 })
	}
}