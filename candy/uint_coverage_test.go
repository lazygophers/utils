package candy

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUintCoverage(t *testing.T) {
	// Test ToUint32 additional coverage
	t.Run("ToUint32", func(t *testing.T) {
		t.Run("float32 conversion", func(t *testing.T) {
			result := ToUint32(float32(42.7))
			assert.Equal(t, uint32(42), result)
		})

		t.Run("float64 conversion", func(t *testing.T) {
			result := ToUint32(float64(123.9))
			assert.Equal(t, uint32(123), result)
		})

		t.Run("string conversion success", func(t *testing.T) {
			result := ToUint32("12345")
			assert.Equal(t, uint32(12345), result)
		})

		t.Run("string conversion failure", func(t *testing.T) {
			result := ToUint32("invalid")
			assert.Equal(t, uint32(0), result)
		})

		t.Run("negative string", func(t *testing.T) {
			result := ToUint32("-123")
			assert.Equal(t, uint32(0), result) // negative numbers fail for uint parsing
		})

		t.Run("bytes conversion success", func(t *testing.T) {
			result := ToUint32([]byte("54321"))
			assert.Equal(t, uint32(54321), result)
		})

		t.Run("bytes conversion failure", func(t *testing.T) {
			result := ToUint32([]byte("invalid"))
			assert.Equal(t, uint32(0), result)
		})

		t.Run("unsupported type", func(t *testing.T) {
			type CustomType struct{}
			result := ToUint32(CustomType{})
			assert.Equal(t, uint32(0), result)
		})

		t.Run("large uint64 truncation", func(t *testing.T) {
			// Test truncation behavior
			largeVal := uint64(0xFFFFFFFF12345678)
			result := ToUint32(largeVal)
			expected := uint32(0x12345678) // lower 32 bits
			assert.Equal(t, expected, result)
		})

		t.Run("negative float", func(t *testing.T) {
			result := ToUint32(float64(-123.5))
			// Note: converting negative float to uint32 may have platform-specific behavior
			// We just ensure it doesn't panic
			_ = result
		})
	})

	// Test ToUint64 additional coverage
	t.Run("ToUint64", func(t *testing.T) {
		t.Run("float32 conversion", func(t *testing.T) {
			result := ToUint64(float32(42.7))
			assert.Equal(t, uint64(42), result)
		})

		t.Run("float64 conversion", func(t *testing.T) {
			result := ToUint64(float64(123.9))
			assert.Equal(t, uint64(123), result)
		})

		t.Run("string conversion success", func(t *testing.T) {
			result := ToUint64("1234567890")
			assert.Equal(t, uint64(1234567890), result)
		})

		t.Run("string conversion failure", func(t *testing.T) {
			result := ToUint64("invalid")
			assert.Equal(t, uint64(0), result)
		})

		t.Run("negative string", func(t *testing.T) {
			result := ToUint64("-123")
			assert.Equal(t, uint64(0), result) // negative numbers fail for uint parsing
		})

		t.Run("bytes conversion success", func(t *testing.T) {
			result := ToUint64([]byte("9876543210"))
			assert.Equal(t, uint64(9876543210), result)
		})

		t.Run("bytes conversion failure", func(t *testing.T) {
			result := ToUint64([]byte("invalid"))
			assert.Equal(t, uint64(0), result)
		})

		t.Run("unsupported type", func(t *testing.T) {
			type CustomType struct{}
			result := ToUint64(CustomType{})
			assert.Equal(t, uint64(0), result)
		})

		t.Run("negative float", func(t *testing.T) {
			result := ToUint64(float64(-123.5))
			// Note: converting negative float to uint64 may have platform-specific behavior
			// We just ensure it doesn't panic
			_ = result
		})

		t.Run("large numbers", func(t *testing.T) {
			result := ToUint64("18446744073709551615") // max uint64
			assert.Equal(t, uint64(18446744073709551615), result)
		})

		t.Run("overflow string", func(t *testing.T) {
			result := ToUint64("99999999999999999999999") // larger than uint64 max
			assert.Equal(t, uint64(0), result) // should fail parsing
		})
	})
}