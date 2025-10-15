package candy

import (
	"testing"
)

// TestDeepCopy 测试 DeepCopy 函数
func TestDeepCopy(t *testing.T) {
	t.Run("copy basic types", func(t *testing.T) {
		// int
		var srcInt = 42
		var dstInt int
		DeepCopy(srcInt, &dstInt)
		if dstInt != 42 {
			t.Errorf("DeepCopy int failed: got %d, want 42", dstInt)
		}

		// string
		var srcStr = "hello"
		var dstStr string
		DeepCopy(srcStr, &dstStr)
		if dstStr != "hello" {
			t.Errorf("DeepCopy string failed: got %s, want hello", dstStr)
		}

		// bool
		var srcBool = true
		var dstBool bool
		DeepCopy(srcBool, &dstBool)
		if !dstBool {
			t.Errorf("DeepCopy bool failed: got false, want true")
		}

		// float64
		var srcFloat = 3.14
		var dstFloat float64
		DeepCopy(srcFloat, &dstFloat)
		if dstFloat != 3.14 {
			t.Errorf("DeepCopy float64 failed: got %f, want 3.14", dstFloat)
		}

		// complex128
		var srcComplex = complex(1, 2)
		var dstComplex complex128
		DeepCopy(srcComplex, &dstComplex)
		if dstComplex != complex(1, 2) {
			t.Errorf("DeepCopy complex128 failed: got %v, want (1+2i)", dstComplex)
		}
	})

	t.Run("copy slice", func(t *testing.T) {
		src := []int{1, 2, 3, 4, 5}
		var dst []int
		DeepCopy(src, &dst)

		if len(dst) != len(src) {
			t.Errorf("DeepCopy slice length mismatch: got %d, want %d", len(dst), len(src))
		}
		for i := range src {
			if dst[i] != src[i] {
				t.Errorf("DeepCopy slice[%d] = %d, want %d", i, dst[i], src[i])
			}
		}

		// Verify independence
		src[0] = 999
		if dst[0] == 999 {
			t.Errorf("DeepCopy slice not independent")
		}
	})

	t.Run("copy nil slice", func(t *testing.T) {
		var src []int
		var dst []int
		DeepCopy(src, &dst)
		if dst != nil {
			t.Errorf("DeepCopy nil slice should result in nil, got %v", dst)
		}
	})

	t.Run("copy map", func(t *testing.T) {
		src := map[string]int{"a": 1, "b": 2, "c": 3}
		var dst map[string]int
		DeepCopy(src, &dst)

		if len(dst) != len(src) {
			t.Errorf("DeepCopy map length mismatch: got %d, want %d", len(dst), len(src))
		}
		for k, v := range src {
			if dst[k] != v {
				t.Errorf("DeepCopy map[%s] = %d, want %d", k, dst[k], v)
			}
		}

		// Verify independence
		src["a"] = 999
		if dst["a"] == 999 {
			t.Errorf("DeepCopy map not independent")
		}
	})

	t.Run("copy nil map", func(t *testing.T) {
		var src map[string]int
		var dst map[string]int
		DeepCopy(src, &dst)
		if dst != nil {
			t.Errorf("DeepCopy nil map should result in nil, got %v", dst)
		}
	})

	t.Run("copy struct", func(t *testing.T) {
		type Person struct {
			Name string
			Age  int
		}
		src := Person{Name: "Alice", Age: 30}
		var dst Person
		DeepCopy(src, &dst)

		if dst.Name != "Alice" || dst.Age != 30 {
			t.Errorf("DeepCopy struct failed: got %+v, want {Alice 30}", dst)
		}
	})

	t.Run("copy nested struct", func(t *testing.T) {
		type Address struct {
			City string
		}
		type Person struct {
			Name    string
			Address Address
		}
		src := Person{Name: "Bob", Address: Address{City: "NYC"}}
		var dst Person
		DeepCopy(src, &dst)

		if dst.Name != "Bob" || dst.Address.City != "NYC" {
			t.Errorf("DeepCopy nested struct failed: got %+v", dst)
		}
	})

	t.Run("copy pointer", func(t *testing.T) {
		val := 42
		src := &val
		var dst *int
		DeepCopy(src, &dst)

		if dst == nil || *dst != 42 {
			t.Errorf("DeepCopy pointer failed")
		}

		// Verify independence
		*src = 999
		if *dst == 999 {
			t.Errorf("DeepCopy pointer not independent")
		}
	})

	t.Run("copy nil pointer", func(t *testing.T) {
		var src *int
		var dst *int
		DeepCopy(src, &dst)
		if dst != nil {
			t.Errorf("DeepCopy nil pointer should result in nil")
		}
	})

	t.Run("copy array", func(t *testing.T) {
		src := [3]int{1, 2, 3}
		var dst [3]int
		DeepCopy(src, &dst)

		for i := range src {
			if dst[i] != src[i] {
				t.Errorf("DeepCopy array[%d] = %d, want %d", i, dst[i], src[i])
			}
		}
	})

	t.Run("copy interface through pointer", func(t *testing.T) {
		// 使用指向接口的指针来复制
		type Container struct {
			Value interface{}
		}
		src := Container{Value: 42}
		var dst Container
		DeepCopy(src, &dst)

		if dst.Value != 42 {
			t.Errorf("DeepCopy interface failed: got %v, want 42", dst.Value)
		}
	})

	t.Run("copy nil interface in struct", func(t *testing.T) {
		type Container struct {
			Value interface{}
		}
		src := Container{Value: nil}
		var dst Container
		DeepCopy(src, &dst)
		if dst.Value != nil {
			t.Errorf("DeepCopy nil interface should result in nil")
		}
	})

	t.Run("copy slice of structs", func(t *testing.T) {
		type Item struct {
			ID   int
			Name string
		}
		src := []Item{{1, "a"}, {2, "b"}}
		var dst []Item
		DeepCopy(src, &dst)

		if len(dst) != 2 || dst[0].ID != 1 || dst[1].Name != "b" {
			t.Errorf("DeepCopy slice of structs failed: got %+v", dst)
		}
	})

	t.Run("copy map of slices", func(t *testing.T) {
		src := map[string][]int{"a": {1, 2}, "b": {3, 4}}
		var dst map[string][]int
		DeepCopy(src, &dst)

		if len(dst) != 2 || len(dst["a"]) != 2 || dst["b"][1] != 4 {
			t.Errorf("DeepCopy map of slices failed: got %+v", dst)
		}
	})
}

