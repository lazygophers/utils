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

// TimeFormats 时间格式化模板。值是传给 time.Time.Format 的 layout
// 字符串（基于 Go 参考时间 2006-01-02 15:04:05）。
type TimeFormats struct {
	Date       string // 2006-01-02 / 2006年1月2日
	Time       string // 15:04:05
	DateTime   string // 2006-01-02 15:04:05 / 2006年1月2日 15:04:05
	Year       string // 2006 / 2006年
	YearMonth  string // 2006-01 / 2006年1月
	MonthDay   string // 01-02 / 1月2日
	Short      string // 最短日期（如 1/2/06 / 06/1/2）
	Long       string // 最长含星期（如 Monday, January 2, 2006）
	Weekday    string // 星期完整（Monday / 星期一）
	WeekdayMin string // 星期缩写（Mon / 周一）
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
