package xtime

import (
	"testing"
	"time"
)

// TestBeginningOfHourOptimization 验证优化方案的正确性
func TestBeginningOfHourOptimization(t *testing.T) {
	testCases := []struct {
		name string
		time time.Time
	}{
		{"2024年6月15日 14:30:45", time.Date(2024, 6, 15, 14, 30, 45, 123456789, time.Local)},
		{"2024年1月1日 00:00:00", time.Date(2024, 1, 1, 0, 0, 0, 0, time.Local)},
		{"2024年12月31日 23:59:59", time.Date(2024, 12, 31, 23, 59, 59, 999999999, time.Local)},
		{"2024年6月15日 00:30:00", time.Date(2024, 6, 15, 0, 30, 0, 0, time.Local)},
		{"2024年6月15日 23:00:00", time.Date(2024, 6, 15, 23, 0, 0, 0, time.Local)},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// 原始实现
			original := With(tc.time).BeginningOfHour()

			// 方案1: Truncate + nil Config
			result1 := &Time{Time: tc.time.Truncate(time.Hour), Config: nil}
			if !result1.Time.Equal(original.Time) {
				t.Errorf("TruncateNil: expected %v, got %v", original.Time, result1.Time)
			}

			// 方案2: Truncate + GlobalConfig
			result2 := &Time{Time: tc.time.Truncate(time.Hour), Config: BeginningOfHourConfig}
			if !result2.Time.Equal(original.Time) {
				t.Errorf("GlobalConfig: expected %v, got %v", original.Time, result2.Time)
			}

			// 方案3: Date 方法
			y, m, d := tc.time.Date()
			h := tc.time.Hour()
			result3 := &Time{Time: time.Date(y, m, d, h, 0, 0, 0, tc.time.Location()), Config: nil}
			if !result3.Time.Equal(original.Time) {
				t.Errorf("Date: expected %v, got %v", original.Time, result3.Time)
			}

			// 验证分钟、秒、纳秒都归零
			if result1.Minute() != 0 || result1.Second() != 0 || result1.Nanosecond() != 0 {
				t.Errorf("Result not truncated to hour: %v", result1.Time)
			}
		})
	}
}

// TestBeginningOfHourTruncateBehavior 验证 Truncate 的行为
func TestBeginningOfHourTruncateBehavior(t *testing.T) {
	now := time.Date(2024, 6, 15, 14, 30, 45, 123456789, time.Local)
	truncated := now.Truncate(time.Hour)

	expected := time.Date(2024, 6, 15, 14, 0, 0, 0, time.Local)
	if !truncated.Equal(expected) {
		t.Errorf("Truncate failed: expected %v, got %v", expected, truncated)
	}

	if truncated.Minute() != 0 || truncated.Second() != 0 || truncated.Nanosecond() != 0 {
		t.Errorf("Truncate did not zero out minute/second/nanosecond: %v", truncated)
	}
}

// TestBeginningOfHourWithCurrentTime 验证使用 time.Now() 的正确性
func TestBeginningOfHourWithCurrentTime(t *testing.T) {
	before := time.Now()
	result := BeginningOfHour()
	after := time.Now()

	// 验证结果的时间戳在合理范围内
	if result.Time.Before(before.Add(-time.Hour)) || result.Time.After(after.Add(time.Hour)) {
		t.Errorf("BeginningOfHour returned unexpected time: %v (between %v and %v)", result.Time, before, after)
	}

	// 验证分钟、秒、纳秒都归零
	if result.Minute() != 0 || result.Second() != 0 || result.Nanosecond() != 0 {
		t.Errorf("BeginningOfHour did not truncate to hour: %v", result.Time)
	}
}
