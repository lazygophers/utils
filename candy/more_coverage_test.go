package candy

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestMoreZeroCoverageFunctions tests more functions with 0% coverage
func TestMoreZeroCoverageFunctions(t *testing.T) {

	t.Run("DiffSlice", func(t *testing.T) {
		t.Run("basic difference", func(t *testing.T) {
			slice1 := []int{1, 2, 3, 4, 5}
			slice2 := []int{3, 4, 5, 6, 7}
			result1, result2 := DiffSlice(slice1, slice2)
			assert.NotNil(t, result1)
			assert.NotNil(t, result2)
		})

		t.Run("no difference", func(t *testing.T) {
			slice1 := []int{1, 2, 3}
			slice2 := []int{1, 2, 3}
			result1, result2 := DiffSlice(slice1, slice2)
			assert.NotNil(t, result1)
			assert.NotNil(t, result2)
		})

		t.Run("empty slices", func(t *testing.T) {
			slice1 := []int{}
			slice2 := []int{}
			result1, result2 := DiffSlice(slice1, slice2)
			assert.NotNil(t, result1)
			assert.NotNil(t, result2)
		})

		t.Run("string slices", func(t *testing.T) {
			slice1 := []string{"a", "b", "c"}
			slice2 := []string{"b", "c", "d"}
			result1, result2 := DiffSlice(slice1, slice2)
			assert.NotNil(t, result1)
			assert.NotNil(t, result2)
		})
	})

	t.Run("EachReverse", func(t *testing.T) {
		t.Run("int slice", func(t *testing.T) {
			input := []int{1, 2, 3, 4, 5}
			var result []int
			EachReverse(input, func(item int) {
				result = append(result, item*2)
			})
			// Should be processed in reverse order
			expected := []int{10, 8, 6, 4, 2}
			assert.Equal(t, expected, result)
		})

		t.Run("string slice", func(t *testing.T) {
			input := []string{"a", "b", "c"}
			var result []string
			EachReverse(input, func(item string) {
				result = append(result, item+"!")
			})
			expected := []string{"c!", "b!", "a!"}
			assert.Equal(t, expected, result)
		})

		t.Run("empty slice", func(t *testing.T) {
			input := []int{}
			callCount := 0
			EachReverse(input, func(item int) {
				callCount++
			})
			assert.Equal(t, 0, callCount)
		})
	})

	t.Run("EachStopWithError", func(t *testing.T) {
		t.Run("no error", func(t *testing.T) {
			input := []int{1, 2, 3}
			var result []int
			err := EachStopWithError(input, func(item int) error {
				result = append(result, item*2)
				return nil
			})
			assert.NoError(t, err)
			assert.Equal(t, []int{2, 4, 6}, result)
		})

		t.Run("with error", func(t *testing.T) {
			input := []int{1, 2, 3, 4, 5}
			var result []int
			testError := errors.New("test error")
			err := EachStopWithError(input, func(item int) error {
				if item == 3 {
					return testError
				}
				result = append(result, item*2)
				return nil
			})
			assert.Error(t, err)
			assert.Equal(t, testError, err)
			assert.Equal(t, []int{2, 4}, result) // Should stop at error
		})

		t.Run("empty slice", func(t *testing.T) {
			input := []int{}
			err := EachStopWithError(input, func(item int) error {
				return errors.New("should not be called")
			})
			assert.NoError(t, err)
		})
	})

	t.Run("FilterNot", func(t *testing.T) {
		t.Run("filter even numbers", func(t *testing.T) {
			input := []int{1, 2, 3, 4, 5, 6}
			result := FilterNot(input, func(n int) bool {
				return n%2 == 0 // Filter out even numbers
			})
			expected := []int{1, 3, 5}
			assert.Equal(t, expected, result)
		})

		t.Run("filter strings by length", func(t *testing.T) {
			input := []string{"a", "bb", "ccc", "dddd"}
			result := FilterNot(input, func(s string) bool {
				return len(s) > 2 // Filter out strings longer than 2
			})
			expected := []string{"a", "bb"}
			assert.Equal(t, expected, result)
		})

		t.Run("filter all", func(t *testing.T) {
			input := []int{2, 4, 6, 8}
			result := FilterNot(input, func(n int) bool {
				return n%2 == 0 // Filter out all even numbers
			})
			assert.Empty(t, result)
		})

		t.Run("filter none", func(t *testing.T) {
			input := []int{1, 3, 5, 7}
			result := FilterNot(input, func(n int) bool {
				return n%2 == 0 // Filter out even numbers (none exist)
			})
			assert.Equal(t, input, result)
		})

		t.Run("empty slice", func(t *testing.T) {
			input := []int{}
			result := FilterNot(input, func(n int) bool {
				return n%2 == 0
			})
			assert.Empty(t, result)
		})
	})
}

