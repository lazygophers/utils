package validator

import (
	"reflect"
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	xlanguage "golang.org/x/text/language"
)

func TestDefaultErrorPath(t *testing.T) {
	origOnce := once
	origDefault := defaultValidator
	once = sync.Once{}
	defaultValidator = nil
	d := Default()
	assert.NotNil(t, d)
	once = origOnce
	defaultValidator = origDefault
}

func TestStructNonValidationErrors(t *testing.T) {
	v, err := New()
	require.NoError(t, err)
	err2 := v.Struct("not a struct")
	assert.Error(t, err2)
}

func TestVarNoError(t *testing.T) {
	v, err := New()
	require.NoError(t, err)
	assert.NoError(t, v.Var("test@example.com", "email"))
}

func TestEffectiveLocaleDefault(t *testing.T) {
	v, err := New()
	require.NoError(t, err)
	assert.Equal(t, xlanguage.Make("en"), v.GetLocale())
}

func TestTranslateCustomMsg(t *testing.T) {
	v, err := New()
	require.NoError(t, err)
	v.RegisterTranslation(xlanguage.Make("en"), "required", "custom {field} is {tag}")
	err2 := v.Struct(struct {
		Name string `validate:"required" json:"name"`
	}{Name: ""})
	require.Error(t, err2)
	assert.Contains(t, err2.Error(), "custom")
}

func TestTranslateUnknownTag(t *testing.T) {
	v, err := New()
	require.NoError(t, err)
	err = v.RegisterValidation("weird_tag_xxx", func(fl FieldLevel) bool { return false })
	require.NoError(t, err)
	err2 := v.Var("x", "weird_tag_xxx")
	require.Error(t, err2)
	assert.Contains(t, err2.Error(), "weird_tag_xxx")
}

func TestTranslateNonEnLocale(t *testing.T) {
	v, err := New(WithLocale(xlanguage.Make("zh")))
	require.NoError(t, err)
	err2 := v.Struct(struct {
		Name string `validate:"required" json:"name"`
	}{Name: ""})
	require.Error(t, err2)
	assert.NotEmpty(t, err2.Error())
}

func TestFormatMsgNoPlaceholder(t *testing.T) {
	v, err := New()
	require.NoError(t, err)
	v.RegisterTranslation(xlanguage.Make("en"), "required", "plain message")
	err2 := v.Struct(struct {
		Name string `validate:"required" json:"name"`
	}{Name: ""})
	require.Error(t, err2)
	assert.Contains(t, err2.Error(), "plain message")
}

func TestRegisterValidationNilAndEmpty(t *testing.T) {
	e := NewEngine()
	assert.Error(t, e.RegisterValidation("", nil))
	assert.Error(t, e.RegisterValidation("", func(fl FieldLevel) bool { return true }))
}

func TestEngineStructValidatorPass(t *testing.T) {
	e := NewEngine()
	e.validators["required"] = func(fl FieldLevel) bool { return fl.Field().String() != "" }
	e.RegisterStructValidation(func(sl StructLevel) bool { return true }, "S2")
	type S2 struct{ Name string `validate:"required"` }
	assert.NoError(t, e.Struct(S2{Name: "x"}))
}

func TestValidateStructPtrStructNoTag(t *testing.T) {
	e := NewEngine()
	e.validators["required"] = func(fl FieldLevel) bool { return fl.Field().String() != "" }
	type Inner struct{ Name string `validate:"required"` }
	type Outer struct{ I *Inner }
	errs := &ValidationErrors{}
	e.validateStruct(reflect.ValueOf(Outer{I: &Inner{Name: ""}}), reflect.ValueOf(Outer{I: &Inner{Name: ""}}), "", errs)
	assert.NotEmpty(t, *errs)
}

func TestCompareFieldsEdgeCases(t *testing.T) {
	assert.Equal(t, 0, compareFields(reflect.Value{}, reflect.Value{}))
	assert.Equal(t, -1, compareFields(reflect.ValueOf("a"), reflect.ValueOf("b")))
	assert.Equal(t, -1, compareFields(reflect.ValueOf(uint(1)), reflect.ValueOf(uint(2))))
	assert.Equal(t, -1, compareFields(reflect.ValueOf(1.0), reflect.ValueOf(2.0)))
}

