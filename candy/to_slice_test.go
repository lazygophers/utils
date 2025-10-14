package candy

import (
	"reflect"
	"testing"
)

func TestToFloat64Slice(t *testing.T) {
	t.Run("nil input", func(t *testing.T) {
		result := ToFloat64Slice(nil)
		if result != nil {
			t.Errorf("ToFloat64Slice(nil) = %v, want nil", result)
		}
	})

	t.Run("bool slice", func(t *testing.T) {
		input := []bool{true, false, true}
		expected := []float64{1, 0, 1}
		result := ToFloat64Slice(input)
		if !reflect.DeepEqual(result, expected) {
			t.Errorf("ToFloat64Slice(%v) = %v, want %v", input, result, expected)
		}
	})

	t.Run("int slice", func(t *testing.T) {
		input := []int{1, 2, 3}
		expected := []float64{1, 2, 3}
		result := ToFloat64Slice(input)
		if !reflect.DeepEqual(result, expected) {
			t.Errorf("ToFloat64Slice(%v) = %v, want %v", input, result, expected)
		}
	})

	t.Run("int8 slice", func(t *testing.T) {
		input := []int8{1, 2, 3}
		expected := []float64{1, 2, 3}
		result := ToFloat64Slice(input)
		if !reflect.DeepEqual(result, expected) {
			t.Errorf("ToFloat64Slice(%v) = %v, want %v", input, result, expected)
		}
	})

	t.Run("int16 slice", func(t *testing.T) {
		input := []int16{100, 200, 300}
		expected := []float64{100, 200, 300}
		result := ToFloat64Slice(input)
		if !reflect.DeepEqual(result, expected) {
			t.Errorf("ToFloat64Slice(%v) = %v, want %v", input, result, expected)
		}
	})

	t.Run("int32 slice", func(t *testing.T) {
		input := []int32{1000, 2000, 3000}
		expected := []float64{1000, 2000, 3000}
		result := ToFloat64Slice(input)
		if !reflect.DeepEqual(result, expected) {
			t.Errorf("ToFloat64Slice(%v) = %v, want %v", input, result, expected)
		}
	})

	t.Run("int64 slice", func(t *testing.T) {
		input := []int64{100000, 200000, 300000}
		expected := []float64{100000, 200000, 300000}
		result := ToFloat64Slice(input)
		if !reflect.DeepEqual(result, expected) {
			t.Errorf("ToFloat64Slice(%v) = %v, want %v", input, result, expected)
		}
	})

	t.Run("uint slice", func(t *testing.T) {
		input := []uint{1, 2, 3}
		expected := []float64{1, 2, 3}
		result := ToFloat64Slice(input)
		if !reflect.DeepEqual(result, expected) {
			t.Errorf("ToFloat64Slice(%v) = %v, want %v", input, result, expected)
		}
	})

	t.Run("uint8 slice", func(t *testing.T) {
		input := []uint8{1, 2, 3}
		expected := []float64{1, 2, 3}
		result := ToFloat64Slice(input)
		if !reflect.DeepEqual(result, expected) {
			t.Errorf("ToFloat64Slice(%v) = %v, want %v", input, result, expected)
		}
	})

	t.Run("uint16 slice", func(t *testing.T) {
		input := []uint16{100, 200, 300}
		expected := []float64{100, 200, 300}
		result := ToFloat64Slice(input)
		if !reflect.DeepEqual(result, expected) {
			t.Errorf("ToFloat64Slice(%v) = %v, want %v", input, result, expected)
		}
	})

	t.Run("uint32 slice", func(t *testing.T) {
		input := []uint32{1000, 2000, 3000}
		expected := []float64{1000, 2000, 3000}
		result := ToFloat64Slice(input)
		if !reflect.DeepEqual(result, expected) {
			t.Errorf("ToFloat64Slice(%v) = %v, want %v", input, result, expected)
		}
	})

	t.Run("uint64 slice", func(t *testing.T) {
		input := []uint64{100000, 200000, 300000}
		expected := []float64{100000, 200000, 300000}
		result := ToFloat64Slice(input)
		if !reflect.DeepEqual(result, expected) {
			t.Errorf("ToFloat64Slice(%v) = %v, want %v", input, result, expected)
		}
	})

	t.Run("float32 slice", func(t *testing.T) {
		input := []float32{1.1, 2.2, 3.3}
		result := ToFloat64Slice(input)
		if len(result) != 3 {
			t.Errorf("ToFloat64Slice(%v) length = %d, want 3", input, len(result))
		}
	})

	t.Run("float64 slice", func(t *testing.T) {
		input := []float64{1.1, 2.2, 3.3}
		expected := []float64{1.1, 2.2, 3.3}
		result := ToFloat64Slice(input)
		if !reflect.DeepEqual(result, expected) {
			t.Errorf("ToFloat64Slice(%v) = %v, want %v", input, result, expected)
		}
	})

	t.Run("string slice", func(t *testing.T) {
		input := []string{"1.5", "2.5", "3.5"}
		expected := []float64{1.5, 2.5, 3.5}
		result := ToFloat64Slice(input)
		if !reflect.DeepEqual(result, expected) {
			t.Errorf("ToFloat64Slice(%v) = %v, want %v", input, result, expected)
		}
	})

	t.Run("byte slice slice", func(t *testing.T) {
		input := [][]byte{[]byte("1"), []byte("2"), []byte("3")}
		expected := []float64{1, 2, 3}
		result := ToFloat64Slice(input)
		if !reflect.DeepEqual(result, expected) {
			t.Errorf("ToFloat64Slice(%v) = %v, want %v", input, result, expected)
		}
	})

	t.Run("interface slice", func(t *testing.T) {
		input := []interface{}{1, "2.5", true}
		expected := []float64{1, 2.5, 1}
		result := ToFloat64Slice(input)
		if !reflect.DeepEqual(result, expected) {
			t.Errorf("ToFloat64Slice(%v) = %v, want %v", input, result, expected)
		}
	})

	t.Run("empty slices", func(t *testing.T) {
		tests := []struct {
			name  string
			input interface{}
		}{
			{"empty int slice", []int{}},
			{"empty string slice", []string{}},
			{"empty float64 slice", []float64{}},
		}

		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				result := ToFloat64Slice(tt.input)
				if len(result) != 0 {
					t.Errorf("ToFloat64Slice(%v) length = %d, want 0", tt.input, len(result))
				}
			})
		}
	})

	t.Run("unsupported type", func(t *testing.T) {
		input := "not a slice"
		result := ToFloat64Slice(input)
		expected := []float64{}
		if !reflect.DeepEqual(result, expected) {
			t.Errorf("ToFloat64Slice(%v) = %v, want %v", input, result, expected)
		}
	})
}

