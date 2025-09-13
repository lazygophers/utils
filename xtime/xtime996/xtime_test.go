package xtime996_test

import (
	"testing"
	"time"

	"github.com/lazygophers/utils/xtime/xtime996"
	"github.com/stretchr/testify/assert"
)

func TestBasicTimeConstants(t *testing.T) {
	t.Run("standard_time_constants", func(t *testing.T) {
		// Test that basic constants match standard library
		assert.Equal(t, time.Nanosecond, xtime996.Nanosecond)
		assert.Equal(t, time.Microsecond, xtime996.Microsecond)
		assert.Equal(t, time.Millisecond, xtime996.Millisecond)
		assert.Equal(t, time.Second, xtime996.Second)
		assert.Equal(t, time.Minute, xtime996.Minute)
	})

	t.Run("extended_constants", func(t *testing.T) {
		// Test extended time constants
		assert.Equal(t, time.Minute*30, xtime996.HalfHour)
		assert.Equal(t, time.Hour, xtime996.Hour)
		assert.Equal(t, time.Hour*24, xtime996.Day)
	})
}

func TestWorkTimeConstants(t *testing.T) {
	t.Run("work_day_constants", func(t *testing.T) {
		// In xtime996, WorkDay is 12 hours (9am-9pm)
		assert.Equal(t, time.Hour*12, xtime996.WorkDay)
		assert.Equal(t, xtime996.Day-xtime996.WorkDay, xtime996.RestDay)
		assert.Equal(t, time.Hour*12, xtime996.RestDay) // 24 - 12 = 12 hours rest
	})

	t.Run("work_week_constants", func(t *testing.T) {
		// Week constants - 6 days work, 1 day rest
		assert.Equal(t, xtime996.Day*7, xtime996.Week)
		assert.Equal(t, xtime996.WorkDay*6, xtime996.WorkWeek) // 6 work days
		assert.Equal(t, xtime996.Week-xtime996.WorkWeek, xtime996.RestWeek)

		// 6 * 12 hours work + 1 * 24 hours weekend + 6 * 12 hours daily rest
		expectedRestWeek := xtime996.Day*1 + xtime996.RestDay*6 // Weekend + daily rest
		assert.Equal(t, expectedRestWeek, xtime996.RestWeek)
	})

	t.Run("work_month_constants", func(t *testing.T) {
		// Month constants
		assert.Equal(t, xtime996.Day*30, xtime996.Month)

		// RestMonth is 4 rest days * 12 hours each
		assert.Equal(t, xtime996.RestDay*4, xtime996.RestMonth)
		assert.Equal(t, time.Hour*12*4, xtime996.RestMonth) // 4 * 12 = 48 hours

		// WorkMonth = 30 days - 4 rest days = 26 full days
		expectedWorkMonth := xtime996.Day*30 - xtime996.RestMonth
		assert.Equal(t, expectedWorkMonth, xtime996.WorkMonth)
	})

	t.Run("work_quarter_constants", func(t *testing.T) {
		// Quarter constants
		assert.Equal(t, xtime996.Day*91, xtime996.Quarter)

		// RestQuarter is 14 rest days * 12 hours each
		assert.Equal(t, xtime996.RestDay*14, xtime996.RestQuarter)
		assert.Equal(t, time.Hour*12*14, xtime996.RestQuarter) // 14 * 12 = 168 hours

		// WorkQuarter = 91 days - 14 rest days = 77 full days
		expectedWorkQuarter := xtime996.Day*91 - xtime996.RestQuarter
		assert.Equal(t, expectedWorkQuarter, xtime996.WorkQuarter)
	})

	t.Run("work_year_constants", func(t *testing.T) {
		// Year constants
		assert.Equal(t, xtime996.Day*365, xtime996.Year)

		// RestYear is 58 rest days * 12 hours each
		assert.Equal(t, xtime996.RestDay*58, xtime996.RestYear)
		assert.Equal(t, time.Hour*12*58, xtime996.RestYear) // 58 * 12 = 696 hours

		// WorkYear = 365 days - 58 rest days = 307 full days
		expectedWorkYear := xtime996.Year - xtime996.RestYear
		assert.Equal(t, expectedWorkYear, xtime996.WorkYear)
	})
}

