package candy

import (
	"testing"
)

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
