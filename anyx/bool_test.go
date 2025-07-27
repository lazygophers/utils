package anyx_test

import (
	"testing"

	"github.com/lazygophers/utils/anyx"

	"github.com/stretchr/testify/assert"
)

func TestToBool(t *testing.T) {
	tests := []struct {
		name     string
		input    interface{}
		expected bool
	}{
		{"bool_true", true, true},
		{"bool_false", false, false},

		{"int_positive", 42, true},
		{"int_zero", 0, false},
		{"int_negative", -1, true},
		{"float_positive", 3.14, true},
		{"float_zero", 0.0, false},

		{"string_true", "true", true},
		{"string_TRUE", "TRUE", true},
		{"string_t", "t", true},
		{"string_yes", "yes", true},
		{"string_on", "on", true},
		{"string_false", "false", false},
		{"string_FALSE", "FALSE", false},
		{"string_f", "f", false},
		{"string_no", "no", false},
		{"string_off", "off", false},
		{"string_empty", "", false},
		{"string_whitespace", "   ", false},
		{"string_random", "hello", true},

		{"bytes_true", []byte("true"), true},
		{"bytes_TRUE", []byte("TRUE"), true},
		{"bytes_false", []byte("false"), false},
		{"bytes_empty", []byte(""), false},
		{"bytes_whitespace", []byte("   "), false},

		{"nil_value", nil, false},
		{"struct_value", struct{}{}, false},
		{"slice_value", []int{}, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.expected, anyx.ToBool(tt.input))
		})
	}
}
