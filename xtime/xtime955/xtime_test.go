package xtime955_test

import (
	"testing"
	"time"

	"github.com/lazygophers/utils/xtime/xtime955"
	"github.com/stretchr/testify/assert"
)

func TestBasicTimeConstants(t *testing.T) {
	t.Run("standard_time_constants", func(t *testing.T) {
		// Test that basic constants match standard library
		assert.Equal(t, time.Nanosecond, xtime955.Nanosecond)
		assert.Equal(t, time.Microsecond, xtime955.Microsecond)
		assert.Equal(t, time.Millisecond, xtime955.Millisecond)
		assert.Equal(t, time.Second, xtime955.Second)
		assert.Equal(t, time.Minute, xtime955.Minute)
	})

	t.Run("extended_constants", func(t *testing.T) {
		// Test extended time constants
		assert.Equal(t, time.Minute*30, xtime955.HalfHour)
		assert.Equal(t, time.Hour, xtime955.Hour)
		assert.Equal(t, time.Hour*24, xtime955.Day)
	})
}

func TestWorkTimeConstants(t *testing.T) {
	t.Run("work_day_constants", func(t *testing.T) {
		// In xtime955, WorkDay is 8 hours (9-to-5)
		assert.Equal(t, time.Hour*8, xtime955.WorkDay)
		assert.Equal(t, xtime955.Day-xtime955.WorkDay, xtime955.RestDay)
		assert.Equal(t, time.Hour*16, xtime955.RestDay) // 24 - 8 = 16 hours rest
	})

	t.Run("work_week_constants", func(t *testing.T) {
		// Week constants
		assert.Equal(t, xtime955.Day*7, xtime955.Week)
		assert.Equal(t, xtime955.WorkDay*5, xtime955.WorkWeek) // 5 work days
		assert.Equal(t, xtime955.Week-xtime955.WorkWeek, xtime955.RestWeek)

		// 5 * 8 hours work + 2 * 24 hours weekend + 5 * 16 hours daily rest
		expectedRestWeek := xtime955.Day*2 + xtime955.RestDay*5 // Weekend + daily rest
		assert.Equal(t, expectedRestWeek, xtime955.RestWeek)
	})

	t.Run("work_month_constants", func(t *testing.T) {
		// Month constants (22 work days)
		assert.Equal(t, xtime955.Day*30, xtime955.Month)
		assert.Equal(t, xtime955.Day*22, xtime955.WorkMonth) // 22 work days
		assert.Equal(t, xtime955.Month-xtime955.WorkMonth, xtime955.RestMonth)
		assert.Equal(t, xtime955.Day*8, xtime955.RestMonth) // 30 - 22 = 8 rest days
	})

	t.Run("work_quarter_constants", func(t *testing.T) {
		// Quarter constants
		assert.Equal(t, xtime955.Day*91, xtime955.Quarter)
		assert.Equal(t, xtime955.WorkMonth*3, xtime955.WorkQuarter) // 3 months of work
		assert.Equal(t, xtime955.Quarter-xtime955.WorkQuarter, xtime955.RestQuarter)

		expectedWorkQuarter := xtime955.Day * 22 * 3 // 22 work days * 3 months
		assert.Equal(t, expectedWorkQuarter, xtime955.WorkQuarter)
	})

	t.Run("work_year_constants", func(t *testing.T) {
		// Year constants (250 work days)
		assert.Equal(t, xtime955.Day*365, xtime955.Year)
		assert.Equal(t, xtime955.WorkDay*250, xtime955.WorkYear) // 250 work days
		assert.Equal(t, xtime955.Year-xtime955.WorkYear, xtime955.RestYear)

		expectedWorkYear := time.Hour * 8 * 250 // 8 hours * 250 days
		assert.Equal(t, expectedWorkYear, xtime955.WorkYear)
	})
}

