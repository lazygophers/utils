package anyx

import (
	"testing"
)

type innerStruct struct {
	X float64
	Y string
}

type testStruct struct {
	A int
	B string
	C bool
	D []int
	E map[string]int
	F *innerStruct
	G []innerStruct
	H map[string]*innerStruct
}

// TestDeepEqual tests the DeepEqual function
func TestDeepEqual(t *testing.T) {
	type testStruct struct {
		A int
		B string
	}

	tests := []struct {
		name string
		x, y interface{}
		want bool
	}{
		{"nil", nil, nil, true},
		{"int", 1, 1, true},
		{"int_false", 1, 2, false},
		{"string", "hello", "hello", true},
		{"string_false", "hello", "world", false},
		{"bool", true, true, true},
		{"bool_false", true, false, false},
		{"slice_int", []int{1, 2}, []int{1, 2}, true},
		{"slice_int_false", []int{1, 2}, []int{2, 1}, false},
		{"slice_int_len_false", []int{1, 2}, []int{1, 2, 3}, false},
		{"slice_nil_vs_empty", []int(nil), []int{}, false}, // Note: nil and empty slices are not equal
		{"map_string_int", map[string]int{"a": 1}, map[string]int{"a": 1}, true},
		{"map_string_int_false", map[string]int{"a": 1}, map[string]int{"a": 2}, false},
		{"map_string_int_key_false", map[string]int{"a": 1}, map[string]int{"b": 1}, false},
		{"map_nil_vs_empty", map[string]int(nil), map[string]int{}, false}, // Note: nil and empty maps are not equal
		{"struct", testStruct{A: 1, B: "a"}, testStruct{A: 1, B: "a"}, true},
		{"struct_false", testStruct{A: 1, B: "a"}, testStruct{A: 2, B: "a"}, false},
		{"ptr_to_int", new(int), new(int), true},
		{"ptr_to_struct", &testStruct{A: 1}, &testStruct{A: 1}, true},
		{"ptr_to_struct_false", &testStruct{A: 1}, &testStruct{A: 2}, false},
		{"complex_nested_struct",
			&struct{ M map[string]*testStruct }{M: map[string]*testStruct{"foo": {A: 1, B: "bar"}}},
			&struct{ M map[string]*testStruct }{M: map[string]*testStruct{"foo": {A: 1, B: "bar"}}},
			true},
		{"complex_nested_struct_false",
			&struct{ M map[string]*testStruct }{M: map[string]*testStruct{"foo": {A: 1, B: "bar"}}},
			&struct{ M map[string]*testStruct }{M: map[string]*testStruct{"foo": {A: 2, B: "bar"}}},
			false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := DeepEqual(tt.x, tt.y); got != tt.want {
				t.Errorf("DeepEqual() = %v, want %v", got, tt.want)
			}
		})
	}
}

// TestDeepCopy tests the DeepCopy function
func TestDeepCopy(t *testing.T) {
	t.Run("struct_copy", func(t *testing.T) {
		type testStruct struct {
			A int
			B *string
		}

		s := "hello"
		src := testStruct{A: 1, B: &s}
		dst := testStruct{}

		// This will fail because the current DeepCopy implementation is wrong
		DeepCopy(&src, &dst)

		if !DeepEqual(src, dst) {
			t.Errorf("DeepCopy() failed, structs are not equal. src: %+v, dst: %+v", src, dst)
		}

		// Modify dst, src should not be affected
		*dst.B = "world"
		if *src.B == *dst.B {
			t.Errorf("DeepCopy() failed, modification to dst affected src. src.B: %s", *src.B)
		}
	})

	t.Run("slice_of_ptr_copy", func(t *testing.T) {
		type testStruct struct {
			Val int
		}
		src := []*testStruct{{Val: 1}, {Val: 2}}
		dst := make([]*testStruct, len(src))

		// This will work because reflect.Copy works for slices
		DeepCopy(&src, &dst)

		if !DeepEqual(src, dst) {
			t.Errorf("DeepCopy() failed, slices are not equal.")
		}

		// Modify dst, src should not be affected
		dst[0].Val = 99
		if src[0].Val == dst[0].Val {
			t.Errorf("DeepCopy() failed, modification to dst affected src. src[0].Val: %d", src[0].Val)
		}
	})

	t.Run("map_copy", func(t *testing.T) {
		src := map[string]string{"a": "original"}
		dst := make(map[string]string)

		DeepCopy(&src, &dst)

		if !DeepEqual(src, dst) {
			t.Errorf("DeepCopy() failed, maps are not equal. src: %+v, dst: %+v", src, dst)
		}

		// Modify dst, src should not be affected
		dst["a"] = "changed"
		if src["a"] == dst["a"] {
			t.Errorf("DeepCopy() failed, modification to dst affected src. src[a]: %s", src["a"])
		}
	})
}

// --- 性能基准测试 ---

// createComplexStruct 是一个辅助函数，用于创建一个复杂的嵌套结构体实例，用于基准测试。
func createComplexStruct() *testStruct {
	return &testStruct{
		A: 100,
		B: "complex string for benchmark",
		C: true,
		D: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0},
		E: map[string]int{"one": 1, "two": 2, "three": 3},
		F: &innerStruct{
			X: 3.14159,
			Y: "pi",
		},
		G: []innerStruct{
			{X: 1.1, Y: "one-one"},
			{X: 2.2, Y: "two-two"},
		},
		H: map[string]*innerStruct{
			"first":  {X: 10.1, Y: "ten-one"},
			"second": {X: 20.2, Y: "twenty-two"},
		},
	}
}

// BenchmarkDeepEqual 测试 DeepEqual 函数在处理复杂结构体时的性能。
func BenchmarkDeepEqual(b *testing.B) {
	s1 := createComplexStruct()
	s2 := createComplexStruct()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = DeepEqual(s1, s2)
	}
}

// BenchmarkDeepCopy 测试 DeepCopy 函数在处理复杂结构体时的性能。
func BenchmarkDeepCopy(b *testing.B) {
	src := createComplexStruct()
	dst := &testStruct{}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		DeepCopy(src, dst)
	}
}
