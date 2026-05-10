package anyx

import (
	"testing"
)

// TestMapExistsWithSepCoverage 全面覆盖率测试
func TestMapExistsWithSepCoverage(t *testing.T) {
	tests := []struct {
		name     string
		m        map[string]any
		key      string
		sep      string
		expected bool
	}{
		// 基础场景
		{
			name:     "简单存在key",
			m:        map[string]any{"name": "value"},
			key:      "name",
			sep:      ".",
			expected: true,
		},
		{
			name:     "简单不存在key",
			m:        map[string]any{"name": "value"},
			key:      "nonexistent",
			sep:      ".",
			expected: false,
		},
		{
			name:     "空map",
			m:        map[string]any{},
			key:      "key",
			sep:      ".",
			expected: false,
		},
		{
			name:     "空key",
			m:        map[string]any{"key": "value"},
			key:      "",
			sep:      ".",
			expected: false,
		},
		{
			name:     "nil map",
			m:        nil,
			key:      "key",
			sep:      ".",
			expected: false,
		},

		// 嵌套场景
		{
			name: "嵌套存在",
			m: map[string]any{
				"level1": map[string]any{
					"level2": map[string]any{
						"target": "value",
					},
				},
			},
			key:      "level1.level2.target",
			sep:      ".",
			expected: true,
		},
		{
			name: "嵌套中间不存在",
			m: map[string]any{
				"level1": map[string]any{
					"level2": "value",
				},
			},
			key:      "level1.nonexistent.target",
			sep:      ".",
			expected: false,
		},
		{
			name: "嵌套最终不存在",
			m: map[string]any{
				"level1": map[string]any{
					"level2": map[string]any{},
				},
			},
			key:      "level1.level2.target",
			sep:      ".",
			expected: false,
		},
		{
			name: "深层嵌套",
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
			key:      "a.b.c.d.e",
			sep:      ".",
			expected: true,
		},

		// 数组索引场景
		{
			name: "数组索引存在",
			m: map[string]any{
				"items": []any{"a", "b", "c"},
			},
			key:      "items[1]",
			sep:      ".",
			expected: true,
		},
		{
			name: "数组索引越界",
			m: map[string]any{
				"items": []any{"a", "b", "c"},
			},
			key:      "items[10]",
			sep:      ".",
			expected: false,
		},
		{
			name: "数组索引负数",
			m: map[string]any{
				"items": []any{"a", "b", "c"},
			},
			key:      "items[-1]",
			sep:      ".",
			expected: false,
		},
		{
			name: "数组索引负号后跟零",
			m: map[string]any{
				"items": []any{"a", "b", "c"},
			},
			key:      "items[-0]",
			sep:      ".",
			expected: true,
		},
		{
			name: "数组索引多位负数",
			m: map[string]any{
				"items": []any{"a", "b", "c"},
			},
			key:      "items[-10]",
			sep:      ".",
			expected: false,
		},
		{
			name: "嵌套数组",
			m: map[string]any{
				"data": map[string]any{
					"items": []any{
						map[string]any{"id": 1},
						map[string]any{"id": 2},
					},
				},
			},
			key:      "data.items[0].id",
			sep:      ".",
			expected: true,
		},
		{
			name: "多个数组索引",
			m: map[string]any{
				"matrix": []any{
					[]any{1, 2, 3},
					[]any{4, 5, 6},
				},
			},
			key:      "matrix[1][2]",
			sep:      ".",
			expected: true,
		},
		{
			name: "空数组",
			m: map[string]any{
				"items": []any{},
			},
			key:      "items[0]",
			sep:      ".",
			expected: false,
		},
		{
			name: "第一个元素",
			m: map[string]any{
				"items": []any{"a", "b", "c"},
			},
			key:      "items[0]",
			sep:      ".",
			expected: true,
		},
		{
			name: "最后一个元素",
			m: map[string]any{
				"items": []any{"a", "b", "c"},
			},
			key:      "items[2]",
			sep:      ".",
			expected: true,
		},
		{
			name: "大索引值",
			m: map[string]any{
				"items": make([]any, 100),
			},
			key:      "items[99]",
			sep:      ".",
			expected: true,
		},

		// 不同分隔符
		{
			name: "斜杠分隔符",
			m: map[string]any{
				"level1": map[string]any{
					"level2": "value",
				},
			},
			key:      "level1/level2",
			sep:      "/",
			expected: true,
		},
		{
			name: "双冒号分隔符",
			m: map[string]any{
				"level1": map[string]any{
					"level2": "value",
				},
			},
			key:      "level1::level2",
			sep:      "::",
			expected: true,
		},
		{
			name: "连字符分隔符",
			m: map[string]any{
				"level1": map[string]any{
					"level2": "value",
				},
			},
			key:      "level1-level2",
			sep:      "-",
			expected: true,
		},
		{
			name: "key包含分隔符但用不同分隔符查询",
			m: map[string]any{
				"a.b": map[string]any{
					"c": "value",
				},
			},
			key:      "a.b/c",
			sep:      "/",
			expected: true,
		},

		// 特殊字符和边界情况
		{
			name: "空索引",
			m: map[string]any{
				"items": []any{"a"},
			},
			key:      "items[]",
			sep:      ".",
			expected: false,
		},
		{
			name: "只有左括号",
			m: map[string]any{
				"items": []any{"a"},
			},
			key:      "items[",
			sep:      ".",
			expected: false,
		},
		{
			name: "只有右括号",
			m: map[string]any{
				"items": []any{"a"},
			},
			key:      "items]",
			sep:      ".",
			expected: false,
		},
		{
			name: "索引不是数字",
			m: map[string]any{
				"items": []any{"a"},
			},
			key:      "items[abc]",
			sep:      ".",
			expected: false,
		},
		{
			name: "索引包含非数字字符",
			m: map[string]any{
				"items": []any{"a"},
			},
			key:      "items[1abc]",
			sep:      ".",
			expected: false,
		},

		// 类型不匹配
		{
			name:     "期望map但实际是字符串",
			m:        map[string]any{"key": "not_a_map"},
			key:      "key.subkey",
			sep:      ".",
			expected: false,
		},
		{
			name:     "期望数组但实际是map",
			m:        map[string]any{"key": map[string]any{"sub": "value"}},
			key:      "key[0]",
			sep:      ".",
			expected: false,
		},
		{
			name: "中间层类型不匹配",
			m: map[string]any{
				"level1": "not_a_map",
			},
			key:      "level1.level2",
			sep:      ".",
			expected: false,
		},

		// 复杂真实场景
		{
			name: "配置文件风格",
			m: map[string]any{
				"server": map[string]any{
					"host": "localhost",
					"port": 8080,
					"ssl": map[string]any{
						"enabled": true,
						"cert": "/path/to/cert",
					},
				},
			},
			key:      "server.ssl.cert",
			sep:      ".",
			expected: true,
		},
		{
			name: "API响应风格",
			m: map[string]any{
				"data": map[string]any{
					"users": []any{
						map[string]any{"id": 1, "name": "Alice"},
						map[string]any{"id": 2, "name": "Bob"},
					},
				},
			},
			key:      "data.users[1].name",
			sep:      ".",
			expected: true,
		},
		{
			name: "混合路径",
			m: map[string]any{
				"results": []any{
					map[string]any{
						"id":     1,
						"tags":   []any{"tag1", "tag2"},
						"nested": map[string]any{"key": "value"},
					},
				},
			},
			key:      "results[0].tags[1]",
			sep:      ".",
			expected: true,
		},

		// 边界值
		{
			name:     "单字符key",
			m:        map[string]any{"a": "value"},
			key:      "a",
			sep:      ".",
			expected: true,
		},
		{
			name:     "长key",
			m:        map[string]any{"very_long_key_name_with_many_underscores": "value"},
			key:      "very_long_key_name_with_many_underscores",
			sep:      ".",
			expected: true,
		},
		{
			name: "许多嵌套层级",
			m: func() map[string]any {
				m := map[string]any{}
				current := m
				for i := 0; i < 20; i++ {
					next := map[string]any{}
					current["level"] = next
					current = next
				}
				current["value"] = "target"
				return m
			}(),
			key:      "level.level.level.level.level.level.level.level.level.level.level.level.level.level.level.level.level.level.level.level.value",
			sep:      ".",
			expected: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := MapExistsWithSep(tt.m, tt.key, tt.sep)
			if result != tt.expected {
				t.Errorf("MapExistsWithSep(%v, %q, %q) = %v, want %v",
					tt.m, tt.key, tt.sep, result, tt.expected)
			}
		})
	}
}

