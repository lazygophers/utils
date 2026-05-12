package xtime

import (
	"fmt"
	"testing"
	"time"
)

// 简单的基准测试结果收集器
func TestBenchmarkBeginningOfHour(t *testing.T) {
	results := []struct {
		name     string
		fn       func()
	}{
		{"Current", func() { _ = BeginningOfHour() }},
		{"TruncateNil", func() {
			t := time.Now()
			_ = &Time{Time: t.Truncate(time.Hour), Config: nil}
		}},
		{"GlobalConfig", func() {
			t := time.Now()
			_ = &Time{Time: t.Truncate(time.Hour), Config: BeginningOfHourConfig}
		}},
		{"Date", func() {
			t := time.Now()
			y, m, d := t.Date()
			h := t.Hour()
			_ = &Time{Time: time.Date(y, m, d, h, 0, 0, 0, t.Location()), Config: nil}
		}},
		{"InlinedTruncate", func() {
			_ = &Time{Time: time.Now().Truncate(time.Hour), Config: BeginningOfHourConfig}
		}},
	}

	for _, r := range results {
		// 预热
		for i := 0; i < 1000; i++ {
			r.fn()
		}

		// 测量
		start := time.Now()
		iterations := 100000
		for i := 0; i < iterations; i++ {
			r.fn()
		}
		duration := time.Since(start)

		nsPerOp := float64(duration.Nanoseconds()) / float64(iterations)
		fmt.Printf("%-20s: %8.1f ns/op\n", r.name, nsPerOp)
	}
}
