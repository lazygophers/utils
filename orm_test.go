package utils

import (
	"database/sql/driver"
	"fmt"
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// Test structs for testing
type TestStruct struct {
	Name string `json:"name" default:"default_name"`
	Age  int    `json:"age" default:"25"`
}

type ComplexStruct struct {
	ID     int            `json:"id" default:"1"`
	Data   TestStruct     `json:"data"`
	Items  []string       `json:"items"`
	Nested *TestStruct    `json:"nested,omitempty"`
	Map    map[string]int `json:"map"`
}

func TestScan(t *testing.T) {
	tests := []struct {
		name        string
		src         interface{}
		dst         interface{}
		expectError bool
		validate    func(t *testing.T, dst interface{})
	}{
		{
			name: "scan_valid_json_object",
			src:  `{"name":"John","age":30}`,
			dst:  &TestStruct{},
			validate: func(t *testing.T, dst interface{}) {
				s := dst.(*TestStruct)
				assert.Equal(t, "John", s.Name)
				assert.Equal(t, 30, s.Age)
			},
		},
		{
			name: "scan_valid_json_array",
			src:  `[{"name":"John","age":30},{"name":"Jane","age":25}]`,
			dst:  &[]TestStruct{},
			validate: func(t *testing.T, dst interface{}) {
				s := dst.(*[]TestStruct)
				require.Len(t, *s, 2)
				assert.Equal(t, "John", (*s)[0].Name)
				assert.Equal(t, 30, (*s)[0].Age)
				assert.Equal(t, "Jane", (*s)[1].Name)
				assert.Equal(t, 25, (*s)[1].Age)
			},
		},
		{
			name: "scan_json_bytes",
			src:  []byte(`{"name":"Alice","age":35}`),
			dst:  &TestStruct{},
			validate: func(t *testing.T, dst interface{}) {
				s := dst.(*TestStruct)
				assert.Equal(t, "Alice", s.Name)
				assert.Equal(t, 35, s.Age)
			},
		},
		{
			name: "scan_empty_string_sets_defaults",
			src:  "",
			dst:  &TestStruct{},
			validate: func(t *testing.T, dst interface{}) {
				s := dst.(*TestStruct)
				assert.Equal(t, "default_name", s.Name)
				assert.Equal(t, 25, s.Age)
			},
		},
		{
			name: "scan_empty_bytes_sets_defaults",
			src:  []byte{},
			dst:  &TestStruct{},
			validate: func(t *testing.T, dst interface{}) {
				s := dst.(*TestStruct)
				assert.Equal(t, "default_name", s.Name)
				assert.Equal(t, 25, s.Age)
			},
		},
		{
			name: "scan_complex_object",
			src:  `{"id":123,"data":{"name":"Test","age":40},"items":["a","b","c"],"map":{"key1":1,"key2":2}}`,
			dst:  &ComplexStruct{},
			validate: func(t *testing.T, dst interface{}) {
				s := dst.(*ComplexStruct)
				assert.Equal(t, 123, s.ID)
				assert.Equal(t, "Test", s.Data.Name)
				assert.Equal(t, 40, s.Data.Age)
				assert.Equal(t, []string{"a", "b", "c"}, s.Items)
				assert.Equal(t, map[string]int{"key1": 1, "key2": 2}, s.Map)
			},
		},
		{
			name:        "scan_invalid_json_string",
			src:         `{"name":"John","age":}`,
			dst:         &TestStruct{},
			expectError: true,
		},
		{
			name:        "scan_invalid_json_bytes",
			src:         []byte(`{"invalid":json}`),
			dst:         &TestStruct{},
			expectError: true,
		},
		{
			name:        "scan_unsupported_type",
			src:         123,
			dst:         &TestStruct{},
			expectError: true,
		},
		{
			name:        "scan_float_type",
			src:         123.45,
			dst:         &TestStruct{},
			expectError: true,
		},
		{
			name:        "scan_bool_type",
			src:         true,
			dst:         &TestStruct{},
			expectError: true,
		},
		{
			name: "scan_json_with_escaped_chars",
			src:  `{"name":"John \"Doe\"","age":30}`,
			dst:  &TestStruct{},
			validate: func(t *testing.T, dst interface{}) {
				s := dst.(*TestStruct)
				assert.Equal(t, `John "Doe"`, s.Name)
				assert.Equal(t, 30, s.Age)
			},
		},
		{
			name: "scan_simple_string_as_json",
			src:  `"simple string"`,
			dst:  new(string),
			validate: func(t *testing.T, dst interface{}) {
				s := dst.(*string)
				assert.Equal(t, "simple string", *s)
			},
		},
		{
			name: "scan_number_as_json",
			src:  `42`,
			dst:  new(int),
			validate: func(t *testing.T, dst interface{}) {
				i := dst.(*int)
				assert.Equal(t, 42, *i)
			},
		},
		{
			name: "scan_single_character_valid_json",
			src:  `"a"`,
			dst:  new(string),
			validate: func(t *testing.T, dst interface{}) {
				s := dst.(*string)
				assert.Equal(t, "a", *s)
			},
		},
		{
			name:        "scan_single_character_buffer_invalid",
			src:         "a",
			dst:         new(string),
			expectError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := Scan(tt.src, tt.dst)
			if tt.expectError {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				if tt.validate != nil {
					tt.validate(t, tt.dst)
				}
			}
		})
	}
}

