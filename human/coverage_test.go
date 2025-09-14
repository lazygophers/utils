package human

import (
	"testing"
	"time"
)

// TestSetLocaleAndGetLocale 测试全局locale设置
func TestSetLocaleAndGetLocale(t *testing.T) {
	originalLocale := GetLocale()
	defer SetLocale(originalLocale) // 恢复原始设置

	SetLocale("zh-CN")
	if GetLocale() != "zh-CN" {
		t.Errorf("Expected locale 'zh-CN', got %s", GetLocale())
	}

	SetLocale("ja")
	if GetLocale() != "ja" {
		t.Errorf("Expected locale 'ja', got %s", GetLocale())
	}
}

// TestSetDefaultPrecision 测试默认精度设置
func TestSetDefaultPrecision(t *testing.T) {
	// 保存原始精度
	original := defaultPrecision
	defer func() { defaultPrecision = original }()

	SetDefaultPrecision(3)
	if defaultPrecision != 3 {
		t.Errorf("Expected precision 3, got %d", defaultPrecision)
	}
}

// TestDefaultOptions 测试默认选项
func TestDefaultOptions(t *testing.T) {
	opts := DefaultOptions()
	if opts.Precision != 1 {
		t.Errorf("Expected default precision 1, got %d", opts.Precision)
	}
	if opts.Locale != defaultLocale {
		t.Errorf("Expected default locale %s, got %s", defaultLocale, opts.Locale)
	}
	if opts.Compact != false {
		t.Errorf("Expected default compact false, got %v", opts.Compact)
	}
}

// TestFunctionalOptions 测试函数式选项
func TestFunctionalOptions(t *testing.T) {
	// Test individual options
	config := applyOptions(WithPrecision(2))
	if config.Precision != 2 {
		t.Errorf("WithPrecision(2) failed, got %d", config.Precision)
	}

	config = applyOptions(WithLocale("zh-CN"))
	if config.Locale != "zh-CN" {
		t.Errorf("WithLocale('zh-CN') failed, got %s", config.Locale)
	}

	config = applyOptions(WithCompact())
	if config.Compact != true {
		t.Errorf("WithCompact() failed, got %v", config.Compact)
	}

	config = applyOptions(WithClockFormat())
	if config.TimeFormat != "clock" {
		t.Errorf("WithClockFormat() failed, got %s", config.TimeFormat)
	}

	// Test multiple options
	config = applyOptions(WithPrecision(3), WithLocale("ja"), WithCompact(), WithClockFormat())
	if config.Precision != 3 || config.Locale != "ja" || !config.Compact || config.TimeFormat != "clock" {
		t.Errorf("Multiple options failed: %+v", config)
	}
}

// TestNegativeValues 测试负值处理
func TestNegativeValues(t *testing.T) {
	// Test negative byte size
	result := ByteSize(-1024)
	expected := "-1 KB"
	if result != expected {
		t.Errorf("ByteSize(-1024) = %s, want %s", result, expected)
	}

	// Test negative speed
	result = Speed(-2048)
	expected = "-2 KB/s"
	if result != expected {
		t.Errorf("Speed(-2048) = %s, want %s", result, expected)
	}

	// Test negative bit speed
	result = BitSpeed(-8000)
	expected = "-8 Kbps"
	if result != expected {
		t.Errorf("BitSpeed(-8000) = %s, want %s", result, expected)
	}

	// Test negative duration
	result = Duration(-90 * time.Second)
	expected = "-1 minute 30 seconds"
	if result != expected {
		t.Errorf("Duration(-90s) = %s, want %s", result, expected)
	}

	// Test negative duration with clock format
	result = Duration(-90*time.Second, WithClockFormat())
	expected = "-1:30"
	if result != expected {
		t.Errorf("Duration(-90s, WithClockFormat()) = %s, want %s", result, expected)
	}
}

// TestLargeValues 测试大数值处理
func TestLargeValues(t *testing.T) {
	// Test very large byte sizes
	tests := []struct {
		input    int64
		expected string
	}{
		{1024 * 1024 * 1024 * 1024 * 1024, "1 PB"},                        // 1 PB
		{1024 * 1024 * 1024 * 1024 * 1024 * 1024, "1024 PB"},             // 1024 PB (beyond units)
		{1000 * 1000 * 1000 * 1000 * 1000, "953.7 PB"},                   // ~1000^5 in binary
	}

	for _, tt := range tests {
		result := ByteSize(tt.input)
		if result != tt.expected {
			t.Errorf("ByteSize(%d) = %s, want %s", tt.input, result, tt.expected)
		}
	}
}

