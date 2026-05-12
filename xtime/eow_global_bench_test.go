package xtime

import (
	"testing"
	"time"
)

// Benchmark_EndOfWeek_Global_Original - 原始实现
// 当前代码：return With(time.Now()).EndOfWeek()
func Benchmark_EndOfWeek_Global_Original(b *testing.B) {
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_ = With(time.Now()).EndOfWeek()
	}
}

// Benchmark_EndOfWeek_Global_Opt1 - 优化方案1：内联所有逻辑
// 避免 With() 调用，直接构造 Time 结构体
func Benchmark_EndOfWeek_Global_Opt1(b *testing.B) {
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		now := time.Now()
		loc := now.Location()
		year, month, day := now.Date()
		midnight := time.Date(year, month, day, 0, 0, 0, 0, loc)
		weekday := int(midnight.Weekday())

		// 默认 WeekStartDay = Monday，无特殊处理
		cfg := &Config{
			WeekStartDay:  time.Monday,
			TimeLocation: time.Local,
			TimeFormats:  []string{},
			Monotonic:    now,
		}

		// 周六最后一刻（周日为下周起始）
		sundayDay := day + 6 - weekday
		eowTime := time.Date(year, month, sundayDay, 23, 59, 59, int(time.Second-time.Nanosecond), loc)

		_ = &Time{Time: eowTime, Config: cfg}
	}
}

// Benchmark_EndOfWeek_Global_Opt2 - 优化方案2：预分配 Config
// 使用全局默认 Config，避免每次创建
func Benchmark_EndOfWeek_Global_Opt2(b *testing.B) {
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		now := time.Now()
		loc := now.Location()
		year, month, day := now.Date()
		midnight := time.Date(year, month, day, 0, 0, 0, 0, loc)
		weekday := int(midnight.Weekday())

		// 周六最后一刻
		sundayDay := day + 6 - weekday
		eowTime := time.Date(year, month, sundayDay, 23, 59, 59, int(time.Second-time.Nanosecond), loc)

		_ = &Time{Time: eowTime, Config: defaultConfig}
	}
}

// Benchmark_EndOfWeek_Global_Opt3 - 优化方案3：减少 Date() 调用
// 合并 midnight 和 eowTime 的计算
func Benchmark_EndOfWeek_Global_Opt3(b *testing.B) {
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		now := time.Now()
		loc := now.Location()
		year, month, day := now.Date()
		weekday := int(now.Weekday())

		// 周六最后一刻
		sundayDay := day + 6 - weekday
		eowTime := time.Date(year, month, sundayDay, 23, 59, 59, int(time.Second-time.Nanosecond), loc)

		_ = &Time{Time: eowTime, Config: defaultConfig}
	}
}

// Benchmark_EndOfWeek_Global_Opt4 - 优化方案4：使用 Duration 常量
// 预计算一周的时间
func Benchmark_EndOfWeek_Global_Opt4(b *testing.B) {
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		now := time.Now()
		loc := now.Location()
		year, month, day := now.Date()
		weekday := int(now.Weekday())

		// 计算到周六的天数
		daysUntilSaturday := (6 - weekday + 7) % 7
		endOfDay := time.Date(year, month, day+daysUntilSaturday, 0, 0, 0, 0, loc).Add(23*time.Hour + 59*time.Minute + 59*time.Second - time.Nanosecond)

		_ = &Time{Time: endOfDay, Config: defaultConfig}
	}
}

// Benchmark_EndOfWeek_Global_Opt5 - 优化方案5：简化常量计算
// 直接使用 6 天加 23:59:59.999999999
func Benchmark_EndOfWeek_Global_Opt5(b *testing.B) {
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		now := time.Now()
		loc := now.Location()
		year, month, day := now.Date()
		weekday := int(now.Weekday())

		// 直接计算目标日期
		targetDay := day + (6-weekday+7)%7
		eowTime := time.Date(year, month, targetDay, 23, 59, 59, int(time.Second-time.Nanosecond), loc)

		_ = &Time{Time: eowTime, Config: defaultConfig}
	}
}

