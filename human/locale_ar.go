//go:build i18n_ar || i18n_all

package human

// 注册阿拉伯语语言支持
func init() {
	RegisterLocale("ar", &Locale{
		Language:      "ar",
		Region:        "SA",
		ByteUnits:     []string{"ب", "كب", "مب", "جب", "تب", "بب"},
		SpeedUnits:    []string{"ب/ث", "كب/ث", "مب/ث", "جب/ث", "تب/ث", "بب/ث"},
		BitSpeedUnits: []string{"بت/ث", "كبت/ث", "مبت/ث", "جبت/ث", "تبت/ث", "ببت/ث"},

		TimeUnits: TimeUnits{
			Nanosecond:  "نث",
			Microsecond: "مث",
			Millisecond: "مث",
			Second:      "ثانية",
			Minute:      "دقيقة",
			Hour:        "ساعة",
			Day:         "يوم",
			Week:        "أسبوع",
			Month:       "شهر",
			Year:        "سنة",

			// 阿拉伯语复数形式
			Seconds: "ثوان",
			Minutes: "دقائق",
			Hours:   "ساعات",
			Days:    "أيام",
			Weeks:   "أسابيع",
			Months:  "أشهر",
			Years:   "سنوات",
		},

		RelativeTime: RelativeTimeStrings{
			JustNow:    "الآن",
			SecondsAgo: "منذ %d ثوان",
			MinutesAgo: "منذ %d دقائق",
			HoursAgo:   "منذ %d ساعات",
			DaysAgo:    "منذ %d أيام",
			WeeksAgo:   "منذ %d أسابيع",
			MonthsAgo:  "منذ %d أشهر",
			YearsAgo:   "منذ %d سنوات",

			In:           "خلال",
			SecondsLater: "خلال %d ثوان",
			MinutesLater: "خلال %d دقائق",
			HoursLater:   "خلال %d ساعات",
			DaysLater:    "خلال %d أيام",
			WeeksLater:   "خلال %d أسابيع",
			MonthsLater:  "خلال %d أشهر",
			YearsLater:   "خلال %d سنوات",
		},

		NumberFormat: NumberFormat{
			DecimalSeparator:  ".",
			ThousandSeparator: ",",
			LargeNumberUnits:  []string{"ألف", "مليون", "مليار", "تريليون"},
		},

		Common: CommonStrings{
			And: "و",
			Or:  "أو",
		},
	})
}