func TestValue(t *testing.T) {
	tests := []struct {
		name        string
		input       interface{}
		expectError bool
		validate    func(t *testing.T, value driver.Value)
	}{
		{
			name:  "value_struct",
			input: TestStruct{Name: "John", Age: 30},
			validate: func(t *testing.T, value driver.Value) {
				bytes, ok := value.([]byte)
				require.True(t, ok)
				assert.Contains(t, string(bytes), `"name":"John"`)
				assert.Contains(t, string(bytes), `"age":30`)
			},
		},
		{
			name:  "value_struct_pointer",
			input: &TestStruct{Name: "Jane", Age: 25},
			validate: func(t *testing.T, value driver.Value) {
				bytes, ok := value.([]byte)
				require.True(t, ok)
				assert.Contains(t, string(bytes), `"name":"Jane"`)
				assert.Contains(t, string(bytes), `"age":25`)
			},
		},
		{
			name:  "value_slice",
			input: []string{"a", "b", "c"},
			validate: func(t *testing.T, value driver.Value) {
				bytes, ok := value.([]byte)
				require.True(t, ok)
				assert.Equal(t, `["a","b","c"]`, string(bytes))
			},
		},
		{
			name:  "value_map",
			input: map[string]int{"key1": 1, "key2": 2},
			validate: func(t *testing.T, value driver.Value) {
				bytes, ok := value.([]byte)
				require.True(t, ok)
				// JSON map ordering can vary, so check both keys exist
				jsonStr := string(bytes)
				assert.Contains(t, jsonStr, `"key1":1`)
				assert.Contains(t, jsonStr, `"key2":2`)
			},
		},
		{
			name:  "value_string",
			input: "test string",
			validate: func(t *testing.T, value driver.Value) {
				bytes, ok := value.([]byte)
				require.True(t, ok)
				assert.Equal(t, `"test string"`, string(bytes))
			},
		},
		{
			name:  "value_int",
			input: 42,
			validate: func(t *testing.T, value driver.Value) {
				bytes, ok := value.([]byte)
				require.True(t, ok)
				assert.Equal(t, "42", string(bytes))
			},
		},
		{
			name:  "value_bool",
			input: true,
			validate: func(t *testing.T, value driver.Value) {
				bytes, ok := value.([]byte)
				require.True(t, ok)
				assert.Equal(t, "true", string(bytes))
			},
		},
		{
			name:  "value_nil",
			input: nil,
			validate: func(t *testing.T, value driver.Value) {
				bytes, ok := value.([]byte)
				require.True(t, ok)
				assert.Equal(t, "null", string(bytes))
			},
		},
		{
			name: "value_complex_struct",
			input: ComplexStruct{
				ID:    123,
				Data:  TestStruct{Name: "Test", Age: 40},
				Items: []string{"a", "b"},
				Map:   map[string]int{"key": 1},
			},
			validate: func(t *testing.T, value driver.Value) {
				bytes, ok := value.([]byte)
				require.True(t, ok)
				jsonStr := string(bytes)
				assert.Contains(t, jsonStr, `"id":123`)
				assert.Contains(t, jsonStr, `"name":"Test"`)
				assert.Contains(t, jsonStr, `"age":40`)
				assert.Contains(t, jsonStr, `["a","b"]`)
				assert.Contains(t, jsonStr, `"key":1`)
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			value, err := Value(tt.input)
			if tt.expectError {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				if tt.validate != nil {
					tt.validate(t, value)
				}
			}
		})
	}
}