// TestEdgeCases 测试边界情况
func TestEdgeCases(t *testing.T) {
	// Test very small durations
	result := Duration(1 * time.Nanosecond)
	if result == "" {
		t.Error("Duration(1ns) should not be empty")
	}

	// Test zero duration with different formats
	result = Duration(0)
	expected := "0 second"
	if result != expected {
		t.Errorf("Duration(0) = %s, want %s", result, expected)
	}

	result = Duration(0, WithClockFormat())
	expected = "0:00"
	if result != expected {
		t.Errorf("Duration(0, WithClockFormat()) = %s, want %s", result, expected)
	}

	// Test very precise values with high precision
	result = ByteSize(1234567, WithPrecision(5))
	if result == "" {
		t.Error("High precision ByteSize should not be empty")
	}
}

// TestFormatFloat 测试浮点数格式化
func TestFormatFloat(t *testing.T) {
	tests := []struct {
		value     float64
		precision int
		expected  string
	}{
		{1.0, 2, "1"},
		{1.5, 1, "1.5"},
		{1.23456, 3, "1.235"},
		{1.000, 3, "1"},
		{0.0, 2, "0"},
		{-1.5, 1, "-1.5"},
	}

	for _, tt := range tests {
		result := formatFloat(tt.value, tt.precision)
		if result != tt.expected {
			t.Errorf("formatFloat(%f, %d) = %s, want %s", tt.value, tt.precision, result, tt.expected)
		}
	}
}

// TestAbsFunction 测试abs函数
func TestAbsFunction(t *testing.T) {
	tests := []struct {
		input    int64
		expected int64
	}{
		{0, 0},
		{5, 5},
		{-5, 5},
		{1024, 1024},
		{-1024, 1024},
	}

	for _, tt := range tests {
		result := abs(tt.input)
		if result != tt.expected {
			t.Errorf("abs(%d) = %d, want %d", tt.input, result, tt.expected)
		}
	}
}

// TestRelativeTimeFuture 测试未来时间
func TestRelativeTimeFuture(t *testing.T) {
	future := time.Now().Add(5 * time.Minute)
	result := RelativeTime(future)
	
	// The exact result may vary, but it should not be empty
	if result == "" {
		t.Error("RelativeTime for future should not be empty")
	}
	
	// Test very distant future
	distantFuture := time.Now().Add(365 * 24 * time.Hour)
	result = RelativeTime(distantFuture)
	if result == "" {
		t.Error("RelativeTime for distant future should not be empty")
	}
}

// TestInvalidUnitType 测试无效单位类型
func TestInvalidUnitType(t *testing.T) {
	config := DefaultConfig()
	result := formatWithUnit(1.0, 0, config, "invalid")
	if result != "-" {
		t.Errorf("Invalid unit type should return '-', got %s", result)
	}
}

// TestConfigToOptions 测试配置转换
func TestConfigToOptions(t *testing.T) {
	config := Config{
		Precision:  2,
		Locale:     "zh-CN",
		Compact:    true,
		TimeFormat: "clock",
	}
	
	options := configToOptions(config)
	
	if options.Precision != config.Precision {
		t.Errorf("Precision conversion failed: %d != %d", options.Precision, config.Precision)
	}
	if options.Locale != config.Locale {
		t.Errorf("Locale conversion failed: %s != %s", options.Locale, config.Locale)
	}
	if options.Compact != config.Compact {
		t.Errorf("Compact conversion failed: %v != %v", options.Compact, config.Compact)
	}
	if options.TimeFormat != config.TimeFormat {
		t.Errorf("TimeFormat conversion failed: %s != %s", options.TimeFormat, config.TimeFormat)
	}
}

// TestCompactFormat 测试紧凑格式
func TestCompactFormat(t *testing.T) {
	tests := []struct {
		function func(int64, ...Option) string
		input    int64
		expected string
	}{
		{ByteSize, 1536, "1.5KB"},
		{Speed, 1536, "1.5KB/s"},
		{BitSpeed, 1500, "1.5Kbps"},
	}

	for _, tt := range tests {
		result := tt.function(tt.input, WithCompact())
		if result != tt.expected {
			t.Errorf("Compact format failed: got %s, want %s", result, tt.expected)
		}
	}
}

