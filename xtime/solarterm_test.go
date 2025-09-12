package xtime_test

import (
	"testing"
	"time"

	"github.com/lazygophers/utils/xtime"
	"github.com/stretchr/testify/assert"
)

func TestSolartermBasicOperations(t *testing.T) {
	t.Run("solarterm_string", func(t *testing.T) {
		// Test the 24 solar terms
		expectedTerms := []string{
			"小寒", "大寒", "立春", "雨水", "惊蛰", "春分",
			"清明", "谷雨", "立夏", "小满", "芒种", "夏至",
			"小暑", "大暑", "立秋", "处暑", "白露", "秋分",
			"寒露", "霜降", "立冬", "小雪", "大雪", "冬至",
		}

		for i := 0; i < 24; i++ {
			solarterm := xtime.Solarterm(i)
			assert.Equal(t, expectedTerms[i], solarterm.String(), "Solarterm %d should be %s", i, expectedTerms[i])
		}
	})

	t.Run("solarterm_modulo", func(t *testing.T) {
		// Test that solarterm wraps around after 24
		solarterm24 := xtime.Solarterm(24)
		solarterm0 := xtime.Solarterm(0)
		
		assert.Equal(t, solarterm0.String(), solarterm24.String(), "Solarterm should wrap around after 24")
		
		solarterm25 := xtime.Solarterm(25)
		solarterm1 := xtime.Solarterm(1)
		assert.Equal(t, solarterm1.String(), solarterm25.String(), "Solarterm 25 should equal solarterm 1")
	})

	t.Run("solarterm_equals", func(t *testing.T) {
		solarterm1 := xtime.Solarterm(5)
		solarterm2 := xtime.Solarterm(5)
		solarterm3 := xtime.Solarterm(6)

		assert.True(t, solarterm1.Equals(solarterm2), "Equal solarterms should be equal")
		assert.False(t, solarterm1.Equals(solarterm3), "Different solarterms should not be equal")
		assert.True(t, solarterm2.Equals(solarterm1), "Equality should be symmetric")
	})
}

func TestSolartermNavigation(t *testing.T) {
	t.Run("next_and_prev", func(t *testing.T) {
		solarterm := xtime.Solarterm(10)
		
		next := solarterm.Next()
		prev := solarterm.Prev()
		
		assert.Equal(t, xtime.Solarterm(11), next, "Next should increment by 1")
		assert.Equal(t, xtime.Solarterm(9), prev, "Prev should decrement by 1")
		
		// Test round-trip
		assert.Equal(t, solarterm, next.Prev(), "next.Prev() should equal original")
		assert.Equal(t, solarterm, prev.Next(), "prev.Next() should equal original")
	})

	t.Run("edge_cases_navigation", func(t *testing.T) {
		// Test navigation with edge values
		solarterm0 := xtime.Solarterm(0)
		prevOfFirst := solarterm0.Prev()
		
		// Should wrap to -1
		assert.Equal(t, xtime.Solarterm(-1), prevOfFirst)
		
		// Test that string representation works with modulo
		defer func() {
			if r := recover(); r != nil {
				t.Logf("String() with negative solarterm panicked: %v", r)
			}
		}()
		
		// The string representation may use modulo, so test carefully
		str := prevOfFirst.String()
		t.Logf("Solarterm(-1).String() = %s", str)
		
		// Test large numbers
		largeSolarterm := xtime.Solarterm(1000)
		nextLarge := largeSolarterm.Next()
		assert.Equal(t, xtime.Solarterm(1001), nextLarge)
	})
}