// TestScanValue_RoundTrip tests that Value and Scan work together
func TestScanValue_RoundTrip(t *testing.T) {
	tests := []struct {
		name     string
		original interface{}
		target   interface{}
		validate func(t *testing.T, original, target interface{})
	}{
		{
			name:     "roundtrip_struct",
			original: TestStruct{Name: "John", Age: 30},
			target:   &TestStruct{},
			validate: func(t *testing.T, original, target interface{}) {
				orig := original.(TestStruct)
				tgt := target.(*TestStruct)
				assert.Equal(t, orig.Name, tgt.Name)
				assert.Equal(t, orig.Age, tgt.Age)
			},
		},
		{
			name:     "roundtrip_slice",
			original: []string{"a", "b", "c"},
			target:   &[]string{},
			validate: func(t *testing.T, original, target interface{}) {
				orig := original.([]string)
				tgt := target.(*[]string)
				assert.Equal(t, orig, *tgt)
			},
		},
		{
			name:     "roundtrip_map",
			original: map[string]int{"key1": 1, "key2": 2},
			target:   &map[string]int{},
			validate: func(t *testing.T, original, target interface{}) {
				orig := original.(map[string]int)
				tgt := target.(*map[string]int)
				assert.Equal(t, orig, *tgt)
			},
		},
		{
			name: "roundtrip_complex",
			original: ComplexStruct{
				ID:    123,
				Data:  TestStruct{Name: "Test", Age: 40},
				Items: []string{"x", "y"},
				Map:   map[string]int{"test": 42},
			},
			target: &ComplexStruct{},
			validate: func(t *testing.T, original, target interface{}) {
				orig := original.(ComplexStruct)
				tgt := target.(*ComplexStruct)
				assert.Equal(t, orig.ID, tgt.ID)
				assert.Equal(t, orig.Data.Name, tgt.Data.Name)
				assert.Equal(t, orig.Data.Age, tgt.Data.Age)
				assert.Equal(t, orig.Items, tgt.Items)
				assert.Equal(t, orig.Map, tgt.Map)
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Convert to driver value
			value, err := Value(tt.original)
			require.NoError(t, err)

			// Scan back to target
			err = Scan(value, tt.target)
			require.NoError(t, err)

			// Validate
			tt.validate(t, tt.original, tt.target)
		})
	}
}

// TestScan_EdgeCases tests edge cases for the Scan function
func TestScan_EdgeCases(t *testing.T) {
	t.Run("scan_nil_source", func(t *testing.T) {
		var target TestStruct
		err := Scan(nil, &target)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "unknown type")
	})

	t.Run("scan_channel_unsupported", func(t *testing.T) {
		ch := make(chan int)
		var target TestStruct
		err := Scan(ch, &target)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "unknown type")
	})

	t.Run("scan_func_unsupported", func(t *testing.T) {
		fn := func() {}
		var target TestStruct
		err := Scan(fn, &target)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "unknown type")
	})
}

// TestValue_EdgeCases tests edge cases for the Value function
func TestValue_EdgeCases(t *testing.T) {
	t.Run("value_channel_unsupported", func(t *testing.T) {
		ch := make(chan int)
		_, err := Value(ch)
		assert.Error(t, err)
	})

	t.Run("value_func_unsupported", func(t *testing.T) {
		fn := func() {}
		_, err := Value(fn)
		assert.Error(t, err)
	})
}

// TestScan_SpecialCases tests additional edge cases for comprehensive coverage
func TestScan_SpecialCases(t *testing.T) {
	t.Run("scan_empty_json_object", func(t *testing.T) {
		var target TestStruct
		err := Scan("{}", &target)
		assert.NoError(t, err)
		assert.Equal(t, "", target.Name)
		assert.Equal(t, 0, target.Age)
	})

	t.Run("scan_empty_json_array", func(t *testing.T) {
		var target []TestStruct
		err := Scan("[]", &target)
		assert.NoError(t, err)
		assert.Len(t, target, 0)
	})

	t.Run("scan_whitespace_only", func(t *testing.T) {
		var target TestStruct
		err := Scan("   ", &target)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "Syntax error")
	})

	t.Run("scan_json_null", func(t *testing.T) {
		var target *TestStruct
		err := Scan("null", &target)
		assert.NoError(t, err)
		assert.Nil(t, target)
	})

	t.Run("scan_malformed_array", func(t *testing.T) {
		var target []TestStruct
		err := Scan("[{\"name\":\"John\",\"age\":30", &target)
		assert.Error(t, err)
	})

	t.Run("scan_nested_json_string", func(t *testing.T) {
		var target map[string]interface{}
		err := Scan(`{"data":"{\"inner\":\"value\"}"}`, &target)
		assert.NoError(t, err)
		assert.Equal(t, "{\"inner\":\"value\"}", target["data"])
	})

	t.Run("scan_large_number", func(t *testing.T) {
		var target struct {
			BigInt int64 `json:"big_int"`
		}
		err := Scan(`{"big_int":9223372036854775807}`, &target)
		assert.NoError(t, err)
		assert.Equal(t, int64(9223372036854775807), target.BigInt)
	})

	t.Run("scan_unicode_characters", func(t *testing.T) {
		var target struct {
			Unicode string `json:"unicode"`
		}
		err := Scan(`{"unicode":"测试🚀"}`, &target)
		assert.NoError(t, err)
		assert.Equal(t, "测试🚀", target.Unicode)
	})
}

