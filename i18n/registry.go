package i18n

// 注册默认的英语配置
func init() {
	RegisterLocale(English, &Locale{
		Language:     English,
		Region:       "US",
		Name:         "English",
		EnglishName:  "English",
		Messages:     make(map[string]string),
		Formats: &Formats{
			DateFormat:        "2006-01-02",
			TimeFormat:        "15:04:05",
			DateTimeFormat:    "2006-01-02 15:04:05",
			NumberFormat:      "%.2f",
			CurrencyFormat:    "%.2f",
			DecimalSeparator:  ".",
			ThousandSeparator: ",",
			Units: &Units{
				ByteUnits:     []string{"B", "KB", "MB", "GB", "TB", "PB"},
				SpeedUnits:    []string{"B/s", "KB/s", "MB/s", "GB/s", "TB/s", "PB/s"},
				BitSpeedUnits: []string{"bps", "Kbps", "Mbps", "Gbps", "Tbps", "Pbps"},
				TimeUnits: map[string]string{
					"nanosecond":  "ns",
					"microsecond": "μs",
					"millisecond": "ms",
					"second":      "second",
					"minute":      "minute",
					"hour":        "hour",
					"day":         "day",
					"week":        "week",
					"month":       "month",
					"year":        "year",
					"seconds":     "seconds",
					"minutes":     "minutes",
					"hours":       "hours",
					"days":        "days",
					"weeks":       "weeks",
					"months":      "months",
					"years":       "years",
				},
				DistanceUnits: []string{"mm", "cm", "m", "km"},
				WeightUnits:   []string{"g", "kg", "t"},
			},
		},
	})
}