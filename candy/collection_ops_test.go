package candy

import (
	"reflect"
	"testing"
)

func TestCollectionOps(t *testing.T) {
	t.Run("All/Any", func(t *testing.T) {
		if !All([]int{}, func(n int) bool { return n > 0 }) {
			t.Fatalf("empty slice should be true")
		}
		if All([]int{1, 2, 3}, func(n int) bool { return n%2 == 0 }) {
			t.Fatalf("expected false")
		}
		if !Any([]int{1, 2, 3}, func(n int) bool { return n == 2 }) {
			t.Fatalf("expected true")
		}
	})

	t.Run("Each/EachReverse", func(t *testing.T) {
		var got []int
		Each([]int{1, 2, 3}, func(v int) { got = append(got, v) })
		if !reflect.DeepEqual(got, []int{1, 2, 3}) {
			t.Fatalf("unexpected: %v", got)
		}

		got = nil
		EachReverse([]int{1, 2, 3}, func(v int) { got = append(got, v) })
		if !reflect.DeepEqual(got, []int{3, 2, 1}) {
			t.Fatalf("unexpected: %v", got)
		}
	})

	t.Run("Map/Reduce", func(t *testing.T) {
		if !reflect.DeepEqual(Map([]int{1, 2, 3}, func(v int) int { return v * 2 }), []int{2, 4, 6}) {
			t.Fatalf("unexpected")
		}
		if got := Reduce([]int{1, 2, 3}, func(a, b int) int { return a + b }); got != 6 {
			t.Fatalf("got=%d want=6", got)
		}
	})

	t.Run("Reverse", func(t *testing.T) {
		if !reflect.DeepEqual(Reverse([]int{1, 2, 3}), []int{3, 2, 1}) {
			t.Fatalf("unexpected")
		}
	})

	t.Run("Shuffle", func(t *testing.T) {
		orig := []int{1, 2, 3, 4, 5}
		got := Shuffle(append([]int(nil), orig...))
		// 不要求顺序，只要求元素集合一致
		if len(got) != len(orig) {
			t.Fatalf("unexpected len: %v", got)
		}
		m := make(map[int]int)
		for _, v := range got {
			m[v]++
		}
		for _, v := range orig {
			m[v]--
		}
		for _, n := range m {
			if n != 0 {
				t.Fatalf("unexpected shuffle result: %v", got)
			}
		}
	})

	t.Run("Sort/SortUsing", func(t *testing.T) {
		if !reflect.DeepEqual(Sort([]int{3, 1, 2}), []int{1, 2, 3}) {
			t.Fatalf("unexpected")
		}
		if !reflect.DeepEqual(SortUsing([]string{"b", "aa", "c"}, func(a, b string) bool { return len(a) < len(b) }), []string{"b", "c", "aa"}) {
			t.Fatalf("unexpected")
		}
	})

	t.Run("Join", func(t *testing.T) {
		if got := Join([]int{1, 2, 3}, "-"); got != "1-2-3" {
			t.Fatalf("got=%q", got)
		}
		if got := Join([]string{"a", "b"}); got != "a,b" {
			t.Fatalf("got=%q", got)
		}
	})
}