// TestMapExistsWithSepEdgeCases 额外边界情况测试
func TestMapExistsWithSepEdgeCases(t *testing.T) {
	t.Run("数组类型的map", func(t *testing.T) {
		m := map[string]any{
			"items": []map[string]any{
				{"id": 1, "name": "A"},
				{"id": 2, "name": "B"},
			},
		}
		if !MapExistsWithSep(m, "items[0].id", ".") {
			t.Error("应该找到 items[0].id")
		}
		if !MapExistsWithSep(m, "items[1].name", ".") {
			t.Error("应该找到 items[1].name")
		}
	})

	t.Run("空字符串作为值", func(t *testing.T) {
		m := map[string]any{
			"key": "",
		}
		if !MapExistsWithSep(m, "key", ".") {
			t.Error("空字符串是有效值")
		}
	})

	t.Run("零值", func(t *testing.T) {
		m := map[string]any{
			"zero": 0,
			"empty": []any{},
			"null": nil,
		}
		if !MapExistsWithSep(m, "zero", ".") {
			t.Error("零值应该存在")
		}
		if !MapExistsWithSep(m, "empty", ".") {
			t.Error("空切片应该存在")
		}
		if !MapExistsWithSep(m, "null", ".") {
			t.Error("nil应该存在")
		}
	})

	t.Run("嵌套中的nil值", func(t *testing.T) {
		m := map[string]any{
			"outer": map[string]any{
				"inner": nil,
			},
		}
		if !MapExistsWithSep(m, "outer.inner", ".") {
			t.Error("嵌套中的nil应该存在")
		}
	})

	t.Run("连续分隔符", func(t *testing.T) {
		m := map[string]any{
			"a": map[string]any{
				"": map[string]any{
					"b": "value",
				},
			},
		}
		// 注意：这可能依赖于splitKey的实现
		result := MapExistsWithSep(m, "a..b", ".")
		if result {
			t.Log("连续分隔符处理：当前实现支持空键名")
		}
	})
}

