package human

import (
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

// GetLocaleConfig 获取语言地区配置。匹配顺序：完整 name → 去地区后的 base
// → "en"。返回 false 仅在 "en" 也未注册时发生。
func GetLocaleConfig(name string) (*Locale, bool) {
	mu.RLock()
	defer mu.RUnlock()

	if locale, ok := locales[name]; ok {
		return locale, true
	}

	if i := strings.IndexByte(name, '-'); i > 0 {
		if locale, ok := locales[name[:i]]; ok {
			return locale, true
		}
	}

	if locale, ok := locales["en"]; ok {
		return locale, true
	}

	return nil, false
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