func TestConstantRelationships(t *testing.T) {
	t.Run("time_unit_progression", func(t *testing.T) {
		// Test that larger units are composed of smaller ones
		assert.Equal(t, xtime996.Microsecond*1000, xtime996.Millisecond)
		assert.Equal(t, xtime996.Millisecond*1000, xtime996.Second)
		assert.Equal(t, xtime996.Second*60, xtime996.Minute)
		assert.Equal(t, xtime996.Minute*60, xtime996.Hour)
		assert.Equal(t, xtime996.Hour*24, xtime996.Day)
	})

	t.Run("work_rest_relationships", func(t *testing.T) {
		// Work + rest = total for each period
		assert.Equal(t, xtime996.WorkDay+xtime996.RestDay, xtime996.Day)
		assert.Equal(t, xtime996.WorkWeek+xtime996.RestWeek, xtime996.Week)
		assert.Equal(t, xtime996.WorkMonth+xtime996.RestMonth, xtime996.Month)
		assert.Equal(t, xtime996.WorkQuarter+xtime996.RestQuarter, xtime996.Quarter)
		assert.Equal(t, xtime996.WorkYear+xtime996.RestYear, xtime996.Year)
	})

	t.Run("half_hour_relationship", func(t *testing.T) {
		// HalfHour should be exactly half of Hour
		assert.Equal(t, xtime996.Hour/2, xtime996.HalfHour)
		assert.Equal(t, xtime996.HalfHour*2, xtime996.Hour)
	})

	t.Run("work_week_composition", func(t *testing.T) {
		// Work week should be 6 work days * 12 hours each
		assert.Equal(t, xtime996.WorkDay*6, xtime996.WorkWeek)
		expectedWorkHours := time.Hour * 72 // 6 days * 12 hours
		assert.Equal(t, expectedWorkHours, xtime996.WorkWeek)
	})

	t.Run("equal_work_rest_day", func(t *testing.T) {
		// In 996 model, work day equals rest day (12 hours each)
		assert.Equal(t, xtime996.WorkDay, xtime996.RestDay)
		assert.Equal(t, time.Hour*12, xtime996.WorkDay)
		assert.Equal(t, time.Hour*12, xtime996.RestDay)
	})
}

func TestConstantValues(t *testing.T) {
	t.Run("positive_durations", func(t *testing.T) {
		// All constants should be positive
		assert.True(t, xtime996.Nanosecond > 0)
		assert.True(t, xtime996.Day > 0)
		assert.True(t, xtime996.WorkDay > 0)
		assert.True(t, xtime996.RestDay > 0)
		assert.True(t, xtime996.Week > 0)
		assert.True(t, xtime996.Month > 0)
		assert.True(t, xtime996.Year > 0)
	})

	t.Run("work_vs_rest_durations", func(t *testing.T) {
		// In 996 model, work and rest day should be equal
		assert.Equal(t, xtime996.WorkDay, xtime996.RestDay, "Work day should equal rest day")

		// In 996 model: WorkWeek = 6*12 = 72 hours, RestWeek = 168-72 = 96 hours
		// So actually RestWeek > WorkWeek
		assert.True(t, xtime996.RestWeek > xtime996.WorkWeek, "Rest week should be longer than work week")

		// Work month/quarter/year should be much longer than rest
		assert.True(t, xtime996.WorkMonth > xtime996.RestMonth, "Work month should be longer than rest month")
		assert.True(t, xtime996.WorkQuarter > xtime996.RestQuarter, "Work quarter should be longer than rest quarter")
		assert.True(t, xtime996.WorkYear > xtime996.RestYear, "Work year should be longer than rest year")
	})

	t.Run("specific_work_durations", func(t *testing.T) {
		// Test specific work duration values
		assert.Equal(t, time.Hour*12, xtime996.WorkDay)  // 12-hour work day
		assert.Equal(t, time.Hour*72, xtime996.WorkWeek) // 72-hour work week (6*12)

		// Month: 30 days - 4 rest days = 26 days
		expectedWorkMonth := xtime996.Day*30 - xtime996.RestDay*4
		assert.Equal(t, expectedWorkMonth, xtime996.WorkMonth)

		// Year: 365 days - 58 rest days = 307 days
		expectedWorkYear := xtime996.Year - xtime996.RestDay*58
		assert.Equal(t, expectedWorkYear, xtime996.WorkYear)
	})

	t.Run("specific_rest_durations", func(t *testing.T) {
		// Test specific rest duration values
		assert.Equal(t, time.Hour*12, xtime996.RestDay) // 12 hours rest per day

		// Weekly rest: 1 full day + 6 half days = 7 * 12 hours
		expectedRestWeek := xtime996.Day*1 + xtime996.RestDay*6
		assert.Equal(t, expectedRestWeek, xtime996.RestWeek)

		assert.Equal(t, xtime996.RestDay*4, xtime996.RestMonth)    // 4 rest days per month
		assert.Equal(t, xtime996.RestDay*14, xtime996.RestQuarter) // 14 rest days per quarter
		assert.Equal(t, xtime996.RestDay*58, xtime996.RestYear)    // 58 rest days per year
	})
}

