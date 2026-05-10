package anyx

import (
	"fmt"
	"strings"
	"sync"
	"testing"
)

// 原始实现（基于 mapGetWithSeparator）
func mapExistsOriginal(m map[string]any, key string) bool {
	_, err := mapGetWithSeparator(m, key, ".")
	return err == nil
}

// 方案1: 直接使用 MapExists（当前实现）
func benchmarkMapExistsCurrent(b *testing.B, m map[string]any, key string) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = MapExists(m, key)
	}
}

// 方案2: 内联单层 key 检查（优化简单场景）
func mapExistsOptimized1(m map[string]any, key string) bool {
	// 快速路径：单层 key
	if !strings.Contains(key, ".") && !strings.Contains(key, "[") {
		_, ok := m[key]
		return ok
	}
	// 复杂场景回退到原实现
	_, err := mapGetWithSeparator(m, key, ".")
	return err == nil
}

// 方案3: 完全内联实现（避免函数调用开销）
func mapExistsOptimized2(m map[string]any, key string) bool {
	if len(m) == 0 || key == "" {
		return false
	}

	// 单层快速路径
	if !strings.Contains(key, ".") && !strings.Contains(key, "[") {
		_, ok := m[key]
		return ok
	}

	// 内联嵌套解析
	parts := splitKey(key, ".")
	if len(parts) == 0 {
		return false
	}

	current := any(m)
	for _, part := range parts {
		// 数组索引检查
		if len(part) > 2 && part[0] == '[' && part[len(part)-1] == ']' {
			// 简化版数组访问（仅 []any）
			indexStr := part[1 : len(part)-1]
			index := 0
			for _, c := range indexStr {
				if c < '0' || c > '9' {
					return false
				}
				index = index*10 + int(c-'0')
			}

			slice, ok := current.([]any)
			if !ok || index < 0 || index >= len(slice) {
				return false
			}
			current = slice[index]
		} else {
			nested, ok := current.(map[string]any)
			if !ok {
				return false
			}
			val, ok := nested[part]
			if !ok {
				return false
			}
			current = val
		}
	}

	return true
}

// 方案4: 预编译 key 路径（适用于重复查询）
type keyPath struct {
	parts []string
}

func (kp *keyPath) exists(m map[string]any) bool {
	current := any(m)
	for _, part := range kp.parts {
		if len(part) > 2 && part[0] == '[' && part[len(part)-1] == ']' {
			indexStr := part[1 : len(part)-1]
			index := 0
			for _, c := range indexStr {
				if c < '0' || c > '9' {
					return false
				}
				index = index*10 + int(c-'0')
			}

			slice, ok := current.([]any)
			if !ok || index < 0 || index >= len(slice) {
				return false
			}
			current = slice[index]
		} else {
			nested, ok := current.(map[string]any)
			if !ok {
				return false
			}
			val, ok := nested[part]
			if !ok {
				return false
			}
			current = val
		}
	}
	return true
}

func compileKeyPath(key string) *keyPath {
	parts := splitKey(key, ".")
	return &keyPath{parts: parts}
}

func mapExistsOptimized4(m map[string]any, key string) bool {
	path := compileKeyPath(key)
	return path.exists(m)
}

// 方案5: 零分配实现（避免内存分配）
func mapExistsOptimized5(m map[string]any, key string) bool {
	if len(m) == 0 || key == "" {
		return false
	}

	// 单层快速路径
	firstDot := strings.IndexByte(key, '.')
	firstBracket := strings.IndexByte(key, '[')

	if firstDot == -1 && firstBracket == -1 {
		_, ok := m[key]
		return ok
	}

	// 手动解析，避免 splitKey 分配
	var current any = m
	start := 0
	for i := 0; i <= len(key); i++ {
		if i == len(key) || key[i] == '.' {
			part := key[start:i]

			// 检查是否是数组索引
			if len(part) > 2 && part[0] == '[' && part[len(part)-1] == ']' {
				indexStr := part[1 : len(part)-1]
				index := 0
				for _, c := range indexStr {
					if c < '0' || c > '9' {
						return false
					}
					index = index*10 + int(c-'0')
				}

				slice, ok := current.([]any)
				if !ok || index < 0 || index >= len(slice) {
					return false
				}
				current = slice[index]
			} else {
				nested, ok := current.(map[string]any)
				if !ok {
					return false
				}
				val, ok := nested[part]
				if !ok {
					return false
				}
				current = val
			}
			start = i + 1
		}
	}

	return true
}

// 方案6: 混合策略（根据 key 长度选择算法）
func mapExistsOptimized6(m map[string]any, key string) bool {
	if len(m) == 0 || key == "" {
		return false
	}

	// 超短 key 直接检查
	if len(key) < 32 && !strings.Contains(key, ".") && !strings.Contains(key, "[") {
		_, ok := m[key]
		return ok
	}

	// 中等长度使用优化内联实现
	if len(key) < 128 {
		return mapExistsOptimized5(m, key)
	}

	// 长 key 使用原实现（正确性优先）
	_, err := mapGetWithSeparator(m, key, ".")
	return err == nil
}

