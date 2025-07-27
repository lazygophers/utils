package anyx

import (
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMapAny_BasicOperations(t *testing.T) {
	t.Run("Set and Get", func(t *testing.T) {
		m := NewMap(nil)
		m.Set("name", "Alice")
		val, err := m.Get("name")
		assert.NoError(t, err)
		assert.Equal(t, "Alice", val)
	})

	t.Run("Exists", func(t *testing.T) {
		m := NewMap(nil)
		m.Set("age", 30)
		assert.True(t, m.Exists("age"))
		assert.False(t, m.Exists("invalid"))
	})

	t.Run("Keys and Values", func(t *testing.T) {
		m := NewMap(map[string]interface{}{
			"k1": "v1",
			"k2": 2,
		})
		keys := []string{}
		values := []interface{}{}
		m.data.Range(func(key, value interface{}) bool {
			keys = append(keys, key.(string))
			values = append(values, value)
			return true
		})
		assert.ElementsMatch(t, []string{"k1", "k2"}, keys)
		assert.ElementsMatch(t, []interface{}{"v1", 2}, values)
	})
}

func TestMapAny_Concurrency(t *testing.T) {
	m := NewMap(nil)
	var wg sync.WaitGroup

	// 并发写入
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func(idx int) {
			defer wg.Done()
			m.Set("key", idx)
		}(i)
	}

	// 并发读取
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			m.Get("key")
		}()
	}

	wg.Wait()
	val, _ := m.Get("key")
	assert.NotNil(t, val)
}

func TestMapAny_BoundaryConditions(t *testing.T) {
	t.Run("Nil Value", func(t *testing.T) {
		m := NewMap(nil)
		m.Set("nil", nil)
		val, err := m.Get("nil")
		assert.NoError(t, err)
		assert.Nil(t, val)
	})

	t.Run("Empty Map", func(t *testing.T) {
		m := NewMap(nil)
		val, err := m.Get("missing")
		assert.ErrorIs(t, err, ErrNotFound)
		assert.Nil(t, val)
	})

	t.Run("Type Conversion", func(t *testing.T) {
		m := NewMap(map[string]interface{}{"num": "not_a_number"})
		assert.Equal(t, 0, m.GetInt("num"))
		assert.Equal(t, "", m.GetString("invalid"))
	})
}

func TestMapAny_TypeAccessors(t *testing.T) {
	m := NewMap(map[string]interface{}{
		"bool":   true,
		"int":    42,
		"string": "hello",
		"float":  3.14,
	})

	tests := []struct {
		key      string
		expected interface{}
	}{
		{"bool", true},
		{"int", 42},
		{"string", "hello"},
		{"float", 3.14},
	}

	for _, tt := range tests {
		t.Run(tt.key, func(t *testing.T) {
			switch tt.expected.(type) {
			case bool:
				assert.Equal(t, tt.expected, m.GetBool(tt.key))
			case int:
				assert.Equal(t, tt.expected, m.GetInt(tt.key))
			case string:
				assert.Equal(t, tt.expected, m.GetString(tt.key))
			case float64:
				assert.Equal(t, tt.expected, m.GetFloat64(tt.key))
			}
		})
	}
}

func TestMapAny_Range(t *testing.T) {
	m := NewMap(map[string]interface{}{
		"a": 1,
		"b": "hello",
		"c": true,
	})

	keys := make(map[string]interface{})
	count := 0

	m.Range(func(key, value interface{}) bool {
		keys[key.(string)] = value
		count++
		return true
	})

	assert.Equal(t, 3, count)
	assert.Equal(t, 1, keys["a"])
	assert.Equal(t, "hello", keys["b"])
	assert.Equal(t, true, keys["c"])

	t.Run("Stop Range", func(t *testing.T) {
		innerCount := 0
		m.Range(func(key, value interface{}) bool {
			innerCount++
			return false // Stop after first item
		})
		assert.Equal(t, 1, innerCount)
	})
}