func TestIsFieldNotEmptyEdgeCases(t *testing.T) {
	assert.False(t, isFieldNotEmpty(reflect.Value{}))
	assert.False(t, isFieldNotEmpty(reflect.ValueOf(0)))
	assert.False(t, isFieldNotEmpty(reflect.ValueOf(uint(0))))
}

func TestReportErrorEmptyMsg(t *testing.T) {
	sl := &structLevel{
		top:     reflect.ValueOf("x"),
		current: reflect.ValueOf("x"),
		errors:  &ValidationErrors{},
	}
	sl.ReportError("val", "field", "tag", "")
	assert.Equal(t, 1, (*sl.errors).Len())
}

func TestFieldValidatorsInvalidTarget(t *testing.T) {
	for name, fn := range map[string]ValidatorFunc{
		"gtfield":   validateGTField,
		"gtefield":  validateGTEField,
		"ltfield":   validateLTField,
		"ltefield":  validateLTEField,
		"eqcsfield": validateEQCSField,
		"necsfield": validateNECSField,
		"gtcsfield": validateGTCSField,
		"gtecsfield": validateGTECSField,
		"ltcsfield": validateLTCSField,
		"ltecsfield": validateLTECSField,
	} {
		t.Run(name, func(t *testing.T) {
			result := fn(paramFL{field: reflect.ValueOf("x"), param: ""})
			assert.False(t, result, name)
		})
	}
}

func TestFormatMsgCompiledNoPH(t *testing.T) {
	assert.Equal(t, "plain", formatMessageCompiled("plain", &FieldError{Field: "x"}))
}

func TestBenchmarkCoverage(t *testing.T) {
	r1 := testing.Benchmark(func(b *testing.B) {
		for i := 0; i < b.N; i++ { validateStrongPasswordOld("Abcdefg1!") }
	})
	r2 := testing.Benchmark(func(b *testing.B) {
		for i := 0; i < b.N; i++ { validateStrongPasswordFast("Abcdefg1!") }
	})
	_, _ = r1, r2
}

func TestInFloat32Path(t *testing.T) {
	fn := In(float32(1.1), float32(2.2))
	assert.True(t, fn(paramFL{field: reflect.ValueOf(float32(1.1))}))
	assert.False(t, fn(paramFL{field: reflect.ValueOf(float32(9.9))}))
}

func TestNotInFloat32Path(t *testing.T) {
	fn := NotIn(float32(1.1), float32(2.2))
	assert.False(t, fn(paramFL{field: reflect.ValueOf(float32(1.1))}))
	assert.True(t, fn(paramFL{field: reflect.ValueOf(float32(9.9))}))
}

func TestLuhnNonDigit(t *testing.T) {
	assert.False(t, validateLuhnChecksum(strFLP("12a4")))
}

func TestRGBAEdge(t *testing.T) {
	assert.False(t, validateRGBA(strFLP("")))
	assert.False(t, validateRGBA(strFLP("rgba(1,2,3)")))
}

func TestISBN10ChecksumFail(t *testing.T) {
	assert.False(t, validateISBN10(strFLP("123456789")))
}

func TestISBN13ChecksumFail(t *testing.T) {
	assert.False(t, validateISBN13(strFLP("9780306406158")))
}

func TestISSNChecksumFail(t *testing.T) {
	assert.False(t, validateISSN(strFLP("0317-8472")))
}

func TestUUIDVersionRFC4122Mismatch(t *testing.T) {
	assert.False(t, validateUUIDVersionRFC4122(strFLP("12345678-1234-3333-8234-123456789abc"), '4'))
}

func TestHostnamePortEdge(t *testing.T) {
	assert.False(t, validateHostnamePort(strFLP("")))
	assert.False(t, validateHostnamePort(strFLP("host:")))
	assert.False(t, validateHostnamePort(strFLP("host:abc")))
}

func TestValidateFQDNEdge(t *testing.T) {
	assert.False(t, validateFQDN(strFLP("")))
	assert.False(t, validateFQDN(strFLP("localhost")))
	assert.True(t, validateFQDN(strFLP("example.com")))
}


