package candy

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestContains 测试 Contains 函数的各种场景
func TestContains(t *testing.T) {
	// 整数类型测试
	t.Run("整数类型", func(t *testing.T) {
		tests := []struct {
			name   string
			slice  []int
			target int
			want   bool
		}{
			{"包含元素", []int{1, 2, 3, 4, 5}, 3, true},
			{"不包含元素", []int{1, 2, 3, 4, 5}, 6, false},
			{"空切片", []int{}, 1, false},
			{"单元素-匹配", []int{42}, 42, true},
			{"单元素-不匹配", []int{42}, 24, false},
			{"重复元素", []int{1, 2, 2, 3, 2}, 2, true},
			{"负数", []int{-1, -2, -3}, -2, true},
			{"零值", []int{0, 1, 2}, 0, true},
		}

		for _, tt := range tests {
			tt := tt
			t.Run(tt.name, func(t *testing.T) {
				t.Parallel()
				got := Contains(tt.slice, tt.target)
				assert.Equal(t, tt.want, got, "Contains() 的结果应与期望值相等")
			})
		}
	})

	// 浮点数类型测试
	t.Run("浮点数类型", func(t *testing.T) {
		tests := []struct {
			name   string
			slice  []float64
			target float64
			want   bool
		}{
			{"包含元素", []float64{1.1, 2.2, 3.3}, 2.2, true},
			{"不包含元素", []float64{1.1, 2.2, 3.3}, 4.4, false},
			{"空切片", []float64{}, 1.1, false},
			{"科学计数法", []float64{1.5e10, 2.3e-5}, 1.5e10, true},
			{"精度测试 - 浮点数精确比较", []float64{0.1 + 0.2}, 0.3, true}, // 浮点数精度问题
		}

		for _, tt := range tests {
			tt := tt
			t.Run(tt.name, func(t *testing.T) {
				t.Parallel()
				got := Contains(tt.slice, tt.target)
				assert.Equal(t, tt.want, got, "Contains() 的结果应与期望值相等")
			})
		}
	})

	// 字符串类型测试
	t.Run("字符串类型", func(t *testing.T) {
		tests := []struct {
			name   string
			slice  []string
			target string
			want   bool
		}{
			{"包含元素", []string{"apple", "banana", "cherry"}, "banana", true},
			{"不包含元素", []string{"apple", "banana", "cherry"}, "orange", false},
			{"空切片", []string{}, "test", false},
			{"空字符串", []string{"", "hello", ""}, "", true},
			{"中文字符串", []string{"苹果", "香蕉", "橙子"}, "香蕉", true},
			{"特殊字符", []string{"a@b.com", "x#y", "test$"}, "x#y", true},
			{"Unicode字符", []string{"café", "naïve", "résumé"}, "naïve", true},
		}

		for _, tt := range tests {
			tt := tt
			t.Run(tt.name, func(t *testing.T) {
				t.Parallel()
				got := Contains(tt.slice, tt.target)
				assert.Equal(t, tt.want, got, "Contains() 的结果应与期望值相等")
			})
		}
	})

	// 边界情况测试
	t.Run("边界情况", func(t *testing.T) {
		tests := []struct {
			name   string
			slice  interface{}
			target interface{}
			want   bool
		}{
			{"大整数切片", []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 10, true},
			{"大字符串切片", []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j"}, "j", true},
			{"nil切片", ([]int)(nil), 1, false},
			{"首元素", []int{1, 2, 3}, 1, true},
			{"末元素", []int{1, 2, 3}, 3, true},
		}

		for _, tt := range tests {
			tt := tt
			t.Run(tt.name, func(t *testing.T) {
				t.Parallel()
				switch s := tt.slice.(type) {
				case []int:
					got := Contains(s, tt.target.(int))
					assert.Equal(t, tt.want, got, "Contains() 的结果应与期望值相等")
				case []string:
					got := Contains(s, tt.target.(string))
					assert.Equal(t, tt.want, got, "Contains() 的结果应与期望值相等")
				}
			})
		}
	})
}

// BenchmarkContains 性能测试
func BenchmarkContains(b *testing.B) {
	// 小切片测试
	b.Run("小切片-存在", func(b *testing.B) {
		slice := []int{1, 2, 3, 4, 5}
		target := 3
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			Contains(slice, target)
		}
	})

	b.Run("小切片-不存在", func(b *testing.B) {
		slice := []int{1, 2, 3, 4, 5}
		target := 99
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			Contains(slice, target)
		}
	})

	// 中等切片测试
	b.Run("中等切片-存在", func(b *testing.B) {
		slice := make([]int, 1000)
		for i := range slice {
			slice[i] = i
		}
		target := 500
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			Contains(slice, target)
		}
	})

	b.Run("中等切片-不存在", func(b *testing.B) {
		slice := make([]int, 1000)
		for i := range slice {
			slice[i] = i
		}
		target := 9999
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			Contains(slice, target)
		}
	})

	// 大切片测试
	b.Run("大切片-存在", func(b *testing.B) {
		slice := make([]int, 100000)
		for i := range slice {
			slice[i] = i
		}
		target := 50000
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			Contains(slice, target)
		}
	})

	b.Run("大切片-不存在", func(b *testing.B) {
		slice := make([]int, 100000)
		for i := range slice {
			slice[i] = i
		}
		target := 999999
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			Contains(slice, target)
		}
	})

	// 字符串切片测试
	b.Run("字符串切片", func(b *testing.B) {
		slice := []string{"apple", "banana", "cherry", "date", "elderberry"}
		target := "cherry"
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			Contains(slice, target)
		}
	})
}