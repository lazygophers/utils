package candy

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestToFloat32(t *testing.T) {
	t.Run("bool values", func(t *testing.T) {
		assert.Equal(t, float32(1), ToFloat32(true))
		assert.Equal(t, float32(0), ToFloat32(false))
	})

	t.Run("int values", func(t *testing.T) {
		assert.Equal(t, float32(0), ToFloat32(int(0)))
		assert.Equal(t, float32(42), ToFloat32(int(42)))
		assert.Equal(t, float32(-42), ToFloat32(int(-42)))
	})

	t.Run("int8 values", func(t *testing.T) {
		assert.Equal(t, float32(127), ToFloat32(int8(127)))
		assert.Equal(t, float32(-128), ToFloat32(int8(-128)))
	})

	t.Run("int16 values", func(t *testing.T) {
		assert.Equal(t, float32(32767), ToFloat32(int16(32767)))
		assert.Equal(t, float32(-32768), ToFloat32(int16(-32768)))
	})

	t.Run("int32 values", func(t *testing.T) {
		assert.Equal(t, float32(100), ToFloat32(int32(100)))
		assert.Equal(t, float32(-100), ToFloat32(int32(-100)))
	})

	t.Run("int64 values", func(t *testing.T) {
		assert.Equal(t, float32(1000), ToFloat32(int64(1000)))
		assert.Equal(t, float32(-1000), ToFloat32(int64(-1000)))
	})

	t.Run("uint values", func(t *testing.T) {
		assert.Equal(t, float32(0), ToFloat32(uint(0)))
		assert.Equal(t, float32(42), ToFloat32(uint(42)))
	})

	t.Run("uint8 values", func(t *testing.T) {
		assert.Equal(t, float32(255), ToFloat32(uint8(255)))
		assert.Equal(t, float32(0), ToFloat32(uint8(0)))
	})

	t.Run("uint16 values", func(t *testing.T) {
		assert.Equal(t, float32(65535), ToFloat32(uint16(65535)))
	})

	t.Run("uint32 values", func(t *testing.T) {
		assert.Equal(t, float32(100), ToFloat32(uint32(100)))
	})

	t.Run("uint64 values", func(t *testing.T) {
		assert.Equal(t, float32(1000), ToFloat32(uint64(1000)))
	})

	t.Run("float32 values", func(t *testing.T) {
		assert.Equal(t, float32(3.14), ToFloat32(float32(3.14)))
		assert.Equal(t, float32(-3.14), ToFloat32(float32(-3.14)))
		assert.Equal(t, float32(0.0), ToFloat32(float32(0.0)))
	})

	t.Run("float64 values", func(t *testing.T) {
		assert.Equal(t, float32(3.14), ToFloat32(float64(3.14)))
		assert.Equal(t, float32(-3.14), ToFloat32(float64(-3.14)))
		assert.Equal(t, float32(0.0), ToFloat32(float64(0.0)))
	})

	t.Run("string values", func(t *testing.T) {
		assert.Equal(t, float32(3.14), ToFloat32("3.14"))
		assert.Equal(t, float32(-3.14), ToFloat32("-3.14"))
		assert.Equal(t, float32(42), ToFloat32("42"))
		assert.Equal(t, float32(0), ToFloat32("0"))
		assert.Equal(t, float32(3.14), ToFloat32("  3.14  ")) // 空格处理
	})

	t.Run("invalid string values", func(t *testing.T) {
		assert.Equal(t, float32(0), ToFloat32("invalid"))
		assert.Equal(t, float32(0), ToFloat32(""))
		assert.Equal(t, float32(0), ToFloat32("abc"))
		assert.Equal(t, float32(0), ToFloat32("3.14.15"))
	})

	t.Run("byte slice values", func(t *testing.T) {
		assert.Equal(t, float32(3.14), ToFloat32([]byte("3.14")))
		assert.Equal(t, float32(-3.14), ToFloat32([]byte("-3.14")))
		assert.Equal(t, float32(42), ToFloat32([]byte("42")))
		assert.Equal(t, float32(3.14), ToFloat32([]byte("  3.14  "))) // 空格处理
	})

	t.Run("invalid byte slice values", func(t *testing.T) {
		assert.Equal(t, float32(0), ToFloat32([]byte("invalid")))
		assert.Equal(t, float32(0), ToFloat32([]byte("")))
		assert.Equal(t, float32(0), ToFloat32([]byte("abc")))
	})

	t.Run("unsupported types", func(t *testing.T) {
		assert.Equal(t, float32(0), ToFloat32(nil))
		assert.Equal(t, float32(0), ToFloat32(struct{}{}))
		assert.Equal(t, float32(0), ToFloat32(map[string]int{}))
		assert.Equal(t, float32(0), ToFloat32([]int{1, 2, 3}))
		assert.Equal(t, float32(0), ToFloat32(func() {}))
	})
}