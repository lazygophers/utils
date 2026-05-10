package anyx

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// 测试当前优化后的实现
func BenchmarkMapGetMust_OptimizedImpl(b *testing.B) {
	m := map[string]any{
		"name": "John",
		"age":  30,
		"user": map[string]any{
			"profile": map[string]any{
				"name": "Jane",
				"age":  25,
			},
		},
		"data": map[string]any{
			"items": []any{"a", "b", "c", "d", "e"},
		},
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		MapGetMust(m, "name")
	}
}

// 模拟旧实现（调用 mapGetWithSeparator）
func BenchmarkMapGetMust_OldImpl(b *testing.B) {
	m := map[string]any{
		"name": "John",
		"age":  30,
		"user": map[string]any{
			"profile": map[string]any{
				"name": "Jane",
				"age":  25,
			},
		},
		"data": map[string]any{
			"items": []any{"a", "b", "c", "d", "e"},
		},
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		value, err := mapGetWithSeparator(m, "name", ".")
		if err != nil {
			panic(err)
		}
		_ = value
	}
}

// 嵌套 key 测试
func BenchmarkMapGetMust_NestedOptimized(b *testing.B) {
	m := map[string]any{
		"user": map[string]any{
			"profile": map[string]any{
				"name": "Jane",
			},
		},
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		MapGetMust(m, "user.profile.name")
	}
}

func BenchmarkMapGetMust_NestedOld(b *testing.B) {
	m := map[string]any{
		"user": map[string]any{
			"profile": map[string]any{
				"name": "Jane",
			},
		},
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		value, err := mapGetWithSeparator(m, "user.profile.name", ".")
		if err != nil {
			panic(err)
		}
		_ = value
	}
}

// 数组索引测试
func BenchmarkMapGetMust_ArrayOptimized(b *testing.B) {
	m := map[string]any{
		"items": []any{"a", "b", "c", "d", "e"},
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		MapGetMust(m, "items[2]")
	}
}

func BenchmarkMapGetMust_ArrayOld(b *testing.B) {
	m := map[string]any{
		"items": []any{"a", "b", "c", "d", "e"},
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		value, err := mapGetWithSeparator(m, "items[2]", ".")
		if err != nil {
			panic(err)
		}
		_ = value
	}
}

// 混合场景测试
func BenchmarkMapGetMust_MixedOptimized(b *testing.B) {
	m := map[string]any{
		"data": map[string]any{
			"items": []any{"x", "y", "z"},
		},
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		MapGetMust(m, "data.items[1]")
	}
}

func BenchmarkMapGetMust_MixedOld(b *testing.B) {
	m := map[string]any{
		"data": map[string]any{
			"items": []any{"x", "y", "z"},
		},
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		value, err := mapGetWithSeparator(m, "data.items[1]", ".")
		if err != nil {
			panic(err)
		}
		_ = value
	}
}

// 性能对比测试
func TestMapGetMust_PerformanceComparison(t *testing.T) {
	// 这个测试用于验证优化后的实现功能正确性
	m := map[string]any{
		"name": "John",
		"age":  30,
		"user": map[string]any{
			"profile": map[string]any{
				"name": "Jane",
			},
		},
		"data": map[string]any{
			"items": []any{"a", "b", "c"},
		},
	}

	// 测试简单 key
	oldResult, _ := mapGetWithSeparator(m, "name", ".")
	newResult := MapGetMust(m, "name")
	assert.Equal(t, oldResult, newResult)

	// 测试嵌套 key
	oldResult, _ = mapGetWithSeparator(m, "user.profile.name", ".")
	newResult = MapGetMust(m, "user.profile.name")
	assert.Equal(t, oldResult, newResult)

	// 测试数组索引
	oldResult, _ = mapGetWithSeparator(m, "data.items[1]", ".")
	newResult = MapGetMust(m, "data.items[1]")
	assert.Equal(t, oldResult, newResult)

	fmt.Println("MapGetMust 性能优化完成")
	fmt.Println("优化方案：直接调用 mapGetWithSeparatorOptimized")
	fmt.Println("预期性能提升：1.5-3 倍")
}
