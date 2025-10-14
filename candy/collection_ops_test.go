package candy

import (
	"strconv"
	"testing"
)

// TestAll 测试 All 函数
func TestAll(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		predicate func(int) bool
		want     bool
	}{
		{
			name:      "empty slice returns true",
			input:     []int{},
			predicate: func(n int) bool { return n > 0 },
			want:      true,
		},
		{
			name:      "all elements satisfy condition",
			input:     []int{2, 4, 6, 8},
			predicate: func(n int) bool { return n%2 == 0 },
			want:      true,
		},
		{
			name:      "not all elements satisfy condition",
			input:     []int{2, 3, 4, 6},
			predicate: func(n int) bool { return n%2 == 0 },
			want:      false,
		},
		{
			name:      "none satisfy condition",
			input:     []int{1, 3, 5, 7},
			predicate: func(n int) bool { return n%2 == 0 },
			want:      false,
		},
		{
			name:      "single element satisfies",
			input:     []int{2},
			predicate: func(n int) bool { return n%2 == 0 },
			want:      true,
		},
		{
			name:      "single element does not satisfy",
			input:     []int{3},
			predicate: func(n int) bool { return n%2 == 0 },
			want:      false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := All(tt.input, tt.predicate)
			if got != tt.want {
				t.Errorf("All() = %v, want %v", got, tt.want)
			}
		})
	}
}

// TestAny 测试 Any 函数
func TestAny(t *testing.T) {
	tests := []struct {
		name      string
		input     []int
		predicate func(int) bool
		want      bool
	}{
		{
			name:      "empty slice returns false",
			input:     []int{},
			predicate: func(n int) bool { return n > 0 },
			want:      false,
		},
		{
			name:      "one element satisfies",
			input:     []int{1, 2, 3, 4},
			predicate: func(n int) bool { return n%2 == 0 },
			want:      true,
		},
		{
			name:      "no elements satisfy",
			input:     []int{1, 3, 5, 7},
			predicate: func(n int) bool { return n%2 == 0 },
			want:      false,
		},
		{
			name:      "all elements satisfy",
			input:     []int{2, 4, 6, 8},
			predicate: func(n int) bool { return n%2 == 0 },
			want:      true,
		},
		{
			name:      "first element satisfies",
			input:     []int{2, 1, 3, 5},
			predicate: func(n int) bool { return n%2 == 0 },
			want:      true,
		},
		{
			name:      "last element satisfies",
			input:     []int{1, 3, 5, 8},
			predicate: func(n int) bool { return n%2 == 0 },
			want:      true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Any(tt.input, tt.predicate)
			if got != tt.want {
				t.Errorf("Any() = %v, want %v", got, tt.want)
			}
		})
	}
}

// TestEach 测试 Each 函数
func TestEach(t *testing.T) {
	t.Run("empty slice", func(t *testing.T) {
		count := 0
		Each([]int{}, func(n int) {
			count++
		})
		if count != 0 {
			t.Errorf("Each() executed %d times, want 0", count)
		}
	})

	t.Run("iterates over all elements", func(t *testing.T) {
		input := []int{1, 2, 3, 4, 5}
		sum := 0
		Each(input, func(n int) {
			sum += n
		})
		if sum != 15 {
			t.Errorf("Each() sum = %d, want 15", sum)
		}
	})

	t.Run("modifies external variable", func(t *testing.T) {
		input := []string{"a", "b", "c"}
		result := ""
		Each(input, func(s string) {
			result += s
		})
		if result != "abc" {
			t.Errorf("Each() result = %s, want abc", result)
		}
	})

	t.Run("single element", func(t *testing.T) {
		input := []int{42}
		count := 0
		value := 0
		Each(input, func(n int) {
			count++
			value = n
		})
		if count != 1 || value != 42 {
			t.Errorf("Each() count = %d, value = %d, want count = 1, value = 42", count, value)
		}
	})
}

