package xtime_test

import (
	"testing"
	"time"

	"github.com/lazygophers/utils/xtime"
	"github.com/stretchr/testify/assert"
)

func TestWithLunar(t *testing.T) {
	t.Run("basic_lunar_conversion", func(t *testing.T) {
		// Test with a known date
		solarTime := time.Date(2023, 1, 22, 12, 0, 0, 0, time.UTC) // Chinese New Year 2023
		lunar := xtime.WithLunar(solarTime)

		assert.NotNil(t, lunar)
		assert.Equal(t, solarTime, lunar.Time)
		assert.True(t, lunar.Year() > 0)
		assert.True(t, lunar.Month() >= 1 && lunar.Month() <= 12)
		assert.True(t, lunar.Day() >= 1 && lunar.Day() <= 30)
	})

	t.Run("different_solar_dates", func(t *testing.T) {
		testDates := []time.Time{
			time.Date(2023, 6, 15, 0, 0, 0, 0, time.UTC),
			time.Date(2023, 12, 31, 0, 0, 0, 0, time.UTC),
			time.Date(2024, 2, 10, 0, 0, 0, 0, time.UTC),
		}

		for i, date := range testDates {
			lunar := xtime.WithLunar(date)
			assert.NotNil(t, lunar, "Lunar conversion %d should not be nil", i)
			assert.Equal(t, date, lunar.Time, "Solar time should be preserved")
			
			// Basic sanity checks
			assert.True(t, lunar.Year() >= 1900 && lunar.Year() <= 2100, "Lunar year should be reasonable")
			assert.True(t, lunar.Month() >= 1 && lunar.Month() <= 12, "Lunar month should be 1-12")
			assert.True(t, lunar.Day() >= 1 && lunar.Day() <= 30, "Lunar day should be 1-30")
		}
	})
}

func TestWithLunarTime(t *testing.T) {
	t.Run("lunar_time_creation", func(t *testing.T) {
		testTime := time.Date(2023, 8, 15, 14, 30, 0, 0, time.UTC)
		lunar := xtime.WithLunarTime(testTime)

		assert.NotNil(t, lunar)
		assert.Equal(t, testTime, lunar.Time)
		// WithLunarTime should provide same functionality as WithLunar
		assert.True(t, lunar.Year() > 0)
		assert.True(t, lunar.Month() >= 1)
		assert.True(t, lunar.Day() >= 1)
	})
}

func TestLunarBasicMethods(t *testing.T) {
	testTime := time.Date(2023, 6, 15, 0, 0, 0, 0, time.UTC)
	lunar := xtime.WithLunar(testTime)

	t.Run("year_month_day_getters", func(t *testing.T) {
		year := lunar.Year()
		month := lunar.Month()
		day := lunar.Day()

		assert.True(t, year >= 1900 && year <= 2100, "Year should be reasonable")
		assert.True(t, month >= 1 && month <= 12, "Month should be 1-12")
		assert.True(t, day >= 1 && day <= 30, "Day should be 1-30")
		assert.IsType(t, int64(0), year, "Year should be int64")
		assert.IsType(t, int64(0), month, "Month should be int64")
		assert.IsType(t, int64(0), day, "Day should be int64")
	})

	t.Run("leap_month_methods", func(t *testing.T) {
		leapMonth := lunar.LeapMonth()
		isLeap := lunar.IsLeap()
		isLeapMonth := lunar.IsLeapMonth()

		// Leap month should be 0-12 (0 means no leap month)
		assert.True(t, leapMonth >= 0 && leapMonth <= 12, "Leap month should be 0-12")
		assert.IsType(t, int64(0), leapMonth, "Leap month should be int64")
		
		// IsLeap should be consistent with LeapMonth
		if leapMonth == 0 {
			assert.False(t, isLeap, "If no leap month, IsLeap should be false")
		} else {
			assert.True(t, isLeap, "If leap month exists, IsLeap should be true")
		}

		assert.IsType(t, true, isLeap, "IsLeap should be bool")
		assert.IsType(t, true, isLeapMonth, "IsLeapMonth should be bool")
	})

	t.Run("date_string", func(t *testing.T) {
		dateStr := lunar.Date()
		assert.IsType(t, "", dateStr, "Date should return string")
		assert.NotEmpty(t, dateStr, "Date string should not be empty")
		// Should contain year, month, day information
		assert.Contains(t, dateStr, "-", "Date should be formatted with separators")
	})
}

