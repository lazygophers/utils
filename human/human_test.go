package human

import (
	"testing"
	"time"

	"github.com/lazygophers/utils/language"
)

// resetState 将包级 flag 恢复默认，避免测试间互相污染。
func resetState() {
	defaultPrecision = 1
	defaultCompact = false
	defaultClockFormat = false
	language.Del()
}

func TestByteSize(t *testing.T) {
	resetState()
	defer resetState()

	type byteSizeCase struct {
		name     string
		input    int64
		expected string
	}
	tests := []byteSizeCase{
		{"0 bytes", 0, "0 B"},
		{"1 byte", 1, "1 B"},
		{"1024 bytes", 1024, "1 KB"},
		{"1536 bytes", 1536, "1.5 KB"},
		{"1MB", 1024 * 1024, "1 MB"},
		{"1.5MB", 1536 * 1024, "1.5 MB"},
		{"1GB", 1024 * 1024 * 1024, "1 GB"},
		{"1TB", 1024 * 1024 * 1024 * 1024, "1 TB"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := ByteSize(tt.input)
			if result != tt.expected {
				t.Errorf("ByteSize(%d) = %s, want %s", tt.input, result, tt.expected)
			}
		})
	}
}

func TestByteSizeWithLocale(t *testing.T) {
	resetState()
	defer resetState()

	language.Set(language.Make("zh-CN"))
	result := ByteSize(1024)
	expected := "1 KB" // 字节单位通常保持英文
	if result != expected {
		t.Errorf("ByteSize(1024) in zh-CN = %s, want %s", result, expected)
	}
}

func TestSpeed(t *testing.T) {
	resetState()
	defer resetState()

	type speedCase struct {
		name     string
		input    int64
		expected string
	}
	tests := []speedCase{
		{"0 B/s", 0, "0 B/s"},
		{"1 B/s", 1, "1 B/s"},
		{"1024 B/s", 1024, "1 KB/s"},
		{"1MB/s", 1024 * 1024, "1 MB/s"},
		{"1.5MB/s", 1536 * 1024, "1.5 MB/s"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Speed(tt.input)
			if result != tt.expected {
				t.Errorf("Speed(%d) = %s, want %s", tt.input, result, tt.expected)
			}
		})
	}
}

func TestBitSpeed(t *testing.T) {
	resetState()
	defer resetState()

	type bitSpeedCase struct {
		name     string
		input    int64
		expected string
	}
	tests := []bitSpeedCase{
		{"0 bps", 0, "0 bps"},
		{"512 bps", 512, "512 bps"},
		{"1000 bps", 1000, "1 Kbps"},
		{"1500 bps", 1500, "1.5 Kbps"},
		{"1,000,000 bps", 1000000, "1 Mbps"},
		{"1,500,000 bps", 1500000, "1.5 Mbps"},
		{"1,000,000,000 bps", 1000000000, "1 Gbps"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := BitSpeed(tt.input)
			if result != tt.expected {
				t.Errorf("BitSpeed(%d) = %s, want %s", tt.input, result, tt.expected)
			}
		})
	}
}

func TestSpeedVsBitSpeedDistinction(t *testing.T) {
	resetState()
	defer resetState()

	testValue := int64(8000)
	byteSpeedResult := Speed(testValue)
	bitSpeedResult := BitSpeed(testValue)
	expectedByteSpeed := "7.8 KB/s"
	expectedBitSpeed := "8 Kbps"

	if byteSpeedResult != expectedByteSpeed {
		t.Errorf("Speed(%d) = %s, want %s", testValue, byteSpeedResult, expectedByteSpeed)
	}
	if bitSpeedResult != expectedBitSpeed {
		t.Errorf("BitSpeed(%d) = %s, want %s", testValue, bitSpeedResult, expectedBitSpeed)
	}
	if byteSpeedResult == bitSpeedResult {
		t.Error("Speed() and BitSpeed() should produce different results for the same input")
	}
}

func TestBitSpeedCompact(t *testing.T) {
	resetState()
	defer resetState()

	SetCompact(true)
	result := BitSpeed(1500)
	if result != "1.5Kbps" {
		t.Errorf("BitSpeed compact = %s, want 1.5Kbps", result)
	}
}

func TestBitSpeedPrecision(t *testing.T) {
	resetState()
	defer resetState()

	SetCompact(true)
	SetDefaultPrecision(2)
	result := BitSpeed(1234)
	if result != "1.23Kbps" {
		t.Errorf("BitSpeed precision/compact = %s, want 1.23Kbps", result)
	}
}

