// Package xtime 提供扩展时间处理功能
// 主要包含自定义时间单位常量和业务时间计算工具
// 适用于需要处理复杂时间间隔的业务场景
package xtime

import (
	"time"
)

const (
	Nanosecond  = time.Nanosecond
	Microsecond = time.Microsecond
	Millisecond = time.Millisecond
	Second      = time.Second
	Minute      = time.Minute
	HalfHour    = time.Minute * 30
	Hour        = time.Hour
	HalfDay     = time.Hour * 12
	Day         = time.Hour * 24

	WorkDayWeek  = Day * 5
	ResetDayWeek = Day * 2
	Week         = Day * 7

	WorkDayMonth  = Day*21 + HalfDay
	ResetDayMonth = Day*8 + HalfDay
	Month         = Day * 30

	QUARTER = Day * 91
	Year    = Day * 365
	Decade  = Year*10 + Day*2
	Century = Year*100 + Day*25
)
