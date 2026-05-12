package xtime

import (
	"testing"
	"time"
)

// 生成测试数据（固定种子保证可重复）
func genEndOfYearTestTimes(n int) []*Time {
	times := make([]*Time, n)
	base := time.Date(2024, 6, 15, 14, 30, 45, 123456789, time.Local)
	for i := 0; i < n; i++ {
		times[i] = With(base.Add(time.Duration(i) * time.Hour))
	}
	return times
}

// Baseline: 当前实现
func BenchmarkEndOfYear_Current(b *testing.B) {
	times := genEndOfYearTestTimes(100)
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_ = times[i%len(times)].EndOfYear()
	}
}

// 方案1: 直接调用 BeginningOfYear + AddDate + Add，避免 With
func BenchmarkEndOfYear_Opt1_DirectWithDate(b *testing.B) {
	times := genEndOfYearTestTimes(100)
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		t := times[i%len(times)]
		// 直接内联 BeginningOfYear 逻辑
		year := t.Time.Year()
		loc := t.Time.Location()
		config := t.Config
		if config == nil {
			config = &Config{}
		}
		// 年初 + 1年 - 1纳秒
		nextYear := time.Date(year, time.January, 1, 0, 0, 0, 0, loc).AddDate(1, 0, 0).Add(-time.Nanosecond)
		_ = &Time{Time: nextYear, Config: config}
	}
}

// 方案2: time.Date 溢出技巧 (year+1, Jan, 1, 0, 0, 0, -1)
func BenchmarkEndOfYear_Opt2_DateOverflow(b *testing.B) {
	times := genEndOfYearTestTimes(100)
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		t := times[i%len(times)]
		year := t.Time.Year()
		loc := t.Time.Location()
		config := t.Config
		if config == nil {
			config = &Config{}
		}
		// 下年1月1日 0点，减1纳秒 = 今年最后一刻
		end := time.Date(year+1, time.January, 1, 0, 0, 0, -1, loc)
		_ = &Time{Time: end, Config: config}
	}
}

// 方案3: 直接构造 12/31 23:59:59.999999999
func BenchmarkEndOfYear_Opt3_DirectDec31(b *testing.B) {
	times := genEndOfYearTestTimes(100)
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		t := times[i%len(times)]
		year := t.Time.Year()
		loc := t.Time.Location()
		config := t.Config
		if config == nil {
			config = &Config{}
		}
		// 直接构造12月31日最后一刻
		end := time.Date(year, time.December, 31, 23, 59, 59, 999999999, loc)
		_ = &Time{Time: end, Config: config}
	}
}

// 方案4: time.Date(0) 溢出技巧 (year+1, Jan, 0)
func BenchmarkEndOfYear_Opt4_DateZeroOverflow(b *testing.B) {
	times := genEndOfYearTestTimes(100)
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		t := times[i%len(times)]
		year := t.Time.Year()
		loc := t.Time.Location()
		config := t.Config
		if config == nil {
			config = &Config{}
		}
		// year+1年1月0日 = 今年12月31日
		end := time.Date(year+1, time.January, 0, 23, 59, 59, 999999999, loc)
		_ = &Time{Time: end, Config: config}
	}
}

// 方案5: 使用 Unix 时间戳计算
func BenchmarkEndOfYear_Opt5_UnixTimestamp(b *testing.B) {
	times := genEndOfYearTestTimes(100)
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		t := times[i%len(times)]
		year := t.Time.Year()
		loc := t.Time.Location()
		config := t.Config
		if config == nil {
			config = &Config{}
		}
		// 计算年末 Unix 时间戳
		end := time.Date(year+1, time.January, 1, 0, 0, 0, 0, loc).Add(-time.Nanosecond)
		_ = &Time{Time: end, Config: config}
	}
}

// 方案6: 预计算年末日期再 Add
func BenchmarkEndOfYear_Opt6_PrecalcDate(b *testing.B) {
	times := genEndOfYearTestTimes(100)
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		t := times[i%len(times)]
		year := t.Time.Year()
		loc := t.Time.Location()
		config := t.Config
		if config == nil {
			config = &Config{}
		}
		// 预计算年末日期
		endDate := time.Date(year, time.December, 31, 0, 0, 0, 0, loc)
		end := endDate.Add(24*time.Hour - time.Nanosecond)
		_ = &Time{Time: end, Config: config}
	}
}

// 方案7: 使用 Truncate + Add
func BenchmarkEndOfYear_Opt7_TruncateAdd(b *testing.B) {
	times := genEndOfYearTestTimes(100)
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		t := times[i%len(times)]
		year := t.Time.Year()
		loc := t.Time.Location()
		config := t.Config
		if config == nil {
			config = &Config{}
		}
		// 下年初 Truncate 到年，减1纳秒
		nextYear := time.Date(year+1, time.January, 1, 0, 0, 0, 0, loc)
		end := nextYear.Add(-time.Nanosecond)
		_ = &Time{Time: end, Config: config}
	}
}

