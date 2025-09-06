package candy

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	
)

// TestUnique 测试 Unique 函数
func TestUnique(t *testing.T) {
	t.Parallel()

	// 测试整数类型切片去重
	t.Run("整数类型切片去重", func(t *testing.T) {
		t.Parallel()
		tests := []struct {
			give []int
			want []int
		}{
			{[]int{1, 2, 3, 2, 1}, []int{1, 2, 3}},
			{[]int{5, 5, 5, 5, 5}, []int{5}},
			{[]int{10, 20, 30, 40, 50}, []int{10, 20, 30, 40, 50}},
			{[]int{1, 3, 2, 4, 3, 2, 1}, []int{1, 3, 2, 4}},
		}

		for _, tt := range tests {
			tt := tt
			t.Run(fmt.Sprintf("%v", tt.give), func(t *testing.T) {
				t.Parallel()
				got := Unique(tt.give)
				assert.Equal(t, tt.want, got, "整数切片去重结果应匹配")
			})
		}
	})

	// 测试浮点数类型切片去重
	t.Run("浮点数类型切片去重", func(t *testing.T) {
		t.Parallel()
		tests := []struct {
			give []float64
			want []float64
		}{
			{[]float64{1.1, 2.2, 3.3, 2.2, 1.1}, []float64{1.1, 2.2, 3.3}},
			{[]float64{5.5, 5.5, 5.5}, []float64{5.5}},
			{[]float64{1.0, 2.0, 3.0, 4.0}, []float64{1.0, 2.0, 3.0, 4.0}},
			{[]float64{0.1, 0.2, 0.1, 0.3}, []float64{0.1, 0.2, 0.3}},
		}

		for _, tt := range tests {
			tt := tt
			t.Run(fmt.Sprintf("%v", tt.give), func(t *testing.T) {
				t.Parallel()
				got := Unique(tt.give)
				assert.Equal(t, tt.want, got, "浮点数切片去重结果应匹配")
			})
		}
	})

	// 测试字符串类型切片去重
	t.Run("字符串类型切片去重", func(t *testing.T) {
		t.Parallel()
		tests := []struct {
			give []string
			want []string
		}{
			{[]string{"a", "b", "c", "b", "a"}, []string{"a", "b", "c"}},
			{[]string{"hello", "hello", "hello"}, []string{"hello"}},
			{[]string{"apple", "banana", "cherry"}, []string{"apple", "banana", "cherry"}},
			{[]string{"go", "python", "go", "java", "python"}, []string{"go", "python", "java"}},
		}

		for _, tt := range tests {
			tt := tt
			t.Run(fmt.Sprintf("%v", tt.give), func(t *testing.T) {
				t.Parallel()
				got := Unique(tt.give)
				assert.Equal(t, tt.want, got, "字符串切片去重结果应匹配")
			})
		}
	})

	// 测试边界情况
	t.Run("边界情况", func(t *testing.T) {
		t.Parallel()
		tests := []struct {
			name string
			give []int
			want []int
		}{
			{"空切片", []int{}, []int{}},
			{"单元素切片", []int{42}, []int{42}},
			{"全部相同元素", []int{7, 7, 7, 7}, []int{7}},
			{"无重复元素", []int{1, 2, 3, 4}, []int{1, 2, 3, 4}},
			{"连续重复", []int{1, 1, 2, 2, 3, 3}, []int{1, 2, 3}},
			{"间隔重复", []int{1, 2, 1, 3, 2, 1}, []int{1, 2, 3}},
		}

		for _, tt := range tests {
			tt := tt
			t.Run(tt.name, func(t *testing.T) {
				t.Parallel()
				got := Unique(tt.give)
				assert.Equal(t, tt.want, got, "边界情况处理应正确")
			})
		}
	})

	// 测试保留原始顺序
	t.Run("保留原始顺序", func(t *testing.T) {
		t.Parallel()
		// 测试去重后的元素保持原始出现顺序
		original := []int{3, 1, 4, 1, 5, 9, 2, 6, 5, 3}
		result := Unique(original)
		expected := []int{3, 1, 4, 5, 9, 2, 6}
		assert.Equal(t, expected, result, "去重后应保留原始顺序")
	})

	// 测试不修改原切片
	t.Run("不修改原切片", func(t *testing.T) {
		t.Parallel()
		original := []int{1, 2, 2, 3}
		originalCopy := make([]int, len(original))
		copy(originalCopy, original)

		result := Unique(original)

		// 确保原切片未被修改
		assert.Equal(t, originalCopy, original, "原切片应保持不变")
		// 确保返回的是新切片
		assert.NotSame(t, &original[0], &result[0], "应返回新切片")
	})
}

// BenchmarkUnique 测试 Unique 函数的基准性能
func BenchmarkUnique(b *testing.B) {
	// 小数据集基准测试
	b.Run("小数据集", func(b *testing.B) {
		data := []int{1, 2, 3, 2, 1, 4, 5, 4, 3}
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			Unique(data)
		}
	})

	// 中等数据集基准测试
	b.Run("中等数据集", func(b *testing.B) {
		data := make([]int, 1000)
		for i := 0; i < 1000; i++ {
			data[i] = i % 100 // 创建有重复的数据
		}
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			Unique(data)
		}
	})

	// 大数据集基准测试
	b.Run("大数据集", func(b *testing.B) {
		data := make([]int, 100000)
		for i := 0; i < 100000; i++ {
			data[i] = i % 1000 // 创建有重复的数据
		}
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			Unique(data)
		}
	})

	// 无重复数据集基准测试
	b.Run("无重复数据集", func(b *testing.B) {
		data := make([]int, 1000)
		for i := 0; i < 1000; i++ {
			data[i] = i // 创建无重复的数据
		}
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			Unique(data)
		}
	})
}

