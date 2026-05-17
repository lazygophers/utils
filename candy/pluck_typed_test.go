package candy

import "testing"

type pluckAllTypes struct {
	I8   int8
	I16  int16
	I32  int32
	I64  int64
	U    uint
	U8   uint8
	U16  uint16
	U32  uint32
	U64  uint64
	F32  float32
	F64  float64
	Flag bool
	Name string
}

var pluckAllTypesData = []pluckAllTypes{
	{1, 1000, 100000, 1 << 40, 10, 200, 50000, 100000, 1 << 50, 3.14, 2.718, true, "a"},
	{-1, -1000, -100000, -(1 << 40), 20, 100, 60000, 200000, 1 << 55, -1.5, 0.001, false, "b"},
}

func testPluckTyped[T comparable](t *testing.T, fn func(interface{}, string) []T, field string, want []T) {
	t.Helper()
	got := fn(pluckAllTypesData, field)
	if len(got) != len(want) {
		t.Fatalf("%s: len %d, want %d", field, len(got), len(want))
	}
	for i := range want {
		if got[i] != want[i] {
			t.Fatalf("%s[%d]: %v, want %v", field, i, got[i], want[i])
		}
	}
}

func TestPluckAllTypes_Basic(t *testing.T) {
	testPluckTyped(t, PluckInt8, "I8", []int8{1, -1})
	testPluckTyped(t, PluckInt16, "I16", []int16{1000, -1000})
	testPluckTyped(t, PluckInt32, "I32", []int32{100000, -100000})
	testPluckTyped(t, PluckInt64, "I64", []int64{1 << 40, -(1 << 40)})
	testPluckTyped(t, PluckUint, "U", []uint{10, 20})
	testPluckTyped(t, PluckUint8, "U8", []uint8{200, 100})
	testPluckTyped(t, PluckUint16, "U16", []uint16{50000, 60000})
	testPluckTyped(t, PluckUint32, "U32", []uint32{100000, 200000})
	testPluckTyped(t, PluckUint64, "U64", []uint64{1 << 50, 1 << 55})
	testPluckTyped(t, PluckFloat32, "F32", []float32{3.14, -1.5})
	testPluckTyped(t, PluckFloat64, "F64", []float64{2.718, 0.001})
	testPluckTyped(t, PluckBool, "Flag", []bool{true, false})
}

func TestPluckAllTypes_EmptySlice(t *testing.T) {
	var empty []pluckAllTypes
	if got := PluckInt32(empty, "I32"); got != nil {
		t.Fatalf("expected nil, got %v", got)
	}
	if got := PluckUint64(empty, "U64"); got != nil {
		t.Fatalf("expected nil, got %v", got)
	}
	if got := PluckBool(empty, "Flag"); got != nil {
		t.Fatalf("expected nil, got %v", got)
	}
}

func TestPluckAllTypes_TypeMismatch(t *testing.T) {
	assertPanicsPluck(t, func() { PluckInt32(pluckAllTypesData, "Name") })
	assertPanicsPluck(t, func() { PluckUint64(pluckAllTypesData, "I64") })
	assertPanicsPluck(t, func() { PluckBool(pluckAllTypesData, "I32") })
	assertPanicsPluck(t, func() { PluckFloat64(pluckAllTypesData, "I64") })
}

func TestPluckAllTypes_FieldNotFound(t *testing.T) {
	assertPanicsPluck(t, func() { PluckInt32(pluckAllTypesData, "NotExist") })
	assertPanicsPluck(t, func() { PluckUint64(pluckAllTypesData, "NotExist") })
}

func TestPluckAllTypes_NonSlice(t *testing.T) {
	assertPanicsPluck(t, func() { PluckInt32(123, "I32") })
	assertPanicsPluck(t, func() { PluckUint64("x", "U64") })
}

func TestPluckAllTypes_PointerSlice(t *testing.T) {
	items := []*pluckAllTypes{
		{I32: 100, U64: 999},
		{I32: 200, U64: 888},
	}
	got32 := PluckInt32(items, "I32")
	if len(got32) != 2 || got32[0] != 100 || got32[1] != 200 {
		t.Fatalf("unexpected: %v", got32)
	}
	got64 := PluckUint64(items, "U64")
	if len(got64) != 2 || got64[0] != 999 || got64[1] != 888 {
		t.Fatalf("unexpected: %v", got64)
	}
}
