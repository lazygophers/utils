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
)
