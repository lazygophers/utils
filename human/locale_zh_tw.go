//go:build lang_zh_tw || lang_all

package human

// 注册繁体中文语言支持
func init() {
	RegisterLocale("zh-TW", &Locale{
		Language:      "zh",
		Region:        "TW",
		ByteUnits:     []string{"B", "KB", "MB", "GB", "TB", "PB"},
		SpeedUnits:    []string{"B/s", "KB/s", "MB/s", "GB/s", "TB/s", "PB/s"},
		BitSpeedUnits: []string{"bps", "Kbps", "Mbps", "Gbps", "Tbps", "Pbps"},

		TimeUnits: TimeUnits{
			Nanosecond:  "奈秒",
			Microsecond: "微秒",
			Millisecond: "毫秒",
			Second:      "秒",
			Minute:      "分鐘",
			Hour:        "小時",
			Day:         "天",
			Week:        "週",
			Month:       "個月",
			Year:        "年",

			// 繁體中文不需要複數形式
			Seconds: "秒",
			Minutes: "分鐘",
			Hours:   "小時",
			Days:    "天",
			Weeks:   "週",
			Months:  "個月",
			Years:   "年",
		},

		RelativeTime: RelativeTimeStrings{
			JustNow:    "剛剛",
			SecondsAgo: "%d秒前",
			MinutesAgo: "%d分鐘前",
			HoursAgo:   "%d小時前",
			DaysAgo:    "%d天前",
			WeeksAgo:   "%d週前",
			MonthsAgo:  "%d個月前",
			YearsAgo:   "%d年前",

			In:           "",
			SecondsLater: "%d秒後",
			MinutesLater: "%d分鐘後",
			HoursLater:   "%d小時後",
			DaysLater:    "%d天後",
			WeeksLater:   "%d週後",
			MonthsLater:  "%d個月後",
			YearsLater:   "%d年後",
		},

		NumberFormat: NumberFormat{
			DecimalSeparator:  ".",
			ThousandSeparator: ",",
			LargeNumberUnits:  []string{"萬", "億"},
		},

		Common: CommonStrings{
			And: "和",
			Or:  "或",
		},
	})
}
