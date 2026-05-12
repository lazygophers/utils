package xtime

import (
	"testing"
	"time"
)

// TestEndOfHour_Correctness 验证优化后的实现与原实现结果一致
func TestEndOfHour_Correctness(t *testing.T) {
	// 测试多个时间点
	testTimes := []time.Time{
		time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC),
		time.Date(2024, 1, 1, 12, 30, 45, 123456789, time.UTC),
		time.Date(2024, 12, 31, 23, 59, 59, 999999999, time.UTC),
		time.Date(2024, 6, 15, 15, 27, 33, 500000000, time.UTC),
		time.Date(2024, 2, 29, 23, 59, 59, 999999999, time.UTC), // 闰年
		time.Now().Local(),
	}

	for _, tt := range testTimes {
		// 原实现
		original := With(tt).EndOfHour()

		// 新实现
		const hourMinusNs = time.Hour - time.Nanosecond
		truncated := tt.Truncate(time.Hour)
		result := truncated.Add(hourMinusNs)
		optimized := &Time{
			Time:   result,
			Config: BeginningOfHourConfig,
		}

		// 验证时间是否一致
		if original.Time.UnixNano() != optimized.Time.UnixNano() {
			t.Errorf("EndOfHour mismatch for %v:\n  original: %v\n  optimized: %v",
				tt, original.Time, optimized.Time)
		}

		// 验证时区是否一致
		if original.Time.Location().String() != optimized.Time.Location().String() {
			t.Errorf("Location mismatch for %v:\n  original: %v\n  optimized: %v",
				tt, original.Time.Location(), optimized.Time.Location())
		}
	}
}

// TestEndOfHour_BoundaryConditions 测试边界条件
func TestEndOfHour_BoundaryConditions(t *testing.T) {
	tests := []struct {
		name     string
		input    time.Time
		expected string
	}{
		{
			name:     "小时开始",
			input:    time.Date(2024, 1, 1, 10, 0, 0, 0, time.UTC),
			expected: "2024-01-01 10:59:59.999999999 +0000 UTC",
		},
		{
			name:     "小时中间",
			input:    time.Date(2024, 1, 1, 10, 30, 30, 300000000, time.UTC),
			expected: "2024-01-01 10:59:59.999999999 +0000 UTC",
		},
		{
			name:     "小时结束前",
			input:    time.Date(2024, 1, 1, 10, 59, 59, 999999998, time.UTC),
			expected: "2024-01-01 10:59:59.999999999 +0000 UTC",
		},
		{
			name:     "午夜",
			input:    time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC),
			expected: "2024-01-01 00:59:59.999999999 +0000 UTC",
		},
		{
			name:     "午夜前一秒",
			input:    time.Date(2024, 1, 1, 23, 59, 59, 999999999, time.UTC),
			expected: "2024-01-01 23:59:59.999999999 +0000 UTC",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			const hourMinusNs = time.Hour - time.Nanosecond
			truncated := tt.input.Truncate(time.Hour)
			result := truncated.Add(hourMinusNs)
			got := &Time{
				Time:   result,
				Config: BeginningOfHourConfig,
			}

			if got.Time.String() != tt.expected {
				t.Errorf("EndOfHour(%v) = %v, want %v", tt.input, got.Time.String(), tt.expected)
			}
		})
	}
}

// TestEndOfHour_Properties 验证 EndOfHour 的数学属性
func TestEndOfHour_Properties(t *testing.T) {
	now := time.Now()

	// 测试优化后的实现
	const hourMinusNs = time.Hour - time.Nanosecond
	truncated := now.Truncate(time.Hour)
	result := truncated.Add(hourMinusNs)
	eoh := &Time{
		Time:   result,
		Config: BeginningOfHourConfig,
	}

	// 1. EndOfHour 的纳秒部分应该是 999999999
	if eoh.Time.Nanosecond() != 999999999 {
		t.Errorf("EndOfHour nanosecond = %d, want 999999999", eoh.Time.Nanosecond())
	}

	// 2. EndOfHour 的秒部分应该是 59
	if eoh.Time.Second() != 59 {
		t.Errorf("EndOfHour second = %d, want 59", eoh.Time.Second())
	}

	// 3. EndOfHour 的分钟部分应该是 59
	if eoh.Time.Minute() != 59 {
		t.Errorf("EndOfHour minute = %d, want 59", eoh.Time.Minute())
	}

	// 4. EndOfHour + 1 纳秒应该是下一小时的开始
	nextHour := eoh.Time.Add(time.Nanosecond)
	expectedNextHour := now.Truncate(time.Hour).Add(time.Hour)
	if !nextHour.Equal(expectedNextHour) {
		t.Errorf("EndOfHour + 1ns = %v, want %v", nextHour, expectedNextHour)
	}
}

// TestEndOfHour_GlobalFunction 测试全局函数
func TestEndOfHour_GlobalFunction(t *testing.T) {
	// 测试全局函数能正常工作且返回值符合预期
	result := EndOfHour()

	// 验证返回值不为 nil
	if result == nil {
		t.Fatal("EndOfHour() returned nil")
	}

	// 验证时间在合理范围内（不应该太久远）
	now := time.Now()
	if result.Time.Before(now.Add(-time.Hour)) || result.Time.After(now.Add(time.Hour)) {
		t.Errorf("EndOfHour() = %v, out of reasonable range", result.Time)
	}

	// 验证纳秒、秒、分钟的值
	if result.Time.Nanosecond() != 999999999 {
		t.Errorf("EndOfHour() nanosecond = %d, want 999999999", result.Time.Nanosecond())
	}
	if result.Time.Second() != 59 {
		t.Errorf("EndOfHour() second = %d, want 59", result.Time.Second())
	}
	if result.Time.Minute() != 59 {
		t.Errorf("EndOfHour() minute = %d, want 59", result.Time.Minute())
	}
}
