package xtime

import (
	"testing"
	"time"
)

// 独立基准测试
func BenchmarkStandalone_Current(b *testing.B) {
	t := With(time.Date(2024, 6, 15, 14, 30, 45, 123456789, time.Local))
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = t.EndOfYear()
	}
}

func BenchmarkStandalone_Optimized(b *testing.B) {
	t := With(time.Date(2024, 6, 15, 14, 30, 45, 123456789, time.Local))
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		year := t.Time.Year()
		loc := t.Time.Location()
		config := t.Config
		if config == nil {
			config = &Config{}
		}
		end := time.Date(year+1, time.January, 0, 23, 59, 59, 999999999, loc)
		_ = &Time{Time: end, Config: config}
	}
}