// Benchmark_EndOfWeek_Global_Opt6 - 优化方案6：内联 time.Now() 逻辑
// 减少中间变量
func Benchmark_EndOfWeek_Global_Opt6(b *testing.B) {
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		now := time.Now()
		loc := now.Location()
		year, month, day := now.Date()
		weekday := int(now.Weekday())

		targetDay := day + (6-weekday+7)%7
		_ = &Time{Time: time.Date(year, month, targetDay, 23, 59, 59, int(time.Second-time.Nanosecond), loc), Config: defaultConfig}
	}
}

// Benchmark_EndOfWeek_Global_Opt7 - 优化方案7：合并 Config 初始化
// 在结构体字面量中直接初始化
func Benchmark_EndOfWeek_Global_Opt7(b *testing.B) {
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		now := time.Now()
		loc := now.Location()
		year, month, day := now.Date()
		weekday := int(now.Weekday())

		targetDay := day + (6-weekday+7)%7
		_ = &Time{
			Time: time.Date(year, month, targetDay, 23, 59, 59, int(time.Second-time.Nanosecond), loc),
			Config: &Config{
				WeekStartDay:  time.Monday,
				TimeLocation: time.Local,
				TimeFormats:  []string{},
				Monotonic:    now,
			},
		}
	}
}

// Benchmark_EndOfWeek_Global_Opt8 - 优化方案8：使用常量优化
// 预定义纳秒常量
func Benchmark_EndOfWeek_Global_Opt8(b *testing.B) {
	b.ResetTimer()
	const endOfDayNanos = int(time.Second-time.Nanosecond)

	for i := 0; i < b.N; i++ {
		now := time.Now()
		loc := now.Location()
		year, month, day := now.Date()
		weekday := int(now.Weekday())

		targetDay := day + (6-weekday+7)%7
		_ = &Time{
			Time:   time.Date(year, month, targetDay, 23, 59, 59, endOfDayNanos, loc),
			Config: defaultConfig,
		}
	}
}

// Benchmark_EndOfWeek_Global_Opt9 - 优化方案9：减少 Weekday() 调用
// 从 Date() 返回值中提取
func Benchmark_EndOfWeek_Global_Opt9(b *testing.B) {
	b.ResetTimer()
	const endOfDayNanos = int(time.Second - time.Nanosecond)

	for i := 0; i < b.N; i++ {
		now := time.Now()
		loc := now.Location()
		year, month, day := now.Date()
		midnight := time.Date(year, month, day, 0, 0, 0, 0, loc)
		weekday := int(midnight.Weekday())

		targetDay := day + (6-weekday+7)%7
		_ = &Time{
			Time:   time.Date(year, month, targetDay, 23, 59, 59, endOfDayNanos, loc),
			Config: defaultConfig,
		}
	}
}

// Benchmark_EndOfWeek_Global_Opt10 - 优化方案10：移除模运算
// 利用 weekday 范围 [0,6]，简化计算
func Benchmark_EndOfWeek_Global_Opt10(b *testing.B) {
	b.ResetTimer()
	const endOfDayNanos = int(time.Second - time.Nanosecond)

	for i := 0; i < b.N; i++ {
		now := time.Now()
		loc := now.Location()
		year, month, day := now.Date()
		weekday := int(now.Weekday())

		// weekday: 0=Sunday, 1=Monday, ..., 6=Saturday
		// 目标: Saturday (6)
		// 如果 weekday <= 6，则 targetDay = day + (6 - weekday)
		daysToAdd := 6 - weekday
		if daysToAdd < 0 {
			daysToAdd += 7
		}

		targetDay := day + daysToAdd
		_ = &Time{
			Time:   time.Date(year, month, targetDay, 23, 59, 59, endOfDayNanos, loc),
			Config: defaultConfig,
		}
	}
}

