package xtime

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestBeginningOfMonth_Dependents(t *testing.T) {
	// 测试依赖 BeginningOfMonth 的函数
	testDate := time.Date(2024, 5, 15, 14, 30, 45, 123456789, time.Local)
	xtime := With(testDate)

	t.Run("BeginningOfQuarter", func(t *testing.T) {
		result := xtime.BeginningOfQuarter()
		// 2024 Q2 starts April 1
		expected := time.Date(2024, 4, 1, 0, 0, 0, 0, time.Local)
		assert.Equal(t, expected, result.Time)
	})

	t.Run("BeginningOfHalf", func(t *testing.T) {
		result := xtime.BeginningOfHalf()
		// 2024 H1 starts January 1
		expected := time.Date(2024, 1, 1, 0, 0, 0, 0, time.Local)
		assert.Equal(t, expected, result.Time)
	})

	t.Run("EndOfMonth", func(t *testing.T) {
		result := xtime.EndOfMonth()
		// May ends at May 31 23:59:59.999999999
		expected := time.Date(2024, 6, 1, 0, 0, 0, 0, time.Local).Add(-time.Nanosecond)
		assert.Equal(t, expected, result.Time)
	})

	t.Run("Q1", func(t *testing.T) {
		q1Date := With(time.Date(2024, 2, 15, 10, 0, 0, 0, time.Local))
		result := q1Date.BeginningOfQuarter()
		expected := time.Date(2024, 1, 1, 0, 0, 0, 0, time.Local)
		assert.Equal(t, expected, result.Time)
	})

	t.Run("Q3", func(t *testing.T) {
		q3Date := With(time.Date(2024, 8, 20, 15, 30, 0, 0, time.Local))
		result := q3Date.BeginningOfQuarter()
		expected := time.Date(2024, 7, 1, 0, 0, 0, 0, time.Local)
		assert.Equal(t, expected, result.Time)
	})

	t.Run("H2", func(t *testing.T) {
		h2Date := With(time.Date(2024, 9, 10, 12, 0, 0, 0, time.Local))
		result := h2Date.BeginningOfHalf()
		expected := time.Date(2024, 7, 1, 0, 0, 0, 0, time.Local)
		assert.Equal(t, expected, result.Time)
	})
}

func TestBeginningOfMonth_Consistency(t *testing.T) {
	// 验证多次调用结果一致
	testDate := time.Date(2024, 5, 15, 14, 30, 45, 123456789, time.Local)
	xtime := With(testDate)

	result1 := xtime.BeginningOfMonth()
	result2 := xtime.BeginningOfMonth()
	result3 := xtime.BeginningOfMonth()

	assert.Equal(t, result1.Time, result2.Time)
	assert.Equal(t, result2.Time, result3.Time)
	assert.Equal(t, xtime.Config, result1.Config)
}
