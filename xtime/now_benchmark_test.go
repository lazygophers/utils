package xtime

import (
	"sync"
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

// 基准测试
func BenchmarkBeginningOfQuarter_Optimized(b *testing.B) {
	t := With(time.Date(2024, 6, 15, 12, 30, 45, 123456789, time.UTC))
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = t.BeginningOfQuarter()
	}
}

// 原始实现对比
func BenchmarkBeginningOfQuarter_Original(b *testing.B) {
	t := With(time.Date(2024, 6, 15, 12, 30, 45, 123456789, time.UTC))
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		// 原始实现：调用 BeginningOfMonth + With
		month := t.BeginningOfMonth()
		offset := (int(month.Month()) - 1) % 3
		_ = With(month.AddDate(0, -offset, 0))
	}
}

// 变体1：内联计算，不提取配置
func BenchmarkBeginningOfQuarter_Variant1(b *testing.B) {
	t := With(time.Date(2024, 6, 15, 12, 30, 45, 123456789, time.UTC))
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		year := t.Year()
		month := int(t.Month())
		quarterStartMonth := ((month - 1) / 3) * 3
		_ = &Time{
			Time:   time.Date(year, time.Month(quarterStartMonth+1), 1, 0, 0, 0, 0, t.Location()),
			Config: t.Config,
		}
	}
}

// 变体2：使用 Month() 直接计算，减少 int 转换
func BenchmarkBeginningOfQuarter_Variant2(b *testing.B) {
	t := With(time.Date(2024, 6, 15, 12, 30, 45, 123456789, time.UTC))
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		year := t.Year()
		month := t.Month()
		quarterStartMonth := ((int(month) - 1) / 3) * 3
		_ = &Time{
			Time:   time.Date(year, time.Month(quarterStartMonth+1), 1, 0, 0, 0, 0, t.Location()),
			Config: t.Config,
		}
	}
}

// 变体3：预计算季度偏移
func BenchmarkBeginningOfQuarter_Variant3(b *testing.B) {
	t := With(time.Date(2024, 6, 15, 12, 30, 45, 123456789, time.UTC))
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		year, month, _ := t.Date()
		quarterOffset := ((int(month) - 1) / 3) * 3
		_ = &Time{
			Time:   time.Date(year, time.Month(int(month)-quarterOffset), 1, 0, 0, 0, 0, t.Location()),
			Config: t.Config,
		}
	}
}

// 变体4：使用 AddDate 计算
func BenchmarkBeginningOfQuarter_Variant4(b *testing.B) {
	t := With(time.Date(2024, 6, 15, 12, 30, 45, 123456789, time.UTC))
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		month := int(t.Month())
		offset := ((month - 1) % 3) * -1
		result := t.BeginningOfMonth().AddDate(0, offset, 0)
		_ = &Time{Time: result, Config: t.Config}
	}
}

// 变体5：查找表
func BenchmarkBeginningOfQuarter_Variant5(b *testing.B) {
	quarterStartMonths := map[int]int{
		1: 1, 2: 1, 3: 1,
		4: 4, 5: 4, 6: 4,
		7: 7, 8: 7, 9: 7,
		10: 10, 11: 10, 12: 10,
	}
	t := With(time.Date(2024, 6, 15, 12, 30, 45, 123456789, time.UTC))
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		year := t.Year()
		month := int(t.Month())
		startMonth := quarterStartMonths[month]
		_ = &Time{
			Time:   time.Date(year, time.Month(startMonth), 1, 0, 0, 0, 0, t.Location()),
			Config: t.Config,
		}
	}
}

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

// 基准测试 - BeginningOfYear 优化方案对比

// 1. Baseline - 当前实现（原始）
func BenchmarkBOY_Original(b *testing.B) {
	t := Now()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		y, _, _ := t.Date()
		_ = With(time.Date(y, time.January, 1, 0, 0, 0, 0, t.Location()))
	}
}

// 2. ConfigReuse - 复用 Config（类似 BeginningOfMonth 模式）
func BenchmarkBOY_ConfigReuse(b *testing.B) {
	t := Now()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_ = &Time{
			Time:   time.Date(t.Year(), time.January, 1, 0, 0, 0, 0, t.Location()),
			Config: t.Config,
		}
	}
}

// 3. DirectStruct - 直接构造结构体（零分配）
func BenchmarkBOY_DirectStruct(b *testing.B) {
	t := Now()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		result := &Time{
			Time:   time.Date(t.Year(), time.January, 1, 0, 0, 0, 0, t.Location()),
			Config: t.Config,
		}
		_ = result
	}
}

// 4. InlineDate - 内联 time.Date，避免多次调用
func BenchmarkBOY_InlineDate(b *testing.B) {
	t := Now()
	loc := t.Location()
	year := t.Year()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_ = With(time.Date(year, time.January, 1, 0, 0, 0, 0, loc))
	}
}

// 5. PreExtract - 预先提取所有需要的值
func BenchmarkBOY_PreExtract(b *testing.B) {
	t := Now()
	loc := t.Location()
	year := t.Year()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_ = &Time{
			Time:   time.Date(year, time.January, 1, 0, 0, 0, 0, loc),
			Config: t.Config,
		}
	}
}

// 6. AddDateMethod - 使用 AddDate（从 BeginningOfMonth 推导）
func BenchmarkBOY_AddDateMethod(b *testing.B) {
	t := Now()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_ = With(t.BeginningOfMonth().AddDate(0, -int(t.Month())+1, 0))
	}
}

// 7. ZeroAlloc - 零分配优化（显式复用）
func BenchmarkBOY_ZeroAlloc(b *testing.B) {
	t := Now()
	config := t.Config
	loc := t.Location()
	year := t.Year()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_ = &Time{
			Time:   time.Date(year, time.January, 1, 0, 0, 0, 0, loc),
			Config: config,
		}
	}
}

// 8. NilConfigCheck - 显式 nil 检查 + 复用
func BenchmarkBOY_NilConfigCheck(b *testing.B) {
	t := Now()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		config := t.Config
		if config == nil {
			config = &Config{
				WeekStartDay: time.Monday,
				TimeLocation: time.Local,
			}
		}
		_ = &Time{
			Time:   time.Date(t.Year(), time.January, 1, 0, 0, 0, 0, t.Location()),
			Config: config,
		}
	}
}

// 9. DirectYMD - 直接使用 Year()，避免 Date()
func BenchmarkBOY_DirectYMD(b *testing.B) {
	t := Now()
	year := t.Year()
	loc := t.Location()
	config := t.Config
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_ = &Time{
			Time:   time.Date(year, time.January, 1, 0, 0, 0, 0, loc),
			Config: config,
		}
	}
}

// 10. Combined - 结合多种优化（预提取 + 直接构造）
func BenchmarkBOY_Combined(b *testing.B) {
	t := Now()
	year := t.Year()
	loc := t.Location()
	config := t.Config
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_ = &Time{
			Time:   time.Date(year, time.January, 1, 0, 0, 0, 0, loc),
			Config: config,
		}
	}
}

// 11. TruncateMethod - 使用 Truncate 方法
func BenchmarkBOY_TruncateMethod(b *testing.B) {
	t := Now()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_ = With(t.Truncate(8760*time.Hour)).AddDate(-int(t.Month())+1, -int(t.Day())+1, 0)
	}
}

// 12. UnixTime - 使用 Unix 时间计算
func BenchmarkBOY_UnixTime(b *testing.B) {
	t := Now()
	loc := t.Location()
	year := t.Year()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		newTime := time.Date(year, time.January, 1, 0, 0, 0, 0, loc)
		_ = &Time{
			Time:   newTime,
			Config: t.Config,
		}
	}
}

// 13. OptimizedDirect - 最优方案（基于 BeginningOfMonth 经验）
func BenchmarkBOY_OptimizedDirect(b *testing.B) {
	t := Now()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_ = &Time{
			Time:   time.Date(t.Year(), time.January, 1, 0, 0, 0, 0, t.Location()),
			Config: t.Config,
		}
	}
}

// 14. MethodChaining - 方法链式调用
func BenchmarkBOY_MethodChaining(b *testing.B) {
	t := Now()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		bom := t.BeginningOfMonth()
		_ = With(bom.AddDate(0, -int(bom.Month())+1, 0))
	}
}

// 15. ExplicitConstruction - 显式构造（避免 With）
func BenchmarkBOY_ExplicitConstruction(b *testing.B) {
	t := Now()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		result := &Time{
			Time: time.Date(
				t.Year(),
				time.January,
				1,
				0, 0, 0, 0,
				t.Location(),
			),
			Config: t.Config,
		}
		_ = result
	}
}

// 生成测试时间（固定种子保证可重复）
func genHalfTestTimes() []*Time {
	times := make([]*Time, 12)
	base := time.Date(2024, 1, 15, 12, 30, 45, 123456789, time.Local)

	for i := 0; i < 12; i++ {
		times[i] = With(base.AddDate(0, i, 0))
	}

	return times
}

// Baseline: 当前实现
func BenchmarkBeginningOfHalf_Current(b *testing.B) {
	times := genHalfTestTimes()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		for _, t := range times {
			_ = t.BeginningOfHalf()
		}
	}
}

// 方案1: 直接计算半年起始月，复用 Config
func BenchmarkBeginningOfHalf_Opt1_DirectCalc(b *testing.B) {
	times := genHalfTestTimes()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		for _, p := range times {
			config := p.Config
			loc := p.Location()
			year := p.Year()
			month := int(p.Month())
			halfStartMonth := ((month - 1) / 6) * 6 // 0 或 6

			_ = &Time{
				Time:   time.Date(year, time.Month(halfStartMonth+1), 1, 0, 0, 0, 0, loc),
				Config: config,
			}
		}
	}
}

// 方案2: 直接计算半年起始月，使用 time.Now()
func BenchmarkBeginningOfHalf_Opt2_NoConfig(b *testing.B) {
	times := genHalfTestTimes()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		for _, p := range times {
			loc := p.Location()
			year := p.Year()
			month := int(p.Month())
			halfStartMonth := ((month - 1) / 6) * 6

			_ = &Time{
				Time: time.Date(year, time.Month(halfStartMonth+1), 1, 0, 0, 0, 0, loc),
			}
		}
	}
}

// 方案3: 预提取所有字段
func BenchmarkBeginningOfHalf_Opt3_PreExtract(b *testing.B) {
	times := genHalfTestTimes()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		for _, p := range times {
			config := p.Config
			loc := p.Location()
			year := p.Year()
			month := p.Month()
			halfStartMonth := ((int(month) - 1) / 6) * 6

			_ = &Time{
				Time:   time.Date(year, time.Month(halfStartMonth+1), 1, 0, 0, 0, 0, loc),
				Config: config,
			}
		}
	}
}

// 方案4: 使用 switch 语句（避免整数除法）
func BenchmarkBeginningOfHalf_Opt4_Switch(b *testing.B) {
	times := genHalfTestTimes()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		for _, p := range times {
			config := p.Config
			loc := p.Location()
			year := p.Year()
			month := p.Month()

			var startMonth time.Month
			switch month {
			case time.January, time.February, time.March, time.April, time.May, time.June:
				startMonth = time.January
			default:
				startMonth = time.July
			}

			_ = &Time{
				Time:   time.Date(year, startMonth, 1, 0, 0, 0, 0, loc),
				Config: config,
			}
		}
	}
}

// 方案5: 使用 if-else（避免整数除法）
func BenchmarkBeginningOfHalf_Opt5_IfElse(b *testing.B) {
	times := genHalfTestTimes()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		for _, p := range times {
			config := p.Config
			loc := p.Location()
			year := p.Year()
			month := p.Month()

			var startMonth time.Month
			if month <= time.June {
				startMonth = time.January
			} else {
				startMonth = time.July
			}

			_ = &Time{
				Time:   time.Date(year, startMonth, 1, 0, 0, 0, 0, loc),
				Config: config,
			}
		}
	}
}

