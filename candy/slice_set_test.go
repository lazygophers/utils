package candy

import (
	"testing"
)

// TestSpare 测试 Spare 函数
func TestSpare(t *testing.T) {
	t.Run("basic difference", func(t *testing.T) {
		ss := []int{1, 2, 3}
		against := []int{2, 3, 4, 5}
		result := Spare(ss, against)
		expected := []int{4, 5}
		if len(result) != len(expected) {
			t.Errorf("Spare length = %d, want %d", len(result), len(expected))
		}
		for i := range expected {
			if result[i] != expected[i] {
				t.Errorf("Spare[%d] = %d, want %d", i, result[i], expected[i])
			}
		}
	})

	t.Run("empty ss", func(t *testing.T) {
		result := Spare([]int{}, []int{1, 2, 3})
		if len(result) != 3 {
			t.Errorf("Spare with empty ss should return all against elements")
		}
	})

	t.Run("empty against", func(t *testing.T) {
		result := Spare([]int{1, 2, 3}, []int{})
		if len(result) != 0 {
			t.Errorf("Spare with empty against should return empty")
		}
	})

	t.Run("no difference", func(t *testing.T) {
		result := Spare([]int{1, 2, 3}, []int{1, 2, 3})
		if len(result) != 0 {
			t.Errorf("Spare with identical slices should return empty")
		}
	})

	t.Run("all different", func(t *testing.T) {
		result := Spare([]int{1, 2, 3}, []int{4, 5, 6})
		if len(result) != 3 {
			t.Errorf("Spare with completely different slices should return all against")
		}
	})
}

// TestRemove 测试 Remove 函数
// TestRemove 测试 Remove 函数
func TestRemove(t *testing.T) {
	t.Run("basic remove", func(t *testing.T) {
		ss := []int{1, 2, 3, 4, 5}
		toRemove := []int{2, 4, 6}
		result := Remove(ss, toRemove)
		expected := []int{1, 3, 5}
		if len(result) != len(expected) {
			t.Errorf("Remove length mismatch")
		}
		for i := range expected {
			if result[i] != expected[i] {
				t.Errorf("Remove[%d] = %d, want %d", i, result[i], expected[i])
			}
		}
	})

	t.Run("empty source", func(t *testing.T) {
		result := Remove([]int{}, []int{1, 2})
		if len(result) != 0 {
			t.Errorf("Remove from empty should return empty")
		}
	})

	t.Run("empty toRemove", func(t *testing.T) {
		result := Remove([]int{1, 2, 3}, []int{})
		if len(result) != 3 {
			t.Errorf("Remove nothing should return all")
		}
	})

	t.Run("remove all", func(t *testing.T) {
		result := Remove([]int{1, 2, 3}, []int{1, 2, 3})
		if len(result) != 0 {
			t.Errorf("Remove all should return empty")
		}
	})
}

// TestRemoveIndex 测试 RemoveIndex 函数
// TestDiff 测试 Diff 函数
func TestDiff(t *testing.T) {
	t.Run("basic diff", func(t *testing.T) {
		ss := []int{1, 2, 3}
		against := []int{2, 3, 4}
		added, removed := Diff(ss, against)
		if len(added) != 1 || added[0] != 4 {
			t.Errorf("Diff added incorrect")
		}
		if len(removed) != 1 || removed[0] != 1 {
			t.Errorf("Diff removed incorrect")
		}
	})

	t.Run("no diff", func(t *testing.T) {
		added, removed := Diff([]int{1, 2, 3}, []int{1, 2, 3})
		if len(added) != 0 || len(removed) != 0 {
			t.Errorf("Diff should return empty for identical slices")
		}
	})

	t.Run("all different", func(t *testing.T) {
		added, removed := Diff([]int{1, 2}, []int{3, 4})
		if len(added) != 2 || len(removed) != 2 {
			t.Errorf("Diff all different failed")
		}
	})
}

// TestDiffSlice 测试 DiffSlice 函数（反射版本）
// TestIndex 测试 Index 函数
func TestIndex(t *testing.T) {
	t.Run("found at beginning", func(t *testing.T) {
		result := Index([]int{1, 2, 3, 4, 5}, 1)
		if result != 0 {
			t.Errorf("Index = %d, want 0", result)
		}
	})

	t.Run("found at middle", func(t *testing.T) {
		result := Index([]int{1, 2, 3, 4, 5}, 3)
		if result != 2 {
			t.Errorf("Index = %d, want 2", result)
		}
	})

	t.Run("found at end", func(t *testing.T) {
		result := Index([]int{1, 2, 3, 4, 5}, 5)
		if result != 4 {
			t.Errorf("Index = %d, want 4", result)
		}
	})

	t.Run("not found", func(t *testing.T) {
		result := Index([]int{1, 2, 3}, 10)
		if result != -1 {
			t.Errorf("Index not found should return -1")
		}
	})

	t.Run("empty slice", func(t *testing.T) {
		result := Index([]int{}, 1)
		if result != -1 {
			t.Errorf("Index empty should return -1")
		}
	})

	t.Run("string slice", func(t *testing.T) {
		result := Index([]string{"a", "b", "c"}, "b")
		if result != 1 {
			t.Errorf("Index strings = %d, want 1", result)
		}
	})
}

// TestSame 测试 Same 函数
// TestSame 测试 Same 函数
func TestSame(t *testing.T) {
	t.Run("basic intersection", func(t *testing.T) {
		result := Same([]int{1, 2, 3, 4}, []int{2, 3, 5})
		if len(result) != 2 {
			t.Errorf("Same length = %d, want 2", len(result))
		}
	})

	t.Run("no intersection", func(t *testing.T) {
		result := Same([]int{1, 2, 3}, []int{4, 5, 6})
		if len(result) != 0 {
			t.Errorf("Same no intersection should return empty")
		}
	})

	t.Run("all same", func(t *testing.T) {
		result := Same([]int{1, 2, 3}, []int{1, 2, 3})
		if len(result) != 3 {
			t.Errorf("Same all same should return all")
		}
	})

	t.Run("empty slices", func(t *testing.T) {
		result := Same([]int{}, []int{1, 2})
		if len(result) != 0 {
			t.Errorf("Same empty should return empty")
		}
	})
}

// TestSliceEqual 测试 SliceEqual 函数
