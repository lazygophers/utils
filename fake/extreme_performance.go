package fake

import (
	"sync/atomic"
	"unsafe"
)

// ExtremePerformanceFaker 极限性能版本 - 追求绝对的性能极限
type ExtremePerformanceFaker struct {
	// 预编译的姓名组合 - 完全消除运行时计算
	precomputedNames []string
	nameCount        int64

	// 原子计数器
	counter int64

	// 预计算的索引序列 - 使用数学序列避免随机数生成
	indexSequence []int
	sequenceLen   int
}

// NewExtremePerformance 创建极限性能生成器
func NewExtremePerformance() *ExtremePerformanceFaker {
	ep := &ExtremePerformanceFaker{}
	ep.precomputeAllCombinations()
	ep.generateIndexSequence()
	return ep
}

// precomputeAllCombinations 预计算所有可能的姓名组合
func (ep *ExtremePerformanceFaker) precomputeAllCombinations() {
	// 高频英文姓名 - 只保留最核心的
	firstNames := []string{
		"James", "John", "Robert", "Michael", "William", "David", "Richard", "Joseph",
		"Mary", "Patricia", "Jennifer", "Linda", "Elizabeth", "Barbara", "Susan", "Jessica",
	}

	lastNames := []string{
		"Smith", "Johnson", "Williams", "Brown", "Jones", "Garcia", "Miller", "Davis",
		"Rodriguez", "Martinez", "Hernandez", "Lopez", "Wilson", "Anderson", "Thomas", "Taylor",
	}

	// 预计算所有组合 - 16x16 = 256 个组合
	combinations := make([]string, 0, len(firstNames)*len(lastNames))

	for _, first := range firstNames {
		for _, last := range lastNames {
			// 使用最快的字符串拼接方法
			totalLen := len(first) + 1 + len(last)
			bytes := make([]byte, totalLen)
			copy(bytes, first)
			bytes[len(first)] = ' '
			copy(bytes[len(first)+1:], last)
			combinations = append(combinations, *(*string)(unsafe.Pointer(&bytes)))
		}
	}

	ep.precomputedNames = combinations
	ep.nameCount = int64(len(combinations))
}

// generateIndexSequence 生成伪随机索引序列
func (ep *ExtremePerformanceFaker) generateIndexSequence() {
	// 使用数学序列模拟随机性，避免真随机数生成的开销
	// 线性同余生成器的简化版本
	const sequenceSize = 1024
	sequence := make([]int, sequenceSize)

	seed := 1
	for i := 0; i < sequenceSize; i++ {
		seed = (seed*1664525 + 1013904223) & 0x7FFFFFFF // LCG算法
		sequence[i] = seed % int(ep.nameCount)
	}

	ep.indexSequence = sequence
	ep.sequenceLen = sequenceSize
}

// ExtremeName 极限性能姓名生成 - 目标 < 10ns
func (ep *ExtremePerformanceFaker) ExtremeName() string {
	// 原子递增计数器
	count := atomic.AddInt64(&ep.counter, 1)

	// 直接查表，避免任何计算
	seqIndex := count & 1023 // 等价于 count % 1024，但使用位运算更快
	nameIndex := ep.indexSequence[seqIndex]

	// 直接返回预计算的字符串
	return ep.precomputedNames[nameIndex]
}

// ZeroAllocExtremeName 零分配极限版本
func (ep *ExtremePerformanceFaker) ZeroAllocExtremeName() string {
	count := atomic.AddInt64(&ep.counter, 1)

	// 使用更简单的伪随机
	index := (count*17 + 7) % ep.nameCount // 使用质数进行简单散列

	return ep.precomputedNames[index]
}

// BatchExtreme 批量极限性能生成
func (ep *ExtremePerformanceFaker) BatchExtreme(count int) []string {
	names := make([]string, count)

	// 获取起始计数器值
	startCount := atomic.AddInt64(&ep.counter, int64(count)) - int64(count)

	// 批量填充，避免原子操作
	for i := 0; i < count; i++ {
		seqIndex := (startCount + int64(i)) & 1023
		nameIndex := ep.indexSequence[seqIndex]
		names[i] = ep.precomputedNames[nameIndex]
	}

	return names
}

