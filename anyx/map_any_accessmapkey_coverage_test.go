package anyx

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestAccessMapKey_Coverage 全面测试 accessMapKey 函数覆盖率
func TestAccessMapKey_Coverage(t *testing.T) {
	tests := []struct {
		name           string
		current        any
		key            string
		expected       any
		wantErr        bool
		errType        error
		description    string
		skipValueCheck bool
	}{
		// ===== map[string]any 测试 =====
		{
			name:        "map[string]any 简单键命中",
			current:     map[string]any{"name": "John", "age": 30},
			key:         "name",
			expected:    "John",
			wantErr:     false,
			description: "验证基本的 map[string]any 键访问",
		},
		{
			name:        "map[string]any 键未命中",
			current:     map[string]any{"name": "John"},
			key:         "nonexistent",
			expected:    nil,
			wantErr:     true,
			errType:     ErrNotFound,
			description: "验证 map[string]any 键未命中返回 ErrNotFound",
		},
		{
			name:        "map[string]any 空字符串键",
			current:     map[string]any{"": "empty_value", "key": "value"},
			key:         "",
			expected:    "empty_value",
			wantErr:     false,
			description: "验证空字符串作为键的访问",
		},
		{
			name:        "map[string]any nil 值",
			current:     map[string]any{"nil_key": nil},
			key:         "nil_key",
			expected:    nil,
			wantErr:     false,
			description: "验证键存在但值为 nil 的情况",
		},
		{
			name:        "map[string]any 空 map",
			current:     map[string]any{},
			key:         "key",
			expected:    nil,
			wantErr:     true,
			errType:     ErrNotFound,
			description: "验证空 map 的访问",
		},

		// ===== map[any]any 测试 =====
		{
			name:        "map[any]any 字符串键命中",
			current:     map[any]any{"name": "John", 42: "answer"},
			key:         "name",
			expected:    "John",
			wantErr:     false,
			description: "验证 map[any]any 字符串键访问",
		},
		{
			name:        "map[any]any 整数键（未命中）",
			current:     map[any]any{"name": "John", 42: "answer"},
			key:         "100",
			expected:    nil,
			wantErr:     true,
			errType:     ErrNotFound,
			description: "验证 map[any]any 中不存在的字符串键",
		},
		{
			name:        "map[any]any 空字符串键",
			current:     map[any]any{"": "empty"},
			key:         "",
			expected:    "empty",
			wantErr:     false,
			description: "验证 map[any]any 空字符串键",
		},
		{
			name:        "map[any]any 空 map",
			current:     map[any]any{},
			key:         "key",
			expected:    nil,
			wantErr:     true,
			errType:     ErrNotFound,
			description: "验证空 map[any]any",
		},

		// ===== 无效类型测试 =====
		{
			name:           "无效类型 - 切片",
			current:        []string{"a", "b", "c"},
			key:            "0",
			expected:       nil,
			wantErr:        true,
			errType:        ErrInvalidMapType,
			description:    "验证切片类型返回 ErrInvalidMapType",
			skipValueCheck: true,
		},
		{
			name:           "无效类型 - 整数",
			current:        42,
			key:            "key",
			expected:       nil,
			wantErr:        true,
			errType:        ErrInvalidMapType,
			description:    "验证整数类型返回 ErrInvalidMapType",
			skipValueCheck: true,
		},
		{
			name:           "无效类型 - 字符串",
			current:        "not a map",
			key:            "key",
			expected:       nil,
			wantErr:        true,
			errType:        ErrInvalidMapType,
			description:    "验证字符串类型返回 ErrInvalidMapType",
			skipValueCheck: true,
		},
		{
			name:           "无效类型 - nil",
			current:        nil,
			key:            "key",
			expected:       nil,
			wantErr:        true,
			errType:        ErrInvalidMapType,
			description:    "验证 nil 输入返回 ErrInvalidMapType",
			skipValueCheck: true,
		},
		{
			name:           "无效类型 - 结构体",
			current:        struct{ Name string }{Name: "test"},
			key:            "Name",
			expected:       nil,
			wantErr:        true,
			errType:        ErrInvalidMapType,
			description:    "验证结构体类型返回 ErrInvalidMapType",
			skipValueCheck: true,
		},

		// ===== 边界情况 =====
		{
			name:        "边界 - 大型 map",
			current:     func() map[string]any { m := make(map[string]any, 1000); for i := 0; i < 1000; i++ { m[string(rune(i))] = i }; return m }(),
			key:         "X",
			expected:    88,
			wantErr:     false,
			description: "验证大型 map 的性能和正确性（ASCII 'X' = 88）",
		},
		{
			name:        "边界 - 复杂值类型",
			current:     map[string]any{"nested": map[string]any{"inner": "value"}},
			key:         "nested",
			expected:    map[string]any{"inner": "value"},
			wantErr:     false,
			description: "验证值类型为嵌套 map 的访问",
		},
		{
			name:           "边界 - 函数值",
			current:        map[string]any{"func": func() {}},
			key:            "func",
			expected:       nil,
			wantErr:        false,
			skipValueCheck: true, // 函数值不能直接比较
			description:    "验证值类型为函数的访问",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := accessMapKey(tt.current, tt.key)

			// 验证错误
			if tt.wantErr {
				assert.Error(t, err, tt.description)
				if tt.errType != nil {
					assert.ErrorIs(t, err, tt.errType, tt.description)
				}
			} else {
				assert.NoError(t, err, tt.description)
			}

			// 验证返回值
			if !tt.skipValueCheck {
				assert.Equal(t, tt.expected, result, tt.description)
			}
			// skipValueCheck 时只验证没有错误即可
		})
	}
}