// TestDeepEqual 测试 DeepEqual 函数
func TestDeepEqual(t *testing.T) {
	t.Run("equal basic types", func(t *testing.T) {
		if !DeepEqual(42, 42) {
			t.Errorf("DeepEqual(42, 42) should be true")
		}
		if DeepEqual(42, 43) {
			t.Errorf("DeepEqual(42, 43) should be false")
		}
		if !DeepEqual("hello", "hello") {
			t.Errorf("DeepEqual strings should be true")
		}
		if !DeepEqual(true, true) {
			t.Errorf("DeepEqual bools should be true")
		}
		if !DeepEqual(3.14, 3.14) {
			t.Errorf("DeepEqual floats should be true")
		}
	})

	t.Run("equal slices", func(t *testing.T) {
		a := []int{1, 2, 3}
		b := []int{1, 2, 3}
		c := []int{1, 2, 4}

		if !DeepEqual(a, b) {
			t.Errorf("DeepEqual equal slices should be true")
		}
		if DeepEqual(a, c) {
			t.Errorf("DeepEqual different slices should be false")
		}
		if DeepEqual(a, []int{1, 2}) {
			t.Errorf("DeepEqual slices with different lengths should be false")
		}
	})

	t.Run("equal nil slices", func(t *testing.T) {
		var a []int
		var b []int
		if !DeepEqual(a, b) {
			t.Errorf("DeepEqual nil slices should be true")
		}

		c := []int{}
		if DeepEqual(a, c) {
			t.Errorf("DeepEqual nil vs empty slice should be false")
		}
	})

	t.Run("equal maps", func(t *testing.T) {
		a := map[string]int{"a": 1, "b": 2}
		b := map[string]int{"a": 1, "b": 2}
		c := map[string]int{"a": 1, "b": 3}

		if !DeepEqual(a, b) {
			t.Errorf("DeepEqual equal maps should be true")
		}
		if DeepEqual(a, c) {
			t.Errorf("DeepEqual different maps should be false")
		}
	})

	t.Run("equal nil maps", func(t *testing.T) {
		var a map[string]int
		var b map[string]int
		if !DeepEqual(a, b) {
			t.Errorf("DeepEqual nil maps should be true")
		}

		c := map[string]int{}
		if DeepEqual(a, c) {
			t.Errorf("DeepEqual nil vs empty map should be false")
		}
	})

	t.Run("equal structs", func(t *testing.T) {
		type Person struct {
			Name string
			Age  int
		}
		a := Person{"Alice", 30}
		b := Person{"Alice", 30}
		c := Person{"Bob", 30}

		if !DeepEqual(a, b) {
			t.Errorf("DeepEqual equal structs should be true")
		}
		if DeepEqual(a, c) {
			t.Errorf("DeepEqual different structs should be false")
		}
	})

	t.Run("equal pointers", func(t *testing.T) {
		val1 := 42
		val2 := 42
		a := &val1
		b := &val2

		if !DeepEqual(a, b) {
			t.Errorf("DeepEqual pointers to equal values should be true")
		}

		val3 := 43
		c := &val3
		if DeepEqual(a, c) {
			t.Errorf("DeepEqual pointers to different values should be false")
		}
	})

	t.Run("equal nil pointers", func(t *testing.T) {
		var a *int
		var b *int
		if !DeepEqual(a, b) {
			t.Errorf("DeepEqual nil pointers should be true")
		}

		val := 42
		c := &val
		if DeepEqual(a, c) {
			t.Errorf("DeepEqual nil vs non-nil pointer should be false")
		}
	})

	t.Run("equal arrays", func(t *testing.T) {
		a := [3]int{1, 2, 3}
		b := [3]int{1, 2, 3}
		c := [3]int{1, 2, 4}

		if !DeepEqual(a, b) {
			t.Errorf("DeepEqual equal arrays should be true")
		}
		if DeepEqual(a, c) {
			t.Errorf("DeepEqual different arrays should be false")
		}
	})

	t.Run("equal interfaces in structs", func(t *testing.T) {
		type Container struct {
			Value interface{}
		}
		a := Container{Value: 42}
		b := Container{Value: 42}
		c := Container{Value: 43}

		if !DeepEqual(a, b) {
			t.Errorf("DeepEqual equal interfaces should be true")
		}
		if DeepEqual(a, c) {
			t.Errorf("DeepEqual different interfaces should be false")
		}
	})

	t.Run("equal nil interfaces in structs", func(t *testing.T) {
		type Container struct {
			Value interface{}
		}
		a := Container{Value: nil}
		b := Container{Value: nil}
		if !DeepEqual(a, b) {
			t.Errorf("DeepEqual nil interfaces should be true")
		}
	})

	t.Run("same slice reference", func(t *testing.T) {
		a := []int{1, 2, 3}
		if !DeepEqual(a, a) {
			t.Errorf("DeepEqual same slice reference should be true")
		}
	})

	t.Run("same map reference", func(t *testing.T) {
		a := map[string]int{"a": 1}
		if !DeepEqual(a, a) {
			t.Errorf("DeepEqual same map reference should be true")
		}
	})
}