func TestConstantMagnitudes(t *testing.T) {
	t.Run("reasonable_magnitudes", func(t *testing.T) {
		// Test that constants have reasonable magnitudes
		assert.True(t, xtime996.Minute > xtime996.Second)
		assert.True(t, xtime996.Hour > xtime996.Minute)
		assert.True(t, xtime996.Day > xtime996.Hour)
		assert.True(t, xtime996.Week > xtime996.Day)
		assert.True(t, xtime996.Month > xtime996.Week)
		assert.True(t, xtime996.Quarter > xtime996.Month)
		assert.True(t, xtime996.Year > xtime996.Quarter)
	})

	t.Run("work_vs_rest_comparisons", func(t *testing.T) {
		// Work day equals rest day in 996
		assert.Equal(t, xtime996.WorkDay, xtime996.RestDay)

		// Actually, rest week is longer than work week (96 vs 72 hours)
		assert.True(t, xtime996.RestWeek > xtime996.WorkWeek)

		// Work periods are longer than rest periods for larger time frames
		assert.True(t, xtime996.WorkMonth > xtime996.RestMonth)
		assert.True(t, xtime996.WorkQuarter > xtime996.RestQuarter)
		assert.True(t, xtime996.WorkYear > xtime996.RestYear)
	})

	t.Run("period_compositions", func(t *testing.T) {
		// Test larger period compositions
		assert.Equal(t, xtime996.Day*7, xtime996.Week)
		assert.Equal(t, xtime996.Day*30, xtime996.Month)
		assert.Equal(t, xtime996.Day*91, xtime996.Quarter)
		assert.Equal(t, xtime996.Day*365, xtime996.Year)
	})
}

