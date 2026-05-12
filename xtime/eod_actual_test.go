package xtime

import (
	"testing"
	"time"
)

// 测试实际的 EndOfDay 方法（优化后）
func BenchmarkEOD_ActualMethod(b *testing.B) {
	times := genEODTestTimes(1000)
	wrapper := make([]*Time, len(times))
	for i := range times {
		wrapper[i] = With(times[i])
	}
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_ = wrapper[i%len(wrapper)].EndOfDay()
	}
}

// 对比旧实现
func BenchmarkEOD_OldMethod(b *testing.B) {
	times := genEODTestTimes(1000)
	wrapper := make([]*Time, len(times))
	for i := range times {
		wrapper[i] = With(times[i])
	}
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		t := wrapper[i%len(wrapper)]
		y, m, d := t.Date()
		_ = With(time.Date(y, m, d, 23, 59, 59, int(time.Second-time.Nanosecond), t.Location()))
	}
}

// 验证 Config 复用导致的内存分配
func BenchmarkEOD_DirectConstructNoAlloc(b *testing.B) {
	times := genEODTestTimes(1000)
	wrapper := make([]*Time, len(times))
	for i := range times {
		wrapper[i] = With(times[i])
	}
	b.ResetTimer()
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		t := wrapper[i%len(wrapper)]
		loc := t.Location()
		year, month, day := t.Date()
		eod := time.Date(year, month, day, 23, 59, 59, int(time.Second-time.Nanosecond), loc)
		// 不创建 Config，直接复用
		_ = &Time{Time: eod, Config: t.Config}
	}
}
