package xtime

import (
	"sync"
	"testing"
	"time"
)

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
	WeekStartDay:  time.Monday,
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
