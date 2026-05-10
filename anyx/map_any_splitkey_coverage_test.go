package anyx

import (
	"reflect"
	"testing"
)

// TestSplitKeyCoverage 全面覆盖率测试
func TestSplitKeyCoverage(t *testing.T) {
	tests := []struct {
		name     string
		key      string
		sep      string
		expected []string
	}{
		// 基础场景
		{"单字符 key", "a", ".", []string{"a"}},
		{"简单点分隔", "a.b.c", ".", []string{"a", "b", "c"}},
		{"无分隔符", "simplekey", ".", []string{"simplekey"}},

		// 数组索引场景
		{"单个数组索引", "items[0]", ".", []string{"items", "[0]"}},
		{"多个数组索引", "matrix[0][1]", ".", []string{"matrix", "[0]", "[1]"}},
		{"嵌套路径带数组", "data.items[0].name", ".", []string{"data", "items", "[0]", "name"}},
		{"数组索引带数字", "items[123]", ".", []string{"items", "[123]"}},
		{"纯数组路径", "[0][1][2]", ".", []string{"[0]", "[1]", "[2]"}},

		// 不同分隔符
		{"斜杠分隔", "api/v1/users", "/", []string{"api", "v1", "users"}},
		{"双冒号分隔", "ns::class::method", "::", []string{"ns", "class", "method"}},
		{"连字符分隔", "level1-level2-level3", "-", []string{"level1", "level2", "level3"}},

		// 边界情况
		{"以分隔符开头", ".a.b.c", ".", []string{"", "a", "b", "c"}},
		{"以分隔符结尾", "a.b.c.", ".", []string{"a", "b", "c", ""}},
		{"连续分隔符", "a..b", ".", []string{"a", "", "b"}},
		{"多个连续分隔符", "a...b", ".", []string{"a", "", "", "b"}},
		{"纯分隔符", "...", ".", []string{"", "", "", ""}},

		// 特殊字符
		{"括号内有点", "items[0].name", ".", []string{"items", "[0]", "name"}},
		{"括号前后有点", ".items[0].name.", ".", []string{"", "items", "[0]", "name", ""}},
		{"只有左括号", "items[", ".", []string{"items", "["}},
		{"括号内是空的", "items[]", ".", []string{"items", "[]"}},
		{"括号内有非数字", "items[abc]", ".", []string{"items", "[abc]"}},

		// 长字符串
		{"长键名", "very_long_key_name_with_many_underscores", ".", []string{"very_long_key_name_with_many_underscores"}},
		{"长路径", "a.b.c.d.e.f.g.h.i.j", ".", []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j"}},

		// 真实场景
		{"配置文件路径", "server.ssl.enabled", ".", []string{"server", "ssl", "enabled"}},
		{"API 响应路径", "data.users[0].profile.settings.theme", ".", []string{"data", "users", "[0]", "profile", "settings", "theme"}},
		{"嵌套数组访问", "results[0].tags[1]", ".", []string{"results", "[0]", "tags", "[1]"}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := splitKey(tt.key, tt.sep)
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("splitKey(%q, %q) = %v, want %v", tt.key, tt.sep, result, tt.expected)
			}
		})
	}
}

// TestSplitKeyEdgeCases 额外边界情况测试
func TestSplitKeyEdgeCases(t *testing.T) {
	t.Run("括号和分隔符混合", func(t *testing.T) {
		// "[0]." 不应该被分割为 "", "[0]"
		result := splitKey("[0].", ".")
		expected := []string{"[0]", ""}
		if !reflect.DeepEqual(result, expected) {
			t.Errorf("splitKey(%q, %q) = %v, want %v", "[0].", ".", result, expected)
		}
	})

	t.Run("分隔符在括号后", func(t *testing.T) {
		result := splitKey("items[0].name", ".")
		expected := []string{"items", "[0]", "name"}
		if !reflect.DeepEqual(result, expected) {
			t.Errorf("splitKey(%q, %q) = %v, want %v", "items[0].name", ".", result, expected)
		}
	})

	t.Run("括号内包含分隔符字符", func(t *testing.T) {
		// 括号内的点不应该被当作分隔符
		result := splitKey("items[0.name]", ".")
		expected := []string{"items", "[0.name]"}
		if !reflect.DeepEqual(result, expected) {
			t.Errorf("splitKey(%q, %q) = %v, want %v", "items[0.name]", ".", result, expected)
		}
	})

	t.Run("嵌套括号场景", func(t *testing.T) {
		result := splitKey("a[0].b[1]", ".")
		expected := []string{"a", "[0]", "b", "[1]"}
		if !reflect.DeepEqual(result, expected) {
			t.Errorf("splitKey(%q, %q) = %v, want %v", "a[0].b[1]", ".", result, expected)
		}
	})

	t.Run("空分隔符", func(t *testing.T) {
		result := splitKey("abc", "")
		// 空分隔符应该返回整个字符串作为一个部分
		expected := []string{"abc"}
		if !reflect.DeepEqual(result, expected) {
			t.Errorf("splitKey(%q, %q) = %v, want %v", "abc", "", result, expected)
		}
	})

	t.Run("单字符长分隔符", func(t *testing.T) {
		result := splitKey("a-b-c", "-")
		expected := []string{"a", "b", "c"}
		if !reflect.DeepEqual(result, expected) {
			t.Errorf("splitKey(%q, %q) = %v, want %v", "a-b-c", "-", result, expected)
		}
	})
}

