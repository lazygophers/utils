package candy

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestToMapInt32String(t *testing.T) {
	t.Run("basic conversion", func(t *testing.T) {
		input := map[int32]string{1: "one", 2: "two", 3: "three"}
		result := ToMapInt32String(input)
		expected := map[int32]string{1: "one", 2: "two", 3: "three"}
		assert.Equal(t, expected, result)
	})

	t.Run("empty map", func(t *testing.T) {
		input := map[int32]string{}
		result := ToMapInt32String(input)
		expected := map[int32]string{}
		assert.Equal(t, expected, result)
	})

	t.Run("non-map input", func(t *testing.T) {
		input := "not a map"
		result := ToMapInt32String(input)
		expected := map[int32]string{}
		assert.Equal(t, expected, result)
	})

	t.Run("nil input", func(t *testing.T) {
		result := ToMapInt32String(nil)
		expected := map[int32]string{}
		assert.Equal(t, expected, result)
	})
}

func TestToMapInt64String(t *testing.T) {
	t.Run("basic conversion", func(t *testing.T) {
		input := map[int64]string{100: "hundred", 200: "two hundred"}
		result := ToMapInt64String(input)
		expected := map[int64]string{100: "hundred", 200: "two hundred"}
		assert.Equal(t, expected, result)
	})

	t.Run("empty map", func(t *testing.T) {
		input := map[int64]string{}
		result := ToMapInt64String(input)
		expected := map[int64]string{}
		assert.Equal(t, expected, result)
	})

	t.Run("non-map input", func(t *testing.T) {
		input := 42
		result := ToMapInt64String(input)
		expected := map[int64]string{}
		assert.Equal(t, expected, result)
	})
}

func TestToMapStringString(t *testing.T) {
	t.Run("basic conversion", func(t *testing.T) {
		input := map[string]string{"key1": "value1", "key2": "value2"}
		result := ToMapStringString(input)
		expected := map[string]string{"key1": "value1", "key2": "value2"}
		assert.Equal(t, expected, result)
	})

	t.Run("empty map", func(t *testing.T) {
		input := map[string]string{}
		result := ToMapStringString(input)
		expected := map[string]string{}
		assert.Equal(t, expected, result)
	})

	t.Run("non-map input", func(t *testing.T) {
		input := []string{"not", "a", "map"}
		result := ToMapStringString(input)
		expected := map[string]string{}
		assert.Equal(t, expected, result)
	})
}

func TestToMapStringInt64(t *testing.T) {
	t.Run("basic conversion", func(t *testing.T) {
		input := map[string]int64{"count": 100, "total": 500}
		result := ToMapStringInt64(input)
		expected := map[string]int64{"count": 100, "total": 500}
		assert.Equal(t, expected, result)
	})

	t.Run("empty map", func(t *testing.T) {
		input := map[string]int64{}
		result := ToMapStringInt64(input)
		expected := map[string]int64{}
		assert.Equal(t, expected, result)
	})

	t.Run("non-map input", func(t *testing.T) {
		input := struct{}{}
		result := ToMapStringInt64(input)
		expected := map[string]int64{}
		assert.Equal(t, expected, result)
	})
}

func TestToMapStringArrayString(t *testing.T) {
	t.Run("basic conversion", func(t *testing.T) {
		input := map[string][]string{
			"colors": {"red", "blue", "green"},
			"fruits": {"apple", "banana"},
		}
		result := ToMapStringArrayString(input)
		expected := map[string][]string{
			"colors": {"red", "blue", "green"},
			"fruits": {"apple", "banana"},
		}
		assert.Equal(t, expected, result)
	})

	t.Run("empty map", func(t *testing.T) {
		input := map[string][]string{}
		result := ToMapStringArrayString(input)
		expected := map[string][]string{}
		assert.Equal(t, expected, result)
	})

	t.Run("non-map input should panic", func(t *testing.T) {
		input := []string{"not", "a", "map"}
		assert.Panics(t, func() {
			ToMapStringArrayString(input)
		})
	})
}