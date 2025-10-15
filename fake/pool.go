package fake

import (
	"fmt"
	"strings"
	"sync"

	"github.com/lazygophers/utils/randx"
)

// 对象池用于减少内存分配和垃圾回收压力

var (
	// 字符串构建器池
	stringBuilderPool = sync.Pool{
		New: func() interface{} {
			return &strings.Builder{}
		},
	}

	// 字符串切片池
	stringSlicePool = sync.Pool{
		New: func() interface{} {
			return make([]string, 0, 16) // 预分配16个容量
		},
	}

	// 整数切片池
	intSlicePool = sync.Pool{
		New: func() interface{} {
			return make([]int, 0, 16)
		},
	}

	// Faker实例池（用于并发场景）
	fakerPool = sync.Pool{
		New: func() interface{} {
			return New()
		},
	}
)

// getStringBuilder 从池中获取字符串构建器
func getStringBuilder() *strings.Builder {
	sb := stringBuilderPool.Get().(*strings.Builder)
	sb.Reset()
	return sb
}

// putStringBuilder 将字符串构建器归还到池中
func putStringBuilder(sb *strings.Builder) {
	if sb.Cap() > 1024 { // 如果容量过大，不放回池中
		return
	}
	stringBuilderPool.Put(sb)
}

// getStringSlice 从池中获取字符串切片
func getStringSlice() []string {
	slice := stringSlicePool.Get().([]string)
	return slice[:0] // 重置长度为0
}

// putStringSlice 将字符串切片归还到池中
func putStringSlice(slice []string) {
	if cap(slice) > 1024 { // 如果容量过大，不放回池中
		return
	}
	stringSlicePool.Put(slice)
}

// getIntSlice 从池中获取整数切片
func getIntSlice() []int {
	slice := intSlicePool.Get().([]int)
	return slice[:0] // 重置长度为0
}

// putIntSlice 将整数切片归还到池中
func putIntSlice(slice []int) {
	if cap(slice) > 1024 { // 如果容量过大，不放回池中
		return
	}
	intSlicePool.Put(slice)
}

// GetFaker 从池中获取Faker实例（并发安全）
func GetFaker() *Faker {
	faker := fakerPool.Get().(*Faker)
	// 重置统计数据
	faker.stats.Lock()
	faker.stats.callCount = 0
	faker.stats.cacheHits = 0
	faker.stats.generatedData = 0
	faker.stats.Unlock()
	return faker
}

// PutFaker 将Faker实例归还到池中
func PutFaker(faker *Faker) {
	if faker == nil {
		return
	}

	// 清空缓存以防止内存泄漏
	faker.ClearCache()
	fakerPool.Put(faker)
}

// WithPooledFaker 使用池化的Faker执行操作
func WithPooledFaker(fn func(*Faker)) {
	faker := GetFaker()
	defer PutFaker(faker)
	fn(faker)
}

// ParallelGenerate 并行生成数据
func ParallelGenerate[T any](count int, generator func(*Faker) T) []T {
	if count <= 0 {
		return []T{}
	}

	// 计算并发数，基于CPU核心数和数据量
	const maxGoroutines = 8
	goroutines := count
	if goroutines > maxGoroutines {
		goroutines = maxGoroutines
	}
	if goroutines > count {
		goroutines = count
	}

	results := make([]T, count)

	if goroutines <= 1 {
		// 单线程处理
		faker := GetFaker()
		defer PutFaker(faker)

		for i := 0; i < count; i++ {
			results[i] = generator(faker)
		}
		return results
	}

	// 多线程处理
	var wg sync.WaitGroup
	itemsPerGoroutine := count / goroutines

	for g := 0; g < goroutines; g++ {
		wg.Add(1)

		start := g * itemsPerGoroutine
		end := start + itemsPerGoroutine
		if g == goroutines-1 {
			end = count // 最后一个协程处理剩余的所有项
		}

		go func(start, end int) {
			defer wg.Done()

			faker := GetFaker()
			defer PutFaker(faker)

			for i := start; i < end; i++ {
				results[i] = generator(faker)
			}
		}(start, end)
	}

	wg.Wait()
	return results
}

// BatchGenerate 批量生成，使用对象池优化
func BatchGenerate[T any](count int, generator func() T) []T {
	if count <= 0 {
		return []T{}
	}

	results := make([]T, count)
	for i := 0; i < count; i++ {
		results[i] = generator()
	}

	return results
}

// ConcurrentGenerate 并发安全的批量生成
func ConcurrentGenerate[T any](count int, generator func(*Faker) T) []T {
	return ParallelGenerate(count, generator)
}

// 优化后的批量生成函数，重写现有的批量方法

