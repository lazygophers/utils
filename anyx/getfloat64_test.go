package anyx

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMapAny_GetFloat64_Optimized(t *testing.T) {
	t.Run("get float64 value", func(t *testing.T) {
		m := NewMap(map[string]interface{}{
			"key": float64(123.456),
		})
		result := m.GetFloat64("key")
		assert.Equal(t, float64(123.456), result)
	})

	t.Run("get float32 value", func(t *testing.T) {
		m := NewMap(map[string]interface{}{
			"key": float32(78.9),
		})
		result := m.GetFloat64("key")
		assert.InDelta(t, float64(78.9), result, 0.0001)
	})

	t.Run("get int value", func(t *testing.T) {
		m := NewMap(map[string]interface{}{
			"key": int(42),
		})
		result := m.GetFloat64("key")
		assert.Equal(t, float64(42), result)
	})

	t.Run("get int8 value", func(t *testing.T) {
		m := NewMap(map[string]interface{}{
			"key": int8(12),
		})
		result := m.GetFloat64("key")
		assert.Equal(t, float64(12), result)
	})

	t.Run("get int16 value", func(t *testing.T) {
		m := NewMap(map[string]interface{}{
			"key": int16(1234),
		})
		result := m.GetFloat64("key")
		assert.Equal(t, float64(1234), result)
	})

	t.Run("get int32 value", func(t *testing.T) {
		m := NewMap(map[string]interface{}{
			"key": int32(5678),
		})
		result := m.GetFloat64("key")
		assert.Equal(t, float64(5678), result)
	})

	t.Run("get int64 value", func(t *testing.T) {
		m := NewMap(map[string]interface{}{
			"key": int64(1234567890),
		})
		result := m.GetFloat64("key")
		assert.Equal(t, float64(1234567890), result)
	})

	t.Run("get uint value", func(t *testing.T) {
		m := NewMap(map[string]interface{}{
			"key": uint(42),
		})
		result := m.GetFloat64("key")
		assert.Equal(t, float64(42), result)
	})

	t.Run("get uint8 value", func(t *testing.T) {
		m := NewMap(map[string]interface{}{
			"key": uint8(255),
		})
		result := m.GetFloat64("key")
		assert.Equal(t, float64(255), result)
	})

	t.Run("get uint16 value", func(t *testing.T) {
		m := NewMap(map[string]interface{}{
			"key": uint16(65535),
		})
		result := m.GetFloat64("key")
		assert.Equal(t, float64(65535), result)
	})

	t.Run("get uint32 value", func(t *testing.T) {
		m := NewMap(map[string]interface{}{
			"key": uint32(123456),
		})
		result := m.GetFloat64("key")
		assert.Equal(t, float64(123456), result)
	})

	t.Run("get uint64 value", func(t *testing.T) {
		m := NewMap(map[string]interface{}{
			"key": uint64(123456789),
		})
		result := m.GetFloat64("key")
		assert.Equal(t, float64(123456789), result)
	})

	t.Run("get bool true value", func(t *testing.T) {
		m := NewMap(map[string]interface{}{
			"key": true,
		})
		result := m.GetFloat64("key")
		assert.Equal(t, float64(1), result)
	})

	t.Run("get bool false value", func(t *testing.T) {
		m := NewMap(map[string]interface{}{
			"key": false,
		})
		result := m.GetFloat64("key")
		assert.Equal(t, float64(0), result)
	})

	t.Run("get string float value", func(t *testing.T) {
		m := NewMap(map[string]interface{}{
			"key": "123.456",
		})
		result := m.GetFloat64("key")
		assert.Equal(t, float64(123.456), result)
	})

	t.Run("get string int value", func(t *testing.T) {
		m := NewMap(map[string]interface{}{
			"key": "12345",
		})
		result := m.GetFloat64("key")
		assert.Equal(t, float64(12345), result)
	})

	t.Run("get []byte value", func(t *testing.T) {
		m := NewMap(map[string]interface{}{
			"key": []byte("99.9"),
		})
		result := m.GetFloat64("key")
		assert.InDelta(t, float64(99.9), result, 0.0001)
	})

	t.Run("get nil value", func(t *testing.T) {
		m := NewMap(map[string]interface{}{
			"key": nil,
		})
		result := m.GetFloat64("key")
		assert.Equal(t, float64(0), result)
	})

	t.Run("get non-existent key", func(t *testing.T) {
		m := NewMap(map[string]interface{}{})
		result := m.GetFloat64("nonexistent")
		assert.Equal(t, float64(0), result)
	})

	t.Run("get invalid string value", func(t *testing.T) {
		m := NewMap(map[string]interface{}{
			"key": "not_a_number",
		})
		result := m.GetFloat64("key")
		assert.Equal(t, float64(0), result)
	})

	t.Run("get invalid type value", func(t *testing.T) {
		m := NewMap(map[string]interface{}{
			"key": map[string]int{},
		})
		result := m.GetFloat64("key")
		assert.Equal(t, float64(0), result)
	})
}

func TestMapAny_GetFloat64_Performance(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping performance test in short mode")
	}

	m := NewMap(map[string]interface{}{
		"float64": float64(123.456),
		"int":     int(42),
		"string":  "99.9",
	})

	t.Run("performance test - float64 key", func(t *testing.T) {
		for i := 0; i < 1000000; i++ {
			_ = m.GetFloat64("float64")
		}
	})

	t.Run("performance test - int key", func(t *testing.T) {
		for i := 0; i < 1000000; i++ {
			_ = m.GetFloat64("int")
		}
	})

	t.Run("performance test - string key", func(t *testing.T) {
		for i := 0; i < 1000000; i++ {
			_ = m.GetFloat64("string")
		}
	})
}
