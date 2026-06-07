//go:build lang_ja || lang_all

package human

import xlanguage "golang.org/x/text/language"

// 注册日文语言支持
func init() {
	RegisterLocale(xlanguage.Japanese, &Locale{
		Language:      "ja",
		Region:        "JP",
		ByteUnits:     []string{"B", "KB", "MB", "GB", "TB", "PB"},
		SpeedUnits:    []string{"B/s", "KB/s", "MB/s", "GB/s", "TB/s", "PB/s"},
		BitSpeedUnits: []string{"bps", "Kbps", "Mbps", "Gbps", "Tbps", "Pbps"},

		TimeUnits: TimeUnits{
			Nanosecond:  "ナノ秒",
			Microsecond: "マイクロ秒",
			Millisecond: "ミリ秒",
			Second:      "秒",
			Minute:      "分",
			Hour:        "時間",
			Day:         "日",
			Week:        "週間",
			Month:       "ヶ月",
			Year:        "年",

			// 日文不需要复数形式
			Seconds: "秒",
			Minutes: "分",
			Hours:   "時間",
			Days:    "日",
			Weeks:   "週間",
			Months:  "ヶ月",
			Years:   "年",
		},

				TimeFormats: TimeFormats{
			Date:       "2006年1月2日",
			Time:       "15:04:05",
			DateTime:   "2006年1月2日 15:04:05",
			Year:       "2006年",
			YearMonth:  "2006年1月",
			MonthDay:   "1月2日",
			Short:      "06/1/2",
			Long:       "2006年1月2日 Monday",
			Weekday:    "Monday",
			WeekdayMin: "Mon",
		},

		RelativeTime: RelativeTimeStrings{
			JustNow:    "たった今",
			SecondsAgo: "%d秒前",
			MinutesAgo: "%d分前",
			HoursAgo:   "%d時間前",
			DaysAgo:    "%d日前",
			WeeksAgo:   "%d週間前",
			MonthsAgo:  "%dヶ月前",
			YearsAgo:   "%d年前",

			In:           "",
			SecondsLater: "%d秒後",
			MinutesLater: "%d分後",
			HoursLater:   "%d時間後",
			DaysLater:    "%d日後",
			WeeksLater:   "%d週間後",
			MonthsLater:  "%dヶ月後",
			YearsLater:   "%d年後",
		},

		NumberFormat: NumberFormat{
			DecimalSeparator:  ".",
			ThousandSeparator: ",",
			LargeNumberUnits:  []string{"千", "万", "十万", "百万", "千万"},
		},

		Common: CommonStrings{
			And: "と",
			Or:  "または",
		},
	})
}
