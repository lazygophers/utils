package fake

import (
	"fmt"
	"math/rand"
	"strings"
	"sync"
	"sync/atomic"
	"time"
)

// HighPerformanceFaker 高性能实现，替代原有的实现
type HighPerformanceFaker struct {
	// 配置
	language Language
	country  Country
	gender   Gender
	
	// 高性能随机数生成器 - 使用本地实例避免全局锁
	rng *rand.Rand
	mu  sync.Mutex // 保护随机数生成器
	
	// 预加载的高频数据
	maleFirstNames   []string
	femaleFirstNames []string
	lastNames        []string
	emailDomains     []string
	
	// 性能统计 - 原子操作
	callCount int64
	
	// 对象池减少分配
	builderPool *sync.Pool
	
	// 预计算的权重表（用于更好的随机分布）
	firstNameWeights []int
	lastNameWeights  []int
}

// NewHighPerformance 创建高性能假数据生成器
func NewHighPerformance(opts ...FakerOption) *HighPerformanceFaker {
	hf := &HighPerformanceFaker{
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
		fakeCopy := &Faker{
			language: hf.language,
			country:  hf.country,
			gender:   hf.gender,
		}
		opt(fakeCopy)
		hf.language = fakeCopy.language
		hf.country = fakeCopy.country
		hf.gender = fakeCopy.gender
	}
	
	// 预加载数据
	hf.loadHighFrequencyData()
	hf.computeWeights()
	
	return hf
}

// loadHighFrequencyData 加载高频数据
func (hf *HighPerformanceFaker) loadHighFrequencyData() {
	switch hf.language {
	case LanguageChineseSimplified:
		// 中文高频姓名数据
		hf.maleFirstNames = []string{
			"伟", "强", "军", "磊", "勇", "峰", "杰", "涛", "明", "超", "辉", "鹏", "华", "建华", "建国",
			"志强", "志明", "晓明", "晓华", "文华", "文明", "国强", "国华", "建军", "建平", "建设",
			"德华", "德明", "志华", "志刚", "永强", "永明", "家华", "家明", "大伟", "大明", "小明",
			"振华", "振明", "海华", "海明", "春华", "春明", "庆华", "庆明", "立华", "立明", "宇华",
		}
		hf.femaleFirstNames = []string{
			"芳", "娜", "敏", "静", "丽", "艳", "秀英", "霞", "秀兰", "秀珍", "红梅", "红霞", "红艳",
			"玲", "燕", "美", "英", "琳", "莉", "欣", "蓉", "雪", "梅", "兰", "菊", "凤", "玉", "萍",
			"小红", "小美", "小英", "小芳", "小丽", "晓红", "晓美", "晓英", "晓芳", "晓丽", "春红",
			"春美", "春英", "春芳", "春丽", "秋红", "秋美", "秋英", "秋芳", "秋丽", "冬红", "冬美",
		}
		hf.lastNames = []string{
			"王", "李", "张", "刘", "陈", "杨", "赵", "黄", "周", "吴", "徐", "孙", "胡", "朱", "高",
			"林", "何", "郭", "马", "罗", "梁", "宋", "郑", "谢", "韩", "唐", "冯", "于", "董", "萧",
			"程", "曹", "袁", "邓", "许", "傅", "沈", "曾", "彭", "吕", "苏", "卢", "蒋", "蔡", "贾",
			"丁", "魏", "薛", "叶", "阎", "余", "潘", "杜", "戴", "夏", "钟", "汪", "田", "任", "姜",
		}
		hf.emailDomains = []string{
			"@qq.com", "@163.com", "@139.com", "@126.com", "@sohu.com", "@sina.com", "@foxmail.com",
			"@yeah.net", "@hotmail.com", "@gmail.com", "@outlook.com",
		}
		
	default: // English and others
		hf.maleFirstNames = []string{
			"James", "Robert", "John", "Michael", "William", "David", "Richard", "Joseph", "Thomas", "Christopher",
			"Charles", "Daniel", "Matthew", "Anthony", "Mark", "Donald", "Steven", "Paul", "Andrew", "Kenneth",
			"Joshua", "Kevin", "Brian", "George", "Timothy", "Ronald", "Jason", "Edward", "Jeffrey", "Ryan",
			"Jacob", "Gary", "Nicholas", "Eric", "Jonathan", "Stephen", "Larry", "Justin", "Scott", "Brandon",
			"Benjamin", "Samuel", "Gregory", "Alexander", "Patrick", "Frank", "Raymond", "Jack", "Dennis", "Jerry",
		}
		hf.femaleFirstNames = []string{
			"Mary", "Patricia", "Jennifer", "Linda", "Elizabeth", "Barbara", "Susan", "Jessica", "Sarah", "Karen",
			"Nancy", "Lisa", "Betty", "Helen", "Sandra", "Donna", "Carol", "Ruth", "Sharon", "Michelle",
			"Laura", "Sarah", "Kimberly", "Deborah", "Dorothy", "Lisa", "Nancy", "Karen", "Betty", "Helen",
			"Sandra", "Donna", "Carol", "Ruth", "Sharon", "Michelle", "Laura", "Emily", "Kimberly", "Deborah",
			"Amy", "Angela", "Ashley", "Brenda", "Emma", "Olivia", "Cynthia", "Marie", "Janet", "Catherine",
		}
		hf.lastNames = []string{
			"Smith", "Johnson", "Williams", "Brown", "Jones", "Garcia", "Miller", "Davis", "Rodriguez", "Martinez",
			"Hernandez", "Lopez", "Gonzalez", "Wilson", "Anderson", "Thomas", "Taylor", "Moore", "Jackson", "Martin",
			"Lee", "Perez", "Thompson", "White", "Harris", "Sanchez", "Clark", "Ramirez", "Lewis", "Robinson",
			"Walker", "Young", "Allen", "King", "Wright", "Scott", "Torres", "Nguyen", "Hill", "Flores",
			"Green", "Adams", "Nelson", "Baker", "Hall", "Rivera", "Campbell", "Mitchell", "Carter", "Roberts",
		}
		hf.emailDomains = []string{
			"@gmail.com", "@yahoo.com", "@hotmail.com", "@outlook.com", "@company.com", "@example.com",
			"@mail.com", "@email.com", "@protonmail.com", "@icloud.com", "@aol.com",
		}
	}
}

