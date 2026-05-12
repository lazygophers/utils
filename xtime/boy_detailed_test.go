package xtime

import (
	"fmt"
	"testing"
	"time"
)

// TestBeginningOfYearGlobal_DetailedPerformance 详细的性能测试
func TestBeginningOfYearGlobal_DetailedPerformance(t *testing.T) {
	iterations := 10000000

	// 预热
	for i := 0; i < 1000; i++ {
		_ = BeginningOfYear()
	}

	// 测试优化后实现
	start := time.Now()
	for i := 0; i < iterations; i++ {
		_ = BeginningOfYear()
	}
	optimizedDuration := time.Since(start)
	avgTime := optimizedDuration.Nanoseconds() / int64(iterations)

	fmt.Printf("\n=== BeginningOfYear Global Detailed Performance ===\n")
	fmt.Printf("Iterations: %d\n", iterations)
	fmt.Printf("Total time: %v\n", optimizedDuration)
	fmt.Printf("Average time per call: %d ns/op\n", avgTime)

	// 性能阈值：应该 < 100 ns/op
	if avgTime > 100 {
		t.Errorf("Performance too slow: %d ns/op, want < 100 ns/op", avgTime)
	} else {
		fmt.Printf("Performance test PASSED\n")
	}
}

// TestBeginningOfYearGlobal_CorrectnessInDetail 详细正确性测试
func TestBeginningOfYearGlobal_CorrectnessInDetail(t *testing.T) {
	testCases := []struct {
		name string
	}{
		{"Test in current year"},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := BeginningOfYear()
			now := time.Now()
			expected := time.Date(now.Year(), time.January, 1, 0, 0, 0, 0, now.Location())

			if result.Time.Unix() != expected.Unix() {
				t.Errorf("BeginningOfYear() = %v, want %v", result.Time, expected)
			}

			// 验证时区
			if result.Time.Location().String() != now.Location().String() {
				t.Errorf("Location mismatch: got %v, want %v", result.Time.Location(), now.Location())
			}

			fmt.Printf("Test: %s PASS\n", tc.name)
		})
	}
}
