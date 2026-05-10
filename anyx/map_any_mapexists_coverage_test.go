package anyx

import (
	"fmt"
	"testing"
)

// TestMapExists_BasicFunctionality 测试基本功能
func TestMapExists_BasicFunctionality(t *testing.T) {
	m := map[string]any{
		"name": "John",
		"age":  30,
	}

	tests := []struct {
		name     string
		key      string
		expected bool
	}{
		{"存在的 key", "name", true},
		{"不存在的 key", "missing", false},
		{"空 key", "", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := MapExists(m, tt.key)
			if result != tt.expected {
				t.Errorf("MapExists(%q) = %v, 期望 %v", tt.key, result, tt.expected)
			}
		})
	}
}

// TestMapExists_NestedKeys 测试嵌套 key
func TestMapExists_NestedKeys(t *testing.T) {
	m := map[string]any{
		"user": map[string]any{
			"name": "John",
			"address": map[string]any{
				"city": "New York",
			},
		},
	}

	tests := []struct {
		name     string
		key      string
		expected bool
	}{
		{"嵌套 key 存在", "user.name", true},
		{"深度嵌套 key 存在", "user.address.city", true},
		{"嵌套 key 不存在", "user.email", false},
		{"部分路径不存在", "admin.name", false},
		{"中间路径不存在", "user.phone.type", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := MapExists(m, tt.key)
			if result != tt.expected {
				t.Errorf("MapExists(%q) = %v, 期望 %v", tt.key, result, tt.expected)
			}
		})
	}
}

// TestMapExists_ArrayIndices 测试数组索引
func TestMapExists_ArrayIndices(t *testing.T) {
	m := map[string]any{
		"items": []any{"a", "b", "c"},
		"nested": map[string]any{
			"array": []any{1, 2, 3},
		},
		"maps": []map[string]any{
			{"id": 1},
			{"id": 2},
		},
	}

	tests := []struct {
		name     string
		key      string
		expected bool
	}{
		{"数组索引存在", "items[0]", true},
		{"数组索引存在（中间）", "items[1]", true},
		{"数组索引存在（末尾）", "items[2]", true},
		{"数组索引不存在", "items[10]", false},
		{"嵌套数组索引", "nested.array[1]", true},
		{"嵌套数组索引不存在", "nested.array[10]", false},
		{"map 数组索引", "maps[0].id", true},
		{"map 数组索引嵌套", "maps[1].id", true},
		{"负数索引", "items[-1]", false},
		{"无效索引", "items[abc]", false},
		{"空索引", "items[]", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := MapExists(m, tt.key)
			if result != tt.expected {
				t.Errorf("MapExists(%q) = %v, 期望 %v", tt.key, result, tt.expected)
			}
		})
	}
}

// TestMapExists_EdgeCases 测试边界情况
func TestMapExists_EdgeCases(t *testing.T) {
	tests := []struct {
		name     string
		m        map[string]any
		key      string
		expected bool
	}{
		{
			name:     "空 map",
			m:        map[string]any{},
			key:      "key",
			expected: false,
		},
		{
			name:     "nil map",
			m:        nil,
			key:      "key",
			expected: false,
		},
		{
			name:     "空 key",
			m:        map[string]any{"key": "value"},
			key:      "",
			expected: false,
		},
		{
			name:     "以点开头",
			m:        map[string]any{"key": "value"},
			key:      ".key",
			expected: false,
		},
		{
			name:     "以点结尾",
			m:        map[string]any{"key": "value"},
			key:      "key.",
			expected: false,
		},
		{
			name:     "连续点",
			m:        map[string]any{"key": "value"},
			key:      "key..value",
			expected: false,
		},
		{
			name:     "只有点",
			m:        map[string]any{"key": "value"},
			key:      ".",
			expected: false,
		},
		{
			name:     "多个点",
			m:        map[string]any{"a": map[string]any{"b": map[string]any{"c": "value"}}},
			key:      "a.b.c",
			expected: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := MapExists(tt.m, tt.key)
			if result != tt.expected {
				t.Errorf("MapExists(%q) = %v, 期望 %v", tt.key, result, tt.expected)
			}
		})
	}
}

// TestMapExists_TypeMismatch 测试类型不匹配
func TestMapExists_TypeMismatch(t *testing.T) {
	m := map[string]any{
		"string": "value",
		"number": 123,
		"bool":   true,
		"slice":  []any{1, 2, 3},
	}

	tests := []struct {
		name     string
		key      string
		expected bool
	}{
		{"字符串尝试嵌套访问", "string.field", false},
		{"数字尝试嵌套访问", "number.field", false},
		{"布尔尝试嵌套访问", "bool.field", false},
		{"切片尝试键访问", "slice.key", false},
		{"切片索引存在", "slice[0]", true},
		{"切片索引不存在", "slice[10]", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := MapExists(m, tt.key)
			if result != tt.expected {
				t.Errorf("MapExists(%q) = %v, 期望 %v", tt.key, result, tt.expected)
			}
		})
	}
}

