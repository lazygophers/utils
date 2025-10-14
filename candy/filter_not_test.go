package candy

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFilterNot(t *testing.T) {
	t.Run("filter not even", func(t *testing.T) {
		input := []int{1, 2, 3, 4, 5, 6}
		result := FilterNot(input, func(v int) bool {
			return v%2 == 0
		})
		expected := []int{1, 3, 5}
		assert.Equal(t, expected, result)
	})
}
