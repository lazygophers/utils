package xtime996

import "time"

// 时间单位常量定义
//
// 该包基于标准库time包的时间单位，定义了工作时间相关的自定义时间常量
// 包含工作日/工作周/工作月/工作年的计算
const (
	// Nanosecond 纳秒时间单位（1e-9秒）
	// 值等于标准库time.Nanosecond
	Nanosecond = time.Nanosecond
	// Microsecond 微秒时间单位（1e-6秒）
	// 值等于标准库time.Microsecond
	Microsecond = time.Microsecond
	// Millisecond 毫秒时间单位（1e-3秒）
	// 值等于标准库time.Millisecond
	Millisecond = time.Millisecond
	// Second 秒时间单位（1秒）
	// 值等于标准库time.Second
	Second = time.Second
	// Minute 分钟时间单位（60秒）
	// 值等于标准库time.Minute
	Minute = time.Minute
	// HalfHour 半小时时长（30分钟）
	// 等于标准库time.Minute * 30
	HalfHour = time.Minute * 30
	// Hour 小时时间单位（60分钟）
	// 值等于标准库time.Hour
	Hour = time.Hour

	// Day 天时间单位（24小时）
	// 值等于标准库time.Hour * 24
	Day = time.Hour * 24
	// WorkDay 工作日时长（12小时）
	// 值等于标准库time.Hour * 12
	WorkDay = time.Hour * 12
	// RestDay 休息日时长（12小时）
	// 值等于Day - WorkDay
	RestDay = Day - WorkDay

	// Week 周时间单位（7天）
	// 等于标准库time.Hour * 24 * 7
	Week = Day * 7
	// WorkWeek 工作周时长（6个工作日）
	// 等于WorkDay * 6（每个工作日12小时）
	WorkWeek = WorkDay * 6
	// RestWeek 休息周时长（1天12小时）
	// 等于Week - WorkWeek
	RestWeek = Week - WorkWeek

	// Month 月时间单位（30天）
	// 等于标准库time.Hour * 24 * 30
	Month = Day * 30
	// RestMonth 休息月时长（4个休息日）
	// 等于RestDay * 4（每个休息日12小时）
	// RestMonth 休息月时长（4个休息日）
	// 等于RestDay * 4（每个休息日12小时）
	RestMonth = RestDay * 4
	// WorkMonth 工作月时长（26天）
	// 等于Month - RestMonth（30天-4个休息日）
	WorkMonth = Day*30 - RestMonth

	// Quarter 季度时间单位（91天）
	Quarter = Day * 91
	// RestQuarter 休息季度时长（14个休息日）
	// 等于RestDay * 14（每个休息日12小时）
	// RestQuarter 休息季度时长（14个休息日）
	// 等于RestDay * 14（每个休息日12小时）
	RestQuarter = RestDay * 14
	// WorkQuarter 工作季度时长（77天）
	// 基于Quarter常量计算，扣除14个标准休息日（91天-14天休息日）
	WorkQuarter = Day*91 - RestQuarter

	// Year 年时间单位（365天）
	// 等于标准库time.Hour * 24 * 365
	// Year 年时间单位（365天）
	// 等于标准库time.Hour * 24 * 365
	Year = Day * 365
	// RestYear 休息年时长（58个休息日）
	// 等于RestDay * 58（每个休息日12小时）
	// RestYear 休息年时长（58个休息日）
	// 等于RestDay * 58（每个休息日12小时）
	RestYear = RestDay * 58
	// WorkYear 工作年时长（307天）
	// 基于Year常量计算，扣除58个标准休息日（365天-58天休息日）
	// WorkYear 表示标准工作年时长（307天）= Year - RestYear
	WorkYear = Year - RestYear
)
