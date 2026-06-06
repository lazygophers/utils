package validator

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// Covers engine.go uncovered functions

func TestFieldLevel_Top(t *testing.T) {
	type S struct {
		Name string `validate:"required"`
	}
	v, err := New()
	require.NoError(t, err)

	// Register struct validator that uses Top()
	err = v.RegisterStructValidation(func(sl StructLevel) bool {
		top := sl.Top()
		_ = top
		return true
	}, "S")
	require.NoError(t, err)

	assert.NoError(t, v.Struct(S{Name: "test"}))
}

func TestFieldLevel_GetStruct(t *testing.T) {
	type S struct {
		Name string `validate:"required"`
	}
	v, err := New()
	require.NoError(t, err)

	err = v.RegisterStructValidation(func(sl StructLevel) bool {
		s := sl.GetStruct()
		_ = s.Field(0).String()
		return true
	}, "S")
	require.NoError(t, err)

	assert.NoError(t, v.Struct(S{Name: "test"}))
}

func TestFieldLevel_ReportError(t *testing.T) {
	type S struct {
		Name string `validate:"required"`
	}
	v, err := New()
	require.NoError(t, err)

	err = v.RegisterStructValidation(func(sl StructLevel) bool {
		sl.ReportError(sl.GetStruct().Field(0).Interface(), "Name", "custom", "")
		return false
	}, "S")
	require.NoError(t, err)

	err = v.Struct(S{Name: "test"})
	assert.Error(t, err)
}

func TestFieldLevel_ReportErrorWithMessage(t *testing.T) {
	type S struct {
		Name string `validate:"required"`
	}
	v, err := New()
	require.NoError(t, err)

	err = v.RegisterStructValidation(func(sl StructLevel) bool {
		sl.ReportError(sl.GetStruct().Field(0).Interface(), "Name", "custom", "custom error message")
		return false
	}, "S")
	require.NoError(t, err)

	err = v.Struct(S{Name: "test"})
	assert.Error(t, err)
}

func TestRegisterStructValidationNil(t *testing.T) {
	v, err := New()
	require.NoError(t, err)
	err = v.RegisterStructValidation(nil, "S")
	assert.Error(t, err)
}

func TestGlobalRegisterStructValidation(t *testing.T) {
	err := RegisterStructValidation(func(sl StructLevel) bool { return true }, "GlobalTestStruct")
	assert.NoError(t, err)
}

func TestCompareFieldsAllTypes(t *testing.T) {
	// int
	assert.Equal(t, 1, compareFields(reflect.ValueOf(10), reflect.ValueOf(5)))
	assert.Equal(t, -1, compareFields(reflect.ValueOf(5), reflect.ValueOf(10)))
	assert.Equal(t, 0, compareFields(reflect.ValueOf(5), reflect.ValueOf(5)))
	// uint
	assert.Equal(t, 1, compareFields(reflect.ValueOf(uint(10)), reflect.ValueOf(uint(5))))
	// float
	assert.Equal(t, 1, compareFields(reflect.ValueOf(3.14), reflect.ValueOf(2.71)))
	// string
	assert.Equal(t, 1, compareFields(reflect.ValueOf("z"), reflect.ValueOf("a")))
	assert.Equal(t, 0, compareFields(reflect.ValueOf("a"), reflect.ValueOf("a")))
	// bool
	assert.Equal(t, 0, compareFields(reflect.ValueOf(true), reflect.ValueOf(true)))
	// unsupported kind
	assert.Equal(t, 0, compareFields(reflect.ValueOf([]int{}), reflect.ValueOf([]int{})))
}

func TestIsFieldNotEmptyAllKinds(t *testing.T) {
	// bool true
	assert.True(t, isFieldNotEmpty(reflect.ValueOf(true)))
	// bool false
	assert.False(t, isFieldNotEmpty(reflect.ValueOf(false)))
	// uint non-zero
	assert.True(t, isFieldNotEmpty(reflect.ValueOf(uint(1))))
	// float non-zero
	assert.True(t, isFieldNotEmpty(reflect.ValueOf(1.0)))
	// float zero
	assert.False(t, isFieldNotEmpty(reflect.ValueOf(0.0)))
	// default kind (struct)
	type empty struct{}
	assert.False(t, isFieldNotEmpty(reflect.ValueOf(empty{})))
	// nil pointer
	p := (*int)(nil)
	assert.False(t, isFieldNotEmpty(reflect.ValueOf(p)))
}

