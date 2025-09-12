package randx

import (
	"time"
)

// TimeDuration4Sleep 高性能版本，使用优化的随机数生成器
func TimeDuration4Sleep(s ...time.Duration) time.Duration {
	start, end := time.Second, time.Second*3
	if len(s) > 1 {
		start = s[0]
		end = s[1]
	} else if len(s) > 0 {
		start = 0
		end = s[0]
	}

	// 保持原始行为兼容性
	r := getFastRand()
	result := time.Duration(r.Int63n(int64(end-start))) + start
	putFastRand(r)
	
	return result
}

// FastTimeDuration4Sleep 使用全局生成器的超快版本
func FastTimeDuration4Sleep(s ...time.Duration) time.Duration {
	start, end := time.Second, time.Second*3
	if len(s) > 1 {
		start = s[0]
		end = s[1]
	} else if len(s) > 0 {
		start = 0
		end = s[0]
	}

	if start >= end {
		return start
	}

	globalMu.Lock()
	result := time.Duration(globalRand.Int63n(int64(end-start))) + start
	globalMu.Unlock()
	
	return result
}

// RandomDuration 在指定范围内生成随机时间间隔 [min, max]
func RandomDuration(min, max time.Duration) time.Duration {
	if min > max {
		return min
	} else if min == max {
		return min
	}

	r := getFastRand()
	result := min + time.Duration(r.Int63n(int64(max-min+1)))
	putFastRand(r)
	
	return result
}

// RandomTime 在指定时间范围内生成随机时间点
func RandomTime(start, end time.Time) time.Time {
	if start.After(end) {
		return start
	} else if start.Equal(end) {
		return start
	}

	diff := end.Sub(start)
	r := getFastRand()
	randomDiff := time.Duration(r.Int63n(int64(diff)))
	putFastRand(r)
	
	return start.Add(randomDiff)
}

// RandomTimeInDay 在指定日期的一天内生成随机时间点
func RandomTimeInDay(date time.Time) time.Time {
	// 获取当天的开始时间
	startOfDay := time.Date(date.Year(), date.Month(), date.Day(), 0, 0, 0, 0, date.Location())
	// 当天结束时间（下一天的开始时间）
	endOfDay := startOfDay.Add(24 * time.Hour)
	
	return RandomTime(startOfDay, endOfDay)
}

// RandomTimeInHour 在指定小时内生成随机时间点
func RandomTimeInHour(baseTime time.Time, hour int) time.Time {
	if hour < 0 || hour > 23 {
		hour = baseTime.Hour()
	}
	
	startOfHour := time.Date(baseTime.Year(), baseTime.Month(), baseTime.Day(), hour, 0, 0, 0, baseTime.Location())
	endOfHour := startOfHour.Add(time.Hour)
	
	return RandomTime(startOfHour, endOfHour)
}

// BatchRandomDuration 批量生成随机时间间隔
func BatchRandomDuration(min, max time.Duration, count int) []time.Duration {
	if count <= 0 {
		return nil
	}
	
	if min > max {
		min, max = max, min
	} else if min == max {
		results := make([]time.Duration, count)
		for i := range results {
			results[i] = min
		}
		return results
	}
	
	results := make([]time.Duration, count)
	r := getFastRand()
	
	diff := int64(max - min + 1)
	for i := 0; i < count; i++ {
		results[i] = min + time.Duration(r.Int63n(diff))
	}
	
	putFastRand(r)
	return results
}

// SleepRandom 随机睡眠指定范围的时间
func SleepRandom(min, max time.Duration) {
	duration := RandomDuration(min, max)
	time.Sleep(duration)
}

// SleepRandomMilliseconds 随机睡眠指定毫秒数范围
func SleepRandomMilliseconds(minMs, maxMs int) {
	min := time.Duration(minMs) * time.Millisecond
	max := time.Duration(maxMs) * time.Millisecond
	SleepRandom(min, max)
}

// Jitter 为时间间隔添加抖动（±jitterPercent%的随机变化）
func Jitter(duration time.Duration, jitterPercent float64) time.Duration {
	if jitterPercent <= 0 {
		return duration
	}
	
	if jitterPercent > 100 {
		jitterPercent = 100
	}
	
	// 计算抖动范围
	jitterRange := time.Duration(float64(duration) * jitterPercent / 100)
	
	// 生成 [-jitterRange, +jitterRange] 的随机变化
	randomJitter := RandomDuration(-jitterRange, jitterRange)
	
	result := duration + randomJitter
	if result < 0 {
		result = 0
	}
	
	return result
}