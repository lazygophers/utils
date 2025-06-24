// xtime955 提供基于标准库time包的时间常量扩展
// 包含工作日/休息日周期定义及时间单位计算

package xtime955

import "time"

// 基础时间单位定义
const (
	// Nanosecond 表示1纳秒（1e-9秒）
	Nanosecond = time.Nanosecond

	// Microsecond 表示1微秒（1e-6秒）
	Microsecond = time.Microsecond

	// Millisecond 表示1毫秒（1e-3秒）
	Millisecond = time.Millisecond

	// Second 表示1标准秒（1e0秒）
	Second = time.Second

	// Minute 表示1标准分钟（60秒）
	Minute = time.Minute
)

// 工作时间周期定义
const (
	// HalfHour 表示30分钟（工作日常用单位）
	HalfHour = time.Minute * 30

	// Hour 表示1标准小时（60分钟）
	Hour = time.Hour

	// Day 表示1标准日（24小时）
	Day = time.Hour * 24

	// WorkDay 表示标准工作日时长（8小时）
	WorkDay = time.Hour * 8

	// RestDay 表示每日休息时间（Day - WorkDay）
	RestDay = Day - WorkDay

	// Week 表示1标准周（7日）
	Week = Day * 7

	// WorkWeek 表示标准工作周时长（5工作日）
	WorkWeek = WorkDay * 5

	// RestWeek 表示周休息时间（Week - WorkWeek）
	RestWeek = Week - WorkWeek
)

// 季度周期定义
const (
	// Month 表示1标准月（30日）
	Month = Day * 30

	// WorkMonth 表示标准工作月时长（22个工作日）
	// 基于Month常量计算，扣除8小时/日休息时间
	// 基于Month常量计算，扣除8小时/日休息时间
	WorkMonth = Day * 22

	// RestMonth 表示月休息时间（Month - WorkMonth）
	RestMonth = Month - WorkMonth

	// Quarter 表示1标准季度（91天）
	Quarter = Day * 91

	// WorkQuarter 表示标准工作季度时长（27个工作日）
	WorkQuarter = WorkMonth * 3

	// RestQuarter 表示季度休息时间（Quarter - WorkQuarter）
	RestQuarter = Quarter - WorkQuarter
)

// 年度周期定义
const (
	// Year 表示1标准年（365日）
	Year = Day * 365

	// WorkYear 表示标准工作年时长（250天）
	WorkYear = WorkDay * 250

	// RestYear 表示年度休息时间（Year - WorkYear）
	RestYear = Year - WorkYear
)
