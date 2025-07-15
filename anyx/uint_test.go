package anyx

import (
	"testing"
)

func TestToUint(t *testing.T) {
	tests := []struct {
		name  string
		input interface{}
		want  uint
	}{
		{"bool true", true, 1},
		{"bool false", false, 0},
		{"int positive", 42, 42},
		{"int negative", -1, 18446744073709551615},
		{"uint", uint(100), 100},
		{"float positive", 3.14, 3},
		{"string valid", "123", 123},
		{"string invalid", "abc", 0},
		{"byte slice valid", []byte("456"), 456},
		{"byte slice invalid", []byte("xyz"), 0},
		{"slice", []int{1, 2}, 0},
		{"map", map[string]int{"a": 1}, 0},
		{"nil pointer", (*int)(nil), 0},
		{"max int", 1<<63 - 1, 9223372036854775807},
		{"min int", -1 << 63, 9223372036854775808},
		{"max uint", ^uint(0), ^uint(0)},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ToUint(tt.input); got != tt.want {
				t.Errorf("ToUint() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestToUint8(t *testing.T) {
	tests := []struct {
		name  string
		input interface{}
		want  uint8
	}{
		{"bool true", true, 1},
		{"int positive", 200, 200},
		{"int overflow", 300, 44},
		{"int negative", -1, 255},
		{"float positive", 100.5, 100},
		{"string valid", "128", 128},
		{"string invalid", "abc", 0},
		{"max uint8", uint8(255), 255},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ToUint8(tt.input); got != tt.want {
				t.Errorf("ToUint8() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestToUint16(t *testing.T) {
	tests := []struct {
		name  string
		input interface{}
		want  uint16
	}{
		{"int positive", 50000, 50000},
		{"int overflow", 70000, 4464},
		{"float negative", -100.5, 65436},
		{"string valid", "65535", 65535},
		{"max uint16", uint16(65535), 65535},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ToUint16(tt.input); got != tt.want {
				t.Errorf("ToUint16() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestToUint32(t *testing.T) {
	tests := []struct {
		name  string
		input interface{}
		want  uint32
	}{
		{"int positive", 3000000000, 3000000000},
		{"int overflow", 5000000000, 705032704},
		{"string valid", "4294967295", 4294967295},
		{"max uint32", uint32(4294967295), 4294967295},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ToUint32(tt.input); got != tt.want {
				t.Errorf("ToUint32() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestToUint64(t *testing.T) {
	tests := []struct {
		name  string
		input interface{}
		want  uint64
	}{
		{"int positive", 9223372036854775807, 9223372036854775807},
		{"int negative", -1, 18446744073709551615},
		{"string valid", "18446744073709551615", 18446744073709551615},
		{"max uint64", uint64(18446744073709551615), 18446744073709551615},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ToUint64(tt.input); got != tt.want {
				t.Errorf("ToUint64() = %v, want %v", got, tt.want)
			}
		})
	}
}

func ptr[T any](v T) *T {
	return &v
}
