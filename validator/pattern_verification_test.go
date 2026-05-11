package validator

import (
	"reflect"
	"testing"
)

// 验证优化后的 Pattern 函数功能正确性
func TestPatternOptimization(t *testing.T) {
	tests := []struct {
		name     string
		pattern  string
		value    string
		expected bool
	}{
		{
			name:     "有效邮箱",
			pattern:  `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`,
			value:    "test@example.com",
			expected: true,
		},
		{
			name:     "无效邮箱-缺少@",
			pattern:  `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`,
			value:    "invalid",
			expected: false,
		},
		{
			name:     "固定长度-5位数字",
			pattern:  `^\d{5}$`,
			value:    "12345",
			expected: true,
		},
		{
			name:     "固定长度-不足5位",
			pattern:  `^\d{5}$`,
			value:    "123",
			expected: false,
		},
		{
			name:     "字面量匹配",
			pattern:  `^hello$`,
			value:    "hello",
			expected: true,
		},
		{
			name:     "字面量不匹配",
			pattern:  `^hello$`,
			value:    "world",
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			validator := Pattern(tt.pattern)
			fl := &testFieldLevel{value: reflect.ValueOf(tt.value)}
			result := validator(fl)
			if result != tt.expected {
				t.Errorf("Pattern() = %v, want %v", result, tt.expected)
			}
		})
	}
}

// 测试用 FieldLevel 实现
type testFieldLevel struct {
	value reflect.Value
}

func (t *testFieldLevel) Top() reflect.Value {
	return t.value
}

func (t *testFieldLevel) Parent() reflect.Value {
	return t.value
}

func (t *testFieldLevel) Field() reflect.Value {
	return t.value
}

func (t *testFieldLevel) FieldName() string {
	return "test"
}

func (t *testFieldLevel) StructFieldName() string {
	return "Test"
}

func (t *testFieldLevel) Param() string {
	return ""
}

func (t *testFieldLevel) GetTag(key string) string {
	return ""
}

func (t *testFieldLevel) GetFieldByName(name string) reflect.Value {
	return reflect.Value{}
}

// 性能对比基准测试
func BenchmarkPatternOptimized(b *testing.B) {
	pattern := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
	validator := Pattern(pattern)
	fl := &testFieldLevel{value: reflect.ValueOf("test@example.com")}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		validator(fl)
	}
}
