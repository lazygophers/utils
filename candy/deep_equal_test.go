package candy

import (
	"reflect"
	"testing"
)

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

func TestDeepEqual_Advanced(t *testing.T) {
	t.Run("arrays", func(t *testing.T) {
		a1 := [3]int{1, 2, 3}
		a2 := [3]int{1, 2, 3}
		if !DeepEqual(a1, a2) {
			t.Fatalf("expected true")
		}
		a3 := [3]int{1, 2, 4}
		if DeepEqual(a1, a3) {
			t.Fatalf("expected false")
		}
	})

	t.Run("nested structs", func(t *testing.T) {
		type inner struct {
			X int
		}
		type outer struct {
			I inner
			Y string
		}
		o1 := outer{I: inner{X: 1}, Y: "test"}
		o2 := outer{I: inner{X: 1}, Y: "test"}
		if !DeepEqual(o1, o2) {
			t.Fatalf("expected true")
		}
	})

	t.Run("interfaces", func(t *testing.T) {
		var i1, i2 interface{} = 1, 1
		if !DeepEqual(i1, i2) {
			t.Fatalf("expected true")
		}
		i3 := interface{}(2)
		if DeepEqual(i1, i3) {
			t.Fatalf("expected false")
		}
	})

	t.Run("nil pointers", func(t *testing.T) {
		var p1, p2 *int
		if !DeepEqual(p1, p2) {
			t.Fatalf("expected true")
		}
		val := 1
		p3 := &val
		if DeepEqual(p1, p3) {
			t.Fatalf("expected false")
		}
	})
}

func TestDeepEqual_EdgeCases(t *testing.T) {
	t.Run("nil slices", func(t *testing.T) {
		var s1 []int
		var s2 []int
		if !DeepEqual(s1, s2) {
			t.Fatalf("expected true for nil slices")
		}
	})

	t.Run("same slice reference", func(t *testing.T) {
		s := []int{1, 2, 3}
		if !DeepEqual(s, s) {
			t.Fatalf("expected true for same reference")
		}
	})

	t.Run("complex numbers", func(t *testing.T) {
		c1 := 1 + 2i
		c2 := 1 + 2i
		if !DeepEqual(c1, c2) {
			t.Fatalf("expected true")
		}
		c3 := 1 + 3i
		if DeepEqual(c1, c3) {
			t.Fatalf("expected false")
		}
	})

	t.Run("nested arrays", func(t *testing.T) {
		a1 := [2][3]int{{1, 2, 3}, {4, 5, 6}}
		a2 := [2][3]int{{1, 2, 3}, {4, 5, 6}}
		if !DeepEqual(a1, a2) {
			t.Fatalf("expected true")
		}
	})
}

func TestDeepValueEqual_EdgeCases(t *testing.T) {
	t.Run("invalid values", func(t *testing.T) {
		var v1 reflect.Value
		v2 := reflect.ValueOf(42)
		// 一个 valid，一个 invalid
		if deepValueEqual(v1, v2) {
			t.Fatalf("expected false")
		}
		// 两个都 invalid
		if !deepValueEqual(v1, reflect.Value{}) {
			t.Fatalf("expected true")
		}
	})

	t.Run("nil vs non-nil map", func(t *testing.T) {
		var m1 map[string]int
		m2 := map[string]int{"a": 1}
		if DeepEqual(m1, m2) {
			t.Fatalf("expected false")
		}
	})

	t.Run("same map reference", func(t *testing.T) {
		m := map[string]int{"a": 1}
		if !DeepEqual(m, m) {
			t.Fatalf("expected true")
		}
	})

	t.Run("same slice reference", func(t *testing.T) {
		s := []int{1, 2, 3}
		if !DeepEqual(s, s) {
			t.Fatalf("expected true")
		}
	})

	t.Run("nil interface values", func(t *testing.T) {
		var i1 interface{}
		var i2 interface{}
		if !DeepEqual(i1, i2) {
			t.Fatalf("expected true")
		}
	})

	t.Run("nil vs non-nil interface", func(t *testing.T) {
		var i1 interface{}
		i2 := interface{}(42)
		if DeepEqual(i1, i2) {
			t.Fatalf("expected false")
		}
	})

	t.Run("slice with nil interface", func(t *testing.T) {
		var s1 []interface{}
		s1 = append(s1, nil)
		s2 := []interface{}{nil}
		if !DeepEqual(s1, s2) {
			t.Fatalf("expected true")
		}
	})
}

