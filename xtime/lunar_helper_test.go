package xtime

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestLunarHelper_GetYearFestivals(t *testing.T) {
	helper := NewLunarHelper()
	festivals := helper.GetYearFestivals(2024)
	assert.NotNil(t, festivals)
	t.Logf("Found %d festivals in 2024", len(festivals))
}

func TestLunarHelper_GetUpcomingFestivals(t *testing.T) {
	helper := NewLunarHelper()
	upcoming := helper.GetUpcomingFestivals(5)
	assert.NotNil(t, upcoming)
	t.Logf("Found %d upcoming festivals", len(upcoming))
}

func TestLunarHelper_GetNextLunarBirthday(t *testing.T) {
	helper := NewLunarHelper()

	// Test with a birthday
	birthday := time.Date(1990, 5, 15, 0, 0, 0, 0, time.UTC)
	now := time.Now()
	next := helper.GetNextLunarBirthday(birthday, now)
	t.Logf("Next lunar birthday: %v", next)
}

func TestLunarHelper_GetFestival(t *testing.T) {
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
}

func TestLunarHelper_FormatFestivalInfo(t *testing.T) {
	helper := NewLunarHelper()

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
}

func TestLunarHelper_IsSpecialDay(t *testing.T) {
	helper := NewLunarHelper()

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
}

func TestLunarHelper_GetLunarInfo(t *testing.T) {
	helper := NewLunarHelper()

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
}

func TestLunarHelper_CompareLunarDates(t *testing.T) {
	helper := NewLunarHelper()

	// Test CompareLunarDates
	date1 := time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC)
	date2 := time.Date(2024, 12, 31, 0, 0, 0, 0, time.UTC)

	result := helper.CompareLunarDates(date1, date2)
	t.Logf("Comparing %v and %v: %s", date1.Format("2006-01-02"), date2.Format("2006-01-02"), result)

	// Test with same dates
	sameResult := helper.CompareLunarDates(date1, date1)
	assert.NotEmpty(t, sameResult)
}

func TestLunarHelper_GetLunarAge(t *testing.T) {
	helper := NewLunarHelper()

	// Test GetLunarAge
	birthday := time.Date(1990, 5, 15, 0, 0, 0, 0, time.UTC)
	now := time.Now()

	age := helper.GetLunarAge(birthday, now)
	assert.True(t, age >= 0)
	t.Logf("Lunar age: %d", age)
}
