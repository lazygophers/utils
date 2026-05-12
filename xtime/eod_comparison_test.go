package xtime

import (
	"testing"
	"time"
)

// 旧实现（用于对比）
func endOfDayOld(p *Time) *Time {
	y, m, d := p.Date()
	return With(time.Date(y, m, d, 23, 59, 59, int(time.Second-time.Nanosecond), p.Location()))
}

// 新实现
func endOfDayNew(p *Time) *Time {
	loc := p.Location()
	year, month, day := p.Date()
	eod := time.Date(year, month, day, 23, 59, 59, int(time.Second-time.Nanosecond), loc)
	cfg := p.Config
	if cfg == nil {
		cfg = &Config{}
	}
	return &Time{Time: eod, Config: cfg}
}

// 性能对比基准测试
func BenchmarkEOD_OldImplementation(b *testing.B) {
	times := genEODTestTimes(1000)
	wrapper := make([]*Time, len(times))
	for i := range times {
		wrapper[i] = With(times[i])
	}
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_ = endOfDayOld(wrapper[i%len(wrapper)])
	}
}

func BenchmarkEOD_NewImplementation(b *testing.B) {
	times := genEODTestTimes(1000)
	wrapper := make([]*Time, len(times))
	for i := range times {
		wrapper[i] = With(times[i])
	}
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_ = endOfDayNew(wrapper[i%len(wrapper)])
	}
}

// 正确性验证
func TestEOD_OldVsNew_Correctness(t *testing.T) {
	testTimes := []time.Time{
		time.Date(2024, 5, 11, 15, 30, 45, 123456789, time.Local),
		time.Date(2024, 1, 1, 0, 0, 0, 0, time.Local),
		time.Date(2024, 12, 31, 23, 59, 59, 999999999, time.Local),
		time.Date(2024, 2, 29, 12, 0, 0, 0, time.Local), // 闰年
	}

	for _, tt := range testTimes {
		wrapper := With(tt)
		oldResult := endOfDayOld(wrapper)
		newResult := endOfDayNew(wrapper)

		// 验证时间相同
		if !oldResult.Time.Equal(newResult.Time) {
			t.Errorf("旧实现和新实现结果不同: 输入=%v\n旧=%v\n新=%v",
				tt, oldResult.Time, newResult.Time)
		}

		// 验证 Config 保留
		if oldResult.Config != nil && newResult.Config == nil {
			t.Errorf("新实现 Config 为 nil，但旧实现不为 nil: 输入=%v", tt)
		}
	}
}
