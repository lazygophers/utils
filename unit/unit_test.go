package unit

import (
	"testing"
	"time"
)

func TestDurationYearMonthDay(t *testing.T) {
	tests := []struct {
		name     string
		duration time.Duration
		expected string
	}{
		{
			name:     "zero duration",
			duration: 0,
			expected: "0天",
		},
		{
			name:     "one day",
			duration: 24 * time.Hour,
			expected: "1天",
		},
		{
			name:     "one month",
			duration: 30 * 24 * time.Hour,
			expected: "1月",
		},
		{
			name:     "one year",
			duration: 365 * 24 * time.Hour,
			expected: "1年",
		},
		{
			name:     "complex duration",
			duration: 395*24*time.Hour + 5*24*time.Hour,
			expected: "1年1月5天",
		},
		{
			name:     "years and days only",
			duration: 370 * 24 * time.Hour,
			expected: "1年5天",
		},
		{
			name:     "months and days",
			duration: 35 * 24 * time.Hour,
			expected: "1月5天",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := DurationYearMonthDay(tt.duration)
			if result != tt.expected {
				t.Errorf("DurationYearMonthDay(%v) = %q, expected %q", tt.duration, result, tt.expected)
			}
		})
	}
}

func TestDurationMonthDayHour(t *testing.T) {
	tests := []struct {
		name     string
		duration time.Duration
		expected string
	}{
		{
			name:     "zero duration",
			duration: 0,
			expected: "0小时",
		},
		{
			name:     "one hour",
			duration: time.Hour,
			expected: "1小时",
		},
		{
			name:     "one day",
			duration: 24 * time.Hour,
			expected: "1天",
		},
		{
			name:     "one month",
			duration: 30 * 24 * time.Hour,
			expected: "1月",
		},
		{
			name:     "complex duration",
			duration: 35*24*time.Hour + 5*time.Hour,
			expected: "1月5天5小时",
		},
		{
			name:     "hours only",
			duration: 5 * time.Hour,
			expected: "5小时",
		},
		{
			name:     "days and hours",
			duration: 2*24*time.Hour + 3*time.Hour,
			expected: "2天3小时",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := DurationMonthDayHour(tt.duration)
			if result != tt.expected {
				t.Errorf("DurationMonthDayHour(%v) = %q, expected %q", tt.duration, result, tt.expected)
			}
		})
	}
}

func TestDurationMinuteSecond(t *testing.T) {
	tests := []struct {
		name     string
		duration time.Duration
		expected string
	}{
		{
			name:     "zero duration",
			duration: 0,
			expected: "0秒",
		},
		{
			name:     "one second",
			duration: time.Second,
			expected: "1秒",
		},
		{
			name:     "one minute",
			duration: time.Minute,
			expected: "0秒",
		},
		{
			name:     "two minutes",
			duration: 2*time.Minute,
			expected: "2分",
		},
		{
			name:     "minutes and seconds",
			duration: 2*time.Minute + 30*time.Second,
			expected: "2分30秒",
		},
		{
			name:     "seconds only",
			duration: 45 * time.Second,
			expected: "45秒",
		},
		{
			name:     "large minutes",
			duration: 90 * time.Minute,
			expected: "90分",
		},
		{
			name:     "subsecond duration",
			duration: 500 * time.Millisecond,
			expected: "0秒",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := DurationMinuteSecond(tt.duration)
			if result != tt.expected {
				t.Errorf("DurationMinuteSecond(%v) = %q, expected %q", tt.duration, result, tt.expected)
			}
		})
	}
}

func TestDurationYearMonthDayHourMinuteSecond(t *testing.T) {
	tests := []struct {
		name     string
		duration time.Duration
		expected string
	}{
		{
			name:     "zero duration",
			duration: 0,
			expected: "0秒",
		},
		{
			name:     "complex full duration",
			duration: 365*24*time.Hour + 30*24*time.Hour + 2*24*time.Hour + 2*time.Hour + 3*time.Minute + 4*time.Second,
			expected: "1年1月2天2小时3分4秒",
		},
		{
			name:     "seconds only",
			duration: 5 * time.Second,
			expected: "5秒",
		},
		{
			name:     "minutes and seconds only",
			duration: 2*time.Minute + 30*time.Second,
			expected: "2分30秒",
		},
		{
			name:     "all units",
			duration: 365*24*time.Hour + 30*24*time.Hour + 10*time.Hour + 30*time.Minute + 45*time.Second,
			expected: "1年1月10小时30分45秒",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := DurationYearMonthDayHourMinuteSecond(tt.duration)
			if result != tt.expected {
				t.Errorf("DurationYearMonthDayHourMinuteSecond(%v) = %q, expected %q", tt.duration, result, tt.expected)
			}
		})
	}
}

