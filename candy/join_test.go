package candy

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestJoin 测试Join连接函数
func TestJoin(t *testing.T) {
	t.Parallel()

	// 整数类型测试
	intTests := []struct {
		name string
		give []int
		glue string
		want string
	}{
		{"默认分隔符", []int{1, 2, 3}, "", "1,2,3"},
		{"自定义分隔符", []int{1, 2, 3}, "-", "1-2-3"},
		{"空分隔符", []int{1, 2, 3}, "", "1,2,3"},
		{"单元素", []int{42}, ",", "42"},
		{"空切片", []int{}, ",", ""},
		{"nil切片", nil, ",", ""},
		{"长分隔符", []int{1, 2, 3}, "->", "1->2->3"},
	}

	for _, tt := range intTests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			var got string
			if tt.glue == "" {
				got = Join(tt.give)
			} else {
				got = Join(tt.give, tt.glue)
			}
			assert.Equal(t, tt.want, got, "Join() 整数的结果应与期望值相等")
		})
	}

	// 字符串类型测试
	stringTests := []struct {
		name string
		give []string
		glue string
		want string
	}{
		{"字符串切片默认分隔符", []string{"a", "b", "c"}, "", "a,b,c"},
		{"字符串切片自定义分隔符", []string{"a", "b", "c"}, " ", "a b c"},
		{"字符串切片单元素", []string{"hello"}, ",", "hello"},
		{"字符串切片空切片", []string{}, ",", ""},
		{"字符串切片nil切片", nil, ",", ""},
	}

	for _, tt := range stringTests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			var got string
			if tt.glue == "" {
				got = Join(tt.give)
			} else {
				got = Join(tt.give, tt.glue)
			}
			assert.Equal(t, tt.want, got, "Join() 字符串的结果应与期望值相等")
		})
	}
}