// 方案7: 基于反射的通用实现（处理更多类型）
func mapExistsOptimized7(m map[string]any, key string) bool {
	if len(m) == 0 || key == "" {
		return false
	}

	if !strings.Contains(key, ".") && !strings.Contains(key, "[") {
		_, ok := m[key]
		return ok
	}

	parts := strings.Split(key, ".")
	current := any(m)

	for _, part := range parts {
		// 处理数组索引
		if idx := strings.IndexByte(part, '['); idx != -1 {
			mapKey := part[:idx]
			if mapKey != "" {
				nested, ok := current.(map[string]any)
				if !ok {
					return false
				}
				val, exists := nested[mapKey]
				if !exists {
					return false
				}
				current = val
			}

			// 解析索引
			indexStr := part[idx+1 : len(part)-1]
			index := 0
			for _, c := range indexStr {
				if c < '0' || c > '9' {
					return false
				}
				index = index*10 + int(c-'0')
			}

			slice, ok := current.([]any)
			if !ok || index < 0 || index >= len(slice) {
				return false
			}
			current = slice[index]
		} else {
			nested, ok := current.(map[string]any)
			if !ok {
				return false
			}
			val, ok := nested[part]
			if !ok {
				return false
			}
			current = val
		}
	}

	return true
}

// 方案8: 使用 sync.Pool 复用 keyPath 对象
var keyPathPool = make(chan *keyPath, 100)

func mapExistsOptimized8(m map[string]any, key string) bool {
	if len(m) == 0 || key == "" {
		return false
	}

	if !strings.Contains(key, ".") && !strings.Contains(key, "[") {
		_, ok := m[key]
		return ok
	}

	// 尝试从池中获取
	select {
	case path := <-keyPathPool:
		// 重用路径对象（简化版，实际需要更复杂的逻辑）
		result := path.exists(m)
		keyPathPool <- path
		return result
	default:
		// 池为空，创建新的
		return mapExistsOptimized4(m, key)
	}
}

// 方案9: 并发安全缓存（缓存解析结果）
type keyPathCache struct {
	sync.RWMutex
	cache map[string]*keyPath
}

var globalKeyPathCache = &keyPathCache{
	cache: make(map[string]*keyPath, 1000),
}

func mapExistsOptimized9(m map[string]any, key string) bool {
	if len(m) == 0 || key == "" {
		return false
	}

	if !strings.Contains(key, ".") && !strings.Contains(key, "[") {
		_, ok := m[key]
		return ok
	}

	// 尝试从缓存获取
	globalKeyPathCache.RLock()
	path, ok := globalKeyPathCache.cache[key]
	globalKeyPathCache.RUnlock()

	if !ok {
		// 编译并缓存
		path = compileKeyPath(key)
		globalKeyPathCache.Lock()
		globalKeyPathCache.cache[key] = path
		globalKeyPathCache.Unlock()
	}

	return path.exists(m)
}

// 方案10: 分支预测优化（基于常见模式）
func mapExistsOptimized10(m map[string]any, key string) bool {
	// 最常见情况：单层 key，直接存在
	if len(m) == 0 || key == "" {
		return false
	}

	// 快速路径：单层 key（80%+ 的场景）
	if key[0] != '.' && key[len(key)-1] != '.' {
		firstDot := strings.IndexByte(key, '.')
		if firstDot == -1 {
			_, ok := m[key]
			return ok
		}
	}

	// 复杂嵌套场景
	_, err := mapGetWithSeparator(m, key, ".")
	return err == nil
}

// ========== 测试数据 ==========

// 创建简单 map（单层）
func createSimpleMap() map[string]any {
	return map[string]any{
		"name":  "John",
		"age":   30,
		"email": "john@example.com",
	}
}

// 创建嵌套 map（2-3 层）
func createNestedMap() map[string]any {
	return map[string]any{
		"user": map[string]any{
			"name": "John",
			"address": map[string]any{
				"city":    "New York",
				"country": "USA",
			},
		},
		"settings": map[string]any{
			"theme": "dark",
			"lang":  "en",
		},
	}
}

// 创建深度嵌套 map（5+ 层）
func createDeepNestedMap() map[string]any {
	return map[string]any{
		"level1": map[string]any{
			"level2": map[string]any{
				"level3": map[string]any{
					"level4": map[string]any{
						"level5": map[string]any{
							"value": "deep",
						},
					},
				},
			},
		},
	}
}