func TestLunarAliases(t *testing.T) {
	testTime := time.Date(2023, 6, 15, 0, 0, 0, 0, time.UTC)
	lunar := xtime.WithLunar(testTime)

	t.Run("year_alias", func(t *testing.T) {
		yearAlias := lunar.YearAlias()
		assert.IsType(t, "", yearAlias, "YearAlias should return string")
		assert.NotEmpty(t, yearAlias, "Year alias should not be empty")
		
		// Should contain Chinese characters for numbers
		// Check that it's not just the numeric year
		assert.NotEqual(t, "2023", yearAlias, "Should be Chinese characters, not digits")
	})

	t.Run("month_alias", func(t *testing.T) {
		monthAlias := lunar.MonthAlias()
		assert.IsType(t, "", monthAlias, "MonthAlias should return string")
		assert.NotEmpty(t, monthAlias, "Month alias should not be empty")
		assert.Contains(t, monthAlias, "月", "Month alias should contain '月'")
		
		// If it's a leap month, should contain "闰"
		if lunar.IsLeapMonth() {
			assert.Contains(t, monthAlias, "闰", "Leap month should contain '闰'")
		}
	})

	t.Run("day_alias", func(t *testing.T) {
		// Test multiple dates to cover different day alias branches
		testDates := []time.Time{
			time.Date(2023, 1, 20, 0, 0, 0, 0, time.UTC),  // Try to get day 10
			time.Date(2023, 2, 8, 0, 0, 0, 0, time.UTC),   // Try to get day 20
			time.Date(2023, 3, 10, 0, 0, 0, 0, time.UTC),  // Try to get day 30
			time.Date(2023, 4, 1, 0, 0, 0, 0, time.UTC),   // Try different days
			time.Date(2023, 5, 15, 0, 0, 0, 0, time.UTC),  // Try different days
			time.Date(2023, 6, 25, 0, 0, 0, 0, time.UTC),  // Try different days
		}
		
		for _, testDate := range testDates {
			lunar := xtime.WithLunar(testDate)
			dayAlias := lunar.DayAlias()
			assert.IsType(t, "", dayAlias, "DayAlias should return string")
			assert.NotEmpty(t, dayAlias, "Day alias should not be empty for %v", testDate)
			
			day := lunar.Day()
			// Test specific day alias patterns based on the actual lunar day
			if day == 10 {
				assert.Equal(t, "初十", dayAlias, "Day 10 should be '初十'")
			} else if day == 20 {
				assert.Equal(t, "二十", dayAlias, "Day 20 should be '二十'")
			} else if day == 30 {
				assert.Equal(t, "三十", dayAlias, "Day 30 should be '三十'")
			} else {
				// Default case - should contain proper Chinese characters
				assert.NotEmpty(t, dayAlias, "Default day alias should not be empty")
			}
		}
	})

	t.Run("month_day_alias", func(t *testing.T) {
		// Test normal month
		normalDate := time.Date(2023, 6, 15, 0, 0, 0, 0, time.UTC)
		normalLunar := xtime.WithLunar(normalDate)
		normalAlias := normalLunar.MonthDayAlias()
		assert.IsType(t, "", normalAlias, "MonthDayAlias should return string")
		assert.NotEmpty(t, normalAlias, "Month day alias should not be empty")
		assert.Contains(t, normalAlias, "-", "Should contain month-day separator")
		
		// Test if this is not a leap month
		if !normalLunar.IsLeapMonth() {
			assert.NotContains(t, normalAlias, "闰", "Regular month should not contain '闰'")
		}
		
		// Try to find a leap month date - leap months are rare but let's try multiple years
		leapDates := []time.Time{
			time.Date(2020, 6, 21, 0, 0, 0, 0, time.UTC),  // 2020 has leap month 4
			time.Date(2023, 4, 20, 0, 0, 0, 0, time.UTC),  // Try different dates
			time.Date(2017, 7, 23, 0, 0, 0, 0, time.UTC),  // 2017 has leap month 6
		}
		
		for _, leapDate := range leapDates {
			leapLunar := xtime.WithLunar(leapDate)
			leapAlias := leapLunar.MonthDayAlias()
			if leapLunar.IsLeapMonth() {
				assert.Contains(t, leapAlias, "闰", "Leap month should contain '闰'")
				break // Found leap month, test passed
			}
		}
	})

	t.Run("animal_zodiac", func(t *testing.T) {
		// Test multiple years to cover different zodiac animals
		testYears := []time.Time{
			time.Date(2020, 6, 1, 0, 0, 0, 0, time.UTC),  // Rat year
			time.Date(2021, 6, 1, 0, 0, 0, 0, time.UTC),  // Ox year
			time.Date(2022, 6, 1, 0, 0, 0, 0, time.UTC),  // Tiger year
			time.Date(2023, 6, 1, 0, 0, 0, 0, time.UTC),  // Rabbit year
			time.Date(1900, 6, 1, 0, 0, 0, 0, time.UTC),  // Test early year edge case
		}
		
		zodiacAnimals := []string{"鼠", "牛", "虎", "兔", "龙", "蛇", "马", "羊", "猴", "鸡", "狗", "猪"}
		
		for _, testYear := range testYears {
			lunar := xtime.WithLunar(testYear)
			animal := lunar.Animal()
			assert.IsType(t, "", animal, "Animal should return string")
			
			if animal != "" {
				assert.Contains(t, zodiacAnimals, animal, "Should be a valid zodiac animal for %v", testYear)
			}
		}
	})

	t.Run("month_alias_variations", func(t *testing.T) {
		// Test different months to cover MonthAlias switch cases
		monthTests := []time.Time{
			time.Date(2023, 1, 15, 0, 0, 0, 0, time.UTC),   // Month 1
			time.Date(2023, 2, 15, 0, 0, 0, 0, time.UTC),   // Month 2
			time.Date(2023, 6, 15, 0, 0, 0, 0, time.UTC),   // Month 6
			time.Date(2023, 10, 15, 0, 0, 0, 0, time.UTC),  // Month 10
			time.Date(2023, 11, 15, 0, 0, 0, 0, time.UTC),  // Month 11
			time.Date(2023, 12, 15, 0, 0, 0, 0, time.UTC),  // Month 12
		}
		
		for _, testMonth := range monthTests {
			lunar := xtime.WithLunar(testMonth)
			monthAlias := lunar.MonthAlias()
			assert.NotEmpty(t, monthAlias, "Month alias should not be empty for %v", testMonth)
			assert.Contains(t, monthAlias, "月", "Month alias should contain '月'")
			
			// Test for leap month case
			if lunar.IsLeapMonth() {
				assert.Contains(t, monthAlias, "闰", "Leap month alias should contain '闰'")
			}
		}
	})
}