func TestDeepEqual_PointerOptimization(t *testing.T) {
	// 测试指针相同的情况
	s := []int{1, 2, 3}
	if !DeepEqual(s, s) {
		t.Fatalf("expected true for same slice reference")
	}

	m := map[string]int{"a": 1}
	if !DeepEqual(m, m) {
		t.Fatalf("expected true for same map reference")
	}
}

func TestDeepEqual_EmptyContainers(t *testing.T) {
	// 测试空容器
	if !DeepEqual([]int{}, []int{}) {
		t.Fatalf("expected true for empty slices")
	}
	if !DeepEqual(map[string]int{}, map[string]int{}) {
		t.Fatalf("expected true for empty maps")
	}
}

func TestDeepEqual_NilContainers(t *testing.T) {
	// 测试 nil 容器
	var s1, s2 []int
	if !DeepEqual(s1, s2) {
		t.Fatalf("expected true for nil slices")
	}

	var m1, m2 map[string]int
	if !DeepEqual(m1, m2) {
		t.Fatalf("expected true for nil maps")
	}
}

func TestDeepEqual_Uncomparable(t *testing.T) {
	// 测试不可比较的类型（函数）
	f1 := func() {}
	f2 := func() {}
	// 函数不能直接比较，DeepEqual 应该返回 false
	if DeepEqual(f1, f2) {
		t.Fatalf("expected false for functions")
	}
}

func TestDeepEqual_ComplexComparisons(t *testing.T) {
	// 测试复数比较
	c1 := complex(1, 2)
	c2 := complex(1, 2)
	if !DeepEqual(c1, c2) {
		t.Fatalf("expected true for equal complex numbers")
	}

	c3 := complex(1, 3)
	if DeepEqual(c1, c3) {
		t.Fatalf("expected false for different complex numbers")
	}
}

func TestDeepEqual_StructWithEmbedded(t *testing.T) {
	type embedded struct {
		X int
	}
	type outer struct {
		embedded
		Y int
	}
	o1 := outer{embedded: embedded{X: 1}, Y: 2}
	o2 := outer{embedded: embedded{X: 1}, Y: 2}
	if !DeepEqual(o1, o2) {
		t.Fatalf("expected true")
	}
}

func TestDeepEqual_SameReferenceOptimization(t *testing.T) {
	// 测试相同引用的优化
	s := []int{1, 2, 3}
	if !DeepEqual(s, s) {
		t.Fatalf("expected true for same slice")
	}

	m := map[string]int{"a": 1}
	if !DeepEqual(m, m) {
		t.Fatalf("expected true for same map")
	}

	// 测试不同内容的比较
	s2 := []int{1, 2, 4}
	if DeepEqual(s, s2) {
		t.Fatalf("expected false for different slices")
	}
}

func TestDeepEqual_SliceElements(t *testing.T) {
	// 测试切片元素的深度比较
	s1 := []interface{}{1, "a", true}
	s2 := []interface{}{1, "a", true}
	if !DeepEqual(s1, s2) {
		t.Fatalf("expected true")
	}

	s3 := []interface{}{1, "a", false}
	if DeepEqual(s1, s3) {
		t.Fatalf("expected false")
	}
}

func TestDeepEqual_MoreComplexTypes(t *testing.T) {
	// 测试更多复杂类型的比较

	// 嵌套 map
	m1 := map[string]map[string]int{"a": {"x": 1}}
	m2 := map[string]map[string]int{"a": {"x": 1}}
	if !DeepEqual(m1, m2) {
		t.Fatalf("expected true for nested maps")
	}

	// map 的 slice
	s1 := []map[string]int{{"a": 1}, {"b": 2}}
	s2 := []map[string]int{{"a": 1}, {"b": 2}}
	if !DeepEqual(s1, s2) {
		t.Fatalf("expected true for slice of maps")
	}

	// 结构体 slice
	type testStruct struct {
		X int
		Y string
	}
	st1 := []testStruct{{X: 1, Y: "a"}}
	st2 := []testStruct{{X: 1, Y: "a"}}
	if !DeepEqual(st1, st2) {
		t.Fatalf("expected true for struct slices")
	}
}

