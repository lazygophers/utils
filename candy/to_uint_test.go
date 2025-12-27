package candy

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestToUint(t *testing.T) {
	t.Run("bool values", func(t *testing.T) {
		assert.Equal(t, uint(1), ToUint(true))
		assert.Equal(t, uint(0), ToUint(false))
	})

	t.Run("int types", func(t *testing.T) {
		assert.Equal(t, uint(42), ToUint(42))
		assert.Equal(t, uint(127), ToUint(int8(127)))
		assert.Equal(t, uint(32767), ToUint(int16(32767)))
		assert.Equal(t, uint(2147483647), ToUint(int32(2147483647)))
		assert.Equal(t, uint(123), ToUint(int64(123)))
	})

	t.Run("negative int types", func(t *testing.T) {
		assert.Equal(t, uint(0), ToUint(-1))
		assert.Equal(t, uint(0), ToUint(int8(-1)))
		assert.Equal(t, uint(0), ToUint(int16(-100)))
		assert.Equal(t, uint(0), ToUint(int32(-500)))
		assert.Equal(t, uint(0), ToUint(int64(-999)))
	})

	t.Run("uint types", func(t *testing.T) {
		assert.Equal(t, uint(42), ToUint(uint(42)))
		assert.Equal(t, uint(255), ToUint(uint8(255)))
		assert.Equal(t, uint(65535), ToUint(uint16(65535)))
		assert.Equal(t, uint(123), ToUint(uint32(123)))
		assert.Equal(t, uint(456), ToUint(uint64(456)))
	})

	t.Run("float types", func(t *testing.T) {
		assert.Equal(t, uint(42), ToUint(float32(42.7)))
		assert.Equal(t, uint(123), ToUint(float64(123.9)))
	})

	t.Run("negative float types", func(t *testing.T) {
		assert.Equal(t, uint(0), ToUint(float32(-42.7)))
		assert.Equal(t, uint(0), ToUint(float64(-123.9)))
	})

	t.Run("string values", func(t *testing.T) {
		assert.Equal(t, uint(123), ToUint("123"))
		assert.Equal(t, uint(4294967295), ToUint("4294967295"))
		assert.Equal(t, uint(0), ToUint("invalid"))
		assert.Equal(t, uint(0), ToUint(""))
		assert.Equal(t, uint(0), ToUint("-1"))
	})

	t.Run("byte slice values", func(t *testing.T) {
		assert.Equal(t, uint(789), ToUint([]byte("789")))
		assert.Equal(t, uint(0), ToUint([]byte("invalid")))
	})

	t.Run("nil and invalid types", func(t *testing.T) {
		assert.Equal(t, uint(0), ToUint(nil))
		assert.Equal(t, uint(0), ToUint(struct{}{}))
		assert.Equal(t, uint(0), ToUint(map[string]int{}))
	})
}

func TestToUint8(t *testing.T) {
	t.Run("bool values", func(t *testing.T) {
		assert.Equal(t, uint8(1), ToUint8(true))
		assert.Equal(t, uint8(0), ToUint8(false))
	})

	t.Run("int types", func(t *testing.T) {
		assert.Equal(t, uint8(42), ToUint8(42))
		assert.Equal(t, uint8(127), ToUint8(int8(127)))
		assert.Equal(t, uint8(100), ToUint8(int16(100)))
		assert.Equal(t, uint8(50), ToUint8(int32(50)))
		assert.Equal(t, uint8(75), ToUint8(int64(75)))
	})

	t.Run("negative int types", func(t *testing.T) {
		assert.Equal(t, uint8(0), ToUint8(-1))
		assert.Equal(t, uint8(0), ToUint8(int8(-1)))
		assert.Equal(t, uint8(0), ToUint8(int16(-100)))
		assert.Equal(t, uint8(0), ToUint8(int32(-50)))
		assert.Equal(t, uint8(0), ToUint8(int64(-75)))
	})

	t.Run("uint types", func(t *testing.T) {
		assert.Equal(t, uint8(42), ToUint8(uint(42)))
		assert.Equal(t, uint8(255), ToUint8(uint8(255)))
		assert.Equal(t, uint8(100), ToUint8(uint16(100)))
		assert.Equal(t, uint8(50), ToUint8(uint32(50)))
		assert.Equal(t, uint8(75), ToUint8(uint64(75)))
	})

	t.Run("float types", func(t *testing.T) {
		assert.Equal(t, uint8(42), ToUint8(float32(42.7)))
		assert.Equal(t, uint8(123), ToUint8(float64(123.9)))
	})

	t.Run("string values", func(t *testing.T) {
		assert.Equal(t, uint8(123), ToUint8("123"))
		assert.Equal(t, uint8(255), ToUint8("255"))
		assert.Equal(t, uint8(0), ToUint8("invalid"))
		assert.Equal(t, uint8(0), ToUint8(""))
		assert.Equal(t, uint8(0), ToUint8("-1"))
	})

	t.Run("byte slice values", func(t *testing.T) {
		assert.Equal(t, uint8(127), ToUint8([]byte("127")))
		assert.Equal(t, uint8(0), ToUint8([]byte("invalid")))
	})

	t.Run("nil and invalid types", func(t *testing.T) {
		assert.Equal(t, uint8(0), ToUint8(nil))
		assert.Equal(t, uint8(0), ToUint8(struct{}{}))
	})
}

