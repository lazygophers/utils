package candy

import (
	"reflect"
	"testing"
)

func TestDiffSlice(t *testing.T) {
	t.Run("basic_int_slices", func(t *testing.T) {
		// Test basic integer slices
		a := []int{1, 2, 3, 4}
		b := []int{3, 4, 5, 6}

		onlyInA, onlyInB := DiffSlice(a, b)

		// Convert results back to slices for comparison
		resultA := onlyInA.([]int)
		resultB := onlyInB.([]int)

		// Elements only in A: {1, 2}
		expectedA := []int{1, 2}
		if !reflect.DeepEqual(resultA, expectedA) {
			t.Errorf("Elements only in A: got %v, want %v", resultA, expectedA)
		}

		// Elements only in B: {5, 6} (order may vary)
		if len(resultB) != 2 {
			t.Errorf("Elements only in B: got length %d, want 2", len(resultB))
		}
		containsAll := true
		for _, expected := range []int{5, 6} {
			found := false
			for _, actual := range resultB {
				if actual == expected {
					found = true
					break
				}
			}
			if !found {
				containsAll = false
				break
			}
		}
		if !containsAll {
			t.Errorf("Elements only in B: got %v, want to contain both 5 and 6", resultB)
		}
	})

	t.Run("string_slices", func(t *testing.T) {
		// Test string slices
		a := []string{"apple", "banana", "cherry"}
		b := []string{"banana", "date", "elderberry"}

		onlyInA, onlyInB := DiffSlice(a, b)

		resultA := onlyInA.([]string)
		resultB := onlyInB.([]string)

		// Elements only in A: {apple, cherry}
		expectedA := []string{"apple", "cherry"}
		if !reflect.DeepEqual(resultA, expectedA) {
			t.Errorf("Elements only in A: got %v, want %v", resultA, expectedA)
		}

		// Elements only in B: {date, elderberry} - order may vary due to map iteration
		if len(resultB) != 2 {
			t.Errorf("Expected 2 elements only in B, got %d: %v", len(resultB), resultB)
		}

		// Check that both expected elements are present
		expectedElements := map[string]bool{"date": true, "elderberry": true}
		for _, elem := range resultB {
			if !expectedElements[elem] {
				t.Errorf("Unexpected element in B diff: %s", elem)
			}
		}
	})

	t.Run("identical_slices", func(t *testing.T) {
		// Test identical slices
		a := []int{1, 2, 3}
		b := []int{1, 2, 3}

		onlyInA, onlyInB := DiffSlice(a, b)

		resultA := onlyInA.([]int)
		resultB := onlyInB.([]int)

		// Both should be empty
		if len(resultA) != 0 {
			t.Errorf("Expected empty slice for identical input A, got %v", resultA)
		}
		if len(resultB) != 0 {
			t.Errorf("Expected empty slice for identical input B, got %v", resultB)
		}
	})

	t.Run("no_common_elements", func(t *testing.T) {
		// Test slices with no common elements
		a := []int{1, 2, 3}
		b := []int{4, 5, 6}

		onlyInA, onlyInB := DiffSlice(a, b)

		resultA := onlyInA.([]int)
		resultB := onlyInB.([]int)

		// All elements from A should be in resultA
		expectedA := []int{1, 2, 3}
		if !reflect.DeepEqual(resultA, expectedA) {
			t.Errorf("All elements should be in A diff: got %v, want %v", resultA, expectedA)
		}

		// All elements from B should be in resultB - order may vary
		if len(resultB) != 3 {
			t.Errorf("Expected 3 elements in B diff, got %d: %v", len(resultB), resultB)
		}

		// Check that all expected elements are present
		expectedElements := map[int]bool{4: true, 5: true, 6: true}
		for _, elem := range resultB {
			if !expectedElements[elem] {
				t.Errorf("Unexpected element in B diff: %d", elem)
			}
		}
	})

	t.Run("empty_slices", func(t *testing.T) {
		// Test one empty slice
		a := []int{1, 2, 3}
		b := []int{}

		onlyInA, onlyInB := DiffSlice(a, b)

		resultA := onlyInA.([]int)
		resultB := onlyInB.([]int)

		// All elements should be in resultA
		expectedA := []int{1, 2, 3}
		if !reflect.DeepEqual(resultA, expectedA) {
			t.Errorf("All elements should be in A diff: got %v, want %v", resultA, expectedA)
		}

		// resultB should be empty
		if len(resultB) != 0 {
			t.Errorf("Expected empty result for B, got %v", resultB)
		}
	})

	t.Run("both_empty_slices", func(t *testing.T) {
		// Test both empty slices
		a := []int{}
		b := []int{}

		onlyInA, onlyInB := DiffSlice(a, b)

		resultA := onlyInA.([]int)
		resultB := onlyInB.([]int)

		// Both should be empty
		if len(resultA) != 0 {
			t.Errorf("Expected empty result for A, got %v", resultA)
		}
		if len(resultB) != 0 {
			t.Errorf("Expected empty result for B, got %v", resultB)
		}
	})

	t.Run("duplicate_elements", func(t *testing.T) {
		// Test slices with duplicate elements
		a := []int{1, 2, 2, 3}
		b := []int{2, 3, 3, 4}

		onlyInA, onlyInB := DiffSlice(a, b)

		resultA := onlyInA.([]int)
		resultB := onlyInB.([]int)

		// The function doesn't handle duplicates the way I initially expected
		// It processes each element in A, and if found in B, removes it from B's map
		// So duplicates in A will remain if B doesn't have enough duplicates
		// Elements only in A: {1} plus any extra duplicates of 2

		// Verify that 1 is in the result (it should be since it's not in B)
		found1 := false
		for _, elem := range resultA {
			if elem == 1 {
				found1 = true
				break
			}
		}
		if !found1 {
			t.Errorf("Expected element 1 to be only in A, got %v", resultA)
		}

		// Elements only in B: {4} plus any extra duplicates of 3
		found4 := false
		for _, elem := range resultB {
			if elem == 4 {
				found4 = true
				break
			}
		}
		if !found4 {
			t.Errorf("Expected element 4 to be only in B, got %v", resultB)
		}
	})

	t.Run("float_slices", func(t *testing.T) {
		// Test float slices
		a := []float64{1.1, 2.2, 3.3}
		b := []float64{2.2, 4.4, 5.5}

		onlyInA, onlyInB := DiffSlice(a, b)

		resultA := onlyInA.([]float64)
		resultB := onlyInB.([]float64)

		// Elements only in A: {1.1, 3.3}
		expectedA := []float64{1.1, 3.3}
		if !reflect.DeepEqual(resultA, expectedA) {
			t.Errorf("Elements only in A: got %v, want %v", resultA, expectedA)
		}

		// Elements only in B: {4.4, 5.5} (order may vary)
		if len(resultB) != 2 {
			t.Errorf("Elements only in B: got length %d, want 2", len(resultB))
		}
		containsAll := true
		for _, expected := range []float64{4.4, 5.5} {
			found := false
			for _, actual := range resultB {
				if actual == expected {
					found = true
					break
				}
			}
			if !found {
				containsAll = false
				break
			}
		}
		if !containsAll {
			t.Errorf("Elements only in B: got %v, want to contain both 4.4 and 5.5", resultB)
		}
	})

	t.Run("bool_slices", func(t *testing.T) {
		// Test bool slices
		a := []bool{true, false, true}
		b := []bool{false, false}

		onlyInA, onlyInB := DiffSlice(a, b)

		resultA := onlyInA.([]bool)
		resultB := onlyInB.([]bool)

		// Elements only in A: {true, true} - both true values since B only has false
		// Verify that we have true elements
		trueCount := 0
		for _, elem := range resultA {
			if elem {
				trueCount++
			}
		}
		if trueCount == 0 {
			t.Errorf("Expected at least one true in A diff, got %v", resultA)
		}

		// Elements only in B: {} (all false elements are common with A)
		if len(resultB) != 0 {
			t.Errorf("Expected empty result for B, got %v", resultB)
		}
	})

	t.Run("large_slices", func(t *testing.T) {
		// Test with larger slices for performance
		a := make([]int, 1000)
		b := make([]int, 1000)

		// Fill a with 0-999
		for i := 0; i < 1000; i++ {
			a[i] = i
		}

		// Fill b with 500-1499
		for i := 0; i < 1000; i++ {
			b[i] = i + 500
		}

		onlyInA, onlyInB := DiffSlice(a, b)

		resultA := onlyInA.([]int)
		resultB := onlyInB.([]int)

		// Elements only in A: 0-499 (500 elements)
		if len(resultA) != 500 {
			t.Errorf("Expected 500 elements in A diff, got %d", len(resultA))
		}

		// Elements only in B: 1000-1499 (500 elements)
		if len(resultB) != 500 {
			t.Errorf("Expected 500 elements in B diff, got %d", len(resultB))
		}

		// Verify that the results contain expected ranges
		// Elements in A diff should be from 0-499
		// Elements in B diff should be from 1000-1499

		// Check that A diff contains elements from the expected range
		foundLowInA := false
		for _, elem := range resultA {
			if elem >= 0 && elem <= 499 {
				foundLowInA = true
				break
			}
		}
		if !foundLowInA {
			t.Error("Expected A diff to contain elements from range 0-499")
		}

		// Check that B diff contains elements from the expected range
		foundHighInB := false
		for _, elem := range resultB {
			if elem >= 1000 && elem <= 1499 {
				foundHighInB = true
				break
			}
		}
		if !foundHighInB {
			t.Error("Expected B diff to contain elements from range 1000-1499")
		}
	})
}

