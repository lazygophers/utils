package xtime

import (
	"testing"
	"time"
)

// Baseline: 当前实现
func BenchmarkNowCalendar_Current(b *testing.B) {
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_ = NowCalendar()
	}
}

// 方案1: 内联 time.Now()
func BenchmarkNowCalendar_InlineTime(b *testing.B) {
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		t := time.Now()
		_ = NewCalendar(t)
	}
}

// 方案2: 延迟计算（仅创建基础结构）
func BenchmarkNowCalendar_Lazy(b *testing.B) {
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_ = newCalendarLazy(time.Now())
	}
}

// 方案3: 缓存 Zodiac 计算（相同年份复用）
func BenchmarkNowCalendar_CachedZodiac(b *testing.B) {
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_ = newCalendarCachedZodiac(time.Now())
	}
}

// 方案4: 简化 Season 计算（移除节气查询）
func BenchmarkNowCalendar_SimpleSeason(b *testing.B) {
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_ = newCalendarSimpleSeason(time.Now())
	}
}

// 方案5: 完全简化（仅基础信息）
func BenchmarkNowCalendar_Minimal(b *testing.B) {
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_ = newCalendarMinimal(time.Now())
	}
}

// 方案6: 预计算常量（优化数组查找）
func BenchmarkNowCalendar_Prealloc(b *testing.B) {
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_ = newCalendarPrealloc(time.Now())
	}
}

// 对比基准：直接创建 Calendar（不含 time.Now()）
func BenchmarkNewCalendar_Only(b *testing.B) {
	t := time.Now()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_ = NewCalendar(t)
	}
}

// 对比基准：仅 time.Now()
func BenchmarkTimeNow_Only(b *testing.B) {
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_ = time.Now()
	}
}

// 对比基准：仅 Lunar 计算
func BenchmarkWithLunar_Only(b *testing.B) {
	t := time.Now()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_ = WithLunar(t)
	}
}

// 对比基准：仅节气查询
func BenchmarkNextSolarterm_Only(b *testing.B) {
	t := time.Now()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_ = NextSolarterm(t)
	}
}
