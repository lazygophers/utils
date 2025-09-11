package candy

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestToInt64(t *testing.T) {
	t.Run("bool values", func(t *testing.T) {
		assert.Equal(t, int64(1), ToInt64(true))
		assert.Equal(t, int64(0), ToInt64(false))
	})

	t.Run("integer types", func(t *testing.T) {
		assert.Equal(t, int64(42), ToInt64(int(42)))
		assert.Equal(t, int64(127), ToInt64(int8(127)))
		assert.Equal(t, int64(32767), ToInt64(int16(32767)))
		assert.Equal(t, int64(100), ToInt64(int32(100)))
		assert.Equal(t, int64(1000), ToInt64(int64(1000)))
		assert.Equal(t, int64(42), ToInt64(uint(42)))
		assert.Equal(t, int64(255), ToInt64(uint8(255)))
		assert.Equal(t, int64(65535), ToInt64(uint16(65535)))
		assert.Equal(t, int64(100), ToInt64(uint32(100)))
		assert.Equal(t, int64(1000), ToInt64(uint64(1000)))
	})

	t.Run("duration values", func(t *testing.T) {
		assert.Equal(t, int64(1000000000), ToInt64(time.Second))
		assert.Equal(t, int64(1000000), ToInt64(time.Millisecond))
	})

	t.Run("float values", func(t *testing.T) {
		assert.Equal(t, int64(3), ToInt64(float32(3.14)))
		assert.Equal(t, int64(3), ToInt64(float64(3.14)))
		assert.Equal(t, int64(-3), ToInt64(float64(-3.14)))
	})

	t.Run("string values", func(t *testing.T) {
		assert.Equal(t, int64(42), ToInt64("42"))
		assert.Equal(t, int64(-42), ToInt64("-42"))
		assert.Equal(t, int64(0), ToInt64("0"))
	})

	t.Run("invalid string values", func(t *testing.T) {
		assert.Equal(t, int64(0), ToInt64("invalid"))
		assert.Equal(t, int64(0), ToInt64(""))
		assert.Equal(t, int64(0), ToInt64("3.14"))
	})

	t.Run("byte slice values", func(t *testing.T) {
		assert.Equal(t, int64(42), ToInt64([]byte("42")))
		assert.Equal(t, int64(-42), ToInt64([]byte("-42")))
	})

	t.Run("unsupported types", func(t *testing.T) {
		assert.Equal(t, int64(0), ToInt64(nil))
		assert.Equal(t, int64(0), ToInt64(struct{}{}))
	})
}

func TestToInt32(t *testing.T) {
	t.Run("basic conversions", func(t *testing.T) {
		assert.Equal(t, int32(1), ToInt32(true))
		assert.Equal(t, int32(0), ToInt32(false))
		assert.Equal(t, int32(42), ToInt32(int(42)))
		assert.Equal(t, int32(42), ToInt32("42"))
		assert.Equal(t, int32(0), ToInt32("invalid"))
	})
}

func TestToInt16(t *testing.T) {
	t.Run("basic conversions", func(t *testing.T) {
		assert.Equal(t, int16(1), ToInt16(true))
		assert.Equal(t, int16(0), ToInt16(false))
		assert.Equal(t, int16(42), ToInt16(int(42)))
		assert.Equal(t, int16(42), ToInt16("42"))
		assert.Equal(t, int16(0), ToInt16("invalid"))
	})
}

func TestToInt8(t *testing.T) {
	t.Run("basic conversions", func(t *testing.T) {
		assert.Equal(t, int8(1), ToInt8(true))
		assert.Equal(t, int8(0), ToInt8(false))
		assert.Equal(t, int8(42), ToInt8(int(42)))
		assert.Equal(t, int8(42), ToInt8("42"))
		assert.Equal(t, int8(0), ToInt8("invalid"))
	})
}