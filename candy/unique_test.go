package candy

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUnique(t *testing.T) {
	t.Run("remove duplicates from integers", func(t *testing.T) {
		input := []int{1, 2, 2, 3, 3, 3, 4, 5, 5}
		result := Unique(input)
		expected := []int{1, 2, 3, 4, 5}
		assert.Equal(t, expected, result)
	})

	t.Run("remove duplicates from strings", func(t *testing.T) {
		input := []string{"a", "b", "a", "c", "b", "d"}
		result := Unique(input)
		expected := []string{"a", "b", "c", "d"}
		assert.Equal(t, expected, result)
	})

	t.Run("no duplicates", func(t *testing.T) {
		input := []int{1, 2, 3, 4}
		result := Unique(input)
		assert.Equal(t, input, result)
	})

	t.Run("empty slice", func(t *testing.T) {
		input := []int{}
		result := Unique(input)
		assert.Empty(t, result)
	})
}
