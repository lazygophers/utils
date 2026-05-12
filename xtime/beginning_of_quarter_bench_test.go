package xtime

import (
	"testing"
	"time"
)

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
