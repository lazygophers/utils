package candy

import (
	"reflect"
	"testing"
)

type TestItem struct {
	ID       int
	ID8      int8
	ID16     int16
	ID32     int32
	ID64     int64
	UID      uint
	UID8     uint8
	UID16    uint16
	UID32    uint32
	UID64    uint64
	Name     string
	Code     string
	Price    float32
	Score    float64
	Active   bool
	Verified bool
}

func TestSlice2Map(t *testing.T) {
	t.Run("string slice", func(t *testing.T) {
		input := []string{"a", "b", "c"}
		result := Slice2Map(input)
		expected := map[string]bool{"a": true, "b": true, "c": true}

		if !reflect.DeepEqual(result, expected) {
			t.Errorf("Slice2Map() = %v, want %v", result, expected)
		}
	})

	t.Run("integer slice", func(t *testing.T) {
		input := []int{1, 2, 3}
		result := Slice2Map(input)
		expected := map[int]bool{1: true, 2: true, 3: true}

		if !reflect.DeepEqual(result, expected) {
			t.Errorf("Slice2Map() = %v, want %v", result, expected)
		}
	})

	t.Run("empty slice", func(t *testing.T) {
		input := []string{}
		result := Slice2Map(input)
		expected := map[string]bool{}

		if !reflect.DeepEqual(result, expected) {
			t.Errorf("Slice2Map() = %v, want %v", result, expected)
		}
	})

	t.Run("slice with duplicates", func(t *testing.T) {
		input := []string{"a", "b", "a", "c"}
		result := Slice2Map(input)
		expected := map[string]bool{"a": true, "b": true, "c": true}

		if !reflect.DeepEqual(result, expected) {
			t.Errorf("Slice2Map() = %v, want %v", result, expected)
		}
	})

	t.Run("float64 slice", func(t *testing.T) {
		input := []float64{1.1, 2.2, 3.3}
		result := Slice2Map(input)
		expected := map[float64]bool{1.1: true, 2.2: true, 3.3: true}

		if !reflect.DeepEqual(result, expected) {
			t.Errorf("Slice2Map() = %v, want %v", result, expected)
		}
	})
}

func TestSliceField2MapString(t *testing.T) {
	t.Run("extract string field", func(t *testing.T) {
		input := []TestItem{
			{Name: "Alice", Code: "A001"},
			{Name: "Bob", Code: "B001"},
			{Name: "Charlie", Code: "C001"},
		}
		result := SliceField2MapString(input, "Name")
		expected := map[string]bool{"Alice": true, "Bob": true, "Charlie": true}

		if !reflect.DeepEqual(result, expected) {
			t.Errorf("SliceField2MapString() = %v, want %v", result, expected)
		}
	})

	t.Run("extract with duplicates", func(t *testing.T) {
		input := []TestItem{
			{Name: "Alice"},
			{Name: "Bob"},
			{Name: "Alice"},
		}
		result := SliceField2MapString(input, "Name")
		expected := map[string]bool{"Alice": true, "Bob": true}

		if !reflect.DeepEqual(result, expected) {
			t.Errorf("SliceField2MapString() = %v, want %v", result, expected)
		}
	})

	t.Run("empty slice", func(t *testing.T) {
		input := []TestItem{}
		result := SliceField2MapString(input, "Name")

		if result != nil {
			t.Errorf("SliceField2MapString() = %v, want nil", result)
		}
	})

	t.Run("with pointer slice", func(t *testing.T) {
		input := []*TestItem{
			{Name: "Alice"},
			{Name: "Bob"},
		}
		result := SliceField2MapString(input, "Name")
		expected := map[string]bool{"Alice": true, "Bob": true}

		if !reflect.DeepEqual(result, expected) {
			t.Errorf("SliceField2MapString() = %v, want %v", result, expected)
		}
	})
}

func TestSliceField2MapInt(t *testing.T) {
	t.Run("extract int field", func(t *testing.T) {
		input := []TestItem{
			{ID: 1, Name: "Alice"},
			{ID: 2, Name: "Bob"},
			{ID: 3, Name: "Charlie"},
		}
		result := SliceField2MapInt(input, "ID")
		expected := map[int]bool{1: true, 2: true, 3: true}

		if !reflect.DeepEqual(result, expected) {
			t.Errorf("SliceField2MapInt() = %v, want %v", result, expected)
		}
	})

	t.Run("extract with duplicates", func(t *testing.T) {
		input := []TestItem{
			{ID: 1},
			{ID: 2},
			{ID: 1},
		}
		result := SliceField2MapInt(input, "ID")
		expected := map[int]bool{1: true, 2: true}

		if !reflect.DeepEqual(result, expected) {
			t.Errorf("SliceField2MapInt() = %v, want %v", result, expected)
		}
	})

	t.Run("empty slice", func(t *testing.T) {
		input := []TestItem{}
		result := SliceField2MapInt(input, "ID")

		if result != nil {
			t.Errorf("SliceField2MapInt() = %v, want nil", result)
		}
	})
}