// 创建包含数组的 map
func createArrayMap() map[string]any {
	return map[string]any{
		"items": []any{"a", "b", "c"},
		"nested": map[string]any{
			"array": []any{1, 2, 3},
		},
		"mixed": []any{
			map[string]any{"name": "item1"},
			map[string]any{"name": "item2"},
		},
	}
}

// 创建大型 map（100+ 键）
func createLargeMap() map[string]any {
	m := make(map[string]any, 100)
	for i := 0; i < 100; i++ {
		m[fmt.Sprintf("key%d", i)] = fmt.Sprintf("value%d", i)
	}
	return m
}

// 创建复杂混合 map
func createComplexMap() map[string]any {
	return map[string]any{
		"simple": "value",
		"nested": map[string]any{
			"deep": map[string]any{
				"value": 123,
			},
			"array": []any{1, 2, 3},
		},
		"list": []any{
			map[string]any{"id": 1},
			map[string]any{"id": 2},
		},
		"users": []map[string]any{
			{"name": "Alice", "age": 25},
			{"name": "Bob", "age": 30},
		},
	}
}

// ========== Benchmark 函数 ==========

// Benchmark 1: 简单 key（单层）- 存在
func BenchmarkMapExists_SimpleKey_Exists_Current(b *testing.B) {
	m := createSimpleMap()
	benchmarkMapExistsCurrent(b, m, "name")
}

func BenchmarkMapExists_SimpleKey_Exists_Opt1(b *testing.B) {
	m := createSimpleMap()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = mapExistsOptimized1(m, "name")
	}
}

func BenchmarkMapExists_SimpleKey_Exists_Opt2(b *testing.B) {
	m := createSimpleMap()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = mapExistsOptimized2(m, "name")
	}
}

func BenchmarkMapExists_SimpleKey_Exists_Opt5(b *testing.B) {
	m := createSimpleMap()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = mapExistsOptimized5(m, "name")
	}
}

func BenchmarkMapExists_SimpleKey_Exists_Opt6(b *testing.B) {
	m := createSimpleMap()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = mapExistsOptimized6(m, "name")
	}
}

func BenchmarkMapExists_SimpleKey_Exists_Opt10(b *testing.B) {
	m := createSimpleMap()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = mapExistsOptimized10(m, "name")
	}
}

// Benchmark 2: 简单 key（单层）- 不存在
func BenchmarkMapExists_SimpleKey_NotExists_Current(b *testing.B) {
	m := createSimpleMap()
	benchmarkMapExistsCurrent(b, m, "missing")
}

func BenchmarkMapExists_SimpleKey_NotExists_Opt5(b *testing.B) {
	m := createSimpleMap()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = mapExistsOptimized5(m, "missing")
	}
}

func BenchmarkMapExists_SimpleKey_NotExists_Opt6(b *testing.B) {
	m := createSimpleMap()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = mapExistsOptimized6(m, "missing")
	}
}

// Benchmark 3: 嵌套 key（2 层）- 存在
func BenchmarkMapExists_Nested2_Exists_Current(b *testing.B) {
	m := createNestedMap()
	benchmarkMapExistsCurrent(b, m, "user.name")
}

func BenchmarkMapExists_Nested2_Exists_Opt1(b *testing.B) {
	m := createNestedMap()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = mapExistsOptimized1(m, "user.name")
	}
}

func BenchmarkMapExists_Nested2_Exists_Opt2(b *testing.B) {
	m := createNestedMap()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = mapExistsOptimized2(m, "user.name")
	}
}

func BenchmarkMapExists_Nested2_Exists_Opt5(b *testing.B) {
	m := createNestedMap()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = mapExistsOptimized5(m, "user.name")
	}
}

func BenchmarkMapExists_Nested2_Exists_Opt6(b *testing.B) {
	m := createNestedMap()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = mapExistsOptimized6(m, "user.name")
	}
}

// Benchmark 4: 深度嵌套 key（5 层）- 存在
func BenchmarkMapExists_Nested5_Exists_Current(b *testing.B) {
	m := createDeepNestedMap()
	benchmarkMapExistsCurrent(b, m, "level1.level2.level3.level4.level5.value")
}

func BenchmarkMapExists_Nested5_Exists_Opt5(b *testing.B) {
	m := createDeepNestedMap()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = mapExistsOptimized5(m, "level1.level2.level3.level4.level5.value")
	}
}

func BenchmarkMapExists_Nested5_Exists_Opt6(b *testing.B) {
	m := createDeepNestedMap()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = mapExistsOptimized6(m, "level1.level2.level3.level4.level5.value")
	}
}

// Benchmark 5: 数组索引 - 存在
func BenchmarkMapExists_ArrayIndex_Exists_Current(b *testing.B) {
	m := createArrayMap()
	benchmarkMapExistsCurrent(b, m, "items[0]")
}

