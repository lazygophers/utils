// Package candy 提供了 Go 语言中常用的语法糖函数和工具方法
package candy

import (
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestFilter 测试 Filter 函数的基本功能
func TestFilter(t *testing.T) {
	t.Parallel()

	// 定义测试结构体
	type Person struct {
		Name string
		Age  int
	}

	tests := []struct {
		name     string
		input    any
		filter   func(any) bool
		expected any
	}{
		{
			name:  "整数切片过滤偶数",
			input: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
			filter: func(x any) bool {
				n := x.(int)
				return n%2 == 0
			},
			expected: []int{2, 4, 6, 8, 10},
		},
		{
			name:  "整数切片过滤大于5的数",
			input: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
			filter: func(x any) bool {
				n := x.(int)
				return n > 5
			},
			expected: []int{6, 7, 8, 9, 10},
		},
		{
			name:  "字符串切片过滤包含特定字符的字符串",
			input: []string{"apple", "banana", "cherry", "date", "elderberry"},
			filter: func(x any) bool {
				s := x.(string)
				return len(s) > 5
			},
			expected: []string{"banana", "cherry", "elderberry"},
		},
		{
			name:  "字符串切片过滤以a开头的字符串",
			input: []string{"apple", "banana", "cherry", "date", "elderberry"},
			filter: func(x any) bool {
				s := x.(string)
				return len(s) > 0 && s[0] == 'a'
			},
			expected: []string{"apple"},
		},
		{
			name: "结构体切片过滤年龄大于18的人",
			input: []Person{
				{Name: "Alice", Age: 25},
				{Name: "Bob", Age: 17},
				{Name: "Charlie", Age: 30},
				{Name: "David", Age: 16},
			},
			filter: func(x any) bool {
				p := x.(Person)
				return p.Age > 18
			},
			expected: []Person{
				{Name: "Alice", Age: 25},
				{Name: "Charlie", Age: 30},
			},
		},
		{
			name: "结构体切片过滤名字以C开头的人",
			input: []Person{
				{Name: "Alice", Age: 25},
				{Name: "Bob", Age: 17},
				{Name: "Charlie", Age: 30},
				{Name: "David", Age: 16},
			},
			filter: func(x any) bool {
				p := x.(Person)
				return len(p.Name) > 0 && p.Name[0] == 'C'
			},
			expected: []Person{
				{Name: "Charlie", Age: 30},
			},
		},
		{
			name:  "指针切片过滤非nil指针",
			input: []*int{nil, intPtr(1), nil, intPtr(2), intPtr(3), nil},
			filter: func(x any) bool {
				p := x.(*int)
				return p != nil
			},
			expected: []*int{intPtr(1), intPtr(2), intPtr(3)},
		},
		{
			name:  "指针切片过滤值大于1的指针",
			input: []*int{intPtr(1), intPtr(2), intPtr(3), intPtr(4), intPtr(5)},
			filter: func(x any) bool {
				p := x.(*int)
				return p != nil && *p > 1
			},
			expected: []*int{intPtr(2), intPtr(3), intPtr(4), intPtr(5)},
		},
		{
			name:     "空切片过滤",
			input:    []int{},
			filter:   func(x any) bool { return x.(int)%2 == 0 },
			expected: []int{},
		},
		{
			name:     "nil切片过滤",
			input:    ([]int)(nil),
			filter:   func(x any) bool { return x.(int)%2 == 0 },
			expected: []int{},
		},
		{
			name:  "浮点数切片过滤大于0.5的数",
			input: []float64{0.1, 0.2, 0.3, 0.6, 0.7, 0.8, 0.9},
			filter: func(x any) bool {
				n := x.(float64)
				return n > 0.5
			},
			expected: []float64{0.6, 0.7, 0.8, 0.9},
		},
		{
			name:  "布尔切片过滤true值",
			input: []bool{true, false, true, false, true},
			filter: func(x any) bool {
				b := x.(bool)
				return b
			},
			expected: []bool{true, true, true},
		},
		{
			name:  "复杂类型：interface切片过滤字符串类型",
			input: []any{"hello", 123, true, "world", 456, false},
			filter: func(x any) bool {
				_, ok := x.(string)
				return ok
			},
			expected: []any{"hello", "world"},
		},
	}

	for _, tt := range tests {
		tt := tt // 避免竞态条件
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			// 使用泛型函数测试
			switch input := tt.input.(type) {
			case []int:
				result := Filter(input, func(n int) bool {
					return tt.filter(n)
				})
				assert.Equal(t, tt.expected, result, "整数切片过滤结果不匹配")
			case []string:
				result := Filter(input, func(s string) bool {
					return tt.filter(s)
				})
				assert.Equal(t, tt.expected, result, "字符串切片过滤结果不匹配")
			case []Person:
				result := Filter(input, func(p Person) bool {
					return tt.filter(p)
				})
				assert.Equal(t, tt.expected, result, "结构体切片过滤结果不匹配")
			case []*int:
				result := Filter(input, func(p *int) bool {
					return tt.filter(p)
				})
				assert.Equal(t, tt.expected, result, "指针切片过滤结果不匹配")
			case []float64:
				result := Filter(input, func(f float64) bool {
					return tt.filter(f)
				})
				assert.Equal(t, tt.expected, result, "浮点数切片过滤结果不匹配")
			case []bool:
				result := Filter(input, func(b bool) bool {
					return tt.filter(b)
				})
				assert.Equal(t, tt.expected, result, "布尔切片过滤结果不匹配")
			case []any:
				result := Filter(input, func(a any) bool {
					return tt.filter(a)
				})
				assert.Equal(t, tt.expected, result, "interface切片过滤结果不匹配")
			default:
				t.Fatalf("不支持的输入类型: %T", input)
			}
		})
	}
}

