package anyx

import (
	"testing"
)

// TestJoinPath_Integration 集成测试：通过实际使用场景验证 joinPath
func TestJoinPath_Integration(t *testing.T) {
	// 模拟 mapGetWithSeparator 中的使用场景
	tests := []struct {
		name     string
		parts    []string
		sep      string
		expected string
	}{
		{
			name:     "空路径",
			parts:    []string{},
			sep:      ".",
			expected: "",
		},
		{
			name:     "单级路径",
			parts:    []string{"root"},
			sep:      ".",
			expected: "root",
		},
		{
			name:     "两级路径",
			parts:    []string{"root", "child"},
			sep:      ".",
			expected: "root.child",
		},
		{
			name:     "三级路径",
			parts:    []string{"root", "child", "grandchild"},
			sep:      ".",
			expected: "root.child.grandchild",
		},
		{
			name:     "斜杠分隔符",
			parts:    []string{"var", "log", "app"},
			sep:      "/",
			expected: "var/log/app",
		},
		{
			name:     "路径片段（切片）",
			parts:    []string{"root", "level1", "level2"},
			sep:      ".",
			expected: "root.level1.level2",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := joinPath(tt.parts, tt.sep)
			if result != tt.expected {
				t.Errorf("joinPath(%v, %q) = %q, want %q",
					tt.parts, tt.sep, result, tt.expected)
			}
		})
	}
}

// TestJoinPath_CoverageAllBranches 覆盖所有分支
func TestJoinPath_CoverageAllBranches(t *testing.T) {
	// 分支 0: 空 slice
	result := joinPath([]string{}, ".")
	if result != "" {
		t.Errorf("Expected empty string, got %q", result)
	}

	// 分支 1: 单元素
	result = joinPath([]string{"single"}, ".")
	if result != "single" {
		t.Errorf("Expected 'single', got %q", result)
	}

	// 分支 2: 双元素（快速路径）
	result = joinPath([]string{"a", "b"}, ".")
	if result != "a.b" {
		t.Errorf("Expected 'a.b', got %q", result)
	}

	// 分支 3: 三元素（快速路径）
	result = joinPath([]string{"a", "b", "c"}, ".")
	if result != "a.b.c" {
		t.Errorf("Expected 'a.b.c', got %q", result)
	}

	// 分支 4: 四元素（Builder 路径）
	result = joinPath([]string{"a", "b", "c", "d"}, ".")
	if result != "a.b.c.d" {
		t.Errorf("Expected 'a.b.c.d', got %q", result)
	}

	// 分支 5: 多元素（Builder 路径）
	parts := make([]string, 10)
	for i := range parts {
		parts[i] = "x"
	}
	result = joinPath(parts, ".")
	expected := "x.x.x.x.x.x.x.x.x.x"
	if result != expected {
		t.Errorf("Expected %q, got %q", expected, result)
	}
}
