package candy

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestToInt(t *testing.T) {
	t.Run("bool values", func(t *testing.T) {
		assert.Equal(t, 1, ToInt(true))
		assert.Equal(t, 0, ToInt(false))
	})

	t.Run("int types", func(t *testing.T) {
		assert.Equal(t, 42, ToInt(42))
		assert.Equal(t, 127, ToInt(int8(127)))
		assert.Equal(t, 32767, ToInt(int16(32767)))
		assert.Equal(t, 2147483647, ToInt(int32(2147483647)))
		assert.Equal(t, 123, ToInt(int64(123)))
	})

	t.Run("uint types", func(t *testing.T) {
		assert.Equal(t, 42, ToInt(uint(42)))
		assert.Equal(t, 255, ToInt(uint8(255)))
		assert.Equal(t, 65535, ToInt(uint16(65535)))
		assert.Equal(t, 123, ToInt(uint32(123)))
		assert.Equal(t, 456, ToInt(uint64(456)))
	})

	t.Run("float types", func(t *testing.T) {
		assert.Equal(t, 42, ToInt(float32(42.7)))
		assert.Equal(t, 123, ToInt(float64(123.9)))
	})

	t.Run("string values", func(t *testing.T) {
		assert.Equal(t, 123, ToInt("123"))
		assert.Equal(t, -456, ToInt("-456"))
		assert.Equal(t, 0, ToInt("invalid"))
		assert.Equal(t, 0, ToInt(""))
	})

	t.Run("byte slice values", func(t *testing.T) {
		assert.Equal(t, 789, ToInt([]byte("789")))
		assert.Equal(t, 0, ToInt([]byte("invalid")))
	})

	t.Run("nil and invalid types", func(t *testing.T) {
		assert.Equal(t, 0, ToInt(nil))
		assert.Equal(t, 0, ToInt(struct{}{}))
		assert.Equal(t, 0, ToInt(map[string]int{}))
	})
}

func TestToInt8(t *testing.T) {
	t.Run("bool values", func(t *testing.T) {
		assert.Equal(t, int8(1), ToInt8(true))
		assert.Equal(t, int8(0), ToInt8(false))
	})

	t.Run("int types", func(t *testing.T) {
		assert.Equal(t, int8(42), ToInt8(42))
		assert.Equal(t, int8(127), ToInt8(int8(127)))
		assert.Equal(t, int8(100), ToInt8(int16(100)))
		assert.Equal(t, int8(50), ToInt8(int32(50)))
		assert.Equal(t, int8(75), ToInt8(int64(75)))
	})

	t.Run("uint types", func(t *testing.T) {
		assert.Equal(t, int8(42), ToInt8(uint(42)))
		assert.Equal(t, int8(-1), ToInt8(uint8(255)))
		assert.Equal(t, int8(100), ToInt8(uint16(100)))
		assert.Equal(t, int8(50), ToInt8(uint32(50)))
		assert.Equal(t, int8(75), ToInt8(uint64(75)))
	})

	t.Run("float types", func(t *testing.T) {
		assert.Equal(t, int8(42), ToInt8(float32(42.7)))
		assert.Equal(t, int8(123), ToInt8(float64(123.9)))
	})

	t.Run("string values", func(t *testing.T) {
		assert.Equal(t, int8(123), ToInt8("123"))
		assert.Equal(t, int8(-127), ToInt8("-127"))
		assert.Equal(t, int8(0), ToInt8("invalid"))
		assert.Equal(t, int8(0), ToInt8(""))
	})

	t.Run("byte slice values", func(t *testing.T) {
		assert.Equal(t, int8(127), ToInt8([]byte("127")))
		assert.Equal(t, int8(0), ToInt8([]byte("invalid")))
	})

	t.Run("nil and invalid types", func(t *testing.T) {
		assert.Equal(t, int8(0), ToInt8(nil))
		assert.Equal(t, int8(0), ToInt8(struct{}{}))
	})
}

func TestToInt16(t *testing.T) {
	t.Run("bool values", func(t *testing.T) {
		assert.Equal(t, int16(1), ToInt16(true))
		assert.Equal(t, int16(0), ToInt16(false))
	})

	t.Run("int types", func(t *testing.T) {
		assert.Equal(t, int16(42), ToInt16(42))
		assert.Equal(t, int16(127), ToInt16(int8(127)))
		assert.Equal(t, int16(32767), ToInt16(int16(32767)))
		assert.Equal(t, int16(1000), ToInt16(int32(1000)))
		assert.Equal(t, int16(2000), ToInt16(int64(2000)))
	})

	t.Run("uint types", func(t *testing.T) {
		assert.Equal(t, int16(42), ToInt16(uint(42)))
		assert.Equal(t, int16(255), ToInt16(uint8(255)))
		assert.Equal(t, int16(-1), ToInt16(uint16(65535)))
		assert.Equal(t, int16(1000), ToInt16(uint32(1000)))
		assert.Equal(t, int16(2000), ToInt16(uint64(2000)))
	})

	t.Run("float types", func(t *testing.T) {
		assert.Equal(t, int16(42), ToInt16(float32(42.7)))
		assert.Equal(t, int16(123), ToInt16(float64(123.9)))
	})

	t.Run("string values", func(t *testing.T) {
		assert.Equal(t, int16(12345), ToInt16("12345"))
		assert.Equal(t, int16(-32767), ToInt16("-32767"))
		assert.Equal(t, int16(0), ToInt16("invalid"))
		assert.Equal(t, int16(0), ToInt16(""))
	})

	t.Run("byte slice values", func(t *testing.T) {
		assert.Equal(t, int16(32767), ToInt16([]byte("32767")))
		assert.Equal(t, int16(0), ToInt16([]byte("invalid")))
	})

	t.Run("nil and invalid types", func(t *testing.T) {
		assert.Equal(t, int16(0), ToInt16(nil))
		assert.Equal(t, int16(0), ToInt16(struct{}{}))
	})
}

