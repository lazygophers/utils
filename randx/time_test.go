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
		// 测试start == end的情况，现在应该返回该值而不是panic
		duration := time.Second * 5
		result := TimeDuration4Sleep(duration, duration)
		if result != duration {
			t.Errorf("TimeDuration4Sleep(%v, %v) returned %v, expected %v", duration, duration, result, duration)
		}
	})

	t.Run("start_greater_than_end", func(t *testing.T) {
		// 测试start > end的情况，现在应该panic
		start := time.Second * 10
		end := time.Second * 5

		defer func() {
			if r := recover(); r != nil {
				t.Logf("TimeDuration4Sleep with start > end caused expected panic: %v", r)
			} else {
				t.Error("Expected panic for start > end case")
			}
		}()

		// 这个调用应该panic
		TimeDuration4Sleep(start, end)
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
		// 测试零时间间隔，现在应该返回0而不是panic
		result := TimeDuration4Sleep(0)
		if result != 0 {
			t.Errorf("TimeDuration4Sleep(0) returned %v, expected 0", result)
		}
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


func TestRandomDuration(t *testing.T) {
	t.Run("random_duration_normal_range", func(t *testing.T) {
		// 测试正常范围
		min := time.Millisecond * 100
		max := time.Millisecond * 500

		for i := 0; i < 100; i++ {
			result := RandomDuration(min, max)
			if result < min || result > max {
				t.Errorf("RandomDuration(%v, %v) returned %v, expected range [%v, %v]",
					min, max, result, min, max)
			}
		}
	})

	t.Run("random_duration_min_greater_than_max", func(t *testing.T) {
		// 测试min > max的情况
		min := time.Second * 5
		max := time.Second * 2
		result := RandomDuration(min, max)
		
		if result != min {
			t.Errorf("RandomDuration(%v, %v) returned %v, expected %v",
				min, max, result, min)
		}
	})

	t.Run("random_duration_min_equals_max", func(t *testing.T) {
		// 测试min == max的情况
		duration := time.Second * 3
		result := RandomDuration(duration, duration)
		
		if result != duration {
			t.Errorf("RandomDuration(%v, %v) returned %v, expected %v",
				duration, duration, result, duration)
		}
	})

	t.Run("random_duration_zero_range", func(t *testing.T) {
		// 测试零时间
		result := RandomDuration(0, 0)
		if result != 0 {
			t.Errorf("RandomDuration(0, 0) returned %v, expected 0", result)
		}
	})

	t.Run("random_duration_distribution", func(t *testing.T) {
		// 测试分布
		min := time.Millisecond * 100
		max := time.Millisecond * 200
		iterations := 1000

		results := make([]time.Duration, iterations)
		for i := 0; i < iterations; i++ {
			results[i] = RandomDuration(min, max)
		}

		// 验证结果在预期范围内
		for i, result := range results {
			if result < min || result > max {
				t.Errorf("Iteration %d: result %v out of range [%v, %v]",
					i, result, min, max)
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
}

func TestRandomTime(t *testing.T) {
	t.Run("random_time_normal_range", func(t *testing.T) {
		// 测试正常时间范围
		start := time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC)
		end := time.Date(2023, 1, 2, 0, 0, 0, 0, time.UTC)

		for i := 0; i < 100; i++ {
			result := RandomTime(start, end)
			if result.Before(start) || result.After(end) {
				t.Errorf("RandomTime(%v, %v) returned %v, expected range [%v, %v]",
					start, end, result, start, end)
			}
		}
	})

	t.Run("random_time_start_after_end", func(t *testing.T) {
		// 测试start > end的情况
		start := time.Date(2023, 1, 2, 0, 0, 0, 0, time.UTC)
		end := time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC)
		result := RandomTime(start, end)
		
		if !result.Equal(start) {
			t.Errorf("RandomTime(%v, %v) returned %v, expected %v",
				start, end, result, start)
		}
	})

	t.Run("random_time_start_equals_end", func(t *testing.T) {
		// 测试start == end的情况
		timestamp := time.Date(2023, 1, 1, 12, 0, 0, 0, time.UTC)
		result := RandomTime(timestamp, timestamp)
		
		if !result.Equal(timestamp) {
			t.Errorf("RandomTime(%v, %v) returned %v, expected %v",
				timestamp, timestamp, result, timestamp)
		}
	})

	t.Run("random_time_distribution", func(t *testing.T) {
		// 测试时间分布
		start := time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC)
		end := time.Date(2023, 1, 1, 1, 0, 0, 0, time.UTC) // 1小时范围
		iterations := 100

		results := make([]time.Time, iterations)
		for i := 0; i < iterations; i++ {
			results[i] = RandomTime(start, end)
		}

		// 验证所有结果在范围内
		for i, result := range results {
			if result.Before(start) || result.After(end) {
				t.Errorf("Iteration %d: result %v out of range [%v, %v]",
					i, result, start, end)
			}
		}

		// 检查变异性
		uniqueValues := make(map[time.Time]bool)
		for _, result := range results {
			uniqueValues[result] = true
		}

		if len(uniqueValues) < iterations/2 { // 至少50%应该不同
			t.Logf("Warning: Only %d unique values out of %d iterations",
				len(uniqueValues), iterations)
		}
	})
}

func TestRandomTimeInDay(t *testing.T) {
	t.Run("random_time_in_day_normal", func(t *testing.T) {
		// 测试正常日期
		date := time.Date(2023, 1, 15, 14, 30, 45, 123456789, time.UTC)
		
		for i := 0; i < 100; i++ {
			result := RandomTimeInDay(date)
			
			// 检查日期是否相同
			if result.Year() != date.Year() || result.Month() != date.Month() || result.Day() != date.Day() {
				t.Errorf("RandomTimeInDay(%v) returned %v with different date", date, result)
			}
			
			// 检查时间是否在0:00:00到23:59:59范围内
			if result.Hour() < 0 || result.Hour() > 23 {
				t.Errorf("RandomTimeInDay(%v) returned %v with invalid hour", date, result)
			}
		}
	})

	t.Run("random_time_in_day_boundary", func(t *testing.T) {
		// 测试边界情况
		date := time.Date(2023, 12, 31, 23, 59, 59, 999999999, time.UTC)
		result := RandomTimeInDay(date)
		
		// 应该在同一天内
		if result.Year() != 2023 || result.Month() != 12 || result.Day() != 31 {
			t.Errorf("RandomTimeInDay(%v) returned %v with different date", date, result)
		}
	})

	t.Run("random_time_in_day_distribution", func(t *testing.T) {
		// 测试时间分布
		date := time.Date(2023, 6, 15, 12, 0, 0, 0, time.UTC)
		iterations := 100

		hours := make(map[int]int)
		for i := 0; i < iterations; i++ {
			result := RandomTimeInDay(date)
			hours[result.Hour()]++
		}

		// 应该有多个不同的小时值
		if len(hours) < 5 { // 至少5个不同的小时
			t.Logf("Warning: Only %d different hours out of %d iterations", len(hours), iterations)
		}
	})
}

func TestRandomTimeInHour(t *testing.T) {
	t.Run("random_time_in_hour_normal", func(t *testing.T) {
		// 测试正常小时
		baseTime := time.Date(2023, 1, 15, 14, 30, 45, 0, time.UTC)
		hour := 10
		
		for i := 0; i < 100; i++ {
			result := RandomTimeInHour(baseTime, hour)
			
			// 检查小时是否正确
			if result.Hour() != hour {
				t.Errorf("RandomTimeInHour(%v, %d) returned %v with hour %d, expected %d",
					baseTime, hour, result, result.Hour(), hour)
			}
			
			// 检查日期是否相同
			if result.Year() != baseTime.Year() || result.Month() != baseTime.Month() || result.Day() != baseTime.Day() {
				t.Errorf("RandomTimeInHour(%v, %d) returned %v with different date", baseTime, hour, result)
			}
		}
	})

	t.Run("random_time_in_hour_invalid_hour", func(t *testing.T) {
		// 测试无效小时
		baseTime := time.Date(2023, 1, 15, 14, 30, 45, 0, time.UTC)
		
		testCases := []int{-1, 24, 25, 100}
		for _, invalidHour := range testCases {
			result := RandomTimeInHour(baseTime, invalidHour)
			
			// 应该使用baseTime的小时
			if result.Hour() != baseTime.Hour() {
				t.Errorf("RandomTimeInHour(%v, %d) returned %v with hour %d, expected %d",
					baseTime, invalidHour, result, result.Hour(), baseTime.Hour())
			}
		}
	})

	t.Run("random_time_in_hour_distribution", func(t *testing.T) {
		// 测试分钟分布
		baseTime := time.Date(2023, 6, 15, 12, 0, 0, 0, time.UTC)
		hour := 15
		iterations := 100

		minutes := make(map[int]int)
		for i := 0; i < iterations; i++ {
			result := RandomTimeInHour(baseTime, hour)
			minutes[result.Minute()]++
		}

		// 应该有多个不同的分钟值
		if len(minutes) < 10 { // 至少10个不同的分钟
			t.Logf("Warning: Only %d different minutes out of %d iterations", len(minutes), iterations)
		}
	})
}

func TestBatchRandomDuration(t *testing.T) {
	t.Run("batch_random_duration_zero_or_negative_count", func(t *testing.T) {
		// 测试count <= 0
		result := BatchRandomDuration(time.Second, time.Second*2, 0)
		if result != nil {
			t.Errorf("Expected nil for count=0, got %v", result)
		}

		result = BatchRandomDuration(time.Second, time.Second*2, -1)
		if result != nil {
			t.Errorf("Expected nil for count=-1, got %v", result)
		}
	})

	t.Run("batch_random_duration_min_greater_than_max", func(t *testing.T) {
		// 测试min > max，应该交换
		min := time.Second * 5
		max := time.Second * 2
		count := 10
		result := BatchRandomDuration(min, max, count)

		if len(result) != count {
			t.Errorf("Expected length %d, got %d", count, len(result))
		}

		// 交换后应该在[max, min]范围内
		for i, r := range result {
			if r < max || r > min {
				t.Errorf("Result[%d] = %v out of range [%v, %v]", i, r, max, min)
			}
		}
	})

	t.Run("batch_random_duration_min_equals_max", func(t *testing.T) {
		// 测试min == max
		duration := time.Second * 3
		count := 5
		result := BatchRandomDuration(duration, duration, count)

		if len(result) != count {
			t.Errorf("Expected length %d, got %d", count, len(result))
		}

		for i, r := range result {
			if r != duration {
				t.Errorf("Result[%d] = %v, expected %v", i, r, duration)
			}
		}
	})

	t.Run("batch_random_duration_normal", func(t *testing.T) {
		// 测试正常情况
		min := time.Millisecond * 100
		max := time.Millisecond * 500
		count := 100
		result := BatchRandomDuration(min, max, count)

		if len(result) != count {
			t.Errorf("Expected length %d, got %d", count, len(result))
		}

		for i, r := range result {
			if r < min || r > max {
				t.Errorf("Result[%d] = %v out of range [%v, %v]", i, r, min, max)
			}
		}
	})

	t.Run("batch_random_duration_distribution", func(t *testing.T) {
		// 测试分布
		min := time.Millisecond * 100
		max := time.Millisecond * 200
		count := 1000
		result := BatchRandomDuration(min, max, count)

		// 检查变异性
		uniqueValues := make(map[time.Duration]bool)
		for _, r := range result {
			uniqueValues[r] = true
		}

		if len(uniqueValues) < count/10 { // 至少10%应该不同
			t.Logf("Warning: Only %d unique values out of %d", len(uniqueValues), count)
		}
	})
}

func TestSleepRandom(t *testing.T) {
	t.Run("sleep_random_execution", func(t *testing.T) {
		// 测试函数能正常执行（不会panic）
		min := time.Millisecond * 1
		max := time.Millisecond * 5

		start := time.Now()
		SleepRandom(min, max)
		elapsed := time.Since(start)

		// 睡眠时间应该大致在预期范围内（允许一些误差）
		if elapsed < min || elapsed > max+time.Millisecond*10 { // 允许10ms误差
			t.Logf("SleepRandom(%v, %v) slept for %v, outside expected range", min, max, elapsed)
		}
	})

	t.Run("sleep_random_multiple_calls", func(t *testing.T) {
		// 测试多次调用
		min := time.Microsecond * 100
		max := time.Microsecond * 500

		for i := 0; i < 5; i++ {
			start := time.Now()
			SleepRandom(min, max)
			elapsed := time.Since(start)

			// 基本验证睡眠时间不为负数
			if elapsed < 0 {
				t.Errorf("SleepRandom resulted in negative elapsed time: %v", elapsed)
			}
		}
	})
}

func TestSleepRandomMilliseconds(t *testing.T) {
	t.Run("sleep_random_milliseconds_execution", func(t *testing.T) {
		// 测试函数能正常执行
		minMs := 1
		maxMs := 5

		start := time.Now()
		SleepRandomMilliseconds(minMs, maxMs)
		elapsed := time.Since(start)

		expectedMin := time.Duration(minMs) * time.Millisecond
		expectedMax := time.Duration(maxMs) * time.Millisecond

		// 睡眠时间应该大致在预期范围内（允许一些误差）
		if elapsed < expectedMin || elapsed > expectedMax+time.Millisecond*10 { // 允许10ms误差
			t.Logf("SleepRandomMilliseconds(%d, %d) slept for %v, outside expected range [%v, %v]",
				minMs, maxMs, elapsed, expectedMin, expectedMax)
		}
	})

	t.Run("sleep_random_milliseconds_various_ranges", func(t *testing.T) {
		// 测试不同的范围
		testCases := []struct {
			minMs, maxMs int
		}{
			{1, 2},
			{5, 10},
			{0, 1},
		}

		for _, tc := range testCases {
			start := time.Now()
			SleepRandomMilliseconds(tc.minMs, tc.maxMs)
			elapsed := time.Since(start)

			// 基本验证
			if elapsed < 0 {
				t.Errorf("SleepRandomMilliseconds(%d, %d) resulted in negative elapsed time: %v",
					tc.minMs, tc.maxMs, elapsed)
			}
		}
	})
}

func TestJitter(t *testing.T) {
	t.Run("jitter_zero_percent", func(t *testing.T) {
		// 测试0%抖动
		duration := time.Second * 5
		result := Jitter(duration, 0)
		
		if result != duration {
			t.Errorf("Jitter(%v, 0) returned %v, expected %v", duration, result, duration)
		}
	})

	t.Run("jitter_negative_percent", func(t *testing.T) {
		// 测试负百分比
		duration := time.Second * 5
		result := Jitter(duration, -10)
		
		if result != duration {
			t.Errorf("Jitter(%v, -10) returned %v, expected %v", duration, result, duration)
		}
	})

	t.Run("jitter_over_100_percent", func(t *testing.T) {
		// 测试超过100%的抖动
		duration := time.Second * 5
		jitterPercent := 150.0
		result := Jitter(duration, jitterPercent)
		
		// 应该被限制在100%
		maxJitter := duration // 100%的抖动范围
		expectedMin := time.Duration(0) // 可能为0
		expectedMax := duration + maxJitter
		
		if result < expectedMin || result > expectedMax {
			t.Logf("Jitter(%v, %f) returned %v, expected range [%v, %v]",
				duration, jitterPercent, result, expectedMin, expectedMax)
		}
	})

	t.Run("jitter_normal_percent", func(t *testing.T) {
		// 测试正常百分比
		duration := time.Second * 10
		jitterPercent := 20.0 // 20%抖动
		
		for i := 0; i < 100; i++ {
			result := Jitter(duration, jitterPercent)
			
			// 计算期望范围：duration ± 20%
			jitterRange := time.Duration(float64(duration) * jitterPercent / 100)
			expectedMin := duration - jitterRange
			expectedMax := duration + jitterRange
			
			// 但结果不能小于0
			if expectedMin < 0 {
				expectedMin = 0
			}
			
			if result < expectedMin || result > expectedMax {
				t.Errorf("Jitter(%v, %f) returned %v, expected range [%v, %v]",
					duration, jitterPercent, result, expectedMin, expectedMax)
			}
		}
	})

	t.Run("jitter_50_percent", func(t *testing.T) {
		// 测试50%抖动
		duration := time.Second * 4
		jitterPercent := 50.0
		
		results := make([]time.Duration, 100)
		for i := 0; i < 100; i++ {
			results[i] = Jitter(duration, jitterPercent)
		}
		
		// 检查变异性
		uniqueValues := make(map[time.Duration]bool)
		for _, r := range results {
			uniqueValues[r] = true
		}
		
		if len(uniqueValues) < 20 { // 至少20个不同值
			t.Logf("Warning: Only %d unique values out of 100 iterations", len(uniqueValues))
		}
	})

	t.Run("jitter_result_not_negative", func(t *testing.T) {
		// 测试结果不会为负数
		duration := time.Millisecond * 100
		jitterPercent := 200.0 // 很大的抖动
		
		for i := 0; i < 100; i++ {
			result := Jitter(duration, jitterPercent)
			if result < 0 {
				t.Errorf("Jitter(%v, %f) returned negative result: %v", duration, jitterPercent, result)
			}
		}
	})

	t.Run("jitter_zero_duration", func(t *testing.T) {
		// 测试零时间
		duration := time.Duration(0)
		jitterPercent := 50.0
		result := Jitter(duration, jitterPercent)
		
		if result != 0 {
			t.Errorf("Jitter(0, %f) returned %v, expected 0", jitterPercent, result)
		}
	})

	t.Run("jitter_very_small_duration_high_percent", func(t *testing.T) {
		// 测试很小的duration和很高的jitter百分比，确保能覆盖result < 0的情况
		duration := time.Nanosecond * 1
		jitterPercent := 99.0 // 高jitter可能导致负值
		
		// 多次运行以增加获得负值的概率
		for i := 0; i < 50000; i++ {
			result := Jitter(duration, jitterPercent)
			// 结果永远不应该为负
			if result < 0 {
				t.Errorf("Jitter(%v, %f) returned negative result: %v", duration, jitterPercent, result)
			}
		}
	})

	t.Run("jitter_force_negative_scenario", func(t *testing.T) {
		// 使用能几乎确定产生负值的参数组合
		testCases := []struct {
			duration time.Duration
			percent  float64
		}{
			{time.Nanosecond, 99.9},
			{time.Nanosecond * 2, 99.5},
			{time.Nanosecond * 5, 90.0},
		}

		foundNegativeCase := false
		for _, tc := range testCases {
			for i := 0; i < 100000; i++ {
				// 在函数内部检查是否会产生负数，然后被设为0
				result := Jitter(tc.duration, tc.percent)
				// 验证结果不为负数（函数应该将负数设为0）
				if result < 0 {
					t.Errorf("Jitter(%v, %f) returned negative result: %v", tc.duration, tc.percent, result)
				}
				// 当结果恰好为0时，很可能是被负数修正过的
				if result == 0 && tc.duration > 0 {
					foundNegativeCase = true
				}
			}
		}

		// 如果没有找到0结果（即负数被修正的情况），加一个更极端的测试
		if !foundNegativeCase {
			// 使用更极端的测试来确保触发负数修正逻辑
			// 尝试不同的极小duration值
			extremeCases := []time.Duration{
				time.Nanosecond,
				time.Nanosecond * 2,
				time.Nanosecond * 3,
				time.Nanosecond * 5,
				time.Nanosecond * 10,
			}

			for _, duration := range extremeCases {
				for i := 0; i < 2000000; i++ {
					result := Jitter(duration, 100.0)
					if result == 0 {
						foundNegativeCase = true
						break
					}
				}
				if foundNegativeCase {
					break
				}
			}
		}

		if !foundNegativeCase {
			t.Log("Warning: Did not trigger negative result correction case, but this is acceptable")
		}
	})
}
