package xtime

import (
	"testing"
	"time"
)

func genWeekTestTimes(n int) []*Time {
	times := make([]*Time, n)
	base := time.Date(2024, 5, 11, 15, 30, 45, 123456789, time.Local)
	for i := 0; i < n; i++ {
		times[i] = With(base.Add(time.Duration(i) * time.Hour))
	}
	return times
}

func BenchmarkBOW_Baseline(b *testing.B) {
	times := genWeekTestTimes(1000)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		t := times[i%len(times)]
		t1 := t.BeginningOfDay()
		weekday := int(t1.Weekday())
		if t.WeekStartDay != time.Sunday {
			weekStartDayInt := int(t.WeekStartDay)
			if weekday < weekStartDayInt {
				weekday = weekday + 7 - weekStartDayInt
			} else {
				weekday = weekday - weekStartDayInt
			}
		}
		_ = With(t1.AddDate(0, 0, -weekday))
	}
}

func BenchmarkBOW_InlineBOD(b *testing.B) {
	times := genWeekTestTimes(1000)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		t := times[i%len(times)]
		year, month, day := t.Date()
		midnight := time.Date(year, month, day, 0, 0, 0, 0, t.Location())
		weekday := int(midnight.Weekday())
		if t.WeekStartDay != time.Sunday {
			weekStartDayInt := int(t.WeekStartDay)
			if weekday < weekStartDayInt {
				weekday = weekday + 7 - weekStartDayInt
			} else {
				weekday = weekday - weekStartDayInt
			}
		}
		_ = With(midnight.AddDate(0, 0, -weekday))
	}
}

func BenchmarkBOW_ConfigReuse(b *testing.B) {
	times := genWeekTestTimes(1000)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		t := times[i%len(times)]
		year, month, day := t.Date()
		midnight := time.Date(year, month, day, 0, 0, 0, 0, t.Location())
		weekday := int(midnight.Weekday())
		if t.WeekStartDay != time.Sunday {
			weekStartDayInt := int(t.WeekStartDay)
			if weekday < weekStartDayInt {
				weekday = weekday + 7 - weekStartDayInt
			} else {
				weekday = weekday - weekStartDayInt
			}
		}
		cfg := t.Config
		if cfg == nil {
			cfg = &Config{}
		}
		_ = &Time{Time: midnight.AddDate(0, 0, -weekday), Config: cfg}
	}
}

func BenchmarkBOW_Modulo(b *testing.B) {
	times := genWeekTestTimes(1000)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		t := times[i%len(times)]
		year, month, day := t.Date()
		midnight := time.Date(year, month, day, 0, 0, 0, 0, t.Location())
		weekday := int(midnight.Weekday())
		if t.WeekStartDay != time.Sunday {
			weekStartDayInt := int(t.WeekStartDay)
			weekday = (weekday - weekStartDayInt + 7) % 7
		}
		cfg := t.Config
		if cfg == nil {
			cfg = &Config{}
		}
		_ = &Time{Time: midnight.AddDate(0, 0, -weekday), Config: cfg}
	}
}

func BenchmarkBOW_Precalc(b *testing.B) {
	times := genWeekTestTimes(1000)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		t := times[i%len(times)]
		year, month, day := t.Date()
		loc := t.Location()
		midnight := time.Date(year, month, day, 0, 0, 0, 0, loc)
		weekday := int(midnight.Weekday())
		if t.WeekStartDay != time.Sunday {
			weekStartDayInt := int(t.WeekStartDay)
			weekday = (weekday - weekStartDayInt + 7) % 7
		}
		cfg := t.Config
		if cfg == nil {
			cfg = &Config{}
		}
		_ = &Time{Time: midnight.AddDate(0, 0, -weekday), Config: cfg}
	}
}

