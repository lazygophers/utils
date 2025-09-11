package candy

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestToFloat64(t *testing.T) {
	t.Run("bool values", func(t *testing.T) {
		assert.Equal(t, 1.0, ToFloat64(true))
		assert.Equal(t, 0.0, ToFloat64(false))
	})

	t.Run("int values", func(t *testing.T) {
		assert.Equal(t, 0.0, ToFloat64(int(0)))
		assert.Equal(t, 42.0, ToFloat64(int(42)))
		assert.Equal(t, -42.0, ToFloat64(int(-42)))
	})

	t.Run("int8 values", func(t *testing.T) {
		assert.Equal(t, 127.0, ToFloat64(int8(127)))
		assert.Equal(t, -128.0, ToFloat64(int8(-128)))
	})

	t.Run("int16 values", func(t *testing.T) {
		assert.Equal(t, 32767.0, ToFloat64(int16(32767)))
		assert.Equal(t, -32768.0, ToFloat64(int16(-32768)))
	})

	t.Run("int32 values", func(t *testing.T) {
		assert.Equal(t, 100.0, ToFloat64(int32(100)))
		assert.Equal(t, -100.0, ToFloat64(int32(-100)))
	})

	t.Run("int64 values", func(t *testing.T) {
		assert.Equal(t, 1000.0, ToFloat64(int64(1000)))
		assert.Equal(t, -1000.0, ToFloat64(int64(-1000)))
	})

	t.Run("uint values", func(t *testing.T) {
		assert.Equal(t, 0.0, ToFloat64(uint(0)))
		assert.Equal(t, 42.0, ToFloat64(uint(42)))
	})

	t.Run("uint8 values", func(t *testing.T) {
		assert.Equal(t, 255.0, ToFloat64(uint8(255)))
		assert.Equal(t, 0.0, ToFloat64(uint8(0)))
	})

	t.Run("uint16 values", func(t *testing.T) {
		assert.Equal(t, 65535.0, ToFloat64(uint16(65535)))
	})

	t.Run("uint32 values", func(t *testing.T) {
		assert.Equal(t, 100.0, ToFloat64(uint32(100)))
	})

	t.Run("uint64 values", func(t *testing.T) {
		assert.Equal(t, 1000.0, ToFloat64(uint64(1000)))
	})

	t.Run("float32 values", func(t *testing.T) {
		assert.Equal(t, 3.140000104904175, ToFloat64(float32(3.14))) // float32 precision
		assert.Equal(t, -3.140000104904175, ToFloat64(float32(-3.14)))
		assert.Equal(t, 0.0, ToFloat64(float32(0.0)))
	})

	t.Run("float64 values", func(t *testing.T) {
		assert.Equal(t, 3.14, ToFloat64(float64(3.14)))
		assert.Equal(t, -3.14, ToFloat64(float64(-3.14)))
		assert.Equal(t, 0.0, ToFloat64(float64(0.0)))
	})

	t.Run("string float values", func(t *testing.T) {
		assert.Equal(t, 3.14, ToFloat64("3.14"))
		assert.Equal(t, -3.14, ToFloat64("-3.14"))
		assert.Equal(t, 0.0, ToFloat64("0"))
		assert.Equal(t, 3.14, ToFloat64("  3.14  ")) // 空格处理
	})

	t.Run("string int values", func(t *testing.T) {
		assert.Equal(t, 42.0, ToFloat64("42"))
		assert.Equal(t, -42.0, ToFloat64("-42"))
		assert.Equal(t, 0.0, ToFloat64("0"))
	})

	t.Run("string hex values", func(t *testing.T) {
		assert.Equal(t, 255.0, ToFloat64("0xFF"))
		assert.Equal(t, 255.0, ToFloat64("0xff"))
		assert.Equal(t, 8.0, ToFloat64("0o10")) // octal
	})

	t.Run("invalid string values", func(t *testing.T) {
		assert.Equal(t, 0.0, ToFloat64("invalid"))
		assert.Equal(t, 0.0, ToFloat64(""))
		assert.Equal(t, 0.0, ToFloat64("abc"))
		assert.Equal(t, 0.0, ToFloat64("3.14.15"))
	})

	t.Run("byte slice values", func(t *testing.T) {
		assert.Equal(t, 3.14, ToFloat64([]byte("3.14")))
		assert.Equal(t, -3.14, ToFloat64([]byte("-3.14")))
		assert.Equal(t, 42.0, ToFloat64([]byte("42")))
		assert.Equal(t, 255.0, ToFloat64([]byte("0xFF")))
		assert.Equal(t, 3.14, ToFloat64([]byte("  3.14  "))) // 空格处理
	})

	t.Run("invalid byte slice values", func(t *testing.T) {
		assert.Equal(t, 0.0, ToFloat64([]byte("invalid")))
		assert.Equal(t, 0.0, ToFloat64([]byte("")))
		assert.Equal(t, 0.0, ToFloat64([]byte("abc")))
	})

	t.Run("unsupported types", func(t *testing.T) {
		assert.Equal(t, 0.0, ToFloat64(nil))
		assert.Equal(t, 0.0, ToFloat64(struct{}{}))
		assert.Equal(t, 0.0, ToFloat64(map[string]int{}))
		assert.Equal(t, 0.0, ToFloat64([]int{1, 2, 3}))
		assert.Equal(t, 0.0, ToFloat64(func() {}))
	})
}