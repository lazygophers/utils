package xtime

import (
	"github.com/lazygophers/utils/randx"
	"time"
)

func Now() time.Time {
	return time.Now()
}

func NowUnix() int64 {
	return time.Now().Unix()
}

func NowUnixMilli() int64 {
	return time.Now().UnixMilli()
}

func RandSleep(s ...time.Duration) {
	time.Sleep(randx.TimeDuration4Sleep(s...))
}
