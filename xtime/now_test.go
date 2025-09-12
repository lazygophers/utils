package xtime_test

import (
	"testing"
	"time"

	"github.com/lazygophers/utils/xtime"
	"github.com/stretchr/testify/assert"
)

func TestNow(t *testing.T) {
	t.Run("now_returns_current_time", func(t *testing.T) {
		before := time.Now()
		now := xtime.Now()
		after := time.Now()

		// Now should return current time wrapped in xtime.Time
		assert.NotNil(t, now)
		assert.NotNil(t, now.Config)
		
		// Time should be between before and after
		assert.True(t, now.Time.After(before.Add(-time.Second)), "Now should be after before time")
		assert.True(t, now.Time.Before(after.Add(time.Second)), "Now should be before after time")
	})

	t.Run("now_config_defaults", func(t *testing.T) {
		now := xtime.Now()
		
		assert.Equal(t, time.Monday, now.Config.WeekStartDay)
		assert.Equal(t, time.Local, now.Config.TimeLocation)
		assert.Empty(t, now.Config.TimeFormats)
	})
}

func TestNowUnix(t *testing.T) {
	t.Run("unix_timestamp", func(t *testing.T) {
		before := time.Now().Unix()
		unixTime := xtime.NowUnix()
		after := time.Now().Unix()

		assert.True(t, unixTime >= before, "Unix time should be >= before time")
		assert.True(t, unixTime <= after, "Unix time should be <= after time")
		assert.True(t, unixTime > 0, "Unix timestamp should be positive")
	})
}

func TestNowUnixMilli(t *testing.T) {
	t.Run("unix_milli_timestamp", func(t *testing.T) {
		before := time.Now().UnixMilli()
		unixMilli := xtime.NowUnixMilli()
		after := time.Now().UnixMilli()

		assert.True(t, unixMilli >= before, "Unix milli should be >= before time")
		assert.True(t, unixMilli <= after, "Unix milli should be <= after time")
		assert.True(t, unixMilli > 0, "Unix milli timestamp should be positive")
		assert.True(t, unixMilli > 1000000000000, "Unix milli should be reasonable size")
	})
}

func TestBeginningMethods(t *testing.T) {
	// Use a specific test time for consistent results
	testTime := time.Date(2023, 6, 15, 14, 30, 45, 123456789, time.UTC)
	xt := xtime.With(testTime)

	t.Run("beginning_of_minute", func(t *testing.T) {
		result := xt.BeginningOfMinute()
		expected := time.Date(2023, 6, 15, 14, 30, 0, 0, time.UTC)
		
		assert.Equal(t, expected, result.Time)
		assert.Equal(t, 0, result.Second())
		assert.Equal(t, 0, result.Nanosecond())
	})

	t.Run("beginning_of_hour", func(t *testing.T) {
		result := xt.BeginningOfHour()
		expected := time.Date(2023, 6, 15, 14, 0, 0, 0, time.UTC)
		
		assert.Equal(t, expected, result.Time)
		assert.Equal(t, 0, result.Minute())
		assert.Equal(t, 0, result.Second())
		assert.Equal(t, 0, result.Nanosecond())
	})

	t.Run("beginning_of_day", func(t *testing.T) {
		result := xt.BeginningOfDay()
		expected := time.Date(2023, 6, 15, 0, 0, 0, 0, time.UTC)
		
		assert.Equal(t, expected, result.Time)
		assert.Equal(t, 0, result.Hour())
		assert.Equal(t, 0, result.Minute())
		assert.Equal(t, 0, result.Second())
	})

	t.Run("beginning_of_week_monday", func(t *testing.T) {
		// 2023-06-15 was a Thursday
		result := xt.BeginningOfWeek() // Default Monday start
		
		// Should go back to Monday 2023-06-12
		assert.Equal(t, time.Monday, result.Weekday())
		assert.Equal(t, 12, result.Day())
		assert.Equal(t, 0, result.Hour())
		assert.Equal(t, 0, result.Minute())
	})

	t.Run("beginning_of_week_sunday", func(t *testing.T) {
		// Test with Sunday as week start
		xt.Config.WeekStartDay = time.Sunday
		result := xt.BeginningOfWeek()
		
		// Should go back to Sunday 2023-06-11
		assert.Equal(t, time.Sunday, result.Weekday())
		assert.Equal(t, 11, result.Day())
		assert.Equal(t, 0, result.Hour())
	})

	t.Run("beginning_of_week_edge_cases", func(t *testing.T) {
		// Test edge case where weekday < weekStartDayInt
		// Use a Saturday (weekday=6) with Wednesday start (weekday=3)
		saturday := time.Date(2023, 6, 17, 14, 30, 0, 0, time.UTC) // Saturday
		satWrapped := xtime.With(saturday)
		satWrapped.Config.WeekStartDay = time.Wednesday // Wednesday = 3
		
		result := satWrapped.BeginningOfWeek()
		// Saturday(6) < Wednesday(3) is false, so should go back to Wednesday 2023-06-14
		assert.Equal(t, time.Wednesday, result.Weekday())
		assert.Equal(t, 14, result.Day())
		
		// Test edge case where weekday < weekStartDayInt is true
		// Use a Sunday (weekday=0) with Tuesday start (weekday=2) 
		sunday := time.Date(2023, 6, 18, 14, 30, 0, 0, time.UTC) // Sunday
		sunWrapped := xtime.With(sunday)
		sunWrapped.Config.WeekStartDay = time.Tuesday // Tuesday = 2
		
		result2 := sunWrapped.BeginningOfWeek()
		// Sunday(0) < Tuesday(2) is true, should use: weekday + 7 - weekStartDayInt = 0 + 7 - 2 = 5
		// This means go back 5 days from Sunday to Tuesday
		assert.Equal(t, time.Tuesday, result2.Weekday()) 
		assert.Equal(t, 13, result2.Day()) // June 13, 2023 was a Tuesday
	})

	t.Run("beginning_of_month", func(t *testing.T) {
		result := xt.BeginningOfMonth()
		expected := time.Date(2023, 6, 1, 0, 0, 0, 0, time.UTC)
		
		assert.Equal(t, expected, result.Time)
		assert.Equal(t, 1, result.Day())
		assert.Equal(t, 0, result.Hour())
	})

	t.Run("beginning_of_quarter", func(t *testing.T) {
		result := xt.BeginningOfQuarter()
		// June is in Q2, so should go to April 1st
		expected := time.Date(2023, 4, 1, 0, 0, 0, 0, time.UTC)
		
		assert.Equal(t, expected, result.Time)
		assert.Equal(t, time.April, result.Month())
		assert.Equal(t, 1, result.Day())
	})

	t.Run("beginning_of_half", func(t *testing.T) {
		result := xt.BeginningOfHalf()
		// June is in second half, so should go to January 1st
		expected := time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC)
		
		assert.Equal(t, expected, result.Time)
		assert.Equal(t, time.January, result.Month())
		assert.Equal(t, 1, result.Day())
	})

	t.Run("beginning_of_year", func(t *testing.T) {
		result := xt.BeginningOfYear()
		expected := time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC)
		
		assert.Equal(t, expected, result.Time)
		assert.Equal(t, time.January, result.Month())
		assert.Equal(t, 1, result.Day())
	})
}

