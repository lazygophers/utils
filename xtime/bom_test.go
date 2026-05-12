package xtime

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestBeginningOfMonth(t *testing.T) {
	tests := []struct {
		name     string
		input    time.Time
		expected time.Time
	}{
		{
			name:     "middle of month",
			input:    time.Date(2024, 5, 15, 14, 30, 45, 123456789, time.Local),
			expected: time.Date(2024, 5, 1, 0, 0, 0, 0, time.Local),
		},
		{
			name:     "first day of month",
			input:    time.Date(2024, 5, 1, 23, 59, 59, 999999999, time.Local),
			expected: time.Date(2024, 5, 1, 0, 0, 0, 0, time.Local),
		},
		{
			name:     "last day of month",
			input:    time.Date(2024, 5, 31, 12, 0, 0, 0, time.Local),
			expected: time.Date(2024, 5, 1, 0, 0, 0, 0, time.Local),
		},
		{
			name:     "different months",
			input:    time.Date(2024, 1, 15, 10, 20, 30, 0, time.Local),
			expected: time.Date(2024, 1, 1, 0, 0, 0, 0, time.Local),
		},
		{
			name:     "february",
			input:    time.Date(2024, 2, 29, 15, 45, 0, 0, time.Local), // leap year
			expected: time.Date(2024, 2, 1, 0, 0, 0, 0, time.Local),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			xtime := With(tt.input)
			result := xtime.BeginningOfMonth()
			assert.Equal(t, tt.expected, result.Time)
			assert.Equal(t, xtime.Config, result.Config, "Config should be preserved")
		})
	}
}

func TestBeginningOfMonth_ConfigPreservation(t *testing.T) {
	customConfig := &Config{
		WeekStartDay:  time.Monday,
		TimeLocation:  time.UTC,
		TimeFormats:   []string{"2006-01-02"},
	}

	xtime := &Time{
		Time:   time.Date(2024, 5, 15, 14, 30, 45, 123456789, time.UTC),
		Config: customConfig,
	}

	result := xtime.BeginningOfMonth()

	assert.Equal(t, time.Date(2024, 5, 1, 0, 0, 0, 0, time.UTC), result.Time)
	assert.Equal(t, customConfig, result.Config, "Config should be the same reference")
}

func TestBeginningOfMonth_NilConfig(t *testing.T) {
	xtime := &Time{
		Time:   time.Date(2024, 5, 15, 14, 30, 45, 123456789, time.Local),
		Config: nil,
	}

	result := xtime.BeginningOfMonth()

	assert.Equal(t, time.Date(2024, 5, 1, 0, 0, 0, 0, time.Local), result.Time)
	assert.Nil(t, result.Config, "Config should remain nil")
}

func TestBeginningOfMonth_DifferentTimeZones(t *testing.T) {
	// Test UTC
	t.Run("UTC", func(t *testing.T) {
		xtime := With(time.Date(2024, 5, 15, 14, 30, 45, 0, time.UTC))
		result := xtime.BeginningOfMonth()
		assert.Equal(t, time.Date(2024, 5, 1, 0, 0, 0, 0, time.UTC), result.Time)
	})

	// Test Local
	t.Run("Local", func(t *testing.T) {
		xtime := With(time.Date(2024, 5, 15, 14, 30, 45, 0, time.Local))
		result := xtime.BeginningOfMonth()
		assert.Equal(t, time.Date(2024, 5, 1, 0, 0, 0, 0, time.Local), result.Time)
	})
}
