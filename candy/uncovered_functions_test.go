package candy

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestUncoveredFunctions tests functions that currently have 0% coverage
func TestUncoveredFunctions(t *testing.T) {

	// Test Chunk function
	t.Run("Chunk", func(t *testing.T) {
		t.Run("basic chunking", func(t *testing.T) {
			input := []int{1, 2, 3, 4, 5, 6}
			result := Chunk(input, 2)
			expected := [][]int{{1, 2}, {3, 4}, {5, 6}}
			assert.Equal(t, expected, result)
		})

		t.Run("chunk size larger than slice", func(t *testing.T) {
			input := []int{1, 2, 3}
			result := Chunk(input, 5)
			expected := [][]int{{1, 2, 3}}
			assert.Equal(t, expected, result)
		})

		t.Run("chunk size 1", func(t *testing.T) {
			input := []int{1, 2, 3}
			result := Chunk(input, 1)
			expected := [][]int{{1}, {2}, {3}}
			assert.Equal(t, expected, result)
		})

		t.Run("empty slice", func(t *testing.T) {
			input := []int{}
			result := Chunk(input, 2)
			assert.Empty(t, result)
		})

		t.Run("string slice", func(t *testing.T) {
			input := []string{"a", "b", "c", "d", "e"}
			result := Chunk(input, 3)
			expected := [][]string{{"a", "b", "c"}, {"d", "e"}}
			assert.Equal(t, expected, result)
		})
	})

	// Test Each function
	t.Run("Each", func(t *testing.T) {
		t.Run("basic iteration", func(t *testing.T) {
			input := []int{1, 2, 3}
			var result []int
			Each(input, func(item int) {
				result = append(result, item*2)
			})
			expected := []int{2, 4, 6}
			assert.Equal(t, expected, result)
		})

		t.Run("empty slice", func(t *testing.T) {
			input := []int{}
			callCount := 0
			Each(input, func(item int) {
				callCount++
			})
			assert.Equal(t, 0, callCount)
		})

		t.Run("string slice", func(t *testing.T) {
			input := []string{"hello", "world"}
			var result []string
			Each(input, func(item string) {
				result = append(result, item+"!")
			})
			expected := []string{"hello!", "world!"}
			assert.Equal(t, expected, result)
		})
	})

	// Test Map function
	t.Run("Map", func(t *testing.T) {
		t.Run("double integers", func(t *testing.T) {
			input := []int{1, 2, 3, 4}
			result := Map(input, func(item int) int {
				return item * 2
			})
			expected := []int{2, 4, 6, 8}
			assert.Equal(t, expected, result)
		})

		t.Run("string to length", func(t *testing.T) {
			input := []string{"hello", "world", "test"}
			result := Map(input, func(item string) int {
				return len(item)
			})
			expected := []int{5, 5, 4}
			assert.Equal(t, expected, result)
		})

		t.Run("empty slice", func(t *testing.T) {
			input := []int{}
			result := Map(input, func(item int) int {
				return item * 2
			})
			assert.Empty(t, result)
		})
	})

	// Test Reverse function
	t.Run("Reverse", func(t *testing.T) {
		t.Run("basic reverse", func(t *testing.T) {
			input := []int{1, 2, 3, 4, 5}
			result := Reverse(input)
			expected := []int{5, 4, 3, 2, 1}
			assert.Equal(t, expected, result)
		})

		t.Run("single element", func(t *testing.T) {
			input := []int{1}
			result := Reverse(input)
			expected := []int{1}
			assert.Equal(t, expected, result)
		})

		t.Run("empty slice", func(t *testing.T) {
			input := []int{}
			result := Reverse(input)
			assert.Empty(t, result)
		})

		t.Run("string slice", func(t *testing.T) {
			input := []string{"a", "b", "c"}
			result := Reverse(input)
			expected := []string{"c", "b", "a"}
			assert.Equal(t, expected, result)
		})
	})

	// Test Unique function
	t.Run("Unique", func(t *testing.T) {
		t.Run("basic deduplication", func(t *testing.T) {
			input := []int{1, 2, 2, 3, 3, 4}
			result := Unique(input)
			expected := []int{1, 2, 3, 4}
			assert.Equal(t, expected, result)
		})

		t.Run("no duplicates", func(t *testing.T) {
			input := []int{1, 2, 3, 4}
			result := Unique(input)
			expected := []int{1, 2, 3, 4}
			assert.Equal(t, expected, result)
		})

		t.Run("all duplicates", func(t *testing.T) {
			input := []int{1, 1, 1, 1}
			result := Unique(input)
			expected := []int{1}
			assert.Equal(t, expected, result)
		})

		t.Run("empty slice", func(t *testing.T) {
			input := []int{}
			result := Unique(input)
			assert.Empty(t, result)
		})

		t.Run("string slice", func(t *testing.T) {
			input := []string{"a", "b", "b", "c", "a"}
			result := Unique(input)
			expected := []string{"a", "b", "c"}
			assert.Equal(t, expected, result)
		})
	})

	// Test Join function
	t.Run("Join", func(t *testing.T) {
		t.Run("basic join", func(t *testing.T) {
			input := []string{"hello", "world", "test"}
			result := Join(input, ", ")
			expected := "hello, world, test"
			assert.Equal(t, expected, result)
		})

		t.Run("empty separator", func(t *testing.T) {
			input := []string{"a", "b", "c"}
			result := Join(input, "")
			expected := "abc"
			assert.Equal(t, expected, result)
		})

		t.Run("single element", func(t *testing.T) {
			input := []string{"hello"}
			result := Join(input, ", ")
			expected := "hello"
			assert.Equal(t, expected, result)
		})

		t.Run("empty slice", func(t *testing.T) {
			input := []string{}
			result := Join(input, ", ")
			expected := ""
			assert.Equal(t, expected, result)
		})

		t.Run("integer slice", func(t *testing.T) {
			input := []int{1, 2, 3}
			result := Join(input, "-")
			expected := "1-2-3"
			assert.Equal(t, expected, result)
		})
	})

	// Test Reduce function
	t.Run("Reduce", func(t *testing.T) {
		t.Run("sum integers", func(t *testing.T) {
			input := []int{1, 2, 3, 4, 5}
			result := Reduce(input, func(acc, item int) int {
				return acc + item
			})
			expected := 15
			assert.Equal(t, expected, result)
		})

		t.Run("concatenate strings", func(t *testing.T) {
			input := []string{"hello", "world", "test"}
			result := Reduce(input, func(acc, item string) string {
				return acc + item
			})
			expected := "helloworldtest"
			assert.Equal(t, expected, result)
		})

		t.Run("find max", func(t *testing.T) {
			input := []int{3, 1, 4, 1, 5, 9, 2, 6}
			result := Reduce(input, func(acc, item int) int {
				if item > acc {
					return item
				}
				return acc
			})
			expected := 9
			assert.Equal(t, expected, result)
		})

		t.Run("empty slice", func(t *testing.T) {
			input := []int{}
			result := Reduce(input, func(acc, item int) int {
				return acc + item
			})
			expected := 0 // should return zero value
			assert.Equal(t, expected, result)
		})
	})

	// Test Convert function
	t.Run("Convert", func(t *testing.T) {
		t.Run("int to float64", func(t *testing.T) {
			result := Convert[int, float64](42)
			expected := 42.0
			assert.Equal(t, expected, result)
		})

		t.Run("string to int", func(t *testing.T) {
			result := Convert[string, int]("123")
			expected := 123
			assert.Equal(t, expected, result)
		})

		t.Run("bool to int", func(t *testing.T) {
			result := Convert[bool, int](true)
			expected := 1
			assert.Equal(t, expected, result)
		})

		t.Run("invalid conversion", func(t *testing.T) {
			result := Convert[string, int]("invalid")
			expected := 0 // should return zero value
			assert.Equal(t, expected, result)
		})
	})

	// Test Drop function
	t.Run("Drop", func(t *testing.T) {
		t.Run("drop from beginning", func(t *testing.T) {
			input := []int{1, 2, 3, 4, 5}
			result := Drop(input, 2)
			expected := []int{3, 4, 5}
			assert.Equal(t, expected, result)
		})

		t.Run("drop more than length", func(t *testing.T) {
			input := []int{1, 2, 3}
			result := Drop(input, 5)
			assert.Empty(t, result)
		})

		t.Run("drop zero", func(t *testing.T) {
			input := []int{1, 2, 3}
			result := Drop(input, 0)
			expected := []int{1, 2, 3}
			assert.Equal(t, expected, result)
		})

		t.Run("empty slice", func(t *testing.T) {
			input := []int{}
			result := Drop(input, 2)
			assert.Empty(t, result)
		})
	})

	// Test Same function
	t.Run("Same", func(t *testing.T) {
		t.Run("intersection of two slices", func(t *testing.T) {
			slice1 := []int{1, 2, 3, 4}
			slice2 := []int{3, 4, 5, 6}
			result := Same(slice1, slice2)
			expected := []int{3, 4}
			assert.Equal(t, expected, result)
		})

		t.Run("no intersection", func(t *testing.T) {
			slice1 := []int{1, 2, 3}
			slice2 := []int{4, 5, 6}
			result := Same(slice1, slice2)
			assert.Empty(t, result)
		})

		t.Run("identical slices", func(t *testing.T) {
			slice1 := []int{1, 2, 3}
			slice2 := []int{1, 2, 3}
			result := Same(slice1, slice2)
			expected := []int{1, 2, 3}
			assert.Equal(t, expected, result)
		})

		t.Run("empty slices", func(t *testing.T) {
			slice1 := []int{}
			slice2 := []int{}
			result := Same(slice1, slice2)
			assert.Empty(t, result)
		})

		t.Run("one empty slice", func(t *testing.T) {
			slice1 := []int{1, 2, 3}
			slice2 := []int{}
			result := Same(slice1, slice2)
			assert.Empty(t, result)
		})

		t.Run("string slices", func(t *testing.T) {
			slice1 := []string{"a", "b", "c"}
			slice2 := []string{"b", "c", "d"}
			result := Same(slice1, slice2)
			expected := []string{"b", "c"}
			assert.Equal(t, expected, result)
		})
	})

	// Test SliceEqual function
	t.Run("SliceEqual", func(t *testing.T) {
		t.Run("equal slices", func(t *testing.T) {
			a := []int{1, 2, 3}
			b := []int{1, 2, 3}
			result := SliceEqual(a, b)
			assert.True(t, result)
		})

		t.Run("different lengths", func(t *testing.T) {
			a := []int{1, 2, 3}
			b := []int{1, 2}
			result := SliceEqual(a, b)
			assert.False(t, result)
		})

		t.Run("different elements", func(t *testing.T) {
			a := []int{1, 2, 3}
			b := []int{1, 2, 4}
			result := SliceEqual(a, b)
			assert.False(t, result)
		})

		t.Run("both empty", func(t *testing.T) {
			a := []int{}
			b := []int{}
			result := SliceEqual(a, b)
			assert.True(t, result)
		})

		t.Run("string slices", func(t *testing.T) {
			a := []string{"hello", "world"}
			b := []string{"hello", "world"}
			result := SliceEqual(a, b)
			assert.True(t, result)
		})
	})

	// Test RemoveIndex function
	t.Run("RemoveIndex", func(t *testing.T) {
		t.Run("remove middle element", func(t *testing.T) {
			input := []int{1, 2, 3, 4, 5}
			result := RemoveIndex(input, 2)
			expected := []int{1, 2, 4, 5}
			assert.Equal(t, expected, result)
		})

		t.Run("remove first element", func(t *testing.T) {
			input := []int{1, 2, 3}
			result := RemoveIndex(input, 0)
			expected := []int{2, 3}
			assert.Equal(t, expected, result)
		})

		t.Run("remove last element", func(t *testing.T) {
			input := []int{1, 2, 3}
			result := RemoveIndex(input, 2)
			expected := []int{1, 2}
			assert.Equal(t, expected, result)
		})

		t.Run("invalid index", func(t *testing.T) {
			input := []int{1, 2, 3}
			result := RemoveIndex(input, 5)
			expected := []int{} // function returns empty slice for invalid index
			assert.Equal(t, expected, result)
		})

		t.Run("negative index", func(t *testing.T) {
			input := []int{1, 2, 3}
			result := RemoveIndex(input, -1)
			expected := []int{} // function returns empty slice for negative index
			assert.Equal(t, expected, result)
		})
	})
}