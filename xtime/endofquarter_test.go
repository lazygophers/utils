package xtime

import (
	"testing"
	"time"
)

func TestEndOfQuarter_Correctness(t *testing.T) {
	tests := []struct {
		name     string
		input    time.Time
		expected time.Time
	}{
		{
			name:     "Q1 - 结束于3月31日",
			input:    time.Date(2024, 1, 15, 12, 0, 0, 0, time.Local),
			expected: time.Date(2024, 3, 31, 23, 59, 59, 999999999, time.Local),
		},
		{
			name:     "Q1 - 边界: 3月1日",
			input:    time.Date(2024, 3, 1, 0, 0, 0, 0, time.Local),
			expected: time.Date(2024, 3, 31, 23, 59, 59, 999999999, time.Local),
		},
		{
			name:     "Q2 - 结束于6月30日",
			input:    time.Date(2024, 4, 15, 12, 0, 0, 0, time.Local),
			expected: time.Date(2024, 6, 30, 23, 59, 59, 999999999, time.Local),
		},
		{
			name:     "Q3 - 结束于9月30日",
			input:    time.Date(2024, 7, 15, 12, 0, 0, 0, time.Local),
			expected: time.Date(2024, 9, 30, 23, 59, 59, 999999999, time.Local),
		},
		{
			name:     "Q4 - 结束于12月31日",
			input:    time.Date(2024, 10, 15, 12, 0, 0, 0, time.Local),
			expected: time.Date(2024, 12, 31, 23, 59, 59, 999999999, time.Local),
		},
		{
			name:     "Q4 - 跨年边界",
			input:    time.Date(2024, 12, 31, 23, 59, 59, 999999998, time.Local),
			expected: time.Date(2024, 12, 31, 23, 59, 59, 999999999, time.Local),
		},
		{
			name:     "Q1 - 年初",
			input:    time.Date(2024, 1, 1, 0, 0, 0, 0, time.Local),
			expected: time.Date(2024, 3, 31, 23, 59, 59, 999999999, time.Local),
		},
		{
			name:     "闰年测试: 2月29日",
			input:    time.Date(2024, 2, 29, 12, 0, 0, 0, time.Local),
			expected: time.Date(2024, 3, 31, 23, 59, 59, 999999999, time.Local),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := With(tt.input).EndOfQuarter()
			if !result.Time.Equal(tt.expected) {
				t.Errorf("EndOfQuarter() = %v, want %v", result.Time, tt.expected)
			}
		})
	}
}

func TestEndOfQuarter_ConfigPreserved(t *testing.T) {
	cfg := &Config{WeekStartDay: time.Monday}
	xtime := &Time{
		Time:   time.Date(2024, 5, 15, 12, 0, 0, 0, time.Local),
		Config: cfg,
	}

	result := xtime.EndOfQuarter()

	if result.Config != cfg {
		t.Errorf("EndOfQuarter() Config not preserved, got %v, want %v", result.Config, cfg)
	}

	if result.Config.WeekStartDay != time.Monday {
		t.Errorf("EndOfQuarter() WeekStartDay not preserved, got %v, want %v", result.Config.WeekStartDay, time.Monday)
	}
}

func TestEndOfQuarter_TimeProperties(t *testing.T) {
	testDate := time.Date(2024, 5, 15, 10, 30, 45, 123456789, time.Local)
	result := With(testDate).EndOfQuarter()

	// 验证是季度的最后一天
	_, month, _ := result.Date()
	if month != time.March && month != time.June && month != time.September && month != time.December {
		t.Errorf("EndOfQuarter() month = %v, want quarter end month (3,6,9,12)", month)
	}

	// 验证是一天的最后一刻
	hour, min, sec := result.Clock()
	if hour != 23 || min != 59 || sec != 59 {
		t.Errorf("EndOfQuarter() time = %d:%d:%d, want 23:59:59", hour, min, sec)
	}

	if result.Nanosecond() != 999999999 {
		t.Errorf("EndOfQuarter() nanosecond = %d, want 999999999", result.Nanosecond())
	}
}

func TestEndOfQuarter_ConsistencyWithBeginningOfQuarter(t *testing.T) {
	testDates := []time.Time{
		time.Date(2024, 1, 15, 12, 0, 0, 0, time.Local),
		time.Date(2024, 4, 15, 12, 0, 0, 0, time.Local),
		time.Date(2024, 7, 15, 12, 0, 0, 0, time.Local),
		time.Date(2024, 10, 15, 12, 0, 0, 0, time.Local),
	}

	for _, tt := range testDates {
		t.Run(tt.Format("2006-01-02"), func(t *testing.T) {
			tt := With(tt)
			beginningOfQuarter := tt.BeginningOfQuarter()
			endOfQuarter := tt.EndOfQuarter()

			// 下一季度开始前1纳秒 = 季度结束
			nextQuarter := beginningOfQuarter.AddDate(0, 3, 0)
			expectedEnd := nextQuarter.Add(-time.Nanosecond)

			if !endOfQuarter.Time.Equal(expectedEnd) {
				t.Errorf("EndOfQuarter() = %v, BeginningOfQuarter()+3mo-1ns = %v", endOfQuarter.Time, expectedEnd)
			}
		})
	}
}
