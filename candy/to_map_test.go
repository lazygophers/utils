package candy

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestToMap(t *testing.T) {
	t.Run("json byte slice", func(t *testing.T) {
		input := []byte(`{"name":"John","age":30}`)
		result := ToMap(input)
		expected := map[string]interface{}{
			"name": "John",
			"age":  float64(30), // JSON numbers are float64
		}
		assert.Equal(t, expected, result)
	})

	t.Run("json string", func(t *testing.T) {
		input := `{"city":"New York","population":8000000}`
		result := ToMap(input)
		expected := map[string]interface{}{
			"city":       "New York",
			"population": float64(8000000),
		}
		assert.Equal(t, expected, result)
	})

	t.Run("invalid json byte slice", func(t *testing.T) {
		input := []byte(`invalid json`)
		result := ToMap(input)
		expected := map[string]interface{}{} // fallback to ToMapStringAny
		assert.Equal(t, expected, result)
	})

	t.Run("invalid json string", func(t *testing.T) {
		input := "invalid json"
		result := ToMap(input)
		expected := map[string]interface{}{} // fallback to ToMapStringAny
		assert.Equal(t, expected, result)
	})

	t.Run("empty json", func(t *testing.T) {
		input := "{}"
		result := ToMap(input)
		expected := map[string]interface{}{}
		assert.Equal(t, expected, result)
	})

	t.Run("nil input", func(t *testing.T) {
		result := ToMap(nil)
		assert.Nil(t, result)
	})

	t.Run("map input", func(t *testing.T) {
		input := map[int]string{1: "one", 2: "two"}
		result := ToMap(input)
		expected := map[string]interface{}{
			"1": "one",
			"2": "two",
		}
		assert.Equal(t, expected, result)
	})

	t.Run("non-map input", func(t *testing.T) {
		input := 42
		result := ToMap(input)
		expected := map[string]interface{}{}
		assert.Equal(t, expected, result)
	})
}

func TestToMapStringAny(t *testing.T) {
	t.Run("nil input", func(t *testing.T) {
		result := ToMapStringAny(nil)
		assert.Nil(t, result)
	})

	t.Run("int key map", func(t *testing.T) {
		input := map[int]string{1: "one", 2: "two", 3: "three"}
		result := ToMapStringAny(input)
		expected := map[string]interface{}{
			"1": "one",
			"2": "two",
			"3": "three",
		}
		assert.Equal(t, expected, result)
	})

	t.Run("string key map", func(t *testing.T) {
		input := map[string]int{"one": 1, "two": 2}
		result := ToMapStringAny(input)
		expected := map[string]interface{}{
			"one": 1,
			"two": 2,
		}
		assert.Equal(t, expected, result)
	})

	t.Run("mixed type map", func(t *testing.T) {
		input := map[interface{}]interface{}{
			"key1": "value1",
			42:     "value2",
			true:   "value3",
		}
		result := ToMapStringAny(input)
		// 检查结果是否包含所有期望的键值对，不依赖遍历顺序
		assert.Len(t, result, 3)
		assert.Equal(t, "value1", result["key1"])
		assert.Equal(t, "value2", result["42"])
		// 需要检查 ToString(true) 的实际返回值
		// 可能是 "1" 而不是 "true"
		hasTrue := result["true"] == "value3"
		hasOne := result["1"] == "value3"
		assert.True(t, hasTrue || hasOne, "Expected either 'true' or '1' key to map to 'value3'")
	})

	t.Run("empty map", func(t *testing.T) {
		input := map[string]string{}
		result := ToMapStringAny(input)
		expected := map[string]interface{}{}
		assert.Equal(t, expected, result)
	})

	t.Run("non-map input", func(t *testing.T) {
		input := "not a map"
		result := ToMapStringAny(input)
		expected := map[string]interface{}{}
		assert.Equal(t, expected, result)
	})

	t.Run("slice input", func(t *testing.T) {
		input := []int{1, 2, 3}
		result := ToMapStringAny(input)
		expected := map[string]interface{}{}
		assert.Equal(t, expected, result)
	})

	t.Run("struct input", func(t *testing.T) {
		input := struct{ Name string }{"test"}
		result := ToMapStringAny(input)
		expected := map[string]interface{}{}
		assert.Equal(t, expected, result)
	})
}
