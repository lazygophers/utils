package xtime

import (
	"testing"
	"time"
)

// TestEndOfWeek_Correctness 验证 EndOfWeek 正确性
func TestEndOfWeek_Correctness(t *testing.T) {
	tests := []struct {
		name     string
		input    time.Time
		expected time.Weekday // 期望结果为周六（周日为周起始）
	}{
		{
			name:     "周一",
			input:    time.Date(2024, 1, 15, 12, 0, 0, 0, time.Local), // 2024-01-15 周一
			expected: time.Saturday,
		},
		{
			name:     "周三",
			input:    time.Date(2024, 1, 17, 12, 0, 0, 0, time.Local), // 2024-01-17 周三
			expected: time.Saturday,
		},
		{
			name:     "周六",
			input:    time.Date(2024, 1, 20, 12, 0, 0, 0, time.Local), // 2024-01-20 周六
			expected: time.Saturday,
		},
		{
			name:     "跨月",
			input:    time.Date(2024, 1, 29, 12, 0, 0, 0, time.Local), // 2024-01-29 周一
			expected: time.Saturday,
		},
		{
			name:     "年底",
			input:    time.Date(2024, 12, 30, 12, 0, 0, 0, time.Local), // 2024-12-30 周一
			expected: time.Saturday,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			bt := &Time{Time: tt.input}
			result := bt.EndOfWeek()

			// 验证结果是周日
			if result.Weekday() != tt.expected {
				t.Errorf("期望星期 %v，得到 %v", tt.expected, result.Weekday())
			}

			// 验证时间是 23:59:59.999999999
			hour, min, sec := result.Clock()
			nsec := result.Nanosecond()
			if hour != 23 || min != 59 || sec != 59 || nsec != 999999999 {
				t.Errorf("期望时间 23:59:59.999999999，得到 %02d:%02d:%02d.%09d", hour, min, sec, nsec)
			}
		})
	}
}

// TestEndOfWeek_WithCustomWeekStart 验证自定义周起始日
func TestEndOfWeek_WithCustomWeekStart(t *testing.T) {
	input := time.Date(2024, 1, 15, 12, 0, 0, 0, time.Local) // 2024-01-15 周一
	bt := &Time{
		Time:   input,
		Config: &Config{WeekStartDay: time.Monday},
	}
	result := bt.EndOfWeek()

	// 周一起始，周日结束
	expectedDay := time.Sunday
	if result.Weekday() != expectedDay {
		t.Errorf("期望星期 %v，得到 %v", expectedDay, result.Weekday())
	}
}

// TestEndOfWeek_ConfigPreservation 验证 Config 保留
func TestEndOfWeek_ConfigPreservation(t *testing.T) {
	cfg := &Config{
		WeekStartDay:  time.Monday,
		TimeLocation:  time.UTC,
		TimeFormats:   []string{"2006-01-02"},
		Monotonic:     time.Now(),
	}
	bt := &Time{
		Time:   time.Now(),
		Config: cfg,
	}

	result := bt.EndOfWeek()

	// 验证 Config 被保留
	if result.Config != cfg {
		t.Error("Config 未被保留")
	}

	if result.Config.WeekStartDay != time.Monday {
		t.Error("Config.WeekStartDay 未被保留")
	}
}

// Benchmark_EndOfWeek_Optimized 优化后的基准测试
func Benchmark_EndOfWeek_Optimized(b *testing.B) {
	t := Now()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_ = t.EndOfWeek()
	}
}

// Benchmark_EndOfWeek_Optimized_Small 小数据集
func Benchmark_EndOfWeek_Optimized_Small(b *testing.B) {
	t := time.Date(2024, 1, 15, 12, 30, 45, 123456789, time.Local)
	bt := &Time{Time: t}
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_ = bt.EndOfWeek()
	}
}

// Benchmark_EndOfWeek_Optimized_Medium 中等数据集
func Benchmark_EndOfWeek_Optimized_Medium(b *testing.B) {
	t := time.Date(2024, 6, 15, 12, 30, 45, 123456789, time.Local)
	bt := &Time{Time: t}
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_ = bt.EndOfWeek()
	}
}

// Benchmark_EndOfWeek_Optimized_Large 大数据集
func Benchmark_EndOfWeek_Optimized_Large(b *testing.B) {
	t := time.Date(2024, 12, 31, 23, 59, 59, 999999999, time.Local)
	bt := &Time{Time: t}
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_ = bt.EndOfWeek()
	}
}

// Benchmark_EndOfWeek_Optimized_Parallel 并发测试
func Benchmark_EndOfWeek_Optimized_Parallel(b *testing.B) {
	t := Now()
	b.ResetTimer()

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			_ = t.EndOfWeek()
		}
	})
}

// Benchmark_EndOfWeek_Optimized_WithConfig 带 Config
func Benchmark_EndOfWeek_Optimized_WithConfig(b *testing.B) {
	t := Now()
	t.Config = &Config{WeekStartDay: time.Monday}
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_ = t.EndOfWeek()
	}
}
