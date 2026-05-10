package anyx

import (
	"strconv"
	"testing"
)

// benchmarkMapExistsWithSepSimple 简单 key 基准测试（最常见场景）
func benchmarkMapExistsWithSepSimple(b *testing.B, sep string) {
	m := map[string]any{
		"name": "value",
		"id":   123,
		"flag": true,
	}

	b.ResetTimer()
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		MapExistsWithSep(m, "name", sep)
	}
}

func BenchmarkMapExistsWithSepSimpleDot(b *testing.B) {
	benchmarkMapExistsWithSepSimple(b, ".")
}

func BenchmarkMapExistsWithSepSimpleSlash(b *testing.B) {
	benchmarkMapExistsWithSepSimple(b, "/")
}

func BenchmarkMapExistsWithSepSimpleColon(b *testing.B) {
	benchmarkMapExistsWithSepSimple(b, "::")
}

func BenchmarkMapExistsWithSepSimpleDash(b *testing.B) {
	benchmarkMapExistsWithSepSimple(b, "-")
}

// benchmarkMapExistsWithSepNested 嵌套 key 基准测试
func benchmarkMapExistsWithSepNested(b *testing.B, sep string, depth int) {
	m := buildNestedMap(depth, sep)

	key := buildNestedKey(depth, sep)

	b.ResetTimer()
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		MapExistsWithSep(m, key, sep)
	}
}

func BenchmarkMapExistsWithSepNested2Dot(b *testing.B) {
	benchmarkMapExistsWithSepNested(b, ".", 2)
}

func BenchmarkMapExistsWithSepNested3Dot(b *testing.B) {
	benchmarkMapExistsWithSepNested(b, ".", 3)
}

func BenchmarkMapExistsWithSepNested5Dot(b *testing.B) {
	benchmarkMapExistsWithSepNested(b, ".", 5)
}

func BenchmarkMapExistsWithSepNested10Dot(b *testing.B) {
	benchmarkMapExistsWithSepNested(b, ".", 10)
}

func BenchmarkMapExistsWithSepNested2Slash(b *testing.B) {
	benchmarkMapExistsWithSepNested(b, "/", 2)
}

func BenchmarkMapExistsWithSepNested5Slash(b *testing.B) {
	benchmarkMapExistsWithSepNested(b, "/", 5)
}

// BenchmarkMapExistsWithSepArrayIndex 数组索引基准测试
func BenchmarkMapExistsWithSepArrayIndex(b *testing.B) {
	m := map[string]any{
		"items": []any{"a", "b", "c", "d", "e"},
	}

	b.ResetTimer()
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		MapExistsWithSep(m, "items[2]", ".")
	}
}

// BenchmarkMapExistsWithSepNestedArrayIndex 嵌套 + 数组索引
func BenchmarkMapExistsWithSepNestedArrayIndex(b *testing.B) {
	m := map[string]any{
		"data": map[string]any{
			"items": []any{
				map[string]any{"id": 1},
				map[string]any{"id": 2},
				map[string]any{"id": 3},
			},
		},
	}

	b.ResetTimer()
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		MapExistsWithSep(m, "data.items[1].id", ".")
	}
}

// BenchmarkMapExistsWithSepMultipleArrayIndexes 多个数组索引
func BenchmarkMapExistsWithSepMultipleArrayIndexes(b *testing.B) {
	m := map[string]any{
		"matrix": []any{
			[]any{1, 2, 3},
			[]any{4, 5, 6},
			[]any{7, 8, 9},
		},
	}

	b.ResetTimer()
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		MapExistsWithSep(m, "matrix[1][2]", ".")
	}
}

// BenchmarkMapExistsWithSepNotFound 不存在的 key
func BenchmarkMapExistsWithSepNotFound(b *testing.B) {
	m := map[string]any{
		"name": "value",
		"nested": map[string]any{
			"key": "value",
		},
	}

	b.ResetTimer()
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		MapExistsWithSep(m, "nonexistent", ".")
	}
}

