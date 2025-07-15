package anyx_test

import (
	"math"
	"testing"

	"github.com/lazygophers/utils/anyx"
	"github.com/stretchr/testify/assert"
)

func TestToFloat32(t *testing.T) {
	testCases := []struct {
		name     string
		input    interface{}
		expected float32
	}{
		{"bool true", true, 1.0},
		{"bool false", false, 0.0},
		{"int", 42, 42.0},
		{"int8", int8(42), 42.0},
		{"int16", int16(42), 42.0},
		{"int32", int32(42), 42.0},
		{"int64", int64(42), 42.0},
		{"uint", uint(42), 42.0},
		{"uint8", uint8(42), 42.0},
		{"uint16", uint16(42), 42.0},
		{"uint32", uint32(42), 42.0},
		{"uint64", uint64(42), 42.0},
		{"float32", float32(42.42), 42.42},
		{"float64", 42.42, float32(42.42)},
		{"valid string", "42.42", 42.42},
		{"valid []byte", []byte("42.42"), 42.42},
		{"invalid string", "invalid", 0.0},
		{"invalid []byte", []byte("invalid"), 0.0},
		{"max int64", int64(math.MaxInt64), float32(math.MaxInt64)},
		{"min float32", math.SmallestNonzeroFloat32, math.SmallestNonzeroFloat32},
		{"nil input", nil, 0.0},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := anyx.ToFloat32(tc.input)
			assert.InDelta(t, tc.expected, result, 0.0001)
		})
	}
}

func TestToFloat64(t *testing.T) {
	testCases := []struct {
		name     string
		input    interface{}
		expected float64
	}{
		{"bool true", true, 1.0},
		{"bool false", false, 0.0},
		{"int", 42, 42.0},
		{"int8", int8(42), 42.0},
		{"int16", int16(42), 42.0},
		{"int32", int32(42), 42.0},
		{"int64", int64(42), 42.0},
		{"uint", uint(42), 42.0},
		{"uint8", uint8(42), 42.0},
		{"uint16", uint16(42), 42.0},
		{"uint32", uint32(42), 42.0},
		{"uint64", uint64(42), 42.0},
		{"float32", float32(42.42), 42.42},
		{"float64", 42.42, 42.42},
		{"valid float string", "42.42", 42.42},
		{"valid int string", "42", 42.0},
		{"valid []byte", []byte("42.42"), 42.42},
		{"invalid string", "invalid", 0.0},
		{"invalid []byte", []byte("invalid"), 0.0},
		{"max int64", int64(math.MaxInt64), float64(math.MaxInt64)},
		{"NaN", math.NaN(), math.NaN()},
		{"Inf", math.Inf(1), math.Inf(1)},
		{"empty string", "", 0.0},
		{"nil input", nil, 0.0},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := anyx.ToFloat64(tc.input)
			if math.IsNaN(tc.expected) {
				assert.True(t, math.IsNaN(result))
			} else if math.IsInf(tc.expected, 0) {
				assert.True(t, math.IsInf(result, int(math.Copysign(1, tc.expected))))
			} else {
				assert.InDelta(t, tc.expected, result, 0.0001)
			}
		})
	}
}

func TestToFloat64Slice(t *testing.T) {
	testCases := []struct {
		name     string
		input    interface{}
		expected []float64
	}{
		{"[]bool", []bool{true, false}, []float64{1.0, 0.0}},
		{"[]int", []int{1, 2, 3}, []float64{1.0, 2.0, 3.0}},
		{"[]int8", []int8{1, 2, 3}, []float64{1.0, 2.0, 3.0}},
		{"[]int16", []int16{1, 2, 3}, []float64{1.0, 2.0, 3.0}},
		{"[]int32", []int32{1, 2, 3}, []float64{1.0, 2.0, 3.0}},
		{"[]int64", []int64{1, 2, 3}, []float64{1.0, 2.0, 3.0}},
		{"[]uint", []uint{1, 2, 3}, []float64{1.0, 2.0, 3.0}},
		{"[]uint8", []uint8{1, 2, 3}, []float64{1.0, 2.0, 3.0}},
		{"[]uint16", []uint16{1, 2, 3}, []float64{1.0, 2.0, 3.0}},
		{"[]uint32", []uint32{1, 2, 3}, []float64{1.0, 2.0, 3.0}},
		{"[]uint64", []uint64{1, 2, 3}, []float64{1.0, 2.0, 3.0}},
		{"[]float32", []float32{1.0, 2.0, 3.0}, []float64{1.0, 2.0, 3.0}},
		{"[]float64", []float64{1.0, 2.0, 3.0}, []float64{1.0, 2.0, 3.0}},
		{"[]string", []string{"1.0", "2.0", "3.0"}, []float64{1.0, 2.0, 3.0}},
		{"[]bool", []bool{true, false}, []float64{1.0, 0.0}},
		{"nil input", nil, nil},
		{"empty slice", []int{}, []float64{}},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := anyx.ToFloat64Slice(tc.input)
			if tc.expected == nil {
				assert.Nil(t, result)
			} else {
				assert.Equal(t, len(tc.expected), len(result))
				for i := range tc.expected {
					assert.InDelta(t, tc.expected[i], result[i], 0.0001)
				}
			}
		})
	}
}
