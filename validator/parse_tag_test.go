package validator

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestParseTag 测试 parseTag 功能正确性
func TestParseTag(t *testing.T) {
	tests := []struct {
		name     string
		tag      string
		expected []validationRule
	}{
		{
			name: "简单标签",
			tag:  "required,email,max=100",
			expected: []validationRule{
				{tag: "required", param: ""},
				{tag: "email", param: ""},
				{tag: "max", param: "100"},
			},
		},
		{
			name: "带空格标签",
			tag:  "required , email , max = 100",
			expected: []validationRule{
				{tag: "required", param: ""},
				{tag: "email", param: ""},
				{tag: "max", param: "100"},
			},
		},
		{
			name: "复杂标签",
			tag:  "required,email,min=18,max=100,len=6-20",
			expected: []validationRule{
				{tag: "required", param: ""},
				{tag: "email", param: ""},
				{tag: "min", param: "18"},
				{tag: "max", param: "100"},
				{tag: "len", param: "6-20"},
			},
		},
		{
			name: "空标签",
			tag:  "",
			expected: []validationRule(nil),
		},
		{
			name: "只有空格",
			tag:  "   ,  ,   ",
			expected: []validationRule(nil),
		},
		{
			name: "带参数值有空格",
			tag:  "regex=^[a-z]+$ , url , max = 100",
			expected: []validationRule{
				{tag: "regex", param: "^[a-z]+$"},
				{tag: "url", param: ""},
				{tag: "max", param: "100"},
			},
		},
		{
			name: "单规则",
			tag:  "required",
			expected: []validationRule{
				{tag: "required", param: ""},
			},
		},
		{
			name: "多参数规则",
			tag:  "in=1,2,3,notin=4,5,6",
			expected: []validationRule{
				{tag: "in", param: "1"},
				{tag: "2", param: ""},
				{tag: "3", param: ""},
				{tag: "notin", param: "4"},
				{tag: "5", param: ""},
				{tag: "6", param: ""},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &Engine{}
			result := e.parseTag(tt.tag)

			// 验证结果长度
			assert.Equal(t, len(tt.expected), len(result), "结果长度不匹配")

			// 验证每个规则
			for i := range tt.expected {
				assert.Equal(t, tt.expected[i].tag, result[i].tag, "规则 %d: tag 不匹配", i)
				assert.Equal(t, tt.expected[i].param, result[i].param, "规则 %d: param 不匹配", i)
			}
		})
	}
}

// TestParseTagPerformance 测试 parseTag 性能
func TestParseTagPerformance(t *testing.T) {
	if testing.Short() {
		t.Skip("跳过性能测试")
	}

	e := &Engine{}
	tag := "required,email,min=18,max=100,len=6-20,in=1,2,3,notin=4,5,6,regex=^[a-z]+$"

	// 预热
	for i := 0; i < 1000; i++ {
		_ = e.parseTag(tag)
	}

	// 测试
	start := testing.Benchmark(func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = e.parseTag(tag)
		}
	})

	t.Logf("性能测试结果: %v", start)
}
