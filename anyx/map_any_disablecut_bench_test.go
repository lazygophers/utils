package anyx

import (
	"sync"
	"sync/atomic"
	"testing"
)

// 方案1：原始实现 - 使用 sync/atomic.StoreUint32
func BenchmarkDisableCut_Original_AtomicStore(b *testing.B) {
	m := NewMap(nil).EnableCut(".")
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		m.DisableCut()
	}
}

// 方案2：直接赋值（不使用原子操作，因为 cut 在单线程模式下设置）
// 风险：如果有并发读可能会出现问题
func BenchmarkDisableCut_DirectAssignment(b *testing.B) {
	m := NewMap(nil).EnableCut(".")
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		m.cut = 0
	}
}

// 方案3：使用 atomic.StoreUint64（假设可以改变字段类型）
// 这需要修改结构体定义，仅作对比参考
func BenchmarkDisableCast_Uint32Store(b *testing.B) {
	m := NewMap(nil).EnableCut(".")
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		atomic.StoreUint32(&m.cut, 0)
	}
}

// 方案4：使用 sync.Mutex 保护普通赋值
func BenchmarkDisableCut_MutexProtected(b *testing.B) {
	m := NewMap(nil).EnableCut(".")
	var mu sync.Mutex
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		mu.Lock()
		m.cut = 0
		mu.Unlock()
	}
}

// 方案5：使用 sync.RWMutex 保护普通赋值
func BenchmarkDisableCut_RWMutexProtected(b *testing.B) {
	m := NewMap(nil).EnableCut(".")
	var mu sync.RWMutex
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		mu.Lock()
		m.cut = 0
		mu.Unlock()
	}
}

// 方案6：使用 atomic.Value 包装 bool
func BenchmarkDisableCut_AtomicValueBool(b *testing.B) {
	type MapAnyAlt struct {
		data map[string]interface{}
		mu   sync.RWMutex
		cut  atomic.Value // bool
		seq  atomic.Value
	}

	m := &MapAnyAlt{
		data: make(map[string]interface{}),
		cut:  atomic.Value{},
	}
	m.cut.Store(true)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		m.cut.Store(false)
	}
}

// 方案7：使用 atomic.Bool（Go 1.19+）
func BenchmarkDisableCut_AtomicBool(b *testing.B) {
	type MapAnyAlt struct {
		data map[string]interface{}
		mu   sync.RWMutex
		cut  atomic.Bool
		seq  atomic.Value
	}

	m := &MapAnyAlt{
		data: make(map[string]interface{}),
	}
	m.cut.Store(true)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		m.cut.Store(false)
	}
}

// 方案8：使用指针间接访问
func BenchmarkDisableCut_IndirectPointer(b *testing.B) {
	cut := new(uint32)
	*cut = 1

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		atomic.StoreUint32(cut, 0)
	}
}

// 方案9：批量操作（测试批量调用的性能）
func BenchmarkDisableCut_Batch(b *testing.B) {
	maps := make([]*MapAny, 100)
	for i := range maps {
		maps[i] = NewMap(nil).EnableCut(".")
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for _, m := range maps {
			m.DisableCut()
		}
	}
}

// 方案10：并发写（测试并发场景下的性能）
func BenchmarkDisableCut_ConcurrentWrites(b *testing.B) {
	m := NewMap(nil).EnableCut(".")
	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			m.DisableCut()
		}
	})
}

// 方案11：使用内存屏障优化（assembly 代码模拟）
func BenchmarkDisableCut_MemoryBarrier(b *testing.B) {
	m := NewMap(nil).EnableCut(".")
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		// 先赋值
		m.cut = 0
		// 内存屏障（确保其他 goroutine 可见）
		atomic.LoadUint32(&m.cut)
	}
}

// 方案12：无操作（作为性能基线对比）
func BenchmarkDisableCut_NoOp(b *testing.B) {
	m := NewMap(nil).EnableCut(".")
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		// 不做任何操作，仅用于基准对比
		_ = m
	}
}

// 方案13：链式调用优化（测试返回值的开销）
func BenchmarkDisableCut_Chaining(b *testing.B) {
	m := NewMap(nil).EnableCut(".")
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = m.DisableCut()
	}
}

