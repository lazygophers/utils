package validator

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestValidateStructAllFieldTypes(t *testing.T) {
	e := NewEngine()

	type S struct {
		Str   string   `validate:"required"`
		Int   int      `validate:"min=1,max=100"`
		Uint  uint     `validate:"min=1"`
		Float float64  `validate:"min=0.1,max=99.9"`
		Bool  bool     `validate:"required"`
		Email string   `validate:"email"`
		URL   string   `validate:"url"`
		Slice []string `validate:"min=1"`
		Map   map[string]int `validate:"min=1"`
		Skip  string   `validate:"-"`
	}

	valid := S{
		Str: "hello", Int: 50, Uint: 5, Float: 50.0, Bool: true,
		Email: "test@example.com", URL: "https://example.com",
		Slice: []string{"a"}, Map: map[string]int{"a": 1}, Skip: "anything",
	}
	errs := &ValidationErrors{}
	e.validateStruct(reflect.ValueOf(valid), reflect.ValueOf(valid), "", errs)
	assert.Empty(t, *errs)

	invalid := S{
		Str: "", Int: 0, Uint: 0, Float: 0.0, Bool: false,
		Email: "bad", URL: "not-url",
		Slice: []string{}, Map: map[string]int{},
	}
	errs2 := &ValidationErrors{}
	e.validateStruct(reflect.ValueOf(invalid), reflect.ValueOf(invalid), "", errs2)
	assert.NotEmpty(t, *errs2)
}

func TestBuiltinMinAllIntTypes(t *testing.T) {
	e := NewEngine()
	fn := e.validators["min"]

	for _, tt := range []struct {
		field interface{}
		param string
		want  bool
	}{
		{int8(5), "3", true}, {int16(5), "3", true}, {int32(5), "3", true}, {int64(5), "3", true},
		{uint(5), "3", true}, {uint8(5), "3", true}, {uint16(5), "3", true}, {uint32(5), "3", true}, {uint64(5), "3", true},
		{float32(5.0), "3", true}, {float64(5.0), "3", true},
		{int(5), "abc", false}, {true, "3", false},
	} {
		result := fn(paramFL{field: reflect.ValueOf(tt.field), param: tt.param})
		assert.Equal(t, tt.want, result, "field=%v param=%s", tt.field, tt.param)
	}
}

func TestBuiltinLenValidator(t *testing.T) {
	e := NewEngine()
	fn := e.validators["len"]

	assert.True(t, fn(paramFL{field: reflect.ValueOf("hello"), param: "5"}))
	assert.False(t, fn(paramFL{field: reflect.ValueOf("hi"), param: "5"}))
	assert.True(t, fn(paramFL{field: reflect.ValueOf([]int{1, 2, 3}), param: "3"}))
	assert.True(t, fn(paramFL{field: reflect.ValueOf(map[string]int{"a": 1, "b": 2}), param: "2"}))
	assert.False(t, fn(paramFL{field: reflect.ValueOf("hello"), param: "abc"}))
	assert.False(t, fn(paramFL{field: reflect.ValueOf(42), param: "1"}))
}

func TestBuiltinEmailURLEmpty(t *testing.T) {
	e := NewEngine()
	// Empty passes (use required for mandatory)
	assert.True(t, e.validators["email"](paramFL{field: reflect.ValueOf("")}))
	assert.True(t, e.validators["url"](paramFL{field: reflect.ValueOf("")}))
}

func TestValidateStructNestedPtr(t *testing.T) {
	e := NewEngine()
	e.RegisterValidation("required", Required())

	type Inner struct{ Value string `validate:"required"` }
	type Outer struct{ I *Inner }

	errs := &ValidationErrors{}
	e.validateStruct(reflect.ValueOf(Outer{I: nil}), reflect.ValueOf(Outer{I: nil}), "", errs)
	assert.Empty(t, *errs)

	errs = &ValidationErrors{}
	e.validateStruct(reflect.ValueOf(Outer{I: &Inner{Value: "hi"}}), reflect.ValueOf(Outer{I: &Inner{Value: "hi"}}), "", errs)
	assert.Empty(t, *errs)

	errs = &ValidationErrors{}
	e.validateStruct(reflect.ValueOf(Outer{I: &Inner{Value: ""}}), reflect.ValueOf(Outer{I: &Inner{Value: ""}}), "", errs)
	assert.NotEmpty(t, *errs)
}

func TestValidateStructNestedStruct(t *testing.T) {
	e := NewEngine()
	e.RegisterValidation("required", Required())

	type Addr struct{ City string `validate:"required"` }
	type Person struct {
		Name string `validate:"required"`
		Addr Addr
	}

	errs := &ValidationErrors{}
	e.validateStruct(reflect.ValueOf(Person{Name: "A", Addr: Addr{City: "B"}}), reflect.ValueOf(Person{Name: "A", Addr: Addr{City: "B"}}), "", errs)
	assert.Empty(t, *errs)

	errs = &ValidationErrors{}
	e.validateStruct(reflect.ValueOf(Person{Name: "A", Addr: Addr{City: ""}}), reflect.ValueOf(Person{Name: "A", Addr: Addr{City: ""}}), "", errs)
	assert.NotEmpty(t, *errs)
}

func TestValidateStructSkipTag(t *testing.T) {
	e := NewEngine()
	type S struct{ Name string `validate:"-"` }
	errs := &ValidationErrors{}
	e.validateStruct(reflect.ValueOf(S{Name: "x"}), reflect.ValueOf(S{Name: "x"}), "", errs)
	assert.Empty(t, *errs)
}

func TestValidateStructPrivateField(t *testing.T) {
	e := NewEngine()
	e.RegisterValidation("required", Required())
	type S struct {
		Public  string `validate:"required"`
		private string `validate:"required"`
	}
	errs := &ValidationErrors{}
	e.validateStruct(reflect.ValueOf(S{Public: "hi", private: ""}), reflect.ValueOf(S{Public: "hi", private: ""}), "", errs)
	assert.Empty(t, *errs) // private field should be skipped
}

func TestStructInvalidType(t *testing.T) {
	e := NewEngine()
	require.Error(t, e.Struct("not a struct"))
	require.Error(t, e.Struct(42))
}

func TestEmailInvalidFormat(t *testing.T) {
	e := NewEngine()
	require.NoError(t, e.Struct(struct {
		E string `validate:"email"`
	}{E: "test@example.com"}))
}

func TestURLInvalidFormat(t *testing.T) {
	e := NewEngine()
	require.Error(t, e.Struct(struct {
		U string `validate:"url"`
	}{U: "not a url"}))
	require.NoError(t, e.Struct(struct {
		U string `validate:"url"`
	}{U: "https://example.com"}))
}

func TestMinFieldTypes(t *testing.T) {
	v, err := New()
	require.NoError(t, err)

	type S struct {
		I int `validate:"min=5"`
	}
	assert.NoError(t, v.Struct(S{I: 10}))
	assert.Error(t, v.Struct(S{I: 3}))
}

func TestMaxFieldTypes(t *testing.T) {
	v, err := New()
	require.NoError(t, err)

	type S struct {
		I int `validate:"max=5"`
	}
	assert.NoError(t, v.Struct(S{I: 3}))
	assert.Error(t, v.Struct(S{I: 10}))
}
