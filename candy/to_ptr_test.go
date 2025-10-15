package candy

import (
	"testing"
)

func TestToPtr(t *testing.T) {
	t.Run("int value", func(t *testing.T) {
		val := 42
		ptr := ToPtr(val)

		if ptr == nil {
			t.Error("ToPtr(42) returned nil")
		}
		if *ptr != val {
			t.Errorf("*ToPtr(%d) = %d, want %d", val, *ptr, val)
		}

		// Verify it's a different address
		if ptr == &val {
			t.Error("ToPtr should return a new pointer, not the original variable's address")
		}
	})

	t.Run("string value", func(t *testing.T) {
		val := "hello"
		ptr := ToPtr(val)

		if ptr == nil {
			t.Error("ToPtr(\"hello\") returned nil")
		}
		if *ptr != val {
			t.Errorf("*ToPtr(%q) = %q, want %q", val, *ptr, val)
		}
	})

	t.Run("bool value", func(t *testing.T) {
		val := true
		ptr := ToPtr(val)

		if ptr == nil {
			t.Error("ToPtr(true) returned nil")
		}
		if *ptr != val {
			t.Errorf("*ToPtr(%v) = %v, want %v", val, *ptr, val)
		}
	})

	t.Run("float64 value", func(t *testing.T) {
		val := 3.14159
		ptr := ToPtr(val)

		if ptr == nil {
			t.Error("ToPtr(3.14159) returned nil")
		}
		if *ptr != val {
			t.Errorf("*ToPtr(%f) = %f, want %f", val, *ptr, val)
		}
	})

	t.Run("struct value", func(t *testing.T) {
		type TestStruct struct {
			Name  string
			Value int
		}
		val := TestStruct{Name: "test", Value: 42}
		ptr := ToPtr(val)

		if ptr == nil {
			t.Error("ToPtr(struct) returned nil")
		}
		if ptr.Name != val.Name || ptr.Value != val.Value {
			t.Errorf("*ToPtr(struct) = %+v, want %+v", *ptr, val)
		}
	})

	t.Run("slice value", func(t *testing.T) {
		val := []int{1, 2, 3}
		ptr := ToPtr(val)

		if ptr == nil {
			t.Error("ToPtr(slice) returned nil")
		}
		if len(*ptr) != len(val) {
			t.Errorf("len(*ToPtr(slice)) = %d, want %d", len(*ptr), len(val))
		}
	})

	t.Run("map value", func(t *testing.T) {
		val := map[string]int{"a": 1, "b": 2}
		ptr := ToPtr(val)

		if ptr == nil {
			t.Error("ToPtr(map) returned nil")
		}
		if len(*ptr) != len(val) {
			t.Errorf("len(*ToPtr(map)) = %d, want %d", len(*ptr), len(val))
		}
	})

	t.Run("zero values", func(t *testing.T) {
		t.Run("zero int", func(t *testing.T) {
			val := 0
			ptr := ToPtr(val)
			if ptr == nil || *ptr != 0 {
				t.Errorf("ToPtr(0) failed")
			}
		})

		t.Run("empty string", func(t *testing.T) {
			val := ""
			ptr := ToPtr(val)
			if ptr == nil || *ptr != "" {
				t.Errorf("ToPtr(\"\") failed")
			}
		})

		t.Run("false bool", func(t *testing.T) {
			val := false
			ptr := ToPtr(val)
			if ptr == nil || *ptr != false {
				t.Errorf("ToPtr(false) failed")
			}
		})

		t.Run("nil slice", func(t *testing.T) {
			var val []int
			ptr := ToPtr(val)
			if ptr == nil {
				t.Error("ToPtr(nil slice) returned nil pointer")
			}
			if *ptr != nil {
				t.Error("*ToPtr(nil slice) should be nil")
			}
		})

		t.Run("nil map", func(t *testing.T) {
			var val map[string]int
			ptr := ToPtr(val)
			if ptr == nil {
				t.Error("ToPtr(nil map) returned nil pointer")
			}
			if *ptr != nil {
				t.Error("*ToPtr(nil map) should be nil")
			}
		})
	})

	t.Run("pointer to pointer", func(t *testing.T) {
		val := 42
		ptr1 := &val
		ptr2 := ToPtr(ptr1)

		if ptr2 == nil {
			t.Error("ToPtr(pointer) returned nil")
		}
		if *ptr2 != ptr1 {
			t.Error("ToPtr should handle pointer values")
		}
	})

	t.Run("interface value", func(t *testing.T) {
		var val interface{} = "test"
		ptr := ToPtr(val)

		if ptr == nil {
			t.Error("ToPtr(interface{}) returned nil")
		}
		if str, ok := (*ptr).(string); !ok || str != "test" {
			t.Errorf("*ToPtr(interface{}) type assertion failed")
		}
	})

	t.Run("byte value", func(t *testing.T) {
		val := byte(255)
		ptr := ToPtr(val)

		if ptr == nil || *ptr != 255 {
			t.Errorf("ToPtr(byte) failed")
		}
	})

	t.Run("rune value", func(t *testing.T) {
		val := 'A'
		ptr := ToPtr(val)

		if ptr == nil || *ptr != 'A' {
			t.Errorf("ToPtr(rune) failed")
		}
	})

	t.Run("complex64 value", func(t *testing.T) {
		val := complex64(1 + 2i)
		ptr := ToPtr(val)

		if ptr == nil || *ptr != val {
			t.Errorf("ToPtr(complex64) failed")
		}
	})

	t.Run("complex128 value", func(t *testing.T) {
		val := complex128(3 + 4i)
		ptr := ToPtr(val)

		if ptr == nil || *ptr != val {
			t.Errorf("ToPtr(complex128) failed")
		}
	})

	t.Run("array value", func(t *testing.T) {
		val := [3]int{1, 2, 3}
		ptr := ToPtr(val)

		if ptr == nil {
			t.Error("ToPtr(array) returned nil")
		}
		if (*ptr)[0] != 1 || (*ptr)[1] != 2 || (*ptr)[2] != 3 {
			t.Errorf("*ToPtr(array) values incorrect")
		}
	})

	t.Run("function value", func(t *testing.T) {
		val := func() int { return 42 }
		ptr := ToPtr(val)

		if ptr == nil {
			t.Error("ToPtr(function) returned nil")
		}
		if (*ptr)() != 42 {
			t.Errorf("*ToPtr(function)() = %d, want 42", (*ptr)())
		}
	})

	t.Run("modification independence", func(t *testing.T) {
		original := 10
		ptr := ToPtr(original)
		*ptr = 20

		if original != 10 {
			t.Errorf("Modifying *ToPtr() affected original value: original = %d, want 10", original)
		}
	})

	t.Run("generic type constraint", func(t *testing.T) {
		// Test with various types to ensure generics work
		intPtr := ToPtr(int(1))
		int8Ptr := ToPtr(int8(2))
		int16Ptr := ToPtr(int16(3))
		int32Ptr := ToPtr(int32(4))
		int64Ptr := ToPtr(int64(5))

		if *intPtr != 1 || *int8Ptr != 2 || *int16Ptr != 3 || *int32Ptr != 4 || *int64Ptr != 5 {
			t.Error("Generic int type constraints failed")
		}

		uintPtr := ToPtr(uint(1))
		uint8Ptr := ToPtr(uint8(2))
		uint16Ptr := ToPtr(uint16(3))
		uint32Ptr := ToPtr(uint32(4))
		uint64Ptr := ToPtr(uint64(5))

		if *uintPtr != 1 || *uint8Ptr != 2 || *uint16Ptr != 3 || *uint32Ptr != 4 || *uint64Ptr != 5 {
			t.Error("Generic uint type constraints failed")
		}

		float32Ptr := ToPtr(float32(1.5))
		float64Ptr := ToPtr(float64(2.5))

		if *float32Ptr != 1.5 || *float64Ptr != 2.5 {
			t.Error("Generic float type constraints failed")
		}
	})

	t.Run("channel value", func(t *testing.T) {
		val := make(chan int, 1)
		val <- 42
		ptr := ToPtr(val)

		if ptr == nil {
			t.Error("ToPtr(channel) returned nil")
		}

		select {
		case v := <-*ptr:
			if v != 42 {
				t.Errorf("Channel value = %d, want 42", v)
			}
		default:
			t.Error("Channel should have a value")
		}
	})
}

func BenchmarkToPtr(b *testing.B) {
	b.Run("int", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = ToPtr(42)
		}
	})

	b.Run("string", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = ToPtr("hello")
		}
	})

	b.Run("struct", func(b *testing.B) {
		type TestStruct struct {
			Name  string
			Value int
		}
		s := TestStruct{Name: "test", Value: 42}
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			_ = ToPtr(s)
		}
	})

	b.Run("slice", func(b *testing.B) {
		slice := []int{1, 2, 3, 4, 5}
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			_ = ToPtr(slice)
		}
	})
}
