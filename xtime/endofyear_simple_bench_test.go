package xtime

import (
	"fmt"
	"testing"
	"time"
)

// 简单基准测试 - 只测试最优方案 vs baseline
func BenchmarkEndOfYear_Simple(b *testing.B) {
	t := With(time.Date(2024, 6, 15, 14, 30, 45, 123456789, time.Local))
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_ = t.EndOfYear()
	}
}

func BenchmarkEndOfYear_Optimized(b *testing.B) {
	t := With(time.Date(2024, 6, 15, 14, 30, 45, 123456789, time.Local))
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		year := t.Time.Year()
		loc := t.Time.Location()
		config := t.Config
		if config == nil {
			config = &Config{}
		}
		// time.Date(0) 溢出技巧
		end := time.Date(year+1, time.January, 0, 23, 59, 59, 999999999, loc)
		_ = &Time{Time: end, Config: config}
	}
}

// 添加简单的测试用例验证正确性
func TestEndOfYear_Correctness(t *testing.T) {
	tests := []struct {
		name     string
		date     time.Time
		wantYear int
		wantMonth time.Month
		wantDay   int
	}{
		{
			name:     "2024年6月15日",
			date:     time.Date(2024, 6, 15, 14, 30, 45, 123456789, time.Local),
			wantYear: 2024, wantMonth: time.December, wantDay: 31,
		},
		{
			name:     "2024年1月1日",
			date:     time.Date(2024, 1, 1, 0, 0, 0, 0, time.Local),
			wantYear: 2024, wantMonth: time.December, wantDay: 31,
		},
		{
			name:     "2024年12月31日中午",
			date:     time.Date(2024, 12, 31, 12, 0, 0, 0, time.Local),
			wantYear: 2024, wantMonth: time.December, wantDay: 31,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := With(tt.date).EndOfYear()
			year, month, day := got.Date()
			hour, min, sec := got.Clock()
			nsec := got.Nanosecond()

			if year != tt.wantYear || month != tt.wantMonth || day != tt.wantDay {
				t.Errorf("EndOfYear() date = %d-%02d-%02d, want %d-%02d-%02d",
					year, month, day, tt.wantYear, tt.wantMonth, tt.wantDay)
			}
			if hour != 23 || min != 59 || sec != 59 || nsec != 999999999 {
				t.Errorf("EndOfYear() time = %d:%d:%d.%d, want 23:59:59.999999999",
					hour, min, sec, nsec)
			}
			fmt.Printf("Test: %s PASS\n", tt.name)
		})
	}
}