func TestToInt64Slice(t *testing.T) {
	t.Run("bool slice", func(t *testing.T) {
		input := []bool{true, false, true}
		expected := []int64{1, 0, 1}
		result := ToInt64Slice(input)
		if !reflect.DeepEqual(result, expected) {
			t.Errorf("ToInt64Slice(%v) = %v, want %v", input, result, expected)
		}
	})

	t.Run("int slice", func(t *testing.T) {
		input := []int{1, 2, 3}
		expected := []int64{1, 2, 3}
		result := ToInt64Slice(input)
		if !reflect.DeepEqual(result, expected) {
			t.Errorf("ToInt64Slice(%v) = %v, want %v", input, result, expected)
		}
	})

	t.Run("int8 slice", func(t *testing.T) {
		input := []int8{1, 2, 3}
		expected := []int64{1, 2, 3}
		result := ToInt64Slice(input)
		if !reflect.DeepEqual(result, expected) {
			t.Errorf("ToInt64Slice(%v) = %v, want %v", input, result, expected)
		}
	})

	t.Run("int16 slice", func(t *testing.T) {
		input := []int16{100, 200, 300}
		expected := []int64{100, 200, 300}
		result := ToInt64Slice(input)
		if !reflect.DeepEqual(result, expected) {
			t.Errorf("ToInt64Slice(%v) = %v, want %v", input, result, expected)
		}
	})

	t.Run("int32 slice", func(t *testing.T) {
		input := []int32{1000, 2000, 3000}
		expected := []int64{1000, 2000, 3000}
		result := ToInt64Slice(input)
		if !reflect.DeepEqual(result, expected) {
			t.Errorf("ToInt64Slice(%v) = %v, want %v", input, result, expected)
		}
	})

	t.Run("int64 slice", func(t *testing.T) {
		input := []int64{100000, 200000, 300000}
		expected := []int64{100000, 200000, 300000}
		result := ToInt64Slice(input)
		if !reflect.DeepEqual(result, expected) {
			t.Errorf("ToInt64Slice(%v) = %v, want %v", input, result, expected)
		}
	})

	t.Run("uint slice", func(t *testing.T) {
		input := []uint{1, 2, 3}
		expected := []int64{1, 2, 3}
		result := ToInt64Slice(input)
		if !reflect.DeepEqual(result, expected) {
			t.Errorf("ToInt64Slice(%v) = %v, want %v", input, result, expected)
		}
	})

	t.Run("uint8 slice", func(t *testing.T) {
		input := []uint8{1, 2, 3}
		expected := []int64{1, 2, 3}
		result := ToInt64Slice(input)
		if !reflect.DeepEqual(result, expected) {
			t.Errorf("ToInt64Slice(%v) = %v, want %v", input, result, expected)
		}
	})

	t.Run("uint16 slice", func(t *testing.T) {
		input := []uint16{100, 200, 300}
		expected := []int64{100, 200, 300}
		result := ToInt64Slice(input)
		if !reflect.DeepEqual(result, expected) {
			t.Errorf("ToInt64Slice(%v) = %v, want %v", input, result, expected)
		}
	})

	t.Run("uint32 slice", func(t *testing.T) {
		input := []uint32{1000, 2000, 3000}
		expected := []int64{1000, 2000, 3000}
		result := ToInt64Slice(input)
		if !reflect.DeepEqual(result, expected) {
			t.Errorf("ToInt64Slice(%v) = %v, want %v", input, result, expected)
		}
	})

	t.Run("uint64 slice", func(t *testing.T) {
		input := []uint64{100000, 200000, 300000}
		expected := []int64{100000, 200000, 300000}
		result := ToInt64Slice(input)
		if !reflect.DeepEqual(result, expected) {
			t.Errorf("ToInt64Slice(%v) = %v, want %v", input, result, expected)
		}
	})

	t.Run("float32 slice", func(t *testing.T) {
		input := []float32{1.5, 2.7, 3.9}
		expected := []int64{1, 2, 3}
		result := ToInt64Slice(input)
		if !reflect.DeepEqual(result, expected) {
			t.Errorf("ToInt64Slice(%v) = %v, want %v", input, result, expected)
		}
	})

	t.Run("float64 slice", func(t *testing.T) {
		input := []float64{1.5, 2.7, 3.9}
		expected := []int64{1, 2, 3}
		result := ToInt64Slice(input)
		if !reflect.DeepEqual(result, expected) {
			t.Errorf("ToInt64Slice(%v) = %v, want %v", input, result, expected)
		}
	})

	t.Run("string slice", func(t *testing.T) {
		input := []string{"1", "2", "3"}
		expected := []int64{1, 2, 3}
		result := ToInt64Slice(input)
		if !reflect.DeepEqual(result, expected) {
			t.Errorf("ToInt64Slice(%v) = %v, want %v", input, result, expected)
		}
	})

	t.Run("byte slice slice", func(t *testing.T) {
		input := [][]byte{[]byte("1"), []byte("2"), []byte("3")}
		expected := []int64{1, 2, 3}
		result := ToInt64Slice(input)
		if !reflect.DeepEqual(result, expected) {
			t.Errorf("ToInt64Slice(%v) = %v, want %v", input, result, expected)
		}
	})

	t.Run("interface slice", func(t *testing.T) {
		input := []interface{}{1, "2", true, 3.5}
		expected := []int64{1, 2, 1, 3}
		result := ToInt64Slice(input)
		if !reflect.DeepEqual(result, expected) {
			t.Errorf("ToInt64Slice(%v) = %v, want %v", input, result, expected)
		}
	})

	t.Run("empty slices", func(t *testing.T) {
		tests := []struct {
			name  string
			input interface{}
		}{
			{"empty int slice", []int{}},
			{"empty string slice", []string{}},
			{"empty float64 slice", []float64{}},
		}

		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				result := ToInt64Slice(tt.input)
				if len(result) != 0 {
					t.Errorf("ToInt64Slice(%v) length = %d, want 0", tt.input, len(result))
				}
			})
		}
	})

	t.Run("unsupported type", func(t *testing.T) {
		input := "not a slice"
		result := ToInt64Slice(input)
		expected := []int64{}
		if !reflect.DeepEqual(result, expected) {
			t.Errorf("ToInt64Slice(%v) = %v, want %v", input, result, expected)
		}
	})
}

