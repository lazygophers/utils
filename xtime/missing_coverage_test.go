package xtime

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

// TestUncoveredFunctions tests functions with 0% coverage
func TestUncoveredFunctions(t *testing.T) {
	t.Run("ExampleUsage", func(t *testing.T) {
		// Test ExampleUsage function
		ExampleUsage()
		t.Log("ExampleUsage completed successfully")
	})

	t.Run("GetYearFestivals", func(t *testing.T) {
		helper := NewLunarHelper()
		festivals := helper.GetYearFestivals(2024)
		assert.NotNil(t, festivals)
		t.Logf("Found %d festivals in 2024", len(festivals))
	})

	t.Run("GetUpcomingFestivals", func(t *testing.T) {
		helper := NewLunarHelper()
		upcoming := helper.GetUpcomingFestivals(5)
		assert.NotNil(t, upcoming)
		t.Logf("Found %d upcoming festivals", len(upcoming))
	})

	t.Run("GetNextLunarBirthday", func(t *testing.T) {
		helper := NewLunarHelper()

		// Test with a birthday
		birthday := time.Date(1990, 5, 15, 0, 0, 0, 0, time.UTC)
		now := time.Now()
		next := helper.GetNextLunarBirthday(birthday, now)
		t.Logf("Next lunar birthday: %v", next)
	})

	t.Run("GetTermsInRange", func(t *testing.T) {
		helper := NewSolarTermHelper()

		start := time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC)
		end := time.Date(2024, 12, 31, 23, 59, 59, 0, time.UTC)

		terms := helper.GetTermsInRange(start, end)
		assert.NotNil(t, terms)
		t.Logf("Found %d solar terms in 2024", len(terms))

		// Should have 24 solar terms in a year
		assert.Equal(t, 24, len(terms))
	})
}

// TestLowCoverageFunctions tests functions with low coverage to improve them
func TestLowCoverageFunctions(t *testing.T) {
	t.Run("DayAlias_edge_cases", func(t *testing.T) {
		// Test DayAlias with various edge cases
		testCases := []int{1, 10, 11, 20, 30, 31}

		for _, day := range testCases {
			lunar := &Lunar{day: int64(day)}
			alias := lunar.DayAlias()
			assert.NotEmpty(t, alias)
			t.Logf("Day %d alias: %s", day, alias)
		}
	})

	t.Run("MonthAlias_edge_cases", func(t *testing.T) {
		// Test MonthAlias with various cases
		testCases := []int{1, 11, 12}

		for _, month := range testCases {
			lunar := &Lunar{month: int64(month)}
			alias := lunar.MonthAlias()
			assert.NotEmpty(t, alias)
			t.Logf("Month %d alias: %s", month, alias)
		}
	})

	t.Run("Animal_edge_cases", func(t *testing.T) {
		lunar := &Lunar{}

		// Test with different years
		testYears := []int{1900, 1984, 2000, 2024}

		for _, year := range testYears {
			lunar.year = int64(year)
			animal := lunar.Animal()
			assert.NotEmpty(t, animal)
			t.Logf("Year %d animal: %s", year, animal)
		}
	})

	t.Run("lunarDays_edge_cases", func(t *testing.T) {
		// Test lunarDays with different years and months
		testCases := []struct {
			year  int
			month int
		}{
			{2024, 1},
			{2024, 12},
			{2020, 4}, // Leap year
		}

		for _, tc := range testCases {
			days := lunarDays(int64(tc.year), int64(tc.month))
			assert.True(t, days >= 29 && days <= 30)
			t.Logf("Year %d Month %d has %d days", tc.year, tc.month, days)
		}
	})

	t.Run("GetFestival_edge_cases", func(t *testing.T) {
		helper := NewLunarHelper()

		// Test with different dates
		testDates := []time.Time{
			time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC),   // New Year
			time.Date(2024, 2, 14, 0, 0, 0, 0, time.UTC),  // Valentine's Day
			time.Date(2024, 12, 25, 0, 0, 0, 0, time.UTC), // Christmas
		}

		for _, date := range testDates {
			festival := helper.GetFestival(date)
			if festival != nil {
				t.Logf("Festival on %v: %s", date.Format("2006-01-02"), festival.Name)
			} else {
				t.Logf("No festival on %v", date.Format("2006-01-02"))
			}
		}
	})

	t.Run("calculateSeason_edge_cases", func(t *testing.T) {
		// Test calculateSeason with different dates via Calendar
		testDates := []time.Time{
			time.Date(2024, 3, 15, 0, 0, 0, 0, time.UTC), // Spring
			time.Date(2024, 6, 15, 0, 0, 0, 0, time.UTC), // Summer
			time.Date(2024, 9, 15, 0, 0, 0, 0, time.UTC), // Autumn
		}

		for _, date := range testDates {
			calendar := NewCalendar(date)
			season := calendar.Season()
			assert.NotEmpty(t, season)
			t.Logf("Season on %v: %s", date.Format("2006-01-02"), season)
		}
	})

	t.Run("GetTodayLucky_edge_cases", func(t *testing.T) {
		// Test GetTodayLucky with different scenarios
		lucky := GetTodayLucky()
		assert.NotEmpty(t, lucky)
		t.Logf("Today's lucky info: %s", lucky)
	})

	t.Run("FormatTodayInfo_edge_cases", func(t *testing.T) {
		// Test FormatTodayInfo
		info := FormatTodayInfo()
		assert.NotEmpty(t, info)
		t.Logf("Today's info: %s", info)
	})

	t.Run("QuickExample_coverage", func(t *testing.T) {
		// Test QuickExample to improve its coverage
		QuickExample()
		t.Log("QuickExample completed successfully")
	})
}

