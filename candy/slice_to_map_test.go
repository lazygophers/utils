package candy

import (
	"reflect"
	"testing"
)

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

func TestSliceToMapWithValue(t *testing.T) {
	t.Run("nil slice", func(t *testing.T) {
		var input []string
		result := SliceToMapWithValue(input, 42)

		if result != nil {
			t.Errorf("SliceToMapWithValue() = %v, want nil", result)
		}
	})

	t.Run("empty slice", func(t *testing.T) {
		input := []string{}
		result := SliceToMapWithValue(input, 42)
		expected := map[string]int{}

		if !reflect.DeepEqual(result, expected) {
			t.Errorf("SliceToMapWithValue() = %v, want %v", result, expected)
		}
	})

	t.Run("string slice with int value", func(t *testing.T) {
		input := []string{"a", "b", "c"}
		result := SliceToMapWithValue(input, 10)
		expected := map[string]int{"a": 10, "b": 10, "c": 10}

		if !reflect.DeepEqual(result, expected) {
			t.Errorf("SliceToMapWithValue() = %v, want %v", result, expected)
		}
	})

	t.Run("int slice with string value", func(t *testing.T) {
		input := []int{1, 2, 3}
		result := SliceToMapWithValue(input, "value")
		expected := map[int]string{1: "value", 2: "value", 3: "value"}

		if !reflect.DeepEqual(result, expected) {
			t.Errorf("SliceToMapWithValue() = %v, want %v", result, expected)
		}
	})

	t.Run("slice with duplicates", func(t *testing.T) {
		input := []string{"a", "b", "a"}
		result := SliceToMapWithValue(input, 100)
		expected := map[string]int{"a": 100, "b": 100}

		if !reflect.DeepEqual(result, expected) {
			t.Errorf("SliceToMapWithValue() = %v, want %v", result, expected)
		}
	})

	t.Run("slice with struct value", func(t *testing.T) {
		type TestStruct struct {
			Name  string
			Value int
		}

		input := []int{1, 2}
		value := TestStruct{Name: "test", Value: 42}
		result := SliceToMapWithValue(input, value)
		expected := map[int]TestStruct{
			1: {Name: "test", Value: 42},
			2: {Name: "test", Value: 42},
		}

		if !reflect.DeepEqual(result, expected) {
			t.Errorf("SliceToMapWithValue() = %v, want %v", result, expected)
		}
	})

	t.Run("slice with pointer value", func(t *testing.T) {
		val := 999
		input := []string{"x", "y"}
		result := SliceToMapWithValue(input, &val)

		if len(result) != 2 {
			t.Errorf("SliceToMapWithValue() length = %d, want 2", len(result))
		}

		if result["x"] != &val || result["y"] != &val {
			t.Errorf("SliceToMapWithValue() pointers don't match expected value")
		}
	})

	t.Run("slice with boolean value", func(t *testing.T) {
		input := []int{5, 10, 15}
		result := SliceToMapWithValue(input, false)
		expected := map[int]bool{5: false, 10: false, 15: false}

		if !reflect.DeepEqual(result, expected) {
			t.Errorf("SliceToMapWithValue() = %v, want %v", result, expected)
		}
	})
}

func TestSliceToMapWithIndex(t *testing.T) {
	t.Run("nil slice", func(t *testing.T) {
		var input []string
		result := SliceToMapWithIndex(input)

		if result != nil {
			t.Errorf("SliceToMapWithIndex() = %v, want nil", result)
		}
	})

	t.Run("empty slice", func(t *testing.T) {
		input := []string{}
		result := SliceToMapWithIndex(input)
		expected := map[string]int{}

		if !reflect.DeepEqual(result, expected) {
			t.Errorf("SliceToMapWithIndex() = %v, want %v", result, expected)
		}
	})

	t.Run("string slice", func(t *testing.T) {
		input := []string{"a", "b", "c"}
		result := SliceToMapWithIndex(input)
		expected := map[string]int{"a": 0, "b": 1, "c": 2}

		if !reflect.DeepEqual(result, expected) {
			t.Errorf("SliceToMapWithIndex() = %v, want %v", result, expected)
		}
	})

	t.Run("integer slice", func(t *testing.T) {
		input := []int{10, 20, 30}
		result := SliceToMapWithIndex(input)
		expected := map[int]int{10: 0, 20: 1, 30: 2}

		if !reflect.DeepEqual(result, expected) {
			t.Errorf("SliceToMapWithIndex() = %v, want %v", result, expected)
		}
	})

	t.Run("slice with duplicates", func(t *testing.T) {
		input := []string{"a", "b", "a", "c"}
		result := SliceToMapWithIndex(input)
		expected := map[string]int{"a": 2, "b": 1, "c": 3}

		if !reflect.DeepEqual(result, expected) {
			t.Errorf("SliceToMapWithIndex() = %v, want %v", result, expected)
		}
	})

	t.Run("single element", func(t *testing.T) {
		input := []int{42}
		result := SliceToMapWithIndex(input)
		expected := map[int]int{42: 0}

		if !reflect.DeepEqual(result, expected) {
			t.Errorf("SliceToMapWithIndex() = %v, want %v", result, expected)
		}
	})

	t.Run("float64 slice", func(t *testing.T) {
		input := []float64{1.1, 2.2, 3.3}
		result := SliceToMapWithIndex(input)
		expected := map[float64]int{1.1: 0, 2.2: 1, 3.3: 2}

		if !reflect.DeepEqual(result, expected) {
			t.Errorf("SliceToMapWithIndex() = %v, want %v", result, expected)
		}
	})
}