// BenchmarkMapExistsWithSepNestedNotFound 嵌套路径中不存在
func BenchmarkMapExistsWithSepNestedNotFound(b *testing.B) {
	m := map[string]any{
		"data": map[string]any{
			"items": []any{"a", "b"},
		},
	}

	b.ResetTimer()
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		MapExistsWithSep(m, "data.nonexistent[0]", ".")
	}
}

// BenchmarkMapExistsWithSepEmptyMap 空 map
func BenchmarkMapExistsWithSepEmptyMap(b *testing.B) {
	m := map[string]any{}

	b.ResetTimer()
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		MapExistsWithSep(m, "key", ".")
	}
}

// BenchmarkMapExistsWithSepEmptyKey 空 key
func BenchmarkMapExistsWithSepEmptyKey(b *testing.B) {
	m := map[string]any{
		"key": "value",
	}

	b.ResetTimer()
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		MapExistsWithSep(m, "", ".")
	}
}

// BenchmarkMapExistsWithSepLargeMap 大型 map（100 个键）
func BenchmarkMapExistsWithSepLargeMap(b *testing.B) {
	m := make(map[string]any, 100)
	for i := 0; i < 100; i++ {
		m["key"+strconv.Itoa(i)] = i
	}

	b.ResetTimer()
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		MapExistsWithSep(m, "key50", ".")
	}
}

// BenchmarkMapExistsWithSepLargeNestedMap 大型嵌套 map
func BenchmarkMapExistsWithSepLargeNestedMap(b *testing.B) {
	m := map[string]any{
		"level1": map[string]any{
			"level2": map[string]any{
				"level3": map[string]any{
					"level4": map[string]any{
						"level5": map[string]any{
							"target": "value",
						},
					},
				},
			},
		},
	}

	b.ResetTimer()
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		MapExistsWithSep(m, "level1.level2.level3.level4.level5.target", ".")
	}
}

// BenchmarkMapExistsWithSepComplexIndex 复杂索引（负数、大数）
func BenchmarkMapExistsWithSepComplexIndex(b *testing.B) {
	m := map[string]any{
		"items": make([]any, 100),
	}
	for i := 0; i < 100; i++ {
		m["items"].([]any)[i] = i
	}

	b.ResetTimer()
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		MapExistsWithSep(m, "items[99]", ".")
	}
}

// BenchmarkMapExistsWithSepSpecialChars 特殊字符（带括号但不完全是索引）
func BenchmarkMapExistsWithSepSpecialChars(b *testing.B) {
	m := map[string]any{
		"data": map[string]any{
			"items[0]": "value",
		},
	}

	b.ResetTimer()
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		MapExistsWithSep(m, "data.items[0]", ".")
	}
}

// BenchmarkMapExistsWithSepConcurrent 并发场景
func BenchmarkMapExistsWithSepConcurrent(b *testing.B) {
	m := map[string]any{
		"name":  "value",
		"nested": map[string]any{
			"key": "value",
		},
		"items": []any{"a", "b", "c"},
	}

	b.ResetTimer()
	b.ReportAllocs()

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			MapExistsWithSep(m, "nested.key", ".")
		}
	})
}

// BenchmarkMapExistsWithSepMixedPathLengths 混合不同路径长度
func BenchmarkMapExistsWithSepMixedPathLengths(b *testing.B) {
	m := map[string]any{
		"a": map[string]any{
			"b": map[string]any{
				"c": map[string]any{
					"d": map[string]any{
						"e": "value",
					},
				},
			},
		},
	}

	tests := []string{
		"a",
		"a.b",
		"a.b.c",
		"a.b.c.d",
		"a.b.c.d.e",
	}

	b.ResetTimer()
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		for _, key := range tests {
			MapExistsWithSep(m, key, ".")
		}
	}
}

