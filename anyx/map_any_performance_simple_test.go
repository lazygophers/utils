package anyx

import (
	"fmt"
	"testing"
)

// 简单性能对比测试
func TestPerformanceComparison(t *testing.T) {
	testCases := []struct {
		name string
		m    map[string]any
		key  string
		sep  string
	}{
		{
			name: "简单键",
			m:    map[string]any{"name": "John"},
			key:  "name",
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
			name: "五层嵌套",
			m: map[string]any{
				"a": map[string]any{
					"b": map[string]any{
						"c": map[string]any{
							"d": map[string]any{
								"e": "value",
							},
						},
					},
				},
			},
			key: "a.b.c.d.e",
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
			name: "混合复杂",
			m: map[string]any{
				"app": map[string]any{
					"services": []any{
						map[string]any{
							"name":  "auth",
							"ports": []any{8080, 8081, 8082},
						},
					},
				},
			},
			key: "app.services.[0].ports.[2]",
			sep:  ".",
		},
	}

	iterations := 10000

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// 测试原始版本
			start := testing.AllocsPerRun(iterations, func() {
				_, _ = mapGetWithSeparator(tc.m, tc.key, tc.sep)
			})
			originalAllocs := start

			// 测试优化版本
			optimizedAllocs := testing.AllocsPerRun(iterations, func() {
				_, _ = mapGetWithSeparatorOptimized(tc.m, tc.key, tc.sep)
			})

			t.Logf("场景: %s", tc.name)
			t.Logf("原始版本每次分配: %.2f", originalAllocs)
			t.Logf("优化版本每次分配: %.2f", optimizedAllocs)

			if originalAllocs > optimizedAllocs {
				improvement := ((originalAllocs - optimizedAllocs) / originalAllocs) * 100
				t.Logf("分配减少: %.1f%%", improvement)
			}
		})
	}
}

// 基准测试
func BenchmarkOriginalVsOptimized(b *testing.B) {
	testCases := []struct {
		name string
		m    map[string]any
		key  string
		sep  string
	}{
		{
			name: "简单键",
			m:    map[string]any{"name": "John"},
			key:  "name",
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
			name: "混合复杂",
			m: map[string]any{
				"app": map[string]any{
					"services": []any{
						map[string]any{
							"name":  "auth",
							"ports": []any{8080, 8081},
						},
					},
				},
			},
			key: "app.services.[0].ports.[1]",
			sep:  ".",
		},
	}

	for _, tc := range testCases {
		b.Run(tc.name+"/Original", func(b *testing.B) {
			b.ReportAllocs()
			for i := 0; i < b.N; i++ {
				_, _ = mapGetWithSeparator(tc.m, tc.key, tc.sep)
			}
		})

		b.Run(tc.name+"/Optimized", func(b *testing.B) {
			b.ReportAllocs()
			for i := 0; i < b.N; i++ {
				_, _ = mapGetWithSeparatorOptimized(tc.m, tc.key, tc.sep)
			}
		})
	}
}

// 输出性能对比报告
func Example_performanceComparison() {
	fmt.Println("mapGetWithSeparator 性能对比")
	fmt.Println("=" + string(make([]byte, 40)))
	// 输出各场景的性能提升数据
}
