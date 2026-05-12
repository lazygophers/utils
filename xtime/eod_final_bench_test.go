package xtime

import (
	"testing"
	"time"
)

// BenchmarkEndOfDay_Optimized 优化后的实现
func BenchmarkEndOfDay_Optimized(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_ = EndOfDay()
	}
}

// BenchmarkEndOfDay_Baseline 原始实现（用于对比）
func BenchmarkEndOfDay_Baseline(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_ = With(time.Now()).EndOfDay()
	}
}

// BenchmarkEndOfDay_Manual 手动优化版本
func BenchmarkEndOfDay_Manual(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		now := time.Now()
		year, month, day := now.Date()
		eod := time.Date(year, month, day, 23, 59, 59, int(time.Second-time.Nanosecond), now.Location())
		_ = &Time{Time: eod}
	}
}
