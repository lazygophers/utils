package anyx

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestKeyByInt64 测试 KeyByInt64 函数
func TestKeyByInt64(t *testing.T) {
	type Person struct {
		ID   int64
		Name string
	}

	tests := []struct {
		name      string
		list     []*Person
		fieldName string
		want      map[int64]*Person
		wantPanic bool
	}{
		{
			name:      "正常情况",
			list:     []*Person{{ID: 1, Name: "Alice"}, {ID: 2, Name: "Bob"}},
			fieldName: "ID",
			want:      map[int64]*Person{1: {ID: 1, Name: "Alice"}, 2: {ID: 2, Name: "Bob"}},
		},
		{
			name:      "空列表",
			list:     []*Person{},
			fieldName: "ID",
			want:      map[int64]*Person{},
		},
		{
			name:      "nil列表",
			list:     nil,
			fieldName: "ID",
			want:      map[int64]*Person{},
		},
		{
			name:      "字段不存在",
			list:     []*Person{{ID: 1, Name: "Alice"}},
			fieldName: "Age",
			wantPanic: true,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			if tt.wantPanic {
				assert.Panics(t, func() {
					KeyByInt64(tt.list, tt.fieldName)
				})
				return
			}

			got := KeyByInt64(tt.list, tt.fieldName)
			assert.Equal(t, tt.want, got)
		})
	}
}