func TestDeepEqual_MixedTypes(t *testing.T) {
	// 测试混合类型的 interface{}
	i1 := []interface{}{1, "a", true, 1.5}
	i2 := []interface{}{1, "a", true, 1.5}
	if !DeepEqual(i1, i2) {
		t.Fatalf("expected true for mixed types")
	}

	i3 := []interface{}{1, "a", false, 1.5}
	if DeepEqual(i1, i3) {
		t.Fatalf("expected false for different mixed types")
	}
}

func TestDeepEqual_DeepNesting(t *testing.T) {
	// 测试深度嵌套结构
	type deep struct {
		Data map[string][]int
	}
	src := []deep{
		{Data: map[string][]int{"a": {1, 2}, "b": {3}}},
		{Data: map[string][]int{"c": {4}}},
	}
	dst := []deep{
		{Data: map[string][]int{"a": {1, 2}, "b": {3}}},
		{Data: map[string][]int{"c": {4}}},
	}
	if !DeepEqual(src, dst) {
		t.Fatalf("expected true for deep nested structures")
	}

	// 修改内容后应该不相等
	dst2 := []deep{
		{Data: map[string][]int{"a": {1, 2}, "b": {3}}},
		{Data: map[string][]int{"c": {5}}},
	}
	if DeepEqual(src, dst2) {
		t.Fatalf("expected false for different deep nested structures")
	}
}

func TestDeepEqual_PointerEquality(t *testing.T) {
	// 测试指针相等情况
	p := &pluckPerson{Name: "a"}
	if !DeepEqual(p, p) {
		t.Fatalf("expected true for same pointer")
	}

	// 不同指针但内容相同
	p2 := &pluckPerson{Name: "a"}
	if !DeepEqual(p, p2) {
		t.Fatalf("expected true for different pointers with same content")
	}

	// 内容不同
	p3 := &pluckPerson{Name: "b"}
	if DeepEqual(p, p3) {
		t.Fatalf("expected false for different pointers with different content")
	}
}

func TestDeepEqual_SpecialCases(t *testing.T) {
	// 测试特殊情况

	// 空结构体
	type empty struct{}
	e1, e2 := empty{}, empty{}
	if !DeepEqual(e1, e2) {
		t.Fatalf("expected true for empty structs")
	}

	// 包含 nil 的 slice
	s1 := []*int{nil, &[]int{1}[0]}
	s2 := []*int{nil, &[]int{1}[0]}
	if !DeepEqual(s1, s2) {
		t.Fatalf("expected true for slice with nil pointer")
	}

	// 复数
	c1 := complex(1.5, 2.5)
	c2 := complex(1.5, 2.5)
	if !DeepEqual(c1, c2) {
		t.Fatalf("expected true for equal complex numbers")
	}

	// 字符串指针
	str := "test"
	if !DeepEqual(&str, &str) {
		t.Fatalf("expected true for same string pointer")
	}
}

func TestDeepValueEqual_AllKinds(t *testing.T) {
	// 测试所有基本类型
	tests := []struct {
		name string
		v1   interface{}
		v2   interface{}
	}{
		{"int", 1, 1},
		{"int8", int8(1), int8(1)},
		{"int16", int16(1), int16(1)},
		{"int32", int32(1), int32(1)},
		{"int64", int64(1), int64(1)},
		{"uint", uint(1), uint(1)},
		{"uint8", uint8(1), uint8(1)},
		{"uint16", uint16(1), uint16(1)},
		{"uint32", uint32(1), uint32(1)},
		{"uint64", uint64(1), uint64(1)},
		{"float32", float32(1.5), float32(1.5)},
		{"float64", float64(1.5), float64(1.5)},
		{"string", "a", "a"},
		{"bool", true, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if !DeepEqual(tt.v1, tt.v2) {
				t.Fatalf("expected true")
			}
		})
	}
}