func TestEndMethods(t *testing.T) {
	testTime := time.Date(2023, 6, 15, 14, 30, 45, 123456789, time.UTC)
	xt := xtime.With(testTime)

	t.Run("end_of_minute", func(t *testing.T) {
		result := xt.EndOfMinute()
		
		assert.Equal(t, 14, result.Hour())
		assert.Equal(t, 30, result.Minute())
		assert.Equal(t, 59, result.Second())
		assert.Equal(t, int(time.Second-time.Nanosecond), result.Nanosecond())
	})

	t.Run("end_of_hour", func(t *testing.T) {
		result := xt.EndOfHour()
		
		assert.Equal(t, 14, result.Hour())
		assert.Equal(t, 59, result.Minute())
		assert.Equal(t, 59, result.Second())
		assert.Equal(t, int(time.Second-time.Nanosecond), result.Nanosecond())
	})

	t.Run("end_of_day", func(t *testing.T) {
		result := xt.EndOfDay()
		
		assert.Equal(t, 15, result.Day())
		assert.Equal(t, 23, result.Hour())
		assert.Equal(t, 59, result.Minute())
		assert.Equal(t, 59, result.Second())
		assert.Equal(t, int(time.Second-time.Nanosecond), result.Nanosecond())
	})

	t.Run("end_of_week", func(t *testing.T) {
		result := xt.EndOfWeek()
		
		// Should be end of Sunday (next week's Sunday minus 1 nanosecond)
		assert.Equal(t, time.Sunday, result.Weekday())
		assert.Equal(t, 18, result.Day()) // 2023-06-18
		assert.Equal(t, 23, result.Hour())
		assert.Equal(t, 59, result.Minute())
		assert.Equal(t, 59, result.Second())
	})

	t.Run("end_of_month", func(t *testing.T) {
		result := xt.EndOfMonth()
		
		assert.Equal(t, time.June, result.Month())
		assert.Equal(t, 30, result.Day()) // June has 30 days
		assert.Equal(t, 23, result.Hour())
		assert.Equal(t, 59, result.Minute())
		assert.Equal(t, 59, result.Second())
	})

	t.Run("end_of_quarter", func(t *testing.T) {
		result := xt.EndOfQuarter()
		
		// Q2 ends in June
		assert.Equal(t, time.June, result.Month())
		assert.Equal(t, 30, result.Day())
		assert.Equal(t, 23, result.Hour())
		assert.Equal(t, 59, result.Minute())
	})

	t.Run("end_of_half", func(t *testing.T) {
		result := xt.EndOfHalf()
		
		// First half ends in June
		assert.Equal(t, time.June, result.Month())
		assert.Equal(t, 30, result.Day())
		assert.Equal(t, 23, result.Hour())
	})

	t.Run("end_of_year", func(t *testing.T) {
		result := xt.EndOfYear()
		
		assert.Equal(t, 2023, result.Year())
		assert.Equal(t, time.December, result.Month())
		assert.Equal(t, 31, result.Day())
		assert.Equal(t, 23, result.Hour())
		assert.Equal(t, 59, result.Minute())
	})
}

