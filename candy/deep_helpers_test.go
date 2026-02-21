package candy

import "testing"

func TestGenericSliceEqual(t *testing.T) {
	if !GenericSliceEqual([]int{1, 2}, []int{1, 2}) {
		t.Fatalf("expected true")
	}
	if GenericSliceEqual([]int{1, 2}, []int{2, 1}) {
		t.Fatalf("expected false")
	}
}

func TestMapEqual(t *testing.T) {
	if !MapEqual(map[string]int{"a": 1}, map[string]int{"a": 1}) {
		t.Fatalf("expected true")
	}
	if MapEqual(map[string]int{"a": 1}, map[string]int{"a": 2}) {
		t.Fatalf("expected false")
	}
}

func TestPointerEqual(t *testing.T) {
	var a = 1
	var b = 1
	if !PointerEqual(&a, &b) {
		t.Fatalf("expected true")
	}
	if PointerEqual[int](nil, &b) {
		t.Fatalf("expected false")
	}
}
