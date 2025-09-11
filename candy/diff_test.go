package candy

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDiff(t *testing.T) {
	t.Run("basic diff", func(t *testing.T) {
		ss := []int{1, 2, 3}
		against := []int{2, 3, 4}
		added, removed := Diff(ss, against)
		// added = Remove(against, ss) = 在 ss 中但不在 against 中 = [1]
		// removed = Remove(ss, against) = 在 against 中但不在 ss 中 = [4]
		assert.Equal(t, []int{1}, added)
		assert.Equal(t, []int{4}, removed)
	})

	t.Run("no diff", func(t *testing.T) {
		ss := []int{1, 2, 3}
		against := []int{1, 2, 3}
		added, removed := Diff(ss, against)
		assert.Equal(t, []int{}, added)
		assert.Equal(t, []int{}, removed)
	})

	t.Run("complete diff", func(t *testing.T) {
		ss := []int{1, 2, 3}
		against := []int{4, 5, 6}
		added, removed := Diff(ss, against)
		// added = Remove(against, ss) = 在 ss 中但不在 against 中 = [1,2,3]
		// removed = Remove(ss, against) = 在 against 中但不在 ss 中 = [4,5,6]
		assert.Equal(t, []int{1, 2, 3}, added)
		assert.Equal(t, []int{4, 5, 6}, removed)
	})

	t.Run("empty slices", func(t *testing.T) {
		ss := []int{}
		against := []int{}
		added, removed := Diff(ss, against)
		assert.Equal(t, []int{}, added)
		assert.Equal(t, []int{}, removed)
	})

	t.Run("first slice empty", func(t *testing.T) {
		ss := []int{}
		against := []int{1, 2, 3}
		added, removed := Diff(ss, against)
		// added = Remove(against, ss) = 在 ss 中但不在 against 中 = []
		// removed = Remove(ss, against) = 在 against 中但不在 ss 中 = [1,2,3]
		assert.Equal(t, []int{}, added)
		assert.Equal(t, []int{1, 2, 3}, removed)
	})

	t.Run("second slice empty", func(t *testing.T) {
		ss := []int{1, 2, 3}
		against := []int{}
		added, removed := Diff(ss, against)
		// added = Remove(against, ss) = 在 ss 中但不在 against 中 = [1,2,3]
		// removed = Remove(ss, against) = 在 against 中但不在 ss 中 = []
		assert.Equal(t, []int{1, 2, 3}, added)
		assert.Equal(t, []int{}, removed)
	})

	t.Run("string slices", func(t *testing.T) {
		ss := []string{"apple", "banana", "cherry"}
		against := []string{"banana", "cherry", "date"}
		added, removed := Diff(ss, against)
		// added = Remove(against, ss) = 在 ss 中但不在 against 中 = ["apple"]
		// removed = Remove(ss, against) = 在 against 中但不在 ss 中 = ["date"]
		assert.Equal(t, []string{"apple"}, added)
		assert.Equal(t, []string{"date"}, removed)
	})

	t.Run("duplicate elements", func(t *testing.T) {
		ss := []int{1, 1, 2, 3}
		against := []int{2, 2, 3, 4}
		added, removed := Diff(ss, against)
		// added = Remove(against, ss) = 在 ss 中但不在 against 中 = [1,1]
		// removed = Remove(ss, against) = 在 against 中但不在 ss 中 = [4] (2已经在ss中存在，所以只有4是新的)
		assert.Equal(t, []int{1, 1}, added)
		assert.Equal(t, []int{4}, removed)
	})
}

func TestRemove(t *testing.T) {
	t.Run("basic remove", func(t *testing.T) {
		ss := []int{1, 2}
		against := []int{2, 3, 4}
		result := Remove(ss, against)
		assert.Equal(t, []int{3, 4}, result)
	})

	t.Run("no elements to remove", func(t *testing.T) {
		ss := []int{1, 2, 3}
		against := []int{1, 2, 3}
		result := Remove(ss, against)
		assert.Equal(t, []int{}, result)
	})

	t.Run("empty ss", func(t *testing.T) {
		ss := []int{}
		against := []int{1, 2, 3}
		result := Remove(ss, against)
		assert.Equal(t, []int{1, 2, 3}, result)
	})

	t.Run("empty against", func(t *testing.T) {
		ss := []int{1, 2, 3}
		against := []int{}
		result := Remove(ss, against)
		assert.Equal(t, []int{}, result)
	})

	t.Run("both empty", func(t *testing.T) {
		ss := []int{}
		against := []int{}
		result := Remove(ss, against)
		assert.Equal(t, []int{}, result)
	})

	t.Run("string slices", func(t *testing.T) {
		ss := []string{"a", "b"}
		against := []string{"b", "c", "d"}
		result := Remove(ss, against)
		assert.Equal(t, []string{"c", "d"}, result)
	})

	t.Run("duplicate elements in against", func(t *testing.T) {
		ss := []int{1, 2}
		against := []int{2, 3, 3, 4}
		result := Remove(ss, against)
		assert.Equal(t, []int{3, 3, 4}, result)
	})

	t.Run("all elements exist in ss", func(t *testing.T) {
		ss := []int{1, 2, 3, 4, 5}
		against := []int{1, 3, 5}
		result := Remove(ss, against)
		assert.Equal(t, []int{}, result)
	})
}