// BenchmarkMapExistsWithSepDifferentSeparators 不同分隔符性能对比
func BenchmarkMapExistsWithSepDifferentSeparators(b *testing.B) {
	m := map[string]any{
		"level1": map[string]any{
			"level2": map[string]any{
				"level3": "value",
			},
		},
	}

	b.Run("dot", func(b *testing.B) {
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			MapExistsWithSep(m, "level1.level2.level3", ".")
		}
	})

	b.Run("slash", func(b *testing.B) {
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			MapExistsWithSep(m, "level1/level2/level3", "/")
		}
	})

	b.Run("double_colon", func(b *testing.B) {
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			MapExistsWithSep(m, "level1::level2::level3", "::")
		}
	})

	b.Run("dash", func(b *testing.B) {
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			MapExistsWithSep(m, "level1-level2-level3", "-")
		}
	})
}

// BenchmarkMapExistsWithSepEdgeCases 边界情况
func BenchmarkMapExistsWithSepEdgeCases(b *testing.B) {
	m := map[string]any{
		"": "empty_key",
		".": map[string]any{
			"": "nested_empty",
		},
	}

	b.Run("empty_key", func(b *testing.B) {
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			MapExistsWithSep(m, "", ".")
		}
	})

	b.Run("dot_separator_key", func(b *testing.B) {
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			MapExistsWithSep(m, ".", ".")
		}
	})

	b.Run("nested_empty_key", func(b *testing.B) {
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			MapExistsWithSep(m, "..", ".")
		}
	})
}

// BenchmarkMapExistsWithSepArrayEdgeCases 数组边界情况
func BenchmarkMapExistsWithSepArrayEdgeCases(b *testing.B) {
	m := map[string]any{
		"items": []any{"a", "b", "c"},
	}

	b.Run("first_element", func(b *testing.B) {
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			MapExistsWithSep(m, "items[0]", ".")
		}
	})

	b.Run("last_element", func(b *testing.B) {
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			MapExistsWithSep(m, "items[2]", ".")
		}
	})

	b.Run("out_of_bounds", func(b *testing.B) {
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			MapExistsWithSep(m, "items[10]", ".")
		}
	})

	b.Run("negative_index", func(b *testing.B) {
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			MapExistsWithSep(m, "items[-1]", ".")
		}
	})
}

// BenchmarkMapExistsWithSepRealWorldScenarios 真实场景模拟
func BenchmarkMapExistsWithSepRealWorldScenarios(b *testing.B) {
	// 模拟配置文件
	config := map[string]any{
		"server": map[string]any{
			"host": "localhost",
			"port": 8080,
			"ssl": map[string]any{
				"enabled": true,
				"cert": "/path/to/cert",
			},
		},
		"database": map[string]any{
			"connections": []any{
				map[string]any{"host": "db1", "port": 5432},
				map[string]any{"host": "db2", "port": 5432},
			},
		},
		"features": []any{"auth", "cache", "logging"},
	}

	b.Run("config_simple", func(b *testing.B) {
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			MapExistsWithSep(config, "server.host", ".")
		}
	})

	b.Run("config_deep", func(b *testing.B) {
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			MapExistsWithSep(config, "server.ssl.cert", ".")
		}
	})

	b.Run("config_array", func(b *testing.B) {
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			MapExistsWithSep(config, "database.connections[0].host", ".")
		}
	})

	b.Run("config_feature", func(b *testing.B) {
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			MapExistsWithSep(config, "features[2]", ".")
		}
	})
}

// 辅助函数
func buildNestedMap(depth int, sep string) map[string]any {
	if depth == 1 {
		return map[string]any{"value": "target"}
	}
	return map[string]any{"level": buildNestedMap(depth-1, sep)}
}

func buildNestedKey(depth int, sep string) string {
	if depth == 1 {
		return "value"
	}
	return "level" + sep + buildNestedKey(depth-1, sep)
}