// TestMapExists_ComplexScenarios 测试复杂场景
func TestMapExists_ComplexScenarios(t *testing.T) {
	m := map[string]any{
		"users": []map[string]any{
			{
				"name": "Alice",
				"contacts": map[string]any{
					"email": "alice@example.com",
					"phone": "123-456-7890",
				},
			},
			{
				"name": "Bob",
				"contacts": map[string]any{
					"email": "bob@example.com",
				},
			},
		},
		"settings": map[string]any{
			"theme": map[string]any{
				"dark": map[string]any{
					"primary": "#000",
					"secondary": "#333",
				},
			},
		},
	}

	tests := []struct {
		name     string
		key      string
		expected bool
	}{
		{"数组 + map 嵌套", "users[0].name", true},
		{"数组 + 深度 map", "users[0].contacts.email", true},
		{"数组索引 + 不存在的字段", "users[0].contacts.address", false},
		{"数组索引越界", "users[10].name", false},
		{"三深度嵌套 map", "settings.theme.dark.primary", true},
		{"三深度嵌套不存在", "settings.theme.light.primary", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := MapExists(m, tt.key)
			if result != tt.expected {
				t.Errorf("MapExists(%q) = %v, 期望 %v", tt.key, result, tt.expected)
			}
		})
	}
}

// TestMapExistsWithSep_Basic 测试 MapExistsWithSep 基本功能
func TestMapExistsWithSep_Basic(t *testing.T) {
	m := map[string]any{
		"user": map[string]any{
			"name": "John",
		},
	}

	tests := []struct {
		name     string
		key      string
		sep      string
		expected bool
	}{
		{"默认分隔符", "user.name", ".", true},
		{"斜杠分隔符", "user/name", "/", true},
		{"连字符分隔符", "user-name", "-", true},
		{"下划线分隔符", "user_name", "_", true},
		{"自定义分隔符不存在", "user|missing", "|", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := MapExistsWithSep(m, tt.key, tt.sep)
			if result != tt.expected {
				t.Errorf("MapExistsWithSep(%q, %q) = %v, 期望 %v",
					tt.key, tt.sep, result, tt.expected)
			}
		})
	}
}

// TestMapExistsWithSep_ArrayWithCustomSeparator 测试自定义分隔符与数组
func TestMapExistsWithSep_ArrayWithCustomSeparator(t *testing.T) {
	m := map[string]any{
		"items": []any{1, 2, 3},
	}

	tests := []struct {
		name     string
		key      string
		sep      string
		expected bool
	}{
		{"点分隔符 + 数组", "items[0]", ".", true},
		{"斜杠分隔符 + 数组", "items[0]", "/", true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := MapExistsWithSep(m, tt.key, tt.sep)
			if result != tt.expected {
				t.Errorf("MapExistsWithSep(%q, %q) = %v, 期望 %v",
					tt.key, tt.sep, result, tt.expected)
			}
		})
	}
}

// TestMapExists_EmptyValues 测试空值情况
func TestMapExists_EmptyValues(t *testing.T) {
	m := map[string]any{
		"emptyString": "",
		"zero":        0,
		"false":       false,
		"nil":         nil,
		"emptyMap":    map[string]any{},
		"emptySlice":  []any{},
	}

	tests := []struct {
		name     string
		key      string
		expected bool
	}{
		{"空字符串存在", "emptyString", true},
		{"零值存在", "zero", true},
		{"假值存在", "false", true},
		{"nil 存在", "nil", true},
		{"空 map 存在", "emptyMap", true},
		{"空切片存在", "emptySlice", true},
		{"空 map 嵌套访问", "emptyMap.key", false},
		{"空切片索引", "emptySlice[0]", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := MapExists(m, tt.key)
			if result != tt.expected {
				t.Errorf("MapExists(%q) = %v, 期望 %v", tt.key, result, tt.expected)
			}
		})
	}
}

// TestMapExists_DeepNesting 测试深度嵌套
func TestMapExists_DeepNesting(t *testing.T) {
	m := map[string]any{
		"a": map[string]any{
			"b": map[string]any{
				"c": map[string]any{
					"d": map[string]any{
						"e": map[string]any{
							"f": "value",
						},
					},
				},
			},
		},
	}

	tests := []struct {
		name     string
		key      string
		expected bool
	}{
		{"6 层嵌套存在", "a.b.c.d.e.f", true},
		{"5 层嵌套存在", "a.b.c.d.e", true},
		{"4 层嵌套存在", "a.b.c.d", true},
		{"6 层嵌套不存在", "a.b.c.d.e.x", false},
		{"错误路径", "a.x.c.d.e.f", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := MapExists(m, tt.key)
			if result != tt.expected {
				t.Errorf("MapExists(%q) = %v, 期望 %v", tt.key, result, tt.expected)
			}
		})
	}
}

