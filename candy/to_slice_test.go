package candy

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestToFloat64Slice(t *testing.T) {
	t.Run("nil_input", func(t *testing.T) {
		result := ToFloat64Slice(nil)
		assert.Nil(t, result)
	})

	t.Run("bool_slice", func(t *testing.T) {
		input := []bool{true, false, true}
		result := ToFloat64Slice(input)
		expected := []float64{1.0, 0.0, 1.0}
		assert.Equal(t, expected, result)
	})

	t.Run("int_slice", func(t *testing.T) {
		input := []int{1, 2, 3}
		result := ToFloat64Slice(input)
		expected := []float64{1.0, 2.0, 3.0}
		assert.Equal(t, expected, result)
	})

	t.Run("int8_slice", func(t *testing.T) {
		input := []int8{1, 2, 3}
		result := ToFloat64Slice(input)
		expected := []float64{1.0, 2.0, 3.0}
		assert.Equal(t, expected, result)
	})

	t.Run("int16_slice", func(t *testing.T) {
		input := []int16{100, 200, 300}
		result := ToFloat64Slice(input)
		expected := []float64{100.0, 200.0, 300.0}
		assert.Equal(t, expected, result)
	})

	t.Run("int32_slice", func(t *testing.T) {
		input := []int32{1000, 2000, 3000}
		result := ToFloat64Slice(input)
		expected := []float64{1000.0, 2000.0, 3000.0}
		assert.Equal(t, expected, result)
	})

	t.Run("int64_slice", func(t *testing.T) {
		input := []int64{10000, 20000, 30000}
		result := ToFloat64Slice(input)
		expected := []float64{10000.0, 20000.0, 30000.0}
		assert.Equal(t, expected, result)
	})

	t.Run("uint_slice", func(t *testing.T) {
		input := []uint{1, 2, 3}
		result := ToFloat64Slice(input)
		expected := []float64{1.0, 2.0, 3.0}
		assert.Equal(t, expected, result)
	})

	t.Run("uint8_slice", func(t *testing.T) {
		input := []uint8{100, 200, 255}
		result := ToFloat64Slice(input)
		expected := []float64{100.0, 200.0, 255.0}
		assert.Equal(t, expected, result)
	})

	t.Run("uint16_slice", func(t *testing.T) {
		input := []uint16{1000, 2000, 65535}
		result := ToFloat64Slice(input)
		expected := []float64{1000.0, 2000.0, 65535.0}
		assert.Equal(t, expected, result)
	})

	t.Run("uint32_slice", func(t *testing.T) {
		input := []uint32{100000, 200000, 300000}
		result := ToFloat64Slice(input)
		expected := []float64{100000.0, 200000.0, 300000.0}
		assert.Equal(t, expected, result)
	})

	t.Run("uint64_slice", func(t *testing.T) {
		input := []uint64{1000000, 2000000, 3000000}
		result := ToFloat64Slice(input)
		expected := []float64{1000000.0, 2000000.0, 3000000.0}
		assert.Equal(t, expected, result)
	})

	t.Run("float32_slice", func(t *testing.T) {
		input := []float32{1.1, 2.2, 3.3}
		result := ToFloat64Slice(input)
		assert.Len(t, result, 3)
		assert.InDelta(t, 1.1, result[0], 1e-6)
		assert.InDelta(t, 2.2, result[1], 1e-6)
		assert.InDelta(t, 3.3, result[2], 1e-6)
	})

	t.Run("string_slice", func(t *testing.T) {
		input := []string{"1.5", "2.7", "invalid"}
		result := ToFloat64Slice(input)
		expected := []float64{1.5, 2.7, 0.0}
		assert.Equal(t, expected, result)
	})

	t.Run("float64_slice", func(t *testing.T) {
		input := []float64{1.1, 2.2, 3.3}
		result := ToFloat64Slice(input)
		expected := []float64{1.1, 2.2, 3.3}
		assert.Equal(t, expected, result)
	})

	t.Run("byte_slice_array", func(t *testing.T) {
		input := [][]byte{{49, 46, 53}, {50, 46, 55}} // "1.5", "2.7"
		result := ToFloat64Slice(input)
		expected := []float64{1.5, 2.7}
		assert.Equal(t, expected, result)
	})

	t.Run("interface_slice", func(t *testing.T) {
		input := []interface{}{1, "2.5", true}
		result := ToFloat64Slice(input)
		expected := []float64{1.0, 2.5, 1.0}
		assert.Equal(t, expected, result)
	})

	t.Run("unsupported_type", func(t *testing.T) {
		input := "not a slice"
		result := ToFloat64Slice(input)
		expected := []float64{}
		assert.Equal(t, expected, result)
	})
}