// 方案6: 使用三元表达式模拟（避免整数除法）
func BenchmarkBeginningOfHalf_Opt6_TernarySim(b *testing.B) {
	times := genHalfTestTimes()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		for _, p := range times {
			config := p.Config
			loc := p.Location()
			year := p.Year()
			month := p.Month()

			offset := 0
			if month > time.June {
				offset = 6
			}

			_ = &Time{
				Time:   time.Date(year, time.Month(offset+1), 1, 0, 0, 0, 0, loc),
				Config: config,
			}
		}
	}
}

// 方案7: 查表法（半年起始月映射）
func BenchmarkBeginningOfHalf_Opt7_LookupTable(b *testing.B) {
	times := genHalfTestTimes()
	b.ResetTimer()

	halfStartMonths := map[time.Month]time.Month{
		time.January:   time.January,
		time.February:  time.January,
		time.March:     time.January,
		time.April:     time.January,
		time.May:       time.January,
		time.June:      time.January,
		time.July:      time.July,
		time.August:    time.July,
		time.September: time.July,
		time.October:   time.July,
		time.November:  time.July,
		time.December:  time.July,
	}

	for i := 0; i < b.N; i++ {
		for _, p := range times {
			config := p.Config
			loc := p.Location()
			year := p.Year()
			month := p.Month()

			_ = &Time{
				Time:   time.Date(year, halfStartMonths[month], 1, 0, 0, 0, 0, loc),
				Config: config,
			}
		}
	}
}

// 方案8: 数组查找（避免 map）
func BenchmarkBeginningOfHalf_Opt8_ArrayLookup(b *testing.B) {
	times := genHalfTestTimes()
	b.ResetTimer()

	halfStartMonths := [12]time.Month{
		time.January, time.January, time.January,
		time.January, time.January, time.January,
		time.July, time.July, time.July,
		time.July, time.July, time.July,
	}

	for i := 0; i < b.N; i++ {
		for _, p := range times {
			config := p.Config
			loc := p.Location()
			year := p.Year()
			month := p.Month()

			_ = &Time{
				Time:   time.Date(year, halfStartMonths[month-1], 1, 0, 0, 0, 0, loc),
				Config: config,
			}
		}
	}
}

// 方案9: 位运算（半年起始月）
func BenchmarkBeginningOfHalf_Opt9_Bitwise(b *testing.B) {
	times := genHalfTestTimes()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		for _, p := range times {
			config := p.Config
			loc := p.Location()
			year := p.Year()
			month := int(p.Month())

			// 位运算：第6位决定半年
			half := (month >> 6) & 1
			halfStartMonth := half * 6

			_ = &Time{
				Time:   time.Date(year, time.Month(halfStartMonth+1), 1, 0, 0, 0, 0, loc),
				Config: config,
			}
		}
	}
}

// 方案10: 混合优化（整数除法 + 预提取）
func BenchmarkBeginningOfHalf_Opt10_Hybrid(b *testing.B) {
	times := genHalfTestTimes()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		for _, p := range times {
			config := p.Config
			loc := p.Location()
			year := p.Year()
			month := int(p.Month())
			halfStartMonth := ((month - 1) / 6) * 6

			_ = &Time{
				Time:   time.Date(year, time.Month(halfStartMonth+1), 1, 0, 0, 0, 0, loc),
				Config: config,
			}
		}
	}
}

// 方案11: 内联 time.Date 调用（最激进）
func BenchmarkBeginningOfHalf_Opt11_Inlined(b *testing.B) {
	times := genHalfTestTimes()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		for _, p := range times {
			config := p.Config
			loc := p.Location()
			year := p.Year()
			month := int(p.Month())
			halfStartMonth := ((month - 1) / 6) * 6

			_ = &Time{
				Time:   time.Date(year, time.Month(halfStartMonth+1), 1, 0, 0, 0, 0, loc),
				Config: config,
			}
		}
	}
}

// 方案12: 使用 BeginningOfQuarter 逻辑扩展
func BenchmarkBeginningOfHalf_Opt12_QuarterLogic(b *testing.B) {
	times := genHalfTestTimes()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		for _, p := range times {
			config := p.Config
			loc := p.Location()
			year := p.Year()
			month := int(p.Month())

			// 类似季度逻辑，但半年
			halfStartMonth := ((month - 1) / 6) * 6

			_ = &Time{
				Time:   time.Date(year, time.Month(halfStartMonth+1), 1, 0, 0, 0, 0, loc),
				Config: config,
			}
		}
	}
}

// 内存分析基准：当前实现
func BenchmarkBeginningOfHalf_Current_Alloc(b *testing.B) {
	times := genHalfTestTimes()
	b.ResetTimer()
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		for _, t := range times {
			_ = t.BeginningOfHalf()
		}
	}
}

// 内存分析基准：最优方案
func BenchmarkBeginningOfHalf_Opt1_Alloc(b *testing.B) {
	times := genHalfTestTimes()
	b.ResetTimer()
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		for _, p := range times {
			config := p.Config
			loc := p.Location()
			year := p.Year()
			month := int(p.Month())
			halfStartMonth := ((month - 1) / 6) * 6

			_ = &Time{
				Time:   time.Date(year, time.Month(halfStartMonth+1), 1, 0, 0, 0, 0, loc),
				Config: config,
			}
		}
	}
}

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

func BenchmarkBeginningOfDay_Global_Current(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_ = BeginningOfDay_Global_V1()
	}
}

func BenchmarkBeginningOfDay_Global_V2(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_ = BeginningOfDay_Global_V2()
	}
}

func BenchmarkBeginningOfDay_Global_V3(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_ = BeginningOfDay_Global_V3()
	}
}

func BenchmarkBeginningOfDay_Global_V4(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_ = BeginningOfDay_Global_V4()
	}
}

func BenchmarkBeginningOfDay_Global_V5(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_ = BeginningOfDay_Global_V5()
	}
}

func BenchmarkBeginningOfDay_Global_V6(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_ = BeginningOfDay_Global_V6()
	}
}

func BenchmarkBeginningOfDay_Global_V7(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_ = BeginningOfDay_Global_V7()
	}
}

func BenchmarkBeginningOfDay_Global_V8(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_ = BeginningOfDay_Global_V8()
	}
}

func BenchmarkBeginningOfDay_Global_V9(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_ = BeginningOfDay_Global_V9()
	}
}

func BenchmarkBeginningOfDay_Global_V10(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_ = BeginningOfDay_Global_V10()
	}
}

func BenchmarkBeginningOfDay_Global_V11(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_ = BeginningOfDay_Global_V11()
	}
}

func BenchmarkBeginningOfDay_Global_V12(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_ = BeginningOfDay_Global_V12()
	}
}

func BenchmarkTimeNow(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_ = time.Now()
	}
}

func BenchmarkTimeNow_Date(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		now := time.Now()
		_, _, _ = now.Date()
	}
}

func BenchmarkTimeNow_Date_Construct(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		now := time.Now()
		year, month, day := now.Date()
		_ = time.Date(year, month, day, 0, 0, 0, 0, now.Location())
	}
}

// ========== 全局 BeginningOfHour() 基准测试 ==========

// Baseline: 当前实现
func BenchmarkBeginningOfHour_Current(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_ = BeginningOfHour()
	}
}

// 方案1: Truncate + nil Config
func BenchmarkBeginningOfHour_TruncateNil(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		t := time.Now()
		_ = &Time{Time: t.Truncate(time.Hour), Config: nil}
	}
}

// 方案2: Truncate + 全局共享 Config
func BenchmarkBeginningOfHour_GlobalConfig(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		t := time.Now()
		_ = &Time{Time: t.Truncate(time.Hour), Config: BeginningOfHourConfig}
	}
}

// 方案3: Truncate + 空结构体 Config
func BenchmarkBeginningOfHour_ZeroConfig(b *testing.B) {
	zeroConfig := &Config{}
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		t := time.Now()
		_ = &Time{Time: t.Truncate(time.Hour), Config: zeroConfig}
	}
}

// 方案4: 完整 Date 构建
func BenchmarkBeginningOfHour_Date(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		t := time.Now()
		y, m, d := t.Date()
		h := t.Hour()
		_ = &Time{Time: time.Date(y, m, d, h, 0, 0, 0, t.Location()), Config: nil}
	}
}

// 方案5: Date + 全局 Config
func BenchmarkBeginningOfHour_DateWithConfig(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		t := time.Now()
		y, m, d := t.Date()
		h := t.Hour()
		_ = &Time{Time: time.Date(y, m, d, h, 0, 0, 0, t.Location()), Config: BeginningOfHourConfig}
	}
}

// 方案6: Add + Subtract 方法
func BenchmarkBeginningOfHour_AddSubtract(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		t := time.Now()
		min, sec := t.Minute(), t.Second()
		ns := t.Nanosecond()
		truncated := t.Add(-time.Duration(min)*time.Minute - time.Duration(sec)*time.Second - time.Duration(ns)*time.Nanosecond)
		_ = &Time{Time: truncated, Config: nil}
	}
}

// 方案7: Unix 时间戳方法
func BenchmarkBeginningOfHour_Unix(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		t := time.Now()
		loc := t.Location()
		unix := t.Unix()
		hourSec := int64(t.Hour()) * 3600
		truncatedUnix := unix - (unix % 3600) - hourSec + int64(t.Hour())*3600
		_ = &Time{Time: time.Unix(truncatedUnix, 0).In(loc), Config: nil}
	}
}

// 方案8: 预先提取 Location
func BenchmarkBeginningOfHour_PreallocLocation(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		t := time.Now()
		loc := t.Location()
		y, m, d := t.Date()
		h := t.Hour()
		_ = &Time{Time: time.Date(y, m, d, h, 0, 0, 0, loc), Config: nil}
	}
}

// 方案9: 完整参数提取 + Config 复用
func BenchmarkBeginningOfHour_FullExtract(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		t := time.Now()
		y, m, d := t.Date()
		h := t.Hour()
		loc := t.Location()
		_ = &Time{Time: time.Date(y, m, d, h, 0, 0, 0, loc), Config: BeginningOfHourConfig}
	}
}

// 方案10: 简化版 Truncate
func BenchmarkBeginningOfHour_Minimal(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		t := time.Now()
		_ = &Time{Time: t.Truncate(time.Hour), Config: BeginningOfHourConfig}
	}
}

// 方案11: 优化版 With（避免重复创建 Config）
func BenchmarkBeginningOfHour_OptimizedWith(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		t := time.Now()
		truncated := t.Truncate(time.Hour)
		_ = &Time{Time: truncated, Config: BeginningOfHourConfig}
	}
}

// 方案12: 使用 Truncate + 嵌入式 Time
func BenchmarkBeginningOfHour_EmbeddedTime(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		t := time.Now()
		result := Time{
			Time:   t.Truncate(time.Hour),
			Config: BeginningOfHourConfig,
		}
		_ = &result
	}
}

// 方案13: 分离 Location 和 Date
func BenchmarkBeginningOfHour_SeparatedLocation(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		t := time.Now()
		loc := t.Location()
		hour := t.Hour()
		y, m, d := t.Date()
		_ = &Time{Time: time.Date(y, m, d, hour, 0, 0, 0, loc), Config: nil}
	}
}

// 方案14: 使用 time.Now().Truncate 直接内联
func BenchmarkBeginningOfHour_InlinedTruncate(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_ = &Time{Time: time.Now().Truncate(time.Hour), Config: BeginningOfHourConfig}
	}
}

// 方案15: 零配置优化（最小内存分配）
func BenchmarkBeginningOfHour_ZeroAlloc(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		t := time.Now()
		_ = &Time{Time: t.Truncate(time.Hour), Config: nil}
	}
}

// 原始实现（用于对比）
func benchmarkBOM_Original(b *testing.B) {
	t := Now()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		y, m, _ := t.Date()
		_ = With(time.Date(y, m, 1, 0, 0, 0, 0, t.Location()))
	}
}

