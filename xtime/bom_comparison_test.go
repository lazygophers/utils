package xtime

import (
	"testing"
	"time"
)

// 原始实现（用于对比）
func benchmarkBOM_Original(b *testing.B) {
	t := Now()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		y, m, _ := t.Date()
		_ = With(time.Date(y, m, 1, 0, 0, 0, 0, t.Location()))
	}
}

// 优化后的实现
func benchmarkBOM_Optimized(b *testing.B) {
	t := Now()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = &Time{
			Time:   time.Date(t.Year(), t.Month(), 1, 0, 0, 0, 0, t.Location()),
			Config: t.Config,
		}
	}
}

func BenchmarkBOM_Comparison_Original(b *testing.B) {
	b.ReportAllocs()
	benchmarkBOM_Original(b)
}

func BenchmarkBOM_Comparison_Optimized(b *testing.B) {
	b.ReportAllocs()
	benchmarkBOM_Optimized(b)
}
