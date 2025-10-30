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