// computeWeights 计算权重分布以提高随机度
func (hf *HighPerformanceFaker) computeWeights() {
	// 创建权重表，让常见名字更频繁，但不会过度集中
	hf.firstNameWeights = make([]int, len(hf.maleFirstNames)+len(hf.femaleFirstNames))
	hf.lastNameWeights = make([]int, len(hf.lastNames))
	
	// 为名字分配权重（前面的权重更高，但差异不会太大）
	totalFirstNames := len(hf.maleFirstNames) + len(hf.femaleFirstNames)
	for i := 0; i < totalFirstNames; i++ {
		// 权重从5递减到1，确保有差异但不会过分集中
		weight := 5 - (i*4)/totalFirstNames
		if weight < 1 {
			weight = 1
		}
		hf.firstNameWeights[i] = weight
	}
	
	// 为姓氏分配权重
	for i := 0; i < len(hf.lastNames); i++ {
		weight := 5 - (i*4)/len(hf.lastNames)
		if weight < 1 {
			weight = 1
		}
		hf.lastNameWeights[i] = weight
	}
}

// weightedRandom 基于权重的随机选择
func (hf *HighPerformanceFaker) weightedRandom(weights []int) int {
	totalWeight := 0
	for _, w := range weights {
		totalWeight += w
	}
	
	hf.mu.Lock()
	r := hf.rng.Intn(totalWeight)
	hf.mu.Unlock()
	
	currentWeight := 0
	for i, w := range weights {
		currentWeight += w
		if r < currentWeight {
			return i
		}
	}
	return len(weights) - 1 // 应该不会到这里
}

// FastName 高性能姓名生成
func (hf *HighPerformanceFaker) FastName() string {
	atomic.AddInt64(&hf.callCount, 1)
	
	// 选择名字
	var firstName string
	switch hf.gender {
	case GenderMale:
		if len(hf.maleFirstNames) > 0 {
			idx := hf.weightedRandom(hf.firstNameWeights[:len(hf.maleFirstNames)])
			firstName = hf.maleFirstNames[idx]
		}
	case GenderFemale:
		if len(hf.femaleFirstNames) > 0 {
			idx := hf.weightedRandom(hf.firstNameWeights[len(hf.maleFirstNames):len(hf.maleFirstNames)+len(hf.femaleFirstNames)])
			firstName = hf.femaleFirstNames[idx]
		}
	default:
		// 随机性别
		hf.mu.Lock()
		isMale := hf.rng.Float32() < 0.5
		hf.mu.Unlock()
		
		if isMale && len(hf.maleFirstNames) > 0 {
			idx := hf.weightedRandom(hf.firstNameWeights[:len(hf.maleFirstNames)])
			firstName = hf.maleFirstNames[idx]
		} else if len(hf.femaleFirstNames) > 0 {
			idx := hf.weightedRandom(hf.firstNameWeights[len(hf.maleFirstNames):len(hf.maleFirstNames)+len(hf.femaleFirstNames)])
			firstName = hf.femaleFirstNames[idx]
		}
	}
	
	// 选择姓氏
	var lastName string
	if len(hf.lastNames) > 0 {
		idx := hf.weightedRandom(hf.lastNameWeights)
		lastName = hf.lastNames[idx]
	}
	
	// 组合姓名
	builder := hf.builderPool.Get().(*strings.Builder)
	defer func() {
		builder.Reset()
		hf.builderPool.Put(builder)
	}()
	
	switch hf.language {
	case LanguageChineseSimplified, LanguageChineseTraditional:
		// 中文：姓+名
		builder.Grow(len(lastName) + len(firstName))
		builder.WriteString(lastName)
		builder.WriteString(firstName)
	default:
		// 英文：名 姓
		builder.Grow(len(firstName) + 1 + len(lastName))
		builder.WriteString(firstName)
		builder.WriteByte(' ')
		builder.WriteString(lastName)
	}
	
	return builder.String()
}