// TestAccessMapKey_ConcurrentAccess 测试并发访问安全性
func TestAccessMapKey_ConcurrentAccess(t *testing.T) {
	// 准备测试数据
	testMaps := []map[string]any{
		{"key1": "value1", "key2": "value2"},
		{"a": 1, "b": 2, "c": 3},
		{"x": map[string]any{"nested": "value"}},
	}

	done := make(chan bool)
	iterations := 1000

	// 多个 goroutine 并发读取
	for i := 0; i < 10; i++ {
		go func(idx int) {
			defer func() { done <- true }()
			m := testMaps[idx%len(testMaps)]
			for j := 0; j < iterations; j++ {
				_, _ = accessMapKey(m, "key1")
				_, _ = accessMapKey(m, "nonexistent")
			}
		}(i)
	}

	// 等待所有 goroutine 完成
	for i := 0; i < 10; i++ {
		<-done
	}

	// 如果没有 panic 或 deadlock，则测试通过
	t.Log("并发访问测试通过")
}

// TestAccessMapKey_NilMaps 测试 nil map 处理
func TestAccessMapKey_NilMaps(t *testing.T) {
	tests := []struct {
		name        string
		current     any
		key         string
		wantErr     bool
		errType     error
		description string
	}{
		{
			name:        "nil map[string]any",
			current:     (map[string]any)(nil),
			key:         "key",
			wantErr:     true,
			errType:     ErrNotFound,
			description: "验证 nil map[string]any 返回 ErrNotFound",
		},
		{
			name:        "nil map[any]any",
			current:     (map[any]any)(nil),
			key:         "key",
			wantErr:     true,
			errType:     ErrNotFound,
			description: "验证 nil map[any]any 返回 ErrNotFound",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := accessMapKey(tt.current, tt.key)

			assert.Error(t, err, tt.description)
			assert.ErrorIs(t, err, tt.errType, tt.description)
			assert.Nil(t, result, tt.description)
		})
	}
}

// TestAccessMapKey_SpecialKeys 测试特殊键值
func TestAccessMapKey_SpecialKeys(t *testing.T) {
	tests := []struct {
		name        string
		current     map[string]any
		key         string
		expected    any
		wantErr     bool
		description string
	}{
		{
			name:        "带空格的键",
			current:     map[string]any{"key with spaces": "value"},
			key:         "key with spaces",
			expected:    "value",
			wantErr:     false,
			description: "验证带空格的键访问",
		},
		{
			name:        "带特殊字符的键",
			current:     map[string]any{"key/with/slashes": "value"},
			key:         "key/with/slashes",
			expected:    "value",
			wantErr:     false,
			description: "验证带特殊字符的键访问",
		},
		{
			name:        "Unicode 键",
			current:     map[string]any{"键": "值"},
			key:         "键",
			expected:    "值",
			wantErr:     false,
			description: "验证 Unicode 键访问",
		},
		{
			name:        "数字字符串键",
			current:     map[string]any{"123": "number key"},
			key:         "123",
			expected:    "number key",
			wantErr:     false,
			description: "验证数字字符串键访问",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := accessMapKey(tt.current, tt.key)

			if tt.wantErr {
				assert.Error(t, err, tt.description)
			} else {
				assert.NoError(t, err, tt.description)
				assert.Equal(t, tt.expected, result, tt.description)
			}
		})
	}
}

// TestAccessMapKey_TypePreservation 测试类型保持
func TestAccessMapKey_TypePreservation(t *testing.T) {
	testCases := []struct {
		name     string
		input    map[string]any
		key      string
		expected any
	}{
		{"字符串值", map[string]any{"str": "hello"}, "str", "hello"},
		{"整数值", map[string]any{"int": 42}, "int", 42},
		{"浮点数值", map[string]any{"float": 3.14}, "float", 3.14},
		{"布尔值", map[string]any{"bool": true}, "bool", true},
		{"切片值", map[string]any{"slice": []int{1, 2, 3}}, "slice", []int{1, 2, 3}},
		{"map 值", map[string]any{"map": map[string]any{"nested": true}}, "map", map[string]any{"nested": true}},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result, err := accessMapKey(tc.input, tc.key)
			assert.NoError(t, err)
			assert.Equal(t, tc.expected, result, "类型应该保持不变")
		})
	}
}
