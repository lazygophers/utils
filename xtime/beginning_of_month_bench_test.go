package xtime

import (
	"testing"
	"time"
)

func genMonthTestTimes(n int) []*Time {
	times := make([]*Time, n)
	base := time.Date(2024, 5, 11, 15, 30, 45, 123456789, time.Local)
	for i := 0; i < n; i++ {
		times[i] = With(base.Add(time.Duration(i) * 24 * time.Hour))
	}
	return times
}

// 方案1: Baseline - 当前实现
func BenchmarkBOM_Baseline(b *testing.B) {
	times := genMonthTestTimes(1000)
	b.ResetTimer()
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		t := times[i%len(times)]
		y, m, _ := t.Date()
		_ = With(time.Date(y, m, 1, 0, 0, 0, 0, t.Location()))
	}
}

// 方案2: ConfigReuse - 复用Config（参考BeginningOfDay优化）
func BenchmarkBOM_ConfigReuse(b *testing.B) {
	times := genMonthTestTimes(1000)
	b.ResetTimer()
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		t := times[i%len(times)]
		y, m, _ := t.Date()
		firstDay := time.Date(y, m, 1, 0, 0, 0, 0, t.Location())
		cfg := t.Config
		if cfg == nil {
			cfg = &Config{}
		}
		_ = &Time{Time: firstDay, Config: cfg}
	}
}

// 方案3: DirectStruct - 直接构造结构体，不调用With
func BenchmarkBOM_DirectStruct(b *testing.B) {
	times := genMonthTestTimes(1000)
	b.ResetTimer()
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		t := times[i%len(times)]
		y, m, _ := t.Date()
		_ = &Time{
			Time:   time.Date(y, m, 1, 0, 0, 0, 0, t.Location()),
			Config: t.Config,
		}
	}
}

// 方案4: InlineDate - 内联time.Date调用
func BenchmarkBOM_InlineDate(b *testing.B) {
	times := genMonthTestTimes(1000)
	b.ResetTimer()
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		t := times[i%len(times)]
		_ = &Time{
			Time:   time.Date(t.Year(), t.Month(), 1, 0, 0, 0, 0, t.Location()),
			Config: t.Config,
		}
	}
}

// 方案5: ZeroAlloc - 零分配优化（复用Config）
func BenchmarkBOM_ZeroAlloc(b *testing.B) {
	times := genMonthTestTimes(1000)
	b.ResetTimer()
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		t := times[i%len(times)]
		cfg := t.Config
		if cfg == nil {
			cfg = &Config{}
		}
		_ = &Time{
			Time:   time.Date(t.Year(), t.Month(), 1, 0, 0, 0, 0, t.Location()),
			Config: cfg,
		}
	}
}

// 方案6: TruncateMethod - 使用Truncate方法
func BenchmarkBOM_TruncateMethod(b *testing.B) {
	times := genMonthTestTimes(1000)
	b.ResetTimer()
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		t := times[i%len(times)]
		y, m, _ := t.Date()
		firstDay := time.Date(y, m, 1, 0, 0, 0, 0, t.Location())
		cfg := t.Config
		if cfg == nil {
			cfg = &Config{}
		}
		_ = &Time{Time: firstDay, Config: cfg}
	}
}

// 方案7: PreallocConfig - 预先检查Config
func BenchmarkBOM_PreallocConfig(b *testing.B) {
	times := genMonthTestTimes(1000)
	b.ResetTimer()
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		t := times[i%len(times)]
		var cfg *Config
		if t.Config != nil {
			cfg = t.Config
		}
		_ = &Time{
			Time:   time.Date(t.Year(), t.Month(), 1, 0, 0, 0, 0, t.Location()),
			Config: cfg,
		}
	}
}

// 方案8: AddDateMethod - 使用AddDate（从 BeginningOfMonth 本身调用）
func BenchmarkBOM_AddDateMethod(b *testing.B) {
	times := genMonthTestTimes(1000)
	b.ResetTimer()
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		t := times[i%len(times)]
		// 先到本月第一天，然后减去天数-1
		d := t.Day()
		firstDay := t.AddDate(0, 0, -(d - 1))
		cfg := t.Config
		if cfg == nil {
			cfg = &Config{}
		}
		_ = &Time{Time: firstDay, Config: cfg}
	}
}

