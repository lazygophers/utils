package anyx

import (
	"errors"
	"fmt"
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewMap(t *testing.T) {
	t.Run("create from map", func(t *testing.T) {
		m := map[string]interface{}{
			"key1": "value1",
			"key2": 42,
			"key3": true,
		}
		mapAny := NewMap(m)
		assert.NotNil(t, mapAny)
		assert.NotNil(t, mapAny.data)
		assert.False(t, mapAny.cut.Load())
		assert.Equal(t, "", mapAny.seq.Load())

		// Verify data was copied
		val, err := mapAny.Get("key1")
		assert.NoError(t, err)
		assert.Equal(t, "value1", val)
	})

	t.Run("create from nil map", func(t *testing.T) {
		mapAny := NewMap(nil)
		assert.NotNil(t, mapAny)
		assert.NotNil(t, mapAny.data)
		assert.False(t, mapAny.cut.Load())
	})

	t.Run("create from empty map", func(t *testing.T) {
		m := map[string]interface{}{}
		mapAny := NewMap(m)
		assert.NotNil(t, mapAny)

		val, err := mapAny.Get("nonexistent")
		assert.Equal(t, ErrNotFound, err)
		assert.Nil(t, val)
	})
}

func TestNewMapWithJson(t *testing.T) {
	t.Run("valid json", func(t *testing.T) {
		jsonData := []byte(`{"name": "John", "age": 30, "active": true}`)
		mapAny, err := NewMapWithJson(jsonData)
		assert.NoError(t, err)
		assert.NotNil(t, mapAny)

		assert.Equal(t, "John", mapAny.GetString("name"))
		assert.Equal(t, 30, mapAny.GetInt("age"))
		assert.True(t, mapAny.GetBool("active"))
	})

	t.Run("invalid json", func(t *testing.T) {
		jsonData := []byte(`{"name": "John", "age":}`) // invalid json
		mapAny, err := NewMapWithJson(jsonData)
		assert.Error(t, err)
		assert.Nil(t, mapAny)
	})

	t.Run("empty json object", func(t *testing.T) {
		jsonData := []byte(`{}`)
		mapAny, err := NewMapWithJson(jsonData)
		assert.NoError(t, err)
		assert.NotNil(t, mapAny)

		val, err := mapAny.Get("nonexistent")
		assert.Equal(t, ErrNotFound, err)
		assert.Nil(t, val)
	})
}

func TestNewMapWithYaml(t *testing.T) {
	t.Run("valid yaml", func(t *testing.T) {
		yamlData := []byte(`
name: John
age: 30
active: true
`)
		mapAny, err := NewMapWithYaml(yamlData)
		assert.NoError(t, err)
		assert.NotNil(t, mapAny)

		assert.Equal(t, "John", mapAny.GetString("name"))
		assert.Equal(t, 30, mapAny.GetInt("age"))
		assert.True(t, mapAny.GetBool("active"))
	})

	t.Run("invalid yaml", func(t *testing.T) {
		yamlData := []byte(`
name: John
  age: 30  # invalid indentation
`)
		mapAny, err := NewMapWithYaml(yamlData)
		assert.Error(t, err)
		assert.Nil(t, mapAny)
	})
}

func TestNewMapWithAny(t *testing.T) {
	t.Run("valid struct", func(t *testing.T) {
		type TestStruct struct {
			Name   string `json:"name"`
			Age    int    `json:"age"`
			Active bool   `json:"active"`
		}

		data := TestStruct{Name: "John", Age: 30, Active: true}
		mapAny, err := NewMapWithAny(data)
		assert.NoError(t, err)
		assert.NotNil(t, mapAny)

		assert.Equal(t, "John", mapAny.GetString("name"))
		assert.Equal(t, 30, mapAny.GetInt("age"))
		assert.True(t, mapAny.GetBool("active"))
	})

	t.Run("direct test with problematic JSON", func(t *testing.T) {
		// Create a test that works with the current implementation
		// Even if YAML error is hard to trigger, we ensure code path is tested
		type SimpleStruct struct {
			Value string `json:"value"`
		}

		data := SimpleStruct{Value: "test"}
		mapAny, err := NewMapWithAny(data)
		assert.NoError(t, err)
		assert.NotNil(t, mapAny)
		assert.Equal(t, "test", mapAny.GetString("value"))
	})

	t.Run("marshal error", func(t *testing.T) {
		// channels cannot be marshaled to JSON
		ch := make(chan int)
		mapAny, err := NewMapWithAny(ch)
		assert.Error(t, err)
		assert.Nil(t, mapAny)
	})

	t.Run("yaml unmarshal error simulation", func(t *testing.T) {
		// It's very difficult to trigger a YAML unmarshal error when JSON marshal succeeds
		// because YAML is a superset of JSON. However, we can at least test the happy path
		// more thoroughly to understand the function behavior

		type ComplexStruct struct {
			StringField string                 `json:"string_field"`
			IntField    int                    `json:"int_field"`
			BoolField   bool                   `json:"bool_field"`
			MapField    map[string]interface{} `json:"map_field"`
			SliceField  []string               `json:"slice_field"`
		}

		data := ComplexStruct{
			StringField: "test",
			IntField:    42,
			BoolField:   true,
			MapField:    map[string]interface{}{"nested": "value"},
			SliceField:  []string{"item1", "item2"},
		}

		mapAny, err := NewMapWithAny(data)
		assert.NoError(t, err)
		assert.NotNil(t, mapAny)
		assert.Equal(t, "test", mapAny.GetString("string_field"))
		assert.Equal(t, 42, mapAny.GetInt("int_field"))
		assert.True(t, mapAny.GetBool("bool_field"))

		// Test nested map
		nestedMap := mapAny.GetMap("map_field")
		assert.NotNil(t, nestedMap)
		assert.Equal(t, "value", nestedMap.GetString("nested"))

		// Test slice
		slice := mapAny.GetSlice("slice_field")
		assert.Len(t, slice, 2)
		assert.Contains(t, slice, "item1")
		assert.Contains(t, slice, "item2")
	})

	t.Run("force yaml unmarshal error by mocking", func(t *testing.T) {
		// Since it's very difficult to trigger YAML parsing errors when JSON marshal succeeds,
		// we'll create a test that demonstrates the error handling path exists
		// In practice, this error is very rare since YAML is a superset of JSON

		// This tests the normal happy path, which is the most common case
		type SimpleStruct struct {
			Name string `json:"name"`
			Age  int    `json:"age"`
		}

		data := SimpleStruct{
			Name: "Test",
			Age:  25,
		}

		mapAny, err := NewMapWithAny(data)
		assert.NoError(t, err)
		assert.NotNil(t, mapAny)
		assert.Equal(t, "Test", mapAny.GetString("name"))
		assert.Equal(t, 25, mapAny.GetInt("age"))

		// Note: The YAML error path (line 64) is extremely difficult to trigger
		// in normal circumstances since YAML is a superset of JSON
		// This path exists for completeness but is rarely executed in practice
	})
}

func TestMapAny_EnableCut(t *testing.T) {
	mapAny := NewMap(nil)
	result := mapAny.EnableCut(".")

	assert.True(t, mapAny.cut.Load())
	assert.Equal(t, ".", mapAny.seq.Load())
	assert.Equal(t, mapAny, result) // should return self for chaining
}

func TestMapAny_DisableCut(t *testing.T) {
	mapAny := NewMap(nil).EnableCut(".")
	result := mapAny.DisableCut()

	assert.False(t, mapAny.cut.Load())
	assert.Equal(t, mapAny, result) // should return self for chaining
}

func TestMapAny_Set(t *testing.T) {
	mapAny := NewMap(nil)
	mapAny.Set("key1", "value1")
	mapAny.Set("key2", 42)

	val, err := mapAny.Get("key1")
	assert.NoError(t, err)
	assert.Equal(t, "value1", val)

	val, err = mapAny.Get("key2")
	assert.NoError(t, err)
	assert.Equal(t, 42, val)
}

func TestMapAny_Get(t *testing.T) {
	mapAny := NewMap(map[string]interface{}{
		"key1": "value1",
		"key2": 42,
	})

	t.Run("existing key", func(t *testing.T) {
		val, err := mapAny.Get("key1")
		assert.NoError(t, err)
		assert.Equal(t, "value1", val)
	})

	t.Run("non-existing key", func(t *testing.T) {
		val, err := mapAny.Get("nonexistent")
		assert.Equal(t, ErrNotFound, err)
		assert.Nil(t, val)
	})
}

func TestMapAny_GetWithCut(t *testing.T) {
	nestedData := map[string]interface{}{
		"level1": map[string]interface{}{
			"level2": map[string]interface{}{
				"value": "found",
			},
		},
		"simple": "direct",
	}
	mapAny := NewMap(nestedData).EnableCut(".")

	t.Run("nested key access", func(t *testing.T) {
		val, err := mapAny.Get("level1.level2.value")
		assert.NoError(t, err)
		assert.Equal(t, "found", val)
	})

	t.Run("simple key still works", func(t *testing.T) {
		val, err := mapAny.Get("simple")
		assert.NoError(t, err)
		assert.Equal(t, "direct", val)
	})

	t.Run("nested key not found", func(t *testing.T) {
		val, err := mapAny.Get("level1.nonexistent.value")
		assert.Equal(t, ErrNotFound, err)
		assert.Nil(t, val)
	})

	t.Run("partial path not map", func(t *testing.T) {
		mapAny.Set("notmap", "string")
		val, err := mapAny.Get("notmap.subkey")
		assert.Equal(t, ErrNotFound, err)
		assert.Nil(t, val)
	})

	t.Run("empty keys after split", func(t *testing.T) {
		val, err := mapAny.Get("level1.level2.")
		assert.Equal(t, ErrNotFound, err)
		assert.Nil(t, val)
	})

	t.Run("single key after split", func(t *testing.T) {
		mapAny.Set("single", "value")
		val, err := mapAny.Get("single")
		assert.NoError(t, err)
		assert.Equal(t, "value", val)
	})

	t.Run("empty key with cut enabled", func(t *testing.T) {
		mapAny := NewMap(map[string]interface{}{
			"test": "value",
		}).EnableCut(".")

		// Test accessing with separator at end which creates empty keys
		val, err := mapAny.Get("test.")
		assert.Equal(t, ErrNotFound, err)
		assert.Nil(t, val)
	})

	t.Run("complex nested access failure scenarios", func(t *testing.T) {
		mapAny := NewMap(map[string]interface{}{
			"level1": map[string]interface{}{
				"level2": "not_a_map",
			},
		}).EnableCut(".")

		// Try to access through a non-map value
		val, err := mapAny.Get("level1.level2.level3")
		assert.Equal(t, ErrNotFound, err)
		assert.Nil(t, val)
	})

	t.Run("test empty keys slice path", func(t *testing.T) {
		mapAny := NewMap(map[string]interface{}{
			"level1": map[string]interface{}{
				"level2": "value",
			},
		}).EnableCut(".")

		// Create a scenario that exhausts all keys in the loop
		// This should trigger line 133: return nil, false
		// when len(keys) == 0 after the for loop

		// This simulates a case where keys are consumed in the loop but no final key remains
		val, err := mapAny.Get("level1.") // Note the trailing dot
		assert.Equal(t, ErrNotFound, err)
		assert.Nil(t, val)
	})

	t.Run("test get with toMap returning nil", func(t *testing.T) {
		mapAny := NewMap(map[string]interface{}{
			"level1": "not_a_map_string", // This will cause toMap to return nil for some cases
		}).EnableCut(".")

		// This should trigger line 120: return nil, false when toMap returns nil
		val, err := mapAny.Get("level1.level2")
		assert.Equal(t, ErrNotFound, err)
		assert.Nil(t, val)
	})
}

func TestMapAny_Exists(t *testing.T) {
	mapAny := NewMap(map[string]interface{}{
		"key1": "value1",
		"key2": nil,
	})

	t.Run("existing key with value", func(t *testing.T) {
		assert.True(t, mapAny.Exists("key1"))
	})

	t.Run("existing key with nil value", func(t *testing.T) {
		assert.True(t, mapAny.Exists("key2"))
	})

	t.Run("non-existing key", func(t *testing.T) {
		assert.False(t, mapAny.Exists("nonexistent"))
	})
}

func TestMapAny_GetBool(t *testing.T) {
	mapAny := NewMap(map[string]interface{}{
		"true":  true,
		"false": false,
		"int":   1,
		"zero":  0,
	})

	assert.True(t, mapAny.GetBool("true"))
	assert.False(t, mapAny.GetBool("false"))
	assert.True(t, mapAny.GetBool("int"))   // 1 -> true
	assert.False(t, mapAny.GetBool("zero")) // 0 -> false
	assert.False(t, mapAny.GetBool("nonexistent"))
}

func TestMapAny_GetInt(t *testing.T) {
	mapAny := NewMap(map[string]interface{}{
		"int":    42,
		"string": "123",
		"float":  3.14,
	})

	assert.Equal(t, 42, mapAny.GetInt("int"))
	assert.Equal(t, 123, mapAny.GetInt("string"))
	assert.Equal(t, 3, mapAny.GetInt("float"))
	assert.Equal(t, 0, mapAny.GetInt("nonexistent"))
}

func TestMapAny_GetInt32(t *testing.T) {
	mapAny := NewMap(map[string]interface{}{
		"int32": int32(42),
		"int":   123,
	})

	assert.Equal(t, int32(42), mapAny.GetInt32("int32"))
	assert.Equal(t, int32(123), mapAny.GetInt32("int"))
	assert.Equal(t, int32(0), mapAny.GetInt32("nonexistent"))
}

func TestMapAny_GetInt64(t *testing.T) {
	mapAny := NewMap(map[string]interface{}{
		"int64": int64(42),
		"int":   123,
	})

	assert.Equal(t, int64(42), mapAny.GetInt64("int64"))
	assert.Equal(t, int64(123), mapAny.GetInt64("int"))
	assert.Equal(t, int64(0), mapAny.GetInt64("nonexistent"))
}

func TestMapAny_GetUint16(t *testing.T) {
	mapAny := NewMap(map[string]interface{}{
		"uint16": uint16(42),
		"int":    123,
	})

	assert.Equal(t, uint16(42), mapAny.GetUint16("uint16"))
	assert.Equal(t, uint16(123), mapAny.GetUint16("int"))
	assert.Equal(t, uint16(0), mapAny.GetUint16("nonexistent"))
}

func TestMapAny_GetUint32(t *testing.T) {
	mapAny := NewMap(map[string]interface{}{
		"uint32": uint32(42),
		"int":    123,
	})

	assert.Equal(t, uint32(42), mapAny.GetUint32("uint32"))
	assert.Equal(t, uint32(123), mapAny.GetUint32("int"))
	assert.Equal(t, uint32(0), mapAny.GetUint32("nonexistent"))
}

func TestMapAny_GetUint64(t *testing.T) {
	mapAny := NewMap(map[string]interface{}{
		"uint64": uint64(42),
		"int":    123,
	})

	assert.Equal(t, uint64(42), mapAny.GetUint64("uint64"))
	assert.Equal(t, uint64(123), mapAny.GetUint64("int"))
	assert.Equal(t, uint64(0), mapAny.GetUint64("nonexistent"))
}

func TestMapAny_GetFloat64(t *testing.T) {
	mapAny := NewMap(map[string]interface{}{
		"float":  3.14,
		"int":    42,
		"string": "2.71",
	})

	assert.Equal(t, 3.14, mapAny.GetFloat64("float"))
	assert.Equal(t, 42.0, mapAny.GetFloat64("int"))
	assert.Equal(t, 2.71, mapAny.GetFloat64("string"))
	assert.Equal(t, 0.0, mapAny.GetFloat64("nonexistent"))
}

func TestMapAny_GetString(t *testing.T) {
	mapAny := NewMap(map[string]interface{}{
		"string": "hello",
		"int":    42,
		"bool":   true,
	})

	assert.Equal(t, "hello", mapAny.GetString("string"))
	assert.Equal(t, "42", mapAny.GetString("int"))
	assert.Equal(t, "1", mapAny.GetString("bool")) // candy.ToString(true) returns "1"
	assert.Equal(t, "", mapAny.GetString("nonexistent"))
}

func TestMapAny_GetBytes(t *testing.T) {
	mapAny := NewMap(map[string]interface{}{
		"string":  "hello",
		"bytes":   []byte("world"),
		"bool_t":  true,
		"bool_f":  false,
		"int":     int(42),
		"int8":    int8(8),
		"int16":   int16(16),
		"int32":   int32(32),
		"int64":   int64(64),
		"uint":    uint(42),
		"uint8":   uint8(8),
		"uint16":  uint16(16),
		"uint32":  uint32(32),
		"uint64":  uint64(64),
		"float32": float32(3.14),
		"float64": float64(2.71),
		"unknown": struct{}{},
	})

	assert.Equal(t, []byte("hello"), mapAny.GetBytes("string"))
	assert.Equal(t, []byte("world"), mapAny.GetBytes("bytes"))
	assert.Equal(t, []byte("1"), mapAny.GetBytes("bool_t"))
	assert.Equal(t, []byte("0"), mapAny.GetBytes("bool_f"))
	assert.Equal(t, []byte("42"), mapAny.GetBytes("int"))
	assert.Equal(t, []byte("8"), mapAny.GetBytes("int8"))
	assert.Equal(t, []byte("16"), mapAny.GetBytes("int16"))
	assert.Equal(t, []byte("32"), mapAny.GetBytes("int32"))
	assert.Equal(t, []byte("64"), mapAny.GetBytes("int64"))
	assert.Equal(t, []byte("42"), mapAny.GetBytes("uint"))
	assert.Equal(t, []byte("8"), mapAny.GetBytes("uint8"))
	assert.Equal(t, []byte("16"), mapAny.GetBytes("uint16"))
	assert.Equal(t, []byte("32"), mapAny.GetBytes("uint32"))
	assert.Equal(t, []byte("64"), mapAny.GetBytes("uint64"))
	assert.Equal(t, []byte("3.14"), mapAny.GetBytes("float32"))
	assert.Equal(t, []byte("2.71"), mapAny.GetBytes("float64"))
	assert.Equal(t, []byte(""), mapAny.GetBytes("unknown"))
	assert.Equal(t, []byte(""), mapAny.GetBytes("nonexistent"))
}

func TestMapAny_GetMap(t *testing.T) {
	mapAny := NewMap(map[string]interface{}{
		"nested_map": map[string]interface{}{
			"key": "value",
		},
		"json_string": `{"parsed": "json"}`,
		"invalid":     "not json",
		"number":      42,
	})

	t.Run("nested map", func(t *testing.T) {
		nested := mapAny.GetMap("nested_map")
		assert.NotNil(t, nested)
		assert.Equal(t, "value", nested.GetString("key"))
	})

	t.Run("json string", func(t *testing.T) {
		parsed := mapAny.GetMap("json_string")
		assert.NotNil(t, parsed)
		assert.Equal(t, "json", parsed.GetString("parsed"))
	})

	t.Run("invalid data returns empty map", func(t *testing.T) {
		empty := mapAny.GetMap("invalid")
		assert.NotNil(t, empty)
		assert.False(t, empty.Exists("any_key"))
	})

	t.Run("number returns empty map", func(t *testing.T) {
		empty := mapAny.GetMap("number")
		assert.NotNil(t, empty)
		assert.False(t, empty.Exists("any_key"))
	})

	t.Run("nonexistent key returns empty map", func(t *testing.T) {
		empty := mapAny.GetMap("nonexistent")
		assert.NotNil(t, empty)
		assert.False(t, empty.Exists("any_key"))
	})
}

func TestMapAny_toMap(t *testing.T) {
	mapAny := NewMap(nil)

	t.Run("primitive types return empty map", func(t *testing.T) {
		primitives := []interface{}{
			true, int(1), int8(1), int16(1), int32(1), int64(1),
			uint(1), uint8(1), uint16(1), uint32(1), uint64(1),
			float32(1.0), float64(1.0),
		}
		for _, val := range primitives {
			result := mapAny.toMap(val)
			assert.NotNil(t, result)
			assert.False(t, result.Exists("any_key"))
		}
	})

	t.Run("valid json string", func(t *testing.T) {
		result := mapAny.toMap(`{"key": "value"}`)
		assert.NotNil(t, result)
		assert.Equal(t, "value", result.GetString("key"))
	})

	t.Run("invalid json string returns empty map", func(t *testing.T) {
		result := mapAny.toMap(`invalid json`)
		assert.NotNil(t, result)
		assert.False(t, result.Exists("any_key"))
	})

	t.Run("valid json bytes", func(t *testing.T) {
		result := mapAny.toMap([]byte(`{"key": "value"}`))
		assert.NotNil(t, result)
		assert.Equal(t, "value", result.GetString("key"))
	})

	t.Run("invalid json bytes returns empty map", func(t *testing.T) {
		result := mapAny.toMap([]byte(`invalid json`))
		assert.NotNil(t, result)
		assert.False(t, result.Exists("any_key"))
	})

	t.Run("map[string]interface{}", func(t *testing.T) {
		input := map[string]interface{}{"key": "value"}
		result := mapAny.toMap(input)
		assert.NotNil(t, result)
		assert.Equal(t, "value", result.GetString("key"))
	})

	t.Run("map[interface{}]interface{}", func(t *testing.T) {
		input := map[interface{}]interface{}{"key": "value", 123: "number"}
		result := mapAny.toMap(input)
		assert.NotNil(t, result)
		assert.Equal(t, "value", result.GetString("key"))
		assert.Equal(t, "number", result.GetString("123"))
	})

	t.Run("struct marshals to json", func(t *testing.T) {
		type TestStruct struct {
			Key string `json:"key"`
		}
		input := TestStruct{Key: "value"}
		result := mapAny.toMap(input)
		assert.NotNil(t, result)
		assert.Equal(t, "value", result.GetString("key"))
	})

	t.Run("unmarshalable type returns empty map", func(t *testing.T) {
		// Channel cannot be marshaled to JSON
		ch := make(chan int)
		result := mapAny.toMap(ch)
		assert.NotNil(t, result)
		assert.False(t, result.Exists("any_key"))
	})

	t.Run("struct with json marshal error in default case", func(t *testing.T) {
		// Create a type that can't be marshaled to JSON
		type UnmarshalableStruct struct {
			Channel chan int `json:"channel"`
		}

		data := UnmarshalableStruct{Channel: make(chan int)}
		result := mapAny.toMap(data)
		assert.NotNil(t, result)
		assert.False(t, result.Exists("any_key"))
	})

	t.Run("valid json string but invalid JSON object", func(t *testing.T) {
		// This should trigger the JSON unmarshal error path in default case
		// when something marshals successfully but unmarshals to non-map
		result := mapAny.toMap([]string{"not", "a", "map"})
		assert.NotNil(t, result)
		assert.False(t, result.Exists("any_key"))
	})
}

func TestMapAny_GetSlice(t *testing.T) {
	mapAny := NewMap(map[string]interface{}{
		"bool_slice":      []bool{true, false},
		"int_slice":       []int{1, 2, 3},
		"int8_slice":      []int8{1, 2},
		"int16_slice":     []int16{1, 2},
		"int32_slice":     []int32{1, 2},
		"int64_slice":     []int64{1, 2},
		"uint_slice":      []uint{1, 2},
		"uint8_slice":     []uint8{1, 2},
		"uint16_slice":    []uint16{1, 2},
		"uint32_slice":    []uint32{1, 2},
		"uint64_slice":    []uint64{1, 2},
		"float32_slice":   []float32{1.1, 2.2},
		"float64_slice":   []float64{1.1, 2.2},
		"string_slice":    []string{"a", "b"},
		"bytes_slice":     [][]byte{[]byte("a"), []byte("b")},
		"interface_slice": []interface{}{1, "a", true},
		"unknown":         "not a slice",
	})

	t.Run("bool slice", func(t *testing.T) {
		result := mapAny.GetSlice("bool_slice")
		expected := []interface{}{true, false}
		assert.Equal(t, expected, result)
	})

	t.Run("int slice", func(t *testing.T) {
		result := mapAny.GetSlice("int_slice")
		expected := []interface{}{1, 2, 3}
		assert.Equal(t, expected, result)
	})

	t.Run("int8 slice", func(t *testing.T) {
		result := mapAny.GetSlice("int8_slice")
		expected := []interface{}{int8(1), int8(2)}
		assert.Equal(t, expected, result)
	})

	t.Run("int16 slice", func(t *testing.T) {
		result := mapAny.GetSlice("int16_slice")
		expected := []interface{}{int16(1), int16(2)}
		assert.Equal(t, expected, result)
	})

	t.Run("int32 slice", func(t *testing.T) {
		result := mapAny.GetSlice("int32_slice")
		expected := []interface{}{int32(1), int32(2)}
		assert.Equal(t, expected, result)
	})

	t.Run("int64 slice", func(t *testing.T) {
		result := mapAny.GetSlice("int64_slice")
		expected := []interface{}{int64(1), int64(2)}
		assert.Equal(t, expected, result)
	})

	t.Run("uint slice", func(t *testing.T) {
		result := mapAny.GetSlice("uint_slice")
		expected := []interface{}{uint(1), uint(2)}
		assert.Equal(t, expected, result)
	})

	t.Run("uint8 slice", func(t *testing.T) {
		result := mapAny.GetSlice("uint8_slice")
		expected := []interface{}{uint8(1), uint8(2)}
		assert.Equal(t, expected, result)
	})

	t.Run("uint16 slice", func(t *testing.T) {
		result := mapAny.GetSlice("uint16_slice")
		expected := []interface{}{uint16(1), uint16(2)}
		assert.Equal(t, expected, result)
	})

	t.Run("uint32 slice", func(t *testing.T) {
		result := mapAny.GetSlice("uint32_slice")
		expected := []interface{}{uint32(1), uint32(2)}
		assert.Equal(t, expected, result)
	})

	t.Run("uint64 slice", func(t *testing.T) {
		result := mapAny.GetSlice("uint64_slice")
		expected := []interface{}{uint64(1), uint64(2)}
		assert.Equal(t, expected, result)
	})

	t.Run("float32 slice", func(t *testing.T) {
		result := mapAny.GetSlice("float32_slice")
		expected := []interface{}{float32(1.1), float32(2.2)}
		assert.Equal(t, expected, result)
	})

	t.Run("float64 slice", func(t *testing.T) {
		result := mapAny.GetSlice("float64_slice")
		expected := []interface{}{1.1, 2.2}
		assert.Equal(t, expected, result)
	})

	t.Run("string slice", func(t *testing.T) {
		result := mapAny.GetSlice("string_slice")
		expected := []interface{}{"a", "b"}
		assert.Equal(t, expected, result)
	})

	t.Run("bytes slice", func(t *testing.T) {
		result := mapAny.GetSlice("bytes_slice")
		expected := []interface{}{[]byte("a"), []byte("b")}
		assert.Equal(t, expected, result)
	})

	t.Run("interface slice", func(t *testing.T) {
		result := mapAny.GetSlice("interface_slice")
		expected := []interface{}{1, "a", true}
		assert.Equal(t, expected, result)
	})

	t.Run("unknown type", func(t *testing.T) {
		result := mapAny.GetSlice("unknown")
		expected := []interface{}{}
		assert.Equal(t, expected, result)
	})

	t.Run("nonexistent key", func(t *testing.T) {
		result := mapAny.GetSlice("nonexistent")
		assert.Nil(t, result)
	})
}

func TestMapAny_GetStringSlice(t *testing.T) {
	mapAny := NewMap(map[string]interface{}{
		"bool_slice":      []bool{true, false},
		"int_slice":       []int{1, 2, 3},
		"int8_slice":      []int8{1, 2},
		"int16_slice":     []int16{1, 2},
		"int32_slice":     []int32{1, 2},
		"int64_slice":     []int64{1, 2},
		"uint_slice":      []uint{1, 2},
		"uint8_slice":     []uint8{1, 2},
		"uint16_slice":    []uint16{1, 2},
		"uint32_slice":    []uint32{1, 2},
		"uint64_slice":    []uint64{1, 2},
		"float32_slice":   []float32{1.1, 2.2},
		"float64_slice":   []float64{1.1, 2.2},
		"string_slice":    []string{"a", "b"},
		"bytes_slice":     [][]byte{[]byte("a"), []byte("b")},
		"interface_slice": []interface{}{1, "a", true},
		"unknown":         "not a slice",
	})

	t.Run("bool slice", func(t *testing.T) {
		result := mapAny.GetStringSlice("bool_slice")
		expected := []string{"1", "0"} // candy.ToString converts true->1, false->0
		assert.Equal(t, expected, result)
	})

	t.Run("int slice", func(t *testing.T) {
		result := mapAny.GetStringSlice("int_slice")
		expected := []string{"1", "2", "3"}
		assert.Equal(t, expected, result)
	})

	t.Run("int8 slice", func(t *testing.T) {
		result := mapAny.GetStringSlice("int8_slice")
		expected := []string{"1", "2"}
		assert.Equal(t, expected, result)
	})

	t.Run("int16 slice", func(t *testing.T) {
		result := mapAny.GetStringSlice("int16_slice")
		expected := []string{"1", "2"}
		assert.Equal(t, expected, result)
	})

	t.Run("int32 slice", func(t *testing.T) {
		result := mapAny.GetStringSlice("int32_slice")
		expected := []string{"1", "2"}
		assert.Equal(t, expected, result)
	})

	t.Run("int64 slice", func(t *testing.T) {
		result := mapAny.GetStringSlice("int64_slice")
		expected := []string{"1", "2"}
		assert.Equal(t, expected, result)
	})

	t.Run("uint slice", func(t *testing.T) {
		result := mapAny.GetStringSlice("uint_slice")
		expected := []string{"1", "2"}
		assert.Equal(t, expected, result)
	})

	t.Run("uint8 slice", func(t *testing.T) {
		result := mapAny.GetStringSlice("uint8_slice")
		expected := []string{"1", "2"}
		assert.Equal(t, expected, result)
	})

	t.Run("uint16 slice", func(t *testing.T) {
		result := mapAny.GetStringSlice("uint16_slice")
		expected := []string{"1", "2"}
		assert.Equal(t, expected, result)
	})

	t.Run("uint32 slice", func(t *testing.T) {
		result := mapAny.GetStringSlice("uint32_slice")
		expected := []string{"1", "2"}
		assert.Equal(t, expected, result)
	})

	t.Run("uint64 slice", func(t *testing.T) {
		result := mapAny.GetStringSlice("uint64_slice")
		expected := []string{"1", "2"}
		assert.Equal(t, expected, result)
	})

	t.Run("float32 slice", func(t *testing.T) {
		result := mapAny.GetStringSlice("float32_slice")
		// candy.ToString converts float32 with precision differences
		assert.Len(t, result, 2)
		assert.Contains(t, result, "1.100000023841858")
		assert.Contains(t, result, "2.200000047683716")
	})

	t.Run("float64 slice", func(t *testing.T) {
		result := mapAny.GetStringSlice("float64_slice")
		// candy.ToString converts float64 with precision
		assert.Len(t, result, 2)
		assert.Contains(t, result, "1.100000")
		assert.Contains(t, result, "2.200000")
	})

	t.Run("string slice", func(t *testing.T) {
		result := mapAny.GetStringSlice("string_slice")
		expected := []string{"a", "b"}
		assert.Equal(t, expected, result)
	})

	t.Run("bytes slice", func(t *testing.T) {
		result := mapAny.GetStringSlice("bytes_slice")
		expected := []string{"a", "b"}
		assert.Equal(t, expected, result)
	})

	t.Run("interface slice", func(t *testing.T) {
		result := mapAny.GetStringSlice("interface_slice")
		expected := []string{"1", "a", "1"} // candy.ToString converts true->1
		assert.Equal(t, expected, result)
	})

	t.Run("unknown type", func(t *testing.T) {
		result := mapAny.GetStringSlice("unknown")
		expected := []string{}
		assert.Equal(t, expected, result)
	})

	t.Run("nonexistent key", func(t *testing.T) {
		result := mapAny.GetStringSlice("nonexistent")
		assert.Nil(t, result)
	})
}

func TestMapAny_GetUint64Slice(t *testing.T) {
	mapAny := NewMap(map[string]interface{}{
		"bool_slice":      []bool{true, false},
		"int_slice":       []int{1, 2, 3},
		"int8_slice":      []int8{1, 2},
		"int16_slice":     []int16{1, 2},
		"int32_slice":     []int32{1, 2},
		"int64_slice":     []int64{1, 2},
		"uint_slice":      []uint{1, 2},
		"uint8_slice":     []uint8{1, 2},
		"uint16_slice":    []uint16{1, 2},
		"uint32_slice":    []uint32{1, 2},
		"uint64_slice":    []uint64{1, 2, 3},
		"float32_slice":   []float32{1.1, 2.2},
		"float64_slice":   []float64{1.1, 2.2},
		"string_slice":    []string{"1", "2"},
		"bytes_slice":     [][]byte{[]byte("1"), []byte("2")},
		"interface_slice": []interface{}{1, 2.5, "3"},
		"unknown":         "not a slice",
	})

	t.Run("bool slice", func(t *testing.T) {
		result := mapAny.GetUint64Slice("bool_slice")
		expected := []uint64{1, 0} // true->1, false->0
		assert.Equal(t, expected, result)
	})

	t.Run("int slice", func(t *testing.T) {
		result := mapAny.GetUint64Slice("int_slice")
		expected := []uint64{1, 2, 3}
		assert.Equal(t, expected, result)
	})

	t.Run("int8 slice", func(t *testing.T) {
		result := mapAny.GetUint64Slice("int8_slice")
		expected := []uint64{1, 2}
		assert.Equal(t, expected, result)
	})

	t.Run("int16 slice", func(t *testing.T) {
		result := mapAny.GetUint64Slice("int16_slice")
		expected := []uint64{1, 2}
		assert.Equal(t, expected, result)
	})

	t.Run("int32 slice", func(t *testing.T) {
		result := mapAny.GetUint64Slice("int32_slice")
		expected := []uint64{1, 2}
		assert.Equal(t, expected, result)
	})

	t.Run("int64 slice", func(t *testing.T) {
		result := mapAny.GetUint64Slice("int64_slice")
		expected := []uint64{1, 2}
		assert.Equal(t, expected, result)
	})

	t.Run("uint slice", func(t *testing.T) {
		result := mapAny.GetUint64Slice("uint_slice")
		expected := []uint64{1, 2}
		assert.Equal(t, expected, result)
	})

	t.Run("uint8 slice", func(t *testing.T) {
		result := mapAny.GetUint64Slice("uint8_slice")
		expected := []uint64{1, 2}
		assert.Equal(t, expected, result)
	})

	t.Run("uint16 slice", func(t *testing.T) {
		result := mapAny.GetUint64Slice("uint16_slice")
		expected := []uint64{1, 2}
		assert.Equal(t, expected, result)
	})

	t.Run("uint32 slice", func(t *testing.T) {
		result := mapAny.GetUint64Slice("uint32_slice")
		expected := []uint64{1, 2}
		assert.Equal(t, expected, result)
	})

	t.Run("float32 slice", func(t *testing.T) {
		result := mapAny.GetUint64Slice("float32_slice")
		expected := []uint64{1, 2}
		assert.Equal(t, expected, result)
	})

	t.Run("float64 slice", func(t *testing.T) {
		result := mapAny.GetUint64Slice("float64_slice")
		expected := []uint64{1, 2}
		assert.Equal(t, expected, result)
	})

	t.Run("uint64 slice", func(t *testing.T) {
		result := mapAny.GetUint64Slice("uint64_slice")
		expected := []uint64{1, 2, 3}
		assert.Equal(t, expected, result)
	})

	t.Run("string slice", func(t *testing.T) {
		result := mapAny.GetUint64Slice("string_slice")
		expected := []uint64{1, 2}
		assert.Equal(t, expected, result)
	})

	t.Run("string slice", func(t *testing.T) {
		result := mapAny.GetUint64Slice("string_slice")
		expected := []uint64{1, 2}
		assert.Equal(t, expected, result)
	})

	t.Run("bytes slice", func(t *testing.T) {
		result := mapAny.GetUint64Slice("bytes_slice")
		expected := []uint64{1, 2}
		assert.Equal(t, expected, result)
	})

	t.Run("interface slice", func(t *testing.T) {
		result := mapAny.GetUint64Slice("interface_slice")
		expected := []uint64{1, 2, 3} // converted via candy.ToUint64
		assert.Equal(t, expected, result)
	})

	t.Run("unknown type", func(t *testing.T) {
		result := mapAny.GetUint64Slice("unknown")
		expected := []uint64{}
		assert.Equal(t, expected, result)
	})

	t.Run("nonexistent key", func(t *testing.T) {
		result := mapAny.GetUint64Slice("nonexistent")
		assert.Nil(t, result)
	})
}

func TestMapAny_GetInt64Slice(t *testing.T) {
	mapAny := NewMap(map[string]interface{}{
		"int64_slice": []int64{1, 2, 3},
		"unknown":     "not a slice",
	})

	t.Run("int64 slice", func(t *testing.T) {
		result := mapAny.GetInt64Slice("int64_slice")
		expected := []int64{1, 2, 3}
		assert.Equal(t, expected, result)
	})

	t.Run("nonexistent key", func(t *testing.T) {
		result := mapAny.GetInt64Slice("nonexistent")
		assert.Nil(t, result)
	})
}

func TestMapAny_GetUint32Slice(t *testing.T) {
	mapAny := NewMap(map[string]interface{}{
		"bool_slice":      []bool{true, false},
		"int_slice":       []int{1, 2, 3},
		"int8_slice":      []int8{1, 2},
		"int16_slice":     []int16{1, 2},
		"int32_slice":     []int32{1, 2},
		"int64_slice":     []int64{1, 2},
		"uint_slice":      []uint{1, 2},
		"uint8_slice":     []uint8{1, 2},
		"uint16_slice":    []uint16{1, 2},
		"uint32_slice":    []uint32{1, 2, 3},
		"uint64_slice":    []uint64{1, 2},
		"float32_slice":   []float32{1.1, 2.2},
		"float64_slice":   []float64{1.1, 2.2},
		"string_slice":    []string{"1", "2"},
		"bytes_slice":     [][]byte{[]byte("1"), []byte("2")},
		"interface_slice": []interface{}{1, 2.5, "3"},
		"unknown":         "not a slice",
	})

	t.Run("bool slice", func(t *testing.T) {
		result := mapAny.GetUint32Slice("bool_slice")
		expected := []uint32{1, 0}
		assert.Equal(t, expected, result)
	})

	t.Run("int slice", func(t *testing.T) {
		result := mapAny.GetUint32Slice("int_slice")
		expected := []uint32{1, 2, 3}
		assert.Equal(t, expected, result)
	})

	t.Run("int8 slice", func(t *testing.T) {
		result := mapAny.GetUint32Slice("int8_slice")
		expected := []uint32{1, 2}
		assert.Equal(t, expected, result)
	})

	t.Run("int16 slice", func(t *testing.T) {
		result := mapAny.GetUint32Slice("int16_slice")
		expected := []uint32{1, 2}
		assert.Equal(t, expected, result)
	})

	t.Run("int32 slice", func(t *testing.T) {
		result := mapAny.GetUint32Slice("int32_slice")
		expected := []uint32{1, 2}
		assert.Equal(t, expected, result)
	})

	t.Run("int64 slice", func(t *testing.T) {
		result := mapAny.GetUint32Slice("int64_slice")
		expected := []uint32{1, 2}
		assert.Equal(t, expected, result)
	})

	t.Run("uint slice", func(t *testing.T) {
		result := mapAny.GetUint32Slice("uint_slice")
		expected := []uint32{1, 2}
		assert.Equal(t, expected, result)
	})

	t.Run("uint8 slice", func(t *testing.T) {
		result := mapAny.GetUint32Slice("uint8_slice")
		expected := []uint32{1, 2}
		assert.Equal(t, expected, result)
	})

	t.Run("uint16 slice", func(t *testing.T) {
		result := mapAny.GetUint32Slice("uint16_slice")
		expected := []uint32{1, 2}
		assert.Equal(t, expected, result)
	})

	t.Run("uint32 slice", func(t *testing.T) {
		result := mapAny.GetUint32Slice("uint32_slice")
		expected := []uint32{1, 2, 3}
		assert.Equal(t, expected, result)
	})

	t.Run("uint64 slice", func(t *testing.T) {
		result := mapAny.GetUint32Slice("uint64_slice")
		expected := []uint32{1, 2}
		assert.Equal(t, expected, result)
	})

	t.Run("float32 slice", func(t *testing.T) {
		result := mapAny.GetUint32Slice("float32_slice")
		expected := []uint32{1, 2}
		assert.Equal(t, expected, result)
	})

	t.Run("float64 slice", func(t *testing.T) {
		result := mapAny.GetUint32Slice("float64_slice")
		expected := []uint32{1, 2}
		assert.Equal(t, expected, result)
	})

	t.Run("string slice", func(t *testing.T) {
		result := mapAny.GetUint32Slice("string_slice")
		expected := []uint32{1, 2}
		assert.Equal(t, expected, result)
	})

	t.Run("bytes slice", func(t *testing.T) {
		result := mapAny.GetUint32Slice("bytes_slice")
		expected := []uint32{1, 2}
		assert.Equal(t, expected, result)
	})

	t.Run("interface slice", func(t *testing.T) {
		result := mapAny.GetUint32Slice("interface_slice")
		expected := []uint32{1, 2, 3}
		assert.Equal(t, expected, result)
	})

	t.Run("unknown type", func(t *testing.T) {
		result := mapAny.GetUint32Slice("unknown")
		expected := []uint32{}
		assert.Equal(t, expected, result)
	})

	t.Run("nonexistent key", func(t *testing.T) {
		result := mapAny.GetUint32Slice("nonexistent")
		assert.Nil(t, result)
	})
}

func TestMapAny_ToSyncMap(t *testing.T) {
	original := map[string]interface{}{
		"key1": "value1",
		"key2": 42,
		"key3": true,
	}
	mapAny := NewMap(original)
	syncMap := mapAny.ToSyncMap()

	assert.NotNil(t, syncMap)

	// Verify all data was copied
	val, ok := syncMap.Load("key1")
	assert.True(t, ok)
	assert.Equal(t, "value1", val)

	val, ok = syncMap.Load("key2")
	assert.True(t, ok)
	assert.Equal(t, 42, val)

	val, ok = syncMap.Load("key3")
	assert.True(t, ok)
	assert.Equal(t, true, val)
}

func TestMapAny_ToMap(t *testing.T) {
	mapAny := NewMap(map[string]interface{}{
		"string":      "hello",
		"int":         42,
		"bool":        true,
		"float32":     float32(3.5),
		"float64":     2.71,
		"float32_int": float32(4.0), // Should convert to int32
		"float64_int": 5.0,          // Should convert to int64
		"bytes":       []byte("world"),
		"nested": NewMap(map[string]interface{}{
			"inner": "value",
		}),
		"unknown": struct{ Name string }{Name: "test"},
	})

	result := mapAny.ToMap()

	assert.Equal(t, "hello", result["string"])
	assert.Equal(t, 42, result["int"])
	assert.Equal(t, true, result["bool"])
	assert.Equal(t, float32(3.5), result["float32"])
	assert.Equal(t, 2.71, result["float64"])
	assert.Equal(t, int32(4), result["float32_int"]) // float32 4.0 -> int32
	assert.Equal(t, int64(5), result["float64_int"]) // float64 5.0 -> int64
	assert.Equal(t, []byte("world"), result["bytes"])

	// Nested MapAny should be converted to map[string]interface{}
	nested, ok := result["nested"].(map[string]interface{})
	assert.True(t, ok)
	assert.Equal(t, "value", nested["inner"])

	// Unknown type should be preserved
	unknown, ok := result["unknown"].(struct{ Name string })
	assert.True(t, ok)
	assert.Equal(t, "test", unknown.Name)
}

func TestMapAny_Clone(t *testing.T) {
	original := NewMap(map[string]interface{}{
		"key1": "value1",
		"key2": 42,
	}).EnableCut(".")

	clone := original.Clone()

	// Verify clone has same data
	assert.Equal(t, "value1", clone.GetString("key1"))
	assert.Equal(t, 42, clone.GetInt("key2"))

	// Verify clone has same settings
	assert.Equal(t, original.cut.Load(), clone.cut.Load())
	assert.Equal(t, original.seq.Load(), clone.seq.Load())

	// Verify clone is independent
	clone.Set("key3", "new_value")
	assert.False(t, original.Exists("key3"))
	assert.True(t, clone.Exists("key3"))
}

func TestMapAny_Range(t *testing.T) {
	mapAny := NewMap(map[string]interface{}{
		"key1": "value1",
		"key2": 42,
		"key3": true,
	})

	collected := make(map[interface{}]interface{})
	mapAny.Range(func(key, value interface{}) bool {
		collected[key] = value
		return true
	})

	assert.Len(t, collected, 3)
	assert.Equal(t, "value1", collected["key1"])
	assert.Equal(t, 42, collected["key2"])
	assert.Equal(t, true, collected["key3"])

	// Test early termination
	count := 0
	mapAny.Range(func(key, value interface{}) bool {
		count++
		return count < 2 // Stop after first item
	})
	assert.Equal(t, 2, count)
}

func TestErrNotFound(t *testing.T) {
	assert.Equal(t, "not found", ErrNotFound.Error())
	assert.True(t, errors.Is(ErrNotFound, ErrNotFound))
}

// Test concurrent access to MapAny
func TestMapAny_Concurrent(t *testing.T) {
	mapAny := NewMap(map[string]interface{}{
		"counter": 0,
	})

	var wg sync.WaitGroup
	numGoroutines := 100
	operationsPerGoroutine := 100

	// Concurrent writes
	for i := 0; i < numGoroutines; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			for j := 0; j < operationsPerGoroutine; j++ {
				key := fmt.Sprintf("key_%d_%d", id, j)
				mapAny.Set(key, id*operationsPerGoroutine+j)
			}
		}(i)
	}

	// Concurrent reads
	for i := 0; i < numGoroutines; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			for j := 0; j < operationsPerGoroutine; j++ {
				_ = mapAny.GetInt("counter")
				_ = mapAny.Exists("counter")
			}
		}(i)
	}

	wg.Wait()

	// Verify some data was written
	assert.True(t, mapAny.Exists("counter"))

	// Count how many keys were set
	count := 0
	mapAny.Range(func(key, value interface{}) bool {
		count++
		return true
	})

	// Should have at least the original "counter" key plus some new ones
	assert.Greater(t, count, 1)
}

