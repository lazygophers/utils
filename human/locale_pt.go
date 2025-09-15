//go:build i18n_pt || i18n_all

package human

// 注册葡萄牙语语言支持
func init() {
	RegisterLocale("pt", &Locale{
		Language:      "pt",
		Region:        "PT",
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
			Day:         "dia",
			Week:        "semana",
			Month:       "mês",
			Year:        "ano",

			Seconds: "segundos",
			Minutes: "minutos",
			Hours:   "horas",
			Days:    "dias",
			Weeks:   "semanas",
			Months:  "meses",
			Years:   "anos",
		},

		RelativeTime: RelativeTimeStrings{
			JustNow:    "agora mesmo",
			SecondsAgo: "há %d segundos",
			MinutesAgo: "há %d minutos",
			HoursAgo:   "há %d horas",
			DaysAgo:    "há %d dias",
			WeeksAgo:   "há %d semanas",
			MonthsAgo:  "há %d meses",
			YearsAgo:   "há %d anos",

			In:           "em",
			SecondsLater: "em %d segundos",
			MinutesLater: "em %d minutos",
			HoursLater:   "em %d horas",
			DaysLater:    "em %d dias",
			WeeksLater:   "em %d semanas",
			MonthsLater:  "em %d meses",
			YearsLater:   "em %d anos",
		},

		NumberFormat: NumberFormat{
			DecimalSeparator:  ",",
			ThousandSeparator: ".",
			LargeNumberUnits:  []string{"K", "M", "B", "T"},
		},

		Common: CommonStrings{
			And: "e",
			Or:  "ou",
		},
	})
}