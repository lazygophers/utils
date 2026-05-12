package xtime

import (
	"testing"
	"time"
)

// 验证优化后的 BeginningOfMinute 性能
func BenchmarkBOM_NewImplementation(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_ = BeginningOfMinute()
	}
}

// 并行测试
func BenchmarkBOM_NewImplementation_Parallel(b *testing.B) {
	b.ReportAllocs()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			_ = BeginningOfMinute()
		}
	})
}

// 对比测试：旧实现
func BenchmarkBOM_OldImplementation(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_ = With(time.Now()).BeginningOfMinute()
	}
}

// 验证正确性
func TestBOMOptimizationCorrectness(t *testing.T) {
	// 测试特定时间
	testTime := time.Date(2024, 1, 15, 14, 32, 45, 123456789, time.Local)
	result := With(testTime).BeginningOfMinute()

	expected := time.Date(2024, 1, 15, 14, 32, 0, 0, time.Local)
	if result.Time != expected {
		t.Errorf("Expected %v, got %v", expected, result.Time)
	}

	// 测试秒和纳秒归零
	if result.Second() != 0 {
		t.Errorf("Expected 0 seconds, got %d", result.Second())
	}
	if result.Nanosecond() != 0 {
		t.Errorf("Expected 0 nanoseconds, got %d", result.Nanosecond())
	}
}
