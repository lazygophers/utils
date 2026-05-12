package xtime

import (
	"testing"
	"time"
)

// TestEndOfQuarterGlobal_Correctness 验证所有变体的正确性
func TestEndOfQuarterGlobal_Correctness(t *testing.T) {
	// 固定测试时间
	testTime := time.Date(2024, time.February, 15, 10, 30, 45, 123456789, time.Local)
	expected := With(testTime).EndOfQuarter()

	// 测试所有变体
	variants := []struct {
		name string
		fn   func() *Time
	}{
		{
			name: "Original",
			fn: func() *Time {
				return With(testTime).EndOfQuarter()
			},
		},
		{
			name: "Variant1",
			fn: func() *Time {
				year, month, _ := testTime.Date()
				quarter := (month-1)/3 + 1
				endQuarterMonth := quarter * 3
				return &Time{
					Time:   time.Date(year, time.Month(endQuarterMonth+1), 0, 23, 59, 59, 999999999, testTime.Location()),
					Config: &Config{WeekStartDay: time.Monday, TimeLocation: time.Local, TimeFormats: []string{}, Monotonic: time.Now()},
				}
			},
		},
		{
			name: "Variant7",
			fn: func() *Time {
				year := testTime.Year()
				month := testTime.Month()
				quarter := (month-1)/3 + 1
				endQuarterMonth := quarter * 3
				return &Time{
					Time:   time.Date(year, endQuarterMonth+1, 0, 23, 59, 59, 999999999, testTime.Location()),
					Config: &Config{WeekStartDay: time.Monday, TimeLocation: testTime.Location()},
				}
			},
		},
		{
			name: "Variant11",
			fn: func() *Time {
				return &Time{
					Time:   time.Date(testTime.Year(), ((testTime.Month()-1)/3+1)*3+1, 0, 23, 59, 59, 999999999, testTime.Location()),
					Config: &Config{WeekStartDay: time.Monday, TimeLocation: testTime.Location()},
				}
			},
		},
	}

	for _, tt := range variants {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.fn()
			if !result.Time.Equal(expected.Time) {
				t.Errorf("%s: Time = %v, want %v", tt.name, result.Time, expected.Time)
			}
		})
	}
}

// TestEndOfQuarterGlobal_AllQuarters 测试所有季度的边界情况
func TestEndOfQuarterGlobal_AllQuarters(t *testing.T) {
	testCases := []struct {
		name     string
		input    time.Time
		expected time.Time
	}{
		{
			name:     "Q1 (Jan)",
			input:    time.Date(2024, time.January, 15, 12, 0, 0, 0, time.Local),
			expected: time.Date(2024, time.March, 31, 23, 59, 59, 999999999, time.Local),
		},
		{
			name:     "Q2 (Apr)",
			input:    time.Date(2024, time.April, 15, 12, 0, 0, 0, time.Local),
			expected: time.Date(2024, time.June, 30, 23, 59, 59, 999999999, time.Local),
		},
		{
			name:     "Q3 (Jul)",
			input:    time.Date(2024, time.July, 15, 12, 0, 0, 0, time.Local),
			expected: time.Date(2024, time.September, 30, 23, 59, 59, 999999999, time.Local),
		},
		{
			name:     "Q4 (Oct)",
			input:    time.Date(2024, time.October, 15, 12, 0, 0, 0, time.Local),
			expected: time.Date(2024, time.December, 31, 23, 59, 59, 999999999, time.Local),
		},
	}

	// 测试最优变体
	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			result := (&Time{Time: tt.input}).EndOfQuarter()
			if !result.Time.Equal(tt.expected) {
				t.Errorf("EndOfQuarter() = %v, want %v", result.Time, tt.expected)
			}
		})
	}
}

// Manual performance test
func TestEndOfQuarterGlobal_Performance(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping performance test in short mode")
	}

	iterations := 100000

	// Original
	t.Run("Original", func(t *testing.T) {
		for i := 0; i < iterations; i++ {
			_ = With(time.Now()).EndOfQuarter()
		}
	})

	// Variant7
	t.Run("Variant7", func(t *testing.T) {
		for i := 0; i < iterations; i++ {
			now := time.Now()
			year := now.Year()
			month := now.Month()
			quarter := (month-1)/3 + 1
			endQuarterMonth := quarter * 3
			_ = &Time{
				Time:   time.Date(year, endQuarterMonth+1, 0, 23, 59, 59, 999999999, now.Location()),
				Config: &Config{WeekStartDay: time.Monday, TimeLocation: now.Location()},
			}
		}
	})

	// Variant11
	t.Run("Variant11", func(t *testing.T) {
		for i := 0; i < iterations; i++ {
			t := time.Now()
			_ = &Time{
				Time:   time.Date(t.Year(), ((t.Month()-1)/3+1)*3+1, 0, 23, 59, 59, 999999999, t.Location()),
				Config: &Config{WeekStartDay: time.Monday, TimeLocation: t.Location()},
			}
		}
	})
}