// TestTypedSliceCopy 测试 TypedSliceCopy 函数
func TestTypedSliceCopy(t *testing.T) {
	t.Run("copy basic type slice", func(t *testing.T) {
		src := []int{1, 2, 3, 4, 5}
		dst := TypedSliceCopy(src)

		if len(dst) != len(src) {
			t.Errorf("TypedSliceCopy length mismatch")
		}
		for i := range src {
			if dst[i] != src[i] {
				t.Errorf("TypedSliceCopy[%d] = %d, want %d", i, dst[i], src[i])
			}
		}

		// Verify independence
		src[0] = 999
		if dst[0] == 999 {
			t.Errorf("TypedSliceCopy not independent")
		}
	})

	t.Run("copy nil slice", func(t *testing.T) {
		var src []int
		dst := TypedSliceCopy(src)
		if dst != nil {
			t.Errorf("TypedSliceCopy nil slice should return nil")
		}
	})

	t.Run("copy empty slice", func(t *testing.T) {
		src := []int{}
		dst := TypedSliceCopy(src)
		if len(dst) != 0 {
			t.Errorf("TypedSliceCopy empty slice should return empty slice")
		}
	})

	t.Run("copy struct slice", func(t *testing.T) {
		type Item struct {
			ID   int
			Name string
		}
		src := []Item{{1, "a"}, {2, "b"}}
		dst := TypedSliceCopy(src)

		if len(dst) != 2 || dst[0].ID != 1 || dst[1].Name != "b" {
			t.Errorf("TypedSliceCopy struct slice failed")
		}
	})

	t.Run("copy string slice", func(t *testing.T) {
		src := []string{"a", "b", "c"}
		dst := TypedSliceCopy(src)

		for i := range src {
			if dst[i] != src[i] {
				t.Errorf("TypedSliceCopy string[%d] mismatch", i)
			}
		}
	})
}

// TestTypedMapCopy 测试 TypedMapCopy 函数
func TestTypedMapCopy(t *testing.T) {
	t.Run("copy basic type map", func(t *testing.T) {
		src := map[string]int{"a": 1, "b": 2, "c": 3}
		dst := TypedMapCopy(src)

		if len(dst) != len(src) {
			t.Errorf("TypedMapCopy length mismatch")
		}
		for k, v := range src {
			if dst[k] != v {
				t.Errorf("TypedMapCopy[%s] = %d, want %d", k, dst[k], v)
			}
		}

		// Verify independence
		src["a"] = 999
		if dst["a"] == 999 {
			t.Errorf("TypedMapCopy not independent")
		}
	})

	t.Run("copy nil map", func(t *testing.T) {
		var src map[string]int
		dst := TypedMapCopy(src)
		if dst != nil {
			t.Errorf("TypedMapCopy nil map should return nil")
		}
	})

	t.Run("copy empty map", func(t *testing.T) {
		src := map[string]int{}
		dst := TypedMapCopy(src)
		if len(dst) != 0 {
			t.Errorf("TypedMapCopy empty map should return empty map")
		}
	})

	t.Run("copy struct value map", func(t *testing.T) {
		type Item struct {
			ID int
		}
		src := map[string]Item{"a": {1}, "b": {2}}
		dst := TypedMapCopy(src)

		if len(dst) != 2 || dst["a"].ID != 1 {
			t.Errorf("TypedMapCopy struct value map failed")
		}
	})
}

// TestGenericSliceEqual 测试 GenericSliceEqual 函数
func TestGenericSliceEqual(t *testing.T) {
	t.Run("equal slices", func(t *testing.T) {
		a := []int{1, 2, 3}
		b := []int{1, 2, 3}
		if !GenericSliceEqual(a, b) {
			t.Errorf("GenericSliceEqual equal slices should be true")
		}
	})

	t.Run("different slices", func(t *testing.T) {
		a := []int{1, 2, 3}
		b := []int{1, 2, 4}
		if GenericSliceEqual(a, b) {
			t.Errorf("GenericSliceEqual different slices should be false")
		}
	})

	t.Run("different lengths", func(t *testing.T) {
		a := []int{1, 2, 3}
		b := []int{1, 2}
		if GenericSliceEqual(a, b) {
			t.Errorf("GenericSliceEqual different lengths should be false")
		}
	})

	t.Run("empty slices", func(t *testing.T) {
		a := []int{}
		b := []int{}
		if !GenericSliceEqual(a, b) {
			t.Errorf("GenericSliceEqual empty slices should be true")
		}
	})

	t.Run("same reference", func(t *testing.T) {
		a := []int{1, 2, 3}
		if !GenericSliceEqual(a, a) {
			t.Errorf("GenericSliceEqual same reference should be true")
		}
	})

	t.Run("string slices", func(t *testing.T) {
		a := []string{"a", "b", "c"}
		b := []string{"a", "b", "c"}
		if !GenericSliceEqual(a, b) {
			t.Errorf("GenericSliceEqual string slices should be true")
		}
	})
}

