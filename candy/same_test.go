package candy

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSame(t *testing.T) {
	t.Run("basic intersection", func(t *testing.T) {
		against := []int{1, 2, 3, 4}
		ss := []int{2, 3, 5, 6}
		result := Same(against, ss)
		assert.Equal(t, []int{2, 3}, result)
	})

	t.Run("no intersection", func(t *testing.T) {
		against := []int{1, 2, 3}
		ss := []int{4, 5, 6}
		result := Same(against, ss)
		assert.Equal(t, []int{}, result)
	})

	t.Run("complete intersection", func(t *testing.T) {
		against := []int{1, 2, 3}
		ss := []int{1, 2, 3}
		result := Same(against, ss)
		assert.Equal(t, []int{1, 2, 3}, result)
	})

	t.Run("empty slices", func(t *testing.T) {
		against := []int{}
		ss := []int{}
		result := Same(against, ss)
		assert.Equal(t, []int{}, result)
	})

	t.Run("first slice empty", func(t *testing.T) {
		against := []int{}
		ss := []int{1, 2, 3}
		result := Same(against, ss)
		assert.Equal(t, []int{}, result)
	})

	t.Run("second slice empty", func(t *testing.T) {
		against := []int{1, 2, 3}
		ss := []int{}
		result := Same(against, ss)
		assert.Equal(t, []int{}, result)
	})

	t.Run("string slices", func(t *testing.T) {
		against := []string{"apple", "banana", "cherry"}
		ss := []string{"banana", "cherry", "date"}
		result := Same(against, ss)
		assert.Equal(t, []string{"banana", "cherry"}, result)
	})

	t.Run("duplicate elements in against", func(t *testing.T) {
		against := []int{1, 2, 2, 3}
		ss := []int{2, 3, 4}
		result := Same(against, ss)
		assert.Equal(t, []int{2, 2, 3}, result) // 保留duplicates
	})

	t.Run("duplicate elements in ss", func(t *testing.T) {
		against := []int{1, 2, 3}
		ss := []int{2, 2, 3, 4}
		result := Same(against, ss)
		assert.Equal(t, []int{2, 3}, result)
	})

	t.Run("single elements", func(t *testing.T) {
		against := []int{42}
		ss := []int{42}
		result := Same(against, ss)
		assert.Equal(t, []int{42}, result)
	})

	t.Run("single elements no match", func(t *testing.T) {
		against := []int{42}
		ss := []int{24}
		result := Same(against, ss)
		assert.Equal(t, []int{}, result)
	})
}