package fake

import (
	"math/rand"
	"strings"
	"sync"
	"sync/atomic"
	"time"
	"unsafe"
)

// PerformanceFirstFaker 性能优先的实现
type PerformanceFirstFaker struct {
	// 配置
	language Language

	// 最小化的数据集 - 只保留最高频的
	firstNames []string
	lastNames  []string

	// 高性能随机数生成
	rng *rand.Rand
	mu  sync.Mutex

	// 原子计数器
	callCount int64

	// 预计算的索引范围
	firstNamesLen int
	lastNamesLen  int

	// 字符串构建器池
	builderPool *sync.Pool

	// 预分配的字节缓冲区池
	bytePool *sync.Pool
}

// NewPerformanceFirst 创建性能优先的生成器
func NewPerformanceFirst(opts ...FakerOption) *PerformanceFirstFaker {
	pf := &PerformanceFirstFaker{
		language: LanguageEnglish,
		rng:      rand.New(rand.NewSource(time.Now().UnixNano())),
		builderPool: &sync.Pool{
			New: func() interface{} {
				return &strings.Builder{}
			},
		},
		bytePool: &sync.Pool{
			New: func() interface{} {
				return make([]byte, 0, 64) // 预分配64字节
			},
		},
	}

	// 应用选项（只处理语言）
	for _, opt := range opts {
		fakeCopy := &Faker{language: pf.language}
		opt(fakeCopy)
		pf.language = fakeCopy.language
	}

	// 加载最小化数据集
	pf.loadMinimalData()
	pf.firstNamesLen = len(pf.firstNames)
	pf.lastNamesLen = len(pf.lastNames)

	return pf
}

// loadMinimalData 加载最小化的高频数据
func (pf *PerformanceFirstFaker) loadMinimalData() {
	switch pf.language {
	case LanguageChineseSimplified:
		// 只保留最高频的20个名字和20个姓氏
		pf.firstNames = []string{
			"伟", "芳", "娜", "敏", "静", "丽", "强", "磊", "军", "洋",
			"勇", "艳", "杰", "涛", "明", "超", "秀英", "霞", "平", "刚",
		}
		pf.lastNames = []string{
			"王", "李", "张", "刘", "陈", "杨", "赵", "黄", "周", "吴",
			"徐", "孙", "胡", "朱", "高", "林", "何", "郭", "马", "罗",
		}
	default:
		// 英文高频名字
		pf.firstNames = []string{
			"James", "John", "Robert", "Michael", "William", "David", "Richard", "Joseph", "Thomas", "Christopher",
			"Mary", "Patricia", "Jennifer", "Linda", "Elizabeth", "Barbara", "Susan", "Jessica", "Sarah", "Karen",
		}
		pf.lastNames = []string{
			"Smith", "Johnson", "Williams", "Brown", "Jones", "Garcia", "Miller", "Davis", "Rodriguez", "Martinez",
			"Hernandez", "Lopez", "Wilson", "Anderson", "Thomas", "Taylor", "Moore", "Jackson", "Martin", "Lee",
		}
	}
}

// UltraFastName 超高性能姓名生成 - 最少分配
func (pf *PerformanceFirstFaker) UltraFastName() string {
	atomic.AddInt64(&pf.callCount, 1)

	// 直接访问，避免边界检查
	pf.mu.Lock()
	firstIdx := pf.rng.Intn(pf.firstNamesLen)
	lastIdx := pf.rng.Intn(pf.lastNamesLen)
	pf.mu.Unlock()

	firstName := pf.firstNames[firstIdx]
	lastName := pf.lastNames[lastIdx]

	// 直接字节拼接，避免字符串构建器的开销
	var result string
	switch pf.language {
	case LanguageChineseSimplified:
		// 中文姓名拼接
		totalLen := len(lastName) + len(firstName)
		bytes := make([]byte, totalLen)
		copy(bytes, lastName)
		copy(bytes[len(lastName):], firstName)
		result = *(*string)(unsafe.Pointer(&bytes))
	default:
		// 英文姓名拼接
		totalLen := len(firstName) + 1 + len(lastName)
		bytes := make([]byte, totalLen)
		copy(bytes, firstName)
		bytes[len(firstName)] = ' '
		copy(bytes[len(firstName)+1:], lastName)
		result = *(*string)(unsafe.Pointer(&bytes))
	}

	return result
}

