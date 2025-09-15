//go:build lang_ru || lang_all

package human

// 注册俄语语言支持
func init() {
	RegisterLocale("ru", &Locale{
		Language:      "ru",
		Region:        "RU",
		ByteUnits:     []string{"Б", "КБ", "МБ", "ГБ", "ТБ", "ПБ"},
		SpeedUnits:    []string{"Б/с", "КБ/с", "МБ/с", "ГБ/с", "ТБ/с", "ПБ/с"},
		BitSpeedUnits: []string{"бит/с", "Кбит/с", "Мбит/с", "Гбит/с", "Тбит/с", "Пбит/с"},

		TimeUnits: TimeUnits{
			Nanosecond:  "нс",
			Microsecond: "мкс",
			Millisecond: "мс",
			Second:      "секунда",
			Minute:      "минута",
			Hour:        "час",
			Day:         "день",
			Week:        "неделя",
			Month:       "месяц",
			Year:        "год",

			// 俄语复数形式（基本复数）
			Seconds: "секунды",
			Minutes: "минуты",
			Hours:   "часа",
			Days:    "дня",
			Weeks:   "недели",
			Months:  "месяца",
			Years:   "года",
		},

		RelativeTime: RelativeTimeStrings{
			JustNow:    "только что",
			SecondsAgo: "%d секунд назад",
			MinutesAgo: "%d минут назад",
			HoursAgo:   "%d часов назад",
			DaysAgo:    "%d дней назад",
			WeeksAgo:   "%d недель назад",
			MonthsAgo:  "%d месяцев назад",
			YearsAgo:   "%d лет назад",

			In:           "через",
			SecondsLater: "через %d секунд",
			MinutesLater: "через %d минут",
			HoursLater:   "через %d часов",
			DaysLater:    "через %d дней",
			WeeksLater:   "через %d недель",
			MonthsLater:  "через %d месяцев",
			YearsLater:   "через %d лет",
		},

		NumberFormat: NumberFormat{
			DecimalSeparator:  ",",
			ThousandSeparator: " ",
			LargeNumberUnits:  []string{"тысяча", "миллион", "миллиард", "триллион"},
		},

		Common: CommonStrings{
			And: "и",
			Or:  "или",
		},
	})
}