// BatchNamesOptimized 优化的批量名字生成
func (f *Faker) BatchNamesOptimized(count int) []string {
	if count <= 0 {
		return []string{}
	}

	// 预先计算需要的数据
	firstNames, firstWeights, _ := getWeightedItems(f.language, "names",
		func() string {
			if f.gender == GenderMale {
				return "first_male"
			} else if f.gender == GenderFemale {
				return "first_female"
			}
			// 随机选择
			if randx.Bool() {
				return "first_male"
			}
			return "first_female"
		}())

	lastNames, lastWeights, _ := getWeightedItems(f.language, "names", "last")

	// 如果数据加载失败，使用默认方式
	if len(firstNames) == 0 || len(lastNames) == 0 {
		return f.batchGenerate(count, f.Name)
	}

	results := make([]string, count)

	// 批量生成，减少函数调用开销
	for i := 0; i < count; i++ {
		firstName := randx.WeightedChoose(firstNames, firstWeights)
		lastName := randx.WeightedChoose(lastNames, lastWeights)

		// 根据语言决定姓名顺序
		switch f.language {
		case LanguageChineseSimplified, LanguageChineseTraditional:
			results[i] = lastName + firstName
		default:
			results[i] = firstName + " " + lastName
		}
	}

	f.stats.Lock()
	f.stats.callCount++
	f.stats.generatedData += int64(count)
	f.stats.Unlock()

	return results
}

// BatchEmailsOptimized 优化的批量邮箱生成
func (f *Faker) BatchEmailsOptimized(count int) []string {
	if count <= 0 {
		return []string{}
	}

	// 预生成域名列表
	domains := []string{
		"gmail.com", "yahoo.com", "hotmail.com", "outlook.com", "icloud.com",
		"aol.com", "mail.com", "protonmail.com", "yandex.com", "zoho.com",
		"live.com", "msn.com", "qq.com", "163.com", "126.com",
		"sina.com", "sohu.com", "yeah.net", "foxmail.com", "139.com",
	}

	results := make([]string, count)
	sb := getStringBuilder()
	defer putStringBuilder(sb)

	for i := 0; i < count; i++ {
		sb.Reset()

		// 生成用户名
		firstName := strings.ToLower(f.FirstName())
		lastName := strings.ToLower(f.LastName())

		// 简化用户名生成逻辑
		patterns := []int{0, 1, 2, 3} // 0: first.last, 1: first_last, 2: firstlast, 3: first123
		pattern := randx.Choose(patterns)

		switch pattern {
		case 0:
			sb.WriteString(firstName)
			sb.WriteByte('.')
			sb.WriteString(lastName)
		case 1:
			sb.WriteString(firstName)
			sb.WriteByte('_')
			sb.WriteString(lastName)
		case 2:
			sb.WriteString(firstName)
			sb.WriteString(lastName)
		case 3:
			sb.WriteString(firstName)
			sb.WriteString(fmt.Sprintf("%d", randx.Intn(999)+1))
		}

		sb.WriteByte('@')
		sb.WriteString(randx.Choose(domains))

		results[i] = sb.String()
	}

	f.stats.Lock()
	f.stats.callCount++
	f.stats.generatedData += int64(count)
	f.stats.Unlock()

	return results
}

// PoolStats 获取对象池统计信息
type PoolStats struct {
	StringBuilderPoolSize int `json:"string_builder_pool_size"`
	StringSlicePoolSize   int `json:"string_slice_pool_size"`
	IntSlicePoolSize      int `json:"int_slice_pool_size"`
	FakerPoolSize         int `json:"faker_pool_size"`
}

// GetPoolStats 获取对象池统计信息（调试用）
func GetPoolStats() PoolStats {
	return PoolStats{
		StringBuilderPoolSize: getPoolSize(&stringBuilderPool),
		StringSlicePoolSize:   getPoolSize(&stringSlicePool),
		IntSlicePoolSize:      getPoolSize(&intSlicePool),
		FakerPoolSize:         getPoolSize(&fakerPool),
	}
}

// getPoolSize 获取池中对象数量的近似值（仅用于调试）
func getPoolSize(pool *sync.Pool) int {
	// 这是一个近似的方法，因为sync.Pool没有提供直接获取大小的方法
	count := 0
	for i := 0; i < 10; i++ { // 尝试获取10次
		obj := pool.Get()
		if obj != nil {
			count++
			pool.Put(obj)
		} else {
			break
		}
	}
	return count
}

// WarmupPools 预热对象池
func WarmupPools() {
	// 预创建一些对象到池中
	for i := 0; i < 10; i++ {
		sb := &strings.Builder{}
		sb.Grow(128)
		stringBuilderPool.Put(sb)

		slice := make([]string, 0, 16)
		stringSlicePool.Put(slice)

		intSlice := make([]int, 0, 16)
		intSlicePool.Put(intSlice)

		faker := New()
		fakerPool.Put(faker)
	}
}
