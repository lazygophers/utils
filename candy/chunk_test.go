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

	t.Run("uneven chunking", func(t *testing.T) {
		input := []int{1, 2, 3, 4, 5}
		result := Chunk(input, 2)
		expected := [][]int{{1, 2}, {3, 4}, {5}}
		assert.Equal(t, expected, result)
	})

	t.Run("chunk size equals slice length", func(t *testing.T) {
		input := []int{1, 2, 3}
		result := Chunk(input, 3)
		expected := [][]int{{1, 2, 3}}
		assert.Equal(t, expected, result)
	})

	t.Run("chunk size greater than slice length", func(t *testing.T) {
		input := []int{1, 2}
		result := Chunk(input, 5)
		expected := [][]int{{1, 2}}
		assert.Equal(t, expected, result)
	})

	t.Run("chunk size is 1", func(t *testing.T) {
		input := []int{1, 2, 3}
		result := Chunk(input, 1)
		expected := [][]int{{1}, {2}, {3}}
		assert.Equal(t, expected, result)
	})

	t.Run("empty slice", func(t *testing.T) {
		input := []int{}
		result := Chunk(input, 2)
		expected := [][]int{}
		assert.Equal(t, expected, result)
	})

	t.Run("zero chunk size", func(t *testing.T) {
		input := []int{1, 2, 3}
		result := Chunk(input, 0)
		expected := [][]int{}
		assert.Equal(t, expected, result)
	})

	t.Run("negative chunk size", func(t *testing.T) {
		input := []int{1, 2, 3}
		result := Chunk(input, -1)
		expected := [][]int{}
		assert.Equal(t, expected, result)
	})

	t.Run("string slice", func(t *testing.T) {
		input := []string{"a", "b", "c", "d", "e"}
		result := Chunk(input, 3)
		expected := [][]string{{"a", "b", "c"}, {"d", "e"}}
		assert.Equal(t, expected, result)
	})
}