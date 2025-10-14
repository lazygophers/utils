package fake

import (
	"math/rand"
	"sync"
	"time"
	"unsafe"
)

// fastName 高性能姓名生成实现
func (f *Faker) fastName() string {
	// 获取或初始化高性能数据
	if f.fastData == nil {
		f.initFastData()
	}

	// 选择名字和姓氏
	var firstName, lastName string

	f.fastData.mu.Lock()
	switch f.gender {
	case GenderMale:
		if len(f.fastData.maleNames) > 0 {
			firstName = f.fastData.maleNames[f.fastData.rng.Intn(len(f.fastData.maleNames))]
		}
	case GenderFemale:
		if len(f.fastData.femaleNames) > 0 {
			firstName = f.fastData.femaleNames[f.fastData.rng.Intn(len(f.fastData.femaleNames))]
		}
	default:
		// 随机性别
		if f.fastData.rng.Float32() < 0.5 && len(f.fastData.maleNames) > 0 {
			firstName = f.fastData.maleNames[f.fastData.rng.Intn(len(f.fastData.maleNames))]
		} else if len(f.fastData.femaleNames) > 0 {
			firstName = f.fastData.femaleNames[f.fastData.rng.Intn(len(f.fastData.femaleNames))]
		}
	}

	if len(f.fastData.surnames) > 0 {
		lastName = f.fastData.surnames[f.fastData.rng.Intn(len(f.fastData.surnames))]
	}
	f.fastData.mu.Unlock()

	// 高效字符串拼接
	switch f.language {
	case LanguageChineseSimplified, LanguageChineseTraditional:
		return lastName + firstName
	default:
		// 使用unsafe进行零拷贝拼接
		totalLen := len(firstName) + 1 + len(lastName)
		bytes := make([]byte, totalLen)
		copy(bytes, firstName)
		bytes[len(firstName)] = ' '
		copy(bytes[len(firstName)+1:], lastName)
		return *(*string)(unsafe.Pointer(&bytes))
	}
}

// FastData 高性能数据结构
type FastData struct {
	maleNames   []string
	femaleNames []string
	surnames    []string
	rng         *rand.Rand
	mu          sync.Mutex
	initialized bool
}

// initFastData 初始化高性能数据
func (f *Faker) initFastData() {
	if f.fastData != nil && f.fastData.initialized {
		return
	}

	f.fastData = &FastData{
		rng: rand.New(rand.NewSource(time.Now().UnixNano())),
	}

	// 根据语言加载优化数据
	switch f.language {
	case LanguageChineseSimplified:
		f.fastData.maleNames = []string{
			"伟", "强", "军", "磊", "勇", "峰", "杰", "涛", "明", "超", "辉", "鹏", "华", "龙", "飞",
			"建华", "建国", "志强", "志明", "晓明", "晓华", "文华", "文明", "国强", "国华",
		}
		f.fastData.femaleNames = []string{
			"芳", "娜", "敏", "静", "丽", "艳", "霞", "玲", "燕", "美", "英", "琳", "莉", "欣", "蓉",
			"秀英", "秀兰", "秀珍", "红梅", "红霞", "红艳", "小红", "小美", "小英", "小芳",
		}
		f.fastData.surnames = []string{
			"王", "李", "张", "刘", "陈", "杨", "赵", "黄", "周", "吴", "徐", "孙", "胡", "朱", "高",
			"林", "何", "郭", "马", "罗", "梁", "宋", "郑", "谢", "韩", "唐", "冯", "于", "董",
		}
	default: // English
		f.fastData.maleNames = []string{
			"James", "Robert", "John", "Michael", "William", "David", "Richard", "Joseph", "Thomas", "Christopher",
			"Charles", "Daniel", "Matthew", "Anthony", "Mark", "Donald", "Steven", "Paul", "Andrew", "Kenneth",
		}
		f.fastData.femaleNames = []string{
			"Mary", "Patricia", "Jennifer", "Linda", "Elizabeth", "Barbara", "Susan", "Jessica", "Sarah", "Karen",
			"Nancy", "Lisa", "Betty", "Helen", "Sandra", "Donna", "Carol", "Ruth", "Sharon", "Michelle",
		}
		f.fastData.surnames = []string{
			"Smith", "Johnson", "Williams", "Brown", "Jones", "Garcia", "Miller", "Davis", "Rodriguez", "Martinez",
			"Hernandez", "Lopez", "Wilson", "Anderson", "Thomas", "Taylor", "Moore", "Jackson", "Martin", "Lee",
		}
	}

	f.fastData.initialized = true
}

// BatchNames 优化的批量生成
func (f *Faker) BatchNames(count int) []string {
	// 使用高性能批量生成
	return f.fastBatchNames(count)
}

// fastBatchNames 高性能批量姓名生成
func (f *Faker) fastBatchNames(count int) []string {
	if f.fastData == nil {
		f.initFastData()
	}

	names := make([]string, count)

	// 预生成随机索引减少锁竞争
	f.fastData.mu.Lock()
	maleIndices := make([]int, count)
	femaleIndices := make([]int, count)
	surnameIndices := make([]int, count)
	genderFlags := make([]bool, count)

	for i := 0; i < count; i++ {
		if len(f.fastData.maleNames) > 0 {
			maleIndices[i] = f.fastData.rng.Intn(len(f.fastData.maleNames))
		}
		if len(f.fastData.femaleNames) > 0 {
			femaleIndices[i] = f.fastData.rng.Intn(len(f.fastData.femaleNames))
		}
		if len(f.fastData.surnames) > 0 {
			surnameIndices[i] = f.fastData.rng.Intn(len(f.fastData.surnames))
		}
		genderFlags[i] = f.fastData.rng.Float32() < 0.5
	}
	f.fastData.mu.Unlock()

	// 无锁生成
	for i := 0; i < count; i++ {
		f.incrementCallCount()

		var firstName, lastName string

		switch f.gender {
		case GenderMale:
			if len(f.fastData.maleNames) > 0 {
				firstName = f.fastData.maleNames[maleIndices[i]]
			}
		case GenderFemale:
			if len(f.fastData.femaleNames) > 0 {
				firstName = f.fastData.femaleNames[femaleIndices[i]]
			}
		default:
			if genderFlags[i] && len(f.fastData.maleNames) > 0 {
				firstName = f.fastData.maleNames[maleIndices[i]]
			} else if len(f.fastData.femaleNames) > 0 {
				firstName = f.fastData.femaleNames[femaleIndices[i]]
			}
		}

		if len(f.fastData.surnames) > 0 {
			lastName = f.fastData.surnames[surnameIndices[i]]
		}

		// 高效拼接
		switch f.language {
		case LanguageChineseSimplified, LanguageChineseTraditional:
			names[i] = lastName + firstName
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
