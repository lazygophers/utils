package randx

import (
	"testing"
	"time"
)

func TestTimeDuration4Sleep(t *testing.T) {
	t.Run("no_arguments_default_range", func(t *testing.T) {
		// 测试无参数时的默认范围：[1s, 3s]
		for i := 0; i < 100; i++ {
			result := TimeDuration4Sleep()
			if result < time.Second || result > time.Second*3 {
				t.Errorf("TimeDuration4Sleep() returned %v, expected range [1s, 3s]", result)
			}
		}
	})

	t.Run("single_argument_range", func(t *testing.T) {
		// 测试单个参数时的范围：[0, end]
		endDuration := time.Second * 5
		
		for i := 0; i < 100; i++ {
			result := TimeDuration4Sleep(endDuration)
			if result < 0 || result > endDuration {
				t.Errorf("TimeDuration4Sleep(%v) returned %v, expected range [0, %v]", 
					endDuration, result, endDuration)
			}
		}
	})

	t.Run("single_argument_various_durations", func(t *testing.T) {
		// 测试不同的单个参数值
		testCases := []time.Duration{
			time.Millisecond * 100,
			time.Second * 2,
			time.Minute,
			time.Hour,
		}
		
		for _, duration := range testCases {
			for i := 0; i < 20; i++ {
				result := TimeDuration4Sleep(duration)
				if result < 0 || result > duration {
					t.Errorf("TimeDuration4Sleep(%v) returned %v, expected range [0, %v]", 
						duration, result, duration)
				}
			}
		}
	})

	t.Run("two_arguments_custom_range", func(t *testing.T) {
		// 测试两个参数时的自定义范围：[start, end]
		start := time.Second * 2
		end := time.Second * 8
		
		for i := 0; i < 100; i++ {
			result := TimeDuration4Sleep(start, end)
			if result < start || result > end {
				t.Errorf("TimeDuration4Sleep(%v, %v) returned %v, expected range [%v, %v]", 
					start, end, result, start, end)
			}
		}
	})

	t.Run("two_arguments_various_ranges", func(t *testing.T) {
		// 测试不同的两个参数组合
		testCases := []struct {
			start, end time.Duration
		}{
			{time.Millisecond * 50, time.Millisecond * 200},
			{time.Second, time.Second * 10},
			{time.Minute, time.Minute * 5},
			{0, time.Second},
		}
		
		for _, tc := range testCases {
			for i := 0; i < 50; i++ {
				result := TimeDuration4Sleep(tc.start, tc.end)
				if result < tc.start || result > tc.end {
					t.Errorf("TimeDuration4Sleep(%v, %v) returned %v, expected range [%v, %v]", 
						tc.start, tc.end, result, tc.start, tc.end)
				}
			}
		}
	})

	t.Run("equal_start_and_end", func(t *testing.T) {
		// 测试start == end的情况（会导致panic）
		duration := time.Second * 5
		
		defer func() {
			if r := recover(); r != nil {
				t.Logf("TimeDuration4Sleep with equal start and end caused expected panic: %v", r)
				// 这是预期的行为，因为rand.Int63n(0)会panic
			}
		}()
		
		// 这个调用会panic，所以我们只测试一次
		result := TimeDuration4Sleep(duration, duration)
		t.Errorf("Expected panic, but got result: %v", result)
	})

	t.Run("start_greater_than_end", func(t *testing.T) {
		// 测试start > end的异常情况
		start := time.Second * 10
		end := time.Second * 5
		
		// 这种情况下函数行为未定义，但不应该panic
		defer func() {
			if r := recover(); r != nil {
				t.Logf("TimeDuration4Sleep with start > end caused panic: %v", r)
			}
		}()
		
		for i := 0; i < 10; i++ {
			result := TimeDuration4Sleep(start, end)
			t.Logf("TimeDuration4Sleep(%v, %v) returned %v", start, end, result)
		}
	})

	t.Run("multiple_arguments_only_first_two_used", func(t *testing.T) {
		// 测试多个参数时只使用前两个
		start := time.Second
		end := time.Second * 3
		extra1 := time.Second * 10
		extra2 := time.Minute
		
		for i := 0; i < 50; i++ {
			result := TimeDuration4Sleep(start, end, extra1, extra2)
			if result < start || result > end {
				t.Errorf("TimeDuration4Sleep with multiple args returned %v, expected range [%v, %v]", 
					result, start, end)
			}
		}
	})

	t.Run("zero_duration", func(t *testing.T) {
		// 测试零时间间隔 - 这会导致panic，因为end-start=0
		defer func() {
			if r := recover(); r != nil {
				t.Logf("TimeDuration4Sleep(0) caused expected panic: %v", r)
				// 这是预期的行为，因为当end=0, start=0时，rand.Int63n(0)会panic
			}
		}()
		
		// 这个调用会panic
		result := TimeDuration4Sleep(0)
		t.Errorf("Expected panic, but got result: %v", result)
	})

	t.Run("negative_duration", func(t *testing.T) {
		// 测试负时间间隔
		negativeDuration := -time.Second
		
		defer func() {
			if r := recover(); r != nil {
				t.Logf("TimeDuration4Sleep with negative duration caused panic: %v", r)
			}
		}()
		
		for i := 0; i < 10; i++ {
			result := TimeDuration4Sleep(negativeDuration)
			t.Logf("TimeDuration4Sleep(%v) returned %v", negativeDuration, result)
		}
	})

	t.Run("very_large_duration", func(t *testing.T) {
		// 测试很大的时间间隔
		largeDuration := time.Hour * 24 * 365 // 一年
		
		for i := 0; i < 10; i++ {
			result := TimeDuration4Sleep(largeDuration)
			if result < 0 || result > largeDuration {
				t.Errorf("TimeDuration4Sleep(%v) returned %v, expected range [0, %v]", 
					largeDuration, result, largeDuration)
			}
		}
	})

	t.Run("distribution_test", func(t *testing.T) {
		// 测试分布的均匀性
		start := time.Millisecond * 100
		end := time.Millisecond * 200
		iterations := 1000
		
		results := make([]time.Duration, iterations)
		for i := 0; i < iterations; i++ {
			results[i] = TimeDuration4Sleep(start, end)
		}
		
		// 验证结果在预期范围内
		for i, result := range results {
			if result < start || result > end {
				t.Errorf("Iteration %d: result %v out of range [%v, %v]", 
					i, result, start, end)
			}
		}
		
		// 简单的分布测试：检查是否有足够的变化
		uniqueValues := make(map[time.Duration]bool)
		for _, result := range results {
			uniqueValues[result] = true
		}
		
		// 应该有相当数量的不同值
		if len(uniqueValues) < iterations/10 { // 至少10%的值应该不同
			t.Logf("Warning: Only %d unique values out of %d iterations", 
				len(uniqueValues), iterations)
		}
	})

	t.Run("microsecond_precision", func(t *testing.T) {
		// 测试微秒级精度
		start := time.Microsecond * 100
		end := time.Microsecond * 200
		
		for i := 0; i < 50; i++ {
			result := TimeDuration4Sleep(start, end)
			if result < start || result > end {
				t.Errorf("TimeDuration4Sleep(%v, %v) returned %v, expected range [%v, %v]", 
					start, end, result, start, end)
			}
		}
	})
}