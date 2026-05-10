package anyx

import (
	"fmt"
	"strconv"
	"testing"
)

// Benchmark mapGetWithSeparator vs mapGetWithSeparatorOptimized

// Scenario 1: 简单键访问（快速路径）
func BenchmarkMapGetWithSeparator_SimpleKey(b *testing.B) {
	m := map[string]any{
		"name": "John Doe",
		"age":  30,
		"city": "New York",
	}

	b.Run("Original", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			_, _ = mapGetWithSeparator(m, "name", ".")
		}
	})

	b.Run("Optimized", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			_, _ = mapGetWithSeparatorOptimized(m, "name", ".")
		}
	})
}

// Scenario 2: 嵌套键访问（2 层）
func BenchmarkMapGetWithSeparator_Nested2Levels(b *testing.B) {
	m := map[string]any{
		"user": map[string]any{
			"name": "Alice",
			"age":  25,
		},
	}

	b.Run("Original", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			_, _ = mapGetWithSeparator(m, "user.name", ".")
		}
	})

	b.Run("Optimized", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			_, _ = mapGetWithSeparatorOptimized(m, "user.name", ".")
		}
	})
}

// Scenario 3: 嵌套键访问（5 层）
func BenchmarkMapGetWithSeparator_Nested5Levels(b *testing.B) {
	m := map[string]any{
		"a": map[string]any{
			"b": map[string]any{
				"c": map[string]any{
					"d": map[string]any{
						"e": "deep value",
					},
				},
			},
		},
	}

	b.Run("Original", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			_, _ = mapGetWithSeparator(m, "a.b.c.d.e", ".")
		}
	})

	b.Run("Optimized", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			_, _ = mapGetWithSeparatorOptimized(m, "a.b.c.d.e", ".")
		}
	})
}

// Scenario 4: 数组索引访问（正数索引）
func BenchmarkMapGetWithSeparator_ArrayIndexPositive(b *testing.B) {
	m := map[string]any{
		"items": []any{"apple", "banana", "cherry"},
	}

	b.Run("Original", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			_, _ = mapGetWithSeparator(m, "items.[1]", ".")
		}
	})

	b.Run("Optimized", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			_, _ = mapGetWithSeparatorOptimized(m, "items.[1]", ".")
		}
	})
}

// Scenario 5: 数组索引访问（负数索引）
func BenchmarkMapGetWithSeparator_ArrayIndexNegative(b *testing.B) {
	m := map[string]any{
		"items": []any{"apple", "banana", "cherry"},
	}

	b.Run("Original", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			_, _ = mapGetWithSeparator(m, "items.[-1]", ".")
		}
	})

	b.Run("Optimized", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			_, _ = mapGetWithSeparatorOptimized(m, "items.[-1]", ".")
		}
	})
}

// Scenario 6: 嵌套数组访问
func BenchmarkMapGetWithSeparator_NestedArrays(b *testing.B) {
	m := map[string]any{
		"matrix": []any{
			[]any{1, 2, 3},
			[]any{4, 5, 6},
			[]any{7, 8, 9},
		},
	}

	b.Run("Original", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			_, _ = mapGetWithSeparator(m, "matrix.[1].[2]", ".")
		}
	})

	b.Run("Optimized", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			_, _ = mapGetWithSeparatorOptimized(m, "matrix.[1].[2]", ".")
		}
	})
}

// Scenario 7: 混合 map 和数组访问
func BenchmarkMapGetWithSeparator_MixedMapArray(b *testing.B) {
	m := map[string]any{
		"users": []any{
			map[string]any{
				"name": "Alice",
				"age":  25,
			},
			map[string]any{
				"name": "Bob",
				"age":  30,
			},
		},
	}

	b.Run("Original", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			_, _ = mapGetWithSeparator(m, "users.[1].name", ".")
		}
	})

	b.Run("Optimized", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			_, _ = mapGetWithSeparatorOptimized(m, "users.[1].name", ".")
		}
	})
}

// Scenario 8: 大型 map（多键值对）
func BenchmarkMapGetWithSeparator_LargeMap(b *testing.B) {
	m := make(map[string]any, 1000)
	for i := 0; i < 1000; i++ {
		key := fmt.Sprintf("key_%d", i)
		m[key] = fmt.Sprintf("value_%d", i)
	}
	m["nested"] = map[string]any{
		"deep": map[string]any{
			"value": "found",
		},
	}

	b.Run("Original", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			_, _ = mapGetWithSeparator(m, "nested.deep.value", ".")
		}
	})

	b.Run("Optimized", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			_, _ = mapGetWithSeparatorOptimized(m, "nested.deep.value", ".")
		}
	})
}

// Scenario 9: 复杂键（多层嵌套 + 数组）
func BenchmarkMapGetWithSeparator_ComplexKey(b *testing.B) {
	m := map[string]any{
		"app": map[string]any{
			"services": []any{
				map[string]any{
					"name": "auth",
					"ports": []any{8080, 8081, 8082},
				},
				map[string]any{
					"name": "database",
					"ports": []any{5432, 5433},
				},
			},
		},
	}

	b.Run("Original", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			_, _ = mapGetWithSeparator(m, "app.services.[0].ports.[2]", ".")
		}
	})

	b.Run("Optimized", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			_, _ = mapGetWithSeparatorOptimized(m, "app.services.[0].ports.[2]", ".")
		}
	})
}

