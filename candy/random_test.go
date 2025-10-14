package candy

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRandom(t *testing.T) {
	t.Run("random from non-empty slice", func(t *testing.T) {
		input := []int{1, 2, 3, 4, 5}
		result := Random(input)
		assert.Contains(t, input, result)
	})

	t.Run("random from single element", func(t *testing.T) {
		input := []string{"only"}
		result := Random(input)
		assert.Equal(t, "only", result)
	})

	t.Run("random from empty slice", func(t *testing.T) {
		input := []int{}
		result := Random(input)
		assert.Equal(t, 0, result) // zero value
	})
}