func TestSliceField2MapInt8(t *testing.T) {
	t.Run("extract int8 field", func(t *testing.T) {
		input := []TestItem{
			{ID8: 1},
			{ID8: 2},
			{ID8: 3},
		}
		result := SliceField2MapInt8(input, "ID8")
		expected := map[int8]bool{1: true, 2: true, 3: true}

		if !reflect.DeepEqual(result, expected) {
			t.Errorf("SliceField2MapInt8() = %v, want %v", result, expected)
		}
	})

	t.Run("empty slice", func(t *testing.T) {
		input := []TestItem{}
		result := SliceField2MapInt8(input, "ID8")

		if result != nil {
			t.Errorf("SliceField2MapInt8() = %v, want nil", result)
		}
	})
}

func TestSliceField2MapInt16(t *testing.T) {
	t.Run("extract int16 field", func(t *testing.T) {
		input := []TestItem{
			{ID16: 100},
			{ID16: 200},
			{ID16: 300},
		}
		result := SliceField2MapInt16(input, "ID16")
		expected := map[int16]bool{100: true, 200: true, 300: true}

		if !reflect.DeepEqual(result, expected) {
			t.Errorf("SliceField2MapInt16() = %v, want %v", result, expected)
		}
	})

	t.Run("empty slice", func(t *testing.T) {
		input := []TestItem{}
		result := SliceField2MapInt16(input, "ID16")

		if result != nil {
			t.Errorf("SliceField2MapInt16() = %v, want nil", result)
		}
	})
}

func TestSliceField2MapInt32(t *testing.T) {
	t.Run("extract int32 field", func(t *testing.T) {
		input := []TestItem{
			{ID32: 1000},
			{ID32: 2000},
			{ID32: 3000},
		}
		result := SliceField2MapInt32(input, "ID32")
		expected := map[int32]bool{1000: true, 2000: true, 3000: true}

		if !reflect.DeepEqual(result, expected) {
			t.Errorf("SliceField2MapInt32() = %v, want %v", result, expected)
		}
	})

	t.Run("empty slice", func(t *testing.T) {
		input := []TestItem{}
		result := SliceField2MapInt32(input, "ID32")

		if result != nil {
			t.Errorf("SliceField2MapInt32() = %v, want nil", result)
		}
	})
}

func TestSliceField2MapInt64(t *testing.T) {
	t.Run("extract int64 field", func(t *testing.T) {
		input := []TestItem{
			{ID64: 10000},
			{ID64: 20000},
			{ID64: 30000},
		}
		result := SliceField2MapInt64(input, "ID64")
		expected := map[int64]bool{10000: true, 20000: true, 30000: true}

		if !reflect.DeepEqual(result, expected) {
			t.Errorf("SliceField2MapInt64() = %v, want %v", result, expected)
		}
	})

	t.Run("empty slice", func(t *testing.T) {
		input := []TestItem{}
		result := SliceField2MapInt64(input, "ID64")

		if result != nil {
			t.Errorf("SliceField2MapInt64() = %v, want nil", result)
		}
	})
}

func TestSliceField2MapUint(t *testing.T) {
	t.Run("extract uint field", func(t *testing.T) {
		input := []TestItem{
			{UID: 1},
			{UID: 2},
			{UID: 3},
		}
		result := SliceField2MapUint(input, "UID")
		expected := map[uint]bool{1: true, 2: true, 3: true}

		if !reflect.DeepEqual(result, expected) {
			t.Errorf("SliceField2MapUint() = %v, want %v", result, expected)
		}
	})

	t.Run("empty slice", func(t *testing.T) {
		input := []TestItem{}
		result := SliceField2MapUint(input, "UID")

		if result != nil {
			t.Errorf("SliceField2MapUint() = %v, want nil", result)
		}
	})
}

func TestSliceField2MapUint8(t *testing.T) {
	t.Run("extract uint8 field", func(t *testing.T) {
		input := []TestItem{
			{UID8: 1},
			{UID8: 2},
			{UID8: 3},
		}
		result := SliceField2MapUint8(input, "UID8")
		expected := map[uint8]bool{1: true, 2: true, 3: true}

		if !reflect.DeepEqual(result, expected) {
			t.Errorf("SliceField2MapUint8() = %v, want %v", result, expected)
		}
	})

	t.Run("empty slice", func(t *testing.T) {
		input := []TestItem{}
		result := SliceField2MapUint8(input, "UID8")

		if result != nil {
			t.Errorf("SliceField2MapUint8() = %v, want nil", result)
		}
	})
}

func TestSliceField2MapUint16(t *testing.T) {
	t.Run("extract uint16 field", func(t *testing.T) {
		input := []TestItem{
			{UID16: 100},
			{UID16: 200},
			{UID16: 300},
		}
		result := SliceField2MapUint16(input, "UID16")
		expected := map[uint16]bool{100: true, 200: true, 300: true}

		if !reflect.DeepEqual(result, expected) {
			t.Errorf("SliceField2MapUint16() = %v, want %v", result, expected)
		}
	})

	t.Run("empty slice", func(t *testing.T) {
		input := []TestItem{}
		result := SliceField2MapUint16(input, "UID16")

		if result != nil {
			t.Errorf("SliceField2MapUint16() = %v, want nil", result)
		}
	})
}