// TestMapEqual 测试 MapEqual 函数
func TestMapEqual(t *testing.T) {
	t.Run("equal maps", func(t *testing.T) {
		a := map[string]int{"a": 1, "b": 2}
		b := map[string]int{"a": 1, "b": 2}
		if !MapEqual(a, b) {
			t.Errorf("MapEqual equal maps should be true")
		}
	})

	t.Run("different values", func(t *testing.T) {
		a := map[string]int{"a": 1, "b": 2}
		b := map[string]int{"a": 1, "b": 3}
		if MapEqual(a, b) {
			t.Errorf("MapEqual different values should be false")
		}
	})

	t.Run("different keys", func(t *testing.T) {
		a := map[string]int{"a": 1, "b": 2}
		b := map[string]int{"a": 1, "c": 2}
		if MapEqual(a, b) {
			t.Errorf("MapEqual different keys should be false")
		}
	})

	t.Run("different lengths", func(t *testing.T) {
		a := map[string]int{"a": 1, "b": 2}
		b := map[string]int{"a": 1}
		if MapEqual(a, b) {
			t.Errorf("MapEqual different lengths should be false")
		}
	})

	t.Run("empty maps", func(t *testing.T) {
		a := map[string]int{}
		b := map[string]int{}
		if !MapEqual(a, b) {
			t.Errorf("MapEqual empty maps should be true")
		}
	})
}

// TestPointerEqual 测试 PointerEqual 函数
func TestPointerEqual(t *testing.T) {
	t.Run("equal pointers", func(t *testing.T) {
		val1 := 42
		val2 := 42
		a := &val1
		b := &val2
		if !PointerEqual(a, b) {
			t.Errorf("PointerEqual equal values should be true")
		}
	})

	t.Run("different pointers", func(t *testing.T) {
		val1 := 42
		val2 := 43
		a := &val1
		b := &val2
		if PointerEqual(a, b) {
			t.Errorf("PointerEqual different values should be false")
		}
	})

	t.Run("both nil", func(t *testing.T) {
		var a *int
		var b *int
		if !PointerEqual(a, b) {
			t.Errorf("PointerEqual both nil should be true")
		}
	})

	t.Run("one nil", func(t *testing.T) {
		val := 42
		a := &val
		var b *int
		if PointerEqual(a, b) {
			t.Errorf("PointerEqual one nil should be false")
		}
	})
}

// TestStructEqual 测试 StructEqual 函数
func TestStructEqual(t *testing.T) {
	type Person struct {
		Name string
		Age  int
	}

	comparer := func(a, b Person) bool {
		return a.Name == b.Name && a.Age == b.Age
	}

	t.Run("equal structs", func(t *testing.T) {
		a := Person{"Alice", 30}
		b := Person{"Alice", 30}
		if !StructEqual(a, b, comparer) {
			t.Errorf("StructEqual equal structs should be true")
		}
	})

	t.Run("different structs", func(t *testing.T) {
		a := Person{"Alice", 30}
		b := Person{"Bob", 30}
		if StructEqual(a, b, comparer) {
			t.Errorf("StructEqual different structs should be false")
		}
	})
}

// TestClone 测试 Clone 函数
func TestClone(t *testing.T) {
	t.Run("clone int", func(t *testing.T) {
		src := 42
		dst := Clone(src)
		if dst != 42 {
			t.Errorf("Clone int failed: got %d, want 42", dst)
		}
	})

	t.Run("clone string", func(t *testing.T) {
		src := "hello"
		dst := Clone(src)
		if dst != "hello" {
			t.Errorf("Clone string failed: got %s, want hello", dst)
		}
	})

	t.Run("clone struct", func(t *testing.T) {
		type Person struct {
			Name string
			Age  int
		}
		src := Person{"Alice", 30}
		dst := Clone(src)
		if dst.Name != "Alice" || dst.Age != 30 {
			t.Errorf("Clone struct failed: got %+v", dst)
		}
	})
}

// TestCloneSlice 测试 CloneSlice 函数
func TestCloneSlice(t *testing.T) {
	t.Run("clone int slice", func(t *testing.T) {
		src := []int{1, 2, 3, 4, 5}
		dst := CloneSlice(src)

		if len(dst) != len(src) {
			t.Errorf("CloneSlice length mismatch")
		}
		for i := range src {
			if dst[i] != src[i] {
				t.Errorf("CloneSlice[%d] mismatch", i)
			}
		}

		// Verify independence
		src[0] = 999
		if dst[0] == 999 {
			t.Errorf("CloneSlice not independent")
		}
	})

	t.Run("clone nil slice", func(t *testing.T) {
		var src []int
		dst := CloneSlice(src)
		if dst != nil {
			t.Errorf("CloneSlice nil should return nil")
		}
	})
}

// TestCloneMap 测试 CloneMap 函数
func TestCloneMap(t *testing.T) {
	t.Run("clone map", func(t *testing.T) {
		src := map[string]int{"a": 1, "b": 2}
		dst := CloneMap(src)

		if len(dst) != len(src) {
			t.Errorf("CloneMap length mismatch")
		}
		for k, v := range src {
			if dst[k] != v {
				t.Errorf("CloneMap[%s] mismatch", k)
			}
		}

		// Verify independence
		src["a"] = 999
		if dst["a"] == 999 {
			t.Errorf("CloneMap not independent")
		}
	})

	t.Run("clone nil map", func(t *testing.T) {
		var src map[string]int
		dst := CloneMap(src)
		if dst != nil {
			t.Errorf("CloneMap nil should return nil")
		}
	})
}

