package validator

import (
	"reflect"
	"testing"
)

type paramFL struct {
	field reflect.Value
	param string
}

func (t paramFL) Field() reflect.Value                { return t.field }
func (t paramFL) Top() reflect.Value                  { return reflect.Value{} }
func (t paramFL) Parent() reflect.Value               { return reflect.Value{} }
func (t paramFL) FieldName() string                   { return "test" }
func (t paramFL) StructFieldName() string             { return "Test" }
func (t paramFL) Param() string                       { return t.param }
func (t paramFL) GetTag(string) string                { return "" }
func (t paramFL) GetFieldByName(string) reflect.Value { return reflect.Value{} }

func TestCompareThresholdAllTypes(t *testing.T) {
	tests := []struct {
		name  string
		field interface{}
		param string
		gt    bool
		gte   bool
		lt    bool
		lte   bool
	}{
		{"string gt", "hello", "4", true, true, false, false},
		{"string lt", "hi", "5", false, false, true, true},
		{"int gt", 10, "5", true, true, false, false},
		{"int lt", 3, "5", false, false, true, true},
		{"uint gt", uint(10), "5", true, true, false, false},
		{"float gt", 3.14, "3", true, true, false, false},
		{"slice gt", []int{1, 2, 3}, "2", true, true, false, false},
		{"map gt", map[string]int{"a": 1, "b": 2}, "1", true, true, false, false},
		{"invalid param", 10, "abc", false, false, false, false},
		{"bool kind", true, "0", false, false, false, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fl := paramFL{field: reflect.ValueOf(tt.field), param: tt.param}
			if got := validateGT(fl); got != tt.gt {
				t.Errorf("gt = %v, want %v", got, tt.gt)
			}
			if got := validateGTE(fl); got != tt.gte {
				t.Errorf("gte = %v, want %v", got, tt.gte)
			}
			if got := validateLT(fl); got != tt.lt {
				t.Errorf("lt = %v, want %v", got, tt.lt)
			}
			if got := validateLTE(fl); got != tt.lte {
				t.Errorf("lte = %v, want %v", got, tt.lte)
			}
		})
	}
}

func TestEqIgnoreCase(t *testing.T) {
	fl := paramFL{field: reflect.ValueOf("Hello"), param: "hello"}
	if !validateEqIgnoreCase(fl) {
		t.Error("expected true for Hello == hello (case insensitive)")
	}
	fl2 := paramFL{field: reflect.ValueOf("Hello"), param: ""}
	if validateEqIgnoreCase(fl2) {
		t.Error("expected false for empty param")
	}
}

func TestNeIgnoreCase(t *testing.T) {
	fl := paramFL{field: reflect.ValueOf("Hello"), param: "world"}
	if !validateNeIgnoreCase(fl) {
		t.Error("expected true for Hello != world")
	}
	fl2 := paramFL{field: reflect.ValueOf("Hello"), param: ""}
	if validateNeIgnoreCase(fl2) {
		t.Error("expected false for empty param")
	}
	fl3 := paramFL{field: reflect.ValueOf("Hello"), param: "hello"}
	if validateNeIgnoreCase(fl3) {
		t.Error("expected false for Hello == hello (case insensitive)")
	}
}
