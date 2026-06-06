package validator

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// ===== registerBuiltinValidators paths =====

func TestBuiltinEqAllTypes(t *testing.T) {
	v, err := New()
	require.NoError(t, err)

	type S struct {
		Str   string  `validate:"eq=hello"`
		Int   int     `validate:"eq=42"`
		Uint  uint    `validate:"eq=10"`
		Float float64 `validate:"eq=3.14"`
	}
	assert.NoError(t, v.Struct(S{Str: "hello", Int: 42, Uint: 10, Float: 3.14}))
	assert.Error(t, v.Struct(S{Str: "no", Int: 1, Uint: 1, Float: 1.0}))
}

func TestBuiltinNe(t *testing.T) {
	v, err := New()
	require.NoError(t, err)
	type S struct{ V string `validate:"ne=hello"` }
	assert.NoError(t, v.Struct(S{V: "world"}))
	assert.Error(t, v.Struct(S{V: "hello"}))
}

func TestBuiltinEqField(t *testing.T) {
	v, err := New()
	require.NoError(t, err)
	type S struct {
		A string `validate:"eqfield=B"`
		B string
	}
	assert.NoError(t, v.Struct(S{A: "x", B: "x"}))
	assert.Error(t, v.Struct(S{A: "x", B: "y"}))
}

func TestBuiltinNeField(t *testing.T) {
	v, err := New()
	require.NoError(t, err)
	type S struct {
		A string `validate:"nefield=B"`
		B string
	}
	assert.NoError(t, v.Struct(S{A: "x", B: "y"}))
	assert.Error(t, v.Struct(S{A: "x", B: "x"}))
}

func TestBuiltinRequiredWith(t *testing.T) {
	v, err := New()
	require.NoError(t, err)
	type S struct {
		A string `validate:"required_with=B"`
		B string
	}
	assert.NoError(t, v.Struct(S{A: "", B: ""}))    // B empty → A not required
	assert.NoError(t, v.Struct(S{A: "x", B: "y"}))  // B has value, A has value
	assert.Error(t, v.Struct(S{A: "", B: "y"}))     // B has value, A empty → error
}

func TestBuiltinRequiredWithout(t *testing.T) {
	v, err := New()
	require.NoError(t, err)
	type S struct {
		A string `validate:"required_without=B"`
		B string
	}
	assert.Error(t, v.Struct(S{A: "", B: ""}))      // B empty → A required
	assert.NoError(t, v.Struct(S{A: "x", B: ""}))   // A has value
	assert.NoError(t, v.Struct(S{A: "", B: "y"}))   // B has value → A not required
}

func TestBuiltinRequiredIf(t *testing.T) {
	v, err := New()
	require.NoError(t, err)
	type S struct {
		Type  string
		Value string `validate:"required_if=Type=admin"`
	}
	assert.NoError(t, v.Struct(S{Type: "user", Value: ""}))    // Type != admin
	assert.Error(t, v.Struct(S{Type: "admin", Value: ""}))     // Type == admin, Value empty
	assert.NoError(t, v.Struct(S{Type: "admin", Value: "x"}))  // Type == admin, Value present
}

func TestBuiltinMinAllTypes(t *testing.T) {
	v, err := New()
	require.NoError(t, err)
	type S struct {
		Str   string  `validate:"min=3"`
		Int   int     `validate:"min=5"`
		Uint  uint    `validate:"min=5"`
		Float float64 `validate:"min=1.0"`
		Slice []int   `validate:"min=2"`
		Map   map[string]int `validate:"min=1"`
	}
	assert.NoError(t, v.Struct(S{Str: "abc", Int: 5, Uint: 5, Float: 1.0, Slice: []int{1, 2}, Map: map[string]int{"a": 1}}))
	assert.Error(t, v.Struct(S{Str: "ab", Int: 4, Uint: 4, Float: 0.5, Slice: []int{1}, Map: map[string]int{}}))
}

