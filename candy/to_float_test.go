package candy

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestToFloat32(t *testing.T) {
	t.Run("bool values", func(t *testing.T) {
		assert.Equal(t, float32(1), ToFloat32(true))
		assert.Equal(t, float32(0), ToFloat32(false))
	})

	t.Run("int types", func(t *testing.T) {
		assert.Equal(t, float32(42), ToFloat32(42))
		assert.Equal(t, float32(127), ToFloat32(int8(127)))
		assert.Equal(t, float32(32767), ToFloat32(int16(32767)))
		assert.Equal(t, float32(2147483647), ToFloat32(int32(2147483647)))
		assert.Equal(t, float32(123), ToFloat32(int64(123)))
	})

	t.Run("uint types", func(t *testing.T) {
		assert.Equal(t, float32(42), ToFloat32(uint(42)))
		assert.Equal(t, float32(255), ToFloat32(uint8(255)))
		assert.Equal(t, float32(65535), ToFloat32(uint16(65535)))
		assert.Equal(t, float32(123), ToFloat32(uint32(123)))
		assert.Equal(t, float32(456), ToFloat32(uint64(456)))
	})

	t.Run("float types", func(t *testing.T) {
		assert.Equal(t, float32(42.5), ToFloat32(float32(42.5)))
		assert.InDelta(t, float32(123.456), ToFloat32(float64(123.456)), 0.001)
	})

	t.Run("string values", func(t *testing.T) {
		assert.Equal(t, float32(123.45), ToFloat32("123.45"))
		assert.Equal(t, float32(-456.78), ToFloat32("-456.78"))
		assert.Equal(t, float32(3.14159), ToFloat32("3.14159"))
		assert.Equal(t, float32(0), ToFloat32("invalid"))
		assert.Equal(t, float32(0), ToFloat32(""))
		assert.Equal(t, float32(123), ToFloat32("  123  "))
		assert.Equal(t, float32(45.67), ToFloat32("  45.67  "))
	})

	t.Run("byte slice values", func(t *testing.T) {
		assert.Equal(t, float32(789.12), ToFloat32([]byte("789.12")))
		assert.Equal(t, float32(0), ToFloat32([]byte("invalid")))
		assert.Equal(t, float32(123), ToFloat32([]byte("  123  ")))
	})

	t.Run("nil and invalid types", func(t *testing.T) {
		assert.Equal(t, float32(0), ToFloat32(nil))
		assert.Equal(t, float32(0), ToFloat32(struct{}{}))
		assert.Equal(t, float32(0), ToFloat32(map[string]int{}))
	})

	t.Run("special float values", func(t *testing.T) {
		assert.True(t, math.IsNaN(float64(ToFloat32("NaN"))))
		assert.True(t, math.IsInf(float64(ToFloat32("Inf")), 1))
		assert.True(t, math.IsInf(float64(ToFloat32("-Inf")), -1))
	})
}

func TestToFloat64(t *testing.T) {
	t.Run("bool values", func(t *testing.T) {
		assert.Equal(t, float64(1), ToFloat64(true))
		assert.Equal(t, float64(0), ToFloat64(false))
	})

	t.Run("int types", func(t *testing.T) {
		assert.Equal(t, float64(42), ToFloat64(42))
		assert.Equal(t, float64(127), ToFloat64(int8(127)))
		assert.Equal(t, float64(32767), ToFloat64(int16(32767)))
		assert.Equal(t, float64(2147483647), ToFloat64(int32(2147483647)))
		assert.Equal(t, float64(9223372036854775807), ToFloat64(int64(9223372036854775807)))
	})

	t.Run("uint types", func(t *testing.T) {
		assert.Equal(t, float64(42), ToFloat64(uint(42)))
		assert.Equal(t, float64(255), ToFloat64(uint8(255)))
		assert.Equal(t, float64(65535), ToFloat64(uint16(65535)))
		assert.Equal(t, float64(4294967295), ToFloat64(uint32(4294967295)))
		assert.Equal(t, float64(123456789), ToFloat64(uint64(123456789)))
	})

	t.Run("float types", func(t *testing.T) {
		assert.InDelta(t, float64(42.5), ToFloat64(float32(42.5)), 0.001)
		assert.Equal(t, float64(123.456789), ToFloat64(float64(123.456789)))
	})

	t.Run("string values", func(t *testing.T) {
		assert.Equal(t, float64(123.45), ToFloat64("123.45"))
		assert.Equal(t, float64(-456.78), ToFloat64("-456.78"))
		assert.Equal(t, float64(3.141592653589793), ToFloat64("3.141592653589793"))
		assert.Equal(t, float64(0), ToFloat64("invalid"))
		assert.Equal(t, float64(0), ToFloat64(""))
		assert.Equal(t, float64(123), ToFloat64("  123  "))
		assert.Equal(t, float64(45.67), ToFloat64("  45.67  "))
	})

	t.Run("string integer values", func(t *testing.T) {
		// ToFloat64 在解析失败时会尝试解析为整数
		assert.Equal(t, float64(123), ToFloat64("123"))
		assert.Equal(t, float64(0xff), ToFloat64("0xff"))
		assert.Equal(t, float64(0b101010), ToFloat64("0b101010"))
	})

	t.Run("byte slice values", func(t *testing.T) {
		assert.Equal(t, float64(789.123456), ToFloat64([]byte("789.123456")))
		assert.Equal(t, float64(0), ToFloat64([]byte("invalid")))
		assert.Equal(t, float64(123), ToFloat64([]byte("  123  ")))
		assert.Equal(t, float64(0xff), ToFloat64([]byte("0xff")))
	})

	t.Run("nil and invalid types", func(t *testing.T) {
		assert.Equal(t, float64(0), ToFloat64(nil))
		assert.Equal(t, float64(0), ToFloat64(struct{}{}))
		assert.Equal(t, float64(0), ToFloat64(map[string]int{}))
	})

	t.Run("special float values", func(t *testing.T) {
		assert.True(t, math.IsNaN(ToFloat64("NaN")))
		assert.True(t, math.IsInf(ToFloat64("Inf"), 1))
		assert.True(t, math.IsInf(ToFloat64("-Inf"), -1))
		assert.True(t, math.IsInf(ToFloat64("+Inf"), 1))
	})

	t.Run("scientific notation", func(t *testing.T) {
		assert.Equal(t, float64(1.23e10), ToFloat64("1.23e10"))
		assert.Equal(t, float64(4.56e-5), ToFloat64("4.56e-5"))
		assert.Equal(t, float64(-7.89e12), ToFloat64("-7.89e12"))
	})
}