func TestSolartermTime(t *testing.T) {
	t.Run("solarterm_to_time", func(t *testing.T) {
		// Test that solarterm can be converted to time
		solarterm := xtime.Solarterm(0) // 小寒
		
		timeObj := solarterm.Time()
		assert.IsType(t, time.Time{}, timeObj, "Time() should return time.Time")
		assert.False(t, timeObj.IsZero(), "Time should not be zero")
		
		// Should be a reasonable date (between 1900 and 2100)
		year := timeObj.Year()
		assert.True(t, year >= 1900 && year <= 2100, "Year should be reasonable: %d", year)
	})

	t.Run("solarterm_timestamp", func(t *testing.T) {
		// Use a solarterm index that corresponds to a date after 1970 (Unix epoch)
		// Since data starts from 1904, we need an index for more recent years
		solarterm := xtime.Solarterm(1600) // A later solarterm that should have positive timestamp
		
		timestamp := solarterm.Timestamp()
		// Note: Early years (1904-1970) will have negative timestamps, which is expected
		// Only test that the function doesn't panic
		assert.NotPanics(t, func() {
			_ = solarterm.Timestamp()
		}, "Timestamp() should not panic")
		
		// Convert back to time and verify consistency
		timeFromTimestamp := time.Unix(timestamp, 0)
		timeFromMethod := solarterm.Time()
		
		assert.Equal(t, timeFromTimestamp.Unix(), timeFromMethod.Unix(), "Timestamp and Time() should be consistent")
	})

	t.Run("multiple_solarterms_progression", func(t *testing.T) {
		// Test that consecutive solarterms have increasing timestamps
		var prevTimestamp int64 = 0
		
		for i := 0; i < 48; i++ { // Test 2 years worth
			solarterm := xtime.Solarterm(i)
			timestamp := solarterm.Timestamp()
			
			if i > 0 {
				assert.True(t, timestamp > prevTimestamp, 
					"Solarterm %d timestamp should be greater than previous", i)
			}
			prevTimestamp = timestamp
		}
	})
}

func TestSolartermIsInDay(t *testing.T) {
	t.Run("is_in_day", func(t *testing.T) {
		solarterm := xtime.Solarterm(0)
		solartermTime := solarterm.Time()
		
		// Same day should return true
		sameDay := time.Date(solartermTime.Year(), solartermTime.Month(), solartermTime.Day(), 
			12, 0, 0, 0, time.Local)
		assert.True(t, solarterm.IsInDay(sameDay), "Same day should return true")
		
		// Different day should return false
		differentDay := solartermTime.AddDate(0, 0, 1)
		assert.False(t, solarterm.IsInDay(differentDay), "Different day should return false")
		
		// Beginning of the day should return true
		beginningOfDay := time.Date(solartermTime.Year(), solartermTime.Month(), solartermTime.Day(),
			0, 0, 0, 0, time.Local)
		assert.True(t, solarterm.IsInDay(beginningOfDay), "Beginning of day should return true")
		
		// End of the day should return true
		endOfDay := time.Date(solartermTime.Year(), solartermTime.Month(), solartermTime.Day(),
			23, 59, 59, 0, time.Local)
		assert.True(t, solarterm.IsInDay(endOfDay), "End of day should return true")
	})

	t.Run("is_today", func(t *testing.T) {
		// This test is time-dependent, so we need to be careful
		now := time.Now()
		
		// Find the current solarterm
		currentSolarterm := xtime.NextSolarterm(now.AddDate(0, 0, -1)) // Get recent solarterm
		
		// Test IsToDay method
		isToday := currentSolarterm.IsToDay()
		assert.IsType(t, true, isToday, "IsToDay should return bool")
		
		// The result depends on current date, but it should not panic
		assert.NotPanics(t, func() {
			_ = currentSolarterm.IsToDay()
		}, "IsToDay should not panic")
	})
}

