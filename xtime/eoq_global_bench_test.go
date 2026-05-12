package xtime

import (
	"testing"
	"time"
)

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
		hour       = 23
		min        = 59
		sec        = 59
		nsec       = 999999999
		weekStart  = time.Monday
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
