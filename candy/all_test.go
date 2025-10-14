package candy

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAll(t *testing.T) {
	t.Run("all elements satisfy condition", func(t *testing.T) {
		input := []int{2, 4, 6, 8, 10}
		result := All(input, func(n int) bool { return n%2 == 0 })
		assert.True(t, result)
	})

	t.Run("not all elements satisfy condition", func(t *testing.T) {
		input := []int{2, 3, 4, 6, 8}
		result := All(input, func(n int) bool { return n%2 == 0 })
		assert.False(t, result)
	})

	t.Run("empty slice", func(t *testing.T) {
		input := []int{}
		result := All(input, func(n int) bool { return n > 0 })
		assert.True(t, result) // empty slice returns true
	})

	t.Run("single element true", func(t *testing.T) {
		input := []int{2}
		result := All(input, func(n int) bool { return n%2 == 0 })
		assert.True(t, result)
	})

	t.Run("single element false", func(t *testing.T) {
		input := []int{3}
		result := All(input, func(n int) bool { return n%2 == 0 })
		assert.False(t, result)
	})

	t.Run("strings all non-empty", func(t *testing.T) {
		input := []string{"hello", "world", "test"}
		result := All(input, func(s string) bool { return len(s) > 0 })
		assert.True(t, result)
	})

	t.Run("strings with empty", func(t *testing.T) {
		input := []string{"hello", "", "test"}
		result := All(input, func(s string) bool { return len(s) > 0 })
		assert.False(t, result)
	})
}
