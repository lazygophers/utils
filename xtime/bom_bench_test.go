package xtime

import (
	"testing"
	"time"
)

// 方案1: Baseline - 当前实现
func BenchmarkBOM_Current(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_ = With(time.Now()).BeginningOfMinute()
	}
}

// 方案2: Truncate + nil Config
func BenchmarkBOM_TruncateNil(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		t := time.Now()
		_ = &Time{Time: t.Truncate(time.Minute), Config: nil}
	}
}

// 方案3: 全局 Config
var bomGlobalConfig = &Config{
	WeekStartDay:  time.Monday,
	TimeLocation: time.Local,
	TimeFormats:  []string{},
}

func BenchmarkBOM_GlobalConfig(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		t := time.Now()
		_ = &Time{Time: t.Truncate(time.Minute), Config: bomGlobalConfig}
	}
}

// 方案4: 使用 Date
func BenchmarkBOM_Date(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		t := time.Now()
		_ = &Time{
			Time:   time.Date(t.Year(), t.Month(), t.Day(), t.Hour(), t.Minute(), 0, 0, t.Location()),
			Config: &Config{},
		}
	}
}

// 方案5: Add 负偏移
func BenchmarkBOM_AddSubtract(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		t := time.Now()
		sec := t.Second()
		nanosec := t.Nanosecond()
		_ = With(t.Add(-time.Duration(sec)*time.Second - time.Duration(nanosec)))
	}
}

// 方案6: Unix 时间戳
func BenchmarkBOM_Unix(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		t := time.Now()
		unix := t.Unix()
		truncatedUnix := unix - (unix % 60)
		_ = With(time.Unix(truncatedUnix, 0).In(t.Location()))
	}
}

// 方案7: 预分配 Location
func BenchmarkBOM_PreallocLocation(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		t := time.Now()
		loc := t.Location()
		_ = &Time{
			Time: time.Date(t.Year(), t.Month(), t.Day(), t.Hour(), t.Minute(), 0, 0, loc),
			Config: &Config{
				WeekStartDay:  time.Monday,
				TimeLocation: loc,
				TimeFormats:  []string{},
			},
		}
	}
}

// 方案8: 空 Config
var bomZeroConfig = &Config{}

func BenchmarkBOM_ZeroConfig(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		t := time.Now()
		_ = &Time{Time: t.Truncate(time.Minute), Config: bomZeroConfig}
	}
}

// 方案9: 完整参数提取
func BenchmarkBOM_FullExtract(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		t := time.Now()
		y, m, d := t.Date()
		h, _, _ := t.Clock()
		_ = &Time{
			Time: time.Date(y, m, d, h, t.Minute(), 0, 0, t.Location()),
			Config: &Config{},
		}
	}
}

// 方案10: 优化版
var bomOptimizedConfig = &Config{
	WeekStartDay:  time.Monday,
	TimeLocation: time.Local,
	TimeFormats:  []string{},
}

func BenchmarkBOM_Optimized(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		t := time.Now()
		_ = &Time{Time: t.Truncate(time.Minute), Config: bomOptimizedConfig}
	}
}

// 方案11: 最小化版本
func BenchmarkBOM_Minimal(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		t := time.Now()
		_ = &Time{Time: t.Truncate(time.Minute), Config: nil}
	}
}

// 并行测试
func BenchmarkBOM_Current_Parallel(b *testing.B) {
	b.ReportAllocs()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			_ = With(time.Now()).BeginningOfMinute()
		}
	})
}

func BenchmarkBOM_Optimized_Parallel(b *testing.B) {
	b.ReportAllocs()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			t := time.Now()
			_ = &Time{Time: t.Truncate(time.Minute), Config: bomOptimizedConfig}
		}
	})
}