// 优化后的实现
func benchmarkBOM_Optimized(b *testing.B) {
	t := Now()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = &Time{
			Time:   time.Date(t.Year(), t.Month(), 1, 0, 0, 0, 0, t.Location()),
			Config: t.Config,
		}
	}
}

func BenchmarkBOM_Comparison_Original(b *testing.B) {
	b.ReportAllocs()
	benchmarkBOM_Original(b)
}

func BenchmarkBOM_Comparison_Optimized(b *testing.B) {
	b.ReportAllocs()
	benchmarkBOM_Optimized(b)
}

// BenchmarkBeginningOfMonth_Optimized 优化后的性能
func BenchmarkBeginningOfMonth_Optimized(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_ = BeginningOfMonth()
	}
}

// BenchmarkBeginningOfMonth_Original 原始实现的性能
func BenchmarkBeginningOfMonth_Original(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_ = With(time.Now()).BeginningOfMonth()
	}
}

// BenchmarkBeginningOfMonth_Variant1 V1: 当前实现
func BenchmarkBeginningOfMonth_Variant1(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_ = With(time.Now()).BeginningOfMonth()
	}
}

// BenchmarkBeginningOfMonth_Variant2 V2: 内联逻辑，完整 Config
func BenchmarkBeginningOfMonth_Variant2(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		now := time.Now()
		year, month, _ := now.Date()
		_ = &Time{
			Time: time.Date(year, month, 1, 0, 0, 0, 0, now.Location()),
			Config: &Config{
				WeekStartDay: time.Monday,
				TimeLocation: time.Local,
				TimeFormats:  []string{},
				Monotonic:    now,
			},
		}
	}
}

// BenchmarkBeginningOfMonth_Variant3 V3: 简化 Config
func BenchmarkBeginningOfMonth_Variant3(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		now := time.Now()
		year, month, _ := now.Date()
		_ = &Time{
			Time: time.Date(year, month, 1, 0, 0, 0, 0, now.Location()),
			Config: &Config{
				TimeLocation: now.Location(),
			},
		}
	}
}

// BenchmarkBeginningOfMonth_Variant4 V4: nil Config
func BenchmarkBeginningOfMonth_Variant4(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		now := time.Now()
		year, month, _ := now.Date()
		_ = &Time{
			Time:   time.Date(year, month, 1, 0, 0, 0, 0, now.Location()),
			Config: nil,
		}
	}
}

// BenchmarkBeginningOfMonth_Variant5 V5: 使用 Year/Month 方法 + nil Config
func BenchmarkBeginningOfMonth_Variant5(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		now := time.Now()
		_ = &Time{
			Time:   time.Date(now.Year(), now.Month(), 1, 0, 0, 0, 0, now.Location()),
			Config: nil,
		}
	}
}

// BenchmarkBeginningOfMonth_Variant6 V6: 最简化（Date + 无 Config）
func BenchmarkBeginningOfMonth_Variant6(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		now := time.Now()
		year, month, _ := now.Date()
		_ = &Time{Time: time.Date(year, month, 1, 0, 0, 0, 0, now.Location())}
	}
}

// BenchmarkBeginningOfMonth_Variant12 V12: 最优方案
func BenchmarkBeginningOfMonth_Variant12(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		now := time.Now()
		_ = &Time{Time: time.Date(now.Year(), now.Month(), 1, 0, 0, 0, 0, now.Location())}
	}
}

func BenchmarkBeginningOfMonth_Global_Optimized(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_ = BeginningOfMonth()
	}
}

func BenchmarkBeginningOfMonth_Global_Original(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_ = With(time.Now()).BeginningOfMonth()
	}
}

func BenchmarkBOM_NewImplementation(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_ = BeginningOfMinute()
	}
}

func BenchmarkBOM_NewImplementation_Parallel(b *testing.B) {
	b.ReportAllocs()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			_ = BeginningOfMinute()
		}
	})
}

func BenchmarkBOM_OldImplementation(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_ = With(time.Now()).BeginningOfMinute()
	}
}

// Baseline: 当前实现
func BenchmarkBeginningOfQuarter_Global_Current(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = BeginningOfQuarter()
	}
}

// 变体1: 直接计算，避免 With() 调用
func BenchmarkBeginningOfQuarter_Global_Variant1(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		now := time.Now()
		year := now.Year()
		month := int(now.Month())
		quarterStartMonth := ((month - 1) / 3) * 3
		_ = &Time{
			Time: time.Date(year, time.Month(quarterStartMonth+1), 1, 0, 0, 0, 0, now.Location()),
			Config: &Config{
				WeekStartDay: time.Monday,
				TimeLocation: time.Local,
				TimeFormats:  []string{},
				Monotonic:    now,
			},
		}
	}
}

// 变体2: 使用 time.Date() 直接构造，减少中间变量
func BenchmarkBeginningOfQuarter_Global_Variant2(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		now := time.Now()
		month := int(now.Month())
		quarterStartMonth := ((month - 1) / 3) * 3
		_ = &Time{
			Time: time.Date(now.Year(), time.Month(quarterStartMonth+1), 1, 0, 0, 0, 0, now.Location()),
			Config: &Config{
				WeekStartDay: time.Monday,
				TimeLocation: time.Local,
				TimeFormats:  []string{},
				Monotonic:    now,
			},
		}
	}
}

// 变体3: 预先创建 Config，每次只创建 Time
func BenchmarkBeginningOfQuarter_Global_Variant3(b *testing.B) {
	config := &Config{
		WeekStartDay: time.Monday,
		TimeLocation: time.Local,
		TimeFormats:  []string{},
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		now := time.Now()
		month := int(now.Month())
		quarterStartMonth := ((month - 1) / 3) * 3
		_ = &Time{
			Time:   time.Date(now.Year(), time.Month(quarterStartMonth+1), 1, 0, 0, 0, 0, now.Location()),
			Config: config,
		}
	}
}

// 变体4: 使用 time.Now().Date() 获取年月日
func BenchmarkBeginningOfQuarter_Global_Variant4(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		now := time.Now()
		year, month, _ := now.Date()
		quarterStartMonth := ((int(month) - 1) / 3) * 3
		_ = &Time{
			Time: time.Date(year, time.Month(quarterStartMonth+1), 1, 0, 0, 0, 0, now.Location()),
			Config: &Config{
				WeekStartDay: time.Monday,
				TimeLocation: time.Local,
				TimeFormats:  []string{},
				Monotonic:    now,
			},
		}
	}
}

// 变体5: 使用查找表
func BenchmarkBeginningOfQuarter_Global_Variant5(b *testing.B) {
	quarterStartMonths := [12]int{0, 0, 0, 3, 3, 3, 6, 6, 6, 9, 9, 9}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		now := time.Now()
		month := now.Month()
		quarterStartMonth := quarterStartMonths[month-1]
		_ = &Time{
			Time: time.Date(now.Year(), time.Month(quarterStartMonth+1), 1, 0, 0, 0, 0, now.Location()),
			Config: &Config{
				WeekStartDay: time.Monday,
				TimeLocation: time.Local,
				TimeFormats:  []string{},
				Monotonic:    now,
			},
		}
	}
}

// 变体6: 使用 switch-case
func BenchmarkBeginningOfQuarter_Global_Variant6(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		now := time.Now()
		month := now.Month()
		var startMonth time.Month
		switch month {
		case time.January, time.February, time.March:
			startMonth = time.January
		case time.April, time.May, time.June:
			startMonth = time.April
		case time.July, time.August, time.September:
			startMonth = time.July
		case time.October, time.November, time.December:
			startMonth = time.October
		}
		_ = &Time{
			Time: time.Date(now.Year(), startMonth, 1, 0, 0, 0, 0, now.Location()),
			Config: &Config{
				WeekStartDay: time.Monday,
				TimeLocation: time.Local,
				TimeFormats:  []string{},
				Monotonic:    now,
			},
		}
	}
}

// 变体7: 使用 if-else 链
func BenchmarkBeginningOfQuarter_Global_Variant7(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		now := time.Now()
		month := now.Month()
		var startMonth time.Month
		if month <= time.March {
			startMonth = time.January
		} else if month <= time.June {
			startMonth = time.April
		} else if month <= time.September {
			startMonth = time.July
		} else {
			startMonth = time.October
		}
		_ = &Time{
			Time: time.Date(now.Year(), startMonth, 1, 0, 0, 0, 0, now.Location()),
			Config: &Config{
				WeekStartDay: time.Monday,
				TimeLocation: time.Local,
				TimeFormats:  []string{},
				Monotonic:    now,
			},
		}
	}
}

// 变体8: 使用位运算计算季度
func BenchmarkBeginningOfQuarter_Global_Variant8(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		now := time.Now()
		month := int(now.Month())
		quarterStartMonth := ((month - 1) / 3) * 3
		_ = &Time{
			Time: time.Date(now.Year(), time.Month(quarterStartMonth+1), 1, 0, 0, 0, 0, now.Location()),
			Config: &Config{
				WeekStartDay: time.Monday,
				TimeLocation: time.Local,
				TimeFormats:  []string{},
				Monotonic:    now,
			},
		}
	}
}

// 变体9: 复用 time.Now() 的结果
func BenchmarkBeginningOfQuarter_Global_Variant9(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		now := time.Now()
		loc := now.Location()
		year := now.Year()
		month := now.Month()
		quarterStartMonth := ((int(month) - 1) / 3) * 3
		_ = &Time{
			Time: time.Date(year, time.Month(quarterStartMonth+1), 1, 0, 0, 0, 0, loc),
			Config: &Config{
				WeekStartDay: time.Monday,
				TimeLocation: loc,
				TimeFormats:  []string{},
				Monotonic:    now,
			},
		}
	}
}

// 变体10: 最简版本 - 只计算必要字段
func BenchmarkBeginningOfQuarter_Global_Variant10(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		now := time.Now()
		month := int(now.Month())
		_ = &Time{
			Time: time.Date(now.Year(), time.Month(((month-1)/3)*3+1), 1, 0, 0, 0, 0, now.Location()),
			Config: &Config{
				WeekStartDay: time.Monday,
				TimeLocation: time.Local,
				TimeFormats:  []string{},
				Monotonic:    now,
			},
		}
	}
}

// 变体11: 使用 nil Config（如果适用）
func BenchmarkBeginningOfQuarter_Global_Variant11(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		now := time.Now()
		month := int(now.Month())
		_ = &Time{
			Time:   time.Date(now.Year(), time.Month(((month-1)/3)*3+1), 1, 0, 0, 0, 0, now.Location()),
			Config: nil, // 延迟初始化
		}
	}
}

// 方案0: 当前实现
func BenchmarkBeginningOfWeekGlobalCurrent(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = BeginningOfWeek()
	}
}

// 方案1: 内联逻辑，避免 With() 调用
func BenchmarkBeginningOfWeekGlobalOpt1(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		now := time.Now()
		year, month, day := now.Date()
		loc := now.Location()
		midnight := time.Date(year, month, day, 0, 0, 0, 0, loc)
		weekday := int(midnight.Weekday())
		_ = &Time{
			Time:   midnight.AddDate(0, 0, -weekday),
			Config: &Config{WeekStartDay: time.Sunday, TimeLocation: time.Local},
		}
	}
}

// 方案2: 使用全局 Config
var bowGlobalConfig = &Config{
	WeekStartDay: time.Sunday,
	TimeLocation: time.Local,
}

func BenchmarkBeginningOfWeekGlobalOpt2(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		now := time.Now()
		midnight := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, now.Location())
		weekday := int(midnight.Weekday())
		_ = &Time{
			Time:   midnight.AddDate(0, 0, -weekday),
			Config: bowGlobalConfig,
		}
	}
}

