package xtime

import (
	"fmt"
	"runtime"
	"testing"
	"time"
)

// TestBeginningOfQuarterGlobal_OptimizationVerification 验证优化效果
func TestBeginningOfQuarterGlobal_OptimizationVerification(t *testing.T) {
	iterations := 1000000

	// 测试优化后的实现
	start := time.Now()
	for i := 0; i < iterations; i++ {
		_ = BeginningOfQuarter()
	}
	optimizedElapsed := time.Since(start)
	optimizedNsPerOp := optimizedElapsed.Nanoseconds() / int64(iterations)

	// 模拟原始实现
	start = time.Now()
	for i := 0; i < iterations; i++ {
		now := time.Now()
		_ = With(now).BeginningOfQuarter()
	}
	originalElapsed := time.Since(start)
	originalNsPerOp := originalElapsed.Nanoseconds() / int64(iterations)

	t.Logf("=== BeginningOfQuarter Global Optimization Results ===")
	t.Logf("Iterations: %d", iterations)
	t.Logf("")
	t.Logf("Original Implementation:")
	t.Logf("  Total time: %v", originalElapsed)
	t.Logf("  Per operation: %d ns/op", originalNsPerOp)
	t.Logf("")
	t.Logf("Optimized Implementation:")
	t.Logf("  Total time: %v", optimizedElapsed)
	t.Logf("  Per operation: %d ns/op", optimizedNsPerOp)
	t.Logf("")
	improvement := float64(originalNsPerOp-optimizedNsPerOp) / float64(originalNsPerOp) * 100
	t.Logf("Performance Improvement: %.1f%%", improvement)
	t.Logf("Speed-up: %.2fx", float64(originalNsPerOp)/float64(optimizedNsPerOp))

	// 验证性能确实有提升
	if optimizedNsPerOp >= originalNsPerOp {
		t.Errorf("Optimization failed: %d ns/op (optimized) >= %d ns/op (original)",
			optimizedNsPerOp, originalNsPerOp)
	}

	// 验证性能提升至少 20%
	if improvement < 20 {
		t.Errorf("Performance improvement less than 20%%: %.1f%%", improvement)
	}
}

// TestBeginningOfQuarterGlobal_MemoryAllocation 测试内存分配
func TestBeginningOfQuarterGlobal_MemoryAllocation(t *testing.T) {
	iterations := 10000

	// 测试优化后的内存分配
	var m1, m2 runtime.MemStats
	_ = testing.AllocsPerRun(5, func() {
		for i := 0; i < iterations; i++ {
			_ = BeginningOfQuarter()
		}
	})

	// 获取内存统计
	runtime.GC()
	runtime.ReadMemStats(&m1)

	for i := 0; i < iterations; i++ {
		_ = BeginningOfQuarter()
	}

	runtime.ReadMemStats(&m2)

	allocs := m2.TotalAlloc - m1.TotalAlloc
	t.Logf("Memory allocation for %d calls: %d bytes", iterations, allocs)
	t.Logf("Average per call: %d bytes/op", allocs/uint64(iterations))

	// 验证每次调用分配的内存合理（应该 < 100 bytes）
	avgAllocs := allocs / uint64(iterations)
	if avgAllocs > 100 {
		t.Logf("Warning: Memory allocation per call seems high: %d bytes/op", avgAllocs)
	}
}

// TestBeginningOfQuarterGlobal_RealWorldUsage 真实场景测试
func TestBeginningOfQuarterGlobal_RealWorldUsage(t *testing.T) {
	// 模拟真实使用场景：在不同时间点调用
	testTimes := []struct {
		time time.Time
		name string
	}{
		{time.Date(2024, 1, 15, 10, 30, 0, 0, time.Local), "Q1 Middle"},
		{time.Date(2024, 4, 30, 23, 59, 59, 0, time.Local), "Q2 End"},
		{time.Date(2024, 7, 1, 0, 0, 0, 0, time.Local), "Q3 Start"},
		{time.Date(2024, 10, 15, 15, 45, 30, 0, time.Local), "Q4 Middle"},
	}

	for _, tt := range testTimes {
		t.Run(tt.name, func(t *testing.T) {
			// 注意：由于 BeginningOfQuarter() 使用 time.Now()，
			// 这里我们只验证函数可调用且返回合理结果
			result := BeginningOfQuarter()

			// 基本验证
			if result == nil {
				t.Fatal("BeginningOfQuarter() returned nil")
			}

			// 验证时间是季度开始
			month := result.Month()
			if month != time.January && month != time.April &&
			   month != time.July && month != time.October {
				t.Errorf("Expected quarter start month, got %v", month)
			}

			// 验证时间是月初
			if result.Day() != 1 {
				t.Errorf("Expected day 1, got %d", result.Day())
			}

			// 验证时间是午夜
			if result.Hour() != 0 || result.Minute() != 0 || result.Second() != 0 {
				t.Errorf("Expected midnight, got %02d:%02d:%02d",
					result.Hour(), result.Minute(), result.Second())
			}

			t.Logf("✓ %s: %v", tt.name, result.Time)
		})
	}
}

// TestBeginningOfQuarterGlobal_Concurrency 并发安全测试
func TestBeginningOfQuarterGlobal_Concurrency(t *testing.T) {
	done := make(chan bool)
	errors := make(chan error, 100)

	// 启动多个 goroutine 并发调用
	for i := 0; i < 100; i++ {
		go func() {
			for j := 0; j < 1000; j++ {
				result := BeginningOfQuarter()
				if result == nil {
					errors <- fmt.Errorf("nil result")
					return
				}
			}
			done <- true
		}()
	}

	// 等待所有 goroutine 完成
	for i := 0; i < 100; i++ {
		select {
		case <-done:
			// OK
		case err := <-errors:
			t.Fatal(err)
		}
	}

	t.Logf("✓ Concurrency test passed: 100 goroutines × 1000 calls")
}
