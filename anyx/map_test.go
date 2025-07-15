package anyx_test

import (
	"testing"

	"github.com/lazygophers/utils/anyx"
	"github.com/stretchr/testify/assert"
)

func TestMapKeysString(t *testing.T) {
	tests := []struct {
		name     string
		input    interface{}
		expected []string
		panics   bool
	}{
		{
			name:     "normal map",
			input:    map[string]int{"a": 1, "b": 2},
			expected: []string{"a", "b"},
		},
		{
			name:     "empty map",
			input:    map[string]int{},
			expected: []string{},
		},
		{
			name:   "nil map",
			input:  map[string]int(nil),
			panics: true,
		},
		{
			name:   "non-map type",
			input:  "not a map",
			panics: true,
		},
		{
			name:   "wrong key type",
			input:  map[int]string{1: "a"},
			panics: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.panics {
				assert.Panics(t, func() { anyx.MapKeysString(tt.input) })
			} else {
				result := anyx.MapKeysString(tt.input)
				assert.ElementsMatch(t, tt.expected, result)
			}
		})
	}
}

func TestMapKeysUint32(t *testing.T) {
	tests := []struct {
		name     string
		input    interface{}
		expected []uint32
		panics   bool
	}{
		{
			name:     "normal map",
			input:    map[uint32]int{1: 10, 2: 20},
			expected: []uint32{1, 2},
		},
		{
			name:     "empty map",
			input:    map[uint32]int{},
			expected: []uint32{},
		},
		{
			name:   "nil map",
			input:  map[uint32]int(nil),
			panics: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.panics {
				assert.Panics(t, func() { anyx.MapKeysUint32(tt.input) })
			} else {
				result := anyx.MapKeysUint32(tt.input)
				assert.ElementsMatch(t, tt.expected, result)
			}
		})
	}
}

func TestMapKeysUint64(t *testing.T) {
	tests := []struct {
		name     string
		input    interface{}
		expected []uint64
		panics   bool
	}{
		{
			name:     "normal map",
			input:    map[uint64]int{100: 10, 200: 20},
			expected: []uint64{100, 200},
		},
		{
			name:     "empty map",
			input:    map[uint64]int{},
			expected: []uint64{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.panics {
				assert.Panics(t, func() { anyx.MapKeysUint64(tt.input) })
			} else {
				result := anyx.MapKeysUint64(tt.input)
				assert.ElementsMatch(t, tt.expected, result)
			}
		})
	}
}

func TestMapKeysInt32(t *testing.T) {
	tests := []struct {
		name     string
		input    interface{}
		expected []int32
		panics   bool
	}{
		{
			name:     "normal map",
			input:    map[int32]int{-1: 10, 2: 20},
			expected: []int32{-1, 2},
		},
		{
			name:     "empty map",
			input:    map[int32]int{},
			expected: []int32{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.panics {
				assert.Panics(t, func() { anyx.MapKeysInt32(tt.input) })
			} else {
				result := anyx.MapKeysInt32(tt.input)
				assert.ElementsMatch(t, tt.expected, result)
			}
		})
	}
}

func TestMapKeysInt64(t *testing.T) {
	tests := []struct {
		name     string
		input    interface{}
		expected []int64
		panics   bool
	}{
		{
			name:     "normal map",
			input:    map[int64]int{-100: 10, 200: 20},
			expected: []int64{-100, 200},
		},
		{
			name:     "empty map",
			input:    map[int64]int{},
			expected: []int64{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.panics {
				assert.Panics(t, func() { anyx.MapKeysInt64(tt.input) })
			} else {
				result := anyx.MapKeysInt64(tt.input)
				assert.ElementsMatch(t, tt.expected, result)
			}
		})
	}
}

func TestMapValues(t *testing.T) {
	tests := []struct {
		name     string
		input    map[string]int
		expected []int
	}{
		{
			name:     "normal map",
			input:    map[string]int{"a": 1, "b": 2},
			expected: []int{1, 2},
		},
		{
			name:     "empty map",
			input:    map[string]int{},
			expected: []int{},
		},
		{
			name:     "nil map",
			input:    nil,
			expected: []int{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := anyx.MapValues(tt.input)
			assert.ElementsMatch(t, tt.expected, result)
		})
	}
}
