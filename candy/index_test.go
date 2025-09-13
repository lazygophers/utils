package candy

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIndex(t *testing.T) {
	t.Run("found at beginning", func(t *testing.T) {
		input := []int{1, 2, 3, 4, 5}
		result := Index(input, 1)
		assert.Equal(t, 0, result)
	})

	t.Run("found in middle", func(t *testing.T) {
		input := []int{1, 2, 3, 4, 5}
		result := Index(input, 3)
		assert.Equal(t, 2, result)
	})

	t.Run("found at end", func(t *testing.T) {
		input := []int{1, 2, 3, 4, 5}
		result := Index(input, 5)
		assert.Equal(t, 4, result)
	})

	t.Run("not found", func(t *testing.T) {
		input := []int{1, 2, 3, 4, 5}
		result := Index(input, 6)
		assert.Equal(t, -1, result)
	})

	t.Run("empty slice", func(t *testing.T) {
		input := []int{}
		result := Index(input, 1)
		assert.Equal(t, -1, result)
	})

	t.Run("single element found", func(t *testing.T) {
		input := []int{42}
		result := Index(input, 42)
		assert.Equal(t, 0, result)
	})

	t.Run("single element not found", func(t *testing.T) {
		input := []int{42}
		result := Index(input, 1)
		assert.Equal(t, -1, result)
	})

	t.Run("string slice found", func(t *testing.T) {
		input := []string{"apple", "banana", "cherry"}
		result := Index(input, "banana")
		assert.Equal(t, 1, result)
	})

	t.Run("string slice not found", func(t *testing.T) {
		input := []string{"apple", "banana", "cherry"}
		result := Index(input, "grape")
		assert.Equal(t, -1, result)
	})

	t.Run("duplicate elements - returns first", func(t *testing.T) {
		input := []int{1, 2, 3, 2, 5}
		result := Index(input, 2)
		assert.Equal(t, 1, result) // 返回第一个匹配的索引
	})

	t.Run("float64 slice", func(t *testing.T) {
		input := []float64{1.1, 2.2, 3.3}
		result := Index(input, 2.2)
		assert.Equal(t, 1, result)
	})
}