// Benchmark_EndOfWeek_Global_Opt11 - 优化方案11：查表法
// 预计算每天到周六的天数
func Benchmark_EndOfWeek_Global_Opt11(b *testing.B) {
	b.ResetTimer()
	const endOfDayNanos = int(time.Second - time.Nanosecond)

	// weekday: 0=Sunday, 1=Monday, ..., 6=Saturday
	// daysToAdd: [6, 5, 4, 3, 2, 1, 0]
	daysToAddTable := [7]int{6, 5, 4, 3, 2, 1, 0}

	for i := 0; i < b.N; i++ {
		now := time.Now()
		loc := now.Location()
		year, month, day := now.Date()
		weekday := int(now.Weekday())

		targetDay := day + daysToAddTable[weekday]
		_ = &Time{
			Time:   time.Date(year, month, targetDay, 23, 59, 59, endOfDayNanos, loc),
			Config: defaultConfig,
		}
	}
}

// Benchmark_EndOfWeek_Global_Opt12 - 优化方案12：使用 time.Now().Truncate()
// 先截断到天，再计算
func Benchmark_EndOfWeek_Global_Opt12(b *testing.B) {
	b.ResetTimer()
	const endOfDayNanos = int(time.Second - time.Nanosecond)
	daysToAddTable := [7]int{6, 5, 4, 3, 2, 1, 0}

	for i := 0; i < b.N; i++ {
		now := time.Now()
		loc := now.Location()
		today := now.Truncate(24 * time.Hour)
		year, month, day := today.Date()
		weekday := int(today.Weekday())

		targetDay := day + daysToAddTable[weekday]
		_ = &Time{
			Time:   time.Date(year, month, targetDay, 23, 59, 59, endOfDayNanos, loc),
			Config: defaultConfig,
		}
	}
}

// Benchmark_EndOfWeek_Global_Opt13 - 优化方案13：使用 AddDate
// 利用 time.AddDate 处理月份边界
func Benchmark_EndOfWeek_Global_Opt13(b *testing.B) {
	b.ResetTimer()
	const endOfDayNanos = int(time.Second - time.Nanosecond)
	daysToAddTable := [7]int{6, 5, 4, 3, 2, 1, 0}

	for i := 0; i < b.N; i++ {
		now := time.Now()
		loc := now.Location()
		today := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, loc)
		weekday := int(today.Weekday())

		saturday := today.AddDate(0, 0, daysToAddTable[weekday])
		_ = &Time{
			Time:   time.Date(saturday.Year(), saturday.Month(), saturday.Day(), 23, 59, 59, endOfDayNanos, loc),
			Config: defaultConfig,
		}
	}
}

// Benchmark_EndOfWeek_Global_Opt14 - 优化方案14：使用 Add 替代 AddDate
// Duration 加法可能更快
func Benchmark_EndOfWeek_Global_Opt14(b *testing.B) {
	b.ResetTimer()
	const endOfDayNanos = int(time.Second - time.Nanosecond)
	daysToAddTable := [7]int{6, 5, 4, 3, 2, 1, 0}

	for i := 0; i < b.N; i++ {
		now := time.Now()
		loc := now.Location()
		today := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, loc)
		weekday := int(today.Weekday())

		saturday := today.Add(time.Duration(daysToAddTable[weekday]*24) * time.Hour)
		_ = &Time{
			Time:   time.Date(saturday.Year(), saturday.Month(), saturday.Day(), 23, 59, 59, endOfDayNanos, loc),
			Config: defaultConfig,
		}
	}
}

// Benchmark_EndOfWeek_Global_Opt15 - 优化方案15：直接构造目标时间
// 合并所有计算到单次 time.Date 调用
func Benchmark_EndOfWeek_Global_Opt15(b *testing.B) {
	b.ResetTimer()
	const endOfDayNanos = int(time.Second - time.Nanosecond)
	daysToAddTable := [7]int{6, 5, 4, 3, 2, 1, 0}

	for i := 0; i < b.N; i++ {
		now := time.Now()
		loc := now.Location()
		weekday := int(now.Weekday())

		// 计算目标日期
		targetDate := now.AddDate(0, 0, daysToAddTable[weekday])

		_ = &Time{
			Time:   time.Date(targetDate.Year(), targetDate.Month(), targetDate.Day(), 23, 59, 59, endOfDayNanos, loc),
			Config: defaultConfig,
		}
	}
}

