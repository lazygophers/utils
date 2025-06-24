// xtime007 包提供扩展的时间间隔定义，基于标准库 time 包的常量进行业务定制
// 该包通过预定义常用时间单位组合，简化定时任务和工时计算场景的开发
//
// 包含以下业务时间模型：
// - 标准工作日/周/月/年模型（全天候制）
// - 休息时间模型（适用于非工作时段计算）
// - 季度和年度复合时间单位
//
// 特别注意：
// 1. 所有时间单位基于 time.Duration 类型
// 2. 月份默认按30天计算，季度按91天（约3个月）
// 3. 工作时间模型均采用全天候定义（24小时制）
package xtime007

import "time"

const (
	Nanosecond  = time.Nanosecond
	Microsecond = time.Microsecond
	Millisecond = time.Millisecond
	Second      = time.Second
	Minute      = time.Minute
	HalfHour    = time.Minute * 30
	Hour        = time.Hour

	Day     = time.Hour * 24
	WorkDay = time.Hour * 24
	RestDay = Day - WorkDay

	Week     = Day * 7
	WorkWeek = WorkDay * 7
	RestWeek = Week - WorkWeek

	Month = Day * 30
	// RestMonth 表示月度休息时间总时长，定义为0
	// 当前实现采用无休息日模型，适用于连续运行场景
	RestMonth = RestDay * 0
	// WorkMonth 表示标准月度工作时间，定义为30天
	// 无休息日模型，适用于连续运行场景
	// 基于Month常量计算，确保与业务时间模型一致
	WorkMonth = Month

	// Quarter 表示季度时间基准（91天）用于业务场景
	Quarter = Day * 91
	// RestQuarter 表示季度休息时间总时长，定义为0
	// 当前实现采用无休息日模型，适用于连续运行场景
	RestQuarter = RestDay * 0
	// WorkQuarter 表示标准季度工作时间，定义为91天
	// 基于Quarter常量计算，确保与业务时间模型一致
	WorkQuarter = Quarter

	// Year 表示年度时间基准，定义为365天
	// 适用于需要按自然年计算的业务场景
	Year = Day * 365
	// RestYear 表示年度休息时间总时长，定义为0
	// 当前实现采用无休息日模型，适用于连续运行场景
	RestYear = RestDay * 0
	// WorkYear 表示标准年度工作时间（365天）
	WorkYear = Year
)
