package xtime

import (
	"testing"
	"time"
)

// TestBeginningOfWeek_Correctness 验证 BeginningOfWeek 功能正确性
func TestBeginningOfWeek_Correctness(t *testing.T) {
	tests := []struct {
		name         string
		date         time.Time
		weekStartDay time.Weekday
		wantDay      int
		wantMonth    time.Month
		wantYear     int
	}{
		{
			name:         "2024-05-11 (周六) 周日起始",
			date:         time.Date(2024, 5, 11, 15, 30, 45, 0, time.Local),
			weekStartDay: time.Sunday,
			wantDay:      5, // 5月5日是周日
			wantMonth:    5,
			wantYear:     2024,
		},
		{
			name:         "2024-05-11 (周六) 周一起始",
			date:         time.Date(2024, 5, 11, 15, 30, 45, 0, time.Local),
			weekStartDay: time.Monday,
			wantDay:      6, // 5月6日是周一
			wantMonth:    5,
			wantYear:     2024,
		},
		{
			name:         "2024-05-06 (周一) 周一起始",
			date:         time.Date(2024, 5, 6, 10, 20, 30, 0, time.Local),
			weekStartDay: time.Monday,
			wantDay:      6, // 5月6日是周一
			wantMonth:    5,
			wantYear:     2024,
		},
		{
			name:         "2024-01-01 (周一) 周一起始",
			date:         time.Date(2024, 1, 1, 12, 0, 0, 0, time.Local),
			weekStartDay: time.Monday,
			wantDay:      1,
			wantMonth:    1,
			wantYear:     2024,
		},
		{
			name:         "2024-12-31 (周二) 周一起始",
			date:         time.Date(2024, 12, 31, 23, 59, 59, 0, time.Local),
			weekStartDay: time.Monday,
			wantDay:      30, // 12月30日是周一
			wantMonth:    12,
			wantYear:     2024,
		},
		{
			name:         "2024-01-07 (周日) 周日起始",
			date:         time.Date(2024, 1, 7, 0, 0, 0, 0, time.Local),
			weekStartDay: time.Sunday,
			wantDay:      7,
			wantMonth:    1,
			wantYear:     2024,
		},
		{
			name:         "2023-12-31 (周日) 周日起始",
			date:         time.Date(2023, 12, 31, 23, 59, 59, 0, time.Local),
			weekStartDay: time.Sunday,
			wantDay:      31,
			wantMonth:    12,
			wantYear:     2023,
		},
		{
			name:         "跨年测试 2024-01-01 (周一) 周日起始",
			date:         time.Date(2024, 1, 1, 0, 0, 0, 0, time.Local),
			weekStartDay: time.Sunday,
			wantDay:      31, // 2023年12月31日是周日
			wantMonth:    12,
			wantYear:     2023,
		},
		{
			name:         "周三 周三起始",
			date:         time.Date(2024, 5, 15, 12, 0, 0, 0, time.Local), // 5月15日是周三
			weekStartDay: time.Wednesday,
			wantDay:      15,
			wantMonth:    5,
			wantYear:     2024,
		},
		{
			name:         "周四 周三起始",
			date:         time.Date(2024, 5, 16, 12, 0, 0, 0, time.Local), // 5月16日是周四
			weekStartDay: time.Wednesday,
			wantDay:      15,
			wantMonth:    5,
			wantYear:     2024,
		},
		{
			name:         "周二 周三起始",
			date:         time.Date(2024, 5, 14, 12, 0, 0, 0, time.Local), // 5月14日是周二
			weekStartDay: time.Wednesday,
			wantDay:      8, // 5月8日是周三
			wantMonth:    5,
			wantYear:     2024,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			xt := With(tt.date)
			xt.WeekStartDay = tt.weekStartDay

			result := xt.BeginningOfWeek()

			if result.Year() != tt.wantYear {
				t.Errorf("Year() = %d, want %d", result.Year(), tt.wantYear)
			}
			if result.Month() != tt.wantMonth {
				t.Errorf("Month() = %d, want %d", result.Month(), tt.wantMonth)
			}
			if result.Day() != tt.wantDay {
				t.Errorf("Day() = %d, want %d", result.Day(), tt.wantDay)
			}

			// 验证时间是午夜
			h, m, s := result.Clock()
			if h != 0 || m != 0 || s != 0 {
				t.Errorf("Clock() = %d:%d:%d, want 0:0:0", h, m, s)
			}

			// 验证 Config 被保留
			if result.WeekStartDay != tt.weekStartDay {
				t.Errorf("WeekStartDay = %d, want %d", result.WeekStartDay, tt.weekStartDay)
			}
		})
	}
}

