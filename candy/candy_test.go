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

func TestTop(t *testing.T) {
	ss := []int{1, 2, 3}

	t.Log(candy.Top(ss, 2))
	t.Log(candy.Top(ss, 3))
	t.Log(candy.Top(ss, 4))
}

func TestBottom(t *testing.T) {
	ss := []int{1, 2, 3}

	t.Log(candy.Bottom(ss, 2))
	t.Log(candy.Bottom(ss, 3))
	t.Log(candy.Bottom(ss, 4))
}
