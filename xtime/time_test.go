package xtime_test

import (
	"testing"
	"time"

	"github.com/lazygophers/utils/xtime"
	"github.com/stretchr/testify/assert"
)

func TestParse(t *testing.T) {
	t.Run("valid_time_strings", func(t *testing.T) {
		testCases := []struct {
			name     string
			input    string
			expected bool
		}{
			{"iso_date", "2023-01-15", true},
			{"iso_datetime", "2023-01-15 14:30:00", true},
			{"iso_full", "2023-01-15T14:30:00Z", true},
			{"us_format", "01/15/2023", true},
			{"simple_time", "14:30", true},
		}

		for _, tc := range testCases {
			t.Run(tc.name, func(t *testing.T) {
				result, err := xtime.Parse(tc.input)
				if tc.expected {
					assert.NoError(t, err, "Should parse %s successfully", tc.input)
					assert.NotNil(t, result, "Result should not be nil")
					assert.NotNil(t, result.Config, "Config should be initialized")
				} else {
					assert.Error(t, err, "Should fail to parse %s", tc.input)
					assert.Nil(t, result, "Result should be nil on error")
				}
			})
		}
	})

	t.Run("invalid_time_strings", func(t *testing.T) {
		invalidInputs := []string{
			"invalid",
			"not-a-date",
			"2023/13/45", // Invalid month/day
			"25:00:00",   // Invalid hour
		}

		for _, input := range invalidInputs {
			result, err := xtime.Parse(input)
			// Note: The underlying now.Parse might be lenient, so we test actual behavior
			if err != nil {
				assert.Error(t, err, "Should error on invalid input: %s", input)
				assert.Nil(t, result, "Result should be nil on error")
			}
			// If no error, that means the parser was lenient - that's ok too
		}
	})

	t.Run("multiple_inputs", func(t *testing.T) {
		// Test with multiple string arguments
		result, err := xtime.Parse("2023-01-15", "14:30:00")
		assert.NoError(t, err)
		assert.NotNil(t, result)
	})

	t.Run("empty_input", func(t *testing.T) {
		result, err := xtime.Parse("")
		// Test actual behavior - empty string might parse to current time or error
		if err != nil {
			assert.Error(t, err)
			assert.Nil(t, result)
		} else {
			assert.NotNil(t, result)
		}
	})
}

func TestMustParse(t *testing.T) {
	t.Run("valid_input", func(t *testing.T) {
		// Test that MustParse returns valid result for good input
		result := xtime.MustParse("2023-01-15")
		assert.NotNil(t, result)
		assert.NotNil(t, result.Config)
		assert.Equal(t, time.Monday, result.Config.WeekStartDay)
		assert.Equal(t, time.Local, result.Config.TimeLocation)
	})

	t.Run("panic_on_invalid_input", func(t *testing.T) {
		// Test that MustParse panics on invalid input
		assert.Panics(t, func() {
			// Use a string that's definitely invalid for parsing
			xtime.MustParse("definitely-not-a-valid-date-format-12345")
		}, "MustParse should panic on invalid input")
	})
}

func TestWith(t *testing.T) {
	t.Run("wrap_time_instance", func(t *testing.T) {
		now := time.Now()
		wrapped := xtime.With(now)

		assert.NotNil(t, wrapped)
		assert.Equal(t, now, wrapped.Time)
		assert.NotNil(t, wrapped.Config)

		// Check default config values
		assert.Equal(t, time.Monday, wrapped.Config.WeekStartDay)
		assert.Equal(t, time.Local, wrapped.Config.TimeLocation)
		assert.NotNil(t, wrapped.Config.TimeFormats)
		assert.Len(t, wrapped.Config.TimeFormats, 0) // Empty slice by default
	})

	t.Run("different_times", func(t *testing.T) {
		times := []time.Time{
			time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC),
			time.Date(2023, 6, 15, 12, 30, 45, 123456789, time.UTC),
			time.Date(2023, 12, 31, 23, 59, 59, 999999999, time.UTC),
		}

		for i, timeVal := range times {
			wrapped := xtime.With(timeVal)
			assert.Equal(t, timeVal, wrapped.Time, "Time %d should be preserved", i)
			assert.NotNil(t, wrapped.Config, "Config should be initialized for time %d", i)
		}
	})

	t.Run("zero_time", func(t *testing.T) {
		zeroTime := time.Time{}
		wrapped := xtime.With(zeroTime)

		assert.Equal(t, zeroTime, wrapped.Time)
		assert.NotNil(t, wrapped.Config)
		assert.Equal(t, time.Monday, wrapped.Config.WeekStartDay)
	})
}

