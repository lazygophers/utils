package candy

import (
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMapKeysString(t *testing.T) {
	t.Run("basic string keys", func(t *testing.T) {
		input := map[string]int{"a": 1, "b": 2, "c": 3}
		result := MapKeysString(input)
		sort.Strings(result)
		expected := []string{"a", "b", "c"}
		assert.Equal(t, expected, result)
	})

	t.Run("empty map", func(t *testing.T) {
		input := map[string]int{}
		result := MapKeysString(input)
		assert.Empty(t, result)
	})

	t.Run("single key", func(t *testing.T) {
		input := map[string]string{"key": "value"}
		result := MapKeysString(input)
		assert.Equal(t, []string{"key"}, result)
	})
}

func TestMapKeysInt(t *testing.T) {
	t.Run("basic int keys", func(t *testing.T) {
		input := map[int]string{1: "one", 2: "two", 3: "three"}
		result := MapKeysInt(input)
		sort.Ints(result)
		expected := []int{1, 2, 3}
		assert.Equal(t, expected, result)
	})

	t.Run("empty map", func(t *testing.T) {
		input := map[int]string{}
		result := MapKeysInt(input)
		assert.Empty(t, result)
	})
}

func TestMapKeysInt64(t *testing.T) {
	t.Run("basic int64 keys", func(t *testing.T) {
		input := map[int64]string{1: "one", 2: "two", 3: "three"}
		result := MapKeysInt64(input)
		sort.Slice(result, func(i, j int) bool { return result[i] < result[j] })
		expected := []int64{1, 2, 3}
		assert.Equal(t, expected, result)
	})
}

func TestMapKeysUint(t *testing.T) {
	t.Run("basic uint keys", func(t *testing.T) {
		input := map[uint]string{1: "one", 2: "two", 3: "three"}
		result := MapKeysUint(input)
		sort.Slice(result, func(i, j int) bool { return result[i] < result[j] })
		expected := []uint{1, 2, 3}
		assert.Equal(t, expected, result)
	})
}

func TestMapKeysUint8(t *testing.T) {
	t.Run("basic uint8 keys", func(t *testing.T) {
		input := map[uint8]string{1: "one", 2: "two", 3: "three"}
		result := MapKeysUint8(input)
		sort.Slice(result, func(i, j int) bool { return result[i] < result[j] })
		expected := []uint8{1, 2, 3}
		assert.Equal(t, expected, result)
	})
}

func TestMapKeysUint16(t *testing.T) {
	t.Run("basic uint16 keys", func(t *testing.T) {
		input := map[uint16]string{1: "one", 2: "two"}
		result := MapKeysUint16(input)
		assert.Len(t, result, 2)
	})
}

func TestMapKeysUint32(t *testing.T) {
	t.Run("basic uint32 keys", func(t *testing.T) {
		input := map[uint32]string{1: "one", 2: "two"}
		result := MapKeysUint32(input)
		assert.Len(t, result, 2)
	})
}

func TestMapKeysUint64(t *testing.T) {
	t.Run("basic uint64 keys", func(t *testing.T) {
		input := map[uint64]string{1: "one", 2: "two"}
		result := MapKeysUint64(input)
		assert.Len(t, result, 2)
	})
}

func TestMapKeysFloat32(t *testing.T) {
	t.Run("basic float32 keys", func(t *testing.T) {
		input := map[float32]string{1.5: "one", 2.5: "two"}
		result := MapKeysFloat32(input)
		assert.Len(t, result, 2)
	})
}

func TestMapKeysFloat64(t *testing.T) {
	t.Run("basic float64 keys", func(t *testing.T) {
		input := map[float64]string{1.5: "one", 2.5: "two"}
		result := MapKeysFloat64(input)
		assert.Len(t, result, 2)
	})
}

func TestMapKeysInterface(t *testing.T) {
	t.Run("basic interface keys", func(t *testing.T) {
		input := map[interface{}]string{"a": "one", 1: "two"}
		result := MapKeysInterface(input)
		assert.Len(t, result, 2)
	})
}

func TestMapKeysAny(t *testing.T) {
	t.Run("basic any keys", func(t *testing.T) {
		input := map[any]string{"a": "one", 1: "two"}
		result := MapKeysAny(input)
		assert.Len(t, result, 2)
	})
}
