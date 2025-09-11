package candy

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestToArrayString 测试 ToArrayString 函数
func TestToArrayString(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		input    interface{}
		expected []string
	}{
		{"string input", "hello", []string{"hello"}},
		{"[]string input", []string{"a", "b", "c"}, []string{"a", "b", "c"}},
		{"empty string", "", []string{""}},
		{"empty []string", []string{}, []string{}},
		{"string with comma", "a,b,c", []string{"a", "b", "c"}},
		{"string with separator", "a|b|c", []string{"a|b|c"}},
		{"single element []string", []string{"single"}, []string{"single"}},
		{"unicode string", "你好", []string{"你好"}},
		{"unicode []string", []string{"你好", "世界"}, []string{"你好", "世界"}},
		{"nil input", nil, nil},
		{"nil string slice", ([]string)(nil), nil},
		{"int slice", []int{1, 2, 3}, []string{"1", "2", "3"}},
		{"bool slice", []bool{true, false}, []string{"1", "0"}}, // ToString converts bool to "1"/"0"
		{"float slice", []float64{1.1, 2.2}, []string{"1.100000", "2.200000"}}, // ToString precision
		{"interface slice", []interface{}{"hello", 42, true}, []string{"hello", "42", "1"}}, // bool becomes "1"
		{"non-slice int", 42, []string{"42"}},
		{"non-slice bool", true, []string{"1"}}, // ToString converts true to "1"
		{"non-slice float", 3.14, []string{"3.140000"}}, // ToString precision
		{"struct input", struct{Name string}{Name: "test"}, []string{`{"Name":"test"}`}}, // JSON serialization
		{"map input", map[string]int{"key": 42}, []string{`{"key":42}`}}, // JSON serialization
		{"slice with nil elements", []interface{}{nil, "test", nil}, []string{"", "test", ""}}, // nil becomes empty string
		{"string with multiple commas", "a,b,c,d,e", []string{"a", "b", "c", "d", "e"}},
		{"string with spaces and commas", "a, b, c", []string{"a", " b", " c"}}, // preserves spaces
		{"empty slice input", []int{}, []string{}}, // empty slice
		{"single element slice", []string{"only"}, []string{"only"}}, // already covered but explicit
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			result := ToArrayString(tt.input)
			assert.Equal(t, tt.expected, result, "ToArrayString() 的结果应与期望值相等")
		})
	}
}