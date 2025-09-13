package candy

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSpare(t *testing.T) {
	t.Run("basic spare", func(t *testing.T) {
		ss := []int{1, 2, 3}
		against := []int{2, 3, 4, 5}
		result := Spare(ss, against)
		assert.Equal(t, []int{4, 5}, result)
	})

	t.Run("no spare elements", func(t *testing.T) {
		ss := []int{1, 2, 3, 4, 5}
		against := []int{1, 2, 3}
		result := Spare(ss, against)
		assert.Equal(t, []int{}, result)
	})

	t.Run("all spare", func(t *testing.T) {
		ss := []int{1, 2, 3}
		against := []int{4, 5, 6}
		result := Spare(ss, against)
		assert.Equal(t, []int{4, 5, 6}, result)
	})

	t.Run("empty slices", func(t *testing.T) {
		ss := []int{}
		against := []int{}
		result := Spare(ss, against)
		assert.Equal(t, []int{}, result)
	})

	t.Run("empty ss", func(t *testing.T) {
		ss := []int{}
		against := []int{1, 2, 3}
		result := Spare(ss, against)
		assert.Equal(t, []int{1, 2, 3}, result)
	})

	t.Run("empty against", func(t *testing.T) {
		ss := []int{1, 2, 3}
		against := []int{}
		result := Spare(ss, against)
		assert.Equal(t, []int{}, result)
	})

	t.Run("string slices", func(t *testing.T) {
		ss := []string{"apple", "banana"}
		against := []string{"banana", "cherry", "date"}
		result := Spare(ss, against)
		assert.Equal(t, []string{"cherry", "date"}, result)
	})

	t.Run("duplicate elements in against", func(t *testing.T) {
		ss := []int{1, 2}
		against := []int{2, 3, 3, 4}
		result := Spare(ss, against)
		assert.Equal(t, []int{3, 3, 4}, result)
	})

	t.Run("duplicate elements in ss", func(t *testing.T) {
		ss := []int{1, 1, 2}
		against := []int{2, 3, 4}
		result := Spare(ss, against)
		assert.Equal(t, []int{3, 4}, result)
	})

	t.Run("identical slices", func(t *testing.T) {
		ss := []int{1, 2, 3}
		against := []int{1, 2, 3}
		result := Spare(ss, against)
		assert.Equal(t, []int{}, result)
	})

	t.Run("single elements", func(t *testing.T) {
		ss := []int{42}
		against := []int{24}
		result := Spare(ss, against)
		assert.Equal(t, []int{24}, result)
	})
}
