package candy

import (
	"reflect"
	"testing"
)

func TestToMap(t *testing.T) {
	t.Run("byte slice with valid JSON", func(t *testing.T) {
		input := []byte(`{"name":"test","value":42}`)
		result := ToMap(input)

		if result == nil {
			t.Error("ToMap(valid JSON bytes) returned nil")
		}
		if result["name"] != "test" {
			t.Errorf("ToMap result[\"name\"] = %v, want \"test\"", result["name"])
		}
		if val, ok := result["value"].(float64); !ok || val != 42 {
			t.Errorf("ToMap result[\"value\"] = %v, want 42", result["value"])
		}
	})

	t.Run("byte slice with invalid JSON", func(t *testing.T) {
		input := []byte("not json")
		result := ToMap(input)

		// Should fall back to ToMapStringAny which returns empty map for non-map type
		if result == nil {
			t.Error("ToMap(invalid JSON bytes) returned nil")
		}
	})

	t.Run("string with valid JSON", func(t *testing.T) {
		input := `{"key":"value","num":123}`
		result := ToMap(input)

		if result == nil {
			t.Error("ToMap(valid JSON string) returned nil")
		}
		if result["key"] != "value" {
			t.Errorf("ToMap result[\"key\"] = %v, want \"value\"", result["key"])
		}
	})

	t.Run("string with invalid JSON", func(t *testing.T) {
		input := "not json"
		result := ToMap(input)

		// Should fall back to ToMapStringAny
		if result == nil {
			t.Error("ToMap(invalid JSON string) returned nil")
		}
	})

	t.Run("map input", func(t *testing.T) {
		input := map[string]interface{}{"a": 1, "b": "test"}
		result := ToMap(input)

		if !reflect.DeepEqual(result, input) {
			t.Errorf("ToMap(map) = %v, want %v", result, input)
		}
	})

	t.Run("nil input", func(t *testing.T) {
		result := ToMap(nil)
		if result != nil {
			t.Errorf("ToMap(nil) = %v, want nil", result)
		}
	})

	t.Run("other types", func(t *testing.T) {
		result := ToMap(42)
		if result == nil {
			t.Error("ToMap(int) should return empty map, not nil")
		}
	})
}

func TestToMapInt32String(t *testing.T) {
	t.Run("map with int keys", func(t *testing.T) {
		input := map[int]string{1: "one", 2: "two", 3: "three"}
		result := ToMapInt32String(input)

		expected := map[int32]string{1: "one", 2: "two", 3: "three"}
		if !reflect.DeepEqual(result, expected) {
			t.Errorf("ToMapInt32String(%v) = %v, want %v", input, result, expected)
		}
	})

	t.Run("map with string keys", func(t *testing.T) {
		input := map[string]int{"1": 100, "2": 200}
		result := ToMapInt32String(input)

		expected := map[int32]string{1: "100", 2: "200"}
		if !reflect.DeepEqual(result, expected) {
			t.Errorf("ToMapInt32String(%v) = %v, want %v", input, result, expected)
		}
	})

	t.Run("map with float keys", func(t *testing.T) {
		input := map[float64]string{1.5: "a", 2.7: "b"}
		result := ToMapInt32String(input)

		expected := map[int32]string{1: "a", 2: "b"}
		if !reflect.DeepEqual(result, expected) {
			t.Errorf("ToMapInt32String(%v) = %v, want %v", input, result, expected)
		}
	})

	t.Run("map with bool values", func(t *testing.T) {
		input := map[int]bool{1: true, 2: false}
		result := ToMapInt32String(input)

		expected := map[int32]string{1: "1", 2: "0"}
		if !reflect.DeepEqual(result, expected) {
			t.Errorf("ToMapInt32String(%v) = %v, want %v", input, result, expected)
		}
	})

	t.Run("empty map", func(t *testing.T) {
		input := map[int]string{}
		result := ToMapInt32String(input)

		if len(result) != 0 {
			t.Errorf("ToMapInt32String(empty map) length = %d, want 0", len(result))
		}
	})

	t.Run("non-map input", func(t *testing.T) {
		result := ToMapInt32String("not a map")
		expected := map[int32]string{}

		if !reflect.DeepEqual(result, expected) {
			t.Errorf("ToMapInt32String(non-map) = %v, want %v", result, expected)
		}
	})

	t.Run("nil input", func(t *testing.T) {
		result := ToMapInt32String(nil)
		expected := map[int32]string{}

		if !reflect.DeepEqual(result, expected) {
			t.Errorf("ToMapInt32String(nil) = %v, want %v", result, expected)
		}
	})
}

