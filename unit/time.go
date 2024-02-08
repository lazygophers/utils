package unit

import (
	"bytes"
	"github.com/lazygophers/utils/xtime"
	"strconv"
	"time"
)

func DurationYearMonthDay(t time.Duration) string {
	years := int(t.Hours() / 24 / 365)
	t = t - time.Duration(years)*xtime.Year

	mouths := int(t.Hours() / 24 / 30)
	t = t - time.Duration(mouths)*xtime.Month

	days := int(t.Hours() / 24)
	t = t - time.Duration(days)*xtime.Day

	var b bytes.Buffer
	if years > 0 {
		b.WriteString(strconv.Itoa(years))
		b.WriteString("年")
	}

	if mouths > 0 {
		b.WriteString(strconv.Itoa(mouths))
		b.WriteString("月")
	}

	if days > 0 {
		b.WriteString(strconv.Itoa(days))
		b.WriteString("天")
	}

	// 如果都为0
	if b.Len() == 0 {
		b.WriteString("0天")
	}

	return b.String()
}

func DurationMonthDayHour(t time.Duration) string {
	mouths := int(t.Hours() / 24 / 30)
	t = t - time.Duration(mouths)*xtime.Month

	days := int(t.Hours() / 24)
	t = t - time.Duration(days)*xtime.Day

	hours := int(t.Hours())
	t = t - time.Duration(hours)*xtime.Hour

	var b bytes.Buffer
	if mouths > 0 {
		b.WriteString(strconv.Itoa(mouths))
		b.WriteString("月")
	}

	if days > 0 {
		b.WriteString(strconv.Itoa(days))
		b.WriteString("天")
	}

	if hours > 0 {
		b.WriteString(strconv.Itoa(hours))
		b.WriteString("小时")
	}

	// 如果都为0
	if b.Len() == 0 {
		b.WriteString("0小时")
	}

	return b.String()
}

func DurationMinuteSecond(t time.Duration) string {
	minutes := int(t.Minutes())
	t = t - time.Duration(minutes)*xtime.Minute

	seconds := t.Seconds()

	var b bytes.Buffer
	if minutes > 1 {
		b.WriteString(strconv.Itoa(minutes))
		b.WriteString("分")
	}

	if seconds > 0 {
		b.WriteString(strconv.Itoa(int(seconds)))
		b.WriteString("秒")
	}

	// 如果都为0
	if b.Len() == 0 {
		b.WriteString("0秒")
	}

	return b.String()
}

func DurationYearMonthDayHourMinuteSecond(t time.Duration) string {
	years := int(t.Hours() / 24 / 365)
	t = t - time.Duration(years)*xtime.Year

	mouths := int(t.Hours() / 24 / 30)
	t = t - time.Duration(mouths)*xtime.Month

	days := int(t.Hours() / 24)
	t = t - time.Duration(days)*xtime.Day

	hours := int(t.Hours())
	t = t - time.Duration(hours)*xtime.Hour

	minutes := int(t.Minutes())
	t = t - time.Duration(minutes)*xtime.Minute

	seconds := t.Seconds()

	var b bytes.Buffer
	if years > 0 {
		b.WriteString(strconv.Itoa(years))
		b.WriteString("年")
	}

	if mouths > 0 {
		b.WriteString(strconv.Itoa(mouths))
		b.WriteString("月")
	}

	if days > 0 {
		b.WriteString(strconv.Itoa(days))
		b.WriteString("天")
	}

	if hours > 0 {
		b.WriteString(strconv.Itoa(hours))
		b.WriteString("小时")
	}

	if minutes > 0 {
		b.WriteString(strconv.Itoa(minutes))
		b.WriteString("分")
	}

	if seconds > 0 {
		b.WriteString(strconv.Itoa(int(seconds)))
		b.WriteString("秒")
	}

	// 如果都为0
	if b.Len() == 0 {
		b.WriteString("0秒")
	}

	return b.String()
}

func TimeYearMonthDayHourMinute(t time.Time) string {
	return t.Format("2006年01月02日15点04")
}

func TimeYearMonthDayHourMinuteSecond(t time.Time) string {
	return t.Format("2006年01月02日15点04分05")
}