func TestToUint16(t *testing.T) {
	t.Run("bool values", func(t *testing.T) {
		assert.Equal(t, uint16(1), ToUint16(true))
		assert.Equal(t, uint16(0), ToUint16(false))
	})

	t.Run("int types", func(t *testing.T) {
		assert.Equal(t, uint16(42), ToUint16(42))
		assert.Equal(t, uint16(127), ToUint16(int8(127)))
		assert.Equal(t, uint16(32767), ToUint16(int16(32767)))
		assert.Equal(t, uint16(1000), ToUint16(int32(1000)))
		assert.Equal(t, uint16(2000), ToUint16(int64(2000)))
	})

	t.Run("negative int types", func(t *testing.T) {
		assert.Equal(t, uint16(0), ToUint16(-1))
		assert.Equal(t, uint16(0), ToUint16(int8(-1)))
		assert.Equal(t, uint16(0), ToUint16(int16(-100)))
		assert.Equal(t, uint16(0), ToUint16(int32(-500)))
		assert.Equal(t, uint16(0), ToUint16(int64(-999)))
	})

	t.Run("uint types", func(t *testing.T) {
		assert.Equal(t, uint16(42), ToUint16(uint(42)))
		assert.Equal(t, uint16(255), ToUint16(uint8(255)))
		assert.Equal(t, uint16(65535), ToUint16(uint16(65535)))
		assert.Equal(t, uint16(1000), ToUint16(uint32(1000)))
		assert.Equal(t, uint16(2000), ToUint16(uint64(2000)))
	})

	t.Run("float types", func(t *testing.T) {
		assert.Equal(t, uint16(42), ToUint16(float32(42.7)))
		assert.Equal(t, uint16(123), ToUint16(float64(123.9)))
	})

	t.Run("string values", func(t *testing.T) {
		assert.Equal(t, uint16(12345), ToUint16("12345"))
		assert.Equal(t, uint16(65535), ToUint16("65535"))
		assert.Equal(t, uint16(0), ToUint16("invalid"))
		assert.Equal(t, uint16(0), ToUint16(""))
		assert.Equal(t, uint16(0), ToUint16("-1"))
	})

	t.Run("byte slice values", func(t *testing.T) {
		assert.Equal(t, uint16(32767), ToUint16([]byte("32767")))
		assert.Equal(t, uint16(0), ToUint16([]byte("invalid")))
	})

	t.Run("nil and invalid types", func(t *testing.T) {
		assert.Equal(t, uint16(0), ToUint16(nil))
		assert.Equal(t, uint16(0), ToUint16(struct{}{}))
	})
}

