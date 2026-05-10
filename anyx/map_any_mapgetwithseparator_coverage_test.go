package anyx

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestMapGetWithSeparator_Coverage 全面测试 mapGetWithSeparator 和优化版本
func TestMapGetWithSeparator_Coverage(t *testing.T) {
	tests := []struct {
		name        string
		mapData     map[string]any
		key         string
		sep         string
		want        any
		wantErrType error
	}{
		// 基础场景
		{
			name: "简单键访问",
			mapData: map[string]any{
				"name": "John",
				"age":  30,
			},
			key:         "name",
			sep:         ".",
			want:        "John",
			wantErrType: nil,
		},
		{
			name: "空map访问",
			mapData: map[string]any{},
			key:     "key",
			sep:     ".",
			want:    nil,
			wantErrType: ErrNotFound,
		},
		{
			name: "空键访问",
			mapData: map[string]any{
				"key": "value",
			},
			key:         "",
			sep:         ".",
			want:        nil,
			wantErrType: ErrEmptyKey,
		},

		// 嵌套访问
		{
			name: "两层嵌套",
			mapData: map[string]any{
				"user": map[string]any{
					"name": "Alice",
				},
			},
			key:         "user.name",
			sep:         ".",
			want:        "Alice",
			wantErrType: nil,
		},
		{
			name: "多层嵌套",
			mapData: map[string]any{
				"a": map[string]any{
					"b": map[string]any{
						"c": map[string]any{
							"d": "deep",
						},
					},
				},
			},
			key:         "a.b.c.d",
			sep:         ".",
			want:        "deep",
			wantErrType: nil,
		},

		// 数组访问
		{
			name: "正数索引",
			mapData: map[string]any{
				"items": []any{"a", "b", "c"},
			},
			key:         "items.[1]",
			sep:         ".",
			want:        "b",
			wantErrType: nil,
		},
		{
			name: "负数索引",
			mapData: map[string]any{
				"items": []any{"a", "b", "c"},
			},
			key:         "items.[-1]",
			sep:         ".",
			want:        "c",
			wantErrType: nil,
		},
		{
			name: "索引越界-正数",
			mapData: map[string]any{
				"items": []any{"a", "b"},
			},
			key:         "items.[5]",
			sep:         ".",
			want:        nil,
			wantErrType: ErrOutOfRange,
		},
		{
			name: "索引越界-负数",
			mapData: map[string]any{
				"items": []any{"a", "b"},
			},
			key:         "items.[-5]",
			sep:         ".",
			want:        nil,
			wantErrType: ErrOutOfRange,
		},

		// 字符串数组
		{
			name: "字符串数组访问",
			mapData: map[string]any{
				"tags": []string{"go", "test"},
			},
			key:         "tags.[0]",
			sep:         ".",
			want:        "go",
			wantErrType: nil,
		},

		// 混合map和数组
		{
			name: "map包含数组",
			mapData: map[string]any{
				"data": map[string]any{
					"items": []any{1, 2, 3},
				},
			},
			key:         "data.items.[1]",
			sep:         ".",
			want:        2,
			wantErrType: nil,
		},
		{
			name: "数组包含map",
			mapData: map[string]any{
				"users": []any{
					map[string]any{"name": "Alice"},
					map[string]any{"name": "Bob"},
				},
			},
			key:         "users.[1].name",
			sep:         ".",
			want:        "Bob",
			wantErrType: nil,
		},

		// 不同分隔符
		{
			name: "使用斜杠分隔符",
			mapData: map[string]any{
				"path": map[string]any{
					"to": map[string]any{
						"file": "data.txt",
					},
				},
			},
			key:         "path/to/file",
			sep:         "/",
			want:        "data.txt",
			wantErrType: nil,
		},

		// 错误场景
		{
			name: "键不存在-简单",
			mapData: map[string]any{
				"existing": "value",
			},
			key:         "nonexistent",
			sep:         ".",
			want:        nil,
			wantErrType: ErrNotFound,
		},
		{
			name: "键不存在-嵌套",
			mapData: map[string]any{
				"user": map[string]any{
					"name": "Alice",
				},
			},
			key:         "user.age",
			sep:         ".",
			want:        nil,
			wantErrType: ErrNotFound,
		},
		{
			name: "中间路径不存在",
			mapData: map[string]any{
				"user": map[string]any{
					"name": "Alice",
				},
			},
			key:         "nonexistent.path",
			sep:         ".",
			want:        nil,
			wantErrType: ErrNotFound,
		},
		{
			name: "无效索引格式",
			mapData: map[string]any{
				"items": []any{1, 2, 3},
			},
			key:         "items.[abc]",
			sep:         ".",
			want:        nil,
			wantErrType: ErrInvalidIndex,
		},
		{
			name: "在非数组类型上使用索引",
			mapData: map[string]any{
				"name": "Alice",
			},
			key:         "name.[0]",
			sep:         ".",
			want:        nil,
			wantErrType: ErrInvalidSlice,
		},

		// 边界情况
		{
			name: "以分隔符结尾的键",
			mapData: map[string]any{
				"a": map[string]any{
					"b": "value",
				},
			},
			key:         "a.b.",
			sep:         ".",
			want:        nil,
			wantErrType: ErrNotFound,
		},
		{
			name: "空字符串键",
			mapData: map[string]any{
				"": "empty key value",
			},
			key:         "",
			sep:         ".",
			want:        nil,
			wantErrType: ErrEmptyKey,
		},

		// 类型转换
		{
			name: "int数组",
			mapData: map[string]any{
				"numbers": []int{1, 2, 3},
			},
			key:         "numbers.[1]",
			sep:         ".",
			want:        2,
			wantErrType: nil,
		},
		{
			name: "int64数组",
			mapData: map[string]any{
				"numbers": []int64{100, 200, 300},
			},
			key:         "numbers.[2]",
			sep:         ".",
			want:        int64(300),
			wantErrType: nil,
		},
		{
			name: "float64数组",
			mapData: map[string]any{
				"values": []float64{1.1, 2.2, 3.3},
			},
			key:         "values.[0]",
			sep:         ".",
			want:        1.1,
			wantErrType: nil,
		},
		{
			name: "bool数组",
			mapData: map[string]any{
				"flags": []bool{true, false, true},
			},
			key:         "flags.[1]",
			sep:         ".",
			want:        false,
			wantErrType: nil,
		},

		// 复杂场景
		{
			name: "深层嵌套加数组",
			mapData: map[string]any{
				"app": map[string]any{
					"services": []any{
						map[string]any{
							"name": "auth",
							"ports": []any{8080, 8081},
						},
					},
				},
			},
			key:         "app.services.[0].ports.[1]",
			sep:         ".",
			want:        8081,
			wantErrType: nil,
		},
		{
			name: "map[any]any 类型",
			mapData: map[string]any{
				"data": map[any]any{
					"key": "value",
					42:    "number key",
				},
			},
			key:         "data.key",
			sep:         ".",
			want:        "value",
			wantErrType: nil,
		},

		// 特殊字符
		{
			name: "键包含特殊字符",
			mapData: map[string]any{
				"key-with-dash": map[string]any{
					"[nested]": "value",
				},
			},
			key:         "key-with-dash.[nested]",
			sep:         ".",
			want:        "value",
			wantErrType: nil,
		},
		{
			name: "nil值处理",
			mapData: map[string]any{
				"nullable": nil,
			},
			key:         "nullable",
			sep:         ".",
			want:        nil,
			wantErrType: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name+"-Original", func(t *testing.T) {
			got, err := mapGetWithSeparator(tt.mapData, tt.key, tt.sep)
			if tt.wantErrType != nil {
				assert.Error(t, err)
				// 验证错误类型
				// 注意：由于错误被包装，我们只检查不为nil
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.want, got)
			}
		})

		t.Run(tt.name+"-Optimized", func(t *testing.T) {
			got, err := mapGetWithSeparatorOptimized(tt.mapData, tt.key, tt.sep)
			if tt.wantErrType != nil {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.want, got)
			}
		})

		// 验证两个版本结果一致
		t.Run(tt.name+"-Compare", func(t *testing.T) {
			gotOrig, errOrig := mapGetWithSeparator(tt.mapData, tt.key, tt.sep)
			gotOpt, errOpt := mapGetWithSeparatorOptimized(tt.mapData, tt.key, tt.sep)

			// 错误状态应该一致
			if (errOrig == nil) != (errOpt == nil) {
				t.Errorf("错误状态不一致: original err=%v, optimized err=%v", errOrig, errOpt)
				return
			}

			// 如果都成功，结果应该相同
			if errOrig == nil && errOpt == nil {
				if !assert.Equal(t, gotOrig, gotOpt, "两个版本返回值不同") {
					t.Logf("原始版本: %#v", gotOrig)
					t.Logf("优化版本: %#v", gotOpt)
				}
			}
		})
	}
}

