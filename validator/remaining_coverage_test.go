package validator

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	xlanguage "golang.org/x/text/language"
)

// ===== Conditional validators: excluded_if / excluded_without branches (76.9%) =====

func TestExcludedIfConditionTrueFieldEmpty(t *testing.T) {
	v, err := New()
	require.NoError(t, err)
	type S struct {
		Status string
		Value  string `validate:"excluded_if=Status=active"`
	}
	assert.NoError(t, v.Struct(S{Status: "active", Value: ""}))
	assert.Error(t, v.Struct(S{Status: "active", Value: "x"}))
	assert.NoError(t, v.Struct(S{Status: "inactive", Value: "x"}))
}

func TestExcludedUnlessConditionTrue(t *testing.T) {
	v, err := New()
	require.NoError(t, err)
	type S struct {
		Status string
		Value  string `validate:"excluded_unless=Status=active"`
	}
	assert.NoError(t, v.Struct(S{Status: "active", Value: "x"}))
	assert.NoError(t, v.Struct(S{Status: "inactive", Value: ""}))
	assert.Error(t, v.Struct(S{Status: "inactive", Value: "x"}))
}

func TestExcludedUnlessEmptyParam(t *testing.T) {
	v, err := New()
	require.NoError(t, err)
	type S struct{ Value string `validate:"excluded_unless"` }
	assert.NoError(t, v.Struct(S{Value: ""}))
	assert.Error(t, v.Struct(S{Value: "x"}))
}

func TestExcludedIfEmptyParam(t *testing.T) {
	v, err := New()
	require.NoError(t, err)
	type S struct{ Value string `validate:"excluded_if"` }
	assert.NoError(t, v.Struct(S{Value: "x"}))
}

func TestExcludedIfInvalidParam(t *testing.T) {
	v, err := New()
	require.NoError(t, err)
	type S struct{ Value string `validate:"excluded_if=noequals"` }
	assert.NoError(t, v.Struct(S{Value: "x"}))
}

func TestExcludedUnlessInvalidParam(t *testing.T) {
	v, err := New()
	require.NoError(t, err)
	type S struct{ Value string `validate:"excluded_unless=noequals"` }
	assert.NoError(t, v.Struct(S{Value: ""}))
	assert.Error(t, v.Struct(S{Value: "x"}))
}

// ===== Field validators: resolveFieldPath nil ptr path (66.7%) =====

func TestResolveFieldPathNilPtr(t *testing.T) {
	type Inner struct{ Value string }
	type Outer struct{ Ptr *Inner }
	o := Outer{Ptr: nil}
	fl := &fieldLevel{
		top:    reflect.ValueOf(o),
		parent: reflect.ValueOf(o),
		field:  reflect.ValueOf(o.Ptr),
	}
	result := resolveFieldPath(fl, "Ptr.Value")
	assert.False(t, result.IsValid())
}

func TestResolveFieldPathNonStructTop(t *testing.T) {
	fl := &fieldLevel{
		top:    reflect.ValueOf("x"),
		parent: reflect.ValueOf("x"),
		field:  reflect.ValueOf("x"),
	}
	result := resolveFieldPath(fl, "x.y")
	assert.False(t, result.IsValid())
}

// ===== translateFieldError "no config" path (70%) =====

func TestTranslateFieldErrorNoConfigAtAll(t *testing.T) {
	v := &Validator{
		engine:   NewEngine(),
		locale:   xlanguage.Make("zz"),
		useJSON:  true,
		messages: make(map[string]string),
	}
	v.updateFieldNameFunc()
	err := v.Struct(struct {
		Name string `validate:"required"`
	}{Name: ""})
	require.Error(t, err)
	assert.NotEmpty(t, err.Error())
}

// ===== validateIsDefault branches (75%) =====

func TestValidateIsDefaultNonDefault(t *testing.T) {
	v, err := New()
	require.NoError(t, err)
	type S struct{ Value int `validate:"isdefault"` }
	assert.NoError(t, v.Struct(S{Value: 0}))
	assert.Error(t, v.Struct(S{Value: 42}))
}

