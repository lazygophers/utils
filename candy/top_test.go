package candy

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTop(t *testing.T) {
	t.Run("basic int slice", func(t *testing.T) {
		input := []int{1, 2, 3, 4, 5}
		result := Top(input, 3)
		expected := []int{1, 2, 3}
		assert.Equal(t, expected, result)
	})

	t.Run("n equals slice length", func(t *testing.T) {
		input := []int{1, 2, 3}
		result := Top(input, 3)
		expected := []int{1, 2, 3}
		assert.Equal(t, expected, result)
	})

	t.Run("n greater than slice length", func(t *testing.T) {
		input := []int{1, 2}
		result := Top(input, 5)
		expected := []int{1, 2}
		assert.Equal(t, expected, result)
	})

	t.Run("n is zero", func(t *testing.T) {
		input := []int{1, 2, 3}
		result := Top(input, 0)
		expected := []int{}
		assert.Equal(t, expected, result)
	})

	t.Run("empty slice", func(t *testing.T) {
		input := []int{}
		result := Top(input, 3)
		expected := []int{}
		assert.Equal(t, expected, result)
	})

	t.Run("string slice", func(t *testing.T) {
		input := []string{"a", "b", "c", "d"}
		result := Top(input, 2)
		expected := []string{"a", "b"}
		assert.Equal(t, expected, result)
	})

	t.Run("single element", func(t *testing.T) {
		input := []int{42}
		result := Top(input, 1)
		expected := []int{42}
		assert.Equal(t, expected, result)
	})

	t.Run("negative n", func(t *testing.T) {
		input := []int{1, 2, 3}
		result := Top(input, -1)
		expected := []int{}
		assert.Equal(t, expected, result)
	})

	t.Run("modifying result doesn't affect original", func(t *testing.T) {
		input := []int{1, 2, 3, 4, 5}
		result := Top(input, 3)
		result[0] = 999
		assert.Equal(t, 1, input[0]) // 原切片不受影响
		assert.Equal(t, 999, result[0]) // 结果切片被修改
	})
}