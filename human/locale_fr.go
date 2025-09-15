//go:build i18n_fr || i18n_all

package human

// 注册法语语言支持
func init() {
	RegisterLocale("fr", &Locale{
		Language:      "fr",
		Region:        "FR",
		ByteUnits:     []string{"o", "Ko", "Mo", "Go", "To", "Po"},
		SpeedUnits:    []string{"o/s", "Ko/s", "Mo/s", "Go/s", "To/s", "Po/s"},
		BitSpeedUnits: []string{"bps", "Kbps", "Mbps", "Gbps", "Tbps", "Pbps"},

		TimeUnits: TimeUnits{
			Nanosecond:  "ns",
			Microsecond: "μs",
			Millisecond: "ms",
			Second:      "seconde",
			Minute:      "minute",
			Hour:        "heure",
			Day:         "jour",
			Week:        "semaine",
			Month:       "mois",
			Year:        "année",

			// 法语复数形式
			Seconds: "secondes",
			Minutes: "minutes",
			Hours:   "heures",
			Days:    "jours",
			Weeks:   "semaines",
			Months:  "mois", // "mois" 在法语中单复数相同
			Years:   "années",
		},

		RelativeTime: RelativeTimeStrings{
			JustNow:    "à l'instant",
			SecondsAgo: "il y a %d secondes",
			MinutesAgo: "il y a %d minutes",
			HoursAgo:   "il y a %d heures",
			DaysAgo:    "il y a %d jours",
			WeeksAgo:   "il y a %d semaines",
			MonthsAgo:  "il y a %d mois",
			YearsAgo:   "il y a %d années",

			In:           "dans",
			SecondsLater: "dans %d secondes",
			MinutesLater: "dans %d minutes",
			HoursLater:   "dans %d heures",
			DaysLater:    "dans %d jours",
			WeeksLater:   "dans %d semaines",
			MonthsLater:  "dans %d mois",
			YearsLater:   "dans %d années",
		},

		NumberFormat: NumberFormat{
			DecimalSeparator:  ",",
			ThousandSeparator: " ",
			LargeNumberUnits:  []string{"mille", "million", "milliard", "billion"},
		},

		Common: CommonStrings{
			And: "et",
			Or:  "ou",
		},
	})
}
