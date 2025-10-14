package candy

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEachStopWithError(t *testing.T) {
	t.Run("stop on error", func(t *testing.T) {
		input := []int{1, 2, 3, 4, 5}
		count := 0
		err := EachStopWithError(input, func(v int) error {
			count++
			if v == 3 {
				return assert.AnError
			}
			return nil
		})
		assert.Error(t, err)
		assert.Equal(t, 3, count)
	})

	t.Run("no error", func(t *testing.T) {
		input := []int{1, 2, 3}
		count := 0
		err := EachStopWithError(input, func(v int) error {
			count++
			return nil
		})
		assert.NoError(t, err)
		assert.Equal(t, 3, count)
	})
}
