package anyx

import (
	"fmt"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestMapGetWithSeparatorCompare 验证两个版本行为一致
func TestMapGetWithSeparatorCompare(t *testing.T) {
	testCases := []struct {
		name string
		m    map[string]any
		key  string
		sep  string
	}{
		{
			name: "简单键",
			m: map[string]any{
				"name": "John",
			},
			key: "name",
			sep:  ".",
		},
		{
			name: "两层嵌套",
			m: map[string]any{
				"user": map[string]any{
					"name": "Alice",
				},
			},
			key: "user.name",
			sep:  ".",
		},
		{
			name: "数组索引",
			m: map[string]any{
				"items": []any{"a", "b", "c"},
			},
			key: "items.[1]",
			sep:  ".",
		},
		{
			name: "负数索引",
			m: map[string]any{
				"items": []any{"a", "b", "c"},
			},
			key: "items.[-1]",
			sep:  ".",
		},
		{
			name: "混合map数组",
			m: map[string]any{
				"users": []any{
					map[string]any{"name": "Alice"},
					map[string]any{"name": "Bob"},
				},
			},
			key: "users.[1].name",
			sep:  ".",
		},
		{
			name: "深层嵌套",
			m: map[string]any{
				"a": map[string]any{
					"b": map[string]any{
						"c": map[string]any{
							"d": "value",
						},
					},
				},
			},
			key: "a.b.c.d",
			sep:  ".",
		},
		{
			name: "错误-键不存在",
			m: map[string]any{
				"existing": "value",
			},
			key: "nonexistent",
			sep:  ".",
		},
		{
			name: "错误-空键",
			m: map[string]any{
				"key": "value",
			},
			key: "",
			sep:  ".",
		},
		{
			name: "错误-空map",
			m:    map[string]any{},
			key:  "key",
			sep:  ".",
		},
		{
			name: "错误-索引越界",
			m: map[string]any{
				"items": []any{1, 2},
			},
			key: "items.[10]",
			sep:  ".",
		},
		{
			name: "错误-无效索引",
			m: map[string]any{
				"items": []any{1, 2},
			},
			key: "items.[abc]",
			sep:  ".",
		},
		{
			name: "字符串数组",
			m: map[string]any{
				"tags": []string{"go", "test"},
			},
			key: "tags.[0]",
			sep:  ".",
		},
		{
			name: "以分隔符结尾",
			m: map[string]any{
				"a": map[string]any{
					"b": "value",
				},
			},
			key: "a.b.",
			sep:  ".",
		},
		{
			name: "不同分隔符",
			m: map[string]any{
				"path": map[string]any{
					"to": map[string]any{
						"file": "data.txt",
					},
				},
			},
			key: "path/to/file",
			sep:  "/",
		},
		{
			name: "int数组",
			m: map[string]any{
				"numbers": []int{1, 2, 3},
			},
			key: "numbers.[1]",
			sep:  ".",
		},
		{
			name: "nil值",
			m: map[string]any{
				"nullable": nil,
			},
			key: "nullable",
			sep:  ".",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			gotOrig, errOrig := mapGetWithSeparator(tc.m, tc.key, tc.sep)
			gotOpt, errOpt := mapGetWithSeparatorOptimized(tc.m, tc.key, tc.sep)

			// 错误状态应该一致
			assert.Equal(t, errOrig == nil, errOpt == nil,
				"错误状态不一致: original err=%v, optimized err=%v", errOrig, errOpt)

			// 如果都成功，结果应该相同
			if errOrig == nil && errOpt == nil {
				assert.Equal(t, gotOrig, gotOpt,
					"返回值不一致: original=%#v, optimized=%#v", gotOrig, gotOpt)
			}
		})
	}
}

// Benchmark scenarios
func BenchmarkSimpleKey(b *testing.B) {
	m := map[string]any{"name": "John"}

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

func BenchmarkNested2Levels(b *testing.B) {
	m := map[string]any{
		"user": map[string]any{
			"name": "Alice",
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

func BenchmarkNested5Levels(b *testing.B) {
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

func BenchmarkArrayIndex(b *testing.B) {
	m := map[string]any{
		"items": []any{"a", "b", "c"},
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

func BenchmarkNegativeIndex(b *testing.B) {
	m := map[string]any{
		"items": []any{"a", "b", "c"},
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

func BenchmarkMixedMapArray(b *testing.B) {
	m := map[string]any{
		"users": []any{
			map[string]any{"name": "Alice"},
			map[string]any{"name": "Bob"},
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

func BenchmarkLargeMap(b *testing.B) {
	m := make(map[string]any, 1000)
	for i := 0; i < 1000; i++ {
		m[fmt.Sprintf("key_%d", i)] = fmt.Sprintf("value_%d", i)
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

func BenchmarkComplexKey(b *testing.B) {
	m := map[string]any{
		"app": map[string]any{
			"services": []any{
				map[string]any{
					"name":  "auth",
					"ports": []any{8080, 8081, 8082},
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

func BenchmarkDifferentSeparator(b *testing.B) {
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

func BenchmarkKeyNotFound(b *testing.B) {
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

func BenchmarkEmptyKey(b *testing.B) {
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

func BenchmarkEmptyMap(b *testing.B) {
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

func BenchmarkDeepNesting(b *testing.B) {
	m := map[string]any{}
	current := m
	for i := 0; i < 10; i++ {
		next := map[string]any{}
		current[fmt.Sprintf("l%d", i)] = next
		current = next
	}
	current["value"] = "deep"
	key := "l0.l1.l2.l3.l4.l5.l6.l7.l8.l9.value"

	b.Run("Original", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			_, _ = mapGetWithSeparator(m, key, ".")
		}
	})

	b.Run("Optimized", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			_, _ = mapGetWithSeparatorOptimized(m, key, ".")
		}
	})
}

func BenchmarkStringArray(b *testing.B) {
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

func BenchmarkTrailingSeparator(b *testing.B) {
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

func BenchmarkSequentialAccess(b *testing.B) {
	m := map[string]any{
		"items": make([]any, 100),
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

func BenchmarkSpecialChars(b *testing.B) {
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

func BenchmarkNestedArrays(b *testing.B) {
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
