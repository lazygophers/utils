package anyx

import (
	"testing"
)

// 基准测试数据
var (
	benchMapSimple = map[string]any{
		"name": "John",
		"age":  30,
	}
	benchMapNested = map[string]any{
		"user": map[string]any{
			"profile": map[string]any{
				"name": "John",
				"age":  30,
			},
		},
	}
	benchMapDeep = map[string]any{
		"l1": map[string]any{
			"l2": map[string]any{
				"l3": map[string]any{
					"l4": map[string]any{
						"l5": "value",
					},
				},
			},
		},
	}
	benchMapArray = map[string]any{
		"items": []any{"a", "b", "c", "d", "e"},
	}
	benchMapNestedArray = map[string]any{
		"data": map[string]any{
			"items": []any{
				map[string]any{"name": "item1"},
				map[string]any{"name": "item2"},
				map[string]any{"name": "item3"},
			},
		},
	}
)

// MapGet 性能基准测试
func BenchmarkMapGet_Simple(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		MapGet(benchMapSimple, "name")
	}
}

func BenchmarkMapGet_Nested(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		MapGet(benchMapNested, "user.profile.name")
	}
}

func BenchmarkMapGet_Deep(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		MapGet(benchMapDeep, "l1.l2.l3.l4.l5")
	}
}

func BenchmarkMapGet_Array(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		MapGet(benchMapArray, "items[2]")
	}
}

func BenchmarkMapGet_NestedArray(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		MapGet(benchMapNestedArray, "data.items[1].name")
	}
}

// 对比基准：原始实现（保留用于对比）
func benchmarkMapGetOriginal(b *testing.B, m map[string]any, key string) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		mapGetWithSeparator(m, key, ".")
	}
}

func BenchmarkMapGet_Original_Simple(b *testing.B) {
	benchmarkMapGetOriginal(b, benchMapSimple, "name")
}

func BenchmarkMapGet_Original_Nested(b *testing.B) {
	benchmarkMapGetOriginal(b, benchMapNested, "user.profile.name")
}

func BenchmarkMapGet_Original_Deep(b *testing.B) {
	benchmarkMapGetOriginal(b, benchMapDeep, "l1.l2.l3.l4.l5")
}

func BenchmarkMapGet_Original_Array(b *testing.B) {
	benchmarkMapGetOriginal(b, benchMapArray, "items[2]")
}

func BenchmarkMapGet_Original_NestedArray(b *testing.B) {
	benchmarkMapGetOriginal(b, benchMapNestedArray, "data.items[1].name")
}