func TestQuarter(t *testing.T) {
	testCases := []struct {
		month    time.Month
		expected uint
	}{
		{time.January, 1},
		{time.February, 1},
		{time.March, 1},
		{time.April, 2},
		{time.May, 2},
		{time.June, 2},
		{time.July, 3},
		{time.August, 3},
		{time.September, 3},
		{time.October, 4},
		{time.November, 4},
		{time.December, 4},
	}

	for _, tc := range testCases {
		t.Run(tc.month.String(), func(t *testing.T) {
			testTime := time.Date(2023, tc.month, 15, 12, 0, 0, 0, time.UTC)
			xt := xtime.With(testTime)
			
			result := xt.Quarter()
			assert.Equal(t, tc.expected, result, "Month %s should be in quarter %d", tc.month, tc.expected)
		})
	}
}

// Test package-level convenience functions
func TestPackageLevelFunctions(t *testing.T) {
	t.Run("beginning_functions", func(t *testing.T) {
		// Test that package level functions work
		beginningOfMinute := xtime.BeginningOfMinute()
		beginningOfHour := xtime.BeginningOfHour()
		beginningOfDay := xtime.BeginningOfDay()
		beginningOfWeek := xtime.BeginningOfWeek()
		beginningOfMonth := xtime.BeginningOfMonth()
		beginningOfQuarter := xtime.BeginningOfQuarter()
		beginningOfYear := xtime.BeginningOfYear()

		// All should return xtime.Time instances
		assert.NotNil(t, beginningOfMinute)
		assert.NotNil(t, beginningOfHour)
		assert.NotNil(t, beginningOfDay)
		assert.NotNil(t, beginningOfWeek)
		assert.NotNil(t, beginningOfMonth)
		assert.NotNil(t, beginningOfQuarter)
		assert.NotNil(t, beginningOfYear)
		
		// Check that they're actually at the beginning of their respective periods
		assert.Equal(t, 0, beginningOfMinute.Second())
		assert.Equal(t, 0, beginningOfHour.Minute())
		assert.Equal(t, 0, beginningOfDay.Hour())
		assert.Equal(t, 1, beginningOfMonth.Day())
		assert.Equal(t, time.January, beginningOfYear.Month())
	})

	t.Run("end_functions", func(t *testing.T) {
		endOfMinute := xtime.EndOfMinute()
		endOfHour := xtime.EndOfHour()
		endOfDay := xtime.EndOfDay()
		endOfWeek := xtime.EndOfWeek()
		endOfMonth := xtime.EndOfMonth()
		endOfQuarter := xtime.EndOfQuarter()
		endOfYear := xtime.EndOfYear()

		// All should return xtime.Time instances
		assert.NotNil(t, endOfMinute)
		assert.NotNil(t, endOfHour)
		assert.NotNil(t, endOfDay)
		assert.NotNil(t, endOfWeek)
		assert.NotNil(t, endOfMonth)
		assert.NotNil(t, endOfQuarter)
		assert.NotNil(t, endOfYear)

		// Check that they're at the end of their periods
		assert.Equal(t, 59, endOfMinute.Second())
		assert.Equal(t, 59, endOfHour.Minute())
		assert.Equal(t, 23, endOfDay.Hour())
		assert.Equal(t, time.December, endOfYear.Month())
	})

	t.Run("quarter_function", func(t *testing.T) {
		quarter := xtime.Quarter()
		assert.True(t, quarter >= 1 && quarter <= 4, "Quarter should be between 1 and 4")
	})
}

func TestEdgeCases(t *testing.T) {
	t.Run("leap_year_february", func(t *testing.T) {
		// Test February in a leap year
		leapTime := time.Date(2024, 2, 15, 12, 0, 0, 0, time.UTC)
		xt := xtime.With(leapTime)
		
		endOfMonth := xt.EndOfMonth()
		assert.Equal(t, 29, endOfMonth.Day()) // Feb 29 in leap year
	})

	t.Run("year_boundary", func(t *testing.T) {
		// Test December 31st
		yearEnd := time.Date(2023, 12, 31, 23, 59, 59, 0, time.UTC)
		xt := xtime.With(yearEnd)
		
		beginningOfYear := xt.BeginningOfYear()
		assert.Equal(t, 2023, beginningOfYear.Year())
		assert.Equal(t, time.January, beginningOfYear.Month())
		assert.Equal(t, 1, beginningOfYear.Day())
	})

	t.Run("different_timezones", func(t *testing.T) {
		// Test with different timezone
		est, _ := time.LoadLocation("America/New_York")
		testTime := time.Date(2023, 6, 15, 14, 30, 0, 0, est)
		xt := xtime.With(testTime)
		
		beginningOfDay := xt.BeginningOfDay()
		assert.Equal(t, est, beginningOfDay.Location())
		assert.Equal(t, 0, beginningOfDay.Hour())
	})
}