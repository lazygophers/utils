package xtime

import (
	"testing"
	"time"
)

// Benchmark_EndOfWeek_Original - 原始实现
// 当前代码：使用 With() 创建新的默认 Config
func Benchmark_EndOfWeek_Original(b *testing.B) {
	t := Now()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_ = With(t.BeginningOfWeek().AddDate(0, 0, 7).Add(-time.Nanosecond))
	}
}

// Benchmark_EndOfWeek_Opt1 - 优化方案1：直接构造 Time 结构体
// 复用 BeginningOfWeek 返回的 Config
func Benchmark_EndOfWeek_Opt1(b *testing.B) {
	t := Now()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		bow := t.BeginningOfWeek()
		eow := bow.AddDate(0, 0, 7).Add(-time.Nanosecond)
		cfg := bow.Config
		if cfg == nil {
			cfg = &Config{}
		}
		_ = &Time{Time: eow, Config: cfg}
	}
}

// Benchmark_EndOfWeek_Opt2 - 优化方案2：合并 AddDate 和 Add 操作
// 一次性计算下周最后一纳秒
func Benchmark_EndOfWeek_Opt2(b *testing.B) {
	t := Now()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		bow := t.BeginningOfWeek()
		// 7天减1纳秒 = 6天23:59:59.999999999
		eow := bow.Add(7*24*time.Hour - time.Nanosecond)
		cfg := bow.Config
		if cfg == nil {
			cfg = &Config{}
		}
		_ = &Time{Time: eow, Config: cfg}
	}
}

// Benchmark_EndOfWeek_Opt3 - 优化方案3：使用 Duration 常量
// 预计算一周减1纳秒的 Duration
func Benchmark_EndOfWeek_Opt3(b *testing.B) {
	t := Now()
	const weekMinusOneNano = 7*24*time.Hour - time.Nanosecond
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		bow := t.BeginningOfWeek()
		eow := bow.Add(weekMinusOneNano)
		cfg := bow.Config
		if cfg == nil {
			cfg = &Config{}
		}
		_ = &Time{Time: eow, Config: cfg}
	}
}

// Benchmark_EndOfWeek_Opt4 - 优化方案4：简化 Config 处理
// 使用 p.Config 直接传递（可能为 nil）
func Benchmark_EndOfWeek_Opt4(b *testing.B) {
	t := Now()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		bow := t.BeginningOfWeek()
		eow := bow.AddDate(0, 0, 7).Add(-time.Nanosecond)
		_ = &Time{Time: eow, Config: t.Config}
	}
}

// Benchmark_EndOfWeek_Opt5 - 优化方案5：合并 Config 处理和 Add
// 单次表达式完成
func Benchmark_EndOfWeek_Opt5(b *testing.B) {
	t := Now()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		bow := t.BeginningOfWeek()
		cfg := bow.Config
		if cfg == nil {
			cfg = &Config{}
		}
		_ = &Time{Time: bow.AddDate(0, 0, 7).Add(-time.Nanosecond), Config: cfg}
	}
}

// Benchmark_EndOfWeek_Opt6 - 优化方案6：内联 BeginningOfWeek 逻辑
// 直接计算周结束时间，避免函数调用
func Benchmark_EndOfWeek_Opt6(b *testing.B) {
	t := Now()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		year, month, day := t.Date()
		loc := t.Location()
		midnight := time.Date(year, month, day, 0, 0, 0, 0, loc)
		weekday := int(midnight.Weekday())

		cfg := t.Config
		if cfg != nil && t.WeekStartDay != time.Sunday {
			weekStartDayInt := int(t.WeekStartDay)
			weekday = (weekday - weekStartDayInt + 7) % 7
		}

		if cfg == nil {
			cfg = &Config{}
		}

		// 周起始 + 7天 - 1纳秒 = 周结束
		bow := midnight.AddDate(0, 0, -weekday)
		eow := bow.AddDate(0, 0, 7).Add(-time.Nanosecond)
		_ = &Time{Time: eow, Config: cfg}
	}
}

// Benchmark_EndOfWeek_Opt7 - 优化方案7：内联 + Duration 常量
// 内联逻辑并使用 Duration 常量
func Benchmark_EndOfWeek_Opt7(b *testing.B) {
	t := Now()
	const weekMinusOneNano = 7*24*time.Hour - time.Nanosecond
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		year, month, day := t.Date()
		loc := t.Location()
		midnight := time.Date(year, month, day, 0, 0, 0, 0, loc)
		weekday := int(midnight.Weekday())

		cfg := t.Config
		if cfg != nil && t.WeekStartDay != time.Sunday {
			weekStartDayInt := int(t.WeekStartDay)
			weekday = (weekday - weekStartDayInt + 7) % 7
		}

		if cfg == nil {
			cfg = &Config{}
		}

		bow := midnight.AddDate(0, 0, -weekday)
		eow := bow.Add(weekMinusOneNano)
		_ = &Time{Time: eow, Config: cfg}
	}
}

