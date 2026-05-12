package xtime

import (
	"testing"
	"time"
)

// ========== 全局 BeginningOfHour() 基准测试 ==========

// Baseline: 当前实现
func BenchmarkBeginningOfHour_Current(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_ = BeginningOfHour()
	}
}

// 方案1: Truncate + nil Config
func BenchmarkBeginningOfHour_TruncateNil(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		t := time.Now()
		_ = &Time{Time: t.Truncate(time.Hour), Config: nil}
	}
}

// 方案2: Truncate + 全局共享 Config
func BenchmarkBeginningOfHour_GlobalConfig(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		t := time.Now()
		_ = &Time{Time: t.Truncate(time.Hour), Config: BeginningOfHourConfig}
	}
}

// 方案3: Truncate + 空结构体 Config
func BenchmarkBeginningOfHour_ZeroConfig(b *testing.B) {
	zeroConfig := &Config{}
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		t := time.Now()
		_ = &Time{Time: t.Truncate(time.Hour), Config: zeroConfig}
	}
}

// 方案4: 完整 Date 构建
func BenchmarkBeginningOfHour_Date(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		t := time.Now()
		y, m, d := t.Date()
		h := t.Hour()
		_ = &Time{Time: time.Date(y, m, d, h, 0, 0, 0, t.Location()), Config: nil}
	}
}

// 方案5: Date + 全局 Config
func BenchmarkBeginningOfHour_DateWithConfig(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		t := time.Now()
		y, m, d := t.Date()
		h := t.Hour()
		_ = &Time{Time: time.Date(y, m, d, h, 0, 0, 0, t.Location()), Config: BeginningOfHourConfig}
	}
}

// 方案6: Add + Subtract 方法
func BenchmarkBeginningOfHour_AddSubtract(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		t := time.Now()
		min, sec := t.Minute(), t.Second()
		ns := t.Nanosecond()
		truncated := t.Add(-time.Duration(min)*time.Minute - time.Duration(sec)*time.Second - time.Duration(ns)*time.Nanosecond)
		_ = &Time{Time: truncated, Config: nil}
	}
}

// 方案7: Unix 时间戳方法
func BenchmarkBeginningOfHour_Unix(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		t := time.Now()
		loc := t.Location()
		unix := t.Unix()
		hourSec := int64(t.Hour()) * 3600
		truncatedUnix := unix - (unix % 3600) - hourSec + int64(t.Hour())*3600
		_ = &Time{Time: time.Unix(truncatedUnix, 0).In(loc), Config: nil}
	}
}

// 方案8: 预先提取 Location
func BenchmarkBeginningOfHour_PreallocLocation(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		t := time.Now()
		loc := t.Location()
		y, m, d := t.Date()
		h := t.Hour()
		_ = &Time{Time: time.Date(y, m, d, h, 0, 0, 0, loc), Config: nil}
	}
}

// 方案9: 完整参数提取 + Config 复用
func BenchmarkBeginningOfHour_FullExtract(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		t := time.Now()
		y, m, d := t.Date()
		h := t.Hour()
		loc := t.Location()
		_ = &Time{Time: time.Date(y, m, d, h, 0, 0, 0, loc), Config: BeginningOfHourConfig}
	}
}

// 方案10: 简化版 Truncate
func BenchmarkBeginningOfHour_Minimal(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		t := time.Now()
		_ = &Time{Time: t.Truncate(time.Hour), Config: BeginningOfHourConfig}
	}
}

// 方案11: 优化版 With（避免重复创建 Config）
func BenchmarkBeginningOfHour_OptimizedWith(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		t := time.Now()
		truncated := t.Truncate(time.Hour)
		_ = &Time{Time: truncated, Config: BeginningOfHourConfig}
	}
}

// 方案12: 使用 Truncate + 嵌入式 Time
func BenchmarkBeginningOfHour_EmbeddedTime(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		t := time.Now()
		result := Time{
			Time:   t.Truncate(time.Hour),
			Config: BeginningOfHourConfig,
		}
		_ = &result
	}
}

// 方案13: 分离 Location 和 Date
func BenchmarkBeginningOfHour_SeparatedLocation(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		t := time.Now()
		loc := t.Location()
		hour := t.Hour()
		y, m, d := t.Date()
		_ = &Time{Time: time.Date(y, m, d, hour, 0, 0, 0, loc), Config: nil}
	}
}

// 方案14: 使用 time.Now().Truncate 直接内联
func BenchmarkBeginningOfHour_InlinedTruncate(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_ = &Time{Time: time.Now().Truncate(time.Hour), Config: BeginningOfHourConfig}
	}
}

// 方案15: 零配置优化（最小内存分配）
func BenchmarkBeginningOfHour_ZeroAlloc(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		t := time.Now()
		_ = &Time{Time: t.Truncate(time.Hour), Config: nil}
	}
}
