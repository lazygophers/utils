package xtime

import (
	"sync"
	"testing"
	"time"
)

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
			WeekStartDay:  time.Monday,
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
	WeekStartDay:  time.Monday,
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