// BatchFastNames 批量生成姓名
func (hf *HighPerformanceFaker) BatchFastNames(count int) []string {
	names := make([]string, count)
	
	// 重用一个builder
	builder := hf.builderPool.Get().(*strings.Builder)
	defer func() {
		builder.Reset()
		hf.builderPool.Put(builder)
	}()
	
	for i := 0; i < count; i++ {
		atomic.AddInt64(&hf.callCount, 1)
		
		// 选择名字（简化逻辑提高性能）
		var firstName string
		if hf.gender == GenderMale && len(hf.maleFirstNames) > 0 {
			idx := hf.weightedRandom(hf.firstNameWeights[:len(hf.maleFirstNames)])
			firstName = hf.maleFirstNames[idx]
		} else if hf.gender == GenderFemale && len(hf.femaleFirstNames) > 0 {
			idx := hf.weightedRandom(hf.firstNameWeights[len(hf.maleFirstNames):])
			firstName = hf.femaleFirstNames[idx]
		} else {
			// 默认随机
			hf.mu.Lock()
			if hf.rng.Float32() < 0.5 && len(hf.maleFirstNames) > 0 {
				idx := hf.rng.Intn(len(hf.maleFirstNames))
				firstName = hf.maleFirstNames[idx]
			} else if len(hf.femaleFirstNames) > 0 {
				idx := hf.rng.Intn(len(hf.femaleFirstNames))
				firstName = hf.femaleFirstNames[idx]
			}
			hf.mu.Unlock()
		}
		
		// 选择姓氏
		var lastName string
		if len(hf.lastNames) > 0 {
			hf.mu.Lock()
			idx := hf.rng.Intn(len(hf.lastNames))
			lastName = hf.lastNames[idx]
			hf.mu.Unlock()
		}
		
		// 组合姓名
		builder.Reset()
		switch hf.language {
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

// FastEmail 高性能邮箱生成
func (hf *HighPerformanceFaker) FastEmail() string {
	atomic.AddInt64(&hf.callCount, 1)
	
	// 生成用户名（基于姓名）
	name := hf.FastName()
	
	builder := hf.builderPool.Get().(*strings.Builder)
	defer func() {
		builder.Reset()
		hf.builderPool.Put(builder)
	}()
	
	// 转换为邮箱用户名
	switch hf.language {
	case LanguageChineseSimplified, LanguageChineseTraditional:
		// 中文名使用拼音或数字
		hf.mu.Lock()
		usernum := hf.rng.Intn(10000)
		hf.mu.Unlock()
		
		builder.WriteString(strings.ToLower(name))
		if usernum > 0 {
			builder.WriteString(fmt.Sprintf("%d", usernum))
		}
	default:
		// 英文名处理
		parts := strings.Split(name, " ")
		if len(parts) >= 2 {
			builder.WriteString(strings.ToLower(parts[0]))
			builder.WriteByte('.')
			builder.WriteString(strings.ToLower(parts[1]))
			
			// 有时添加数字
			hf.mu.Lock()
			if hf.rng.Float32() < 0.3 {
				num := hf.rng.Intn(1000)
				builder.WriteString(fmt.Sprintf("%d", num))
			}
			hf.mu.Unlock()
		} else {
			builder.WriteString(strings.ToLower(name))
		}
	}
	
	// 添加域名
	if len(hf.emailDomains) > 0 {
		hf.mu.Lock()
		domain := hf.emailDomains[hf.rng.Intn(len(hf.emailDomains))]
		hf.mu.Unlock()
		builder.WriteString(domain)
	}
	
	return builder.String()
}

// Stats 获取统计信息
func (hf *HighPerformanceFaker) Stats() map[string]int64 {
	return map[string]int64{
		"call_count": atomic.LoadInt64(&hf.callCount),
	}
}

// Clone 创建克隆实例
func (hf *HighPerformanceFaker) Clone() *HighPerformanceFaker {
	return NewHighPerformance(
		WithLanguage(hf.language),
		WithCountry(hf.country),
		WithGender(hf.gender),
	)
}