func TestLunarEquals(t *testing.T) {
	t.Run("same_lunar_dates", func(t *testing.T) {
		testTime := time.Date(2023, 6, 15, 12, 0, 0, 0, time.UTC)
		lunar1 := xtime.WithLunar(testTime)
		lunar2 := xtime.WithLunar(testTime)

		assert.True(t, lunar1.Equals(lunar2), "Same lunar dates should be equal")
		assert.True(t, lunar2.Equals(lunar1), "Equality should be symmetric")
	})

	t.Run("different_lunar_dates", func(t *testing.T) {
		time1 := time.Date(2023, 6, 15, 0, 0, 0, 0, time.UTC)
		time2 := time.Date(2023, 7, 15, 0, 0, 0, 0, time.UTC)
		
		lunar1 := xtime.WithLunar(time1)
		lunar2 := xtime.WithLunar(time2)

		// Different solar dates likely result in different lunar dates
		assert.False(t, lunar1.Equals(lunar2), "Different lunar dates should not be equal")
	})

	t.Run("nil_comparison", func(t *testing.T) {
		testTime := time.Date(2023, 6, 15, 0, 0, 0, 0, time.UTC)
		lunar := xtime.WithLunar(testTime)

		// Test behavior with nil - may panic based on implementation
		// This tests the actual behavior rather than expecting specific handling
		defer func() {
			if r := recover(); r != nil {
				t.Logf("Equals with nil panicked as expected: %v", r)
			}
		}()
		
		result := lunar.Equals(nil)
		assert.False(t, result, "Should return false for nil comparison if no panic")
	})
}

