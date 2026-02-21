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
	assertSameElements(t, MapKeysInt16(map[int16]string{1: "a"}), []int16{1})
	assertSameElements(t, MapKeysInt32(map[int32]string{1: "a"}), []int32{1})
	assertSameElements(t, MapKeysInt64(map[int64]string{1: "a"}), []int64{1})
}

func TestMapKeysUintFamily(t *testing.T) {
	assertSameElements(t, MapKeysUint(map[uint]string{1: "a"}), []uint{1})
	assertSameElements(t, MapKeysUint8(map[uint8]string{1: "a"}), []uint8{1})
	assertSameElements(t, MapKeysUint16(map[uint16]string{1: "a"}), []uint16{1})
	assertSameElements(t, MapKeysUint32(map[uint32]string{1: "a"}), []uint32{1})
	assertSameElements(t, MapKeysUint64(map[uint64]string{1: "a"}), []uint64{1})
}

func TestMapKeysFloatFamily(t *testing.T) {
	assertSameElements(t, MapKeysFloat32(map[float32]string{1.5: "a"}), []float32{1.5})
	assertSameElements(t, MapKeysFloat64(map[float64]string{1.5: "a"}), []float64{1.5})
}

func TestMapKeysMisc(t *testing.T) {
	assertSameElements(t, MapKeysString(map[string]int{"a": 1, "b": 2}), []string{"a", "b"})

	// MapKeysBytes 目前不会收集 keys，仅验证不会 panic 且返回 slice。
	if got := MapKeysBytes(map[string]int{"a": 1}); got == nil {
		t.Fatalf("expected non-nil slice")
	}

	gotAny := MapKeysAny(map[any]any{1: "a", "b": 2})
	if len(gotAny) != 2 {
		t.Fatalf("unexpected: %v", gotAny)
	}
}

func TestMapValues(t *testing.T) {
	got := MapValues(map[int]string{1: "a", 2: "b"})
	assertSameElements(t, got, []string{"a", "b"})
}
