package xtime

import (
	"fmt"
	"testing"
	"time"
)

// TestEndOfMonthGlobal_Correctness 验证优化后的函数正确性
func TestEndOfMonthGlobal_Correctness(t *testing.T) {
	tests := []struct {
		name string
		date time.Time
		want string
	}{
		{
			name: "2024年1月15日",
			date: time.Date(2024, 1, 15, 10, 30, 0, 0, time.Local),
			want: "2024-01-31 23:59:59.999999999 +0800 CST",
		},
		{
			name: "2024年2月10日（闰年）",
			date: time.Date(2024, 2, 10, 14, 20, 0, 0, time.Local),
			want: "2024-02-29 23:59:59.999999999 +0800 CST",
		},
		{
			name: "2024年12月31日",
			date: time.Date(2024, 12, 31, 23, 59, 59, 0, time.Local),
			want: "2024-12-31 23:59:59.999999999 +0800 CST",
		},
		{
			name: "2023年2月10日（非闰年）",
			date: time.Date(2023, 2, 10, 14, 20, 0, 0, time.Local),
			want: "2023-02-28 23:59:59.999999999 +0800 CST",
		},
		{
			name: "2024年6月30日",
			date: time.Date(2024, 6, 30, 12, 0, 0, 0, time.Local),
			want: "2024-06-30 23:59:59.999999999 +0800 CST",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// 使用固定时间进行测试
			result := With(tt.date).EndOfMonth()
			got := result.Time.String()

			// 只比较时间部分，忽略时区差异
			wantTime, _ := time.Parse("2006-01-02 15:04:05.999999999 -0700 MST", tt.want)

			if !result.Time.Equal(wantTime) {
				t.Errorf("EndOfMonth() = %v, want %v", got, tt.want)
			} else {
				fmt.Printf("Test: %s PASS\n", tt.name)
			}
		})
	}
}

// TestEndOfMonthGlobal_Performance 验证性能优化效果
func TestEndOfMonthGlobal_Performance(t *testing.T) {
	if testing.Short() {
		t.Skip("跳过性能测试")
	}

	iterations := 1000000
	start := time.Now()

	for i := 0; i < iterations; i++ {
		_ = EndOfMonth()
	}

	elapsed := time.Since(start)
	avgTime := elapsed.Nanoseconds() / int64(iterations)

	t.Logf("Total time for %d calls: %v", iterations, elapsed)
	t.Logf("Average time per call: %d ns/op", avgTime)

	// 验证性能阈值：< 100 ns/op
	if avgTime > 100 {
		t.Errorf("Performance too slow: %d ns/op, want < 100 ns/op", avgTime)
	}
}

// TestEndOfMonthGlobal_ZeroAllocation 验证零内存分配
func TestEndOfMonthGlobal_ZeroAllocation(t *testing.T) {
	if testing.Short() {
		t.Skip("跳过内存分配测试")
	}

	// 使用 testing.AllocsPerRun 测试内存分配
	allocs := testing.AllocsPerRun(1000, func() {
		_ = EndOfMonth()
	})

	t.Logf("Allocs per call: %.2f", allocs)

	// 验证零分配
	if allocs > 0.5 { // 允许一些浮点误差
		t.Errorf("Expected zero allocations, got %.2f allocs/op", allocs)
	} else {
		fmt.Println("Zero allocation test PASSED")
	}
}

// TestEndOfMonthGlobal_MonthBoundaries 测试月份边界
func TestEndOfMonthGlobal_MonthBoundaries(t *testing.T) {
	testMonths := []struct {
		year  int
		month time.Month
		day   int
	}{
		{2024, 1, 31},   // 一月有31天
		{2024, 2, 29},   // 闰年二月有29天
		{2023, 2, 28},   // 非闰年二月有28天
		{2024, 4, 30},   // 四月有30天
		{2024, 6, 30},   // 六月有30天
		{2024, 9, 30},   // 九月有30天
		{2024, 11, 30},  // 十一月有30天
		{2024, 12, 31},  // 十二月有31天
	}

	for _, tm := range testMonths {
		t.Run(fmt.Sprintf("%d-%02d", tm.year, tm.month), func(t *testing.T) {
			testDate := time.Date(tm.year, tm.month, tm.day, 0, 0, 0, 0, time.Local)
			result := With(testDate).EndOfMonth()

			expectedDay := tm.day
			if result.Day() != expectedDay {
				t.Errorf("%d-%02d: EndOfMonth().Day() = %d, want %d",
					tm.year, tm.month, result.Day(), expectedDay)
			} else {
				t.Logf("%d-%02d: 正确返回第 %d 天", tm.year, tm.month, expectedDay)
			}
		})
	}
}

// TestEndOfMonthGlobal_YearTransition 测试年份过渡
func TestEndOfMonthGlobal_YearTransition(t *testing.T) {
	// 测试12月的月末不应该影响年份
	dec31 := time.Date(2024, 12, 31, 23, 59, 59, 0, time.Local)
	result := With(dec31).EndOfMonth()

	if result.Year() != 2024 {
		t.Errorf("12月31日年末: Year = %d, want 2024", result.Year())
	}

	if result.Month() != 12 {
		t.Errorf("12月31日年末: Month = %d, want 12", result.Month())
	}

	if result.Day() != 31 {
		t.Errorf("12月31日年末: Day = %d, want 31", result.Day())
	}

	t.Log("年份过渡测试通过")
}