// TestEqual 测试 Equal 函数
func TestEqual(t *testing.T) {
	t.Run("equal ints", func(t *testing.T) {
		if !Equal(42, 42) {
			t.Errorf("Equal(42, 42) should be true")
		}
		if Equal(42, 43) {
			t.Errorf("Equal(42, 43) should be false")
		}
	})

	t.Run("equal strings", func(t *testing.T) {
		if !Equal("hello", "hello") {
			t.Errorf("Equal strings should be true")
		}
		if Equal("hello", "world") {
			t.Errorf("Equal different strings should be false")
		}
	})
}

// TestEqualSlice 测试 EqualSlice 函数
func TestEqualSlice(t *testing.T) {
	t.Run("equal slices", func(t *testing.T) {
		a := []int{1, 2, 3}
		b := []int{1, 2, 3}
		if !EqualSlice(a, b) {
			t.Errorf("EqualSlice equal slices should be true")
		}
	})

	t.Run("different slices", func(t *testing.T) {
		a := []int{1, 2, 3}
		b := []int{1, 2, 4}
		if EqualSlice(a, b) {
			t.Errorf("EqualSlice different slices should be false")
		}
	})
}

// TestEqualMap 测试 EqualMap 函数
func TestEqualMap(t *testing.T) {
	t.Run("equal maps", func(t *testing.T) {
		a := map[string]int{"a": 1, "b": 2}
		b := map[string]int{"a": 1, "b": 2}
		if !EqualMap(a, b) {
			t.Errorf("EqualMap equal maps should be true")
		}
	})

	t.Run("different maps", func(t *testing.T) {
		a := map[string]int{"a": 1, "b": 2}
		b := map[string]int{"a": 1, "b": 3}
		if EqualMap(a, b) {
			t.Errorf("EqualMap different maps should be false")
		}
	})
}

// TestDeepCopyEdgeCases 测试边界情况
func TestDeepCopyEdgeCases(t *testing.T) {
	t.Run("copy uint types", func(t *testing.T) {
		var srcUint uint = 42
		var dstUint uint
		DeepCopy(srcUint, &dstUint)
		if dstUint != 42 {
			t.Errorf("DeepCopy uint failed")
		}

		var srcUint8 uint8 = 255
		var dstUint8 uint8
		DeepCopy(srcUint8, &dstUint8)
		if dstUint8 != 255 {
			t.Errorf("DeepCopy uint8 failed")
		}
	})

	t.Run("copy int types", func(t *testing.T) {
		var srcInt8 int8 = -128
		var dstInt8 int8
		DeepCopy(srcInt8, &dstInt8)
		if dstInt8 != -128 {
			t.Errorf("DeepCopy int8 failed")
		}

		var srcInt16 int16 = -32768
		var dstInt16 int16
		DeepCopy(srcInt16, &dstInt16)
		if dstInt16 != -32768 {
			t.Errorf("DeepCopy int16 failed")
		}
	})

	t.Run("copy float types", func(t *testing.T) {
		var srcFloat32 float32 = 3.14
		var dstFloat32 float32
		DeepCopy(srcFloat32, &dstFloat32)
		if dstFloat32 != 3.14 {
			t.Errorf("DeepCopy float32 failed")
		}
	})

	t.Run("copy complex types", func(t *testing.T) {
		var srcComplex64 complex64 = complex(1, 2)
		var dstComplex64 complex64
		DeepCopy(srcComplex64, &dstComplex64)
		if dstComplex64 != complex(1, 2) {
			t.Errorf("DeepCopy complex64 failed")
		}
	})

	t.Run("copy nested pointers", func(t *testing.T) {
		val := 42
		ptr := &val
		src := &ptr
		var dst **int
		DeepCopy(src, &dst)

		if dst == nil || *dst == nil || **dst != 42 {
			t.Errorf("DeepCopy nested pointers failed")
		}
	})

	t.Run("copy struct with unexported fields", func(t *testing.T) {
		type Private struct {
			Public  string
			private int
		}
		src := Private{Public: "test", private: 42}
		var dst Private
		DeepCopy(src, &dst)

		if dst.Public != "test" {
			t.Errorf("DeepCopy struct with unexported fields failed")
		}
	})

	t.Run("copy empty struct", func(t *testing.T) {
		type Empty struct{}
		src := Empty{}
		var dst Empty
		DeepCopy(src, &dst)
		// Should not panic
	})

	t.Run("copy map with interface values", func(t *testing.T) {
		type Container struct {
			Data map[string]interface{}
		}
		src := Container{
			Data: map[string]interface{}{
				"int":    42,
				"string": "hello",
				"slice":  []int{1, 2, 3},
			},
		}
		var dst Container
		DeepCopy(src, &dst)

		if dst.Data["int"] != 42 {
			t.Errorf("DeepCopy map with interface values failed")
		}
	})

	t.Run("copy slice with pointer elements", func(t *testing.T) {
		val1, val2 := 1, 2
		src := []*int{&val1, &val2, nil}
		var dst []*int
		DeepCopy(src, &dst)

		if len(dst) != 3 || *dst[0] != 1 || *dst[1] != 2 || dst[2] != nil {
			t.Errorf("DeepCopy slice with pointer elements failed")
		}
	})
}

