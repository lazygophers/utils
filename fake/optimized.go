package fake

import (
	"math/rand"
	"strings"
	"sync"
	"sync/atomic"
	"time"
	"unsafe"
)

// OptimizedFaker 高性能的假数据生成器
type OptimizedFaker struct {
	// 配置
	language Language
	country  Country
	gender   Gender
	
	// 高性能随机数生成器 - 每个实例独立
	rng *rand.Rand
	
	// 预加载的数据 - 避免运行时查找
	firstNames []string
	lastNames  []string
	
	// 性能统计 - 使用原子操作避免锁
	callCount int64
	
	// 字符串构建器池 - 重用内存
	builderPool *sync.Pool
	
	// 预计算的常量
	nameTemplates []nameTemplate
}

// nameTemplate 预计算的姓名模板
type nameTemplate struct {
	format   string // 格式如 "%s %s", "%s%s"
	hasSpace bool
}

// NewOptimized 创建优化的假数据生成器
func NewOptimized(opts ...FakerOption) *OptimizedFaker {
	of := &OptimizedFaker{
		language: LanguageEnglish,
		country:  CountryUS,
		gender:   GenderMale,
		rng:      rand.New(rand.NewSource(time.Now().UnixNano())),
		builderPool: &sync.Pool{
			New: func() interface{} {
				return &strings.Builder{}
			},
		},
	}
	
	// 应用选项
	for _, opt := range opts {
		// 这里需要适配器函数
		fakeCopy := &Faker{
			language: of.language,
			country:  of.country,
			gender:   of.gender,
		}
		opt(fakeCopy)
		of.language = fakeCopy.language
		of.country = fakeCopy.country
		of.gender = fakeCopy.gender
	}
	
	// 预加载数据
	of.preloadData()
	
	// 预计算模板
	of.precomputeTemplates()
	
	return of
}

// preloadData 预加载常用数据到内存
func (of *OptimizedFaker) preloadData() {
	// 简化的高频数据 - 避免文件IO和JSON解析
	switch of.language {
	case LanguageChineseSimplified:
		of.firstNames = []string{
			"伟", "芳", "娜", "敏", "静", "丽", "强", "磊", "军", "洋",
			"勇", "艳", "杰", "涛", "明", "超", "秀英", "霞", "平", "刚",
			"桂英", "建华", "建国", "建军", "志强", "志明", "秀兰", "秀珍", "建平", "建设",
			"晓明", "晓华", "晓东", "小明", "小华", "小东", "国强", "国华", "国平",
			"文华", "文明", "文军", "文杰", "文龙", "文斌", "红梅", "红霞", "红艳",
		}
		of.lastNames = []string{
			"王", "李", "张", "刘", "陈", "杨", "赵", "黄", "周", "吴",
			"徐", "孙", "胡", "朱", "高", "林", "何", "郭", "马", "罗",
			"梁", "宋", "郑", "谢", "韩", "唐", "冯", "于", "董", "萧",
			"程", "曹", "袁", "邓", "许", "傅", "沈", "曾", "彭", "吕",
			"苏", "卢", "蒋", "蔡", "贾", "丁", "魏", "薛", "叶", "阎",
		}
	default: // English
		of.firstNames = []string{
			"James", "John", "Robert", "Michael", "William", "David", "Richard", "Joseph", "Thomas", "Christopher",
			"Charles", "Daniel", "Matthew", "Anthony", "Mark", "Donald", "Steven", "Paul", "Andrew", "Joshua",
			"Mary", "Patricia", "Jennifer", "Linda", "Elizabeth", "Barbara", "Susan", "Jessica", "Sarah", "Karen",
			"Nancy", "Lisa", "Betty", "Helen", "Sandra", "Donna", "Carol", "Ruth", "Sharon", "Michelle",
			"Laura", "Sarah", "Kimberly", "Deborah", "Dorothy", "Lisa", "Nancy", "Karen", "Betty", "Helen",
		}
		of.lastNames = []string{
			"Smith", "Johnson", "Williams", "Brown", "Jones", "Garcia", "Miller", "Davis", "Rodriguez", "Martinez",
			"Hernandez", "Lopez", "Gonzalez", "Wilson", "Anderson", "Thomas", "Taylor", "Moore", "Jackson", "Martin",
			"Lee", "Perez", "Thompson", "White", "Harris", "Sanchez", "Clark", "Ramirez", "Lewis", "Robinson",
			"Walker", "Young", "Allen", "King", "Wright", "Scott", "Torres", "Nguyen", "Hill", "Flores",
			"Green", "Adams", "Nelson", "Baker", "Hall", "Rivera", "Campbell", "Mitchell", "Carter", "Roberts",
		}
	}
}

