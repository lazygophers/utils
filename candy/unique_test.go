package candy

import (
	"reflect"
	"testing"
)

func TestUnique(t *testing.T) {
	t.Run("integer slice", func(t *testing.T) {
		input := []int{1, 2, 2, 3, 4, 4, 5}
		expected := []int{1, 2, 3, 4, 5}
		result := Unique(input)

		if !reflect.DeepEqual(result, expected) {
			t.Errorf("Unique(%v) = %v, want %v", input, result, expected)
		}
	})

	t.Run("string slice", func(t *testing.T) {
		input := []string{"Alice", "Bob", "Alice", "Charlie", "Bob"}
		expected := []string{"Alice", "Bob", "Charlie"}
		result := Unique(input)

		if !reflect.DeepEqual(result, expected) {
			t.Errorf("Unique(%v) = %v, want %v", input, result, expected)
		}
	})

	t.Run("float64 slice", func(t *testing.T) {
		input := []float64{1.1, 2.2, 1.1, 3.3, 2.2}
		expected := []float64{1.1, 2.2, 3.3}
		result := Unique(input)

		if !reflect.DeepEqual(result, expected) {
			t.Errorf("Unique(%v) = %v, want %v", input, result, expected)
		}
	})

	t.Run("empty slice", func(t *testing.T) {
		input := []int{}
		expected := []int{}
		result := Unique(input)

		if !reflect.DeepEqual(result, expected) {
			t.Errorf("Unique(%v) = %v, want %v", input, result, expected)
		}

		// Verify it returns empty slice, not nil
		if result == nil {
			t.Error("Unique(empty) should return empty slice, not nil")
		}
	})

	t.Run("nil slice", func(t *testing.T) {
		var input []int
		result := Unique(input)

		// Should return empty slice, not nil
		if result == nil {
			t.Error("Unique(nil) should return empty slice, not nil")
		}

		if len(result) != 0 {
			t.Errorf("Unique(nil) length = %d, want 0", len(result))
		}
	})

	t.Run("single element", func(t *testing.T) {
		input := []string{"only"}
		expected := []string{"only"}
		result := Unique(input)

		if !reflect.DeepEqual(result, expected) {
			t.Errorf("Unique(%v) = %v, want %v", input, result, expected)
		}
	})

	t.Run("all duplicates", func(t *testing.T) {
		input := []int{5, 5, 5, 5, 5}
		expected := []int{5}
		result := Unique(input)

		if !reflect.DeepEqual(result, expected) {
			t.Errorf("Unique(%v) = %v, want %v", input, result, expected)
		}
	})

	t.Run("no duplicates", func(t *testing.T) {
		input := []int{1, 2, 3, 4, 5}
		expected := []int{1, 2, 3, 4, 5}
		result := Unique(input)

		if !reflect.DeepEqual(result, expected) {
			t.Errorf("Unique(%v) = %v, want %v", input, result, expected)
		}
	})

	t.Run("preserve order", func(t *testing.T) {
		input := []int{5, 3, 2, 4, 3, 1, 5, 2}
		expected := []int{5, 3, 2, 4, 1}
		result := Unique(input)

		if !reflect.DeepEqual(result, expected) {
			t.Errorf("Unique(%v) = %v, want %v (order preserved)", input, result, expected)
		}
	})

	t.Run("int8 type", func(t *testing.T) {
		input := []int8{1, 2, 1, 3, 2}
		expected := []int8{1, 2, 3}
		result := Unique(input)

		if !reflect.DeepEqual(result, expected) {
			t.Errorf("Unique(%v) = %v, want %v", input, result, expected)
		}
	})

	t.Run("int16 type", func(t *testing.T) {
		input := []int16{100, 200, 100, 300}
		expected := []int16{100, 200, 300}
		result := Unique(input)

		if !reflect.DeepEqual(result, expected) {
			t.Errorf("Unique(%v) = %v, want %v", input, result, expected)
		}
	})

	t.Run("int32 type", func(t *testing.T) {
		input := []int32{1000, 2000, 1000, 3000}
		expected := []int32{1000, 2000, 3000}
		result := Unique(input)

		if !reflect.DeepEqual(result, expected) {
			t.Errorf("Unique(%v) = %v, want %v", input, result, expected)
		}
	})

	t.Run("int64 type", func(t *testing.T) {
		input := []int64{100000, 200000, 100000, 300000}
		expected := []int64{100000, 200000, 300000}
		result := Unique(input)

		if !reflect.DeepEqual(result, expected) {
			t.Errorf("Unique(%v) = %v, want %v", input, result, expected)
		}
	})

	t.Run("uint type", func(t *testing.T) {
		input := []uint{1, 2, 1, 3, 2}
		expected := []uint{1, 2, 3}
		result := Unique(input)

		if !reflect.DeepEqual(result, expected) {
			t.Errorf("Unique(%v) = %v, want %v", input, result, expected)
		}
	})

	t.Run("float32 type", func(t *testing.T) {
		input := []float32{1.1, 2.2, 1.1, 3.3}
		expected := []float32{1.1, 2.2, 3.3}
		result := Unique(input)

		if !reflect.DeepEqual(result, expected) {
			t.Errorf("Unique(%v) = %v, want %v", input, result, expected)
		}
	})

	t.Run("with negative numbers", func(t *testing.T) {
		input := []int{-1, 0, 1, -1, 0, 2}
		expected := []int{-1, 0, 1, 2}
		result := Unique(input)

		if !reflect.DeepEqual(result, expected) {
			t.Errorf("Unique(%v) = %v, want %v", input, result, expected)
		}
	})

	t.Run("large slice", func(t *testing.T) {
		input := make([]int, 1000)
		for i := 0; i < 1000; i++ {
			input[i] = i % 100 // Creates duplicates
		}
		result := Unique(input)

		// Should have exactly 100 unique elements
		if len(result) != 100 {
			t.Errorf("Unique(large slice) length = %d, want 100", len(result))
		}

		// Verify order is preserved (first occurrence)
		for i := 0; i < 100; i++ {
			if result[i] != i {
				t.Errorf("Unique(large slice)[%d] = %d, want %d", i, result[i], i)
				break
			}
		}
	})
}

