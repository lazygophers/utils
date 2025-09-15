//go:build i18n_it || i18n_all

package human

// 注册意大利语语言支持
func init() {
	RegisterLocale("it", &Locale{
		Language:      "it",
		Region:        "IT",
		ByteUnits:     []string{"B", "KB", "MB", "GB", "TB", "PB"},
		SpeedUnits:    []string{"B/s", "KB/s", "MB/s", "GB/s", "TB/s", "PB/s"},
		BitSpeedUnits: []string{"bps", "Kbps", "Mbps", "Gbps", "Tbps", "Pbps"},

		TimeUnits: TimeUnits{
			Nanosecond:  "ns",
			Microsecond: "μs",
			Millisecond: "ms",
			Second:      "secondo",
			Minute:      "minuto",
			Hour:        "ora",
			Day:         "giorno",
			Week:        "settimana",
			Month:       "mese",
			Year:        "anno",

			Seconds: "secondi",
			Minutes: "minuti",
			Hours:   "ore",
			Days:    "giorni",
			Weeks:   "settimane",
			Months:  "mesi",
			Years:   "anni",
		},

		RelativeTime: RelativeTimeStrings{
			JustNow:    "proprio ora",
			SecondsAgo: "%d secondi fa",
			MinutesAgo: "%d minuti fa",
			HoursAgo:   "%d ore fa",
			DaysAgo:    "%d giorni fa",
			WeeksAgo:   "%d settimane fa",
			MonthsAgo:  "%d mesi fa",
			YearsAgo:   "%d anni fa",

			In:           "tra",
			SecondsLater: "tra %d secondi",
			MinutesLater: "tra %d minuti",
			HoursLater:   "tra %d ore",
			DaysLater:    "tra %d giorni",
			WeeksLater:   "tra %d settimane",
			MonthsLater:  "tra %d mesi",
			YearsLater:   "tra %d anni",
		},

		NumberFormat: NumberFormat{
			DecimalSeparator:  ",",
			ThousandSeparator: ".",
			LargeNumberUnits:  []string{"K", "M", "B", "T"},
		},

		Common: CommonStrings{
			And: "e",
			Or:  "o",
		},
	})
}