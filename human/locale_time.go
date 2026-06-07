package human

// TimeUnits 时间单位
type TimeUnits struct {
	Nanosecond  string
	Microsecond string
	Millisecond string
	Second      string
	Minute      string
	Hour        string
	Day         string
	Week        string
	Month       string
	Year        string

	// 复数形式
	Seconds string
	Minutes string
	Hours   string
	Days    string
	Weeks   string
	Months  string
	Years   string
}

// RelativeTimeStrings 相对时间表达
type RelativeTimeStrings struct {
	JustNow    string
	SecondsAgo string
	MinutesAgo string
	HoursAgo   string
	DaysAgo    string
	WeeksAgo   string
	MonthsAgo  string
	YearsAgo   string

	In           string
	SecondsLater string
	MinutesLater string
	HoursLater   string
	DaysLater    string
	WeeksLater   string
	MonthsLater  string
	YearsLater   string
}
