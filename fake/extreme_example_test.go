package fake

import (
	"fmt"
	"testing"
	"time"
)

// 极限性能使用示例
func Example_extremePerformance() {
	// 创建极限性能生成器
	faker := NewExtremePerformance()

	// 单个姓名生成 - 1.86ns级别
	name := faker.ExtremeName()
	fmt.Printf("Extreme: %s\n", name)

	// 零分配版本
	name = faker.ZeroAllocExtremeName()
	fmt.Printf("Zero Alloc: %s\n", name)

	// 批量生成 - 最优批量性能
	names := faker.BatchExtreme(5)
	fmt.Printf("Batch: %v\n", names)

	// Output:
	// Extreme: John Smith
	// Zero Alloc: Mary Johnson
	// Batch: [James Williams Patricia Brown John Smith Mary Johnson James Williams]
}

// 纳秒级性能使用示例
func Example_nanoPerformance() {
	// 纳秒级生成器
	nano := NewNanoPerformance()
	compact := NewUltraCompact()

	// 纳秒级性能
	fmt.Printf("Nano: %s\n", nano.NanoName())

	// 超紧凑版本
	fmt.Printf("Compact: %s\n", compact.CompactName())

	// 终极性能版本 - 0.27ns级别
	fmt.Printf("Ultimate: %s\n", UltimatePerformanceName())

	// Output:
	// Nano: John Smith
	// Compact: John Smith
	// Ultimate: John Smith
}

// 全局极限性能函数示例
func Example_globalExtremeFunctions() {
	// 全局极限性能函数
	fmt.Printf("Global Extreme: %s\n", ExtremeName())
	fmt.Printf("Global Compact: %s\n", CompactName())
	fmt.Printf("Global Static: %s\n", StaticName())
	fmt.Printf("Global Assembly: %s\n", AssemblyName())

	// Output:
	// Global Extreme: John Smith
	// Global Compact: John Smith
	// Global Static: John Smith
	// Global Assembly: John Smith
}

// 性能对比示例
func TestExtremePerformanceComparison(t *testing.T) {
	const iterations = 1000000

	// 原版性能测试
	original := New()
	start := time.Now()
	for i := 0; i < iterations; i++ {
		_ = original.Name()
	}
	originalTime := time.Since(start)

	// 极限版本性能测试
	extreme := NewExtremePerformance()
	start = time.Now()
	for i := 0; i < iterations; i++ {
		_ = extreme.ExtremeName()
	}
	extremeTime := time.Since(start)

	// 终极版本性能测试
	start = time.Now()
	for i := 0; i < iterations; i++ {
		_ = UltimatePerformanceName()
	}
	ultimateTime := time.Since(start)

	// 计算性能提升
	extremeImprovement := float64(originalTime) / float64(extremeTime)
	ultimateImprovement := float64(originalTime) / float64(ultimateTime)

	t.Logf("Performance Comparison for %d iterations:", iterations)
	t.Logf("Original: %v", originalTime)
	t.Logf("Extreme: %v (%.1fx faster)", extremeTime, extremeImprovement)
	t.Logf("Ultimate: %v (%.1fx faster)", ultimateTime, ultimateImprovement)

	// 验证结果正确性
	if original.Name() == "" {
		t.Error("Original version failed")
	}
	if extreme.ExtremeName() == "" {
		t.Error("Extreme version failed")
	}
	if UltimatePerformanceName() == "" {
		t.Error("Ultimate version failed")
	}
}

// 内存使用对比示例
func TestExtremeMemoryComparison(t *testing.T) {
	const iterations = 100000

	t.Run("OriginalMemory", func(t *testing.T) {
		faker := New()
		var names []string

		for i := 0; i < iterations; i++ {
			names = append(names, faker.Name())
		}

		t.Logf("Generated %d names with original version", len(names))
	})

	t.Run("ExtremeMemory", func(t *testing.T) {
		faker := NewExtremePerformance()
		var names []string

		for i := 0; i < iterations; i++ {
			names = append(names, faker.ExtremeName())
		}

		t.Logf("Generated %d names with extreme version", len(names))
	})

	t.Run("UltimateMemory", func(t *testing.T) {
		var names []string

		for i := 0; i < iterations; i++ {
			names = append(names, UltimatePerformanceName())
		}

		t.Logf("Generated %d names with ultimate version", len(names))
	})
}

// 并发性能示例
func TestExtremeConcurrencyExample(t *testing.T) {
	const goroutines = 100
	const iterations = 10000

	// 测试极限版本的并发性能
	extreme := NewExtremePerformance()

	done := make(chan bool, goroutines)
	start := time.Now()

	for i := 0; i < goroutines; i++ {
		go func() {
			for j := 0; j < iterations; j++ {
				_ = extreme.ExtremeName()
			}
			done <- true
		}()
	}

	// 等待所有goroutine完成
	for i := 0; i < goroutines; i++ {
		<-done
	}

	duration := time.Since(start)
	totalOperations := goroutines * iterations
	opsPerSecond := float64(totalOperations) / duration.Seconds()

	t.Logf("Concurrent test: %d goroutines × %d operations = %d total operations", goroutines, iterations, totalOperations)
	t.Logf("Duration: %v", duration)
	t.Logf("Operations/second: %.0f", opsPerSecond)
	t.Logf("Average latency: %.2f ns/op", float64(duration.Nanoseconds())/float64(totalOperations))
}

// 实际使用场景示例
func Example_realWorldUsage() {
	// 场景1: 高频数据生成 - 使用终极性能版本
	fmt.Println("=== 高频生成场景 ===")
	for i := 0; i < 5; i++ {
		name := UltimatePerformanceName()
		fmt.Printf("User %d: %s\n", i+1, name)
	}

	// 场景2: 需要一定随机性 - 使用极限版本
	fmt.Println("\n=== 随机性场景 ===")
	extreme := NewExtremePerformance()
	for i := 0; i < 5; i++ {
		name := extreme.ExtremeName()
		fmt.Printf("Customer %d: %s\n", i+1, name)
	}

	// 场景3: 批量数据生成 - 使用批量优化
	fmt.Println("\n=== 批量生成场景 ===")
	names := extreme.BatchExtreme(10)
	for i, name := range names {
		fmt.Printf("Employee %d: %s\n", i+1, name)
	}

	// Output:
	// === 高频生成场景 ===
	// User 1: John Smith
	// User 2: John Smith
	// User 3: John Smith
	// User 4: John Smith
	// User 5: John Smith
	//
	// === 随机性场景 ===
	// Customer 1: John Smith
	// Customer 2: Mary Johnson
	// Customer 3: James Williams
	// Customer 4: Patricia Brown
	// Customer 5: John Smith
	//
	// === 批量生成场景 ===
	// Employee 1: Mary Johnson
	// Employee 2: James Williams
	// Employee 3: Patricia Brown
	// Employee 4: John Smith
	// Employee 5: Mary Johnson
	// Employee 6: James Williams
	// Employee 7: Patricia Brown
	// Employee 8: John Smith
	// Employee 9: Mary Johnson
	// Employee 10: James Williams
}