func TestToInt32(t *testing.T) {
	t.Run("bool values", func(t *testing.T) {
		assert.Equal(t, int32(1), ToInt32(true))
		assert.Equal(t, int32(0), ToInt32(false))
	})

	t.Run("int types", func(t *testing.T) {
		assert.Equal(t, int32(42), ToInt32(42))
		assert.Equal(t, int32(127), ToInt32(int8(127)))
		assert.Equal(t, int32(32767), ToInt32(int16(32767)))
		assert.Equal(t, int32(2147483647), ToInt32(int32(2147483647)))
		assert.Equal(t, int32(123456), ToInt32(int64(123456)))
	})

	t.Run("uint types", func(t *testing.T) {
		assert.Equal(t, int32(42), ToInt32(uint(42)))
		assert.Equal(t, int32(255), ToInt32(uint8(255)))
		assert.Equal(t, int32(65535), ToInt32(uint16(65535)))
		assert.Equal(t, int32(-1), ToInt32(uint32(4294967295)))
		assert.Equal(t, int32(123456), ToInt32(uint64(123456)))
	})

	t.Run("float types", func(t *testing.T) {
		assert.Equal(t, int32(42), ToInt32(float32(42.7)))
		assert.Equal(t, int32(123), ToInt32(float64(123.9)))
	})

	t.Run("string values", func(t *testing.T) {
		assert.Equal(t, int32(123456789), ToInt32("123456789"))
		assert.Equal(t, int32(-2147483647), ToInt32("-2147483647"))
		assert.Equal(t, int32(0), ToInt32("invalid"))
		assert.Equal(t, int32(0), ToInt32(""))
	})

	t.Run("byte slice values", func(t *testing.T) {
		assert.Equal(t, int32(2147483647), ToInt32([]byte("2147483647")))
		assert.Equal(t, int32(0), ToInt32([]byte("invalid")))
	})

	t.Run("nil and invalid types", func(t *testing.T) {
		assert.Equal(t, int32(0), ToInt32(nil))
		assert.Equal(t, int32(0), ToInt32(struct{}{}))
	})
}

func TestToInt64(t *testing.T) {
	t.Run("bool values", func(t *testing.T) {
		assert.Equal(t, int64(1), ToInt64(true))
		assert.Equal(t, int64(0), ToInt64(false))
	})

	t.Run("int types", func(t *testing.T) {
		assert.Equal(t, int64(42), ToInt64(42))
		assert.Equal(t, int64(127), ToInt64(int8(127)))
		assert.Equal(t, int64(32767), ToInt64(int16(32767)))
		assert.Equal(t, int64(2147483647), ToInt64(int32(2147483647)))
		assert.Equal(t, int64(9223372036854775807), ToInt64(int64(9223372036854775807)))
	})

	t.Run("uint types", func(t *testing.T) {
		assert.Equal(t, int64(42), ToInt64(uint(42)))
		assert.Equal(t, int64(255), ToInt64(uint8(255)))
		assert.Equal(t, int64(65535), ToInt64(uint16(65535)))
		assert.Equal(t, int64(4294967295), ToInt64(uint32(4294967295)))
		assert.Equal(t, int64(123456789), ToInt64(uint64(123456789)))
	})

	t.Run("time.Duration", func(t *testing.T) {
		assert.Equal(t, int64(1000000000), ToInt64(time.Second))
		assert.Equal(t, int64(60000000000), ToInt64(time.Minute))
	})

	t.Run("float types", func(t *testing.T) {
		assert.Equal(t, int64(42), ToInt64(float32(42.7)))
		assert.Equal(t, int64(123), ToInt64(float64(123.9)))
	})

	t.Run("string values", func(t *testing.T) {
		assert.Equal(t, int64(123456789012345), ToInt64("123456789012345"))
		assert.Equal(t, int64(-9223372036854775807), ToInt64("-9223372036854775807"))
		assert.Equal(t, int64(0), ToInt64("invalid"))
		assert.Equal(t, int64(0), ToInt64(""))
	})

	t.Run("byte slice values", func(t *testing.T) {
		assert.Equal(t, int64(9223372036854775807), ToInt64([]byte("9223372036854775807")))
		assert.Equal(t, int64(0), ToInt64([]byte("invalid")))
	})

	t.Run("nil and invalid types", func(t *testing.T) {
		assert.Equal(t, int64(0), ToInt64(nil))
		assert.Equal(t, int64(0), ToInt64(struct{}{}))
		assert.Equal(t, int64(0), ToInt64(map[string]int{}))
	})
}
