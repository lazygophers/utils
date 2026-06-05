package validator

import (
	"reflect"
	"testing"
)

// 测试用的复杂嵌套结构体
type TestAddress struct {
	Street  string `validate:"required"`
	City    string `validate:"required"`
	ZipCode string `validate:"len=5"`
}

type TestPerson struct {
	Name    string      `validate:"required"`
	Age     int         `validate:"gte=0,lte=150"`
	Email   string      `validate:"email"`
	Address TestAddress `validate:"required"`
	Tags    []string    `validate:"dive,omitempty"`
}

// 测试 validateStruct 性能
func BenchmarkValidateStruct_Optimized(b *testing.B) {
	v, _ := New()
	person := TestPerson{
		Name:  "John Doe",
		Age:   30,
		Email: "john@example.com",
		Address: TestAddress{
			Street:  "123 Main St",
			City:    "New York",
			ZipCode: "12345",
		},
		Tags: []string{"tag1", "tag2", "tag3"},
	}

	b.ResetTimer()
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		var errs ValidationErrors
		v.engine.validateStruct(reflect.ValueOf(person), reflect.ValueOf(person), "", &errs)
	}
}

// 测试直接调用 Struct 方法
func BenchmarkValidator_Struct_Optimized(b *testing.B) {
	v, _ := New()
	person := TestPerson{
		Name:  "John Doe",
		Age:   30,
		Email: "john@example.com",
		Address: TestAddress{
			Street:  "123 Main St",
			City:    "New York",
			ZipCode: "12345",
		},
		Tags: []string{"tag1", "tag2", "tag3"},
	}

	b.ResetTimer()
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		_ = v.Struct(person)
	}
}
