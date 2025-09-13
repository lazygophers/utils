package xtime_test

import (
	"testing"
	"time"

	"github.com/lazygophers/utils/xtime"
	"github.com/stretchr/testify/assert"
)

func TestConstants(t *testing.T) {
	t.Run("basic_time_constants", func(t *testing.T) {
		assert.Equal(t, time.Nanosecond, xtime.Nanosecond)
		assert.Equal(t, time.Microsecond, xtime.Microsecond)
		assert.Equal(t, time.Millisecond, xtime.Millisecond)
		assert.Equal(t, time.Second, xtime.Second)
		assert.Equal(t, time.Minute, xtime.Minute)
		assert.Equal(t, time.Hour, xtime.Hour)
	})

	t.Run("extended_time_constants", func(t *testing.T) {
		assert.Equal(t, time.Minute*30, xtime.HalfHour)
		assert.Equal(t, time.Hour*12, xtime.HalfDay)
		assert.Equal(t, time.Hour*24, xtime.Day)
	})

	t.Run("work_week_constants", func(t *testing.T) {
		assert.Equal(t, xtime.Day*5, xtime.WorkDayWeek)
		assert.Equal(t, xtime.Day*2, xtime.ResetDayWeek)
		assert.Equal(t, xtime.Day*7, xtime.Week)
	})

	t.Run("work_month_constants", func(t *testing.T) {
		expectedWorkMonth := xtime.Day*21 + xtime.HalfDay
		assert.Equal(t, expectedWorkMonth, xtime.WorkDayMonth)

		expectedResetMonth := xtime.Day*8 + xtime.HalfDay
		assert.Equal(t, expectedResetMonth, xtime.ResetDayMonth)

		assert.Equal(t, xtime.Day*30, xtime.Month)
	})

	t.Run("long_period_constants", func(t *testing.T) {
		assert.Equal(t, xtime.Day*91, xtime.QUARTER)
		assert.Equal(t, xtime.Day*365, xtime.Year)
		assert.Equal(t, xtime.Year*10+xtime.Day*2, xtime.Decade)
		assert.Equal(t, xtime.Year*100+xtime.Day*25, xtime.Century)
	})

	t.Run("constants_relationships", func(t *testing.T) {
		// Verify basic relationships
		assert.Equal(t, xtime.WorkDayWeek+xtime.ResetDayWeek, xtime.Week)
		assert.True(t, xtime.HalfHour*2 == xtime.Hour)
		assert.True(t, xtime.HalfDay*2 == xtime.Day)

		// Verify larger units are composed correctly
		assert.True(t, xtime.Year > xtime.Month*12)
		assert.True(t, xtime.Decade > xtime.Year*10)
		assert.True(t, xtime.Century > xtime.Year*100)
	})
}

func TestConstantsEdgeCases(t *testing.T) {
	t.Run("zero_duration", func(t *testing.T) {
		// Test that constants are non-zero
		assert.NotEqual(t, time.Duration(0), xtime.Nanosecond)
		assert.NotEqual(t, time.Duration(0), xtime.Day)
		assert.NotEqual(t, time.Duration(0), xtime.Year)
	})

	t.Run("positive_durations", func(t *testing.T) {
		// All constants should be positive
		assert.True(t, xtime.Nanosecond > 0)
		assert.True(t, xtime.Day > 0)
		assert.True(t, xtime.WorkDayWeek > 0)
		assert.True(t, xtime.Century > 0)
	})

	t.Run("reasonable_magnitudes", func(t *testing.T) {
		// Basic sanity checks
		assert.True(t, xtime.Minute > xtime.Second)
		assert.True(t, xtime.Hour > xtime.Minute)
		assert.True(t, xtime.Day > xtime.Hour)
		assert.True(t, xtime.Week > xtime.Day)
		assert.True(t, xtime.Month > xtime.Week)
		assert.True(t, xtime.Year > xtime.Month)
	})
}

// Benchmark constants access
func BenchmarkConstants(b *testing.B) {
	b.Run("basic_constants", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = xtime.Second
			_ = xtime.Minute
			_ = xtime.Hour
			_ = xtime.Day
		}
	})

	b.Run("computed_constants", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = xtime.WorkDayWeek
			_ = xtime.WorkDayMonth
			_ = xtime.QUARTER
			_ = xtime.Year
		}
	})
}
