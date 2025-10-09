// Package candy provides convenient utility functions
// filter_test.go tests the Filter function comprehensively
package candy

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestFilter tests the Filter function with various slice types and conditions
// 测试 Filter 函数处理各种切片类型和条件
func TestFilter(t *testing.T) {
	t.Run("IntegerSlice", func(t *testing.T) {
		t.Run("filter even numbers", func(t *testing.T) {
			input := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
			result := Filter(input, func(n int) bool { return n%2 == 0 })
			expected := []int{2, 4, 6, 8, 10}
			assert.Equal(t, expected, result)
		})

		t.Run("filter odd numbers", func(t *testing.T) {
			input := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
			result := Filter(input, func(n int) bool { return n%2 != 0 })
			expected := []int{1, 3, 5, 7, 9}
			assert.Equal(t, expected, result)
		})

		t.Run("filter greater than 5", func(t *testing.T) {
			input := []int{1, 3, 5, 7, 9, 11}
			result := Filter(input, func(n int) bool { return n > 5 })
			expected := []int{7, 9, 11}
			assert.Equal(t, expected, result)
		})

		t.Run("filter negative numbers", func(t *testing.T) {
			input := []int{-5, -3, 0, 2, 4}
			result := Filter(input, func(n int) bool { return n < 0 })
			expected := []int{-5, -3}
			assert.Equal(t, expected, result)
		})

		t.Run("all elements match", func(t *testing.T) {
			input := []int{2, 4, 6, 8, 10}
			result := Filter(input, func(n int) bool { return n%2 == 0 })
			assert.Equal(t, input, result)
		})

		t.Run("no elements match", func(t *testing.T) {
			input := []int{1, 3, 5, 7, 9}
			result := Filter(input, func(n int) bool { return n%2 == 0 })
			assert.Empty(t, result)
			assert.NotNil(t, result)
		})

		t.Run("empty slice", func(t *testing.T) {
			input := []int{}
			result := Filter(input, func(n int) bool { return n > 0 })
			assert.Empty(t, result)
			assert.NotNil(t, result)
		})

		t.Run("single element match", func(t *testing.T) {
			input := []int{42}
			result := Filter(input, func(n int) bool { return n > 0 })
			assert.Equal(t, []int{42}, result)
		})

		t.Run("single element no match", func(t *testing.T) {
			input := []int{42}
			result := Filter(input, func(n int) bool { return n < 0 })
			assert.Empty(t, result)
		})
	})

	t.Run("StringSlice", func(t *testing.T) {
		t.Run("filter by length", func(t *testing.T) {
			input := []string{"a", "ab", "abc", "abcd", "abcde"}
			result := Filter(input, func(s string) bool { return len(s) > 2 })
			expected := []string{"abc", "abcd", "abcde"}
			assert.Equal(t, expected, result)
		})

		t.Run("filter by prefix", func(t *testing.T) {
			input := []string{"apple", "banana", "apricot", "cherry", "avocado"}
			result := Filter(input, func(s string) bool { return len(s) > 0 && s[0] == 'a' })
			expected := []string{"apple", "apricot", "avocado"}
			assert.Equal(t, expected, result)
		})

		t.Run("filter non-empty strings", func(t *testing.T) {
			input := []string{"", "hello", "", "world", ""}
			result := Filter(input, func(s string) bool { return len(s) > 0 })
			expected := []string{"hello", "world"}
			assert.Equal(t, expected, result)
		})

		t.Run("filter by substring", func(t *testing.T) {
			input := []string{"golang", "python", "javascript", "typescript"}
			result := Filter(input, func(s string) bool {
				for i := 0; i < len(s)-2; i++ {
					if s[i:i+3] == "ang" {
						return true
					}
				}
				return false
			})
			expected := []string{"golang"}
			assert.Equal(t, expected, result)
		})

		t.Run("empty string slice", func(t *testing.T) {
			input := []string{}
			result := Filter(input, func(s string) bool { return len(s) > 0 })
			assert.Empty(t, result)
			assert.NotNil(t, result)
		})
	})

	t.Run("FloatSlice", func(t *testing.T) {
		t.Run("filter by value range", func(t *testing.T) {
			input := []float64{0.5, 1.5, 2.5, 3.5, 4.5}
			result := Filter(input, func(f float64) bool { return f >= 2.0 && f <= 4.0 })
			expected := []float64{2.5, 3.5}
			assert.Equal(t, expected, result)
		})

		t.Run("filter positive floats", func(t *testing.T) {
			input := []float64{-2.5, -1.0, 0.0, 1.5, 2.5}
			result := Filter(input, func(f float64) bool { return f > 0 })
			expected := []float64{1.5, 2.5}
			assert.Equal(t, expected, result)
		})
	})

	t.Run("StructSlice", func(t *testing.T) {
		type Person struct {
			Name string
			Age  int
		}

		t.Run("filter by age", func(t *testing.T) {
			input := []Person{
				{"Alice", 25},
				{"Bob", 17},
				{"Charlie", 30},
				{"David", 16},
				{"Eve", 22},
			}
			result := Filter(input, func(p Person) bool { return p.Age >= 18 })
			expected := []Person{
				{"Alice", 25},
				{"Charlie", 30},
				{"Eve", 22},
			}
			assert.Equal(t, expected, result)
		})

		t.Run("filter by name length", func(t *testing.T) {
			input := []Person{
				{"Al", 25},
				{"Bob", 30},
				{"Charlie", 35},
			}
			result := Filter(input, func(p Person) bool { return len(p.Name) > 3 })
			expected := []Person{
				{"Charlie", 35},
			}
			assert.Equal(t, expected, result)
		})
	})

	t.Run("BooleanSlice", func(t *testing.T) {
		t.Run("filter true values", func(t *testing.T) {
			input := []bool{true, false, true, false, true}
			result := Filter(input, func(b bool) bool { return b })
			expected := []bool{true, true, true}
			assert.Equal(t, expected, result)
		})

		t.Run("filter false values", func(t *testing.T) {
			input := []bool{true, false, true, false, true}
			result := Filter(input, func(b bool) bool { return !b })
			expected := []bool{false, false}
			assert.Equal(t, expected, result)
		})
	})

	t.Run("ComplexConditions", func(t *testing.T) {
		t.Run("multiple conditions", func(t *testing.T) {
			input := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12}
			result := Filter(input, func(n int) bool {
				return n%2 == 0 && n%3 == 0
			})
			expected := []int{6, 12}
			assert.Equal(t, expected, result)
		})

		t.Run("complex logic", func(t *testing.T) {
			input := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
			result := Filter(input, func(n int) bool {
				// Divisible by 2 or 3, but not both
				return (n%2 == 0) != (n%3 == 0)
			})
			expected := []int{2, 3, 4, 8, 9, 10}
			assert.Equal(t, expected, result)
		})
	})

	t.Run("NilSlice", func(t *testing.T) {
		var input []int
		result := Filter(input, func(n int) bool { return n > 0 })
		assert.Empty(t, result)
		assert.NotNil(t, result)
	})

	t.Run("PreservesOrder", func(t *testing.T) {
		input := []int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1}
		result := Filter(input, func(n int) bool { return n%2 == 0 })
		expected := []int{10, 8, 6, 4, 2}
		assert.Equal(t, expected, result, "Filter should preserve element order")
	})

	t.Run("DoesNotModifyOriginal", func(t *testing.T) {
		input := []int{1, 2, 3, 4, 5}
		original := make([]int, len(input))
		copy(original, input)

		_ = Filter(input, func(n int) bool { return n%2 == 0 })

		assert.Equal(t, original, input, "Filter should not modify original slice")
	})
}