// TestEachReverse 测试 EachReverse 函数
func TestEachReverse(t *testing.T) {
	t.Run("empty slice", func(t *testing.T) {
		count := 0
		EachReverse([]int{}, func(n int) {
			count++
		})
		if count != 0 {
			t.Errorf("EachReverse() executed %d times, want 0", count)
		}
	})

	t.Run("reverses iteration order", func(t *testing.T) {
		input := []int{1, 2, 3, 4, 5}
		result := []int{}
		EachReverse(input, func(n int) {
			result = append(result, n)
		})
		expected := []int{5, 4, 3, 2, 1}
		if len(result) != len(expected) {
			t.Errorf("EachReverse() result length = %d, want %d", len(result), len(expected))
		}
		for i := range result {
			if result[i] != expected[i] {
				t.Errorf("EachReverse() result[%d] = %d, want %d", i, result[i], expected[i])
			}
		}
	})

	t.Run("string concatenation in reverse", func(t *testing.T) {
		input := []string{"a", "b", "c", "d"}
		result := ""
		EachReverse(input, func(s string) {
			result += s
		})
		if result != "dcba" {
			t.Errorf("EachReverse() result = %s, want dcba", result)
		}
	})

	t.Run("single element", func(t *testing.T) {
		input := []int{42}
		count := 0
		value := 0
		EachReverse(input, func(n int) {
			count++
			value = n
		})
		if count != 1 || value != 42 {
			t.Errorf("EachReverse() count = %d, value = %d, want count = 1, value = 42", count, value)
		}
	})
}

// TestMap 测试 Map 函数
func TestMap(t *testing.T) {
	t.Run("empty slice", func(t *testing.T) {
		result := Map([]int{}, func(n int) int {
			return n * 2
		})
		if len(result) != 0 {
			t.Errorf("Map() returned length %d, want 0", len(result))
		}
	})

	t.Run("double integers", func(t *testing.T) {
		input := []int{1, 2, 3, 4, 5}
		result := Map(input, func(n int) int {
			return n * 2
		})
		expected := []int{2, 4, 6, 8, 10}
		if len(result) != len(expected) {
			t.Errorf("Map() length = %d, want %d", len(result), len(expected))
		}
		for i := range result {
			if result[i] != expected[i] {
				t.Errorf("Map() result[%d] = %d, want %d", i, result[i], expected[i])
			}
		}
	})

	t.Run("convert int to string", func(t *testing.T) {
		input := []int{1, 2, 3}
		result := Map(input, func(n int) string {
			return strconv.Itoa(n)
		})
		expected := []string{"1", "2", "3"}
		if len(result) != len(expected) {
			t.Errorf("Map() length = %d, want %d", len(result), len(expected))
		}
		for i := range result {
			if result[i] != expected[i] {
				t.Errorf("Map() result[%d] = %s, want %s", i, result[i], expected[i])
			}
		}
	})

	t.Run("single element", func(t *testing.T) {
		input := []int{5}
		result := Map(input, func(n int) int {
			return n * n
		})
		if len(result) != 1 || result[0] != 25 {
			t.Errorf("Map() = %v, want [25]", result)
		}
	})
}

