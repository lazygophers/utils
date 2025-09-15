//go:build i18n_es || i18n_all

package human

// 注册西班牙语语言支持
func init() {
	RegisterLocale("es", &Locale{
		Language:      "es",
		Region:        "ES",
		ByteUnits:     []string{"B", "KB", "MB", "GB", "TB", "PB"},
		SpeedUnits:    []string{"B/s", "KB/s", "MB/s", "GB/s", "TB/s", "PB/s"},
		BitSpeedUnits: []string{"bps", "Kbps", "Mbps", "Gbps", "Tbps", "Pbps"},

		TimeUnits: TimeUnits{
			Nanosecond:  "ns",
			Microsecond: "μs",
			Millisecond: "ms",
			Second:      "segundo",
			Minute:      "minuto",
			Hour:        "hora",
			Day:         "día",
			Week:        "semana",
			Month:       "mes",
			Year:        "año",

			// 西班牙语复数形式
			Seconds: "segundos",
			Minutes: "minutos",
			Hours:   "horas",
			Days:    "días",
			Weeks:   "semanas",
			Months:  "meses",
			Years:   "años",
		},

		RelativeTime: RelativeTimeStrings{
			JustNow:    "ahora mismo",
			SecondsAgo: "hace %d segundos",
			MinutesAgo: "hace %d minutos",
			HoursAgo:   "hace %d horas",
			DaysAgo:    "hace %d días",
			WeeksAgo:   "hace %d semanas",
			MonthsAgo:  "hace %d meses",
			YearsAgo:   "hace %d años",

			In:           "en",
			SecondsLater: "en %d segundos",
			MinutesLater: "en %d minutos",
			HoursLater:   "en %d horas",
			DaysLater:    "en %d días",
			WeeksLater:   "en %d semanas",
			MonthsLater:  "en %d meses",
			YearsLater:   "en %d años",
		},

		NumberFormat: NumberFormat{
			DecimalSeparator:  ",",
			ThousandSeparator: ".",
			LargeNumberUnits:  []string{"mil", "millón", "mil millones", "billón"},
		},

		Common: CommonStrings{
			And: "y",
			Or:  "o",
		},
	})
}