func TestToInt64Slice(t *testing.T) {
	t.Run("bool_slice", func(t *testing.T) {
		input := []bool{true, false, true}
		result := ToInt64Slice(input)
		expected := []int64{1, 0, 1}
		assert.Equal(t, expected, result)
	})

	t.Run("int_slice", func(t *testing.T) {
		input := []int{1, 2, 3}
		result := ToInt64Slice(input)
		expected := []int64{1, 2, 3}
		assert.Equal(t, expected, result)
	})

	t.Run("int8_slice", func(t *testing.T) {
		input := []int8{1, 2, 3}
		result := ToInt64Slice(input)
		expected := []int64{1, 2, 3}
		assert.Equal(t, expected, result)
	})

	t.Run("int16_slice", func(t *testing.T) {
		input := []int16{100, 200, 300}
		result := ToInt64Slice(input)
		expected := []int64{100, 200, 300}
		assert.Equal(t, expected, result)
	})

	t.Run("int32_slice", func(t *testing.T) {
		input := []int32{1000, 2000, 3000}
		result := ToInt64Slice(input)
		expected := []int64{1000, 2000, 3000}
		assert.Equal(t, expected, result)
	})

	t.Run("int64_slice", func(t *testing.T) {
		input := []int64{100, 200, 300}
		result := ToInt64Slice(input)
		expected := []int64{100, 200, 300}
		assert.Equal(t, expected, result)
	})

	t.Run("uint_slice", func(t *testing.T) {
		input := []uint{1, 2, 3}
		result := ToInt64Slice(input)
		expected := []int64{1, 2, 3}
		assert.Equal(t, expected, result)
	})

	t.Run("uint8_slice", func(t *testing.T) {
		input := []uint8{100, 200, 255}
		result := ToInt64Slice(input)
		expected := []int64{100, 200, 255}
		assert.Equal(t, expected, result)
	})

	t.Run("uint16_slice", func(t *testing.T) {
		input := []uint16{1000, 2000, 65535}
		result := ToInt64Slice(input)
		expected := []int64{1000, 2000, 65535}
		assert.Equal(t, expected, result)
	})

	t.Run("uint32_slice", func(t *testing.T) {
		input := []uint32{100000, 200000, 300000}
		result := ToInt64Slice(input)
		expected := []int64{100000, 200000, 300000}
		assert.Equal(t, expected, result)
	})

	t.Run("uint64_slice", func(t *testing.T) {
		input := []uint64{1000000, 2000000, 3000000}
		result := ToInt64Slice(input)
		expected := []int64{1000000, 2000000, 3000000}
		assert.Equal(t, expected, result)
	})

	t.Run("float32_slice", func(t *testing.T) {
		input := []float32{1.5, 2.7, 3.9}
		result := ToInt64Slice(input)
		expected := []int64{1, 2, 3} // truncated
		assert.Equal(t, expected, result)
	})

	t.Run("float64_slice", func(t *testing.T) {
		input := []float64{1.5, 2.7, 3.9}
		result := ToInt64Slice(input)
		expected := []int64{1, 2, 3} // truncated
		assert.Equal(t, expected, result)
	})

	t.Run("string_slice", func(t *testing.T) {
		input := []string{"42", "-10", "invalid"}
		result := ToInt64Slice(input)
		expected := []int64{42, -10, 0}
		assert.Equal(t, expected, result)
	})

	t.Run("byte_slice_array", func(t *testing.T) {
		input := [][]byte{{52, 50}, {45, 49, 48}} // "42", "-10"
		result := ToInt64Slice(input)
		expected := []int64{42, -10}
		assert.Equal(t, expected, result)
	})

	t.Run("interface_slice", func(t *testing.T) {
		input := []interface{}{1, "42", true}
		result := ToInt64Slice(input)
		expected := []int64{1, 42, 1}
		assert.Equal(t, expected, result)
	})

	t.Run("unsupported_type", func(t *testing.T) {
		input := "not a slice"
		result := ToInt64Slice(input)
		expected := []int64{}
		assert.Equal(t, expected, result)
	})
}