// TestSplitKeyBracketHandling 括号处理专项测试
func TestSplitKeyBracketHandling(t *testing.T) {
	tests := []struct {
		name     string
		key      string
		sep      string
		expected []string
	}{
		{
			name:     "开括号前有内容",
			key:      "items[0]",
			sep:      ".",
			expected: []string{"items", "[0]"},
		},
		{
			name:     "闭括号后有分隔符",
			key:      "items[0].name",
			sep:      ".",
			expected: []string{"items", "[0]", "name"},
		},
		{
			name:     "闭括号后是字符串结尾",
			key:      "items[0]",
			sep:      ".",
			expected: []string{"items", "[0]"},
		},
		{
			name:     "只有开括号",
			key:      "items[",
			sep:      ".",
			expected: []string{"items", "["},
		},
		{
			name:     "只有闭括号",
			key:      "items]",
			sep:      ".",
			expected: []string{"items", "]"},
		},
		{
			name:     "括号在开头",
			key:      "[0]",
			sep:      ".",
			expected: []string{"[0]"},
		},
		{
			name:     "括号在结尾",
			key:      "items[0]",
			sep:      ".",
			expected: []string{"items", "[0]"},
		},
		{
			name:     "多个括号连续",
			key:      "[0][1][2]",
			sep:      ".",
			expected: []string{"[0]", "[1]", "[2]"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := splitKey(tt.key, tt.sep)
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("splitKey(%q, %q) = %v, want %v", tt.key, tt.sep, result, tt.expected)
			}
		})
	}
}

// TestSplitKeySeparatorCases 分隔符专项测试
func TestSplitKeySeparatorCases(t *testing.T) {
	tests := []struct {
		name     string
		key      string
		sep      string
		expected []string
	}{
		{
			name:     "单字符分隔符",
			key:      "a.b.c",
			sep:      ".",
			expected: []string{"a", "b", "c"},
		},
		{
			name:     "双字符分隔符",
			key:      "a::b::c",
			sep:      "::",
			expected: []string{"a", "b", "c"},
		},
		{
			name:     "三字符分隔符",
			key:      "a:::b:::c",
			sep:      ":::",
			expected: []string{"a", "b", "c"},
		},
		{
			name:     "分隔符不在字符串中",
			key:      "abcdefgh",
			sep:      "::",
			expected: []string{"abcdefgh"},
		},
		{
			name:     "分隔符在开头",
			key:      "::a::b",
			sep:      "::",
			expected: []string{"", "a", "b"},
		},
		{
			name:     "分隔符在结尾",
			key:      "a::b::",
			sep:      "::",
			expected: []string{"a", "b", ""},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := splitKey(tt.key, tt.sep)
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("splitKey(%q, %q) = %v, want %v", tt.key, tt.sep, result, tt.expected)
			}
		})
	}
}

// TestSplitKeyPerformanceCorrectness 性能优化后的正确性验证
func TestSplitKeyPerformanceCorrectness(t *testing.T) {
	// 验证所有变体实现的结果一致性
	testCases := []struct {
		key string
		sep string
	}{
		{"a.b.c", "."},
		{"items[0].name", "."},
		{"data.results[1].user.profile[2].settings", "."},
		{"api/v1/users", "/"},
		{"namespace::class::method", "::"},
		{".a.b.c.", "."},
		{"a..b", "."},
		{"[0][1][2]", "."},
		{"very_long_key_name_with_many_underscores.another_one", "."},
	}

	for _, tc := range testCases {
		expected := splitKeyCurrent(tc.key, tc.sep)
		result := splitKey(tc.key, tc.sep)

		if !reflect.DeepEqual(result, expected) {
			t.Errorf("splitKey(%q, %q) = %v, current impl gives %v",
				tc.key, tc.sep, result, expected)
		}
	}
}

// TestSplitKeyCornerCases 极端情况测试
func TestSplitKeyCornerCases(t *testing.T) {
	t.Run("超长单段", func(t *testing.T) {
		longStr := string(make([]byte, 1000))
		for i := range longStr {
			longStr = longStr[:i] + "a" + longStr[i+1:]
		}
		result := splitKey(longStr, ".")
		if len(result) != 1 || result[0] != longStr {
			t.Error("超长单段处理失败")
		}
	})

	t.Run("超多短段", func(t *testing.T) {
		key := "a.b.c.d.e.f.g.h.i.j.k.l.m.n.o.p.q.r.s.t.u.v.w.x.y.z"
		result := splitKey(key, ".")
		if len(result) != 26 {
			t.Errorf("期望 26 段，得到 %d", len(result))
		}
	})

	t.Run("所有字符都是分隔符", func(t *testing.T) {
		result := splitKey("....", ".")
		expected := []string{"", "", "", "", ""}
		if !reflect.DeepEqual(result, expected) {
			t.Errorf("splitKey(%q, %q) = %v, want %v", "....", ".", result, expected)
		}
	})
}