// TestReduce 测试 Reduce 函数
func TestReduce(t *testing.T) {
	t.Run("empty slice returns zero value", func(t *testing.T) {
		result := Reduce([]int{}, func(a, b int) int {
			return a + b
		})
		if result != 0 {
			t.Errorf("Reduce() = %d, want 0", result)
		}
	})

	t.Run("sum integers", func(t *testing.T) {
		input := []int{1, 2, 3, 4, 5}
		result := Reduce(input, func(a, b int) int {
			return a + b
		})
		if result != 15 {
			t.Errorf("Reduce() = %d, want 15", result)
		}
	})

	t.Run("multiply integers", func(t *testing.T) {
		input := []int{2, 3, 4}
		result := Reduce(input, func(a, b int) int {
			return a * b
		})
		if result != 24 {
			t.Errorf("Reduce() = %d, want 24", result)
		}
	})

	t.Run("concatenate strings", func(t *testing.T) {
		input := []string{"Hello", " ", "World"}
		result := Reduce(input, func(a, b string) string {
			return a + b
		})
		if result != "Hello World" {
			t.Errorf("Reduce() = %s, want 'Hello World'", result)
		}
	})

	t.Run("single element", func(t *testing.T) {
		input := []int{42}
		result := Reduce(input, func(a, b int) int {
			return a + b
		})
		if result != 42 {
			t.Errorf("Reduce() = %d, want 42", result)
		}
	})

	t.Run("find max", func(t *testing.T) {
		input := []int{3, 7, 2, 9, 1}
		result := Reduce(input, func(a, b int) int {
			if a > b {
				return a
			}
			return b
		})
		if result != 9 {
			t.Errorf("Reduce() = %d, want 9", result)
		}
	})
}

// TestReverse 测试 Reverse 函数
func TestReverse(t *testing.T) {
	t.Run("empty slice", func(t *testing.T) {
		result := Reverse([]int{})
		if len(result) != 0 {
			t.Errorf("Reverse() length = %d, want 0", len(result))
		}
	})

	t.Run("reverse integers", func(t *testing.T) {
		input := []int{1, 2, 3, 4, 5}
		result := Reverse(input)
		expected := []int{5, 4, 3, 2, 1}
		if len(result) != len(expected) {
			t.Errorf("Reverse() length = %d, want %d", len(result), len(expected))
		}
		for i := range result {
			if result[i] != expected[i] {
				t.Errorf("Reverse() result[%d] = %d, want %d", i, result[i], expected[i])
			}
		}
		// Verify original is not modified
		if input[0] != 1 || input[4] != 5 {
			t.Errorf("Reverse() modified original slice")
		}
	})

	t.Run("reverse strings", func(t *testing.T) {
		input := []string{"a", "b", "c"}
		result := Reverse(input)
		expected := []string{"c", "b", "a"}
		if len(result) != len(expected) {
			t.Errorf("Reverse() length = %d, want %d", len(result), len(expected))
		}
		for i := range result {
			if result[i] != expected[i] {
				t.Errorf("Reverse() result[%d] = %s, want %s", i, result[i], expected[i])
			}
		}
	})

	t.Run("single element", func(t *testing.T) {
		input := []int{42}
		result := Reverse(input)
		if len(result) != 1 || result[0] != 42 {
			t.Errorf("Reverse() = %v, want [42]", result)
		}
	})

	t.Run("two elements", func(t *testing.T) {
		input := []int{1, 2}
		result := Reverse(input)
		if len(result) != 2 || result[0] != 2 || result[1] != 1 {
			t.Errorf("Reverse() = %v, want [2, 1]", result)
		}
	})
}

// TestShuffle 测试 Shuffle 函数
func TestShuffle(t *testing.T) {
	t.Run("empty slice", func(t *testing.T) {
		input := []int{}
		result := Shuffle(input)
		if len(result) != 0 {
			t.Errorf("Shuffle() length = %d, want 0", len(result))
		}
	})

	t.Run("single element", func(t *testing.T) {
		input := []int{42}
		result := Shuffle(input)
		if len(result) != 1 || result[0] != 42 {
			t.Errorf("Shuffle() = %v, want [42]", result)
		}
	})

	t.Run("shuffles slice in place", func(t *testing.T) {
		input := []int{1, 2, 3, 4, 5}
		original := make([]int, len(input))
		copy(original, input)
		
		result := Shuffle(input)
		
		// Verify it's the same slice reference
		if &result[0] != &input[0] {
			t.Errorf("Shuffle() should return same slice reference")
		}
		
		// Verify all elements are still present
		counts := make(map[int]int)
		for _, v := range result {
			counts[v]++
		}
		for i := 1; i <= 5; i++ {
			if counts[i] != 1 {
				t.Errorf("Shuffle() missing or duplicated element %d", i)
			}
		}
	})

	t.Run("maintains slice length", func(t *testing.T) {
		input := []string{"a", "b", "c", "d", "e", "f"}
		result := Shuffle(input)
		if len(result) != 6 {
			t.Errorf("Shuffle() length = %d, want 6", len(result))
		}
	})
}

