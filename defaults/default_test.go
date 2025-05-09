package defaults_test

import (
	"github.com/lazygophers/utils/defaults"
	"testing"
)

type A struct {
	Name string `default:"name"`
}

type B struct {
	Value int `default:"100"`
}

type C struct {
	FloatValue float64 `default:"3.14"`
}

type D struct {
	BoolValue bool `default:"true"`
}

type E struct {
	PtrValue *int `default:"5"`
}

type F struct {
	StructValue A `default:""`
}

type G struct {
	MapValue map[string]string `default:""`
}

type H struct {
	InterfaceValue interface{} `default:""`
}

type I struct {
	SliceValue []string `default:""`
}

func TestStruct(t *testing.T) {
	var s A
	defaults.SetDefaults(&s)
	if s.Name != "name" {
		t.Errorf("Expected Name to be 'name', got '%s'", s.Name)
	}
}

func TestInt(t *testing.T) {
	var b B
	defaults.SetDefaults(&b)
	if b.Value != 100 {
		t.Errorf("Expected Value to be 100, got %d", b.Value)
	}
}

func TestFloat(t *testing.T) {
	var c C
	defaults.SetDefaults(&c)
	if c.FloatValue != 3.14 {
		t.Errorf("Expected FloatValue to be 3.14, got %f", c.FloatValue)
	}
}

func TestBool(t *testing.T) {
	var d D
	defaults.SetDefaults(&d)
	if !d.BoolValue {
		t.Errorf("Expected BoolValue to be true, got %v", d.BoolValue)
	}
}

func TestPointer(t *testing.T) {
	var e E
	defaults.SetDefaults(&e)
	if e.PtrValue == nil || *e.PtrValue != 5 {
		t.Errorf("Expected PtrValue to be 5, got %v", e.PtrValue)
	}

	// 验证指针类型初始化逻辑
	type NestedPointer struct {
		Ptr *int `default:"10"`
	}
	var np NestedPointer
	defaults.SetDefaults(&np)
	if np.Ptr == nil || *np.Ptr != 10 {
		t.Errorf("Expected Ptr to be 10, got %v", np.Ptr)
	}

	// 新增测试：验证嵌套指针的初始化逻辑
	type DoubleNestedPointer struct {
		Ptr *NestedPointer `default:""`
	}
	var dnp DoubleNestedPointer
	defaults.SetDefaults(&dnp)
	if dnp.Ptr == nil || dnp.Ptr.Ptr == nil || *dnp.Ptr.Ptr != 10 {
		t.Errorf("Expected DoubleNestedPointer.Ptr.Ptr to be 10, got %v", dnp.Ptr.Ptr)
	}
}

func TestPointerNil(t *testing.T) {
	var ptr *int
	defaults.SetDefaults(&ptr)
	if ptr == nil {
		t.Errorf("Expected ptr to be initialized, got nil")
	}
}

func TestStructWithNilFields(t *testing.T) {
	type Nested struct {
		Field *int `default:"10"`
	}
	var s struct {
		Nested *Nested `default:""`
	}
	defaults.SetDefaults(&s)
	if s.Nested == nil || s.Nested.Field == nil || *s.Nested.Field != 10 {
		t.Errorf("Expected Nested.Field to be 10, got %v", s.Nested.Field)
	}
}

func TestNestedStruct(t *testing.T) {
	var f F
	defaults.SetDefaults(&f)
	if f.StructValue.Name != "name" {
		t.Errorf("Expected StructValue.Name to be 'name', got '%s'", f.StructValue.Name)
	}
}

func TestInvalidDefault(t *testing.T) {
	type Invalid struct {
		Value int `default:"invalid"`
	}
	var i Invalid
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("Expected panic for invalid default value, but got no panic")
		}
	}()
	defaults.SetDefaults(&i)
}

func TestZeroValue(t *testing.T) {
	type Zero struct {
		Value int `default:"0"`
	}
	var z Zero
	defaults.SetDefaults(&z)
	if z.Value != 0 {
		t.Errorf("Expected Value to be 0, got %d", z.Value)
	}
}

func TestPanicOnInvalidKind(t *testing.T) {
	type InvalidKind struct {
		Field chan int `default:""`
	}
	var i InvalidKind
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("Expected panic for invalid kind, but got no panic")
		}
	}()
	defaults.SetDefaults(&i)
}

func TestPanicOnNilStruct(t *testing.T) {
	var s *struct{}
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("Expected panic for nil struct, but got no panic")
		}
	}()
	defaults.SetDefaults(s)
}

func TestPanicOnInvalidDefault(t *testing.T) {
	type InvalidDefault struct {
		Field int `default:"invalid"`
	}
	var i InvalidDefault
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("Expected panic for invalid default value, but got no panic")
		}
	}()
	defaults.SetDefaults(&i)
}

func TestPanicOnInvalidUint(t *testing.T) {
	type InvalidUint struct {
		Field uint `default:"invalid"`
	}
	var i InvalidUint
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("Expected panic for invalid uint default value, but got no panic")
		}
	}()
	defaults.SetDefaults(&i)
}

func TestPanicOnInvalidInt(t *testing.T) {
	type InvalidInt struct {
		Field int `default:"invalid"`
	}
	var i InvalidInt
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("Expected panic for invalid int default value, but got no panic")
		}
	}()
	defaults.SetDefaults(&i)
}

func TestPanicOnInvalidFloat(t *testing.T) {
	type InvalidFloat struct {
		Field float64 `default:"invalid"`
	}
	var i InvalidFloat
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("Expected panic for invalid float default value, but got no panic")
		}
	}()
	defaults.SetDefaults(&i)
}

func TestPanicOnInvalidBool(t *testing.T) {
	type InvalidBool struct {
		Field bool `default:"invalid"`
	}
	var i InvalidBool
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("Expected panic for invalid bool default value, but got no panic")
		}
	}()
	defaults.SetDefaults(&i)
}

func TestInterface(t *testing.T) {
	var h H
	defaults.SetDefaults(&h)
	if h.InterfaceValue != nil {
		t.Errorf("Expected InterfaceValue to be nil, got %v", h.InterfaceValue)
	}
}

func TestSlice(t *testing.T) {
	var i I
	defaults.SetDefaults(&i)
	if i.SliceValue == nil {
		t.Errorf("Expected SliceValue to be initialized, got nil")
	}
}

func TestNilValue(t *testing.T) {
	var value *int
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("Expected panic for nil value, but got no panic")
		}
	}()
	defaults.SetDefaults(value)
}
