package anyx

import (
	"strings"
	"testing"
)

// 方案1：当前实现 - 调用 mapGetWithSeparator
func BenchmarkMapGetMust_Current(b *testing.B) {
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
		"a": map[string]any{
			"b": map[string]any{
				"c": map[string]any{
					"d": map[string]any{
						"e": map[string]any{
							"f": "deep",
						},
					},
				},
			},
		},
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		MapGetMust(m, "name")
	}
}

// 方案2：直接调用 mapGetWithSeparatorOptimized（跳过错误检查）
func BenchmarkMapGetMust_OptimizedDirect(b *testing.B) {
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
		"a": map[string]any{
			"b": map[string]any{
				"c": map[string]any{
					"d": map[string]any{
						"e": map[string]any{
							"f": "deep",
						},
					},
				},
			},
		},
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		val, _ := mapGetWithSeparatorOptimized(m, "name", ".")
		_ = val
	}
}

// 方案3：内联快速路径（简单 key 不调用函数）
func BenchmarkMapGetMust_InlineFastPath(b *testing.B) {
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
		"a": map[string]any{
			"b": map[string]any{
				"c": map[string]any{
					"d": map[string]any{
						"e": map[string]any{
							"f": "deep",
						},
					},
				},
			},
		},
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		// 快速路径：简单 key
		if val, ok := m["name"]; ok {
			_ = val
			continue
		}
		// 复杂路径：调用优化版本
		val, _ := mapGetWithSeparatorOptimized(m, "name", ".")
		_ = val
	}
}

// 方案4：预检查分隔符（避免重复检查）
func BenchmarkMapGetMust_PreCheckSeparator(b *testing.B) {
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
		"a": map[string]any{
			"b": map[string]any{
				"c": map[string]any{
					"d": map[string]any{
						"e": map[string]any{
							"f": "deep",
						},
					},
				},
			},
		},
	}
	key := "name"
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		if strings.IndexByte(key, '.') == -1 && strings.IndexByte(key, '[') == -1 {
			if val, ok := m[key]; ok {
				_ = val
				continue
			}
		}
		val, _ := mapGetWithSeparatorOptimized(m, key, ".")
		_ = val
	}
}

// 方案5：嵌套 key - 2 层
func BenchmarkMapGetMust_Nested2Layers(b *testing.B) {
	m := map[string]any{
		"user": map[string]any{
			"profile": map[string]any{
				"name": "Jane",
				"age":  25,
			},
		},
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		MapGetMust(m, "user.profile.name")
	}
}

// 方案6：嵌套 key - 3 层
func BenchmarkMapGetMust_Nested3Layers(b *testing.B) {
	m := map[string]any{
		"level1": map[string]any{
			"level2": map[string]any{
				"level3": "value",
			},
		},
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		MapGetMust(m, "level1.level2.level3")
	}
}

// 方案7：深度嵌套 - 6 层
func BenchmarkMapGetMust_DeepNested6Layers(b *testing.B) {
	m := map[string]any{
		"a": map[string]any{
			"b": map[string]any{
				"c": map[string]any{
					"d": map[string]any{
						"e": map[string]any{
							"f": "deep",
						},
					},
				},
			},
		},
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		MapGetMust(m, "a.b.c.d.e.f")
	}
}

// 方案8：数组索引访问
func BenchmarkMapGetMust_ArrayIndex(b *testing.B) {
	m := map[string]any{
		"items": []any{"a", "b", "c", "d", "e"},
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		MapGetMust(m, "items[2]")
	}
}

// 方案9：混合场景 - 嵌套 + 数组
func BenchmarkMapGetMust_MixedNestedArray(b *testing.B) {
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

// 方案10：大型 map - 1000 条目
func BenchmarkMapGetMust_LargeMap1000(b *testing.B) {
	m := make(map[string]any, 1000)
	for i := 0; i < 1000; i++ {
		m[string(rune(i))] = i
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		// 使用存在的 key
		MapGetMust(m, string(rune(500)))
	}
}

// 方案11：并发访问 - 多 goroutine
func BenchmarkMapGetMust_Concurrent(b *testing.B) {
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
		"a": map[string]any{
			"b": map[string]any{
				"c": map[string]any{
					"d": map[string]any{
						"e": map[string]any{
							"f": "deep",
						},
					},
				},
			},
		},
	}
	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			MapGetMust(m, "name")
		}
	})
}

// 方案12：边界测试 - 不存在的 key (会panic，移除以避免日志过大)
// func BenchmarkMapGetMust_KeyNotFound(b *testing.B) {
// 	m := map[string]any{
// 		"name": "John",
// 	}
// 	b.ResetTimer()
// 	for i := 0; i < b.N; i++ {
// 		func() {
// 			defer func() {
// 				_ = recover()
// 			}()
// 			MapGetMust(m, "nonexistent")
// 		}()
// 	}
// }

// 方案13：边界测试 - 空值处理
func BenchmarkMapGetMust_NullValue(b *testing.B) {
	m := map[string]any{
		"value": nil,
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		MapGetMust(m, "value")
	}
}

// 方案14：零拷贝优化 - 直接返回切片引用
func BenchmarkMapGetMust_ZeroCopy(b *testing.B) {
	m := map[string]any{
		"data": []any{"a", "b", "c"},
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		val, _ := mapGetWithSeparatorOptimized(m, "data[0]", ".")
		_ = val
	}
}

