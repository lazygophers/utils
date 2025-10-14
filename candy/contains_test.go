package candy

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestContains(t *testing.T) {
	t.Run("contains integer", func(t *testing.T) {
		input := []int{1, 2, 3, 4, 5}
		assert.True(t, Contains(input, 3))
		assert.False(t, Contains(input, 10))
	})

	t.Run("contains string", func(t *testing.T) {
		input := []string{"apple", "banana", "cherry"}
		assert.True(t, Contains(input, "banana"))
		assert.False(t, Contains(input, "orange"))
	})

	t.Run("empty slice", func(t *testing.T) {
		input := []int{}
		assert.False(t, Contains(input, 1))
	})
}
