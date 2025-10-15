package candy

import (
	"math"
	"strings"
	"testing"
)

func TestToBool(t *testing.T) {
	t.Run("bool values", func(t *testing.T) {
		tests := []struct {
			name     string
			input    bool
			expected bool
		}{
			{"true", true, true},
			{"false", false, false},
		}

		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				result := ToBool(tt.input)
				if result != tt.expected {
					t.Errorf("ToBool(%v) = %v, want %v", tt.input, result, tt.expected)
				}
			})
		}
	})

	t.Run("int values", func(t *testing.T) {
		tests := []struct {
			name     string
			input    int
			expected bool
		}{
			{"zero", 0, false},
			{"positive", 1, true},
			{"negative", -1, true},
			{"large positive", 1000, true},
			{"large negative", -1000, true},
		}

		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				result := ToBool(tt.input)
				if result != tt.expected {
					t.Errorf("ToBool(%v) = %v, want %v", tt.input, result, tt.expected)
				}
			})
		}
	})

	t.Run("int8 values", func(t *testing.T) {
		tests := []struct {
			name     string
			input    int8
			expected bool
		}{
			{"zero", int8(0), false},
			{"positive", int8(1), true},
			{"negative", int8(-1), true},
			{"max", int8(127), true},
			{"min", int8(-128), true},
		}

		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				result := ToBool(tt.input)
				if result != tt.expected {
					t.Errorf("ToBool(%v) = %v, want %v", tt.input, result, tt.expected)
				}
			})
		}
	})

	t.Run("int16 values", func(t *testing.T) {
		tests := []struct {
			name     string
			input    int16
			expected bool
		}{
			{"zero", int16(0), false},
			{"positive", int16(100), true},
			{"negative", int16(-100), true},
		}

		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				result := ToBool(tt.input)
				if result != tt.expected {
					t.Errorf("ToBool(%v) = %v, want %v", tt.input, result, tt.expected)
				}
			})
		}
	})

	t.Run("int32 values", func(t *testing.T) {
		tests := []struct {
			name     string
			input    int32
			expected bool
		}{
			{"zero", int32(0), false},
			{"positive", int32(1000), true},
			{"negative", int32(-1000), true},
		}

		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				result := ToBool(tt.input)
				if result != tt.expected {
					t.Errorf("ToBool(%v) = %v, want %v", tt.input, result, tt.expected)
				}
			})
		}
	})

	t.Run("int64 values", func(t *testing.T) {
		tests := []struct {
			name     string
			input    int64
			expected bool
		}{
			{"zero", int64(0), false},
			{"positive", int64(100000), true},
			{"negative", int64(-100000), true},
		}

		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				result := ToBool(tt.input)
				if result != tt.expected {
					t.Errorf("ToBool(%v) = %v, want %v", tt.input, result, tt.expected)
				}
			})
		}
	})

	t.Run("uint values", func(t *testing.T) {
		tests := []struct {
			name     string
			input    uint
			expected bool
		}{
			{"zero", uint(0), false},
			{"positive", uint(1), true},
			{"large", uint(1000), true},
		}

		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				result := ToBool(tt.input)
				if result != tt.expected {
					t.Errorf("ToBool(%v) = %v, want %v", tt.input, result, tt.expected)
				}
			})
		}
	})

	t.Run("uint8 values", func(t *testing.T) {
		tests := []struct {
			name     string
			input    uint8
			expected bool
		}{
			{"zero", uint8(0), false},
			{"positive", uint8(1), true},
			{"max", uint8(255), true},
		}

		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				result := ToBool(tt.input)
				if result != tt.expected {
					t.Errorf("ToBool(%v) = %v, want %v", tt.input, result, tt.expected)
				}
			})
		}
	})

	t.Run("uint16 values", func(t *testing.T) {
		tests := []struct {
			name     string
			input    uint16
			expected bool
		}{
			{"zero", uint16(0), false},
			{"positive", uint16(100), true},
		}

		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				result := ToBool(tt.input)
				if result != tt.expected {
					t.Errorf("ToBool(%v) = %v, want %v", tt.input, result, tt.expected)
				}
			})
		}
	})

	t.Run("uint32 values", func(t *testing.T) {
		tests := []struct {
			name     string
			input    uint32
			expected bool
		}{
			{"zero", uint32(0), false},
			{"positive", uint32(1000), true},
		}

		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				result := ToBool(tt.input)
				if result != tt.expected {
					t.Errorf("ToBool(%v) = %v, want %v", tt.input, result, tt.expected)
				}
			})
		}
	})

	t.Run("uint64 values", func(t *testing.T) {
		tests := []struct {
			name     string
			input    uint64
			expected bool
		}{
			{"zero", uint64(0), false},
			{"positive", uint64(100000), true},
		}

		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				result := ToBool(tt.input)
				if result != tt.expected {
					t.Errorf("ToBool(%v) = %v, want %v", tt.input, result, tt.expected)
				}
			})
		}
	})

	t.Run("float32 values", func(t *testing.T) {
		tests := []struct {
			name     string
			input    float32
			expected bool
		}{
			{"zero", float32(0.0), false},
			{"positive", float32(1.5), true},
			{"negative", float32(-1.5), true},
			{"NaN", float32(math.NaN()), false},
		}

		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				result := ToBool(tt.input)
				if result != tt.expected {
					t.Errorf("ToBool(%v) = %v, want %v", tt.input, result, tt.expected)
				}
			})
		}
	})

	t.Run("float64 values", func(t *testing.T) {
		tests := []struct {
			name     string
			input    float64
			expected bool
		}{
			{"zero", 0.0, false},
			{"positive", 1.5, true},
			{"negative", -1.5, true},
			{"NaN", math.NaN(), false},
			{"small positive", 0.0001, true},
			{"small negative", -0.0001, true},
		}

		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				result := ToBool(tt.input)
				if result != tt.expected {
					t.Errorf("ToBool(%v) = %v, want %v", tt.input, result, tt.expected)
				}
			})
		}
	})

	t.Run("string true values", func(t *testing.T) {
		trueStrings := []string{
			"true", "TRUE", "True",
			"1",
			"t", "T",
			"y", "Y",
			"yes", "YES", "Yes",
			"on", "ON", "On",
			"  true  ", // with spaces
			"anything", // non-empty non-false string
			"random",
		}

		for _, str := range trueStrings {
			t.Run(str, func(t *testing.T) {
				result := ToBool(str)
				if !result {
					t.Errorf("ToBool(%q) = false, want true", str)
				}
			})
		}
	})

	t.Run("string false values", func(t *testing.T) {
		falseStrings := []string{
			"false", "FALSE", "False",
			"0",
			"f", "F",
			"n", "N",
			"no", "NO", "No",
			"off", "OFF", "Off",
			"", // empty string
			"  ", // whitespace only
			"  false  ", // with spaces
		}

		for _, str := range falseStrings {
			t.Run(str, func(t *testing.T) {
				result := ToBool(str)
				if result {
					t.Errorf("ToBool(%q) = true, want false", str)
				}
			})
		}
	})

	t.Run("byte slice true values", func(t *testing.T) {
		trueBytes := [][]byte{
			[]byte("true"),
			[]byte("1"),
			[]byte("t"),
			[]byte("y"),
			[]byte("yes"),
			[]byte("on"),
			[]byte("  TRUE  "),
			[]byte("something"),
		}

		for _, b := range trueBytes {
			t.Run(string(b), func(t *testing.T) {
				result := ToBool(b)
				if !result {
					t.Errorf("ToBool(%q) = false, want true", b)
				}
			})
		}
	})

	t.Run("byte slice false values", func(t *testing.T) {
		falseBytes := [][]byte{
			[]byte("false"),
			[]byte("0"),
			[]byte("f"),
			[]byte("n"),
			[]byte("no"),
			[]byte("off"),
			[]byte(""),
			[]byte("  "),
			[]byte("  FALSE  "),
		}

		for _, b := range falseBytes {
			t.Run(string(b), func(t *testing.T) {
				result := ToBool(b)
				if result {
					t.Errorf("ToBool(%q) = true, want false", b)
				}
			})
		}
	})

	t.Run("nil value", func(t *testing.T) {
		result := ToBool(nil)
		if result {
			t.Errorf("ToBool(nil) = true, want false")
		}
	})

	t.Run("other types default to false", func(t *testing.T) {
		tests := []struct {
			name  string
			input interface{}
		}{
			{"struct", struct{ Name string }{Name: "test"}},
			{"slice", []int{1, 2, 3}},
			{"map", map[string]int{"a": 1}},
			{"function", func() {}},
			{"channel", make(chan int)},
			{"pointer", new(int)},
		}

		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				result := ToBool(tt.input)
				if result {
					t.Errorf("ToBool(%v) = true, want false", tt.input)
				}
			})
		}
	})

	t.Run("edge cases", func(t *testing.T) {
		t.Run("empty byte slice", func(t *testing.T) {
			result := ToBool([]byte{})
			if result {
				t.Error("ToBool(empty byte slice) = true, want false")
			}
		})

		t.Run("byte slice with only spaces", func(t *testing.T) {
			result := ToBool([]byte("   \t\n   "))
			if result {
				t.Error("ToBool(whitespace byte slice) = true, want false")
			}
		})

		t.Run("string with mixed case", func(t *testing.T) {
			result := ToBool("TrUe")
			if !result {
				t.Error("ToBool(\"TrUe\") = false, want true")
			}
		})

		t.Run("string with leading/trailing spaces", func(t *testing.T) {
			result := ToBool("   YES   ")
			if !result {
				t.Error("ToBool(\"   YES   \") = false, want true")
			}
		})
	})

	t.Run("all true variations", func(t *testing.T) {
		variations := []string{"true", "1", "t", "y", "yes", "on"}
		for _, v := range variations {
			// Test lowercase
			if !ToBool(v) {
				t.Errorf("ToBool(%q) = false, want true", v)
			}
			// Test uppercase
			if !ToBool(strings.ToUpper(v)) {
				t.Errorf("ToBool(%q) = false, want true", strings.ToUpper(v))
			}
		}
	})

	t.Run("all false variations", func(t *testing.T) {
		variations := []string{"false", "0", "f", "n", "no", "off"}
		for _, v := range variations {
			// Test lowercase
			if ToBool(v) {
				t.Errorf("ToBool(%q) = true, want false", v)
			}
			// Test uppercase
			if ToBool(strings.ToUpper(v)) {
				t.Errorf("ToBool(%q) = true, want false", strings.ToUpper(v))
			}
		}
	})
}

func BenchmarkToBool(b *testing.B) {
	b.Run("bool", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			ToBool(true)
		}
	})

	b.Run("int", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			ToBool(1)
		}
	})

	b.Run("string true", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			ToBool("true")
		}
	})

	b.Run("string random", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			ToBool("random")
		}
	})

	b.Run("float64", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			ToBool(1.5)
		}
	})
}