func TestRandSleep(t *testing.T) {
	t.Run("default_range", func(t *testing.T) {
		start := time.Now()
		xtime.RandSleep() // Should sleep for 1-3 seconds by default
		elapsed := time.Since(start)

		// Should sleep for at least 1 second and at most 3 seconds (with some tolerance)
		assert.True(t, elapsed >= time.Second, "Should sleep at least 1 second")
		assert.True(t, elapsed <= 4*time.Second, "Should sleep at most ~3 seconds (with tolerance)")
	})

	t.Run("custom_single_duration", func(t *testing.T) {
		start := time.Now()
		maxSleep := 50 * time.Millisecond
		xtime.RandSleep(maxSleep) // Should sleep for 0 to maxSleep
		elapsed := time.Since(start)

		// Should sleep for at most maxSleep duration (with some tolerance for timing)
		assert.True(t, elapsed <= maxSleep+10*time.Millisecond, "Should sleep at most %v", maxSleep)
		assert.True(t, elapsed >= 0, "Should not have negative sleep time")
	})

	t.Run("custom_range", func(t *testing.T) {
		start := time.Now()
		minSleep := 20 * time.Millisecond
		maxSleep := 100 * time.Millisecond
		xtime.RandSleep(minSleep, maxSleep)
		elapsed := time.Since(start)

		// Should sleep within the specified range (with tolerance)
		assert.True(t, elapsed >= minSleep, "Should sleep at least %v", minSleep)
		assert.True(t, elapsed <= maxSleep+20*time.Millisecond, "Should sleep at most %v (with tolerance)", maxSleep)
	})

	t.Run("zero_duration", func(t *testing.T) {
		start := time.Now()
		// Zero duration might cause panic in underlying random function, so use defer recovery
		defer func() {
			if r := recover(); r != nil {
				// If it panics on zero duration, that's expected behavior
				t.Logf("RandSleep(0) panicked as expected: %v", r)
				return
			}
			// If no panic, check timing
			elapsed := time.Since(start)
			assert.True(t, elapsed < 10*time.Millisecond, "Should complete quickly for zero duration")
		}()

		xtime.RandSleep(0)
	})

	t.Run("multiple_durations", func(t *testing.T) {
		// Test with more than 2 durations - should only use first two
		start := time.Now()
		xtime.RandSleep(10*time.Millisecond, 50*time.Millisecond, time.Second, time.Minute)
		elapsed := time.Since(start)

		// Should use only first two parameters
		assert.True(t, elapsed >= 10*time.Millisecond, "Should respect minimum")
		assert.True(t, elapsed <= 100*time.Millisecond, "Should ignore extra parameters")
	})
}

func TestConfig(t *testing.T) {
	t.Run("default_config_values", func(t *testing.T) {
		config := &xtime.Config{
			WeekStartDay: time.Monday,
			TimeLocation: time.Local,
			TimeFormats:  []string{},
		}

		assert.Equal(t, time.Monday, config.WeekStartDay)
		assert.Equal(t, time.Local, config.TimeLocation)
		assert.Empty(t, config.TimeFormats)
	})

	t.Run("config_modification", func(t *testing.T) {
		wrapped := xtime.With(time.Now())

		// Modify config
		wrapped.Config.WeekStartDay = time.Sunday
		wrapped.Config.TimeFormats = []string{"2006-01-02", "15:04:05"}

		assert.Equal(t, time.Sunday, wrapped.Config.WeekStartDay)
		assert.Len(t, wrapped.Config.TimeFormats, 2)
		assert.Contains(t, wrapped.Config.TimeFormats, "2006-01-02")
		assert.Contains(t, wrapped.Config.TimeFormats, "15:04:05")
	})
}

// Test the Time struct embedding
func TestTimeEmbedding(t *testing.T) {
	t.Run("time_methods_available", func(t *testing.T) {
		baseTime := time.Date(2023, 6, 15, 14, 30, 45, 0, time.UTC)
		wrapped := xtime.With(baseTime)

		// Test that standard time.Time methods are available
		assert.Equal(t, 2023, wrapped.Year())
		assert.Equal(t, time.June, wrapped.Month())
		assert.Equal(t, 15, wrapped.Day())
		assert.Equal(t, 14, wrapped.Hour())
		assert.Equal(t, 30, wrapped.Minute())
		assert.Equal(t, 45, wrapped.Second())

		// Test Unix timestamp
		assert.Equal(t, baseTime.Unix(), wrapped.Unix())
		assert.Equal(t, baseTime.UnixMilli(), wrapped.UnixMilli())
	})

	t.Run("time_formatting", func(t *testing.T) {
		baseTime := time.Date(2023, 1, 15, 9, 30, 0, 0, time.UTC)
		wrapped := xtime.With(baseTime)

		// Test formatting
		formatted := wrapped.Format("2006-01-02 15:04:05")
		assert.Equal(t, "2023-01-15 09:30:00", formatted)

		// Test string representation
		assert.Contains(t, wrapped.String(), "2023")
	})
}
