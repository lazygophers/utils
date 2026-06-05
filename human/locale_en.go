package human

// Register English language configuration
func init() {
	RegisterLocale("en", &Locale{
		Language:      "en",
		Region:        "US",
		ByteUnits:     []string{"B", "KB", "MB", "GB", "TB", "PB"},
		SpeedUnits:    []string{"B/s", "KB/s", "MB/s", "GB/s", "TB/s", "PB/s"},
		BitSpeedUnits: []string{"bps", "Kbps", "Mbps", "Gbps", "Tbps", "Pbps"},

		TimeUnits: TimeUnits{
			Nanosecond:  "ns",
			Microsecond: "μs",
			Millisecond: "ms",
			Second:      "second",
			Minute:      "minute",
			Hour:        "hour",
			Day:         "day",
			Week:        "week",
			Month:       "month",
			Year:        "year",

			Seconds: "seconds",
			Minutes: "minutes",
			Hours:   "hours",
			Days:    "days",
			Weeks:   "weeks",
			Months:  "months",
			Years:   "years",
		},

		RelativeTime: RelativeTimeStrings{
			JustNow:    "just now",
			SecondsAgo: "%d seconds ago",
			MinutesAgo: "%d minutes ago",
			HoursAgo:   "%d hours ago",
			DaysAgo:    "%d days ago",
			WeeksAgo:   "%d weeks ago",
			MonthsAgo:  "%d months ago",
			YearsAgo:   "%d years ago",

			In:           "in",
			SecondsLater: "in %d seconds",
			MinutesLater: "in %d minutes",
			HoursLater:   "in %d hours",
			DaysLater:    "in %d days",
			WeeksLater:   "in %d weeks",
			MonthsLater:  "in %d months",
			YearsLater:   "in %d years",
		},

		NumberFormat: NumberFormat{
			DecimalSeparator:  ".",
			ThousandSeparator: ",",
			LargeNumberUnits:  []string{"K", "M", "B", "T"},
		},

		Common: CommonStrings{
			And: "and",
			Or:  "or",
		},
	})
}
