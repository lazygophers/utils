package candy

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReverse(t *testing.T) {
	t.Run("reverse integers", func(t *testing.T) {
		input := []int{1, 2, 3, 4, 5}
		result := Reverse(input)
		expected := []int{5, 4, 3, 2, 1}
		assert.Equal(t, expected, result)
	})

	t.Run("reverse strings", func(t *testing.T) {
		input := []string{"a", "b", "c"}
		result := Reverse(input)
		expected := []string{"c", "b", "a"}
		assert.Equal(t, expected, result)
	})

	t.Run("single element", func(t *testing.T) {
		input := []int{42}
		result := Reverse(input)
		expected := []int{42}
		assert.Equal(t, expected, result)
	})

	t.Run("empty slice", func(t *testing.T) {
		input := []int{}
		result := Reverse(input)
		assert.Empty(t, result)
	})
}
