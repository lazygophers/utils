package xtime

import (
	"testing"
	"time"
)

func TestEndOfHalf_Correctness(t *testing.T) {
	tests := []struct {
		name        string
		input       string
		expectedMonth time.Month
		expectedDay  int
	}{
		{
			name:        "上半年开始",
			input:       "2024-01-01 00:00:00",
			expectedMonth: time.June,
			expectedDay:  30,
		},
		{
			name:        "上半年中间",
			input:       "2024-03-15 14:30:45",
			expectedMonth: time.June,
			expectedDay:  30,
		},
		{
			name:        "上半年结束",
			input:       "2024-06-30 23:59:59",
			expectedMonth: time.June,
			expectedDay:  30,
		},
		{
			name:        "下半年开始",
			input:       "2024-07-01 00:00:00",
			expectedMonth: time.December,
			expectedDay:  31,
		},
		{
			name:        "下半年中间",
			input:       "2024-09-15 14:30:45",
			expectedMonth: time.December,
			expectedDay:  31,
		},
		{
			name:        "下半年结束",
			input:       "2024-12-31 23:59:59",
			expectedMonth: time.December,
			expectedDay:  31,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			input := MustParse(tt.input)
			result := input.EndOfHalf()

			// 验证月份
			if result.Month() != tt.expectedMonth {
				t.Errorf("EndOfHalf() month = %v, want %v", result.Month(), tt.expectedMonth)
			}

			// 验证日期
			if result.Day() != tt.expectedDay {
				t.Errorf("EndOfHalf() day = %d, want %d", result.Day(), tt.expectedDay)
			}

			// 验证时间为一天的最后时刻
			if result.Hour() != 23 || result.Minute() != 59 || result.Second() != 59 {
				t.Errorf("EndOfHalf() time = %d:%d:%d, want 23:59:59", result.Hour(), result.Minute(), result.Second())
			}

			// 验证纳秒数
			if result.Nanosecond() != 999999999 {
				t.Errorf("EndOfHalf() nanoseconds = %d, want 999999999", result.Nanosecond())
			}
		})
	}
}

func TestEndOfHalf_ConfigPreserved(t *testing.T) {
	customConfig := &Config{
		WeekStartDay:  time.Sunday,
		TimeLocation:  time.UTC,
		TimeFormats:   []string{"2006-01-02"},
		Monotonic:     time.Now(),
	}

	testTime := &Time{
		Time:   time.Date(2024, 3, 15, 14, 30, 45, 0, time.UTC),
		Config: customConfig,
	}

	result := testTime.EndOfHalf()

	if result.Config != customConfig {
		t.Error("EndOfHalf() did not preserve Config")
	}

	if result.WeekStartDay != time.Sunday {
		t.Error("EndOfHalf() did not preserve WeekStartDay")
	}

	if result.Location() != time.UTC {
		t.Error("EndOfHalf() did not preserve Location")
	}
}
