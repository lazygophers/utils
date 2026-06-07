package human

import "time"

// Date formats t as a localized date (e.g. "2026-06-07" / "2026年6月7日").
func Date(t time.Time) string { return t.Format(timeLayout(localeTimeFormatDate)) }

// Time formats t as a localized time-of-day (e.g. "15:04:05" / "15时04分05秒").
func Time(t time.Time) string { return t.Format(timeLayout(localeTimeFormatTime)) }

// DateTime formats t as a localized date + time (e.g. "2026-06-07 15:04:05").
func DateTime(t time.Time) string { return t.Format(timeLayout(localeTimeFormatDateTime)) }

// Year formats t as a localized year (e.g. "2026" / "2026年").
func Year(t time.Time) string { return t.Format(timeLayout(localeTimeFormatYear)) }

// YearMonth formats t as a localized year + month (e.g. "2026-06" / "2026年6月").
func YearMonth(t time.Time) string { return t.Format(timeLayout(localeTimeFormatYearMonth)) }

// MonthDay formats t as a localized month + day (e.g. "06-07" / "6月7日").
func MonthDay(t time.Time) string { return t.Format(timeLayout(localeTimeFormatMonthDay)) }

// DateShort formats t with the locale's short date layout (e.g. "06-06-07" /
// "6/7/26").
func DateShort(t time.Time) string { return t.Format(timeLayout(localeTimeFormatShort)) }

// DateLong formats t with the locale's long layout (e.g. "Sunday, June 7, 2026"
// / "2026年6月7日 星期日").
func DateLong(t time.Time) string { return t.Format(timeLayout(localeTimeFormatLong)) }

// Weekday returns the localized full weekday name for t (e.g. "Sunday" /
// "星期日").
func Weekday(t time.Time) string { return t.Format(timeLayout(localeTimeFormatWeekday)) }

// WeekdayMin returns the localized short weekday name (e.g. "Sun" / "周日").
func WeekdayMin(t time.Time) string { return t.Format(timeLayout(localeTimeFormatWeekdayMin)) }

// timeFormatField enumerates the named entries on Locale.TimeFormats.
type timeFormatField int

const (
	localeTimeFormatDate timeFormatField = iota
	localeTimeFormatTime
	localeTimeFormatDateTime
	localeTimeFormatYear
	localeTimeFormatYearMonth
	localeTimeFormatMonthDay
	localeTimeFormatShort
	localeTimeFormatLong
	localeTimeFormatWeekday
	localeTimeFormatWeekdayMin
)

// fallback layouts when the current locale (and English) leave a field empty.
var defaultTimeLayouts = map[timeFormatField]string{
	localeTimeFormatDate:       "2006-01-02",
	localeTimeFormatTime:       "15:04:05",
	localeTimeFormatDateTime:   "2006-01-02 15:04:05",
	localeTimeFormatYear:       "2006",
	localeTimeFormatYearMonth:  "2006-01",
	localeTimeFormatMonthDay:   "01-02",
	localeTimeFormatShort:      "1/2/06",
	localeTimeFormatLong:       "Monday, January 2, 2006",
	localeTimeFormatWeekday:    "Monday",
	localeTimeFormatWeekdayMin: "Mon",
}

// timeLayout resolves a layout string for the current goroutine's locale.
// Falls back from current → English → package default.
func timeLayout(field timeFormatField) string {
	if locale, _ := GetLocaleConfig(currentTag()); locale != nil {
		if v := pickTimeLayout(locale.TimeFormats, field); v != "" {
			return v
		}
	}
	return defaultTimeLayouts[field]
}

func pickTimeLayout(f TimeFormats, field timeFormatField) string {
	switch field {
	case localeTimeFormatDate:
		return f.Date
	case localeTimeFormatTime:
		return f.Time
	case localeTimeFormatDateTime:
		return f.DateTime
	case localeTimeFormatYear:
		return f.Year
	case localeTimeFormatYearMonth:
		return f.YearMonth
	case localeTimeFormatMonthDay:
		return f.MonthDay
	case localeTimeFormatShort:
		return f.Short
	case localeTimeFormatLong:
		return f.Long
	case localeTimeFormatWeekday:
		return f.Weekday
	case localeTimeFormatWeekdayMin:
		return f.WeekdayMin
	}
	return ""
}
