package candy

import (
	"reflect"
	"unsafe"
)

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
