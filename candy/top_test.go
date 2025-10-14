package candy

import (
	"reflect"
	"testing"
)

func TestTop(t *testing.T) {
	t.Run("normal case - take first 3 elements", func(t *testing.T) {
		input := []int{1, 2, 3, 4, 5}
		expected := []int{1, 2, 3}
		result := Top(input, 3)

		if !reflect.DeepEqual(result, expected) {
			t.Errorf("Top(%v, 3) = %v, want %v", input, result, expected)
		}
	})

	t.Run("n equals slice length", func(t *testing.T) {
		input := []int{1, 2, 3, 4, 5}
		expected := []int{1, 2, 3, 4, 5}
		result := Top(input, 5)

		if !reflect.DeepEqual(result, expected) {
			t.Errorf("Top(%v, 5) = %v, want %v", input, result, expected)
		}
	})

	t.Run("n greater than slice length", func(t *testing.T) {
		input := []int{1, 2, 3}
		expected := []int{1, 2, 3}
		result := Top(input, 10)

		if !reflect.DeepEqual(result, expected) {
			t.Errorf("Top(%v, 10) = %v, want %v", input, result, expected)
		}
	})

	t.Run("n is zero", func(t *testing.T) {
		input := []int{1, 2, 3, 4, 5}
		expected := []int{}
		result := Top(input, 0)

		if !reflect.DeepEqual(result, expected) {
			t.Errorf("Top(%v, 0) = %v, want %v", input, result, expected)
		}

		// Verify it returns empty slice, not nil
		if result == nil {
			t.Error("Top(slice, 0) should return empty slice, not nil")
		}
	})

	t.Run("n is negative", func(t *testing.T) {
		input := []int{1, 2, 3, 4, 5}
		expected := []int{}
		result := Top(input, -5)

		if !reflect.DeepEqual(result, expected) {
			t.Errorf("Top(%v, -5) = %v, want %v", input, result, expected)
		}

		// Verify it returns empty slice, not nil
		if result == nil {
			t.Error("Top(slice, negative) should return empty slice, not nil")
		}
	})

	t.Run("empty slice", func(t *testing.T) {
		input := []int{}
		expected := []int{}
		result := Top(input, 5)

		if !reflect.DeepEqual(result, expected) {
			t.Errorf("Top(%v, 5) = %v, want %v", input, result, expected)
		}
	})

	t.Run("nil slice", func(t *testing.T) {
		var input []int
		expected := []int{}
		result := Top(input, 5)

		if !reflect.DeepEqual(result, expected) {
			t.Errorf("Top(nil, 5) = %v, want %v", result, expected)
		}
	})

	t.Run("single element slice with n=1", func(t *testing.T) {
		input := []string{"only"}
		expected := []string{"only"}
		result := Top(input, 1)

		if !reflect.DeepEqual(result, expected) {
			t.Errorf("Top(%v, 1) = %v, want %v", input, result, expected)
		}
	})

	t.Run("single element slice with n>1", func(t *testing.T) {
		input := []string{"only"}
		expected := []string{"only"}
		result := Top(input, 10)

		if !reflect.DeepEqual(result, expected) {
			t.Errorf("Top(%v, 10) = %v, want %v", input, result, expected)
		}
	})

	t.Run("string slice", func(t *testing.T) {
		input := []string{"apple", "banana", "cherry", "date", "elderberry"}
		expected := []string{"apple", "banana"}
		result := Top(input, 2)

		if !reflect.DeepEqual(result, expected) {
			t.Errorf("Top(%v, 2) = %v, want %v", input, result, expected)
		}
	})

	t.Run("float slice", func(t *testing.T) {
		input := []float64{1.1, 2.2, 3.3, 4.4, 5.5}
		expected := []float64{1.1, 2.2, 3.3}
		result := Top(input, 3)

		if !reflect.DeepEqual(result, expected) {
			t.Errorf("Top(%v, 3) = %v, want %v", input, result, expected)
		}
	})

	t.Run("returns new slice - modification independence", func(t *testing.T) {
		input := []int{1, 2, 3, 4, 5}
		result := Top(input, 3)

		// Modify the result
		result[0] = 999

		// Original should be unchanged
		if input[0] != 1 {
			t.Errorf("Modifying Top result affected original slice: input[0] = %d, want 1", input[0])
		}
	})

	t.Run("struct slice", func(t *testing.T) {
		type User struct {
			ID   int
			Name string
		}
		input := []User{{1, "Alice"}, {2, "Bob"}, {3, "Charlie"}, {4, "David"}}
		expected := []User{{1, "Alice"}, {2, "Bob"}}
		result := Top(input, 2)

		if !reflect.DeepEqual(result, expected) {
			t.Errorf("Top() = %v, want %v", result, expected)
		}
	})

	t.Run("pointer slice", func(t *testing.T) {
		a, b, c := 1, 2, 3
		input := []*int{&a, &b, &c}
		result := Top(input, 2)

		if len(result) != 2 {
			t.Errorf("Top() length = %d, want 2", len(result))
		}

		if result[0] != &a || result[1] != &b {
			t.Error("Top() did not return correct pointers")
		}
	})

	t.Run("byte slice", func(t *testing.T) {
		input := []byte{0x01, 0x02, 0x03, 0x04, 0x05}
		expected := []byte{0x01, 0x02}
		result := Top(input, 2)

		if !reflect.DeepEqual(result, expected) {
			t.Errorf("Top(%v, 2) = %v, want %v", input, result, expected)
		}
	})

	t.Run("rune slice", func(t *testing.T) {
		input := []rune{'a', 'b', 'c', 'd', 'e'}
		expected := []rune{'a', 'b', 'c'}
		result := Top(input, 3)

		if !reflect.DeepEqual(result, expected) {
			t.Errorf("Top(%v, 3) = %v, want %v", input, result, expected)
		}
	})

	t.Run("large slice", func(t *testing.T) {
		input := make([]int, 10000)
		for i := 0; i < 10000; i++ {
			input[i] = i
		}

		result := Top(input, 100)

		if len(result) != 100 {
			t.Errorf("Top(large slice, 100) length = %d, want 100", len(result))
		}

		// Verify first and last elements
		if result[0] != 0 {
			t.Errorf("Top(large slice, 100)[0] = %d, want 0", result[0])
		}
		if result[99] != 99 {
			t.Errorf("Top(large slice, 100)[99] = %d, want 99", result[99])
		}
	})

	t.Run("n=1 boundary", func(t *testing.T) {
		input := []int{10, 20, 30}
		expected := []int{10}
		result := Top(input, 1)

		if !reflect.DeepEqual(result, expected) {
			t.Errorf("Top(%v, 1) = %v, want %v", input, result, expected)
		}
	})

	t.Run("interface slice", func(t *testing.T) {
		input := []interface{}{1, "two", 3.0, true}
		expected := []interface{}{1, "two"}
		result := Top(input, 2)

		if !reflect.DeepEqual(result, expected) {
			t.Errorf("Top() = %v, want %v", result, expected)
		}
	})

	t.Run("slice of slices", func(t *testing.T) {
		input := [][]int{{1, 2}, {3, 4}, {5, 6}, {7, 8}}
		expected := [][]int{{1, 2}, {3, 4}}
		result := Top(input, 2)

		if !reflect.DeepEqual(result, expected) {
			t.Errorf("Top() = %v, want %v", result, expected)
		}
	})

	t.Run("map slice", func(t *testing.T) {
		input := []map[string]int{
			{"a": 1},
			{"b": 2},
			{"c": 3},
		}
		result := Top(input, 2)

		if len(result) != 2 {
			t.Errorf("Top() length = %d, want 2", len(result))
		}
	})
}

func BenchmarkTop(b *testing.B) {
	input := make([]int, 1000)
	for i := 0; i < 1000; i++ {
		input[i] = i
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Top(input, 100)
	}
}

func BenchmarkTop_SmallN(b *testing.B) {
	input := make([]int, 1000)
	for i := 0; i < 1000; i++ {
		input[i] = i
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Top(input, 10)
	}
}

func BenchmarkTop_LargeN(b *testing.B) {
	input := make([]int, 1000)
	for i := 0; i < 1000; i++ {
		input[i] = i
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Top(input, 900)
	}
}
