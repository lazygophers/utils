package xtime

import (
	"testing"
	"time"
)

// 方案1: 当前实现（baseline）
func BenchmarkEndOfHour_Current(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_ = EndOfHour()
	}
}

// 方案2: 直接使用 time.Date 构建
func BenchmarkEndOfHour_DirectDate(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		now := time.Now()
		result := time.Date(now.Year(), now.Month(), now.Day(), now.Hour(), 59, 59, 999999999, now.Location())
		_ = With(result)
	}
}

// 方案3: 预先计算常量
func BenchmarkEndOfHour_PreComputed(b *testing.B) {
	b.ReportAllocs()
	const endMinute = 59
	const endSecond = 59
	const endNano = 999999999

	for i := 0; i < b.N; i++ {
		now := time.Now()
		result := time.Date(now.Year(), now.Month(), now.Day(), now.Hour(), endMinute, endSecond, endNano, now.Location())
		_ = With(result)
	}
}

// 方案4: 使用 Truncate 后加 1 小时减 1 纳秒
func BenchmarkEndOfHour_Truncate(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		now := time.Now()
		truncated := now.Truncate(time.Hour)
		result := truncated.Add(time.Hour - time.Nanosecond)
		_ = With(result)
	}
}

// 方案5: 使用 Add 替代部分 Date 调用
func BenchmarkEndOfHour_AddVersion(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		now := time.Now()
		truncated := now.Truncate(time.Hour)
		result := truncated.Add(59*time.Minute + 59*time.Second + 999999999*time.Nanosecond)
		_ = With(result)
	}
}

// 方案6: 内联 With 逻辑
func BenchmarkEndOfHour_InlineWith(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		now := time.Now()
		result := time.Date(now.Year(), now.Month(), now.Day(), now.Hour(), 59, 59, 999999999, now.Location())
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
func BenchmarkEndOfHour_SingleTimeNow(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		now := time.Now()
		result := time.Date(now.Year(), now.Month(), now.Day(), now.Hour(), 59, 59, 999999999, now.Location())
		_ = &Time{
			Time: result,
			Config: &Config{
				WeekStartDay:  time.Monday,
				TimeLocation: time.Local,
				TimeFormats:  []string{},
				Monotonic:    now,
			},
		}
	}
}

// 方案8: 使用全局 Config
func BenchmarkEndOfHour_GlobalConfig(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		now := time.Now()
		result := time.Date(now.Year(), now.Month(), now.Day(), now.Hour(), 59, 59, 999999999, now.Location())
		_ = &Time{
			Time:   result,
			Config: BeginningOfHourConfig,
		}
	}
}

// 方案9: 零分配版本（直接返回 Time，不分配 Config）
func BenchmarkEndOfHour_ZeroAlloc(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		now := time.Now()
		result := time.Date(now.Year(), now.Month(), now.Day(), now.Hour(), 59, 59, 999999999, now.Location())
		_ = &Time{Time: result}
	}
}

// 方案10: 复用 BeginningOfHour 逻辑
func BenchmarkEndOfHour_ReuseBeginning(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		now := time.Now()
		y, m, d := now.Date()
		beginning := time.Date(y, m, d, now.Hour(), 0, 0, 0, now.Location())
		result := beginning.Add(time.Hour - time.Nanosecond)
		_ = With(result)
	}
}

// 方案11: 使用 BeginningOfHour() 函数
func BenchmarkEndOfHour_CallBeginningOfHour(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		now := time.Now()
		beginning := With(now).BeginningOfHour()
		result := beginning.Time.Add(time.Hour - time.Nanosecond)
		_ = With(result)
	}
}

// 方案12: 内联 BeginningOfHour 逻辑 + Add
func BenchmarkEndOfHour_InlineBeginningAdd(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		now := time.Now()
		y, m, d := now.Date()
		beginning := time.Date(y, m, d, now.Hour(), 0, 0, 0, now.Location())
		result := beginning.Add(time.Hour - time.Nanosecond)
		_ = &Time{
			Time:   result,
			Config: BeginningOfHourConfig,
		}
	}
}

// 方案13: 预计算 Hour - 1ns 常量
func BenchmarkEndOfHour_PreComputedHourMinusNs(b *testing.B) {
	b.ReportAllocs()
	const hourMinusNs = time.Hour - time.Nanosecond

	for i := 0; i < b.N; i++ {
		now := time.Now()
		truncated := now.Truncate(time.Hour)
		result := truncated.Add(hourMinusNs)
		_ = &Time{
			Time:   result,
			Config: BeginningOfHourConfig,
		}
	}
}

// 方案14: 使用 Truncate + 全局 Config
func BenchmarkEndOfHour_TruncateWithGlobalConfig(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		now := time.Now()
		truncated := now.Truncate(time.Hour)
		result := truncated.Add(time.Hour - time.Nanosecond)
		_ = &Time{
			Time:   result,
			Config: BeginningOfHourConfig,
		}
	}
}

// 方案15: 完全内联版本
func BenchmarkEndOfHour_FullyInline(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		now := time.Now()
		result := time.Date(now.Year(), now.Month(), now.Day(), now.Hour(), 59, 59, 999999999, now.Location())
		_ = &Time{
			Time: result,
			Config: &Config{
				WeekStartDay:  time.Monday,
				TimeLocation: time.Local,
				TimeFormats:  []string{},
			},
		}
	}
}
