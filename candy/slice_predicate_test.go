package candy

import (
	"testing"
)

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

	t.Run("uncomparable elements", func(t *testing.T) {
		defer func() {
			if r := recover(); r == nil {
				t.Errorf("SliceEqual with uncomparable elements should panic")
			}
		}()
		a := []func(){func() {}}
		b := []func(){func() {}}
		result := SliceEqual(a, b)
		_ = result
	})

	t.Run("equal slices with duplicates", func(t *testing.T) {
		a := []int{1, 2, 2, 3}
		b := []int{1, 2, 2, 3}
		if !SliceEqual(a, b) {
			t.Errorf("SliceEqual with duplicates should be true")
		}
	})

	t.Run("different order but same elements", func(t *testing.T) {
		a := []int{1, 2, 3}
		b := []int{3, 2, 1}
		if !SliceEqual(a, b) {
			t.Errorf("SliceEqual with different order should be true")
		}
	})

	t.Run("elements with different counts", func(t *testing.T) {
		// Test count != 0 branch in SliceEqual
		a := []int{1, 2, 2, 3}
		b := []int{1, 2, 3}
		if SliceEqual(a, b) {
			t.Errorf("SliceEqual with different element counts should be false")
		}
	})

	t.Run("all elements same count", func(t *testing.T) {
		// Test count == 0 branch in SliceEqual
		a := []int{1, 2, 3}
		b := []int{1, 2, 3}
		if !SliceEqual(a, b) {
			t.Errorf("SliceEqual with same elements should be true")
		}
	})

	t.Run("elements with different counts", func(t *testing.T) {
		// Test count != 0 branch in SliceEqual
		a := []int{1, 2, 2, 3}
		b := []int{1, 2, 3}
		if SliceEqual(a, b) {
			t.Errorf("SliceEqual with different element counts should be false")
		}
	})

	t.Run("nil elements in slices", func(t *testing.T) {
		a := []*int{nil, nil, nil}
		b := []*int{nil, nil, nil}
		if !SliceEqual(a, b) {
			t.Errorf("SliceEqual with nil elements should be true")
		}
	})

	t.Run("different nil element counts", func(t *testing.T) {
		a := []*int{nil, nil}
		b := []*int{nil, nil, nil}
		if SliceEqual(a, b) {
			t.Errorf("SliceEqual with different nil counts should be false")
		}
	})

	t.Run("element not in second slice", func(t *testing.T) {
		a := []int{1, 2, 3}
		b := []int{1, 2, 4}
		if SliceEqual(a, b) {
			t.Errorf("SliceEqual with element not in second slice should be false")
		}
	})

	t.Run("element count mismatch", func(t *testing.T) {
		a := []int{1, 2, 2, 3}
		b := []int{1, 2, 3, 3}
		if SliceEqual(a, b) {
			t.Errorf("SliceEqual with element count mismatch should be false")
		}
	})

	t.Run("element count mismatch 2", func(t *testing.T) {
		a := []int{1, 2, 2, 3}
		b := []int{1, 2, 3, 3}
		if SliceEqual(a, b) {
			t.Errorf("SliceEqual with element count mismatch 2 should be false")
		}
	})

	t.Run("element count mismatch 3", func(t *testing.T) {
		a := []int{1, 2, 3}
		b := []int{1, 2, 2, 3}
		if SliceEqual(a, b) {
			t.Errorf("SliceEqual with element count mismatch 3 should be false")
		}
	})

	t.Run("element appears more in second slice", func(t *testing.T) {
		a := []int{1, 2, 3}
		b := []int{1, 2, 2, 3}
		if SliceEqual(a, b) {
			t.Errorf("SliceEqual with element appearing more in second slice should be false")
		}
	})

	t.Run("element appears less in second slice", func(t *testing.T) {
		a := []int{1, 2, 2, 3}
		b := []int{1, 2, 3}
		if SliceEqual(a, b) {
			t.Errorf("SliceEqual with element appearing less in second slice should be false")
		}
	})

	t.Run("element count mismatch", func(t *testing.T) {
		a := []int{1, 2, 2, 3}
		b := []int{1, 2, 3, 3}
		if SliceEqual(a, b) {
			t.Errorf("SliceEqual with element count mismatch should be false")
		}
	})

	t.Run("uncomparable elements", func(t *testing.T) {
		defer func() {
			if r := recover(); r == nil {
				t.Errorf("SliceEqual with uncomparable elements should panic")
			}
		}()
		a := []func(){func() {}}
		b := []func(){func() {}}
		result := SliceEqual(a, b)
		_ = result
	})
}

// TestDrop 测试 Drop 函数
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
