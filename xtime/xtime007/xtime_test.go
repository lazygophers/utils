package xtime007_test

import (
	"testing"
	"time"

	"github.com/lazygophers/utils/xtime/xtime007"
	"github.com/stretchr/testify/assert"
)

func TestBasicTimeConstants(t *testing.T) {
	t.Run("standard_time_constants", func(t *testing.T) {
		// Test that basic constants match standard library
		assert.Equal(t, time.Nanosecond, xtime007.Nanosecond)
		assert.Equal(t, time.Microsecond, xtime007.Microsecond)
		assert.Equal(t, time.Millisecond, xtime007.Millisecond)
		assert.Equal(t, time.Second, xtime007.Second)
		assert.Equal(t, time.Minute, xtime007.Minute)
		assert.Equal(t, time.Hour, xtime007.Hour)
	})

	t.Run("extended_constants", func(t *testing.T) {
		// Test extended time constants
		assert.Equal(t, time.Minute*30, xtime007.HalfHour)
		assert.Equal(t, time.Hour*24, xtime007.Day)
	})
}

func TestWorkTimeConstants(t *testing.T) {
	t.Run("work_day_constants", func(t *testing.T) {
		// In xtime007, WorkDay equals full Day (24/7 model)
		assert.Equal(t, xtime007.Day, xtime007.WorkDay)
		assert.Equal(t, time.Hour*24, xtime007.WorkDay)
		
		// RestDay should be Day - WorkDay = 0 (24/7 model)
		assert.Equal(t, xtime007.Day-xtime007.WorkDay, xtime007.RestDay)
		assert.Equal(t, time.Duration(0), xtime007.RestDay)
	})

	t.Run("work_week_constants", func(t *testing.T) {
		// Week constants
		assert.Equal(t, xtime007.Day*7, xtime007.Week)
		assert.Equal(t, xtime007.WorkDay*7, xtime007.WorkWeek)
		assert.Equal(t, xtime007.Week-xtime007.WorkWeek, xtime007.RestWeek)
		
		// In 24/7 model, WorkWeek equals Week
		assert.Equal(t, xtime007.Week, xtime007.WorkWeek)
		assert.Equal(t, time.Duration(0), xtime007.RestWeek)
	})

	t.Run("work_month_constants", func(t *testing.T) {
		// Month constants
		assert.Equal(t, xtime007.Day*30, xtime007.Month)
		assert.Equal(t, xtime007.Month, xtime007.WorkMonth)
		assert.Equal(t, time.Duration(0), xtime007.RestMonth)
	})

	t.Run("work_quarter_constants", func(t *testing.T) {
		// Quarter constants
		assert.Equal(t, xtime007.Day*91, xtime007.Quarter)
		assert.Equal(t, xtime007.Quarter, xtime007.WorkQuarter)
		assert.Equal(t, time.Duration(0), xtime007.RestQuarter)
	})

	t.Run("work_year_constants", func(t *testing.T) {
		// Year constants
		assert.Equal(t, xtime007.Day*365, xtime007.Year)
		assert.Equal(t, xtime007.Year, xtime007.WorkYear)
		assert.Equal(t, time.Duration(0), xtime007.RestYear)
	})
}

func TestConstantRelationships(t *testing.T) {
	t.Run("time_unit_progression", func(t *testing.T) {
		// Test that larger units are composed of smaller ones
		assert.Equal(t, xtime007.Microsecond*1000, xtime007.Millisecond)
		assert.Equal(t, xtime007.Millisecond*1000, xtime007.Second)
		assert.Equal(t, xtime007.Second*60, xtime007.Minute)
		assert.Equal(t, xtime007.Minute*60, xtime007.Hour)
		assert.Equal(t, xtime007.Hour*24, xtime007.Day)
	})

	t.Run("work_rest_relationships", func(t *testing.T) {
		// In 24/7 model, work + rest = total for each period
		assert.Equal(t, xtime007.WorkDay+xtime007.RestDay, xtime007.Day)
		assert.Equal(t, xtime007.WorkWeek+xtime007.RestWeek, xtime007.Week)
		assert.Equal(t, xtime007.WorkMonth+xtime007.RestMonth, xtime007.Month)
		assert.Equal(t, xtime007.WorkQuarter+xtime007.RestQuarter, xtime007.Quarter)
		assert.Equal(t, xtime007.WorkYear+xtime007.RestYear, xtime007.Year)
	})

	t.Run("period_compositions", func(t *testing.T) {
		// Test larger period compositions
		assert.Equal(t, xtime007.Day*7, xtime007.Week)
		assert.Equal(t, xtime007.Day*30, xtime007.Month)
		assert.Equal(t, xtime007.Day*91, xtime007.Quarter)
		assert.Equal(t, xtime007.Day*365, xtime007.Year)
	})

	t.Run("half_hour_relationship", func(t *testing.T) {
		// HalfHour should be exactly half of Hour
		assert.Equal(t, xtime007.Hour/2, xtime007.HalfHour)
		assert.Equal(t, xtime007.HalfHour*2, xtime007.Hour)
	})
}

