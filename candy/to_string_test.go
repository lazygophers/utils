package candy

import (
	"errors"
	"testing"
	"time"
)

func TestToString(t *testing.T) {
	t.Run("boolean values", func(t *testing.T) {
		tests := []struct {
			name     string
			input    bool
			expected string
		}{
			{"true", true, "1"},
			{"false", false, "0"},
		}

		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				result := ToString(tt.input)
				if result != tt.expected {
					t.Errorf("ToString(%v) = %q, want %q", tt.input, result, tt.expected)
				}
			})
		}
	})

	t.Run("integer types", func(t *testing.T) {
		tests := []struct {
			name     string
			input    interface{}
			expected string
		}{
			{"int positive", int(42), "42"},
			{"int negative", int(-42), "-42"},
			{"int zero", int(0), "0"},
			{"int8 max", int8(127), "127"},
			{"int8 min", int8(-128), "-128"},
			{"int16 positive", int16(1000), "1000"},
			{"int16 negative", int16(-1000), "-1000"},
			{"int32 positive", int32(100000), "100000"},
			{"int32 negative", int32(-100000), "-100000"},
			{"int64 positive", int64(9223372036854775807), "9223372036854775807"},
			{"int64 negative", int64(-9223372036854775808), "-9223372036854775808"},
		}

		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				result := ToString(tt.input)
				if result != tt.expected {
					t.Errorf("ToString(%v) = %q, want %q", tt.input, result, tt.expected)
				}
			})
		}
	})

	t.Run("unsigned integer types", func(t *testing.T) {
		tests := []struct {
			name     string
			input    interface{}
			expected string
		}{
			{"uint zero", uint(0), "0"},
			{"uint positive", uint(42), "42"},
			{"uint8 max", uint8(255), "255"},
			{"uint8 zero", uint8(0), "0"},
			{"uint16 max", uint16(65535), "65535"},
			{"uint16 positive", uint16(1000), "1000"},
			{"uint32 max", uint32(4294967295), "4294967295"},
			{"uint32 positive", uint32(100000), "100000"},
			{"uint64 max", uint64(18446744073709551615), "18446744073709551615"},
			{"uint64 positive", uint64(1000000), "1000000"},
		}

		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				result := ToString(tt.input)
				if result != tt.expected {
					t.Errorf("ToString(%v) = %q, want %q", tt.input, result, tt.expected)
				}
			})
		}
	})

	t.Run("float32 values", func(t *testing.T) {
		tests := []struct {
			name     string
			input    float32
			expected string
		}{
			{"float32 zero", float32(0.0), "0"},
			{"float32 integer", float32(42.0), "42"},
			{"float32 decimal", float32(3.14), "3.140000104904175"},
			{"float32 negative", float32(-3.14), "-3.140000104904175"},
			{"float32 small", float32(0.001), "0.001000000047497"},
		}

		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				result := ToString(tt.input)
				if result != tt.expected {
					t.Errorf("ToString(%v) = %q, want %q", tt.input, result, tt.expected)
				}
			})
		}
	})

	t.Run("float64 values", func(t *testing.T) {
		tests := []struct {
			name     string
			input    float64
			expected string
		}{
			{"float64 zero", 0.0, "0"},
			{"float64 integer", 42.0, "42"},
			{"float64 decimal", 3.14159, "3.141590"},
			{"float64 negative", -3.14159, "-3.141590"},
			{"float64 small", 0.000001, "0.000001"},
			{"float64 large", 123456789.123456, "123456789.123456"},
		}

		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				result := ToString(tt.input)
				if result != tt.expected {
					t.Errorf("ToString(%v) = %q, want %q", tt.input, result, tt.expected)
				}
			})
		}
	})

	t.Run("time.Duration values", func(t *testing.T) {
		tests := []struct {
			name     string
			input    time.Duration
			expected string
		}{
			{"zero duration", time.Duration(0), "0s"},
			{"nanoseconds", 100 * time.Nanosecond, "100ns"},
			{"microseconds", 100 * time.Microsecond, "100µs"},
			{"milliseconds", 100 * time.Millisecond, "100ms"},
			{"seconds", 5 * time.Second, "5s"},
			{"minutes", 2 * time.Minute, "2m0s"},
			{"hours", 3 * time.Hour, "3h0m0s"},
			{"mixed", 1*time.Hour + 30*time.Minute + 45*time.Second, "1h30m45s"},
			{"negative", -5 * time.Second, "-5s"},
		}

		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				result := ToString(tt.input)
				if result != tt.expected {
					t.Errorf("ToString(%v) = %q, want %q", tt.input, result, tt.expected)
				}
			})
		}
	})

	t.Run("string values", func(t *testing.T) {
		tests := []struct {
			name     string
			input    string
			expected string
		}{
			{"empty string", "", ""},
			{"simple string", "hello", "hello"},
			{"string with spaces", "hello world", "hello world"},
			{"unicode string", "你好世界", "你好世界"},
			{"string with special chars", "test\n\t\r", "test\n\t\r"},
		}

		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				result := ToString(tt.input)
				if result != tt.expected {
					t.Errorf("ToString(%q) = %q, want %q", tt.input, result, tt.expected)
				}
			})
		}
	})

	t.Run("byte slice values", func(t *testing.T) {
		tests := []struct {
			name     string
			input    []byte
			expected string
		}{
			{"nil bytes", nil, ""},
			{"empty bytes", []byte{}, ""},
			{"simple bytes", []byte("test"), "test"},
			{"bytes with nulls", []byte{0, 1, 2, 3}, string([]byte{0, 1, 2, 3})},
			{"unicode bytes", []byte("你好"), "你好"},
		}

		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				result := ToString(tt.input)
				if result != tt.expected {
					t.Errorf("ToString(%v) = %q, want %q", tt.input, result, tt.expected)
				}
			})
		}
	})

	t.Run("nil value", func(t *testing.T) {
		result := ToString(nil)
		if result != "" {
			t.Errorf("ToString(nil) = %q, want empty string", result)
		}
	})

	t.Run("error values", func(t *testing.T) {
		tests := []struct {
			name     string
			input    error
			expected string
		}{
			{"simple error", errors.New("test error"), "test error"},
			{"custom error", errors.New("custom: something went wrong"), "custom: something went wrong"},
		}

		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				result := ToString(tt.input)
				if result != tt.expected {
					t.Errorf("ToString(%v) = %q, want %q", tt.input, result, tt.expected)
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
			expected string
		}{
			{"simple struct", TestStruct{Name: "test", Value: 42}, `{"name":"test","value":42}`},
			{"empty struct", TestStruct{}, `{"name":"","value":0}`},
			{"struct pointer", &TestStruct{Name: "ptr", Value: 100}, `{"name":"ptr","value":100}`},
		}

		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				result := ToString(tt.input)
				if result != tt.expected {
					t.Errorf("ToString(%v) = %q, want %q", tt.input, result, tt.expected)
				}
			})
		}
	})

	t.Run("map values with JSON serialization", func(t *testing.T) {
		t.Run("empty map", func(t *testing.T) {
			input := map[string]int{}
			expected := "{}"
			result := ToString(input)

			if result != expected {
				t.Errorf("ToString(%v) = %q, want %q", input, result, expected)
			}
		})

		t.Run("simple map", func(t *testing.T) {
			input := map[string]int{"a": 1}
			expected := `{"a":1}`
			result := ToString(input)

			if result != expected {
				t.Errorf("ToString(%v) = %q, want %q", input, result, expected)
			}
		})

		t.Run("nested map", func(t *testing.T) {
			input := map[string]interface{}{"outer": map[string]int{"inner": 42}}
			expected := `{"outer":{"inner":42}}`
			result := ToString(input)

			if result != expected {
				t.Errorf("ToString(%v) = %q, want %q", input, result, expected)
			}
		})
	})

	t.Run("slice values with JSON serialization", func(t *testing.T) {
		tests := []struct {
			name     string
			input    interface{}
			expected string
		}{
			{"int slice", []int{1, 2, 3}, `[1,2,3]`},
			{"string slice", []string{"a", "b", "c"}, `["a","b","c"]`},
			{"empty slice", []int{}, `[]`},
		}

		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				result := ToString(tt.input)
				if result != tt.expected {
					t.Errorf("ToString(%v) = %q, want %q", tt.input, result, tt.expected)
				}
			})
		}
	})

	t.Run("unsupported types that fail JSON marshaling", func(t *testing.T) {
		// Channel type cannot be marshaled to JSON
		ch := make(chan int)
		result := ToString(ch)
		if result != "" {
			t.Errorf("ToString(channel) = %q, want empty string", result)
		}

		// Function type cannot be marshaled to JSON
		fn := func() {}
		result = ToString(fn)
		if result != "" {
			t.Errorf("ToString(function) = %q, want empty string", result)
		}
	})

	t.Run("edge cases", func(t *testing.T) {
		t.Run("large numbers", func(t *testing.T) {
			largeInt := int64(9223372036854775807)
			result := ToString(largeInt)
			expected := "9223372036854775807"
			if result != expected {
				t.Errorf("ToString(large int) = %q, want %q", result, expected)
			}

			largeUint := uint64(18446744073709551615)
			result = ToString(largeUint)
			expected = "18446744073709551615"
			if result != expected {
				t.Errorf("ToString(large uint) = %q, want %q", result, expected)
			}
		})

		t.Run("very small floats", func(t *testing.T) {
			smallFloat := 0.000000000001
			result := ToString(smallFloat)
			// Should have precision handling
			if len(result) == 0 {
				t.Error("ToString(small float) returned empty string")
			}
		})

		t.Run("complex nested structures", func(t *testing.T) {
			type Nested struct {
				Deep map[string][]int `json:"deep"`
			}
			input := Nested{Deep: map[string][]int{"array": {1, 2, 3}}}
			result := ToString(input)
			expected := `{"deep":{"array":[1,2,3]}}`
			if result != expected {
				t.Errorf("ToString(nested) = %q, want %q", result, expected)
			}
		})
	})

	t.Run("special numeric cases", func(t *testing.T) {
		t.Run("negative zero", func(t *testing.T) {
			result := ToString(int(-0))
			expected := "0"
			if result != expected {
				t.Errorf("ToString(-0) = %q, want %q", result, expected)
			}
		})

		t.Run("int boundary values", func(t *testing.T) {
			tests := []struct {
				name     string
				input    interface{}
				expected string
			}{
				{"int8 min", int8(-128), "-128"},
				{"int8 max", int8(127), "127"},
				{"int16 min", int16(-32768), "-32768"},
				{"int16 max", int16(32767), "32767"},
				{"int32 min", int32(-2147483648), "-2147483648"},
				{"int32 max", int32(2147483647), "2147483647"},
				{"uint8 max", uint8(255), "255"},
				{"uint16 max", uint16(65535), "65535"},
				{"uint32 max", uint32(4294967295), "4294967295"},
			}

			for _, tt := range tests {
				t.Run(tt.name, func(t *testing.T) {
					result := ToString(tt.input)
					if result != tt.expected {
						t.Errorf("ToString(%v) = %q, want %q", tt.input, result, tt.expected)
					}
				})
			}
		})

		t.Run("float edge cases", func(t *testing.T) {
			// Float that looks like integer
			result := ToString(float64(100.0))
			expected := "100"
			if result != expected {
				t.Errorf("ToString(100.0) = %q, want %q", result, expected)
			}

			// Very precise float
			result = ToString(float64(0.123456789))
			if len(result) == 0 {
				t.Error("ToString(precise float) returned empty string")
			}
		})
	})

	t.Run("pointer values", func(t *testing.T) {
		t.Run("pointer to struct", func(t *testing.T) {
			type User struct {
				Name string `json:"name"`
			}
			user := &User{Name: "Alice"}
			result := ToString(user)
			expected := `{"name":"Alice"}`
			if result != expected {
				t.Errorf("ToString(pointer) = %q, want %q", result, expected)
			}
		})

		t.Run("nil pointer", func(t *testing.T) {
			var ptr *int
			result := ToString(ptr)
			expected := "null"
			if result != expected {
				t.Errorf("ToString(nil pointer) = %q, want %q", result, expected)
			}
		})
	})

	t.Run("interface slice", func(t *testing.T) {
		input := []interface{}{1, "two", 3.0, true}
		result := ToString(input)
		expected := `[1,"two",3,true]`
		if result != expected {
			t.Errorf("ToString() = %q, want %q", result, expected)
		}
	})
}

func BenchmarkToString(b *testing.B) {
	b.Run("int", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			ToString(42)
		}
	})

	b.Run("string", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			ToString("hello world")
		}
	})

	b.Run("float64", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			ToString(3.14159)
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
			ToString(s)
		}
	})
}
