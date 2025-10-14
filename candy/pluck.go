package candy

import (
	"fmt"
	"reflect"
)

// Pluck 从结构体切片中提取指定字段的值（泛型版本）
// 使用函数选择器而不是反射，提供类型安全和高性能
func Pluck[T any, U any](slice []T, selector func(T) U) []U {
	if len(slice) == 0 {
		return nil
	}

	result := make([]U, len(slice))
	for i, item := range slice {
		result[i] = selector(item)
	}
	return result
}

// PluckPtr 安全地从指针结构体切片中提取字段值
// 自动处理 nil 指针，提供默认值
func PluckPtr[T any, U any](slice []*T, selector func(*T) U, defaultVal U) []U {
	if len(slice) == 0 {
		return nil
	}

	result := make([]U, len(slice))
	for i, item := range slice {
		if item != nil {
			result[i] = selector(item)
		} else {
			result[i] = defaultVal
		}
	}
	return result
}

// PluckFilter 从结构体切片中提取字段值，同时进行过滤
func PluckFilter[T any, U any](slice []T, selector func(T) U, filter func(T) bool) []U {
	if len(slice) == 0 {
		return nil
	}

	var result []U
	for _, item := range slice {
		if filter(item) {
			result = append(result, selector(item))
		}
	}
	return result
}

// PluckUnique 从结构体切片中提取字段值并去重
func PluckUnique[T any, U comparable](slice []T, selector func(T) U) []U {
	if len(slice) == 0 {
		return nil
	}

	seen := make(map[U]struct{}, len(slice))
	var result []U

	for _, item := range slice {
		value := selector(item)
		if _, exists := seen[value]; !exists {
			seen[value] = struct{}{}
			result = append(result, value)
		}
	}
	return result
}

// PluckMap 从结构体切片中提取键值对，构建 map
func PluckMap[T any, K comparable, V any](slice []T, keySelector func(T) K, valueSelector func(T) V) map[K]V {
	if len(slice) == 0 {
		return nil
	}

	result := make(map[K]V, len(slice))
	for _, item := range slice {
		key := keySelector(item)
		value := valueSelector(item)
		result[key] = value
	}
	return result
}

// PluckGroupBy 按指定字段对结构体切片进行分组
func PluckGroupBy[T any, K comparable](slice []T, keySelector func(T) K) map[K][]T {
	if len(slice) == 0 {
		return nil
	}

	result := make(map[K][]T)
	for _, item := range slice {
		key := keySelector(item)
		result[key] = append(result[key], item)
	}
	return result
}

// 为向后兼容性保留的函数，使用新的泛型实现

// PluckIntGeneric 从结构体切片中提取 int 字段（泛型版本）
func PluckIntGeneric[T any](slice []T, selector func(T) int) []int {
	return Pluck(slice, selector)
}

// PluckStringGeneric 从结构体切片中提取 string 字段（泛型版本）
func PluckStringGeneric[T any](slice []T, selector func(T) string) []string {
	return Pluck(slice, selector)
}

// PluckInt32Generic 从结构体切片中提取 int32 字段（泛型版本）
func PluckInt32Generic[T any](slice []T, selector func(T) int32) []int32 {
	return Pluck(slice, selector)
}

// PluckInt64Generic 从结构体切片中提取 int64 字段（泛型版本）
func PluckInt64Generic[T any](slice []T, selector func(T) int64) []int64 {
	return Pluck(slice, selector)
}

// PluckUint32Generic 从结构体切片中提取 uint32 字段（泛型版本）
func PluckUint32Generic[T any](slice []T, selector func(T) uint32) []uint32 {
	return Pluck(slice, selector)
}

// PluckUint64Generic 从结构体切片中提取 uint64 字段（泛型版本）
func PluckUint64Generic[T any](slice []T, selector func(T) uint64) []uint64 {
	return Pluck(slice, selector)
}

// 基于反射的旧版 Pluck 实现，用于向后兼容

func pluck(list interface{}, fieldName string, deferVal interface{}) interface{} {
	v := reflect.ValueOf(list)
	switch v.Kind() {
	case reflect.Array, reflect.Slice:
		if v.Len() == 0 {
			return deferVal
		}

		ev := v.Type().Elem()
		evs := ev
		for evs.Kind() == reflect.Ptr {
			evs = evs.Elem()
		}

		switch evs.Kind() {
		case reflect.Struct:
			field, ok := evs.FieldByName(fieldName)
			if !ok {
				panic(fmt.Sprintf("field %s not found", fieldName))
			}

			result := reflect.MakeSlice(reflect.SliceOf(field.Type), v.Len(), v.Len())

			for i := 0; i < v.Len(); i++ {
				ev := v.Index(i)
				for ev.Kind() == reflect.Ptr {
					ev = ev.Elem()
				}
				if ev.Kind() != reflect.Struct {
					panic("element is not a struct")
				}
				if !ev.IsValid() {
					continue
				}
				result.Index(i).Set(ev.FieldByIndex(field.Index))
			}

			return result.Interface()
		case reflect.Slice, reflect.Array:
			var ev reflect.Value
			var c int
			for i := 0; i < v.Len(); i++ {
				ev = v.Index(i)
				for i := 0; i < ev.Len(); i++ {
					c += ev.Index(i).Len()
				}
			}

			result := reflect.MakeSlice(ev.Type(), c, c)
			var idx int
			for i := 0; i < v.Len(); i++ {
				ev := v.Index(i)
				for i := 0; i < ev.Len(); i++ {
					result.Index(idx).Set(ev.Index(i))
					idx++
				}
			}

			return result.Interface()
		default:
			panic("list element type is not supported")
		}

	default:
		panic("list must be an array or slice")
	}
}

// PluckInt 从结构体切片中提取指定字段的 int 值
func PluckInt(list interface{}, fieldName string) []int {
	return pluck(list, fieldName, []int{}).([]int)
}

// PluckInt32 从结构体切片中提取指定字段的 int32 值
func PluckInt32(list interface{}, fieldName string) []int32 {
	return pluck(list, fieldName, []int32{}).([]int32)
}

// PluckInt64 从结构体切片中提取指定字段的 int64 值
func PluckInt64(list interface{}, fieldName string) []int64 {
	return pluck(list, fieldName, []int64{}).([]int64)
}

// PluckString 从结构体切片中提取指定字段的 string 值
func PluckString(list interface{}, fieldName string) []string {
	return pluck(list, fieldName, []string{}).([]string)
}

// PluckUint32 从结构体切片中提取指定字段的 uint32 值
func PluckUint32(list interface{}, fileName string) []uint32 {
	return pluck(list, fileName, []uint32{}).([]uint32)
}

// PluckUint64 从结构体切片中提取指定字段的 uint64 值
func PluckUint64(list interface{}, fieldName string) []uint64 {
	return pluck(list, fieldName, []uint64{}).([]uint64)
}

// PluckStringSlice 从结构体切片中提取指定字段的 []string 值
func PluckStringSlice(list interface{}, fieldName string) [][]string {
	return pluck(list, fieldName, [][]string{}).([][]string)
}
