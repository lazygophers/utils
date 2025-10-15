package xtime

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestSolarTermHelper_GetTermsInRange(t *testing.T) {
	helper := NewSolarTermHelper()

	start := time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC)
	end := time.Date(2024, 12, 31, 23, 59, 59, 0, time.UTC)

	terms := helper.GetTermsInRange(start, end)
	assert.NotNil(t, terms)
	t.Logf("Found %d solar terms in 2024", len(terms))

	// Should have 24 solar terms in a year
	assert.Equal(t, 24, len(terms))
}

func TestSolarTermHelper_GetCurrentTerm(t *testing.T) {
	helper := NewSolarTermHelper()

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
}

func TestSolarTermHelper_GetNextTerm(t *testing.T) {
	helper := NewSolarTermHelper()

	testDates := []time.Time{
		time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC),
		time.Date(2024, 12, 31, 0, 0, 0, 0, time.UTC),
	}

	for _, date := range testDates {
		term := helper.GetNextTerm(date)
		assert.NotNil(t, term)
		t.Logf("Next term after %v: %v", date.Format("2006-01-02"), term)
	}
}

func TestSolarTermHelper_FindTermByName(t *testing.T) {
	helper := NewSolarTermHelper()

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
}

func TestSolarTermHelper_FormatTermInfo(t *testing.T) {
	helper := NewSolarTermHelper()

	// Test FormatTermInfo with different terms
	terms := helper.GetYearTerms(2024)
	if len(terms) > 0 {
		info := helper.FormatTermInfo(&terms[0])
		assert.NotEmpty(t, info)
		t.Logf("Term info: %s", info)
	}
}

func TestSolarTermHelper_DaysUntilTerm(t *testing.T) {
	helper := NewSolarTermHelper()

	// Test DaysUntilTerm
	now := time.Now()
	days := helper.DaysUntilTerm(now, "立春")
	t.Logf("Days until term: %d", days)
}