// TestFilterEdgeCases 测试 Filter 函数的边界情况
func TestFilterEdgeCases(t *testing.T) {
	t.Parallel()

	// 测试空切片
	t.Run("空切片", func(t *testing.T) {
		t.Parallel()
		result := Filter([]int{}, func(n int) bool { return n%2 == 0 })
		assert.Empty(t, result, "空切片过滤结果应该为空")
		assert.NotNil(t, result, "空切片过滤结果不应该为nil")
	})

	// 测试nil切片
	t.Run("nil切片", func(t *testing.T) {
		t.Parallel()
		var nilSlice []int
		result := Filter(nilSlice, func(n int) bool { return n%2 == 0 })
		assert.Empty(t, result, "nil切片过滤结果应该为空")
		assert.NotNil(t, result, "nil切片过滤结果不应该为nil")
	})

	// 测试所有元素都不满足条件
	t.Run("所有元素都不满足条件", func(t *testing.T) {
		t.Parallel()
		result := Filter([]int{1, 3, 5, 7, 9}, func(n int) bool { return n%2 == 0 })
		assert.Empty(t, result, "所有元素都不满足条件时结果应该为空")
		assert.NotNil(t, result, "结果不应该为nil")
	})

	// 测试所有元素都满足条件
	t.Run("所有元素都满足条件", func(t *testing.T) {
		t.Parallel()
		input := []int{2, 4, 6, 8, 10}
		result := Filter(input, func(n int) bool { return n%2 == 0 })
		assert.Equal(t, input, result, "所有元素都满足条件时结果应该等于输入")
	})

	// 测试大切片
	t.Run("大切片性能测试", func(t *testing.T) {
		t.Parallel()
		largeSlice := make([]int, 10000)
		for i := range largeSlice {
			largeSlice[i] = i
		}
		result := Filter(largeSlice, func(n int) bool { return n%2 == 0 })
		assert.Len(t, result, 5000, "大切片过滤结果长度应该正确")
	})
}

// TestFilterConcurrentSafety 测试 Filter 函数的并发安全性
func TestFilterConcurrentSafety(t *testing.T) {
	t.Parallel()

	// 准备测试数据：创建5000个元素的切片，确保并发测试的数据量足够大
	input := make([]int, 5000)
	for i := range input {
		input[i] = i
	}

	// 创建多个goroutine并发调用Filter，验证线程安全性
	var wg sync.WaitGroup
	concurrentCalls := 10

	for i := 0; i < concurrentCalls; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			result := Filter(input, func(n int) bool { return n%2 == 0 })
			assert.Len(t, result, 2500, "并发调用Filter结果应该正确")
		}()
	}

	wg.Wait()
}