func TestTimeYearMonthDayHourMinute(t *testing.T) {
	testTime := time.Date(2023, 12, 25, 14, 30, 45, 0, time.UTC)
	expected := "2023年12月25日14点30"
	result := TimeYearMonthDayHourMinute(testTime)
	
	if result != expected {
		t.Errorf("TimeYearMonthDayHourMinute(%v) = %q, expected %q", testTime, result, expected)
	}
}

func TestTimeYearMonthDayHourMinuteSecond(t *testing.T) {
	testTime := time.Date(2023, 12, 25, 14, 30, 45, 0, time.UTC)
	expected := "2023年12月25日14点30分45"
	result := TimeYearMonthDayHourMinuteSecond(testTime)
	
	if result != expected {
		t.Errorf("TimeYearMonthDayHourMinuteSecond(%v) = %q, expected %q", testTime, result, expected)
	}
}

func TestFormatSpeed(t *testing.T) {
	tests := []struct {
		name     string
		speed    float64
		expected string
	}{
		{
			name:     "zero speed",
			speed:    0,
			expected: "——",
		},
		{
			name:     "negative speed",
			speed:    -10,
			expected: "——",
		},
		{
			name:     "bps",
			speed:    100,
			expected: "800.00 bps",
		},
		{
			name:     "Kbps",
			speed:    float64(Kb/8 + 100),
			expected: "1.10 Kbps",
		},
		{
			name:     "Mbps",
			speed:    float64(Mb/8),
			expected: "1.00 Mbps",
		},
		{
			name:     "Gbps",
			speed:    float64(Gb/8),
			expected: "1.00 Gbps",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := FormatSpeed(tt.speed)
			if result != tt.expected {
				t.Errorf("FormatSpeed(%f) = %q, expected %q", tt.speed, result, tt.expected)
			}
		})
	}
}

func TestFormat2bps(t *testing.T) {
	tests := []struct {
		name     string
		speed    float64
		expected string
	}{
		{
			name:     "zero speed",
			speed:    0,
			expected: "——",
		},
		{
			name:     "negative speed",
			speed:    -10,
			expected: "——",
		},
		{
			name:     "small bps",
			speed:    10,
			expected: "80.00 bps",
		},
		{
			name:     "Kbps range",
			speed:    float64(Kb/8 + 100),
			expected: "1.10 Kbps",
		},
		{
			name:     "Mbps range",
			speed:    float64(Mb/8),
			expected: "1.00 Mbps",
		},
		{
			name:     "very large speed",
			speed:    float64(Eb),
			expected: "8.00 Ebps",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Format2bps(tt.speed)
			if result != tt.expected {
				t.Errorf("Format2bps(%f) = %q, expected %q", tt.speed, result, tt.expected)
			}
		})
	}
}

func TestFormat2Bs(t *testing.T) {
	tests := []struct {
		name     string
		speed    float64
		expected string
	}{
		{
			name:     "zero speed",
			speed:    0,
			expected: "——",
		},
		{
			name:     "negative speed",
			speed:    -5,
			expected: "——",
		},
		{
			name:     "bytes per second",
			speed:    500,
			expected: "500.00 B/s",
		},
		{
			name:     "KB per second",
			speed:    float64(2 * Kb),
			expected: "2.00 KB/s",
		},
		{
			name:     "MB per second",
			speed:    float64(5 * Mb),
			expected: "5.00 MB/s",
		},
		{
			name:     "GB per second",
			speed:    float64(3 * Gb),
			expected: "3.00 GB/s",
		},
		{
			name:     "TB per second",
			speed:    float64(2 * Tb),
			expected: "2.00 TB/s",
		},
		{
			name:     "PB per second",
			speed:    float64(1 * Pb),
			expected: "1.00 PB/s",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Format2Bs(tt.speed)
			if result != tt.expected {
				t.Errorf("Format2Bs(%f) = %q, expected %q", tt.speed, result, tt.expected)
			}
		})
	}
}