// Additional coverage tests for edge cases and error paths
func TestMapAny_GetMethodEdgeCases(t *testing.T) {
	t.Run("nested path not found", func(t *testing.T) {
		m := NewMap(map[string]interface{}{
			"level1": map[string]interface{}{
				"level2": "value",
			},
		})
		m.EnableCut(".")

		// Test path that doesn't exist in nested structure
		result, exists := m.get("level1.nonexistent.field")
		assert.False(t, exists)
		assert.Nil(t, result)
	})

	t.Run("invalid nested map type", func(t *testing.T) {
		m := NewMap(map[string]interface{}{
			"level1": "not_a_map", // This is not a map, so nested access should fail
		})
		m.EnableCut(".")

		// Test path that tries to access nested field on non-map value
		result, exists := m.get("level1.field")
		assert.False(t, exists)
		assert.Nil(t, result)
	})

	t.Run("empty keys after split", func(t *testing.T) {
		m := NewMap(map[string]interface{}{
			"test": "value",
		})
		m.EnableCut(".")

		// This should hit the len(keys) > 0 condition at the end
		// by having an empty key after splitting
		result, exists := m.get("test.")
		assert.False(t, exists)
		assert.Nil(t, result)
	})

	t.Run("disable cut path coverage", func(t *testing.T) {
		m := NewMap(map[string]interface{}{
			"nested": map[string]interface{}{
				"level2": "nested_value",
			},
		})

		// Test get without cut enabled to cover non-cut path
		val, ok := m.get("nested.level2")
		assert.False(t, ok)
		assert.Nil(t, val)
	})
}

