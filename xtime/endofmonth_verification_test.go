package xtime

import (
	"testing"
	"time"
)

// TestEndOfMonth_Correctness 验证 EndOfMonth 的正确性
func TestEndOfMonth_Correctness(t *testing.T) {
	tests := []struct {
		name     string
		input    time.Time
		expected time.Time
	}{
		{
			name:     "2024年1月",
			input:    time.Date(2024, 1, 15, 12, 30, 45, 123456789, time.Local),
			expected: time.Date(2024, 1, 31, 23, 59, 59, 999999999, time.Local),
		},
		{
			name:     "2024年2月（闰年）",
			input:    time.Date(2024, 2, 10, 0, 0, 0, 0, time.Local),
			expected: time.Date(2024, 2, 29, 23, 59, 59, 999999999, time.Local),
		},
		{
			name:     "2023年2月（非闰年）",
			input:    time.Date(2023, 2, 15, 12, 0, 0, 0, time.Local),
			expected: time.Date(2023, 2, 28, 23, 59, 59, 999999999, time.Local),
		},
		{
			name:     "2024年12月（跨年）",
			input:    time.Date(2024, 12, 25, 10, 0, 0, 0, time.Local),
			expected: time.Date(2024, 12, 31, 23, 59, 59, 999999999, time.Local),
		},
		{
			name:     "月末当天",
			input:    time.Date(2024, 1, 31, 23, 59, 59, 999999999, time.Local),
			expected: time.Date(2024, 1, 31, 23, 59, 59, 999999999, time.Local),
		},
		{
			name:     "月初当天",
			input:    time.Date(2024, 1, 1, 0, 0, 0, 0, time.Local),
			expected: time.Date(2024, 1, 31, 23, 59, 59, 999999999, time.Local),
		},
		{
			name:     "4月（30天）",
			input:    time.Date(2024, 4, 15, 12, 0, 0, 0, time.Local),
			expected: time.Date(2024, 4, 30, 23, 59, 59, 999999999, time.Local),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			xTime := With(tt.input)
			result := xTime.EndOfMonth()

			if !result.Time.Equal(tt.expected) {
				t.Errorf("EndOfMonth() = %v, want %v", result.Time, tt.expected)
			}

			// 验证时间是当月最后一刻
			expectedNextDay := tt.expected.Add(1)
			if expectedNextDay.Day() != 1 {
				t.Errorf("不是月末最后一刻: %v", result.Time)
			}

			// 验证纳秒是 999999999
			if result.Time.Nanosecond() != 999999999 {
				t.Errorf("Nanosecond = %d, want 999999999", result.Time.Nanosecond())
			}

			// 验证时分秒是 23:59:59
			h, m, s := result.Time.Clock()
			if h != 23 || m != 59 || s != 59 {
				t.Errorf("Clock = %d:%d:%d, want 23:59:59", h, m, s)
			}
		})
	}
}

// TestEndOfMonth_Consistency 验证不同方式调用结果一致
func TestEndOfMonth_Consistency(t *testing.T) {
	now := time.Now()

	// 使用新实现
	newImpl := With(now).EndOfMonth()

	// 使用原始逻辑计算
	oldImpl := With(With(now).BeginningOfMonth().AddDate(0, 1, 0).Add(-time.Nanosecond))

	if !newImpl.Time.Equal(oldImpl.Time) {
		t.Errorf("新旧实现不一致: new=%v, old=%v", newImpl.Time, oldImpl.Time)
	}
}

// TestEndOfMonth_EdgeCases 边界情况测试
func TestEndOfMonth_EdgeCases(t *testing.T) {
	// 测试不同时区
	locations := []*time.Location{
		time.UTC,
		time.Local,
		time.FixedZone("EST", -5*3600),
		time.FixedZone("JST", 9*3600),
	}

	for _, loc := range locations {
		t.Run(loc.String(), func(t *testing.T) {
			testTime := time.Date(2024, 2, 15, 12, 0, 0, 0, loc)
			result := With(testTime).EndOfMonth()

			// 验证时区正确
			if result.Location().String() != loc.String() {
				t.Errorf("Location = %v, want %v", result.Location(), loc)
			}

			// 验证是2月最后一天（闰年）
			if result.Month() != time.February || result.Day() != 29 {
				t.Errorf("Date = %v, want 2024-02-29", result.Time)
			}
		})
	}
}

// Benchmark_EndOfMonth_Optimized 验证优化后的性能
func Benchmark_EndOfMonth_Optimized(b *testing.B) {
	t := Now()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_ = t.EndOfMonth()
	}
}

// Benchmark_EndOfMonth_Final 最终性能测试（无 Config）
func Benchmark_EndOfMonth_Final(b *testing.B) {
	t := &Time{Time: time.Date(2024, 2, 15, 12, 0, 0, 0, time.Local)}
	b.ResetTimer()
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		_ = t.EndOfMonth()
	}
}
