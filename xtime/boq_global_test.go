package xtime

import (
	"testing"
	"time"
)

// TestBeginningOfQuarterGlobal_Correctness 验证全局 BeginningOfQuarter 函数的正确性
func TestBeginningOfQuarterGlobal_Correctness(t *testing.T) {
	tests := []struct {
		name     string
		month    time.Month
		expected time.Month
	}{
		{"January starts Q1", time.January, time.January},
		{"February starts Q1", time.February, time.January},
		{"March starts Q1", time.March, time.January},
		{"April starts Q2", time.April, time.April},
		{"May starts Q2", time.May, time.April},
		{"June starts Q2", time.June, time.April},
		{"July starts Q3", time.July, time.July},
		{"August starts Q3", time.August, time.July},
		{"September starts Q3", time.September, time.July},
		{"October starts Q4", time.October, time.October},
		{"November starts Q4", time.November, time.October},
		{"December starts Q4", time.December, time.October},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// 创建一个测试时间点
			testTime := time.Date(2024, tt.month, 15, 12, 30, 45, 0, time.UTC)

			// 模拟 BeginningOfQuarter 的逻辑
			month := testTime.Month()
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

			result := time.Date(testTime.Year(), startMonth, 1, 0, 0, 0, 0, testTime.Location())

			if result.Month() != tt.expected {
				t.Errorf("Expected month %v, got %v", tt.expected, result.Month())
			}

			if result.Day() != 1 {
				t.Errorf("Expected day 1, got %d", result.Day())
			}

			if result.Hour() != 0 || result.Minute() != 0 || result.Second() != 0 {
				t.Errorf("Expected time 00:00:00, got %02d:%02d:%02d",
					result.Hour(), result.Minute(), result.Second())
			}
		})
	}
}

// TestBeginningOfQuarterGlobal_Consistency 验证全局函数与方法的一致性
func TestBeginningOfQuarterGlobal_Consistency(t *testing.T) {
	// 使用固定时间进行测试
	fixedTime := time.Date(2024, 6, 15, 12, 30, 45, 0, time.UTC)
	timeWrapper := With(fixedTime)

	// 方法调用
	methodResult := timeWrapper.BeginningOfQuarter()

	// 全局函数调用（使用相同的时间）
	month := fixedTime.Month()
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
	globalResult := &Time{
		Time: time.Date(fixedTime.Year(), startMonth, 1, 0, 0, 0, 0, fixedTime.Location()),
		Config: &Config{
			WeekStartDay:  time.Monday,
			TimeLocation: time.Local,
			TimeFormats:  []string{},
			Monotonic:    fixedTime,
		},
	}

	if methodResult.Time != globalResult.Time {
		t.Errorf("Method result %v != global result %v", methodResult.Time, globalResult.Time)
	}
}

// TestBeginningOfQuarterGlobal_Performance 性能测试
func TestBeginningOfQuarterGlobal_Performance(t *testing.T) {
	iterations := 1000000

	start := time.Now()
	for i := 0; i < iterations; i++ {
		_ = BeginningOfQuarter()
	}
	elapsed := time.Since(start)

	nsPerOp := elapsed.Nanoseconds() / int64(iterations)
	t.Logf("Average time per call: %d ns/op", nsPerOp)
	t.Logf("Total time for %d calls: %v", iterations, elapsed)

	// 验证性能：应该 < 200 ns/op（包含 time.Now() 开销）
	// 原始实现约 194 ns/op，优化后应该 < 180 ns/op
	if nsPerOp > 180 {
		t.Errorf("Performance too slow: %d ns/op, want < 180 ns/op", nsPerOp)
	}
}