// TestClockTimeEdgeCases 测试时钟格式边界情况
func TestClockTimeEdgeCases(t *testing.T) {
	tests := []struct {
		duration time.Duration
		expected string
	}{
		{0, "0:00"},
		{30 * time.Second, "0:30"},
		{90 * time.Second, "1:30"},
		{1 * time.Hour, "1:00"},
		{61 * time.Minute, "1:01"},
		{3661 * time.Second, "1:01:01"}, // 1h 1m 1s
		{7200 * time.Second, "2:00"},    // exactly 2h
	}

	for _, tt := range tests {
		result := formatClockTime(tt.duration)
		if result != tt.expected {
			t.Errorf("formatClockTime(%v) = %s, want %s", tt.duration, result, tt.expected)
		}
	}
}

// TestLocaleRegistration 测试locale注册和获取
func TestLocaleRegistration(t *testing.T) {
	// Register a test locale
	testLocale := &Locale{
		Language:      "test",
		Region:        "TEST",
		ByteUnits:     []string{"B", "KB", "MB"},
		SpeedUnits:    []string{"B/s", "KB/s", "MB/s"},
		BitSpeedUnits: []string{"bps", "Kbps", "Mbps"},
	}
	
	RegisterLocale("test", testLocale)
	
	// Get the registered locale
	retrieved, ok := GetLocaleConfig("test")
	if !ok {
		t.Error("Failed to retrieve registered locale")
	}
	
	if retrieved.Language != "test" {
		t.Errorf("Retrieved locale language = %s, want test", retrieved.Language)
	}
	
	// Test fallback to language without region
	retrieved, ok = GetLocaleConfig("test-REGION")
	if !ok {
		t.Error("Failed to fallback to language locale")
	}
	
	// Test fallback to English
	retrieved, ok = GetLocaleConfig("nonexistent")
	if !ok {
		t.Error("Failed to fallback to English locale")
	}
	if retrieved.Language != "en" {
		t.Errorf("Fallback locale language = %s, want en", retrieved.Language)
	}
}

// TestFormatWithLocale 测试locale格式化
func TestFormatWithLocale(t *testing.T) {
	locale, _ := GetLocaleConfig("en")
	
	result := formatWithLocale(locale, "Test %s", "value")
	expected := "Test value"
	if result != expected {
		t.Errorf("formatWithLocale failed: %s != %s", result, expected)
	}
	
	// Test with nil locale (should fallback to English)
	result = formatWithLocale(nil, "Test %s", "value")
	if result != expected {
		t.Errorf("formatWithLocale with nil locale failed: %s != %s", result, expected)
	}
}

// TestGetTimeUnit 测试时间单位获取
func TestGetTimeUnit(t *testing.T) {
	locale, _ := GetLocaleConfig("en")
	
	// Test singular
	result := getTimeUnit(locale, locale.TimeUnits.Second, 1)
	if result != "second" {
		t.Errorf("getTimeUnit(1 second) = %s, want second", result)
	}
	
	// Test plural
	result = getTimeUnit(locale, locale.TimeUnits.Second, 2)
	if result != "seconds" {
		t.Errorf("getTimeUnit(2 seconds) = %s, want seconds", result)
	}
	
	// Test with nil locale
	result = getTimeUnit(nil, "second", 2)
	if result != "seconds" {
		t.Errorf("getTimeUnit with nil locale failed: %s != seconds", result)
	}
}

// TestCompatibilityFunctions 测试兼容性函数
func TestCompatibilityFunctions(t *testing.T) {
	// Test BitSpeedWithOptions
	result := BitSpeedWithOptions(8000, Options{Precision: 1})
	expected := "8 Kbps"
	if result != expected {
		t.Errorf("BitSpeedWithOptions = %s, want %s", result, expected)
	}
	
	// Test ClockDuration
	result = ClockDuration(90 * time.Second)
	expected = "1:30"
	if result != expected {
		t.Errorf("ClockDuration = %s, want %s", result, expected)
	}
	
	// Test DurationWithOptions
	result = DurationWithOptions(90*time.Second, Options{TimeFormat: "clock"})
	if result != expected {
		t.Errorf("DurationWithOptions = %s, want %s", result, expected)
	}
	
	// Test ByteSizeWithOptions
	result = ByteSizeWithOptions(1536, Options{Precision: 1})
	expected = "1.5 KB"
	if result != expected {
		t.Errorf("ByteSizeWithOptions = %s, want %s", result, expected)
	}
}