func TestInvalidInputs(t *testing.T) {
	resetState()
	defer resetState()

	result := formatWithUnit(1.0, 0, "invalid")
	if result != "-" {
		t.Errorf("Invalid input should return \"-\", got %s", result)
	}
}

func TestDuration(t *testing.T) {
	resetState()
	defer resetState()

	type durationCase struct {
		name     string
		input    time.Duration
		expected string
	}
	tests := []durationCase{
		{"0 seconds", 0, "0 second"},
		{"1 second", time.Second, "1 second"},
		{"30 seconds", 30 * time.Second, "30 seconds"},
		{"1 minute", time.Minute, "1 minute"},
		{"90 seconds", 90 * time.Second, "1 minute 30 seconds"},
		{"1 hour", time.Hour, "1 hour"},
		{"90 minutes", 90 * time.Minute, "1 hour 30 minutes"},
		{"25 hours", 25 * time.Hour, "1 day 1 hour"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Duration(tt.input)
			if result != tt.expected {
				t.Errorf("Duration(%v) = %s, want %s", tt.input, result, tt.expected)
			}
		})
	}
}

func TestClockDuration(t *testing.T) {
	resetState()
	defer resetState()

	type clockDurationCase struct {
		name     string
		input    time.Duration
		expected string
	}
	tests := []clockDurationCase{
		{"0 seconds", 0, "0:00"},
		{"30 seconds", 30 * time.Second, "0:30"},
		{"1 minute", time.Minute, "1:00"},
		{"90 seconds", 90 * time.Second, "1:30"},
		{"1 hour", time.Hour, "1:00"},
		{"1 hour 10 minutes", time.Hour + 10*time.Minute, "1:10"},
		{"2 hours 30 minutes 45 seconds", 2*time.Hour + 30*time.Minute + 45*time.Second, "2:30:45"},
		{"4 hours 5 minutes", 4*time.Hour + 5*time.Minute, "4:05"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := ClockDuration(tt.input)
			if result != tt.expected {
				t.Errorf("ClockDuration(%v) = %s, want %s", tt.input, result, tt.expected)
			}
		})
	}
}

func TestDurationWithClockFormat(t *testing.T) {
	resetState()
	defer resetState()

	SetClockFormat(true)

	type clockFmtCase struct {
		name     string
		input    time.Duration
		expected string
	}
	tests := []clockFmtCase{
		{"30 seconds with clock format", 30 * time.Second, "0:30"},
		{"1 hour 10 minutes with clock format", time.Hour + 10*time.Minute, "1:10"},
		{"negative duration", -90 * time.Second, "-1:30"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Duration(tt.input)
			if result != tt.expected {
				t.Errorf("Duration(%v, clock) = %s, want %s", tt.input, result, tt.expected)
			}
		})
	}
}

func TestRelativeTime(t *testing.T) {
	resetState()
	defer resetState()

	now := time.Now()

	type relativeCase struct {
		name  string
		input time.Time
	}
	tests := []relativeCase{
		{"just now", now.Add(-5 * time.Second)},
		{"30 seconds ago", now.Add(-30 * time.Second)},
		{"2 minutes ago", now.Add(-2 * time.Minute)},
		{"1 hour ago", now.Add(-1 * time.Hour)},
		{"1 day ago", now.Add(-24 * time.Hour)},
		{"in 5 minutes", now.Add(5 * time.Minute)},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := RelativeTime(tt.input)
			if result == "" {
				t.Errorf("RelativeTime(%v) returned empty string", tt.input)
			}
		})
	}
}

func TestByteSizeCompactPrecision(t *testing.T) {
	resetState()
	defer resetState()

	SetDefaultPrecision(2)
	SetCompact(true)
	result := ByteSize(1536)
	expected := "1.5KB"
	if result != expected {
		t.Errorf("ByteSize(1536, compact) = %s, want %s", result, expected)
	}
}

// 基准测试
func BenchmarkByteSize(b *testing.B) {
	resetState()
	for i := 0; i < b.N; i++ {
		ByteSize(1073741824)
	}
}

func BenchmarkSpeed(b *testing.B) {
	resetState()
	for i := 0; i < b.N; i++ {
		Speed(104857600)
	}
}

