package anyx

import (
	"math"
	"testing"
)

func TestToBool(t *testing.T) {
	cases := []struct {
		name     string
		input    interface{}
		expected bool
	}{
		// bool
		{"bool true", true, true},
		{"bool false", false, false},

		// int
		{"int 0", 0, false},
		{"int 1", 1, true},
		{"int -1", -1, true},
		{"int8 0", int8(0), false},
		{"int16 10", int16(10), true},
		{"int32 -10", int32(-10), true},
		{"int64 100", int64(100), true},

		// uint
		{"uint 0", uint(0), false},
		{"uint 1", uint(1), true},
		{"uint8 0", uint8(0), false},
		{"uint16 10", uint16(10), true},
		{"uint32 10", uint32(10), true},
		{"uint64 100", uint64(100), true},

		// float
		{"float32 0.0", float32(0.0), false},
		{"float32 1.23", float32(1.23), true},
		{"float32 -4.56", float32(-4.56), true},
		{"float32 NaN", float32(math.NaN()), false},
		{"float64 0.0", 0.0, false},
		{"float64 1.23", 1.23, true},
		{"float64 -4.56", -4.56, true},
		{"float64 NaN", math.NaN(), false},

		// string (true values)
		{"string true", "true", true},
		{"string TRUE", "TRUE", true},
		{"string 1", "1", true},
		{"string t", "t", true},
		{"string T", "T", true},
		{"string y", "y", true},
		{"string Y", "Y", true},
		{"string yes", "yes", true},
		{"string YES", "YES", true},
		{"string on", "on", true},
		{"string ON", "ON", true},

		// string (false values)
		{"string false", "false", false},
		{"string FALSE", "FALSE", false},
		{"string 0", "0", false},
		{"string f", "f", false},
		{"string F", "F", false},
		{"string n", "n", false},
		{"string N", "N", false},
		{"string no", "no", false},
		{"string NO", "NO", false},
		{"string off", "off", false},
		{"string OFF", "OFF", false},

		// string (other non-empty)
		{"string hello", "hello", true},
		{"string with spaces", "  hello  ", true},

		// string (empty)
		{"string empty", "", false},
		{"string space only", "   ", false},

		// []byte (true values)
		{"[]byte true", []byte("true"), true},
		{"[]byte TRUE", []byte("TRUE"), true},
		{"[]byte 1", []byte("1"), true},

		// []byte (false values)
		{"[]byte false", []byte("false"), false},
		{"[]byte FALSE", []byte("FALSE"), false},
		{"[]byte 0", []byte("0"), false},

		// []byte (other non-empty)
		{"[]byte hello", []byte("hello"), true},

		// []byte (empty)
		{"[]byte empty", []byte(""), false},
		{"[]byte space only", []byte("   "), false},

		// nil
		{"nil", nil, false},

		// unsupported types
		{"unsupported struct", struct{}{}, false},
		{"unsupported map", make(map[int]int), false},
		{"unsupported slice", []int{1, 2}, false},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			if got := ToBool(tc.input); got != tc.expected {
				t.Errorf("ToBool(%v) = %v; want %v", tc.input, got, tc.expected)
			}
		})
	}
}
