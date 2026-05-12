package xtime

import (
	"testing"
	"time"
)

// 生成测试时间
func genTestTimes(n int) []time.Time {
	times := make([]time.Time, n)
	base := time.Date(2024, 5, 11, 15, 30, 45, 123456789, time.Local)
	for i := 0; i < n; i++ {
		times[i] = base.Add(time.Duration(i) * time.Hour)
	}
	return times
}

// ========== 12种优化方案基准测试 ==========

// 方案1: Baseline - 当前实现 (Date + With)
func BenchmarkBOD_Baseline(b *testing.B) {
	times := genTestTimes(1000)
	wrapper := make([]*Time, len(times))
	for i := range times {
		wrapper[i] = With(times[i])
	}
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		t := wrapper[i%len(wrapper)]
		y, m, d := t.Date()
		_ = With(time.Date(y, m, d, 0, 0, 0, 0, t.Time.Location()))
	}
}

// 方案2: 使用 Truncate
func BenchmarkBOD_Truncate(b *testing.B) {
	times := genTestTimes(1000)
	wrapper := make([]*Time, len(times))
	for i := range times {
		wrapper[i] = With(times[i])
	}
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_ = With(wrapper[i%len(wrapper)].Truncate(24 * time.Hour))
	}
}

// 方案3: 直接使用 Date，不调用 With
func BenchmarkBOD_DateNoWith(b *testing.B) {
	times := genTestTimes(1000)
	wrapper := make([]*Time, len(times))
	for i := range times {
		wrapper[i] = With(times[i])
	}
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		t := wrapper[i%len(wrapper)]
		loc := t.Location()
		year, month, day := t.Date()
		midnight := time.Date(year, month, day, 0, 0, 0, 0, loc)
		_ = &Time{Time: midnight, Config: t.Config}
	}
}

// 方案4: 使用 Add 向下取整
func BenchmarkBOD_AddRound(b *testing.B) {
	times := genTestTimes(1000)
	wrapper := make([]*Time, len(times))
	for i := range times {
		wrapper[i] = With(times[i])
	}
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		t := wrapper[i%len(wrapper)]
		h, m, s := t.Clock()
		nanos := t.Nanosecond()
		duration := time.Duration(h)*time.Hour + time.Duration(m)*time.Minute + time.Duration(s)*time.Second + time.Duration(nanos)*time.Nanosecond
		_ = With(t.Add(-duration))
	}
}

// 方案5: 使用 In + Truncate
func BenchmarkBOD_InTruncate(b *testing.B) {
	times := genTestTimes(1000)
	wrapper := make([]*Time, len(times))
	for i := range times {
		wrapper[i] = With(times[i])
	}
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		t := wrapper[i%len(wrapper)]
		loc := t.Location()
		truncated := t.In(loc).Truncate(24 * time.Hour)
		_ = &Time{Time: truncated, Config: t.Config}
	}
}

// 方案6: 减去当天已过时间
func BenchmarkBOD_Subtract(b *testing.B) {
	times := genTestTimes(1000)
	wrapper := make([]*Time, len(times))
	for i := range times {
		wrapper[i] = With(times[i])
	}
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		t := wrapper[i%len(wrapper)]
		h, m, s := t.Clock()
		ns := t.Nanosecond()
		elapsed := time.Duration(h)*time.Hour + time.Duration(m)*time.Minute + time.Duration(s)*time.Second + time.Duration(ns)*time.Nanosecond
		_ = &Time{Time: t.Add(-elapsed), Config: t.Config}
	}
}

// 方案7: Unix + Date 组合
func BenchmarkBOD_UnixDate(b *testing.B) {
	times := genTestTimes(1000)
	wrapper := make([]*Time, len(times))
	for i := range times {
		wrapper[i] = With(times[i])
	}
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		t := wrapper[i%len(wrapper)]
		loc := t.Location()
		year, month, day := t.Date()
		midnight := time.Date(year, month, day, 0, 0, 0, 0, loc)
		_ = With(midnight)
	}
}