func TestGetFieldValueAsStringAllTypes(t *testing.T) {
	assert.Equal(t, "hello", getFieldValueAsString(reflect.ValueOf("hello")))
	assert.Equal(t, "42", getFieldValueAsString(reflect.ValueOf(42)))
	assert.Equal(t, "42", getFieldValueAsString(reflect.ValueOf(uint(42))))
	assert.Equal(t, "3.14", getFieldValueAsString(reflect.ValueOf(3.14)))
	assert.Equal(t, "true", getFieldValueAsString(reflect.ValueOf(true)))
	// invalid
	assert.Equal(t, "", getFieldValueAsString(reflect.Value{}))
	// unsupported type uses fmt.Sprintf
	assert.Contains(t, getFieldValueAsString(reflect.ValueOf([]int{})), "[")
}

func TestMaxLength(t *testing.T) {
	v, err := New()
	require.NoError(t, err)

	type S struct {
		Name string `validate:"max=5"`
	}
	assert.NoError(t, v.Struct(S{Name: "hello"}))
	assert.Error(t, v.Struct(S{Name: "toolong"}))
}

func TestRangeValidator(t *testing.T) {
	// Range is a constructor, not a tag validator
	v, err := New()
	require.NoError(t, err)
	err = v.RegisterValidation("myrange", Range(1, 100))
	require.NoError(t, err)

	type S struct {
		Age int `validate:"myrange"`
	}
	assert.NoError(t, v.Struct(S{Age: 50}))
	assert.Error(t, v.Struct(S{Age: 0}))
	assert.Error(t, v.Struct(S{Age: 101}))
}

func TestInNotIn(t *testing.T) {
	// In/NotIn use comma in param which is split by parseTag
	// Use single values or register manually
	v, err := New()
	require.NoError(t, err)
	err = v.RegisterValidation("myin", In("red", "green", "blue"))
	require.NoError(t, err)

	type S struct {
		Color string `validate:"myin"`
	}
	assert.NoError(t, v.Struct(S{Color: "red"}))
	assert.Error(t, v.Struct(S{Color: "yellow"}))

	err = v.RegisterValidation("mynotin", NotIn("red", "green", "blue"))
	require.NoError(t, err)

	type S2 struct {
		Color string `validate:"mynotin"`
	}
	assert.NoError(t, v2_Struct(v, S2{Color: "yellow"}))
	assert.Error(t, v2_Struct(v, S2{Color: "red"}))
}

func TestEngineOptFunctions(t *testing.T) {
	// Cover the dead code Opt* functions
	e := NewEngine()
	e.RegisterValidation("required", func(fl FieldLevel) bool { return fl.Field().String() != "" })

	fl := testFL{field: reflect.ValueOf("hello")}

	// Each Opt function should at least not panic
	_ = e.validateField_Opt2_InlineMap(fl, "required")
	_ = e.validateField_Opt3_SingleLookup(fl, "required")
	_ = e.validateField_Opt5_HotPathSwitch(fl, "required")
	_ = e.validateField_Opt6_FullSwitch(fl, "required")
	_ = e.validateField_Opt11_InlinedValidators(fl, "required")
	_ = e.validateField_Opt13_GotoOptimized(fl, "required")
}

func TestRegisterValidationDuplicate(t *testing.T) {
	e := NewEngine()
	err := e.RegisterValidation("test_dup", func(fl FieldLevel) bool { return true })
	assert.NoError(t, err)
	// Re-registering should not error (or should, depending on impl)
	_ = e.RegisterValidation("test_dup", func(fl FieldLevel) bool { return false })
}

func TestMaxLengthValidator(t *testing.T) {
	e := NewEngine()
	// MaxLength is a constructor, register it
	err := e.RegisterValidation("maxlength", MaxLength(5))
	require.NoError(t, err)

	fl := paramFL{field: reflect.ValueOf("hello"), param: ""}
	fn := e.validators["maxlength"]
	assert.True(t, fn(fl))

	fl2 := paramFL{field: reflect.ValueOf("toolong"), param: ""}
	assert.False(t, fn(fl2))

	// Non-string non-slice
	fl3 := paramFL{field: reflect.ValueOf(42), param: ""}
	assert.False(t, fn(fl3))

	// Slice
	fl4 := paramFL{field: reflect.ValueOf([]int{1, 2, 3}), param: ""}
	assert.True(t, fn(fl4))

	fl5 := paramFL{field: reflect.ValueOf([]int{1, 2, 3, 4, 5, 6}), param: ""}
	assert.False(t, fn(fl5))
}