func TestFormatSize(t *testing.T) {
	tests := []struct {
		name     string
		size     int64
		expected string
	}{
		{
			name:     "negative size",
			size:     -100,
			expected: "——",
		},
		{
			name:     "zero size",
			size:     0,
			expected: "0.00 KB",
		},
		{
			name:     "small size",
			size:     500,
			expected: "500.00 KB",
		},
		{
			name:     "KB size",
			size:     2048,
			expected: "2.00 KB",
		},
		{
			name:     "MB size",
			size:     5 * MB,
			expected: "5.00 MB",
		},
		{
			name:     "GB size",
			size:     3 * GB,
			expected: "3.00 GB",
		},
		{
			name:     "TB size",
			size:     2 * TB,
			expected: "2.00 TB",
		},
		{
			name:     "PB size",
			size:     1 * PB,
			expected: "1.00 PB",
		},
		{
			name:     "EB size",
			size:     1 * EB,
			expected: "1.00 EB",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := FormatSize(tt.size)
			if result != tt.expected {
				t.Errorf("FormatSize(%d) = %q, expected %q", tt.size, result, tt.expected)
			}
		})
	}
}

func TestFormat2b(t *testing.T) {
	tests := []struct {
		name     string
		size     int64
		expected string
	}{
		{
			name:     "negative size",
			size:     -100,
			expected: "——",
		},
		{
			name:     "zero size",
			size:     0,
			expected: "0.00b",
		},
		{
			name:     "small size",
			size:     500,
			expected: "500.00b",
		},
		{
			name:     "Kb size",
			size:     int64(2 * Kb),
			expected: "2.00Kb",
		},
		{
			name:     "Mb size",
			size:     int64(5 * Mb),
			expected: "5.00Mb",
		},
		{
			name:     "Gb size",
			size:     int64(3 * Gb),
			expected: "3.00Gb",
		},
		{
			name:     "Tb size",
			size:     int64(2 * Tb),
			expected: "2.00Tb",
		},
		{
			name:     "Pb size",
			size:     int64(1 * Pb),
			expected: "1.00Pb",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Format2b(tt.size)
			if result != tt.expected {
				t.Errorf("Format2b(%d) = %q, expected %q", tt.size, result, tt.expected)
			}
		})
	}
}

func TestFormat2B(t *testing.T) {
	tests := []struct {
		name     string
		size     int64
		expected string
	}{
		{
			name:     "negative size",
			size:     -100,
			expected: "——",
		},
		{
			name:     "zero size",
			size:     0,
			expected: "0.00 KB",
		},
		{
			name:     "small size",
			size:     500,
			expected: "500.00 KB",
		},
		{
			name:     "KB size",
			size:     2 * KB,
			expected: "2.00 KB",
		},
		{
			name:     "MB size",
			size:     5 * MB,
			expected: "5.00 MB",
		},
		{
			name:     "GB size",
			size:     3 * GB,
			expected: "3.00 GB",
		},
		{
			name:     "TB size",
			size:     2 * TB,
			expected: "2.00 TB",
		},
		{
			name:     "PB size",
			size:     1 * PB,
			expected: "1.00 PB",
		},
		{
			name:     "EB size",
			size:     1 * EB,
			expected: "1.00 EB",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Format2B(tt.size)
			if result != tt.expected {
				t.Errorf("Format2B(%d) = %q, expected %q", tt.size, result, tt.expected)
			}
		})
	}
}

// Test constants
func TestConstants(t *testing.T) {
	// Test Byte constants
	if Byte != 1 {
		t.Errorf("Byte = %d, expected 1", Byte)
	}
	if KB != 1024 {
		t.Errorf("KB = %d, expected 1024", KB)
	}
	if MB != 1024*1024 {
		t.Errorf("MB = %d, expected %d", MB, 1024*1024)
	}
	
	// Test bit constants
	if Bit != 8 {
		t.Errorf("Bit = %d, expected 8", Bit)
	}
	if Kb != 8*1024 {
		t.Errorf("Kb = %d, expected %d", Kb, 8*1024)
	}
}

// Benchmark tests
func BenchmarkDurationYearMonthDay(b *testing.B) {
	duration := 400*24*time.Hour + 35*24*time.Hour + 5*24*time.Hour
	for i := 0; i < b.N; i++ {
		DurationYearMonthDay(duration)
	}
}

func BenchmarkFormat2B(b *testing.B) {
	size := int64(5 * GB)
	for i := 0; i < b.N; i++ {
		Format2B(size)
	}
}

func BenchmarkFormat2bps(b *testing.B) {
	speed := float64(1000000)
	for i := 0; i < b.N; i++ {
		Format2bps(speed)
	}
}