// TestDeepEqualEdgeCases 测试 DeepEqual 的边界情况
func TestDeepEqualEdgeCases(t *testing.T) {
	t.Run("compare empty structs", func(t *testing.T) {
		type Empty struct{}
		a := Empty{}
		b := Empty{}
		if !DeepEqual(a, b) {
			t.Errorf("DeepEqual empty structs should be true")
		}
	})

	t.Run("compare slices with nil elements", func(t *testing.T) {
		a := []*int{nil, nil}
		b := []*int{nil, nil}
		if !DeepEqual(a, b) {
			t.Errorf("DeepEqual slices with nil elements should be true")
		}

		val := 1
		c := []*int{&val, nil}
		if DeepEqual(a, c) {
			t.Errorf("DeepEqual different slices with nil elements should be false")
		}
	})

	t.Run("compare maps with nil values", func(t *testing.T) {
		a := map[string]*int{"a": nil, "b": nil}
		b := map[string]*int{"a": nil, "b": nil}
		if !DeepEqual(a, b) {
			t.Errorf("DeepEqual maps with nil values should be true")
		}
	})

	t.Run("compare nested nil slices", func(t *testing.T) {
		type Container struct {
			Slice []int
		}
		a := Container{Slice: nil}
		b := Container{Slice: nil}
		if !DeepEqual(a, b) {
			t.Errorf("DeepEqual nested nil slices should be true")
		}
	})

	t.Run("compare structs with interface fields", func(t *testing.T) {
		type Container struct {
			Value interface{}
		}
		a := Container{Value: nil}
		b := Container{Value: nil}
		if !DeepEqual(a, b) {
			t.Errorf("DeepEqual structs with nil interface fields should be true")
		}

		c := Container{Value: 42}
		if DeepEqual(a, c) {
			t.Errorf("DeepEqual nil vs non-nil interface should be false")
		}
	})

	t.Run("compare complex nested structures", func(t *testing.T) {
		type Inner struct {
			Value int
		}
		type Middle struct {
			Inners []*Inner
		}
		type Outer struct {
			Middles map[string]*Middle
		}

		a := Outer{
			Middles: map[string]*Middle{
				"m1": {Inners: []*Inner{{1}, {2}}},
			},
		}
		b := Outer{
			Middles: map[string]*Middle{
				"m1": {Inners: []*Inner{{1}, {2}}},
			},
		}

		if !DeepEqual(a, b) {
			t.Errorf("DeepEqual complex nested structures should be true")
		}
	})

	t.Run("compare different length maps", func(t *testing.T) {
		a := map[string]int{"a": 1}
		b := map[string]int{"a": 1, "b": 2}
		if DeepEqual(a, b) {
			t.Errorf("DeepEqual different length maps should be false")
		}
	})

	t.Run("compare maps with missing keys", func(t *testing.T) {
		a := map[string]int{"a": 1, "b": 2}
		b := map[string]int{"a": 1, "c": 2}
		if DeepEqual(a, b) {
			t.Errorf("DeepEqual maps with different keys should be false")
		}
	})
}

// TestDeepCopyComplexScenarios 测试复杂场景
func TestDeepCopyComplexScenarios(t *testing.T) {
	t.Run("copy multi-level nested map", func(t *testing.T) {
		src := map[string]map[string][]int{
			"outer1": {
				"inner1": {1, 2, 3},
				"inner2": {4, 5, 6},
			},
			"outer2": {
				"inner3": {7, 8, 9},
			},
		}
		var dst map[string]map[string][]int
		DeepCopy(src, &dst)

		if len(dst) != 2 || len(dst["outer1"]) != 2 {
			t.Errorf("DeepCopy multi-level nested map failed")
		}

		// Verify independence
		src["outer1"]["inner1"][0] = 999
		if dst["outer1"]["inner1"][0] == 999 {
			t.Errorf("DeepCopy multi-level nested map not independent")
		}
	})

	t.Run("copy slice of slices", func(t *testing.T) {
		src := [][]int{{1, 2}, {3, 4}, {5, 6}}
		var dst [][]int
		DeepCopy(src, &dst)

		if len(dst) != 3 || len(dst[0]) != 2 {
			t.Errorf("DeepCopy slice of slices failed")
		}

		// Verify independence
		src[0][0] = 999
		if dst[0][0] == 999 {
			t.Errorf("DeepCopy slice of slices not independent")
		}
	})

	t.Run("copy array of structs", func(t *testing.T) {
		type Point struct {
			X, Y int
		}
		src := [3]Point{{1, 2}, {3, 4}, {5, 6}}
		var dst [3]Point
		DeepCopy(src, &dst)

		if dst[0].X != 1 || dst[2].Y != 6 {
			t.Errorf("DeepCopy array of structs failed")
		}
	})
}