func TestBuiltinMaxAllTypes(t *testing.T) {
	v, err := New()
	require.NoError(t, err)
	type S struct {
		Str   string  `validate:"max=5"`
		Int   int     `validate:"max=10"`
		Uint  uint    `validate:"max=10"`
		Float float64 `validate:"max=99.9"`
		Slice []int   `validate:"max=3"`
		Map   map[string]int `validate:"max=2"`
	}
	assert.NoError(t, v.Struct(S{Str: "hello", Int: 10, Uint: 10, Float: 50.0, Slice: []int{1, 2, 3}, Map: map[string]int{"a": 1, "b": 2}}))
	assert.Error(t, v.Struct(S{Str: "toolong", Int: 11, Uint: 11, Float: 100.0, Slice: []int{1, 2, 3, 4}, Map: map[string]int{"a": 1, "b": 2, "c": 3}}))
}

func TestBuiltinLenAllTypes(t *testing.T) {
	v, err := New()
	require.NoError(t, err)
	type S struct {
		Str   string `validate:"len=5"`
		Slice []int  `validate:"len=3"`
	}
	assert.NoError(t, v.Struct(S{Str: "hello", Slice: []int{1, 2, 3}}))
	assert.Error(t, v.Struct(S{Str: "hi", Slice: []int{1}}))
}

func TestBuiltinNumeric(t *testing.T) {
	v, err := New()
	require.NoError(t, err)
	type S struct{ V string `validate:"numeric"` }
	assert.NoError(t, v.Struct(S{V: "123.45"}))
	assert.Error(t, v.Struct(S{V: "abc"}))
	assert.NoError(t, v.Struct(S{V: ""})) // empty passes
}

func TestBuiltinAlpha(t *testing.T) {
	v, err := New()
	require.NoError(t, err)
	type S struct{ V string `validate:"alpha"` }
	assert.NoError(t, v.Struct(S{V: "hello"}))
	assert.Error(t, v.Struct(S{V: "hello123"}))
}

func TestBuiltinAlphanum(t *testing.T) {
	v, err := New()
	require.NoError(t, err)
	type S struct{ V string `validate:"alphanum"` }
	assert.NoError(t, v.Struct(S{V: "hello123"}))
	assert.Error(t, v.Struct(S{V: "hello!"}))
}

// ===== compareFields pointer/interface paths =====

func TestCompareFieldsPtr(t *testing.T) {
	a := "x"
	b := "y"
	nilPtr := (*string)(nil)
	assert.Equal(t, 0, compareFields(reflect.ValueOf(nilPtr), reflect.ValueOf(nilPtr)))
	assert.Equal(t, -1, compareFields(reflect.ValueOf(nilPtr), reflect.ValueOf(&b)))
	assert.Equal(t, 1, compareFields(reflect.ValueOf(&a), reflect.ValueOf(nilPtr)))
	assert.Equal(t, 0, compareFields(reflect.ValueOf(&a), reflect.ValueOf(&a)))

	var iface1, iface2 interface{} = 42, 42
	var ifaceNil interface{} = nil
	assert.Equal(t, 0, compareFields(reflect.ValueOf(iface1), reflect.ValueOf(iface2)))
	assert.Equal(t, 0, compareFields(reflect.ValueOf(ifaceNil), reflect.ValueOf(nilPtr)))
}

func TestCompareFieldsInvalid(t *testing.T) {
	assert.Equal(t, 0, compareFields(reflect.Value{}, reflect.Value{}))
}

func TestCompareFieldsDefaultKind(t *testing.T) {
	// struct falls into default case
	type S struct{ X int }
	assert.Equal(t, 0, compareFields(reflect.ValueOf(S{X: 1}), reflect.ValueOf(S{X: 1})))
	assert.Equal(t, -1, compareFields(reflect.ValueOf(S{X: 1}), reflect.ValueOf(S{X: 2})))
}

// ===== validateStruct dive tag =====

func TestValidateStructDiveSlice(t *testing.T) {
	e := NewEngine()
	e.validators["min"] = func(fl FieldLevel) bool {
		return float64(len(fl.Field().String())) >= 2
	}

	type S struct {
		Items []string `validate:"dive=min=2"`
	}
	errs := &ValidationErrors{}
	e.validateStruct(reflect.ValueOf(S{Items: []string{"ab", "cd"}}), reflect.ValueOf(S{Items: []string{"ab", "cd"}}), "", errs)
	assert.Empty(t, *errs)

	errs = &ValidationErrors{}
	e.validateStruct(reflect.ValueOf(S{Items: []string{"a"}}), reflect.ValueOf(S{Items: []string{"a"}}), "", errs)
	assert.NotEmpty(t, *errs)
}

