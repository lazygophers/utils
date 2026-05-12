package xtime

import (
	"testing"
	"time"
)

// 生成测试时间（EndOfDay 专用）
func genEODTestTimes(n int) []time.Time {
	times := make([]time.Time, n)
	base := time.Date(2024, 5, 11, 15, 30, 45, 123456789, time.Local)
	for i := 0; i < n; i++ {
		times[i] = base.Add(time.Duration(i) * time.Hour)
	}
	return times
}

// ========== 12种优化方案基准测试 ==========

// 方案1: Baseline - 当前实现 (Date + With)
func BenchmarkEOD_Baseline(b *testing.B) {
	times := genEODTestTimes(1000)
	wrapper := make([]*Time, len(times))
	for i := range times {
		wrapper[i] = With(times[i])
	}
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		t := wrapper[i%len(wrapper)]
		y, m, d := t.Date()
		_ = With(time.Date(y, m, d, 23, 59, 59, int(time.Second-time.Nanosecond), t.Time.Location()))
	}
}

// 方案2: 直接构造 Time 结构体，复用 Config
func BenchmarkEOD_DirectConstruct(b *testing.B) {
	times := genEODTestTimes(1000)
	wrapper := make([]*Time, len(times))
	for i := range times {
		wrapper[i] = With(times[i])
	}
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		t := wrapper[i%len(wrapper)]
		loc := t.Location()
		year, month, day := t.Date()
		eod := time.Date(year, month, day, 23, 59, 59, int(time.Second-time.Nanosecond), loc)
		_ = &Time{Time: eod, Config: t.Config}
	}
}

// 方案3: 使用 BeginningOfDay + Add
func BenchmarkEOD_BoDAdd(b *testing.B) {
	times := genEODTestTimes(1000)
	wrapper := make([]*Time, len(times))
	for i := range times {
		wrapper[i] = With(times[i])
	}
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		t := wrapper[i%len(wrapper)]
		year, month, day := t.Date()
		bod := time.Date(year, month, day, 0, 0, 0, 0, t.Location())
		eod := bod.Add(24*time.Hour - time.Nanosecond)
		_ = &Time{Time: eod, Config: t.Config}
	}
}

// 方案4: 使用 Truncate + Add
func BenchmarkEOD_TruncateAdd(b *testing.B) {
	times := genEODTestTimes(1000)
	wrapper := make([]*Time, len(times))
	for i := range times {
		wrapper[i] = With(times[i])
	}
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		t := wrapper[i%len(wrapper)]
		truncated := t.Truncate(24 * time.Hour)
		eod := truncated.Add(24*time.Hour - time.Nanosecond)
		_ = &Time{Time: eod, Config: t.Config}
	}
}

// 方案5: 使用 Add 向上取整
func BenchmarkEOD_AddRoundUp(b *testing.B) {
	times := genEODTestTimes(1000)
	wrapper := make([]*Time, len(times))
	for i := range times {
		wrapper[i] = With(times[i])
	}
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		t := wrapper[i%len(wrapper)]
		h, m, s := t.Clock()
		nanos := t.Nanosecond()
		// 计算到当天的剩余时间
		remaining := (24-time.Duration(h))*time.Hour -
			time.Duration(m)*time.Minute -
			time.Duration(s)*time.Second -
			time.Duration(nanos)*time.Nanosecond
		_ = &Time{Time: t.Add(remaining - time.Nanosecond), Config: t.Config}
	}
}

// 方案6: 使用 AddDate + Add
func BenchmarkEOD_AddDate(b *testing.B) {
	times := genEODTestTimes(1000)
	wrapper := make([]*Time, len(times))
	for i := range times {
		wrapper[i] = With(times[i])
	}
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		t := wrapper[i%len(wrapper)]
		year, month, day := t.Date()
		tomorrow := time.Date(year, month, day, 0, 0, 0, 0, t.Location()).AddDate(0, 0, 1)
		eod := tomorrow.Add(-time.Nanosecond)
		_ = &Time{Time: eod, Config: t.Config}
	}
}