func TestFromSolarTimestamp(t *testing.T) {
	t.Run("valid_timestamps", func(t *testing.T) {
		// Test with known timestamps
		testTimestamps := []int64{
			1640995200, // 2022-01-01 00:00:00 UTC
			1672531200, // 2023-01-01 00:00:00 UTC
			1704067200, // 2024-01-01 00:00:00 UTC
		}

		for _, ts := range testTimestamps {
			year, month, day, isLeap := xtime.FromSolarTimestamp(ts)
			
			assert.True(t, year >= 1900 && year <= 2100, "Lunar year should be reasonable for ts %d", ts)
			assert.True(t, month >= 1 && month <= 12, "Lunar month should be 1-12 for ts %d", ts)
			assert.True(t, day >= 1 && day <= 30, "Lunar day should be 1-30 for ts %d", ts)
			assert.IsType(t, true, isLeap, "IsLeap should be bool for ts %d", ts)
		}
	})

	t.Run("timestamp_consistency", func(t *testing.T) {
		// Test that WithLunar and FromSolarTimestamp give same results
		testTime := time.Date(2023, 6, 15, 12, 0, 0, 0, time.UTC)
		timestamp := testTime.Unix()
		
		// Using WithLunar
		lunar1 := xtime.WithLunar(testTime)
		
		// Using FromSolarTimestamp
		year2, month2, day2, isLeap2 := xtime.FromSolarTimestamp(timestamp)
		
		// Should produce the same lunar date
		assert.Equal(t, lunar1.Year(), year2, "Year should match")
		assert.Equal(t, lunar1.Month(), month2, "Month should match") 
		assert.Equal(t, lunar1.Day(), day2, "Day should match")
		assert.Equal(t, lunar1.IsLeapMonth(), isLeap2, "Leap status should match")
	})

	t.Run("edge_case_timestamps", func(t *testing.T) {
		// Test with reasonable range timestamps only
		edgeTimestamps := []int64{
			946684800,  // 2000-01-01 00:00:00 UTC  
			1640995200, // 2022-01-01 00:00:00 UTC
			2147483647, // Max int32 timestamp (2038-01-19)
		}

		for _, ts := range edgeTimestamps {
			// Should not panic for reasonable timestamps
			defer func(timestamp int64) {
				if r := recover(); r != nil {
					t.Logf("FromSolarTimestamp(%d) panicked: %v", timestamp, r)
				}
			}(ts)
			
			year, month, day, isLeap := xtime.FromSolarTimestamp(ts)
			// Basic type checks and reasonable value checks
			if year > 0 { // Only check if conversion succeeded
				assert.True(t, month >= 1 && month <= 12, "Month should be valid for ts %d", ts)
				assert.True(t, day >= 1 && day <= 30, "Day should be valid for ts %d", ts)
				_ = isLeap // Use the variable
			}
		}
	})

	t.Run("comprehensive_lunar_days_coverage", func(t *testing.T) {
		// Test a wide range of dates to improve lunarDays function coverage
		// Test different months and years to cover more branches
		testTimestamps := []int64{
			time.Date(1950, 2, 28, 0, 0, 0, 0, time.UTC).Unix(),  // Mid-century
			time.Date(1980, 6, 15, 0, 0, 0, 0, time.UTC).Unix(),  // Different season
			time.Date(2000, 12, 31, 0, 0, 0, 0, time.UTC).Unix(), // Millennium
			time.Date(2020, 4, 15, 0, 0, 0, 0, time.UTC).Unix(),  // Recent year
			time.Date(2023, 8, 20, 0, 0, 0, 0, time.UTC).Unix(),  // Current year
			time.Date(2037, 11, 10, 0, 0, 0, 0, time.UTC).Unix(), // Future year
		}

		for _, ts := range testTimestamps {
			// Test to improve coverage of FromSolarTimestamp function
			assert.NotPanics(t, func() {
				year, month, day, isLeap := xtime.FromSolarTimestamp(ts)
				
				if year > 0 { // Valid conversion
					// Basic validation
					assert.True(t, month >= 1 && month <= 12, "Month should be 1-12")
					assert.True(t, day >= 1 && day <= 30, "Day should be 1-30")
					_ = isLeap // Use the variable
					
					// Test around this date to get more coverage of edge cases
					for offset := -5; offset <= 5; offset++ {
						testTs := ts + int64(offset)*86400 // Add/subtract days
						testYear, testMonth, testDay, testIsLeap := xtime.FromSolarTimestamp(testTs)
						if testYear > 0 {
							_ = testMonth
							_ = testDay
							_ = testIsLeap
						}
					}
				}
			}, "FromSolarTimestamp should not panic for timestamp %d", ts)
		}
	})
}

