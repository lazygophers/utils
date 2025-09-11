package candy

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRemoveSlice(t *testing.T) {
	t.Run("basic removal", func(t *testing.T) {
		src := []int{1, 2, 3, 4, 5}
		rm := []int{2, 4}
		result := RemoveSlice(src, rm)
		expected := []int{1, 3, 5}
		assert.Equal(t, expected, result)
	})

	t.Run("remove all elements", func(t *testing.T) {
		src := []int{1, 2, 3}
		rm := []int{1, 2, 3}
		result := RemoveSlice(src, rm)
		expected := []int{}
		assert.Equal(t, expected, result)
	})

	t.Run("remove none", func(t *testing.T) {
		src := []int{1, 2, 3}
		rm := []int{4, 5, 6}
		result := RemoveSlice(src, rm)
		expected := []int{1, 2, 3}
		assert.Equal(t, expected, result)
	})

	t.Run("empty source", func(t *testing.T) {
		src := []int{}
		rm := []int{1, 2}
		result := RemoveSlice(src, rm)
		expected := []int{}
		assert.Equal(t, expected, result)
	})

	t.Run("empty removal", func(t *testing.T) {
		src := []int{1, 2, 3}
		rm := []int{}
		result := RemoveSlice(src, rm)
		expected := []int{1, 2, 3}
		assert.Equal(t, expected, result)
	})

	t.Run("string slices", func(t *testing.T) {
		src := []string{"apple", "banana", "cherry", "date"}
		rm := []string{"banana", "date"}
		result := RemoveSlice(src, rm)
		expected := []string{"apple", "cherry"}
		assert.Equal(t, expected, result)
	})

	t.Run("duplicate elements in source", func(t *testing.T) {
		src := []int{1, 2, 2, 3, 2, 4}
		rm := []int{2}
		result := RemoveSlice(src, rm)
		expected := []int{1, 3, 4} // only removes unique elements
		assert.Equal(t, expected, result)
	})

	t.Run("duplicate elements in removal", func(t *testing.T) {
		src := []int{1, 2, 3, 4}
		rm := []int{2, 2, 3, 3}
		result := RemoveSlice(src, rm)
		expected := []int{1, 4}
		assert.Equal(t, expected, result)
	})

	t.Run("panic on non-slice source", func(t *testing.T) {
		assert.Panics(t, func() {
			RemoveSlice(42, []int{1, 2})
		})
	})

	t.Run("panic on non-slice removal", func(t *testing.T) {
		assert.Panics(t, func() {
			RemoveSlice([]int{1, 2, 3}, 42)
		})
	})

	t.Run("panic on different types", func(t *testing.T) {
		assert.Panics(t, func() {
			RemoveSlice([]int{1, 2, 3}, []string{"a", "b"})
		})
	})
	
	// Additional test cases to improve coverage
	t.Run("mixed integers", func(t *testing.T) {
		src := []int{-1, 0, 1, 2, -2}
		rm := []int{-1, 1}
		result := RemoveSlice(src, rm)
		expected := []int{0, 2, -2}
		assert.Equal(t, expected, result)
	})
	
	t.Run("large slice", func(t *testing.T) {
		src := make([]int, 100)
		for i := 0; i < 100; i++ {
			src[i] = i
		}
		rm := []int{10, 20, 30}
		result := RemoveSlice(src, rm).([]int)
		assert.Len(t, result, 97)
		assert.NotContains(t, result, 10)
		assert.NotContains(t, result, 20)
		assert.NotContains(t, result, 30)
	})
}