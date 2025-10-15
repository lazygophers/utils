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
func TestFirst(t *testing.T) {
	t.Run("normal slice", func(t *testing.T) {
		result := First([]int{1, 2, 3})
		if result != 1 {
			t.Errorf("First = %d, want 1", result)
		}
	})

	t.Run("empty slice", func(t *testing.T) {
		result := First([]int{})
		if result != 0 {
			t.Errorf("First empty should return zero value")
		}
	})

	t.Run("single element", func(t *testing.T) {
		result := First([]string{"hello"})
		if result != "hello" {
			t.Errorf("First single = %s, want hello", result)
		}
	})
}

// TestFirstOr 测试 FirstOr 函数
func TestFirstOr(t *testing.T) {
	t.Run("normal slice", func(t *testing.T) {
		result := FirstOr([]int{1, 2, 3}, 999)
		if result != 1 {
			t.Errorf("FirstOr = %d, want 1", result)
		}
	})

	t.Run("empty slice", func(t *testing.T) {
		result := FirstOr([]int{}, 999)
		if result != 999 {
			t.Errorf("FirstOr empty should return default")
		}
	})

	t.Run("empty string slice", func(t *testing.T) {
		result := FirstOr([]string{}, "default")
		if result != "default" {
			t.Errorf("FirstOr empty strings should return default")
		}
	})
}

// TestLast 测试 Last 函数
func TestLast(t *testing.T) {
	t.Run("normal slice", func(t *testing.T) {
		result := Last([]int{1, 2, 3, 4, 5})
		if result != 5 {
			t.Errorf("Last = %d, want 5", result)
		}
	})

	t.Run("empty slice", func(t *testing.T) {
		result := Last([]string{})
		if result != "" {
			t.Errorf("Last empty should return zero value")
		}
	})

	t.Run("single element", func(t *testing.T) {
		result := Last([]int{42})
		if result != 42 {
			t.Errorf("Last single = %d, want 42", result)
		}
	})
}

// TestLastOr 测试 LastOr 函数
func TestLastOr(t *testing.T) {
	t.Run("normal slice", func(t *testing.T) {
		result := LastOr([]int{1, 2, 3, 4, 5}, 999)
		if result != 5 {
			t.Errorf("LastOr = %d, want 5", result)
		}
	})

	t.Run("empty slice", func(t *testing.T) {
		result := LastOr([]int{}, 999)
		if result != 999 {
			t.Errorf("LastOr empty should return default")
		}
	})

	t.Run("empty string slice", func(t *testing.T) {
		result := LastOr([]string{}, "default")
		if result != "default" {
			t.Errorf("LastOr empty strings should return default")
		}
	})
}

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
func TestSliceEqual(t *testing.T) {
	t.Run("equal slices", func(t *testing.T) {
		if !SliceEqual([]int{1, 2, 3}, []int{1, 2, 3}) {
			t.Errorf("SliceEqual should return true for equal slices")
		}
	})

	t.Run("different order same elements", func(t *testing.T) {
		if !SliceEqual([]int{1, 2, 3}, []int{3, 2, 1}) {
			t.Errorf("SliceEqual should return true regardless of order")
		}
	})

	t.Run("different length", func(t *testing.T) {
		if SliceEqual([]int{1, 2}, []int{1, 2, 3}) {
			t.Errorf("SliceEqual should return false for different lengths")
		}
	})

	t.Run("different elements", func(t *testing.T) {
		if SliceEqual([]int{1, 2, 3}, []int{1, 2, 4}) {
			t.Errorf("SliceEqual should return false for different elements")
		}
	})

	t.Run("both nil", func(t *testing.T) {
		var a, b []int
		if !SliceEqual(a, b) {
			t.Errorf("SliceEqual should return true for both nil")
		}
	})

	t.Run("one nil one empty", func(t *testing.T) {
		var a []int
		b := []int{}
		if SliceEqual(a, b) {
			t.Errorf("SliceEqual should return false for nil vs empty")
		}
	})

	t.Run("duplicate elements", func(t *testing.T) {
		if !SliceEqual([]int{1, 1, 2, 2}, []int{2, 1, 2, 1}) {
			t.Errorf("SliceEqual should handle duplicates correctly")
		}
	})

	t.Run("different duplicate counts", func(t *testing.T) {
		if SliceEqual([]int{1, 1, 2}, []int{1, 2, 2}) {
			t.Errorf("SliceEqual should return false for different duplicate counts")
		}
	})

	t.Run("empty slices", func(t *testing.T) {
		if !SliceEqual([]int{}, []int{}) {
			t.Errorf("SliceEqual should return true for empty slices")
		}
	})

	t.Run("one nil other not", func(t *testing.T) {
		var a []int
		b := []int{1, 2, 3}
		if SliceEqual(a, b) {
			t.Errorf("SliceEqual should return false when one is nil")
		}
		if SliceEqual(b, a) {
			t.Errorf("SliceEqual should return false when one is nil (reversed)")
		}
	})
}

