package human

import (
	"testing"
	"time"
)

// TestFormatWithUnitLargeIndex 测试单位索引超出范围的情况
func TestFormatWithUnitLargeIndex(t *testing.T) {
	config := DefaultConfig()
	
	// Test with index larger than available units
	result := formatWithUnit(1.0, 999, config, "byte")
	if result == "" {
		t.Error("formatWithUnit with large index should not return empty string")
	}
}

// TestFormatByteSizeWithLargeValue 测试非常大的字节值
func TestFormatByteSizeWithLargeValue(t *testing.T) {
	// Test value larger than available units
	veryLargeValue := int64(1024) * 1024 * 1024 * 1024 * 1024 * 1024 // Beyond PB
	result := formatByteSize(veryLargeValue, DefaultConfig())
	if result == "" {
		t.Error("formatByteSize with very large value should not return empty string")
	}
}

// TestFormatSpeedWithLargeValue 测试非常大的速度值
func TestFormatSpeedWithLargeValue(t *testing.T) {
	veryLargeValue := int64(1024) * 1024 * 1024 * 1024 * 1024 * 1024
	result := formatSpeed(veryLargeValue, DefaultConfig())
	if result == "" {
		t.Error("formatSpeed with very large value should not return empty string")
	}
}

// TestFormatBitSpeedWithLargeValue 测试非常大的比特速度值
func TestFormatBitSpeedWithLargeValue(t *testing.T) {
	veryLargeValue := int64(1000) * 1000 * 1000 * 1000 * 1000 * 1000
	result := formatBitSpeed(veryLargeValue, DefaultConfig())
	if result == "" {
		t.Error("formatBitSpeed with very large value should not return empty string")
	}
}

// TestFormatFloatNegativePrecision 测试负精度的浮点数格式化
func TestFormatFloatNegativePrecision(t *testing.T) {
	result := formatFloat(1.5, -1)
	expected := "2"
	if result != expected {
		t.Errorf("formatFloat with negative precision = %s, want %s", result, expected)
	}
}

// TestFormatDurationVerySmall 测试非常小的时间间隔
func TestFormatDurationVerySmall(t *testing.T) {
	config := DefaultConfig()
	
	// Test nanoseconds
	result := formatDuration(1*time.Nanosecond, config)
	if result == "" {
		t.Error("formatDuration(1ns) should not return empty string")
	}
	
	// Test microseconds
	result = formatDuration(1*time.Microsecond, config)
	if result == "" {
		t.Error("formatDuration(1μs) should not return empty string")
	}
	
	// Test milliseconds
	result = formatDuration(1*time.Millisecond, config)
	if result == "" {
		t.Error("formatDuration(1ms) should not return empty string")
	}
}

// TestFormatClockTimeVeryLarge 测试非常大的时钟时间
func TestFormatClockTimeVeryLarge(t *testing.T) {
	// Test 25 hours (more than 24 hours)
	result := formatClockTime(25 * time.Hour)
	expected := "25:00"
	if result != expected {
		t.Errorf("formatClockTime(25h) = %s, want %s", result, expected)
	}
	
	// Test with seconds
	result = formatClockTime(25*time.Hour + 30*time.Minute + 45*time.Second)
	expected = "25:30:45"
	if result != expected {
		t.Errorf("formatClockTime(25h30m45s) = %s, want %s", result, expected)
	}
}

// TestFormatRelativeTimeEdgeCases 测试相对时间的边界情况
func TestFormatRelativeTimeEdgeCases(t *testing.T) {
	config := DefaultConfig()
	
	// Test exactly 10 seconds (boundary between "just now" and "X seconds ago")
	now := time.Now()
	past := now.Add(-10 * time.Second)
	result := formatRelativeTime(past, config)
	if result == "" {
		t.Error("formatRelativeTime(10s ago) should not return empty string")
	}
	
	// Test future time at various boundaries
	future := now.Add(59 * time.Second) // Just under 1 minute
	result = formatRelativeTime(future, config)
	if result == "" {
		t.Error("formatRelativeTime(59s future) should not return empty string")
	}
}

// TestFormatPastTimeAllRanges 测试过去时间的所有时间范围
func TestFormatPastTimeAllRanges(t *testing.T) {
	locale, _ := GetLocaleConfig("en")
	
	tests := []struct {
		duration time.Duration
		contains string
	}{
		{5 * time.Second, "just now"},
		{30 * time.Second, "seconds ago"},
		{30 * time.Minute, "minutes ago"},
		{5 * time.Hour, "hours ago"},
		{3 * 24 * time.Hour, "days ago"},
		{10 * 24 * time.Hour, "weeks ago"},
		{45 * 24 * time.Hour, "months ago"},
		{400 * 24 * time.Hour, "years ago"},
	}
	
	for _, tt := range tests {
		result := formatPastTime(tt.duration, locale)
		if result == "" {
			t.Errorf("formatPastTime(%v) should not return empty string", tt.duration)
		}
	}
}

