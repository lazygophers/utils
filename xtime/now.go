package xtime

import "time"

// Now 获取当前时间并封装为xtime.Time
// 返回包装后的当前时间对象
func Now() *Time {
	return With(time.Now())
}

func NowUnix() int64 {
	return Now().Unix()
}

func NowUnixMilli() int64 {
	return Now().UnixMilli()
}

// BeginningOfMinute 获取当前分钟的起始时间
func (p *Time) BeginningOfMinute() *Time {
	return With(p.Truncate(time.Minute))
}

// BeginningOfHour 获取当前小时的起始时间
func (p *Time) BeginningOfHour() *Time {
	y, m, d := p.Date()
	return With(time.Date(y, m, d, p.Time.Hour(), 0, 0, 0, p.Time.Location()))
}

// BeginningOfDay 获取当前日期的起始时间（00:00:00）
func (p *Time) BeginningOfDay() *Time {
	y, m, d := p.Date()
	return With(time.Date(y, m, d, 0, 0, 0, 0, p.Time.Location()))
}

// BeginningOfWeek 获取当前周的起始时间（根据WeekStartDay参数决定周起始日，默认周日）
// WeekStartDay determines the starting day of the week (default: Sunday)
// If the week doesn't start on Sunday, adjusts the calculation accordingly
func (p *Time) BeginningOfWeek() *Time {
	t := p.BeginningOfDay()
	weekday := int(t.Weekday())

	if p.WeekStartDay != time.Sunday {
		weekStartDayInt := int(p.WeekStartDay)

		if weekday < weekStartDayInt {
			weekday = weekday + 7 - weekStartDayInt
		} else {
			weekday = weekday - weekStartDayInt
		}
	}
	return With(t.AddDate(0, 0, -weekday))
}

// BeginningOfMonth 获取当前月份的起始时间（1号00:00:00）
func (p *Time) BeginningOfMonth() *Time {
	y, m, _ := p.Date()
	return With(time.Date(y, m, 1, 0, 0, 0, 0, p.Location()))
}

func (p *Time) BeginningOfQuarter() *Time {
	month := p.BeginningOfMonth()
	offset := (int(month.Month()) - 1) % 3
	return With(month.AddDate(0, -offset, 0))
}

func (p *Time) BeginningOfHalf() *Time {
	month := p.BeginningOfMonth()
	offset := (int(month.Month()) - 1) % 6
	return With(month.AddDate(0, -offset, 0))
}

func (p *Time) BeginningOfYear() *Time {
	y, _, _ := p.Date()
	return With(time.Date(y, time.January, 1, 0, 0, 0, 0, p.Location()))
}

func (p *Time) EndOfMinute() *Time {
	return With(p.BeginningOfMinute().Add(time.Minute - time.Nanosecond))
}

// EndOfHour 获取当前小时的结束时间（下一小时开始前1纳秒）
// 设置时间为当前小时的23:59:59.999999999
func (p *Time) EndOfHour() *Time {
	return With(p.BeginningOfHour().Add(time.Hour - time.Nanosecond))
}

// EndOfDay 获取当前日期的结束时间（次日00:00前1纳秒）
// 设置时间为当天23:59:59.999999999
func (p *Time) EndOfDay() *Time {
	y, m, d := p.Date()
	return With(time.Date(y, m, d, 23, 59, 59, int(time.Second-time.Nanosecond), p.Location()))
}

// EndOfWeek 获取当前周的结束时间（下周起始日前1纳秒）
// 返回下周第一天的前一纳秒（星期日为最后一天）
func (p *Time) EndOfWeek() *Time {
	return With(p.BeginningOfWeek().AddDate(0, 0, 7).Add(-time.Nanosecond))
}

// EndOfMonth 获取当前月份的结束时间（下月1日前1纳秒）
// 返回下个月第一天的前一纳秒
func (p *Time) EndOfMonth() *Time {
	return With(p.BeginningOfMonth().AddDate(0, 1, 0).Add(-time.Nanosecond))
}

// EndOfQuarter 获取当前季度的结束时间（下一季度首日前1纳秒）
func (p *Time) EndOfQuarter() *Time {
	return With(p.BeginningOfQuarter().AddDate(0, 3, 0).Add(-time.Nanosecond))
}

// EndOfHalf 获取当前半年的结束时间（下半年首日前1纳秒）
func (p *Time) EndOfHalf() *Time {
	return With(p.BeginningOfHalf().AddDate(0, 6, 0).Add(-time.Nanosecond))
}

// EndOfYear 获取当前年的结束时间（下年首日前1纳秒）
// 返回下一年第一天的前一纳秒
func (p *Time) EndOfYear() *Time {
	return With(p.BeginningOfYear().AddDate(1, 0, 0).Add(-time.Nanosecond))
}

// Quarter 获取当前时间所属的季度编号（1-4季度）
func (p *Time) Quarter() uint {
	return (uint(p.Month())-1)/3 + 1
}

func BeginningOfMinute() *Time {
	return With(time.Now()).BeginningOfMinute()
}

func BeginningOfHour() *Time {
	return With(time.Now()).BeginningOfHour()
}

func BeginningOfDay() *Time {
	return With(time.Now()).BeginningOfDay()
}

func BeginningOfWeek() *Time {
	return With(time.Now()).BeginningOfWeek()
}

func BeginningOfMonth() *Time {
	return With(time.Now()).BeginningOfMonth()
}

func BeginningOfQuarter() *Time {
	return With(time.Now()).BeginningOfQuarter()
}

func BeginningOfYear() *Time {
	return With(time.Now()).BeginningOfYear()
}

func EndOfMinute() *Time {
	return With(time.Now()).EndOfMinute()
}

func EndOfHour() *Time {
	return With(time.Now()).EndOfHour()
}

func EndOfDay() *Time {
	return With(time.Now()).EndOfDay()
}

func EndOfWeek() *Time {
	return With(time.Now()).EndOfWeek()
}

func EndOfMonth() *Time {
	return With(time.Now()).EndOfMonth()
}

func EndOfQuarter() *Time {
	return With(time.Now()).EndOfQuarter()
}

// EndOfYear 获取当前年的结束时间（下年首日前1纳秒）
func EndOfYear() *Time {
	return With(time.Now()).EndOfYear()
}

func Quarter() uint {
	return With(time.Now()).Quarter()
}