// TestValue_SpecialCases tests additional edge cases for Value function
func TestValue_SpecialCases(t *testing.T) {
	t.Run("value_empty_struct", func(t *testing.T) {
		input := struct{}{}
		value, err := Value(input)
		assert.NoError(t, err)
		bytes, ok := value.([]byte)
		require.True(t, ok)
		assert.Equal(t, "{}", string(bytes))
	})

	t.Run("value_struct_with_tags", func(t *testing.T) {
		input := struct {
			Name string `json:"custom_name"`
			Age  int    `json:"-"`
		}{
			Name: "Test",
			Age:  30,
		}
		value, err := Value(input)
		assert.NoError(t, err)
		bytes, ok := value.([]byte)
		require.True(t, ok)
		jsonStr := string(bytes)
		assert.Contains(t, jsonStr, `"custom_name":"Test"`)
		assert.NotContains(t, jsonStr, "Age")
	})

	t.Run("value_pointer_to_struct", func(t *testing.T) {
		input := &struct {
			Data string `json:"data"`
		}{
			Data: "pointer test",
		}
		value, err := Value(input)
		assert.NoError(t, err)
		bytes, ok := value.([]byte)
		require.True(t, ok)
		assert.Contains(t, string(bytes), `"data":"pointer test"`)
	})

	t.Run("value_nested_pointers", func(t *testing.T) {
		type Inner struct {
			Value string `json:"value"`
		}
		type Outer struct {
			Inner *Inner `json:"inner"`
		}
		input := &Outer{
			Inner: &Inner{Value: "nested"},
		}
		value, err := Value(input)
		assert.NoError(t, err)
		bytes, ok := value.([]byte)
		require.True(t, ok)
		assert.Contains(t, string(bytes), `"inner":{"value":"nested"}`)
	})

	t.Run("value_float_precision", func(t *testing.T) {
		input := struct {
			Float32 float32 `json:"float32"`
			Float64 float64 `json:"float64"`
		}{
			Float32: 3.14159,
			Float64: 3.141592653589793,
		}
		value, err := Value(input)
		assert.NoError(t, err)
		bytes, ok := value.([]byte)
		require.True(t, ok)
		jsonStr := string(bytes)
		assert.Contains(t, jsonStr, "float32")
		assert.Contains(t, jsonStr, "float64")
	})
}

// TestScanValue_ConcurrentAccess tests thread safety
func TestScanValue_ConcurrentAccess(t *testing.T) {
	const numGoroutines = 10
	const numOperations = 100

	var wg sync.WaitGroup
	errCh := make(chan error, numGoroutines*numOperations*2)

	for i := 0; i < numGoroutines; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			for j := 0; j < numOperations; j++ {
				// Test Scan
				src := fmt.Sprintf(`{"name":"Goroutine_%d_%d","age":%d}`, id, j, j%100)
				var target TestStruct
				if err := Scan(src, &target); err != nil {
					errCh <- fmt.Errorf("goroutine %d scan error: %v", id, err)
					return
				}

				// Test Value
				input := TestStruct{Name: fmt.Sprintf("Test_%d_%d", id, j), Age: j % 50}
				if _, err := Value(input); err != nil {
					errCh <- fmt.Errorf("goroutine %d value error: %v", id, err)
					return
				}
			}
		}(i)
	}

	wg.Wait()
	close(errCh)

	for err := range errCh {
		t.Errorf("Concurrent access error: %v", err)
	}
}

// Benchmark tests for performance
func BenchmarkScan(b *testing.B) {
	src := `{"name":"John","age":30}`
	for i := 0; i < b.N; i++ {
		var target TestStruct
		_ = Scan(src, &target)
	}
}

func BenchmarkValue(b *testing.B) {
	input := TestStruct{Name: "John", Age: 30}
	for i := 0; i < b.N; i++ {
		_, _ = Value(input)
	}
}

func BenchmarkScanValue_RoundTrip(b *testing.B) {
	original := ComplexStruct{
		ID:    123,
		Data:  TestStruct{Name: "Benchmark", Age: 25},
		Items: []string{"a", "b", "c"},
		Map:   map[string]int{"key1": 1, "key2": 2},
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		value, _ := Value(original)
		var target ComplexStruct
		_ = Scan(value, &target)
	}
}