// 方案15：缓存优化 - 预编译 key parts（模拟）
func BenchmarkMapGetMust_CachedKeyParts(b *testing.B) {
	m := map[string]any{
		"user": map[string]any{
			"profile": map[string]any{
				"name": "Jane",
			},
		},
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		// 模拟缓存 key parts 的场景
		val, _ := mapGetWithSeparatorOptimized(m, "user.profile.name", ".")
		_ = val
	}
}

// 方案16：字符串拼接优化（避免）
func BenchmarkMapGetMust_AvoidStringConcat(b *testing.B) {
	m := map[string]any{
		"name": "John",
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		// 直接访问，避免字符串拼接
		if val, ok := m["name"]; ok {
			_ = val
		}
	}
}

// 方案17：批量操作 - 多个 key
func BenchmarkMapGetMust_BatchKeys(b *testing.B) {
	m := map[string]any{
		"name":  "John",
		"age":   30,
		"email": "john@example.com",
		"user": map[string]any{
			"profile": map[string]any{
				"name": "Jane",
				"age":  25,
			},
		},
	}
	keys := []string{"name", "age", "email", "user.profile.name"}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for _, key := range keys {
			MapGetMust(m, key)
		}
	}
}

// 方案18：负数索引支持
func BenchmarkMapGetMust_NegativeIndex(b *testing.B) {
	m := map[string]any{
		"items": []any{"a", "b", "c", "d", "e"},
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		MapGetMust(m, "items[-1]")
	}
}

// 方案19：复杂嵌套 - 多层数组
func BenchmarkMapGetMust_ComplexMultiLevelArrays(b *testing.B) {
	m := map[string]any{
		"data": map[string]any{
			"items": []any{
				map[string]any{
					"value": "a",
				},
				map[string]any{
					"value": "b",
				},
			},
		},
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		MapGetMust(m, "data.items[1].value")
	}
}

// 方案20：错误路径优化（快速失败）(会panic，移除以避免日志过大)
// func BenchmarkMapGetMust_FastFail(b *testing.B) {
// 	m := map[string]any{}
// 	b.ResetTimer()
// 	for i := 0; i < b.N; i++ {
// 		func() {
// 			defer func() {
// 				_ = recover()
// 			}()
// 			MapGetMust(m, "deep.nested.path")
// 		}()
// 	}
// }

// 方案21：类型断言优化 - 直接类型匹配
func BenchmarkMapGetMust_DirectTypeAssertion(b *testing.B) {
	m := map[string]any{
		"name": "John",
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		if val, ok := m["name"]; ok {
			if str, ok := val.(string); ok {
				_ = str
			}
		}
	}
}

// 方案22：内存池优化（模拟）
func BenchmarkMapGetMust_MemoryPool(b *testing.B) {
	m := map[string]any{
		"name": "John",
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		val, _ := mapGetWithSeparatorOptimized(m, "name", ".")
		_ = val
	}
}

// 方案23：分支预测优化
func BenchmarkMapGetMust_BranchPrediction(b *testing.B) {
	m := map[string]any{
		"name":  "John",
		"age":   30,
		"email": "john@example.com",
	}
	keys := []string{"name", "age", "email"}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		// 循环访问不同的 key，帮助分支预测
		key := keys[i%3]
		MapGetMust(m, key)
	}
}

// 方案24：内联所有逻辑（零函数调用）
func BenchmarkMapGetMust_FullyInlined(b *testing.B) {
	m := map[string]any{
		"name": "John",
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		// 完全内联的简单 key 访问
		if val, ok := m["name"]; ok {
			_ = val
		} else {
			panic("key not found")
		}
	}
}

// 方案25：使用 defer recover 的开销
func BenchmarkMapGetMust_DeferRecoverOverhead(b *testing.B) {
	m := map[string]any{
		"name": "John",
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		func() {
			defer func() {
				_ = recover()
			}()
			if val, ok := m["name"]; ok {
				_ = val
			}
		}()
	}
}

// 方案26：自定义 panic 处理
func BenchmarkMapGetMust_CustomPanic(b *testing.B) {
	m := map[string]any{
		"name": "John",
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		val, err := mapGetWithSeparator(m, "name", ".")
		if err != nil {
			panic(err)
		}
		_ = val
	}
}

// 方案27：map[any]any 类型
func BenchmarkMapGetMust_MapAnyAny(b *testing.B) {
	m := map[string]any{
		"data": map[any]any{
			"key": "value",
		},
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		MapGetMust(m, "data.key")
	}
}

// 方案28：slice 边界检查优化
func BenchmarkMapGetMust_SliceBoundsCheck(b *testing.B) {
	m := map[string]any{
		"items": make([]any, 100),
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		MapGetMust(m, "items[50]")
	}
}

// 方案29：字符串比较优化
func BenchmarkMapGetMust_StringCompareOptimization(b *testing.B) {
	m := map[string]any{
		"very_long_key_name_that_should_be_optimized": "value",
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		MapGetMust(m, "very_long_key_name_that_should_be_optimized")
	}
}

// 方案30：热路径优化（重复访问相同 key）
func BenchmarkMapGetMust_HotPath(b *testing.B) {
	m := map[string]any{
		"name": "John",
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		// 重复访问相同的 key，模拟热路径
		for j := 0; j < 10; j++ {
			MapGetMust(m, "name")
		}
	}
}
