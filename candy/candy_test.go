package candy_test

import (
	"github.com/lazygophers/utils/candy"
	"testing"
)

func TestSpare(t *testing.T) {
	t.Log(candy.Spare([]int{1, 2, 3}, []int{2, 3, 4})) // 4
}

func TestDiff(t *testing.T) {
	t.Log(candy.Diff([]int{1, 2, 3}, []int{2, 3, 4})) // 4
}
