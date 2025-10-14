package candy

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMergeMap(t *testing.T) {
	t.Run("merge two maps", func(t *testing.T) {
		m1 := map[string]int{"a": 1, "b": 2}
		m2 := map[string]int{"c": 3, "d": 4}
		result := MergeMap(m1, m2)
		expected := map[string]int{"a": 1, "b": 2, "c": 3, "d": 4}
		assert.Equal(t, expected, result)
	})

	t.Run("overlapping keys", func(t *testing.T) {
		m1 := map[string]int{"a": 1, "b": 2}
		m2 := map[string]int{"b": 3, "c": 4}
		result := MergeMap(m1, m2)
		assert.Equal(t, 3, result["b"]) // m2 value overwrites m1
		assert.Equal(t, 1, result["a"])
		assert.Equal(t, 4, result["c"])
	})

	t.Run("empty maps", func(t *testing.T) {
		m1 := map[string]int{}
		m2 := map[string]int{}
		result := MergeMap(m1, m2)
		assert.Empty(t, result)
	})
}

func TestMergeMapGeneric(t *testing.T) {
	t.Run("merge two generic maps", func(t *testing.T) {
		m1 := map[string]string{"a": "one", "b": "two"}
		m2 := map[string]string{"c": "three", "d": "four"}
		result := MergeMapGeneric(m1, m2)
		assert.Len(t, result, 4)
		assert.Equal(t, "one", result["a"])
		assert.Equal(t, "three", result["c"])
	})

	t.Run("overlapping keys", func(t *testing.T) {
		m1 := map[int]string{1: "one", 2: "two"}
		m2 := map[int]string{2: "TWO", 3: "three"}
		result := MergeMapGeneric(m1, m2)
		assert.Equal(t, "TWO", result[2])
		assert.Equal(t, "one", result[1])
	})
}

func TestCloneMapShallow(t *testing.T) {
	t.Run("clone map", func(t *testing.T) {
		original := map[string]int{"a": 1, "b": 2, "c": 3}
		cloned := CloneMapShallow(original)
		assert.Equal(t, original, cloned)
		
		// Verify it's a shallow copy
		cloned["d"] = 4
		assert.NotContains(t, original, "d")
	})

	t.Run("empty map", func(t *testing.T) {
		original := map[string]int{}
		cloned := CloneMapShallow(original)
		assert.Empty(t, cloned)
	})
}