// UltraCompactFaker 超紧凑版本 - 最小内存占用
type UltraCompactFaker struct {
	// 只保留4个最高频的姓名组合
	name1, name2, name3, name4 string
	counter                    uint32
}

// NewUltraCompact 创建超紧凑生成器
func NewUltraCompact() *UltraCompactFaker {
	return &UltraCompactFaker{
		name1: "John Smith",
		name2: "Mary Johnson",
		name3: "James Williams",
		name4: "Patricia Brown",
	}
}

// CompactName 超紧凑姓名生成 - 最小内存访问
func (uc *UltraCompactFaker) CompactName() string {
	count := atomic.AddUint32(&uc.counter, 1)

	// 2位选择器，4个选项
	switch count & 3 {
	case 0:
		return uc.name1
	case 1:
		return uc.name2
	case 2:
		return uc.name3
	default:
		return uc.name4
	}
}

// InlineFaker 内联优化版本 - 完全内联计算
type InlineFaker struct {
	counter uint64
}

// NewInline 创建内联生成器
func NewInline() *InlineFaker {
	return &InlineFaker{}
}

// InlineName 完全内联的姓名生成
func (inf *InlineFaker) InlineName() string {
	count := atomic.AddUint64(&inf.counter, 1)

	// 使用位运算直接构造索引
	firstIdx := (count >> 2) & 7 // 3位 = 8个选择
	lastIdx := count & 7         // 3位 = 8个选择

	// 内联字符串数组 - 编译时常量
	var firstName, lastName string

	switch firstIdx {
	case 0:
		firstName = "James"
	case 1:
		firstName = "John"
	case 2:
		firstName = "Robert"
	case 3:
		firstName = "Michael"
	case 4:
		firstName = "Mary"
	case 5:
		firstName = "Patricia"
	case 6:
		firstName = "Jennifer"
	default:
		firstName = "Linda"
	}

	switch lastIdx {
	case 0:
		lastName = "Smith"
	case 1:
		lastName = "Johnson"
	case 2:
		lastName = "Williams"
	case 3:
		lastName = "Brown"
	case 4:
		lastName = "Jones"
	case 5:
		lastName = "Garcia"
	case 6:
		lastName = "Miller"
	default:
		lastName = "Davis"
	}

	// 最快的字符串拼接
	totalLen := len(firstName) + 1 + len(lastName)
	bytes := make([]byte, totalLen)
	copy(bytes, firstName)
	bytes[len(firstName)] = ' '
	copy(bytes[len(firstName)+1:], lastName)

	return *(*string)(unsafe.Pointer(&bytes))
}

// AssemblyOptimizedFaker 汇编级优化（Go实现）
type AssemblyOptimizedFaker struct {
	// 使用固定大小数组减少指针间接
	names   [256]string
	counter uint64
}

// NewAssemblyOptimized 创建汇编级优化生成器
func NewAssemblyOptimized() *AssemblyOptimizedFaker {
	ao := &AssemblyOptimizedFaker{}

	// 预填充固定数组
	firstNames := [8]string{"James", "John", "Robert", "Michael", "Mary", "Patricia", "Jennifer", "Linda"}
	lastNames := [8]string{"Smith", "Johnson", "Williams", "Brown", "Jones", "Garcia", "Miller", "Davis"}

	idx := 0
	for i := 0; i < 8; i++ {
		for j := 0; j < 8; j++ {
			totalLen := len(firstNames[i]) + 1 + len(lastNames[j])
			bytes := make([]byte, totalLen)
			copy(bytes, firstNames[i])
			bytes[len(firstNames[i])] = ' '
			copy(bytes[len(firstNames[i])+1:], lastNames[j])
			ao.names[idx] = *(*string)(unsafe.Pointer(&bytes))
			idx++
		}
	}

	// 填充剩余位置（重复模式）
	for idx < 256 {
		ao.names[idx] = ao.names[idx-64]
		idx++
	}

	return ao
}

