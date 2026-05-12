package xtime

import (
	"testing"
	"time"
)

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
				WeekStartDay:  time.Monday,
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
				WeekStartDay:  time.Monday,
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
		WeekStartDay:  time.Monday,
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
				WeekStartDay:  time.Monday,
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
				WeekStartDay:  time.Monday,
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
				WeekStartDay:  time.Monday,
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
				WeekStartDay:  time.Monday,
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
				WeekStartDay:  time.Monday,
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
			Time:     time.Date(year, time.Month(quarterStartMonth+1), 1, 0, 0, 0, 0, loc),
			Config: &Config{
				WeekStartDay:  time.Monday,
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
				WeekStartDay:  time.Monday,
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
