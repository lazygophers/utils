package candy

import (
	"errors"
	"reflect"
	"testing"
	"time"
)

func TestToBytes(t *testing.T) {
	t.Run("boolean values", func(t *testing.T) {
		tests := []struct {
			name     string
			input    bool
			expected []byte
		}{
			{"true", true, []byte("1")},
			{"false", false, []byte("0")},
		}

		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				result := ToBytes(tt.input)
				if !reflect.DeepEqual(result, tt.expected) {
					t.Errorf("ToBytes(%v) = %v, want %v", tt.input, result, tt.expected)
				}
			})
		}
	})

	t.Run("integer types", func(t *testing.T) {
		tests := []struct {
			name     string
			input    interface{}
			expected []byte
		}{
			{"int positive", int(42), []byte("42")},
			{"int negative", int(-42), []byte("-42")},
			{"int zero", int(0), []byte("0")},
			{"int8 max", int8(127), []byte("127")},
			{"int8 min", int8(-128), []byte("-128")},
			{"int16 positive", int16(1000), []byte("1000")},
			{"int16 negative", int16(-1000), []byte("-1000")},
			{"int32 positive", int32(100000), []byte("100000")},
			{"int32 negative", int32(-100000), []byte("-100000")},
			{"int64 positive", int64(9223372036854775807), []byte("9223372036854775807")},
			{"int64 negative", int64(-9223372036854775808), []byte("-9223372036854775808")},
		}

		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				result := ToBytes(tt.input)
				if !reflect.DeepEqual(result, tt.expected) {
					t.Errorf("ToBytes(%v) = %s, want %s", tt.input, result, tt.expected)
				}
			})
		}
	})

	t.Run("unsigned integer types", func(t *testing.T) {
		tests := []struct {
			name     string
			input    interface{}
			expected []byte
		}{
			{"uint zero", uint(0), []byte("0")},
			{"uint positive", uint(42), []byte("42")},
			{"uint8 max", uint8(255), []byte("255")},
			{"uint8 zero", uint8(0), []byte("0")},
			{"uint16 max", uint16(65535), []byte("65535")},
			{"uint16 positive", uint16(1000), []byte("1000")},
			{"uint32 max", uint32(4294967295), []byte("4294967295")},
			{"uint32 positive", uint32(100000), []byte("100000")},
			{"uint64 max", uint64(18446744073709551615), []byte("18446744073709551615")},
			{"uint64 positive", uint64(1000000), []byte("1000000")},
		}

		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				result := ToBytes(tt.input)
				if !reflect.DeepEqual(result, tt.expected) {
					t.Errorf("ToBytes(%v) = %s, want %s", tt.input, result, tt.expected)
				}
			})
		}
	})

	t.Run("float32 values", func(t *testing.T) {
		tests := []struct {
			name     string
			input    float32
			expected []byte
		}{
			{"float32 zero", float32(0.0), []byte("0")},
			{"float32 integer", float32(42.0), []byte("42")},
			{"float32 decimal", float32(3.14), []byte("3.140000104904175")},
			{"float32 negative", float32(-3.14), []byte("-3.140000104904175")},
			{"float32 small", float32(0.001), []byte("0.001000000047497")},
		}

		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				result := ToBytes(tt.input)
				if !reflect.DeepEqual(result, tt.expected) {
					t.Errorf("ToBytes(%v) = %s, want %s", tt.input, result, tt.expected)
				}
			})
		}
	})

	t.Run("float64 values", func(t *testing.T) {
		tests := []struct {
			name     string
			input    float64
			expected []byte
		}{
			{"float64 zero", 0.0, []byte("0")},
			{"float64 integer", 42.0, []byte("42")},
			{"float64 decimal", 3.14159, []byte("3.141590")},
			{"float64 negative", -3.14159, []byte("-3.141590")},
			{"float64 small", 0.000001, []byte("0.000001")},
			{"float64 large", 123456789.123456, []byte("123456789.123456")},
		}

		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				result := ToBytes(tt.input)
				if !reflect.DeepEqual(result, tt.expected) {
					t.Errorf("ToBytes(%v) = %s, want %s", tt.input, result, tt.expected)
				}
			})
		}
	})

	t.Run("time.Duration values", func(t *testing.T) {
		tests := []struct {
			name     string
			input    time.Duration
			expected []byte
		}{
			{"zero duration", time.Duration(0), []byte("0s")},
			{"nanoseconds", 100 * time.Nanosecond, []byte("100ns")},
			{"microseconds", 100 * time.Microsecond, []byte("100µs")},
			{"milliseconds", 100 * time.Millisecond, []byte("100ms")},
			{"seconds", 5 * time.Second, []byte("5s")},
			{"minutes", 2 * time.Minute, []byte("2m0s")},
			{"hours", 3 * time.Hour, []byte("3h0m0s")},
			{"mixed", 1*time.Hour + 30*time.Minute + 45*time.Second, []byte("1h30m45s")},
			{"negative", -5 * time.Second, []byte("-5s")},
		}

		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				result := ToBytes(tt.input)
				if !reflect.DeepEqual(result, tt.expected) {
					t.Errorf("ToBytes(%v) = %s, want %s", tt.input, result, tt.expected)
				}
			})
		}
	})

	t.Run("string values", func(t *testing.T) {
		tests := []struct {
			name     string
			input    string
			expected []byte
		}{
			{"empty string", "", []byte("")},
			{"simple string", "hello", []byte("hello")},
			{"string with spaces", "hello world", []byte("hello world")},
			{"unicode string", "你好世界", []byte("你好世界")},
			{"string with special chars", "test\n\t\r", []byte("test\n\t\r")},
		}

		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				result := ToBytes(tt.input)
				if !reflect.DeepEqual(result, tt.expected) {
					t.Errorf("ToBytes(%q) = %v, want %v", tt.input, result, tt.expected)
				}
			})
		}
	})

	t.Run("byte slice values", func(t *testing.T) {
		tests := []struct {
			name     string
			input    []byte
			expected []byte
		}{
			{"nil bytes", nil, nil},
			{"empty bytes", []byte{}, []byte{}},
			{"simple bytes", []byte("test"), []byte("test")},
			{"bytes with nulls", []byte{0, 1, 2, 3}, []byte{0, 1, 2, 3}},
		}

		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				result := ToBytes(tt.input)
				if !reflect.DeepEqual(result, tt.expected) {
					t.Errorf("ToBytes(%v) = %v, want %v", tt.input, result, tt.expected)
				}
			})
		}
	})

	t.Run("nil value", func(t *testing.T) {
		result := ToBytes(nil)
		if result != nil {
			t.Errorf("ToBytes(nil) = %v, want nil", result)
		}
	})

	t.Run("error values", func(t *testing.T) {
		tests := []struct {
			name     string
			input    error
			expected []byte
		}{
			{"simple error", errors.New("test error"), []byte("test error")},
			{"custom error", errors.New("custom: something went wrong"), []byte("custom: something went wrong")},
		}

		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				result := ToBytes(tt.input)
				if !reflect.DeepEqual(result, tt.expected) {
					t.Errorf("ToBytes(%v) = %s, want %s", tt.input, result, tt.expected)
				}
			})
		}
	})

	t.Run("struct values with JSON serialization", func(t *testing.T) {
		type TestStruct struct {
			Name  string `json:"name"`
			Value int    `json:"value"`
		}

		tests := []struct {
			name     string
			input    interface{}
			expected []byte
		}{
			{"simple struct", TestStruct{Name: "test", Value: 42}, []byte(`{"name":"test","value":42}`)},
			{"empty struct", TestStruct{}, []byte(`{"name":"","value":0}`)},
			{"struct pointer", &TestStruct{Name: "ptr", Value: 100}, []byte(`{"name":"ptr","value":100}`)},
		}

		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				result := ToBytes(tt.input)
				if !reflect.DeepEqual(result, tt.expected) {
					t.Errorf("ToBytes(%v) = %s, want %s", tt.input, result, tt.expected)
				}
			})
		}
	})

	t.Run("map values with JSON serialization", func(t *testing.T) {
		t.Run("empty map", func(t *testing.T) {
			input := map[string]int{}
			result := ToBytes(input)
			expected := []byte(`{}`)
			if !reflect.DeepEqual(result, expected) {
				t.Errorf("ToBytes(%v) = %s, want %s", input, result, expected)
			}
		})

		t.Run("simple map", func(t *testing.T) {
			input := map[string]int{"a": 1}
			result := ToBytes(input)
			expected := []byte(`{"a":1}`)
			if !reflect.DeepEqual(result, expected) {
				t.Errorf("ToBytes(%v) = %s, want %s", input, result, expected)
			}
		})

		t.Run("nested map", func(t *testing.T) {
			input := map[string]interface{}{"outer": map[string]int{"inner": 42}}
			result := ToBytes(input)
			expected := []byte(`{"outer":{"inner":42}}`)
			if !reflect.DeepEqual(result, expected) {
				t.Errorf("ToBytes(%v) = %s, want %s", input, result, expected)
			}
		})
	})

	t.Run("slice values with JSON serialization", func(t *testing.T) {
		tests := []struct {
			name     string
			input    interface{}
			expected []byte
		}{
			{"int slice", []int{1, 2, 3}, []byte(`[1,2,3]`)},
			{"string slice", []string{"a", "b", "c"}, []byte(`["a","b","c"]`)},
			{"empty slice", []int{}, []byte(`[]`)},
		}

		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				result := ToBytes(tt.input)
				if !reflect.DeepEqual(result, tt.expected) {
					t.Errorf("ToBytes(%v) = %s, want %s", tt.input, result, tt.expected)
				}
			})
		}
	})

	t.Run("unsupported types that fail JSON marshaling", func(t *testing.T) {
		// Channel type cannot be marshaled to JSON
		ch := make(chan int)
		result := ToBytes(ch)
		if result != nil {
			t.Errorf("ToBytes(channel) = %v, want nil", result)
		}

		// Function type cannot be marshaled to JSON
		fn := func() {}
		result = ToBytes(fn)
		if result != nil {
			t.Errorf("ToBytes(function) = %v, want nil", result)
		}
	})
}