func BenchmarkDuration(b *testing.B) {
	resetState()
	for i := 0; i < b.N; i++ {
		Duration(90 * time.Minute)
	}
}

func BenchmarkRelativeTime(b *testing.B) {
	resetState()
	t := time.Now().Add(-5 * time.Minute)
	for i := 0; i < b.N; i++ {
		RelativeTime(t)
	}
}

// TestCompleteCodeCoverage 确保所有代码路径被覆盖
func TestCompleteCodeCoverage(t *testing.T) {
	t.Run("All exported functions", func(t *testing.T) {
		resetState()
		defer resetState()

		_ = ByteSize(1024)
		_ = Speed(1024)
		_ = BitSpeed(1000)
		_ = Duration(time.Minute)
		_ = RelativeTime(time.Now())
		_ = ClockDuration(time.Minute)

		SetDefaultPrecision(3)
		SetCompact(true)
		_ = ByteSize(1024)
		SetClockFormat(true)
		_ = Duration(time.Minute)
	})

	t.Run("All internal functions", func(t *testing.T) {
		resetState()
		defer resetState()

		_ = formatByteSize(1024)
		_ = formatSpeed(1024)
		_ = formatBitSpeed(1000)
		_ = formatDuration(time.Minute)
		_ = formatRelativeTime(time.Now())

		_ = formatWithUnit(1.5, 1, "byte")
		_ = formatWithUnit(1.5, 1, "speed")
		_ = formatWithUnit(1.5, 1, "bitspeed")
		_ = formatWithUnit(1.5, 1, "invalid")

		_ = formatFloat(1.5, 2)
		_ = formatFloat(1.0, 2)
		_ = formatFloat(1.5, -1)

		_ = formatClockTime(time.Hour + 30*time.Minute + 45*time.Second)
		_ = formatClockTime(-time.Minute)

		_ = abs(5)
		_ = abs(-5)
		_ = abs(0)
	})

	t.Run("Edge cases", func(t *testing.T) {
		resetState()
		defer resetState()

		_ = formatByteSize(0)
		_ = formatSpeed(0)
		_ = formatBitSpeed(0)
		_ = formatDuration(0)

		SetClockFormat(true)
		_ = formatDuration(0)
		_ = formatDuration(-time.Minute)
		SetClockFormat(false)

		_ = formatByteSize(-1024)
		_ = formatSpeed(-1024)
		_ = formatBitSpeed(-1000)
		_ = formatDuration(-time.Minute)

		veryLarge := int64(1024) * 1024 * 1024 * 1024 * 1024 * 1024
		_ = formatByteSize(veryLarge)
		_ = formatSpeed(veryLarge)
		_ = formatBitSpeed(veryLarge)

		_ = formatDuration(time.Nanosecond)
		_ = formatDuration(time.Microsecond)
		_ = formatDuration(time.Millisecond)

		_ = formatWithUnit(1.0, 999, "byte")

		now := time.Now()
		past := []time.Duration{
			5 * time.Second,
			30 * time.Second,
			30 * time.Minute,
			5 * time.Hour,
			3 * 24 * time.Hour,
			10 * 24 * time.Hour,
			45 * 24 * time.Hour,
			400 * 24 * time.Hour,
		}
		for _, d := range past {
			_ = formatRelativeTime(now.Add(-d))
			_ = formatRelativeTime(now.Add(d))
		}

		locale, _ := GetLocaleConfig("en")
		_ = formatWithLocale(locale, "test %s", "value")
		_ = formatWithLocale(nil, "test %s", "value")

		_ = getTimeUnit(locale, locale.TimeUnits.Second, 1)
		_ = getTimeUnit(locale, locale.TimeUnits.Second, 2)
		_ = getTimeUnit(locale, locale.TimeUnits.Minute, 1)
		_ = getTimeUnit(locale, locale.TimeUnits.Hour, 1)
		_ = getTimeUnit(locale, locale.TimeUnits.Day, 1)
		_ = getTimeUnit(locale, locale.TimeUnits.Week, 1)
		_ = getTimeUnit(locale, locale.TimeUnits.Month, 1)
		_ = getTimeUnit(locale, locale.TimeUnits.Year, 1)
		_ = getTimeUnit(nil, "second", 2)
	})

	t.Run("Locale management", func(t *testing.T) {
		resetState()
		defer resetState()

		testLocale := &Locale{
			Language:      "test",
			Region:        "TEST",
			ByteUnits:     []string{"B", "KB", "MB"},
			SpeedUnits:    []string{"B/s", "KB/s", "MB/s"},
			BitSpeedUnits: []string{"bps", "Kbps", "Mbps"},
		}
		RegisterLocale("test", testLocale)

		_, ok := GetLocaleConfig("test")
		if !ok {
			t.Error("Should find registered locale")
		}
		_, ok = GetLocaleConfig("test-REGION")
		if !ok {
			t.Error("Should fallback to language")
		}
		_, ok = GetLocaleConfig("nonexistent")
		if !ok {
			t.Error("Should fallback to English")
		}
		_, ok = GetLocaleConfig("")
		if !ok {
			t.Error("Should fallback to English for empty string")
		}
	})
}

