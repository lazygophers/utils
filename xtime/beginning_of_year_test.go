package xtime

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

// TestBeginningOfYear_Correctness 测试 BeginningOfYear 的正确性
func TestBeginningOfYear_Correctness(t *testing.T) {
	tests := []struct {
		name     string
		input    time.Time
		expected time.Time
	}{
		{
			name:     "2024-06-15 (年中)",
			input:    time.Date(2024, 6, 15, 14, 30, 45, 123456789, time.Local),
			expected: time.Date(2024, 1, 1, 0, 0, 0, 0, time.Local),
		},
		{
			name:     "2024-01-01 (年初)",
			input:    time.Date(2024, 1, 1, 0, 0, 0, 0, time.Local),
			expected: time.Date(2024, 1, 1, 0, 0, 0, 0, time.Local),
		},
		{
			name:     "2024-12-31 23:59:59 (年末)",
			input:    time.Date(2024, 12, 31, 23, 59, 59, 999999999, time.Local),
			expected: time.Date(2024, 1, 1, 0, 0, 0, 0, time.Local),
		},
		{
			name:     "2023-03-15 (闰年后)",
			input:    time.Date(2023, 3, 15, 10, 20, 30, 0, time.Local),
			expected: time.Date(2023, 1, 1, 0, 0, 0, 0, time.Local),
		},
		{
			name:     "2020-02-29 (闰日)",
			input:    time.Date(2020, 2, 29, 12, 0, 0, 0, time.Local),
			expected: time.Date(2020, 1, 1, 0, 0, 0, 0, time.Local),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			x := With(tt.input)
			result := x.BeginningOfYear()
			assert.Equal(t, tt.expected, result.Time, "时间应该等于年初")
		})
	}
}

// TestBeginningOfYear_ConfigPreservation 测试 Config 是否正确复用
func TestBeginningOfYear_ConfigPreservation(t *testing.T) {
	config := &Config{
		WeekStartDay:  time.Sunday,
		TimeLocation:  time.UTC,
		TimeFormats:   []string{"2006-01-02", "15:04:05", "2006-01-02 15:04:05"},
	}

	x := &Time{
		Time:   time.Date(2024, 6, 15, 14, 30, 45, 0, time.Local),
		Config: config,
	}

	result := x.BeginningOfYear()

	assert.Same(t, config, result.Config, "Config 应该被复用")
	assert.Equal(t, time.Sunday, result.Config.WeekStartDay)
	assert.Equal(t, time.UTC, result.Config.TimeLocation)
}

// TestBeginningOfYear_NilConfig 测试 nil Config 处理
func TestBeginningOfYear_NilConfig(t *testing.T) {
	x := &Time{
		Time:   time.Date(2024, 6, 15, 14, 30, 45, 0, time.Local),
		Config: nil,
	}

	result := x.BeginningOfYear()

	// nil Config 应该保持 nil（不自动创建）
	assert.Nil(t, result.Config, "nil Config 应该保持 nil")
	assert.Equal(t, time.Date(2024, 1, 1, 0, 0, 0, 0, time.Local), result.Time)
}

// TestBeginningOfYear_DifferentTimeZones 测试不同时区
func TestBeginningOfYear_DifferentTimeZones(t *testing.T) {
	// UTC
	t1 := With(time.Date(2024, 6, 15, 14, 30, 45, 0, time.UTC))
	result1 := t1.BeginningOfYear()
	assert.Equal(t, time.UTC, result1.Location())
	assert.Equal(t, 2024, result1.Year())
	assert.Equal(t, time.January, result1.Month())
	assert.Equal(t, 1, result1.Day())

	// America/New_York
	ny, _ := time.LoadLocation("America/New_York")
	t2 := With(time.Date(2024, 6, 15, 14, 30, 45, 0, ny))
	result2 := t2.BeginningOfYear()
	assert.Equal(t, ny, result2.Location())
	assert.Equal(t, 2024, result2.Year())
	assert.Equal(t, time.January, result2.Month())
	assert.Equal(t, 1, result2.Day())

	// Asia/Shanghai
	sh, _ := time.LoadLocation("Asia/Shanghai")
	t3 := With(time.Date(2024, 6, 15, 14, 30, 45, 0, sh))
	result3 := t3.BeginningOfYear()
	assert.Equal(t, sh, result3.Location())
	assert.Equal(t, 2024, result3.Year())
	assert.Equal(t, time.January, result3.Month())
	assert.Equal(t, 1, result3.Day())
}

// TestBeginningOfYear_ZeroTime 测试零值时间
func TestBeginningOfYear_ZeroTime(t *testing.T) {
	var x *Time // nil 指针
	if x != nil {
		result := x.BeginningOfYear()
		assert.NotNil(t, result)
	}
}

// TestBeginningOfYear_Immutable 测试不可变性（原对象不应被修改）
func TestBeginningOfYear_Immutable(t *testing.T) {
	original := With(time.Date(2024, 6, 15, 14, 30, 45, 123456789, time.Local))
	originalTime := original.Time

	result := original.BeginningOfYear()

	// 原对象不应被修改
	assert.Equal(t, originalTime, original.Time)
	assert.Equal(t, time.Date(2024, 6, 15, 14, 30, 45, 123456789, time.Local), original.Time)

	// 返回新对象
	assert.NotSame(t, original, result)
	assert.Equal(t, time.Date(2024, 1, 1, 0, 0, 0, 0, time.Local), result.Time)
}

// TestBeginningOfYear_LeapYear 测试闰年
func TestBeginningOfYear_LeapYear(t *testing.T) {
	leapYears := []int{2000, 2004, 2020, 2024}
	for _, year := range leapYears {
		t.Run("leap_year_"+string(rune(year)), func(t *testing.T) {
			x := With(time.Date(year, 6, 15, 12, 0, 0, 0, time.Local))
			result := x.BeginningOfYear()
			assert.Equal(t, year, result.Year())
			assert.Equal(t, time.January, result.Month())
			assert.Equal(t, 1, result.Day())
		})
	}
}

// TestBeginningOfYear_Consistency 测试多次调用一致性
func TestBeginningOfYear_Consistency(t *testing.T) {
	x := With(time.Date(2024, 6, 15, 14, 30, 45, 0, time.Local))

	result1 := x.BeginningOfYear()
	result2 := x.BeginningOfYear()
	result3 := x.BeginningOfYear()

	// 多次调用应该返回相同结果
	assert.Equal(t, result1.Time, result2.Time)
	assert.Equal(t, result2.Time, result3.Time)
	assert.Same(t, result1.Config, result2.Config)
	assert.Same(t, result2.Config, result3.Config)
}

// TestBeginningOfYear_Dependents 测试依赖函数（如 EndOfYear）
func TestBeginningOfYear_Dependents(t *testing.T) {
	x := With(time.Date(2024, 6, 15, 14, 30, 45, 0, time.Local))

	boy := x.BeginningOfYear()
	eoy := x.EndOfYear()

	// 年初应该是 1 月 1 日
	assert.Equal(t, time.January, boy.Month())
	assert.Equal(t, 1, boy.Day())

	// 年末应该是 12 月 31 日
	assert.Equal(t, time.December, eoy.Month())
	assert.Equal(t, 31, eoy.Day())

	// 年初和年末应该同一年
	assert.Equal(t, boy.Year(), eoy.Year())
}