func TestConstantUsageScenarios(t *testing.T) {
	t.Run("work_calculations", func(t *testing.T) {
		// Calculate work hours in different periods
		workHoursInDay := xtime996.WorkDay / xtime996.Hour
		assert.Equal(t, time.Duration(12), workHoursInDay) // 12-hour work day

		workHoursInWeek := xtime996.WorkWeek / xtime996.Hour
		assert.Equal(t, time.Duration(72), workHoursInWeek) // 72-hour work week

		// Work days in month = (30 days - 4 rest days) = 26 days
		expectedWorkDaysInMonth := (xtime996.Month - xtime996.RestMonth) / xtime996.Day
		actualWorkDaysInMonth := xtime996.WorkMonth / xtime996.Day
		assert.Equal(t, expectedWorkDaysInMonth, actualWorkDaysInMonth)

		// Work days in year = (365 days - 58 rest days) = 307 days
		expectedWorkDaysInYear := (xtime996.Year - xtime996.RestYear) / xtime996.Day
		actualWorkDaysInYear := xtime996.WorkYear / xtime996.Day
		assert.Equal(t, expectedWorkDaysInYear, actualWorkDaysInYear)
	})

	t.Run("rest_calculations", func(t *testing.T) {
		// Calculate rest hours in different periods
		restHoursInDay := xtime996.RestDay / xtime996.Hour
		assert.Equal(t, time.Duration(12), restHoursInDay) // 12 hours rest per day

		restDaysInMonth := xtime996.RestMonth / xtime996.RestDay
		assert.Equal(t, time.Duration(4), restDaysInMonth) // 4 rest days per month

		restDaysInQuarter := xtime996.RestQuarter / xtime996.RestDay
		assert.Equal(t, time.Duration(14), restDaysInQuarter) // 14 rest days per quarter

		restDaysInYear := xtime996.RestYear / xtime996.RestDay
		assert.Equal(t, time.Duration(58), restDaysInYear) // 58 rest days per year
	})

	t.Run("efficiency_calculations", func(t *testing.T) {
		// Calculate work efficiency ratios
		dailyWorkRatio := float64(xtime996.WorkDay) / float64(xtime996.Day)
		assert.InDelta(t, 0.5, dailyWorkRatio, 0.01) // 50% work time per day

		weeklyWorkRatio := float64(xtime996.WorkWeek) / float64(xtime996.Week)
		expectedWeeklyRatio := 72.0 / 168.0 // 72 hours / 168 hours per week
		assert.InDelta(t, expectedWeeklyRatio, weeklyWorkRatio, 0.01)

		// Monthly work ratio
		monthlyWorkRatio := float64(xtime996.WorkMonth) / float64(xtime996.Month)
		expectedMonthlyRatio := (30.0*24.0 - 4.0*12.0) / (30.0 * 24.0) // (Total - Rest) / Total
		assert.InDelta(t, expectedMonthlyRatio, monthlyWorkRatio, 0.01)
	})

	t.Run("duration_arithmetic", func(t *testing.T) {
		// Test arithmetic operations with constants
		overtimeHours := time.Hour * 2
		extendedWorkDay := xtime996.WorkDay + overtimeHours
		assert.Equal(t, time.Hour*14, extendedWorkDay) // 12 + 2 = 14 hours

		weekendRest := xtime996.Day * 1                 // One day rest per week
		assert.True(t, weekendRest < xtime996.RestWeek) // Weekend is part of rest week

		// Double work week calculation
		doubleWorkWeek := xtime996.WorkWeek * 2
		assert.Equal(t, time.Hour*144, doubleWorkWeek) // 72 * 2 = 144 hours
	})

	t.Run("burnout_risk_calculations", func(t *testing.T) {
		// 996 is known for high work intensity
		workToRestRatio := float64(xtime996.WorkWeek) / float64(xtime996.RestWeek)

		// Work week (72h) vs Rest week (1*24h + 6*12h = 96h)
		expectedRatio := 72.0 / 96.0 // 0.75
		assert.InDelta(t, expectedRatio, workToRestRatio, 0.01)

		// Monthly work intensity
		monthlyWorkHours := xtime996.WorkMonth / xtime996.Hour
		monthlyTotalHours := xtime996.Month / xtime996.Hour
		monthlyWorkIntensity := float64(monthlyWorkHours) / float64(monthlyTotalHours)

		expectedMonthlyIntensity := (30.0*24.0 - 4.0*12.0) / (30.0 * 24.0)
		assert.InDelta(t, expectedMonthlyIntensity, monthlyWorkIntensity, 0.01)
	})
}

func TestConstantTypes(t *testing.T) {
	t.Run("duration_types", func(t *testing.T) {
		// All constants should be of type time.Duration
		assert.IsType(t, time.Duration(0), xtime996.Nanosecond)
		assert.IsType(t, time.Duration(0), xtime996.Day)
		assert.IsType(t, time.Duration(0), xtime996.WorkDay)
		assert.IsType(t, time.Duration(0), xtime996.RestDay)
		assert.IsType(t, time.Duration(0), xtime996.WorkWeek)
		assert.IsType(t, time.Duration(0), xtime996.RestWeek)
		assert.IsType(t, time.Duration(0), xtime996.WorkMonth)
		assert.IsType(t, time.Duration(0), xtime996.RestMonth)
		assert.IsType(t, time.Duration(0), xtime996.WorkQuarter)
		assert.IsType(t, time.Duration(0), xtime996.RestQuarter)
		assert.IsType(t, time.Duration(0), xtime996.WorkYear)
		assert.IsType(t, time.Duration(0), xtime996.RestYear)
	})
}