// 方案7: 使用 In + Date
func BenchmarkEOD_InDate(b *testing.B) {
	times := genEODTestTimes(1000)
	wrapper := make([]*Time, len(times))
	for i := range times {
		wrapper[i] = With(times[i])
	}
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		t := wrapper[i%len(wrapper)]
		loc := t.Location()
		year, month, day := t.In(loc).Date()
		eod := time.Date(year, month, day, 23, 59, 59, int(time.Second-time.Nanosecond), loc)
		_ = &Time{Time: eod, Config: t.Config}
	}
}

// 方案8: 预计算常量
func BenchmarkEOD_PrecomputedConst(b *testing.B) {
	times := genEODTestTimes(1000)
	wrapper := make([]*Time, len(times))
	for i := range times {
		wrapper[i] = With(times[i])
	}
	eodTime := time.Date(0, 0, 0, 23, 59, 59, int(time.Second-time.Nanosecond), time.UTC)
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		t := wrapper[i%len(wrapper)]
		year, month, day := t.Date()
		eod := time.Date(year, month, day, 0, 0, 0, 0, t.Location()).Add(eodTime.Sub(time.Date(0, 0, 0, 0, 0, 0, 0, time.UTC)))
		_ = &Time{Time: eod, Config: t.Config}
	}
}

// 方案9: 使用 BeginningOfDay 方法 + Add
func BenchmarkEOD_BoDMethod(b *testing.B) {
	times := genEODTestTimes(1000)
	wrapper := make([]*Time, len(times))
	for i := range times {
		wrapper[i] = With(times[i])
	}
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		t := wrapper[i%len(wrapper)]
		year, month, day := t.Date()
		bod := &Time{Time: time.Date(year, month, day, 0, 0, 0, 0, t.Location()), Config: t.Config}
		eod := bod.Time.Add(24*time.Hour - time.Nanosecond)
		_ = &Time{Time: eod, Config: t.Config}
	}
}

// 方案10: 直接构造，使用 Config 复用（检查 nil）
func BenchmarkEOD_DirectConstructWithNilCheck(b *testing.B) {
	times := genEODTestTimes(1000)
	wrapper := make([]*Time, len(times))
	for i := range times {
		wrapper[i] = With(times[i])
	}
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		t := wrapper[i%len(wrapper)]
		loc := t.Location()
		year, month, day := t.Date()
		eod := time.Date(year, month, day, 23, 59, 59, int(time.Second-time.Nanosecond), loc)
		cfg := t.Config
		if cfg == nil {
			cfg = &Config{}
		}
		_ = &Time{Time: eod, Config: cfg}
	}
}

// 方案11: 使用 Unix 时间戳
func BenchmarkEOD_UnixTimestamp(b *testing.B) {
	times := genEODTestTimes(1000)
	wrapper := make([]*Time, len(times))
	for i := range times {
		wrapper[i] = With(times[i])
	}
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		t := wrapper[i%len(wrapper)]
		year, month, day := t.Date()
		eod := time.Date(year, month, day, 23, 59, 59, int(time.Second-time.Nanosecond), t.Location())
		_ = &Time{Time: eod, Config: t.Config}
	}
}

// 方案12: 组合优化 - Date + Config 复用
func BenchmarkEOD_CombinedOptimized(b *testing.B) {
	times := genEODTestTimes(1000)
	wrapper := make([]*Time, len(times))
	for i := range times {
		wrapper[i] = With(times[i])
	}
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		t := wrapper[i%len(wrapper)]
		year, month, day := t.Date()
		eod := time.Date(year, month, day, 23, 59, 59, int(time.Second-time.Nanosecond), t.Location())
		cfg := t.Config
		if cfg == nil {
			cfg = &Config{}
		}
		_ = &Time{Time: eod, Config: cfg}
	}
}
