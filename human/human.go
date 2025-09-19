package human

import (
	"fmt"
	"math"
	"strconv"
	"strings"
	"time"
)

// 全局默认配置变量
var (
	// defaultLocale 默认语言地区代码，用于格式化输出
	defaultLocale    = "en"
	// defaultPrecision 默认精度，表示小数点后保留的位数
	defaultPrecision = 1
)

// SetLocale 设置默认语言地区
func SetLocale(locale string) {
	defaultLocale = locale
}

// GetLocale 获取当前语言地区
func GetLocale() string {
	return defaultLocale
}

// SetDefaultPrecision 设置默认精度
func SetDefaultPrecision(precision int) {
	defaultPrecision = precision
}

// ByteSize 格式化字节大小为人类友好形式
func ByteSize(bytes int64, opts ...Option) string {
	config := applyOptions(opts...)
	return formatByteSize(bytes, config)
}

// Speed 格式化字节速度为人类友好形式
func Speed(bytesPerSecond int64, opts ...Option) string {
	config := applyOptions(opts...)
	return formatSpeed(bytesPerSecond, config)
}

// BitSpeed 格式化比特速度为人类友好形式，使用十进制换算
func BitSpeed(bitsPerSecond int64, opts ...Option) string {
	config := applyOptions(opts...)
	return formatBitSpeed(bitsPerSecond, config)
}

// Duration 格式化时间间隔为人类友好形式
func Duration(d time.Duration, opts ...Option) string {
	config := applyOptions(opts...)
	return formatDuration(d, config)
}

// RelativeTime 格式化相对时间为人类友好形式
func RelativeTime(t time.Time, opts ...Option) string {
	config := applyOptions(opts...)
	return formatRelativeTime(t, config)
}

// configToOptions 配置结构转换，用于向后兼容
func configToOptions(config Config) Options {
	return Options{
		Precision:  config.Precision,
		Locale:     config.Locale,
		Compact:    config.Compact,
		TimeFormat: config.TimeFormat,
	}
}

// Options 格式化选项，向后兼容用
type Options struct {
	Precision  int    // 精度
	Locale     string // 语言地区
	Unit       string // 强制单位
	Compact    bool   // 紧凑模式
	TimeFormat string // 时间格式
}

// DefaultOptions 返回默认选项
func DefaultOptions() Options {
	return Options{
		Precision: 1,
		Locale:    defaultLocale,
		Compact:   false,
	}
}

// Formatter 格式化器接口
type Formatter interface {
	Format(value interface{}) string
	FormatWithOptions(value interface{}, options Options) string
}

// 兼容性函数

// BitSpeedWithOptions 格式化比特速度，兼容旧版本
func BitSpeedWithOptions(bitsPerSecond int64, opts Options) string {
	config := Config{
		Precision:  opts.Precision,
		Locale:     opts.Locale,
		Compact:    opts.Compact,
		TimeFormat: opts.TimeFormat,
	}
	return formatBitSpeed(bitsPerSecond, config)
}

// ClockDuration 格式化时间为时钟格式
func ClockDuration(d time.Duration) string {
	return Duration(d, WithClockFormat())
}

// DurationWithOptions 格式化时间，兼容旧版本
func DurationWithOptions(d time.Duration, opts Options) string {
	config := Config{
		Precision:  opts.Precision,
		Locale:     opts.Locale,
		Compact:    opts.Compact,
		TimeFormat: opts.TimeFormat,
	}
	return formatDuration(d, config)
}

// ByteSizeWithOptions 格式化字节大小，兼容旧版本
func ByteSizeWithOptions(bytes int64, opts Options) string {
	config := Config{
		Precision: opts.Precision,
		Locale:    opts.Locale,
		Compact:   opts.Compact,
	}
	return formatByteSize(bytes, config)
}

// 内部格式化函数

// formatByteSize 格式化字节大小，使用二进制换算 (1024)
func formatByteSize(bytes int64, config Config) string {
	if bytes == 0 {
		return formatWithUnit(0, 0, config, "byte")
	}
	
	absBytes := abs(bytes)
	const unit = 1024
	
	locale, _ := GetLocaleConfig(config.Locale)
	units := locale.ByteUnits
	
	if absBytes < unit {
		return formatWithUnit(float64(bytes), 0, config, "byte")
	}
	
	// 计算合适的单位级别
	exp := int(math.Floor(math.Log(float64(absBytes)) / math.Log(unit)))
	if exp >= len(units) {
		exp = len(units) - 1
	}
	
	// 计算数值
	value := float64(bytes) / math.Pow(unit, float64(exp))
	
	return formatWithUnit(value, exp, config, "byte")
}