// TestSetCurrentLocale 测试 goroutine-local locale 切换
func TestSetCurrentLocale(t *testing.T) {
	resetState()
	defer resetState()

	language.Set(language.Make("zh-CN"))
	if got := currentLocaleName(); got != "zh" {
		t.Errorf("currentLocaleName after Set(zh-CN) = %s, want zh", got)
	}

	language.Set(language.Make("ja"))
	if got := currentLocaleName(); got != "ja" {
		t.Errorf("currentLocaleName after Set(ja) = %s, want ja", got)
	}

	language.Del()
	// 默认 fallback
	if got := currentLocaleName(); got != "en" {
		t.Errorf("currentLocaleName after Del = %s, want en", got)
	}
}

// TestSetDefaultPrecision 测试默认精度设置
func TestSetDefaultPrecision(t *testing.T) {
	resetState()
	defer resetState()

	SetDefaultPrecision(3)
	if defaultPrecision != 3 {
		t.Errorf("Expected precision 3, got %d", defaultPrecision)
	}
}

// TestSetCompact 测试紧凑模式 setter
func TestSetCompact(t *testing.T) {
	resetState()
	defer resetState()

	SetCompact(true)
	if !defaultCompact {
		t.Error("SetCompact(true) failed")
	}
	SetCompact(false)
	if defaultCompact {
		t.Error("SetCompact(false) failed")
	}
}

// TestSetClockFormat 测试时钟格式 setter
func TestSetClockFormat(t *testing.T) {
	resetState()
	defer resetState()

	SetClockFormat(true)
	if !defaultClockFormat {
		t.Error("SetClockFormat(true) failed")
	}
}

// TestNegativeValues 测试负值处理
func TestNegativeValues(t *testing.T) {
	resetState()
	defer resetState()

	type negCase struct {
		got      string
		expected string
		name     string
	}

	cases := []negCase{
		{ByteSize(-1024), "-1 KB", "ByteSize"},
		{Speed(-2048), "-2 KB/s", "Speed"},
		{BitSpeed(-8000), "-8 Kbps", "BitSpeed"},
		{Duration(-90 * time.Second), "-1 minute 30 seconds", "Duration"},
	}

	for _, c := range cases {
		if c.got != c.expected {
			t.Errorf("%s: got %s, want %s", c.name, c.got, c.expected)
		}
	}

	SetClockFormat(true)
	result := Duration(-90 * time.Second)
	if result != "-1:30" {
		t.Errorf("Duration(-90s, clock) = %s, want -1:30", result)
	}
}

