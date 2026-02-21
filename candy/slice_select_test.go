package candy

import (
	"testing"
)

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