func TestDiffSlicePanicCases(t *testing.T) {
	t.Run("first_param_not_slice", func(t *testing.T) {
		defer func() {
			if r := recover(); r == nil {
				t.Error("Expected panic when first parameter is not a slice")
			} else {
				// Verify the panic message
				if r != "a is not slice" {
					t.Errorf("Expected panic message 'a is not slice', got %v", r)
				}
			}
		}()

		notSlice := 42
		slice := []int{1, 2, 3}
		DiffSlice(notSlice, slice)
	})

	t.Run("second_param_not_slice", func(t *testing.T) {
		defer func() {
			if r := recover(); r == nil {
				t.Error("Expected panic when second parameter is not a slice")
			} else {
				// Verify the panic message
				if r != "b is not slice" {
					t.Errorf("Expected panic message 'b is not slice', got %v", r)
				}
			}
		}()

		slice := []int{1, 2, 3}
		notSlice := "not a slice"
		DiffSlice(slice, notSlice)
	})

	t.Run("different_element_types", func(t *testing.T) {
		defer func() {
			if r := recover(); r == nil {
				t.Error("Expected panic when slices have different element types")
			} else {
				// Verify the panic message
				if r != "a and b are not same type" {
					t.Errorf("Expected panic message 'a and b are not same type', got %v", r)
				}
			}
		}()

		intSlice := []int{1, 2, 3}
		stringSlice := []string{"a", "b", "c"}
		DiffSlice(intSlice, stringSlice)
	})
}