// TestDeepEqualComplexScenarios 测试 DeepEqual 的复杂场景
func TestDeepEqualComplexScenarios(t *testing.T) {
	t.Run("equal multi-level nested structures", func(t *testing.T) {
		type Level3 struct {
			Value int
		}
		type Level2 struct {
			L3 *Level3
		}
		type Level1 struct {
			L2 map[string]*Level2
		}

		a := Level1{
			L2: map[string]*Level2{
				"key": {L3: &Level3{Value: 42}},
			},
		}
		b := Level1{
			L2: map[string]*Level2{
				"key": {L3: &Level3{Value: 42}},
			},
		}

		if !DeepEqual(a, b) {
			t.Errorf("DeepEqual multi-level nested structures should be true")
		}

		c := Level1{
			L2: map[string]*Level2{
				"key": {L3: &Level3{Value: 43}},
			},
		}
		if DeepEqual(a, c) {
			t.Errorf("DeepEqual different nested values should be false")
		}
	})

	t.Run("equal arrays of arrays", func(t *testing.T) {
		a := [2][3]int{{1, 2, 3}, {4, 5, 6}}
		b := [2][3]int{{1, 2, 3}, {4, 5, 6}}
		c := [2][3]int{{1, 2, 3}, {4, 5, 7}}

		if !DeepEqual(a, b) {
			t.Errorf("DeepEqual arrays of arrays should be true")
		}
		if DeepEqual(a, c) {
			t.Errorf("DeepEqual different arrays of arrays should be false")
		}
	})

	t.Run("equal slices of pointers to arrays", func(t *testing.T) {
		arr1 := [2]int{1, 2}
		arr2 := [2]int{1, 2}
		a := []*[2]int{&arr1}
		b := []*[2]int{&arr2}

		if !DeepEqual(a, b) {
			t.Errorf("DeepEqual slices of pointers to arrays should be true")
		}
	})

	t.Run("compare structs with function fields", func(t *testing.T) {
		type Container struct {
			Fn   func()
			Data int
		}

		fn1 := func() {}
		fn2 := func() {}

		a := Container{Fn: fn1, Data: 42}
		b := Container{Fn: fn2, Data: 42}

		// Functions are not comparable, but struct comparison should handle it
		// This tests the panic recovery path in deepValueEqual
		result := DeepEqual(a, b)
		// Result depends on whether functions are considered equal
		_ = result // Just ensure no panic
	})

	t.Run("compare structs with channel fields", func(t *testing.T) {
		type Container struct {
			Ch   chan int
			Data int
		}

		a := Container{Ch: make(chan int), Data: 42}
		b := Container{Ch: make(chan int), Data: 42}

		// Channels are comparable by identity
		result := DeepEqual(a, b)
		if result {
			t.Errorf("DeepEqual different channels should be false")
		}
	})

	t.Run("compare same channel references", func(t *testing.T) {
		type Container struct {
			Ch chan int
		}

		ch := make(chan int)
		a := Container{Ch: ch}
		b := Container{Ch: ch}

		if !DeepEqual(a, b) {
			t.Errorf("DeepEqual same channel reference should be true")
		}
	})
}

// TestDeepCopyPanicCases 测试会引发 panic 的情况
func TestDeepCopyPanicCases(t *testing.T) {
	t.Run("type mismatch should panic", func(t *testing.T) {
		defer func() {
			if r := recover(); r == nil {
				t.Error("DeepCopy with type mismatch should panic")
			}
		}()

		src := 42
		var dst string
		DeepCopy(src, &dst)
	})

	t.Run("unsupported type should panic", func(t *testing.T) {
		defer func() {
			if r := recover(); r == nil {
				t.Error("DeepCopy with unsupported type should panic")
			}
		}()

		type Unsupported struct {
			Ch chan int
		}
		src := Unsupported{Ch: make(chan int)}
		var dst Unsupported
		DeepCopy(src, &dst)
	})

	t.Run("function type should panic", func(t *testing.T) {
		defer func() {
			if r := recover(); r == nil {
				t.Error("DeepCopy with function type should panic")
			}
		}()

		type Container struct {
			Fn func()
		}
		src := Container{Fn: func() {}}
		var dst Container
		DeepCopy(src, &dst)
	})
}

// TestDeepCopyInvalidValues 测试 invalid 值的处理
func TestDeepCopyInvalidValues(t *testing.T) {
	t.Run("copy with invalid source", func(t *testing.T) {
		// Create an invalid reflect.Value
		var invalidVal interface{}
		var dst int
		DeepCopy(invalidVal, &dst)
		// Should not panic, just return without doing anything
	})

	t.Run("copy all integer types", func(t *testing.T) {
		// Test int8
		var srcInt8 int8 = 127
		var dstInt8 int8
		DeepCopy(srcInt8, &dstInt8)
		if dstInt8 != 127 {
			t.Errorf("int8 copy failed")
		}

		// Test int16
		var srcInt16 int16 = 32767
		var dstInt16 int16
		DeepCopy(srcInt16, &dstInt16)
		if dstInt16 != 32767 {
			t.Errorf("int16 copy failed")
		}

		// Test int32
		var srcInt32 int32 = 2147483647
		var dstInt32 int32
		DeepCopy(srcInt32, &dstInt32)
		if dstInt32 != 2147483647 {
			t.Errorf("int32 copy failed")
		}

		// Test int64
		var srcInt64 int64 = 9223372036854775807
		var dstInt64 int64
		DeepCopy(srcInt64, &dstInt64)
		if dstInt64 != 9223372036854775807 {
			t.Errorf("int64 copy failed")
		}
	})

	t.Run("copy all unsigned integer types", func(t *testing.T) {
		// Test uint8
		var srcUint8 uint8 = 255
		var dstUint8 uint8
		DeepCopy(srcUint8, &dstUint8)
		if dstUint8 != 255 {
			t.Errorf("uint8 copy failed")
		}

		// Test uint16
		var srcUint16 uint16 = 65535
		var dstUint16 uint16
		DeepCopy(srcUint16, &dstUint16)
		if dstUint16 != 65535 {
			t.Errorf("uint16 copy failed")
		}

		// Test uint32
		var srcUint32 uint32 = 4294967295
		var dstUint32 uint32
		DeepCopy(srcUint32, &dstUint32)
		if dstUint32 != 4294967295 {
			t.Errorf("uint32 copy failed")
		}

		// Test uint64
		var srcUint64 uint64 = 18446744073709551615
		var dstUint64 uint64
		DeepCopy(srcUint64, &dstUint64)
		if dstUint64 != 18446744073709551615 {
			t.Errorf("uint64 copy failed")
		}
	})

	t.Run("copy complex64", func(t *testing.T) {
		var srcComplex64 complex64 = complex(float32(1.5), float32(2.5))
		var dstComplex64 complex64
		DeepCopy(srcComplex64, &dstComplex64)
		if dstComplex64 != srcComplex64 {
			t.Errorf("complex64 copy failed")
		}
	})

	t.Run("copy pointer to nil pointer", func(t *testing.T) {
		var nilPtr *int
		var dst *int
		DeepCopy(&nilPtr, &dst)
		if dst != nil {
			t.Errorf("copy pointer to nil pointer should result in nil")
		}
	})

	t.Run("copy invalid kind in struct", func(t *testing.T) {
		// Test the Invalid kind branch
		type Container struct {
			Valid int
		}
		src := Container{Valid: 42}
		var dst Container
		DeepCopy(src, &dst)
		if dst.Valid != 42 {
			t.Errorf("copy struct with valid fields failed")
		}
	})
}