// 方案3: 最小化变量
func BenchmarkBeginningOfWeekGlobalOpt3(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		t := time.Now()
		_ = &Time{
			Time:   time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, t.Location()).AddDate(0, 0, -int(t.Weekday())),
			Config: bowGlobalConfig,
		}
	}
}

// 方案4: 预计算 time.Local
func BenchmarkBeginningOfWeekGlobalOpt4(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		t := time.Now()
		loc := time.Local
		_ = &Time{
			Time:   time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, loc).AddDate(0, 0, -int(t.Weekday())),
			Config: bowGlobalConfig,
		}
	}
}

// 方案5: 避免重复 weekday 计算
func BenchmarkBeginningOfWeekGlobalOpt5(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		t := time.Now()
		weekday := int(t.Weekday())
		_ = &Time{
			Time:   time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, t.Location()).AddDate(0, 0, -weekday),
			Config: bowGlobalConfig,
		}
	}
}

// 方案6: 组合优化
func BenchmarkBeginningOfWeekGlobalOpt6(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		t := time.Now()
		loc := time.Local
		wd := int(t.Weekday())
		_ = &Time{
			Time:   time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, loc).AddDate(0, 0, -wd),
			Config: bowGlobalConfig,
		}
	}
}

// 方案7: 使用 sync.Pool
var bowTimePool = &sync.Pool{
	New: func() interface{} {
		return &Time{
			Config: &Config{
				WeekStartDay: time.Sunday,
				TimeLocation: time.Local,
			},
		}
	},
}

func BenchmarkBeginningOfWeekGlobalOpt7(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		t := time.Now()
		result := bowTimePool.Get().(*Time)
		result.Time = time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, t.Location()).AddDate(0, 0, -int(t.Weekday()))
		_ = result
	}
}

// 方案8: sync.Pool + 预计算变量
func BenchmarkBeginningOfWeekGlobalOpt8(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		t := time.Now()
		loc := time.Local
		wd := int(t.Weekday())
		result := bowTimePool.Get().(*Time)
		result.Time = time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, loc).AddDate(0, 0, -wd)
		_ = result
	}
}

// 方案9: 延迟初始化 Config
var (
	bowLazyConfig     *Config
	bowLazyConfigOnce sync.Once
)

func getBowLazyConfig() *Config {
	bowLazyConfigOnce.Do(func() {
		bowLazyConfig = &Config{
			WeekStartDay: time.Sunday,
			TimeLocation: time.Local,
		}
	})
	return bowLazyConfig
}

func BenchmarkBeginningOfWeekGlobalOpt9(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		t := time.Now()
		loc := time.Local
		wd := int(t.Weekday())
		_ = &Time{
			Time:   time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, loc).AddDate(0, 0, -wd),
			Config: getBowLazyConfig(),
		}
	}
}

// 方案10: 直接构造（最小化调用）
func BenchmarkBeginningOfWeekGlobalOpt10(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		t := time.Now()
		y, m, d := t.Date()
		_ = &Time{
			Time:   time.Date(y, m, d, 0, 0, 0, 0, t.Location()).AddDate(0, 0, -int(t.Weekday())),
			Config: bowGlobalConfig,
		}
	}
}

// 方案1: 当前实现（Baseline）
func BenchmarkBeginningOfYear_Global_Current(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_ = BeginningOfYear_Global_V1()
	}
}

func BeginningOfYear_Global_V1() *Time {
	return With(time.Now()).BeginningOfYear()
}

// 方案2: 内联 With + BeginningOfYear 逻辑
func BenchmarkBeginningOfYear_Global_V2(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_ = BeginningOfYear_Global_V2()
	}
}

func BeginningOfYear_Global_V2() *Time {
	now := time.Now()
	year := now.Year()
	janFirst := time.Date(year, time.January, 1, 0, 0, 0, 0, now.Location())
	return &Time{
		Time: janFirst,
		Config: &Config{
			WeekStartDay: time.Monday,
			TimeLocation: time.Local,
			TimeFormats:  []string{},
			Monotonic:    now,
		},
	}
}

// 方案3: 简化 Config，只设置必要字段
func BenchmarkBeginningOfYear_Global_V3(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_ = BeginningOfYear_Global_V3()
	}
}

func BeginningOfYear_Global_V3() *Time {
	now := time.Now()
	return &Time{
		Time: time.Date(now.Year(), time.January, 1, 0, 0, 0, 0, now.Location()),
		Config: &Config{
			TimeLocation: now.Location(),
		},
	}
}

// 方案4: 零 Config（使用 nil）
func BenchmarkBeginningOfYear_Global_V4(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_ = BeginningOfYear_Global_V4()
	}
}

func BeginningOfYear_Global_V4() *Time {
	now := time.Now()
	return &Time{
		Time:   time.Date(now.Year(), time.January, 1, 0, 0, 0, 0, now.Location()),
		Config: nil,
	}
}

// 方案5: 使用全局 Config（避免每次分配）
var BeginningOfYearConfig = &Config{
	WeekStartDay: time.Monday,
	TimeLocation: time.Local,
	TimeFormats:  []string{},
}

func BenchmarkBeginningOfYear_Global_V5(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_ = BeginningOfYear_Global_V5()
	}
}

func BeginningOfYear_Global_V5() *Time {
	now := time.Now()
	return &Time{
		Time:   time.Date(now.Year(), time.January, 1, 0, 0, 0, 0, now.Location()),
		Config: BeginningOfYearConfig,
	}
}

// 方案6: 预先计算 Year，避免多次调用 now.Year()
func BenchmarkBeginningOfYear_Global_V6(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_ = BeginningOfYear_Global_V6()
	}
}

func BeginningOfYear_Global_V6() *Time {
	now := time.Now()
	year := now.Year()
	return &Time{
		Time:   time.Date(year, time.January, 1, 0, 0, 0, 0, now.Location()),
		Config: BeginningOfYearConfig,
	}
}

// 方案7: 省略 time.Date 的零值参数（秒、纳秒）
func BenchmarkBeginningOfYear_Global_V7(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_ = BeginningOfYear_Global_V7()
	}
}

func BeginningOfYear_Global_V7() *Time {
	now := time.Now()
	year := now.Year()
	return &Time{
		Time:   time.Date(year, time.January, 1, 0, 0, 0, 0, now.Location()),
		Config: BeginningOfYearConfig,
	}
}

// 方案8: 使用 time.Local 代替 now.Location()（假设本地时区）
func BenchmarkBeginningOfYear_Global_V8(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_ = BeginningOfYear_Global_V8()
	}
}

func BeginningOfYear_Global_V8() *Time {
	now := time.Now()
	year := now.Year()
	return &Time{
		Time:   time.Date(year, time.January, 1, 0, 0, 0, 0, time.Local),
		Config: BeginningOfYearConfig,
	}
}

// 方案9: 完全省略 Config（零分配）
func BenchmarkBeginningOfYear_Global_V9(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_ = BeginningOfYear_Global_V9()
	}
}

func BeginningOfYear_Global_V9() *Time {
	now := time.Now()
	year := now.Year()
	return &Time{Time: time.Date(year, time.January, 1, 0, 0, 0, 0, now.Location())}
}

// 方案10: 直接返回 time.Time，包装为 Time（最简洁）
func BenchmarkBeginningOfYear_Global_V10(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_ = BeginningOfYear_Global_V10()
	}
}

func BeginningOfYear_Global_V10() *Time {
	now := time.Now()
	year := now.Year()
	janFirst := time.Date(year, time.January, 1, 0, 0, 0, 0, now.Location())
	return &Time{Time: janFirst}
}

// 方案11: 使用 sync.Pool 复用 Time 对象
var boyTimePool = sync.Pool{
	New: func() interface{} {
		return &Time{}
	},
}

func BenchmarkBeginningOfYear_Global_V11(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_ = BeginningOfYear_Global_V11()
	}
}

func BeginningOfYear_Global_V11() *Time {
	t := boyTimePool.Get().(*Time)
	now := time.Now()
	year := now.Year()
	t.Time = time.Date(year, time.January, 1, 0, 0, 0, 0, now.Location())
	t.Config = nil
	result := *t
	boyTimePool.Put(t)
	return &result
}

// 方案12: 组合优化：省略中间变量
func BenchmarkBeginningOfYear_Global_V12(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_ = BeginningOfYear_Global_V12()
	}
}

func BeginningOfYear_Global_V12() *Time {
	now := time.Now()
	return &Time{Time: time.Date(now.Year(), time.January, 1, 0, 0, 0, 0, now.Location())}
}

// 方案13: 使用 Truncate - 存在闰年问题，仅作性能参考
func BenchmarkBeginningOfYear_Global_V13(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_ = BeginningOfYear_Global_V13()
	}
}

func BeginningOfYear_Global_V13() *Time {
	now := time.Now()
	yearDuration := time.Hour * 24 * 365
	truncated := now.Truncate(yearDuration)
	return &Time{
		Time:   truncated,
		Config: BeginningOfYearConfig,
	}
}

// 方案14: 使用 AddDate 负数回退（不推荐，性能差）
func BenchmarkBeginningOfYear_Global_V14(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_ = BeginningOfYear_Global_V14()
	}
}

func BeginningOfYear_Global_V14() *Time {
	now := time.Now()
	yearStart := now.AddDate(0, 1-int(now.Month()), 1-int(now.Day()))
	return &Time{
		Time:   yearStart,
		Config: BeginningOfYearConfig,
	}
}

// 方案15: 单次 time.Now() 调用 + 内联所有逻辑
func BenchmarkBeginningOfYear_Global_V15(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_ = BeginningOfYear_Global_V15()
	}
}

func BeginningOfYear_Global_V15() *Time {
	now := time.Now()
	loc := now.Location()
	year := now.Year()
	return &Time{
		Time: time.Date(year, time.January, 1, 0, 0, 0, 0, loc),
	}
}

func BenchmarkBeginningOfYear_Optimized(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_ = BeginningOfYear()
	}
}

// 方案1: 当前实现（baseline）
func BenchmarkEndOfHour_Current(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_ = EndOfHour()
	}
}

// 方案2: 直接使用 time.Date 构建
func BenchmarkEndOfHour_DirectDate(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		now := time.Now()
		result := time.Date(now.Year(), now.Month(), now.Day(), now.Hour(), 59, 59, 999999999, now.Location())
		_ = With(result)
	}
}

// 方案3: 预先计算常量
func BenchmarkEndOfHour_PreComputed(b *testing.B) {
	b.ReportAllocs()
	const endMinute = 59
	const endSecond = 59
	const endNano = 999999999

	for i := 0; i < b.N; i++ {
		now := time.Now()
		result := time.Date(now.Year(), now.Month(), now.Day(), now.Hour(), endMinute, endSecond, endNano, now.Location())
		_ = With(result)
	}
}

// 方案4: 使用 Truncate 后加 1 小时减 1 纳秒
func BenchmarkEndOfHour_Truncate(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		now := time.Now()
		truncated := now.Truncate(time.Hour)
		result := truncated.Add(time.Hour - time.Nanosecond)
		_ = With(result)
	}
}

// 方案5: 使用 Add 替代部分 Date 调用
func BenchmarkEndOfHour_AddVersion(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		now := time.Now()
		truncated := now.Truncate(time.Hour)
		result := truncated.Add(59*time.Minute + 59*time.Second + 999999999*time.Nanosecond)
		_ = With(result)
	}
}

// 方案6: 内联 With 逻辑
func BenchmarkEndOfHour_InlineWith(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		now := time.Now()
		result := time.Date(now.Year(), now.Month(), now.Day(), now.Hour(), 59, 59, 999999999, now.Location())
		_ = &Time{
			Time: result,
			Config: &Config{
				WeekStartDay: time.Monday,
				TimeLocation: time.Local,
				TimeFormats:  []string{},
				Monotonic:    time.Now(),
			},
		}
	}
}

