//go:build lang_en || lang_all

package human

// Register English language support
func init() {
	RegisterLocale("en", &Locale{
		Language:      "en",
		Region:        "US",
		ByteUnits:     []string{"B", "KB", "MB", "GB", "TB", "PB"},
		SpeedUnits:    []string{"B/s", "KB/s", "MB/s", "GB/s", "TB/s", "PB/s"},
		BitSpeedUnits: []string{"bps", "Kbps", "Mbps", "Gbps", "Tbps", "Pbps"},

		TimeUnits: TimeUnits{
			Nanosecond:  "ns",
			Microsecond: "µs",
			Millisecond: "ms",
			Second:      "s",
			Minute:      "min",
			Hour:        "h",
			Day:         "day",
			Week:        "week",
			Month:       "month",
			Year:        "year",
		},

		RelativeTimeUnits: RelativeTimeUnits{
			JustNow:    "just now",
			SecondsAgo: "seconds ago",
			MinuteAgo:  "minute ago",
			MinutesAgo: "minutes ago",
			HourAgo:    "hour ago",
			HoursAgo:   "hours ago",
			DayAgo:     "day ago",
			DaysAgo:    "days ago",
			WeekAgo:    "week ago",
			WeeksAgo:   "weeks ago",
			MonthAgo:   "month ago",
			MonthsAgo:  "months ago",
			YearAgo:    "year ago",
			YearsAgo:   "years ago",
		},

		NumberUnits: NumberUnits{
			Thousand: "thousand",
			Million:  "million",
			Billion:  "billion",
			Trillion: "trillion",
		},
	})
}