func TestDeepValueEqual_NilChecks(t *testing.T) {
	// 测试 nil 检查的所有情况
	var nilSlice []int
	var nilMap map[string]int
	var nilPtr *int
	var nilInterface interface{}

	// 两个 nil 值应该相等
	if !DeepEqual(nilSlice, nilSlice) {
		t.Fatalf("expected true for nil slices")
	}
	if !DeepEqual(nilMap, nilMap) {
		t.Fatalf("expected true for nil maps")
	}
	if !DeepEqual(nilPtr, nilPtr) {
		t.Fatalf("expected true for nil pointers")
	}
	if !DeepEqual(nilInterface, nilInterface) {
		t.Fatalf("expected true for nil interfaces")
	}

	// nil vs 非-nil 应该不相等
	notEmpty := []int{1}
	if DeepEqual(nilSlice, notEmpty) {
		t.Fatalf("expected false for nil vs non-nil slice")
	}
}

func TestDeepValueEqual_Comparisons(t *testing.T) {
	// 测试各种比较情况

	// map 的元素比较
	m1 := map[string][]int{"a": {1, 2}, "b": {3, 4}}
	m2 := map[string][]int{"a": {1, 2}, "b": {3, 4}}
	if !DeepEqual(m1, m2) {
		t.Fatalf("expected true for maps with slice values")
	}

	// slice 的元素比较
	s1 := []map[string]int{{"a": 1}, {"b": 2}}
	s2 := []map[string]int{{"a": 1}, {"b": 2}}
	if !DeepEqual(s1, s2) {
		t.Fatalf("expected true for slices with map values")
	}

	// 嵌套指针比较
	val := 42
	p1 := &val
	p2 := &p1
	p3 := &p1
	if !DeepEqual(p2, p3) {
		t.Fatalf("expected true for same nested pointer")
	}
}

func TestDeepEqual_AllTypes(t *testing.T) {
	// 测试所有类型的比较

	// 复数
	c1 := complex(1, 2)
	c2 := complex(1, 2)
	c3 := complex(2, 3)
	if !DeepEqual(c1, c2) {
		t.Fatalf("expected true for equal complex")
	}
	if DeepEqual(c1, c3) {
		t.Fatalf("expected false for different complex")
	}

	// 布尔值
	if !DeepEqual(true, true) {
		t.Fatalf("expected true for equal bool")
	}
	if DeepEqual(true, false) {
		t.Fatalf("expected false for different bool")
	}

	// 字符串
	if !DeepEqual("test", "test") {
		t.Fatalf("expected true for equal strings")
	}
	if DeepEqual("test", "other") {
		t.Fatalf("expected false for different strings")
	}
}

func TestDeepValueEqual_RemainingCases(t *testing.T) {
	// 测试剩余未覆盖的情况

	t.Run("slice with nil elements", func(t *testing.T) {
		s1 := []*int{nil, &[]int{1}[0], nil}
		s2 := []*int{nil, &[]int{1}[0], nil}
		if !DeepEqual(s1, s2) {
			t.Fatalf("expected true")
		}

		// 改变一个元素
		s3 := []*int{nil, &[]int{2}[0], nil}
		if DeepEqual(s1, s3) {
			t.Fatalf("expected false")
		}
	})

	t.Run("map with nil key values", func(t *testing.T) {
		m1 := map[string]*int{"a": nil, "b": &[]int{1}[0]}
		m2 := map[string]*int{"a": nil, "b": &[]int{1}[0]}
		if !DeepEqual(m1, m2) {
			t.Fatalf("expected true")
		}

		m3 := map[string]*int{"a": &[]int{1}[0]}
		if DeepEqual(m1, m3) {
			t.Fatalf("expected false")
		}
	})

	t.Run("deeply nested structures", func(t *testing.T) {
		type level3 struct{ X int }
		type level2 struct{ L3 level3 }
		type level1 struct{ L2 level2 }
		l1 := level1{L2: level2{L3: level3{X: 1}}}
		l2 := level1{L2: level2{L3: level3{X: 1}}}
		if !DeepEqual(l1, l2) {
			t.Fatalf("expected true")
		}

		l3 := level1{L2: level2{L3: level3{X: 2}}}
		if DeepEqual(l1, l3) {
			t.Fatalf("expected false")
		}
	})
}

