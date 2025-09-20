package candy

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewZeroCoverageFunctions(t *testing.T) {
	// Test RemoveSlice function
	t.Run("RemoveSlice", func(t *testing.T) {
		t.Run("remove existing elements", func(t *testing.T) {
			source := []int{1, 2, 3, 4, 5}
			toRemove := []int{2, 4}
			result := RemoveSlice(source, toRemove)
			expected := []int{1, 3, 5}
			assert.Equal(t, expected, result)
		})

		t.Run("remove non-existing elements", func(t *testing.T) {
			source := []int{1, 2, 3}
			toRemove := []int{4, 5}
			result := RemoveSlice(source, toRemove)
			expected := []int{1, 2, 3}
			assert.Equal(t, expected, result)
		})

		t.Run("remove all elements", func(t *testing.T) {
			source := []int{1, 2, 3}
			toRemove := []int{1, 2, 3}
			result := RemoveSlice(source, toRemove)
			assert.Empty(t, result)
		})

		t.Run("empty source", func(t *testing.T) {
			source := []int{}
			toRemove := []int{1, 2}
			result := RemoveSlice(source, toRemove)
			assert.Empty(t, result)
		})

		t.Run("empty remove slice", func(t *testing.T) {
			source := []int{1, 2, 3}
			toRemove := []int{}
			result := RemoveSlice(source, toRemove)
			expected := []int{1, 2, 3}
			assert.Equal(t, expected, result)
		})

		t.Run("string slice", func(t *testing.T) {
			source := []string{"a", "b", "c", "d"}
			toRemove := []string{"b", "d"}
			result := RemoveSlice(source, toRemove)
			expected := []string{"a", "c"}
			assert.Equal(t, expected, result)
		})
	})

	// Test Slice2Map function
	t.Run("Slice2Map", func(t *testing.T) {
		t.Run("string slice to map", func(t *testing.T) {
			input := []string{"a", "b", "c"}
			result := Slice2Map(input)
			expected := map[string]bool{"a": true, "b": true, "c": true}
			assert.Equal(t, expected, result)
		})

		t.Run("int slice to map", func(t *testing.T) {
			input := []int{1, 2, 3}
			result := Slice2Map(input)
			expected := map[int]bool{1: true, 2: true, 3: true}
			assert.Equal(t, expected, result)
		})

		t.Run("empty slice", func(t *testing.T) {
			input := []string{}
			result := Slice2Map(input)
			expected := map[string]bool{}
			assert.Equal(t, expected, result)
		})

		t.Run("duplicate elements", func(t *testing.T) {
			input := []string{"a", "b", "a", "c"}
			result := Slice2Map(input)
			expected := map[string]bool{"a": true, "b": true, "c": true}
			assert.Equal(t, expected, result)
		})
	})

	// Test SliceToMapGeneric function
	t.Run("SliceToMapGeneric", func(t *testing.T) {
		t.Run("string slice", func(t *testing.T) {
			input := []string{"a", "b", "c"}
			result := SliceToMapGeneric(input)
			expected := map[string]bool{"a": true, "b": true, "c": true}
			assert.Equal(t, expected, result)
		})

		t.Run("int slice", func(t *testing.T) {
			input := []int{1, 2, 3}
			result := SliceToMapGeneric(input)
			expected := map[int]bool{1: true, 2: true, 3: true}
			assert.Equal(t, expected, result)
		})

		t.Run("nil slice", func(t *testing.T) {
			var input []string
			result := SliceToMapGeneric(input)
			assert.Nil(t, result)
		})
	})

	// Test SliceToMapWithValue function
	t.Run("SliceToMapWithValue", func(t *testing.T) {
		t.Run("string slice with int value", func(t *testing.T) {
			input := []string{"hello", "world", "test"}
			result := SliceToMapWithValue(input, 42)
			expected := map[string]int{"hello": 42, "world": 42, "test": 42}
			assert.Equal(t, expected, result)
		})

		t.Run("int slice with string value", func(t *testing.T) {
			input := []int{1, 2, 3}
			result := SliceToMapWithValue(input, "value")
			expected := map[int]string{1: "value", 2: "value", 3: "value"}
			assert.Equal(t, expected, result)
		})

		t.Run("nil slice", func(t *testing.T) {
			var input []string
			result := SliceToMapWithValue(input, 42)
			assert.Nil(t, result)
		})
	})

	// Test SliceToMapWithIndex function
	t.Run("SliceToMapWithIndex", func(t *testing.T) {
		t.Run("string slice with index", func(t *testing.T) {
			input := []string{"a", "b", "c"}
			result := SliceToMapWithIndex(input)
			expected := map[string]int{"a": 0, "b": 1, "c": 2}
			assert.Equal(t, expected, result)
		})

		t.Run("int slice with index", func(t *testing.T) {
			input := []int{10, 20, 30}
			result := SliceToMapWithIndex(input)
			expected := map[int]int{10: 0, 20: 1, 30: 2}
			assert.Equal(t, expected, result)
		})

		t.Run("nil slice", func(t *testing.T) {
			var input []string
			result := SliceToMapWithIndex(input)
			assert.Nil(t, result)
		})
	})

	// Test Spare function
	t.Run("Spare", func(t *testing.T) {
		t.Run("basic difference operation", func(t *testing.T) {
			ss := []int{1, 2, 3}
			against := []int{2, 3, 4, 5}
			result := Spare(ss, against)
			expected := []int{4, 5}
			assert.Equal(t, expected, result)
		})

		t.Run("no difference", func(t *testing.T) {
			ss := []int{1, 2, 3, 4, 5}
			against := []int{1, 2, 3}
			result := Spare(ss, against)
			assert.Empty(t, result)
		})

		t.Run("all different", func(t *testing.T) {
			ss := []int{1, 2, 3}
			against := []int{4, 5, 6}
			result := Spare(ss, against)
			expected := []int{4, 5, 6}
			assert.Equal(t, expected, result)
		})

		t.Run("empty ss", func(t *testing.T) {
			ss := []int{}
			against := []int{1, 2, 3}
			result := Spare(ss, against)
			expected := []int{1, 2, 3}
			assert.Equal(t, expected, result)
		})

		t.Run("empty against", func(t *testing.T) {
			ss := []int{1, 2, 3}
			against := []int{}
			result := Spare(ss, against)
			assert.Empty(t, result)
		})

		t.Run("string slices", func(t *testing.T) {
			ss := []string{"a", "b"}
			against := []string{"b", "c", "d"}
			result := Spare(ss, against)
			expected := []string{"c", "d"}
			assert.Equal(t, expected, result)
		})
	})

	// Test UniqueUsing function
	t.Run("UniqueUsing", func(t *testing.T) {
		t.Run("string slice with length comparator", func(t *testing.T) {
			input := []string{"a", "bb", "c", "dd", "e"}
			result := UniqueUsing(input, func(s string) any {
				return len(s)
			})
			expected := []string{"a", "bb"}
			assert.Equal(t, expected, result)
		})

		t.Run("int slice with modulo comparator", func(t *testing.T) {
			input := []int{1, 2, 3, 4, 5, 6}
			result := UniqueUsing(input, func(i int) any {
				return i % 3
			})
			expected := []int{1, 2, 3}
			assert.Equal(t, expected, result)
		})

		t.Run("empty slice", func(t *testing.T) {
			input := []string{}
			result := UniqueUsing(input, func(s string) any {
				return len(s)
			})
			assert.Empty(t, result)
		})

		t.Run("all same after transformation", func(t *testing.T) {
			input := []string{"hello", "world", "tests"}
			result := UniqueUsing(input, func(s string) any {
				return 5 // all strings have same "transformed" value
			})
			expected := []string{"hello"}
			assert.Equal(t, expected, result)
		})

		t.Run("struct slice with field comparator", func(t *testing.T) {
			type Person struct {
				Name string
				Age  int
			}
			input := []Person{
				{Name: "Alice", Age: 25},
				{Name: "Bob", Age: 25},
				{Name: "Charlie", Age: 30},
			}
			result := UniqueUsing(input, func(p Person) any {
				return p.Age
			})
			expected := []Person{
				{Name: "Alice", Age: 25},
				{Name: "Charlie", Age: 30},
			}
			assert.Equal(t, expected, result)
		})
	})
}