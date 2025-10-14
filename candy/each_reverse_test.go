package candy

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEachReverse(t *testing.T) {
	t.Run("basic reverse iteration", func(t *testing.T) {
		input := []int{1, 2, 3, 4, 5}
		var result []int
		EachReverse(input, func(v int) {
			result = append(result, v)
		})
		expected := []int{5, 4, 3, 2, 1}
		assert.Equal(t, expected, result)
	})
}
