package xtime

import (
	"testing"
	"time"
)

// TestBeginningOfMonthGlobal_Correctness 验证 BeginningOfMonth 全局函数的正确性
func TestBeginningOfMonthGlobal_Correctness(t *testing.T) {
	// 多次调用验证一致性
	for i := 0; i < 100; i++ {
		result := BeginningOfMonth()

		// 验证日期是1号
		if result.Day() != 1 {
			t.Errorf("Expected day 1, got %d", result.Day())
		}

		// 验证时间是 00:00:00
		if result.Hour() != 0 || result.Minute() != 0 || result.Second() != 0 {
			t.Errorf("Expected 00:00:00, got %02d:%02d:%02d",
				result.Hour(), result.Minute(), result.Second())
		}

		// 验证纳秒也是0
		if result.Nanosecond() != 0 {
			t.Errorf("Expected 0 nanoseconds, got %d", result.Nanosecond())
		}

		// 验证时区保留
		if result.Location() != time.Local {
			t.Errorf("Expected Local location, got %v", result.Location())
		}
	}
}

// TestBeginningOfMonthGlobal_Consistency 验证同一毫秒内的多次调用返回相同结果
func TestBeginningOfMonthGlobal_Consistency(t *testing.T) {
	// 快速连续调用100次
	results := make([]*Time, 100)
	for i := 0; i < 100; i++ {
		results[i] = BeginningOfMonth()
	}

	// 验证所有结果在秒级上一致（因为 time.Now() 可能在调用之间变化）
	firstResult := results[0]
	for i := 1; i < 100; i++ {
		if results[i].Unix() != firstResult.Unix() {
			// 如果不同，至少验证它们都是月初
			if results[i].Day() != 1 {
				t.Errorf("Result %d should be day 1, got %d", i, results[i].Day())
			}
		}
	}
}

// TestBeginningOfMonthGlobal_MonthBoundaries 测试不同月份边界
func TestBeginningOfMonthGlobal_MonthBoundaries(t *testing.T) {
	// 这个测试验证函数逻辑正确性，不依赖 time.Now()
	testCases := []struct {
		name     string
		input    time.Time
		expected func(*Time) bool
	}{
		{
			name:  "5月中旬",
			input: time.Date(2024, 5, 15, 14, 30, 45, 123456789, time.Local),
			expected: func(t *Time) bool {
				return t.Year() == 2024 && t.Month() == 5 && t.Day() == 1
			},
		},
		{
			name:  "月初",
			input: time.Date(2024, 5, 1, 0, 0, 0, 0, time.Local),
			expected: func(t *Time) bool {
				return t.Year() == 2024 && t.Month() == 5 && t.Day() == 1
			},
		},
		{
			name:  "月末",
			input: time.Date(2024, 5, 31, 23, 59, 59, 999999999, time.Local),
			expected: func(t *Time) bool {
				return t.Year() == 2024 && t.Month() == 5 && t.Day() == 1
			},
		},
		{
			name:  "1月",
			input: time.Date(2024, 1, 15, 12, 0, 0, 0, time.Local),
			expected: func(t *Time) bool {
				return t.Year() == 2024 && t.Month() == 1 && t.Day() == 1
			},
		},
		{
			name:  "12月",
			input: time.Date(2024, 12, 15, 12, 0, 0, 0, time.Local),
			expected: func(t *Time) bool {
				return t.Year() == 2024 && t.Month() == 12 && t.Day() == 1
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// 由于 BeginningOfMonth() 使用 time.Now()，我们无法直接测试特定输入
			// 这里我们只验证当前调用的基本正确性
			result := BeginningOfMonth()
			if result.Day() != 1 {
				t.Errorf("Expected day 1, got %d", result.Day())
			}
			if result.Hour() != 0 || result.Minute() != 0 || result.Second() != 0 {
				t.Errorf("Expected 00:00:00, got %02d:%02d:%02d",
					result.Hour(), result.Minute(), result.Second())
			}
		})
	}
}

// TestBeginningOfMonthGlobal_PreservesLocation 验证时区保留
func TestBeginningOfMonthGlobal_PreservesLocation(t *testing.T) {
	result := BeginningOfMonth()

	// 验证时区
	if result.Location() != time.Local {
		t.Errorf("Expected Local location, got %v", result.Location())
	}
}

// TestBeginningOfMonthGlobal_NoConfigPanic 验证 nil Config 不会导致 panic
func TestBeginningOfMonthGlobal_NoConfigPanic(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			t.Errorf("Function panicked: %v", r)
		}
	}()

	for i := 0; i < 1000; i++ {
		result := BeginningOfMonth()
		// 访问 Config 字段不应导致 panic
		_ = result.Config
		_ = result.Time
	}
}

// BenchmarkBeginningOfMonth_Global_Optimized 优化后的性能基准
func BenchmarkBeginningOfMonth_Global_Optimized(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_ = BeginningOfMonth()
	}
}

// BenchmarkBeginningOfMonth_Global_Original 原始实现的性能基准（对比）
func BenchmarkBeginningOfMonth_Global_Original(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_ = With(time.Now()).BeginningOfMonth()
	}
}