func TestConstantRelationships(t *testing.T) {
	t.Run("time_unit_progression", func(t *testing.T) {
		// Test that larger units are composed of smaller ones
		assert.Equal(t, xtime955.Microsecond*1000, xtime955.Millisecond)
		assert.Equal(t, xtime955.Millisecond*1000, xtime955.Second)
		assert.Equal(t, xtime955.Second*60, xtime955.Minute)
		assert.Equal(t, xtime955.Minute*60, xtime955.Hour)
		assert.Equal(t, xtime955.Hour*24, xtime955.Day)
	})

	t.Run("work_rest_relationships", func(t *testing.T) {
		// Work + rest = total for each period
		assert.Equal(t, xtime955.WorkDay+xtime955.RestDay, xtime955.Day)
		assert.Equal(t, xtime955.WorkWeek+xtime955.RestWeek, xtime955.Week)
		assert.Equal(t, xtime955.WorkMonth+xtime955.RestMonth, xtime955.Month)
		assert.Equal(t, xtime955.WorkQuarter+xtime955.RestQuarter, xtime955.Quarter)
		assert.Equal(t, xtime955.WorkYear+xtime955.RestYear, xtime955.Year)
	})

	t.Run("half_hour_relationship", func(t *testing.T) {
		// HalfHour should be exactly half of Hour
		assert.Equal(t, xtime955.Hour/2, xtime955.HalfHour)
		assert.Equal(t, xtime955.HalfHour*2, xtime955.Hour)
	})

	t.Run("work_week_composition", func(t *testing.T) {
		// Work week should be 5 work days
		assert.Equal(t, xtime955.WorkDay*5, xtime955.WorkWeek)
		expectedWorkHours := time.Hour * 40 // 5 days * 8 hours
		assert.Equal(t, expectedWorkHours, xtime955.WorkWeek)
	})
}

func TestConstantValues(t *testing.T) {
	t.Run("positive_durations", func(t *testing.T) {
		// All constants should be positive
		assert.True(t, xtime955.Nanosecond > 0)
		assert.True(t, xtime955.Day > 0)
		assert.True(t, xtime955.WorkDay > 0)
		assert.True(t, xtime955.RestDay > 0)
		assert.True(t, xtime955.Week > 0)
		assert.True(t, xtime955.Month > 0)
		assert.True(t, xtime955.Year > 0)
	})

	t.Run("work_vs_rest_durations", func(t *testing.T) {
		// In 955 model, rest time should be significant
		assert.True(t, xtime955.RestDay > 0, "Should have rest time in a day")
		assert.True(t, xtime955.RestWeek > xtime955.WorkWeek, "Rest week should be longer than work week")
		assert.True(t, xtime955.RestMonth > 0, "Should have rest days in a month")
		assert.True(t, xtime955.RestYear > 0, "Should have rest days in a year")
	})

	t.Run("specific_work_durations", func(t *testing.T) {
		// Test specific work duration values
		assert.Equal(t, time.Hour*8, xtime955.WorkDay)       // 8-hour work day
		assert.Equal(t, time.Hour*40, xtime955.WorkWeek)     // 40-hour work week
		assert.Equal(t, time.Hour*24*22, xtime955.WorkMonth) // 22 work days (full days, not just work hours)
		assert.Equal(t, time.Hour*8*250, xtime955.WorkYear)  // 250 work days
	})

	t.Run("specific_rest_durations", func(t *testing.T) {
		// Test specific rest duration values
		assert.Equal(t, time.Hour*16, xtime955.RestDay)         // 16 hours rest per day
		expectedRestWeek := xtime955.Day*2 + xtime955.RestDay*5 // Weekend + daily rest
		assert.Equal(t, expectedRestWeek, xtime955.RestWeek)
		assert.Equal(t, xtime955.Day*8, xtime955.RestMonth) // 8 rest days per month
	})
}

