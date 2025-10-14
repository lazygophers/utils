package candy

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// Tests for generic conversion functions
func TestToBoolGeneric(t *testing.T) {
	t.Run("int types", func(t *testing.T) {
		assert.True(t, ToBoolGeneric(1))
		assert.False(t, ToBoolGeneric(0))
		assert.True(t, ToBoolGeneric(int8(5)))
		assert.True(t, ToBoolGeneric(int16(-1)))
	})

	t.Run("uint types", func(t *testing.T) {
		assert.True(t, ToBoolGeneric(uint(1)))
		assert.False(t, ToBoolGeneric(uint(0)))
		assert.True(t, ToBoolGeneric(uint32(100)))
	})

	t.Run("float types", func(t *testing.T) {
		assert.True(t, ToBoolGeneric(1.5))
		assert.False(t, ToBoolGeneric(0.0))
		assert.True(t, ToBoolGeneric(float32(-2.5)))
	})

	t.Run("string types", func(t *testing.T) {
		assert.True(t, ToBoolGeneric("true"))
		assert.True(t, ToBoolGeneric("1"))
		assert.False(t, ToBoolGeneric("false"))
		assert.False(t, ToBoolGeneric("0"))
	})

	t.Run("bool type", func(t *testing.T) {
		assert.True(t, ToBoolGeneric(true))
		assert.False(t, ToBoolGeneric(false))
	})
}

func TestToStringGeneric(t *testing.T) {
	t.Run("basic types", func(t *testing.T) {
		assert.Equal(t, "123", ToStringGeneric(123))
		assert.Equal(t, "true", ToStringGeneric(true))
		assert.Equal(t, "3.14", ToStringGeneric(3.14))
		assert.Equal(t, "hello", ToStringGeneric("hello"))
	})

	t.Run("various int types", func(t *testing.T) {
		assert.Equal(t, "10", ToStringGeneric(int8(10)))
		assert.Equal(t, "20", ToStringGeneric(int16(20)))
		assert.Equal(t, "30", ToStringGeneric(int32(30)))
		assert.Equal(t, "40", ToStringGeneric(int64(40)))
	})

	t.Run("various uint types", func(t *testing.T) {
		assert.Equal(t, "10", ToStringGeneric(uint8(10)))
		assert.Equal(t, "20", ToStringGeneric(uint16(20)))
		assert.Equal(t, "30", ToStringGeneric(uint32(30)))
		assert.Equal(t, "40", ToStringGeneric(uint64(40)))
	})
}

func TestToSlice(t *testing.T) {
	t.Run("slice to slice with converter", func(t *testing.T) {
		input := []int{1, 2, 3, 4}
		result := ToSlice(input, func(v any) string {
			return ToString(v)
		})
		expected := []string{"1", "2", "3", "4"}
		assert.Equal(t, expected, result)
	})

	t.Run("single value to slice", func(t *testing.T) {
		result := ToSlice(42, func(v any) int {
			return ToInt(v)
		})
		expected := []int{42}
		assert.Equal(t, expected, result)
	})

	t.Run("array to slice", func(t *testing.T) {
		input := [3]string{"a", "b", "c"}
		result := ToSlice(input, func(v any) string {
			return ToString(v)
		})
		expected := []string{"a", "b", "c"}
		assert.Equal(t, expected, result)
	})

	t.Run("empty slice", func(t *testing.T) {
		input := []int{}
		result := ToSlice(input, func(v any) string {
			return ToString(v)
		})
		assert.Empty(t, result)
	})
}