func TestSliceField2MapUint32(t *testing.T) {
	t.Run("extract uint32 field", func(t *testing.T) {
		input := []TestItem{
			{UID32: 1000},
			{UID32: 2000},
			{UID32: 3000},
		}
		result := SliceField2MapUint32(input, "UID32")
		expected := map[uint32]bool{1000: true, 2000: true, 3000: true}

		if !reflect.DeepEqual(result, expected) {
			t.Errorf("SliceField2MapUint32() = %v, want %v", result, expected)
		}
	})

	t.Run("empty slice", func(t *testing.T) {
		input := []TestItem{}
		result := SliceField2MapUint32(input, "UID32")

		if result != nil {
			t.Errorf("SliceField2MapUint32() = %v, want nil", result)
		}
	})
}

func TestSliceField2MapUint64(t *testing.T) {
	t.Run("extract uint64 field", func(t *testing.T) {
		input := []TestItem{
			{UID64: 10000},
			{UID64: 20000},
			{UID64: 30000},
		}
		result := SliceField2MapUint64(input, "UID64")
		expected := map[uint64]bool{10000: true, 20000: true, 30000: true}

		if !reflect.DeepEqual(result, expected) {
			t.Errorf("SliceField2MapUint64() = %v, want %v", result, expected)
		}
	})

	t.Run("empty slice", func(t *testing.T) {
		input := []TestItem{}
		result := SliceField2MapUint64(input, "UID64")

		if result != nil {
			t.Errorf("SliceField2MapUint64() = %v, want nil", result)
		}
	})
}

func TestSliceField2MapFloat32(t *testing.T) {
	t.Run("extract float32 field", func(t *testing.T) {
		input := []TestItem{
			{Price: 1.5},
			{Price: 2.5},
			{Price: 3.5},
		}
		result := SliceField2MapFloat32(input, "Price")
		expected := map[float32]bool{1.5: true, 2.5: true, 3.5: true}

		if !reflect.DeepEqual(result, expected) {
			t.Errorf("SliceField2MapFloat32() = %v, want %v", result, expected)
		}
	})

	t.Run("empty slice", func(t *testing.T) {
		input := []TestItem{}
		result := SliceField2MapFloat32(input, "Price")

		if result != nil {
			t.Errorf("SliceField2MapFloat32() = %v, want nil", result)
		}
	})
}

func TestSliceField2MapFloat64(t *testing.T) {
	t.Run("extract float64 field", func(t *testing.T) {
		input := []TestItem{
			{Score: 90.5},
			{Score: 85.5},
			{Score: 95.5},
		}
		result := SliceField2MapFloat64(input, "Score")
		expected := map[float64]bool{90.5: true, 85.5: true, 95.5: true}

		if !reflect.DeepEqual(result, expected) {
			t.Errorf("SliceField2MapFloat64() = %v, want %v", result, expected)
		}
	})

	t.Run("empty slice", func(t *testing.T) {
		input := []TestItem{}
		result := SliceField2MapFloat64(input, "Score")

		if result != nil {
			t.Errorf("SliceField2MapFloat64() = %v, want nil", result)
		}
	})
}

func TestSliceField2MapBool(t *testing.T) {
	t.Run("extract bool field", func(t *testing.T) {
		input := []TestItem{
			{Active: true},
			{Active: false},
			{Active: true},
		}
		result := SliceField2MapBool(input, "Active")
		expected := map[bool]bool{true: true, false: true}

		if !reflect.DeepEqual(result, expected) {
			t.Errorf("SliceField2MapBool() = %v, want %v", result, expected)
		}
	})

	t.Run("empty slice", func(t *testing.T) {
		input := []TestItem{}
		result := SliceField2MapBool(input, "Active")

		if result != nil {
			t.Errorf("SliceField2MapBool() = %v, want nil", result)
		}
	})
}

func TestSliceField2MapPanic(t *testing.T) {
	t.Run("field not found", func(t *testing.T) {
		defer func() {
			if r := recover(); r == nil {
				t.Errorf("Expected panic for non-existent field")
			}
		}()

		input := []TestItem{{Name: "Alice"}}
		SliceField2MapString(input, "NonExistent")
	})

	t.Run("wrong field type", func(t *testing.T) {
		defer func() {
			if r := recover(); r == nil {
				t.Errorf("Expected panic for wrong field type")
			}
		}()

		input := []TestItem{{ID: 1}}
		SliceField2MapString(input, "ID") // ID is int, not string
	})

	t.Run("not a struct", func(t *testing.T) {
		defer func() {
			if r := recover(); r == nil {
				t.Errorf("Expected panic for non-struct input")
			}
		}()

		input := []string{"a", "b"}
		SliceField2MapString(input, "Name")
	})
}
