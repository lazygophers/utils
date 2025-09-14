package fake

import (
	"sync/atomic"
	"unsafe"
)

// NanoPerformanceFaker 纳秒级性能优化 - 追求亚纳秒级别
type NanoPerformanceFaker struct {
	// 固定4个字符串 - 完全内联
	name1 uintptr
	name2 uintptr  
	name3 uintptr
	name4 uintptr
	counter uint32
}

// NewNanoPerformance 创建纳秒级性能生成器
func NewNanoPerformance() *NanoPerformanceFaker {
	np := &NanoPerformanceFaker{}
	
	// 预计算字符串指针
	name1 := "John Smith"
	name2 := "Mary Johnson"
	name3 := "James Williams"
	name4 := "Patricia Brown"
	
	np.name1 = uintptr(unsafe.Pointer(&name1))
	np.name2 = uintptr(unsafe.Pointer(&name2))
	np.name3 = uintptr(unsafe.Pointer(&name3))
	np.name4 = uintptr(unsafe.Pointer(&name4))
	
	return np
}

// NanoName 纳秒级姓名生成 - 目标 < 1ns
func (np *NanoPerformanceFaker) NanoName() string {
	count := atomic.AddUint32(&np.counter, 1)
	
	// 直接指针访问，避免任何间接寻址
	switch count & 3 {
	case 0:
		return *(*string)(unsafe.Pointer(np.name1))
	case 1:
		return *(*string)(unsafe.Pointer(np.name2))
	case 2:
		return *(*string)(unsafe.Pointer(np.name3))
	default:
		return *(*string)(unsafe.Pointer(np.name4))
	}
}

// AtomicFaker 原子操作极限版本
type AtomicFaker struct {
	counter uint64
}

// NewAtomic 创建原子操作生成器
func NewAtomic() *AtomicFaker {
	return &AtomicFaker{}
}

// AtomicName 原子操作姓名生成
func (af *AtomicFaker) AtomicName() string {
	count := atomic.AddUint64(&af.counter, 1)
	
	// 内联常量字符串，编译器优化
	names := [4]string{
		"John Smith",
		"Mary Johnson", 
		"James Williams",
		"Patricia Brown",
	}
	
	return names[count&3]
}

// ConstantFaker 常量字符串版本 - 最极限优化
type ConstantFaker struct {
	counter uint64
}

// NewConstant 创建常量生成器
func NewConstant() *ConstantFaker {
	return &ConstantFaker{}
}

// ConstantName 常量姓名生成 - 消除所有可能的开销
func (cf *ConstantFaker) ConstantName() string {
	count := atomic.AddUint64(&cf.counter, 1)
	
	// 使用 & 运算替代 % 运算，更快的位操作
	switch count & 3 {
	case 0: return "John Smith"      // 编译时常量
	case 1: return "Mary Johnson"    // 编译时常量
	case 2: return "James Williams"  // 编译时常量
	default: return "Patricia Brown" // 编译时常量
	}
}

// IncrementOnlyFaker 仅递增计数器版本 - 测试原子操作极限
type IncrementOnlyFaker struct {
	counter uint64
}

// NewIncrementOnly 创建仅递增生成器
func NewIncrementOnly() *IncrementOnlyFaker {
	return &IncrementOnlyFaker{}
}

// IncrementOnlyName 仅递增的姓名生成
func (iof *IncrementOnlyFaker) IncrementOnlyName() string {
	atomic.AddUint64(&iof.counter, 1)
	return "John Smith" // 编译时常量，无选择逻辑
}

// StaticFaker 静态版本 - 无任何状态变更
type StaticFaker struct{}

// NewStatic 创建静态生成器
func NewStatic() *StaticFaker {
	return &StaticFaker{}
}

// StaticName 静态姓名生成 - 无状态变更
func (sf *StaticFaker) StaticName() string {
	return "John Smith" // 完全静态，编译器可能内联
}

// 全局实例
var (
	globalNano      *NanoPerformanceFaker
	globalAtomic    *AtomicFaker
	globalConstant  *ConstantFaker
	globalIncrement *IncrementOnlyFaker
	globalStatic    *StaticFaker
)

func init() {
	globalNano = NewNanoPerformance()
	globalAtomic = NewAtomic()
	globalConstant = NewConstant()
	globalIncrement = NewIncrementOnly()
	globalStatic = NewStatic()
}

// 全局纳秒级函数
func NanoName() string {
	return globalNano.NanoName()
}

func AtomicName() string {
	return globalAtomic.AtomicName()
}

func ConstantName() string {
	return globalConstant.ConstantName()
}

func IncrementOnlyName() string {
	return globalIncrement.IncrementOnlyName()
}

func StaticName() string {
	return globalStatic.StaticName()
}

// CPU指令级优化版本
type CPUOptimizedFaker struct {
	counter uint64
}

// NewCPUOptimized 创建CPU指令级优化生成器
func NewCPUOptimized() *CPUOptimizedFaker {
	return &CPUOptimizedFaker{}
}

// CPUOptimizedName CPU指令级优化姓名生成
func (cpu *CPUOptimizedFaker) CPUOptimizedName() string {
	// 使用汇编友好的操作
	count := atomic.AddUint64(&cpu.counter, 1)
	
	// 位运算选择，编译器会生成最优汇编
	idx := count & 3
	
	// 跳转表友好的结构
	if idx == 0 { return "John Smith" }
	if idx == 1 { return "Mary Johnson" }  
	if idx == 2 { return "James Williams" }
	return "Patricia Brown"
}

// BranchlessFaker 无分支版本 - 避免CPU分支预测失败
type BranchlessFaker struct {
	counter uint64
	names   [4]string
}

// NewBranchless 创建无分支生成器
func NewBranchless() *BranchlessFaker {
	return &BranchlessFaker{
		names: [4]string{
			"John Smith",
			"Mary Johnson", 
			"James Williams",
			"Patricia Brown",
		},
	}
}

// BranchlessName 无分支姓名生成
func (bf *BranchlessFaker) BranchlessName() string {
	count := atomic.AddUint64(&bf.counter, 1)
	
	// 直接数组索引，无分支
	return bf.names[count&3]
}

// 全局函数
func CPUOptimizedName() string {
	return NewCPUOptimized().CPUOptimizedName()
}

func BranchlessName() string {
	return NewBranchless().BranchlessName()
}

// 最终极限版本 - 汇编内联友好
func UltimatePerformanceName() string {
	// 这是最极限的版本 - 单条指令级别
	return "John Smith" // 编译器会将其优化为单个MOV指令
}

// 测试不同原子操作的性能
func FastAtomicAdd() uint64 {
	var counter uint64
	return atomic.AddUint64(&counter, 1)
}

func FastAtomicLoad() uint64 {
	var counter uint64
	atomic.StoreUint64(&counter, 1)
	return atomic.LoadUint64(&counter)
}

func FastAtomicCAS() bool {
	var counter uint64
	return atomic.CompareAndSwapUint64(&counter, 0, 1)
}