package candy

import (
	"errors"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

// stringTestStruct 测试用的结构体
type stringTestStruct struct {
	ID   int
	Name string
	Age  int
}

// jsonMarshalError 用于测试 JSON 序列化错误的类型
type jsonMarshalError struct {
	Msg string
}

// Error 实现 error 接口
func (e *jsonMarshalError) Error() string {
	return e.Msg
}

// TestToString 测试 ToString 函数
func TestToString(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		input    interface{}
		expected string
	}{
		{"bool true", true, "1"},
		{"bool false", false, "0"},
		{"int", 42, "42"},
		{"int8", int8(8), "8"},
		{"int16", int16(16), "16"},
		{"int32", int32(32), "32"},
		{"int64", int64(64), "64"},
		{"uint", uint(42), "42"},
		{"uint8", uint8(8), "8"},
		{"uint16", uint16(16), "16"},
		{"uint32", uint32(32), "32"},
		{"uint64", uint64(64), "64"},
		{"float32 integer", float32(42), "42"},
		{"float32 fraction", float32(3.14), "3.140000104904175"},
		{"float64 integer", float64(42), "42"},
		{"float64 fraction", float64(3.14), "3.140000"},
		{"time.Duration", time.Second, "1s"},
		{"string", "hello", "hello"},
		{"[]byte", []byte("hello"), "hello"},
		{"nil", nil, ""},
		{"error", errors.New("test error"), "test error"},
		{"struct", stringTestStruct{ID: 1, Name: "test", Age: 25}, `{"ID":1,"Name":"test","Age":25}`},
		{"slice", []int{1, 2, 3}, "[1,2,3]"},
		{"map", map[string]int{"key": 42}, `{"key":42}`},
		{"channel", make(chan int), ""},  // channels cannot be marshaled to JSON, should return empty string
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			result := ToString(tt.input)
			assert.Equal(t, tt.expected, result, "ToString() 的结果应与期望值相等")
		})
	}
}

// BenchmarkToString 性能基准测试
func BenchmarkToString(b *testing.B) {
	b.Run("int", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = ToString(42)
		}
	})

	b.Run("string", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = ToString("hello")
		}
	})

	b.Run("[]byte", func(b *testing.B) {
		data := []byte("hello")
		for i := 0; i < b.N; i++ {
			_ = ToString(data)
		}
	})

	b.Run("float64", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = ToString(3.14)
		}
	})

	b.Run("struct", func(b *testing.B) {
		s := stringTestStruct{ID: 1, Name: "test", Age: 25}
		for i := 0; i < b.N; i++ {
			_ = ToString(s)
		}
	})
}