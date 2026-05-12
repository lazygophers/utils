package xtime

import (
	"fmt"
	"testing"
	"time"
)

// TestEndOfYearGlobal_OptimizationVerification 验证 EndOfYear 全局函数优化效果
func TestEndOfYearGlobal_OptimizationVerification(t *testing.T) {
	const iterations = 1000000

	t.Log("=== EndOfYear Global Optimization Results ===")
	t.Logf("Iterations: %d", iterations)
	t.Log("")

	// 测试原始实现
	originalFunc := func() *Time {
		return With(time.Now()).EndOfYear()
	}

	start := time.Now()
	for i := 0; i < iterations; i++ {
		_ = originalFunc()
	}
	originalTime := time.Since(start)

	t.Logf("Original Implementation:")
	t.Logf("  Total time: %v", originalTime)
	t.Logf("  Avg time: %.2f ns/op", float64(originalTime.Nanoseconds())/float64(iterations))
	t.Log("")

	// 测试优化实现
	optimizedFunc := func() *Time {
		now := time.Now()
		year := now.Year()
		return &Time{
			Time:   time.Date(year+1, time.January, 0, 23, 59, 59, 999999999, now.Location()),
			Config: nil,
		}
	}

	start = time.Now()
	for i := 0; i < iterations; i++ {
		_ = optimizedFunc()
	}
	optimizedTime := time.Since(start)

	t.Logf("Optimized Implementation:")
	t.Logf("  Total time: %v", optimizedTime)
	t.Logf("  Avg time: %.2f ns/op", float64(optimizedTime.Nanoseconds())/float64(iterations))
	t.Log("")

	improvement := float64(originalTime-optimizedTime) / float64(originalTime) * 100

	t.Logf("Performance Improvement: %.2f%%", improvement)
	t.Log("")

	if optimizedTime >= originalTime {
		t.Errorf("Optimized version should be faster: %v >= %v", optimizedTime, originalTime)
	}

	// 验证结果正确性
	now := time.Now()
	originalResult := With(now).EndOfYear()
	optimizedResult := optimizedFunc()

	if originalResult.Time.UnixNano() != optimizedResult.Time.UnixNano() {
		t.Logf("Original: %v", originalResult.Time)
		t.Logf("Optimized: %v", optimizedResult.Time)
		t.Errorf("Results don't match within same second")
	}

	// 验证年份、月份、日期相同
	if originalResult.Time.Year() != optimizedResult.Time.Year() ||
		originalResult.Time.Month() != optimizedResult.Time.Month() ||
		originalResult.Time.Day() != optimizedResult.Time.Day() {
		t.Errorf("Date components don't match")
	}

	// 验证时间都是23:59:59.999999999
	h, m, s := optimizedResult.Time.Clock()
	ns := optimizedResult.Time.Nanosecond()
	if h != 23 || m != 59 || s != 59 || ns != 999999999 {
		t.Errorf("Time should be 23:59:59.999999999, got %d:%d:%d.%d", h, m, s, ns)
	}
}

// TestEndOfYearGlobal_Correctness 验证 EndOfYear 全局函数正确性
func TestEndOfYearGlobal_Correctness(t *testing.T) {
	// 测试不同年份
	testCases := []struct {
		year        int
		expectedDay int
		expectedMon int
	}{
		{2024, 31, 12}, // 闰年
		{2023, 31, 12},
		{2020, 31, 12}, // 闰年
		{2025, 31, 12},
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("Year_%d", tc.year), func(t *testing.T) {
			testTime := time.Date(tc.year, time.July, 15, 12, 0, 0, 0, time.Local)
			result := With(testTime).EndOfYear()

			if result.Time.Year() != tc.year {
				t.Errorf("Expected year %d, got %d", tc.year, result.Time.Year())
			}
			if result.Time.Month() != time.Month(tc.expectedMon) {
				t.Errorf("Expected month %d, got %d", tc.expectedMon, result.Time.Month())
			}
			if result.Time.Day() != tc.expectedDay {
				t.Errorf("Expected day %d, got %d", tc.expectedDay, result.Time.Day())
			}

			// 验证时间是23:59:59.999999999
			h, m, s := result.Time.Clock()
			ns := result.Time.Nanosecond()
			if h != 23 || m != 59 || s != 59 || ns != 999999999 {
				t.Errorf("Expected 23:59:59.999999999, got %d:%d:%d.%d", h, m, s, ns)
			}
		})
	}

	// 测试跨年
	t.Run("CrossYearBoundary", func(t *testing.T) {
		// 2024-12-31 23:59:59
		testTime := time.Date(2024, time.December, 31, 23, 59, 59, 0, time.Local)
		result := With(testTime).EndOfYear()

		// 应该返回2024年结束，不是2025年
		if result.Time.Year() != 2024 {
			t.Errorf("Expected year 2024, got %d", result.Time.Year())
		}
		if result.Time.Month() != time.December {
			t.Errorf("Expected December, got %v", result.Time.Month())
		}
		if result.Time.Day() != 31 {
			t.Errorf("Expected day 31, got %d", result.Time.Day())
		}
	})
}

// TestEndOfYearGlobal_GlobalFunction 测试全局函数正确性
func TestEndOfYearGlobal_GlobalFunction(t *testing.T) {
	result := EndOfYear()

	if result == nil {
		t.Fatal("EndOfYear() should not return nil")
	}

	if result.Time.IsZero() {
		t.Error("EndOfYear() should not return zero time")
	}

	// 验证返回的是当前年的结束时间
	currentYear := time.Now().Year()
	if result.Time.Year() != currentYear {
		t.Errorf("Expected year %d, got %d", currentYear, result.Time.Year())
	}

	if result.Time.Month() != time.December {
		t.Errorf("Expected December, got %v", result.Time.Month())
	}

	if result.Time.Day() != 31 {
		t.Errorf("Expected day 31, got %d", result.Time.Day())
	}

	// 验证时间是23:59:59.999999999
	h, m, s := result.Time.Clock()
	ns := result.Time.Nanosecond()
	if h != 23 || m != 59 || s != 59 || ns != 999999999 {
		t.Errorf("Expected 23:59:59.999999999, got %d:%d:%d.%d", h, m, s, ns)
	}
}
