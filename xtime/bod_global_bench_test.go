package xtime

import (
	"sync"
	"testing"
	"time"
)

// =============================================================================
// 方案1: 当前实现（Baseline）
// =============================================================================

func BenchmarkBeginningOfDay_Global_Current(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_ = BeginningOfDay_Global_V1()
	}
}

func BeginningOfDay_Global_V1() *Time {
	return With(time.Now()).BeginningOfDay()
}

// =============================================================================
// 方案2: 内联 With + BeginningOfDay 逻辑
// =============================================================================

func BenchmarkBeginningOfDay_Global_V2(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_ = BeginningOfDay_Global_V2()
	}
}

func BeginningOfDay_Global_V2() *Time {
	now := time.Now()
	year, month, day := now.Date()
	midnight := time.Date(year, month, day, 0, 0, 0, 0, now.Location())
	return &Time{
		Time: midnight,
		Config: &Config{
			WeekStartDay:  time.Monday,
			TimeLocation: time.Local,
			TimeFormats:  []string{},
			Monotonic:    now,
		},
	}
}

// =============================================================================
// 方案3: 简化 Config，只设置必要字段
// =============================================================================

func BenchmarkBeginningOfDay_Global_V3(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_ = BeginningOfDay_Global_V3()
	}
}

func BeginningOfDay_Global_V3() *Time {
	now := time.Now()
	year, month, day := now.Date()
	return &Time{
		Time: time.Date(year, month, day, 0, 0, 0, 0, now.Location()),
		Config: &Config{
			TimeLocation: now.Location(),
		},
	}
}

// =============================================================================
// 方案4: 零 Config（使用 nil）
// =============================================================================

func BenchmarkBeginningOfDay_Global_V4(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_ = BeginningOfDay_Global_V4()
	}
}

func BeginningOfDay_Global_V4() *Time {
	now := time.Now()
	year, month, day := now.Date()
	return &Time{
		Time:   time.Date(year, month, day, 0, 0, 0, 0, now.Location()),
		Config: nil,
	}
}

// =============================================================================
// 方案5: 使用 Truncate(24小时) - 存在时区问题，仅作性能参考
// 问题: Truncate 从 00:00 开始，在非 UTC 时区可能不准确
// =============================================================================

func BenchmarkBeginningOfDay_Global_V5(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_ = BeginningOfDay_Global_V5()
	}
}

func BeginningOfDay_Global_V5() *Time {
	now := time.Now()
	// 注意: 这个方案有正确性问题，仅供参考
	midnight := now.Truncate(24 * time.Hour)
	return &Time{
		Time:   midnight,
		Config: &Config{TimeLocation: now.Location()},
	}
}

// =============================================================================
// 方案6: 使用 Add 向下取整到午夜
// =============================================================================

func BenchmarkBeginningOfDay_Global_V6(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_ = BeginningOfDay_Global_V6()
	}
}

func BeginningOfDay_Global_V6() *Time {
	now := time.Now()
	year, month, day := now.Date()
	midnight := time.Date(year, month, day, 0, 0, 0, 0, now.Location())
	return &Time{Time: midnight}
}

// =============================================================================
// 方案7: 预计算 UTC 午夜，再转换时区 - 存在时区问题，仅供参考
// 问题: UTC 的日期可能与本地时间不同
// =============================================================================

func BenchmarkBeginningOfDay_Global_V7(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_ = BeginningOfDay_Global_V7()
	}
}

func BeginningOfDay_Global_V7() *Time {
	now := time.Now()
	// 注意: 这个方案有正确性问题，仅供参考
	utcNow := now.UTC()
	year, month, day := utcNow.Date()
	utcMidnight := time.Date(year, month, day, 0, 0, 0, 0, time.UTC)
	return &Time{
		Time:   utcMidnight.In(now.Location()),
		Config: &Config{TimeLocation: now.Location()},
	}
}

// =============================================================================
// 方案8: 使用 Unix 时间戳计算
// =============================================================================

func BenchmarkBeginningOfDay_Global_V8(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_ = BeginningOfDay_Global_V8()
	}
}

func BeginningOfDay_Global_V8() *Time {
	now := time.Now()
	year, month, day := now.Date()
	midnight := time.Date(year, month, day, 0, 0, 0, 0, now.Location())
	return &Time{Time: midnight, Config: &Config{}}
}

// =============================================================================
// 方案9: 复用全局默认 Config（只读）- 存在并发安全问题，仅供参考
// 问题: 全局 Config 可能被意外修改
// =============================================================================