// TestDeepEqualInvalidValues 测试 DeepEqual 的 invalid 值处理
func TestDeepEqualInvalidValues(t *testing.T) {
	t.Run("compare invalid values", func(t *testing.T) {
		var a, b interface{}
		if !DeepEqual(a, b) {
			t.Errorf("DeepEqual nil interfaces should be true")
		}
	})

	t.Run("compare one invalid value", func(t *testing.T) {
		var a interface{}
		var b interface{} = 42
		if DeepEqual(a, b) {
			t.Errorf("DeepEqual nil vs value should be false")
		}
	})
}

// TestTypedSliceCopyAllTypes 测试所有基本类型的切片复制
func TestTypedSliceCopyAllTypes(t *testing.T) {
	t.Run("copy bool slice", func(t *testing.T) {
		src := []bool{true, false, true}
		dst := TypedSliceCopy(src)
		if len(dst) != 3 || dst[0] != true || dst[1] != false {
			t.Errorf("bool slice copy failed")
		}
	})

	t.Run("copy uint slice", func(t *testing.T) {
		src := []uint{1, 2, 3}
		dst := TypedSliceCopy(src)
		if len(dst) != 3 || dst[0] != 1 {
			t.Errorf("uint slice copy failed")
		}
	})

	t.Run("copy float32 slice", func(t *testing.T) {
		src := []float32{1.1, 2.2, 3.3}
		dst := TypedSliceCopy(src)
		if len(dst) != 3 || dst[0] != 1.1 {
			t.Errorf("float32 slice copy failed")
		}
	})

	t.Run("copy float64 slice", func(t *testing.T) {
		src := []float64{1.1, 2.2, 3.3}
		dst := TypedSliceCopy(src)
		if len(dst) != 3 || dst[0] != 1.1 {
			t.Errorf("float64 slice copy failed")
		}
	})

	t.Run("copy complex128 slice", func(t *testing.T) {
		src := []complex128{complex(1, 2), complex(3, 4)}
		dst := TypedSliceCopy(src)
		if len(dst) != 2 || dst[0] != complex(1, 2) {
			t.Errorf("complex128 slice copy failed")
		}
	})
}

// TestTypedMapCopyAllTypes 测试所有基本类型的 map 复制
func TestTypedMapCopyAllTypes(t *testing.T) {
	t.Run("copy bool value map", func(t *testing.T) {
		src := map[string]bool{"a": true, "b": false}
		dst := TypedMapCopy(src)
		if len(dst) != 2 || dst["a"] != true {
			t.Errorf("bool value map copy failed")
		}
	})

	t.Run("copy uint value map", func(t *testing.T) {
		src := map[string]uint{"a": 1, "b": 2}
		dst := TypedMapCopy(src)
		if len(dst) != 2 || dst["a"] != 1 {
			t.Errorf("uint value map copy failed")
		}
	})

	t.Run("copy float32 value map", func(t *testing.T) {
		src := map[string]float32{"a": 1.1, "b": 2.2}
		dst := TypedMapCopy(src)
		if len(dst) != 2 || dst["a"] != 1.1 {
			t.Errorf("float32 value map copy failed")
		}
	})

	t.Run("copy float64 value map", func(t *testing.T) {
		src := map[string]float64{"a": 1.1, "b": 2.2}
		dst := TypedMapCopy(src)
		if len(dst) != 2 || dst["a"] != 1.1 {
			t.Errorf("float64 value map copy failed")
		}
	})

	t.Run("copy complex128 value map", func(t *testing.T) {
		src := map[string]complex128{"a": complex(1, 2)}
		dst := TypedMapCopy(src)
		if len(dst) != 1 || dst["a"] != complex(1, 2) {
			t.Errorf("complex128 value map copy failed")
		}
	})
}