func TestToMapInt64String(t *testing.T) {
	t.Run("map with int keys", func(t *testing.T) {
		input := map[int]string{1: "one", 2: "two", 3: "three"}
		result := ToMapInt64String(input)

		expected := map[int64]string{1: "one", 2: "two", 3: "three"}
		if !reflect.DeepEqual(result, expected) {
			t.Errorf("ToMapInt64String(%v) = %v, want %v", input, result, expected)
		}
	})

	t.Run("map with string keys", func(t *testing.T) {
		input := map[string]int{"10": 100, "20": 200}
		result := ToMapInt64String(input)

		expected := map[int64]string{10: "100", 20: "200"}
		if !reflect.DeepEqual(result, expected) {
			t.Errorf("ToMapInt64String(%v) = %v, want %v", input, result, expected)
		}
	})

	t.Run("map with large int64 keys", func(t *testing.T) {
		input := map[int64]string{9223372036854775807: "max"}
		result := ToMapInt64String(input)

		expected := map[int64]string{9223372036854775807: "max"}
		if !reflect.DeepEqual(result, expected) {
			t.Errorf("ToMapInt64String(%v) = %v, want %v", input, result, expected)
		}
	})

	t.Run("empty map", func(t *testing.T) {
		input := map[int]string{}
		result := ToMapInt64String(input)

		if len(result) != 0 {
			t.Errorf("ToMapInt64String(empty map) length = %d, want 0", len(result))
		}
	})

	t.Run("non-map input", func(t *testing.T) {
		result := ToMapInt64String(123)
		expected := map[int64]string{}

		if !reflect.DeepEqual(result, expected) {
			t.Errorf("ToMapInt64String(non-map) = %v, want %v", result, expected)
		}
	})

	t.Run("nil input", func(t *testing.T) {
		result := ToMapInt64String(nil)
		expected := map[int64]string{}

		if !reflect.DeepEqual(result, expected) {
			t.Errorf("ToMapInt64String(nil) = %v, want %v", result, expected)
		}
	})
}

func TestToMapStringAny(t *testing.T) {
	t.Run("map with various types", func(t *testing.T) {
		input := map[string]interface{}{"a": 1, "b": "test", "c": true}
		result := ToMapStringAny(input)

		if !reflect.DeepEqual(result, input) {
			t.Errorf("ToMapStringAny(%v) = %v, want %v", input, result, input)
		}
	})

	t.Run("map with int keys", func(t *testing.T) {
		input := map[int]string{1: "one", 2: "two"}
		result := ToMapStringAny(input)

		expected := map[string]interface{}{"1": "one", "2": "two"}
		if !reflect.DeepEqual(result, expected) {
			t.Errorf("ToMapStringAny(%v) = %v, want %v", input, result, expected)
		}
	})

	t.Run("empty map", func(t *testing.T) {
		input := map[string]interface{}{}
		result := ToMapStringAny(input)

		if len(result) != 0 {
			t.Errorf("ToMapStringAny(empty map) length = %d, want 0", len(result))
		}
	})

	t.Run("nil input", func(t *testing.T) {
		result := ToMapStringAny(nil)

		if result != nil {
			t.Errorf("ToMapStringAny(nil) = %v, want nil", result)
		}
	})

	t.Run("non-map input", func(t *testing.T) {
		result := ToMapStringAny("not a map")
		expected := map[string]interface{}{}

		if !reflect.DeepEqual(result, expected) {
			t.Errorf("ToMapStringAny(non-map) = %v, want %v", result, expected)
		}
	})

	t.Run("map with complex values", func(t *testing.T) {
		input := map[string]interface{}{
			"slice": []int{1, 2, 3},
			"map":   map[string]int{"nested": 42},
		}
		result := ToMapStringAny(input)

		if !reflect.DeepEqual(result, input) {
			t.Errorf("ToMapStringAny(complex) = %v, want %v", result, input)
		}
	})
}

