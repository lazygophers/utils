package candy

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