func TestDeepValueEqual_EdgeCases2(t *testing.T) {
	// 测试更多边缘情况

	t.Run("maps with complex keys", func(t *testing.T) {
		type keyStruct struct {
			X int
			Y int
		}
		m1 := map[keyStruct]string{{1, 2}: "a"}
		m2 := map[keyStruct]string{{1, 2}: "a"}
		if !DeepEqual(m1, m2) {
			t.Fatalf("expected true")
		}

		m3 := map[keyStruct]string{{1, 2}: "b"}
		if DeepEqual(m1, m3) {
			t.Fatalf("expected false")
		}
	})

	t.Run("arrays of pointers", func(t *testing.T) {
		vals := []int{1, 2, 3}
		a1 := [3]*int{&vals[0], &vals[1], &vals[2]}
		a2 := [3]*int{&vals[0], &vals[1], &vals[2]}
		if !DeepEqual(a1, a2) {
			t.Fatalf("expected true for same pointer array")
		}
	})

	t.Run("nested interface comparisons", func(t *testing.T) {
		type inner struct {
			X int
		}
		i1 := interface{}(inner{X: 1})
		i2 := interface{}(inner{X: 1})
		if !DeepEqual(i1, i2) {
			t.Fatalf("expected true")
		}

		i3 := interface{}(inner{X: 2})
		if DeepEqual(i1, i3) {
			t.Fatalf("expected false")
		}
	})

	t.Run("mixed type slices", func(t *testing.T) {
		s1 := []interface{}{int(1), int8(2), int16(3), int32(4), int64(5)}
		s2 := []interface{}{int(1), int8(2), int16(3), int32(4), int64(5)}
		if !DeepEqual(s1, s2) {
			t.Fatalf("expected true")
		}
	})
}

func TestDeepValueEqual_AllComparisons(t *testing.T) {
	// 测试所有可能的比较情况

	t.Run("empty vs non-empty", func(t *testing.T) {
		if DeepEqual([]int{}, []int{1}) {
			t.Fatalf("expected false for empty vs non-empty slice")
		}
		if DeepEqual(map[string]int{}, map[string]int{"a": 1}) {
			t.Fatalf("expected false for empty vs non-empty map")
		}
	})

	t.Run("different length containers", func(t *testing.T) {
		s1 := []int{1, 2, 3}
		s2 := []int{1, 2}
		if DeepEqual(s1, s2) {
			t.Fatalf("expected false for different length slices")
		}

		m1 := map[string]int{"a": 1, "b": 2}
		m2 := map[string]int{"a": 1}
		if DeepEqual(m1, m2) {
			t.Fatalf("expected false for different size maps")
		}
	})
	t.Run("pointers to pointers", func(t *testing.T) {
		val := 42
		p1 := &val
		p2 := &p1
		p3 := &p1

		// 相同的指针链
		if !DeepEqual(p2, p3) {
			t.Fatalf("expected true")
		}

		// 不同的指针链，指向相同的值
		val2 := 42
		pp4 := &val2
		p4 := &pp4
		if !DeepEqual(p2, p4) {
			t.Fatalf("expected true for different pointers with same value")
		}
	})
}

func TestDeepValueEqual_SpecialCases(t *testing.T) {
	// 测试特殊情况

	t.Run("nested pointers with nil", func(t *testing.T) {
		var p1 *int
		var p2 **int
		var p3 ***int

		// 所有都是nil，应该相等
		if !DeepEqual(p1, p1) {
			t.Fatalf("expected true for nil pointers")
		}
		if !DeepEqual(p2, p2) {
			t.Fatalf("expected true for nil double pointers")
		}
		if !DeepEqual(p3, p3) {
			t.Fatalf("expected true for nil triple pointers")
		}
	})

	t.Run("interface with different concrete types", func(t *testing.T) {
		i1 := interface{}(int(1))
		i2 := interface{}("1")
		if DeepEqual(i1, i2) {
			t.Fatalf("expected false for different types")
		}
	})

	t.Run("complex nested structures", func(t *testing.T) {
		type inner struct {
			Data []int
		}
		type middle struct {
			Items []inner
		}
		type outer struct {
			Middle middle
		}

		o1 := outer{Middle: middle{Items: []inner{{Data: []int{1, 2}}}}}
		o2 := outer{Middle: middle{Items: []inner{{Data: []int{1, 2}}}}}
		if !DeepEqual(o1, o2) {
			t.Fatalf("expected true")
		}
	})
}
