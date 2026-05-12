package xtime

import (
	"testing"
	"time"
)

// 保存原始实现用于性能对比
func (p *Time) EndOfHalf_Original() *Time {
	return With(p.BeginningOfHalf().AddDate(0, 6, 0).Add(-time.Nanosecond))
}

func BenchmarkEndOfHalf_Original(b *testing.B) {
	b.ReportAllocs()
	t := MustParse("2024-03-15 14:30:45")
	for i := 0; i < b.N; i++ {
		_ = t.EndOfHalf_Original()
	}
}

func BenchmarkEndOfHalf_Optimized(b *testing.B) {
	b.ReportAllocs()
	t := MustParse("2024-03-15 14:30:45")
	for i := 0; i < b.N; i++ {
		_ = t.EndOfHalf()
	}
}

func BenchmarkEndOfHalf_Original_H2(b *testing.B) {
	b.ReportAllocs()
	t := MustParse("2024-09-15 14:30:45")
	for i := 0; i < b.N; i++ {
		_ = t.EndOfHalf_Original()
	}
}

func BenchmarkEndOfHalf_Optimized_H2(b *testing.B) {
	b.ReportAllocs()
	t := MustParse("2024-09-15 14:30:45")
	for i := 0; i < b.N; i++ {
		_ = t.EndOfHalf()
	}
}