// TestBeginningOfWeek_ConfigNil 验证 nil Config 处理
func TestBeginningOfWeek_ConfigNil(t *testing.T) {
	date := time.Date(2024, 5, 11, 15, 30, 45, 0, time.Local)
	xt := &Time{Time: date, Config: nil}

	result := xt.BeginningOfWeek()

	if result.Config == nil {
		t.Error("Config should not be nil after BeginningOfWeek")
	}
}

// TestBeginningOfWeek_Timezone 验证时区正确性
func TestBeginningOfWeek_Timezone(t *testing.T) {
	locations := []struct {
		name string
		loc  *time.Location
	}{
		{"UTC", time.UTC},
		{"Local", time.Local},
		{"America/New_York", time.FixedZone("EST", -5*3600)},
		{"Asia/Shanghai", time.FixedZone("CST", 8*3600)},
	}

	date := time.Date(2024, 5, 11, 15, 30, 45, 0, time.UTC) // UTC 时间

	for _, loc := range locations {
		t.Run(loc.name, func(t *testing.T) {
			xt := With(date.In(loc.loc))
			xt.WeekStartDay = time.Sunday

			result := xt.BeginningOfWeek()

			// 验证时区被保留
			if result.Location().String() != loc.loc.String() {
				t.Errorf("Location = %s, want %s", result.Location(), loc.loc)
			}
		})
	}
}

// TestBeginningOfWeek_Monotonic 验证 Monotonic 时间被保留
func TestBeginningOfWeek_Monotonic(t *testing.T) {
	date := time.Date(2024, 5, 11, 15, 30, 45, 0, time.Local)
	monotonic := time.Now().Add(-time.Hour)

	xt := &Time{
		Time: date,
		Config: &Config{
			Monotonic: monotonic,
		},
	}

	result := xt.BeginningOfWeek()

	if result.Monotonic.IsZero() {
		t.Error("Monotonic should not be zero")
	}
}

// TestBeginningOfWeek_BeginningOfDayConsistency 验证与 BeginningOfDay 的一致性
func TestBeginningOfWeek_BeginningOfDayConsistency(t *testing.T) {
	date := time.Date(2024, 5, 11, 15, 30, 45, 0, time.Local)
	xt := With(date)
	xt.WeekStartDay = time.Sunday

	bow := xt.BeginningOfWeek()
	bod := bow.BeginningOfDay()

	// BeginningOfWeek 的结果应该已经是午夜，再次调用 BeginningOfDay 不应该改变
	if !bow.Time.Equal(bod.Time) {
		t.Errorf("BeginningOfWeek().BeginningOfDay() should equal BeginningOfWeek()")
	}
}

// TestBeginningOfWeek_Performance 性能对比测试
func TestBeginningOfWeek_Performance(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping performance test in short mode")
	}

	date := time.Date(2024, 5, 11, 15, 30, 45, 0, time.Local)
	xt := With(date)
	xt.WeekStartDay = time.Sunday

	// 预热
	for i := 0; i < 1000; i++ {
		xt.BeginningOfWeek()
	}

	// 基准测试
	iterations := 100000
	start := time.Now()
	for i := 0; i < iterations; i++ {
		xt.BeginningOfWeek()
	}
	elapsed := time.Since(start)

	avgNs := elapsed.Nanoseconds() / int64(iterations)
	t.Logf("Average time per call: %d ns/op", avgNs)

	// 验证性能目标：应该 < 200 ns/op（测试环境可能比基准测试慢）
	if avgNs > 200 {
		t.Errorf("Performance too slow: %d ns/op, want < 200 ns/op", avgNs)
	}
}