// 方案7: 单次 time.Now() 调用（用于 Monotonic）
func BenchmarkEndOfHour_SingleTimeNow(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		now := time.Now()
		result := time.Date(now.Year(), now.Month(), now.Day(), now.Hour(), 59, 59, 999999999, now.Location())
		_ = &Time{
			Time: result,
			Config: &Config{
				WeekStartDay: time.Monday,
				TimeLocation: time.Local,
				TimeFormats:  []string{},
				Monotonic:    now,
			},
		}
	}
}

// 方案8: 使用全局 Config
func BenchmarkEndOfHour_GlobalConfig(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		now := time.Now()
		result := time.Date(now.Year(), now.Month(), now.Day(), now.Hour(), 59, 59, 999999999, now.Location())
		_ = &Time{
			Time:   result,
			Config: BeginningOfHourConfig,
		}
	}
}

// 方案9: 零分配版本（直接返回 Time，不分配 Config）
func BenchmarkEndOfHour_ZeroAlloc(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		now := time.Now()
		result := time.Date(now.Year(), now.Month(), now.Day(), now.Hour(), 59, 59, 999999999, now.Location())
		_ = &Time{Time: result}
	}
}

// 方案10: 复用 BeginningOfHour 逻辑
func BenchmarkEndOfHour_ReuseBeginning(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		now := time.Now()
		y, m, d := now.Date()
		beginning := time.Date(y, m, d, now.Hour(), 0, 0, 0, now.Location())
		result := beginning.Add(time.Hour - time.Nanosecond)
		_ = With(result)
	}
}

// 方案11: 使用 BeginningOfHour() 函数
func BenchmarkEndOfHour_CallBeginningOfHour(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		now := time.Now()
		beginning := With(now).BeginningOfHour()
		result := beginning.Time.Add(time.Hour - time.Nanosecond)
		_ = With(result)
	}
}

// 方案12: 内联 BeginningOfHour 逻辑 + Add
func BenchmarkEndOfHour_InlineBeginningAdd(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		now := time.Now()
		y, m, d := now.Date()
		beginning := time.Date(y, m, d, now.Hour(), 0, 0, 0, now.Location())
		result := beginning.Add(time.Hour - time.Nanosecond)
		_ = &Time{
			Time:   result,
			Config: BeginningOfHourConfig,
		}
	}
}

// 方案13: 预计算 Hour - 1ns 常量
func BenchmarkEndOfHour_PreComputedHourMinusNs(b *testing.B) {
	b.ReportAllocs()
	const hourMinusNs = time.Hour - time.Nanosecond

	for i := 0; i < b.N; i++ {
		now := time.Now()
		truncated := now.Truncate(time.Hour)
		result := truncated.Add(hourMinusNs)
		_ = &Time{
			Time:   result,
			Config: BeginningOfHourConfig,
		}
	}
}

// 方案14: 使用 Truncate + 全局 Config
func BenchmarkEndOfHour_TruncateWithGlobalConfig(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		now := time.Now()
		truncated := now.Truncate(time.Hour)
		result := truncated.Add(time.Hour - time.Nanosecond)
		_ = &Time{
			Time:   result,
			Config: BeginningOfHourConfig,
		}
	}
}

// 方案15: 完全内联版本
func BenchmarkEndOfHour_FullyInline(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		now := time.Now()
		result := time.Date(now.Year(), now.Month(), now.Day(), now.Hour(), 59, 59, 999999999, now.Location())
		_ = &Time{
			Time: result,
			Config: &Config{
				WeekStartDay: time.Monday,
				TimeLocation: time.Local,
				TimeFormats:  []string{},
			},
		}
	}
}

// 方案1: 当前实现（baseline）
func BenchmarkEndOfMinute_Current(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_ = With(time.Now()).EndOfMinute()
	}
}

// 方案2: 直接使用 time.Date 构建
func BenchmarkEndOfMinute_DirectDate(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		now := time.Now()
		result := time.Date(now.Year(), now.Month(), now.Day(), now.Hour(), now.Minute(), 59, 999999999, now.Location())
		_ = With(result)
	}
}

// 方案3: 预先计算常量
func BenchmarkEndOfMinute_PreComputed(b *testing.B) {
	b.ReportAllocs()
	const endSecond = 59
	const endNano = 999999999

	for i := 0; i < b.N; i++ {
		now := time.Now()
		result := time.Date(now.Year(), now.Month(), now.Day(), now.Hour(), now.Minute(), endSecond, endNano, now.Location())
		_ = With(result)
	}
}

// 方案4: 使用 Truncate 后加 1 分钟减 1 纳秒
func BenchmarkEndOfMinute_Truncate(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		now := time.Now()
		truncated := now.Truncate(time.Minute)
		result := truncated.Add(time.Minute - time.Nanosecond)
		_ = With(result)
	}
}

// 方案5: 使用 Add 替代部分 Date 调用
func BenchmarkEndOfMinute_AddVersion(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		now := time.Now()
		truncated := now.Truncate(time.Minute)
		result := truncated.Add(59*time.Second + 999999999*time.Nanosecond)
		_ = With(result)
	}
}

// 方案6: 内联 With 逻辑
func BenchmarkEndOfMinute_InlineWith(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		now := time.Now()
		result := time.Date(now.Year(), now.Month(), now.Day(), now.Hour(), now.Minute(), 59, 999999999, now.Location())
		_ = &Time{
			Time: result,
			Config: &Config{
				WeekStartDay: time.Monday,
				TimeLocation: time.Local,
				TimeFormats:  []string{},
				Monotonic:    time.Now(),
			},
		}
	}
}

// 方案7: 单次 time.Now() 调用（用于 Monotonic）
func BenchmarkEndOfMinute_SingleTimeNow(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		now := time.Now()
		result := time.Date(now.Year(), now.Month(), now.Day(), now.Hour(), now.Minute(), 59, 999999999, now.Location())
		_ = &Time{
			Time: result,
			Config: &Config{
				WeekStartDay: time.Monday,
				TimeLocation: time.Local,
				TimeFormats:  []string{},
				Monotonic:    now,
			},
		}
	}
}

// 方案8: 使用 Unix 时间戳计算
func BenchmarkEndOfMinute_Unix(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		now := time.Now()
		nanos := now.UnixNano()
		minuteNanos := int64(time.Minute)
		aligned := (nanos / minuteNanos) * minuteNanos
		result := time.Unix(0, aligned+minuteNanos-1).In(now.Location())
		_ = With(result)
	}
}

// 方案9: 最简版本（nil TimeFormats，无 Monotonic）
func BenchmarkEndOfMinute_Minimal(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		now := time.Now()
		result := time.Date(now.Year(), now.Month(), now.Day(), now.Hour(), now.Minute(), 59, 999999999, now.Location())
		_ = &Time{
			Time: result,
			Config: &Config{
				WeekStartDay: time.Monday,
				TimeLocation: time.Local,
				TimeFormats:  nil,
			},
		}
	}
}

// 方案10: 组合优化（单次 time.Now + 预计算常量 + nil TimeFormats）
func BenchmarkEndOfMinute_Combined(b *testing.B) {
	b.ReportAllocs()
	const (
		endSecond    = 59
		endNano      = 999999999
		weekStartDay = time.Monday
	)

	for i := 0; i < b.N; i++ {
		now := time.Now()
		result := time.Date(now.Year(), now.Month(), now.Day(), now.Hour(), now.Minute(), endSecond, endNano, now.Location())
		_ = &Time{
			Time: result,
			Config: &Config{
				WeekStartDay: weekStartDay,
				TimeLocation: time.Local,
				TimeFormats:  nil,
				Monotonic:    now,
			},
		}
	}
}

// 方案11: 使用 time.Add 直接计算
func BenchmarkEndOfMinute_AddDirect(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		now := time.Now()
		nextMinute := now.Truncate(time.Minute).Add(time.Minute)
		result := nextMinute.Add(-time.Nanosecond)
		_ = With(result)
	}
}

// 方案12: 优化的 Truncate 版本（单次 time.Now）
func BenchmarkEndOfMinute_OptimizedTruncate(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		now := time.Now()
		result := now.Truncate(time.Minute).Add(time.Minute - time.Nanosecond)
		_ = &Time{
			Time: result,
			Config: &Config{
				WeekStartDay: time.Monday,
				TimeLocation: time.Local,
				TimeFormats:  nil,
				Monotonic:    now,
			},
		}
	}
}

// 方案13: 完全内联版本
func BenchmarkEndOfMinute_FullyInline(b *testing.B) {
	b.ReportAllocs()
	const (
		endSecond    = 59
		endNano      = 999999999
		weekStartDay = time.Monday
	)

	for i := 0; i < b.N; i++ {
		now := time.Now()
		result := time.Date(now.Year(), now.Month(), now.Day(), now.Hour(), now.Minute(), endSecond, endNano, now.Location())
		t := &Time{
			Time: result,
			Config: &Config{
				WeekStartDay: weekStartDay,
				TimeLocation: time.Local,
				TimeFormats:  nil,
				Monotonic:    now,
			},
		}
		_ = t
	}
}

// 方案14: 使用 time.Date 但避免重复字段提取
func BenchmarkEndOfMinute_FieldReuse(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		now := time.Now()
		y, m, d := now.Date()
		h, min, _ := now.Clock()
		result := time.Date(y, m, d, h, min, 59, 999999999, now.Location())
		_ = With(result)
	}
}

// 验证新实现 vs 旧实现性能
func BenchmarkEndOfMinute_NewImplementation(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_ = EndOfMinute()
	}
}

// 旧实现（用于对比）
func BenchmarkEndOfMinute_OldImplementation(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_ = With(time.Now()).EndOfMinute()
	}
}

// 使用固定时间测试（避免 time.Now() 开销）
func BenchmarkEndOfMinute_FixedTime(b *testing.B) {
	b.ReportAllocs()
	testTime := time.Date(2024, 6, 15, 14, 30, 45, 123456789, time.Local)
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		now := testTime
		result := now.Truncate(time.Minute).Add(time.Minute - time.Nanosecond)
		_ = &Time{
			Time: result,
			Config: &Config{
				WeekStartDay: time.Monday,
				TimeLocation: time.Local,
				TimeFormats:  nil,
				Monotonic:    now,
			},
		}
	}
}

// 旧实现（固定时间）
func BenchmarkEndOfMinute_OldFixedTime(b *testing.B) {
	b.ReportAllocs()
	testTime := time.Date(2024, 6, 15, 14, 30, 45, 123456789, time.Local)
	t := With(testTime)
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_ = t.EndOfMinute()
	}
}

// 保存原始实现用于性能对比
func (p *Time) EndOfHalf_Original() *Time {
	return With(p.BeginningOfHalf().AddDate(0, 6, 0).Add(-time.Nanosecond))
}

func BenchmarkEndOfHalf_Original(b *testing.B) {
	b.ReportAllocs()
	t := MustParse("2024-03-15 14:30:45")
	for i := 0; i < b.N; i++ {
		_ = t.EndOfHalf_Original()
	}
}

func BenchmarkEndOfHalf_Optimized(b *testing.B) {
	b.ReportAllocs()
	t := MustParse("2024-03-15 14:30:45")
	for i := 0; i < b.N; i++ {
		_ = t.EndOfHalf()
	}
}

func BenchmarkEndOfHalf_Original_H2(b *testing.B) {
	b.ReportAllocs()
	t := MustParse("2024-09-15 14:30:45")
	for i := 0; i < b.N; i++ {
		_ = t.EndOfHalf_Original()
	}
}

func BenchmarkEndOfHalf_Optimized_H2(b *testing.B) {
	b.ReportAllocs()
	t := MustParse("2024-09-15 14:30:45")
	for i := 0; i < b.N; i++ {
		_ = t.EndOfHalf()
	}
}

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

func Benchmark_EndOfMonth_Optimized(b *testing.B) {
	t := Now()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_ = t.EndOfMonth()
	}
}

