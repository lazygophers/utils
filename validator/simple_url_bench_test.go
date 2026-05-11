package validator

import (
	"reflect"
	"testing"
)

type mockURLFieldLevel struct {
	s string
}

func (m *mockURLFieldLevel) Field() reflect.Value {
	return reflect.ValueOf(m.s)
}

func (m *mockURLFieldLevel) FieldName() string {
	return ""
}

func (m *mockURLFieldLevel) StructFieldName() string {
	return ""
}

func (m *mockURLFieldLevel) Param() string {
	return ""
}

func (m *mockURLFieldLevel) GetTag(key string) string {
	return ""
}

func (m *mockURLFieldLevel) GetFieldByName(name string) reflect.Value {
	return reflect.Value{}
}

func (m *mockURLFieldLevel) Top() reflect.Value {
	return reflect.Value{}
}

func (m *mockURLFieldLevel) Parent() reflect.Value {
	return reflect.Value{}
}

func BenchmarkURL(b *testing.B) {
	urls := []string{
		"http://example.com",
		"https://example.com",
		"ftp://example.com",
		"ws://example.com",
		"wss://example.com",
		"invalid",
		"",
	}
	
	fn := URL()
	for i := 0; i < b.N; i++ {
		for _, u := range urls {
			_ = fn(&mockURLFieldLevel{s: u})
		}
	}
}
