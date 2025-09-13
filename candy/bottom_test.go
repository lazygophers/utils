package candy

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBottom(t *testing.T) {
	t.Run("basic int slice", func(t *testing.T) {
		input := []int{1, 2, 3, 4, 5}
		result := Bottom(input, 3)
		expected := []int{3, 4, 5}
		assert.Equal(t, expected, result)
	})

	t.Run("n equals slice length", func(t *testing.T) {
		input := []int{1, 2, 3}
		result := Bottom(input, 3)
		expected := []int{1, 2, 3}
		assert.Equal(t, expected, result)
	})

	t.Run("n greater than slice length", func(t *testing.T) {
		input := []int{1, 2}
		result := Bottom(input, 5)
		expected := []int{1, 2}
		assert.Equal(t, expected, result)
	})

	t.Run("n is zero", func(t *testing.T) {
		input := []int{1, 2, 3}
		result := Bottom(input, 0)
		expected := []int{}
		assert.Equal(t, expected, result)
	})

	t.Run("empty slice", func(t *testing.T) {
		input := []int{}
		result := Bottom(input, 3)
		expected := []int{}
		assert.Equal(t, expected, result)
	})

	t.Run("string slice", func(t *testing.T) {
		input := []string{"a", "b", "c", "d"}
		result := Bottom(input, 2)
		expected := []string{"c", "d"}
		assert.Equal(t, expected, result)
	})

	t.Run("single element", func(t *testing.T) {
		input := []int{42}
		result := Bottom(input, 1)
		expected := []int{42}
		assert.Equal(t, expected, result)
	})

	t.Run("negative n", func(t *testing.T) {
		input := []int{1, 2, 3}
		result := Bottom(input, -1)
		expected := []int{}
		assert.Equal(t, expected, result)
	})
}
