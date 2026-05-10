package anyx

import (
	"fmt"
)

// navigateToValueOptimized navigateToValue 的优化版本
// 优化点：
// 1. 内联 parseIndex 逻辑，减少函数调用开销
// 2. 预先检查字符串长度，避免多次 len() 调用
// 3. 使用更高效的模式匹配
// TestNavigateToValueOptimized 导出用于测试
func TestNavigateToValueOptimized(current any, part string) (any, error) {
	return navigateToValueOptimized(current, part)
}

func navigateToValueOptimized(current any, part string) (any, error) {
	// 快速路径：检查是否是数组索引模式
	partLen := len(part)
	if partLen > 2 && part[0] == '[' && part[partLen-1] == ']' {
		// 内联索引解析和访问
		indexStr := part[1 : partLen-1]

		// 内联 parseIndex 逻辑
		indexStrLen := len(indexStr)
		if indexStrLen == 0 {
			return nil, fmt.Errorf("%w: empty index", ErrInvalidIndex)
		}

		// 处理负号
		negative := false
		startPos := 0
		if indexStr[0] == '-' {
			negative = true
			startPos = 1
			if indexStrLen == 1 {
				// 只有负号，返回错误（bug fix）
				return nil, fmt.Errorf("%w: '%s'", ErrInvalidIndex, indexStr)
			}
		}

		// 快速解析数字
		var index int
		for i := startPos; i < indexStrLen; i++ {
			c := indexStr[i]
			if c < '0' || c > '9' {
				return nil, fmt.Errorf("%w: '%s'", ErrInvalidIndex, indexStr)
			}
			index = index*10 + int(c-'0')
		}

		if negative {
			index = -index
		}

		// 内联数组访问
		return accessArrayIndexOptimized(current, index)
	}

	// Map 键访问
	return accessMapKeyOptimized(current, part)
}

// accessArrayIndexOptimized accessArrayIndex 的优化版本
// 优化点：
// 1. 直接使用解析好的索引，避免重复解析
// 2. 减少边界检查次数
// 3. 更紧凑的分支预测
func accessArrayIndexOptimized(current any, index int) (any, error) {
	// 负数索引转换
	if index < 0 {
		switch v := current.(type) {
		case []any:
			if -index > len(v) {
				return nil, fmt.Errorf("%w: index %d out of range [-%d, %d)", ErrOutOfRange, index, len(v), len(v))
			}
			index = len(v) + index
		case []string:
			if -index > len(v) {
				return nil, fmt.Errorf("%w: index %d out of range [-%d, %d)", ErrOutOfRange, index, len(v), len(v))
			}
			index = len(v) + index
		case []int:
			if -index > len(v) {
				return nil, fmt.Errorf("%w: index %d out of range [-%d, %d)", ErrOutOfRange, index, len(v), len(v))
			}
			index = len(v) + index
		case []int64:
			if -index > len(v) {
				return nil, fmt.Errorf("%w: index %d out of range [-%d, %d)", ErrOutOfRange, index, len(v), len(v))
			}
			index = len(v) + index
		case []float64:
			if -index > len(v) {
				return nil, fmt.Errorf("%w: index %d out of range [-%d, %d)", ErrOutOfRange, index, len(v), len(v))
			}
			index = len(v) + index
		case []bool:
			if -index > len(v) {
				return nil, fmt.Errorf("%w: index %d out of range [-%d, %d)", ErrOutOfRange, index, len(v), len(v))
			}
			index = len(v) + index
		case []map[string]any:
			if -index > len(v) {
				return nil, fmt.Errorf("%w: index %d out of range [-%d, %d)", ErrOutOfRange, index, len(v), len(v))
			}
			index = len(v) + index
		default:
			return accessGenericSlice(v, index)
		}
	}

	// 正数索引访问
	switch v := current.(type) {
	case []any:
		if index < 0 || index >= len(v) {
			return nil, fmt.Errorf("%w: index %d out of range [0, %d)", ErrOutOfRange, index, len(v))
		}
		return v[index], nil
	case []string:
		if index < 0 || index >= len(v) {
			return nil, fmt.Errorf("%w: index %d out of range [0, %d)", ErrOutOfRange, index, len(v))
		}
		return v[index], nil
	case []int:
		if index < 0 || index >= len(v) {
			return nil, fmt.Errorf("%w: index %d out of range [0, %d)", ErrOutOfRange, index, len(v))
		}
		return v[index], nil
	case []int64:
		if index < 0 || index >= len(v) {
			return nil, fmt.Errorf("%w: index %d out of range [0, %d)", ErrOutOfRange, index, len(v))
		}
		return v[index], nil
	case []float64:
		if index < 0 || index >= len(v) {
			return nil, fmt.Errorf("%w: index %d out of range [0, %d)", ErrOutOfRange, index, len(v))
		}
		return v[index], nil
	case []bool:
		if index < 0 || index >= len(v) {
			return nil, fmt.Errorf("%w: index %d out of range [0, %d)", ErrOutOfRange, index, len(v))
		}
		return v[index], nil
	case []map[string]any:
		if index < 0 || index >= len(v) {
			return nil, fmt.Errorf("%w: index %d out of range [0, %d)", ErrOutOfRange, index, len(v))
		}
		return v[index], nil
	default:
		return accessGenericSlice(v, index)
	}
}

