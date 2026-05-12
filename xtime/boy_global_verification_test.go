package xtime

import (
	"fmt"
	"runtime"
	"testing"
	"time"
)

// TestBeginningOfYearGlobal_Correctness 验证优化后的正确性
func TestBeginningOfYearGlobal_Correctness(t *testing.T) {
	now := time.Now()
	expected := time.Date(now.Year(), time.January, 1, 0, 0, 0, 0, now.Location())
	result := BeginningOfYear()

	if result.Time.Unix() != expected.Unix() {
		t.Errorf("BeginningOfYear() = %v, want %v", result.Time, expected)
	}

	// 验证时区
	if result.Time.Location().String() != now.Location().String() {
		t.Errorf("Location mismatch: got %v, want %v", result.Time.Location(), now.Location())
	}
}

// TestBeginningOfYearGlobal_Performance 性能测试
func TestBeginningOfYearGlobal_Performance(t *testing.T) {
	iterations := 1000000

	// 测试优化后实现
	start := time.Now()
	for i := 0; i < iterations; i++ {
		_ = BeginningOfYear()
	}
	optimizedDuration := time.Since(start)
	avgTime := optimizedDuration.Nanoseconds() / int64(iterations)

	fmt.Printf("\n=== BeginningOfYear Global Performance ===\n")
	fmt.Printf("Iterations: %d\n", iterations)
	fmt.Printf("Total time: %v\n", optimizedDuration)
	fmt.Printf("Average time per call: %d ns/op\n", avgTime)

	// 性能阈值：应该 < 100 ns/op
	if avgTime > 100 {
		t.Errorf("Performance too slow: %d ns/op, want < 100 ns/op", avgTime)
	}
}

// TestBeginningOfYearGlobal_ZeroAllocation 验证零内存分配
func TestBeginningOfYearGlobal_ZeroAllocation(t *testing.T) {
	iterations := 1000

	// 强制 GC
	runtime.GC()

	allocs := testing.AllocsPerRun(iterations, func() {
		_ = BeginningOfYear()
	})

	fmt.Printf("\n=== BeginningOfYear Global Memory Test ===\n")
	fmt.Printf("Allocations per run: %.2f allocs/op\n", allocs)

	// 验证最小分配（&Time{} 必然有1次分配）
	if allocs > 1.1 {
		t.Errorf("Memory allocation too high: %.2f allocs/op, want ~1", allocs)
	} else {
		fmt.Printf("Minimal allocation test PASSED (1 alloc for &Time{})\n")
	}
}