// TestDrop 测试 Drop 函数
func TestDrop(t *testing.T) {
	t.Run("drop first 2", func(t *testing.T) {
		result := Drop([]int{1, 2, 3, 4, 5}, 2)
		expected := []int{3, 4, 5}
		if len(result) != len(expected) {
			t.Errorf("Drop length mismatch")
		}
		for i := range expected {
			if result[i] != expected[i] {
				t.Errorf("Drop[%d] = %d, want %d", i, result[i], expected[i])
			}
		}
	})

	t.Run("drop zero", func(t *testing.T) {
		result := Drop([]string{"a", "b", "c"}, 0)
		if len(result) != 3 {
			t.Errorf("Drop 0 should return all")
		}
	})

	t.Run("drop negative", func(t *testing.T) {
		result := Drop([]int{1, 2, 3}, -1)
		if len(result) != 3 {
			t.Errorf("Drop negative should return all")
		}
	})

	t.Run("drop more than length", func(t *testing.T) {
		result := Drop([]int{1, 2, 3}, 5)
		if len(result) != 0 {
			t.Errorf("Drop more than length should return empty")
		}
	})

	t.Run("drop all", func(t *testing.T) {
		result := Drop([]int{1, 2, 3}, 3)
		if len(result) != 0 {
			t.Errorf("Drop all should return empty")
		}
	})
}

// TestFilter 测试 Filter 函数
func TestFilter(t *testing.T) {
	t.Run("filter even numbers", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5, 6}
		result := Filter(numbers, func(n int) bool {
			return n%2 == 0
		})
		expected := []int{2, 4, 6}
		if len(result) != len(expected) {
			t.Errorf("Filter length = %d, want %d", len(result), len(expected))
		}
		for i := range expected {
			if result[i] != expected[i] {
				t.Errorf("Filter[%d] = %d, want %d", i, result[i], expected[i])
			}
		}
	})

	t.Run("filter none match", func(t *testing.T) {
		result := Filter([]int{1, 3, 5}, func(n int) bool {
			return n%2 == 0
		})
		if len(result) != 0 {
			t.Errorf("Filter no match should return empty")
		}
	})

	t.Run("filter all match", func(t *testing.T) {
		result := Filter([]int{2, 4, 6}, func(n int) bool {
			return n%2 == 0
		})
		if len(result) != 3 {
			t.Errorf("Filter all match should return all")
		}
	})

	t.Run("filter empty slice", func(t *testing.T) {
		result := Filter([]int{}, func(n int) bool {
			return n > 0
		})
		if len(result) != 0 {
			t.Errorf("Filter empty should return empty")
		}
	})
}

// TestFilterNot 测试 FilterNot 函数
func TestFilterNot(t *testing.T) {
	t.Run("filter not even", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5, 6}
		result := FilterNot(numbers, func(n int) bool {
			return n%2 == 0
		})
		expected := []int{1, 3, 5}
		if len(result) != len(expected) {
			t.Errorf("FilterNot length mismatch")
		}
		for i := range expected {
			if result[i] != expected[i] {
				t.Errorf("FilterNot[%d] = %d, want %d", i, result[i], expected[i])
			}
		}
	})

	t.Run("filter not empty strings", func(t *testing.T) {
		strings := []string{"hello", "", "world", ""}
		result := FilterNot(strings, func(s string) bool {
			return s == ""
		})
		if len(result) != 2 {
			t.Errorf("FilterNot empty strings length mismatch")
		}
	})

	t.Run("filter not none match", func(t *testing.T) {
		result := FilterNot([]int{2, 4, 6}, func(n int) bool {
			return n%2 == 0
		})
		if len(result) != 0 {
			t.Errorf("FilterNot all excluded should return empty")
		}
	})

	t.Run("filter not all match", func(t *testing.T) {
		result := FilterNot([]int{1, 3, 5}, func(n int) bool {
			return n%2 == 0
		})
		if len(result) != 3 {
			t.Errorf("FilterNot no exclusion should return all")
		}
	})
}

// TestContains 测试 Contains 函数
func TestContains(t *testing.T) {
	t.Run("contains element", func(t *testing.T) {
		if !Contains([]int{1, 2, 3, 4, 5}, 3) {
			t.Errorf("Contains should return true")
		}
	})

	t.Run("does not contain", func(t *testing.T) {
		if Contains([]int{1, 2, 3}, 10) {
			t.Errorf("Contains should return false")
		}
	})

	t.Run("contains string", func(t *testing.T) {
		if !Contains([]string{"a", "b", "c"}, "b") {
			t.Errorf("Contains string should return true")
		}
	})

	t.Run("empty slice", func(t *testing.T) {
		if Contains([]int{}, 1) {
			t.Errorf("Contains empty should return false")
		}
	})
}

