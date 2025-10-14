package candy

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDiffSlice(t *testing.T) {
	t.Run("basic diff", func(t *testing.T) {
		s1 := []int{1, 2, 3, 4, 5}
		s2 := []int{3, 4, 5, 6, 7}
		result, _ := DiffSlice(s1, s2)
		expected := []int{1, 2}
		assert.Equal(t, expected, result)
	})

	t.Run("no common elements", func(t *testing.T) {
		s1 := []int{1, 2, 3}
		s2 := []int{4, 5, 6}
		result, _ := DiffSlice(s1, s2)
		assert.Equal(t, s1, result)
	})

	t.Run("all common", func(t *testing.T) {
		s1 := []int{1, 2, 3}
		s2 := []int{1, 2, 3}
		result, _ := DiffSlice(s1, s2)
		assert.Empty(t, result)
	})
}
