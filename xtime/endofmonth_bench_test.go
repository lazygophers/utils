package xtime

import (
	"testing"
	"time"
)

func Benchmark_EndOfMonth_Original(b *testing.B) {
	t := Now()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = With(t.BeginningOfMonth().AddDate(0, 1, 0).Add(-time.Nanosecond))
	}
}

func Benchmark_EndOfMonth_Opt1(b *testing.B) {
	t := Now()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		bom := t.BeginningOfMonth()
		eom := bom.AddDate(0, 1, 0).Add(-time.Nanosecond)
		cfg := bom.Config
		if cfg == nil {
			cfg = &Config{}
		}
		_ = &Time{Time: eom, Config: cfg}
	}
}

func Benchmark_EndOfMonth_Opt2(b *testing.B) {
	t := Now()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		bom := t.BeginningOfMonth()
		nextMonth := bom.AddDate(0, 1, 0)
		eom := nextMonth.Add(-time.Nanosecond)
		cfg := bom.Config
		if cfg == nil {
			cfg = &Config{}
		}
		_ = &Time{Time: eom, Config: cfg}
	}
}

func Benchmark_EndOfMonth_Opt3(b *testing.B) {
	t := Now()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		bom := t.BeginningOfMonth()
		eom := bom.AddDate(0, 1, 0).Add(-time.Nanosecond)
		_ = &Time{Time: eom, Config: t.Config}
	}
}

func Benchmark_EndOfMonth_Opt4(b *testing.B) {
	t := Now()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		bom := t.BeginningOfMonth()
		eom := bom.AddDate(0, 1, 0).Add(-time.Nanosecond)
		_ = &Time{Time: eom, Config: bom.Config}
	}
}

func Benchmark_EndOfMonth_Opt5(b *testing.B) {
	t := Now()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		bom := t.BeginningOfMonth()
		cfg := bom.Config
		if cfg == nil {
			cfg = &Config{}
		}
		_ = &Time{Time: bom.AddDate(0, 1, 0).Add(-time.Nanosecond), Config: cfg}
	}
}

func Benchmark_EndOfMonth_Opt6(b *testing.B) {
	t := Now()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		year, month, _ := t.Date()
		nextMonth := month + 1
		nextYear := year
		if nextMonth > time.December {
			nextMonth = time.January
			nextYear = year + 1
		}
		eomTime := time.Date(nextYear, nextMonth, 1, 0, 0, 0, 0, t.Location()).Add(-time.Nanosecond)
		_ = &Time{Time: eomTime, Config: t.Config}
	}
}

func Benchmark_EndOfMonth_Opt7(b *testing.B) {
	t := Now()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		year, month, _ := t.Date()
		nextMonth := month + 1
		nextYear := year
		if nextMonth > time.December {
			nextMonth = time.January
			nextYear = year + 1
		}
		eomTime := time.Date(nextYear, nextMonth, 1, 0, 0, 0, 0, t.Location()).Add(-time.Nanosecond)
		_ = &Time{Time: eomTime, Config: t.Config}
	}
}

func Benchmark_EndOfMonth_Opt8(b *testing.B) {
	t := Now()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		bom := t.BeginningOfMonth()
		eom := bom.AddDate(0, 1, -1)
		eom = time.Date(eom.Year(), eom.Month(), eom.Day(), 23, 59, 59, 999999999, eom.Location())
		cfg := bom.Config
		if cfg == nil {
			cfg = &Config{}
		}
		_ = &Time{Time: eom, Config: cfg}
	}
}

func Benchmark_EndOfMonth_Opt9(b *testing.B) {
	t := Now()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		year, month, _ := t.Date()
		eomTime := time.Date(year, month+1, 0, 23, 59, 59, 999999999, t.Location())
		_ = &Time{Time: eomTime, Config: t.Config}
	}
}

func Benchmark_EndOfMonth_Opt10(b *testing.B) {
	t := Now()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		year, month, _ := t.Date()
		firstOfMonth := time.Date(year, month, 1, 0, 0, 0, 0, t.Location())
		eomTime := firstOfMonth.AddDate(0, 1, 0).Add(-time.Nanosecond)
		_ = &Time{Time: eomTime, Config: t.Config}
	}
}

func Benchmark_EndOfMonth_Opt11(b *testing.B) {
	t := Now()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		year, month, _ := t.Date()
		eomTime := time.Date(year, month+1, 0, 23, 59, 59, 999999999, t.Location())
		_ = &Time{Time: eomTime, Config: t.Config}
	}
}

func Benchmark_EndOfMonth_Opt12(b *testing.B) {
	t := Now()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = &Time{
			Time:   time.Date(t.Year(), t.Month()+1, 0, 23, 59, 59, 999999999, t.Location()),
			Config: t.Config,
		}
	}
}