// formatSpeed 格式化字节速度，使用二进制换算 (1024)
func formatSpeed(bytesPerSecond int64, config Config) string {
	if bytesPerSecond == 0 {
		return formatWithUnit(0, 0, config, "speed")
	}
	
	absBytes := abs(bytesPerSecond)
	const unit = 1024
	
	locale, _ := GetLocaleConfig(config.Locale)
	units := locale.SpeedUnits
	
	if absBytes < unit {
		return formatWithUnit(float64(bytesPerSecond), 0, config, "speed")
	}
	
	// 计算合适的单位级别
	exp := int(math.Floor(math.Log(float64(absBytes)) / math.Log(unit)))
	if exp >= len(units) {
		exp = len(units) - 1
	}
	
	// 计算数值
	value := float64(bytesPerSecond) / math.Pow(unit, float64(exp))
	
	return formatWithUnit(value, exp, config, "speed")
}

// formatBitSpeed 格式化比特速度，使用十进制换算 (1000)
func formatBitSpeed(bitsPerSecond int64, config Config) string {
	if bitsPerSecond == 0 {
		return formatWithUnit(0, 0, config, "bitspeed")
	}
	
	absBits := abs(bitsPerSecond)
	const unit = 1000 // 网络速度通常使用十进制
	
	locale, _ := GetLocaleConfig(config.Locale)
	units := locale.BitSpeedUnits
	
	if absBits < unit {
		return formatWithUnit(float64(bitsPerSecond), 0, config, "bitspeed")
	}
	
	// 计算合适的单位级别
	exp := int(math.Floor(math.Log(float64(absBits)) / math.Log(unit)))
	if exp >= len(units) {
		exp = len(units) - 1
	}
	
	// 计算数值
	value := float64(bitsPerSecond) / math.Pow(unit, float64(exp))
	
	return formatWithUnit(value, exp, config, "bitspeed")
}

// formatWithUnit 格式化数值和单位
func formatWithUnit(value float64, unitIndex int, config Config, unitType string) string {
	locale, _ := GetLocaleConfig(config.Locale)
	
	var units []string
	switch unitType {
	case "byte":
		units = locale.ByteUnits
	case "speed":
		units = locale.SpeedUnits
	case "bitspeed":
		units = locale.BitSpeedUnits
	default:
		return "-"
	}
	
	if unitIndex >= len(units) {
		unitIndex = len(units) - 1
	}
	
	unit := units[unitIndex]
	
	// 格式化数值
	var formattedValue string
	if value == math.Trunc(value) {
		formattedValue = strconv.FormatFloat(value, 'f', 0, 64)
	} else {
		precision := config.Precision
		if precision < 0 {
			precision = defaultPrecision
		}
		formattedValue = formatFloat(value, precision)
	}
	
	// 紧凑模式不加空格
	if config.Compact {
		return formattedValue + unit
	}
	
	return formattedValue + " " + unit
}

// 工具函数

// formatFloat 格式化浮点数，去除尾随零
func formatFloat(f float64, precision int) string {
	if precision < 0 {
		precision = 0
	}

	str := strconv.FormatFloat(f, 'f', precision, 64)

	// 去除尾随零和多余的小数点
	if strings.Contains(str, ".") {
		str = strings.TrimRight(str, "0")
		str = strings.TrimRight(str, ".")
	}

	return str
}

// formatDuration 格式化时间间隔
func formatDuration(d time.Duration, config Config) string {
	if d == 0 {
		if config.TimeFormat == "clock" {
			return "0:00"
		}
		locale, _ := GetLocaleConfig(config.Locale)
		return "0 " + locale.TimeUnits.Second
	}
	
	// 时钟格式
	if config.TimeFormat == "clock" {
		return formatClockTime(d)
	}
	
	locale, _ := GetLocaleConfig(config.Locale)
	
	// 处理负数
	negative := d < 0
	if negative {
		d = -d
	}
	
	var parts []string
	
	// 计算各个时间单位
	days := d / (24 * time.Hour)
	d %= 24 * time.Hour
	hours := d / time.Hour
	d %= time.Hour
	minutes := d / time.Minute
	d %= time.Minute
	seconds := d / time.Second
	
	// 根据最大的时间单位来决定显示精度
	if days > 0 {
		parts = append(parts, fmt.Sprintf("%d %s", days, getTimeUnit(locale, locale.TimeUnits.Day, int64(days))))
		if hours > 0 {
			parts = append(parts, fmt.Sprintf("%d %s", hours, getTimeUnit(locale, locale.TimeUnits.Hour, int64(hours))))
		}
	} else if hours > 0 {
		parts = append(parts, fmt.Sprintf("%d %s", hours, getTimeUnit(locale, locale.TimeUnits.Hour, int64(hours))))
		if minutes > 0 {
			parts = append(parts, fmt.Sprintf("%d %s", minutes, getTimeUnit(locale, locale.TimeUnits.Minute, int64(minutes))))
		}
	} else if minutes > 0 {
		parts = append(parts, fmt.Sprintf("%d %s", minutes, getTimeUnit(locale, locale.TimeUnits.Minute, int64(minutes))))
		if seconds > 0 {
			parts = append(parts, fmt.Sprintf("%d %s", seconds, getTimeUnit(locale, locale.TimeUnits.Second, int64(seconds))))
		}
	} else {
		parts = append(parts, fmt.Sprintf("%d %s", seconds, getTimeUnit(locale, locale.TimeUnits.Second, int64(seconds))))
	}
	
	result := strings.Join(parts, " ")
	if negative {
		result = "-" + result
	}
	
	return result
}