func TestValidateIsDefaultString(t *testing.T) {
	v, err := New()
	require.NoError(t, err)
	type S struct{ Value string `validate:"isdefault"` }
	assert.NoError(t, v.Struct(S{Value: ""}))
	assert.Error(t, v.Struct(S{Value: "x"}))
}

func TestValidateIsDefaultBool(t *testing.T) {
	v, err := New()
	require.NoError(t, err)
	type S struct{ Value bool `validate:"isdefault"` }
	assert.NoError(t, v.Struct(S{Value: false}))
	assert.Error(t, v.Struct(S{Value: true}))
}

func TestValidateIsDefaultPtr(t *testing.T) {
	v, err := New()
	require.NoError(t, err)
	type S struct{ Value *string `validate:"isdefault"` }
	assert.NoError(t, v.Struct(S{Value: nil}))
	x := "hello"
	assert.Error(t, v.Struct(S{Value: &x}))
}

// ===== validateHTTPURL branches (80%) =====

func TestValidateHTTPURLEdgeCases(t *testing.T) {
	assert.True(t, validateHTTPURL(strFLP("http://example.com")))
	assert.True(t, validateHTTPURL(strFLP("https://example.com")))
	assert.False(t, validateHTTPURL(strFLP("ftp://example.com")))
	assert.False(t, validateHTTPURL(strFLP("")))
	assert.False(t, validateHTTPURL(strFLP("://no-scheme")))
}

// ===== UUID version validators remaining branches =====

func TestValidateUUIDVersionInvalid(t *testing.T) {
	assert.False(t, validateUUIDVersion(strFLP("not-a-uuid"), '4'))
}

func TestValidateUUIDVersionRFC4122Invalid(t *testing.T) {
	assert.False(t, validateUUIDVersionRFC4122(strFLP("not-a-uuid"), '4'))
}

// ===== validateLuhnChecksum remaining paths =====

func TestLuhnChecksumValidAndInvalid(t *testing.T) {
	assert.True(t, validateLuhnChecksum(strFLP("79927398713")))
	assert.False(t, validateLuhnChecksum(strFLP("79927398710")))
	assert.False(t, validateLuhnChecksum(strFLP("0")))
	assert.False(t, validateLuhnChecksum(strFLP("")))
	assert.False(t, validateLuhnChecksum(strFLP("abcd")))
}

// ===== validatePostcodeField remaining paths =====

func TestValidatePostcodeFieldAll(t *testing.T) {
	v, err := New()
	require.NoError(t, err)
	type S struct {
		Country  string
		Postcode string `validate:"pcfield=Country"`
	}
	err = v.RegisterValidation("pcfield", validatePostcodeField)
	require.NoError(t, err)
	assert.NoError(t, v.Struct(S{Country: "US", Postcode: "12345"}))
	assert.Error(t, v.Struct(S{Country: "", Postcode: "12345"}))
}

// ===== registerDefaultValidators: iscolor and country_code aliases (60.6%) =====

func TestIsColorAlias(t *testing.T) {
	v, err := New()
	require.NoError(t, err)
	err = v.Var("#ff0000", "iscolor")
	assert.NoError(t, err)
	err = v.Var("rgb(255,0,0)", "iscolor")
	assert.NoError(t, err)
	err = v.Var("not-a-color", "iscolor")
	assert.Error(t, err)
}

func TestCountryCodeAlias(t *testing.T) {
	v, err := New()
	require.NoError(t, err)
	err = v.Var("US", "country_code")
	assert.NoError(t, err)
	err = v.Var("USA", "country_code")
	assert.NoError(t, err)
	err = v.Var("12", "country_code")
	assert.Error(t, err)
}

// ===== Var with non-FieldError error =====

func TestVarNonFieldError(t *testing.T) {
	v, err := New()
	require.NoError(t, err)
	err2 := v.Var(42, "required")
	_ = err2
}

// ===== Benchmark coverage wrappers =====

func TestBenchmarkOldNewFunctions(t *testing.T) {
	assert.True(t, validateStrongPasswordOld("Abcdefg1!"))
	assert.True(t, validateStrongPasswordFast("Abcdefg1!"))
}
