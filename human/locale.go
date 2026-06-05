package human

import (
	"fmt"
	"strings"
	"sync"
)

// Locale 语言地区配置
type Locale struct {
	Language string
	Region   string

	ByteUnits     []string // 字节单位
	SpeedUnits    []string // 速度单位
	BitSpeedUnits []string // 比特速度单位

	TimeUnits    TimeUnits           // 时间单位
	RelativeTime RelativeTimeStrings // 相对时间表达
	NumberFormat NumberFormat        // 数字格式
	Common       CommonStrings       // 常用词汇
}

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

// NumberFormat 数字格式
type NumberFormat struct {
	DecimalSeparator  string   // 小数分隔符
	ThousandSeparator string   // 千位分隔符
	LargeNumberUnits  []string // 大数字单位
}

// CommonStrings 常用字符串
type CommonStrings struct {
	And string
	Or  string
}

var (
	locales = make(map[string]*Locale)
	mu      sync.RWMutex
)

// RegisterLocale 注册语言地区
func RegisterLocale(name string, locale *Locale) {
	mu.Lock()
	defer mu.Unlock()
	locales[name] = locale
}

// GetLocaleConfig 获取语言地区配置
func GetLocaleConfig(name string) (*Locale, bool) {
	mu.RLock()
	defer mu.RUnlock()

	// 尝试完整匹配
	if locale, ok := locales[name]; ok {
		return locale, true
	}

	// 尝试语言匹配（忽略地区）
	lang := strings.Split(name, "-")[0]
	if locale, ok := locales[lang]; ok {
		return locale, true
	}

	// 默认英文
	if locale, ok := locales["en"]; ok {
		return locale, true
	}

	return nil, false
}


// formatWithLocale 使用地区配置格式化字符串
func formatWithLocale(locale *Locale, format string, args ...interface{}) string {
	if locale == nil {
		locale, _ = GetLocaleConfig("en")
	}

	// 这里可以根据需要进行更复杂的本地化处理
	// 比如处理复数形式、语序调整等

	return fmt.Sprintf(format, args...)
}

// getTimeUnit 获取时间单位的正确形式
func getTimeUnit(locale *Locale, unit string, count int64) string {
	if locale == nil {
		locale, _ = GetLocaleConfig("en")
	}

	// 对于英文，需要处理单复数
	if locale.Language == "en" && count != 1 {
		switch unit {
		case locale.TimeUnits.Second:
			return locale.TimeUnits.Seconds
		case locale.TimeUnits.Minute:
			return locale.TimeUnits.Minutes
		case locale.TimeUnits.Hour:
			return locale.TimeUnits.Hours
		case locale.TimeUnits.Day:
			return locale.TimeUnits.Days
		case locale.TimeUnits.Week:
			return locale.TimeUnits.Weeks
		case locale.TimeUnits.Month:
			return locale.TimeUnits.Months
		case locale.TimeUnits.Year:
			return locale.TimeUnits.Years
		}
	}

	return unit
}
