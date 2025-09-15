//go:build i18n_de || i18n_all

package human

// 注册德语语言支持
func init() {
	RegisterLocale("de", &Locale{
		Language:      "de",
		Region:        "DE",
		ByteUnits:     []string{"B", "KB", "MB", "GB", "TB", "PB"},
		SpeedUnits:    []string{"B/s", "KB/s", "MB/s", "GB/s", "TB/s", "PB/s"},
		BitSpeedUnits: []string{"bps", "Kbps", "Mbps", "Gbps", "Tbps", "Pbps"},

		TimeUnits: TimeUnits{
			Nanosecond:  "ns",
			Microsecond: "μs",
			Millisecond: "ms",
			Second:      "Sekunde",
			Minute:      "Minute",
			Hour:        "Stunde",
			Day:         "Tag",
			Week:        "Woche",
			Month:       "Monat",
			Year:        "Jahr",

			Seconds: "Sekunden",
			Minutes: "Minuten",
			Hours:   "Stunden",
			Days:    "Tage",
			Weeks:   "Wochen",
			Months:  "Monate",
			Years:   "Jahre",
		},

		RelativeTime: RelativeTimeStrings{
			JustNow:    "gerade eben",
			SecondsAgo: "vor %d Sekunden",
			MinutesAgo: "vor %d Minuten",
			HoursAgo:   "vor %d Stunden",
			DaysAgo:    "vor %d Tagen",
			WeeksAgo:   "vor %d Wochen",
			MonthsAgo:  "vor %d Monaten",
			YearsAgo:   "vor %d Jahren",

			In:           "in",
			SecondsLater: "in %d Sekunden",
			MinutesLater: "in %d Minuten",
			HoursLater:   "in %d Stunden",
			DaysLater:    "in %d Tagen",
			WeeksLater:   "in %d Wochen",
			MonthsLater:  "in %d Monaten",
			YearsLater:   "in %d Jahren",
		},

		NumberFormat: NumberFormat{
			DecimalSeparator:  ",",
			ThousandSeparator: ".",
			LargeNumberUnits:  []string{"K", "M", "B", "T"},
		},

		Common: CommonStrings{
			And: "und",
			Or:  "oder",
		},
	})
}