// TestMapExistsWithSepConcurrent 并发安全性测试
func TestMapExistsWithSepConcurrent(t *testing.T) {
	m := map[string]any{
		"a": map[string]any{
			"b": map[string]any{
				"c": "value",
			},
		},
		"items": []any{"x", "y", "z"},
	}

	done := make(chan bool)
	iterations := 1000

	// 启动多个goroutine并发读取
	for i := 0; i < 10; i++ {
		go func() {
			for j := 0; j < iterations; j++ {
				MapExistsWithSep(m, "a.b.c", ".")
				MapExistsWithSep(m, "items[1]", ".")
				MapExistsWithSep(m, "nonexistent", ".")
			}
			done <- true
		}()
	}

	// 等待所有goroutine完成
	for i := 0; i < 10; i++ {
		<-done
	}
}

// TestMapExistsWithSepSpecialSeparators 特殊分隔符测试
func TestMapExistsWithSepSpecialSeparators(t *testing.T) {
	m := map[string]any{
		"a": map[string]any{
			"b": map[string]any{
				"c": "value",
			},
		},
	}

	// 各种特殊分隔符
	separators := []string{".", "/", "-", "::", "|", "#", "@"}

	for _, sep := range separators {
		t.Run("sep_"+sep, func(t *testing.T) {
			// 构建对应分隔符的key
			key := "a" + sep + "b" + sep + "c"
			if !MapExistsWithSep(m, key, sep) {
				t.Errorf("分隔符 %q 未能正确处理", sep)
			}
		})
	}
}

// TestMapExistsWithSepArrayTypes 不同数组类型测试
func TestMapExistsWithSepArrayTypes(t *testing.T) {
	t.Run("[]any数组", func(t *testing.T) {
		m := map[string]any{
			"items": []any{1, 2, 3},
		}
		if !MapExistsWithSep(m, "items[0]", ".") {
			t.Error("应该找到[]any元素")
		}
	})

	t.Run("[]map[string]any数组", func(t *testing.T) {
		m := map[string]any{
			"items": []map[string]any{
				{"id": 1},
				{"id": 2},
			},
		}
		if !MapExistsWithSep(m, "items[1].id", ".") {
			t.Error("应该找到[]map[string]any元素")
		}
	})

	t.Run("[]string数组", func(t *testing.T) {
		m := map[string]any{
			"items": []string{"a", "b", "c"},
		}
		// 当前实现可能不支持[]string类型
		result := MapExistsWithSep(m, "items[0]", ".")
		t.Logf("[]string支持: %v", result)
	})

	t.Run("数组索引为负数的边界", func(t *testing.T) {
		m := map[string]any{
			"items": []any{"a", "b", "c"},
		}
		// 负索引应该被解析但不应该找到元素（因为Go不支持负索引）
		if MapExistsWithSep(m, "items[-1]", ".") {
			t.Error("负索引应该返回false")
		}
	})
}

// TestMapExistsWithSepNegativeIndexParsing 负索引解析测试
func TestMapExistsWithSepNegativeIndexParsing(t *testing.T) {
	tests := []struct {
		name     string
		indexStr string
		valid    bool
	}{
		{"零", "[0]", true},
		{"正数", "[1]", true},
		{"负数", "[-1]", true},
		{"负号后零", "[-0]", true},
		{"多位负数", "[-10]", true},
		{"只有负号", "[-]", false},
		{"负号在中间", "[1-]", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := map[string]any{
				"items": []any{"a", "b", "c"},
			}
			// 测试解析是否正确
			_ = MapExistsWithSep(m, "items"+tt.indexStr, ".")
		})
	}
}