func BenchmarkBOW_FastPathSunday(b *testing.B) {
	times := genWeekTestTimes(1000)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		t := times[i%len(times)]
		year, month, day := t.Date()
		loc := t.Location()
		midnight := time.Date(year, month, day, 0, 0, 0, 0, loc)
		var offset int
		if t.WeekStartDay == time.Sunday {
			offset = int(midnight.Weekday())
		} else {
			weekday := int(midnight.Weekday())
			weekStartDayInt := int(t.WeekStartDay)
			offset = (weekday - weekStartDayInt + 7) % 7
		}
		cfg := t.Config
		if cfg == nil {
			cfg = &Config{}
		}
		_ = &Time{Time: midnight.AddDate(0, 0, -offset), Config: cfg}
	}
}

func BenchmarkBOW_ZeroAlloc(b *testing.B) {
	times := genWeekTestTimes(1000)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		t := times[i%len(times)]
		year, month, day := t.Date()
		loc := t.Location()
		midnight := time.Date(year, month, day, 0, 0, 0, 0, loc)
		weekday := int(midnight.Weekday())
		if t.WeekStartDay != time.Sunday {
			weekStartDayInt := int(t.WeekStartDay)
			weekday = (weekday - weekStartDayInt + 7) % 7
		}
		_ = &Time{Time: midnight.AddDate(0, 0, -weekday), Config: t.Config}
	}
}

func BenchmarkBOW_SinceLogic(b *testing.B) {
	times := genWeekTestTimes(1000)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		t := times[i%len(times)]
		year, month, day := t.Date()
		loc := t.Location()
		midnight := time.Date(year, month, day, 0, 0, 0, 0, loc)
		weekday := midnight.Weekday()
		var daysToAdd int
		if t.WeekStartDay == time.Sunday {
			daysToAdd = -int(weekday)
		} else {
			daysToAdd = int(t.WeekStartDay) - int(weekday)
			if daysToAdd > 0 {
				daysToAdd -= 7
			}
		}
		cfg := t.Config
		if cfg == nil {
			cfg = &Config{}
		}
		_ = &Time{Time: midnight.AddDate(0, 0, daysToAdd), Config: cfg}
	}
}

func BenchmarkBOW_FullyInline(b *testing.B) {
	times := genWeekTestTimes(1000)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		t := times[i%len(times)]
		year, month, day := t.Date()
		loc := t.Location()
		midnight := time.Date(year, month, day, 0, 0, 0, 0, loc)
		weekday := int(midnight.Weekday())
		if t.WeekStartDay != time.Sunday {
			weekStartDayInt := int(t.WeekStartDay)
			if weekday < weekStartDayInt {
				weekday = weekday + 7 - weekStartDayInt
			} else {
				weekday = weekday - weekStartDayInt
			}
		}
		cfg := t.Config
		if cfg == nil {
			cfg = &Config{}
		}
		_ = &Time{Time: midnight.AddDate(0, 0, -weekday), Config: cfg}
	}
}

func BenchmarkBOW_UnixCalc(b *testing.B) {
	times := genWeekTestTimes(1000)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		t := times[i%len(times)]
		year, month, day := t.Date()
		loc := t.Location()
		midnight := time.Date(year, month, day, 0, 0, 0, 0, loc)
		weekday := int(midnight.Weekday())
		if t.WeekStartDay != time.Sunday {
			weekStartDayInt := int(t.WeekStartDay)
			weekday = (weekday - weekStartDayInt + 7) % 7
		}
		cfg := t.Config
		if cfg == nil {
			cfg = &Config{}
		}
		result := midnight.Add(-time.Duration(weekday) * 24 * time.Hour)
		_ = &Time{Time: result, Config: cfg}
	}
}

func BenchmarkBOW_Optimized(b *testing.B) {
	times := genWeekTestTimes(1000)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		t := times[i%len(times)]
		year, month, day := t.Date()
		loc := t.Location()
		midnight := time.Date(year, month, day, 0, 0, 0, 0, loc)
		weekday := int(midnight.Weekday())
		if t.WeekStartDay != time.Sunday {
			weekStartDayInt := int(t.WeekStartDay)
			weekday = (weekday - weekStartDayInt + 7) % 7
		}
		cfg := t.Config
		if cfg == nil {
			cfg = &Config{}
		}
		_ = &Time{Time: midnight.AddDate(0, 0, -weekday), Config: cfg}
	}
}