func TestToUint32(t *testing.T) {
	t.Run("bool values", func(t *testing.T) {
		assert.Equal(t, uint32(1), ToUint32(true))
		assert.Equal(t, uint32(0), ToUint32(false))
	})

	t.Run("int types", func(t *testing.T) {
		assert.Equal(t, uint32(42), ToUint32(42))
		assert.Equal(t, uint32(127), ToUint32(int8(127)))
		assert.Equal(t, uint32(32767), ToUint32(int16(32767)))
		assert.Equal(t, uint32(2147483647), ToUint32(int32(2147483647)))
		assert.Equal(t, uint32(123456), ToUint32(int64(123456)))
	})

	t.Run("negative int types", func(t *testing.T) {
		assert.Equal(t, uint32(0), ToUint32(-1))
		assert.Equal(t, uint32(0), ToUint32(int8(-1)))
		assert.Equal(t, uint32(0), ToUint32(int16(-100)))
		assert.Equal(t, uint32(0), ToUint32(int32(-500)))
		assert.Equal(t, uint32(0), ToUint32(int64(-999)))
	})

	t.Run("uint types", func(t *testing.T) {
		assert.Equal(t, uint32(42), ToUint32(uint(42)))
		assert.Equal(t, uint32(255), ToUint32(uint8(255)))
		assert.Equal(t, uint32(65535), ToUint32(uint16(65535)))
		assert.Equal(t, uint32(4294967295), ToUint32(uint32(4294967295)))
		assert.Equal(t, uint32(123456), ToUint32(uint64(123456)))
	})

	t.Run("float types", func(t *testing.T) {
		assert.Equal(t, uint32(42), ToUint32(float32(42.7)))
		assert.Equal(t, uint32(123), ToUint32(float64(123.9)))
	})

	t.Run("string values", func(t *testing.T) {
		assert.Equal(t, uint32(123456789), ToUint32("123456789"))
		assert.Equal(t, uint32(4294967295), ToUint32("4294967295"))
		assert.Equal(t, uint32(0), ToUint32("invalid"))
		assert.Equal(t, uint32(0), ToUint32(""))
		assert.Equal(t, uint32(0), ToUint32("-1"))
	})

	t.Run("byte slice values", func(t *testing.T) {
		assert.Equal(t, uint32(2147483647), ToUint32([]byte("2147483647")))
		assert.Equal(t, uint32(0), ToUint32([]byte("invalid")))
	})

	t.Run("nil and invalid types", func(t *testing.T) {
		assert.Equal(t, uint32(0), ToUint32(nil))
		assert.Equal(t, uint32(0), ToUint32(struct{}{}))
	})
}

func TestToUint64(t *testing.T) {
	t.Run("bool values", func(t *testing.T) {
		assert.Equal(t, uint64(1), ToUint64(true))
		assert.Equal(t, uint64(0), ToUint64(false))
	})

	t.Run("int types", func(t *testing.T) {
		assert.Equal(t, uint64(42), ToUint64(42))
		assert.Equal(t, uint64(127), ToUint64(int8(127)))
		assert.Equal(t, uint64(32767), ToUint64(int16(32767)))
		assert.Equal(t, uint64(2147483647), ToUint64(int32(2147483647)))
		assert.Equal(t, uint64(9223372036854775807), ToUint64(int64(9223372036854775807)))
	})

	t.Run("negative int types", func(t *testing.T) {
		assert.Equal(t, uint64(0), ToUint64(-1))
		assert.Equal(t, uint64(0), ToUint64(int8(-1)))
		assert.Equal(t, uint64(0), ToUint64(int16(-100)))
		assert.Equal(t, uint64(0), ToUint64(int32(-500)))
		assert.Equal(t, uint64(0), ToUint64(int64(-999)))
	})

	t.Run("uint types", func(t *testing.T) {
		assert.Equal(t, uint64(42), ToUint64(uint(42)))
		assert.Equal(t, uint64(255), ToUint64(uint8(255)))
		assert.Equal(t, uint64(65535), ToUint64(uint16(65535)))
		assert.Equal(t, uint64(4294967295), ToUint64(uint32(4294967295)))
		assert.Equal(t, uint64(18446744073709551615), ToUint64(uint64(18446744073709551615)))
	})

	t.Run("float types", func(t *testing.T) {
		assert.Equal(t, uint64(42), ToUint64(float32(42.7)))
		assert.Equal(t, uint64(123), ToUint64(float64(123.9)))
	})

	t.Run("string values", func(t *testing.T) {
		assert.Equal(t, uint64(123456789012345), ToUint64("123456789012345"))
		assert.Equal(t, uint64(18446744073709551615), ToUint64("18446744073709551615"))
		assert.Equal(t, uint64(0), ToUint64("invalid"))
		assert.Equal(t, uint64(0), ToUint64(""))
		assert.Equal(t, uint64(0), ToUint64("-1"))
	})

	t.Run("byte slice values", func(t *testing.T) {
		assert.Equal(t, uint64(9223372036854775807), ToUint64([]byte("9223372036854775807")))
		assert.Equal(t, uint64(0), ToUint64([]byte("invalid")))
	})

	t.Run("nil and invalid types", func(t *testing.T) {
		assert.Equal(t, uint64(0), ToUint64(nil))
		assert.Equal(t, uint64(0), ToUint64(struct{}{}))
		assert.Equal(t, uint64(0), ToUint64(map[string]int{}))
	})
}
