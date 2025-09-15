//go:build lang_zh_cn || lang_all

package human

// 注册中文语言支持
func init() {
	RegisterLocale("zh", &Locale{
		Language:      "zh",
		Region:        "CN",
		ByteUnits:     []string{"B", "KB", "MB", "GB", "TB", "PB"},
		SpeedUnits:    []string{"B/s", "KB/s", "MB/s", "GB/s", "TB/s", "PB/s"},
		BitSpeedUnits: []string{"bps", "Kbps", "Mbps", "Gbps", "Tbps", "Pbps"},

		TimeUnits: TimeUnits{
			Nanosecond:  "纳秒",
			Microsecond: "微秒",
			Millisecond: "毫秒",
			Second:      "秒",
			Minute:      "分钟",
			Hour:        "小时",
			Day:         "天",
			Week:        "周",
			Month:       "个月",
			Year:        "年",

			// 中文不需要复数形式
			Seconds: "秒",
			Minutes: "分钟",
			Hours:   "小时",
			Days:    "天",
			Weeks:   "周",
			Months:  "个月",
			Years:   "年",
		},

		RelativeTime: RelativeTimeStrings{
			JustNow:    "刚刚",
			SecondsAgo: "%d秒前",
			MinutesAgo: "%d分钟前",
			HoursAgo:   "%d小时前",
			DaysAgo:    "%d天前",
			WeeksAgo:   "%d周前",
			MonthsAgo:  "%d个月前",
			YearsAgo:   "%d年前",

			In:           "",
			SecondsLater: "%d秒后",
			MinutesLater: "%d分钟后",
			HoursLater:   "%d小时后",
			DaysLater:    "%d天后",
			WeeksLater:   "%d周后",
			MonthsLater:  "%d个月后",
			YearsLater:   "%d年后",
		},

		NumberFormat: NumberFormat{
			DecimalSeparator:  ".",
			ThousandSeparator: ",",
			LargeNumberUnits:  []string{"万", "亿"},
		},

		Common: CommonStrings{
			And: "和",
			Or:  "或",
		},
	})

	// 同时注册 zh-CN
	RegisterLocale("zh-CN", &Locale{
		Language:      "zh",
		Region:        "CN",
		ByteUnits:     []string{"B", "KB", "MB", "GB", "TB", "PB"},
		SpeedUnits:    []string{"B/s", "KB/s", "MB/s", "GB/s", "TB/s", "PB/s"},
		BitSpeedUnits: []string{"bps", "Kbps", "Mbps", "Gbps", "Tbps", "Pbps"},

		TimeUnits: TimeUnits{
			Nanosecond:  "纳秒",
			Microsecond: "微秒",
			Millisecond: "毫秒",
			Second:      "秒",
			Minute:      "分钟",
			Hour:        "小时",
			Day:         "天",
			Week:        "周",
			Month:       "个月",
			Year:        "年",

			Seconds: "秒",
			Minutes: "分钟",
			Hours:   "小时",
			Days:    "天",
			Weeks:   "周",
			Months:  "个月",
			Years:   "年",
		},

		RelativeTime: RelativeTimeStrings{
			JustNow:    "刚刚",
			SecondsAgo: "%d秒前",
			MinutesAgo: "%d分钟前",
			HoursAgo:   "%d小时前",
			DaysAgo:    "%d天前",
			WeeksAgo:   "%d周前",
			MonthsAgo:  "%d个月前",
			YearsAgo:   "%d年前",

			In:           "",
			SecondsLater: "%d秒后",
			MinutesLater: "%d分钟后",
			HoursLater:   "%d小时后",
			DaysLater:    "%d天后",
			WeeksLater:   "%d周后",
			MonthsLater:  "%d个月后",
			YearsLater:   "%d年后",
		},

		NumberFormat: NumberFormat{
			DecimalSeparator:  ".",
			ThousandSeparator: ",",
			LargeNumberUnits:  []string{"万", "亿"},
		},

		Common: CommonStrings{
			And: "和",
			Or:  "或",
		},
	})
}