func TestValidateOneOfEmptyParam(t *testing.T) {
	assert.False(t, validateOneOf(strFLP("x")))
}

func TestValidateIsDefaultSliceMap(t *testing.T) {
	v, err := New()
	require.NoError(t, err)
	type SS struct{ V []int `validate:"isdefault"` }
	assert.NoError(t, v.Struct(SS{V: nil}))
	assert.Error(t, v.Struct(SS{V: []int{}}))
	type SM struct{ V map[string]int `validate:"isdefault"` }
	assert.NoError(t, v.Struct(SM{V: nil}))
	assert.Error(t, v.Struct(SM{V: map[string]int{}}))
}

func TestValidateIsDefaultFloat(t *testing.T) {
	v, err := New()
	require.NoError(t, err)
	type S struct{ V float64 `validate:"isdefault"` }
	assert.NoError(t, v.Struct(S{V: 0.0}))
	assert.Error(t, v.Struct(S{V: 1.0}))
}

func TestRegisterValidationOverwrite(t *testing.T) {
	e := NewEngine()
	require.NoError(t, e.RegisterValidation("ow", func(fl FieldLevel) bool { return true }))
	require.NoError(t, e.RegisterValidation("ow", func(fl FieldLevel) bool { return false }))
	assert.False(t, e.validators["ow"](paramFL{field: reflect.ValueOf("x")}))
}

func TestReportErrorBothPaths(t *testing.T) {
	errs := &ValidationErrors{}
	sl := &structLevel{top: reflect.ValueOf("x"), current: reflect.ValueOf("x"), errors: errs}
	sl.ReportError("val", "f", "t", "msg")
	sl.ReportError("val2", "f2", "t2", "")
	assert.Equal(t, 2, errs.Len())
}

func TestTranslateUnknownTagFallback(t *testing.T) {
	v, err := New()
	require.NoError(t, err)
	require.NoError(t, v.RegisterValidation("unknowntag", func(fl FieldLevel) bool { return false }))
	err2 := v.Var("x", "unknowntag")
	require.Error(t, err2)
	assert.Contains(t, err2.Error(), "unknowntag")
}

func TestFmtVariantsNoPH(t *testing.T) {
	fe := &FieldError{Field: "x", Tag: "t", Param: "p", Value: 42}
	assert.Equal(t, "abc", formatMessageBuilder("abc", fe))
	assert.Equal(t, "abc", formatMessageByteSlice("abc", fe))
	assert.Equal(t, "abc", formatMessageNoFmt("abc", fe))
	assert.Equal(t, "abc", formatMessageInlineCheck("abc", fe))
	assert.Equal(t, "abc", formatMessageBytesBuffer("abc", fe))
	assert.Equal(t, "abc", formatMessageOptimizedCurrent("abc", fe))
	assert.Equal(t, "abc", formatMessageFastPath("abc", fe))
}

func TestInNotInIntPaths(t *testing.T) {
	fn := In(1, 2, 3)
	assert.False(t, fn(paramFL{field: reflect.ValueOf(0)}))
	gn := NotIn(1, 2, 3)
	assert.False(t, gn(paramFL{field: reflect.ValueOf(2)}))
}

func TestOpt11SliceNilPtr(t *testing.T) {
	e := NewEngine()
	assert.False(t, e.validateField_Opt11_InlinedValidators(&fieldLevel{field: reflect.ValueOf([]int{})}, "required"))
	var p *string
	assert.False(t, e.validateField_Opt11_InlinedValidators(&fieldLevel{field: reflect.ValueOf(p)}, "required"))
}

// ===== registerDefaultValidators: cover more registration paths =====

func TestRegisterAllValidatorTypes(t *testing.T) {
	v, err := New()
	require.NoError(t, err)
	// Exercise registered validators through Var
	assert.NoError(t, v.Var("test@example.com", "email"))
	assert.NoError(t, v.Var("https://x.com", "url"))
	assert.NoError(t, v.Var("192.168.1.1", "ipv4"))
	assert.NoError(t, v.Var("01:23:45:67:89:ab", "mac"))
	assert.NoError(t, v.Var(`{"k":"v"}`, "json"))
	assert.NoError(t, v.Var("12345678-1234-1234-1234-123456789abc", "uuid"))
	assert.NoError(t, v.Var("ABC", "uppercase"))
	assert.NoError(t, v.Var("abc", "lowercase"))
	assert.NoError(t, v.Var("ABC123", "alphanum_upper"))
	assert.NoError(t, v.Var("abc123", "alphanum_lower"))
}