// 方案8: 内联 BeginningOfYear 完整逻辑
func BenchmarkEndOfYear_Opt8_InlineBeginningOfYear(b *testing.B) {
	times := genEndOfYearTestTimes(100)
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		t := times[i%len(times)]
		year := t.Time.Year()
		loc := t.Time.Location()
		config := t.Config
		if config == nil {
			config = &Config{}
		}
		// 完全内联
		beginningOfYear := time.Date(year, time.January, 1, 0, 0, 0, 0, loc)
		end := beginningOfYear.AddDate(1, 0, 0).Add(-time.Nanosecond)
		_ = &Time{Time: end, Config: config}
	}
}

// 方案9: 使用 AddDate 直接加1年到下年初
func BenchmarkEndOfYear_Opt9_AddDateOnly(b *testing.B) {
	times := genEndOfYearTestTimes(100)
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		t := times[i%len(times)]
		loc := t.Time.Location()
		config := t.Config
		if config == nil {
			config = &Config{}
		}
		// 当前时间 + (1年 - 当前年内已过时间) - 1纳秒
		// 简化：直接用当前年份
		year := t.Time.Year()
		nextYearStart := time.Date(year+1, time.January, 1, 0, 0, 0, 0, loc)
		end := nextYearStart.Add(-time.Nanosecond)
		_ = &Time{Time: end, Config: config}
	}
}

// 方案10: 最优方案 - time.Date(0) 溢出 + 预提取 Config
func BenchmarkEndOfYear_Opt10_Best(b *testing.B) {
	times := genEndOfYearTestTimes(100)
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		t := times[i%len(times)]
		year := t.Time.Year()
		loc := t.Time.Location()
		config := t.Config
		if config == nil {
			config = &Config{}
		}
		// year+1年1月0日 = 今年12月31日
		end := time.Date(year+1, time.January, 0, 23, 59, 59, 999999999, loc)
		_ = &Time{Time: end, Config: config}
	}
}

// 方案11: 使用 time.Date 的负纳秒参数
func BenchmarkEndOfYear_Opt11_NegativeNanosec(b *testing.B) {
	times := genEndOfYearTestTimes(100)
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		t := times[i%len(times)]
		year := t.Time.Year()
		loc := t.Time.Location()
		config := t.Config
		if config == nil {
			config = &Config{}
		}
		// 下年初，纳秒=-1
		end := time.Date(year+1, time.January, 1, 0, 0, 0, -1, loc)
		_ = &Time{Time: end, Config: config}
	}
}

// 方案12: 混合方案 - 预检查 Config
func BenchmarkEndOfYear_Opt12_ConfigCheck(b *testing.B) {
	times := genEndOfYearTestTimes(100)
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		t := times[i%len(times)]
		year := t.Time.Year()
		loc := t.Time.Location()
		config := t.Config
		// 预检查，避免内联判断
		end := time.Date(year+1, time.January, 1, 0, 0, 0, -1, loc)
		_ = &Time{Time: end, Config: config}
	}
}

// 对比组：不同大小数据集
func BenchmarkEndOfYear_Current_Small(b *testing.B) {
	times := genEndOfYearTestTimes(10)
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_ = times[i%len(times)].EndOfYear()
	}
}

func BenchmarkEndOfYear_Opt2_Small(b *testing.B) {
	times := genEndOfYearTestTimes(10)
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		t := times[i%len(times)]
		year := t.Time.Year()
		loc := t.Time.Location()
		config := t.Config
		if config == nil {
			config = &Config{}
		}
		end := time.Date(year+1, time.January, 1, 0, 0, 0, -1, loc)
		_ = &Time{Time: end, Config: config}
	}
}

func BenchmarkEndOfYear_Current_Medium(b *testing.B) {
	times := genEndOfYearTestTimes(100)
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_ = times[i%len(times)].EndOfYear()
	}
}

func BenchmarkEndOfYear_Opt2_Medium(b *testing.B) {
	times := genEndOfYearTestTimes(100)
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		t := times[i%len(times)]
		year := t.Time.Year()
		loc := t.Time.Location()
		config := t.Config
		if config == nil {
			config = &Config{}
		}
		end := time.Date(year+1, time.January, 1, 0, 0, 0, -1, loc)
		_ = &Time{Time: end, Config: config}
	}
}

func BenchmarkEndOfYear_Current_Large(b *testing.B) {
	times := genEndOfYearTestTimes(1000)
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_ = times[i%len(times)].EndOfYear()
	}
}

func BenchmarkEndOfYear_Opt2_Large(b *testing.B) {
	times := genEndOfYearTestTimes(1000)
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		t := times[i%len(times)]
		year := t.Time.Year()
		loc := t.Time.Location()
		config := t.Config
		if config == nil {
			config = &Config{}
		}
		end := time.Date(year+1, time.January, 1, 0, 0, 0, -1, loc)
		_ = &Time{Time: end, Config: config}
	}
}