// formatClockTime 格式化为时钟格式
func formatClockTime(d time.Duration) string {
	// 处理负数
	negative := d < 0
	if negative {
		d = -d
	}
	
	totalSeconds := int64(d.Seconds())
	hours := totalSeconds / 3600
	minutes := (totalSeconds % 3600) / 60
	seconds := totalSeconds % 60
	
	var result string
	if hours > 0 {
		// 有小时：H:MM:SS 或 H:MM (如果秒数为0)
		if seconds > 0 {
			result = fmt.Sprintf("%d:%02d:%02d", hours, minutes, seconds)
		} else {
			result = fmt.Sprintf("%d:%02d", hours, minutes)
		}
	} else if minutes > 0 {
		// 只有分钟：M:SS
		result = fmt.Sprintf("%d:%02d", minutes, seconds)
	} else {
		// 只有秒：0:SS
		result = fmt.Sprintf("0:%02d", seconds)
	}
	
	if negative {
		result = "-" + result
	}
	
	return result
}

// formatRelativeTime 格式化相对时间
func formatRelativeTime(t time.Time, config Config) string {
	now := time.Now()
	diff := now.Sub(t)
	
	locale, _ := GetLocaleConfig(config.Locale)
	
	// 处理未来时间
	if diff < 0 {
		diff = -diff
		return formatFutureTime(diff, locale)
	}
	
	// 处理过去时间
	return formatPastTime(diff, locale)
}

// formatPastTime 格式化过去时间
func formatPastTime(diff time.Duration, locale *Locale) string {
	if diff < 10*time.Second {
		return locale.RelativeTime.JustNow
	}
	
	if diff < time.Minute {
		seconds := int(diff.Seconds())
		return fmt.Sprintf(locale.RelativeTime.SecondsAgo, seconds)
	}
	
	if diff < time.Hour {
		minutes := int(diff.Minutes())
		return fmt.Sprintf(locale.RelativeTime.MinutesAgo, minutes)
	}
	
	if diff < 24*time.Hour {
		hours := int(diff.Hours())
		return fmt.Sprintf(locale.RelativeTime.HoursAgo, hours)
	}
	
	if diff < 7*24*time.Hour {
		days := int(diff.Hours() / 24)
		return fmt.Sprintf(locale.RelativeTime.DaysAgo, days)
	}
	
	if diff < 30*24*time.Hour {
		weeks := int(diff.Hours() / (7 * 24))
		return fmt.Sprintf(locale.RelativeTime.WeeksAgo, weeks)
	}
	
	if diff < 365*24*time.Hour {
		months := int(diff.Hours() / (30 * 24))
		return fmt.Sprintf(locale.RelativeTime.MonthsAgo, months)
	}
	
	years := int(diff.Hours() / (365 * 24))
	return fmt.Sprintf(locale.RelativeTime.YearsAgo, years)
}

// formatFutureTime 格式化未来时间
func formatFutureTime(diff time.Duration, locale *Locale) string {
	if diff < time.Minute {
		seconds := int(diff.Seconds())
		return fmt.Sprintf(locale.RelativeTime.SecondsLater, seconds)
	}
	
	if diff < time.Hour {
		minutes := int(diff.Minutes())
		return fmt.Sprintf(locale.RelativeTime.MinutesLater, minutes)
	}
	
	if diff < 24*time.Hour {
		hours := int(diff.Hours())
		return fmt.Sprintf(locale.RelativeTime.HoursLater, hours)
	}
	
	if diff < 7*24*time.Hour {
		days := int(diff.Hours() / 24)
		return fmt.Sprintf(locale.RelativeTime.DaysLater, days)
	}
	
	if diff < 30*24*time.Hour {
		weeks := int(diff.Hours() / (7 * 24))
		return fmt.Sprintf(locale.RelativeTime.WeeksLater, weeks)
	}
	
	if diff < 365*24*time.Hour {
		months := int(diff.Hours() / (30 * 24))
		return fmt.Sprintf(locale.RelativeTime.MonthsLater, months)
	}
	
	years := int(diff.Hours() / (365 * 24))
	return fmt.Sprintf(locale.RelativeTime.YearsLater, years)
}

// abs 返回整数的绝对值
func abs(n int64) int64 {
	if n < 0 {
		return -n
	}
	return n
}
