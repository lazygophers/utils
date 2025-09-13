package candy

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLast(t *testing.T) {
	t.Run("non-empty int slice", func(t *testing.T) {
		input := []int{1, 2, 3, 4, 5}
		result := Last(input)
		assert.Equal(t, 5, result)
	})

	t.Run("empty int slice", func(t *testing.T) {
		input := []int{}
		result := Last(input)
		assert.Equal(t, 0, result) // int 的零值
	})

	t.Run("single element", func(t *testing.T) {
		input := []string{"hello"}
		result := Last(input)
		assert.Equal(t, "hello", result)
	})

	t.Run("empty string slice", func(t *testing.T) {
		input := []string{}
		result := Last(input)
		assert.Equal(t, "", result) // string 的零值
	})

	t.Run("float64 slice", func(t *testing.T) {
		input := []float64{3.14, 2.71, 1.41}
		result := Last(input)
		assert.Equal(t, 1.41, result)
	})

	t.Run("bool slice", func(t *testing.T) {
		input := []bool{true, false, true}
		result := Last(input)
		assert.Equal(t, true, result)
	})

	t.Run("empty bool slice", func(t *testing.T) {
		input := []bool{}
		result := Last(input)
		assert.Equal(t, false, result) // bool 的零值
	})
}

func TestLastOr(t *testing.T) {
	t.Run("non-empty int slice", func(t *testing.T) {
		input := []int{1, 2, 3, 4, 5}
		result := LastOr(input, 99)
		assert.Equal(t, 5, result)
	})

	t.Run("empty int slice with default", func(t *testing.T) {
		input := []int{}
		result := LastOr(input, 99)
		assert.Equal(t, 99, result)
	})

	t.Run("single element", func(t *testing.T) {
		input := []string{"hello"}
		result := LastOr(input, "default")
		assert.Equal(t, "hello", result)
	})

	t.Run("empty string slice with default", func(t *testing.T) {
		input := []string{}
		result := LastOr(input, "default")
		assert.Equal(t, "default", result)
	})

	t.Run("float64 slice", func(t *testing.T) {
		input := []float64{3.14, 2.71, 1.41}
		result := LastOr(input, 1.0)
		assert.Equal(t, 1.41, result)
	})

	t.Run("empty float64 slice", func(t *testing.T) {
		input := []float64{}
		result := LastOr(input, 1.0)
		assert.Equal(t, 1.0, result)
	})

	t.Run("bool slice", func(t *testing.T) {
		input := []bool{true, false, true}
		result := LastOr(input, false)
		assert.Equal(t, true, result)
	})

	t.Run("empty bool slice with default true", func(t *testing.T) {
		input := []bool{}
		result := LastOr(input, true)
		assert.Equal(t, true, result)
	})
}