// 方案14：条件写入（仅在值不同时写入）
func BenchmarkDisableCut_ConditionalWrite(b *testing.B) {
	m := NewMap(nil).EnableCut(".")
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		if atomic.LoadUint32(&m.cut) != 0 {
			atomic.StoreUint32(&m.cut, 0)
		}
	}
}

// 方案15：使用 CAS（Compare-And-Swap）
func BenchmarkDisableCut_CAS(b *testing.B) {
	m := NewMap(nil).EnableCut(".")
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for {
			if atomic.CompareAndSwapUint32(&m.cut, 1, 0) {
				break
			}
		}
	}
}

// 方案16：位操作优化（使用位掩码）
func BenchmarkDisableCut_BitMask(b *testing.B) {
	type MapAnyAlt struct {
		data  map[string]interface{}
		mu    sync.RWMutex
		flags uint32 // bit 0 = cut enabled
		seq   atomic.Value
	}

	m := &MapAnyAlt{
		data:  make(map[string]interface{}),
		flags: 0x01, // cut enabled
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		// 清除 bit 0
		atomic.StoreUint32(&m.flags, m.flags&0xFFFFFFFE)
	}
}

// 方案17：使用 unsafe 指针（高风险，仅供研究）
func BenchmarkDisableCut_Unsafe(b *testing.B) {
	m := NewMap(nil).EnableCut(".")
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		// 直接写入，绕过原子操作
		*(*uint32)(&m.cut) = 0
	}
}

// 方案18：缓存的原子操作指针（减少地址计算）
func BenchmarkDisableCut_CachedPointer(b *testing.B) {
	m := NewMap(nil).EnableCut(".")
	cutPtr := &m.cut
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		atomic.StoreUint32(cutPtr, 0)
	}
}

// 方案19：内联优化测试（小函数更容易内联）
func BenchmarkDisableCut_InlineFriendly(b *testing.B) {
	m := NewMap(nil).EnableCut(".")
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		// 简单赋值，编译器可能内联
		atomic.StoreUint32(&m.cut, 0)
	}
}

// 方案20：本地变量先操作再写回
func BenchmarkDisableCut_LocalVariable(b *testing.B) {
	m := NewMap(nil).EnableCut(".")
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		cut := uint32(0)
		atomic.StoreUint32(&m.cut, cut)
	}
}

// 方案21：使用 sync/atomic.Pointer（Go 1.18+）
func BenchmarkDisableCut_AtomicPointer(b *testing.B) {
	type MapAnyAlt struct {
		data map[string]interface{}
		mu   sync.RWMutex
		cut  atomic.Pointer[uint32]
		seq  atomic.Value
	}

	cutVal := uint32(1)
	m := &MapAnyAlt{
		data: make(map[string]interface{}),
	}
	m.cut.Store(&cutVal)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		newVal := uint32(0)
		m.cut.Store(&newVal)
	}
}

// 方案22：使用 channel 通信（完全不同的思路）
func BenchmarkDisableCut_Channel(b *testing.B) {
	type MapAnyAlt struct {
		data  map[string]interface{}
		mu    sync.RWMutex
		cutCh chan uint32
		seq   atomic.Value
	}

	m := &MapAnyAlt{
		data:  make(map[string]interface{}),
		cutCh: make(chan uint32, 10),
	}
	go func() {
		for v := range m.cutCh {
			_ = v
		}
	}()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		m.cutCh <- 0
	}
	close(m.cutCh)
}

// 方案23：预计算值（存储常见值）
func BenchmarkDisableCut_Precomputed(b *testing.B) {
	m := NewMap(nil).EnableCut(".")
	zero := uint32(0)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		atomic.StoreUint32(&m.cut, zero)
	}
}

// 方案24：混合策略：先检查再写入
func BenchmarkDisableCut_Hybrid(b *testing.B) {
	m := NewMap(nil).EnableCut(".")
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		// 快速路径：如果已经是 0，跳过
		if atomic.LoadUint32(&m.cut) == 0 {
			continue
		}
		atomic.StoreUint32(&m.cut, 0)
	}
}