// BatchUltraFastNames 批量超高性能生成
func (pf *PerformanceFirstFaker) BatchUltraFastNames(count int) []string {
	names := make([]string, count)

	// 预锁定随机数生成器减少锁开销
	pf.mu.Lock()
	indices := make([][2]int, count)
	for i := 0; i < count; i++ {
		indices[i][0] = pf.rng.Intn(pf.firstNamesLen)
		indices[i][1] = pf.rng.Intn(pf.lastNamesLen)
	}
	pf.mu.Unlock()

	// 无锁生成名字
	for i := 0; i < count; i++ {
		atomic.AddInt64(&pf.callCount, 1)

		firstName := pf.firstNames[indices[i][0]]
		lastName := pf.lastNames[indices[i][1]]

		// 直接字节拼接
		switch pf.language {
		case LanguageChineseSimplified:
			totalLen := len(lastName) + len(firstName)
			bytes := make([]byte, totalLen)
			copy(bytes, lastName)
			copy(bytes[len(lastName):], firstName)
			names[i] = *(*string)(unsafe.Pointer(&bytes))
		default:
			totalLen := len(firstName) + 1 + len(lastName)
			bytes := make([]byte, totalLen)
			copy(bytes, firstName)
			bytes[len(firstName)] = ' '
			copy(bytes[len(firstName)+1:], lastName)
			names[i] = *(*string)(unsafe.Pointer(&bytes))
		}
	}

	return names
}

// PrecomputedFastName 预计算版本 - 使用查表法
func (pf *PerformanceFirstFaker) PrecomputedFastName() string {
	atomic.AddInt64(&pf.callCount, 1)

	// 使用简单的哈希避免随机数生成开销
	count := atomic.LoadInt64(&pf.callCount)
	firstIdx := int(count*7) % pf.firstNamesLen // 使用质数避免规律
	lastIdx := int(count*11) % pf.lastNamesLen  // 使用不同质数

	firstName := pf.firstNames[firstIdx]
	lastName := pf.lastNames[lastIdx]

	// 最快的字符串拼接
	switch pf.language {
	case LanguageChineseSimplified:
		return lastName + firstName
	default:
		return firstName + " " + lastName
	}
}

// NoAllocName 零额外分配版本（实验性）
func (pf *PerformanceFirstFaker) NoAllocName() string {
	atomic.AddInt64(&pf.callCount, 1)

	// 从池中获取字节缓冲区
	bytes := pf.bytePool.Get().([]byte)
	defer pf.bytePool.Put(bytes[:0]) // 重置长度但保持容量

	// 简化的随机选择
	count := atomic.LoadInt64(&pf.callCount)
	firstIdx := int(count*13) % pf.firstNamesLen
	lastIdx := int(count*17) % pf.lastNamesLen

	firstName := pf.firstNames[firstIdx]
	lastName := pf.lastNames[lastIdx]

	// 重用字节缓冲区
	switch pf.language {
	case LanguageChineseSimplified:
		bytes = append(bytes, lastName...)
		bytes = append(bytes, firstName...)
	default:
		bytes = append(bytes, firstName...)
		bytes = append(bytes, ' ')
		bytes = append(bytes, lastName...)
	}

	// 转换为字符串
	return string(bytes)
}

// Stats 获取统计信息
func (pf *PerformanceFirstFaker) Stats() map[string]int64 {
	return map[string]int64{
		"call_count": atomic.LoadInt64(&pf.callCount),
	}
}

// 全局性能优先实例
var (
	globalPerformanceFaker *PerformanceFirstFaker
	globalPerformanceOnce  sync.Once
)

// GetGlobalPerformanceFaker 获取全局性能优先实例
func GetGlobalPerformanceFaker() *PerformanceFirstFaker {
	globalPerformanceOnce.Do(func() {
		globalPerformanceFaker = NewPerformanceFirst()
	})
	return globalPerformanceFaker
}

// UltraFastName 全局超高性能姓名生成函数
func UltraFastName() string {
	return GetGlobalPerformanceFaker().UltraFastName()
}

// PrecomputedFastName 全局预计算快速姓名生成函数
func PrecomputedFastName() string {
	return GetGlobalPerformanceFaker().PrecomputedFastName()
}
