package xtime

import (
	"testing"
	"time"
)

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
