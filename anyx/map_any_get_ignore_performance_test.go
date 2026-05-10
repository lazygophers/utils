package anyx

import (
	"testing"
)

// 性能对比测试 - 验证优化后的性能提升
func Benchmark_MapGetIgnore_OldImplementation(b *testing.B) {
	// 旧实现：直接调用 mapGetWithSeparator
	data := map[string]any{
		"a":      1,
		"nested": map[string]any{"x": int(10)},
		"deep":   map[string]any{"a": map[string]any{"b": map[string]any{"c": int(100)}}},
	}

	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		// 模拟旧实现
		_, _ = mapGetWithSeparator(data, "a", ".")
	}
}

func Benchmark_MapGetIgnore_NewImplementation(b *testing.B) {
	// 新实现：快速路径优化
	data := map[string]any{
		"a":      1,
		"nested": map[string]any{"x": int(10)},
		"deep":   map[string]any{"a": map[string]any{"b": map[string]any{"c": int(100)}}},
	}

	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_ = MapGetIgnore(data, "a")
	}
}

// 场景化基准测试
func Benchmark_MapGetIgnore_Scenarios(b *testing.B) {
	data := map[string]any{
		"a":           1,
		"b":           "string",
		"nested":      map[string]any{"x": int(10), "y": int(20)},
		"deep":        map[string]any{"a": map[string]any{"b": map[string]any{"c": int(100)}}},
		"arr":         []any{int(1), int(2), int(3), int(4), int(5)},
		"nested_arr":  map[string]any{"items": []any{map[string]any{"id": int(1)}, map[string]any{"id": int(2)}}},
		"complex_arr": map[string]any{"data": []map[string]any{{"x": []any{int(1), int(2), int(3)}}}},
	}

	scenarios := []struct {
		name string
		key  string
	}{
		{"SimpleKey", "a"},
		{"NestedKey", "nested.x"},
		{"DeepNest", "deep.a.b.c"},
		{"ArrayIndex", "arr.[2]"},
		{"NestedArray", "nested_arr.items.[0]"},
		{"ComplexArray", "complex_arr.data.[0].x.[1]"},
	}

	for _, scenario := range scenarios {
		b.Run(scenario.name, func(b *testing.B) {
			b.ReportAllocs()
			for i := 0; i < b.N; i++ {
				_ = MapGetIgnore(data, scenario.key)
			}
		})
	}
}

// 对比测试：旧实现 vs 新实现
func Benchmark_MapGetIgnore_Comparison(b *testing.B) {
	data := map[string]any{
		"a":      1,
		"nested": map[string]any{"x": int(10)},
		"deep":   map[string]any{"a": map[string]any{"b": map[string]any{"c": int(100)}}},
	}

	b.Run("Old_Simple", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			_, _ = mapGetWithSeparator(data, "a", ".")
		}
	})

	b.Run("New_Simple", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			_ = MapGetIgnore(data, "a")
		}
	})

	b.Run("Old_Nested", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			_, _ = mapGetWithSeparator(data, "nested.x", ".")
		}
	})

	b.Run("New_Nested", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			_ = MapGetIgnore(data, "nested.x")
		}
	})

	b.Run("Old_Deep", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			_, _ = mapGetWithSeparator(data, "deep.a.b.c", ".")
		}
	})

	b.Run("New_Deep", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			_ = MapGetIgnore(data, "deep.a.b.c")
		}
	})
}

// 快速路径效果测试
func Benchmark_MapGetIgnore_FastPathEffectiveness(b *testing.B) {
	data := map[string]any{
		"a":      1,
		"b":      2,
		"c":      3,
		"nested": map[string]any{"x": int(10)},
	}

	b.Run("FastPath_Simple", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			_ = MapGetIgnore(data, "a")
			_ = MapGetIgnore(data, "b")
			_ = MapGetIgnore(data, "c")
		}
	})

	b.Run("SlowPath_Nested", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			_ = MapGetIgnore(data, "nested.x")
		}
	})
}

// 零分配测试
func Benchmark_MapGetIgnore_ZeroAlloc(b *testing.B) {
	data := map[string]any{
		"a": 1,
	}

	b.Run("SimpleKey", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			_ = MapGetIgnore(data, "a")
		}
	})
}

// 预热测试（模拟真实使用场景）
func Benchmark_MapGetIgnore_RealWorld(b *testing.B) {
	// 模拟真实配置数据
	config := map[string]any{
		"app": map[string]any{
			"name":    "test",
			"version": "1.0.0",
			"debug":   false,
		},
		"database": map[string]any{
			"host":     "localhost",
			"port":     5432,
			"username": "user",
			"password": "pass",
		},
		"cache": map[string]any{
			"enabled": true,
			"ttl":     3600,
		},
	}

	// 模拟配置读取操作
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_ = MapGetIgnore(config, "app.name")
		_ = MapGetIgnore(config, "app.version")
		_ = MapGetIgnore(config, "database.host")
		_ = MapGetIgnore(config, "database.port")
		_ = MapGetIgnore(config, "cache.enabled")
		_ = MapGetIgnore(config, "cache.ttl")
	}
}

// 边界情况性能测试
func Benchmark_MapGetIgnore_EdgeCases(b *testing.B) {
	emptyMap := map[string]any{}
	normalMap := map[string]any{"a": 1}
	deepMap := map[string]any{
		"a": map[string]any{
			"b": map[string]any{
				"c": map[string]any{
					"d": map[string]any{
						"e": 1,
					},
				},
			},
		},
	}

	b.Run("EmptyMap", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			_ = MapGetIgnore(emptyMap, "a")
		}
	})

	b.Run("NonExistentKey", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			_ = MapGetIgnore(normalMap, "nonexistent")
		}
	})

	b.Run("VeryDeepNest", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			_ = MapGetIgnore(deepMap, "a.b.c.d.e")
		}
	})
}