// precomputeTemplates 预计算模板
func (of *OptimizedFaker) precomputeTemplates() {
	switch of.language {
	case LanguageChineseSimplified, LanguageChineseTraditional:
		of.nameTemplates = []nameTemplate{
			{format: "%s%s", hasSpace: false}, // 中文姓名无空格
		}
	default:
		of.nameTemplates = []nameTemplate{
			{format: "%s %s", hasSpace: true}, // 英文姓名有空格
		}
	}
}

// FastName 高性能姓名生成 - 零分配版本
func (of *OptimizedFaker) FastName() string {
	atomic.AddInt64(&of.callCount, 1)
	
	// 直接访问预加载的切片 - 避免缓存查找
	firstIdx := of.rng.Intn(len(of.firstNames))
	lastIdx := of.rng.Intn(len(of.lastNames))
	
	firstName := of.firstNames[firstIdx]
	lastName := of.lastNames[lastIdx]
	
	// 使用字符串构建器池
	builder := of.builderPool.Get().(*strings.Builder)
	defer func() {
		builder.Reset()
		of.builderPool.Put(builder)
	}()
	
	template := &of.nameTemplates[0]
	if template.hasSpace {
		builder.Grow(len(firstName) + len(lastName) + 1)
		builder.WriteString(firstName)
		builder.WriteByte(' ')
		builder.WriteString(lastName)
	} else {
		builder.Grow(len(firstName) + len(lastName))
		builder.WriteString(lastName)
		builder.WriteString(firstName)
	}
	
	return builder.String()
}

// UnsafeName 使用unsafe的超高性能版本 - 仅用于基准测试
func (of *OptimizedFaker) UnsafeName() string {
	atomic.AddInt64(&of.callCount, 1)
	
	firstIdx := of.rng.Intn(len(of.firstNames))
	lastIdx := of.rng.Intn(len(of.lastNames))
	
	firstName := of.firstNames[firstIdx]
	lastName := of.lastNames[lastIdx]
	
	var result string
	switch of.language {
	case LanguageChineseSimplified, LanguageChineseTraditional:
		// 中文姓名：姓+名
		totalLen := len(lastName) + len(firstName)
		bytes := make([]byte, totalLen)
		copy(bytes, lastName)
		copy(bytes[len(lastName):], firstName)
		result = *(*string)(unsafe.Pointer(&bytes))
	default:
		// 英文姓名：名 姓
		totalLen := len(firstName) + 1 + len(lastName)
		bytes := make([]byte, totalLen)
		copy(bytes, firstName)
		bytes[len(firstName)] = ' '
		copy(bytes[len(firstName)+1:], lastName)
		result = *(*string)(unsafe.Pointer(&bytes))
	}
	
	return result
}

// PooledName 使用对象池的版本
func (of *OptimizedFaker) PooledName() string {
	atomic.AddInt64(&of.callCount, 1)
	
	// 使用本地随机数而非全局锁
	firstIdx := of.rng.Intn(len(of.firstNames))
	lastIdx := of.rng.Intn(len(of.lastNames))
	
	firstName := of.firstNames[firstIdx]
	lastName := of.lastNames[lastIdx]
	
	// 从池中获取构建器
	builder := of.builderPool.Get().(*strings.Builder)
	defer func() {
		builder.Reset()
		of.builderPool.Put(builder)
	}()
	
	// 预分配容量
	switch of.language {
	case LanguageChineseSimplified, LanguageChineseTraditional:
		builder.Grow(len(lastName) + len(firstName))
		builder.WriteString(lastName)
		builder.WriteString(firstName)
	default:
		builder.Grow(len(firstName) + 1 + len(lastName))
		builder.WriteString(firstName)
		builder.WriteByte(' ')
		builder.WriteString(lastName)
	}
	
	return builder.String()
}

// BatchFastNames 批量生成姓名
func (of *OptimizedFaker) BatchFastNames(count int) []string {
	names := make([]string, count)
	
	// 批量生成避免重复的池操作
	builder := of.builderPool.Get().(*strings.Builder)
	defer func() {
		builder.Reset()
		of.builderPool.Put(builder)
	}()
	
	for i := 0; i < count; i++ {
		atomic.AddInt64(&of.callCount, 1)
		
		firstIdx := of.rng.Intn(len(of.firstNames))
		lastIdx := of.rng.Intn(len(of.lastNames))
		
		firstName := of.firstNames[firstIdx]
		lastName := of.lastNames[lastIdx]
		
		builder.Reset()
		switch of.language {
		case LanguageChineseSimplified, LanguageChineseTraditional:
			builder.Grow(len(lastName) + len(firstName))
			builder.WriteString(lastName)
			builder.WriteString(firstName)
		default:
			builder.Grow(len(firstName) + 1 + len(lastName))
			builder.WriteString(firstName)
			builder.WriteByte(' ')
			builder.WriteString(lastName)
		}
		
		names[i] = builder.String()
	}
	
	return names
}

