package xtime

import (
	"testing"
	"time"
)

func Benchmark_EndOfQuarter_Original(b *testing.B) {
	t := Now()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = With(t.BeginningOfQuarter().AddDate(0, 3, 0).Add(-time.Nanosecond))
	}
}

func Benchmark_EndOfQuarter_Variant1(b *testing.B) {
	t := Now()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		boq := t.BeginningOfQuarter()
		eoq := boq.AddDate(0, 3, 0).Add(-time.Nanosecond)
		_ = &Time{Time: eoq, Config: boq.Config}
	}
}

func Benchmark_EndOfQuarter_Variant2(b *testing.B) {
	t := Now()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		boq := t.BeginningOfQuarter()
		cfg := boq.Config
		if cfg == nil {
			cfg = &Config{}
		}
		_ = &Time{Time: boq.AddDate(0, 3, 0).Add(-time.Nanosecond), Config: cfg}
	}
}

func Benchmark_EndOfQuarter_Variant3(b *testing.B) {
	t := Now()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		boq := t.BeginningOfQuarter()
		_ = &Time{Time: boq.AddDate(0, 3, 0).Add(-time.Nanosecond), Config: t.Config}
	}
}

func Benchmark_EndOfQuarter_Variant4(b *testing.B) {
	t := Now()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		boq := t.BeginningOfQuarter()
		nextQuarter := boq.AddDate(0, 3, 0)
		eoq := nextQuarter.Add(-time.Nanosecond)
		_ = &Time{Time: eoq, Config: t.Config}
	}
}

func Benchmark_EndOfQuarter_Variant5(b *testing.B) {
	t := Now()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		year, month, _ := t.Date()
		quarter := (month-1)/3 + 1
		if quarter > 4 {
			quarter = 4
		}
		nextQuarterMonth := quarter*3 + 1
		nextYear := year
		if nextQuarterMonth > time.December {
			nextQuarterMonth = time.January
			nextYear = year + 1
		}
		eoqTime := time.Date(nextYear, nextQuarterMonth, 1, 0, 0, 0, 0, t.Location()).Add(-time.Nanosecond)
		_ = &Time{Time: eoqTime, Config: t.Config}
	}
}

func Benchmark_EndOfQuarter_Variant6(b *testing.B) {
	t := Now()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		year, month, _ := t.Date()
		quarter := (month-1)/3 + 1
		endQuarterMonth := quarter * 3
		eoqTime := time.Date(year, endQuarterMonth+1, 0, 23, 59, 59, 999999999, t.Location())
		_ = &Time{Time: eoqTime, Config: t.Config}
	}
}

func Benchmark_EndOfQuarter_Variant7(b *testing.B) {
	t := Now()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		loc := t.Location()
		year := t.Year()
		month := int(t.Month())
		quarterStartMonth := ((month - 1) / 3) * 3
		boq := time.Date(year, time.Month(quarterStartMonth+1), 1, 0, 0, 0, 0, loc)
		eoqTime := boq.AddDate(0, 3, 0).Add(-time.Nanosecond)
		_ = &Time{Time: eoqTime, Config: t.Config}
	}
}

func Benchmark_EndOfQuarter_Variant8(b *testing.B) {
	t := Now()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		loc := t.Location()
		year := t.Year()
		month := int(t.Month())
		quarterStartMonth := ((month - 1) / 3) * 3
		nextQuarterMonth := quarterStartMonth + 4
		nextYear := year
		if nextQuarterMonth > 9 {
			nextQuarterMonth = nextQuarterMonth - 12
			nextYear = year + 1
		}
		eoqTime := time.Date(nextYear, time.Month(nextQuarterMonth+1), 1, 0, 0, 0, 0, loc).Add(-time.Nanosecond)
		_ = &Time{Time: eoqTime, Config: t.Config}
	}
}

func Benchmark_EndOfQuarter_Variant9(b *testing.B) {
	t := Now()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		boq := t.BeginningOfQuarter()
		eoq := boq.AddDate(0, 3, -1)
		eoq = time.Date(eoq.Year(), eoq.Month(), eoq.Day(), 23, 59, 59, 999999999, eoq.Location())
		cfg := boq.Config
		if cfg == nil {
			cfg = &Config{}
		}
		_ = &Time{Time: eoq, Config: cfg}
	}
}

func Benchmark_EndOfQuarter_Variant10(b *testing.B) {
	t := Now()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		loc := t.Location()
		year, month, _ := t.Date()
		quarter := (month-1)/3 + 1
		endQuarterMonth := quarter * 3
		eoqTime := time.Date(year, endQuarterMonth+1, 0, 23, 59, 59, 999999999, loc)
		cfg := t.Config
		if cfg == nil {
			cfg = &Config{}
		}
		_ = &Time{Time: eoqTime, Config: cfg}
	}
}

func Benchmark_EndOfQuarter_Variant11(b *testing.B) {
	t := Now()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		year, month, _ := t.Date()
		quarter := (month-1)/3 + 1
		endQuarterMonth := quarter * 3
		_ = &Time{
			Time:   time.Date(year, endQuarterMonth+1, 0, 23, 59, 59, 999999999, t.Location()),
			Config: t.Config,
		}
	}
}

func Benchmark_EndOfQuarter_Variant12(b *testing.B) {
	t := Now()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		year, month, _ := t.Date()
		quarter := (month-1)/3 + 1
		_ = &Time{
			Time:   time.Date(year, time.Month(quarter*3)+1, 0, 23, 59, 59, 999999999, t.Location()),
			Config: t.Config,
		}
	}
}
