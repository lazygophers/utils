package candy

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestChunk(t *testing.T) {
	t.Run("basic chunking", func(t *testing.T) {
		input := []int{1, 2, 3, 4, 5, 6}
		result := Chunk(input, 2)
		expected := [][]int{{1, 2}, {3, 4}, {5, 6}}
		assert.Equal(t, expected, result)
	})

	t.Run("chunk size larger than slice", func(t *testing.T) {
		input := []int{1, 2, 3}
		result := Chunk(input, 5)
		expected := [][]int{{1, 2, 3}}
		assert.Equal(t, expected, result)
	})

	t.Run("chunk size 1", func(t *testing.T) {
		input := []int{1, 2, 3}
		result := Chunk(input, 1)
		expected := [][]int{{1}, {2}, {3}}
		assert.Equal(t, expected, result)
	})

	t.Run("empty slice", func(t *testing.T) {
		input := []int{}
		result := Chunk(input, 2)
		assert.Empty(t, result)
	})

	t.Run("string slice", func(t *testing.T) {
		input := []string{"a", "b", "c", "d", "e"}
		result := Chunk(input, 3)
		expected := [][]string{{"a", "b", "c"}, {"d", "e"}}
		assert.Equal(t, expected, result)
	})
}
