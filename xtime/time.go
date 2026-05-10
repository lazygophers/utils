package xtime

import (
	"time"

	"github.com/jinzhu/now"
	"github.com/lazygophers/utils/randx"
)

type Config struct {
	WeekStartDay  time.Weekday
	TimeLocation *time.Location
	TimeFormats  []string
	Monotonic    time.Time
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

	return With(t), nil
}

func With(t time.Time) *Time {
	return &Time{
		Time: t,
		Config: &Config{
			WeekStartDay:  time.Monday,
			TimeLocation: time.Local,
			TimeFormats:  []string{},
			Monotonic:    time.Now(),
		},
	}
}

func RandSleep(s ...time.Duration) {
	time.Sleep(randx.TimeDuration4Sleep(s...))
}

// Elapsed 返回从创建时间到现在经过的时间（基于单调时钟）
func (t *Time) Elapsed() time.Duration {
	if t.Config != nil && !t.Monotonic.IsZero() {
		return time.Since(t.Monotonic)
	}
	return time.Since(t.Time)
}

// UTC 转换为 UTC 时区时间
func (t *Time) UTC() *Time {
	return &Time{
		Time:   t.Time.UTC(),
		Config: t.Config,
	}
}

// Local 转换为本地时区时间
func (t *Time) Local() *Time {
	return &Time{
		Time:   t.Time.Local(),
		Config: t.Config,
	}
}

// In 转换为指定时区时间
func (t *Time) In(loc *time.Location) *Time {
	return &Time{
		Time:   t.Time.In(loc),
		Config: t.Config,
	}
}

// Since 返回从 t 到现在经过的时间
func (t *Time) Since(d time.Duration) time.Duration {
	return time.Since(t.Time) - d
}

// Until 返回从现在到 t 的时间
func (t *Time) Until(d time.Duration) time.Duration {
	return t.Time.Sub(time.Now()) - d
}