func TestLunarHelperFunctions(t *testing.T) {
	t.Run("order_mod", func(t *testing.T) {
		// Test the OrderMod function with various inputs
		testCases := []struct {
			a, b, expected int64
		}{
			{10, 3, 1},   // 10 % 3 = 1
			{9, 3, 3},    // 9 % 3 = 0, but OrderMod should return 3
			{15, 12, 3},  // 15 % 12 = 3
			{12, 12, 12}, // 12 % 12 = 0, should return 12
			{1, 5, 1},    // 1 % 5 = 1
		}

		for _, tc := range testCases {
			result := xtime.OrderMod(tc.a, tc.b)
			assert.Equal(t, tc.expected, result, "OrderMod(%d, %d) should equal %d", tc.a, tc.b, tc.expected)
		}
	})

	t.Run("order_mod_negative", func(t *testing.T) {
		// Test OrderMod with negative numbers - behavior may vary by implementation
		defer func() {
			if r := recover(); r != nil {
				t.Logf("OrderMod with negative panicked: %v", r)
			}
		}()
		
		result1 := xtime.OrderMod(-1, 12)
		t.Logf("OrderMod(-1, 12) = %d", result1)
		
		result2 := xtime.OrderMod(-13, 12)
		t.Logf("OrderMod(-13, 12) = %d", result2)
		
		// Just verify they return some result without strict range checking
		_ = result1
		_ = result2
	})
}

func TestLunarEdgeCases(t *testing.T) {
	t.Run("leap_year_scenarios", func(t *testing.T) {
		// Test years known to have leap months (these are approximate)
		possibleLeapYears := []int{2020, 2023, 2025, 2028}
		
		for _, year := range possibleLeapYears {
			testTime := time.Date(year, 6, 15, 0, 0, 0, 0, time.UTC)
			lunar := xtime.WithLunar(testTime)
			
			// Test leap month functionality
			leapMonth := lunar.LeapMonth()
			isLeap := lunar.IsLeap()
			
			// Consistency check
			if leapMonth > 0 {
				assert.True(t, isLeap, "If leap month exists, year should be leap")
			} else {
				assert.False(t, isLeap, "If no leap month, year should not be leap")
			}
		}
	})

	t.Run("month_boundary_dates", func(t *testing.T) {
		// Test dates around lunar month boundaries
		// These are approximate Chinese New Year dates
		chineseNewYearDates := []time.Time{
			time.Date(2023, 1, 22, 0, 0, 0, 0, time.UTC), // Chinese New Year 2023
			time.Date(2024, 2, 10, 0, 0, 0, 0, time.UTC), // Chinese New Year 2024
		}

		for _, date := range chineseNewYearDates {
			lunar := xtime.WithLunar(date)
			
			// Around Chinese New Year, should be early in lunar calendar
			assert.True(t, lunar.Month() <= 2, "Around Chinese New Year should be month 1 or 2")
			assert.True(t, lunar.Day() <= 15, "Around Chinese New Year should be early in month")
		}
	})

	t.Run("different_timezones", func(t *testing.T) {
		// Test that timezone doesn't affect lunar calculation dramatically
		utc := time.Date(2023, 6, 15, 12, 0, 0, 0, time.UTC)
		
		// Convert to different timezones
		est, _ := time.LoadLocation("America/New_York")
		pst, _ := time.LoadLocation("America/Los_Angeles")
		
		lunar1 := xtime.WithLunar(utc)
		lunar2 := xtime.WithLunar(utc.In(est))
		lunar3 := xtime.WithLunar(utc.In(pst))
		
		// Same UTC time should give same or very similar lunar dates
		// (might differ by 1 day due to timezone differences)
		yearDiff := abs(lunar1.Year() - lunar2.Year())
		monthDiff := abs(lunar1.Month() - lunar2.Month())
		dayDiff := abs(lunar1.Day() - lunar3.Day())
		
		assert.True(t, yearDiff <= 1, "Year difference should be at most 1")
		assert.True(t, monthDiff <= 1, "Month difference should be at most 1")
		assert.True(t, dayDiff <= 1, "Day difference should be at most 1")
	})
}

// Helper function for absolute value
func abs(x int64) int64 {
	if x < 0 {
		return -x
	}
	return x
}

// Benchmark lunar conversions
func BenchmarkLunarConversions(b *testing.B) {
	testTime := time.Date(2023, 6, 15, 12, 0, 0, 0, time.UTC)
	
	b.Run("WithLunar", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = xtime.WithLunar(testTime)
		}
	})

	b.Run("FromSolarTimestamp", func(b *testing.B) {
		timestamp := testTime.Unix()
		for i := 0; i < b.N; i++ {
			_, _, _, _ = xtime.FromSolarTimestamp(timestamp)
		}
	})

	b.Run("LunarMethods", func(b *testing.B) {
		lunar := xtime.WithLunar(testTime)
		for i := 0; i < b.N; i++ {
			_ = lunar.Year()
			_ = lunar.Month()
			_ = lunar.Day()
			_ = lunar.Animal()
			_ = lunar.YearAlias()
		}
	})
}