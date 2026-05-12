package xtime

import (
	"testing"
	"time"
)

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
				WeekStartDay:  time.Monday,
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
		_ = With(t.Truncate(8760 * time.Hour)).AddDate(-int(t.Month())+1, -int(t.Day())+1, 0)
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