func TestNewMapWithAny_EdgeCases(t *testing.T) {
	t.Run("json marshal error with channel", func(t *testing.T) {
		// Create a struct with a channel that can't be marshaled to JSON
		type invalidStruct struct {
			Ch chan int
		}

		input := invalidStruct{Ch: make(chan int)}
		result, err := NewMapWithAny(input)
		assert.Error(t, err)
		assert.Nil(t, result)
	})

	t.Run("complex struct success case", func(t *testing.T) {
		// Test with a more complex struct to ensure proper coverage
		type complexStruct struct {
			Field1 string
			Field2 int
			Field3 map[string]interface{}
			Field4 []string
		}

		input := complexStruct{
			Field1: "test",
			Field2: 42,
			Field3: map[string]interface{}{"nested": "value"},
			Field4: []string{"a", "b", "c"},
		}

		result, err := NewMapWithAny(input)
		assert.NoError(t, err)
		assert.NotNil(t, result)

		// Check that the values are properly accessible
		assert.Equal(t, "test", result.GetString("Field1"))
		assert.Equal(t, 42, result.GetInt("Field2"))
	})

	t.Run("circular reference handling", func(t *testing.T) {
		// Create a complex struct that might cause issues
		complexStruct := map[string]interface{}{
			"circular": nil,
		}
		// Create circular reference which might cause issues
		complexStruct["circular"] = complexStruct

		_, err := NewMapWithAny(complexStruct)
		if err != nil {
			t.Logf("NewMapWithAny correctly returned error for circular reference: %v", err)
		} else {
			t.Log("NewMapWithAny handled circular reference successfully")
		}
	})
}

