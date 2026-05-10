package anyx

import (
	"fmt"
	"testing"
)

// ============================================================
// accessMapKey 性能优化 - 多方案对比测试
// ============================================================

// 测试数据准备
var (
	// 小型 map[string]any
	benchMapStringAnySmall = map[string]any{
		"name": "John",
		"age":  30,
	}

	// 中型 map[string]any
	benchMapStringAnyMedium = func() map[string]any {
		m := make(map[string]any, 50)
		for i := 0; i < 50; i++ {
			m[fmt.Sprintf("key%d", i)] = i
		}
		return m
	}()

	// 大型 map[string]any
	benchMapStringAnyLarge = func() map[string]any {
		m := make(map[string]any, 1000)
		for i := 0; i < 1000; i++ {
			m[fmt.Sprintf("key%d", i)] = i
		}
		return m
	}()

	// map[any]any
	benchMapAnyAny = map[any]any{
		"name": "John",
		"age":  30,
		42:     "answer",
	}

	// 无效类型（用于错误路径测试）
	benchInvalidType = []string{"a", "b", "c"}
)

// ============================================================
// 方案 0: 原始实现（baseline）
// ============================================================

func accessMapKeyOriginal(current any, key string) (any, error) {
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
// 方案 3: 内联返回（减少局部变量）
// ============================================================

func accessMapKeyInline(current any, key string) (any, error) {
	switch v := current.(type) {
	case map[string]any:
		if val, ok := v[key]; ok {
			return val, nil
		}
		return nil, ErrNotFound
	case map[any]any:
		if val, ok := v[key]; ok {
			return val, nil
		}
		return nil, ErrNotFound
	default:
		return nil, fmt.Errorf("%w: cannot access key '%s' on type %T", ErrInvalidMapType, key, current)
	}
}

// ============================================================
// 方案 6: 简化错误消息（减少格式化开销）
// ============================================================

func accessMapKeySimpleError(current any, key string) (any, error) {
	switch v := current.(type) {
	case map[string]any:
		if val, ok := v[key]; ok {
			return val, nil
		}
		return nil, ErrNotFound
	case map[any]any:
		if val, ok := v[key]; ok {
			return val, nil
		}
		return nil, ErrNotFound
	default:
		return nil, ErrInvalidMapType
	}
}

// ============================================================
// 方案 10: 快速失败（提前返回错误）
// ============================================================

func accessMapKeyFastFail(current any, key string) (any, error) {
	// 快速失败：检查 key 是否为空
	if key == "" {
		return nil, ErrNotFound
	}

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
// Benchmark 测试用例
// ============================================================

// 场景 1: 小型 map[string]any，键命中
func BenchmarkAccessMapKey_SmallMapHit_Original(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = accessMapKeyOriginal(benchMapStringAnySmall, "name")
	}
}

func BenchmarkAccessMapKey_SmallMapHit_Inline(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = accessMapKeyInline(benchMapStringAnySmall, "name")
	}
}

func BenchmarkAccessMapKey_SmallMapHit_SimpleError(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = accessMapKeySimpleError(benchMapStringAnySmall, "name")
	}
}

func BenchmarkAccessMapKey_SmallMapHit_FastFail(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = accessMapKeyFastFail(benchMapStringAnySmall, "name")
	}
}

// 场景 2: 小型 map[string]any，键未命中
func BenchmarkAccessMapKey_SmallMapMiss_Original(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = accessMapKeyOriginal(benchMapStringAnySmall, "nonexistent")
	}
}

func BenchmarkAccessMapKey_SmallMapMiss_Inline(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = accessMapKeyInline(benchMapStringAnySmall, "nonexistent")
	}
}

func BenchmarkAccessMapKey_SmallMapMiss_SimpleError(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = accessMapKeySimpleError(benchMapStringAnySmall, "nonexistent")
	}
}

// 场景 3: 中型 map[string]any
func BenchmarkAccessMapKey_MediumMap_Original(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = accessMapKeyOriginal(benchMapStringAnyMedium, "key25")
	}
}

func BenchmarkAccessMapKey_MediumMap_Inline(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = accessMapKeyInline(benchMapStringAnyMedium, "key25")
	}
}

