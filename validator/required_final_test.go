package validator

import (
	"reflect"
	"testing"
)

// 测试数据
var (
	testCases = []struct {
		name  string
		field FieldLevel
		want  bool
	}{
		{"empty_string", testFL{reflect.ValueOf("")}, false},
		{"nonempty_string", testFL{reflect.ValueOf("hello")}, true},
		{"empty_slice", testFL{reflect.ValueOf([]int{})}, false},
		{"nonempty_slice", testFL{reflect.ValueOf([]int{1})}, true},
		{"empty_map", testFL{reflect.ValueOf(map[string]int{})}, false},
		{"nonempty_map", testFL{reflect.ValueOf(map[string]int{"a": 1})}, true},
		{"nil_ptr", testFL{reflect.ValueOf((*int)(nil))}, false},
		{"nonnil_ptr", testFL{reflect.ValueOf(ptr(42))}, true},
		{"zero_int", testFL{reflect.ValueOf(0)}, false},
		{"nonzero_int", testFL{reflect.ValueOf(42)}, true},
		{"nil_interface", testFL{reflect.ValueOf(nil)}, false},
		{"nonnil_interface", testFL{reflect.ValueOf(42)}, true},
	}
)

type testFL struct{ field reflect.Value }

func (t testFL) Field() reflect.Value                            { return t.field }
func (t testFL) Top() reflect.Value                               { return reflect.Value{} }
func (t testFL) Parent() reflect.Value                            { return reflect.Value{} }
func (t testFL) FieldName() string                                { return "" }
func (t testFL) StructFieldName() string                          { return "" }
func (t testFL) Param() string                                    { return "" }
func (t testFL) GetTag(string) string                             { return "" }
func (t testFL) GetFieldByName(string) reflect.Value              { return reflect.Value{} }

func ptr(v int) *int { return &v }

// 原始实现
func reqOrig(fl FieldLevel) bool {
	field := fl.Field()
	switch field.Kind() {
	case reflect.String:
		return field.String() != ""
	case reflect.Slice, reflect.Map, reflect.Array:
		return field.Len() > 0
	case reflect.Ptr, reflect.Interface:
		return !field.IsNil()
	default:
		return field.IsValid() && !field.IsZero()
	}
}

// FastPath 优化
func reqFast(fl FieldLevel) bool {
	field := fl.Field()
	kind := field.Kind()
	if kind == reflect.String {
		return field.String() != ""
	}
	if kind == reflect.Slice || kind == reflect.Map || kind == reflect.Array {
		return field.Len() > 0
	}
	if kind == reflect.Ptr || kind == reflect.Interface {
		return !field.IsNil()
	}
	return field.IsValid() && !field.IsZero()
}

// 分离变量优化
func reqSep(fl FieldLevel) bool {
	field := fl.Field()
	kind := field.Kind()
	if kind == reflect.String {
		s := field.String()
		return s != ""
	}
	if kind == reflect.Slice || kind == reflect.Map || kind == reflect.Array {
		l := field.Len()
		return l > 0
	}
	if kind == reflect.Ptr || kind == reflect.Interface {
		return !field.IsNil()
	}
	return field.IsValid() && !field.IsZero()
}

func BenchmarkRequired(b *testing.B) {
	b.Run("Original", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			for _, tc := range testCases {
				reqOrig(tc.field)
			}
		}
	})

	b.Run("FastPath", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			for _, tc := range testCases {
				reqFast(tc.field)
			}
		}
	})

	b.Run("SeparatedVars", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			for _, tc := range testCases {
				reqSep(tc.field)
			}
		}
	})
}

func TestRequiredCorrectness(t *testing.T) {
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if got := reqOrig(tc.field); got != tc.want {
				t.Errorf("reqOrig() = %v, want %v", got, tc.want)
			}
			if got := reqFast(tc.field); got != tc.want {
				t.Errorf("reqFast() = %v, want %v", got, tc.want)
			}
			if got := reqSep(tc.field); got != tc.want {
				t.Errorf("reqSep() = %v, want %v", got, tc.want)
			}
		})
	}
}
