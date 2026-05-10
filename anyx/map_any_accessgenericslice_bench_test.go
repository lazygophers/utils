package anyx

import (
	"fmt"
	"reflect"
	"testing"
)

// 方案1：当前实现 - 直接返回错误
func accessGenericSlice_Current(slice any, index int) (any, error) {
	return nil, fmt.Errorf("%w: cannot access index %d on type %T", ErrInvalidSlice, index, slice)
}

// 方案2：使用 reflect.Value 访问
func accessGenericSlice_ReflectValue(slice any, index int) (any, error) {
	v := reflect.ValueOf(slice)
	if v.Kind() != reflect.Slice {
		return nil, fmt.Errorf("%w: cannot access index %d on type %T", ErrInvalidSlice, index, slice)
	}
	if index < 0 || index >= v.Len() {
		return nil, ErrOutOfRange
	}
	return v.Index(index).Interface(), nil
}

// 方案3：使用 reflect.Value + 预检查
func accessGenericSlice_ReflectValuePreCheck(slice any, index int) (any, error) {
	if slice == nil {
		return nil, fmt.Errorf("%w: cannot access index %d on type %T", ErrInvalidSlice, index, slice)
	}
	v := reflect.ValueOf(slice)
	if v.Kind() != reflect.Slice {
		return nil, fmt.Errorf("%w: cannot access index %d on type %T", ErrInvalidSlice, index, slice)
	}
	if index < 0 || index >= v.Len() {
		return nil, ErrOutOfRange
	}
	return v.Index(index).Interface(), nil
}

// 方案4：使用 reflect.Value + Unsafe（仅测试用，不推荐生产）
func accessGenericSlice_ReflectUnsafe(slice any, index int) (any, error) {
	v := reflect.ValueOf(slice)
	if v.Kind() != reflect.Slice {
		return nil, fmt.Errorf("%w: cannot access index %d on type %T", ErrInvalidSlice, index, slice)
	}
	if index < 0 || index >= v.Len() {
		return nil, ErrOutOfRange
	}
	elem := v.Index(index)
	return elem.Interface(), nil
}

// 方案5：使用 reflect.Value + 缓存 Kind
func accessGenericSlice_ReflectCachedKind(slice any, index int) (any, error) {
	v := reflect.ValueOf(slice)
	kind := v.Kind()
	if kind != reflect.Slice {
		return nil, fmt.Errorf("%w: cannot access index %d on type %T", ErrInvalidSlice, index, slice)
	}
	if index < 0 || index >= v.Len() {
		return nil, ErrOutOfRange
	}
	return v.Index(index).Interface(), nil
}

// 方案6：先类型断言再 reflect（针对常见类型优化）
func accessGenericSlice_TypeAssertFirst(slice any, index int) (any, error) {
	switch v := slice.(type) {
	case []uint:
		if uint(index) >= uint(len(v)) {
			return nil, ErrOutOfRange
		}
		return v[index], nil
	case []float32:
		if uint(index) >= uint(len(v)) {
			return nil, ErrOutOfRange
		}
		return v[index], nil
	case []int32:
		if uint(index) >= uint(len(v)) {
			return nil, ErrOutOfRange
		}
		return v[index], nil
	default:
		// Fallback to reflection
		vReflect := reflect.ValueOf(slice)
		if vReflect.Kind() != reflect.Slice {
			return nil, fmt.Errorf("%w: cannot access index %d on type %T", ErrInvalidSlice, index, slice)
		}
		if index < 0 || index >= vReflect.Len() {
			return nil, ErrOutOfRange
		}
		return vReflect.Index(index).Interface(), nil
	}
}

// 方案7：完全 reflect，无错误检查（仅用于性能对比）
func accessGenericSlice_ReflectNoCheck(slice any, index int) (any, error) {
	v := reflect.ValueOf(slice)
	// 注意：这不检查 Kind 和边界，仅用于性能测试
	return v.Index(index).Interface(), nil
}

// 方案8：使用 reflect.Value + 切片长度缓存
func accessGenericSlice_ReflectCachedLen(slice any, index int) (any, error) {
	v := reflect.ValueOf(slice)
	if v.Kind() != reflect.Slice {
		return nil, fmt.Errorf("%w: cannot access index %d on type %T", ErrInvalidSlice, index, slice)
	}
	length := v.Len()
	if index < 0 || index >= length {
		return nil, ErrOutOfRange
	}
	return v.Index(index).Interface(), nil
}

// 方案9：使用 reflect + 边界检查优化（uint 转换）
func accessGenericSlice_ReflectUintCheck(slice any, index int) (any, error) {
	v := reflect.ValueOf(slice)
	if v.Kind() != reflect.Slice {
		return nil, fmt.Errorf("%w: cannot access index %d on type %T", ErrInvalidSlice, index, slice)
	}
	if uint(index) >= uint(v.Len()) {
		return nil, ErrOutOfRange
	}
	return v.Index(index).Interface(), nil
}

// 方案10：简化错误消息（减少格式化开销）
func accessGenericSlice_SimpleError(slice any, index int) (any, error) {
	v := reflect.ValueOf(slice)
	if v.Kind() != reflect.Slice {
		return nil, ErrInvalidSlice
	}
	if uint(index) >= uint(v.Len()) {
		return nil, ErrOutOfRange
	}
	return v.Index(index).Interface(), nil
}

