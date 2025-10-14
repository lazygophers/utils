package candy

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDrop(t *testing.T) {
	t.Run("drop first elements", func(t *testing.T) {
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

	t.Run("drop zero elements", func(t *testing.T) {
		input := []int{1, 2, 3}
		result := Drop(input, 0)
		assert.Equal(t, input, result)
	})

	t.Run("drop all elements", func(t *testing.T) {
		input := []string{"a", "b", "c"}
		result := Drop(input, 3)
		assert.Empty(t, result)
	})

	t.Run("drop from single element", func(t *testing.T) {
		input := []int{42}
		result := Drop(input, 1)
		assert.Empty(t, result)
	})

	t.Run("empty slice", func(t *testing.T) {
		input := []int{}
		result := Drop(input, 1)
		assert.Empty(t, result)
	})

	t.Run("drop negative count", func(t *testing.T) {
		input := []int{1, 2, 3}
		result := Drop(input, -1)
		assert.Equal(t, input, result)
	})

	t.Run("drop strings", func(t *testing.T) {
		input := []string{"apple", "banana", "cherry", "date"}
		result := Drop(input, 1)
		expected := []string{"banana", "cherry", "date"}
		assert.Equal(t, expected, result)
	})
}