// accessMapKeyOptimized accessMapKey 的优化版本
// 优化点：
// 1. 减少类型切换开销
// 2. 更紧凑的代码路径
func accessMapKeyOptimized(current any, key string) (any, error) {
	switch v := current.(type) {
	case map[string]any:
		val, ok := v[key]
		if !ok {
			return nil, ErrNotFound
		}
		return val, nil
	case map[any]any:
		val, ok := v[key]
		if !ok {
			return nil, ErrNotFound
		}
		return val, nil
	default:
		return nil, fmt.Errorf("%w: cannot access key '%s' on type %T", ErrInvalidMapType, key, current)
	}
}

// ============================================================
// 方案 2: 更激进的优化 - 完全内联所有逻辑
// ============================================================

// TestNavigateToValueOptimizedV2 导出用于测试
func TestNavigateToValueOptimizedV2(current any, part string) (any, error) {
	return navigateToValueOptimizedV2(current, part)
}

func navigateToValueOptimizedV2(current any, part string) (any, error) {
	partLen := len(part)

	// 检查是否是数组索引
	if partLen > 2 && part[0] == '[' && part[partLen-1] == ']' {
		indexStr := part[1 : partLen-1]
		indexStrLen := len(indexStr)

		// 解析索引
		if indexStrLen == 0 {
			return nil, fmt.Errorf("%w: empty index", ErrInvalidIndex)
		}

		negative := indexStr[0] == '-'
		startPos := 0
		if negative {
			startPos = 1
			if indexStrLen == 1 {
				return accessArrayIndexOptimized(current, 0)
			}
		}

		var index int
		for i := startPos; i < indexStrLen; i++ {
			c := indexStr[i]
			if c < '0' || c > '9' {
				return nil, fmt.Errorf("%w: '%s'", ErrInvalidIndex, indexStr)
			}
			index = index*10 + int(c-'0')
		}

		if negative {
			index = -index
		}

		return accessArrayIndexOptimized(current, index)
	}

	// Map 键访问
	switch v := current.(type) {
	case map[string]any:
		val, ok := v[part]
		if !ok {
			return nil, ErrNotFound
		}
		return val, nil
	case map[any]any:
		val, ok := v[part]
		if !ok {
			return nil, ErrNotFound
		}
		return val, nil
	default:
		return nil, fmt.Errorf("%w: cannot access key '%s' on type %T", ErrInvalidMapType, part, current)
	}
}

// ============================================================
// 方案 3: 专门针对常见场景的优化
// ============================================================

// TestNavigateToValueOptimizedV3 导出用于测试
func TestNavigateToValueOptimizedV3(current any, part string) (any, error) {
	return navigateToValueOptimizedV3(current, part)
}

func navigateToValueOptimizedV3(current any, part string) (any, error) {
	// 快速路径：短字符串键直接访问
	if len(part) > 0 && (part[0] != '[' || part[len(part)-1] != ']') {
		switch v := current.(type) {
		case map[string]any:
			if val, ok := v[part]; ok {
				return val, nil
			}
			return nil, ErrNotFound
		case map[any]any:
			if val, ok := v[part]; ok {
				return val, nil
			}
			return nil, ErrNotFound
		default:
			return nil, fmt.Errorf("%w: cannot access key '%s' on type %T", ErrInvalidMapType, part, current)
		}
	}

	// 数组索引路径（使用内联逻辑）
	partLen := len(part)
	if partLen > 2 && part[0] == '[' && part[partLen-1] == ']' {
		indexStr := part[1 : partLen-1]
		indexStrLen := len(indexStr)

		if indexStrLen == 0 {
			return nil, fmt.Errorf("%w: empty index", ErrInvalidIndex)
		}

		negative := indexStr[0] == '-'
		startPos := 0
		if negative {
			startPos = 1
			if indexStrLen == 1 {
				return accessArrayIndexOptimized(current, 0)
			}
		}

		var index int
		for i := startPos; i < indexStrLen; i++ {
			c := indexStr[i]
			if c < '0' || c > '9' {
				return nil, fmt.Errorf("%w: '%s'", ErrInvalidIndex, indexStr)
			}
			index = index*10 + int(c-'0')
		}

		if negative {
			index = -index
		}

		return accessArrayIndexOptimized(current, index)
	}

	// 默认路径
	return accessMapKeyOptimized(current, part)
}