func TestToMapStringArrayString(t *testing.T) {
	t.Run("map with string slice values", func(t *testing.T) {
		input := map[string][]string{"a": {"1", "2"}, "b": {"3", "4"}}
		result := ToMapStringArrayString(input)

		if !reflect.DeepEqual(result, input) {
			t.Errorf("ToMapStringArrayString(%v) = %v, want %v", input, result, input)
		}
	})

	t.Run("map with int slice values", func(t *testing.T) {
		input := map[string][]int{"a": {1, 2}, "b": {3, 4}}
		result := ToMapStringArrayString(input)

		expected := map[string][]string{"a": {"1", "2"}, "b": {"3", "4"}}
		if !reflect.DeepEqual(result, expected) {
			t.Errorf("ToMapStringArrayString(%v) = %v, want %v", input, result, expected)
		}
	})

	t.Run("map with single values converted to slice", func(t *testing.T) {
		input := map[string]int{"a": 1, "b": 2}
		result := ToMapStringArrayString(input)

		expected := map[string][]string{"a": {"1"}, "b": {"2"}}
		if !reflect.DeepEqual(result, expected) {
			t.Errorf("ToMapStringArrayString(%v) = %v, want %v", input, result, expected)
		}
	})

	t.Run("map with int keys", func(t *testing.T) {
		input := map[int][]string{1: {"a", "b"}, 2: {"c", "d"}}
		result := ToMapStringArrayString(input)

		expected := map[string][]string{"1": {"a", "b"}, "2": {"c", "d"}}
		if !reflect.DeepEqual(result, expected) {
			t.Errorf("ToMapStringArrayString(%v) = %v, want %v", input, result, expected)
		}
	})

	t.Run("empty map", func(t *testing.T) {
		input := map[string][]string{}
		result := ToMapStringArrayString(input)

		if len(result) != 0 {
			t.Errorf("ToMapStringArrayString(empty map) length = %d, want 0", len(result))
		}
	})

	t.Run("nil input", func(t *testing.T) {
		result := ToMapStringArrayString(nil)

		if result != nil {
			t.Errorf("ToMapStringArrayString(nil) = %v, want nil", result)
		}
	})

	t.Run("non-map input panics", func(t *testing.T) {
		defer func() {
			if r := recover(); r == nil {
				t.Error("ToMapStringArrayString(non-map) should panic")
			}
		}()
		ToMapStringArrayString("not a map")
	})
}

func TestToMapStringInt64(t *testing.T) {
	t.Run("map with int values", func(t *testing.T) {
		input := map[string]int{"a": 1, "b": 2, "c": 3}
		result := ToMapStringInt64(input)

		expected := map[string]int64{"a": 1, "b": 2, "c": 3}
		if !reflect.DeepEqual(result, expected) {
			t.Errorf("ToMapStringInt64(%v) = %v, want %v", input, result, expected)
		}
	})

	t.Run("map with string values", func(t *testing.T) {
		input := map[string]string{"a": "100", "b": "200"}
		result := ToMapStringInt64(input)

		expected := map[string]int64{"a": 100, "b": 200}
		if !reflect.DeepEqual(result, expected) {
			t.Errorf("ToMapStringInt64(%v) = %v, want %v", input, result, expected)
		}
	})

	t.Run("map with float values", func(t *testing.T) {
		input := map[string]float64{"a": 1.5, "b": 2.7}
		result := ToMapStringInt64(input)

		expected := map[string]int64{"a": 1, "b": 2}
		if !reflect.DeepEqual(result, expected) {
			t.Errorf("ToMapStringInt64(%v) = %v, want %v", input, result, expected)
		}
	})

	t.Run("map with bool values", func(t *testing.T) {
		input := map[string]bool{"a": true, "b": false}
		result := ToMapStringInt64(input)

		expected := map[string]int64{"a": 1, "b": 0}
		if !reflect.DeepEqual(result, expected) {
			t.Errorf("ToMapStringInt64(%v) = %v, want %v", input, result, expected)
		}
	})

	t.Run("map with int keys", func(t *testing.T) {
		input := map[int]int{1: 100, 2: 200}
		result := ToMapStringInt64(input)

		expected := map[string]int64{"1": 100, "2": 200}
		if !reflect.DeepEqual(result, expected) {
			t.Errorf("ToMapStringInt64(%v) = %v, want %v", input, result, expected)
		}
	})

	t.Run("empty map", func(t *testing.T) {
		input := map[string]int{}
		result := ToMapStringInt64(input)

		if len(result) != 0 {
			t.Errorf("ToMapStringInt64(empty map) length = %d, want 0", len(result))
		}
	})

	t.Run("non-map input", func(t *testing.T) {
		result := ToMapStringInt64(42)
		expected := map[string]int64{}

		if !reflect.DeepEqual(result, expected) {
			t.Errorf("ToMapStringInt64(non-map) = %v, want %v", result, expected)
		}
	})

	t.Run("nil input", func(t *testing.T) {
		result := ToMapStringInt64(nil)
		expected := map[string]int64{}

		if !reflect.DeepEqual(result, expected) {
			t.Errorf("ToMapStringInt64(nil) = %v, want %v", result, expected)
		}
	})
}