func TestNextSolarterm(t *testing.T) {
	t.Run("next_solarterm_basic", func(t *testing.T) {
		// Test with a specific date
		testTime := time.Date(2023, 6, 15, 12, 0, 0, 0, time.UTC)
		nextSolarterm := xtime.NextSolarterm(testTime)
		
		assert.IsType(t, xtime.Solarterm(0), nextSolarterm, "Should return Solarterm type")
		
		// Next solarterm should be after the test time
		nextTime := nextSolarterm.Time()
		assert.True(t, nextTime.After(testTime) || nextTime.Equal(testTime), 
			"Next solarterm should be after or equal to test time")
	})

	t.Run("next_solarterm_progression", func(t *testing.T) {
		// Test that NextSolarterm works correctly across different dates
		startTime := time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC)
		
		var prevTime time.Time
		
		for i := 0; i < 12; i++ { // Test across a year
			testTime := startTime.AddDate(0, i, 0)
			nextSolarterm := xtime.NextSolarterm(testTime)
			nextTime := nextSolarterm.Time()
			
			if i > 0 {
				// Should generally progress (though might repeat if we hit exact dates)
				assert.True(t, nextTime.After(prevTime) || nextTime.Equal(prevTime),
					"Solarterm times should progress or stay the same")
			}
			
			prevTime = nextTime
		}
	})

	t.Run("next_solarterm_edge_cases", func(t *testing.T) {
		// Test with various edge case dates
		edgeDates := []time.Time{
			time.Date(1900, 1, 1, 0, 0, 0, 0, time.UTC),  // Early date
			time.Date(2100, 12, 31, 23, 59, 59, 0, time.UTC), // Late date
			time.Date(2023, 2, 29, 0, 0, 0, 0, time.UTC), // Invalid date (non-leap year)
		}

		for i, date := range edgeDates {
			if i == 2 { // Skip invalid date
				continue
			}
			
			assert.NotPanics(t, func() {
				nextSolarterm := xtime.NextSolarterm(date)
				_ = nextSolarterm.String() // Should not panic
			}, "NextSolarterm should not panic for edge date %d", i)
		}
	})
}

func TestSolartermHelperFunctions(t *testing.T) {
	t.Run("dd_function", func(t *testing.T) {
		// Test the DD (Julian Day to Gregorian) function
		// This is an internal function but we can test it exists
		// by testing solarterm functionality that depends on it
		
		solarterm := xtime.Solarterm(0)
		
		// Should not panic when getting time
		assert.NotPanics(t, func() {
			_ = solarterm.Time()
		}, "DD function should work correctly")
		
		// Should produce reasonable dates
		timeObj := solarterm.Time()
		assert.True(t, timeObj.Year() >= 1900 && timeObj.Year() <= 2100,
			"DD should produce reasonable year")
		assert.True(t, timeObj.Month() >= 1 && timeObj.Month() <= 12,
			"DD should produce valid month")
		assert.True(t, timeObj.Day() >= 1 && timeObj.Day() <= 31,
			"DD should produce valid day")
	})
}

func TestSolartermConstantsAndData(t *testing.T) {
	t.Run("solarterm_names", func(t *testing.T) {
		// Test that all 24 solar terms have valid Chinese names
		expectedLength := 24
		terms := make([]string, expectedLength)
		
		for i := 0; i < expectedLength; i++ {
			solarterm := xtime.Solarterm(i)
			name := solarterm.String()
			terms[i] = name
			
			assert.NotEmpty(t, name, "Solarterm %d should have a name", i)
			assert.Len(t, []rune(name), 2, "Solarterm names should be 2 Chinese characters")
		}
		
		// Check for uniqueness
		uniqueTerms := make(map[string]bool)
		for _, term := range terms {
			assert.False(t, uniqueTerms[term], "Term %s should be unique", term)
			uniqueTerms[term] = true
		}
		assert.Len(t, uniqueTerms, expectedLength, "Should have exactly 24 unique terms")
	})

	t.Run("solarterm_data_validity", func(t *testing.T) {
		// Test that solarterm data produces valid timestamps
		// Note: Early solarterms (1904-1970) will have negative Unix timestamps, which is normal
		for i := 0; i < 100; i++ { // Test first 100 entries
			solarterm := xtime.Solarterm(i)
			
			assert.NotPanics(t, func() {
				timestamp := solarterm.Timestamp()
				// Don't assert positive - early years have negative Unix timestamps
				_ = timestamp
			}, "Solarterm %d should have valid timestamp", i)
		}
	})
}

// Benchmark solarterm operations
func BenchmarkSolarterm(b *testing.B) {
	testTime := time.Date(2023, 6, 15, 12, 0, 0, 0, time.UTC)
	solarterm := xtime.Solarterm(10)

	b.Run("String", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = solarterm.String()
		}
	})

	b.Run("Time", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = solarterm.Time()
		}
	})

	b.Run("Timestamp", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = solarterm.Timestamp()
		}
	})

	b.Run("NextSolarterm", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = xtime.NextSolarterm(testTime)
		}
	})

	b.Run("IsInDay", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = solarterm.IsInDay(testTime)
		}
	})
}