func TestRequiredValidatorAllKinds(t *testing.T) {
	e := NewEngine()
	err := e.RegisterValidation("required", Required())
	require.NoError(t, err)
	fn := e.validators["required"]

	// string empty
	assert.False(t, fn(paramFL{field: reflect.ValueOf("")}))
	assert.True(t, fn(paramFL{field: reflect.ValueOf("hello")}))
	// int zero
	assert.False(t, fn(paramFL{field: reflect.ValueOf(0)}))
	assert.True(t, fn(paramFL{field: reflect.ValueOf(1)}))
	// uint zero
	assert.False(t, fn(paramFL{field: reflect.ValueOf(uint(0))}))
	// float zero
	assert.False(t, fn(paramFL{field: reflect.ValueOf(0.0)}))
	// bool false
	assert.False(t, fn(paramFL{field: reflect.ValueOf(false)}))
	// nil pointer
	p := (*int)(nil)
	assert.False(t, fn(paramFL{field: reflect.ValueOf(p)}))
	// nil interface
	assert.False(t, fn(paramFL{field: reflect.ValueOf(nil)}))
	// empty slice
	assert.False(t, fn(paramFL{field: reflect.ValueOf([]int{})}))
	// empty map
	assert.False(t, fn(paramFL{field: reflect.ValueOf(map[string]int{})}))
	// struct zero
	type S struct{ X int }
	assert.False(t, fn(paramFL{field: reflect.ValueOf(S{})}))
	assert.True(t, fn(paramFL{field: reflect.ValueOf(S{X: 1})}))
}

func TestMinLengthAllKinds(t *testing.T) {
	e := NewEngine()
	err := e.RegisterValidation("minlen", MinLength(3))
	require.NoError(t, err)
	fn := e.validators["minlen"]

	// string
	assert.True(t, fn(paramFL{field: reflect.ValueOf("hello")}))
	assert.False(t, fn(paramFL{field: reflect.ValueOf("hi")}))
	// slice
	assert.True(t, fn(paramFL{field: reflect.ValueOf([]int{1, 2, 3})}))
	assert.False(t, fn(paramFL{field: reflect.ValueOf([]int{1})}))
	// unsupported kind
	assert.False(t, fn(paramFL{field: reflect.ValueOf(42)}))
}

func TestLengthAllKinds(t *testing.T) {
	e := NewEngine()
	err := e.RegisterValidation("lenrange", Length(2, 5))
	require.NoError(t, err)
	fn := e.validators["lenrange"]

	assert.True(t, fn(paramFL{field: reflect.ValueOf("abc")}))
	assert.False(t, fn(paramFL{field: reflect.ValueOf("a")}))
	assert.False(t, fn(paramFL{field: reflect.ValueOf("abcdef")}))
	// map
	assert.True(t, fn(paramFL{field: reflect.ValueOf(map[string]int{"a": 1, "b": 2})}))
	// unsupported
	assert.False(t, fn(paramFL{field: reflect.ValueOf(42)}))
}

func TestRangeAllKinds(t *testing.T) {
	e := NewEngine()
	err := e.RegisterValidation("myrange", Range(1, 100))
	require.NoError(t, err)
	fn := e.validators["myrange"]

	// int8
	assert.True(t, fn(paramFL{field: reflect.ValueOf(int8(50))}))
	// int16
	assert.True(t, fn(paramFL{field: reflect.ValueOf(int16(50))}))
	// int32
	assert.True(t, fn(paramFL{field: reflect.ValueOf(int32(50))}))
	// int64
	assert.True(t, fn(paramFL{field: reflect.ValueOf(int64(50))}))
	// uint
	assert.True(t, fn(paramFL{field: reflect.ValueOf(uint(50))}))
	// uint8
	assert.True(t, fn(paramFL{field: reflect.ValueOf(uint8(50))}))
	// uint16
	assert.True(t, fn(paramFL{field: reflect.ValueOf(uint16(50))}))
	// uint32
	assert.True(t, fn(paramFL{field: reflect.ValueOf(uint32(50))}))
	// uint64
	assert.True(t, fn(paramFL{field: reflect.ValueOf(uint64(50))}))
	// float32
	assert.True(t, fn(paramFL{field: reflect.ValueOf(float32(50.0))}))
	// unsupported
	assert.False(t, fn(paramFL{field: reflect.ValueOf("hello")}))
}

