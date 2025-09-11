package candy

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestToInt(t *testing.T) {
	t.Run("bool values", func(t *testing.T) {
		assert.Equal(t, 1, ToInt(true))
		assert.Equal(t, 0, ToInt(false))
	})

	t.Run("int values", func(t *testing.T) {
		assert.Equal(t, 0, ToInt(int(0)))
		assert.Equal(t, 42, ToInt(int(42)))
		assert.Equal(t, -42, ToInt(int(-42)))
	})

	t.Run("int8 values", func(t *testing.T) {
		assert.Equal(t, 127, ToInt(int8(127)))
		assert.Equal(t, -128, ToInt(int8(-128)))
	})

	t.Run("int16 values", func(t *testing.T) {
		assert.Equal(t, 32767, ToInt(int16(32767)))
		assert.Equal(t, -32768, ToInt(int16(-32768)))
	})

	t.Run("int32 values", func(t *testing.T) {
		assert.Equal(t, 100, ToInt(int32(100)))
		assert.Equal(t, -100, ToInt(int32(-100)))
	})

	t.Run("int64 values", func(t *testing.T) {
		assert.Equal(t, 1000, ToInt(int64(1000)))
		assert.Equal(t, -1000, ToInt(int64(-1000)))
	})

	t.Run("uint values", func(t *testing.T) {
		assert.Equal(t, 0, ToInt(uint(0)))
		assert.Equal(t, 42, ToInt(uint(42)))
	})

	t.Run("uint8 values", func(t *testing.T) {
		assert.Equal(t, 255, ToInt(uint8(255)))
		assert.Equal(t, 0, ToInt(uint8(0)))
	})

	t.Run("uint16 values", func(t *testing.T) {
		assert.Equal(t, 65535, ToInt(uint16(65535)))
	})

	t.Run("uint32 values", func(t *testing.T) {
		assert.Equal(t, 100, ToInt(uint32(100)))
	})

	t.Run("uint64 values", func(t *testing.T) {
		assert.Equal(t, 1000, ToInt(uint64(1000)))
	})

	t.Run("float32 values", func(t *testing.T) {
		assert.Equal(t, 3, ToInt(float32(3.14)))
		assert.Equal(t, -3, ToInt(float32(-3.14)))
		assert.Equal(t, 0, ToInt(float32(0.0)))
		assert.Equal(t, 0, ToInt(float32(0.9))) // 截断小数部分
	})

	t.Run("float64 values", func(t *testing.T) {
		assert.Equal(t, 3, ToInt(float64(3.14)))
		assert.Equal(t, -3, ToInt(float64(-3.14)))
		assert.Equal(t, 0, ToInt(float64(0.0)))
		assert.Equal(t, 0, ToInt(float64(0.9))) // 截断小数部分
	})

	t.Run("string values", func(t *testing.T) {
		assert.Equal(t, 42, ToInt("42"))
		assert.Equal(t, 0, ToInt("0"))
		assert.Equal(t, 123, ToInt("123"))
	})

	t.Run("invalid string values", func(t *testing.T) {
		// 注意：ToInt使用ParseUint，不支持负数字符串
		assert.Equal(t, 0, ToInt("-42"))
		assert.Equal(t, 0, ToInt("invalid"))
		assert.Equal(t, 0, ToInt(""))
		assert.Equal(t, 0, ToInt("abc"))
		assert.Equal(t, 0, ToInt("3.14"))
	})

	t.Run("byte slice values", func(t *testing.T) {
		assert.Equal(t, 42, ToInt([]byte("42")))
		assert.Equal(t, 0, ToInt([]byte("0")))
		assert.Equal(t, 123, ToInt([]byte("123")))
	})

	t.Run("invalid byte slice values", func(t *testing.T) {
		assert.Equal(t, 0, ToInt([]byte("-42")))
		assert.Equal(t, 0, ToInt([]byte("invalid")))
		assert.Equal(t, 0, ToInt([]byte("")))
		assert.Equal(t, 0, ToInt([]byte("abc")))
	})

	t.Run("unsupported types", func(t *testing.T) {
		assert.Equal(t, 0, ToInt(nil))
		assert.Equal(t, 0, ToInt(struct{}{}))
		assert.Equal(t, 0, ToInt(map[string]int{}))
		assert.Equal(t, 0, ToInt([]int{1, 2, 3}))
		assert.Equal(t, 0, ToInt(func() {}))
	})
}