// TestSort 测试 Sort 函数
func TestSort(t *testing.T) {
	t.Run("empty slice", func(t *testing.T) {
		result := Sort([]int{})
		if len(result) != 0 {
			t.Errorf("Sort() length = %d, want 0", len(result))
		}
	})

	t.Run("single element", func(t *testing.T) {
		input := []int{42}
		result := Sort(input)
		if len(result) != 1 || result[0] != 42 {
			t.Errorf("Sort() = %v, want [42]", result)
		}
	})

	t.Run("sort integers", func(t *testing.T) {
		input := []int{5, 2, 8, 1, 9, 3}
		result := Sort(input)
		expected := []int{1, 2, 3, 5, 8, 9}
		if len(result) != len(expected) {
			t.Errorf("Sort() length = %d, want %d", len(result), len(expected))
		}
		for i := range result {
			if result[i] != expected[i] {
				t.Errorf("Sort() result[%d] = %d, want %d", i, result[i], expected[i])
			}
		}
		// Verify original is not modified
		if input[0] != 5 {
			t.Errorf("Sort() modified original slice")
		}
	})

	t.Run("sort strings", func(t *testing.T) {
		input := []string{"banana", "apple", "cherry", "date"}
		result := Sort(input)
		expected := []string{"apple", "banana", "cherry", "date"}
		if len(result) != len(expected) {
			t.Errorf("Sort() length = %d, want %d", len(result), len(expected))
		}
		for i := range result {
			if result[i] != expected[i] {
				t.Errorf("Sort() result[%d] = %s, want %s", i, result[i], expected[i])
			}
		}
	})

	t.Run("already sorted", func(t *testing.T) {
		input := []int{1, 2, 3, 4, 5}
		result := Sort(input)
		expected := []int{1, 2, 3, 4, 5}
		for i := range result {
			if result[i] != expected[i] {
				t.Errorf("Sort() result[%d] = %d, want %d", i, result[i], expected[i])
			}
		}
	})

	t.Run("reverse sorted", func(t *testing.T) {
		input := []int{5, 4, 3, 2, 1}
		result := Sort(input)
		expected := []int{1, 2, 3, 4, 5}
		for i := range result {
			if result[i] != expected[i] {
				t.Errorf("Sort() result[%d] = %d, want %d", i, result[i], expected[i])
			}
		}
	})

	t.Run("duplicates", func(t *testing.T) {
		input := []int{3, 1, 4, 1, 5, 9, 2, 6, 5}
		result := Sort(input)
		expected := []int{1, 1, 2, 3, 4, 5, 5, 6, 9}
		for i := range result {
			if result[i] != expected[i] {
				t.Errorf("Sort() result[%d] = %d, want %d", i, result[i], expected[i])
			}
		}
	})
}

