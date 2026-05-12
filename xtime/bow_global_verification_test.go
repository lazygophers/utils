package xtime

import (
	"testing"
	"time"
)

// TestBeginningOfWeek_Global_Correctness 验证全局 BeginningOfWeek 函数正确性
func TestBeginningOfWeek_Global_Correctness(t *testing.T) {
	testCases := []struct {
		name     string
		expected time.Weekday
	}{
		{"周日", time.Sunday},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := BeginningOfWeek()

			// 验证返回值不为 nil
			if result == nil {
				t.Fatal("BeginningOfWeek() 返回 nil")
			}

			// 验证周起始日是周日
			if result.Weekday() != tc.expected {
				t.Errorf("期望 %v, 实际 %v", tc.expected, result.Weekday())
			}

			// 验证时间是午夜（00:00:00）
			h, m, s := result.Clock()
			if h != 0 || m != 0 || s != 0 {
				t.Errorf("期望午夜时间 00:00:00, 实际 %02d:%02d:%02d", h, m, s)
			}

			// 验证 Config 不为 nil
			if result.Config == nil {
				t.Error("Config 不应为 nil")
			}

			// 验证默认周起始日是周日
			if result.WeekStartDay != time.Sunday {
				t.Errorf("期望 WeekStartDay = Sunday, 实际 %v", result.WeekStartDay)
			}
		})
	}
}

// TestBeginningOfWeek_Global_MultipleCalls 验证多次调用的一致性
func TestBeginningOfWeek_Global_MultipleCalls(t *testing.T) {
	// 快速连续调用两次，验证返回类型和 Config 一致性
	result1 := BeginningOfWeek()
	result2 := BeginningOfWeek()

	if result1 == nil || result2 == nil {
		t.Fatal("BeginningOfWeek() 返回 nil")
	}

	// 验证 Config 设置一致
	if result1.WeekStartDay != result2.WeekStartDay {
		t.Errorf("WeekStartDay 不一致: %v vs %v", result1.WeekStartDay, result2.WeekStartDay)
	}

	if result1.TimeLocation != result2.TimeLocation {
		t.Errorf("TimeLocation 不一致: %v vs %v", result1.TimeLocation, result2.TimeLocation)
	}
}

// TestBeginningOfWeek_Global_ConsistencyWithMethod 验证全局函数与方法的一致性
func TestBeginningOfWeek_Global_ConsistencyWithMethod(t *testing.T) {
	globalResult := BeginningOfWeek()

	// 创建一个与全局函数相同配置的 Time 对象
	methodResult := With(time.Now())
	methodResult.Config.WeekStartDay = time.Sunday // 设置为与全局函数一致
	methodResult = methodResult.BeginningOfWeek()

	// 验证返回类型相同
	if globalResult == nil || methodResult == nil {
		t.Fatal("返回 nil")
	}

	// 验证 Config 设置一致
	if globalResult.WeekStartDay != methodResult.WeekStartDay {
		t.Errorf("WeekStartDay 不一致: %v vs %v", globalResult.WeekStartDay, methodResult.WeekStartDay)
	}

	// 验证都是周日
	if globalResult.Weekday() != time.Sunday || methodResult.Weekday() != time.Sunday {
		t.Errorf("期望都是周日, 全局函数: %v, 方法: %v", globalResult.Weekday(), methodResult.Weekday())
	}
}