// Benchmark_EndOfWeek_Opt8 - 优化方案8：直接计算周日23:59:59.999999999
// 跳过周起始，直接计算周结束的最后一刻
func Benchmark_EndOfWeek_Opt8(b *testing.B) {
	t := Now()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		year, month, day := t.Date()
		loc := t.Location()
		midnight := time.Date(year, month, day, 0, 0, 0, 0, loc)
		weekday := int(midnight.Weekday())

		cfg := t.Config
		if cfg != nil && t.WeekStartDay != time.Sunday {
			weekStartDayInt := int(t.WeekStartDay)
			weekday = (weekday - weekStartDayInt + 7) % 7
		}

		if cfg == nil {
			cfg = &Config{}
		}

		// 周结束 = (当前 - weekday) + 6天
		eod := midnight.AddDate(0, 0, -weekday+6)
		eowTime := time.Date(eod.Year(), eod.Month(), eod.Day(), 23, 59, 59, int(time.Second-time.Nanosecond), loc)
		_ = &Time{Time: eowTime, Config: cfg}
	}
}

// Benchmark_EndOfWeek_Opt9 - 优化方案9：使用 AddDate 计算周日
// 计算本周日，然后设置为 23:59:59.999999999
func Benchmark_EndOfWeek_Opt9(b *testing.B) {
	t := Now()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		year, month, day := t.Date()
		loc := t.Location()
		midnight := time.Date(year, month, day, 0, 0, 0, 0, loc)
		weekday := int(midnight.Weekday())

		cfg := t.Config
		if cfg != nil && t.WeekStartDay != time.Sunday {
			weekStartDayInt := int(t.WeekStartDay)
			weekday = (weekday - weekStartDayInt + 7) % 7
		}

		if cfg == nil {
			cfg = &Config{}
		}

		// 当前 + (6 - weekday) 天 = 周日
		sunday := midnight.AddDate(0, 0, 6-weekday)
		eowTime := time.Date(sunday.Year(), sunday.Month(), sunday.Day(), 23, 59, 59, int(time.Second-time.Nanosecond), loc)
		_ = &Time{Time: eowTime, Config: cfg}
	}
}

// Benchmark_EndOfWeek_Opt10 - 优化方案10：简化 Config 判断
// 假设 Config 不为 nil（常见情况）
func Benchmark_EndOfWeek_Opt10(b *testing.B) {
	t := Now()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		bow := t.BeginningOfWeek()
		eow := bow.AddDate(0, 0, 7).Add(-time.Nanosecond)
		_ = &Time{Time: eow, Config: bow.Config}
	}
}

// Benchmark_EndOfWeek_Opt11 - 优化方案11：使用 EndOfDay 模式
// 获取周日后调用 EndOfDay
func Benchmark_EndOfWeek_Opt11(b *testing.B) {
	t := Now()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		year, month, day := t.Date()
		loc := t.Location()
		midnight := time.Date(year, month, day, 0, 0, 0, 0, loc)
		weekday := int(midnight.Weekday())

		cfg := t.Config
		if cfg != nil && t.WeekStartDay != time.Sunday {
			weekStartDayInt := int(t.WeekStartDay)
			weekday = (weekday - weekStartDayInt + 7) % 7
		}

		if cfg == nil {
			cfg = &Config{}
		}

		// 计算周日
		sunday := midnight.AddDate(0, 0, 6-weekday)
		// 设置为周日结束
		eowTime := time.Date(sunday.Year(), sunday.Month(), sunday.Day(), 23, 59, 59, int(time.Second-time.Nanosecond), loc)
		_ = &Time{Time: eowTime, Config: cfg}
	}
}

// Benchmark_EndOfWeek_Opt12 - 优化方案12：完全内联并优化
// 最简洁的内联实现
func Benchmark_EndOfWeek_Opt12(b *testing.B) {
	t := Now()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		loc := t.Location()
		year, month, day := t.Date()
		midnight := time.Date(year, month, day, 0, 0, 0, 0, loc)
		weekday := int(midnight.Weekday())

		cfg := t.Config
		if cfg != nil && t.WeekStartDay != time.Sunday {
			weekStartDayInt := int(t.WeekStartDay)
			weekday = (weekday - weekStartDayInt + 7) % 7
		}

		if cfg == nil {
			cfg = &Config{}
		}

		// 周日 = 当前 + (6-weekday)天
		sundayDay := day + 6 - weekday
		eowTime := time.Date(year, month, sundayDay, 23, 59, 59, int(time.Second-time.Nanosecond), loc)
		_ = &Time{Time: eowTime, Config: cfg}
	}
}

