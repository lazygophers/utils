package xtime_test

import (
	"github.com/lazygophers/utils/xtime"
	"testing"
	"time"
)

func TestBegin(t *testing.T) {
	t.Log(xtime.BeginningOfDay().Format(time.DateTime))
	t.Log(xtime.EndOfDay().Format(time.DateTime))
	t.Log(xtime.BeginningOfHour().Format(time.DateTime))
	t.Log(xtime.EndOfHour().Format(time.DateTime))
	t.Log(xtime.BeginningOfQuarter().Format(time.DateTime))
	t.Log(xtime.EndOfQuarter().Format(time.DateTime))
	t.Log(xtime.BeginningOfWeek().Format(time.DateTime))
	t.Log(xtime.EndOfWeek().Format(time.DateTime))
}
