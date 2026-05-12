package xtime

import (
	"testing"
	"time"
)

// TestBeginningOfHalf_Correctness 验证 BeginningOfHalf 正确性
func TestBeginningOfHalf_Correctness(t *testing.T) {
	tests := []struct {
		name     string
		input    time.Time
		expected time.Time
	}{
		{
			name:     "H1起始 - 1月",
			input:    time.Date(2024, 1, 15, 12, 30, 45, 0, time.Local),
			expected: time.Date(2024, 1, 1, 0, 0, 0, 0, time.Local),
		},
		{
			name:     "H1中间 - 3月",
			input:    time.Date(2024, 3, 15, 12, 30, 45, 0, time.Local),
			expected: time.Date(2024, 1, 1, 0, 0, 0, 0, time.Local),
		},
		{
			name:     "H1结束 - 6月",
			input:    time.Date(2024, 6, 15, 12, 30, 45, 0, time.Local),
			expected: time.Date(2024, 1, 1, 0, 0, 0, 0, time.Local),
		},
		{
			name:     "H2起始 - 7月",
			input:    time.Date(2024, 7, 15, 12, 30, 45, 0, time.Local),
			expected: time.Date(2024, 7, 1, 0, 0, 0, 0, time.Local),
		},
		{
			name:     "H2中间 - 9月",
			input:    time.Date(2024, 9, 15, 12, 30, 45, 0, time.Local),
			expected: time.Date(2024, 7, 1, 0, 0, 0, 0, time.Local),
		},
		{
			name:     "H2结束 - 12月",
			input:    time.Date(2024, 12, 15, 12, 30, 45, 0, time.Local),
			expected: time.Date(2024, 7, 1, 0, 0, 0, 0, time.Local),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			pt := With(tt.input)
			result := pt.BeginningOfHalf()

			if result.Time.Year() != tt.expected.Year() ||
				result.Time.Month() != tt.expected.Month() ||
				result.Time.Day() != tt.expected.Day() ||
				result.Time.Hour() != tt.expected.Hour() ||
				result.Time.Minute() != tt.expected.Minute() ||
				result.Time.Second() != tt.expected.Second() {
				t.Errorf("BeginningOfHalf() = %v, want %v", result.Time, tt.expected)
			}
		})
	}
}

// TestBeginningOfHalf_完整性 测试所有月份
func TestBeginningOfHalf_Completeness(t *testing.T) {
	for year := 2020; year <= 2025; year++ {
		for month := 1; month <= 12; month++ {
			input := time.Date(year, time.Month(month), 15, 12, 30, 45, 0, time.Local)
			pt := With(input)
			result := pt.BeginningOfHalf()

			expectedMonth := time.Month(1)
			if month > 6 {
				expectedMonth = 7
			}

			if result.Time.Year() != year || result.Time.Month() != expectedMonth {
				t.Errorf("Year %d Month %d: got year=%d month=%d, want year=%d month=%d",
					year, month, result.Time.Year(), result.Time.Month(), year, expectedMonth)
			}

			if result.Time.Day() != 1 {
				t.Errorf("Year %d Month %d: got day=%d, want 1", year, month, result.Time.Day())
			}

			if result.Time.Hour() != 0 || result.Time.Minute() != 0 || result.Time.Second() != 0 {
				t.Errorf("Year %d Month %d: got time=%d:%d:%d, want 0:0:0",
					year, month, result.Time.Hour(), result.Time.Minute(), result.Time.Second())
			}
		}
	}
}
