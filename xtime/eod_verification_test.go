package xtime

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

// TestEndOfDay_Correctness 验证 EndOfDay 功能正确性
func TestEndOfDay_Correctness(t *testing.T) {
	testCases := []struct {
		name     string
		input    time.Time
		expected string // ISO 8601 格式
	}{
		{
			name:     "中午时间",
			input:    time.Date(2024, 5, 11, 15, 30, 45, 123456789, time.Local),
			expected: "2024-05-11T23:59:59.999999999",
		},
		{
			name:     "午夜时间",
			input:    time.Date(2024, 5, 11, 0, 0, 0, 0, time.Local),
			expected: "2024-05-11T23:59:59.999999999",
		},
		{
			name:     "当天最后一秒",
			input:    time.Date(2024, 5, 11, 23, 59, 59, 999999999, time.Local),
			expected: "2024-05-11T23:59:59.999999999",
		},
		{
			name:     "跨月边界",
			input:    time.Date(2024, 1, 31, 12, 0, 0, 0, time.Local),
			expected: "2024-01-31T23:59:59.999999999",
		},
		{
			name:     "闰年2月",
			input:    time.Date(2024, 2, 29, 10, 30, 0, 0, time.Local),
			expected: "2024-02-29T23:59:59.999999999",
		},
		{
			name:     "夏令时边界",
			input:    time.Date(2024, 3, 10, 1, 30, 0, 0, time.Local),
			expected: "2024-03-10T23:59:59.999999999",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			wrapper := With(tc.input)
			result := wrapper.EndOfDay()

			// 验证时间部分
			hour, min, sec := result.Clock()
			assert.Equal(t, 23, hour, "小时应为23")
			assert.Equal(t, 59, min, "分钟应为59")
			assert.Equal(t, 59, sec, "秒应为59")

			// 验证纳秒
			nanos := result.Nanosecond()
			assert.Equal(t, 999999999, nanos, "纳秒应为999999999")

			// 验证日期部分不变
			year, month, day := result.Date()
			expYear, expMonth, expDay := tc.input.Date()
			assert.Equal(t, expYear, year, "年份应相同")
			assert.Equal(t, expMonth, month, "月份应相同")
			assert.Equal(t, expDay, day, "日期应相同")

			// 验证 ISO 格式
			actual := result.Format("2006-01-02T15:04:05.999999999")
			assert.Equal(t, tc.expected, actual)
		})
	}
}

// TestEndOfDay_BeforeEndOfNextDay 验证 EndOfDay 结果在次日00:00之前
func TestEndOfDay_BeforeEndOfNextDay(t *testing.T) {
	base := time.Date(2024, 5, 11, 15, 30, 45, 123456789, time.Local)
	wrapper := With(base)
	eod := wrapper.EndOfDay()

	// 获取次日开始时间
	nextDay := eod.Add(time.Nanosecond)

	// 验证次日开始时间为00:00:00
	hour, min, sec := nextDay.Clock()
	assert.Equal(t, 0, hour, "次日应为00点")
	assert.Equal(t, 0, min, "次日应为00分")
	assert.Equal(t, 0, sec, "次日应为00秒")

	// 验证日期已递增
	_, _, eodDay := eod.Date()
	_, _, nextDayDay := nextDay.Date()
	assert.Equal(t, eodDay+1, nextDayDay, "日期应递增1天")
}

// TestEndOfDay_ConfigPreservation 验证 Config 被正确保留
func TestEndOfDay_ConfigPreservation(t *testing.T) {
	base := time.Date(2024, 5, 11, 15, 30, 45, 123456789, time.Local)
	wrapper := With(base)

	// 修改 Config
	wrapper.Config.WeekStartDay = time.Sunday
	wrapper.Config.TimeLocation = time.UTC

	result := wrapper.EndOfDay()

	assert.NotNil(t, result.Config, "Config 不应为 nil")
	assert.Equal(t, time.Sunday, result.Config.WeekStartDay, "WeekStartDay 应保留")
	assert.Equal(t, time.UTC, result.Config.TimeLocation, "TimeLocation 应保留")
}

// TestEndOfDay_NilConfig 安全处理 nil Config
func TestEndOfDay_NilConfig(t *testing.T) {
	base := time.Date(2024, 5, 11, 15, 30, 45, 123456789, time.Local)

	// 直接构造 Time，Config 为 nil
	wrapper := &Time{Time: base, Config: nil}

	result := wrapper.EndOfDay()

	assert.NotNil(t, result.Config, "应创建新 Config")
	assert.NotNil(t, result.Time, "Time 应正确设置")
}

// TestEndOfDay_BeginningOfDayConsistency 与 BeginningOfDay 的一致性
func TestEndOfDay_BeginningOfDayConsistency(t *testing.T) {
	base := time.Date(2024, 5, 11, 15, 30, 45, 123456789, time.Local)
	wrapper := With(base)

	bod := wrapper.BeginningOfDay()
	eod := wrapper.EndOfDay()

	// 验证日期相同
	bodYear, bodMonth, bodDay := bod.Date()
	eodYear, eodMonth, eodDay := eod.Date()

	assert.Equal(t, bodYear, eodYear, "年份应相同")
	assert.Equal(t, bodMonth, eodMonth, "月份应相同")
	assert.Equal(t, bodDay, eodDay, "日期应相同")

	// 验证 EndOfDay - BeginningOfDay ≈ 24小时
	diff := eod.Time.Sub(bod.Time)
	expectedDiff := 24*time.Hour - time.Nanosecond
	assert.Equal(t, expectedDiff, diff, "时间差应为24小时减1纳秒")
}

// TestEndOfDay_PerformanceComparison 性能对比基准
func TestEndOfDay_PerformanceComparison(t *testing.T) {
	if testing.Short() {
		t.Skip("跳过性能测试")
	}

	// 这个测试仅用于验证性能提升，实际基准测试在 eod_bench_test.go
	times := genEODTestTimes(10000)
	wrapper := make([]*Time, len(times))
	for i := range times {
		wrapper[i] = With(times[i])
	}

	// 验证不崩溃且结果正确
	for i := 0; i < 100; i++ {
		result := wrapper[i].EndOfDay()
		hour, _, _ := result.Clock()
		if hour != 23 {
			t.Errorf("第 %d 次: 小时应为23，得到 %d", i, hour)
		}
	}
}
