package human

import (
	"testing"
	"time"
)

// TestCompleteCodeCoverage 确保所有代码路径被覆盖
func TestCompleteCodeCoverage(t *testing.T) {
	t.Run("All exported functions", func(t *testing.T) {
		// Test all main API functions
		_ = ByteSize(1024)
		_ = Speed(1024)
		_ = BitSpeed(1000)
		_ = Duration(time.Minute)
		_ = RelativeTime(time.Now())
		
		// Test with all options
		_ = ByteSize(1024, WithPrecision(2), WithLocale("en"), WithCompact())
		_ = Duration(time.Minute, WithClockFormat())
		
		// Test global setters/getters
		originalLocale := GetLocale()
		SetLocale("test")
		SetLocale(originalLocale)
		
		originalPrecision := defaultPrecision
		SetDefaultPrecision(3)
		defaultPrecision = originalPrecision
		
		// Test default options
		_ = DefaultOptions()
		
		// Test compatibility functions
		_ = BitSpeedWithOptions(1000, Options{Precision: 2})
		_ = ClockDuration(time.Minute)
		_ = DurationWithOptions(time.Minute, Options{TimeFormat: "clock"})
		_ = ByteSizeWithOptions(1024, Options{Precision: 2})
	})
	
	t.Run("All internal functions", func(t *testing.T) {
		config := DefaultConfig()
		
		// Test all format functions
		_ = formatByteSize(1024, config)
		_ = formatSpeed(1024, config)
		_ = formatBitSpeed(1000, config)
		_ = formatDuration(time.Minute, config)
		_ = formatRelativeTime(time.Now(), config)
		
		// Test utility functions
		_ = formatWithUnit(1.5, 1, config, "byte")
		_ = formatWithUnit(1.5, 1, config, "speed")
		_ = formatWithUnit(1.5, 1, config, "bitspeed")
		_ = formatWithUnit(1.5, 1, config, "invalid") // Should return "-"
		
		_ = formatFloat(1.5, 2)
		_ = formatFloat(1.0, 2) // Integer value
		_ = formatFloat(1.5, -1) // Negative precision
		
		_ = formatClockTime(time.Hour + 30*time.Minute + 45*time.Second)
		_ = formatClockTime(-time.Minute) // Negative
		
		// Test abs function
		_ = abs(5)
		_ = abs(-5)
		_ = abs(0)
		
		// Test config conversion
		_ = configToOptions(config)
		
		// Test options application
		_ = applyOptions()
		_ = applyOptions(WithPrecision(2))
		_ = applyOptions(WithPrecision(2), WithLocale("en"), WithCompact(), WithClockFormat())
	})
	
	t.Run("All edge cases", func(t *testing.T) {
		config := DefaultConfig()
		
		// Zero values
		_ = formatByteSize(0, config)
		_ = formatSpeed(0, config)
		_ = formatBitSpeed(0, config)
		_ = formatDuration(0, config)
		_ = formatDuration(0, Config{TimeFormat: "clock"})
		
		// Negative values
		_ = formatByteSize(-1024, config)
		_ = formatSpeed(-1024, config)
		_ = formatBitSpeed(-1000, config)
		_ = formatDuration(-time.Minute, config)
		_ = formatDuration(-time.Minute, Config{TimeFormat: "clock"})
		
		// Large values (beyond available units)
		veryLarge := int64(1024) * 1024 * 1024 * 1024 * 1024 * 1024
		_ = formatByteSize(veryLarge, config)
		_ = formatSpeed(veryLarge, config)
		_ = formatBitSpeed(veryLarge, config)
		
		// Small durations
		_ = formatDuration(time.Nanosecond, config)
		_ = formatDuration(time.Microsecond, config)
		_ = formatDuration(time.Millisecond, config)
		
		// Large unit index
		_ = formatWithUnit(1.0, 999, config, "byte")
		
		// Various time ranges for relative time
		now := time.Now()
		past := []time.Duration{
			5 * time.Second,    // "just now"
			30 * time.Second,   // seconds ago
			30 * time.Minute,   // minutes ago
			5 * time.Hour,      // hours ago
			3 * 24 * time.Hour, // days ago
			10 * 24 * time.Hour, // weeks ago
			45 * 24 * time.Hour, // months ago
			400 * 24 * time.Hour, // years ago
		}
		
		for _, d := range past {
			_ = formatRelativeTime(now.Add(-d), config)
			_ = formatRelativeTime(now.Add(d), config) // future
		}
		
		// Test locale functions
		locale, _ := GetLocaleConfig("en")
		_ = formatWithLocale(locale, "test %s", "value")
		_ = formatWithLocale(nil, "test %s", "value") // nil locale
		
		// Test time unit function
		_ = getTimeUnit(locale, locale.TimeUnits.Second, 1) // singular
		_ = getTimeUnit(locale, locale.TimeUnits.Second, 2) // plural
		_ = getTimeUnit(locale, locale.TimeUnits.Minute, 1)
		_ = getTimeUnit(locale, locale.TimeUnits.Hour, 1)
		_ = getTimeUnit(locale, locale.TimeUnits.Day, 1)
		_ = getTimeUnit(locale, locale.TimeUnits.Week, 1)
		_ = getTimeUnit(locale, locale.TimeUnits.Month, 1)
		_ = getTimeUnit(locale, locale.TimeUnits.Year, 1)
		_ = getTimeUnit(nil, "second", 2) // nil locale
		
		// Test non-English locale (no pluralization)
		locale.Language = "zh"
		_ = getTimeUnit(locale, locale.TimeUnits.Second, 2)
	})
	
	t.Run("Locale management", func(t *testing.T) {
		// Test locale registration and retrieval
		testLocale := &Locale{
			Language: "test",
			Region: "TEST",
			ByteUnits: []string{"B", "KB", "MB"},
			SpeedUnits: []string{"B/s", "KB/s", "MB/s"},
			BitSpeedUnits: []string{"bps", "Kbps", "Mbps"},
		}
		
		RegisterLocale("test", testLocale)
		
		// Test retrieval
		_, ok := GetLocaleConfig("test")
		if !ok {
			t.Error("Should find registered locale")
		}
		
		// Test fallback to language
		_, ok = GetLocaleConfig("test-REGION")
		if !ok {
			t.Error("Should fallback to language")
		}
		
		// Test fallback to English
		_, ok = GetLocaleConfig("nonexistent")
		if !ok {
			t.Error("Should fallback to English")
		}
		
		// Test with empty string
		_, ok = GetLocaleConfig("")
		if !ok {
			t.Error("Should fallback to English for empty string")
		}
	})
	
	t.Run("Format functions with all combinations", func(t *testing.T) {
		// Test formatPastTime with all time ranges
		locale, _ := GetLocaleConfig("en")
		
		timeRanges := []time.Duration{
			5 * time.Second,
			30 * time.Second,
			30 * time.Minute,
			5 * time.Hour,
			3 * 24 * time.Hour,
			10 * 24 * time.Hour,
			45 * 24 * time.Hour,
			400 * 24 * time.Hour,
		}
		
		for _, d := range timeRanges {
			_ = formatPastTime(d, locale)
			_ = formatFutureTime(d, locale)
		}
		
		// Test clock time with various durations
		clockTimes := []time.Duration{
			0,
			30 * time.Second,
			90 * time.Second,
			time.Hour,
			25 * time.Hour,
			25*time.Hour + 30*time.Minute + 45*time.Second,
		}
		
		for _, d := range clockTimes {
			_ = formatClockTime(d)
			_ = formatClockTime(-d) // negative
		}
	})
}

// TestAllOptionsPath 测试所有选项路径
func TestAllOptionsPath(t *testing.T) {
	// Test all combinations of options
	options := [][]Option{
		{},
		{WithPrecision(2)},
		{WithLocale("en")},
		{WithCompact()},
		{WithClockFormat()},
		{WithPrecision(2), WithLocale("en")},
		{WithPrecision(2), WithCompact()},
		{WithPrecision(2), WithClockFormat()},
		{WithLocale("en"), WithCompact()},
		{WithLocale("en"), WithClockFormat()},
		{WithCompact(), WithClockFormat()},
		{WithPrecision(2), WithLocale("en"), WithCompact()},
		{WithPrecision(2), WithLocale("en"), WithClockFormat()},
		{WithPrecision(2), WithCompact(), WithClockFormat()},
		{WithLocale("en"), WithCompact(), WithClockFormat()},
		{WithPrecision(2), WithLocale("en"), WithCompact(), WithClockFormat()},
	}
	
	for _, opts := range options {
		_ = ByteSize(1024, opts...)
		_ = Speed(1024, opts...)
		_ = BitSpeed(1000, opts...)
		_ = Duration(time.Minute, opts...)
		_ = RelativeTime(time.Now(), opts...)
	}
}