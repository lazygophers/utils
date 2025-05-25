package candy_test

import (
	"fmt"
	"github.com/lazygophers/utils/candy"
	"sort"
	"testing"
)

func TestAll(t *testing.T) {
	t.Run("empty slice", func(t *testing.T) {
		var input []int
		result := candy.All(input, func(i int) bool { return i > 0 })
		if !result {
			t.Errorf("expected true")
		}
	})

	t.Run("not all elements match", func(t *testing.T) {
		input := []int{1, 2, -3}
		result := candy.All(input, func(i int) bool { return i > 0 })
		if result {
			t.Errorf("expected false")
		}
	})

	t.Run("all elements match", func(t *testing.T) {
		input := []int{1, 2, 3}
		result := candy.All(input, func(i int) bool { return i > 0 })
		if !result {
			t.Errorf("expected true")
		}
	})
}

func TestShuffle(t *testing.T) {
	t.Run("multiple elements", func(t *testing.T) {
		input := []int{1, 2, 3, 4, 5}
		// 复制原始输入并排序
		expected := make([]int, len(input))
		copy(expected, input)
		sort.Ints(expected)

		result := candy.Shuffle(input)
		// 排序结果应与排序后的原始输入一致
		sortedResult := make([]int, len(result))
		copy(sortedResult, result)
		sort.Ints(sortedResult)

		if !candy.SliceEqual(sortedResult, expected) {
			t.Errorf("expected permutation of %v, got %v", expected, sortedResult)
		}
	})
}

func TestMax(t *testing.T) {
	t.Run("empty slice", func(t *testing.T) {
		var input []int
		result := candy.Max(input)
		if result != 0 {
			t.Errorf("expected 0, got %v", result)
		}
	})

	t.Run("positive numbers", func(t *testing.T) {
		input := []int{1, 3, 2}
		result := candy.Max(input)
		if result != 3 {
			t.Errorf("expected 3, got %v", result)
		}
	})
}

func TestMin(t *testing.T) {
	t.Run("empty slice", func(t *testing.T) {
		var input []int
		result := candy.Min(input)
		if result != 0 {
			t.Errorf("expected 0, got %v", result)
		}
	})

	t.Run("positive numbers", func(t *testing.T) {
		input := []int{5, 1, 3}
		result := candy.Min(input)
		if result != 1 {
			t.Errorf("expected 1, got %v", result)
		}
	})
}

func TestRandom(t *testing.T) {
	t.Run("empty slice", func(t *testing.T) {
		var input []int
		result := candy.Random(input)
		if result != 0 {
			t.Errorf("expected 0, got %v", result)
		}
	})

	t.Run("single element", func(t *testing.T) {
		input := []int{42}
		result := candy.Random(input)
		if result != 42 {
			t.Errorf("expected 42, got %v", result)
		}
	})
}

func TestEachStopWithError(t *testing.T) {
	t.Run("no error", func(t *testing.T) {
		input := []int{1, 2, 3}
		err := candy.EachStopWithError(input, func(i int) error {
			return nil
		})
		if err != nil {
			t.Errorf("expected nil, got %v", err)
		}
	})

	t.Run("with error", func(t *testing.T) {
		input := []int{1, 2, 3}
		expectedErr := fmt.Errorf("test error")
		err := candy.EachStopWithError(input, func(i int) error {
			if i == 2 {
				return expectedErr
			}
			return nil
		})
		if err != expectedErr {
			t.Errorf("expected %v, got %v", expectedErr, err)
		}
	})
}