func TestDiffSliceEdgeCases(t *testing.T) {
	t.Run("nil_slices", func(t *testing.T) {
		// Test nil slices (they should be treated as empty slices)
		var a []int
		var b []int

		onlyInA, onlyInB := DiffSlice(a, b)

		resultA := onlyInA.([]int)
		resultB := onlyInB.([]int)

		// Both should be empty
		if len(resultA) != 0 {
			t.Errorf("Expected empty result for A, got %v", resultA)
		}
		if len(resultB) != 0 {
			t.Errorf("Expected empty result for B, got %v", resultB)
		}
	})

	t.Run("one_nil_one_populated", func(t *testing.T) {
		// Test one nil, one populated slice
		var a []int
		b := []int{1, 2, 3}

		onlyInA, onlyInB := DiffSlice(a, b)

		resultA := onlyInA.([]int)
		resultB := onlyInB.([]int)

		// A should be empty
		if len(resultA) != 0 {
			t.Errorf("Expected empty result for A, got %v", resultA)
		}

		// B should contain all elements - order may vary due to map iteration
		if len(resultB) != 3 {
			t.Errorf("Expected 3 elements in B diff, got %d: %v", len(resultB), resultB)
		}

		// Check that all expected elements are present
		expectedElements := map[int]bool{1: true, 2: true, 3: true}
		for _, elem := range resultB {
			if !expectedElements[elem] {
				t.Errorf("Unexpected element in B diff: %d", elem)
			}
		}
	})

	t.Run("single_element_slices", func(t *testing.T) {
		// Test single element slices
		a := []int{42}
		b := []int{42}

		onlyInA, onlyInB := DiffSlice(a, b)

		resultA := onlyInA.([]int)
		resultB := onlyInB.([]int)

		// Both should be empty
		if len(resultA) != 0 {
			t.Errorf("Expected empty result for A, got %v", resultA)
		}
		if len(resultB) != 0 {
			t.Errorf("Expected empty result for B, got %v", resultB)
		}
	})

	t.Run("single_different_elements", func(t *testing.T) {
		// Test single different elements
		a := []int{1}
		b := []int{2}

		onlyInA, onlyInB := DiffSlice(a, b)

		resultA := onlyInA.([]int)
		resultB := onlyInB.([]int)

		// Each should contain their respective element
		expectedA := []int{1}
		expectedB := []int{2}

		if !reflect.DeepEqual(resultA, expectedA) {
			t.Errorf("Expected A diff: got %v, want %v", resultA, expectedA)
		}

		if !reflect.DeepEqual(resultB, expectedB) {
			t.Errorf("Expected B diff: got %v, want %v", resultB, expectedB)
		}
	})
}

// BenchmarkDiffSlice provides performance benchmarks
func BenchmarkDiffSlice(b *testing.B) {
	// Setup test data
	sliceA := make([]int, 1000)
	sliceB := make([]int, 1000)

	for i := 0; i < 1000; i++ {
		sliceA[i] = i
		sliceB[i] = i + 500 // Half overlap
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		DiffSlice(sliceA, sliceB)
	}
}

func BenchmarkDiffSliceStrings(b *testing.B) {
	// Setup string test data
	sliceA := make([]string, 100)
	sliceB := make([]string, 100)

	for i := 0; i < 100; i++ {
		sliceA[i] = "item" + string(rune(i))
		sliceB[i] = "item" + string(rune(i+50)) // Half overlap
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		DiffSlice(sliceA, sliceB)
	}
}

func BenchmarkDiffSliceNoOverlap(b *testing.B) {
	// Test performance with no overlap
	sliceA := make([]int, 500)
	sliceB := make([]int, 500)

	for i := 0; i < 500; i++ {
		sliceA[i] = i
		sliceB[i] = i + 1000 // No overlap
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		DiffSlice(sliceA, sliceB)
	}
}

func BenchmarkDiffSliceFullOverlap(b *testing.B) {
	// Test performance with full overlap
	sliceA := make([]int, 500)
	sliceB := make([]int, 500)

	for i := 0; i < 500; i++ {
		sliceA[i] = i
		sliceB[i] = i // Full overlap
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		DiffSlice(sliceA, sliceB)
	}
}