func TestConstantMagnitudes(t *testing.T) {
	t.Run("reasonable_magnitudes", func(t *testing.T) {
		// Test that constants have reasonable magnitudes
		assert.True(t, xtime955.Minute > xtime955.Second)
		assert.True(t, xtime955.Hour > xtime955.Minute)
		assert.True(t, xtime955.Day > xtime955.Hour)
		assert.True(t, xtime955.Week > xtime955.Day)
		assert.True(t, xtime955.Month > xtime955.Week)
		assert.True(t, xtime955.Quarter > xtime955.Month)
		assert.True(t, xtime955.Year > xtime955.Quarter)
	})

	t.Run("work_vs_rest_comparisons", func(t *testing.T) {
		// Work should be less than rest for most periods in 955 model
		assert.True(t, xtime955.WorkDay < xtime955.RestDay, "Work day should be shorter than rest day")
		assert.True(t, xtime955.WorkWeek < xtime955.RestWeek, "Work week should be shorter than rest week")
		// Month and year depend on specific calculations
	})

	t.Run("period_compositions", func(t *testing.T) {
		// Test larger period compositions
		assert.Equal(t, xtime955.Day*7, xtime955.Week)
		assert.Equal(t, xtime955.Day*30, xtime955.Month)
		assert.Equal(t, xtime955.Day*91, xtime955.Quarter)
		assert.Equal(t, xtime955.Day*365, xtime955.Year)
	})
}

func TestConstantUsageScenarios(t *testing.T) {
	t.Run("work_calculations", func(t *testing.T) {
		// Calculate work hours in different periods
		workHoursInWeek := xtime955.WorkWeek / xtime955.Hour
		assert.Equal(t, time.Duration(40), workHoursInWeek) // 40-hour work week

		workHoursInMonth := xtime955.WorkMonth / xtime955.Hour
		assert.Equal(t, time.Duration(22*24), workHoursInMonth) // 22 work days * 24 hours per day

		workDaysInYear := xtime955.WorkYear / xtime955.WorkDay
		assert.Equal(t, time.Duration(250), workDaysInYear) // 250 work days
	})

	t.Run("rest_calculations", func(t *testing.T) {
		// Calculate rest hours in different periods
		restHoursInDay := xtime955.RestDay / xtime955.Hour
		assert.Equal(t, time.Duration(16), restHoursInDay) // 16 hours rest per day

		restDaysInMonth := xtime955.RestMonth / xtime955.Day
		assert.Equal(t, time.Duration(8), restDaysInMonth) // 8 rest days per month

		// Rest time in year calculation
		// Year = 365 days * 24 hours/day = 8760 hours
		// WorkYear = 250 work days * 8 hours/day = 2000 hours
		// RestYear = Year - WorkYear = 8760 - 2000 = 6760 hours
		restHoursInYear := (xtime955.Year - xtime955.WorkYear) / xtime955.Hour
		expectedRestHours := 365*24 - 250*8 // Total hours - work hours
		assert.Equal(t, time.Duration(expectedRestHours), restHoursInYear)
	})

	t.Run("efficiency_calculations", func(t *testing.T) {
		// Calculate work efficiency ratios
		dailyWorkRatio := float64(xtime955.WorkDay) / float64(xtime955.Day)
		assert.InDelta(t, 1.0/3.0, dailyWorkRatio, 0.01) // ~33% work time per day

		weeklyWorkRatio := float64(xtime955.WorkWeek) / float64(xtime955.Week)
		expectedWeeklyRatio := 40.0 / 168.0 // 40 hours / 168 hours per week
		assert.InDelta(t, expectedWeeklyRatio, weeklyWorkRatio, 0.01)

		yearlyWorkRatio := float64(xtime955.WorkYear) / float64(xtime955.Year)
		expectedYearlyRatio := 250.0 * 8.0 / (365.0 * 24.0) // Work hours / total hours
		assert.InDelta(t, expectedYearlyRatio, yearlyWorkRatio, 0.01)
	})

	t.Run("duration_arithmetic", func(t *testing.T) {
		// Test arithmetic operations with constants
		overtimeHours := time.Hour * 2
		extendedWorkDay := xtime955.WorkDay + overtimeHours
		assert.Equal(t, time.Hour*10, extendedWorkDay)

		weekendRest := xtime955.Day * 2
		assert.True(t, weekendRest < xtime955.RestWeek) // Weekend is part of rest week

		monthlyOvertime := xtime955.WorkMonth + xtime955.WorkWeek
		expectedOvertimeHours := 22*24 + 40 // Regular month + extra week
		assert.Equal(t, time.Hour*time.Duration(expectedOvertimeHours), monthlyOvertime)
	})
}

