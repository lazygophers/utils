// Package candy 提供实用的语法糖函数，简化常见的编程操作
package candy

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestFilterNot 测试 FilterNot 函数
func TestFilterNot(t *testing.T) {
	// 整数类型测试
	t.Run("整数类型", func(t *testing.T) {
		tests := []struct {
			name string
			give []int
			f    func(int) bool
			want []int
		}{
			{
				name: "过滤偶数保留奇数",
				give: []int{1, 2, 3, 4, 5, 6},
				f: func(n int) bool {
					return n%2 == 0
				},
				want: []int{1, 3, 5},
			},
			{
				name: "过滤正数保留负数和零",
				give: []int{-1, 0, 1, -2, 2},
				f: func(n int) bool {
					return n > 0
				},
				want: []int{-1, 0, -2},
			},
			{
				name: "过滤大数保留小数",
				give: []int{1, 10, 100, 1000},
				f: func(n int) bool {
					return n > 50
				},
				want: []int{1, 10},
			},
			{
				name: "空切片输入",
				give: []int{},
				f: func(n int) bool {
					return n > 0
				},
				want: []int{},
			},
			{
				name: "全部元素都被过滤",
				give: []int{2, 4, 6, 8},
				f: func(n int) bool {
					return n%2 == 0
				},
				want: []int{},
			},
			{
				name: "没有元素被过滤",
				give: []int{1, 3, 5, 7},
				f: func(n int) bool {
					return n%2 == 0
				},
				want: []int{1, 3, 5, 7},
			},
		}

		for _, tt := range tests {
			tt := tt // 避免竞态
			t.Run(tt.name, func(t *testing.T) {
				got := FilterNot(tt.give, tt.f)
				assert.Equal(t, tt.want, got, "FilterNot() 的结果应与期望值相等")
			})
		}
	})
}
