package xtime

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestBeginningOfQuarter_Correctness(t *testing.T) {
	tests := []struct {
		name     string
		input    time.Time
		expected monthDay
	}{
		{"Q1 Jan", time.Date(2024, 1, 15, 0, 0, 0, 0, time.UTC), monthDay{1, 1}},
		{"Q1 Feb", time.Date(2024, 2, 15, 0, 0, 0, 0, time.UTC), monthDay{1, 1}},
		{"Q1 Mar", time.Date(2024, 3, 15, 0, 0, 0, 0, time.UTC), monthDay{1, 1}},
		{"Q2 Apr", time.Date(2024, 4, 15, 0, 0, 0, 0, time.UTC), monthDay{4, 1}},
		{"Q2 May", time.Date(2024, 5, 15, 0, 0, 0, 0, time.UTC), monthDay{4, 1}},
		{"Q2 Jun", time.Date(2024, 6, 15, 0, 0, 0, 0, time.UTC), monthDay{4, 1}},
		{"Q3 Jul", time.Date(2024, 7, 15, 0, 0, 0, 0, time.UTC), monthDay{7, 1}},
		{"Q3 Aug", time.Date(2024, 8, 15, 0, 0, 0, 0, time.UTC), monthDay{7, 1}},
		{"Q3 Sep", time.Date(2024, 9, 15, 0, 0, 0, 0, time.UTC), monthDay{7, 1}},
		{"Q4 Oct", time.Date(2024, 10, 15, 0, 0, 0, 0, time.UTC), monthDay{10, 1}},
		{"Q4 Nov", time.Date(2024, 11, 15, 0, 0, 0, 0, time.UTC), monthDay{10, 1}},
		{"Q4 Dec", time.Date(2024, 12, 15, 0, 0, 0, 0, time.UTC), monthDay{10, 1}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := With(tt.input).BeginningOfQuarter()
			assert.Equal(t, tt.expected.month, int(result.Month()))
			assert.Equal(t, tt.expected.day, result.Day())
			assert.Equal(t, 0, result.Hour())
			assert.Equal(t, 0, result.Minute())
			assert.Equal(t, 0, result.Second())
		})
	}
}

type monthDay struct {
	month int
	day   int
}
