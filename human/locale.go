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
	
	// 字节单位
	ByteUnits     []string
	SpeedUnits    []string
	BitSpeedUnits []string
	
	// 时间单位
	TimeUnits    TimeUnits
	
	// 相对时间表达
	RelativeTime RelativeTimeStrings
	
	// 数字格式
	NumberFormat NumberFormat
	
	// 其他常用词汇
	Common       CommonStrings
}

// TimeUnits 时间单位表达
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
	
	// 复数形式（英文需要）
	Seconds     string
	Minutes     string
	Hours       string
	Days        string
	Weeks       string
	Months      string
	Years       string
}

// RelativeTimeStrings 相对时间表达
type RelativeTimeStrings struct {
	JustNow     string
	SecondsAgo  string // %d 秒前
	MinutesAgo  string // %d 分钟前
	HoursAgo    string // %d 小时前
	DaysAgo     string // %d 天前
	WeeksAgo    string // %d 周前
	MonthsAgo   string // %d 个月前
	YearsAgo    string // %d 年前
	
	In          string // "在" (中文) 或 "in" (英文)
	SecondsLater string // %d 秒后
	MinutesLater string // %d 分钟后
	HoursLater   string // %d 小时后
	DaysLater    string // %d 天后
	WeeksLater   string // %d 周后
	MonthsLater  string // %d 个月后
	YearsLater   string // %d 年后
}

// NumberFormat 数字格式配置
type NumberFormat struct {
	DecimalSeparator  string   // 小数分隔符
	ThousandSeparator string   // 千位分隔符
	LargeNumberUnits  []string // 大数字单位：万、十万、百万...
}

// CommonStrings 常用字符串
type CommonStrings struct {
	And string // "和"、"and"
	Or  string // "或"、"or"
}

// 全局地区管理器
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

// 注册默认英文地区
func init() {
	RegisterLocale("en", &Locale{
		Language:   "en",
		Region:     "US",
		ByteUnits:     []string{"B", "KB", "MB", "GB", "TB", "PB"},
		SpeedUnits:    []string{"B/s", "KB/s", "MB/s", "GB/s", "TB/s", "PB/s"},
		BitSpeedUnits: []string{"bps", "Kbps", "Mbps", "Gbps", "Tbps", "Pbps"},
		
		TimeUnits: TimeUnits{
			Nanosecond:  "ns",
			Microsecond: "μs", 
			Millisecond: "ms",
			Second:      "second",
			Minute:      "minute",
			Hour:        "hour",
			Day:         "day",
			Week:        "week",
			Month:       "month",
			Year:        "year",
			
			Seconds: "seconds",
			Minutes: "minutes",
			Hours:   "hours",
			Days:    "days",
			Weeks:   "weeks",
			Months:  "months",
			Years:   "years",
		},
		
		RelativeTime: RelativeTimeStrings{
			JustNow:      "just now",
			SecondsAgo:   "%d seconds ago",
			MinutesAgo:   "%d minutes ago",
			HoursAgo:     "%d hours ago",
			DaysAgo:      "%d days ago",
			WeeksAgo:     "%d weeks ago",
			MonthsAgo:    "%d months ago",
			YearsAgo:     "%d years ago",
			
			In:           "in",
			SecondsLater: "in %d seconds",
			MinutesLater: "in %d minutes",
			HoursLater:   "in %d hours",
			DaysLater:    "in %d days",
			WeeksLater:   "in %d weeks",
			MonthsLater:  "in %d months",
			YearsLater:   "in %d years",
		},
		
		NumberFormat: NumberFormat{
			DecimalSeparator:  ".",
			ThousandSeparator: ",",
			LargeNumberUnits:  []string{"K", "M", "B", "T"},
		},
		
		Common: CommonStrings{
			And: "and",
			Or:  "or",
		},
	})
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