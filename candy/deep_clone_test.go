package candy

import "testing"

func TestClone(t *testing.T) {
	src := deepCopyStruct{A: 1, C: []int{1}}
	dst := Clone(src)
	src.C[0] = 9
	if dst.C[0] == 9 {
		t.Fatalf("expected deep copy")
	}
}

func TestCloneSlice(t *testing.T) {
	src := []deepCopyStruct{{A: 1, C: []int{1}}}
	dst := CloneSlice(src)
	src[0].C[0] = 9
	if dst[0].C[0] == 9 {
		t.Fatalf("expected deep copy")
	}
}

func TestCloneMap(t *testing.T) {
	src := map[string]deepCopyStruct{"a": {A: 1, C: []int{1}}}
	dst := CloneMap(src)
	src["a"] = deepCopyStruct{A: 2}
	if dst["a"].A != 1 {
		t.Fatalf("unexpected: %v", dst)
	}
}