// TestMapExists_LargeMap 测试大型 map
func TestMapExists_LargeMap(t *testing.T) {
	m := make(map[string]any, 1000)
	for i := 0; i < 1000; i++ {
		m[fmt.Sprintf("key%d", i)] = fmt.Sprintf("value%d", i)
	}

	tests := []struct {
		name     string
		key      string
		expected bool
	}{
		{"开头 key", "key0", true},
		{"中间 key", "key500", true},
		{"结尾 key", "key999", true},
		{"不存在的 key", "key1000", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := MapExists(m, tt.key)
			if result != tt.expected {
				t.Errorf("MapExists(%q) = %v, 期望 %v", tt.key, result, tt.expected)
			}
		})
	}
}

// TestMapExists_ConcurrentAccess 测试并发访问
func TestMapExists_ConcurrentAccess(t *testing.T) {
	m := map[string]any{
		"user": map[string]any{
			"name": "John",
		},
	}

	done := make(chan bool)
	iterations := 1000

	// 多个 goroutine 并发读取
	for i := 0; i < 10; i++ {
		go func() {
			for j := 0; j < iterations; j++ {
				_ = MapExists(m, "user.name")
				_ = MapExists(m, "user.missing")
			}
			done <- true
		}()
	}

	// 等待所有 goroutine 完成
	for i := 0; i < 10; i++ {
		<-done
	}
}

// TestMapExists_SpecialCharacters 测试特殊字符
func TestMapExists_SpecialCharacters(t *testing.T) {
	m := map[string]any{
		"key-with-dash":    "value1",
		"key_with_underscore": "value2",
		"key.with.dots":    "value3",
		"key@symbol":       "value4",
		"key$美元":         "value5",
	}

	tests := []struct {
		name     string
		key      string
		expected bool
	}{
		{"带连字符", "key-with-dash", true},
		{"带下划线", "key_with_underscore", true},
		{"带点", "key.with.dots", false}, // 点会被当作分隔符，所以这是嵌套访问
		{"带 @ 符号", "key@symbol", true},
		{"带 $ 符号", "key$美元", true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := MapExists(m, tt.key)
			if result != tt.expected {
				t.Errorf("MapExists(%q) = %v, 期望 %v", tt.key, result, tt.expected)
			}
		})
	}
}

// TestMapExists_MapExistsWithSep_Equivalence 测试 MapExists 和 MapExistsWithSep 等价性
func TestMapExists_MapExistsWithSep_Equivalence(t *testing.T) {
	m := map[string]any{
		"user": map[string]any{
			"name": "John",
			"age":  30,
		},
	}

	testKeys := []string{
		"user",
		"user.name",
		"user.age",
		"user.missing",
		"missing",
		"",
	}

	for _, key := range testKeys {
		result1 := MapExists(m, key)
		result2 := MapExistsWithSep(m, key, ".")
		if result1 != result2 {
			t.Errorf("MapExists(%q) = %v, MapExistsWithSep(%q, \".\") = %v, 不一致",
				key, result1, key, result2)
		}
	}
}

// TestMapExists_BracketNotation 测试括号表示法
func TestMapExists_BracketNotation(t *testing.T) {
	m := map[string]any{
		"items": []any{
			map[string]any{"name": "item1"},
			map[string]any{"name": "item2"},
			map[string]any{"name": "item3"},
		},
		"nested": map[string]any{
			"array": []any{1, 2, 3},
		},
	}

	tests := []struct {
		name     string
		key      string
		expected bool
	}{
		{"简单数组索引", "items[0]", true},
		{"数组索引 + 嵌套", "items[0].name", true},
		{"嵌套 + 数组索引", "nested.array[1]", true},
		{"多个数组索引", "items[1].name", true},
		{"索引不存在", "items[10].name", false},
		{"中间数组索引不存在", "items[1].missing", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := MapExists(m, tt.key)
			if result != tt.expected {
				t.Errorf("MapExists(%q) = %v, 期望 %v", tt.key, result, tt.expected)
			}
		})
	}
}

// TestMapExists_ZeroIndexNegativeIndex 测试零索引和负索引
func TestMapExists_ZeroIndexNegativeIndex(t *testing.T) {
	m := map[string]any{
		"items": []any{"a", "b", "c"},
	}

	tests := []struct {
		name     string
		key      string
		expected bool
	}{
		{"零索引", "items[0]", true},
		{"负索引（不支持）", "items[-1]", false},
		{"索引为 0 的字符串", "items[abc]", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := MapExists(m, tt.key)
			if result != tt.expected {
				t.Errorf("MapExists(%q) = %v, 期望 %v", tt.key, result, tt.expected)
			}
		})
	}
}
