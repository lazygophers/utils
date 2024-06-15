package xtime_test

import (
	"github.com/lazygophers/utils/xtime"
	"testing"
	"time"
)

func TestLunar(t *testing.T) {
	t.Log(xtime.WithLunarTime(xtime.BeginningOfDay()).Animal())
	t.Log(xtime.WithLunarTime(xtime.BeginningOfYear()).Animal())

	t.Log(xtime.CalcSolarterm(xtime.BeginningOfYear().Time).String())
	t.Log(xtime.CalcSolarterm(xtime.BeginningOfYear().Time).Prev().String())

	t.Log(xtime.CalcSolarterm(time.Now()).IsInDay(time.Now()))
	t.Log(xtime.CalcSolarterm(time.Now()).String())
	t.Log(xtime.CalcSolarterm(time.Unix(1716134400, 0)).String())
	t.Log(xtime.CalcSolarterm(time.Unix(1716134400, 0)).Next().String())
}
