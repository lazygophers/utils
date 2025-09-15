package utils

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestMustOk(t *testing.T) {
	tests := []struct {
		name     string
		value    interface{}
		ok       bool
		expected interface{}
		panics   bool
	}{
		{
			name:     "ok_true_string",
			value:    "test",
			ok:       true,
			expected: "test",
			panics:   false,
		},
		{
			name:     "ok_true_int",
			value:    42,
			ok:       true,
			expected: 42,
			panics:   false,
		},
		{
			name:     "ok_true_nil",
			value:    nil,
			ok:       true,
			expected: nil,
			panics:   false,
		},
		{
			name:   "ok_false_string",
			value:  "test",
			ok:     false,
			panics: true,
		},
		{
			name:   "ok_false_int",
			value:  42,
			ok:     false,
			panics: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.panics {
				assert.Panics(t, func() {
					MustOk(tt.value, tt.ok)
				})
			} else {
				result := MustOk(tt.value, tt.ok)
				assert.Equal(t, tt.expected, result)
			}
		})
	}
}

func TestMustSuccess(t *testing.T) {
	tests := []struct {
		name   string
		err    error
		panics bool
	}{
		{
			name:   "nil_error",
			err:    nil,
			panics: false,
		},
		{
			name:   "non_nil_error",
			err:    errors.New("test error"),
			panics: true,
		},
		{
			name:   "wrapped_error",
			err:    errors.New("wrapped: test error"),
			panics: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.panics {
				assert.Panics(t, func() {
					MustSuccess(tt.err)
				})
			} else {
				assert.NotPanics(t, func() {
					MustSuccess(tt.err)
				})
			}
		})
	}
}

func TestMust(t *testing.T) {
	tests := []struct {
		name     string
		value    interface{}
		err      error
		expected interface{}
		panics   bool
	}{
		{
			name:     "success_string",
			value:    "test",
			err:      nil,
			expected: "test",
			panics:   false,
		},
		{
			name:     "success_int",
			value:    42,
			err:      nil,
			expected: 42,
			panics:   false,
		},
		{
			name:     "success_nil_value",
			value:    nil,
			err:      nil,
			expected: nil,
			panics:   false,
		},
		{
			name:     "success_struct",
			value:    struct{ Name string }{Name: "test"},
			err:      nil,
			expected: struct{ Name string }{Name: "test"},
			panics:   false,
		},
		{
			name:   "error_string",
			value:  "test",
			err:    errors.New("test error"),
			panics: true,
		},
		{
			name:   "error_int",
			value:  42,
			err:    errors.New("test error"),
			panics: true,
		},
		{
			name:   "error_nil_value",
			value:  nil,
			err:    errors.New("test error"),
			panics: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.panics {
				assert.Panics(t, func() {
					Must(tt.value, tt.err)
				})
			} else {
				result := Must(tt.value, tt.err)
				assert.Equal(t, tt.expected, result)
			}
		})
	}
}

func TestIgnore(t *testing.T) {
	tests := []struct {
		name     string
		value    interface{}
		ignored  interface{}
		expected interface{}
	}{
		{
			name:     "string_value_error_ignored",
			value:    "test",
			ignored:  errors.New("ignored error"),
			expected: "test",
		},
		{
			name:     "int_value_string_ignored",
			value:    42,
			ignored:  "ignored",
			expected: 42,
		},
		{
			name:     "nil_value_error_ignored",
			value:    nil,
			ignored:  errors.New("ignored error"),
			expected: nil,
		},
		{
			name:     "struct_value_bool_ignored",
			value:    struct{ Name string }{Name: "test"},
			ignored:  true,
			expected: struct{ Name string }{Name: "test"},
		},
		{
			name:     "slice_value_nil_ignored",
			value:    []int{1, 2, 3},
			ignored:  nil,
			expected: []int{1, 2, 3},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Ignore(tt.value, tt.ignored)
			assert.Equal(t, tt.expected, result)
		})
	}
}

// TestMustFunctions_TypeSafety tests that generic functions work with different types
func TestMustFunctions_TypeSafety(t *testing.T) {
	t.Run("string_type", func(t *testing.T) {
		result := Must("hello", nil)
		require.Equal(t, "hello", result)

		okResult := MustOk("world", true)
		require.Equal(t, "world", okResult)

		ignoreResult := Ignore("test", errors.New("ignored"))
		require.Equal(t, "test", ignoreResult)
	})

	t.Run("int_type", func(t *testing.T) {
		result := Must(123, nil)
		require.Equal(t, 123, result)

		okResult := MustOk(456, true)
		require.Equal(t, 456, okResult)

		ignoreResult := Ignore(789, "ignored")
		require.Equal(t, 789, ignoreResult)
	})

	t.Run("slice_type", func(t *testing.T) {
		slice := []string{"a", "b", "c"}
		result := Must(slice, nil)
		require.Equal(t, slice, result)

		okResult := MustOk(slice, true)
		require.Equal(t, slice, okResult)

		ignoreResult := Ignore(slice, nil)
		require.Equal(t, slice, ignoreResult)
	})

	t.Run("map_type", func(t *testing.T) {
		m := map[string]int{"key": 42}
		result := Must(m, nil)
		require.Equal(t, m, result)

		okResult := MustOk(m, true)
		require.Equal(t, m, okResult)

		ignoreResult := Ignore(m, true)
		require.Equal(t, m, ignoreResult)
	})
}