func Benchmark_EndOfMonth_Final(b *testing.B) {
	t := &Time{Time: time.Date(2024, 2, 15, 12, 0, 0, 0, time.Local)}
	b.ResetTimer()
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		_ = t.EndOfMonth()
	}
}

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

func Benchmark_EndOfWeek_Optimized(b *testing.B) {
	t := Now()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_ = t.EndOfWeek()
	}
}

func Benchmark_EndOfWeek_Optimized_Small(b *testing.B) {
	t := time.Date(2024, 1, 15, 12, 30, 45, 123456789, time.Local)
	bt := &Time{Time: t}
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_ = bt.EndOfWeek()
	}
}

func Benchmark_EndOfWeek_Optimized_Medium(b *testing.B) {
	t := time.Date(2024, 6, 15, 12, 30, 45, 123456789, time.Local)
	bt := &Time{Time: t}
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_ = bt.EndOfWeek()
	}
}

func Benchmark_EndOfWeek_Optimized_Large(b *testing.B) {
	t := time.Date(2024, 12, 31, 23, 59, 59, 999999999, time.Local)
	bt := &Time{Time: t}
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_ = bt.EndOfWeek()
	}
}

func Benchmark_EndOfWeek_Optimized_Parallel(b *testing.B) {
	t := Now()
	b.ResetTimer()

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			_ = t.EndOfWeek()
		}
	})
}

func Benchmark_EndOfWeek_Optimized_WithConfig(b *testing.B) {
	t := Now()
	t.Config = &Config{WeekStartDay: time.Monday}
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_ = t.EndOfWeek()
	}
}

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

func BenchmarkEndOfYear_Simple(b *testing.B) {
	t := With(time.Date(2024, 6, 15, 14, 30, 45, 123456789, time.Local))
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_ = t.EndOfYear()
	}
}

func BenchmarkEndOfYear_Optimized(b *testing.B) {
	t := With(time.Date(2024, 6, 15, 14, 30, 45, 123456789, time.Local))
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		year := t.Time.Year()
		loc := t.Time.Location()
		config := t.Config
		if config == nil {
			config = &Config{}
		}
		// time.Date(0) 溢出技巧
		end := time.Date(year+1, time.January, 0, 23, 59, 59, 999999999, loc)
		_ = &Time{Time: end, Config: config}
	}
}

// 测试实际的 EndOfDay 方法（优化后）
func BenchmarkEOD_ActualMethod(b *testing.B) {
	times := genEODTestTimes(1000)
	wrapper := make([]*Time, len(times))
	for i := range times {
		wrapper[i] = With(times[i])
	}
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_ = wrapper[i%len(wrapper)].EndOfDay()
	}
}

// 对比旧实现
func BenchmarkEOD_OldMethod(b *testing.B) {
	times := genEODTestTimes(1000)
	wrapper := make([]*Time, len(times))
	for i := range times {
		wrapper[i] = With(times[i])
	}
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		t := wrapper[i%len(wrapper)]
		y, m, d := t.Date()
		_ = With(time.Date(y, m, d, 23, 59, 59, int(time.Second-time.Nanosecond), t.Location()))
	}
}

// 验证 Config 复用导致的内存分配
func BenchmarkEOD_DirectConstructNoAlloc(b *testing.B) {
	times := genEODTestTimes(1000)
	wrapper := make([]*Time, len(times))
	for i := range times {
		wrapper[i] = With(times[i])
	}
	b.ResetTimer()
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		t := wrapper[i%len(wrapper)]
		loc := t.Location()
		year, month, day := t.Date()
		eod := time.Date(year, month, day, 23, 59, 59, int(time.Second-time.Nanosecond), loc)
		// 不创建 Config，直接复用
		_ = &Time{Time: eod, Config: t.Config}
	}
}

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

func BenchmarkEOD_OldImplementation(b *testing.B) {
	times := genEODTestTimes(1000)
	wrapper := make([]*Time, len(times))
	for i := range times {
		wrapper[i] = With(times[i])
	}
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_ = endOfDayOld(wrapper[i%len(wrapper)])
	}
}

func BenchmarkEOD_NewImplementation(b *testing.B) {
	times := genEODTestTimes(1000)
	wrapper := make([]*Time, len(times))
	for i := range times {
		wrapper[i] = With(times[i])
	}
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_ = endOfDayNew(wrapper[i%len(wrapper)])
	}
}

// BenchmarkEndOfDay_Optimized 优化后的实现
func BenchmarkEndOfDay_Optimized(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_ = EndOfDay()
	}
}

// BenchmarkEndOfDay_Baseline 原始实现（用于对比）
func BenchmarkEndOfDay_Baseline(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_ = With(time.Now()).EndOfDay()
	}
}

// BenchmarkEndOfDay_Manual 手动优化版本
func BenchmarkEndOfDay_Manual(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		now := time.Now()
		year, month, day := now.Date()
		eod := time.Date(year, month, day, 23, 59, 59, int(time.Second-time.Nanosecond), now.Location())
		_ = &Time{Time: eod}
	}
}

// =============================================================================
// 方案1: 当前实现（Baseline） - With(time.Now()).EndOfMonth()
// =============================================================================

func BenchmarkEndOfMonth_Global_V1_Current(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_ = EndOfMonth_Global_V1()
	}
}

func EndOfMonth_Global_V1() *Time {
	return With(time.Now()).EndOfMonth()
}

// =============================================================================
// 方案2: 内联逻辑 - 复用 With 和 EndOfMonth 的完整逻辑
// =============================================================================

func BenchmarkEndOfMonth_Global_V2(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_ = EndOfMonth_Global_V2()
	}
}

func EndOfMonth_Global_V2() *Time {
	now := time.Now()
	year, month, _ := now.Date()
	eom := time.Date(year, month+1, 0, 23, 59, 59, 999999999, now.Location())
	return &Time{
		Time: eom,
		Config: &Config{
			WeekStartDay: time.Monday,
			TimeLocation: time.Local,
			TimeFormats:  []string{},
			Monotonic:    now,
		},
	}
}

// =============================================================================
// 方案3: 简化 Config - 只设置必要字段
// =============================================================================

func BenchmarkEndOfMonth_Global_V3(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_ = EndOfMonth_Global_V3()
	}
}

func EndOfMonth_Global_V3() *Time {
	now := time.Now()
	year, month, _ := now.Date()
	return &Time{
		Time: time.Date(year, month+1, 0, 23, 59, 59, 999999999, now.Location()),
		Config: &Config{
			TimeLocation: now.Location(),
		},
	}
}

// =============================================================================
// 方案4: 零 Config - 使用 nil Config
// =============================================================================

func BenchmarkEndOfMonth_Global_V4(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_ = EndOfMonth_Global_V4()
	}
}

func EndOfMonth_Global_V4() *Time {
	now := time.Now()
	year, month, _ := now.Date()
	return &Time{
		Time:   time.Date(year, month+1, 0, 23, 59, 59, 999999999, now.Location()),
		Config: nil,
	}
}

// =============================================================================
// 方案5: 预计算常量 - 使用预定义的月末时间常量
// =============================================================================

func BenchmarkEndOfMonth_Global_V5(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_ = EndOfMonth_Global_V5()
	}
}

var endOfDayNanos = int64(999999999)

func EndOfMonth_Global_V5() *Time {
	now := time.Now()
	year, month, _ := now.Date()
	return &Time{
		Time:   time.Date(year, month+1, 0, 23, 59, 59, int(endOfDayNanos), now.Location()),
		Config: nil,
	}
}

// =============================================================================
// 方案6: 直接构造 Time struct - 最小化操作
// =============================================================================

func BenchmarkEndOfMonth_Global_V6(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_ = EndOfMonth_Global_V6()
	}
}

func EndOfMonth_Global_V6() *Time {
	now := time.Now()
	year, month, _ := now.Date()
	return &Time{
		Time:   time.Date(year, month+1, 0, 23, 59, 59, 999999999, now.Location()),
		Config: nil,
	}
}

// =============================================================================
// 方案7: 使用 sync.Pool 复用 Time 对象
// =============================================================================

func BenchmarkEndOfMonth_Global_V7(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_ = EndOfMonth_Global_V7()
	}
}

var eomTimePool = sync.Pool{
	New: func() any {
		return &Time{}
	},
}

func EndOfMonth_Global_V7() *Time {
	now := time.Now()
	year, month, _ := now.Date()
	t := eomTimePool.Get().(*Time)
	t.Time = time.Date(year, month+1, 0, 23, 59, 59, 999999999, now.Location())
	t.Config = nil
	return t
}

// =============================================================================
// 方案8: 使用闭包捕获 time.Now() 结果
// =============================================================================

func BenchmarkEndOfMonth_Global_V8(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_ = EndOfMonth_Global_V8()
	}
}

func EndOfMonth_Global_V8() *Time {
	return func() *Time {
		now := time.Now()
		year, month, _ := now.Date()
		return &Time{
			Time:   time.Date(year, month+1, 0, 23, 59, 59, 999999999, now.Location()),
			Config: nil,
		}
	}()
}

// =============================================================================
// 方案9: 分离日期计算和对象构造
// =============================================================================

func BenchmarkEndOfMonth_Global_V9(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_ = EndOfMonth_Global_V9()
	}
}

func EndOfMonth_Global_V9() *Time {
	now := time.Now()
	year, month, _ := now.Date()
	eomTime := time.Date(year, month+1, 0, 23, 59, 59, 999999999, now.Location())
	return &Time{Time: eomTime, Config: nil}
}

// =============================================================================
// 方案10: 使用全局默认 Config（避免每次创建）
// =============================================================================

func BenchmarkEndOfMonth_Global_V10(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_ = EndOfMonth_Global_V10()
	}
}

var eomDefaultConfig = &Config{
	WeekStartDay: time.Monday,
	TimeLocation: time.Local,
	TimeFormats:  []string{},
}

func EndOfMonth_Global_V10() *Time {
	now := time.Now()
	year, month, _ := now.Date()
	return &Time{
		Time:   time.Date(year, month+1, 0, 23, 59, 59, 999999999, now.Location()),
		Config: eomDefaultConfig,
	}
}

// =============================================================================
// 方案11: 使用 time.Now().In() 确保时区一致性
// =============================================================================

func BenchmarkEndOfMonth_Global_V11(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_ = EndOfMonth_Global_V11()
	}
}

func EndOfMonth_Global_V11() *Time {
	now := time.Now().In(time.Local)
	year, month, _ := now.Date()
	return &Time{
		Time:   time.Date(year, month+1, 0, 23, 59, 59, 999999999, now.Location()),
		Config: nil,
	}
}

// =============================================================================
// 方案12: 使用指针返回优化（避免逃逸到堆）
// =============================================================================

func BenchmarkEndOfMonth_Global_V12(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_ = EndOfMonth_Global_V12()
	}
}

//go:noinline
func EndOfMonth_Global_V12() *Time {
	now := time.Now()
	year, month, _ := now.Date()
	return &Time{
		Time:   time.Date(year, month+1, 0, 23, 59, 59, 999999999, now.Location()),
		Config: nil,
	}
}

// =============================================================================
// 综合对比测试 - 所有方案在同一测试中对比
// =============================================================================

func BenchmarkEndOfMonth_Global_Comparison(b *testing.B) {
	b.Run("V1_Current", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			_ = EndOfMonth_Global_V1()
		}
	})

	b.Run("V2_FullInline", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			_ = EndOfMonth_Global_V2()
		}
	})

	b.Run("V3_MinimalConfig", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			_ = EndOfMonth_Global_V3()
		}
	})

	b.Run("V4_NilConfig", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			_ = EndOfMonth_Global_V4()
		}
	})

	b.Run("V5_ConstantNanos", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			_ = EndOfMonth_Global_V5()
		}
	})

	b.Run("V6_DirectConstruct", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			_ = EndOfMonth_Global_V6()
		}
	})

	b.Run("V7_SyncPool", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			_ = EndOfMonth_Global_V7()
		}
	})

	b.Run("V8_Closure", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			_ = EndOfMonth_Global_V8()
		}
	})

	b.Run("V9_SeparatedCalc", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			_ = EndOfMonth_Global_V9()
		}
	})

	b.Run("V10_GlobalConfig", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			_ = EndOfMonth_Global_V10()
		}
	})

	b.Run("V11_ExplicitLocal", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			_ = EndOfMonth_Global_V11()
		}
	})

	b.Run("V12_NoInline", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			_ = EndOfMonth_Global_V12()
		}
	})
}

