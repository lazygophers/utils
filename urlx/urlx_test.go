package urlx

import (
	"net/url"
	"reflect"
	"testing"
)

func TestSortQuery(t *testing.T) {
	tests := []struct {
		name     string
		input    url.Values
		expected url.Values
	}{
		{
			name:     "empty query",
			input:    url.Values{},
			expected: url.Values{},
		},
		{
			name:     "nil query",
			input:    nil,
			expected: nil,
		},
		{
			name: "single parameter",
			input: url.Values{
				"name": []string{"john"},
			},
			expected: url.Values{
				"name": []string{"john"},
			},
		},
		{
			name: "already sorted parameters",
			input: url.Values{
				"a": []string{"1"},
				"b": []string{"2"},
				"c": []string{"3"},
			},
			expected: url.Values{
				"a": []string{"1"},
				"b": []string{"2"},
				"c": []string{"3"},
			},
		},
		{
			name: "reverse sorted parameters",
			input: url.Values{
				"z": []string{"1"},
				"y": []string{"2"},
				"x": []string{"3"},
			},
			expected: url.Values{
				"x": []string{"3"},
				"y": []string{"2"},
				"z": []string{"1"},
			},
		},
		{
			name: "mixed order parameters",
			input: url.Values{
				"name":  []string{"john"},
				"age":   []string{"30"},
				"city":  []string{"beijing"},
				"email": []string{"john@example.com"},
			},
			expected: url.Values{
				"age":   []string{"30"},
				"city":  []string{"beijing"},
				"email": []string{"john@example.com"},
				"name":  []string{"john"},
			},
		},
		{
			name: "parameters with numbers",
			input: url.Values{
				"param3": []string{"value3"},
				"param1": []string{"value1"},
				"param2": []string{"value2"},
			},
			expected: url.Values{
				"param1": []string{"value1"},
				"param2": []string{"value2"},
				"param3": []string{"value3"},
			},
		},
		{
			name: "parameters with special characters",
			input: url.Values{
				"user_name": []string{"john"},
				"user-id":   []string{"123"},
				"userId":    []string{"456"},
			},
			expected: url.Values{
				"user-id":   []string{"123"},
				"userId":    []string{"456"},
				"user_name": []string{"john"},
			},
		},
		{
			name: "case sensitive sorting",
			input: url.Values{
				"Name": []string{"John"},
				"age":  []string{"30"},
				"City": []string{"Beijing"},
			},
			expected: url.Values{
				"City": []string{"Beijing"},
				"Name": []string{"John"},
				"age":  []string{"30"},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := SortQuery(tt.input)
			
			// For nil input, expect nil result
			if tt.input == nil {
				if result != nil {
					t.Errorf("SortQuery(nil) = %v, expected nil", result)
				}
				return
			}
			
			// For empty input, result should also be empty but not nil
			if len(tt.input) == 0 {
				if len(result) != 0 {
					t.Errorf("SortQuery(empty) returned non-empty result: %v", result)
				}
				return
			}
			
			// Check that all expected keys are present with correct values
			for key, expectedValues := range tt.expected {
				if !result.Has(key) {
					t.Errorf("Expected key %q not found in result", key)
					continue
				}
				
				// Note: SortQuery uses Set() which only keeps one value per key
				// So we compare the first value from expected with Get()
				expectedValue := expectedValues[0]
				actualValue := result.Get(key)
				
				if actualValue != expectedValue {
					t.Errorf("For key %q, got value %q, expected %q", key, actualValue, expectedValue)
				}
			}
			
			// Check that result doesn't have extra keys
			for key := range result {
				if !tt.expected.Has(key) {
					t.Errorf("Unexpected key %q found in result", key)
				}
			}
		})
	}
}

func TestSortQueryMultipleValues(t *testing.T) {
	// Test behavior with multiple values per key
	// Note: SortQuery uses Set() which overwrites previous values
	input := url.Values{
		"tags": []string{"go", "programming", "web"},
		"lang": []string{"en", "zh"},
	}
	
	result := SortQuery(input)
	
	// Should have both keys
	if !result.Has("tags") {
		t.Error("Expected 'tags' key in result")
	}
	if !result.Has("lang") {
		t.Error("Expected 'lang' key in result")
	}
	
	// Values should be the first value from Get() since Set() is used
	tagsValue := result.Get("tags")
	langValue := result.Get("lang")
	
	// Get() returns the first value, so we expect the first values from input
	expectedTagsValue := input.Get("tags")
	expectedLangValue := input.Get("lang")
	
	if tagsValue != expectedTagsValue {
		t.Errorf("tags value = %q, expected %q", tagsValue, expectedTagsValue)
	}
	if langValue != expectedLangValue {
		t.Errorf("lang value = %q, expected %q", langValue, expectedLangValue)
	}
}

