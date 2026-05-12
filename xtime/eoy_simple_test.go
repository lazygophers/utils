package xtime

import (
	"testing"
	"time"
)

// Benchmark: 当前实现
func BenchmarkEOY_Current(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_ = With(time.Now()).EndOfYear()
	}
}

// Benchmark: 优化版本 - 直接构造
func BenchmarkEOY_Optimized(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		now := time.Now()
		_ = &Time{
			Time:   time.Date(now.Year()+1, time.January, 0, 23, 59, 59, 999999999, now.Location()),
			Config: nil,
		}
	}
}

// Benchmark: 优化版本2 - 使用变量
func BenchmarkEOY_OptimizedV2(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		now := time.Now()
		year := now.Year()
		_ = &Time{
			Time:   time.Date(year+1, time.January, 0, 23, 59, 59, 999999999, now.Location()),
			Config: nil,
		}
	}
}

// Benchmark: 优化版本3 - AddDate
func BenchmarkEOY_OptimizedV3(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		now := time.Now()
		nextYearStart := time.Date(now.Year()+1, time.January, 1, 0, 0, 0, 0, now.Location())
		_ = &Time{
			Time:   nextYearStart.Add(-time.Nanosecond),
			Config: nil,
		}
	}
}

// 测试验证优化效果
func TestEOYOptimizationEffect(t *testing.T) {
	const iterations = 50000

	// 测试当前实现
	start := time.Now()
	for i := 0; i < iterations; i++ {
		_ = With(time.Now()).EndOfYear()
	}
	currentTime := time.Since(start)

	// 测试优化实现
	start = time.Now()
	for i := 0; i < iterations; i++ {
		now := time.Now()
		_ = &Time{
			Time:   time.Date(now.Year()+1, time.January, 0, 23, 59, 59, 999999999, now.Location()),
			Config: nil,
		}
	}
	optimizedTime := time.Since(start)

	improvement := float64(currentTime-optimizedTime) / float64(currentTime) * 100

	t.Logf("Current: %v for %d iterations (%.2f ns/op)",
		currentTime, iterations, float64(currentTime.Nanoseconds())/float64(iterations))
	t.Logf("Optimized: %v for %d iterations (%.2f ns/op)",
		optimizedTime, iterations, float64(optimizedTime.Nanoseconds())/float64(iterations))
	t.Logf("Improvement: %.2f%%", improvement)

	if optimizedTime >= currentTime {
		t.Errorf("Optimized version should be faster: %v >= %v", optimizedTime, currentTime)
	}

	// 验证结果正确性
	now := time.Now()
	currentResult := With(now).EndOfYear()
	optimizedResult := &Time{
		Time:   time.Date(now.Year()+1, time.January, 0, 23, 59, 59, 999999999, now.Location()),
		Config: nil,
	}

	if currentResult.Time.Year() != optimizedResult.Time.Year() ||
		currentResult.Time.Month() != optimizedResult.Time.Month() ||
		currentResult.Time.Day() != optimizedResult.Time.Day() {
		t.Errorf("Results don't match: current=%v, optimized=%v",
			currentResult.Time, optimizedResult.Time)
	}
}
