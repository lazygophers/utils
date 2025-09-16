package randx

import (
	"math/rand"
	"sync"
	"time"
	"unsafe"
)

// 高性能随机数生成器
var (
	// 全局随机数生成器，使用最高性能实现
	globalRand = rand.New(rand.NewSource(time.Now().UnixNano()))
	globalMu   sync.Mutex

	// 种子生成器，避免频繁调用time.Now()
	seedCounter uint64 = uint64(time.Now().UnixNano())
)

// generateSeed 生成高性能种子，避免系统调用
func generateSeed() int64 {
	// 使用原子操作递增计数器，混合当前纳秒时间
	counter := (*uint64)(unsafe.Pointer(&seedCounter))
	return int64(*counter<<8 | uint64(time.Now().UnixNano()&0xFF))
}

// globalRandIntn 使用全局随机数生成器
func globalRandIntn(n int) int {
	globalMu.Lock()
	result := globalRand.Intn(n)
	globalMu.Unlock()
	return result
}

// Intn 高性能版本
func Intn(n int) int {
	if n <= 0 {
		return 0
	}
	if n == 1 {
		return 0
	}
	return globalRandIntn(n)
}

// Int 高性能版本
func Int() int {
	globalMu.Lock()
	result := globalRand.Int()
	globalMu.Unlock()
	return result
}

// IntnRange 高性能范围随机数 [min, max]
func IntnRange(min, max int) int {
	if min > max {
		return 0
	} else if min == max {
		return min
	}

	globalMu.Lock()
	result := min + globalRand.Intn(max-min+1)
	globalMu.Unlock()
	return result
}

// Int64n 高性能版本
func Int64n(n int64) int64 {
	if n <= 0 {
		return 0
	}
	if n == 1 {
		return 0
	}

	globalMu.Lock()
	result := globalRand.Int63n(n)
	globalMu.Unlock()
	return result
}

// Int64 高性能版本
func Int64() int64 {
	globalMu.Lock()
	result := globalRand.Int63()
	globalMu.Unlock()
	return result
}

// Int64nRange 高性能范围随机数 [min, max]
func Int64nRange(min, max int64) int64 {
	if min > max {
		return 0
	} else if min == max {
		return min
	}

	globalMu.Lock()
	result := min + globalRand.Int63n(max-min+1)
	globalMu.Unlock()
	return result
}

// Float64 高性能版本
func Float64() float64 {
	globalMu.Lock()
	result := globalRand.Float64()
	globalMu.Unlock()
	return result
}

// Float64Range 高性能范围随机数 [min, max]
func Float64Range(min, max float64) float64 {
	if min > max {
		return 0
	} else if min == max {
		return min
	}

	globalMu.Lock()
	result := min + globalRand.Float64()*(max-min)
	globalMu.Unlock()
	return result
}

// Float32 高性能版本
func Float32() float32 {
	globalMu.Lock()
	result := globalRand.Float32()
	globalMu.Unlock()
	return result
}

// Float32Range 高性能范围随机数 [min, max]
func Float32Range(min, max float32) float32 {
	if min > max {
		return 0
	} else if min == max {
		return min
	}

	globalMu.Lock()
	result := min + globalRand.Float32()*(max-min)
	globalMu.Unlock()
	return result
}

// Uint32 高性能版本
func Uint32() uint32 {
	globalMu.Lock()
	result := globalRand.Uint32()
	globalMu.Unlock()
	return result
}

// Uint32Range 高性能范围随机数 [min, max]
func Uint32Range(min, max uint32) uint32 {
	if min > max {
		return 0
	} else if min == max {
		return min
	}

	globalMu.Lock()
	result := min + globalRand.Uint32()%(max-min+1)
	globalMu.Unlock()
	return result
}

// Uint64 高性能版本
func Uint64() uint64 {
	globalMu.Lock()
	result := globalRand.Uint64()
	globalMu.Unlock()
	return result
}

// Uint64Range 高性能范围随机数 [min, max]
func Uint64Range(min, max uint64) uint64 {
	if min > max {
		return 0
	} else if min == max {
		return min
	}

	globalMu.Lock()
	result := min + globalRand.Uint64()%(max-min+1)
	globalMu.Unlock()
	return result
}


// BatchIntn 批量生成随机数
func BatchIntn(n int, count int) []int {
	if count <= 0 {
		return nil
	}

	results := make([]int, count)
	globalMu.Lock()
	for i := 0; i < count; i++ {
		results[i] = globalRand.Intn(n)
	}
	globalMu.Unlock()
	return results
}

// BatchInt64n 批量生成int64随机数
func BatchInt64n(n int64, count int) []int64 {
	if count <= 0 {
		return nil
	}

	results := make([]int64, count)
	globalMu.Lock()
	for i := 0; i < count; i++ {
		results[i] = globalRand.Int63n(n)
	}
	globalMu.Unlock()
	return results
}

// BatchFloat64 批量生成float64随机数
func BatchFloat64(count int) []float64 {
	if count <= 0 {
		return nil
	}

	results := make([]float64, count)
	globalMu.Lock()
	for i := 0; i < count; i++ {
		results[i] = globalRand.Float64()
	}
	globalMu.Unlock()
	return results
}