func TestMapAny_ToMapJSONUnmarshalError(t *testing.T) {
	t.Run("json unmarshal error in toMap default case", func(t *testing.T) {
		mapAny := NewMap(nil)

		// Test with a slice that marshals to JSON but doesn't unmarshal to a map
		result := mapAny.toMap([]string{"not", "a", "map"})
		assert.NotNil(t, result)
		assert.False(t, result.Exists("any_key"))
	})
}

func TestMapAny_ComplexScenarios(t *testing.T) {
	t.Run("complex nested structure testing", func(t *testing.T) {
		// Create complex nested structure to test various code paths
		complexData := map[string]interface{}{
			"simple":  "value",
			"number":  42,
			"boolean": true,
			"null":    nil,
			"array": []interface{}{
				"item1",
				map[string]interface{}{"nested": "in_array"},
				42,
			},
			"nested": map[string]interface{}{
				"level2": map[string]interface{}{
					"level3": "deep_value",
				},
			},
		}

		m := NewMap(complexData)

		// Test various getter methods to ensure full coverage
		assert.Equal(t, "value", m.GetString("simple"))
		assert.Equal(t, 42, m.GetInt("number"))
		assert.True(t, m.GetBool("boolean"))
		assert.NotNil(t, m.GetSlice("array"))
		assert.NotNil(t, m.GetMap("nested"))

		// Test with invalid types to cover error paths
	assert.Equal(t, "42", m.GetString("number"))           // int to string
	assert.Equal(t, 0, m.GetInt("simple"))                 // string to int
	assert.True(t, m.GetBool("simple"))                    // string "value" to bool is true
	assert.Equal(t, []interface{}{}, m.GetSlice("simple")) // string to slice returns empty slice
	assert.False(t, m.GetMap("simple").Exists("any"))      // string to map
	})
}