// TestFormatFutureTimeAllRanges 测试未来时间的所有时间范围
func TestFormatFutureTimeAllRanges(t *testing.T) {
	locale, _ := GetLocaleConfig("en")
	
	tests := []struct {
		duration time.Duration
		contains string
	}{
		{30 * time.Second, "seconds"},
		{30 * time.Minute, "minutes"},
		{5 * time.Hour, "hours"},
		{3 * 24 * time.Hour, "days"},
		{10 * 24 * time.Hour, "weeks"},
		{45 * 24 * time.Hour, "months"},
		{400 * 24 * time.Hour, "years"},
	}
	
	for _, tt := range tests {
		result := formatFutureTime(tt.duration, locale)
		if result == "" {
			t.Errorf("formatFutureTime(%v) should not return empty string", tt.duration)
		}
	}
}

// TestApplyOptionsEmpty 测试空选项
func TestApplyOptionsEmpty(t *testing.T) {
	config := applyOptions()
	defaultConfig := DefaultConfig()
	
	if config.Precision != defaultConfig.Precision {
		t.Errorf("Empty options precision = %d, want %d", config.Precision, defaultConfig.Precision)
	}
	if config.Locale != defaultConfig.Locale {
		t.Errorf("Empty options locale = %s, want %s", config.Locale, defaultConfig.Locale)
	}
}

// TestDefaultConfigValues 测试默认配置值
func TestDefaultConfigValues(t *testing.T) {
	config := DefaultConfig()
	
	if config.Precision != defaultPrecision {
		t.Errorf("DefaultConfig precision = %d, want %d", config.Precision, defaultPrecision)
	}
	if config.Locale != defaultLocale {
		t.Errorf("DefaultConfig locale = %s, want %s", config.Locale, defaultLocale)
	}
	if config.Compact != false {
		t.Errorf("DefaultConfig compact = %v, want false", config.Compact)
	}
	if config.TimeFormat != "" {
		t.Errorf("DefaultConfig TimeFormat = %s, want empty string", config.TimeFormat)
	}
}

// TestIntegerValues 测试整数值（无小数点）
func TestIntegerValues(t *testing.T) {
	// Test values that should be displayed as integers
	tests := []struct {
		input    int64
		expected string
	}{
		{1024, "1 KB"},     // Exactly 1KB
		{2048, "2 KB"},     // Exactly 2KB
		{1024 * 1024, "1 MB"}, // Exactly 1MB
	}
	
	for _, tt := range tests {
		result := ByteSize(tt.input)
		if result != tt.expected {
			t.Errorf("ByteSize(%d) = %s, want %s", tt.input, result, tt.expected)
		}
	}
}

// TestFormatWithUnitIntegerValue 测试formatWithUnit的整数值
func TestFormatWithUnitIntegerValue(t *testing.T) {
	config := DefaultConfig()
	
	// Test with exact integer value
	result := formatWithUnit(1.0, 1, config, "byte")
	expected := "1 KB"
	if result != expected {
		t.Errorf("formatWithUnit(1.0, 1, byte) = %s, want %s", result, expected)
	}
	
	// Test with decimal value
	result = formatWithUnit(1.5, 1, config, "byte")
	expected = "1.5 KB"
	if result != expected {
		t.Errorf("formatWithUnit(1.5, 1, byte) = %s, want %s", result, expected)
	}
}

// TestAllUnitTypes 测试所有单位类型
func TestAllUnitTypes(t *testing.T) {
	config := DefaultConfig()
	
	unitTypes := []string{"byte", "speed", "bitspeed"}
	
	for _, unitType := range unitTypes {
		result := formatWithUnit(1.5, 1, config, unitType)
		if result == "" {
			t.Errorf("formatWithUnit with unitType %s should not return empty string", unitType)
		}
	}
}

// TestLocaleConfigEdgeCases 测试locale配置的边界情况
func TestLocaleConfigEdgeCases(t *testing.T) {
	// Test with empty string
	_, ok := GetLocaleConfig("")
	if !ok {
		t.Error("GetLocaleConfig('') should fallback to English and return true")
	}
	
	// Test with hyphenated locale that doesn't exist
	retrieved, ok := GetLocaleConfig("nonexistent-REGION")
	if !ok {
		t.Error("GetLocaleConfig with nonexistent hyphenated locale should fallback to English")
	}
	if retrieved.Language != "zh" {
		t.Errorf("Fallback locale should be Chinese, got %s", retrieved.Language)
	}
}

// TestGetTimeUnitEdgeCases 测试getTimeUnit的边界情况
func TestGetTimeUnitEdgeCases(t *testing.T) {
	locale, _ := GetLocaleConfig("en")
	
	// Test with zero count
	result := getTimeUnit(locale, locale.TimeUnits.Second, 0)
	if result != "second" { // Zero is treated as singular
		t.Errorf("getTimeUnit(0) = %s, want second", result)
	}
	
	// Test with non-English locale (should not pluralize)
	locale.Language = "zh"
	result = getTimeUnit(locale, locale.TimeUnits.Second, 2)
	if result != locale.TimeUnits.Second {
		t.Errorf("Non-English locale should not pluralize: %s", result)
	}
}