func TestToString_HelperFunc(t *testing.T) {
	t.Run("convert bytes to string", func(t *testing.T) {
		tests := []struct {
			name  string
			input []byte
		}{
			{"empty bytes", []byte{}},
			{"simple bytes", []byte("hello")},
			{"unicode bytes", []byte("你好")},
			{"bytes with nulls", []byte{0, 1, 2}},
			{"bytes with special chars", []byte("test\n\t\r")},
		}

		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				result := toString(tt.input)
				expected := string(tt.input)
				if result != expected {
					t.Errorf("toString(%v) = %q, want %q", tt.input, result, expected)
				}
				// Verify the conversion is correct by converting back
				if !reflect.DeepEqual([]byte(result), tt.input) {
					t.Errorf("toString round-trip failed for %v", tt.input)
				}
			})
		}
	})
}

func TestToBytes_HelperFunc(t *testing.T) {
	t.Run("convert string to bytes", func(t *testing.T) {
		tests := []struct {
			name  string
			input string
		}{
			{"empty string", ""},
			{"simple string", "hello"},
			{"unicode string", "你好"},
		}

		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				result := toBytes(tt.input)
				expected := []byte(tt.input)
				if string(result) != tt.input {
					t.Errorf("toBytes(%q) = %v, want %v", tt.input, result, expected)
				}
			})
		}
	})
}