// ============================================================
// 方案 4: 零分配优化（避免错误路径的字符串格式化）
// ============================================================

var (
	errInvalidIndexCached   = fmt.Errorf("%w: invalid index", ErrInvalidIndex)
	errOutOfRangeCached     = fmt.Errorf("%w: index out of range", ErrOutOfRange)
	errInvalidSliceCached   = fmt.Errorf("%w: invalid slice type", ErrInvalidSlice)
	errInvalidMapTypeCached = fmt.Errorf("%w: invalid map type", ErrInvalidMapType)
	errNotFoundCached       = ErrNotFound
)

// TestNavigateToValueOptimizedV4 导出用于测试
func TestNavigateToValueOptimizedV4(current any, part string) (any, error) {
	return navigateToValueOptimizedV4(current, part)
}

func navigateToValueOptimizedV4(current any, part string) (any, error) {
	partLen := len(part)

	if partLen > 2 && part[0] == '[' && part[partLen-1] == ']' {
		indexStr := part[1 : partLen-1]
		indexStrLen := len(indexStr)

		if indexStrLen == 0 {
			return nil, errInvalidIndexCached
		}

		negative := indexStr[0] == '-'
		startPos := 0
		if negative {
			startPos = 1
			if indexStrLen == 1 {
				return accessArrayIndexOptimizedV4(current, 0)
			}
		}

		var index int
		for i := startPos; i < indexStrLen; i++ {
			c := indexStr[i]
			if c < '0' || c > '9' {
				return nil, errInvalidIndexCached
			}
			index = index*10 + int(c-'0')
		}

		if negative {
			index = -index
		}

		return accessArrayIndexOptimizedV4(current, index)
	}

	switch v := current.(type) {
	case map[string]any:
		if val, ok := v[part]; ok {
			return val, nil
		}
		return nil, errNotFoundCached
	case map[any]any:
		if val, ok := v[part]; ok {
			return val, nil
		}
		return nil, errNotFoundCached
	default:
		return nil, errInvalidMapTypeCached
	}
}

func accessArrayIndexOptimizedV4(current any, index int) (any, error) {
	if index < 0 {
		switch v := current.(type) {
		case []any:
			if -index > len(v) {
				return nil, errOutOfRangeCached
			}
			index = len(v) + index
			return v[index], nil
		case []string:
			if -index > len(v) {
				return nil, errOutOfRangeCached
			}
			index = len(v) + index
			return v[index], nil
		case []int:
			if -index > len(v) {
				return nil, errOutOfRangeCached
			}
			index = len(v) + index
			return v[index], nil
		case []int64:
			if -index > len(v) {
				return nil, errOutOfRangeCached
			}
			index = len(v) + index
			return v[index], nil
		case []float64:
			if -index > len(v) {
				return nil, errOutOfRangeCached
			}
			index = len(v) + index
			return v[index], nil
		case []bool:
			if -index > len(v) {
				return nil, errOutOfRangeCached
			}
			index = len(v) + index
			return v[index], nil
		case []map[string]any:
			if -index > len(v) {
				return nil, errOutOfRangeCached
			}
			index = len(v) + index
			return v[index], nil
		default:
			return nil, errInvalidSliceCached
		}
	}

	switch v := current.(type) {
	case []any:
		if uint(index) < uint(len(v)) {
			return v[index], nil
		}
		return nil, errOutOfRangeCached
	case []string:
		if uint(index) < uint(len(v)) {
			return v[index], nil
		}
		return nil, errOutOfRangeCached
	case []int:
		if uint(index) < uint(len(v)) {
			return v[index], nil
		}
		return nil, errOutOfRangeCached
	case []int64:
		if uint(index) < uint(len(v)) {
			return v[index], nil
		}
		return nil, errOutOfRangeCached
	case []float64:
		if uint(index) < uint(len(v)) {
			return v[index], nil
		}
		return nil, errOutOfRangeCached
	case []bool:
		if uint(index) < uint(len(v)) {
			return v[index], nil
		}
		return nil, errOutOfRangeCached
	case []map[string]any:
		if uint(index) < uint(len(v)) {
			return v[index], nil
		}
		return nil, errOutOfRangeCached
	default:
		return nil, errInvalidSliceCached
	}
}
