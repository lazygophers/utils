package candy

import (
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMapValues(t *testing.T) {
	t.Run("extract values from map", func(t *testing.T) {
		m := map[string]int{"a": 1, "b": 2, "c": 3}
		result := MapValues(m)
		sort.Ints(result)
		expected := []int{1, 2, 3}
		assert.Equal(t, expected, result)
	})
}

func TestMapValuesGeneric(t *testing.T) {
	t.Run("extract string values", func(t *testing.T) {
		m := map[int]string{1: "one", 2: "two", 3: "three"}
		result := MapValuesGeneric(m)
		assert.Len(t, result, 3)
		assert.Contains(t, result, "one")
		assert.Contains(t, result, "two")
	})
}

func TestMapValuesAny(t *testing.T) {
	t.Run("extract any values", func(t *testing.T) {
		m := map[string]interface{}{"a": 1, "b": "two", "c": 3.0}
		result := MapValuesAny(m)
		assert.Len(t, result, 3)
	})
}

func TestMapValuesString(t *testing.T) {
	t.Run("extract string values", func(t *testing.T) {
		m := map[int]string{1: "one", 2: "two"}
		result := MapValuesString(m)
		assert.Len(t, result, 2)
		assert.Contains(t, result, "one")
	})
}

func TestMapValuesInt(t *testing.T) {
	t.Run("extract int values", func(t *testing.T) {
		m := map[string]int{"a": 1, "b": 2}
		result := MapValuesInt(m)
		sort.Ints(result)
		expected := []int{1, 2}
		assert.Equal(t, expected, result)
	})
}

func TestMapValuesFloat64(t *testing.T) {
	t.Run("extract float64 values", func(t *testing.T) {
		m := map[string]float64{"a": 1.5, "b": 2.5}
		result := MapValuesFloat64(m)
		assert.Len(t, result, 2)
		assert.Contains(t, result, 1.5)
	})
}