// TestFilterTypeSafety 测试 Filter 函数的类型安全性
func TestFilterTypeSafety(t *testing.T) {
	t.Parallel()

	// 测试不同类型的过滤器函数
	t.Run("类型安全测试", func(t *testing.T) {
		t.Parallel()

		// 整数切片
		intSlice := []int{1, 2, 3, 4, 5}
		result := Filter(intSlice, func(n int) bool { return n > 2 })
		assert.Equal(t, []int{3, 4, 5}, result)

		// 字符串切片
		strSlice := []string{"a", "bb", "ccc", "dddd"}
		strResult := Filter(strSlice, func(s string) bool { return len(s) > 1 })
		assert.Equal(t, []string{"bb", "ccc", "dddd"}, strResult)
	})
}

// BenchmarkFilter 整数切片基准测试
func BenchmarkFilter(b *testing.B) {
	// 小切片基准测试
	b.Run("小切片(10元素)", func(b *testing.B) {
		data := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			Filter(data, func(n int) bool { return n%2 == 0 })
		}
	})

	// 中等切片基准测试
	b.Run("中等切片(1000元素)", func(b *testing.B) {
		data := make([]int, 1000)
		for i := range data {
			data[i] = i
		}
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			Filter(data, func(n int) bool { return n%2 == 0 })
		}
	})

	// 大切片基准测试
	b.Run("大切片(100000元素)", func(b *testing.B) {
		data := make([]int, 100000)
		for i := range data {
			data[i] = i
		}
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			Filter(data, func(n int) bool { return n%2 == 0 })
		}
	})
}

// BenchmarkFilterString 字符串切片基准测试
func BenchmarkFilterString(b *testing.B) {
	// 小字符串切片基准测试
	b.Run("小字符串切片(10元素)", func(b *testing.B) {
		data := []string{"a", "bb", "ccc", "dddd", "eeeee", "ffffff", "ggggggg", "hhhhhhhh", "iiiiiiiii", "jjjjjjjjjj"}
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			Filter(data, func(s string) bool { return len(s) > 3 })
		}
	})

	// 中等字符串切片基准测试
	b.Run("中等字符串切片(1000元素)", func(b *testing.B) {
		data := make([]string, 1000)
		for i := range data {
			data[i] = string(rune('a' + i%26))
		}
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			Filter(data, func(s string) bool { return len(s) > 0 })
		}
	})
}

// BenchmarkFilterStruct 结构体切片基准测试
func BenchmarkFilterStruct(b *testing.B) {
	type User struct {
		ID   int
		Name string
		Age  int
	}

	// 小结构体切片基准测试
	b.Run("小结构体切片(10元素)", func(b *testing.B) {
		data := make([]User, 10)
		for i := range data {
			data[i] = User{ID: i + 1, Name: "User", Age: 20 + i}
		}
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			Filter(data, func(u User) bool { return u.Age > 25 })
		}
	})

	// 中等结构体切片基准测试
	b.Run("中等结构体切片(1000元素)", func(b *testing.B) {
		data := make([]User, 1000)
		for i := range data {
			data[i] = User{ID: i + 1, Name: "User", Age: 20 + i%50}
		}
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			Filter(data, func(u User) bool { return u.Age > 25 })
		}
	})
}

// BenchmarkFilterDifferentSelectivity 不同选择性的基准测试
func BenchmarkFilterDifferentSelectivity(b *testing.B) {
	data := make([]int, 10000)
	for i := range data {
		data[i] = i
	}

	// 低选择性（大部分元素保留）
	b.Run("低选择性(保留90%)", func(b *testing.B) {
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			Filter(data, func(n int) bool { return n < 9000 })
		}
	})

	// 中等选择性（保留50%）
	b.Run("中等选择性(保留50%)", func(b *testing.B) {
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			Filter(data, func(n int) bool { return n%2 == 0 })
		}
	})

	// 高选择性（保留10%）
	b.Run("高选择性(保留10%)", func(b *testing.B) {
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			Filter(data, func(n int) bool { return n < 1000 })
		}
	})

	// 极高选择性（保留1%）
	b.Run("极高选择性(保留1%)", func(b *testing.B) {
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			Filter(data, func(n int) bool { return n < 100 })
		}
	})
}

// intPtr 是一个辅助函数，用于创建int指针
func intPtr(n int) *int {
	return &n
}