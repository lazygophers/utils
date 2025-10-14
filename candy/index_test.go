package candy

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIndex(t *testing.T) {
	t.Run("find index of integer", func(t *testing.T) {
		input := []int{10, 20, 30, 40, 50}
		assert.Equal(t, 0, Index(input, 10))
		assert.Equal(t, 2, Index(input, 30))
		assert.Equal(t, 4, Index(input, 50))
		assert.Equal(t, -1, Index(input, 99))
	})

	t.Run("find index of string", func(t *testing.T) {
		input := []string{"apple", "banana", "cherry"}
		assert.Equal(t, 1, Index(input, "banana"))
		assert.Equal(t, -1, Index(input, "orange"))
	})

	t.Run("empty slice", func(t *testing.T) {
		input := []int{}
		assert.Equal(t, -1, Index(input, 1))
	})
}
