package candy

import "testing"

func TestTypedSliceCopy(t *testing.T) {
	src := []int{1, 2, 3}
	dst := TypedSliceCopy(src)
	src[0] = 9
	if dst[0] == 9 {
		t.Fatalf("expected independent copy")
	}
}

func TestTypedMapCopy(t *testing.T) {
	src := map[string]int{"a": 1}
	dst := TypedMapCopy(src)
	src["a"] = 9
	if dst["a"] == 9 {
		t.Fatalf("expected independent copy")
	}
}

func TestTypedSliceCopy_Complex(t *testing.T) {
	src := []deepCopyStruct{{A: 1, C: []int{1}}}
	dst := TypedSliceCopy(src)
	src[0].C[0] = 9
	if dst[0].C[0] == 9 {
		t.Fatalf("expected deep copy")
	}
}