// AssemblyName 汇编级优化姓名生成
func (ao *AssemblyOptimizedFaker) AssemblyName() string {
	count := atomic.AddUint64(&ao.counter, 1)

	// 8位掩码直接索引
	index := count & 255

	// 直接数组访问 - 编译器会优化为最优汇编
	return ao.names[index]
}

// MemoryMappedFaker 内存映射优化版本
type MemoryMappedFaker struct {
	// 使用unsafe直接访问内存块
	nameData unsafe.Pointer
	offsets  [64]int
	lengths  [64]int
	counter  uint64
}

// NewMemoryMapped 创建内存映射生成器
func NewMemoryMapped() *MemoryMappedFaker {
	mm := &MemoryMappedFaker{}
	mm.setupMemoryMappedData()
	return mm
}

// setupMemoryMappedData 设置内存映射数据
func (mm *MemoryMappedFaker) setupMemoryMappedData() {
	names := []string{
		"James Smith", "John Johnson", "Robert Williams", "Michael Brown",
		"Mary Jones", "Patricia Garcia", "Jennifer Miller", "Linda Davis",
		"James Smith", "John Johnson", "Robert Williams", "Michael Brown",
		"Mary Jones", "Patricia Garcia", "Jennifer Miller", "Linda Davis",
		"James Smith", "John Johnson", "Robert Williams", "Michael Brown",
		"Mary Jones", "Patricia Garcia", "Jennifer Miller", "Linda Davis",
		"James Smith", "John Johnson", "Robert Williams", "Michael Brown",
		"Mary Jones", "Patricia Garcia", "Jennifer Miller", "Linda Davis",
		"James Smith", "John Johnson", "Robert Williams", "Michael Brown",
		"Mary Jones", "Patricia Garcia", "Jennifer Miller", "Linda Davis",
		"James Smith", "John Johnson", "Robert Williams", "Michael Brown",
		"Mary Jones", "Patricia Garcia", "Jennifer Miller", "Linda Davis",
		"James Smith", "John Johnson", "Robert Williams", "Michael Brown",
		"Mary Jones", "Patricia Garcia", "Jennifer Miller", "Linda Davis",
		"James Smith", "John Johnson", "Robert Williams", "Michael Brown",
		"Mary Jones", "Patricia Garcia", "Jennifer Miller", "Linda Davis",
	}

	// 计算总长度
	totalLen := 0
	for _, name := range names {
		totalLen += len(name)
	}

	// 分配连续内存块
	data := make([]byte, totalLen)
	offset := 0

	for i, name := range names {
		copy(data[offset:], name)
		mm.offsets[i] = offset
		mm.lengths[i] = len(name)
		offset += len(name)
	}

	mm.nameData = unsafe.Pointer(&data[0])
}

// MemoryMappedName 内存映射姓名生成
func (mm *MemoryMappedFaker) MemoryMappedName() string {
	count := atomic.AddUint64(&mm.counter, 1)

	// 6位索引 = 64个选项
	index := count & 63

	// 直接从内存块构造字符串
	offset := mm.offsets[index]
	length := mm.lengths[index]

	// 使用unsafe直接从内存块创建字符串
	dataPtr := unsafe.Pointer(uintptr(mm.nameData) + uintptr(offset))
	return *(*string)(unsafe.Pointer(&struct {
		data unsafe.Pointer
		len  int
	}{dataPtr, length}))
}

// 全局极限性能实例
var (
	globalExtreme      *ExtremePerformanceFaker
	globalCompact      *UltraCompactFaker
	globalInline       *InlineFaker
	globalAssembly     *AssemblyOptimizedFaker
	globalMemoryMapped *MemoryMappedFaker
)

func init() {
	globalExtreme = NewExtremePerformance()
	globalCompact = NewUltraCompact()
	globalInline = NewInline()
	globalAssembly = NewAssemblyOptimized()
	globalMemoryMapped = NewMemoryMapped()
}

// 全局极限性能函数
func ExtremeName() string {
	return globalExtreme.ExtremeName()
}

func CompactName() string {
	return globalCompact.CompactName()
}

func InlineName() string {
	return globalInline.InlineName()
}

func AssemblyName() string {
	return globalAssembly.AssemblyName()
}

func MemoryMappedName() string {
	return globalMemoryMapped.MemoryMappedName()
}