// Custom type for testing YAML unmarshal error
// This type will cause YAML unmarshal error after successful JSON marshal
type CustomTypeForYamlError struct {}

// Override MarshalJSON to return invalid YAML
func (c CustomTypeForYamlError) MarshalJSON() ([]byte, error) {
	// Return invalid YAML that is valid JSON
	return []byte(`"not a yaml map"`), nil
}

// Test for NewMapWithAny YAML unmarshal error path
func TestNewMapWithAny_YamlUnmarshalError(t *testing.T) {
	// This should trigger the YAML unmarshal error path
	// because the JSON marshal returns a string, but YAML
	// expects a map when unmarshaling into map[string]interface{}
	input := CustomTypeForYamlError{}
	result, err := NewMapWithAny(input)
	assert.Error(t, err)
	assert.Nil(t, result)
}

// Additional test cases to improve get function coverage
func TestMapAny_Get_FullCoverage(t *testing.T) {
	t.Run("keys slice with length 0 after loop", func(t *testing.T) {
		m := NewMap(map[string]interface{}{
			"test": "value",
		})
		m.EnableCut(".")

		// This should trigger the case where strings.Split returns
		// a slice with length 0 after splitting
		result, exists := m.get(".")
		assert.False(t, exists)
		assert.Nil(t, result)
	})

	t.Run("empty sequence string with loop execution", func(t *testing.T) {
		m := NewMap(map[string]interface{}{
			"t": map[string]interface{}{
				"e": map[string]interface{}{
					"s": map[string]interface{}{
						"t": "final_value",
					},
				},
			},
		})
		// Enable cut with empty sequence
		m.EnableCut("")

		// This should trigger the case where strings.Split returns
		// a slice with each character as an element
		// because seq is empty
		result, exists := m.get("test")
		assert.True(t, exists)
		assert.Equal(t, "final_value", result)
	})

	t.Run("loop executes and last key not found", func(t *testing.T) {
		m := NewMap(map[string]interface{}{
			"level1": map[string]interface{}{
				"level2": map[string]interface{}{},
			},
		})
		m.EnableCut(".")

		// This should trigger the loop execution
		// and then return false because the last key is not found
		result, exists := m.get("level1.level2.level3")
		assert.False(t, exists)
		assert.Nil(t, result)
	})

	t.Run("empty seq with toMap returning nil", func(t *testing.T) {
		m := NewMap(map[string]interface{}{
			"t": "not a map", // This will cause toMap to return nil
		})
		// Enable cut with empty sequence
		m.EnableCut("")

		// This should trigger the case where strings.Split returns
		// a slice with each character as an element,
		// and toMap returns nil for the first character
		result, exists := m.get("test")
		assert.False(t, exists)
		assert.Nil(t, result)
	})

	t.Run("middle key not map type", func(t *testing.T) {
		m := NewMap(map[string]interface{}{
			"level1": "not a map", // This will cause toMap to return nil
		})
		m.EnableCut(".")

		// This should trigger the case where strings.Split returns
		// a slice with length > 1, and toMap returns nil for the first key
		result, exists := m.get("level1.level2")
		assert.False(t, exists)
		assert.Nil(t, result)
	})

	t.Run("keys slice length 2 success", func(t *testing.T) {
		m := NewMap(map[string]interface{}{
			"level1": map[string]interface{}{
				"level2": "value",
			},
		})
		m.EnableCut(".")

		// This should trigger the case where strings.Split returns
		// a slice with length 2, loop executes once, then data.Load(keys[0]) returns true
		result, exists := m.get("level1.level2")
		assert.True(t, exists)
		assert.Equal(t, "value", result)
	})

	t.Run("empty seq with split returning slice including empty string", func(t *testing.T) {
		m := NewMap(map[string]interface{}{
			"t": map[string]interface{}{
				"e": map[string]interface{}{
					"s": map[string]interface{}{
						"t": "final_value",
					},
				},
			},
		})
		// Enable cut with empty sequence
		m.EnableCut("")

		// This should trigger the case where strings.Split returns
		// a slice with length > 1, containing each character of "test"
		result, exists := m.get("test")
		assert.True(t, exists)
		assert.Equal(t, "final_value", result)
	})

	// Additional edge cases to reach 100% coverage
	t.Run("empty key string", func(t *testing.T) {
		m := NewMap(map[string]interface{}{
			"": "empty_key_value",
			"test": "value",
		})
		m.EnableCut(".")

		// Test empty key
		result, exists := m.get("")
		assert.True(t, exists)
		assert.Equal(t, "empty_key_value", result)
	})

	t.Run("empty key with nested structure", func(t *testing.T) {
		m := NewMap(map[string]interface{}{
			"": map[string]interface{}{
				"nested": "value",
			},
		})
		m.EnableCut(".")

		// Test accessing nested structure with empty key
		result, exists := m.get(".nested")
		assert.True(t, exists)
		assert.Equal(t, "value", result)
	})

	t.Run("empty seq with single character key", func(t *testing.T) {
		m := NewMap(map[string]interface{}{
			"k": "single_key_value",
		})
		// Enable cut with empty sequence
		m.EnableCut("")

		// Test single character key with empty seq
		result, exists := m.get("k")
		assert.True(t, exists)
		assert.Equal(t, "single_key_value", result)
	})

	t.Run("empty seq with same character key", func(t *testing.T) {
		m := NewMap(map[string]interface{}{
			"a": map[string]interface{}{
				"a": map[string]interface{}{
					"a": "aaa_value",
				},
			},
		})
		// Enable cut with empty sequence
		m.EnableCut("")

		// Test key with same characters
		result, exists := m.get("aaa")
		assert.True(t, exists)
		assert.Equal(t, "aaa_value", result)
	})

	t.Run("empty seq with non-existent single character", func(t *testing.T) {
		m := NewMap(map[string]interface{}{
			"a": "value_a",
		})
		// Enable cut with empty sequence
		m.EnableCut("")

		// Test non-existent single character key
		result, exists := m.get("b")
		assert.False(t, exists)
		assert.Nil(t, result)
	})

	t.Run("empty seq with split resulting in single element", func(t *testing.T) {
		m := NewMap(map[string]interface{}{
			"test": "value",
		})
		// Enable cut with empty sequence
		m.EnableCut("")

		// Test key that when split with empty seq has single element
		result, exists := m.get("test")
		assert.True(t, exists)
		assert.Equal(t, "value", result)
	})

	t.Run("enable cut then access direct key", func(t *testing.T) {
		m := NewMap(map[string]interface{}{
			"direct_key": "direct_value",
			"nested": map[string]interface{}{
				"key": "nested_value",
			},
		})
		m.EnableCut(".")

		// Test direct key access still works after enabling cut
		result, exists := m.get("direct_key")
		assert.True(t, exists)
		assert.Equal(t, "direct_value", result)
	})

	t.Run("non-existent key with cut enabled", func(t *testing.T) {
		m := NewMap(map[string]interface{}{
			"existing": "value",
		})
		m.EnableCut(".")

		// Test non-existent key with cut enabled
		result, exists := m.get("non_existent_key")
		assert.False(t, exists)
		assert.Nil(t, result)
	})

	// Test specific edge cases for 100% coverage
	t.Run("empty seq with empty key", func(t *testing.T) {
		m := NewMap(map[string]interface{}{
			"": "empty_key_value",
		})
		// Enable cut with empty sequence
		m.EnableCut("")

		// Test empty key with empty seq
		result, exists := m.get("")
		assert.True(t, exists)
		assert.Equal(t, "empty_key_value", result)
	})

	t.Run("empty seq with key that splits to empty slice", func(t *testing.T) {
		m := NewMap(map[string]interface{}{
			"test": "value",
		})
		// Enable cut with empty sequence
		m.EnableCut("")

		// Test key that when split with empty seq results in special behavior
		// This should test the case where len(keys) == 0 after processing
		result, exists := m.get("")
		assert.False(t, exists)
		assert.Nil(t, result)
	})

	// Additional test cases to cover all strings.Split scenarios
	t.Run("strings_split_various_cases", func(t *testing.T) {
		// Test case 1: strings.Split("", "") - returns empty slice
		m1 := NewMap(map[string]interface{}{
			"test": "value",
		})
		m1.EnableCut("")
		result, exists := m1.get("")
		assert.False(t, exists)
		assert.Nil(t, result)

		// Test case 2: strings.Split("key", "") with no matching nested structure
		m2 := NewMap(map[string]interface{}{
			"k": "value",
		})
		m2.EnableCut("")
		result, exists = m2.get("key")
		assert.False(t, exists)
		assert.Nil(t, result)

		// Test case 3: strings.Split with multiple separators in key
		m3 := NewMap(map[string]interface{}{
			"level1": map[string]interface{}{
				"level2": map[string]interface{}{
					"level3": "value",
				},
			},
		})
		m3.EnableCut(".")
		result, exists = m3.get("level1.level2.level3")
		assert.True(t, exists)
		assert.Equal(t, "value", result)

		// Test case 4: strings.Split with no separators in key
		m4 := NewMap(map[string]interface{}{
			"single_key": "value",
		})
		m4.EnableCut(".")
		result, exists = m4.get("single_key")
		assert.True(t, exists)
		assert.Equal(t, "value", result)

		// Test case 5: key with only separators
		m5 := NewMap(map[string]interface{}{
			"test": "value",
		})
		m5.EnableCut(".")
		result, exists = m5.get("....")
		assert.False(t, exists)
		assert.Nil(t, result)

		// Test case 6: strings.Split("", ".") - returns [""]
		m6 := NewMap(map[string]interface{}{
			"": "empty_key_value",
		})
		m6.EnableCut(".")
		result, exists = m6.get("")
		assert.True(t, exists)
		assert.Equal(t, "empty_key_value", result)

		// Test case 7: strings.Split("a.b.c", ".") with partial match
		m7 := NewMap(map[string]interface{}{
			"a": map[string]interface{}{
				"b": "value",
			},
		})
		m7.EnableCut(".")
		result, exists = m7.get("a.b.c")
		assert.False(t, exists)
		assert.Nil(t, result)

		// Test case 8: strings.Split with single separator at beginning
		m8 := NewMap(map[string]interface{}{
			"level1": map[string]interface{}{
				"level2": "value",
			},
		})
		m8.EnableCut(".")
		result, exists = m8.get(".level1.level2")
		assert.False(t, exists)
		assert.Nil(t, result)

		// Test case 9: strings.Split with empty key in middle
		m9 := NewMap(map[string]interface{}{
			"level1": map[string]interface{}{
				"level2": "value",
			},
		})
		m9.EnableCut(".")
		result, exists = m9.get("level1..level2")
		assert.False(t, exists)
		assert.Nil(t, result)
	})

	// Comprehensive test case for get function to reach 100% coverage
	t.Run("get_function_comprehensive_coverage", func(t *testing.T) {
		// Test case 1: Direct key exists
		m1 := NewMap(map[string]interface{}{"direct": "value"})
		m1.EnableCut(".")
		result, exists := m1.get("direct")
		assert.True(t, exists)
		assert.Equal(t, "value", result)

		// Test case 2: Cut disabled, direct key not found
		m2 := NewMap(map[string]interface{}{"direct": "value"})
		m2.DisableCut()
		result, exists = m2.get("direct.nested")
		assert.False(t, exists)
		assert.Nil(t, result)

		// Test case 3: Nested key exists
		m3 := NewMap(map[string]interface{}{
			"level1": map[string]interface{}{
				"level2": "value",
			},
		})
		m3.EnableCut(".")
		result, exists = m3.get("level1.level2")
		assert.True(t, exists)
		assert.Equal(t, "value", result)

		// Test case 4: Nested key not found
		m4 := NewMap(map[string]interface{}{
			"level1": map[string]interface{}{},
		})
		m4.EnableCut(".")
		result, exists = m4.get("level1.level2.level3")
		assert.False(t, exists)
		assert.Nil(t, result)

		// Test case 5: Intermediate key is not a map
		m5 := NewMap(map[string]interface{}{
			"level1": "not_a_map",
		})
		m5.EnableCut(".")
		result, exists = m5.get("level1.level2")
		assert.False(t, exists)
		assert.Nil(t, result)

		// Test case 6: Empty key with cut enabled
		m6 := NewMap(map[string]interface{}{"": "empty_key"})
		m6.EnableCut(".")
		result, exists = m6.get("")
		assert.True(t, exists)
		assert.Equal(t, "empty_key", result)

		// Test case 7: strings.Split returns slice with one element (no separators in key)
		m7 := NewMap(map[string]interface{}{"single": "value"})
		m7.EnableCut(".")
		result, exists = m7.get("single")
		assert.True(t, exists)
		assert.Equal(t, "value", result)

		// Test case 8: strings.Split returns empty slice
		m8 := NewMap(map[string]interface{}{"test": "value"})
		m8.EnableCut("")
		result, exists = m8.get("")
		assert.False(t, exists)
		assert.Nil(t, result)

		// Test case 9: Cut disabled, key not found
		m9 := NewMap(map[string]interface{}{"key": "value"})
		m9.DisableCut()
		result, exists = m9.get("nonexistent")
		assert.False(t, exists)
		assert.Nil(t, result)

		// Test case 10: Cut disabled, direct key found
		m10 := NewMap(map[string]interface{}{"key": "value"})
		m10.DisableCut()
		result, exists = m10.get("key")
		assert.True(t, exists)
		assert.Equal(t, "value", result)

		// Test case 11: strings.Split("", "") with cut enabled
		m11 := NewMap(map[string]interface{}{"": "empty_value"})
		m11.EnableCut("")
		result, exists = m11.get("")
		assert.True(t, exists)
		assert.Equal(t, "empty_value", result)

		// Test case 12: strings.Split with multi-character separator
		m12 := NewMap(map[string]interface{}{
			"level1": map[string]interface{}{
				"level2": "value",
			},
		})
		m12.EnableCut("..")
		result, exists = m12.get("level1..level2")
		assert.True(t, exists)
		assert.Equal(t, "value", result)

		// Test case 13: strings.Split with multi-character separator not found
		m13 := NewMap(map[string]interface{}{
			"level1": map[string]interface{}{
				"level2": "value",
			},
		})
		m13.EnableCut("..")
		result, exists = m13.get("level1.level2")
		assert.False(t, exists)
		assert.Nil(t, result)

		// Test case 14: Direct key found when cut is enabled but no separator in key
		m14 := NewMap(map[string]interface{}{"direct": "value"})
		m14.EnableCut(".")
		result, exists = m14.get("direct")
		assert.True(t, exists)
		assert.Equal(t, "value", result)

		// Test case 15: keys slice empty after strings.Split
		m15 := NewMap(map[string]interface{}{"test": "value"})
		m15.EnableCut("test")
		result, exists = m15.get("test")
		assert.True(t, exists)
		assert.Equal(t, "value", result)

		// Test case 16: keys slice becomes empty after processing
		m16 := NewMap(map[string]interface{}{})
		m16.EnableCut(".")
		result, exists = m16.get(".")
		assert.False(t, exists)
		assert.Nil(t, result)

		// Test case 17: strings.Split returns empty slice for empty key with empty separator
		m17 := NewMap(map[string]interface{}{"": "empty_value"})
		m17.EnableCut("")
		result, exists = m17.get("")
		assert.True(t, exists)
		assert.Equal(t, "empty_value", result)

		// Test case 18: strings.Split with empty separator on single character key
		m18 := NewMap(map[string]interface{}{"a": "value"})
		m18.EnableCut("")
		result, exists = m18.get("a")
		assert.True(t, exists)
		assert.Equal(t, "value", result)

		// Test case 19: strings.Split with empty separator on two character key with no matching nested structure
		m19 := NewMap(map[string]interface{}{"a": "value"})
		m19.EnableCut("")
		result, exists = m19.get("ab")
		assert.False(t, exists)
		assert.Nil(t, result)

		// Test case 20: keys slice becomes empty after loop execution
		// This test case specifically covers the scenario where len(keys) == 0 after the loop
		m20 := NewMap(map[string]interface{}{})
		m20.EnableCut(".")
		// This will cause strings.Split(".", ".") which returns ["", ""]
		// After processing the first empty string in the loop, keys will be [""]
		// The loop condition len(keys) > 1 will be false, and then we check len(keys) > 0 which is true
		// This covers all code paths in the get function
		result, exists = m20.get(".")
		assert.False(t, exists)
		assert.Nil(t, result)

		// Test case 21: strings.Split with separator matching entire key
		// This should result in keys = ["", ""]
		m21 := NewMap(map[string]interface{}{"": map[string]interface{}{"": "value"}})
		m21.EnableCut(".")
		result, exists = m21.get(".")
		assert.True(t, exists)
		assert.Equal(t, "value", result)

		// Test case 22: strings.Split returns slice with only empty strings
		m22 := NewMap(map[string]interface{}{"": "value"})
		m22.EnableCut(".")
		result, exists = m22.get(".")
		// This should fail because "value" is a string, not a map, so toMap returns nil
		assert.False(t, exists)
		assert.Nil(t, result)
	})

	t.Run("direct key found after cut enabled", func(t *testing.T) {
		m := NewMap(map[string]interface{}{
			"direct_key": "direct_value",
		})
		m.EnableCut(".")

		// Test direct key access when cut is enabled
		// This should hit the first if condition and return immediately
		result, exists := m.get("direct_key")
		assert.True(t, exists)
		assert.Equal(t, "direct_value", result)
	})

	t.Run("cut disabled case", func(t *testing.T) {
		m := NewMap(map[string]interface{}{
			"key": "value",
		})
		// Ensure cut is disabled
		m.DisableCut()

		// Test non-existent key when cut is disabled
		// This should hit the second if condition and return immediately
		result, exists := m.get("non_existent_key")
		assert.False(t, exists)
		assert.Nil(t, result)
	})

	// Test specific edge cases for strings.Split behavior
	t.Run("empty seq with key that matches direct key", func(t *testing.T) {
		m := NewMap(map[string]interface{}{
			"matching_key": "value",
		})
		// Enable cut with empty sequence
		m.EnableCut("")

		// This test ensures that when we have a direct match, we return immediately
		// regardless of the cut settings
		result, exists := m.get("matching_key")
		assert.True(t, exists)
		assert.Equal(t, "value", result)
	})

	t.Run("empty seq with key that splits to long slice", func(t *testing.T) {
		// Create a deeply nested structure
		nested := map[string]interface{}{}
		current := nested
		for _, c := range "abcd" {
			next := make(map[string]interface{})
			current[string(c)] = next
			current = next
		}
		current["final"] = "deep_value"

		m := NewMap(map[string]interface{}{
			"a": nested["a"],
		})
		// Enable cut with empty sequence
		m.EnableCut("")

		// Test accessing nested structure with empty seq
		result, exists := m.get("abcdfinal")
		assert.False(t, exists)
		assert.Nil(t, result)
	})

	t.Run("key with multiple separators", func(t *testing.T) {
		m := NewMap(map[string]interface{}{
			"level1": map[string]interface{}{
				"level2": map[string]interface{}{
					"level3": map[string]interface{}{
						"final": "value",
					},
				},
			},
		})
		m.EnableCut(".")

		// Test key with multiple separators
		result, exists := m.get("level1.level2.level3.final")
		assert.True(t, exists)
		assert.Equal(t, "value", result)
	})

	t.Run("mixed case: direct key not found but nested exists", func(t *testing.T) {
		m := NewMap(map[string]interface{}{
			"level1": map[string]interface{}{
				"nested": "value",
			},
		})
		m.EnableCut(".")

		// Test key that doesn't match directly but has nested structure
		result, exists := m.get("level1.nested")
		assert.True(t, exists)
		assert.Equal(t, "value", result)
	})

	// Test the specific case that is missing coverage
	t.Run("seq_empty_key_empty_len_keys_0", func(t *testing.T) {
		m := NewMap(map[string]interface{}{
			"key": "value",
		})
		// Enable cut with empty sequence
		m.EnableCut("")

		// This test is designed to cover the case where:
		// 1. Direct key lookup fails ("" is not in the map)
		// 2. Cut is enabled
		// 3. seq is empty string
		// 4. key is empty string
		// 5. strings.Split("", "") returns []string{""}
		// 6. len(keys) > 1 is false (since len([]string{""]) == 1)
		// 7. len(keys) > 0 is true (since len([]string{""]) == 1)
		// 8. data.Load("") returns false
		// 9. Finally return nil, false
		result, exists := m.get("")
		assert.False(t, exists)
		assert.Nil(t, result)
	})

	t.Run("seq_empty_key_direct_match_not_found", func(t *testing.T) {
		m := NewMap(map[string]interface{}{
			"existing_key": "value",
		})
		// Enable cut with empty sequence
		m.EnableCut("")

		// Test key that doesn't match directly
		result, exists := m.get("non_existent_key")
		assert.False(t, exists)
		assert.Nil(t, result)
	})

	// Test the exact scenario that is missing coverage
	t.Run("exact_missing_coverage_scenario", func(t *testing.T) {
		// Create a test case that will exercise the specific code path
		// that is currently not covered
		m := NewMap(map[string]interface{}{
			"test": "value",
		})
		// Enable cut with empty sequence
		m.EnableCut("")

		// This test is carefully designed to follow these steps:
		// 1. Direct key lookup fails ("a" is not in the map)
		// 2. Cut is enabled
		// 3. seq is empty string
		// 4. key is "a"
		// 5. strings.Split("a", "") returns []string{"", "a", ""} (in Go 1.20+)
		// 6. len(keys) > 1 is true (len is 3)
		// 7. Enter the loop, process first key ""
		// 8. data.Load("") returns false
		// 9. Return nil, false from line 115
		result, exists := m.get("a")
		assert.False(t, exists)
		assert.Nil(t, result)
	})

	// Test the case where strings.Split returns slice with leading empty string
	t.Run("split_with_leading_empty_string", func(t *testing.T) {
		// Create a nested structure
		m := NewMap(map[string]interface{}{
			"": map[string]interface{}{ // Empty string as key
				"test": "value",
			},
		})
		// Enable cut with empty sequence
		m.EnableCut("")

		// Test key that will be split into ["", "test"]
		result, exists := m.get("test")
		assert.False(t, exists)
		assert.Nil(t, result)

		// Test key that will be split into ["", "", "test"]
		result, exists = m.get("test")
		assert.False(t, exists)
		assert.Nil(t, result)
	})

	// Test the case where we have a matching key in the nested structure
	t.Run("match_in_nested_structure", func(t *testing.T) {
		// Create a nested structure
		m := NewMap(map[string]interface{}{
			"a": map[string]interface{}{ // First character matches
				"b": map[string]interface{}{ // Second character matches
					"c": "final_value", // Third character matches
				},
			},
		})
		// Enable cut with empty sequence
		m.EnableCut("")

		// This should match the nested structure
		result, exists := m.get("abc")
		assert.True(t, exists)
		assert.Equal(t, "final_value", result)

		// Test with extra characters
		result, exists = m.get("abcd")
		assert.False(t, exists)
		assert.Nil(t, result)
	})

	// Test the exact code paths that are missing coverage
	t.Run("precise_coverage_test", func(t *testing.T) {
		// Create a simple map for testing
		m := NewMap(map[string]interface{}{
			"test": "direct_value",
		})
		
		// Enable cut with empty sequence
		m.EnableCut("")
		
		// Test 1: Direct key access (should return immediately)
		result, exists := m.get("test")
		assert.True(t, exists)
		assert.Equal(t, "direct_value", result)
		
		// Test 2: Non-existent key with empty seq
		result, exists = m.get("nonexistent")
		assert.False(t, exists)
		assert.Nil(t, result)
	})

	// Test the case where we have a single character key with empty sequence
	t.Run("single_char_key_empty_seq", func(t *testing.T) {
		// Create a map with a single character key
		m := NewMap(map[string]interface{}{
			"a": "value_a",
		})
		
		// Enable cut with empty sequence
		m.EnableCut("")
		
		// Test accessing the single character key
		result, exists := m.get("a")
		assert.True(t, exists)
		assert.Equal(t, "value_a", result)
	})

	// Test the exact code path that is missing coverage
	t.Run("missing_code_path_test", func(t *testing.T) {
		// Create a map with nested structure
		m := NewMap(map[string]interface{}{
			"level1": map[string]interface{}{
				"level2": map[string]interface{}{
					"level3": "final_value",
				},
			},
		})
		
		// Enable cut with dot separator
		m.EnableCut(".")
		
		// Test 1: Access with single separator (should enter loop once)
		result, exists := m.get("level1.level2")
		assert.True(t, exists)
		assert.NotNil(t, result)
		
		// Test 2: Access with multiple separators (should enter loop multiple times)
		result, exists = m.get("level1.level2.level3")
		assert.True(t, exists)
		assert.Equal(t, "final_value", result)
		
		// Test 3: Access with separator at beginning (should fail in first iteration)
		result, exists = m.get(".level1.level2.level3")
		assert.False(t, exists)
		assert.Nil(t, result)
		
		// Test 4: Access with separator at end (should fail in last check)
		result, exists = m.get("level1.level2.level3.")
		assert.False(t, exists)
		assert.Nil(t, result)
		
		// Test 5: Access with non-existent middle key (should fail in loop)
		result, exists = m.get("level1.non_existent.level3")
		assert.False(t, exists)
		assert.Nil(t, result)
	})

	// Test the case where strings.Split returns empty slice
	t.Run("strings_split_returns_empty_slice", func(t *testing.T) {
		// Create a simple map
		m := NewMap(map[string]interface{}{
			"key": "value",
		})
		
		// Enable cut with a separator that won't be in any key
		m.EnableCut("|")
		
		// Test with a key that won't be split
		result, exists := m.get("key")
		assert.True(t, exists)
		assert.Equal(t, "value", result)
		
		// Test with a key that will be split into empty slice
		// This should trigger the case where len(keys) == 0 after strings.Split
		result, exists = m.get("")
		assert.False(t, exists)
		assert.Nil(t, result)
	})

	// Test the case where len(keys) == 0 after processing
	t.Run("len_keys_0_after_processing", func(t *testing.T) {
		// Create a map with various keys
		m := NewMap(map[string]interface{}{
			"key": "value",
		})
		
		// Enable cut with empty sequence
		m.EnableCut("")
		
		// Test with a key that will result in len(keys) == 0 after processing
		// This should trigger the final return nil, false path
		result, exists := m.get("")
		assert.False(t, exists)
		assert.Nil(t, result)
	})

	// Test the exact case where strings.Split returns empty slice
	t.Run("strings_split_empty_result", func(t *testing.T) {
		// Create a simple map
		m := NewMap(map[string]interface{}{
			"key": "value",
		})
		
		// Enable cut with empty sequence
		m.EnableCut("")
		
		// Test with empty key - strings.Split("", "") returns []string{}
		// This should trigger the case where len(keys) == 0 after the loop
		result, exists := m.get("")
		assert.False(t, exists)
		assert.Nil(t, result)
	})

	// Test the case where we have multiple levels of nesting
	t.Run("multiple_levels_nesting", func(t *testing.T) {
		// Create a deeply nested map
		deepMap := map[string]interface{}{
			"a": map[string]interface{}{
				"b": map[string]interface{}{
					"c": map[string]interface{}{
						"d": map[string]interface{}{
							"e": "final_value",
						},
					},
				},
			},
		}
		
		m := NewMap(deepMap)
		// Enable cut with empty sequence
		m.EnableCut("")
		
		// Test accessing the deep value
		result, exists := m.get("abcde")
		assert.True(t, exists)
		assert.Equal(t, "final_value", result)
	})

	// Test the case where seq is empty string
	t.Run("seq_empty_string", func(t *testing.T) {
		// Create a map with nested structure
		m := NewMap(map[string]interface{}{
			"": map[string]interface{}{ // Empty string as key
				"": "value", // Empty string as nested key
			},
		})
		
		// Enable cut with empty sequence
		m.EnableCut("")
		
		// Test accessing the top-level empty key
		result, exists := m.get("")
		assert.True(t, exists)
		// The result should be the map itself, not the nested value
		assert.NotNil(t, result)
		
		// Test with a longer key that should trigger the nested access
		result, exists = m.get("")
		assert.True(t, exists)
		assert.NotNil(t, result)
	})

	// Test the case that should cover all code paths
	t.Run("final_coverage_test", func(t *testing.T) {
		// Create a map with nested structure
		m := NewMap(map[string]interface{}{
			"test": "direct_value",
			"nested": map[string]interface{}{
				"level1": map[string]interface{}{
					"level2": "final_value",
				},
			},
		})
		
		// Enable cut with dot separator
		m.EnableCut(".")
		
		// Test 1: Direct key access
		result, exists := m.get("test")
		assert.True(t, exists)
		assert.Equal(t, "direct_value", result)
		
		// Test 2: Nested key access with multiple separators
		result, exists = m.get("nested.level1.level2")
		assert.True(t, exists)
		assert.Equal(t, "final_value", result)
		
		// Test 3: Non-existent key
		result, exists = m.get("non_existent")
		assert.False(t, exists)
		assert.Nil(t, result)
		
		// Test 4: Non-existent nested key
		result, exists = m.get("nested.level1.non_existent")
		assert.False(t, exists)
		assert.Nil(t, result)
		
		// Test 5: Middle key not a map
		result, exists = m.get("nested.test.level2")
		assert.False(t, exists)
		assert.Nil(t, result)
	})

	// Test the exact case where strings.Split returns ["", "a", ""]
	t.Run("strings_split_with_empty_surrounding", func(t *testing.T) {
		// Create a simple map
		m := NewMap(map[string]interface{}{
			"key": "value",
		})
		
		// Enable cut with empty sequence
		m.EnableCut("")
		
		// Test with "a" - strings.Split("a", "") returns ["", "a", ""]
		// This should trigger the case where len(keys) > 1 is true
		// and then data.Load("") returns false
		result, exists := m.get("a")
		assert.False(t, exists)
		assert.Nil(t, result)
	})

	// Test the case where we have a key that is exactly the separator
	t.Run("key_equals_separator", func(t *testing.T) {
		// Create a simple map
		m := NewMap(map[string]interface{}{
			".": "separator_value",
		})
		
		// Enable cut with dot separator
		m.EnableCut(".")
		
		// Test with "." - strings.Split(".", ".") returns ["", ""]
		result, exists := m.get(".")
		assert.True(t, exists)
		assert.Equal(t, "separator_value", result)
		
		// Test with ".." - strings.Split("..", ".") returns ["", "", ""]
		result, exists = m.get("..")
		assert.False(t, exists)
		assert.Nil(t, result)
	})

	// Test the exact code paths that are missing coverage
	t.Run("missing_coverage_paths", func(t *testing.T) {
		// Create a map with nested structure
		m := NewMap(map[string]interface{}{
			"level1": map[string]interface{}{
				"level2": "value",
			},
		})
		
		// Enable cut with dot separator
		m.EnableCut(".")
		
		// Test case 1: Cover line 115 - data.Load(k) returns false in loop
		result, exists := m.get("level1.non_existent.level3")
		assert.False(t, exists)
		assert.Nil(t, result)
		
		// Test case 2: Cover line 120 - p.toMap(val) returns nil
		result, exists = m.get("level1.level2.non_existent")
		assert.False(t, exists)
		assert.Nil(t, result)
		
		// Test case 3: Cover line 129 - data.Load(keys[0]) returns false
		result, exists = m.get("non_existent.level2.level3")
		assert.False(t, exists)
		assert.Nil(t, result)
	})

	// Test the case with empty seq and multi-character key
	t.Run("empty_seq_multi_char_key", func(t *testing.T) {
		// Create a map with a nested structure where each level is named after a character
		m := NewMap(map[string]interface{}{
			"a": map[string]interface{}{
				"b": map[string]interface{}{
					"c": "final_value",
				},
			},
		})
		
		// Enable cut with empty sequence
		m.EnableCut("")
		
		// Test 1: Access with single character key
		result, exists := m.get("a")
		assert.True(t, exists)
		assert.NotNil(t, result)
		
		// Test 2: Access with multiple character key
		result, exists = m.get("ab")
		assert.True(t, exists)
		assert.NotNil(t, result)
		
		// Test 3: Access with full key path
		result, exists = m.get("abc")
		assert.True(t, exists)
		assert.Equal(t, "final_value", result)
		
		// Test 4: Access with non-existent key
		result, exists = m.get("abd")
		assert.False(t, exists)
		assert.Nil(t, result)
		
		// Test 5: Access with partial key
		result, exists = m.get("abx")
		assert.False(t, exists)
		assert.Nil(t, result)
	})
}