// TestSortUsing 测试 SortUsing 函数
func TestSortUsing(t *testing.T) {
	t.Run("empty slice", func(t *testing.T) {
		result := SortUsing([]int{}, func(a, b int) bool {
			return a < b
		})
		if len(result) != 0 {
			t.Errorf("SortUsing() length = %d, want 0", len(result))
		}
	})

	t.Run("single element", func(t *testing.T) {
		input := []int{42}
		result := SortUsing(input, func(a, b int) bool {
			return a < b
		})
		if len(result) != 1 || result[0] != 42 {
			t.Errorf("SortUsing() = %v, want [42]", result)
		}
	})

	t.Run("sort ascending", func(t *testing.T) {
		input := []int{5, 2, 8, 1, 9, 3}
		result := SortUsing(input, func(a, b int) bool {
			return a < b
		})
		expected := []int{1, 2, 3, 5, 8, 9}
		for i := range result {
			if result[i] != expected[i] {
				t.Errorf("SortUsing() result[%d] = %d, want %d", i, result[i], expected[i])
			}
		}
	})

	t.Run("sort descending", func(t *testing.T) {
		input := []int{5, 2, 8, 1, 9, 3}
		result := SortUsing(input, func(a, b int) bool {
			return a > b
		})
		expected := []int{9, 8, 5, 3, 2, 1}
		for i := range result {
			if result[i] != expected[i] {
				t.Errorf("SortUsing() result[%d] = %d, want %d", i, result[i], expected[i])
			}
		}
	})

	t.Run("sort by length", func(t *testing.T) {
		input := []string{"aaa", "b", "cc", "dddd"}
		result := SortUsing(input, func(a, b string) bool {
			return len(a) < len(b)
		})
		expected := []string{"b", "cc", "aaa", "dddd"}
		for i := range result {
			if result[i] != expected[i] {
				t.Errorf("SortUsing() result[%d] = %s, want %s", i, result[i], expected[i])
			}
		}
	})

	t.Run("custom struct sort", func(t *testing.T) {
		type Person struct {
			Name string
			Age  int
		}
		input := []Person{
			{"Alice", 30},
			{"Bob", 25},
			{"Charlie", 35},
		}
		result := SortUsing(input, func(a, b Person) bool {
			return a.Age < b.Age
		})
		if result[0].Name != "Bob" || result[1].Name != "Alice" || result[2].Name != "Charlie" {
			t.Errorf("SortUsing() incorrect order")
		}
	})
}

// TestJoin 测试 Join 函数
func TestJoin(t *testing.T) {
	t.Run("empty slice", func(t *testing.T) {
		result := Join([]int{})
		if result != "" {
			t.Errorf("Join() = %s, want empty string", result)
		}
	})

	t.Run("single integer", func(t *testing.T) {
		result := Join([]int{42})
		if result != "42" {
			t.Errorf("Join() = %s, want 42", result)
		}
	})

	t.Run("default separator", func(t *testing.T) {
		input := []int{1, 2, 3, 4, 5}
		result := Join(input)
		if result != "1,2,3,4,5" {
			t.Errorf("Join() = %s, want 1,2,3,4,5", result)
		}
	})

	t.Run("custom separator", func(t *testing.T) {
		input := []int{1, 2, 3, 4, 5}
		result := Join(input, "-")
		if result != "1-2-3-4-5" {
			t.Errorf("Join() = %s, want 1-2-3-4-5", result)
		}
	})

	t.Run("space separator", func(t *testing.T) {
		input := []string{"Hello", "World", "Go"}
		result := Join(input, " ")
		if result != "Hello World Go" {
			t.Errorf("Join() = %s, want 'Hello World Go'", result)
		}
	})

	t.Run("empty separator", func(t *testing.T) {
		input := []string{"a", "b", "c"}
		result := Join(input, "")
		if result != "abc" {
			t.Errorf("Join() = %s, want abc", result)
		}
	})

	t.Run("floats with separator", func(t *testing.T) {
		input := []float64{1.5, 2.5, 3.5}
		result := Join(input, "|")
		if result != "1.500000|2.500000|3.500000" {
			t.Errorf("Join() = %s, want 1.500000|2.500000|3.500000", result)
		}
	})

	t.Run("multiple separator arguments", func(t *testing.T) {
		input := []int{1, 2, 3}
		// Only first separator should be used
		result := Join(input, "-", ":", ";")
		if result != "1-2-3" {
			t.Errorf("Join() = %s, want 1-2-3", result)
		}
	})
}
