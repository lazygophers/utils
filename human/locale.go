package human

import (
	"sync"

	xlanguage "golang.org/x/text/language"
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

var (
	locales = make(map[xlanguage.Tag]*Locale)
	mu      sync.RWMutex
)

// RegisterLocale 注册语言地区配置。tag 使用 stdlib
// golang.org/x/text/language.Tag。
func RegisterLocale(tag xlanguage.Tag, locale *Locale) {
	mu.Lock()
	defer mu.Unlock()
	locales[tag] = locale
}

// GetLocaleConfig 获取语言地区配置。匹配顺序：完整 tag → base 语言（zh-CN → zh）
// → English。仅当 English 也未注册时返回 false。
func GetLocaleConfig(tag xlanguage.Tag) (*Locale, bool) {
	mu.RLock()
	defer mu.RUnlock()

	if locale, ok := locales[tag]; ok {
		return locale, true
	}

	base, _ := tag.Base()
	baseTag := xlanguage.Make(base.String())
	if locale, ok := locales[baseTag]; ok {
		return locale, true
	}

	if locale, ok := locales[xlanguage.English]; ok {
		return locale, true
	}

	return nil, false
}

// getTimeUnit 获取时间单位的正确形式
func getTimeUnit(locale *Locale, unit string, count int64) string {
	if locale == nil {
		locale, _ = GetLocaleConfig(xlanguage.English)
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