// ============================================================================
// Benchmarks
// ============================================================================

// Benchmark 1: 当前实现（错误返回）
func BenchmarkAccessGenericSlice_Current_ErrorCase(b *testing.B) {
	slice := []uint{1, 2, 3}
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = accessGenericSlice_Current(slice, 0)
	}
}

// Benchmark 2: Reflect.Value 基础实现
func BenchmarkAccessGenericSlice_ReflectValue_Valid(b *testing.B) {
	slice := []uint{1, 2, 3}
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = accessGenericSlice_ReflectValue(slice, 0)
	}
}

// Benchmark 3: Reflect + 预检查 nil
func BenchmarkAccessGenericSlice_ReflectValuePreCheck_Valid(b *testing.B) {
	slice := []uint{1, 2, 3}
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = accessGenericSlice_ReflectValuePreCheck(slice, 0)
	}
}

// Benchmark 4: 类型断言优先（命中 fast path）
func BenchmarkAccessGenericSlice_TypeAssertFirst_Hit(b *testing.B) {
	slice := []uint{1, 2, 3}
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = accessGenericSlice_TypeAssertFirst(slice, 0)
	}
}

// Benchmark 5: 类型断言优先（未命中，fallback to reflect）
func BenchmarkAccessGenericSlice_TypeAssertFirst_Miss(b *testing.B) {
	type customSlice []int
	slice := customSlice{1, 2, 3}
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = accessGenericSlice_TypeAssertFirst(slice, 0)
	}
}

// Benchmark 6: Reflect + 缓存 Kind
func BenchmarkAccessGenericSlice_ReflectCachedKind_Valid(b *testing.B) {
	slice := []uint{1, 2, 3}
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = accessGenericSlice_ReflectCachedKind(slice, 0)
	}
}

// Benchmark 7: Reflect + 缓存长度
func BenchmarkAccessGenericSlice_ReflectCachedLen_Valid(b *testing.B) {
	slice := []uint{1, 2, 3}
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = accessGenericSlice_ReflectCachedLen(slice, 0)
	}
}

// Benchmark 8: Reflect + Uint 边界检查
func BenchmarkAccessGenericSlice_ReflectUintCheck_Valid(b *testing.B) {
	slice := []uint{1, 2, 3}
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = accessGenericSlice_ReflectUintCheck(slice, 0)
	}
}

// Benchmark 9: 简化错误消息
func BenchmarkAccessGenericSlice_SimpleError_Valid(b *testing.B) {
	slice := []uint{1, 2, 3}
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = accessGenericSlice_SimpleError(slice, 0)
	}
}

// Benchmark 10: Reflect.Value（负索引边界情况）
func BenchmarkAccessGenericSlice_ReflectValue_NegativeIndex(b *testing.B) {
	slice := []uint{1, 2, 3}
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = accessGenericSlice_ReflectValue(slice, -1)
	}
}

// Benchmark 11: 不同切片类型性能对比
func BenchmarkAccessGenericSlice_Types_Uint(b *testing.B) {
	slice := []uint{1, 2, 3}
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_, _ = accessGenericSlice_ReflectValue(slice, 1)
	}
}

func BenchmarkAccessGenericSlice_Types_Float32(b *testing.B) {
	slice := []float32{1.1, 2.2, 3.3}
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_, _ = accessGenericSlice_ReflectValue(slice, 1)
	}
}

func BenchmarkAccessGenericSlice_Types_Int32(b *testing.B) {
	slice := []int32{1, 2, 3}
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_, _ = accessGenericSlice_ReflectValue(slice, 1)
	}
}

func BenchmarkAccessGenericSlice_Types_CustomStruct(b *testing.B) {
	type Point struct{ X, Y int }
	slice := []Point{{1, 2}, {3, 4}, {5, 6}}
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_, _ = accessGenericSlice_ReflectValue(slice, 1)
	}
}

// Benchmark 12: 大切片访问性能
func BenchmarkAccessGenericSlice_LargeSlice_Middle(b *testing.B) {
	slice := make([]uint, 1000)
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = accessGenericSlice_ReflectValue(slice, 500)
	}
}

func BenchmarkAccessGenericSlice_LargeSlice_Last(b *testing.B) {
	slice := make([]uint, 1000)
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = accessGenericSlice_ReflectValue(slice, 999)
	}
}

// Benchmark 13: 错误路径性能对比
func BenchmarkAccessGenericSlice_Error_NotSlice(b *testing.B) {
	notSlice := "not a slice"
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = accessGenericSlice_ReflectValue(notSlice, 0)
	}
}

func BenchmarkAccessGenericSlice_Error_OutOfRange(b *testing.B) {
	slice := []uint{1, 2, 3}
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = accessGenericSlice_ReflectValue(slice, 10)
	}
}

// Benchmark 14: 当前实现 vs Reflect 实现错误路径
func BenchmarkAccessGenericSlice_Current_NotSlice(b *testing.B) {
	notSlice := "not a slice"
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = accessGenericSlice_Current(notSlice, 0)
	}
}

// Benchmark 15: 并发访问性能
func BenchmarkAccessGenericSlice_Concurrent_Parallel(b *testing.B) {
	slice := []uint{1, 2, 3, 4, 5}
	b.ReportAllocs()
	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		i := 0
		for pb.Next() {
			_, _ = accessGenericSlice_ReflectValue(slice, i%5)
			i++
		}
	})
}