// 方案8: 缓存 Location 引用
func BenchmarkBOD_CacheLocation(b *testing.B) {
	times := genTestTimes(1000)
	wrapper := make([]*Time, len(times))
	for i := range times {
		wrapper[i] = With(times[i])
	}
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		t := wrapper[i%len(wrapper)]
		loc := t.Location()
		year, month, day := t.Date()
		midnight := time.Date(year, month, day, 0, 0, 0, 0, loc)
		cfg := t.Config
		if cfg == nil {
			cfg = &Config{}
		}
		_ = &Time{Time: midnight, Config: cfg}
	}
}

// 方案9: 零分配优化 - 直接构造 Time
func BenchmarkBOD_ZeroAlloc(b *testing.B) {
	times := genTestTimes(1000)
	wrapper := make([]*Time, len(times))
	for i := range times {
		wrapper[i] = With(times[i])
	}
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		t := wrapper[i%len(wrapper)]
		year, month, day := t.Date()
		midnight := time.Date(year, month, day, 0, 0, 0, 0, t.Location())
		_ = &Time{Time: midnight, Config: t.Config}
	}
}

// 方案10: 直接返回，处理 nil Config
func BenchmarkBOD_DirectReturn(b *testing.B) {
	times := genTestTimes(1000)
	wrapper := make([]*Time, len(times))
	for i := range times {
		wrapper[i] = With(times[i])
	}
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		t := wrapper[i%len(wrapper)]
		year, month, day := t.Date()
		midnight := time.Date(year, month, day, 0, 0, 0, 0, t.Location())
		if t.Config == nil {
			_ = &Time{Time: midnight, Config: &Config{}}
		} else {
			_ = &Time{Time: midnight, Config: t.Config}
		}
	}
}

// 方案11: Date 优化 - 单次 Location 调用
func BenchmarkBOD_DateOptimized(b *testing.B) {
	times := genTestTimes(1000)
	wrapper := make([]*Time, len(times))
	for i := range times {
		wrapper[i] = With(times[i])
	}
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		t := wrapper[i%len(wrapper)]
		loc := t.Location()
		year, month, day := t.Date()
		t1 := time.Date(year, month, day, 0, 0, 0, 0, loc)
		cfg := t.Config
		if cfg == nil {
			cfg = &Config{}
		}
		_ = &Time{Time: t1, Config: cfg}
	}
}

// 方案12: Truncate + Config 复用（最优方案）
func BenchmarkBOD_Optimized(b *testing.B) {
	times := genTestTimes(1000)
	wrapper := make([]*Time, len(times))
	for i := range times {
		wrapper[i] = With(times[i])
	}
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		t := wrapper[i%len(wrapper)]
		timeVal := t.Truncate(24 * time.Hour)
		cfg := t.Config
		if cfg == nil {
			cfg = &Config{}
		}
		_ = &Time{Time: timeVal, Config: cfg}
	}
}

// ========== 内存分配基准 ==========

func BenchmarkBOD_Baseline_Alloc(b *testing.B) {
	t := With(time.Date(2024, 5, 11, 15, 30, 45, 123456789, time.Local))
	b.ResetTimer()
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		y, m, d := t.Date()
		_ = With(time.Date(y, m, d, 0, 0, 0, 0, t.Time.Location()))
	}
}

func BenchmarkBOD_Truncate_Alloc(b *testing.B) {
	t := With(time.Date(2024, 5, 11, 15, 30, 45, 123456789, time.Local))
	b.ResetTimer()
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		_ = With(t.Truncate(24 * time.Hour))
	}
}

func BenchmarkBOD_DateNoWith_Alloc(b *testing.B) {
	t := With(time.Date(2024, 5, 11, 15, 30, 45, 123456789, time.Local))
	b.ResetTimer()
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		loc := t.Location()
		year, month, day := t.Date()
		midnight := time.Date(year, month, day, 0, 0, 0, 0, loc)
		_ = &Time{Time: midnight, Config: t.Config}
	}
}