// TestSolarTermHelperLowCoverage tests low coverage functions in solar term helper
func TestSolarTermHelperLowCoverage(t *testing.T) {
	helper := NewSolarTermHelper()

	t.Run("GetCurrentTerm_edge_cases", func(t *testing.T) {
		// Test with different times of year
		testDates := []time.Time{
			time.Date(2024, 1, 5, 0, 0, 0, 0, time.UTC),
			time.Date(2024, 6, 15, 0, 0, 0, 0, time.UTC),
			time.Date(2024, 12, 15, 0, 0, 0, 0, time.UTC),
		}

		for _, date := range testDates {
			term := helper.GetCurrentTerm(date)
			assert.NotNil(t, term)
			t.Logf("Current term on %v: %v", date.Format("2006-01-02"), term)
		}
	})

	t.Run("GetNextTerm_edge_cases", func(t *testing.T) {
		testDates := []time.Time{
			time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC),
			time.Date(2024, 12, 31, 0, 0, 0, 0, time.UTC),
		}

		for _, date := range testDates {
			term := helper.GetNextTerm(date)
			assert.NotNil(t, term)
			t.Logf("Next term after %v: %v", date.Format("2006-01-02"), term)
		}
	})

	t.Run("FindTermByName_edge_cases", func(t *testing.T) {
		// Test finding terms by different names
		testNames := []string{
			"立春", "春分", "立夏", "夏至",
			"立秋", "秋分", "立冬", "冬至",
			"nonexistent", // This should return nil
		}

		for _, name := range testNames {
			term := helper.FindTermByName(2024, name)
			if name == "nonexistent" {
				assert.Nil(t, term)
			} else {
				assert.NotNil(t, term)
				t.Logf("Found term %s: %v", name, term)
			}
		}
	})

	t.Run("FormatTermInfo_edge_cases", func(t *testing.T) {
		// Test FormatTermInfo with different terms
		terms := helper.GetYearTerms(2024)
		if len(terms) > 0 {
			info := helper.FormatTermInfo(&terms[0])
			assert.NotEmpty(t, info)
			t.Logf("Term info: %s", info)
		}
	})

	t.Run("DaysUntilTerm_edge_cases", func(t *testing.T) {
		// Test DaysUntilTerm
		now := time.Now()
		days := helper.DaysUntilTerm(now, "立春")
		t.Logf("Days until term: %d", days)
	})
}

// TestLunarHelperLowCoverage tests low coverage functions in lunar helper
func TestLunarHelperLowCoverage(t *testing.T) {
	helper := NewLunarHelper()

	t.Run("FormatFestivalInfo_edge_cases", func(t *testing.T) {
		// Test FormatFestivalInfo with an actual festival
		festival := helper.GetTodayFestival()
		if festival != nil {
			info := helper.FormatFestivalInfo(festival)
			assert.NotEmpty(t, info)
			t.Logf("Festival info: %s", info)
		}

		// Test with nil festival
		emptyInfo := helper.FormatFestivalInfo(nil)
		t.Logf("Empty festival info: %s", emptyInfo)
	})

	t.Run("IsSpecialDay_edge_cases", func(t *testing.T) {
		// Test IsSpecialDay with various dates
		testDates := []time.Time{
			time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC),
			time.Date(2024, 2, 14, 0, 0, 0, 0, time.UTC),
			time.Date(2024, 10, 31, 0, 0, 0, 0, time.UTC),
		}

		for _, date := range testDates {
			isSpecial, reason := helper.IsSpecialDay(date)
			t.Logf("Is %v special? %v (reason: %s)", date.Format("2006-01-02"), isSpecial, reason)
		}
	})

	t.Run("GetLunarInfo_edge_cases", func(t *testing.T) {
		// Test GetLunarInfo with different dates
		testDates := []time.Time{
			time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC),
			time.Date(2024, 6, 15, 0, 0, 0, 0, time.UTC),
		}

		for _, date := range testDates {
			info := helper.GetLunarInfo(date)
			assert.NotEmpty(t, info)
			t.Logf("Lunar info for %v: %s", date.Format("2006-01-02"), info)
		}
	})

	t.Run("CompareLunarDates_edge_cases", func(t *testing.T) {
		// Test CompareLunarDates
		date1 := time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC)
		date2 := time.Date(2024, 12, 31, 0, 0, 0, 0, time.UTC)

		result := helper.CompareLunarDates(date1, date2)
		t.Logf("Comparing %v and %v: %s", date1.Format("2006-01-02"), date2.Format("2006-01-02"), result)

		// Test with same dates
		sameResult := helper.CompareLunarDates(date1, date1)
		assert.NotEmpty(t, sameResult)
	})

	t.Run("GetLunarAge_edge_cases", func(t *testing.T) {
		// Test GetLunarAge
		birthday := time.Date(1990, 5, 15, 0, 0, 0, 0, time.UTC)
		now := time.Now()

		age := helper.GetLunarAge(birthday, now)
		assert.True(t, age >= 0)
		t.Logf("Lunar age: %d", age)
	})
}