// TestLargeValues 测试大数值处理
func TestLargeValues(t *testing.T) {
	resetState()
	defer resetState()

	type largeByteCase struct {
		input    int64
		expected string
	}
	tests := []largeByteCase{
		{1024 * 1024 * 1024 * 1024 * 1024, "1 PB"},
		{1024 * 1024 * 1024 * 1024 * 1024 * 1024, "1024 PB"},
		{1000 * 1000 * 1000 * 1000 * 1000, "909.5 TB"},
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
	resetState()
	defer resetState()

	result := Duration(1 * time.Nanosecond)
	if result == "" {
		t.Error("Duration(1ns) should not be empty")
	}

	result = Duration(0)
	if result != "0 second" {
		t.Errorf("Duration(0) = %s, want 0 second", result)
	}

	SetClockFormat(true)
	result = Duration(0)
	if result != "0:00" {
		t.Errorf("Duration(0, clock) = %s, want 0:00", result)
	}
	SetClockFormat(false)

	SetDefaultPrecision(5)
	result = ByteSize(1234567)
	if result == "" {
		t.Error("High precision ByteSize should not be empty")
	}
}

// TestFormatFloat 测试浮点数格式化
func TestFormatFloat(t *testing.T) {
	type formatFloatCase struct {
		value     float64
		precision int
		expected  string
	}
	tests := []formatFloatCase{
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
	type absCase struct {
		input    int64
		expected int64
	}
	tests := []absCase{
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
	resetState()
	defer resetState()

	future := time.Now().Add(5 * time.Minute)
	result := RelativeTime(future)
	if result == "" {
		t.Error("RelativeTime for future should not be empty")
	}

	distantFuture := time.Now().Add(365 * 24 * time.Hour)
	result = RelativeTime(distantFuture)
	if result == "" {
		t.Error("RelativeTime for distant future should not be empty")
	}
}

// TestCompactFormat 测试紧凑格式
func TestCompactFormat(t *testing.T) {
	resetState()
	defer resetState()

	type compactCase struct {
		fn       func(int64) string
		input    int64
		expected string
	}
	tests := []compactCase{
		{ByteSize, 1536, "1.5KB"},
		{Speed, 1536, "1.5KB/s"},
		{BitSpeed, 1500, "1.5Kbps"},
	}

	SetCompact(true)
	for _, tt := range tests {
		result := tt.fn(tt.input)
		if result != tt.expected {
			t.Errorf("Compact format failed: got %s, want %s", result, tt.expected)
		}
	}
}

// TestClockTimeEdgeCases 测试时钟格式边界情况
func TestClockTimeEdgeCases(t *testing.T) {
	resetState()
	defer resetState()

	type clockEdgeCase struct {
		duration time.Duration
		expected string
	}
	tests := []clockEdgeCase{
		{0, "0:00"},
		{30 * time.Second, "0:30"},
		{90 * time.Second, "1:30"},
		{1 * time.Hour, "1:00"},
		{61 * time.Minute, "1:01"},
		{3661 * time.Second, "1:01:01"},
		{7200 * time.Second, "2:00"},
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
	resetState()
	defer resetState()

	testLocale := &Locale{
		Language:      "regtest",
		Region:        "TEST",
		ByteUnits:     []string{"B", "KB", "MB"},
		SpeedUnits:    []string{"B/s", "KB/s", "MB/s"},
		BitSpeedUnits: []string{"bps", "Kbps", "Mbps"},
	}
	RegisterLocale("regtest", testLocale)

	retrieved, ok := GetLocaleConfig("regtest")
	if !ok {
		t.Error("Failed to retrieve registered locale")
	}
	if retrieved.Language != "regtest" {
		t.Errorf("Retrieved locale language = %s, want regtest", retrieved.Language)
	}

	_, ok = GetLocaleConfig("regtest-REGION")
	if !ok {
		t.Error("Failed to fallback to language locale")
	}

	_, ok = GetLocaleConfig("nonexistent")
	if !ok {
		t.Error("Failed to fallback to English locale")
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

	result = formatWithLocale(nil, "Test %s", "value")
	if result != expected {
		t.Errorf("formatWithLocale with nil locale failed: %s != %s", result, expected)
	}
}

// TestGetTimeUnit 测试时间单位获取
func TestGetTimeUnit(t *testing.T) {
	locale, _ := GetLocaleConfig("en")

	result := getTimeUnit(locale, locale.TimeUnits.Second, 1)
	if result != "second" {
		t.Errorf("getTimeUnit(1 second) = %s, want second", result)
	}

	// nil locale falls back to en, so 2 pluralizes
	result = getTimeUnit(nil, "second", 2)
	if result != "seconds" {
		t.Errorf("getTimeUnit with nil locale failed: %s != seconds", result)
	}
}

// TestGetTimeUnitAllCases 测试getTimeUnit的所有switch情况
func TestGetTimeUnitAllCases(t *testing.T) {
	enLocale := &Locale{
		Language: "en",
		TimeUnits: TimeUnits{
			Second:  "second",
			Minute:  "minute",
			Hour:    "hour",
			Day:     "day",
			Week:    "week",
			Month:   "month",
			Year:    "year",
			Seconds: "seconds",
			Minutes: "minutes",
			Hours:   "hours",
			Days:    "days",
			Weeks:   "weeks",
			Months:  "months",
			Years:   "years",
		},
	}

	type timeUnitCase struct {
		unit     string
		count    int64
		expected string
	}
	testCases := []timeUnitCase{
		{"second", 2, "seconds"},
		{"minute", 2, "minutes"},
		{"hour", 2, "hours"},
		{"day", 2, "days"},
		{"week", 2, "weeks"},
		{"month", 2, "months"},
		{"year", 2, "years"},
		{"second", 1, "second"},
		{"unknown", 2, "unknown"},
	}

	for _, tc := range testCases {
		result := getTimeUnit(enLocale, tc.unit, tc.count)
		if result != tc.expected {
			t.Errorf("getTimeUnit(%s, %d) = %s, want %s", tc.unit, tc.count, result, tc.expected)
		}
	}

	zhLocale := &Locale{Language: "zh"}
	result := getTimeUnit(zhLocale, "second", 2)
	if result != "second" {
		t.Errorf("getTimeUnit with zh locale should not pluralize: got %s", result)
	}
}

// TestFormatWithUnitLargeIndex 测试单位索引超出范围的情况
func TestFormatWithUnitLargeIndex(t *testing.T) {
	resetState()
	defer resetState()

	result := formatWithUnit(1.0, 999, "byte")
	if result == "" {
		t.Error("formatWithUnit with large index should not return empty string")
	}
}

// TestFormatByteSizeWithLargeValue 测试非常大的字节值
func TestFormatByteSizeWithLargeValue(t *testing.T) {
	resetState()
	defer resetState()

	veryLargeValue := int64(1024) * 1024 * 1024 * 1024 * 1024 * 1024
	result := formatByteSize(veryLargeValue)
	if result == "" {
		t.Error("formatByteSize with very large value should not return empty string")
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
	resetState()
	defer resetState()

	result := formatDuration(1 * time.Nanosecond)
	if result == "" {
		t.Error("formatDuration(1ns) should not return empty string")
	}
	result = formatDuration(1 * time.Microsecond)
	if result == "" {
		t.Error("formatDuration(1μs) should not return empty string")
	}
	result = formatDuration(1 * time.Millisecond)
	if result == "" {
		t.Error("formatDuration(1ms) should not return empty string")
	}
}

// TestFormatClockTimeVeryLarge 测试非常大的时钟时间
func TestFormatClockTimeVeryLarge(t *testing.T) {
	result := formatClockTime(25 * time.Hour)
	if result != "25:00" {
		t.Errorf("formatClockTime(25h) = %s, want 25:00", result)
	}

	result = formatClockTime(25*time.Hour + 30*time.Minute + 45*time.Second)
	if result != "25:30:45" {
		t.Errorf("formatClockTime(25h30m45s) = %s, want 25:30:45", result)
	}
}

// TestFormatRelativeTimeEdgeCases 测试相对时间的边界情况
func TestFormatRelativeTimeEdgeCases(t *testing.T) {
	resetState()
	defer resetState()

	now := time.Now()
	past := now.Add(-10 * time.Second)
	result := formatRelativeTime(past)
	if result == "" {
		t.Error("formatRelativeTime(10s ago) should not return empty string")
	}

	future := now.Add(59 * time.Second)
	result = formatRelativeTime(future)
	if result == "" {
		t.Error("formatRelativeTime(59s future) should not return empty string")
	}
}

// TestFormatPastTimeAllRanges 测试过去时间的所有时间范围
func TestFormatPastTimeAllRanges(t *testing.T) {
	locale, _ := GetLocaleConfig("en")

	durations := []time.Duration{
		5 * time.Second,
		30 * time.Second,
		30 * time.Minute,
		5 * time.Hour,
		3 * 24 * time.Hour,
		10 * 24 * time.Hour,
		45 * 24 * time.Hour,
		400 * 24 * time.Hour,
	}

	for _, d := range durations {
		result := formatPastTime(d, locale)
		if result == "" {
			t.Errorf("formatPastTime(%v) should not return empty string", d)
		}
	}
}

// TestFormatFutureTimeAllRanges 测试未来时间的所有时间范围
func TestFormatFutureTimeAllRanges(t *testing.T) {
	locale, _ := GetLocaleConfig("en")

	durations := []time.Duration{
		30 * time.Second,
		30 * time.Minute,
		5 * time.Hour,
		3 * 24 * time.Hour,
		10 * 24 * time.Hour,
		45 * 24 * time.Hour,
		400 * 24 * time.Hour,
	}

	for _, d := range durations {
		result := formatFutureTime(d, locale)
		if result == "" {
			t.Errorf("formatFutureTime(%v) should not return empty string", d)
		}
	}
}

// TestIntegerValues 测试整数值（无小数点）
func TestIntegerValues(t *testing.T) {
	resetState()
	defer resetState()

	type integerValueCase struct {
		input    int64
		expected string
	}
	tests := []integerValueCase{
		{1024, "1 KB"},
		{2048, "2 KB"},
		{1024 * 1024, "1 MB"},
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
	resetState()
	defer resetState()

	result := formatWithUnit(1.0, 1, "byte")
	if result != "1 KB" {
		t.Errorf("formatWithUnit(1.0, 1, byte) = %s, want 1 KB", result)
	}

	result = formatWithUnit(1.5, 1, "byte")
	if result != "1.5 KB" {
		t.Errorf("formatWithUnit(1.5, 1, byte) = %s, want 1.5 KB", result)
	}
}

// TestAllUnitTypes 测试所有单位类型
func TestAllUnitTypes(t *testing.T) {
	resetState()
	defer resetState()

	unitTypes := []string{"byte", "speed", "bitspeed"}
	for _, unitType := range unitTypes {
		result := formatWithUnit(1.5, 1, unitType)
		if result == "" {
			t.Errorf("formatWithUnit with unitType %s should not return empty string", unitType)
		}
	}
}

// TestGetTimeUnitEdgeCases 测试getTimeUnit的边界情况
func TestGetTimeUnitEdgeCases(t *testing.T) {
	locale, _ := GetLocaleConfig("en")

	// count != 1 triggers pluralization in en, so 0 → "seconds"
	result := getTimeUnit(locale, locale.TimeUnits.Second, 0)
	if result != "seconds" {
		t.Errorf("getTimeUnit(0) = %s, want seconds", result)
	}

	zhLocale := &Locale{Language: "zh", TimeUnits: TimeUnits{Second: "秒"}}
	result = getTimeUnit(zhLocale, zhLocale.TimeUnits.Second, 2)
	if result != zhLocale.TimeUnits.Second {
		t.Errorf("Non-English locale should not pluralize: %s", result)
	}
}

// TestDirectFunctionAPI 测试直接函数API设计
func TestDirectFunctionAPI(t *testing.T) {
	t.Run("ByteSize direct function", func(t *testing.T) {
		resetState()
		defer resetState()
		if result := ByteSize(1024); result != "1 KB" {
			t.Errorf("ByteSize(1024) = %s, want 1 KB", result)
		}
	})

	t.Run("ByteSize with global setters", func(t *testing.T) {
		resetState()
		defer resetState()
		SetDefaultPrecision(2)
		SetCompact(true)
		result := ByteSize(1536)
		if result != "1.5KB" {
			t.Errorf("ByteSize(1536) = %s, want 1.5KB", result)
		}
	})

	t.Run("BitSpeed direct function (1000-based)", func(t *testing.T) {
		resetState()
		defer resetState()
		if result := BitSpeed(8000); result != "8 Kbps" {
			t.Errorf("BitSpeed(8000) = %s, want 8 Kbps", result)
		}
	})

	t.Run("Speed vs BitSpeed distinction", func(t *testing.T) {
		resetState()
		defer resetState()
		if Speed(8000) != "7.8 KB/s" {
			t.Errorf("Speed(8000) wrong")
		}
		if BitSpeed(8000) != "8 Kbps" {
			t.Errorf("BitSpeed(8000) wrong")
		}
	})

	t.Run("Duration direct function", func(t *testing.T) {
		resetState()
		defer resetState()
		result := Duration(90 * time.Second)
		expected := "1 minute 30 seconds"
		if result != expected {
			t.Errorf("Duration(90s) = %s, want %s", result, expected)
		}
	})

	t.Run("Duration with clock format setter", func(t *testing.T) {
		resetState()
		defer resetState()
		SetClockFormat(true)
		result := Duration(90 * time.Second)
		if result != "1:30" {
			t.Errorf("Duration(90s, clock) = %s, want 1:30", result)
		}
	})

	t.Run("Zero value handling", func(t *testing.T) {
		resetState()
		defer resetState()
		if result := ByteSize(0); result != "0 B" {
			t.Errorf("ByteSize(0) = %s, want 0 B", result)
		}
		SetClockFormat(true)
		if result := Duration(0); result != "0:00" {
			t.Errorf("Duration(0, clock) = %s, want 0:00", result)
		}
	})
}
