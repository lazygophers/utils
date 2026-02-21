package candy

import (
	"testing"
)

// TestRemoveIndex 测试 RemoveIndex 函数
func TestRemoveIndex(t *testing.T) {
	t.Run("remove first", func(t *testing.T) {
		result := RemoveIndex([]int{1, 2, 3, 4}, 0)
		expected := []int{2, 3, 4}
		if len(result) != len(expected) {
			t.Errorf("RemoveIndex first length mismatch")
		}
		for i := range expected {
			if result[i] != expected[i] {
				t.Errorf("RemoveIndex[%d] = %d, want %d", i, result[i], expected[i])
			}
		}
	})

	t.Run("remove last", func(t *testing.T) {
		result := RemoveIndex([]int{1, 2, 3, 4}, 3)
		expected := []int{1, 2, 3}
		if len(result) != len(expected) {
			t.Errorf("RemoveIndex last length mismatch")
		}
	})

	t.Run("remove middle", func(t *testing.T) {
		result := RemoveIndex([]int{1, 2, 3, 4}, 2)
		expected := []int{1, 2, 4}
		if len(result) != len(expected) {
			t.Errorf("RemoveIndex middle length mismatch")
		}
		for i := range expected {
			if result[i] != expected[i] {
				t.Errorf("RemoveIndex middle[%d] = %d, want %d", i, result[i], expected[i])
			}
		}
	})

	t.Run("invalid index negative", func(t *testing.T) {
		result := RemoveIndex([]int{1, 2, 3}, -1)
		if len(result) != 0 {
			t.Errorf("RemoveIndex with negative index should return empty")
		}
	})

	t.Run("invalid index too large", func(t *testing.T) {
		result := RemoveIndex([]int{1, 2, 3}, 10)
		if len(result) != 0 {
			t.Errorf("RemoveIndex with large index should return empty")
		}
	})

	t.Run("empty slice", func(t *testing.T) {
		result := RemoveIndex([]int{}, 0)
		if len(result) != 0 {
			t.Errorf("RemoveIndex from empty should return empty")
		}
	})
}

// TestRemoveSlice 测试 RemoveSlice 函数（反射版本）
// TestRemoveSlice 测试 RemoveSlice 函数（反射版本）
func TestRemoveSlice(t *testing.T) {
	t.Run("basic remove", func(t *testing.T) {
		result := RemoveSlice([]int{1, 2, 3, 4, 5}, []int{2, 4}).([]int)
		if len(result) != 3 {
			t.Errorf("RemoveSlice length = %d, want 3", len(result))
		}
	})

	t.Run("remove strings", func(t *testing.T) {
		result := RemoveSlice([]string{"a", "b", "c"}, []string{"b"}).([]string)
		if len(result) != 2 || result[0] != "a" || result[1] != "c" {
			t.Errorf("RemoveSlice strings failed")
		}
	})

	t.Run("panic on non-slice source", func(t *testing.T) {
		defer func() {
			if r := recover(); r == nil {
				t.Errorf("RemoveSlice should panic on non-slice source")
			}
		}()
		RemoveSlice("not a slice", []int{1})
	})

	t.Run("panic on non-slice remove", func(t *testing.T) {
		defer func() {
			if r := recover(); r == nil {
				t.Errorf("RemoveSlice should panic on non-slice remove")
			}
		}()
		RemoveSlice([]int{1, 2}, "not a slice")
	})

	t.Run("panic on type mismatch", func(t *testing.T) {
		defer func() {
			if r := recover(); r == nil {
				t.Errorf("RemoveSlice should panic on type mismatch")
			}
		}()
		RemoveSlice([]int{1, 2}, []string{"a"})
	})
}

// TestDiff 测试 Diff 函数
// TestDiffSlice 测试 DiffSlice 函数（反射版本）
func TestDiffSlice(t *testing.T) {
	t.Run("basic diff", func(t *testing.T) {
		removed, added := DiffSlice([]int{1, 2, 3}, []int{2, 3, 4})
		removedSlice := removed.([]int)
		addedSlice := added.([]int)
		if len(removedSlice) != 1 || removedSlice[0] != 1 {
			t.Errorf("DiffSlice removed incorrect")
		}
		if len(addedSlice) != 1 || addedSlice[0] != 4 {
			t.Errorf("DiffSlice added incorrect")
		}
	})

	t.Run("panic on non-slice a", func(t *testing.T) {
		defer func() {
			if r := recover(); r == nil {
				t.Errorf("DiffSlice should panic on non-slice a")
			}
		}()
		DiffSlice("not a slice", []int{1})
	})

	t.Run("panic on non-slice b", func(t *testing.T) {
		defer func() {
			if r := recover(); r == nil {
				t.Errorf("DiffSlice should panic on non-slice b")
			}
		}()
		DiffSlice([]int{1}, "not a slice")
	})

	t.Run("panic on type mismatch", func(t *testing.T) {
		defer func() {
			if r := recover(); r == nil {
				t.Errorf("DiffSlice should panic on type mismatch")
			}
		}()
		DiffSlice([]int{1, 2}, []string{"a"})
	})
}

// TestFirst 测试 First 函数
