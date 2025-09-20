package candy

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConvertCoverage(t *testing.T) {
	// Test ConvertWithDefault function
	t.Run("ConvertWithDefault", func(t *testing.T) {
		t.Run("nil pointer", func(t *testing.T) {
			var ptr *int
			result := ConvertWithDefault(ptr, 42)
			assert.Equal(t, 42, result)
		})

		t.Run("invalid value", func(t *testing.T) {
			var val interface{}
			result := ConvertWithDefault(val, 99)
			assert.Equal(t, 99, result)
		})

		t.Run("bool true", func(t *testing.T) {
			result := ConvertWithDefault(true, 0)
			assert.Equal(t, 1, result)
		})

		t.Run("bool false", func(t *testing.T) {
			result := ConvertWithDefault(false, 99)
			assert.Equal(t, 0, result)
		})

		t.Run("int types", func(t *testing.T) {
			assert.Equal(t, int64(42), ConvertWithDefault(int(42), int64(0)))
			assert.Equal(t, int64(42), ConvertWithDefault(int8(42), int64(0)))
			assert.Equal(t, int64(42), ConvertWithDefault(int16(42), int64(0)))
			assert.Equal(t, int64(42), ConvertWithDefault(int32(42), int64(0)))
			assert.Equal(t, int64(42), ConvertWithDefault(int64(42), int64(0)))
		})

		t.Run("uint types", func(t *testing.T) {
			assert.Equal(t, uint64(42), ConvertWithDefault(uint(42), uint64(0)))
			assert.Equal(t, uint64(42), ConvertWithDefault(uint8(42), uint64(0)))
			assert.Equal(t, uint64(42), ConvertWithDefault(uint16(42), uint64(0)))
			assert.Equal(t, uint64(42), ConvertWithDefault(uint32(42), uint64(0)))
			assert.Equal(t, uint64(42), ConvertWithDefault(uint64(42), uint64(0)))
		})

		t.Run("float types", func(t *testing.T) {
			assert.Equal(t, float64(42.5), ConvertWithDefault(float32(42.5), float64(0)))
			assert.Equal(t, float64(42.5), ConvertWithDefault(float64(42.5), float64(0)))
		})

		t.Run("string to int with default", func(t *testing.T) {
			assert.Equal(t, 123, ConvertWithDefault("123", 0))
			assert.Equal(t, 99, ConvertWithDefault("invalid", 99))
		})

		t.Run("string to uint with default", func(t *testing.T) {
			assert.Equal(t, uint64(123), ConvertWithDefault("123", uint64(0)))
			assert.Equal(t, uint64(99), ConvertWithDefault("invalid", uint64(99)))
			assert.Equal(t, uint64(99), ConvertWithDefault("-123", uint64(99))) // negative to uint
		})

		t.Run("string to float with default", func(t *testing.T) {
			assert.Equal(t, 123.45, ConvertWithDefault("123.45", 0.0))
			assert.Equal(t, 99.0, ConvertWithDefault("invalid", 99.0))
		})

		t.Run("bytes with default", func(t *testing.T) {
			result := ConvertWithDefault([]byte("123"), 0)
			assert.Equal(t, 123, result)

			result = ConvertWithDefault([]byte("invalid"), 99)
			assert.Equal(t, 99, result)
		})

		t.Run("unsupported type", func(t *testing.T) {
			type CustomType struct{ value int }
			custom := CustomType{value: 42}
			result := ConvertWithDefault(custom, 99)
			assert.Equal(t, 99, result)
		})
	})

	// Test more Convert function paths
	t.Run("Convert additional paths", func(t *testing.T) {
		t.Run("bytes conversion", func(t *testing.T) {
			result := Convert[[]byte, int]([]byte("123"))
			assert.Equal(t, 123, result)

			result = Convert[[]byte, int]([]byte("invalid"))
			assert.Equal(t, 0, result)
		})

		t.Run("string to different numeric types", func(t *testing.T) {
			// Test int types
			assert.Equal(t, int(123), Convert[string, int]("123"))
			assert.Equal(t, int8(123), Convert[string, int8]("123"))
			assert.Equal(t, int16(123), Convert[string, int16]("123"))
			assert.Equal(t, int32(123), Convert[string, int32]("123"))
			assert.Equal(t, int64(123), Convert[string, int64]("123"))

			// Test uint types
			assert.Equal(t, uint(123), Convert[string, uint]("123"))
			assert.Equal(t, uint8(123), Convert[string, uint8]("123"))
			assert.Equal(t, uint16(123), Convert[string, uint16]("123"))
			assert.Equal(t, uint32(123), Convert[string, uint32]("123"))
			assert.Equal(t, uint64(123), Convert[string, uint64]("123"))

			// Test float types
			assert.Equal(t, float32(123.45), Convert[string, float32]("123.45"))
			assert.Equal(t, float64(123.45), Convert[string, float64]("123.45"))
		})

		t.Run("invalid string conversions", func(t *testing.T) {
			assert.Equal(t, 0, Convert[string, int]("invalid"))
			assert.Equal(t, uint64(0), Convert[string, uint64]("invalid"))
			assert.Equal(t, 0.0, Convert[string, float64]("invalid"))
		})

		t.Run("empty string conversion", func(t *testing.T) {
			// Test an edge case with Convert function
			result := Convert[string, int]("")
			assert.Equal(t, 0, result)
		})
	})

	// Test additional coverage for ToStringGeneric
	t.Run("ToStringGeneric edge cases", func(t *testing.T) {
		t.Run("basic types", func(t *testing.T) {
			assert.Equal(t, "true", ToStringGeneric(true))
			assert.Equal(t, "false", ToStringGeneric(false))
			assert.Equal(t, "42", ToStringGeneric(42))
			assert.Equal(t, "test", ToStringGeneric("test"))
		})
	})
}