// Scenario 10: 不同分隔符（使用 / 而非 .）
func BenchmarkMapGetWithSeparator_DifferentSeparator(b *testing.B) {
	m := map[string]any{
		"path": map[string]any{
			"to": map[string]any{
				"resource": "file.txt",
			},
		},
	}

	b.Run("Original", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			_, _ = mapGetWithSeparator(m, "path/to/resource", "/")
		}
	})

	b.Run("Optimized", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			_, _ = mapGetWithSeparatorOptimized(m, "path/to/resource", "/")
		}
	})
}

// Scenario 11: 错误路径（键不存在）
func BenchmarkMapGetWithSeparator_KeyNotFound(b *testing.B) {
	m := map[string]any{
		"existing": "value",
	}

	b.Run("Original", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			_, _ = mapGetWithSeparator(m, "nonexistent", ".")
		}
	})

	b.Run("Optimized", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			_, _ = mapGetWithSeparatorOptimized(m, "nonexistent", ".")
		}
	})
}

// Scenario 12: 空键访问
func BenchmarkMapGetWithSeparator_EmptyKey(b *testing.B) {
	m := map[string]any{
		"key": "value",
	}

	b.Run("Original", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			_, _ = mapGetWithSeparator(m, "", ".")
		}
	})

	b.Run("Optimized", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			_, _ = mapGetWithSeparatorOptimized(m, "", ".")
		}
	})
}

// Scenario 13: 空map访问
func BenchmarkMapGetWithSeparator_EmptyMap(b *testing.B) {
	m := map[string]any{}

	b.Run("Original", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			_, _ = mapGetWithSeparator(m, "key", ".")
		}
	})

	b.Run("Optimized", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			_, _ = mapGetWithSeparatorOptimized(m, "key", ".")
		}
	})
}

// Scenario 14: 并发场景（使用 Parallel）
func BenchmarkMapGetWithSeparator_Concurrent(b *testing.B) {
	m := map[string]any{
		"a": map[string]any{
			"b": map[string]any{
				"c": "value",
			},
		},
	}

	b.Run("Original", func(b *testing.B) {
		b.ReportAllocs()
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				_, _ = mapGetWithSeparator(m, "a.b.c", ".")
			}
		})
	})

	b.Run("Optimized", func(b *testing.B) {
		b.ReportAllocs()
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				_, _ = mapGetWithSeparatorOptimized(m, "a.b.c", ".")
			}
		})
	})
}

// Scenario 15: 长键路径（10 层嵌套）
func BenchmarkMapGetWithSeparator_DeepNesting(b *testing.B) {
	// 构建深度嵌套结构
	m := make(map[string]any)
	current := m
	for i := 0; i < 10; i++ {
		next := make(map[string]any)
		current[fmt.Sprintf("level%d", i)] = next
		current = next
	}
	current["value"] = "deep"

	b.Run("Original", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			_, _ = mapGetWithSeparator(m, "level0.level1.level2.level3.level4.level5.level6.level7.level8.level9.value", ".")
		}
	})

	b.Run("Optimized", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			_, _ = mapGetWithSeparatorOptimized(m, "level0.level1.level2.level3.level4.level5.level6.level7.level8.level9.value", ".")
		}
	})
}

// Scenario 16: 字符串数组访问
func BenchmarkMapGetWithSeparator_StringArray(b *testing.B) {
	m := map[string]any{
		"tags": []string{"go", "performance", "benchmark"},
	}

	b.Run("Original", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			_, _ = mapGetWithSeparator(m, "tags.[1]", ".")
		}
	})

	b.Run("Optimized", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			_, _ = mapGetWithSeparatorOptimized(m, "tags.[1]", ".")
		}
	})
}

// Scenario 17: 多个索引访问（循环场景）
func BenchmarkMapGetWithSeparator_SequentialAccess(b *testing.B) {
	m := map[string]any{
		"items": []any{},
	}
	for i := 0; i < 100; i++ {
		m["items"] = append(m["items"].([]any), i)
	}

	b.Run("Original", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			idx := i % 100
			_, _ = mapGetWithSeparator(m, "items.["+strconv.Itoa(idx)+"]", ".")
		}
	})

	b.Run("Optimized", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			idx := i % 100
			_, _ = mapGetWithSeparatorOptimized(m, "items.["+strconv.Itoa(idx)+"]", ".")
		}
	})
}

// Scenario 18: 以分隔符结尾的键
func BenchmarkMapGetWithSeparator_TrailingSeparator(b *testing.B) {
	m := map[string]any{
		"a": map[string]any{
			"b": "value",
		},
	}

	b.Run("Original", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			_, _ = mapGetWithSeparator(m, "a.b.", ".")
		}
	})

	b.Run("Optimized", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			_, _ = mapGetWithSeparatorOptimized(m, "a.b.", ".")
		}
	})
}

// Scenario 19: 特殊字符键
func BenchmarkMapGetWithSeparator_SpecialChars(b *testing.B) {
	m := map[string]any{
		"key-with-dash": map[string]any{
			"[nested]": "value",
		},
	}

	b.Run("Original", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			_, _ = mapGetWithSeparator(m, "key-with-dash.[nested]", ".")
		}
	})

	b.Run("Optimized", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			_, _ = mapGetWithSeparatorOptimized(m, "key-with-dash.[nested]", ".")
		}
	})
}

// Scenario 20: 混合类型数组元素
func BenchmarkMapGetWithSeparator_MixedArrayTypes(b *testing.B) {
	m := map[string]any{
		"mixed": []any{
			"string",
			42,
			3.14,
			true,
			nil,
			map[string]any{"key": "value"},
		},
	}

	b.Run("Original", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			_, _ = mapGetWithSeparator(m, "mixed.[5]", ".")
		}
	})

	b.Run("Optimized", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			_, _ = mapGetWithSeparatorOptimized(m, "mixed.[5]", ".")
		}
	})
}
