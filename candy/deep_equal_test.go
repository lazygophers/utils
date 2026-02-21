package candy

import "testing"

func TestDeepEqual(t *testing.T) {
	if !DeepEqual(1, 1) {
		t.Fatalf("expected true")
	}
	if DeepEqual(1, 2) {
		t.Fatalf("expected false")
	}

	if !DeepEqual([]int{1, 2}, []int{1, 2}) {
		t.Fatalf("expected true")
	}
	if DeepEqual([]int{1, 2}, []int{2, 1}) {
		t.Fatalf("expected false")
	}

	if !DeepEqual(map[string]int{"a": 1}, map[string]int{"a": 1}) {
		t.Fatalf("expected true")
	}
	if DeepEqual(map[string]int{"a": 1}, map[string]int{"a": 2}) {
		t.Fatalf("expected false")
	}

	var p1 = &deepCopyStruct{A: 1}
	var p2 = &deepCopyStruct{A: 1}
	if !DeepEqual(p1, p2) {
		t.Fatalf("expected true")
	}
}
