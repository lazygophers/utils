package anyx

import (
	"testing"
	"time"
)

func TestToInt(t *testing.T) {
	tests := []struct {
		name  string
		input interface{}
		want  int
	}{
		{"bool true", true, 1},
		{"bool false", false, 0},
		{"int", 42, 42},
		{"int8", int8(42), 42},
		{"int16", int16(42), 42},
		{"int32", int32(42), 42},
		{"int64", int64(42), 42},
		{"uint", uint(42), 42},
		{"uint8", uint8(42), 42},
		{"uint16", uint16(42), 42},
		{"uint32", uint32(42), 42},
		{"uint64", uint64(42), 42},
		{"float32", float32(42.5), 42},
		{"float64", 42.5, 42},
		{"valid string", "42", 42},
		{"invalid string", "abc", 0},
		{"empty string", "", 0},
		{"valid []byte", []byte("42"), 42},
		{"invalid []byte", []byte("abc"), 0},
		{"nil", nil, 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ToInt(tt.input); got != tt.want {
				t.Errorf("ToInt() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestToInt64(t *testing.T) {
	tests := []struct {
		name  string
		input interface{}
		want  int64
	}{
		{"bool true", true, 1},
		{"bool false", false, 0},
		{"int", 42, 42},
		{"int8", int8(42), 42},
		{"int16", int16(42), 42},
		{"int32", int32(42), 42},
		{"int64", int64(42), 42},
		{"uint", uint(42), 42},
		{"uint8", uint8(42), 42},
		{"uint16", uint16(42), 42},
		{"uint32", uint32(42), 42},
		{"uint64", uint64(42), 42},
		{"time.Duration", time.Second, int64(time.Second)},
		{"float32", float32(42.5), 42},
		{"float64", 42.5, 42},
		{"valid string", "42", 42},
		{"invalid string", "abc", 0},
		{"empty string", "", 0},
		{"max int64", "9223372036854775807", 9223372036854775807},
		{"min int64", "-9223372036854775808", -9223372036854775808},
		{"overflow positive", "9223372036854775808", 0},
		{"overflow negative", "-9223372036854775809", 0},
		{"valid []byte", []byte("42"), 42},
		{"invalid []byte", []byte("abc"), 0},
		{"nil", nil, 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ToInt64(tt.input); got != tt.want {
				t.Errorf("ToInt64() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestToInt8(t *testing.T) {
	tests := []struct {
		name  string
		input interface{}
		want  int8
	}{
		{"bool true", true, 1},
		{"bool false", false, 0},
		{"int", 42, 42},
		{"int8", int8(42), 42},
		{"int16", int16(42), 42},
		{"int32", int32(42), 42},
		{"int64", int64(42), 42},
		{"uint", uint(42), 42},
		{"uint8", uint8(42), 42},
		{"uint16", uint16(42), 42},
		{"uint32", uint32(42), 42},
		{"uint64", uint64(42), 42},
		{"float32", float32(42.5), 42},
		{"float64", 42.5, 42},
		{"valid string", "42", 42},
		{"invalid string", "abc", 0},
		{"empty string", "", 0},
		{"valid []byte", []byte("42"), 42},
		{"invalid []byte", []byte("abc"), 0},
		{"nil", nil, 0},
		{"max int8", "127", 127},
		{"min int8", "-128", 0},
		{"overflow max", "128", -128},
		{"overflow min", "-129", 0},
		{"overflow min", "-129", 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ToInt8(tt.input); got != tt.want {
				t.Errorf("ToInt8() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestToInt16(t *testing.T) {
	tests := []struct {
		name  string
		input interface{}
		want  int16
	}{
		{"bool true", true, 1},
		{"bool false", false, 0},
		{"int", 42, 42},
		{"int8", int8(42), 42},
		{"int16", int16(42), 42},
		{"int32", int32(42), 42},
		{"int64", int64(42), 42},
		{"uint", uint(42), 42},
		{"uint8", uint8(42), 42},
		{"uint16", uint16(42), 42},
		{"uint32", uint32(42), 42},
		{"uint64", uint64(42), 42},
		{"float32", float32(42.5), 42},
		{"float64", 42.5, 42},
		{"valid string", "42", 42},
		{"invalid string", "abc", 0},
		{"empty string", "", 0},
		{"valid []byte", []byte("42"), 42},
		{"invalid []byte", []byte("abc"), 0},
		{"nil", nil, 0},
		{"max int16", "32767", 32767},
		{"min int16", "-32768", 0},
		{"overflow max", "32768", -32768},
		{"overflow min", "-32769", 0},
		{"overflow min", "-32769", 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ToInt16(tt.input); got != tt.want {
				t.Errorf("ToInt16() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestToInt32(t *testing.T) {
	tests := []struct {
		name  string
		input interface{}
		want  int32
	}{
		{"bool true", true, 1},
		{"bool false", false, 0},
		{"int", 42, 42},
		{"int8", int8(42), 42},
		{"int16", int16(42), 42},
		{"int32", int32(42), 42},
		{"int64", int64(42), 42},
		{"uint", uint(42), 42},
		{"uint8", uint8(42), 42},
		{"uint16", uint16(42), 42},
		{"uint32", uint32(42), 42},
		{"uint64", uint64(42), 42},
		{"float32", float32(42.5), 42},
		{"float64", 42.5, 42},
		{"valid string", "42", 42},
		{"invalid string", "abc", 0},
		{"empty string", "", 0},
		{"valid []byte", []byte("42"), 42},
		{"invalid []byte", []byte("abc"), 0},
		{"nil", nil, 0},
		{"max int32", "2147483647", 2147483647},
		{"min int32", "-2147483648", 0},
		{"overflow max", "2147483648", -2147483648},
		{"overflow min", "-2147483649", 0},
		{"overflow min", "-2147483649", 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ToInt32(tt.input); got != tt.want {
				t.Errorf("ToInt32() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestToInt64Slice(t *testing.T) {
	tests := []struct {
		name  string
		input interface{}
		want  []int64
	}{
		{"[]bool", []bool{true, false}, []int64{1, 0}},
		{"[]int", []int{1, 2, 3}, []int64{1, 2, 3}},
		{"[]int8", []int8{1, 2, 3}, []int64{1, 2, 3}},
		{"[]int16", []int16{1, 2, 3}, []int64{1, 2, 3}},
		{"[]int32", []int32{1, 2, 3}, []int64{1, 2, 3}},
		{"[]int64", []int64{1, 2, 3}, []int64{1, 2, 3}},
		{"[]uint", []uint{1, 2, 3}, []int64{1, 2, 3}},
		{"[]uint8", []uint8{1, 2, 3}, []int64{1, 2, 3}},
		{"[]uint16", []uint16{1, 2, 3}, []int64{1, 2, 3}},
		{"[]uint32", []uint32{1, 2, 3}, []int64{1, 2, 3}},
		{"[]uint64", []uint64{1, 2, 3}, []int64{1, 2, 3}},
		{"[]float32", []float32{1.1, 2.2, 3.3}, []int64{1, 2, 3}},
		{"[]float64", []float64{1.1, 2.2, 3.3}, []int64{1, 2, 3}},
		{"[]string", []string{"1", "2", "3"}, []int64{1, 2, 3}},
		{"invalid []string", []string{"a", "b", "c"}, []int64{0, 0, 0}},
		{"[][]byte", [][]byte{[]byte("1"), []byte("2"), []byte("3")}, []int64{1, 2, 3}},
		{"invalid [][]byte", [][]byte{[]byte("a"), []byte("b"), []byte("c")}, []int64{0, 0, 0}},
		{"[]interface{}", []interface{}{1, "2", true}, []int64{1, 2, 1}},
		{"empty slice", []int{}, []int64{}},
		{"nil", nil, []int64{}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ToInt64Slice(tt.input)
			if len(got) != len(tt.want) {
				t.Errorf("ToInt64Slice() length = %d, want %d", len(got), len(tt.want))
				return
			}
			for i := range got {
				if got[i] != tt.want[i] {
					t.Errorf("ToInt64Slice()[%d] = %v, want %v", i, got[i], tt.want[i])
				}
			}
		})
	}
}