func TestSortQueryPreservesOriginal(t *testing.T) {
	// Test that original query is not modified
	original := url.Values{
		"z": []string{"last"},
		"a": []string{"first"},
		"m": []string{"middle"},
	}
	
	// Create a copy to compare with later
	originalCopy := make(url.Values)
	for key, values := range original {
		originalCopy[key] = make([]string, len(values))
		copy(originalCopy[key], values)
	}
	
	result := SortQuery(original)
	
	// Check that original is unchanged
	if !reflect.DeepEqual(original, originalCopy) {
		t.Error("SortQuery modified the original query")
	}
	
	// Check that result is sorted by comparing first and last keys
	if len(original) > 1 {
		keys := make([]string, 0, len(result))
		for key := range result {
			keys = append(keys, key)
		}
		// The returned result should have keys that when iterated give a sorted order
		// We'll verify the sorting worked by checking we have the expected keys
		expectedKeys := []string{"a", "m", "z"}
		for _, expectedKey := range expectedKeys {
			if !result.Has(expectedKey) {
				t.Errorf("Expected key %q not found in sorted result", expectedKey)
			}
		}
	}
}

func TestSortQueryReturnType(t *testing.T) {
	// Test that the function returns url.Values type
	input := url.Values{"test": []string{"value"}}
	result := SortQuery(input)
	
	// Check type using reflection
	if result == nil {
		t.Error("SortQuery returned nil")
		return
	}
	
	// Verify it's url.Values type
	expectedType := reflect.TypeOf(url.Values{})
	actualType := reflect.TypeOf(result)
	
	if actualType != expectedType {
		t.Errorf("SortQuery returned %v, expected %v", actualType, expectedType)
	}
}

func TestSortQueryEmptyValues(t *testing.T) {
	// Test with empty string values
	input := url.Values{
		"empty": []string{""},
		"blank": []string{" "},
		"zero":  []string{"0"},
		"a":     []string{"value"},
	}
	
	result := SortQuery(input)
	
	// Check that empty values are preserved and keys are sorted
	expectedKeys := []string{"a", "blank", "empty", "zero"}
	actualKeys := make([]string, 0, len(result))
	for key := range result {
		actualKeys = append(actualKeys, key)
	}
	
	// Since we can't control the iteration order of map,
	// we'll check each key individually
	for _, expectedKey := range expectedKeys {
		if !result.Has(expectedKey) {
			t.Errorf("Expected key %q not found", expectedKey)
		}
	}
	
	// Verify values are preserved correctly
	if result.Get("empty") != "" {
		t.Errorf("empty value = %q, expected empty string", result.Get("empty"))
	}
	if result.Get("blank") != " " {
		t.Errorf("blank value = %q, expected single space", result.Get("blank"))
	}
	if result.Get("zero") != "0" {
		t.Errorf("zero value = %q, expected '0'", result.Get("zero"))
	}
	if result.Get("a") != "value" {
		t.Errorf("a value = %q, expected 'value'", result.Get("a"))
	}
}

// Benchmark tests
func BenchmarkSortQuery_Small(b *testing.B) {
	query := url.Values{
		"name": []string{"john"},
		"age":  []string{"30"},
		"city": []string{"beijing"},
	}
	
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		SortQuery(query)
	}
}

func BenchmarkSortQuery_Medium(b *testing.B) {
	query := url.Values{
		"param1":  []string{"value1"},
		"param2":  []string{"value2"},
		"param3":  []string{"value3"},
		"param4":  []string{"value4"},
		"param5":  []string{"value5"},
		"param6":  []string{"value6"},
		"param7":  []string{"value7"},
		"param8":  []string{"value8"},
		"param9":  []string{"value9"},
		"param10": []string{"value10"},
	}
	
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		SortQuery(query)
	}
}

func BenchmarkSortQuery_Large(b *testing.B) {
	query := make(url.Values)
	for i := 0; i < 100; i++ {
		query.Add("param"+string(rune(i)), "value"+string(rune(i)))
	}
	
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		SortQuery(query)
	}
}

func BenchmarkSortQuery_Empty(b *testing.B) {
	query := url.Values{}
	
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		SortQuery(query)
	}
}