func TestToStringSlice(t *testing.T) {
	t.Run("nil input", func(t *testing.T) {
		result := ToStringSlice(nil)
		if result != nil {
			t.Errorf("ToStringSlice(nil) = %v, want nil", result)
		}
	})

	t.Run("string with comma", func(t *testing.T) {
		input := "a,b,c"
		expected := []string{"a", "b", "c"}
		result := ToStringSlice(input)
		if !reflect.DeepEqual(result, expected) {
			t.Errorf("ToStringSlice(%q) = %v, want %v", input, result, expected)
		}
	})

	t.Run("string without comma", func(t *testing.T) {
		input := "hello"
		expected := []string{"hello"}
		result := ToStringSlice(input)
		if !reflect.DeepEqual(result, expected) {
			t.Errorf("ToStringSlice(%q) = %v, want %v", input, result, expected)
		}
	})

	t.Run("empty string", func(t *testing.T) {
		input := ""
		expected := []string{""}
		result := ToStringSlice(input)
		if !reflect.DeepEqual(result, expected) {
			t.Errorf("ToStringSlice(%q) = %v, want %v", input, result, expected)
		}
	})

	t.Run("int slice", func(t *testing.T) {
		input := []int{1, 2, 3}
		expected := []string{"1", "2", "3"}
		result := ToStringSlice(input)
		if !reflect.DeepEqual(result, expected) {
			t.Errorf("ToStringSlice(%v) = %v, want %v", input, result, expected)
		}
	})

	t.Run("string slice", func(t *testing.T) {
		input := []string{"a", "b", "c"}
		expected := []string{"a", "b", "c"}
		result := ToStringSlice(input)
		if !reflect.DeepEqual(result, expected) {
			t.Errorf("ToStringSlice(%v) = %v, want %v", input, result, expected)
		}
	})

	t.Run("float slice", func(t *testing.T) {
		input := []float64{1.1, 2.2, 3.3}
		result := ToStringSlice(input)
		if len(result) != 3 {
			t.Errorf("ToStringSlice(%v) length = %d, want 3", input, len(result))
		}
	})

	t.Run("bool slice", func(t *testing.T) {
		input := []bool{true, false, true}
		expected := []string{"1", "0", "1"}
		result := ToStringSlice(input)
		if !reflect.DeepEqual(result, expected) {
			t.Errorf("ToStringSlice(%v) = %v, want %v", input, result, expected)
		}
	})

	t.Run("interface slice", func(t *testing.T) {
		input := []interface{}{1, "hello", true, 3.14}
		result := ToStringSlice(input)
		if len(result) != 4 {
			t.Errorf("ToStringSlice(%v) length = %d, want 4", input, len(result))
		}
	})

	t.Run("nil slice", func(t *testing.T) {
		var input []int
		result := ToStringSlice(input)
		if result != nil {
			t.Errorf("ToStringSlice(nil slice) = %v, want nil", result)
		}
	})

	t.Run("empty slice", func(t *testing.T) {
		input := []int{}
		result := ToStringSlice(input)
		if len(result) != 0 {
			t.Errorf("ToStringSlice(empty slice) length = %d, want 0", len(result))
		}
	})

	t.Run("single element slice", func(t *testing.T) {
		input := []int{42}
		expected := []string{"42"}
		result := ToStringSlice(input)
		if !reflect.DeepEqual(result, expected) {
			t.Errorf("ToStringSlice(%v) = %v, want %v", input, result, expected)
		}
	})

	t.Run("non-slice non-string type", func(t *testing.T) {
		input := 42
		expected := []string{"42"}
		result := ToStringSlice(input)
		if !reflect.DeepEqual(result, expected) {
			t.Errorf("ToStringSlice(%v) = %v, want %v", input, result, expected)
		}
	})

	t.Run("struct type", func(t *testing.T) {
		type TestStruct struct {
			Name  string `json:"name"`
			Value int    `json:"value"`
		}
		input := TestStruct{Name: "test", Value: 42}
		result := ToStringSlice(input)
		if len(result) != 1 {
			t.Errorf("ToStringSlice(struct) length = %d, want 1", len(result))
		}
	})
}

func TestToArrayString(t *testing.T) {
	t.Run("alias function works", func(t *testing.T) {
		input := []int{1, 2, 3}
		expected := []string{"1", "2", "3"}
		result := ToArrayString(input)
		if !reflect.DeepEqual(result, expected) {
			t.Errorf("ToArrayString(%v) = %v, want %v", input, result, expected)
		}
	})

	t.Run("same as ToStringSlice", func(t *testing.T) {
		input := "a,b,c"
		result1 := ToStringSlice(input)
		result2 := ToArrayString(input)
		if !reflect.DeepEqual(result1, result2) {
			t.Errorf("ToArrayString and ToStringSlice produce different results")
		}
	})
}

func BenchmarkToFloat64Slice(b *testing.B) {
	input := []int{1, 2, 3, 4, 5}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		ToFloat64Slice(input)
	}
}

func BenchmarkToInt64Slice(b *testing.B) {
	input := []float64{1.1, 2.2, 3.3, 4.4, 5.5}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		ToInt64Slice(input)
	}
}

func BenchmarkToStringSlice(b *testing.B) {
	input := []int{1, 2, 3, 4, 5}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		ToStringSlice(input)
	}
}