func TestConstantValues(t *testing.T) {
	t.Run("positive_durations", func(t *testing.T) {
		// All positive constants should be positive
		assert.True(t, xtime007.Nanosecond > 0)
		assert.True(t, xtime007.Day > 0)
		assert.True(t, xtime007.WorkDay > 0)
		assert.True(t, xtime007.Week > 0)
		assert.True(t, xtime007.Month > 0)
		assert.True(t, xtime007.Year > 0)
	})

	t.Run("zero_rest_durations", func(t *testing.T) {
		// All rest durations should be zero in 24/7 model
		assert.Equal(t, time.Duration(0), xtime007.RestDay)
		assert.Equal(t, time.Duration(0), xtime007.RestWeek)
		assert.Equal(t, time.Duration(0), xtime007.RestMonth)
		assert.Equal(t, time.Duration(0), xtime007.RestQuarter)
		assert.Equal(t, time.Duration(0), xtime007.RestYear)
	})

	t.Run("work_equals_total", func(t *testing.T) {
		// In 24/7 model, work time equals total time
		assert.Equal(t, xtime007.Day, xtime007.WorkDay)
		assert.Equal(t, xtime007.Week, xtime007.WorkWeek)
		assert.Equal(t, xtime007.Month, xtime007.WorkMonth)
		assert.Equal(t, xtime007.Quarter, xtime007.WorkQuarter)
		assert.Equal(t, xtime007.Year, xtime007.WorkYear)
	})
}

func TestConstantMagnitudes(t *testing.T) {
	t.Run("reasonable_magnitudes", func(t *testing.T) {
		// Test that constants have reasonable magnitudes
		assert.True(t, xtime007.Minute > xtime007.Second)
		assert.True(t, xtime007.Hour > xtime007.Minute)
		assert.True(t, xtime007.Day > xtime007.Hour)
		assert.True(t, xtime007.Week > xtime007.Day)
		assert.True(t, xtime007.Month > xtime007.Week)
		assert.True(t, xtime007.Quarter > xtime007.Month)
		assert.True(t, xtime007.Year > xtime007.Quarter)
	})

	t.Run("specific_durations", func(t *testing.T) {
		// Test specific duration values
		assert.Equal(t, time.Duration(30)*time.Minute, xtime007.HalfHour)
		assert.Equal(t, time.Duration(24)*time.Hour, xtime007.Day)
		assert.Equal(t, time.Duration(7)*xtime007.Day, xtime007.Week)
		assert.Equal(t, time.Duration(30)*xtime007.Day, xtime007.Month)
		assert.Equal(t, time.Duration(91)*xtime007.Day, xtime007.Quarter)
		assert.Equal(t, time.Duration(365)*xtime007.Day, xtime007.Year)
	})
}

func TestConstantUsageScenarios(t *testing.T) {
	t.Run("time_calculations", func(t *testing.T) {
		// Test using constants in typical calculations
		
		// Calculate work hours in a month (24/7 model)
		workHoursInMonth := xtime007.WorkMonth / xtime007.Hour
		expectedHours := 30 * 24 // 30 days * 24 hours
		assert.Equal(t, time.Duration(expectedHours), workHoursInMonth)
		
		// Calculate work days in a year
		workDaysInYear := xtime007.WorkYear / xtime007.WorkDay
		assert.Equal(t, time.Duration(365), workDaysInYear)
		
		// Calculate rest time (should be zero)
		totalRestInYear := xtime007.RestYear
		assert.Equal(t, time.Duration(0), totalRestInYear)
	})

	t.Run("duration_arithmetic", func(t *testing.T) {
		// Test arithmetic operations with constants
		twoWeeks := xtime007.Week * 2
		assert.Equal(t, xtime007.Day*14, twoWeeks)
		
		quarterYear := xtime007.Year / 4
		// Should be close to Quarter, but not exactly equal due to rounding
		assert.True(t, quarterYear > xtime007.Quarter-xtime007.Day)
		assert.True(t, quarterYear < xtime007.Quarter+xtime007.Day*3)
		
		// Half day calculations
		halfDay := xtime007.Day / 2
		assert.Equal(t, xtime007.Hour*12, halfDay)
	})

	t.Run("comparison_operations", func(t *testing.T) {
		// Test comparison operations
		assert.True(t, xtime007.WorkDay >= xtime007.RestDay)
		assert.True(t, xtime007.WorkWeek >= xtime007.RestWeek)
		assert.True(t, xtime007.WorkMonth >= xtime007.RestMonth)
		assert.True(t, xtime007.WorkYear >= xtime007.RestYear)
		
		// All work periods should be greater than zero
		assert.True(t, xtime007.WorkDay > 0)
		assert.True(t, xtime007.WorkWeek > 0)
		assert.True(t, xtime007.WorkMonth > 0)
		assert.True(t, xtime007.WorkYear > 0)
	})
}

func TestConstantTypes(t *testing.T) {
	t.Run("duration_types", func(t *testing.T) {
		// All constants should be of type time.Duration
		assert.IsType(t, time.Duration(0), xtime007.Nanosecond)
		assert.IsType(t, time.Duration(0), xtime007.Day)
		assert.IsType(t, time.Duration(0), xtime007.WorkDay)
		assert.IsType(t, time.Duration(0), xtime007.RestDay)
		assert.IsType(t, time.Duration(0), xtime007.Week)
		assert.IsType(t, time.Duration(0), xtime007.Month)
		assert.IsType(t, time.Duration(0), xtime007.Quarter)
		assert.IsType(t, time.Duration(0), xtime007.Year)
	})
}

// Benchmark constant access (should be very fast)
func BenchmarkConstants(b *testing.B) {
	b.Run("basic_constants", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = xtime007.Second
			_ = xtime007.Minute
			_ = xtime007.Hour
			_ = xtime007.Day
		}
	})

	b.Run("work_constants", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = xtime007.WorkDay
			_ = xtime007.WorkWeek
			_ = xtime007.WorkMonth
			_ = xtime007.WorkYear
		}
	})

	b.Run("rest_constants", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = xtime007.RestDay
			_ = xtime007.RestWeek
			_ = xtime007.RestMonth
			_ = xtime007.RestYear
		}
	})

	b.Run("calculations", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = xtime007.WorkYear / xtime007.WorkDay
			_ = xtime007.WorkMonth / xtime007.WorkDay
			_ = xtime007.WorkWeek / xtime007.WorkDay
		}
	})
}