// Benchmark_EndOfWeek_Small - 小数据集测试
func Benchmark_EndOfWeek_Original_Small(b *testing.B) {
	t := time.Date(2024, 1, 15, 12, 30, 45, 123456789, time.Local)
	bt := &Time{Time: t}
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_ = With(bt.BeginningOfWeek().AddDate(0, 0, 7).Add(-time.Nanosecond))
	}
}

func Benchmark_EndOfWeek_Opt1_Small(b *testing.B) {
	t := time.Date(2024, 1, 15, 12, 30, 45, 123456789, time.Local)
	bt := &Time{Time: t}
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		bow := bt.BeginningOfWeek()
		eow := bow.AddDate(0, 0, 7).Add(-time.Nanosecond)
		cfg := bow.Config
		if cfg == nil {
			cfg = &Config{}
		}
		_ = &Time{Time: eow, Config: cfg}
	}
}

// Benchmark_EndOfWeek_Medium - 中等数据集测试
func Benchmark_EndOfWeek_Original_Medium(b *testing.B) {
	t := time.Date(2024, 6, 15, 12, 30, 45, 123456789, time.Local)
	bt := &Time{Time: t}
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_ = With(bt.BeginningOfWeek().AddDate(0, 0, 7).Add(-time.Nanosecond))
	}
}

func Benchmark_EndOfWeek_Opt1_Medium(b *testing.B) {
	t := time.Date(2024, 6, 15, 12, 30, 45, 123456789, time.Local)
	bt := &Time{Time: t}
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		bow := bt.BeginningOfWeek()
		eow := bow.AddDate(0, 0, 7).Add(-time.Nanosecond)
		cfg := bow.Config
		if cfg == nil {
			cfg = &Config{}
		}
		_ = &Time{Time: eow, Config: cfg}
	}
}

// Benchmark_EndOfWeek_Large - 大数据集测试
func Benchmark_EndOfWeek_Original_Large(b *testing.B) {
	t := time.Date(2024, 12, 31, 23, 59, 59, 999999999, time.Local)
	bt := &Time{Time: t}
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_ = With(bt.BeginningOfWeek().AddDate(0, 0, 7).Add(-time.Nanosecond))
	}
}

func Benchmark_EndOfWeek_Opt1_Large(b *testing.B) {
	t := time.Date(2024, 12, 31, 23, 59, 59, 999999999, time.Local)
	bt := &Time{Time: t}
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		bow := bt.BeginningOfWeek()
		eow := bow.AddDate(0, 0, 7).Add(-time.Nanosecond)
		cfg := bow.Config
		if cfg == nil {
			cfg = &Config{}
		}
		_ = &Time{Time: eow, Config: cfg}
	}
}

// Benchmark_EndOfWeek_Parallel - 并发测试
func Benchmark_EndOfWeek_Original_Parallel(b *testing.B) {
	t := Now()
	b.ResetTimer()

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			_ = With(t.BeginningOfWeek().AddDate(0, 0, 7).Add(-time.Nanosecond))
		}
	})
}

func Benchmark_EndOfWeek_Opt1_Parallel(b *testing.B) {
	t := Now()
	b.ResetTimer()

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			bow := t.BeginningOfWeek()
			eow := bow.AddDate(0, 0, 7).Add(-time.Nanosecond)
			cfg := bow.Config
			if cfg == nil {
				cfg = &Config{}
			}
			_ = &Time{Time: eow, Config: cfg}
		}
	})
}

// Benchmark_EndOfWeek_WithConfig - 带 Config 的测试
func Benchmark_EndOfWeek_Original_WithConfig(b *testing.B) {
	t := Now()
	t.Config = &Config{WeekStartDay: time.Monday}
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_ = With(t.BeginningOfWeek().AddDate(0, 0, 7).Add(-time.Nanosecond))
	}
}

func Benchmark_EndOfWeek_Opt1_WithConfig(b *testing.B) {
	t := Now()
	t.Config = &Config{WeekStartDay: time.Monday}
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		bow := t.BeginningOfWeek()
		eow := bow.AddDate(0, 0, 7).Add(-time.Nanosecond)
		cfg := bow.Config
		if cfg == nil {
			cfg = &Config{}
		}
		_ = &Time{Time: eow, Config: cfg}
	}
}
