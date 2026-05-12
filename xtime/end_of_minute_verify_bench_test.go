package xtime

import (
	"testing"
	"time"
)

// 验证新实现 vs 旧实现性能
func BenchmarkEndOfMinute_NewImplementation(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_ = EndOfMinute()
	}
}

// 旧实现（用于对比）
func BenchmarkEndOfMinute_OldImplementation(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_ = With(time.Now()).EndOfMinute()
	}
}

// 使用固定时间测试（避免 time.Now() 开销）
func BenchmarkEndOfMinute_FixedTime(b *testing.B) {
	b.ReportAllocs()
	testTime := time.Date(2024, 6, 15, 14, 30, 45, 123456789, time.Local)
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		now := testTime
		result := now.Truncate(time.Minute).Add(time.Minute - time.Nanosecond)
		_ = &Time{
			Time:   result,
			Config: &Config{
				WeekStartDay:  time.Monday,
				TimeLocation: time.Local,
				TimeFormats:  nil,
				Monotonic:    now,
			},
		}
	}
}

// 旧实现（固定时间）
func BenchmarkEndOfMinute_OldFixedTime(b *testing.B) {
	b.ReportAllocs()
	testTime := time.Date(2024, 6, 15, 14, 30, 45, 123456789, time.Local)
	t := With(testTime)
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_ = t.EndOfMinute()
	}
}
