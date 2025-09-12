package randx

// Choose 高性能版本，使用优化的随机数生成器
func Choose[T any](s []T) T {
	if len(s) == 0 {
		return *new(T)
	}

	if len(s) == 1 {
		return s[0]
	}

	return s[FastIntn(len(s))]
}

// FastChoose 使用全局生成器的超快版本
func FastChoose[T any](s []T) T {
	if len(s) == 0 {
		return *new(T)
	}

	if len(s) == 1 {
		return s[0]
	}

	globalMu.Lock()
	idx := globalRand.Intn(len(s))
	globalMu.Unlock()
	
	return s[idx]
}

// ChooseN 从切片中选择N个不重复的元素（高性能版本）
func ChooseN[T any](s []T, n int) []T {
	if len(s) == 0 || n <= 0 {
		return []T{}
	}
	
	if n >= len(s) {
		// 返回所有元素的副本
		result := make([]T, len(s))
		copy(result, s)
		return result
	}
	
	// 使用Fisher-Yates洗牌算法选择前N个
	sCopy := make([]T, len(s))
	copy(sCopy, s)
	
	r := getFastRand()
	
	for i := 0; i < n; i++ {
		j := i + r.Intn(len(sCopy)-i)
		sCopy[i], sCopy[j] = sCopy[j], sCopy[i]
	}
	
	putFastRand(r)
	
	return sCopy[:n]
}

// Shuffle 随机打乱切片（高性能版本）
func Shuffle[T any](s []T) {
	if len(s) <= 1 {
		return
	}
	
	r := getFastRand()
	
	// Fisher-Yates 洗牌算法
	for i := len(s) - 1; i > 0; i-- {
		j := r.Intn(i + 1)
		s[i], s[j] = s[j], s[i]
	}
	
	putFastRand(r)
}

// FastShuffle 使用全局生成器的超快版本
func FastShuffle[T any](s []T) {
	if len(s) <= 1 {
		return
	}
	
	globalMu.Lock()
	defer globalMu.Unlock()
	
	for i := len(s) - 1; i > 0; i-- {
		j := globalRand.Intn(i + 1)
		s[i], s[j] = s[j], s[i]
	}
}

// WeightedChoose 加权选择（权重数组对应元素选择概率）
func WeightedChoose[T any](items []T, weights []float64) T {
	if len(items) == 0 || len(items) != len(weights) {
		return *new(T)
	}
	
	if len(items) == 1 {
		return items[0]
	}
	
	// 计算权重总和
	totalWeight := 0.0
	for _, w := range weights {
		totalWeight += w
	}
	
	if totalWeight <= 0 {
		return items[FastIntn(len(items))]
	}
	
	// 生成随机数
	r := FastFloat64() * totalWeight
	
	// 找到对应的元素
	accumWeight := 0.0
	for i, weight := range weights {
		accumWeight += weight
		if r <= accumWeight {
			return items[i]
		}
	}
	
	// 理论上不应该到达这里，但为安全起见
	return items[len(items)-1]
}

// BatchChoose 批量从切片中选择元素
func BatchChoose[T any](s []T, count int) []T {
	if len(s) == 0 || count <= 0 {
		return []T{}
	}
	
	results := make([]T, count)
	
	if len(s) == 1 {
		for i := range results {
			results[i] = s[0]
		}
		return results
	}
	
	r := getFastRand()
	
	for i := 0; i < count; i++ {
		results[i] = s[r.Intn(len(s))]
	}
	
	putFastRand(r)
	return results
}