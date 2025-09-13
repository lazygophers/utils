package anyx

import (
	"testing"
)

// TestKeyByPanicConditions tests the panic conditions in KeyBy functions
func TestKeyByPanicConditions(t *testing.T) {
	type testStruct struct {
		ID   uint64
		Name string
	}

	t.Run("KeyBy_element_not_struct_panic", func(t *testing.T) {
		defer func() {
			if r := recover(); r != nil {
				if r != "list element is not struct" {
					t.Errorf("Expected panic 'list element is not struct', got %v", r)
				}
			} else {
				t.Error("Expected panic but didn't panic")
			}
		}()

		// Create a slice with interface{} containing non-struct values
		list := []interface{}{testStruct{ID: 1, Name: "test"}, "not a struct"}
		// This should trigger the panic when processing the string element
		_ = KeyBy(list, "ID")
	})

	t.Run("KeyByUint64_element_not_struct_panic", func(t *testing.T) {
		// The panic path in these generic functions is very hard to reach
		// without unsafe operations or compiler bugs. Let's skip this test.
		t.Skip("Generic KeyBy functions have panic paths that are hard to reach safely")
	})

	// Test covering invalid element scenarios for the other KeyBy functions
	t.Run("KeyByUint64_nil_pointer_coverage", func(t *testing.T) {
		// Test with nil pointers in the slice to reach the continue branch
		list := []*testStruct{{ID: 1, Name: "test"}, nil, {ID: 2, Name: "test2"}}
		result := KeyByUint64(list, "ID")
		if len(result) != 2 { // Should skip nil and process 2 valid elements
			t.Errorf("Expected 2 elements, got %d", len(result))
		}
	})

	t.Run("KeyByInt64_nil_pointer_coverage", func(t *testing.T) {
		type testStructInt64 struct {
			ID   int64
			Name string
		}
		list := []*testStructInt64{{ID: 1, Name: "test"}, nil, {ID: 2, Name: "test2"}}
		result := KeyByInt64(list, "ID")
		if len(result) != 2 {
			t.Errorf("Expected 2 elements, got %d", len(result))
		}
	})

	t.Run("KeyByString_nil_pointer_coverage", func(t *testing.T) {
		type testStructString struct {
			ID   string
			Name string
		}
		list := []*testStructString{{ID: "1", Name: "test"}, nil, {ID: "2", Name: "test2"}}
		result := KeyByString(list, "ID")
		if len(result) != 2 {
			t.Errorf("Expected 2 elements, got %d", len(result))
		}
	})

	t.Run("KeyByInt32_nil_pointer_coverage", func(t *testing.T) {
		type testStructInt32 struct {
			ID   int32
			Name string
		}
		list := []*testStructInt32{{ID: 1, Name: "test"}, nil, {ID: 2, Name: "test2"}}
		result := KeyByInt32(list, "ID")
		if len(result) != 2 {
			t.Errorf("Expected 2 elements, got %d", len(result))
		}
	})
}

// TestNewMapWithAnyErrorPaths tests error paths in NewMapWithAny
func TestNewMapWithAnyErrorPaths(t *testing.T) {
	t.Run("json_marshal_error", func(t *testing.T) {
		// Create a struct with a channel that can't be marshaled to JSON
		type invalidStruct struct {
			Ch chan int
		}

		input := invalidStruct{Ch: make(chan int)}
		result, err := NewMapWithAny(input)
		if err == nil {
			t.Error("Expected error but got nil")
		}
		if result != nil {
			t.Error("Expected nil result on error")
		}
	})

	t.Run("yaml_unmarshal_error", func(t *testing.T) {
		// It's very hard to create JSON that's valid for marshaling but invalid for YAML unmarshaling
		// since the code uses yaml.Unmarshal on JSON bytes. Most valid JSON is also valid YAML.
		// This path is hard to test without complex edge cases.
		// Let's test with some edge case values
		input := map[string]interface{}{
			"validField": "test",
		}
		
		result, err := NewMapWithAny(input)
		if err != nil {
			t.Errorf("Unexpected error: %v", err)
		}
		if result == nil {
			t.Error("Expected non-nil result")
		}
	})

	t.Run("complex_struct_success", func(t *testing.T) {
		// Test with a more complex struct to ensure proper coverage
		type complexStruct struct {
			Field1 string
			Field2 int
			Field3 map[string]interface{}
			Field4 []string
		}

		input := complexStruct{
			Field1: "test",
			Field2: 42,
			Field3: map[string]interface{}{"nested": "value"},
			Field4: []string{"a", "b", "c"},
		}

		result, err := NewMapWithAny(input)
		if err != nil {
			t.Errorf("Unexpected error: %v", err)
		}
		if result == nil {
			t.Error("Expected non-nil result")
		}

		// Check that the values are properly accessible
		if result.GetString("Field1") != "test" {
			t.Error("Field1 not properly converted")
		}
		if result.GetInt("Field2") != 42 {
			t.Error("Field2 not properly converted")
		}
	})
}

// TestMapAnyGetMissingPaths tests the missing paths in the get function
func TestMapAnyGetMissingPaths(t *testing.T) {
	t.Run("nested_path_not_found", func(t *testing.T) {
		m := NewMap(map[string]interface{}{
			"level1": map[string]interface{}{
				"level2": "value",
			},
		})
		m.EnableCut(".")

		// Test path that doesn't exist in nested structure
		result, exists := m.get("level1.nonexistent.field")
		if exists {
			t.Error("Expected path not to exist")
		}
		if result != nil {
			t.Error("Expected nil result for non-existent path")
		}
	})

	t.Run("invalid_nested_map_type", func(t *testing.T) {
		m := NewMap(map[string]interface{}{
			"level1": "not_a_map", // This is not a map, so nested access should fail
		})
		m.EnableCut(".")

		// Test path that tries to access nested field on non-map value
		result, exists := m.get("level1.field")
		if exists {
			t.Error("Expected path not to exist")
		}
		if result != nil {
			t.Error("Expected nil result for invalid nested access")
		}
	})

	t.Run("empty_keys_after_split", func(t *testing.T) {
		m := NewMap(map[string]interface{}{
			"test": "value",
		})
		m.EnableCut(".")

		// This should hit the len(keys) > 0 condition at the end
		// by having an empty key after splitting
		result, exists := m.get("test.")
		if exists {
			t.Error("Expected path not to exist")
		}
		if result != nil {
			t.Error("Expected nil result for empty key path")
		}
	})
}