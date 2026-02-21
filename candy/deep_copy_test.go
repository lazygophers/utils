package candy

import "testing"

type deepCopyStruct struct {
	A int
	B string
	C []int
	M map[string]int
}

func TestDeepCopy(t *testing.T) {
	t.Run("basic", func(t *testing.T) {
		var dst int
		DeepCopy(42, &dst)
		if dst != 42 {
			t.Fatalf("got=%d want=42", dst)
		}
	})

	t.Run("slice", func(t *testing.T) {
		src := []int{1, 2, 3}
		var dst []int
		DeepCopy(src, &dst)
		if len(dst) != 3 || dst[0] != 1 {
			t.Fatalf("unexpected: %v", dst)
		}
		src[0] = 9
		if dst[0] == 9 {
			t.Fatalf("expected independent copy")
		}
	})

	t.Run("map", func(t *testing.T) {
		src := map[string]int{"a": 1}
		var dst map[string]int
		DeepCopy(src, &dst)
		if dst["a"] != 1 {
			t.Fatalf("unexpected: %v", dst)
		}
		src["a"] = 9
		if dst["a"] == 9 {
			t.Fatalf("expected independent copy")
		}
	})

	t.Run("struct", func(t *testing.T) {
		src := deepCopyStruct{A: 1, B: "x", C: []int{1}, M: map[string]int{"a": 1}}
		var dst deepCopyStruct
		DeepCopy(src, &dst)
		if dst.A != 1 || dst.B != "x" || dst.C[0] != 1 || dst.M["a"] != 1 {
			t.Fatalf("unexpected: %+v", dst)
		}
		src.C[0] = 9
		src.M["a"] = 9
		if dst.C[0] == 9 || dst.M["a"] == 9 {
			t.Fatalf("expected independent copy")
		}
	})

	t.Run("nil slice/map", func(t *testing.T) {
		var srcS []int
		var dstS []int
		DeepCopy(srcS, &dstS)
		if dstS != nil {
			t.Fatalf("expected nil, got %v", dstS)
		}

		var srcM map[string]int
		var dstM map[string]int
		DeepCopy(srcM, &dstM)
		if dstM != nil {
			t.Fatalf("expected nil, got %v", dstM)
		}
	})
}
