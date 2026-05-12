package xtime

import (
	"testing"
	"time"
)

// BenchmarkBeginningOfMonth_Optimized 优化后的性能
func BenchmarkBeginningOfMonth_Optimized(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_ = BeginningOfMonth()
	}
}

// BenchmarkBeginningOfMonth_Original 原始实现的性能
func BenchmarkBeginningOfMonth_Original(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_ = With(time.Now()).BeginningOfMonth()
	}
}

// BenchmarkBeginningOfMonth_Variant1 V1: 当前实现
func BenchmarkBeginningOfMonth_Variant1(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_ = With(time.Now()).BeginningOfMonth()
	}
}

// BenchmarkBeginningOfMonth_Variant2 V2: 内联逻辑，完整 Config
func BenchmarkBeginningOfMonth_Variant2(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		now := time.Now()
		year, month, _ := now.Date()
		_ = &Time{
			Time: time.Date(year, month, 1, 0, 0, 0, 0, now.Location()),
			Config: &Config{
				WeekStartDay:  time.Monday,
				TimeLocation: time.Local,
				TimeFormats:  []string{},
				Monotonic:    now,
			},
		}
	}
}

// BenchmarkBeginningOfMonth_Variant3 V3: 简化 Config
func BenchmarkBeginningOfMonth_Variant3(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		now := time.Now()
		year, month, _ := now.Date()
		_ = &Time{
			Time: time.Date(year, month, 1, 0, 0, 0, 0, now.Location()),
			Config: &Config{
				TimeLocation: now.Location(),
			},
		}
	}
}

// BenchmarkBeginningOfMonth_Variant4 V4: nil Config
func BenchmarkBeginningOfMonth_Variant4(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		now := time.Now()
		year, month, _ := now.Date()
		_ = &Time{
			Time:   time.Date(year, month, 1, 0, 0, 0, 0, now.Location()),
			Config: nil,
		}
	}
}

// BenchmarkBeginningOfMonth_Variant5 V5: 使用 Year/Month 方法 + nil Config
func BenchmarkBeginningOfMonth_Variant5(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		now := time.Now()
		_ = &Time{
			Time:   time.Date(now.Year(), now.Month(), 1, 0, 0, 0, 0, now.Location()),
			Config: nil,
		}
	}
}

// BenchmarkBeginningOfMonth_Variant6 V6: 最简化（Date + 无 Config）
func BenchmarkBeginningOfMonth_Variant6(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		now := time.Now()
		year, month, _ := now.Date()
		_ = &Time{Time: time.Date(year, month, 1, 0, 0, 0, 0, now.Location())}
	}
}

// BenchmarkBeginningOfMonth_Variant12 V12: 最优方案
func BenchmarkBeginningOfMonth_Variant12(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		now := time.Now()
		_ = &Time{Time: time.Date(now.Year(), now.Month(), 1, 0, 0, 0, 0, now.Location())}
	}
}