// TestContainsUsing 测试 ContainsUsing 函数
func TestContainsUsing(t *testing.T) {
	t.Run("contains using predicate", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5}
		found := ContainsUsing(numbers, func(v int) bool {
			return v > 3
		})
		if !found {
			t.Errorf("ContainsUsing should return true")
		}
	})

	t.Run("does not contain using predicate", func(t *testing.T) {
		numbers := []int{1, 2, 3}
		found := ContainsUsing(numbers, func(v int) bool {
			return v > 10
		})
		if found {
			t.Errorf("ContainsUsing should return false")
		}
	})

	t.Run("contains using custom struct", func(t *testing.T) {
		type Person struct {
			Name string
			Age  int
		}
		persons := []Person{
			{Name: "Alice", Age: 30},
			{Name: "Bob", Age: 25},
		}
		found := ContainsUsing(persons, func(p Person) bool {
			return p.Age > 28
		})
		if !found {
			t.Errorf("ContainsUsing custom struct should return true")
		}
	})

	t.Run("empty slice", func(t *testing.T) {
		found := ContainsUsing([]int{}, func(v int) bool {
			return v > 0
		})
		if found {
			t.Errorf("ContainsUsing empty should return false")
		}
	})
}

// TestBottom 测试 Bottom 函数
func TestBottom(t *testing.T) {
	t.Run("bottom 2 elements", func(t *testing.T) {
		result := Bottom([]int{1, 2, 3, 4, 5}, 2)
		expected := []int{4, 5}
		if len(result) != len(expected) {
			t.Errorf("Bottom length mismatch")
		}
		for i := range expected {
			if result[i] != expected[i] {
				t.Errorf("Bottom[%d] = %d, want %d", i, result[i], expected[i])
			}
		}
	})

	t.Run("bottom zero", func(t *testing.T) {
		result := Bottom([]int{1, 2, 3}, 0)
		if len(result) != 0 {
			t.Errorf("Bottom 0 should return empty")
		}
	})

	t.Run("bottom negative", func(t *testing.T) {
		result := Bottom([]int{1, 2, 3}, -1)
		if len(result) != 0 {
			t.Errorf("Bottom negative should return empty")
		}
	})

	t.Run("bottom more than length", func(t *testing.T) {
		result := Bottom([]int{1, 2, 3}, 5)
		if len(result) != 3 {
			t.Errorf("Bottom more than length should return all")
		}
	})

	t.Run("bottom all", func(t *testing.T) {
		result := Bottom([]int{1, 2, 3}, 3)
		if len(result) != 3 {
			t.Errorf("Bottom all should return all")
		}
	})
}

// TestChunk 测试 Chunk 函数
func TestChunk(t *testing.T) {
	t.Run("chunk size 2", func(t *testing.T) {
		result := Chunk([]int{1, 2, 3, 4, 5, 6}, 2)
		if len(result) != 3 {
			t.Errorf("Chunk length = %d, want 3", len(result))
		}
		if len(result[0]) != 2 || result[0][0] != 1 || result[0][1] != 2 {
			t.Errorf("Chunk[0] incorrect")
		}
		if len(result[2]) != 2 || result[2][0] != 5 || result[2][1] != 6 {
			t.Errorf("Chunk[2] incorrect")
		}
	})

	t.Run("chunk uneven", func(t *testing.T) {
		result := Chunk([]int{1, 2, 3, 4, 5}, 2)
		if len(result) != 3 {
			t.Errorf("Chunk uneven length = %d, want 3", len(result))
		}
		if len(result[2]) != 1 || result[2][0] != 5 {
			t.Errorf("Chunk last piece incorrect")
		}
	})

	t.Run("chunk size larger than slice", func(t *testing.T) {
		result := Chunk([]int{1, 2, 3}, 10)
		if len(result) != 1 || len(result[0]) != 3 {
			t.Errorf("Chunk large size should return single chunk")
		}
	})

	t.Run("chunk empty slice", func(t *testing.T) {
		result := Chunk([]int{}, 2)
		if len(result) != 0 {
			t.Errorf("Chunk empty should return empty")
		}
	})

	t.Run("chunk size zero", func(t *testing.T) {
		result := Chunk([]int{1, 2, 3}, 0)
		if len(result) != 0 {
			t.Errorf("Chunk size 0 should return empty")
		}
	})

	t.Run("chunk size negative", func(t *testing.T) {
		result := Chunk([]int{1, 2, 3}, -1)
		if len(result) != 0 {
			t.Errorf("Chunk negative size should return empty")
		}
	})

	t.Run("chunk size 1", func(t *testing.T) {
		result := Chunk([]int{1, 2, 3}, 1)
		if len(result) != 3 {
			t.Errorf("Chunk size 1 length = %d, want 3", len(result))
		}
		for i := range result {
			if len(result[i]) != 1 {
				t.Errorf("Chunk[%d] should have length 1", i)
			}
		}
	})
}
