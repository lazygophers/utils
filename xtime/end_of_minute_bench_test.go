package xtime

import (
	"testing"
	"time"
)

// 方案1: 当前实现（baseline）
func BenchmarkEndOfMinute_Current(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_ = With(time.Now()).EndOfMinute()
	}
}

// 方案2: 直接使用 time.Date 构建
func BenchmarkEndOfMinute_DirectDate(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		now := time.Now()
		result := time.Date(now.Year(), now.Month(), now.Day(), now.Hour(), now.Minute(), 59, 999999999, now.Location())
		_ = With(result)
	}
}

// 方案3: 预先计算常量
func BenchmarkEndOfMinute_PreComputed(b *testing.B) {
	b.ReportAllocs()
	const endSecond = 59
	const endNano = 999999999

	for i := 0; i < b.N; i++ {
		now := time.Now()
		result := time.Date(now.Year(), now.Month(), now.Day(), now.Hour(), now.Minute(), endSecond, endNano, now.Location())
		_ = With(result)
	}
}

// 方案4: 使用 Truncate 后加 1 分钟减 1 纳秒
func BenchmarkEndOfMinute_Truncate(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		now := time.Now()
		truncated := now.Truncate(time.Minute)
		result := truncated.Add(time.Minute - time.Nanosecond)
		_ = With(result)
	}
}

// 方案5: 使用 Add 替代部分 Date 调用
func BenchmarkEndOfMinute_AddVersion(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		now := time.Now()
		truncated := now.Truncate(time.Minute)
		result := truncated.Add(59*time.Second + 999999999*time.Nanosecond)
		_ = With(result)
	}
}

// 方案6: 内联 With 逻辑
func BenchmarkEndOfMinute_InlineWith(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		now := time.Now()
		result := time.Date(now.Year(), now.Month(), now.Day(), now.Hour(), now.Minute(), 59, 999999999, now.Location())
		_ = &Time{
			Time: result,
			Config: &Config{
				WeekStartDay:  time.Monday,
				TimeLocation: time.Local,
				TimeFormats:  []string{},
				Monotonic:    time.Now(),
			},
		}
	}
}

// 方案7: 单次 time.Now() 调用（用于 Monotonic）
func BenchmarkEndOfMinute_SingleTimeNow(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		now := time.Now()
		result := time.Date(now.Year(), now.Month(), now.Day(), now.Hour(), now.Minute(), 59, 999999999, now.Location())
		_ = &Time{
			Time:   result,
			Config: &Config{
				WeekStartDay:  time.Monday,
				TimeLocation: time.Local,
				TimeFormats:  []string{},
				Monotonic:    now,
			},
		}
	}
}

// 方案8: 使用 Unix 时间戳计算
func BenchmarkEndOfMinute_Unix(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		now := time.Now()
		nanos := now.UnixNano()
		minuteNanos := int64(time.Minute)
		aligned := (nanos / minuteNanos) * minuteNanos
		result := time.Unix(0, aligned+minuteNanos-1).In(now.Location())
		_ = With(result)
	}
}

// 方案9: 最简版本（nil TimeFormats，无 Monotonic）
func BenchmarkEndOfMinute_Minimal(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		now := time.Now()
		result := time.Date(now.Year(), now.Month(), now.Day(), now.Hour(), now.Minute(), 59, 999999999, now.Location())
		_ = &Time{
			Time: result,
			Config: &Config{
				WeekStartDay:  time.Monday,
				TimeLocation: time.Local,
				TimeFormats:  nil,
			},
		}
	}
}

// 方案10: 组合优化（单次 time.Now + 预计算常量 + nil TimeFormats）
func BenchmarkEndOfMinute_Combined(b *testing.B) {
	b.ReportAllocs()
	const (
		endSecond    = 59
		endNano      = 999999999
		weekStartDay = time.Monday
	)

	for i := 0; i < b.N; i++ {
		now := time.Now()
		result := time.Date(now.Year(), now.Month(), now.Day(), now.Hour(), now.Minute(), endSecond, endNano, now.Location())
		_ = &Time{
			Time:   result,
			Config: &Config{
				WeekStartDay:  weekStartDay,
				TimeLocation: time.Local,
				TimeFormats:  nil,
				Monotonic:    now,
			},
		}
	}
}

// 方案11: 使用 time.Add 直接计算
func BenchmarkEndOfMinute_AddDirect(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		now := time.Now()
		nextMinute := now.Truncate(time.Minute).Add(time.Minute)
		result := nextMinute.Add(-time.Nanosecond)
		_ = With(result)
	}
}

// 方案12: 优化的 Truncate 版本（单次 time.Now）
func BenchmarkEndOfMinute_OptimizedTruncate(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		now := time.Now()
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

// 方案13: 完全内联版本
func BenchmarkEndOfMinute_FullyInline(b *testing.B) {
	b.ReportAllocs()
	const (
		endSecond    = 59
		endNano      = 999999999
		weekStartDay = time.Monday
	)

	for i := 0; i < b.N; i++ {
		now := time.Now()
		result := time.Date(now.Year(), now.Month(), now.Day(), now.Hour(), now.Minute(), endSecond, endNano, now.Location())
		t := &Time{
			Time:   result,
			Config: &Config{
				WeekStartDay:  weekStartDay,
				TimeLocation: time.Local,
				TimeFormats:  nil,
				Monotonic:    now,
			},
		}
		_ = t
	}
}

// 方案14: 使用 time.Date 但避免重复字段提取
func BenchmarkEndOfMinute_FieldReuse(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		now := time.Now()
		y, m, d := now.Date()
		h, min, _ := now.Clock()
		result := time.Date(y, m, d, h, min, 59, 999999999, now.Location())
		_ = With(result)
	}
}