// 原始实现
func Benchmark_EndOfQuarter_Global_Original(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = With(time.Now()).EndOfQuarter()
	}
}

// Variant1: 内联逻辑，避免 With() 调用
func Benchmark_EndOfQuarter_Global_Variant1(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		t := time.Now()
		year, month, _ := t.Date()
		quarter := (month-1)/3 + 1
		endQuarterMonth := quarter * 3
		_ = &Time{
			Time:   time.Date(year, time.Month(endQuarterMonth+1), 0, 23, 59, 59, 999999999, t.Location()),
			Config: &Config{WeekStartDay: time.Monday, TimeLocation: time.Local, TimeFormats: []string{}, Monotonic: time.Now()},
		}
	}
}

// Variant2: 预先创建 Config
func Benchmark_EndOfQuarter_Global_Variant2(b *testing.B) {
	config := &Config{WeekStartDay: time.Monday, TimeLocation: time.Local, TimeFormats: []string{}, Monotonic: time.Now()}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		t := time.Now()
		year, month, _ := t.Date()
		quarter := (month-1)/3 + 1
		endQuarterMonth := quarter * 3
		_ = &Time{
			Time:   time.Date(year, time.Month(endQuarterMonth+1), 0, 23, 59, 59, 999999999, t.Location()),
			Config: config,
		}
	}
}

// Variant3: 内联季度计算，减少中间变量
func Benchmark_EndOfQuarter_Global_Variant3(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		t := time.Now()
		year, month, _ := t.Date()
		_ = &Time{
			Time:   time.Date(year, time.Month(((month-1)/3+1)*3)+1, 0, 23, 59, 59, 999999999, t.Location()),
			Config: &Config{WeekStartDay: time.Monday, TimeLocation: time.Local, TimeFormats: []string{}, Monotonic: time.Now()},
		}
	}
}

// Variant4: 复用 time.Now() 结果
func Benchmark_EndOfQuarter_Global_Variant4(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		now := time.Now()
		year, month, _ := now.Date()
		quarter := (month-1)/3 + 1
		endQuarterMonth := quarter * 3
		_ = &Time{
			Time:   time.Date(year, time.Month(endQuarterMonth+1), 0, 23, 59, 59, 999999999, now.Location()),
			Config: &Config{WeekStartDay: time.Monday, TimeLocation: now.Location(), TimeFormats: []string{}, Monotonic: now},
		}
	}
}

// Variant5: 直接使用 Year() 和 Month() 方法
func Benchmark_EndOfQuarter_Global_Variant5(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		now := time.Now()
		year := now.Year()
		month := int(now.Month())
		quarter := (month-1)/3 + 1
		endQuarterMonth := quarter * 3
		_ = &Time{
			Time:   time.Date(year, time.Month(endQuarterMonth+1), 0, 23, 59, 59, 999999999, now.Location()),
			Config: &Config{WeekStartDay: time.Monday, TimeLocation: now.Location(), TimeFormats: []string{}, Monotonic: now},
		}
	}
}

// Variant6: 简化 Config，只设置必要字段
func Benchmark_EndOfQuarter_Global_Variant6(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		now := time.Now()
		year, month, _ := now.Date()
		quarter := (month-1)/3 + 1
		endQuarterMonth := quarter * 3
		_ = &Time{
			Time:   time.Date(year, time.Month(endQuarterMonth+1), 0, 23, 59, 59, 999999999, now.Location()),
			Config: &Config{WeekStartDay: time.Monday, TimeLocation: now.Location()},
		}
	}
}

// Variant7: 使用 month 直接计算，避免类型转换
func Benchmark_EndOfQuarter_Global_Variant7(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		now := time.Now()
		year := now.Year()
		month := now.Month()
		quarter := (month-1)/3 + 1
		endQuarterMonth := quarter * 3
		_ = &Time{
			Time:   time.Date(year, endQuarterMonth+1, 0, 23, 59, 59, 999999999, now.Location()),
			Config: &Config{WeekStartDay: time.Monday, TimeLocation: now.Location()},
		}
	}
}

// Variant8: 预先计算常用常量
func Benchmark_EndOfQuarter_Global_Variant8(b *testing.B) {
	const (
		hour      = 23
		min       = 59
		sec       = 59
		nsec      = 999999999
		weekStart = time.Monday
	)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		now := time.Now()
		year := now.Year()
		month := now.Month()
		quarter := (month-1)/3 + 1
		endQuarterMonth := quarter * 3
		loc := now.Location()
		_ = &Time{
			Time:   time.Date(year, endQuarterMonth+1, 0, hour, min, sec, nsec, loc),
			Config: &Config{WeekStartDay: weekStart, TimeLocation: loc},
		}
	}
}

// Variant9: 内联所有计算，最小化变量
func Benchmark_EndOfQuarter_Global_Variant9(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		now := time.Now()
		_ = &Time{
			Time:   time.Date(now.Year(), ((now.Month()-1)/3+1)*3+1, 0, 23, 59, 59, 999999999, now.Location()),
			Config: &Config{WeekStartDay: time.Monday, TimeLocation: now.Location()},
		}
	}
}

// Variant10: 使用 sync.Pool 复用 Config（如果适用）
func Benchmark_EndOfQuarter_Global_Variant10(b *testing.B) {
	configPool := &Config{WeekStartDay: time.Monday, TimeLocation: time.Local}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		now := time.Now()
		year, month, _ := now.Date()
		quarter := (month-1)/3 + 1
		endQuarterMonth := quarter * 3
		loc := now.Location()
		_ = &Time{
			Time:   time.Date(year, time.Month(endQuarterMonth+1), 0, 23, 59, 59, 999999999, loc),
			Config: &Config{WeekStartDay: configPool.WeekStartDay, TimeLocation: loc},
		}
	}
}

// Variant11: 完全内联，零中间变量
func Benchmark_EndOfQuarter_Global_Variant11(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		t := time.Now()
		_ = &Time{
			Time:   time.Date(t.Year(), ((t.Month()-1)/3+1)*3+1, 0, 23, 59, 59, 999999999, t.Location()),
			Config: &Config{WeekStartDay: time.Monday, TimeLocation: t.Location()},
		}
	}
}

// Variant12: 分离 time.Now() 调用，优化 Config 创建
func Benchmark_EndOfQuarter_Global_Variant12(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		now := time.Now()
		loc := now.Location()
		year := now.Year()
		month := now.Month()
		quarter := (month - 1) / 3
		endMonth := (quarter+1)*3 + 1
		_ = &Time{
			Time:   time.Date(year, time.Month(endMonth), 0, 23, 59, 59, 999999999, loc),
			Config: &Config{WeekStartDay: time.Monday, TimeLocation: loc},
		}
	}
}

// Variant13: 使用结构体字面量一次性创建
func Benchmark_EndOfQuarter_Global_Variant13(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		now := time.Now()
		t := now
		_ = &Time{
			Time:   time.Date(t.Year(), time.Month(((int(t.Month())-1)/3+1)*3)+1, 0, 23, 59, 59, 999999999, t.Location()),
			Config: &Config{WeekStartDay: time.Monday, TimeLocation: t.Location()},
		}
	}
}

// Variant14: 提取 quarter 计算为独立步骤
func Benchmark_EndOfQuarter_Global_Variant14(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		now := time.Now()
		month := now.Month()
		quarterEndMonth := ((month-1)/3+1)*3 + 1
		_ = &Time{
			Time:   time.Date(now.Year(), quarterEndMonth, 0, 23, 59, 59, 999999999, now.Location()),
			Config: &Config{WeekStartDay: time.Monday, TimeLocation: now.Location()},
		}
	}
}

// Variant15: 最简化版本，完全内联
func Benchmark_EndOfQuarter_Global_Variant15(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		t := time.Now()
		l := t.Location()
		m := t.Month()
		_ = &Time{
			Time:   time.Date(t.Year(), ((m-1)/3+1)*3+1, 0, 23, 59, 59, 999999999, l),
			Config: &Config{WeekStartDay: time.Monday, TimeLocation: l},
		}
	}
}

func Benchmark_EndOfQuarter_Global_Old(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		// 模拟原始实现
		_ = With(time.Now()).EndOfQuarter()
	}
}

func Benchmark_EndOfQuarter_Global_New(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		// 新实现已直接在 now.go 中
		_ = EndOfQuarter()
	}
}

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
			WeekStartDay: time.Monday,
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
				WeekStartDay: time.Monday,
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
	const endOfDayNanos = int(time.Second - time.Nanosecond)

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
			Time: time.Date(year, month, day+daysToAddTable[weekday], 23, 59, 59, endOfDayNanos, loc),
			Config: &Config{
				WeekStartDay: time.Monday,
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

func Benchmark_EndOfWeekGlobal_Optimized(b *testing.B) {
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_ = EndOfWeek()
	}
}

// =============================================================================
// 方案1: 当前实现（Baseline） - With(time.Now()).EndOfYear()
// =============================================================================

func BenchmarkEndOfYear_Global_V1_Current(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_ = EndOfYear_Global_V1()
	}
}

func EndOfYear_Global_V1() *Time {
	return With(time.Now()).EndOfYear()
}

// =============================================================================
// 方案2: 内联逻辑 - 复用 With 和 EndOfYear 的完整逻辑
// =============================================================================

func BenchmarkEndOfYear_Global_V2(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_ = EndOfYear_Global_V2()
	}
}

func EndOfYear_Global_V2() *Time {
	now := time.Now()
	year := now.Year()
	eoy := time.Date(year+1, time.January, 0, 23, 59, 59, 999999999, now.Location())
	return &Time{
		Time: eoy,
		Config: &Config{
			WeekStartDay: time.Monday,
			TimeLocation: time.Local,
			TimeFormats:  []string{},
			Monotonic:    now,
		},
	}
}

// =============================================================================
// 方案3: 简化 Config - 只设置必要字段
// =============================================================================

func BenchmarkEndOfYear_Global_V3(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_ = EndOfYear_Global_V3()
	}
}

func EndOfYear_Global_V3() *Time {
	now := time.Now()
	year := now.Year()
	return &Time{
		Time: time.Date(year+1, time.January, 0, 23, 59, 59, 999999999, now.Location()),
		Config: &Config{
			TimeLocation: now.Location(),
		},
	}
}

// =============================================================================
// 方案4: 零 Config - 使用 nil Config
// =============================================================================

func BenchmarkEndOfYear_Global_V4(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_ = EndOfYear_Global_V4()
	}
}

func EndOfYear_Global_V4() *Time {
	now := time.Now()
	year := now.Year()
	return &Time{
		Time:   time.Date(year+1, time.January, 0, 23, 59, 59, 999999999, now.Location()),
		Config: nil,
	}
}

// =============================================================================
// 方案5: 直接构造 Time - 避免中间变量
// =============================================================================

func BenchmarkEndOfYear_Global_V5(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_ = EndOfYear_Global_V5()
	}
}

func EndOfYear_Global_V5() *Time {
	now := time.Now()
	return &Time{
		Time:   time.Date(now.Year()+1, time.January, 0, 23, 59, 59, 999999999, now.Location()),
		Config: nil,
	}
}

// =============================================================================
// 方案6: 使用 UTC 时间 - 避免 Location 查询
// =============================================================================

func BenchmarkEndOfYear_Global_V6(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_ = EndOfYear_Global_V6()
	}
}

