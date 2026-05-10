package anyx

import (
	"testing"

	"github.com/lazygophers/utils/candy"
)

// GetUint64SliceLegacy 优化前的原始实现（用于性能对比）
func (p *MapAny) GetUint64SliceLegacy(key string) []uint64 {
	val, ok := p.get(key)
	if !ok {
		return nil
	}

	return candy.ToUint64Slice(val)
}

// Benchmark 对比：优化前 vs 优化后
func BenchmarkGetUint64SliceLegacy_Uint64(b *testing.B) {
	mapAny := NewMap(map[string]interface{}{
		"uint64_slice": []uint64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
	})
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = mapAny.GetUint64SliceLegacy("uint64_slice")
	}
}

func BenchmarkGetUint64SliceCurrent_Uint64(b *testing.B) {
	mapAny := NewMap(map[string]interface{}{
		"uint64_slice": []uint64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
	})
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = mapAny.GetUint64Slice("uint64_slice")
	}
}

// 对比不同输入类型的性能
func BenchmarkGetUint64SliceLegacy_Int64(b *testing.B) {
	mapAny := NewMap(map[string]interface{}{
		"int64_slice": []int64{1, 2, 3, 4, 5},
	})
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = mapAny.GetUint64SliceLegacy("int64_slice")
	}
}

func BenchmarkGetUint64SliceCurrent_Int64(b *testing.B) {
	mapAny := NewMap(map[string]interface{}{
		"int64_slice": []int64{1, 2, 3, 4, 5},
	})
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = mapAny.GetUint64Slice("int64_slice")
	}
}

func BenchmarkGetUint64SliceLegacy_Int(b *testing.B) {
	mapAny := NewMap(map[string]interface{}{
		"int_slice": []int{1, 2, 3, 4, 5},
	})
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = mapAny.GetUint64SliceLegacy("int_slice")
	}
}

func BenchmarkGetUint64SliceCurrent_Int(b *testing.B) {
	mapAny := NewMap(map[string]interface{}{
		"int_slice": []int{1, 2, 3, 4, 5},
	})
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = mapAny.GetUint64Slice("int_slice")
	}
}

func BenchmarkGetUint64SliceLegacy_Large(b *testing.B) {
	mapAny := NewMap(map[string]interface{}{
		"large_slice": make([]uint64, 1000),
	})
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = mapAny.GetUint64SliceLegacy("large_slice")
	}
}

func BenchmarkGetUint64SliceCurrent_Large(b *testing.B) {
	mapAny := NewMap(map[string]interface{}{
		"large_slice": make([]uint64, 1000),
	})
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = mapAny.GetUint64Slice("large_slice")
	}
}