func TestUniqueUsing(t *testing.T) {
	type User struct {
		ID   int
		Name string
	}

	t.Run("struct slice by ID", func(t *testing.T) {
		input := []User{{1, "Alice"}, {2, "Bob"}, {1, "Alice2"}, {3, "Charlie"}}
		expected := []User{{1, "Alice"}, {2, "Bob"}, {3, "Charlie"}}
		result := UniqueUsing(input, func(u User) any { return u.ID })

		if !reflect.DeepEqual(result, expected) {
			t.Errorf("UniqueUsing() = %v, want %v", result, expected)
		}
	})

	t.Run("struct slice by Name", func(t *testing.T) {
		input := []User{{1, "Alice"}, {2, "Bob"}, {3, "Alice"}, {4, "Charlie"}}
		expected := []User{{1, "Alice"}, {2, "Bob"}, {4, "Charlie"}}
		result := UniqueUsing(input, func(u User) any { return u.Name })

		if !reflect.DeepEqual(result, expected) {
			t.Errorf("UniqueUsing() = %v, want %v", result, expected)
		}
	})

	t.Run("string slice by length", func(t *testing.T) {
		input := []string{"apple", "banana", "kiwi", "grape"}
		// apple=5, banana=6, kiwi=4, grape=5 -> unique lengths: 5, 6, 4
		expected := []string{"apple", "banana", "kiwi"}
		result := UniqueUsing(input, func(s string) any { return len(s) })

		if !reflect.DeepEqual(result, expected) {
			t.Errorf("UniqueUsing() = %v, want %v", result, expected)
		}
	})

	t.Run("string slice by first letter", func(t *testing.T) {
		input := []string{"Alice", "Bob", "Anna", "Charlie", "Bob"}
		expected := []string{"Alice", "Bob", "Charlie"}
		result := UniqueUsing(input, func(s string) any {
			if len(s) > 0 {
				return s[0]
			}
			return ""
		})

		if !reflect.DeepEqual(result, expected) {
			t.Errorf("UniqueUsing() = %v, want %v", result, expected)
		}
	})

	t.Run("empty slice", func(t *testing.T) {
		input := []User{}
		expected := []User{}
		result := UniqueUsing(input, func(u User) any { return u.ID })

		if !reflect.DeepEqual(result, expected) {
			t.Errorf("UniqueUsing(empty) = %v, want %v", result, expected)
		}

		// Verify it returns empty slice, not nil
		if result == nil {
			t.Error("UniqueUsing(empty) should return empty slice, not nil")
		}
	})

	t.Run("nil slice", func(t *testing.T) {
		var input []User
		result := UniqueUsing(input, func(u User) any { return u.ID })

		// Should return empty slice, not nil
		if result == nil {
			t.Error("UniqueUsing(nil) should return empty slice, not nil")
		}

		if len(result) != 0 {
			t.Errorf("UniqueUsing(nil) length = %d, want 0", len(result))
		}
	})

	t.Run("single element", func(t *testing.T) {
		input := []User{{1, "Alice"}}
		expected := []User{{1, "Alice"}}
		result := UniqueUsing(input, func(u User) any { return u.ID })

		if !reflect.DeepEqual(result, expected) {
			t.Errorf("UniqueUsing() = %v, want %v", result, expected)
		}
	})

	t.Run("all duplicates", func(t *testing.T) {
		input := []User{{1, "Alice"}, {1, "Alice2"}, {1, "Alice3"}}
		expected := []User{{1, "Alice"}}
		result := UniqueUsing(input, func(u User) any { return u.ID })

		if !reflect.DeepEqual(result, expected) {
			t.Errorf("UniqueUsing() = %v, want %v", result, expected)
		}
	})

	t.Run("no duplicates", func(t *testing.T) {
		input := []User{{1, "Alice"}, {2, "Bob"}, {3, "Charlie"}}
		expected := []User{{1, "Alice"}, {2, "Bob"}, {3, "Charlie"}}
		result := UniqueUsing(input, func(u User) any { return u.ID })

		if !reflect.DeepEqual(result, expected) {
			t.Errorf("UniqueUsing() = %v, want %v", result, expected)
		}
	})

	t.Run("preserve order", func(t *testing.T) {
		input := []User{{5, "E"}, {3, "C"}, {2, "B"}, {4, "D"}, {3, "C2"}, {1, "A"}}
		expected := []User{{5, "E"}, {3, "C"}, {2, "B"}, {4, "D"}, {1, "A"}}
		result := UniqueUsing(input, func(u User) any { return u.ID })

		if !reflect.DeepEqual(result, expected) {
			t.Errorf("UniqueUsing() = %v, want %v (order preserved)", result, expected)
		}
	})

	t.Run("integer slice with mod function", func(t *testing.T) {
		input := []int{1, 11, 21, 2, 12, 22, 3}
		expected := []int{1, 2, 3}
		result := UniqueUsing(input, func(n int) any { return n % 10 })

		if !reflect.DeepEqual(result, expected) {
			t.Errorf("UniqueUsing() = %v, want %v", result, expected)
		}
	})

	t.Run("nil key values", func(t *testing.T) {
		type Item struct {
			Value *int
		}

		val1 := 1
		val2 := 2
		input := []Item{{&val1}, {nil}, {&val2}, {nil}, {&val1}}
		result := UniqueUsing(input, func(i Item) any { return i.Value })

		// Should have 3 unique: val1 pointer, nil, val2 pointer
		if len(result) != 3 {
			t.Errorf("UniqueUsing() length = %d, want 3", len(result))
		}
	})

	t.Run("string key function", func(t *testing.T) {
		input := []int{123, 456, 123, 789}
		expected := []int{123, 456, 789}
		result := UniqueUsing(input, func(n int) any { return n })

		if !reflect.DeepEqual(result, expected) {
			t.Errorf("UniqueUsing() = %v, want %v", result, expected)
		}
	})

	t.Run("complex key function", func(t *testing.T) {
		type Point struct {
			X, Y int
		}
		input := []Point{{1, 2}, {3, 4}, {1, 2}, {5, 6}}
		expected := []Point{{1, 2}, {3, 4}, {5, 6}}
		result := UniqueUsing(input, func(p Point) any { return [2]int{p.X, p.Y} })

		if !reflect.DeepEqual(result, expected) {
			t.Errorf("UniqueUsing() = %v, want %v", result, expected)
		}
	})

	t.Run("empty string handling", func(t *testing.T) {
		input := []string{"", "a", "", "b", ""}
		expected := []string{"", "a", "b"}
		result := UniqueUsing(input, func(s string) any { return s })

		if !reflect.DeepEqual(result, expected) {
			t.Errorf("UniqueUsing() = %v, want %v", result, expected)
		}
	})
}

func BenchmarkUnique(b *testing.B) {
	input := make([]int, 1000)
	for i := 0; i < 1000; i++ {
		input[i] = i % 100
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Unique(input)
	}
}

func BenchmarkUniqueUsing(b *testing.B) {
	type User struct {
		ID   int
		Name string
	}

	input := make([]User, 1000)
	for i := 0; i < 1000; i++ {
		input[i] = User{ID: i % 100, Name: "User"}
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		UniqueUsing(input, func(u User) any { return u.ID })
	}
}