func BenchmarkBOD_AddRound_Alloc(b *testing.B) {
	t := With(time.Date(2024, 5, 11, 15, 30, 45, 123456789, time.Local))
	b.ResetTimer()
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		h, m, s := t.Clock()
		nanos := t.Nanosecond()
		duration := time.Duration(h)*time.Hour + time.Duration(m)*time.Minute + time.Duration(s)*time.Second + time.Duration(nanos)*time.Nanosecond
		_ = With(t.Add(-duration))
	}
}

func BenchmarkBOD_InTruncate_Alloc(b *testing.B) {
	t := With(time.Date(2024, 5, 11, 15, 30, 45, 123456789, time.Local))
	b.ResetTimer()
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		loc := t.Location()
		truncated := t.In(loc).Truncate(24 * time.Hour)
		_ = &Time{Time: truncated, Config: t.Config}
	}
}

func BenchmarkBOD_Subtract_Alloc(b *testing.B) {
	t := With(time.Date(2024, 5, 11, 15, 30, 45, 123456789, time.Local))
	b.ResetTimer()
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		h, m, s := t.Clock()
		ns := t.Nanosecond()
		elapsed := time.Duration(h)*time.Hour + time.Duration(m)*time.Minute + time.Duration(s)*time.Second + time.Duration(ns)*time.Nanosecond
		_ = &Time{Time: t.Add(-elapsed), Config: t.Config}
	}
}

func BenchmarkBOD_UnixDate_Alloc(b *testing.B) {
	t := With(time.Date(2024, 5, 11, 15, 30, 45, 123456789, time.Local))
	b.ResetTimer()
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		loc := t.Location()
		year, month, day := t.Date()
		midnight := time.Date(year, month, day, 0, 0, 0, 0, loc)
		_ = With(midnight)
	}
}

func BenchmarkBOD_CacheLocation_Alloc(b *testing.B) {
	t := With(time.Date(2024, 5, 11, 15, 30, 45, 123456789, time.Local))
	b.ResetTimer()
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		loc := t.Location()
		year, month, day := t.Date()
		midnight := time.Date(year, month, day, 0, 0, 0, 0, loc)
		cfg := t.Config
		if cfg == nil {
			cfg = &Config{}
		}
		_ = &Time{Time: midnight, Config: cfg}
	}
}

func BenchmarkBOD_ZeroAlloc_Alloc(b *testing.B) {
	t := With(time.Date(2024, 5, 11, 15, 30, 45, 123456789, time.Local))
	b.ResetTimer()
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		year, month, day := t.Date()
		midnight := time.Date(year, month, day, 0, 0, 0, 0, t.Location())
		_ = &Time{Time: midnight, Config: t.Config}
	}
}

func BenchmarkBOD_DirectReturn_Alloc(b *testing.B) {
	t := With(time.Date(2024, 5, 11, 15, 30, 45, 123456789, time.Local))
	b.ResetTimer()
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		year, month, day := t.Date()
		midnight := time.Date(year, month, day, 0, 0, 0, 0, t.Location())
		if t.Config == nil {
			_ = &Time{Time: midnight, Config: &Config{}}
		} else {
			_ = &Time{Time: midnight, Config: t.Config}
		}
	}
}

func BenchmarkBOD_DateOptimized_Alloc(b *testing.B) {
	t := With(time.Date(2024, 5, 11, 15, 30, 45, 123456789, time.Local))
	b.ResetTimer()
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		loc := t.Location()
		year, month, day := t.Date()
		t1 := time.Date(year, month, day, 0, 0, 0, 0, loc)
		cfg := t.Config
		if cfg == nil {
			cfg = &Config{}
		}
		_ = &Time{Time: t1, Config: cfg}
	}
}

func BenchmarkBOD_Optimized_Alloc(b *testing.B) {
	t := With(time.Date(2024, 5, 11, 15, 30, 45, 123456789, time.Local))
	b.ResetTimer()
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		timeVal := t.Truncate(24 * time.Hour)
		cfg := t.Config
		if cfg == nil {
			cfg = &Config{}
		}
		_ = &Time{Time: timeVal, Config: cfg}
	}
}