func BenchmarkBOW_Baseline_Alloc(b *testing.B) {
	t := With(time.Date(2024, 5, 11, 15, 30, 45, 123456789, time.Local))
	b.ResetTimer()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		t1 := t.BeginningOfDay()
		weekday := int(t1.Weekday())
		if t.WeekStartDay != time.Sunday {
			weekStartDayInt := int(t.WeekStartDay)
			if weekday < weekStartDayInt {
				weekday = weekday + 7 - weekStartDayInt
			} else {
				weekday = weekday - weekStartDayInt
			}
		}
		_ = With(t1.AddDate(0, 0, -weekday))
	}
}

func BenchmarkBOW_Optimized_Alloc(b *testing.B) {
	t := With(time.Date(2024, 5, 11, 15, 30, 45, 123456789, time.Local))
	b.ResetTimer()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		year, month, day := t.Date()
		loc := t.Location()
		midnight := time.Date(year, month, day, 0, 0, 0, 0, loc)
		weekday := int(midnight.Weekday())
		if t.WeekStartDay != time.Sunday {
			weekStartDayInt := int(t.WeekStartDay)
			weekday = (weekday - weekStartDayInt + 7) % 7
		}
		cfg := t.Config
		if cfg == nil {
			cfg = &Config{}
		}
		_ = &Time{Time: midnight.AddDate(0, 0, -weekday), Config: cfg}
	}
}

func BenchmarkBOW_Baseline_MondayStart(b *testing.B) {
	t := With(time.Date(2024, 5, 11, 15, 30, 45, 123456789, time.Local))
	t.WeekStartDay = time.Monday
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		t1 := t.BeginningOfDay()
		weekday := int(t1.Weekday())
		if t.WeekStartDay != time.Sunday {
			weekStartDayInt := int(t.WeekStartDay)
			if weekday < weekStartDayInt {
				weekday = weekday + 7 - weekStartDayInt
			} else {
				weekday = weekday - weekStartDayInt
			}
		}
		_ = With(t1.AddDate(0, 0, -weekday))
	}
}

func BenchmarkBOW_Optimized_MondayStart(b *testing.B) {
	t := With(time.Date(2024, 5, 11, 15, 30, 45, 123456789, time.Local))
	t.WeekStartDay = time.Monday
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		year, month, day := t.Date()
		loc := t.Location()
		midnight := time.Date(year, month, day, 0, 0, 0, 0, loc)
		weekday := int(midnight.Weekday())
		if t.WeekStartDay != time.Sunday {
			weekStartDayInt := int(t.WeekStartDay)
			weekday = (weekday - weekStartDayInt + 7) % 7
		}
		cfg := t.Config
		if cfg == nil {
			cfg = &Config{}
		}
		_ = &Time{Time: midnight.AddDate(0, 0, -weekday), Config: cfg}
	}
}

func BenchmarkBOW_Baseline_Small(b *testing.B) {
	times := genWeekTestTimes(10)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		t := times[i%len(times)]
		t1 := t.BeginningOfDay()
		weekday := int(t1.Weekday())
		if t.WeekStartDay != time.Sunday {
			weekStartDayInt := int(t.WeekStartDay)
			if weekday < weekStartDayInt {
				weekday = weekday + 7 - weekStartDayInt
			} else {
				weekday = weekday - weekStartDayInt
			}
		}
		_ = With(t1.AddDate(0, 0, -weekday))
	}
}

func BenchmarkBOW_Optimized_Small(b *testing.B) {
	times := genWeekTestTimes(10)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		t := times[i%len(times)]
		year, month, day := t.Date()
		loc := t.Location()
		midnight := time.Date(year, month, day, 0, 0, 0, 0, loc)
		weekday := int(midnight.Weekday())
		if t.WeekStartDay != time.Sunday {
			weekStartDayInt := int(t.WeekStartDay)
			weekday = (weekday - weekStartDayInt + 7) % 7
		}
		cfg := t.Config
		if cfg == nil {
			cfg = &Config{}
		}
		_ = &Time{Time: midnight.AddDate(0, 0, -weekday), Config: cfg}
	}
}