// BenchmarkFilter benchmarks the Filter function with various slice sizes
// 对 Filter 函数进行不同切片大小的性能基准测试
func BenchmarkFilter(b *testing.B) {
	b.Run("SmallSlice_10", func(b *testing.B) {
		input := make([]int, 10)
		for i := range input {
			input[i] = i
		}
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			_ = Filter(input, func(n int) bool { return n%2 == 0 })
		}
	})

	b.Run("MediumSlice_100", func(b *testing.B) {
		input := make([]int, 100)
		for i := range input {
			input[i] = i
		}
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			_ = Filter(input, func(n int) bool { return n%2 == 0 })
		}
	})

	b.Run("LargeSlice_1000", func(b *testing.B) {
		input := make([]int, 1000)
		for i := range input {
			input[i] = i
		}
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			_ = Filter(input, func(n int) bool { return n%2 == 0 })
		}
	})

	b.Run("VeryLargeSlice_10000", func(b *testing.B) {
		input := make([]int, 10000)
		for i := range input {
			input[i] = i
		}
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			_ = Filter(input, func(n int) bool { return n%2 == 0 })
		}
	})

	b.Run("AllMatch", func(b *testing.B) {
		input := make([]int, 1000)
		for i := range input {
			input[i] = i * 2 // All even
		}
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			_ = Filter(input, func(n int) bool { return n%2 == 0 })
		}
	})

	b.Run("NoneMatch", func(b *testing.B) {
		input := make([]int, 1000)
		for i := range input {
			input[i] = i*2 + 1 // All odd
		}
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			_ = Filter(input, func(n int) bool { return n%2 == 0 })
		}
	})

	b.Run("StringSlice", func(b *testing.B) {
		input := make([]string, 1000)
		for i := range input {
			if i%2 == 0 {
				input[i] = "short"
			} else {
				input[i] = "verylongstring"
			}
		}
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			_ = Filter(input, func(s string) bool { return len(s) > 5 })
		}
	})
}