func TestWorkLifeBalance996(t *testing.T) {
	t.Run("daily_balance", func(t *testing.T) {
		// 12 hours work, 12 hours personal time (including sleep)
		totalDailyTime := xtime996.WorkDay + xtime996.RestDay
		assert.Equal(t, xtime996.Day, totalDailyTime)

		workPercentage := float64(xtime996.WorkDay) * 100 / float64(xtime996.Day)
		assert.InDelta(t, 50.0, workPercentage, 0.1) // 50% work time
	})

	t.Run("weekly_balance", func(t *testing.T) {
		// 72 hours work, 96 hours personal time
		totalWeeklyTime := xtime996.WorkWeek + xtime996.RestWeek
		assert.Equal(t, xtime996.Week, totalWeeklyTime)

		workPercentage := float64(xtime996.WorkWeek) * 100 / float64(xtime996.Week)
		expectedPercentage := 72.0 * 100 / 168.0 // 72 hours out of 168
		assert.InDelta(t, expectedPercentage, workPercentage, 0.1)
	})

	t.Run("annual_work_load", func(t *testing.T) {
		// Calculate annual work hours
		annualWorkHours := xtime996.WorkYear / xtime996.Hour

		// 307 work days * 24 hours per day = 7368 hours per year
		// But WorkYear = Year - RestYear = 8760 - 696 = 8064 hours
		expectedAnnualHours := 365*24 - 58*12 // Full year hours - rest day hours
		assert.Equal(t, time.Duration(expectedAnnualHours), annualWorkHours)

		// Compare to standard work year (40 hours * 50 weeks = 2000 hours)
		standardWorkYear := time.Hour * 2000
		intensityRatio := float64(xtime996.WorkYear) / float64(standardWorkYear)
		assert.True(t, intensityRatio > 1.8, "996 should be significantly more intense than standard work")
	})

	t.Run("rest_adequacy", func(t *testing.T) {
		// Check if rest periods are adequate
		dailyRestHours := xtime996.RestDay / xtime996.Hour
		assert.Equal(t, time.Duration(12), dailyRestHours)

		// Assuming 8 hours sleep, only 4 hours for personal activities
		sleepTime := time.Hour * 8
		personalTime := xtime996.RestDay - sleepTime
		personalHours := personalTime / xtime996.Hour
		assert.Equal(t, time.Duration(4), personalHours)

		// Weekly rest includes one full day off
		weeklyRestDays := xtime996.RestWeek / xtime996.Day
		// 1 full day + 6 half days = 4 equivalent full rest days
		expectedRestDays := 1.0 + 6.0*0.5
		assert.InDelta(t, expectedRestDays, float64(weeklyRestDays), 0.1)
	})
}

// Benchmark constant access
func BenchmarkConstants996(b *testing.B) {
	b.Run("basic_constants", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = xtime996.Second
			_ = xtime996.Minute
			_ = xtime996.Hour
			_ = xtime996.Day
		}
	})

	b.Run("work_constants", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = xtime996.WorkDay
			_ = xtime996.WorkWeek
			_ = xtime996.WorkMonth
			_ = xtime996.WorkYear
		}
	})

	b.Run("rest_constants", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = xtime996.RestDay
			_ = xtime996.RestWeek
			_ = xtime996.RestMonth
			_ = xtime996.RestYear
		}
	})

	b.Run("intensity_calculations", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = float64(xtime996.WorkYear) / float64(xtime996.Year)
			_ = xtime996.WorkWeek / xtime996.Hour
			_ = xtime996.WorkMonth / xtime996.Day
		}
	})
}
