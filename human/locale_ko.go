//go:build i18n_ko || i18n_all

package human

// 注册韩文语言支持
func init() {
	RegisterLocale("ko", &Locale{
		Language:      "ko",
		Region:        "KR",
		ByteUnits:     []string{"B", "KB", "MB", "GB", "TB", "PB"},
		SpeedUnits:    []string{"B/s", "KB/s", "MB/s", "GB/s", "TB/s", "PB/s"},
		BitSpeedUnits: []string{"bps", "Kbps", "Mbps", "Gbps", "Tbps", "Pbps"},

		TimeUnits: TimeUnits{
			Nanosecond:  "나노초",
			Microsecond: "마이크로초",
			Millisecond: "밀리초",
			Second:      "초",
			Minute:      "분",
			Hour:        "시간",
			Day:         "일",
			Week:        "주",
			Month:       "개월",
			Year:        "년",

			// 韩文不需要复数形式
			Seconds: "초",
			Minutes: "분",
			Hours:   "시간",
			Days:    "일",
			Weeks:   "주",
			Months:  "개월",
			Years:   "년",
		},

		RelativeTime: RelativeTimeStrings{
			JustNow:    "방금",
			SecondsAgo: "%d초 전",
			MinutesAgo: "%d분 전",
			HoursAgo:   "%d시간 전",
			DaysAgo:    "%d일 전",
			WeeksAgo:   "%d주 전",
			MonthsAgo:  "%d개월 전",
			YearsAgo:   "%d년 전",

			In:           "",
			SecondsLater: "%d초 후",
			MinutesLater: "%d분 후",
			HoursLater:   "%d시간 후",
			DaysLater:    "%d일 후",
			WeeksLater:   "%d주 후",
			MonthsLater:  "%d개월 후",
			YearsLater:   "%d년 후",
		},

		NumberFormat: NumberFormat{
			DecimalSeparator:  ".",
			ThousandSeparator: ",",
			LargeNumberUnits:  []string{"천", "만", "십만", "백만", "천만"},
		},

		Common: CommonStrings{
			And: "그리고",
			Or:  "또는",
		},
	})
}