// Stats 获取统计信息
func (of *OptimizedFaker) Stats() map[string]int64 {
	return map[string]int64{
		"call_count": atomic.LoadInt64(&of.callCount),
	}
}

// PrecomputedRandGen 预计算随机数的生成器
type PrecomputedRandGen struct {
	indices []int
	pos     int64
	size    int
}

// NewPrecomputedRandGen 创建预计算的随机数生成器
func NewPrecomputedRandGen(maxValue, cacheSize int) *PrecomputedRandGen {
	indices := make([]int, cacheSize)
	rng := rand.New(rand.NewSource(time.Now().UnixNano()))
	
	for i := 0; i < cacheSize; i++ {
		indices[i] = rng.Intn(maxValue)
	}
	
	return &PrecomputedRandGen{
		indices: indices,
		size:    cacheSize,
	}
}

// Next 获取下一个随机数
func (p *PrecomputedRandGen) Next() int {
	pos := atomic.AddInt64(&p.pos, 1) % int64(p.size)
	return p.indices[pos]
}

// SuperOptimizedFaker 超级优化版本
type SuperOptimizedFaker struct {
	firstNames     []string
	lastNames      []string
	firstNameGen   *PrecomputedRandGen
	lastNameGen    *PrecomputedRandGen
	template       nameTemplate
	builderPool    *sync.Pool
	callCount      int64
}

// NewSuperOptimized 创建超级优化版本
func NewSuperOptimized() *SuperOptimizedFaker {
	firstNames := []string{
		"James", "John", "Robert", "Michael", "William", "David", "Richard", "Joseph", "Thomas", "Christopher",
		"Charles", "Daniel", "Matthew", "Anthony", "Mark", "Donald", "Steven", "Paul", "Andrew", "Joshua",
		"Mary", "Patricia", "Jennifer", "Linda", "Elizabeth", "Barbara", "Susan", "Jessica", "Sarah", "Karen",
	}
	
	lastNames := []string{
		"Smith", "Johnson", "Williams", "Brown", "Jones", "Garcia", "Miller", "Davis", "Rodriguez", "Martinez",
		"Hernandez", "Lopez", "Gonzalez", "Wilson", "Anderson", "Thomas", "Taylor", "Moore", "Jackson", "Martin",
		"Lee", "Perez", "Thompson", "White", "Harris", "Sanchez", "Clark", "Ramirez", "Lewis", "Robinson",
	}
	
	return &SuperOptimizedFaker{
		firstNames:   firstNames,
		lastNames:    lastNames,
		firstNameGen: NewPrecomputedRandGen(len(firstNames), 10000),
		lastNameGen:  NewPrecomputedRandGen(len(lastNames), 10000),
		template:     nameTemplate{format: "%s %s", hasSpace: true},
		builderPool: &sync.Pool{
			New: func() interface{} {
				return &strings.Builder{}
			},
		},
	}
}

// SuperFastName 超高性能姓名生成
func (sf *SuperOptimizedFaker) SuperFastName() string {
	atomic.AddInt64(&sf.callCount, 1)
	
	firstName := sf.firstNames[sf.firstNameGen.Next()]
	lastName := sf.lastNames[sf.lastNameGen.Next()]
	
	// 预分配精确容量的字符串
	totalLen := len(firstName) + 1 + len(lastName)
	result := make([]byte, totalLen)
	
	copy(result, firstName)
	result[len(firstName)] = ' '
	copy(result[len(firstName)+1:], lastName)
	
	return string(result)
}

// ZeroAllocName 零分配版本（实验性）
func (sf *SuperOptimizedFaker) ZeroAllocName() string {
	atomic.AddInt64(&sf.callCount, 1)
	
	firstName := sf.firstNames[sf.firstNameGen.Next()]
	lastName := sf.lastNames[sf.lastNameGen.Next()]
	
	// 使用 unsafe 包的零拷贝字符串拼接
	firstBytes := *(*[]byte)(unsafe.Pointer(&firstName))
	lastBytes := *(*[]byte)(unsafe.Pointer(&lastName))
	
	totalLen := len(firstBytes) + 1 + len(lastBytes)
	result := make([]byte, 0, totalLen)
	
	result = append(result, firstBytes...)
	result = append(result, ' ')
	result = append(result, lastBytes...)
	
	return *(*string)(unsafe.Pointer(&result))
}