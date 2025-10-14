package candy

import (
	"reflect"
	"unsafe"
)

// Comparable 定义可比较的类型约束
type Comparable interface {
	comparable
}

// Copyable 定义可复制的基本类型约束
type Copyable interface {
	~bool | ~string |
		~int | ~int8 | ~int16 | ~int32 | ~int64 |
		~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 |
		~float32 | ~float64 |
		~complex64 | ~complex128
}

// TypedSliceCopy 类型安全的切片复制
func TypedSliceCopy[T any](src []T) []T {
	if src == nil {
		return nil
	}

	dst := make([]T, len(src))

	// 检查是否是基本类型
	if len(src) > 0 {
		switch any(src[0]).(type) {
		case bool, string,
			int, int8, int16, int32, int64,
			uint, uint8, uint16, uint32, uint64,
			float32, float64,
			complex64, complex128:
			// 对于基本类型，使用快速复制
			copy(dst, src)
			return dst
		}
	}

	// 对于复杂类型，逐个深度复制
	for i := range src {
		DeepCopy(src[i], &dst[i])
	}
	return dst
}

// TypedMapCopy 类型安全的 map 复制
func TypedMapCopy[K comparable, V any](src map[K]V) map[K]V {
	if src == nil {
		return nil
	}

	dst := make(map[K]V, len(src))

	// 检查值类型是否是基本类型
	var sampleValue V
	switch any(sampleValue).(type) {
	case bool, string,
		int, int8, int16, int32, int64,
		uint, uint8, uint16, uint32, uint64,
		float32, float64,
		complex64, complex128:
		// 对于基本类型，直接复制
		for k, v := range src {
			dst[k] = v
		}
		return dst
	}

	// 对于复杂类型，深度复制值
	for k, v := range src {
		var dstV V
		DeepCopy(v, &dstV)
		dst[k] = dstV
	}
	return dst
}

// GenericSliceEqual 类型安全的切片比较
func GenericSliceEqual[T comparable](a, b []T) bool {
	if len(a) != len(b) {
		return false
	}

	// 检查是否指向相同内存
	if len(a) > 0 && len(b) > 0 {
		aPtr := (*reflect.SliceHeader)(unsafe.Pointer(&a)).Data
		bPtr := (*reflect.SliceHeader)(unsafe.Pointer(&b)).Data
		if aPtr == bPtr {
			return true
		}
	}

	// 逐个比较元素
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

// MapEqual 类型安全的 map 比较
func MapEqual[K, V comparable](a, b map[K]V) bool {
	if len(a) != len(b) {
		return false
	}

	for k, av := range a {
		if bv, ok := b[k]; !ok || av != bv {
			return false
		}
	}
	return true
}

// PointerEqual 安全的指针比较
func PointerEqual[T comparable](a, b *T) bool {
	if a == nil && b == nil {
		return true
	}
	if a == nil || b == nil {
		return false
	}
	return *a == *b
}

// StructEqual 结构体字段比较（需要手动实现）
// 这个函数展示了如何为特定结构体类型实现高性能比较
func StructEqual[T any](a, b T, comparer func(T, T) bool) bool {
	return comparer(a, b)
}

// Clone 通用克隆函数，根据类型选择最佳策略
func Clone[T any](src T) T {
	var dst T
	DeepCopy(src, &dst)
	return dst
}

// CloneSlice 切片克隆的便捷函数
func CloneSlice[T any](src []T) []T {
	return TypedSliceCopy(src)
}

// CloneMap map 克隆的便捷函数
func CloneMap[K comparable, V any](src map[K]V) map[K]V {
	return TypedMapCopy(src)
}

// Equal 通用相等性检查，根据类型选择最佳策略
func Equal[T comparable](a, b T) bool {
	return a == b
}

// EqualSlice 切片相等性检查的便捷函数
func EqualSlice[T comparable](a, b []T) bool {
	return SliceEqual(a, b)
}

// EqualMap map 相等性检查的便捷函数
func EqualMap[K, V comparable](a, b map[K]V) bool {
	return MapEqual(a, b)
}