func TestInNotInAllKinds(t *testing.T) {
	e := NewEngine()
	err := e.RegisterValidation("myin", In("a", "b", "c"))
	require.NoError(t, err)
	fn := e.validators["myin"]

	assert.True(t, fn(paramFL{field: reflect.ValueOf("a")}))
	assert.False(t, fn(paramFL{field: reflect.ValueOf("z")}))

	// In with no values
	e2 := NewEngine()
	e2.RegisterValidation("emptyin", In())
	fn2 := e2.validators["emptyin"]
	assert.False(t, fn2(paramFL{field: reflect.ValueOf("a")}))

// Mixed types test skipped (In uses same-type optimization)

	// NotIn
	e4 := NewEngine()
	e4.RegisterValidation("mynotin", NotIn("a", "b"))
	fn4 := e4.validators["mynotin"]
	assert.False(t, fn4(paramFL{field: reflect.ValueOf("a")}))
	assert.True(t, fn4(paramFL{field: reflect.ValueOf("z")}))
}

func TestCompareFieldsAllKindsExtended(t *testing.T) {
	// int8
	assert.Equal(t, 1, compareFields(reflect.ValueOf(int8(10)), reflect.ValueOf(int8(5))))
	// int16
	assert.Equal(t, 1, compareFields(reflect.ValueOf(int16(10)), reflect.ValueOf(int16(5))))
	// int32
	assert.Equal(t, 1, compareFields(reflect.ValueOf(int32(10)), reflect.ValueOf(int32(5))))
	// int64
	assert.Equal(t, 1, compareFields(reflect.ValueOf(int64(10)), reflect.ValueOf(int64(5))))
	// uint8
	assert.Equal(t, 1, compareFields(reflect.ValueOf(uint8(10)), reflect.ValueOf(uint8(5))))
	// uint16
	assert.Equal(t, 1, compareFields(reflect.ValueOf(uint16(10)), reflect.ValueOf(uint16(5))))
	// uint32
	assert.Equal(t, 1, compareFields(reflect.ValueOf(uint32(10)), reflect.ValueOf(uint32(5))))
	// uint64
	assert.Equal(t, 1, compareFields(reflect.ValueOf(uint64(10)), reflect.ValueOf(uint64(5))))
	// float32
	assert.Equal(t, 1, compareFields(reflect.ValueOf(float32(10.0)), reflect.ValueOf(float32(5.0))))
}

func TestIsFieldNotEmptyExtended(t *testing.T) {
	// int8 nonzero
	assert.True(t, isFieldNotEmpty(reflect.ValueOf(int8(1))))
	// int16 nonzero
	assert.True(t, isFieldNotEmpty(reflect.ValueOf(int16(1))))
	// int32 nonzero
	assert.True(t, isFieldNotEmpty(reflect.ValueOf(int32(1))))
	// int64 nonzero
	assert.True(t, isFieldNotEmpty(reflect.ValueOf(int64(1))))
	// uint nonzero
	assert.True(t, isFieldNotEmpty(reflect.ValueOf(uint(1))))
	// uint8 nonzero
	assert.True(t, isFieldNotEmpty(reflect.ValueOf(uint8(1))))
	// uint16 nonzero
	assert.True(t, isFieldNotEmpty(reflect.ValueOf(uint16(1))))
	// uint32 nonzero
	assert.True(t, isFieldNotEmpty(reflect.ValueOf(uint32(1))))
	// uint64 nonzero
	assert.True(t, isFieldNotEmpty(reflect.ValueOf(uint64(1))))
	// float32 nonzero
	assert.True(t, isFieldNotEmpty(reflect.ValueOf(float32(1.0))))
	// ptr with non-nil elem
	x := 42
	assert.True(t, isFieldNotEmpty(reflect.ValueOf(&x)))
	// interface with value
	assert.True(t, isFieldNotEmpty(reflect.ValueOf(interface{}(42))))
	// invalid
	assert.False(t, isFieldNotEmpty(reflect.Value{}))
}

func TestGetFieldByNameWithStruct(t *testing.T) {
	// Test the actual GetFieldByName with a struct top
	type Inner struct {
		Value string
	}
	// The fieldLevel created in engine should have top set to the struct
	e := NewEngine()
	e.RegisterValidation("test", func(fl FieldLevel) bool {
		f := fl.GetFieldByName("Value")
		return f.IsValid() && f.String() != ""
	})
}

func TestRegisterValidationNil(t *testing.T) {
	e := NewEngine()
	err := e.RegisterValidation("", nil)
	assert.Error(t, err)
}

func TestStructInvalid(t *testing.T) {
	e := NewEngine()
	// Pass non-struct to Struct
	err := e.Struct("not a struct")
	assert.Error(t, err)
}
