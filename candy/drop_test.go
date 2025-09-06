// Package drop 提供 Drop 函数的测试用例
package candy

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestDrop 测试 Drop 函数的基本功能
func TestDrop(t *testing.T) {
	tests := []struct {
		name string
		give []int
		n    int
		want []int
	}{
		{
			name: "正常情况",
			give: []int{1, 2, 3, 4, 5},
			n:    2,
			want: []int{3, 4, 5},
		},
		{
			name: "n为0",
			give: []int{1, 2, 3},
			n:    0,
			want: []int{1, 2, 3},
		},
		{
			name: "n为负数",
			give: []int{1, 2, 3},
			n:    -1,
			want: []int{1, 2, 3},
		},
		{
			name: "n等于切片长度",
			give: []int{1, 2, 3},
			n:    3,
			want: []int{},
		},
		{
			name: "n大于切片长度",
			give: []int{1, 2, 3},
			n:    5,
			want: []int{},
		},
		{
			name: "空切片",
			give: []int{},
			n:    3,
			want: []int{},
		},
		{
			name: "单元素切片",
			give: []int{1},
			n:    1,
			want: []int{},
		},
		{
			name: "单元素切片不丢弃",
			give: []int{1},
			n:    0,
			want: []int{1},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			got := Drop(tt.give, tt.n)
			assert.Equal(t, tt.want, got, "Drop() 的结果应与期望值相等")
		})
	}
}

// TestDropString 测试 Drop 函数对字符串切片的处理
func TestDropString(t *testing.T) {
	tests := []struct {
		name string
		give []string
		n    int
		want []string
	}{
		{
			name: "字符串切片",
			give: []string{"a", "b", "c", "d"},
			n:    2,
			want: []string{"c", "d"},
		},
		{
			name: "空字符串切片",
			give: []string{},
			n:    1,
			want: []string{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			got := Drop(tt.give, tt.n)
			assert.Equal(t, tt.want, got, "Drop() 对字符串切片的处理应正确")
		})
	}
}

// TestDropFloat 测试 Drop 函数对浮点数切片的处理
func TestDropFloat(t *testing.T) {
	tests := []struct {
		name string
		give []float64
		n    int
		want []float64
	}{
		{
			name: "浮点数切片",
			give: []float64{1.1, 2.2, 3.3, 4.4},
			n:    2,
			want: []float64{3.3, 4.4},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			got := Drop(tt.give, tt.n)
			assert.Equal(t, tt.want, got, "Drop() 对浮点数切片的处理应正确")
		})
	}
}

// TestDropOriginalSliceUnmodified 测试 Drop 函数不修改原切片
func TestDropOriginalSliceUnmodified(t *testing.T) {
	original := []int{1, 2, 3, 4, 5}
	originalCopy := make([]int, len(original))
	copy(originalCopy, original)

	result := Drop(original, 2)

	// 验证原切片未被修改
	assert.Equal(t, originalCopy, original, "Drop() 不应修改原切片")

	// 验证结果正确
	assert.Equal(t, []int{3, 4, 5}, result, "Drop() 的结果应正确")
}

// BenchmarkDrop 性能基准测试
func BenchmarkDrop(b *testing.B) {
	data := make([]int, 1000)
	for i := range data {
		data[i] = i
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Drop(data, 500)
	}
}
