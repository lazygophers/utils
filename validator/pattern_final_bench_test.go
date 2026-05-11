package validator

import (
	"reflect"
	"testing"
)

// 专门的性能基准测试
func BenchmarkPattern_Email_Valid(b *testing.B) {
	pattern := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
	validator := Pattern(pattern)
	fl := &testFieldLevel{value: reflect.ValueOf("test@example.com")}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		validator(fl)
	}
}

func BenchmarkPattern_Email_Invalid(b *testing.B) {
	pattern := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
	validator := Pattern(pattern)
	fl := &testFieldLevel{value: reflect.ValueOf("invalid")}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		validator(fl)
	}
}

func BenchmarkPattern_FixedLength_Valid(b *testing.B) {
	pattern := `^\d{5}$`
	validator := Pattern(pattern)
	fl := &testFieldLevel{value: reflect.ValueOf("12345")}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		validator(fl)
	}
}

func BenchmarkPattern_Literal_Valid(b *testing.B) {
	pattern := `^hello$`
	validator := Pattern(pattern)
	fl := &testFieldLevel{value: reflect.ValueOf("hello")}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		validator(fl)
	}
}
