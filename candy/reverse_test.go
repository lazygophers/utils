package candy

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReverse(t *testing.T) {
	t.Run("basic int slice", func(t *testing.T) {
		input := []int{1, 2, 3, 4, 5}
		result := Reverse(input)
		expected := []int{5, 4, 3, 2, 1}
		assert.Equal(t, expected, result)
	})

	t.Run("empty slice", func(t *testing.T) {
		input := []int{}
		result := Reverse(input)
		expected := []int{}
		assert.Equal(t, expected, result)
	})

	t.Run("single element", func(t *testing.T) {
		input := []int{42}
		result := Reverse(input)
		expected := []int{42}
		assert.Equal(t, expected, result)
	})

	t.Run("two elements", func(t *testing.T) {
		input := []int{1, 2}
		result := Reverse(input)
		expected := []int{2, 1}
		assert.Equal(t, expected, result)
	})

	t.Run("string slice", func(t *testing.T) {
		input := []string{"a", "b", "c", "d"}
		result := Reverse(input)
		expected := []string{"d", "c", "b", "a"}
		assert.Equal(t, expected, result)
	})

	t.Run("float slice", func(t *testing.T) {
		input := []float64{1.1, 2.2, 3.3}
		result := Reverse(input)
		expected := []float64{3.3, 2.2, 1.1}
		assert.Equal(t, expected, result)
	})

	t.Run("original slice unchanged", func(t *testing.T) {
		input := []int{1, 2, 3}
		original := make([]int, len(input))
		copy(original, input)
		result := Reverse(input)
		assert.Equal(t, original, input)        // 原切片不变
		assert.Equal(t, []int{3, 2, 1}, result) // 结果是反转的
	})

	t.Run("bool slice", func(t *testing.T) {
		input := []bool{true, false, true, false}
		result := Reverse(input)
		expected := []bool{false, true, false, true}
		assert.Equal(t, expected, result)
	})
}
