package candy

import "testing"

func TestGenericSliceEqual(t *testing.T) {
	if !GenericSliceEqual([]int{1, 2}, []int{1, 2}) {
		t.Fatalf("expected true")
	}
	s := []int{1, 2}
	if !GenericSliceEqual(s, s) {
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
	if MapEqual(map[string]int{"a": 1}, map[string]int{"a": 1, "b": 2}) {
		t.Fatalf("expected false")
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
	if !PointerEqual[int](nil, nil) {
		t.Fatalf("expected true")
	}
	if PointerEqual[int](nil, &b) {
		t.Fatalf("expected false")
	}
}

func TestStructEqualAndEqualHelpers(t *testing.T) {
	type testStruct struct {
		A int
		B string
	}

	a := testStruct{A: 1, B: "x"}
	b := testStruct{A: 1, B: "x"}
	if !StructEqual(a, b, func(x, y testStruct) bool { return x.A == y.A && x.B == y.B }) {
		t.Fatalf("expected true")
	}

	if !Equal(1, 1) {
		t.Fatalf("expected true")
	}
	if Equal(1, 2) {
		t.Fatalf("expected false")
	}

	if !EqualSlice([]int{1, 2}, []int{2, 1}) {
		t.Fatalf("expected true (order-insensitive)")
	}
	if EqualSlice([]int{1, 2}, []int{1, 3}) {
		t.Fatalf("expected false")
	}

	if !EqualMap(map[string]int{"a": 1}, map[string]int{"a": 1}) {
		t.Fatalf("expected true")
	}
	if EqualMap(map[string]int{"a": 1}, map[string]int{"a": 2}) {
		t.Fatalf("expected false")
	}
}
