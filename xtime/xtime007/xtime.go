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

	Month     = Day * 30
	RestMonth = RestDay * 0
	WorkMonth = Day - RestMonth

	Quarter     = Day * 91
	RestQuarter = RestDay * 0
	WorkQuarter = Day - RestQuarter

	Year     = Day * 365
	RestYear = RestDay * 0
	WorkYear = Year - RestYear
)
