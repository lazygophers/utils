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
		assert.Equal(t, int64(0), ToInt64([]byte("invalid")))
		assert.Equal(t, int64(0), ToInt64([]byte("")))
		assert.Equal(t, int64(0), ToInt64([]byte("3.14")))
	})

	t.Run("unsupported types", func(t *testing.T) {
		assert.Equal(t, int64(0), ToInt64(nil))
		assert.Equal(t, int64(0), ToInt64(struct{}{}))
	})
}

func TestToInt32(t *testing.T) {
	t.Run("bool values", func(t *testing.T) {
		assert.Equal(t, int32(1), ToInt32(true))
		assert.Equal(t, int32(0), ToInt32(false))
	})

	t.Run("integer types", func(t *testing.T) {
		assert.Equal(t, int32(42), ToInt32(int(42)))
		assert.Equal(t, int32(127), ToInt32(int8(127)))
		assert.Equal(t, int32(32767), ToInt32(int16(32767)))
		assert.Equal(t, int32(100), ToInt32(int32(100)))
		assert.Equal(t, int32(1000), ToInt32(int64(1000)))
		assert.Equal(t, int32(42), ToInt32(uint(42)))
		assert.Equal(t, int32(255), ToInt32(uint8(255)))
		assert.Equal(t, int32(65535), ToInt32(uint16(65535)))
		assert.Equal(t, int32(100), ToInt32(uint32(100)))
		assert.Equal(t, int32(1000), ToInt32(uint64(1000)))
	})

	t.Run("float values", func(t *testing.T) {
		assert.Equal(t, int32(3), ToInt32(float32(3.14)))
		assert.Equal(t, int32(3), ToInt32(float64(3.14)))
		assert.Equal(t, int32(-3), ToInt32(float64(-3.14)))
	})

	t.Run("string values", func(t *testing.T) {
		assert.Equal(t, int32(42), ToInt32("42"))
		assert.Equal(t, int32(0), ToInt32("-42")) // negative strings not supported
		assert.Equal(t, int32(0), ToInt32("0"))
	})

	t.Run("invalid string values", func(t *testing.T) {
		assert.Equal(t, int32(0), ToInt32("invalid"))
		assert.Equal(t, int32(0), ToInt32(""))
		assert.Equal(t, int32(0), ToInt32("3.14"))
	})

	t.Run("byte slice values", func(t *testing.T) {
		assert.Equal(t, int32(42), ToInt32([]byte("42")))
		assert.Equal(t, int32(0), ToInt32([]byte("-42"))) // negative strings not supported
	})

	t.Run("unsupported types", func(t *testing.T) {
		assert.Equal(t, int32(0), ToInt32(nil))
		assert.Equal(t, int32(0), ToInt32(struct{}{}))
	})
}

func TestToInt16(t *testing.T) {
	t.Run("bool values", func(t *testing.T) {
		assert.Equal(t, int16(1), ToInt16(true))
		assert.Equal(t, int16(0), ToInt16(false))
	})

	t.Run("integer types", func(t *testing.T) {
		assert.Equal(t, int16(42), ToInt16(int(42)))
		assert.Equal(t, int16(127), ToInt16(int8(127)))
		assert.Equal(t, int16(1000), ToInt16(int16(1000)))
		assert.Equal(t, int16(100), ToInt16(int32(100)))
		assert.Equal(t, int16(1000), ToInt16(int64(1000)))
		assert.Equal(t, int16(42), ToInt16(uint(42)))
		assert.Equal(t, int16(255), ToInt16(uint8(255)))
		assert.Equal(t, int16(1000), ToInt16(uint16(1000)))
		assert.Equal(t, int16(100), ToInt16(uint32(100)))
		assert.Equal(t, int16(1000), ToInt16(uint64(1000)))
	})

	t.Run("float values", func(t *testing.T) {
		assert.Equal(t, int16(3), ToInt16(float32(3.14)))
		assert.Equal(t, int16(3), ToInt16(float64(3.14)))
	})

	t.Run("string values", func(t *testing.T) {
		assert.Equal(t, int16(42), ToInt16("42"))
		assert.Equal(t, int16(0), ToInt16("-42")) // negative strings not supported due to ParseUint
		assert.Equal(t, int16(0), ToInt16(""))
		assert.Equal(t, int16(0), ToInt16("3.14"))
	})

	t.Run("byte slice values", func(t *testing.T) {
		assert.Equal(t, int16(42), ToInt16([]byte("42")))
		assert.Equal(t, int16(0), ToInt16([]byte("-42"))) // negative not supported
		assert.Equal(t, int16(0), ToInt16([]byte("")))
		assert.Equal(t, int16(0), ToInt16([]byte("invalid")))
	})

	t.Run("invalid values", func(t *testing.T) {
		assert.Equal(t, int16(0), ToInt16("invalid"))
		assert.Equal(t, int16(0), ToInt16(nil))
		assert.Equal(t, int16(0), ToInt16(struct{}{}))
	})
}

func TestToInt8(t *testing.T) {
	t.Run("bool values", func(t *testing.T) {
		assert.Equal(t, int8(1), ToInt8(true))
		assert.Equal(t, int8(0), ToInt8(false))
	})

	t.Run("integer types", func(t *testing.T) {
		assert.Equal(t, int8(42), ToInt8(int(42)))
		assert.Equal(t, int8(127), ToInt8(int8(127)))
		assert.Equal(t, int8(100), ToInt8(int16(100)))
		assert.Equal(t, int8(100), ToInt8(int32(100)))
		assert.Equal(t, int8(100), ToInt8(int64(100)))
		assert.Equal(t, int8(42), ToInt8(uint(42)))
		assert.Equal(t, int8(100), ToInt8(uint8(100)))
		assert.Equal(t, int8(100), ToInt8(uint16(100)))
		assert.Equal(t, int8(100), ToInt8(uint32(100)))
		assert.Equal(t, int8(100), ToInt8(uint64(100)))
	})

	t.Run("float values", func(t *testing.T) {
		assert.Equal(t, int8(3), ToInt8(float32(3.14)))
		assert.Equal(t, int8(3), ToInt8(float64(3.14)))
	})

	t.Run("string values", func(t *testing.T) {
		assert.Equal(t, int8(42), ToInt8("42"))
		assert.Equal(t, int8(0), ToInt8("-42")) // negative strings not supported due to ParseUint
		assert.Equal(t, int8(0), ToInt8(""))
		assert.Equal(t, int8(0), ToInt8("3.14"))
	})

	t.Run("byte slice values", func(t *testing.T) {
		assert.Equal(t, int8(42), ToInt8([]byte("42")))
		assert.Equal(t, int8(0), ToInt8([]byte("-42"))) // negative not supported
		assert.Equal(t, int8(0), ToInt8([]byte("")))
		assert.Equal(t, int8(0), ToInt8([]byte("invalid")))
	})

	t.Run("invalid values", func(t *testing.T) {
		assert.Equal(t, int8(0), ToInt8("invalid"))
		assert.Equal(t, int8(0), ToInt8(nil))
		assert.Equal(t, int8(0), ToInt8(struct{}{}))
	})
}
