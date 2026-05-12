package main

import (
	"fmt"
	"time"

	"github.com/lazygophers/utils/xtime"
)

func main() {
	fmt.Println("BeginningOfHour 性能测试")
	fmt.Println("========================")

	results := []struct {
		name string
		fn   func() *xtime.Time
	}{
		{"Current (Baseline)", func() *xtime.Time {
			return xtime.BeginningOfHour()
		}},
		{"TruncateNil", func() *xtime.Time {
			t := time.Now()
			return &xtime.Time{Time: t.Truncate(time.Hour), Config: nil}
		}},
		{"GlobalConfig", func() *xtime.Time {
			t := time.Now()
			return &xtime.Time{Time: t.Truncate(time.Hour), Config: xtime.BeginningOfHourConfig}
		}},
		{"ZeroConfig", func() *xtime.Time {
			t := time.Now()
			cfg := &xtime.Config{}
			return &xtime.Time{Time: t.Truncate(time.Hour), Config: cfg}
		}},
		{"Date", func() *xtime.Time {
			t := time.Now()
			y, m, d := t.Date()
			h := t.Hour()
			return &xtime.Time{Time: time.Date(y, m, d, h, 0, 0, 0, t.Location()), Config: nil}
		}},
		{"DateWithConfig", func() *xtime.Time {
			t := time.Now()
			y, m, d := t.Date()
			h := t.Hour()
			return &xtime.Time{Time: time.Date(y, m, d, h, 0, 0, 0, t.Location()), Config: xtime.BeginningOfHourConfig}
		}},
		{"PreallocLocation", func() *xtime.Time {
			t := time.Now()
			loc := t.Location()
			y, m, d := t.Date()
			h := t.Hour()
			return &xtime.Time{Time: time.Date(y, m, d, h, 0, 0, 0, loc), Config: nil}
		}},
		{"InlinedTruncate", func() *xtime.Time {
			return &xtime.Time{Time: time.Now().Truncate(time.Hour), Config: xtime.BeginningOfHourConfig}
		}},
		{"OptimizedWith", func() *xtime.Time {
			t := time.Now()
			truncated := t.Truncate(time.Hour)
			return &xtime.Time{Time: truncated, Config: xtime.BeginningOfHourConfig}
		}},
		{"Minimal", func() *xtime.Time {
			t := time.Now()
			return &xtime.Time{Time: t.Truncate(time.Hour), Config: xtime.BeginningOfHourConfig}
		}},
	}

	for _, r := range results {
		// 预热
		for i := 0; i < 10000; i++ {
			_ = r.fn()
		}

		// 测量
		iterations := 1000000
		start := time.Now()

		for i := 0; i < iterations; i++ {
			_ = r.fn()
		}

		duration := time.Since(start)
		nsPerOp := float64(duration.Nanoseconds()) / float64(iterations)
		msPerOp := nsPerOp / 1000000.0

		fmt.Printf("%-20s: %8.1f ns/op (%6.3f µs/op)\n", r.name, nsPerOp, msPerOp)
	}

	fmt.Println("\n推荐方案: InlinedTruncate 或 Minimal (性能最优)")
}