func TestToBytes_EdgeCases(t *testing.T) {
	t.Run("large numbers", func(t *testing.T) {
		largeInt := int64(9223372036854775807)
		result := ToBytes(largeInt)
		expected := []byte("9223372036854775807")
		if !reflect.DeepEqual(result, expected) {
			t.Errorf("ToBytes(large int) = %s, want %s", result, expected)
		}

		largeUint := uint64(18446744073709551615)
		result = ToBytes(largeUint)
		expected = []byte("18446744073709551615")
		if !reflect.DeepEqual(result, expected) {
			t.Errorf("ToBytes(large uint) = %s, want %s", result, expected)
		}
	})

	t.Run("very small floats", func(t *testing.T) {
		smallFloat := 0.000000000001
		result := ToBytes(smallFloat)
		// Should have precision handling
		if len(result) == 0 {
			t.Error("ToBytes(small float) returned empty bytes")
		}
	})

	t.Run("complex nested structures", func(t *testing.T) {
		type Nested struct {
			Deep map[string][]int `json:"deep"`
		}
		input := Nested{Deep: map[string][]int{"array": {1, 2, 3}}}
		result := ToBytes(input)
		expected := []byte(`{"deep":{"array":[1,2,3]}}`)
		if !reflect.DeepEqual(result, expected) {
			t.Errorf("ToBytes(nested) = %s, want %s", result, expected)
		}
	})
}
