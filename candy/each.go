package candy

import "reflect"

// Each 遍历切片、数组或映射的每个元素，并对每个元素执行指定的函数
// 此函数使用反射实现，保持与原始 API 的完全兼容性
func Each(collection interface{}, fn func(index int, value interface{})) {
	val := reflect.ValueOf(collection)

	switch val.Kind() {
	case reflect.Slice, reflect.Array:
		for i := 0; i < val.Len(); i++ {
			fn(i, val.Index(i).Interface())
		}
	case reflect.Map:
		keys := val.MapKeys()
		for i, key := range keys {
			fn(i, val.MapIndex(key).Interface())
		}
	default:
		panic("Each: collection must be a slice, array, or map")
	}
}

// EachSlice 遍历切片的每个元素，只传递值（高性能泛型版本）
func EachSlice[T any](values []T, fn func(value T)) {
	for _, value := range values {
		fn(value)
	}
}

// EachSliceIndexed 遍历切片的每个元素，传递索引和值（类型安全版本）
func EachSliceIndexed[T any](values []T, fn func(index int, value T)) {
	for i, value := range values {
		fn(i, value)
	}
}

// EachMap 遍历映射的每个键值对（类型安全版本）
func EachMap[K comparable, V any](m map[K]V, fn func(key K, value V)) {
	for k, v := range m {
		fn(k, v)
	}
}

// EachMapIndexed 遍历映射的每个键值对，带索引（类型安全版本）
func EachMapIndexed[K comparable, V any](m map[K]V, fn func(index int, key K, value V)) {
	i := 0
	for k, v := range m {
		fn(i, k, v)
		i++
	}
}
