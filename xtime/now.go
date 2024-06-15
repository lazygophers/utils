package xtime

import "time"

func Now() *Time {
	return With(time.Now())
}

func NowUnix() int64 {
	return Now().Unix()
}

func NowUnixMilli() int64 {
	return Now().UnixMilli()
}

func (p *Time) BeginningOfMinute() *Time {
	return With(p.Truncate(time.Minute))
}

func (p *Time) BeginningOfHour() *Time {
	y, m, d := p.Date()
	return With(time.Date(y, m, d, p.Time.Hour(), 0, 0, 0, p.Time.Location()))
}

func (p *Time) BeginningOfDay() *Time {
	y, m, d := p.Date()
	return With(time.Date(y, m, d, 0, 0, 0, 0, p.Time.Location()))
}

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

// EndOfHour end of hour
func (p *Time) EndOfHour() *Time {
	return With(p.BeginningOfHour().Add(time.Hour - time.Nanosecond))
}

// EndOfDay end of day
func (p *Time) EndOfDay() *Time {
	y, m, d := p.Date()
	return With(time.Date(y, m, d, 23, 59, 59, int(time.Second-time.Nanosecond), p.Location()))
}

// EndOfWeek end of week
func (p *Time) EndOfWeek() *Time {
	return With(p.BeginningOfWeek().AddDate(0, 0, 7).Add(-time.Nanosecond))
}

// EndOfMonth end of month
func (p *Time) EndOfMonth() *Time {
	return With(p.BeginningOfMonth().AddDate(0, 1, 0).Add(-time.Nanosecond))
}

// EndOfQuarter end of quarter
func (p *Time) EndOfQuarter() *Time {
	return With(p.BeginningOfQuarter().AddDate(0, 3, 0).Add(-time.Nanosecond))
}

// EndOfHalf end of half year
func (p *Time) EndOfHalf() *Time {
	return With(p.BeginningOfHalf().AddDate(0, 6, 0).Add(-time.Nanosecond))
}

// EndOfYear end of year
func (p *Time) EndOfYear() *Time {
	return With(p.BeginningOfYear().AddDate(1, 0, 0).Add(-time.Nanosecond))
}

// Quarter returns the yearly quarter
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

// EndOfYear end of year
func EndOfYear() *Time {
	return With(time.Now()).EndOfYear()
}

func Quarter() uint {
	return With(time.Now()).Quarter()
}