func TestToMapStringString(t *testing.T) {
	t.Run("map with string keys and values", func(t *testing.T) {
		input := map[string]string{"a": "one", "b": "two"}
		result := ToMapStringString(input)

		if !reflect.DeepEqual(result, input) {
			t.Errorf("ToMapStringString(%v) = %v, want %v", input, result, input)
		}
	})

	t.Run("map with int values", func(t *testing.T) {
		input := map[string]int{"a": 1, "b": 2}
		result := ToMapStringString(input)

		expected := map[string]string{"a": "1", "b": "2"}
		if !reflect.DeepEqual(result, expected) {
			t.Errorf("ToMapStringString(%v) = %v, want %v", input, result, expected)
		}
	})

	t.Run("map with bool values", func(t *testing.T) {
		input := map[string]bool{"a": true, "b": false}
		result := ToMapStringString(input)

		expected := map[string]string{"a": "1", "b": "0"}
		if !reflect.DeepEqual(result, expected) {
			t.Errorf("ToMapStringString(%v) = %v, want %v", input, result, expected)
		}
	})

	t.Run("map with int keys", func(t *testing.T) {
		input := map[int]string{1: "one", 2: "two"}
		result := ToMapStringString(input)

		expected := map[string]string{"1": "one", "2": "two"}
		if !reflect.DeepEqual(result, expected) {
			t.Errorf("ToMapStringString(%v) = %v, want %v", input, result, expected)
		}
	})

	t.Run("map with float keys and values", func(t *testing.T) {
		input := map[float64]float64{1.5: 2.5, 3.7: 4.8}
		result := ToMapStringString(input)

		if len(result) != 2 {
			t.Errorf("ToMapStringString(float map) length = %d, want 2", len(result))
		}
	})

	t.Run("empty map", func(t *testing.T) {
		input := map[string]string{}
		result := ToMapStringString(input)

		if len(result) != 0 {
			t.Errorf("ToMapStringString(empty map) length = %d, want 0", len(result))
		}
	})

	t.Run("non-map input", func(t *testing.T) {
		result := ToMapStringString("not a map")
		expected := map[string]string{}

		if !reflect.DeepEqual(result, expected) {
			t.Errorf("ToMapStringString(non-map) = %v, want %v", result, expected)
		}
	})

	t.Run("nil input", func(t *testing.T) {
		result := ToMapStringString(nil)
		expected := map[string]string{}

		if !reflect.DeepEqual(result, expected) {
			t.Errorf("ToMapStringString(nil) = %v, want %v", result, expected)
		}
	})
}

func BenchmarkToMap(b *testing.B) {
	input := []byte(`{"name":"test","value":42}`)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		ToMap(input)
	}
}

func BenchmarkToMapInt32String(b *testing.B) {
	input := map[int]string{1: "one", 2: "two", 3: "three"}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		ToMapInt32String(input)
	}
}

func BenchmarkToMapStringString(b *testing.B) {
	input := map[string]int{"a": 1, "b": 2, "c": 3}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		ToMapStringString(input)
	}
}
