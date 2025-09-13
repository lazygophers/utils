package unit

import (
	"testing"
	"time"
)

// TestFormat2bpsMissingBranches tests the missing coverage branches in Format2bps
func TestFormat2bpsMissingBranches(t *testing.T) {
	tests := []struct {
		name     string
		speed    float64
		expected string
	}{
		{
			name:     "Gbps range",
			speed:    float64(Gb/8 + 1000),     // Just over Gb/8 threshold
			expected: "1.00 Gbps",
		},
		{
			name:     "Tbps range",
			speed:    float64(Tb/8 + 100000),   // Just over Tb/8 threshold
			expected: "1.00 Tbps",
		},
		{
			name:     "Pbps range",
			speed:    float64(Pb/8 + 10000000), // Just over Pb/8 threshold
			expected: "1.00 Pbps",
		},
		{
			name:     "boundary_at_Gb",
			speed:    float64(Gb / 8),          // Exactly at Gb/8 boundary
			expected: "1.00 Gbps",
		},
		{
			name:     "boundary_at_Tb",
			speed:    float64(Tb / 8),          // Exactly at Tb/8 boundary
			expected: "1.00 Tbps",
		},
		{
			name:     "boundary_at_Pb",
			speed:    float64(Pb / 8),          // Exactly at Pb/8 boundary
			expected: "1.00 Pbps",
		},
		{
			name:     "boundary_at_Eb",
			speed:    float64(Eb / 8),          // Exactly at Eb/8 boundary
			expected: "1.00 Ebps",
		},
		{
			name:     "just_under_Gb",
			speed:    float64(Gb/8 - 1),        // Just under Gb/8 threshold
			expected: "1024.00 Mbps",
		},
		{
			name:     "just_under_Tb", 
			speed:    float64(Tb/8 - 1000),     // Just under Tb/8 threshold
			expected: "1024.00 Gbps",
		},
		{
			name:     "just_under_Pb",
			speed:    float64(Pb/8 - 100000),   // Just under Pb/8 threshold  
			expected: "1024.00 Tbps",
		},
		{
			name:     "just_under_Eb",
			speed:    float64(Eb/8 - 10000000), // Just under Eb/8 threshold
			expected: "1024.00 Pbps",
		},
		{
			name:     "extremely_large_value",
			speed:    float64(Eb),              // Very large value beyond Eb/8
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

// TestExtremeValues tests edge cases with very small and very large values
func TestExtremeValues(t *testing.T) {
	tests := []struct {
		name     string
		speed    float64
		desc     string
	}{
		{
			name:     "tiny_positive",
			speed:    0.0001,
			desc:     "Very small positive value",
		},
		{
			name:     "just_above_zero",
			speed:    0.000001,
			desc:     "Just above zero",
		},
		{
			name:     "fractional_byte",
			speed:    0.5,
			desc:     "Half byte per second",
		},
		{
			name:     "one_byte",
			speed:    1.0,
			desc:     "One byte per second",
		},
		{
			name:     "huge_value",
			speed:    float64(Eb * 10),
			desc:     "Extremely large value",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Format2bps(tt.speed)
			// Just ensure it doesn't panic and returns a non-empty result
			if result == "" {
				t.Errorf("Format2bps(%f) returned empty string", tt.speed)
			}
			t.Logf("%s: Format2bps(%f) = %q", tt.desc, tt.speed, result)
		})
	}
}

// TestAllNetworkFormattingFunctions ensures all network formatting functions are covered
func TestAllNetworkFormattingFunctions(t *testing.T) {
	testValues := []float64{
		0,
		-1,
		0.5,
		1,
		128,     // 1 Kbps equivalent
		1024,    // 8 Kbps  
		131072,  // 1 Mbps equivalent
		float64(Gb / 8),
		float64(Tb / 8),
		float64(Pb / 8),
		float64(Eb / 8),
		float64(Eb),
	}

	for _, value := range testValues {
		t.Run("comprehensive_test", func(t *testing.T) {
			// Test all network formatting functions
			r1 := FormatSpeed(value)
			r2 := Format2bps(value)
			r3 := Format2Bs(value)
			r4 := FormatSize(int64(value))
			r5 := Format2b(int64(value))
			r6 := Format2B(int64(value))

			// Ensure none return empty strings for valid inputs
			if r1 == "" || r2 == "" || r3 == "" || r4 == "" || r5 == "" || r6 == "" {
				t.Errorf("Empty result for value %f: FormatSpeed=%q, Format2bps=%q, Format2Bs=%q, FormatSize=%q, Format2b=%q, Format2B=%q",
					value, r1, r2, r3, r4, r5, r6)
			}

			t.Logf("Value %f: Speed=%q, bps=%q, Bs=%q, Size=%q, b=%q, B=%q",
				value, r1, r2, r3, r4, r5, r6)
		})
	}
}

// TestTimeFormattingCompleteCoverage ensures all time formatting functions are well tested
func TestTimeFormattingCompleteCoverage(t *testing.T) {
	// Test various duration values
	testDurations := []int64{
		0,
		1,
		59,
		60,
		3599,
		3600,
		86399,
		86400,
		2592000,   // 30 days
		31536000,  // 365 days
		63072000,  // 2 years
	}

	for _, duration := range testDurations {
		t.Run("time_formatting", func(t *testing.T) {
			// Test all time formatting functions
			r1 := DurationYearMonthDay(time.Duration(duration) * time.Second)
			r2 := DurationMonthDayHour(time.Duration(duration) * time.Second)
			r3 := DurationMinuteSecond(time.Duration(duration) * time.Second)
			r4 := DurationYearMonthDayHourMinuteSecond(time.Duration(duration) * time.Second)

			// Ensure none return empty strings
			if r1 == "" || r2 == "" || r3 == "" || r4 == "" {
				t.Errorf("Empty result for duration %d: YMD=%q, MDH=%q, MS=%q, YMDHMS=%q",
					duration, r1, r2, r3, r4)
			}

			t.Logf("Duration %d: YMD=%q, MDH=%q, MS=%q, YMDHMS=%q",
				duration, r1, r2, r3, r4)
		})
	}
}

// TestTimeFormattingFunctions tests time formatting functions
func TestTimeFormattingFunctions(t *testing.T) {
	// Test TimeYearMonthDayHourMinute and TimeYearMonthDayHourMinuteSecond
	testTime := time.Unix(1609459200, 0) // 2021-01-01 00:00:00 UTC
	
	result1 := TimeYearMonthDayHourMinute(testTime)
	result2 := TimeYearMonthDayHourMinuteSecond(testTime)
	
	if result1 == "" {
		t.Error("TimeYearMonthDayHourMinute returned empty string")
	}
	if result2 == "" {
		t.Error("TimeYearMonthDayHourMinuteSecond returned empty string")
	}
	
	t.Logf("TimeYearMonthDayHourMinute(%v) = %q", testTime, result1)
	t.Logf("TimeYearMonthDayHourMinuteSecond(%v) = %q", testTime, result2)
}