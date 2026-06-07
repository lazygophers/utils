package human

import (
	"fmt"
	"strings"
	"time"
)

// Duration formats a time.Duration as a human-readable string with up to two
// significant units (e.g. "1 day 3 hours"). Honors the clock-format toggle
// for HH:MM:SS rendering.
func Duration(d time.Duration) string { return formatDuration(d) }

// ClockDuration always renders d in HH:MM:SS / M:SS form regardless of the
// global SetClockFormat toggle.
func ClockDuration(d time.Duration) string { return formatClockTime(d) }

func formatDuration(d time.Duration) string {
	if d == 0 {
		if defaultClockFormat {
			return "0:00"
		}
		locale, _ := GetLocaleConfig(currentLocaleName())
		return "0 " + locale.TimeUnits.Second
	}

	if defaultClockFormat {
		return formatClockTime(d)
	}

	locale, _ := GetLocaleConfig(currentLocaleName())

	negative := d < 0
	if negative {
		d = -d
	}

	var parts []string

	days := d / (24 * time.Hour)
	d %= 24 * time.Hour
	hours := d / time.Hour
	d %= time.Hour
	minutes := d / time.Minute
	d %= time.Minute
	seconds := d / time.Second

	if days > 0 {
		parts = append(parts, fmt.Sprintf("%d %s", days, getTimeUnit(locale, locale.TimeUnits.Day, int64(days))))
		if hours > 0 {
			parts = append(parts, fmt.Sprintf("%d %s", hours, getTimeUnit(locale, locale.TimeUnits.Hour, int64(hours))))
		}
	} else if hours > 0 {
		parts = append(parts, fmt.Sprintf("%d %s", hours, getTimeUnit(locale, locale.TimeUnits.Hour, int64(hours))))
		if minutes > 0 {
			parts = append(parts, fmt.Sprintf("%d %s", minutes, getTimeUnit(locale, locale.TimeUnits.Minute, int64(minutes))))
		}
	} else if minutes > 0 {
		parts = append(parts, fmt.Sprintf("%d %s", minutes, getTimeUnit(locale, locale.TimeUnits.Minute, int64(minutes))))
		if seconds > 0 {
			parts = append(parts, fmt.Sprintf("%d %s", seconds, getTimeUnit(locale, locale.TimeUnits.Second, int64(seconds))))
		}
	} else {
		parts = append(parts, fmt.Sprintf("%d %s", seconds, getTimeUnit(locale, locale.TimeUnits.Second, int64(seconds))))
	}

	result := strings.Join(parts, " ")
	if negative {
		result = "-" + result
	}

	return result
}

func formatClockTime(d time.Duration) string {
	negative := d < 0
	if negative {
		d = -d
	}

	totalSeconds := int64(d.Seconds())
	hours := totalSeconds / 3600
	minutes := (totalSeconds % 3600) / 60
	seconds := totalSeconds % 60

	var result string
	if hours > 0 {
		if seconds > 0 {
			result = fmt.Sprintf("%d:%02d:%02d", hours, minutes, seconds)
		} else {
			result = fmt.Sprintf("%d:%02d", hours, minutes)
		}
	} else if minutes > 0 {
		result = fmt.Sprintf("%d:%02d", minutes, seconds)
	} else {
		result = fmt.Sprintf("0:%02d", seconds)
	}

	if negative {
		result = "-" + result
	}

	return result
}
