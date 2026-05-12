package xtime

import (
	"fmt"
	"testing"
	"time"
)

// TestEndOfYearGlobalOptimization 验证 EndOfYear 全局函数优化效果
func TestEndOfYearGlobalOptimization(t *testing.T) {
	const iterations = 100000

	type result struct {
		name        string
		totalNs     int64
		avgNs       float64
		improvement float64
	}

	var results []result

	// Baseline: 当前实现
	{
		start := time.Now()
		for i := 0; i < iterations; i++ {
			_ = EndOfYear()
		}
		elapsed := time.Since(start)
		results = append(results, result{
			name:    "Current",
			totalNs: elapsed.Nanoseconds(),
			avgNs:   float64(elapsed.Nanoseconds()) / float64(iterations),
		})
	}

	// 变体1: 直接内联优化
	{
		start := time.Now()
		for i := 0; i < iterations; i++ {
			now := time.Now()
			_ = &Time{
				Time:   time.Date(now.Year()+1, time.January, 0, 23, 59, 59, 999999999, now.Location()),
				Config: nil,
			}
		}
		elapsed := time.Since(start)
		results = append(results, result{
			name:    "Optimized",
			totalNs: elapsed.Nanoseconds(),
			avgNs:   float64(elapsed.Nanoseconds()) / float64(iterations),
		})
	}

	// 计算性能提升
	baselineAvg := results[0].avgNs
	for i := range results {
		results[i].improvement = ((baselineAvg - results[i].avgNs) / baselineAvg) * 100
	}

	// 输出结果
	fmt.Println("=== EndOfYear Global Optimization Results ===")
	fmt.Printf("Iterations: %d\n\n", iterations)

	fmt.Println("| Variant    | Total Time | Avg Time/op | Improvement |")
	fmt.Println("|------------|------------|-------------|-------------|")
	for _, r := range results {
		fmt.Printf("| %-10s | %10v | %11.2f ns | %10.2f%% |\n",
			r.name, time.Duration(r.totalNs), r.avgNs, r.improvement)
	}

	// 验证优化版本确实更快
	if results[1].avgNs >= results[0].avgNs {
		t.Errorf("Optimized version (%.2f ns/op) should be faster than current (%.2f ns/op)",
			results[1].avgNs, results[0].avgNs)
	}
}
