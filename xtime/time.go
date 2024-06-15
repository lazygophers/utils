package xtime

import (
	"github.com/jinzhu/now"
	"github.com/lazygophers/utils/randx"
	"time"
)

type Config struct {
	WeekStartDay time.Weekday
	TimeLocation *time.Location
	TimeFormats  []string
}

type Time struct {
	time.Time
	*Config
}

func MustParse(str ...string) *Time {
	t, err := Parse(str...)
	if err != nil {
		panic(err)
	}
	return t
}

func Parse(strs ...string) (*Time, error) {
	t, err := now.Parse(strs...)
	if err != nil {
		return nil, err
	}

	now.BeginningOfDay()

	return With(t), nil
}

func With(t time.Time) *Time {
	return &Time{
		Time: t,
		Config: &Config{
			WeekStartDay: time.Monday,
			TimeLocation: time.Local,
			TimeFormats:  []string{},
		},
	}
}

func RandSleep(s ...time.Duration) {
	time.Sleep(randx.TimeDuration4Sleep(s...))
}
