package candy

import (
	"fmt"
	"reflect"
	"sync"
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

// PluckPtr 安全地从指针结构体切片中提取字段值（优化版本）
// 自动处理 nil 指针，提供默认值
func PluckPtr[T any, U any](slice []*T, selector func(*T) U, defaultVal U) []U {
	if len(slice) == 0 {
		return nil
	}

	result := make([]U, len(slice))
	length := len(slice)

	for i := 0; i < length; i++ {
		item := slice[i]
		if item != nil {
			result[i] = selector(item)
		} else {
			result[i] = defaultVal
		}
	}
	return result
}

// PluckUnique 从结构体切片中提取字段值并去重（优化版本）
func PluckUnique[T any, U comparable](slice []T, selector func(T) U) []U {
	if len(slice) == 0 {
		return nil
	}

	seen := make(map[U]struct{}, len(slice))
	result := make([]U, 0, len(slice))

	for _, item := range slice {
		value := selector(item)
		if _, exists := seen[value]; !exists {
			seen[value] = struct{}{}
			result = append(result, value)
		}
	}
	return result
}

// PluckMap 从结构体切片中提取键值对，构建 map（优化版本）
func PluckMap[T any, K comparable, V any](slice []T, keySelector func(T) K, valueSelector func(T) V) map[K]V {
	if len(slice) == 0 {
		return nil
	}

	result := make(map[K]V, len(slice))
	length := len(slice)

	for i := 0; i < length; i++ {
		item := slice[i]
		key := keySelector(item)
		value := valueSelector(item)
		result[key] = value
	}
	return result
}

// PluckGroupBy 按指定字段对结构体切片进行分组（优化版本）
func PluckGroupBy[T any, K comparable](slice []T, keySelector func(T) K) map[K][]T {
	if len(slice) == 0 {
		return nil
	}

	// 预估分组数量
	estimatedGroups := len(slice) / 10
	if estimatedGroups < 4 {
		estimatedGroups = 4
	}

	result := make(map[K][]T, estimatedGroups)

	for _, item := range slice {
		key := keySelector(item)
		result[key] = append(result[key], item)
	}
	return result
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

// PluckInt 从结构体切片中提取指定字段的 int 值（优化版本）
func PluckInt(list interface{}, fieldName string) []int {
	return pluckIntOptimized(list, fieldName)
}

// PluckInt32 从结构体切片中提取指定字段的 int32 值（优化版本）
func PluckInt32(list interface{}, fieldName string) []int32 {
	return pluckInt32Optimized(list, fieldName)
}

// PluckInt64 从结构体切片中提取指定字段的 int64 值（优化版本）
func PluckInt64(list interface{}, fieldName string) []int64 {
	return pluckInt64Optimized(list, fieldName)
}

// PluckString 从结构体切片中提取指定字段的 string 值（优化版本）
func PluckString(list interface{}, fieldName string) []string {
	return pluckStringOptimized(list, fieldName)
}

// pluckStringOptimized 优化版的 PluckString 实现
// 使用缓存反射避免重复的字段查找，提升性能
func pluckStringOptimized(list interface{}, fieldName string) []string {
	v := reflect.ValueOf(list)
	if v.Kind() != reflect.Slice && v.Kind() != reflect.Array {
		panic("list must be an array or slice")
	}

	if v.Len() == 0 {
		return []string{}
	}

	// 获取元素类型
	elemType := v.Type().Elem()

	// 尝试获取缓存的字段索引
	fieldIndex, fieldValueType, ok := getPluckStringFieldIndex(elemType, fieldName)
	if !ok {
		panic(fmt.Sprintf("field %s not found", fieldName))
	}

	// 验证字段类型是否为 string
	if fieldValueType.Kind() != reflect.String {
		panic(fmt.Sprintf("field %s is not of type string", fieldName))
	}

	// 预分配结果切片
	result := make([]string, v.Len())

	// 遍历切片并提取字段值
	for i := 0; i < v.Len(); i++ {
		elem := v.Index(i)

		// 处理指针类型
		for elem.Kind() == reflect.Ptr {
			elem = elem.Elem()
		}

		if elem.Kind() != reflect.Struct {
			panic("element is not a struct")
		}

		// 使用缓存的字段索引直接访问字段
		result[i] = elem.FieldByIndex(fieldIndex).String()
	}

	return result
}

// pluckStringFieldCache 字段索引缓存
var pluckStringFieldCache struct {
	sync.RWMutex
	cache map[reflect.Type]map[string][]int
}

// 初始化缓存
func init() {
	pluckStringFieldCache.cache = make(map[reflect.Type]map[string][]int)
}

// getPluckStringFieldIndex 获取字段的索引和类型，使用缓存避免重复反射
func getPluckStringFieldIndex(elemType reflect.Type, fieldName string) ([]int, reflect.Type, bool) {
	// 尝试从缓存中读取
	pluckStringFieldCache.RLock()
	if fields, ok := pluckStringFieldCache.cache[elemType]; ok {
		if index, exists := fields[fieldName]; exists {
			pluckStringFieldCache.RUnlock()
			// 从缓存获取后，还需要获取字段类型进行验证
			actualType := elemType
			for actualType.Kind() == reflect.Ptr {
				actualType = actualType.Elem()
			}
			if field, found := actualType.FieldByName(fieldName); found {
				return index, field.Type, true
			}
		}
	}
	pluckStringFieldCache.RUnlock()

	// 缓存未命中，获取字段索引并缓存
	pluckStringFieldCache.Lock()
	defer pluckStringFieldCache.Unlock()

	// 双重检查
	if fields, ok := pluckStringFieldCache.cache[elemType]; ok {
		if index, exists := fields[fieldName]; exists {
			actualType := elemType
			for actualType.Kind() == reflect.Ptr {
				actualType = actualType.Elem()
			}
			if field, found := actualType.FieldByName(fieldName); found {
				return index, field.Type, true
			}
		}
	}

	// 解析指针类型
	actualType := elemType
	for actualType.Kind() == reflect.Ptr {
		actualType = actualType.Elem()
	}

	if actualType.Kind() != reflect.Struct {
		return nil, nil, false
	}

	// 查找字段
	field, found := actualType.FieldByName(fieldName)
	if !found {
		return nil, nil, false
	}

	// 初始化类型缓存
	if pluckStringFieldCache.cache[elemType] == nil {
		pluckStringFieldCache.cache[elemType] = make(map[string][]int)
	}

	// 缓存字段索引
	pluckStringFieldCache.cache[elemType][fieldName] = field.Index
	return field.Index, field.Type, true
}

// PluckUint32 从结构体切片中提取指定字段的 uint32 值（优化版本）
func PluckUint32(list interface{}, fileName string) []uint32 {
	return pluckUint32Optimized(list, fileName)
}

// PluckUint64 从结构体切片中提取指定字段的 uint64 值（优化版本）
func PluckUint64(list interface{}, fieldName string) []uint64 {
	return pluckUint64Optimized(list, fieldName)
}

// PluckStringSlice 从结构体切片中提取指定字段的 []string 值（优化版本）
func PluckStringSlice(list interface{}, fieldName string) [][]string {
	return pluckStringSliceOptimized(list, fieldName)
}

// ==================== PluckInt 优化实现 ====================

// pluckIntFieldCache 字段索引缓存
var pluckIntFieldCache struct {
	sync.RWMutex
	cache map[reflect.Type]map[string]int
}

// 初始化缓存
func init() {
	pluckIntFieldCache.cache = make(map[reflect.Type]map[string]int)
}

// getPluckIntFieldIndex 获取字段的索引和类型，使用缓存避免重复反射
func getPluckIntFieldIndex(elemType reflect.Type, fieldName string) (int, reflect.Type, bool) {
	// 尝试从缓存中读取
	pluckIntFieldCache.RLock()
	if fields, ok := pluckIntFieldCache.cache[elemType]; ok {
		if index, exists := fields[fieldName]; exists {
			pluckIntFieldCache.RUnlock()
			// 从缓存获取后，还需要获取字段类型进行验证
			actualType := elemType
			for actualType.Kind() == reflect.Ptr {
				actualType = actualType.Elem()
			}
			if field, found := actualType.FieldByName(fieldName); found {
				return index, field.Type, true
			}
		}
	}
	pluckIntFieldCache.RUnlock()

	// 缓存未命中，获取字段索引并缓存
	pluckIntFieldCache.Lock()
	defer pluckIntFieldCache.Unlock()

	// 双重检查
	if fields, ok := pluckIntFieldCache.cache[elemType]; ok {
		if index, exists := fields[fieldName]; exists {
			actualType := elemType
			for actualType.Kind() == reflect.Ptr {
				actualType = actualType.Elem()
			}
			if field, found := actualType.FieldByName(fieldName); found {
				return index, field.Type, true
			}
		}
	}

	// 解析指针类型
	actualType := elemType
	for actualType.Kind() == reflect.Ptr {
		actualType = actualType.Elem()
	}

	if actualType.Kind() != reflect.Struct {
		return 0, nil, false
	}

	field, found := actualType.FieldByName(fieldName)
	if !found {
		return 0, nil, false
	}

	// 缓存字段索引（只缓存第一个索引，适用于非嵌套字段）
	if pluckIntFieldCache.cache[elemType] == nil {
		pluckIntFieldCache.cache[elemType] = make(map[string]int)
	}
	pluckIntFieldCache.cache[elemType][fieldName] = field.Index[0]

	return field.Index[0], field.Type, true
}

// pluckIntOptimized PluckInt 的优化实现，使用缓存反射提高性能
func pluckIntOptimized(list interface{}, fieldName string) []int {
	v := reflect.ValueOf(list)
	if v.Kind() != reflect.Slice && v.Kind() != reflect.Array {
		panic("list must be an array or slice")
	}

	if v.Len() == 0 {
		return []int{}
	}

	// 获取元素类型
	elemType := v.Type().Elem()

	// 尝试获取缓存的字段索引
	fieldIndex, fieldValueType, ok := getPluckIntFieldIndex(elemType, fieldName)
	if !ok {
		panic(fmt.Sprintf("field %s not found", fieldName))
	}

	// 验证字段类型是否为 int
	if fieldValueType.Kind() != reflect.Int {
		panic(fmt.Sprintf("field %s is not of type int", fieldName))
	}

	// 预分配结果切片
	result := make([]int, v.Len())

	// 遍历切片并提取字段值
	for i := 0; i < v.Len(); i++ {
		elem := v.Index(i)

		// 处理指针类型
		for elem.Kind() == reflect.Ptr {
			elem = elem.Elem()
		}

		if elem.Kind() != reflect.Struct {
			panic("element is not a struct")
		}

		// 使用缓存的字段索引直接访问字段
		result[i] = int(elem.Field(fieldIndex).Int())
	}

	return result
}
