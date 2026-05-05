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

	t.Run("struct with pointer fields", func(t *testing.T) {
		type testStruct struct {
			Value  int
			Ptr    *int
			Slice  []int
			SliceP *[]int
		}

		val := 42
		slice := []int{1, 2, 3}
		src := testStruct{
			Value:  1,
			Ptr:    &val,
			Slice:  []int{4, 5},
			SliceP: &slice,
		}
		var dst testStruct
		DeepCopy(src, &dst)
		if dst.Value != 1 || dst.Ptr == nil || *dst.Ptr != 42 {
			t.Fatalf("unexpected: %+v", dst)
		}
		if len(dst.Slice) != 2 || dst.Slice[0] != 4 {
			t.Fatalf("unexpected slice: %v", dst.Slice)
		}
		if dst.SliceP == nil || len(*dst.SliceP) != 3 {
			t.Fatalf("unexpected slice pointer: %v", dst.SliceP)
		}
	})

	t.Run("nested map with interface values", func(t *testing.T) {
		src := map[string]interface{}{
			"a": []int{1, 2, 3},
			"b": map[string]int{"x": 1},
		}
		var dst map[string]interface{}
		DeepCopy(src, &dst)
		if len(dst) != 2 {
			t.Fatalf("unexpected: %v", dst)
		}
	})

	t.Run("nested slice with interface elements", func(t *testing.T) {
		src := []interface{}{1, "test", []int{1, 2}, map[string]int{"a": 1}}
		var dst []interface{}
		DeepCopy(src, &dst)
		if len(dst) != 4 {
			t.Fatalf("unexpected: %v", dst)
		}
	})

	t.Run("nil interface values", func(t *testing.T) {
		src := []interface{}{nil, 1, nil, "test", nil}
		var dst []interface{}
		DeepCopy(src, &dst)
		if len(dst) != 5 || dst[0] != nil || dst[2] != nil || dst[4] != nil {
			t.Fatalf("unexpected: %v", dst)
		}
	})

	t.Run("nested struct pointers", func(t *testing.T) {
		type innerStruct struct {
			Value int
		}
		type outerStruct struct {
			Inner *innerStruct
		}

		inner := &innerStruct{Value: 42}
		src := outerStruct{Inner: inner}
		var dst outerStruct
		DeepCopy(src, &dst)
		if dst.Inner == nil || dst.Inner.Value != 42 {
			t.Fatalf("unexpected: %+v", dst)
		}
		src.Inner.Value = 99
		if dst.Inner.Value == 99 {
			t.Fatalf("expected independent copy")
		}
	})

	t.Run("array copy", func(t *testing.T) {
		src := [3]int{1, 2, 3}
		var dst [3]int
		DeepCopy(src, &dst)
		if dst != [3]int{1, 2, 3} {
			t.Fatalf("unexpected: %v", dst)
		}
	})

	t.Run("nested array in struct", func(t *testing.T) {
		type arrayStruct struct {
			Fixed [3]int
		}
		src := arrayStruct{Fixed: [3]int{1, 2, 3}}
		var dst arrayStruct
		DeepCopy(src, &dst)
		if dst.Fixed != [3]int{1, 2, 3} {
			t.Fatalf("unexpected: %+v", dst)
		}
	})

	t.Run("nil source pointer", func(t *testing.T) {
		type ptrStruct struct {
			Value int
		}
		var src *ptrStruct
		var dst *ptrStruct
		DeepCopy(src, &dst)
		if dst != nil {
			t.Fatalf("expected nil, got %v", dst)
		}
	})

	t.Run("nil source pointer in struct", func(t *testing.T) {
		type testStruct struct {
			Ptr *int
		}
		src := testStruct{Ptr: nil}
		var dst testStruct
		DeepCopy(src, &dst)
		if dst.Ptr != nil {
			t.Fatalf("expected nil pointer, got %v", dst.Ptr)
		}
	})

	t.Run("triple nested pointers", func(t *testing.T) {
		val := 42
		ptr1 := &val
		ptr2 := &ptr1
		ptr3 := &ptr2
		var dst ***int
		DeepCopy(&ptr3, &dst)
		if ***dst != 42 {
			t.Fatalf("unexpected: %d", ***dst)
		}
	})

	t.Run("map with nil value", func(t *testing.T) {
		src := map[string]*int{"a": nil, "b": nil}
		var dst map[string]*int
		DeepCopy(src, &dst)
		if len(dst) != 2 {
			t.Fatalf("unexpected: %v", dst)
		}
	})

	t.Run("slice with nil elements", func(t *testing.T) {
		src := []*int{nil, nil, nil}
		var dst []*int
		DeepCopy(src, &dst)
		if len(dst) != 3 {
			t.Fatalf("unexpected: %v", dst)
		}
	})

	t.Run("nil interface in map", func(t *testing.T) {
		src := map[string]interface{}{"a": nil, "b": 1}
		var dst map[string]interface{}
		DeepCopy(src, &dst)
		if len(dst) != 2 || dst["a"] != nil {
			t.Fatalf("unexpected: %v", dst)
		}
	})

	t.Run("deeply nested map in map", func(t *testing.T) {
		src := map[string]map[string]int{
			"a": {"x": 1, "y": 2},
			"b": {"z": 3},
		}
		var dst map[string]map[string]int
		DeepCopy(src, &dst)
		if len(dst) != 2 || dst["a"]["x"] != 1 {
			t.Fatalf("unexpected: %v", dst)
		}
	})

	t.Run("deeply nested slice in slice", func(t *testing.T) {
		src := [][]int{{1, 2}, {3, 4, 5}, {6}}
		var dst [][]int
		DeepCopy(src, &dst)
		if len(dst) != 3 || len(dst[0]) != 2 || dst[0][0] != 1 {
			t.Fatalf("unexpected: %v", dst)
		}
	})

	t.Run("struct with slice of structs", func(t *testing.T) {
		type innerStruct struct {
			Value int
		}
		type outerStruct struct {
			Items []innerStruct
		}
		src := outerStruct{
			Items: []innerStruct{{Value: 1}, {Value: 2}},
		}
		var dst outerStruct
		DeepCopy(src, &dst)
		if len(dst.Items) != 2 || dst.Items[0].Value != 1 {
			t.Fatalf("unexpected: %+v", dst)
		}
	})

	t.Run("nil source interface value", func(t *testing.T) {
		var src interface{} = nil
		var dst interface{}
		DeepCopy(src, &dst)
		if dst != nil {
			t.Fatalf("expected nil, got %v", dst)
		}
	})

	t.Run("empty map", func(t *testing.T) {
		src := map[string]int{}
		var dst map[string]int
		DeepCopy(src, &dst)
		if len(dst) != 0 {
			t.Fatalf("expected empty map, got %v", dst)
		}
	})

	t.Run("empty slice", func(t *testing.T) {
		src := []int{}
		var dst []int
		DeepCopy(src, &dst)
		if len(dst) != 0 {
			t.Fatalf("expected empty slice, got %v", dst)
		}
	})

	t.Run("uint types", func(t *testing.T) {
		type uintStruct struct {
			U8  uint8
			U16 uint16
			U32 uint32
			U64 uint64
			U   uint
		}
		src := uintStruct{U8: 1, U16: 2, U32: 3, U64: 4, U: 5}
		var dst uintStruct
		DeepCopy(src, &dst)
		if dst.U8 != 1 || dst.U16 != 2 || dst.U32 != 3 || dst.U64 != 4 || dst.U != 5 {
			t.Fatalf("unexpected: %+v", dst)
		}
	})

	t.Run("float types", func(t *testing.T) {
		type floatStruct struct {
			F32 float32
			F64 float64
		}
		src := floatStruct{F32: 1.5, F64: 2.5}
		var dst floatStruct
		DeepCopy(src, &dst)
		if dst.F32 != 1.5 || dst.F64 != 2.5 {
			t.Fatalf("unexpected: %+v", dst)
		}
	})

	t.Run("complex types", func(t *testing.T) {
		type complexStruct struct {
			C64  complex64
			C128 complex128
		}
		src := complexStruct{C64: 1 + 2i, C128: 3 + 4i}
		var dst complexStruct
		DeepCopy(src, &dst)
		if dst.C64 != 1+2i || dst.C128 != 3+4i {
			t.Fatalf("unexpected: %+v", dst)
		}
	})

	t.Run("bool type", func(t *testing.T) {
		type boolStruct struct {
			Flag1 bool
			Flag2 bool
		}
		src := boolStruct{Flag1: true, Flag2: false}
		var dst boolStruct
		DeepCopy(src, &dst)
		if dst.Flag1 != true || dst.Flag2 != false {
			t.Fatalf("unexpected: %+v", dst)
		}
	})

	t.Run("pointer to nil pointer", func(t *testing.T) {
		var src *int = nil
		var dst **int
		DeepCopy(&src, &dst)
		if dst != nil {
			t.Fatalf("expected nil, got %v", dst)
		}
	})

	t.Run("nil map value with nil key value", func(t *testing.T) {
		src := map[int]*int{1: nil, 2: nil}
		var dst map[int]*int
		DeepCopy(src, &dst)
		if len(dst) != 2 {
			t.Fatalf("unexpected: %v", dst)
		}
	})

	t.Run("deeply nested pointers with nil", func(t *testing.T) {
		var src **int
		var dst **int
		DeepCopy(&src, &dst)
		if dst != nil {
			t.Fatalf("expected nil, got %v", dst)
		}
	})

	t.Run("complex nested structure", func(t *testing.T) {
		type innerStruct struct {
			Slice []*int
			Map   map[string]*int
		}
		type outerStruct struct {
			Inner *innerStruct
			Data  []interface{}
		}
		val1 := 1
		val2 := 2
		src := outerStruct{
			Inner: &innerStruct{
				Slice: []*int{&val1, nil, &val2},
				Map:   map[string]*int{"a": &val1, "b": nil},
			},
			Data: []interface{}{nil, 1, "test", []int{1, 2}},
		}
		var dst outerStruct
		DeepCopy(src, &dst)
		if dst.Inner == nil || len(dst.Inner.Slice) != 3 {
			t.Fatalf("unexpected: %+v", dst)
		}
	})

	t.Run("nested maps with slices", func(t *testing.T) {
		src := map[string][]int{
			"a": {1, 2, 3},
			"b": {4, 5},
		}
		var dst map[string][]int
		DeepCopy(src, &dst)
		if len(dst) != 2 || len(dst["a"]) != 3 {
			t.Fatalf("unexpected: %v", dst)
		}
	})

	t.Run("nested slices with maps", func(t *testing.T) {
		src := []map[string]int{
			{"a": 1, "b": 2},
			{"c": 3},
		}
		var dst []map[string]int
		DeepCopy(src, &dst)
		if len(dst) != 2 || dst[0]["a"] != 1 {
			t.Fatalf("unexpected: %v", dst)
		}
	})

	t.Run("array with nested structures", func(t *testing.T) {
		type itemStruct struct {
			Name  string
			Value int
		}
		src := [2]itemStruct{{Name: "a", Value: 1}, {Name: "b", Value: 2}}
		var dst [2]itemStruct
		DeepCopy(src, &dst)
		if dst[0].Name != "a" || dst[1].Value != 2 {
			t.Fatalf("unexpected: %+v", dst)
		}
	})

	t.Run("map with struct pointer keys", func(t *testing.T) {
		type keyStruct struct {
			ID int
		}
		src := map[keyStruct]string{{ID: 1}: "a", {ID: 2}: "b"}
		var dst map[keyStruct]string
		DeepCopy(src, &dst)
		if len(dst) != 2 {
			t.Fatalf("unexpected: %v", dst)
		}
	})

	t.Run("slice with interface values", func(t *testing.T) {
		src := []interface{}{1, "test", true, nil}
		var dst []interface{}
		DeepCopy(src, &dst)
		if len(dst) != 4 {
			t.Fatalf("unexpected: %v", dst)
		}
	})

	t.Run("map with interface keys and values", func(t *testing.T) {
		type keyStruct struct {
			ID int
		}
		src := map[keyStruct]interface{}{
			{ID: 1}: "a",
			{ID: 2}: 123,
		}
		var dst map[keyStruct]interface{}
		DeepCopy(src, &dst)
		if len(dst) != 2 {
			t.Fatalf("unexpected: %v", dst)
		}
	})

	t.Run("int8 and int16 types", func(t *testing.T) {
		type smallIntStruct struct {
			I8  int8
			I16 int16
		}
		src := smallIntStruct{I8: -128, I16: 32767}
		var dst smallIntStruct
		DeepCopy(src, &dst)
		if dst.I8 != -128 || dst.I16 != 32767 {
			t.Fatalf("unexpected: %+v", dst)
		}
	})

	t.Run("empty interface with nil", func(t *testing.T) {
		var src interface{} = nil
		var dst interface{}
		DeepCopy(&src, &dst)
		if dst != nil {
			t.Fatalf("expected nil, got %v", dst)
		}
	})

	t.Run("struct with multiple pointer levels", func(t *testing.T) {
		type level3 struct {
			Value int
		}
		type level2 struct {
			L3 *level3
		}
		type level1 struct {
			L2 *level2
		}
		val := level3{Value: 42}
		l2 := level2{L3: &val}
		l1 := level1{L2: &l2}
		var dst level1
		DeepCopy(l1, &dst)
		if dst.L2 == nil || dst.L2.L3 == nil {
			t.Fatalf("unexpected: %+v", dst)
		}
	})

	t.Run("int type boundary values", func(t *testing.T) {
		type intStruct struct {
			Max int
			Min int
		}
		src := intStruct{Max: 2147483647, Min: -2147483648}
		var dst intStruct
		DeepCopy(src, &dst)
		if dst.Max != 2147483647 || dst.Min != -2147483648 {
			t.Fatalf("unexpected: %+v", dst)
		}
	})

	t.Run("slice of empty structs", func(t *testing.T) {
		type emptyStruct struct{}
		src := []emptyStruct{{}, {}, {}}
		var dst []emptyStruct
		DeepCopy(src, &dst)
		if len(dst) != 3 {
			t.Fatalf("unexpected: %v", dst)
		}
	})

	t.Run("map with string keys and nil values", func(t *testing.T) {
		src := map[string]*int{"a": nil, "b": nil, "c": nil}
		var dst map[string]*int
		DeepCopy(src, &dst)
		if len(dst) != 3 {
			t.Fatalf("unexpected: %v", dst)
		}
	})

	t.Run("copy to same type with pointer", func(t *testing.T) {
		type testStruct struct {
			Value int
		}
		src := &testStruct{Value: 42}
		dst := new(testStruct)
		DeepCopy(src, dst)
		if dst.Value != 42 {
			t.Fatalf("unexpected: %+v", dst)
		}
	})

	t.Run("nil pointer to pointer", func(t *testing.T) {
		var src *int = nil
		var dst *int
		DeepCopy(src, &dst)
		if dst != nil {
			t.Fatalf("expected nil, got %v", dst)
		}
	})

	t.Run("struct with slice of pointers", func(t *testing.T) {
		type testStruct struct {
			Items []*int
		}
		val1, val2 := 1, 2
		src := testStruct{Items: []*int{&val1, &val2}}
		var dst testStruct
		DeepCopy(src, &dst)
		if len(dst.Items) != 2 || *dst.Items[0] != 1 {
			t.Fatalf("unexpected: %+v", dst)
		}
	})

	t.Run("map with pointer keys", func(t *testing.T) {
		type keyStruct struct {
			ID int
		}
		key1 := &keyStruct{ID: 1}
		key2 := &keyStruct{ID: 2}
		src := map[*keyStruct]string{key1: "a", key2: "b"}
		var dst map[*keyStruct]string
		DeepCopy(src, &dst)
		if len(dst) != 2 {
			t.Fatalf("unexpected: %v", dst)
		}
	})

	t.Run("empty struct slice", func(t *testing.T) {
		type emptyStruct struct{}
		src := []emptyStruct{{}}
		var dst []emptyStruct
		DeepCopy(src, &dst)
		if len(dst) != 1 {
			t.Fatalf("unexpected: %v", dst)
		}
	})
}