func BenchmarkAccessMapKey_MediumMap_SimpleError(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = accessMapKeySimpleError(benchMapStringAnyMedium, "key25")
	}
}

// 场景 4: 大型 map[string]any
func BenchmarkAccessMapKey_LargeMap_Original(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = accessMapKeyOriginal(benchMapStringAnyLarge, "key500")
	}
}

func BenchmarkAccessMapKey_LargeMap_Inline(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = accessMapKeyInline(benchMapStringAnyLarge, "key500")
	}
}

func BenchmarkAccessMapKey_LargeMap_SimpleError(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = accessMapKeySimpleError(benchMapStringAnyLarge, "key500")
	}
}

// 场景 5: map[any]any
func BenchmarkAccessMapKey_MapAnyAny_Original(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = accessMapKeyOriginal(benchMapAnyAny, "name")
	}
}

func BenchmarkAccessMapKey_MapAnyAny_Inline(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = accessMapKeyInline(benchMapAnyAny, "name")
	}
}

func BenchmarkAccessMapKey_MapAnyAny_SimpleError(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = accessMapKeySimpleError(benchMapAnyAny, "name")
	}
}

// 场景 6: 错误路径（无效类型）
func BenchmarkAccessMapKey_ErrorPath_Original(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = accessMapKeyOriginal(benchInvalidType, "key")
	}
}

func BenchmarkAccessMapKey_ErrorPath_Inline(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = accessMapKeyInline(benchInvalidType, "key")
	}
}

func BenchmarkAccessMapKey_ErrorPath_SimpleError(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = accessMapKeySimpleError(benchInvalidType, "key")
	}
}

// 场景 7: 并发访问（测试安全性）
func BenchmarkAccessMapKey_Concurrent_Original(b *testing.B) {
	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			_, _ = accessMapKeyOriginal(benchMapStringAnyMedium, "key25")
		}
	})
}

func BenchmarkAccessMapKey_Concurrent_Inline(b *testing.B) {
	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			_, _ = accessMapKeyInline(benchMapStringAnyMedium, "key25")
		}
	})
}

func BenchmarkAccessMapKey_Concurrent_SimpleError(b *testing.B) {
	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			_, _ = accessMapKeySimpleError(benchMapStringAnyMedium, "key25")
		}
	})
}

// 场景 8: 空键（边界情况）
func BenchmarkAccessMapKey_EmptyKey_Original(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = accessMapKeyOriginal(benchMapStringAnySmall, "")
	}
}

func BenchmarkAccessMapKey_EmptyKey_FastFail(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = accessMapKeyFastFail(benchMapStringAnySmall, "")
	}
}

// 场景 9: 混合命中/未命中（模拟真实负载）
func BenchmarkAccessMapKey_Mixed_Original(b *testing.B) {
	keys := []string{"name", "age", "nonexistent1", "nonexistent2"}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = accessMapKeyOriginal(benchMapStringAnySmall, keys[i%4])
	}
}

func BenchmarkAccessMapKey_Mixed_Inline(b *testing.B) {
	keys := []string{"name", "age", "nonexistent1", "nonexistent2"}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = accessMapKeyInline(benchMapStringAnySmall, keys[i%4])
	}
}

func BenchmarkAccessMapKey_Mixed_SimpleError(b *testing.B) {
	keys := []string{"name", "age", "nonexistent1", "nonexistent2"}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = accessMapKeySimpleError(benchMapStringAnySmall, keys[i%4])
	}
}

// 场景 10: 不同键长度（测试字符串比较性能）
func BenchmarkAccessMapKey_LongKey_Original(b *testing.B) {
	longKey := "this_is_a_very_long_key_name_with_many_characters"
	m := map[string]any{longKey: "value"}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = accessMapKeyOriginal(m, longKey)
	}
}

func BenchmarkAccessMapKey_LongKey_Inline(b *testing.B) {
	longKey := "this_is_a_very_long_key_name_with_many_characters"
	m := map[string]any{longKey: "value"}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = accessMapKeyInline(m, longKey)
	}
}

func BenchmarkAccessMapKey_LongKey_SimpleError(b *testing.B) {
	longKey := "this_is_a_very_long_key_name_with_many_characters"
	m := map[string]any{longKey: "value"}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = accessMapKeySimpleError(m, longKey)
	}
}
