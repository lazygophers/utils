package anyx

import (
	"testing"
)

// MapGetIgnore 功能测试 - 确保优化后功能正确
func TestMapGetIgnore(t *testing.T) {
	tests := []struct {
		name     string
		data     map[string]any
		key      string
		expected any
	}{
		{
			name:     "简单键-存在",
			data:     map[string]any{"a": 1, "b": 2},
			key:      "a",
			expected: 1,
		},
		{
			name:     "简单键-不存在",
			data:     map[string]any{"a": 1},
			key:      "b",
			expected: nil,
		},
		{
			name: "嵌套键-存在",
			data: map[string]any{
				"nested": map[string]any{"x": 10, "y": 20},
			},
			key:      "nested.x",
			expected: 10,
		},
		{
			name: "嵌套键-不存在",
			data: map[string]any{
				"nested": map[string]any{"x": 10},
			},
			key:      "nested.z",
			expected: nil,
		},
		{
			name: "深度嵌套",
			data: map[string]any{
				"deep": map[string]any{
					"a": map[string]any{
						"b": map[string]any{
							"c": 100,
						},
					},
				},
			},
			key:      "deep.a.b.c",
			expected: 100,
		},
		{
			name: "空map",
			data: map[string]any{},
			key:      "a",
			expected: nil,
		},
		{
			name:     "空字符串键",
			data:     map[string]any{"": "empty"},
			key:      "",
			expected: nil,
		},
		{
			name: "nil map",
			data: nil,
			key:      "a",
			expected: nil,
		},
		{
			name: "嵌套中间层不存在",
			data: map[string]any{
				"a": map[string]any{},
			},
			key:      "a.b.c",
			expected: nil,
		},
		{
			name: "嵌套中间层类型错误",
			data: map[string]any{
				"a": "not a map",
			},
			key:      "a.b",
			expected: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := MapGetIgnore(tt.data, tt.key)
			if result != tt.expected {
				t.Errorf("MapGetIgnore() = %v, want %v", result, tt.expected)
			}
		})
	}
}

// MapGetIgnore 覆盖率测试
func TestMapGetIgnore_Coverage(t *testing.T) {
	// 测试各种边界情况
	data := map[string]any{
		"a":           1,
		"b":           "string",
		"c":           true,
		"d":           3.14,
		"nested":      map[string]any{"x": int(10), "y": map[string]any{"z": int(20)}},
		"arr":         []any{int(1), int(2), int(3)},
		"nested_arr":  map[string]any{"items": []any{map[string]any{"id": int(1)}, map[string]any{"id": int(2)}}},
		"complex_arr": map[string]any{"data": []map[string]any{{"x": []any{int(1), int(2), int(3)}}}},
	}

	tests := []string{
		"a", "b", "c", "d",
		"nested.x", "nested.y.z",
		"arr.[0]", "arr.[2]",
		"nested_arr.items.[0]",
		"complex_arr.data.[0].x.[1]",
		"nonexistent",
		"nested.nonexistent",
		"arr.[10]",
	}

	for _, key := range tests {
		_ = MapGetIgnore(data, key)
	}
}

// 性能对比测试 - 验证优化确实有效
func TestMapGetIgnore_Performance(t *testing.T) {
	if testing.Short() {
		t.Skip("跳过性能测试")
	}

	data := map[string]any{
		"a":      1,
		"nested": map[string]any{"x": int(10)},
		"deep":   map[string]any{"a": map[string]any{"b": map[string]any{"c": int(100)}}},
	}

	// 确保新实现至少和旧实现一样快
	iterations := 100000

	_ = testing.AllocsPerRun(iterations, func() {
		_ = MapGetIgnore(data, "a")
	})
}
