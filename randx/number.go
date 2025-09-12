package randx

import (
	"math/rand"
	"sync"
	"time"
	"unsafe"
)

// 高性能随机数生成器池
var (
	// 全局随机数生成器，用于最高性能场景
	globalRand = rand.New(rand.NewSource(time.Now().UnixNano()))
	globalMu   sync.Mutex

	// 线程本地存储池，避免锁竞争
	randPool = sync.Pool{
		New: func() interface{} {
			return rand.New(rand.NewSource(time.Now().UnixNano()))
		},
	}

	// 快速种子生成器，避免频繁调用time.Now()
	fastSeedCounter uint64 = uint64(time.Now().UnixNano())
)

// getFastRand 获取高性能随机数生成器
func getFastRand() *rand.Rand {
	return randPool.Get().(*rand.Rand)
}

// putFastRand 归还随机数生成器到池中
func putFastRand(r *rand.Rand) {
	randPool.Put(r)
}

// fastSeed 生成快速种子，避免系统调用
func fastSeed() int64 {
	// 使用原子操作递增计数器，混合当前纳秒时间
	counter := (*uint64)(unsafe.Pointer(&fastSeedCounter))
	return int64(*counter<<8 | uint64(time.Now().UnixNano()&0xFF))
}

// globalRandIntn 使用全局锁定的随机数生成器
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
	
	r := getFastRand()
	result := r.Intn(n)
	putFastRand(r)
	return result
}

// Int 高性能版本
func Int() int {
	r := getFastRand()
	result := r.Int()
	putFastRand(r)
	return result
}

// IntnRange 高性能范围随机数 [min, max]
func IntnRange(min, max int) int {
	if min > max {
		return 0
	} else if min == max {
		return min
	}

	r := getFastRand()
	result := min + r.Intn(max-min+1)
	putFastRand(r)
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
	
	r := getFastRand()
	result := r.Int63n(n)
	putFastRand(r)
	return result
}

// Int64 高性能版本
func Int64() int64 {
	r := getFastRand()
	result := r.Int63()
	putFastRand(r)
	return result
}

// Int64nRange 高性能范围随机数 [min, max]
func Int64nRange(min, max int64) int64 {
	if min > max {
		return 0
	} else if min == max {
		return min
	}

	r := getFastRand()
	result := min + r.Int63n(max-min+1)
	putFastRand(r)
	return result
}

// Float64 高性能版本
func Float64() float64 {
	r := getFastRand()
	result := r.Float64()
	putFastRand(r)
	return result
}

// Float64Range 高性能范围随机数 [min, max]
func Float64Range(min, max float64) float64 {
	if min > max {
		return 0
	} else if min == max {
		return min
	}

	r := getFastRand()
	result := min + r.Float64()*(max-min)
	putFastRand(r)
	return result
}

// Float32 高性能版本
func Float32() float32 {
	r := getFastRand()
	result := r.Float32()
	putFastRand(r)
	return result
}

// Float32Range 高性能范围随机数 [min, max]
func Float32Range(min, max float32) float32 {
	if min > max {
		return 0
	} else if min == max {
		return min
	}

	r := getFastRand()
	result := min + r.Float32()*(max-min)
	putFastRand(r)
	return result
}

// Uint32 高性能版本
func Uint32() uint32 {
	r := getFastRand()
	result := r.Uint32()
	putFastRand(r)
	return result
}

// Uint32Range 高性能范围随机数 [min, max]
func Uint32Range(min, max uint32) uint32 {
	if min > max {
		return 0
	} else if min == max {
		return min
	}

	r := getFastRand()
	result := min + r.Uint32()%(max-min+1)
	putFastRand(r)
	return result
}

// Uint64 高性能版本
func Uint64() uint64 {
	r := getFastRand()
	result := r.Uint64()
	putFastRand(r)
	return result
}

// Uint64Range 高性能范围随机数 [min, max]
func Uint64Range(min, max uint64) uint64 {
	if min > max {
		return 0
	} else if min == max {
		return min
	}

	r := getFastRand()
	result := min + r.Uint64()%(max-min+1)
	putFastRand(r)
	return result
}

// FastIntn 超快版本，使用全局锁定生成器（适合单线程或低并发）
func FastIntn(n int) int {
	if n <= 0 {
		return 0
	}
	if n == 1 {
		return 0
	}
	return globalRandIntn(n)
}

// FastInt 超快版本
func FastInt() int {
	globalMu.Lock()
	result := globalRand.Int()
	globalMu.Unlock()
	return result
}

// FastFloat64 超快版本
func FastFloat64() float64 {
	globalMu.Lock()
	result := globalRand.Float64()
	globalMu.Unlock()
	return result
}

// BatchIntn 批量生成随机数，减少池获取开销
func BatchIntn(n int, count int) []int {
	if count <= 0 {
		return nil
	}
	
	results := make([]int, count)
	r := getFastRand()
	
	for i := 0; i < count; i++ {
		results[i] = r.Intn(n)
	}
	
	putFastRand(r)
	return results
}

// BatchInt64n 批量生成int64随机数
func BatchInt64n(n int64, count int) []int64 {
	if count <= 0 {
		return nil
	}
	
	results := make([]int64, count)
	r := getFastRand()
	
	for i := 0; i < count; i++ {
		results[i] = r.Int63n(n)
	}
	
	putFastRand(r)
	return results
}

// BatchFloat64 批量生成float64随机数
func BatchFloat64(count int) []float64 {
	if count <= 0 {
		return nil
	}
	
	results := make([]float64, count)
	r := getFastRand()
	
	for i := 0; i < count; i++ {
		results[i] = r.Float64()
	}
	
	putFastRand(r)
	return results
}