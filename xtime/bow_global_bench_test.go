package xtime

import (
	"sync"
	"testing"
	"time"
)

// 方案0: 当前实现
func BenchmarkBeginningOfWeekGlobalCurrent(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = BeginningOfWeek()
	}
}

// 方案1: 内联逻辑，避免 With() 调用
func BenchmarkBeginningOfWeekGlobalOpt1(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		now := time.Now()
		year, month, day := now.Date()
		loc := now.Location()
		midnight := time.Date(year, month, day, 0, 0, 0, 0, loc)
		weekday := int(midnight.Weekday())
		_ = &Time{
			Time:   midnight.AddDate(0, 0, -weekday),
			Config: &Config{WeekStartDay: time.Sunday, TimeLocation: time.Local},
		}
	}
}

// 方案2: 使用全局 Config
var bowGlobalConfig = &Config{
	WeekStartDay:  time.Sunday,
	TimeLocation: time.Local,
}

func BenchmarkBeginningOfWeekGlobalOpt2(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		now := time.Now()
		midnight := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, now.Location())
		weekday := int(midnight.Weekday())
		_ = &Time{
			Time:   midnight.AddDate(0, 0, -weekday),
			Config: bowGlobalConfig,
		}
	}
}

// 方案3: 最小化变量
func BenchmarkBeginningOfWeekGlobalOpt3(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		t := time.Now()
		_ = &Time{
			Time:   time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, t.Location()).AddDate(0, 0, -int(t.Weekday())),
			Config: bowGlobalConfig,
		}
	}
}

// 方案4: 预计算 time.Local
func BenchmarkBeginningOfWeekGlobalOpt4(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		t := time.Now()
		loc := time.Local
		_ = &Time{
			Time:   time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, loc).AddDate(0, 0, -int(t.Weekday())),
			Config: bowGlobalConfig,
		}
	}
}

// 方案5: 避免重复 weekday 计算
func BenchmarkBeginningOfWeekGlobalOpt5(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		t := time.Now()
		weekday := int(t.Weekday())
		_ = &Time{
			Time:   time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, t.Location()).AddDate(0, 0, -weekday),
			Config: bowGlobalConfig,
		}
	}
}

// 方案6: 组合优化
func BenchmarkBeginningOfWeekGlobalOpt6(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		t := time.Now()
		loc := time.Local
		wd := int(t.Weekday())
		_ = &Time{
			Time:   time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, loc).AddDate(0, 0, -wd),
			Config: bowGlobalConfig,
		}
	}
}

// 方案7: 使用 sync.Pool
var bowTimePool = &sync.Pool{
	New: func() interface{} {
		return &Time{
			Config: &Config{
				WeekStartDay:  time.Sunday,
				TimeLocation: time.Local,
			},
		}
	},
}

func BenchmarkBeginningOfWeekGlobalOpt7(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		t := time.Now()
		result := bowTimePool.Get().(*Time)
		result.Time = time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, t.Location()).AddDate(0, 0, -int(t.Weekday()))
		_ = result
	}
}

// 方案8: sync.Pool + 预计算变量
func BenchmarkBeginningOfWeekGlobalOpt8(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		t := time.Now()
		loc := time.Local
		wd := int(t.Weekday())
		result := bowTimePool.Get().(*Time)
		result.Time = time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, loc).AddDate(0, 0, -wd)
		_ = result
	}
}

// 方案9: 延迟初始化 Config
var (
	bowLazyConfig     *Config
	bowLazyConfigOnce sync.Once
)

func getBowLazyConfig() *Config {
	bowLazyConfigOnce.Do(func() {
		bowLazyConfig = &Config{
			WeekStartDay:  time.Sunday,
			TimeLocation: time.Local,
		}
	})
	return bowLazyConfig
}

func BenchmarkBeginningOfWeekGlobalOpt9(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		t := time.Now()
		loc := time.Local
		wd := int(t.Weekday())
		_ = &Time{
			Time:   time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, loc).AddDate(0, 0, -wd),
			Config: getBowLazyConfig(),
		}
	}
}

// 方案10: 直接构造（最小化调用）
func BenchmarkBeginningOfWeekGlobalOpt10(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		t := time.Now()
		y, m, d := t.Date()
		_ = &Time{
			Time:   time.Date(y, m, d, 0, 0, 0, 0, t.Location()).AddDate(0, 0, -int(t.Weekday())),
			Config: bowGlobalConfig,
		}
	}
}
