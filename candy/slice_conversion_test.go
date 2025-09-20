package candy

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSliceConversionCoverage(t *testing.T) {
	// Test ToFloat64Slice function coverage
	t.Run("ToFloat64Slice", func(t *testing.T) {
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

		t.Run("int8 slice", func(t *testing.T) {
			input := []int8{1, 2, 3}
			result := ToFloat64Slice(input)
			expected := []float64{1.0, 2.0, 3.0}
			assert.Equal(t, expected, result)
		})

		t.Run("int16 slice", func(t *testing.T) {
			input := []int16{10, 20, 30}
			result := ToFloat64Slice(input)
			expected := []float64{10.0, 20.0, 30.0}
			assert.Equal(t, expected, result)
		})

		t.Run("int32 slice", func(t *testing.T) {
			input := []int32{100, 200, 300}
			result := ToFloat64Slice(input)
			expected := []float64{100.0, 200.0, 300.0}
			assert.Equal(t, expected, result)
		})

		t.Run("int64 slice", func(t *testing.T) {
			input := []int64{1000, 2000, 3000}
			result := ToFloat64Slice(input)
			expected := []float64{1000.0, 2000.0, 3000.0}
			assert.Equal(t, expected, result)
		})

		t.Run("uint slice", func(t *testing.T) {
			input := []uint{1, 2, 3}
			result := ToFloat64Slice(input)
			expected := []float64{1.0, 2.0, 3.0}
			assert.Equal(t, expected, result)
		})

		t.Run("uint8 slice", func(t *testing.T) {
			input := []uint8{10, 20, 30}
			result := ToFloat64Slice(input)
			expected := []float64{10.0, 20.0, 30.0}
			assert.Equal(t, expected, result)
		})

		t.Run("uint16 slice", func(t *testing.T) {
			input := []uint16{100, 200, 300}
			result := ToFloat64Slice(input)
			expected := []float64{100.0, 200.0, 300.0}
			assert.Equal(t, expected, result)
		})

		t.Run("uint32 slice", func(t *testing.T) {
			input := []uint32{1000, 2000, 3000}
			result := ToFloat64Slice(input)
			expected := []float64{1000.0, 2000.0, 3000.0}
			assert.Equal(t, expected, result)
		})

		t.Run("uint64 slice", func(t *testing.T) {
			input := []uint64{10000, 20000, 30000}
			result := ToFloat64Slice(input)
			expected := []float64{10000.0, 20000.0, 30000.0}
			assert.Equal(t, expected, result)
		})

		t.Run("float32 slice", func(t *testing.T) {
			input := []float32{1.5, 2.5, 3.5}
			result := ToFloat64Slice(input)
			expected := []float64{1.5, 2.5, 3.5}
			assert.Equal(t, expected, result)
		})

		t.Run("float64 slice", func(t *testing.T) {
			input := []float64{1.1, 2.2, 3.3}
			result := ToFloat64Slice(input)
			expected := []float64{1.1, 2.2, 3.3}
			assert.Equal(t, expected, result)
		})

		t.Run("string slice", func(t *testing.T) {
			input := []string{"1.5", "2.5", "3.5"}
			result := ToFloat64Slice(input)
			expected := []float64{1.5, 2.5, 3.5}
			assert.Equal(t, expected, result)
		})

		t.Run("bytes slice", func(t *testing.T) {
			input := [][]byte{[]byte("1.5"), []byte("2.5"), []byte("3.5")}
			result := ToFloat64Slice(input)
			expected := []float64{1.5, 2.5, 3.5}
			assert.Equal(t, expected, result)
		})

		t.Run("interface slice", func(t *testing.T) {
			input := []interface{}{1, "2.5", true, 3.5}
			result := ToFloat64Slice(input)
			expected := []float64{1.0, 2.5, 1.0, 3.5}
			assert.Equal(t, expected, result)
		})

		t.Run("unsupported type", func(t *testing.T) {
			input := "not a slice"
			result := ToFloat64Slice(input)
			expected := []float64{}
			assert.Equal(t, expected, result)
		})
	})

	// Test ToInt64Slice function coverage
	t.Run("ToInt64Slice", func(t *testing.T) {
		t.Run("nil input", func(t *testing.T) {
			result := ToInt64Slice(nil)
			expected := []int64{}
			assert.Equal(t, expected, result)
		})

		t.Run("bool slice", func(t *testing.T) {
			input := []bool{true, false, true}
			result := ToInt64Slice(input)
			expected := []int64{1, 0, 1}
			assert.Equal(t, expected, result)
		})

		t.Run("int8 slice", func(t *testing.T) {
			input := []int8{1, 2, 3}
			result := ToInt64Slice(input)
			expected := []int64{1, 2, 3}
			assert.Equal(t, expected, result)
		})

		t.Run("int16 slice", func(t *testing.T) {
			input := []int16{10, 20, 30}
			result := ToInt64Slice(input)
			expected := []int64{10, 20, 30}
			assert.Equal(t, expected, result)
		})

		t.Run("int32 slice", func(t *testing.T) {
			input := []int32{100, 200, 300}
			result := ToInt64Slice(input)
			expected := []int64{100, 200, 300}
			assert.Equal(t, expected, result)
		})

		t.Run("int64 slice", func(t *testing.T) {
			input := []int64{1000, 2000, 3000}
			result := ToInt64Slice(input)
			expected := []int64{1000, 2000, 3000}
			assert.Equal(t, expected, result)
		})

		t.Run("uint slice", func(t *testing.T) {
			input := []uint{1, 2, 3}
			result := ToInt64Slice(input)
			expected := []int64{1, 2, 3}
			assert.Equal(t, expected, result)
		})

		t.Run("uint8 slice", func(t *testing.T) {
			input := []uint8{10, 20, 30}
			result := ToInt64Slice(input)
			expected := []int64{10, 20, 30}
			assert.Equal(t, expected, result)
		})

		t.Run("uint16 slice", func(t *testing.T) {
			input := []uint16{100, 200, 300}
			result := ToInt64Slice(input)
			expected := []int64{100, 200, 300}
			assert.Equal(t, expected, result)
		})

		t.Run("uint32 slice", func(t *testing.T) {
			input := []uint32{1000, 2000, 3000}
			result := ToInt64Slice(input)
			expected := []int64{1000, 2000, 3000}
			assert.Equal(t, expected, result)
		})

		t.Run("uint64 slice", func(t *testing.T) {
			input := []uint64{10000, 20000, 30000}
			result := ToInt64Slice(input)
			expected := []int64{10000, 20000, 30000}
			assert.Equal(t, expected, result)
		})

		t.Run("float32 slice", func(t *testing.T) {
			input := []float32{1.5, 2.5, 3.5}
			result := ToInt64Slice(input)
			expected := []int64{1, 2, 3} // float to int truncates
			assert.Equal(t, expected, result)
		})

		t.Run("float64 slice", func(t *testing.T) {
			input := []float64{1.9, 2.9, 3.9}
			result := ToInt64Slice(input)
			expected := []int64{1, 2, 3} // float to int truncates
			assert.Equal(t, expected, result)
		})

		t.Run("string slice", func(t *testing.T) {
			input := []string{"1", "2", "3"}
			result := ToInt64Slice(input)
			expected := []int64{1, 2, 3}
			assert.Equal(t, expected, result)
		})

		t.Run("bytes slice", func(t *testing.T) {
			input := [][]byte{[]byte("1"), []byte("2"), []byte("3")}
			result := ToInt64Slice(input)
			expected := []int64{1, 2, 3}
			assert.Equal(t, expected, result)
		})

		t.Run("interface slice", func(t *testing.T) {
			input := []interface{}{1, "2", true, 3.5}
			result := ToInt64Slice(input)
			expected := []int64{1, 2, 1, 3}
			assert.Equal(t, expected, result)
		})

		t.Run("unsupported type", func(t *testing.T) {
			input := "not a slice"
			result := ToInt64Slice(input)
			expected := []int64{}
			assert.Equal(t, expected, result)
		})
	})
}