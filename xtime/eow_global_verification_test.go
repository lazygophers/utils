package xtime

import (
	"testing"
	"time"
)

// TestEndOfWeekGlobal_Correctness 验证全局 EndOfWeek() 函数正确性
func TestEndOfWeekGlobal_Correctness(t *testing.T) {
	tests := []struct {
		name           string
		year, month, day int
		expectedWeekday time.Weekday
	}{
		{"2024年6月15日 (周六)", 2024, 6, 15, time.Sunday},
		{"2024年6月16日 (周日)", 2024, 6, 16, time.Sunday},
		{"2024年6月17日 (周一)", 2024, 6, 17, time.Sunday},
		{"2024年1月1日 (周一)", 2024, 1, 1, time.Sunday},
		{"2024年12月31日 (周二)", 2024, 12, 31, time.Sunday},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// 使用 With() 创建测试时间
			testTime := With(time.Date(tt.year, time.Month(tt.month), tt.day, 12, 30, 45, 0, time.Local))

			// 调用 EndOfWeek() 方法
			result := testTime.EndOfWeek()

			// 验证结果是周六
			if result.Weekday() != tt.expectedWeekday {
				t.Errorf("EndOfWeek() weekday = %v, want %v", result.Weekday(), tt.expectedWeekday)
			}

			// 验证时间是 23:59:59.999999999
			h, m, s := result.Clock()
			if h != 23 || m != 59 || s != 59 {
				t.Errorf("EndOfWeek() time = %d:%d:%d, want 23:59:59", h, m, s)
			}

			ns := result.Nanosecond()
			if ns != int(time.Second-time.Nanosecond) {
				t.Errorf("EndOfWeek() nanos = %d, want %d", ns, int(time.Second-time.Nanosecond))
			}

			t.Logf("✓ Test: %s PASS - Result: %s", tt.name, result.Format("2006-01-02 15:04:05.999999999"))
		})
	}
}

// TestEndOfWeekGlobal_RealTime 验证全局函数在真实时间下的行为
func TestEndOfWeekGlobal_RealTime(t *testing.T) {
	result := EndOfWeek()

	// 验证结果是周日
	if result.Weekday() != time.Sunday {
		t.Errorf("EndOfWeek() weekday = %v, want %v", result.Weekday(), time.Sunday)
	}

	// 验证时间是 23:59:59.999999999
	h, m, s := result.Clock()
	if h != 23 || m != 59 || s != 59 {
		t.Errorf("EndOfWeek() time = %d:%d:%d, want 23:59:59", h, m, s)
	}

	ns := result.Nanosecond()
	if ns != int(time.Second-time.Nanosecond) {
		t.Errorf("EndOfWeek() nanos = %d, want %d", ns, int(time.Second-time.Nanosecond))
	}

	// 验证 Config 存在
	if result.Config == nil {
		t.Error("EndOfWeek() Config is nil")
	}

	t.Logf("✓ Current time EndOfWeek: %s", result.Format("2006-01-02 15:04:05.999999999"))
}

// TestEndOfWeekGlobal_Performance 性能测试
func TestEndOfWeekGlobal_Performance(t *testing.T) {
	iterations := 1000000

	start := time.Now()
	for i := 0; i < iterations; i++ {
		_ = EndOfWeek()
	}
	elapsed := time.Since(start)

	avgTime := elapsed.Nanoseconds() / int64(iterations)

	t.Logf("Average time per call: %d ns/op", avgTime)
	t.Logf("Total time for %d calls: %v", iterations, elapsed)

	// 验证性能 < 100 ns/op
	if avgTime > 100 {
		t.Errorf("Performance too slow: %d ns/op, want < 100 ns/op", avgTime)
	}
}

// Benchmark_EndOfWeekGlobal_Optimized 优化后的基准测试
func Benchmark_EndOfWeekGlobal_Optimized(b *testing.B) {
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_ = EndOfWeek()
	}
}