func EndOfYear_Global_V6() *Time {
	now := time.Now().UTC()
	year := now.Year()
	return &Time{
		Time:   time.Date(year+1, time.January, 0, 23, 59, 59, 999999999, time.UTC),
		Config: nil,
	}
}

// =============================================================================
// 方案7: 预定义常量 Config - 复用全局 Config
// =============================================================================

var globalEOYConfig = &Config{
	WeekStartDay: time.Monday,
	TimeLocation: time.Local,
	TimeFormats:  []string{},
}

func BenchmarkEndOfYear_Global_V7(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_ = EndOfYear_Global_V7()
	}
}

func EndOfYear_Global_V7() *Time {
	now := time.Now()
	year := now.Year()
	return &Time{
		Time:   time.Date(year+1, time.January, 0, 23, 59, 59, 999999999, now.Location()),
		Config: globalEOYConfig,
	}
}

// =============================================================================
// 方案8: 闭包优化 - 使用闭包减少重复计算
// =============================================================================

func BenchmarkEndOfYear_Global_V8(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_ = EndOfYear_Global_V8()
	}
}

func EndOfYear_Global_V8() *Time {
	now := time.Now()
	var year int
	var loc *time.Location

	func() {
		year = now.Year()
		loc = now.Location()
	}()

	return &Time{
		Time:   time.Date(year+1, time.January, 0, 23, 59, 59, 999999999, loc),
		Config: nil,
	}
}

// =============================================================================
// 方案9: 使用 sync.Pool 复用 Time 对象
// =============================================================================

var eoyTimePool = sync.Pool{
	New: func() interface{} {
		return &Time{}
	},
}

func BenchmarkEndOfYear_Global_V9(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_ = EndOfYear_Global_V9()
	}
}

func EndOfYear_Global_V9() *Time {
	now := time.Now()
	year := now.Year()
	t := eoyTimePool.Get().(*Time)
	t.Time = time.Date(year+1, time.January, 0, 23, 59, 59, 999999999, now.Location())
	t.Config = nil
	return t
}

// =============================================================================
// 方案10: 内联所有逻辑 - 最极致优化
// =============================================================================

func BenchmarkEndOfYear_Global_V10(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_ = EndOfYear_Global_V10()
	}
}

func EndOfYear_Global_V10() *Time {
	now := time.Now()
	return &Time{
		Time: time.Date(
			now.Year()+1,
			time.January,
			0,
			23,
			59,
			59,
			999999999,
			now.Location(),
		),
		Config: nil,
	}
}

// =============================================================================
// 方案11: 使用结构体字面量直接返回
// =============================================================================

func BenchmarkEndOfYear_Global_V11(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_ = EndOfYear_Global_V11()
	}
}

func EndOfYear_Global_V11() *Time {
	now := time.Now()
	year := now.Year()
	loc := now.Location()
	endTime := time.Date(year+1, time.January, 0, 23, 59, 59, 999999999, loc)
	return &Time{Time: endTime, Config: nil}
}

// =============================================================================
// 方案12: 零分配变体 - 使用空结构体 Config
// =============================================================================

var emptyConfigEOY = &Config{}

func BenchmarkEndOfYear_Global_V12(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_ = EndOfYear_Global_V12()
	}
}

func EndOfYear_Global_V12() *Time {
	now := time.Now()
	year := now.Year()
	return &Time{
		Time:   time.Date(year+1, time.January, 0, 23, 59, 59, 999999999, now.Location()),
		Config: emptyConfigEOY,
	}
}

// =============================================================================
// 方案13: 使用 time.Date 优化 - 减少参数解析
// =============================================================================

func BenchmarkEndOfYear_Global_V13(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_ = EndOfYear_Global_V13()
	}
}

func EndOfYear_Global_V13() *Time {
	now := time.Now()
	return &Time{
		Time:   now.AddDate(1, -int(now.Month()), 0).Add(time.Hour*24 - time.Nanosecond),
		Config: nil,
	}
}

// =============================================================================
// 方案14: 双重 nil Config - 测试 nil vs empty Config 性能差异
// =============================================================================

func BenchmarkEndOfYear_Global_V14(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_ = EndOfYear_Global_V14()
	}
}

func EndOfYear_Global_V14() *Time {
	now := time.Now()
	t := &Time{}
	t.Time = time.Date(now.Year()+1, time.January, 0, 23, 59, 59, 999999999, now.Location())
	return t
}

// =============================================================================
// 方案15: 使用 AddDate 计算明年1月1日再减1纳秒
// =============================================================================

func BenchmarkEndOfYear_Global_V15(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_ = EndOfYear_Global_V15()
	}
}

func EndOfYear_Global_V15() *Time {
	now := time.Now()
	// 明年1月1日
	nextYearStart := time.Date(now.Year()+1, time.January, 1, 0, 0, 0, 0, now.Location())
	// 减1纳秒 = 今年最后一刻
	return &Time{
		Time:   nextYearStart.Add(-time.Nanosecond),
		Config: nil,
	}
}

func BenchmarkEOY_Current(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_ = With(time.Now()).EndOfYear()
	}
}

func BenchmarkEOY_Optimized(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		now := time.Now()
		_ = &Time{
			Time:   time.Date(now.Year()+1, time.January, 0, 23, 59, 59, 999999999, now.Location()),
			Config: nil,
		}
	}
}

func BenchmarkEOY_OptimizedV2(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		now := time.Now()
		year := now.Year()
		_ = &Time{
			Time:   time.Date(year+1, time.January, 0, 23, 59, 59, 999999999, now.Location()),
			Config: nil,
		}
	}
}

func BenchmarkEOY_OptimizedV3(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		now := time.Now()
		nextYearStart := time.Date(now.Year()+1, time.January, 1, 0, 0, 0, 0, now.Location())
		_ = &Time{
			Time:   nextYearStart.Add(-time.Nanosecond),
			Config: nil,
		}
	}
}

// Baseline: 当前实现
func BenchmarkNowCalendar_Current(b *testing.B) {
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_ = NowCalendar()
	}
}

// 方案1: 内联 time.Now()
func BenchmarkNowCalendar_InlineTime(b *testing.B) {
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		t := time.Now()
		_ = NewCalendar(t)
	}
}

// 方案2: 延迟计算（仅创建基础结构）
func BenchmarkNowCalendar_Lazy(b *testing.B) {
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_ = newCalendarLazy(time.Now())
	}
}

// 方案3: 缓存 Zodiac 计算（相同年份复用）
func BenchmarkNowCalendar_CachedZodiac(b *testing.B) {
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_ = newCalendarCachedZodiac(time.Now())
	}
}

// 方案4: 简化 Season 计算（移除节气查询）
func BenchmarkNowCalendar_SimpleSeason(b *testing.B) {
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_ = newCalendarSimpleSeason(time.Now())
	}
}

// 方案5: 完全简化（仅基础信息）
func BenchmarkNowCalendar_Minimal(b *testing.B) {
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_ = newCalendarMinimal(time.Now())
	}
}

// 方案6: 预计算常量（优化数组查找）
func BenchmarkNowCalendar_Prealloc(b *testing.B) {
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_ = newCalendarPrealloc(time.Now())
	}
}

// 对比基准：直接创建 Calendar（不含 time.Now()）
func BenchmarkNewCalendar_Only(b *testing.B) {
	t := time.Now()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_ = NewCalendar(t)
	}
}

// 对比基准：仅 time.Now()
func BenchmarkTimeNow_Only(b *testing.B) {
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_ = time.Now()
	}
}

// 对比基准：仅 Lunar 计算
func BenchmarkWithLunar_Only(b *testing.B) {
	t := time.Now()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_ = WithLunar(t)
	}
}

// 对比基准：仅节气查询
func BenchmarkNextSolarterm_Only(b *testing.B) {
	t := time.Now()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_ = NextSolarterm(t)
	}
}

// ========== 优化变体方法 ==========

// Quarter_Original 当前实现
func (p *Time) Quarter_Original() uint {
	return (uint(p.Month())-1)/3 + 1
}

// Quarter_V2: 消除中间减法
func (p *Time) Quarter_V2_NoMinus() uint {
	month := int(p.Month())
	return uint(month/3 + 1)
}

// Quarter_V3: 位运算优化
func (p *Time) Quarter_V3_BitOps() uint {
	month := int(p.Month())
	return uint((month >> 1) + (month >> 3))
}

// Quarter_V5: 全局查找表
var globalQuarterTable = [12]uint{1, 1, 1, 2, 2, 2, 3, 3, 3, 4, 4, 4}

func (p *Time) Quarter_V5_GlobalLookup() uint {
	return globalQuarterTable[int(p.Month())-1]
}

// Quarter_V6: Switch 语句
func (p *Time) Quarter_V6_Switch() uint {
	switch p.Month() {
	case time.January, time.February, time.March:
		return 1
	case time.April, time.May, time.June:
		return 2
	case time.July, time.August, time.September:
		return 3
	default:
		return 4
	}
}

// Quarter_V7: If-Else 链
func (p *Time) Quarter_V7_IfElse() uint {
	m := p.Month()
	if m <= time.March {
		return 1
	}
	if m <= time.June {
		return 2
	}
	if m <= time.September {
		return 3
	}
	return 4
}

// Quarter_V8: 直接访问 time.Time
func (p *Time) Quarter_V8_DirectTime() uint {
	return uint(p.Time.Month()/3) + 1
}

// ========== 基准测试 ==========

func genQuarterTestTimes(n int) []*Time {
	times := make([]*Time, n)
	for i := 0; i < n; i++ {
		m := time.Month((i % 12) + 1)
		times[i] = &Time{Time: time.Date(2024, m, 1, 0, 0, 0, 0, time.UTC)}
	}
	return times
}

func BenchmarkQuarter_Original(b *testing.B) {
	times := genQuarterTestTimes(12)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for _, t := range times {
			_ = t.Quarter_Original()
		}
	}
}

func BenchmarkQuarter_V2_NoMinus(b *testing.B) {
	times := genQuarterTestTimes(12)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for _, t := range times {
			_ = t.Quarter_V2_NoMinus()
		}
	}
}

func BenchmarkQuarter_V3_BitOps(b *testing.B) {
	times := genQuarterTestTimes(12)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for _, t := range times {
			_ = t.Quarter_V3_BitOps()
		}
	}
}

func BenchmarkQuarter_V5_GlobalLookup(b *testing.B) {
	times := genQuarterTestTimes(12)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for _, t := range times {
			_ = t.Quarter_V5_GlobalLookup()
		}
	}
}

func BenchmarkQuarter_V6_Switch(b *testing.B) {
	times := genQuarterTestTimes(12)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for _, t := range times {
			_ = t.Quarter_V6_Switch()
		}
	}
}

func BenchmarkQuarter_V7_IfElse(b *testing.B) {
	times := genQuarterTestTimes(12)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for _, t := range times {
			_ = t.Quarter_V7_IfElse()
		}
	}
}

func BenchmarkQuarter_V8_DirectTime(b *testing.B) {
	times := genQuarterTestTimes(12)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for _, t := range times {
			_ = t.Quarter_V8_DirectTime()
		}
	}
}

// ========== 内存分配测试 ==========

func BenchmarkQuarter_Original_Alloc(b *testing.B) {
	times := genQuarterTestTimes(12)
	b.ResetTimer()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		for _, t := range times {
			_ = t.Quarter_Original()
		}
	}
}

func BenchmarkQuarter_V5_GlobalLookup_Alloc(b *testing.B) {
	times := genQuarterTestTimes(12)
	b.ResetTimer()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		for _, t := range times {
			_ = t.Quarter_V5_GlobalLookup()
		}
	}
}

func BenchmarkQuarter_V6_Switch_Alloc(b *testing.B) {
	times := genQuarterTestTimes(12)
	b.ResetTimer()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		for _, t := range times {
			_ = t.Quarter_V6_Switch()
		}
	}
}

func BenchmarkSimple(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = i
	}
}

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
