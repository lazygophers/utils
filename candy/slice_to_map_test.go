package candy

import (
	"reflect"
	"testing"
)

type sliceToMapItem struct {
	Name  string
	I     int
	I8    int8
	I16   int16
	I32   int32
	I64   int64
	U     uint
	U8    uint8
	U16   uint16
	U32   uint32
	U64   uint64
	F32   float32
	F64   float64
	Bool  bool
	Code  string
	Alias string
}

func assertMapBoolKeys[T comparable](t *testing.T, got map[T]bool, want []T) {
	t.Helper()

	wantSet := make(map[T]struct{}, len(want))
	for _, v := range want {
		wantSet[v] = struct{}{}
	}

	if len(got) != len(wantSet) {
		t.Fatalf("len(got)=%d want=%d (unique)", len(got), len(wantSet))
	}
	for k := range wantSet {
		if !got[k] {
			t.Fatalf("missing key: %v", k)
		}
	}
}

func assertPanics(t *testing.T, fn func()) {
	t.Helper()
	defer func() {
		if recover() == nil {
			t.Fatalf("expected panic")
		}
	}()
	fn()
}

func TestSlice2Map(t *testing.T) {
	tests := []struct {
		name string
		got  any
		want any
	}{
		{
			name: "string",
			got:  Slice2Map([]string{"a", "b", "c"}),
			want: map[string]bool{"a": true, "b": true, "c": true},
		},
		{
			name: "int",
			got:  Slice2Map([]int{1, 2, 3}),
			want: map[int]bool{1: true, 2: true, 3: true},
		},
		{
			name: "empty",
			got:  Slice2Map([]string{}),
			want: map[string]bool{},
		},
		{
			name: "duplicates",
			got:  Slice2Map([]string{"a", "b", "a", "c"}),
			want: map[string]bool{"a": true, "b": true, "c": true},
		},
		{
			name: "float64",
			got:  Slice2Map([]float64{1.1, 2.2, 3.3}),
			want: map[float64]bool{1.1: true, 2.2: true, 3.3: true},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if !reflect.DeepEqual(tt.got, tt.want) {
				t.Fatalf("got=%v want=%v", tt.got, tt.want)
			}
		})
	}
}

func TestSliceField2Map_Basic(t *testing.T) {
	items := []sliceToMapItem{
		{Name: "a", I: 1, I8: 2, I16: 3, I32: 4, I64: 5, U: 6, U8: 7, U16: 8, U32: 9, U64: 10, F32: 1.5, F64: 2.5, Bool: true},
		{Name: "b", I: 2, I8: 3, I16: 4, I32: 5, I64: 6, U: 7, U8: 8, U16: 9, U32: 10, U64: 11, F32: 2.5, F64: 3.5, Bool: false},
		{Name: "a", I: 1, I8: 2, I16: 3, I32: 4, I64: 5, U: 6, U8: 7, U16: 8, U32: 9, U64: 10, F32: 1.5, F64: 2.5, Bool: true},
	}

	assertMapBoolKeys(t, SliceField2MapString(items, "Name"), []string{"a", "b"})
	assertMapBoolKeys(t, SliceField2MapInt(items, "I"), []int{1, 2})
	assertMapBoolKeys(t, SliceField2MapInt8(items, "I8"), []int8{2, 3})
	assertMapBoolKeys(t, SliceField2MapInt16(items, "I16"), []int16{3, 4})
	assertMapBoolKeys(t, SliceField2MapInt32(items, "I32"), []int32{4, 5})
	assertMapBoolKeys(t, SliceField2MapInt64(items, "I64"), []int64{5, 6})
	assertMapBoolKeys(t, SliceField2MapUint(items, "U"), []uint{6, 7})
	assertMapBoolKeys(t, SliceField2MapUint8(items, "U8"), []uint8{7, 8})
	assertMapBoolKeys(t, SliceField2MapUint16(items, "U16"), []uint16{8, 9})
	assertMapBoolKeys(t, SliceField2MapUint32(items, "U32"), []uint32{9, 10})
	assertMapBoolKeys(t, SliceField2MapUint64(items, "U64"), []uint64{10, 11})
	assertMapBoolKeys(t, SliceField2MapFloat32(items, "F32"), []float32{1.5, 2.5})
	assertMapBoolKeys(t, SliceField2MapFloat64(items, "F64"), []float64{2.5, 3.5})
	assertMapBoolKeys(t, SliceField2MapBool(items, "Bool"), []bool{true, false})
}

func TestSliceField2Map_Pointers(t *testing.T) {
	items := []*sliceToMapItem{{Name: "a"}, {Name: "b"}, {Name: "a"}}
	assertMapBoolKeys(t, SliceField2MapString(items, "Name"), []string{"a", "b"})
}

func TestSliceField2Map_Empty(t *testing.T) {
	if got := SliceField2MapString([]sliceToMapItem{}, "Name"); got != nil {
		t.Fatalf("expected nil, got %v", got)
	}
}

func TestSliceField2Map_Panics(t *testing.T) {
	t.Run("field not found", func(t *testing.T) {
		assertPanics(t, func() {
			SliceField2MapString([]sliceToMapItem{{Name: "a"}}, "Nope")
		})
	})

	t.Run("wrong field type", func(t *testing.T) {
		assertPanics(t, func() {
			SliceField2MapInt([]sliceToMapItem{{Name: "a"}}, "Name")
		})
	})

	t.Run("non-struct element", func(t *testing.T) {
		assertPanics(t, func() {
			SliceField2MapInt([]int{1, 2, 3}, "X")
		})
	})

	t.Run("pointer to non-struct", func(t *testing.T) {
		assertPanics(t, func() {
			SliceField2MapInt([]*int{new(int)}, "X")
		})
	})
}
