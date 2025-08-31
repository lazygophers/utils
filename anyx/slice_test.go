package anyx

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestPluckInt32 测试 PluckInt32 函数
func TestPluckInt32(t *testing.T) {
	type User struct {
		ID   int32
		Name string
	}

	tests := []struct {
		name      string
		list     interface{}
		fieldName string
		want      []int32
		wantPanic bool
	}{
		{
			name:      "正常情况",
			list:     []*User{{ID: 1, Name: "Alice"}, {ID: 2, Name: "Bob"}},
			fieldName: "ID",
			want:      []int32{1, 2},
		},
		{
			name:      "空列表",
			list:     []*User{},
			fieldName: "ID",
			want:      []int32{},
		},
		{
			name:      "nil列表",
			list:     nil,
			fieldName: "ID",
			wantPanic: true,
		},
		{
			name:      "字段不存在",
			list:     []*User{{ID: 1, Name: "Alice"}},
			fieldName: "Age",
			wantPanic: true,
		},
		{
			name:      "非结构体列表",
			list:     []int{1, 2, 3},
			fieldName: "ID",
			wantPanic: true,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			if tt.wantPanic {
				assert.Panics(t, func() {
					PluckInt32(tt.list, tt.fieldName)
				})
				return
			}

			got := PluckInt32(tt.list, tt.fieldName)
			assert.Equal(t, tt.want, got)
		})
	}
}

// TestPluckUint32 测试 PluckUint32 函数
func TestPluckUint32(t *testing.T) {
	type Product struct {
		ID    uint32
		Name  string
		Price float64
	}

	tests := []struct {
		name      string
		list     interface{}
		fieldName string
		want      []uint32
		wantPanic bool
	}{
		{
			name:      "正常情况",
			list:     []*Product{{ID: 100, Name: "Phone"}, {ID: 200, Name: "Laptop"}},
			fieldName: "ID",
			want:      []uint32{100, 200},
		},
		{
			name:      "空列表",
			list:     []*Product{},
			fieldName: "ID",
			want:      []uint32{},
		},
		{
			name:      "nil列表",
			list:     nil,
			fieldName: "ID",
			wantPanic: true,
		},
		{
			name:      "字段不存在",
			list:     []*Product{{ID: 100, Name: "Phone"}},
			fieldName: "Code",
			wantPanic: true,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			if tt.wantPanic {
				assert.Panics(t, func() {
					PluckUint32(tt.list, tt.fieldName)
				})
				return
			}

			got := PluckUint32(tt.list, tt.fieldName)
			assert.Equal(t, tt.want, got)
		})
	}
}

// TestPluckInt64 测试 PluckInt64 函数
func TestPluckInt64(t *testing.T) {
	type Order struct {
		ID        int64
		UserID    int64
		ProductID string
	}

	tests := []struct {
		name      string
		list     interface{}
		fieldName string
		want      []int64
		wantPanic bool
	}{
		{
			name:      "正常情况",
			list:     []*Order{{ID: 1001, UserID: 2001}, {ID: 1002, UserID: 2002}},
			fieldName: "ID",
			want:      []int64{1001, 1002},
		},
		{
			name:      "空列表",
			list:     []*Order{},
			fieldName: "ID",
			want:      []int64{},
		},
		{
			name:      "nil列表",
			list:     nil,
			fieldName: "ID",
			wantPanic: true,
		},
		{
			name:      "字段不存在",
			list:     []*Order{{ID: 1001, UserID: 2001}},
			fieldName: "CreatedAt",
			wantPanic: true,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			if tt.wantPanic {
				assert.Panics(t, func() {
					PluckInt64(tt.list, tt.fieldName)
				})
				return
			}

			got := PluckInt64(tt.list, tt.fieldName)
			assert.Equal(t, tt.want, got)
		})
	}
}

// TestPluckUint64 测试 PluckUint64 函数
func TestPluckUint64(t *testing.T) {
	type Transaction struct {
		ID     uint64
		Amount uint64
		Status string
	}

	tests := []struct {
		name      string
		list     interface{}
		fieldName string
		want      []uint64
		wantPanic bool
	}{
		{
			name:      "正常情况",
			list:     []*Transaction{{ID: 9007199254740991, Amount: 1000}, {ID: 9007199254740992, Amount: 2000}},
			fieldName: "ID",
			want:      []uint64{9007199254740991, 9007199254740992},
		},
		{
			name:      "空列表",
			list:     []*Transaction{},
			fieldName: "ID",
			want:      []uint64{},
		},
		{
			name:      "nil列表",
			list:     nil,
			fieldName: "ID",
			wantPanic: true,
		},
		{
			name:      "字段不存在",
			list:     []*Transaction{{ID: 9007199254740991, Amount: 1000}},
			fieldName: "Timestamp",
			wantPanic: true,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			if tt.wantPanic {
				assert.Panics(t, func() {
					PluckUint64(tt.list, tt.fieldName)
				})
				return
			}

			got := PluckUint64(tt.list, tt.fieldName)
			assert.Equal(t, tt.want, got)
		})
	}
}