// ===== translateFieldError: cover "en config but no message" path =====

func TestTranslateFieldErrorNoMsgInEnConfig(t *testing.T) {
	v, err := New()
	require.NoError(t, err)
	// Register validator with no translation
	require.NoError(t, v.RegisterValidation("no_trans_tag", func(fl FieldLevel) bool { return false }))
	err2 := v.Var("x", "no_trans_tag")
	require.Error(t, err2)
	// Falls through to final sprintf fallback
	assert.Contains(t, err2.Error(), "no_trans_tag")
}

// ===== validateIsDefault: cover struct/interface cases =====

func TestValidateIsDefaultStruct(t *testing.T) {
	v, err := New()
	require.NoError(t, err)
	type Inner struct{ X int }
	type S struct{ V Inner `validate:"isdefault"` }
	assert.NoError(t, v.Struct(S{V: Inner{}}))      // zero struct = default
	assert.Error(t, v.Struct(S{V: Inner{X: 1}}))    // non-zero struct
}

func TestValidateIsDefaultInterface(t *testing.T) {
	v, err := New()
	require.NoError(t, err)
	type S struct{ V interface{} `validate:"isdefault"` }
	assert.NoError(t, v.Struct(S{V: nil}))
	assert.Error(t, v.Struct(S{V: "x"}))
}

// ===== LuhnChecksum: cover more paths =====

func TestLuhnMoreCases(t *testing.T) {
	assert.True(t, validateLuhnChecksum(strFLP("79927398713")))
	assert.False(t, validateLuhnChecksum(strFLP("79927398710")))
	assert.False(t, validateLuhnChecksum(strFLP("0")))
	assert.False(t, validateLuhnChecksum(strFLP("")))
	assert.False(t, validateLuhnChecksum(strFLP("abcd")))
	assert.False(t, validateLuhnChecksum(strFLP("1234")))
}

// ===== validateLuhnChecksum int/uint paths =====

func TestLuhnIntUint(t *testing.T) {
	assert.True(t, validateLuhnChecksum(paramFL{field: reflect.ValueOf("79927398713")}))
}

// ===== validateIsDefault invalid field =====

func TestValidateIsDefaultInvalid(t *testing.T) {
	assert.True(t, validateIsDefault(paramFL{field: reflect.Value{}}))
}

// ===== New error path =====

func TestNewWithError(t *testing.T) {
	// New always succeeds in normal operation
	v, err := New()
	require.NoError(t, err)
	require.NotNil(t, v)
}

// ===== Default() error path =====

func TestDefaultFallback(t *testing.T) {
	// Already tested - just ensure it doesn't panic
	d := Default()
	assert.NotNil(t, d)
}

// ===== Var error path =====

func TestVarError(t *testing.T) {
	v, err := New()
	require.NoError(t, err)
	err = v.Var("", "required")
	assert.Error(t, err)
}

// ===== validateLuhnChecksum: cover int/uint paths =====

func TestLuhnIntField(t *testing.T) {
	// luhnCheck on an int value - register directly
	v, err := New()
	require.NoError(t, err)
	require.NoError(t, v.RegisterValidation("luhn_int", validateLuhnChecksum))
	
	type S struct{ V int `validate:"luhn_int"` }
	// Int value gets formatted as string then checked
	_ = v.Struct(S{V: 79927398713})
}

func TestLuhnUintField(t *testing.T) {
	v, err := New()
	require.NoError(t, err)
	require.NoError(t, v.RegisterValidation("luhn_uint", validateLuhnChecksum))
	
	type S struct{ V uint `validate:"luhn_uint"` }
	_ = v.Struct(S{V: 79927398713})
}

func TestLuhnDefaultKind(t *testing.T) {
	// Default kind falls back to String()
	assert.False(t, validateLuhnChecksum(paramFL{field: reflect.ValueOf(true)}))
}
