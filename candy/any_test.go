package candy

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAny(t *testing.T) {
	t.Run("some elements satisfy condition", func(t *testing.T) {
		input := []int{1, 3, 4, 7, 9}
		result := Any(input, func(n int) bool { return n%2 == 0 })
		assert.True(t, result)
	})

	t.Run("no elements satisfy condition", func(t *testing.T) {
		input := []int{1, 3, 5, 7, 9}
		result := Any(input, func(n int) bool { return n%2 == 0 })
		assert.False(t, result)
	})

	t.Run("empty slice", func(t *testing.T) {
		input := []int{}
		result := Any(input, func(n int) bool { return n > 0 })
		assert.False(t, result)
	})

	t.Run("single element true", func(t *testing.T) {
		input := []int{2}
		result := Any(input, func(n int) bool { return n%2 == 0 })
		assert.True(t, result)
	})

	t.Run("single element false", func(t *testing.T) {
		input := []int{3}
		result := Any(input, func(n int) bool { return n%2 == 0 })
		assert.False(t, result)
	})

	t.Run("strings with non-empty", func(t *testing.T) {
		input := []string{"", "", "test"}
		result := Any(input, func(s string) bool { return len(s) > 0 })
		assert.True(t, result)
	})

	t.Run("strings all empty", func(t *testing.T) {
		input := []string{"", "", ""}
		result := Any(input, func(s string) bool { return len(s) > 0 })
		assert.False(t, result)
	})
}
