package xtime

import (
	"sync"
	"testing"
	"time"
)

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
			WeekStartDay:  time.Monday,
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
	WeekStartDay:  time.Monday,
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
