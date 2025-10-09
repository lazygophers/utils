// Package candy provides convenient utility functions
// each_test.go tests the Each function comprehensively
package candy

import (
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestEach tests the Each function with various slice types
// 测试 Each 函数处理各种切片类型
func TestEach(t *testing.T) {
	t.Run("IntegerSlice", func(t *testing.T) {
		t.Run("basic iteration", func(t *testing.T) {
			input := []int{1, 2, 3, 4, 5}
			var result []int
			Each(input, func(n int) {
				result = append(result, n*2)
			})
			expected := []int{2, 4, 6, 8, 10}
			assert.Equal(t, expected, result)
		})

		t.Run("sum elements", func(t *testing.T) {
			input := []int{1, 2, 3, 4, 5}
			sum := 0
			Each(input, func(n int) {
				sum += n
			})
			assert.Equal(t, 15, sum)
		})

		t.Run("count elements", func(t *testing.T) {
			input := []int{1, 2, 3, 4, 5}
			count := 0
			Each(input, func(n int) {
				count++
			})
			assert.Equal(t, 5, count)
		})

		t.Run("empty slice", func(t *testing.T) {
			input := []int{}
			executed := false
			Each(input, func(n int) {
				executed = true
			})
			assert.False(t, executed, "Function should not execute for empty slice")
		})

		t.Run("nil slice", func(t *testing.T) {
			var input []int
			executed := false
			Each(input, func(n int) {
				executed = true
			})
			assert.False(t, executed, "Function should not execute for nil slice")
		})

		t.Run("single element", func(t *testing.T) {
			input := []int{42}
			var result []int
			Each(input, func(n int) {
				result = append(result, n)
			})
			assert.Equal(t, []int{42}, result)
		})
	})

	t.Run("StringSlice", func(t *testing.T) {
		t.Run("concatenate strings", func(t *testing.T) {
			input := []string{"Hello", " ", "World", "!"}
			result := ""
			Each(input, func(s string) {
				result += s
			})
			assert.Equal(t, "Hello World!", result)
		})

		t.Run("count string lengths", func(t *testing.T) {
			input := []string{"a", "bb", "ccc", "dddd"}
			totalLen := 0
			Each(input, func(s string) {
				totalLen += len(s)
			})
			assert.Equal(t, 10, totalLen)
		})

		t.Run("collect uppercase", func(t *testing.T) {
			input := []string{"hello", "world", "test"}
			var result []string
			Each(input, func(s string) {
				// Simple uppercase conversion for ASCII
				upper := ""
				for _, c := range s {
					if c >= 'a' && c <= 'z' {
						upper += string(c - 32)
					} else {
						upper += string(c)
					}
				}
				result = append(result, upper)
			})
			expected := []string{"HELLO", "WORLD", "TEST"}
			assert.Equal(t, expected, result)
		})
	})

	t.Run("FloatSlice", func(t *testing.T) {
		t.Run("sum floats", func(t *testing.T) {
			input := []float64{1.1, 2.2, 3.3, 4.4, 5.5}
			sum := 0.0
			Each(input, func(f float64) {
				sum += f
			})
			assert.InDelta(t, 16.5, sum, 0.001)
		})

		t.Run("multiply by factor", func(t *testing.T) {
			input := []float64{1.0, 2.0, 3.0}
			var result []float64
			Each(input, func(f float64) {
				result = append(result, f*2.5)
			})
			expected := []float64{2.5, 5.0, 7.5}
			for i, v := range expected {
				assert.InDelta(t, v, result[i], 0.001)
			}
		})
	})

	t.Run("StructSlice", func(t *testing.T) {
		type Person struct {
			Name string
			Age  int
		}

		t.Run("collect names", func(t *testing.T) {
			input := []Person{
				{"Alice", 25},
				{"Bob", 30},
				{"Charlie", 35},
			}
			var names []string
			Each(input, func(p Person) {
				names = append(names, p.Name)
			})
			expected := []string{"Alice", "Bob", "Charlie"}
			assert.Equal(t, expected, names)
		})

		t.Run("sum ages", func(t *testing.T) {
			input := []Person{
				{"Alice", 25},
				{"Bob", 30},
				{"Charlie", 35},
			}
			totalAge := 0
			Each(input, func(p Person) {
				totalAge += p.Age
			})
			assert.Equal(t, 90, totalAge)
		})

		t.Run("filter and collect", func(t *testing.T) {
			input := []Person{
				{"Alice", 25},
				{"Bob", 17},
				{"Charlie", 30},
				{"David", 16},
			}
			var adults []Person
			Each(input, func(p Person) {
				if p.Age >= 18 {
					adults = append(adults, p)
				}
			})
			assert.Len(t, adults, 2)
			assert.Equal(t, "Alice", adults[0].Name)
			assert.Equal(t, "Charlie", adults[1].Name)
		})
	})

	t.Run("BooleanSlice", func(t *testing.T) {
		t.Run("count true values", func(t *testing.T) {
			input := []bool{true, false, true, true, false}
			trueCount := 0
			Each(input, func(b bool) {
				if b {
					trueCount++
				}
			})
			assert.Equal(t, 3, trueCount)
		})
	})

	t.Run("SideEffects", func(t *testing.T) {
		t.Run("modify external counter", func(t *testing.T) {
			input := []int{1, 2, 3, 4, 5}
			counter := 0
			Each(input, func(n int) {
				counter += n
			})
			assert.Equal(t, 15, counter)
		})

		t.Run("append to external slice", func(t *testing.T) {
			input := []int{1, 2, 3, 4, 5}
			var squares []int
			Each(input, func(n int) {
				squares = append(squares, n*n)
			})
			expected := []int{1, 4, 9, 16, 25}
			assert.Equal(t, expected, squares)
		})

		t.Run("populate map", func(t *testing.T) {
			input := []string{"a", "bb", "ccc"}
			lengths := make(map[string]int)
			Each(input, func(s string) {
				lengths[s] = len(s)
			})
			assert.Equal(t, 3, len(lengths))
			assert.Equal(t, 1, lengths["a"])
			assert.Equal(t, 2, lengths["bb"])
			assert.Equal(t, 3, lengths["ccc"])
		})
	})

	t.Run("OrderPreservation", func(t *testing.T) {
		t.Run("maintains iteration order", func(t *testing.T) {
			input := []int{5, 4, 3, 2, 1}
			var result []int
			Each(input, func(n int) {
				result = append(result, n)
			})
			assert.Equal(t, input, result, "Each should preserve element order")
		})
	})

	t.Run("DoesNotModifyOriginal", func(t *testing.T) {
		input := []int{1, 2, 3, 4, 5}
		original := make([]int, len(input))
		copy(original, input)

		Each(input, func(n int) {
			// Perform operations that don't modify the original
			_ = n * 2
		})

		assert.Equal(t, original, input, "Each should not modify original slice")
	})

	t.Run("EmptyFunction", func(t *testing.T) {
		t.Run("no-op function", func(t *testing.T) {
			input := []int{1, 2, 3}
			Each(input, func(n int) {
				// Do nothing
			})
			// Should not panic or error
		})
	})

	t.Run("LargeSlice", func(t *testing.T) {
		t.Run("handles large slice", func(t *testing.T) {
			input := make([]int, 10000)
			for i := range input {
				input[i] = i
			}
			sum := 0
			Each(input, func(n int) {
				sum += n
			})
			expectedSum := (10000 * 9999) / 2 // Sum of 0 to 9999
			assert.Equal(t, expectedSum, sum)
		})
	})
}

// TestEachConcurrency tests the Each function behavior in concurrent scenarios
// 测试 Each 函数在并发场景下的行为
func TestEachConcurrency(t *testing.T) {
	t.Run("concurrent access safe", func(t *testing.T) {
		// Note: Each is not designed to be concurrency-safe within the function
		// but concurrent calls to Each on different slices should be safe
		var wg sync.WaitGroup
		results := make([]int, 10)

		for i := 0; i < 10; i++ {
			wg.Add(1)
			go func(idx int) {
				defer wg.Done()
				input := []int{1, 2, 3}
				sum := 0
				Each(input, func(n int) {
					sum += n
				})
				results[idx] = sum
			}(i)
		}

		wg.Wait()

		for i, sum := range results {
			assert.Equal(t, 6, sum, "Goroutine %d should have sum of 6", i)
		}
	})
}

// BenchmarkEach benchmarks the Each function with various slice sizes
// 对 Each 函数进行不同切片大小的性能基准测试
func BenchmarkEach(b *testing.B) {
	b.Run("SmallSlice_10", func(b *testing.B) {
		input := make([]int, 10)
		for i := range input {
			input[i] = i
		}
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			sum := 0
			Each(input, func(n int) {
				sum += n
			})
		}
	})

	b.Run("MediumSlice_100", func(b *testing.B) {
		input := make([]int, 100)
		for i := range input {
			input[i] = i
		}
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			sum := 0
			Each(input, func(n int) {
				sum += n
			})
		}
	})

	b.Run("LargeSlice_1000", func(b *testing.B) {
		input := make([]int, 1000)
		for i := range input {
			input[i] = i
		}
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			sum := 0
			Each(input, func(n int) {
				sum += n
			})
		}
	})

	b.Run("VeryLargeSlice_10000", func(b *testing.B) {
		input := make([]int, 10000)
		for i := range input {
			input[i] = i
		}
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			sum := 0
			Each(input, func(n int) {
				sum += n
			})
		}
	})

	b.Run("StringSlice", func(b *testing.B) {
		input := make([]string, 1000)
		for i := range input {
			input[i] = "test string"
		}
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			count := 0
			Each(input, func(s string) {
				count += len(s)
			})
		}
	})

	b.Run("NoOp", func(b *testing.B) {
		input := make([]int, 1000)
		for i := range input {
			input[i] = i
		}
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			Each(input, func(n int) {
				// No operation
			})
		}
	})

	b.Run("ComplexOperation", func(b *testing.B) {
		input := make([]int, 1000)
		for i := range input {
			input[i] = i
		}
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			var results []int
			Each(input, func(n int) {
				results = append(results, n*n+n*2)
			})
		}
	})
}
