package candy

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRemoveIndex(t *testing.T) {
	t.Run("remove first element", func(t *testing.T) {
		input := []int{1, 2, 3, 4, 5}
		result := RemoveIndex(input, 0)
		expected := []int{2, 3, 4, 5}
		assert.Equal(t, expected, result)
	})

	t.Run("remove middle element", func(t *testing.T) {
		input := []int{1, 2, 3, 4, 5}
		result := RemoveIndex(input, 2)
		expected := []int{1, 2, 4, 5}
		assert.Equal(t, expected, result)
	})

	t.Run("remove last element", func(t *testing.T) {
		input := []int{1, 2, 3, 4, 5}
		result := RemoveIndex(input, 4)
		expected := []int{1, 2, 3, 4}
		assert.Equal(t, expected, result)
	})

	t.Run("single element", func(t *testing.T) {
		input := []int{42}
		result := RemoveIndex(input, 0)
		expected := []int{}
		assert.Equal(t, expected, result)
	})

	t.Run("empty slice", func(t *testing.T) {
		input := []int{}
		result := RemoveIndex(input, 0)
		expected := []int{}
		assert.Equal(t, expected, result)
	})

	t.Run("negative index", func(t *testing.T) {
		input := []int{1, 2, 3}
		result := RemoveIndex(input, -1)
		expected := []int{}
		assert.Equal(t, expected, result)
	})

	t.Run("index out of range", func(t *testing.T) {
		input := []int{1, 2, 3}
		result := RemoveIndex(input, 5)
		expected := []int{}
		assert.Equal(t, expected, result)
	})

	t.Run("string slice", func(t *testing.T) {
		input := []string{"a", "b", "c", "d"}
		result := RemoveIndex(input, 1)
		expected := []string{"a", "c", "d"}
		assert.Equal(t, expected, result)
	})

	t.Run("two elements remove first", func(t *testing.T) {
		input := []int{1, 2}
		result := RemoveIndex(input, 0)
		expected := []int{2}
		assert.Equal(t, expected, result)
	})

	t.Run("two elements remove second", func(t *testing.T) {
		input := []int{1, 2}
		result := RemoveIndex(input, 1)
		expected := []int{1}
		assert.Equal(t, expected, result)
	})
}