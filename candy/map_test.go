package candy

import (
	"testing"
)

func assertSameElements[T comparable](t *testing.T, got, want []T) {
	t.Helper()

	wantSet := make(map[T]int, len(want))
	for _, v := range want {
		wantSet[v]++
	}

	for _, v := range got {
		if wantSet[v] == 0 {
			t.Fatalf("unexpected element: %v (got=%v want=%v)", v, got, want)
		}
		wantSet[v]--
	}

	for v, n := range wantSet {
		if n != 0 {
			t.Fatalf("missing element: %v (got=%v want=%v)", v, got, want)
		}
	}
}

func TestMapKeys(t *testing.T) {
	got := MapKeys(map[int]string{1: "a", 2: "b", 3: "c"})
	assertSameElements(t, got, []int{1, 2, 3})
}

func TestMapKeysIntFamily(t *testing.T) {
	if got := MapKeysInt(nil); got != nil {
		t.Fatalf("expected nil, got %v", got)
	}
	if got := MapKeysInt("x"); got != nil {
		t.Fatalf("expected nil, got %v", got)
	}
	assertSameElements(t, MapKeysInt(map[int]string{1: "a", 2: "b"}), []int{1, 2})
	if got := MapKeysInt(map[string]int{"a": 1}); len(got) != 0 {
		t.Fatalf("expected empty, got %v", got)
	}

	assertSameElements(t, MapKeysInt8(map[int8]string{1: "a"}), []int8{1})
	if got := MapKeysInt8(nil); got != nil {
		t.Fatalf("expected nil, got %v", got)
	}
	if got := MapKeysInt8("x"); got != nil {
		t.Fatalf("expected nil, got %v", got)
	}

	assertSameElements(t, MapKeysInt16(map[int16]string{1: "a"}), []int16{1})
	if got := MapKeysInt16(nil); got != nil {
		t.Fatalf("expected nil, got %v", got)
	}
	if got := MapKeysInt16("x"); got != nil {
		t.Fatalf("expected nil, got %v", got)
	}

	assertSameElements(t, MapKeysInt32(map[int32]string{1: "a"}), []int32{1})
	if got := MapKeysInt32(nil); got != nil {
		t.Fatalf("expected nil, got %v", got)
	}
	if got := MapKeysInt32("x"); got != nil {
		t.Fatalf("expected nil, got %v", got)
	}

	assertSameElements(t, MapKeysInt64(map[int64]string{1: "a"}), []int64{1})
	if got := MapKeysInt64(nil); got != nil {
		t.Fatalf("expected nil, got %v", got)
	}
	if got := MapKeysInt64("x"); got != nil {
		t.Fatalf("expected nil, got %v", got)
	}
}

func TestMapKeysUintFamily(t *testing.T) {
	assertSameElements(t, MapKeysUint(map[uint]string{1: "a"}), []uint{1})
	if got := MapKeysUint(nil); got != nil {
		t.Fatalf("expected nil, got %v", got)
	}
	if got := MapKeysUint("x"); got != nil {
		t.Fatalf("expected nil, got %v", got)
	}

	assertSameElements(t, MapKeysUint8(map[uint8]string{1: "a"}), []uint8{1})
	if got := MapKeysUint8(nil); got != nil {
		t.Fatalf("expected nil, got %v", got)
	}
	if got := MapKeysUint8("x"); got != nil {
		t.Fatalf("expected nil, got %v", got)
	}

	assertSameElements(t, MapKeysUint16(map[uint16]string{1: "a"}), []uint16{1})
	if got := MapKeysUint16(nil); got != nil {
		t.Fatalf("expected nil, got %v", got)
	}
	if got := MapKeysUint16("x"); got != nil {
		t.Fatalf("expected nil, got %v", got)
	}

	assertSameElements(t, MapKeysUint32(map[uint32]string{1: "a"}), []uint32{1})
	if got := MapKeysUint32(nil); got != nil {
		t.Fatalf("expected nil, got %v", got)
	}
	if got := MapKeysUint32("x"); got != nil {
		t.Fatalf("expected nil, got %v", got)
	}

	assertSameElements(t, MapKeysUint64(map[uint64]string{1: "a"}), []uint64{1})
	if got := MapKeysUint64(nil); got != nil {
		t.Fatalf("expected nil, got %v", got)
	}
	if got := MapKeysUint64("x"); got != nil {
		t.Fatalf("expected nil, got %v", got)
	}
}

func TestMapKeysFloatFamily(t *testing.T) {
	assertSameElements(t, MapKeysFloat32(map[float32]string{1.5: "a"}), []float32{1.5})
	if got := MapKeysFloat32(nil); got != nil {
		t.Fatalf("expected nil, got %v", got)
	}
	if got := MapKeysFloat32("x"); got != nil {
		t.Fatalf("expected nil, got %v", got)
	}

	assertSameElements(t, MapKeysFloat64(map[float64]string{1.5: "a"}), []float64{1.5})
	if got := MapKeysFloat64(nil); got != nil {
		t.Fatalf("expected nil, got %v", got)
	}
	if got := MapKeysFloat64("x"); got != nil {
		t.Fatalf("expected nil, got %v", got)
	}
}

func TestMapKeysMisc(t *testing.T) {
	assertSameElements(t, MapKeysString(map[string]int{"a": 1, "b": 2}), []string{"a", "b"})
	if got := MapKeysString(nil); got != nil {
		t.Fatalf("expected nil, got %v", got)
	}
	if got := MapKeysString(123); got != nil {
		t.Fatalf("expected nil, got %v", got)
	}

	gotAny := MapKeysAny(map[any]any{1: "a", "b": 2})
	if len(gotAny) != 2 {
		t.Fatalf("unexpected: %v", gotAny)
	}
	if got := MapKeysAny(nil); got != nil {
		t.Fatalf("expected nil, got %v", got)
	}
	if got := MapKeysAny(123); got != nil {
		t.Fatalf("expected nil, got %v", got)
	}
}

func TestMapValues(t *testing.T) {
	got := MapValues(map[int]string{1: "a", 2: "b"})
	assertSameElements(t, got, []string{"a", "b"})
}