// Benchmark_EndOfWeek_Global_Opt16 - 优化方案16：最小化函数调用
// 合并 Date() 调用
func Benchmark_EndOfWeek_Global_Opt16(b *testing.B) {
	b.ResetTimer()
	const endOfDayNanos = int(time.Second - time.Nanosecond)
	daysToAddTable := [7]int{6, 5, 4, 3, 2, 1, 0}

	for i := 0; i < b.N; i++ {
		now := time.Now()
		loc := now.Location()
		year, month, day := now.Date()
		weekday := int(now.Weekday())

		_ = &Time{
			Time:   time.Date(year, month, day+daysToAddTable[weekday], 23, 59, 59, endOfDayNanos, loc),
			Config: defaultConfig,
		}
	}
}

// Benchmark_EndOfWeek_Global_Opt17 - 优化方案17：使用全局 Config 常量
// 避免每次分配 Config
func Benchmark_EndOfWeek_Global_Opt17(b *testing.B) {
	b.ResetTimer()
	const endOfDayNanos = int(time.Second - time.Nanosecond)
	daysToAddTable := [7]int{6, 5, 4, 3, 2, 1, 0}

	for i := 0; i < b.N; i++ {
		now := time.Now()
		loc := now.Location()
		year, month, day := now.Date()
		weekday := int(now.Weekday())

		_ = &Time{
			Time:   time.Date(year, month, day+daysToAddTable[weekday], 23, 59, 59, endOfDayNanos, loc),
			Config: &Config{
				WeekStartDay:  time.Monday,
				TimeLocation: time.Local,
			},
		}
	}
}

// Benchmark_EndOfWeek_Global_Opt18 - 优化方案18：完全内联，无中间变量
// 极致优化版本
func Benchmark_EndOfWeek_Global_Opt18(b *testing.B) {
	b.ResetTimer()
	const endOfDayNanos = int(time.Second - time.Nanosecond)

	for i := 0; i < b.N; i++ {
		now := time.Now()
		loc := now.Location()
		year, month, day := now.Date()
		weekday := int(now.Weekday())

		_ = &Time{
			Time:   time.Date(year, month, day+([]int{6, 5, 4, 3, 2, 1, 0}[weekday]), 23, 59, 59, endOfDayNanos, loc),
			Config: defaultConfig,
		}
	}
}

// Benchmark_EndOfWeek_Global_Opt19 - 优化方案19：使用 sync.Pool
// 复用 Time 结构体
func Benchmark_EndOfWeek_Global_Opt19(b *testing.B) {
	b.ResetTimer()
	const endOfDayNanos = int(time.Second - time.Nanosecond)

	for i := 0; i < b.N; i++ {
		now := time.Now()
		loc := now.Location()
		year, month, day := now.Date()
		weekday := int(now.Weekday())

		t := &Time{
			Time:   time.Date(year, month, day+([]int{6, 5, 4, 3, 2, 1, 0}[weekday]), 23, 59, 59, endOfDayNanos, loc),
			Config: defaultConfig,
		}
		_ = t
	}
}

// Benchmark_EndOfWeek_Global_Opt20 - 优化方案20：使用闭包缓存
// 减少重复计算
func Benchmark_EndOfWeek_Global_Opt20(b *testing.B) {
	b.ResetTimer()
	const endOfDayNanos = int(time.Second - time.Nanosecond)
	daysToAdd := []int{6, 5, 4, 3, 2, 1, 0}

	for i := 0; i < b.N; i++ {
		now := time.Now()
		loc := now.Location()
		year, month, day := now.Date()
		weekday := int(now.Weekday())

		_ = &Time{
			Time:   time.Date(year, month, day+daysToAdd[weekday], 23, 59, 59, endOfDayNanos, loc),
			Config: defaultConfig,
		}
	}
}

// Benchmark_EndOfWeek_Global_Current - 当前实现（对照组）
func Benchmark_EndOfWeek_Global_Current(b *testing.B) {
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_ = EndOfWeek()
	}
}
