package candy

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestToFloat64Slice(t *testing.T) {
	t.Run("nil input", func(t *testing.T) {
		result := ToFloat64Slice(nil)
		assert.Nil(t, result)
	})

	t.Run("bool slice", func(t *testing.T) {
		input := []bool{true, false, true}
		result := ToFloat64Slice(input)
		expected := []float64{1.0, 0.0, 1.0}
		assert.Equal(t, expected, result)
	})

	t.Run("int slice", func(t *testing.T) {
		input := []int{1, 2, 3}
		result := ToFloat64Slice(input)
		expected := []float64{1.0, 2.0, 3.0}
		assert.Equal(t, expected, result)
	})

	t.Run("string slice", func(t *testing.T) {
		input := []string{"1.5", "2.7", "invalid"}
		result := ToFloat64Slice(input)
		expected := []float64{1.5, 2.7, 0.0}
		assert.Equal(t, expected, result)
	})

	t.Run("float64 slice", func(t *testing.T) {
		input := []float64{1.1, 2.2, 3.3}
		result := ToFloat64Slice(input)
		expected := []float64{1.1, 2.2, 3.3}
		assert.Equal(t, expected, result)
	})

	t.Run("interface slice", func(t *testing.T) {
		input := []interface{}{1, "2.5", true}
		result := ToFloat64Slice(input)
		expected := []float64{1.0, 2.5, 1.0}
		assert.Equal(t, expected, result)
	})

	t.Run("unsupported type", func(t *testing.T) {
		input := "not a slice"
		result := ToFloat64Slice(input)
		expected := []float64{}
		assert.Equal(t, expected, result)
	})
}

func TestToInt64Slice(t *testing.T) {
	t.Run("bool slice", func(t *testing.T) {
		input := []bool{true, false, true}
		result := ToInt64Slice(input)
		expected := []int64{1, 0, 1}
		assert.Equal(t, expected, result)
	})

	t.Run("int slice", func(t *testing.T) {
		input := []int{1, 2, 3}
		result := ToInt64Slice(input)
		expected := []int64{1, 2, 3}
		assert.Equal(t, expected, result)
	})

	t.Run("float32 slice", func(t *testing.T) {
		input := []float32{1.5, 2.7, 3.9}
		result := ToInt64Slice(input)
		expected := []int64{1, 2, 3} // truncated
		assert.Equal(t, expected, result)
	})

	t.Run("string slice", func(t *testing.T) {
		input := []string{"42", "-10", "invalid"}
		result := ToInt64Slice(input)
		expected := []int64{42, -10, 0}
		assert.Equal(t, expected, result)
	})

	t.Run("int64 slice", func(t *testing.T) {
		input := []int64{100, 200, 300}
		result := ToInt64Slice(input)
		expected := []int64{100, 200, 300}
		assert.Equal(t, expected, result)
	})

	t.Run("interface slice", func(t *testing.T) {
		input := []interface{}{1, "42", true}
		result := ToInt64Slice(input)
		expected := []int64{1, 42, 1}
		assert.Equal(t, expected, result)
	})

	t.Run("unsupported type", func(t *testing.T) {
		input := "not a slice"
		result := ToInt64Slice(input)
		expected := []int64{}
		assert.Equal(t, expected, result)
	})
}