func TestConstantTypes(t *testing.T) {
	t.Run("duration_types", func(t *testing.T) {
		// All constants should be of type time.Duration
		assert.IsType(t, time.Duration(0), xtime955.Nanosecond)
		assert.IsType(t, time.Duration(0), xtime955.Day)
		assert.IsType(t, time.Duration(0), xtime955.WorkDay)
		assert.IsType(t, time.Duration(0), xtime955.RestDay)
		assert.IsType(t, time.Duration(0), xtime955.WorkWeek)
		assert.IsType(t, time.Duration(0), xtime955.RestWeek)
		assert.IsType(t, time.Duration(0), xtime955.WorkMonth)
		assert.IsType(t, time.Duration(0), xtime955.RestMonth)
		assert.IsType(t, time.Duration(0), xtime955.WorkQuarter)
		assert.IsType(t, time.Duration(0), xtime955.RestQuarter)
		assert.IsType(t, time.Duration(0), xtime955.WorkYear)
		assert.IsType(t, time.Duration(0), xtime955.RestYear)
	})
}

func TestWorkLifeBalance(t *testing.T) {
	t.Run("daily_balance", func(t *testing.T) {
		// 8 hours work, 16 hours personal time (including sleep)
		totalDailyTime := xtime955.WorkDay + xtime955.RestDay
		assert.Equal(t, xtime955.Day, totalDailyTime)

		workPercentage := float64(xtime955.WorkDay) * 100 / float64(xtime955.Day)
		assert.InDelta(t, 33.33, workPercentage, 0.1) // ~33% work time
	})

	t.Run("weekly_balance", func(t *testing.T) {
		// 40 hours work, 128 hours personal time
		totalWeeklyTime := xtime955.WorkWeek + xtime955.RestWeek
		assert.Equal(t, xtime955.Week, totalWeeklyTime)

		workPercentage := float64(xtime955.WorkWeek) * 100 / float64(xtime955.Week)
		expectedPercentage := 40.0 * 100 / 168.0 // 40 hours out of 168
		assert.InDelta(t, expectedPercentage, workPercentage, 0.1)
	})

	t.Run("yearly_balance", func(t *testing.T) {
		// 250 work days, 115 rest days
		totalYearlyTime := xtime955.WorkYear + xtime955.RestYear
		assert.Equal(t, xtime955.Year, totalYearlyTime)

		workDays := xtime955.WorkYear / xtime955.WorkDay
		assert.Equal(t, time.Duration(250), workDays)

		// Rest includes weekends, holidays, and daily rest time
		// This is more complex but should be consistent
	})
}

// Benchmark constant access
func BenchmarkConstants955(b *testing.B) {
	b.Run("basic_constants", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = xtime955.Second
			_ = xtime955.Minute
			_ = xtime955.Hour
			_ = xtime955.Day
		}
	})

	b.Run("work_constants", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = xtime955.WorkDay
			_ = xtime955.WorkWeek
			_ = xtime955.WorkMonth
			_ = xtime955.WorkYear
		}
	})

	b.Run("rest_constants", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = xtime955.RestDay
			_ = xtime955.RestWeek
			_ = xtime955.RestMonth
			_ = xtime955.RestYear
		}
	})

	b.Run("calculations", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = xtime955.WorkYear / xtime955.WorkDay
			_ = xtime955.RestWeek / xtime955.RestDay
			_ = float64(xtime955.WorkDay) / float64(xtime955.Day)
		}
	})
}