func BenchmarkMapExists_ArrayIndex_Exists_Opt5(b *testing.B) {
	m := createArrayMap()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = mapExistsOptimized5(m, "items[0]")
	}
}

func BenchmarkMapExists_ArrayIndex_Exists_Opt6(b *testing.B) {
	m := createArrayMap()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = mapExistsOptimized6(m, "items[0]")
	}
}

// Benchmark 6: 嵌套 + 数组混合
func BenchmarkMapExists_NestedArray_Exists_Current(b *testing.B) {
	m := createArrayMap()
	benchmarkMapExistsCurrent(b, m, "nested.array[1]")
}

func BenchmarkMapExists_NestedArray_Exists_Opt5(b *testing.B) {
	m := createArrayMap()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = mapExistsOptimized5(m, "nested.array[1]")
	}
}

func BenchmarkMapExists_NestedArray_Exists_Opt6(b *testing.B) {
	m := createArrayMap()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = mapExistsOptimized6(m, "nested.array[1]")
	}
}

// Benchmark 7: 大型 map（100 键）
func BenchmarkMapExists_LargeMap_Current(b *testing.B) {
	m := createLargeMap()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = MapExists(m, "key50")
	}
}

func BenchmarkMapExists_LargeMap_Opt5(b *testing.B) {
	m := createLargeMap()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = mapExistsOptimized5(m, "key50")
	}
}

func BenchmarkMapExists_LargeMap_Opt6(b *testing.B) {
	m := createLargeMap()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = mapExistsOptimized6(m, "key50")
	}
}

// Benchmark 8: 空 map
func BenchmarkMapExists_EmptyMap_Current(b *testing.B) {
	m := map[string]any{}
	benchmarkMapExistsCurrent(b, m, "key")
}

func BenchmarkMapExists_EmptyMap_Opt5(b *testing.B) {
	m := map[string]any{}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = mapExistsOptimized5(m, "key")
	}
}

func BenchmarkMapExists_EmptyMap_Opt6(b *testing.B) {
	m := map[string]any{}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = mapExistsOptimized6(m, "key")
	}
}

// Benchmark 9: 复杂混合场景
func BenchmarkMapExists_Complex_Current(b *testing.B) {
	m := createComplexMap()
	benchmarkMapExistsCurrent(b, m, "nested.array[1]")
}

func BenchmarkMapExists_Complex_Opt5(b *testing.B) {
	m := createComplexMap()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = mapExistsOptimized5(m, "nested.array[1]")
	}
}

func BenchmarkMapExists_Complex_Opt6(b *testing.B) {
	m := createComplexMap()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = mapExistsOptimized6(m, "nested.array[1]")
	}
}

// Benchmark 10: 预编译路径（方案 4）
func BenchmarkMapExists_Precompiled_Path(b *testing.B) {
	m := createNestedMap()
	path := compileKeyPath("user.name")
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = path.exists(m)
	}
}

// Benchmark 11: 并发场景
func BenchmarkMapExists_Concurrent_Current(b *testing.B) {
	m := createNestedMap()
	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			_ = MapExists(m, "user.name")
		}
	})
}

func BenchmarkMapExists_Concurrent_Opt5(b *testing.B) {
	m := createNestedMap()
	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			_ = mapExistsOptimized5(m, "user.name")
		}
	})
}

func BenchmarkMapExists_Concurrent_Opt6(b *testing.B) {
	m := createNestedMap()
	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			_ = mapExistsOptimized6(m, "user.name")
		}
	})
}

// Benchmark 12: 对比所有方案（简单 key）
func BenchmarkMapExists_AllOptions_Simple(b *testing.B) {
	m := createSimpleMap()

	b.Run("Current", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = MapExists(m, "name")
		}
	})

	b.Run("Opt1", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = mapExistsOptimized1(m, "name")
		}
	})

	b.Run("Opt2", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = mapExistsOptimized2(m, "name")
		}
	})

	b.Run("Opt5", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = mapExistsOptimized5(m, "name")
		}
	})

	b.Run("Opt6", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = mapExistsOptimized6(m, "name")
		}
	})

	b.Run("Opt10", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = mapExistsOptimized10(m, "name")
		}
	})
}

// Benchmark 13: 对比所有方案（嵌套 key）
func BenchmarkMapExists_AllOptions_Nested(b *testing.B) {
	m := createNestedMap()

	b.Run("Current", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = MapExists(m, "user.address.city")
		}
	})

	b.Run("Opt1", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = mapExistsOptimized1(m, "user.address.city")
		}
	})

	b.Run("Opt2", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = mapExistsOptimized2(m, "user.address.city")
		}
	})

	b.Run("Opt5", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = mapExistsOptimized5(m, "user.address.city")
		}
	})

	b.Run("Opt6", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = mapExistsOptimized6(m, "user.address.city")
		}
	})
}
