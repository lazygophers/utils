package candy

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTypedSliceCopy(t *testing.T) {
	t.Run("copy int slice", func(t *testing.T) {
		original := []int{1, 2, 3, 4, 5}
		copied := TypedSliceCopy(original)
		assert.Equal(t, original, copied)
		// Verify independence
		copied[0] = 99
		assert.Equal(t, 1, original[0])
	})

	t.Run("copy string slice", func(t *testing.T) {
		original := []string{"a", "b", "c"}
		copied := TypedSliceCopy(original)
		assert.Equal(t, original, copied)
		copied[0] = "z"
		assert.Equal(t, "a", original[0])
	})

	t.Run("empty slice", func(t *testing.T) {
		original := []int{}
		copied := TypedSliceCopy(original)
		assert.Empty(t, copied)
	})
}

func TestTypedMapCopy(t *testing.T) {
	t.Run("copy string-int map", func(t *testing.T) {
		original := map[string]int{"a": 1, "b": 2, "c": 3}
		copied := TypedMapCopy(original)
		assert.Equal(t, original, copied)
		// Verify independence
		copied["d"] = 4
		assert.NotContains(t, original, "d")
	})

	t.Run("copy int-string map", func(t *testing.T) {
		original := map[int]string{1: "one", 2: "two"}
		copied := TypedMapCopy(original)
		assert.Equal(t, original, copied)
	})

	t.Run("empty map", func(t *testing.T) {
		original := map[string]int{}
		copied := TypedMapCopy(original)
		assert.Empty(t, copied)
	})
}

func TestGenericSliceEqual(t *testing.T) {
	t.Run("equal int slices", func(t *testing.T) {
		s1 := []int{1, 2, 3}
		s2 := []int{1, 2, 3}
		assert.True(t, GenericSliceEqual(s1, s2))
	})

	t.Run("unequal int slices", func(t *testing.T) {
		s1 := []int{1, 2, 3}
		s2 := []int{1, 2, 4}
		assert.False(t, GenericSliceEqual(s1, s2))
	})

	t.Run("different length slices", func(t *testing.T) {
		s1 := []int{1, 2}
		s2 := []int{1, 2, 3}
		assert.False(t, GenericSliceEqual(s1, s2))
	})

	t.Run("string slices", func(t *testing.T) {
		s1 := []string{"a", "b"}
		s2 := []string{"a", "b"}
		s3 := []string{"a", "c"}
		assert.True(t, GenericSliceEqual(s1, s2))
		assert.False(t, GenericSliceEqual(s1, s3))
	})
}

func TestMapEqual(t *testing.T) {
	t.Run("equal maps", func(t *testing.T) {
		m1 := map[string]int{"a": 1, "b": 2}
		m2 := map[string]int{"a": 1, "b": 2}
		assert.True(t, MapEqual(m1, m2))
	})

	t.Run("unequal maps", func(t *testing.T) {
		m1 := map[string]int{"a": 1, "b": 2}
		m2 := map[string]int{"a": 1, "b": 3}
		assert.False(t, MapEqual(m1, m2))
	})

	t.Run("different size maps", func(t *testing.T) {
		m1 := map[string]int{"a": 1}
		m2 := map[string]int{"a": 1, "b": 2}
		assert.False(t, MapEqual(m1, m2))
	})
}

func TestPointerEqual(t *testing.T) {
	t.Run("equal pointers", func(t *testing.T) {
		x := 42
		p1 := &x
		p2 := &x
		assert.True(t, PointerEqual(p1, p2))
	})

	t.Run("different pointers same value", func(t *testing.T) {
		x, y := 42, 42
		p1 := &x
		p2 := &y
		assert.True(t, PointerEqual(p1, p2))
	})

	t.Run("different values", func(t *testing.T) {
		x, y := 42, 43
		p1 := &x
		p2 := &y
		assert.False(t, PointerEqual(p1, p2))
	})
}

func TestStructEqual(t *testing.T) {
	type Person struct {
		Name string
		Age  int
	}

	comparer := func(a, b Person) bool {
		return a.Name == b.Name && a.Age == b.Age
	}

	t.Run("equal structs", func(t *testing.T) {
		p1 := Person{Name: "Alice", Age: 30}
		p2 := Person{Name: "Alice", Age: 30}
		assert.True(t, StructEqual(p1, p2, comparer))
	})

	t.Run("unequal structs", func(t *testing.T) {
		p1 := Person{Name: "Alice", Age: 30}
		p2 := Person{Name: "Bob", Age: 25}
		assert.False(t, StructEqual(p1, p2, comparer))
	})
}

func TestCloneFunctions(t *testing.T) {
	t.Run("Clone int", func(t *testing.T) {
		original := 42
		cloned := Clone(original)
		assert.Equal(t, original, cloned)
	})

	t.Run("CloneSlice", func(t *testing.T) {
		original := []int{1, 2, 3}
		cloned := CloneSlice(original)
		assert.Equal(t, original, cloned)
		cloned[0] = 99
		assert.NotEqual(t, original[0], cloned[0])
	})

	t.Run("CloneMap", func(t *testing.T) {
		original := map[string]int{"a": 1, "b": 2}
		cloned := CloneMap(original)
		assert.Equal(t, original, cloned)
		cloned["c"] = 3
		assert.NotContains(t, original, "c")
	})
}

func TestEqualFunctions(t *testing.T) {
	t.Run("Equal", func(t *testing.T) {
		assert.True(t, Equal(42, 42))
		assert.False(t, Equal(42, 43))
	})

	t.Run("EqualSlice", func(t *testing.T) {
		assert.True(t, EqualSlice([]int{1, 2, 3}, []int{1, 2, 3}))
		assert.False(t, EqualSlice([]int{1, 2, 3}, []int{1, 2, 4}))
	})

	t.Run("EqualMap", func(t *testing.T) {
		m1 := map[string]int{"a": 1, "b": 2}
		m2 := map[string]int{"a": 1, "b": 2}
		assert.True(t, EqualMap(m1, m2))
		m3 := map[string]int{"a": 1, "b": 3}
		assert.False(t, EqualMap(m1, m3))
	})
}
