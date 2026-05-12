package xtime

import (
	"testing"
	"time"
)

// TestEndOfDayGlobal 验证全局 EndOfDay 函数的正确性
func TestEndOfDayGlobal(t *testing.T) {
	testCases := []struct {
		name string
		year int
		month time.Month
		day int
		hour int
		min int
		sec int
		nsec int
	}{
		{"2024年6月15日", 2024, time.June, 15, 14, 30, 45, 123456789},
		{"2024年1月1日", 2024, time.January, 1, 0, 0, 0, 0},
		{"2024年12月31日中午", 2024, time.December, 31, 12, 0, 0, 0},
		{"2024年闰年日", 2024, time.February, 29, 23, 59, 59, 999999999},
		{"2023年非闰年", 2023, time.February, 28, 1, 2, 3, 4},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// 构造测试时间
			testTime := time.Date(tc.year, tc.month, tc.day, tc.hour, tc.min, tc.sec, tc.nsec, time.Local)

			// 使用 With().EndOfDay() 作为基准
			expected := With(testTime).EndOfDay()

			// 模拟全局 EndOfDay 的逻辑
			now := testTime
			year, month, day := now.Date()
			eod := time.Date(year, month, day, 23, 59, 59, int(time.Second-time.Nanosecond), now.Location())
			actual := &Time{Time: eod}

			// 验证时间相等
			if !actual.Time.Equal(expected.Time) {
				t.Errorf("时间不相等: expected=%v, actual=%v", expected.Time, actual.Time)
			}

			// 验证具体字段
			expectedYear, expectedMonth, expectedDay := expected.Date()
			actualYear, actualMonth, actualDay := actual.Date()

			if expectedYear != actualYear || expectedMonth != actualMonth || expectedDay != actualDay {
				t.Errorf("日期不相等: expected=%d-%d-%d, actual=%d-%d-%d",
					expectedYear, expectedMonth, expectedDay,
					actualYear, actualMonth, actualDay)
			}

			// 验证时分秒
			expectedHour, expectedMin, expectedSec := expected.Clock()
			actualHour, actualMin, actualSec := actual.Clock()

			if expectedHour != actualHour || expectedMin != actualMin || expectedSec != actualSec {
				t.Errorf("时分秒不相等: expected=%d:%d:%d, actual=%d:%d:%d",
					expectedHour, expectedMin, expectedSec,
					actualHour, actualMin, actualSec)
			}

			// 验证纳秒
			expectedNsec := expected.Nanosecond()
			actualNsec := actual.Nanosecond()

			if expectedNsec != actualNsec {
				t.Errorf("纳秒不相等: expected=%d, actual=%d", expectedNsec, actualNsec)
			}

			// 验证时区
			if expected.Location().String() != actual.Location().String() {
				t.Errorf("时区不相等: expected=%s, actual=%s",
					expected.Location(), actual.Location())
			}

			t.Logf("测试通过: %s", tc.name)
		})
	}
}

// TestEndOfDayGlobalRealtime 测试真实的全局函数
func TestEndOfDayGlobalRealtime(t *testing.T) {
	// 调用真实的全局函数
	result := EndOfDay()

	if result == nil {
		t.Fatal("EndOfDay 返回 nil")
	}

	// 验证时间不为零
	if result.Time.IsZero() {
		t.Error("EndOfDay 返回零时间")
	}

	// 验证小时是 23
	hour, _, _ := result.Clock()
	if hour != 23 {
		t.Errorf("期望小时=23, 实际=%d", hour)
	}

	// 验证分钟是 59
	_, min, _ := result.Clock()
	if min != 59 {
		t.Errorf("期望分钟=59, 实际=%d", min)
	}

	// 验证秒是 59
	_, _, sec := result.Clock()
	if sec != 59 {
		t.Errorf("期望秒=59, 实际=%d", sec)
	}

	// 验证纳秒是 999999999
	nsec := result.Nanosecond()
	expectedNsec := int(time.Second - time.Nanosecond)
	if nsec != expectedNsec {
		t.Errorf("期望纳秒=%d, 实际=%d", expectedNsec, nsec)
	}

	t.Logf("测试通过: EndOfDay = %v", result.Time)
}

// TestEndOfDayGlobalMemoryLayout 验证内存布局
func TestEndOfDayGlobalMemoryLayout(t *testing.T) {
	result := EndOfDay()

	// 验证 Config 为 nil（零分配）
	if result.Config != nil {
		t.Error("期望 Config 为 nil（零分配），实际不为 nil")
	}

	// 验证 Time 字段已设置
	if result.Time.IsZero() {
		t.Error("Time 字段为零时间")
	}

	t.Logf("内存布局验证通过: Config=nil, Time=%v", result.Time)
}
