package candy

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMap(t *testing.T) {
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
}