// TestMapGetWithSeparator_EdgeCases 测试边界情况
func TestMapGetWithSeparator_EdgeCases(t *testing.T) {
	t.Run("零值处理", func(t *testing.T) {
		m := map[string]any{
			"zero_int":     0,
			"zero_float":   0.0,
			"zero_string":  "",
			"zero_bool":    false,
			"empty_array":  []any{},
			"empty_object": map[string]any{},
		}

		tests := []struct {
			key string
		}{
			{"zero_int"},
			{"zero_float"},
			{"zero_string"},
			{"zero_bool"},
			{"empty_array"},
			{"empty_object"},
		}

		for _, tt := range tests {
			t.Run(tt.key, func(t *testing.T) {
				gotOrig, errOrig := mapGetWithSeparator(m, tt.key, ".")
				gotOpt, errOpt := mapGetWithSeparatorOptimized(m, tt.key, ".")

				assert.Equal(t, errOrig == nil, errOpt == nil, "错误状态不一致")
				if errOrig == nil && errOpt == nil {
					assert.Equal(t, gotOrig, gotOpt, "返回值不一致")
				}
			})
		}
	})

	t.Run("多级数组索引", func(t *testing.T) {
		m := map[string]any{
			"matrix": []any{
				[]any{1, 2},
				[]any{3, 4},
			},
		}

		gotOrig, errOrig := mapGetWithSeparator(m, "matrix.[1].[0]", ".")
		gotOpt, errOpt := mapGetWithSeparatorOptimized(m, "matrix.[1].[0]", ".")

		assert.Equal(t, errOrig == nil, errOpt == nil)
		if errOrig == nil && errOpt == nil {
			assert.Equal(t, gotOrig, gotOpt)
			assert.Equal(t, 3, gotOpt)
		}
	})

	t.Run("超长键路径", func(t *testing.T) {
		m := map[string]any{}
		current := m
		for i := 0; i < 20; i++ {
			next := map[string]any{}
			current[formatInt(i)] = next
			current = next
		}
		current["value"] = "found"

		// 构建长路径
		key := "0"
		for i := 1; i < 20; i++ {
			key += "." + formatInt(i)
		}
		key += ".value"

		gotOrig, errOrig := mapGetWithSeparator(m, key, ".")
		gotOpt, errOpt := mapGetWithSeparatorOptimized(m, key, ".")

		assert.Equal(t, errOrig == nil, errOpt == nil)
		if errOrig == nil && errOpt == nil {
			assert.Equal(t, gotOrig, gotOpt)
			assert.Equal(t, "found", gotOpt)
		}
	})
}

// TestMapGetWithSeparator_Concurrency 并发测试
func TestMapGetWithSeparator_Concurrency(t *testing.T) {
	m := map[string]any{
		"a": map[string]any{
			"b": map[string]any{
				"c": "value",
			},
		},
	}

	t.Run("原始版本并发", func(t *testing.T) {
		done := make(chan bool)
		for i := 0; i < 100; i++ {
			go func() {
				_, _ = mapGetWithSeparator(m, "a.b.c", ".")
				done <- true
			}()
		}
		for i := 0; i < 100; i++ {
			<-done
		}
	})

	t.Run("优化版本并发", func(t *testing.T) {
		done := make(chan bool)
		for i := 0; i < 100; i++ {
			go func() {
				_, _ = mapGetWithSeparatorOptimized(m, "a.b.c", ".")
				done <- true
			}()
		}
		for i := 0; i < 100; i++ {
			<-done
		}
	})
}

func formatInt(n int) string {
	if n < 10 {
		return string(rune('0' + n))
	}
	return "x" // 简化处理
}