// 方案9: Combined - 结合多种优化
func BenchmarkBOM_Combined(b *testing.B) {
	times := genMonthTestTimes(1000)
	b.ResetTimer()
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		t := times[i%len(times)]
		cfg := t.Config
		loc := t.Location()
		_ = &Time{
			Time:   time.Date(t.Year(), t.Month(), 1, 0, 0, 0, 0, loc),
			Config: cfg,
		}
	}
}

// 方案10: UnixTime - 使用Unix时间计算（实验性）
func BenchmarkBOM_UnixTime(b *testing.B) {
	times := genMonthTestTimes(1000)
	b.ResetTimer()
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		t := times[i%len(times)]
		y, m, _ := t.Date()
		cfg := t.Config
		_ = &Time{
			Time:   time.Date(y, m, 1, 0, 0, 0, 0, t.Location()),
			Config: cfg,
		}
	}
}

// 方案11: DirectYMD - 直接使用Year/Month/Day
func BenchmarkBOM_DirectYMD(b *testing.B) {
	times := genMonthTestTimes(1000)
	b.ResetTimer()
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		t := times[i%len(times)]
		_ = &Time{
			Time:   time.Date(t.Year(), t.Month(), 1, 0, 0, 0, 0, t.Location()),
			Config: t.Config,
		}
	}
}

// 方案12: NilConfigCheck - 显式nil检查
func BenchmarkBOM_NilConfigCheck(b *testing.B) {
	times := genMonthTestTimes(1000)
	b.ResetTimer()
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		t := times[i%len(times)]
		cfg := t.Config
		if cfg == nil {
			cfg = &Config{}
		}
		y, m, _ := t.Date()
		_ = &Time{
			Time:   time.Date(y, m, 1, 0, 0, 0, 0, t.Location()),
			Config: cfg,
		}
	}
}

// 对比测试：不同数据大小
func BenchmarkBOM_Baseline_Small(b *testing.B) {
	times := genMonthTestTimes(10)
	b.ResetTimer()
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		t := times[i%len(times)]
		y, m, _ := t.Date()
		_ = With(time.Date(y, m, 1, 0, 0, 0, 0, t.Location()))
	}
}

func BenchmarkBOM_DirectStruct_Small(b *testing.B) {
	times := genMonthTestTimes(10)
	b.ResetTimer()
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		t := times[i%len(times)]
		y, m, _ := t.Date()
		_ = &Time{
			Time:   time.Date(y, m, 1, 0, 0, 0, 0, t.Location()),
			Config: t.Config,
		}
	}
}

func BenchmarkBOM_Baseline_Medium(b *testing.B) {
	times := genMonthTestTimes(100)
	b.ResetTimer()
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		t := times[i%len(times)]
		y, m, _ := t.Date()
		_ = With(time.Date(y, m, 1, 0, 0, 0, 0, t.Location()))
	}
}

func BenchmarkBOM_DirectStruct_Medium(b *testing.B) {
	times := genMonthTestTimes(100)
	b.ResetTimer()
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		t := times[i%len(times)]
		y, m, _ := t.Date()
		_ = &Time{
			Time:   time.Date(y, m, 1, 0, 0, 0, 0, t.Location()),
			Config: t.Config,
		}
	}
}

func BenchmarkBOM_Baseline_Large(b *testing.B) {
	times := genMonthTestTimes(1000)
	b.ResetTimer()
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		t := times[i%len(times)]
		y, m, _ := t.Date()
		_ = With(time.Date(y, m, 1, 0, 0, 0, 0, t.Location()))
	}
}

func BenchmarkBOM_DirectStruct_Large(b *testing.B) {
	times := genMonthTestTimes(1000)
	b.ResetTimer()
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		t := times[i%len(times)]
		y, m, _ := t.Date()
		_ = &Time{
			Time:   time.Date(y, m, 1, 0, 0, 0, 0, t.Location()),
			Config: t.Config,
		}
	}
}
