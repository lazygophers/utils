package candy

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFirst(t *testing.T) {
	t.Run("non-empty int slice", func(t *testing.T) {
		input := []int{1, 2, 3}
		result := First(input)
		assert.Equal(t, 1, result)
	})

	t.Run("empty int slice", func(t *testing.T) {
		input := []int{}
		result := First(input)
		assert.Equal(t, 0, result) // int 的零值
	})

	t.Run("single element", func(t *testing.T) {
		input := []string{"hello"}
		result := First(input)
		assert.Equal(t, "hello", result)
	})

	t.Run("empty string slice", func(t *testing.T) {
		input := []string{}
		result := First(input)
		assert.Equal(t, "", result) // string 的零值
	})

	t.Run("float64 slice", func(t *testing.T) {
		input := []float64{3.14, 2.71}
		result := First(input)
		assert.Equal(t, 3.14, result)
	})

	t.Run("bool slice", func(t *testing.T) {
		input := []bool{true, false}
		result := First(input)
		assert.Equal(t, true, result)
	})

	t.Run("empty bool slice", func(t *testing.T) {
		input := []bool{}
		result := First(input)
		assert.Equal(t, false, result) // bool 的零值
	})
}

func TestFirstOr(t *testing.T) {
	t.Run("non-empty int slice", func(t *testing.T) {
		input := []int{1, 2, 3}
		result := FirstOr(input, 99)
		assert.Equal(t, 1, result)
	})

	t.Run("empty int slice with default", func(t *testing.T) {
		input := []int{}
		result := FirstOr(input, 99)
		assert.Equal(t, 99, result)
	})

	t.Run("single element", func(t *testing.T) {
		input := []string{"hello"}
		result := FirstOr(input, "default")
		assert.Equal(t, "hello", result)
	})

	t.Run("empty string slice with default", func(t *testing.T) {
		input := []string{}
		result := FirstOr(input, "default")
		assert.Equal(t, "default", result)
	})

	t.Run("float64 slice", func(t *testing.T) {
		input := []float64{3.14, 2.71}
		result := FirstOr(input, 1.0)
		assert.Equal(t, 3.14, result)
	})

	t.Run("empty float64 slice", func(t *testing.T) {
		input := []float64{}
		result := FirstOr(input, 1.0)
		assert.Equal(t, 1.0, result)
	})

	t.Run("bool slice", func(t *testing.T) {
		input := []bool{true, false}
		result := FirstOr(input, false)
		assert.Equal(t, true, result)
	})

	t.Run("empty bool slice with default true", func(t *testing.T) {
		input := []bool{}
		result := FirstOr(input, true)
		assert.Equal(t, true, result)
	})
}