package human

import (
	"fmt"
	"math"
	"strconv"
	"strings"
	"time"

	"github.com/lazygophers/utils/language"
)

// 全局默认配置变量
var (
	// defaultPrecision 默认精度，表示小数点后保留的位数
	defaultPrecision = 1
	// defaultCompact 是否启用紧凑模式（不带空格）
	defaultCompact = false
	// defaultClockFormat 是否启用时钟格式（HH:MM:SS）
	defaultClockFormat = false
)

// currentLocaleName 返回当前 goroutine 的 locale 名称（小写 base 语言代码）
// 解析顺序：goroutine-local → 全局默认 → "en"
func currentLocaleName() string {
	t := language.Get()
	if t == nil {
		return defaultLocaleName()
	}
	base, _ := t.Tag().Base()
	return base.String()
}

// defaultLocaleName 返回全局默认 locale 名称
func defaultLocaleName() string {
	t := language.Default()
	if t == nil {
		return "en"
	}
	base, _ := t.Tag().Base()
	return base.String()
}

// SetDefaultPrecision 设置默认精度
func SetDefaultPrecision(precision int) {
	defaultPrecision = precision
}

// SetCompact 设置紧凑模式开关
func SetCompact(compact bool) {
	defaultCompact = compact
}

// SetClockFormat 设置时钟格式开关
func SetClockFormat(clock bool) {
	defaultClockFormat = clock
}

// ByteSize 格式化字节大小为人类友好形式
func ByteSize(bytes int64) string {
	return formatByteSize(bytes)
}

// Speed 格式化字节速度为人类友好形式
func Speed(bytesPerSecond int64) string {
	return formatSpeed(bytesPerSecond)
}

// BitSpeed 格式化比特速度为人类友好形式，使用十进制换算
func BitSpeed(bitsPerSecond int64) string {
	return formatBitSpeed(bitsPerSecond)
}

// Duration 格式化时间间隔为人类友好形式
func Duration(d time.Duration) string {
	return formatDuration(d)
}

// RelativeTime 格式化相对时间为人类友好形式
func RelativeTime(t time.Time) string {
	return formatRelativeTime(t)
}

// ClockDuration 格式化时间为时钟格式 (HH:MM:SS / M:SS)
func ClockDuration(d time.Duration) string {
	return formatClockTime(d)
}

// 内部格式化函数

// formatByteSize 格式化字节大小，使用二进制换算 (1024)
func formatByteSize(bytes int64) string {
	if bytes == 0 {
		return formatWithUnit(0, 0, "byte")
	}

	absBytes := abs(bytes)
	const unit = 1024

	locale, _ := GetLocaleConfig(currentLocaleName())
	units := locale.ByteUnits

	if absBytes < unit {
		return formatWithUnit(float64(bytes), 0, "byte")
	}

	// 计算合适的单位级别
	exp := int(math.Floor(math.Log(float64(absBytes)) / math.Log(unit)))
	if exp >= len(units) {
		exp = len(units) - 1
	}

	// 计算数值
	value := float64(bytes) / math.Pow(unit, float64(exp))

	return formatWithUnit(value, exp, "byte")
}

// formatSpeed 格式化字节速度，使用二进制换算 (1024)
func formatSpeed(bytesPerSecond int64) string {
	if bytesPerSecond == 0 {
		return formatWithUnit(0, 0, "speed")
	}

	absBytes := abs(bytesPerSecond)
	const unit = 1024

	locale, _ := GetLocaleConfig(currentLocaleName())
	units := locale.SpeedUnits

	if absBytes < unit {
		return formatWithUnit(float64(bytesPerSecond), 0, "speed")
	}

	// 计算合适的单位级别
	exp := int(math.Floor(math.Log(float64(absBytes)) / math.Log(unit)))
	if exp >= len(units) {
		exp = len(units) - 1
	}

	// 计算数值
	value := float64(bytesPerSecond) / math.Pow(unit, float64(exp))

	return formatWithUnit(value, exp, "speed")
}

// formatBitSpeed 格式化比特速度，使用十进制换算 (1000)
func formatBitSpeed(bitsPerSecond int64) string {
	if bitsPerSecond == 0 {
		return formatWithUnit(0, 0, "bitspeed")
	}

	absBits := abs(bitsPerSecond)
	const unit = 1000 // 网络速度通常使用十进制

	locale, _ := GetLocaleConfig(currentLocaleName())
	units := locale.BitSpeedUnits

	if absBits < unit {
		return formatWithUnit(float64(bitsPerSecond), 0, "bitspeed")
	}

	// 计算合适的单位级别
	exp := int(math.Floor(math.Log(float64(absBits)) / math.Log(unit)))
	if exp >= len(units) {
		exp = len(units) - 1
	}

	// 计算数值
	value := float64(bitsPerSecond) / math.Pow(unit, float64(exp))

	return formatWithUnit(value, exp, "bitspeed")
}

// formatWithUnit 格式化数值和单位
func formatWithUnit(value float64, unitIndex int, unitType string) string {
	locale, _ := GetLocaleConfig(currentLocaleName())

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
		precision := defaultPrecision
		if precision < 0 {
			precision = 1
		}
		formattedValue = formatFloat(value, precision)
	}

	// 紧凑模式不加空格
	if defaultCompact {
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
func formatDuration(d time.Duration) string {
	if d == 0 {
		if defaultClockFormat {
			return "0:00"
		}
		locale, _ := GetLocaleConfig(currentLocaleName())
		return "0 " + locale.TimeUnits.Second
	}

	// 时钟格式
	if defaultClockFormat {
		return formatClockTime(d)
	}

	locale, _ := GetLocaleConfig(currentLocaleName())

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
func formatRelativeTime(t time.Time) string {
	now := time.Now()
	diff := now.Sub(t)

	locale, _ := GetLocaleConfig(currentLocaleName())

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