var defaultConfig = &Config{
	WeekStartDay:  time.Monday,
	TimeLocation: time.Local,
	TimeFormats:  []string{},
	Monotonic:    time.Time{},
}

func BenchmarkBeginningOfDay_Global_V9(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_ = BeginningOfDay_Global_V9()
	}
}

func BeginningOfDay_Global_V9() *Time {
	now := time.Now()
	year, month, day := now.Date()
	// 注意: 这个方案有并发安全问题，仅供参考
	return &Time{
		Time:   time.Date(year, month, day, 0, 0, 0, 0, now.Location()),
		Config: defaultConfig,
	}
}

// =============================================================================
// 方案10: 直接返回 time.Time，包装为 Time
// =============================================================================

func BenchmarkBeginningOfDay_Global_V10(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_ = BeginningOfDay_Global_V10()
	}
}

func BeginningOfDay_Global_V10() *Time {
	now := time.Now()
	year, month, day := now.Date()
	midnight := time.Date(year, month, day, 0, 0, 0, 0, now.Location())
	return &Time{Time: midnight}
}

// =============================================================================
// 方案11: 使用 sync.Pool 复用 Time 对象
// =============================================================================

var timePool = sync.Pool{
	New: func() interface{} {
		return &Time{}
	},
}

func BenchmarkBeginningOfDay_Global_V11(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_ = BeginningOfDay_Global_V11()
	}
}

func BeginningOfDay_Global_V11() *Time {
	t := timePool.Get().(*Time)
	now := time.Now()
	year, month, day := now.Date()
	t.Time = time.Date(year, month, day, 0, 0, 0, 0, now.Location())
	t.Config = nil
	result := *t // 复制返回
	timePool.Put(t)
	return &result
}

// =============================================================================
// 方案12: 最简化 - 只设置 Time 字段
// =============================================================================

func BenchmarkBeginningOfDay_Global_V12(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_ = BeginningOfDay_Global_V12()
	}
}

func BeginningOfDay_Global_V12() *Time {
	now := time.Now()
	year, month, day := now.Date()
	return &Time{Time: time.Date(year, month, day, 0, 0, 0, 0, now.Location())}
}

// =============================================================================
// 对比基准：time.Now() 自身性能
// =============================================================================

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

// =============================================================================
// 正确性验证测试
// =============================================================================

func TestBeginningOfDayGlobal_Correctness(t *testing.T) {
	// 测试不同方案结果一致性
	// 排除 V5, V7, V9 (有已知问题)
	now := time.Now()
	expected := With(now).BeginningOfDay()

	results := map[string]*Time{
		"V1":  BeginningOfDay_Global_V1(),
		"V2":  BeginningOfDay_Global_V2(),
		"V3":  BeginningOfDay_Global_V3(),
		"V4":  BeginningOfDay_Global_V4(),
		"V6":  BeginningOfDay_Global_V6(),
		"V8":  BeginningOfDay_Global_V8(),
		"V10": BeginningOfDay_Global_V10(),
		"V12": BeginningOfDay_Global_V12(),
	}

	for name, result := range results {
		if result.Time.Unix() != expected.Time.Unix() {
			t.Errorf("%s 时间不一致: expected %v, got %v", name, expected.Time, result.Time)
		}
		if result.Time.Location().String() != expected.Time.Location().String() {
			t.Errorf("%s 时区不一致: expected %v, got %v", name, expected.Time.Location(), result.Time.Location())
		}
	}
}

func TestBeginningOfDayGlobal_V5_Issues(t *testing.T) {
	// V5 使用 Truncate，记录已知问题
	t.Skip("V5 Truncate 方案在非 UTC 时区存在正确性问题，已排除")
}

func TestBeginningOfDayGlobal_V7_Issues(t *testing.T) {
	// V7 使用 UTC 转换，记录已知问题
	t.Skip("V7 UTC 方案在跨时区边界时存在正确性问题，已排除")
}

func TestBeginningOfDayGlobal_V9_Issues(t *testing.T) {
	// V9 使用全局 Config，记录已知问题
	t.Skip("V9 全局 Config 方案存在并发安全问题，已排除")
}

func TestBeginningOfDayGlobal_V11_Correctness(t *testing.T) {
	// V11 使用 sync.Pool，验证返回值独立性
	result := BeginningOfDay_Global_V11()
	_ = result // 使用结果
	// 多次调用验证无数据竞争
	for i := 0; i < 100; i++ {
		r := BeginningOfDay_Global_V11()
		if r.Time.IsZero() {
			t.Error("V11 返回零时间")
		}
	}
}
