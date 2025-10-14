package randx

// Bool 高性能版本，使用优化的随机数生成器
func Bool() bool {
	globalMu.Lock()
	result := globalRand.Intn(2) == 0
	globalMu.Unlock()
	return result
}

// Booln 高性能概率布尔值生成器
func Booln(n float64) bool {
	if n >= 100 {
		return true
	} else if n <= 0 {
		return false
	}

	globalMu.Lock()
	result := globalRand.Float64()*100 < n
	globalMu.Unlock()
	return result
}

// WeightedBool 加权布尔值，weight为true的权重(0.0-1.0)
func WeightedBool(weight float64) bool {
	if weight >= 1.0 {
		return true
	} else if weight <= 0.0 {
		return false
	}
	return Float64() < weight
}

// BatchBool 批量生成布尔值
func BatchBool(count int) []bool {
	if count <= 0 {
		return nil
	}

	results := make([]bool, count)
	globalMu.Lock()
	for i := 0; i < count; i++ {
		results[i] = globalRand.Intn(2) == 0
	}
	globalMu.Unlock()
	return results
}

// BatchBooln 批量生成概率布尔值
func BatchBooln(n float64, count int) []bool {
	if count <= 0 {
		return nil
	}

	if n >= 100 {
		results := make([]bool, count)
		for i := range results {
			results[i] = true
		}
		return results
	} else if n <= 0 {
		return make([]bool, count) // 默认为false
	}

	results := make([]bool, count)
	globalMu.Lock()
	for i := 0; i < count; i++ {
		results[i] = globalRand.Float64()*100 < n
	}
	globalMu.Unlock()
	return results
}
