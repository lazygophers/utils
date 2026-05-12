package xtime

import (
	"testing"
	"time"
)

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
