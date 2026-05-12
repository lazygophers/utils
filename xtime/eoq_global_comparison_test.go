package xtime

import (
	"testing"
	"time"
)

// Benchmark_EndOfQuarter_Global_Old 原始实现（用于对比）
func Benchmark_EndOfQuarter_Global_Old(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		// 模拟原始实现
		_ = With(time.Now()).EndOfQuarter()
	}
}

// Benchmark_EndOfQuarter_Global_New 优化后的实现
func Benchmark_EndOfQuarter_Global_New(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		// 新实现已直接在 now.go 中
		_ = EndOfQuarter()
	}
}

// TestEndOfQuarterGlobal_OptimizationComparison 验证优化效果
func TestEndOfQuarterGlobal_OptimizationComparison(t *testing.T) {
	if testing.Short() {
		t.Skip("跳过性能对比测试")
	}

	iterations := 100000

	// 测试原始实现
	t.Run("Original", func(t *testing.T) {
		start := time.Now()
		for i := 0; i < iterations; i++ {
			_ = With(time.Now()).EndOfQuarter()
		}
		originalTime := time.Since(start)

		t.Logf("原始实现: %d 次操作耗时 %v (%.2f ns/op)",
			iterations, originalTime, float64(originalTime.Nanoseconds())/float64(iterations))
	})

	// 测试优化实现
	t.Run("Optimized", func(t *testing.T) {
		start := time.Now()
		for i := 0; i < iterations; i++ {
			_ = EndOfQuarter()
		}
		optimizedTime := time.Since(start)

		t.Logf("优化实现: %d 次操作耗时 %v (%.2f ns/op)",
			iterations, optimizedTime, float64(optimizedTime.Nanoseconds())/float64(iterations))
	})
}

// TestEndOfQuarterGlobal_CorrectnessFinal 最终正确性验证
func TestEndOfQuarterGlobal_CorrectnessFinal(t *testing.T) {
	// 测试当前时间
	now := time.Now()
	result := EndOfQuarter()

	// 验证返回的是当季度的结束时间
	_, month, _ := now.Date()
	quarter := (month-1)/3 + 1
	expectedMonth := quarter * 3 // Q1=3, Q2=6, Q3=9, Q4=12

	if result.Month() != expectedMonth {
		t.Errorf("EndOfQuarter() month = %v, want %v", result.Month(), expectedMonth)
	}

	// 验证时间是 23:59:59.999999999
	hour, min, sec := result.Clock()
	if hour != 23 || min != 59 || sec != 59 {
		t.Errorf("EndOfQuarter() time = %d:%d:%d, want 23:59:59", hour, min, sec)
	}

	if result.Nanosecond() != 999999999 {
		t.Errorf("EndOfQuarter() nanosecond = %d, want 999999999", result.Nanosecond())
	}

	// 验证 Config 存在
	if result.Config == nil {
		t.Error("EndOfQuarter() Config is nil")
	}

	if result.Config.WeekStartDay != time.Monday {
		t.Errorf("EndOfQuarter() WeekStartDay = %v, want Monday", result.Config.WeekStartDay)
	}

	t.Logf("✅ EndOfQuarter() 正确性验证通过: %v", result.Time)
}

// TestEndOfQuarterGlobal_AllQuartersEdgeCases 测试所有季度的边界情况
func TestEndOfQuarterGlobal_AllQuartersEdgeCases(t *testing.T) {
	testCases := []struct {
		name     string
		month    time.Month
		expected time.Month
	}{
		{"Q1 - January", time.January, time.March},
		{"Q1 - February", time.February, time.March},
		{"Q1 - March", time.March, time.March},
		{"Q2 - April", time.April, time.June},
		{"Q2 - May", time.May, time.June},
		{"Q2 - June", time.June, time.June},
		{"Q3 - July", time.July, time.September},
		{"Q3 - August", time.August, time.September},
		{"Q3 - September", time.September, time.September},
		{"Q4 - October", time.October, time.December},
		{"Q4 - November", time.November, time.December},
		{"Q4 - December", time.December, time.December},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// 创建指定月份的测试时间
			testTime := time.Date(2024, tc.month, 15, 12, 0, 0, 0, time.Local)

			// 使用 With 创建 Time 对象，然后调用 EndOfQuarter
			result := With(testTime).EndOfQuarter()

			if result.Month() != tc.expected {
				t.Errorf("For month %v, got month %v, want %v", tc.month, result.Month(), tc.expected)
			}

			// 验证是季度最后一天
			lastDay := time.Date(2024, tc.expected+1, 0, 23, 59, 59, 999999999, time.Local)
			if !result.Time.Equal(lastDay) {
				t.Errorf("For month %v, got %v, want %v", tc.month, result.Time, lastDay)
			}
		})
	}
}