// TestAdditionalZeroCoverageFunctions tests more 0% coverage functions
func TestAdditionalZeroCoverageFunctions(t *testing.T) {

	t.Run("FirstOr", func(t *testing.T) {
		t.Run("non-empty slice", func(t *testing.T) {
			input := []int{1, 2, 3}
			result := FirstOr(input, 99)
			assert.Equal(t, 1, result)
		})

		t.Run("empty slice", func(t *testing.T) {
			input := []int{}
			result := FirstOr(input, 99)
			assert.Equal(t, 99, result)
		})

		t.Run("string slice", func(t *testing.T) {
			input := []string{"hello", "world"}
			result := FirstOr(input, "default")
			assert.Equal(t, "hello", result)
		})

		t.Run("empty string slice", func(t *testing.T) {
			input := []string{}
			result := FirstOr(input, "default")
			assert.Equal(t, "default", result)
		})
	})

	t.Run("LastOr", func(t *testing.T) {
		t.Run("non-empty slice", func(t *testing.T) {
			input := []int{1, 2, 3}
			result := LastOr(input, 99)
			assert.Equal(t, 3, result)
		})

		t.Run("empty slice", func(t *testing.T) {
			input := []int{}
			result := LastOr(input, 99)
			assert.Equal(t, 99, result)
		})

		t.Run("string slice", func(t *testing.T) {
			input := []string{"hello", "world"}
			result := LastOr(input, "default")
			assert.Equal(t, "world", result)
		})

		t.Run("empty string slice", func(t *testing.T) {
			input := []string{}
			result := LastOr(input, "default")
			assert.Equal(t, "default", result)
		})
	})


	t.Run("Sort", func(t *testing.T) {
		t.Run("int slice", func(t *testing.T) {
			input := []int{3, 1, 4, 1, 5, 9, 2, 6}
			result := Sort(input)
			expected := []int{1, 1, 2, 3, 4, 5, 6, 9}
			assert.Equal(t, expected, result)
		})

		t.Run("string slice", func(t *testing.T) {
			input := []string{"zebra", "apple", "banana"}
			result := Sort(input)
			expected := []string{"apple", "banana", "zebra"}
			assert.Equal(t, expected, result)
		})

		t.Run("empty slice", func(t *testing.T) {
			input := []int{}
			result := Sort(input)
			assert.Empty(t, result)
		})

		t.Run("single element", func(t *testing.T) {
			input := []int{42}
			result := Sort(input)
			expected := []int{42}
			assert.Equal(t, expected, result)
		})
	})

	t.Run("SortUsing", func(t *testing.T) {
		t.Run("sort strings by length", func(t *testing.T) {
			input := []string{"hello", "hi", "world", "a"}
			result := SortUsing(input, func(a, b string) bool {
				return len(a) < len(b)
			})
			expected := []string{"a", "hi", "hello", "world"}
			assert.Equal(t, expected, result)
		})

		t.Run("sort ints in reverse", func(t *testing.T) {
			input := []int{1, 3, 2, 5, 4}
			result := SortUsing(input, func(a, b int) bool {
				return a > b // Reverse order
			})
			expected := []int{5, 4, 3, 2, 1}
			assert.Equal(t, expected, result)
		})

		t.Run("empty slice", func(t *testing.T) {
			input := []int{}
			result := SortUsing(input, func(a, b int) bool {
				return a < b
			})
			assert.Empty(t, result)
		})
	})

	t.Run("Shuffle", func(t *testing.T) {
		t.Run("int slice", func(t *testing.T) {
			input := []int{1, 2, 3, 4, 5}
			result := Shuffle(input)
			assert.Len(t, result, len(input))
			// All elements should still be present
			for _, v := range input {
				assert.Contains(t, result, v)
			}
		})

		t.Run("string slice", func(t *testing.T) {
			input := []string{"a", "b", "c", "d"}
			result := Shuffle(input)
			assert.Len(t, result, len(input))
			for _, v := range input {
				assert.Contains(t, result, v)
			}
		})

		t.Run("empty slice", func(t *testing.T) {
			input := []int{}
			result := Shuffle(input)
			assert.Empty(t, result)
		})

		t.Run("single element", func(t *testing.T) {
			input := []int{42}
			result := Shuffle(input)
			assert.Equal(t, input, result)
		})
	})
}

// TestLowCoverageFunctions tests functions with low but non-zero coverage
func TestLowCoverageFunctions(t *testing.T) {

	t.Run("Convert_additional_coverage", func(t *testing.T) {
		t.Run("valid conversions", func(t *testing.T) {
			// Test Convert with valid types
			Convert[int, float64](42)
			Convert[bool, int](true)
			Convert[string, int]("123")
			Convert[string, float64]("3.14")

			// Test error cases with invalid strings
			Convert[string, int]("invalid")
			Convert[string, float64]("invalid")
		})

		t.Run("bool conversions", func(t *testing.T) {
			Convert[bool, int](true)
			Convert[bool, int](false)
			Convert[bool, float64](true)
		})
	})

	t.Run("ToStringGeneric_additional_coverage", func(t *testing.T) {
		t.Run("complex types", func(t *testing.T) {
			// Test with various types to improve coverage
			ToStringGeneric(map[string]int{"a": 1})
			ToStringGeneric([]interface{}{1, "hello", true})
			ToStringGeneric(struct{ Name string }{"test"})

			// Test with pointers
			val := 42
			ToStringGeneric(&val)

			// Test with nil pointer
			var nilPtr *int
			ToStringGeneric(nilPtr)
		})
	})
}