func TestValidateStructDiveStructSlice(t *testing.T) {
	e := NewEngine()
	e.validators["required"] = func(fl FieldLevel) bool { return fl.Field().String() != "" }

	type Inner struct{ Name string `validate:"required"` }
	type Outer struct{ Items []Inner `validate:"dive"` }

	errs := &ValidationErrors{}
	e.validateStruct(reflect.ValueOf(Outer{Items: []Inner{{Name: "a"}}}), reflect.ValueOf(Outer{Items: []Inner{{Name: "a"}}}), "", errs)
	assert.Empty(t, *errs)

	errs = &ValidationErrors{}
	e.validateStruct(reflect.ValueOf(Outer{Items: []Inner{{Name: ""}}}), reflect.ValueOf(Outer{Items: []Inner{{Name: ""}}}), "", errs)
	assert.NotEmpty(t, *errs)
}

func TestValidateStructDivePtrSlice(t *testing.T) {
	e := NewEngine()
	e.validators["required"] = func(fl FieldLevel) bool { return fl.Field().String() != "" }

	type Inner struct{ Name string `validate:"required"` }
	type Outer struct{ Items []*Inner `validate:"dive"` }

	x := Inner{Name: "a"}
	errs := &ValidationErrors{}
	e.validateStruct(reflect.ValueOf(Outer{Items: []*Inner{&x}}), reflect.ValueOf(Outer{Items: []*Inner{&x}}), "", errs)
	assert.Empty(t, *errs)

	emptyInner := Inner{Name: ""}
	errs = &ValidationErrors{}
	e.validateStruct(reflect.ValueOf(Outer{Items: []*Inner{&emptyInner}}), reflect.ValueOf(Outer{Items: []*Inner{&emptyInner}}), "", errs)
	assert.NotEmpty(t, *errs)

	// nil ptr in slice - should skip
	errs = &ValidationErrors{}
	e.validateStruct(reflect.ValueOf(Outer{Items: []*Inner{nil}}), reflect.ValueOf(Outer{Items: []*Inner{nil}}), "", errs)
	assert.Empty(t, *errs)
}

func TestValidateStructEmptyDive(t *testing.T) {
	e := NewEngine()
	type S struct{ Items []string `validate:"dive"` }
	errs := &ValidationErrors{}
	e.validateStruct(reflect.ValueOf(S{Items: []string{}}), reflect.ValueOf(S{Items: []string{}}), "", errs)
	assert.Empty(t, *errs)
}

// ===== struct validators (RegisterStructValidation) =====

func TestStructValidatorOnEngine(t *testing.T) {
	e := NewEngine()
	e.validators["required"] = func(fl FieldLevel) bool { return fl.Field().String() != "" }
	e.RegisterStructValidation(func(sl StructLevel) bool {
		s := sl.GetStruct()
		sl.ReportError(s.Interface(), "test", "custom_struct", "custom validation message")
		return false
	}, "TestStructForValidator")

	type TestStructForValidator struct{ Name string `validate:"required"` }
	err := e.Struct(TestStructForValidator{Name: "x"})
	assert.Error(t, err) // struct validator always reports error
}

// ===== Email/URL constructor functions (80%) =====

func TestEmailConstructor(t *testing.T) {
	fn := Email()
	assert.True(t, fn(paramFL{field: reflect.ValueOf("test@example.com")}))
	assert.False(t, fn(paramFL{field: reflect.ValueOf("invalid")}))
	assert.True(t, fn(paramFL{field: reflect.ValueOf("")})) // empty passes
}

func TestURLConstructor(t *testing.T) {
	fn := URL()
	assert.True(t, fn(paramFL{field: reflect.ValueOf("https://example.com")}))
	assert.False(t, fn(paramFL{field: reflect.ValueOf("not-a-url")}))
	assert.True(t, fn(paramFL{field: reflect.ValueOf("")}))
}

// ===== GetFieldByName on non-struct (66.7%) =====

func TestGetFieldByNameNonStruct(t *testing.T) {
	fl := &fieldLevel{
		top:    reflect.ValueOf("x"),
		parent: reflect.ValueOf("x"),
		field:  reflect.ValueOf("x"),
	}
	result := fl.GetFieldByName("anything")
	assert.False(t, result.IsValid())
}
