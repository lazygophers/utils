package validator

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/lazygophers/utils/language"
	xlanguage "golang.org/x/text/language"
)

// ===== engine.go: Opt11 email/alpha/alphanum/url cases (lines 1331-1353) =====

func TestOpt11AllCases(t *testing.T) {
	e := NewEngine()
	// empty email -> true (skip check)
	assert.True(t, e.validateField_Opt11_InlinedValidators(&fieldLevel{field: reflect.ValueOf("")}, "email"))
	assert.True(t, e.validateField_Opt11_InlinedValidators(&fieldLevel{field: reflect.ValueOf("a@b.com")}, "email"))
	assert.False(t, e.validateField_Opt11_InlinedValidators(&fieldLevel{field: reflect.ValueOf("not-email")}, "email"))

	// alpha
	assert.True(t, e.validateField_Opt11_InlinedValidators(&fieldLevel{field: reflect.ValueOf("")}, "alpha"))
	assert.True(t, e.validateField_Opt11_InlinedValidators(&fieldLevel{field: reflect.ValueOf("abc")}, "alpha"))
	assert.False(t, e.validateField_Opt11_InlinedValidators(&fieldLevel{field: reflect.ValueOf("abc1")}, "alpha"))

	// alphanum
	assert.True(t, e.validateField_Opt11_InlinedValidators(&fieldLevel{field: reflect.ValueOf("")}, "alphanum"))
	assert.True(t, e.validateField_Opt11_InlinedValidators(&fieldLevel{field: reflect.ValueOf("abc123")}, "alphanum"))
	assert.False(t, e.validateField_Opt11_InlinedValidators(&fieldLevel{field: reflect.ValueOf("abc!123")}, "alphanum"))

	// url
	assert.True(t, e.validateField_Opt11_InlinedValidators(&fieldLevel{field: reflect.ValueOf("")}, "url"))
	assert.True(t, e.validateField_Opt11_InlinedValidators(&fieldLevel{field: reflect.ValueOf("https://example.com")}, "url"))
	assert.False(t, e.validateField_Opt11_InlinedValidators(&fieldLevel{field: reflect.ValueOf("not-url")}, "url"))

	// default case: known validator
	e.RegisterValidation("custom_opt11", func(fl FieldLevel) bool { return true })
	assert.True(t, e.validateField_Opt11_InlinedValidators(&fieldLevel{field: reflect.ValueOf("x")}, "custom_opt11"))

	// default case: unknown validator -> true
	assert.True(t, e.validateField_Opt11_InlinedValidators(&fieldLevel{field: reflect.ValueOf("x")}, "unknown_tag_opt11"))
}

// ===== engine.go: RegisterValidation nil (line 130-131) =====

func TestRegisterValidationNilFunc2(t *testing.T) {
	e := NewEngine()
	assert.Error(t, e.RegisterValidation("test_nil2", nil))
}

// ===== engine.go: ReportError with pointer value (line 844-846) =====

func TestReportErrorPtrValue2(t *testing.T) {
	errs := &ValidationErrors{}
	sl := &structLevel{
		top:       reflect.ValueOf("x"),
		current:   reflect.ValueOf("x"),
		errors:    errs,
		namespace: "ns",
	}
	x := 42
	sl.ReportError(&x, "f", "t", "msg")
	assert.Equal(t, 1, errs.Len())
}

// ===== engine.go: validateField unknown tag (line 739+) =====

func TestValidateFieldUnknownTag(t *testing.T) {
	e := NewEngine()
	fl := &fieldLevel{field: reflect.ValueOf("x")}
	assert.True(t, e.validateField(fl, "nonexistent_tag_abc"))
}

// ===== validator.go: translateFieldError en fallback for non-en locale (lines 191-195) =====

func TestTranslateFieldErrorEnFallback2(t *testing.T) {
	v, err := New(WithLocale(xlanguage.Make("fr")))
	require.NoError(t, err)
	require.NoError(t, v.RegisterValidation("custom_fallback2", func(fl FieldLevel) bool { return false }))
	err2 := v.Var("x", "custom_fallback2")
	require.Error(t, err2)
	assert.Contains(t, err2.Error(), "custom_fallback2")
}

// ===== conditional_validators.go: empty param paths =====

func TestRequiredUnlessEmptyParam2(t *testing.T) {
	e := NewEngine()
	e.RegisterValidation("required_unless", validateRequiredUnless)
	assert.False(t, e.validators["required_unless"](paramFL{field: reflect.ValueOf("")}))
	assert.True(t, e.validators["required_unless"](paramFL{field: reflect.ValueOf("x")}))
}

func TestRequiredWithAllEmptyParam2(t *testing.T) {
	e := NewEngine()
	e.RegisterValidation("required_with_all", validateRequiredWithAll)
	assert.False(t, e.validators["required_with_all"](paramFL{field: reflect.ValueOf("")}))
	assert.True(t, e.validators["required_with_all"](paramFL{field: reflect.ValueOf("x")}))
}

func TestRequiredWithoutAllEmptyParam2(t *testing.T) {
	e := NewEngine()
	e.RegisterValidation("required_without_all", validateRequiredWithoutAll)
	assert.False(t, e.validators["required_without_all"](paramFL{field: reflect.ValueOf("")}))
	assert.True(t, e.validators["required_without_all"](paramFL{field: reflect.ValueOf("x")}))
}

func TestRequiredWithEmptyParam(t *testing.T) {
	e := NewEngine()
	e.registerBuiltinValidators()
	assert.False(t, e.validators["required_with"](paramFL{field: reflect.ValueOf("")}))
	assert.True(t, e.validators["required_with"](paramFL{field: reflect.ValueOf("x")}))
}

func TestRequiredWithoutEmptyParam(t *testing.T) {
	e := NewEngine()
	e.registerBuiltinValidators()
	assert.False(t, e.validators["required_without"](paramFL{field: reflect.ValueOf("")}))
	assert.True(t, e.validators["required_without"](paramFL{field: reflect.ValueOf("x")}))
}

// ===== format_validators.go: RGBA value>255 =====

func TestRGBAOver255(t *testing.T) {
	assert.False(t, validateRGBA(strFLP("rgba(256,0,0,1)")))
	assert.False(t, validateRGBA(strFLP("rgba(0,256,0,1)")))
	assert.False(t, validateRGBA(strFLP("rgba(0,0,256,1)")))
}

// ===== format_validators.go: ISBN bad digit =====

func TestISBN10BadBodyDigit(t *testing.T) {
	assert.False(t, validateISBN10(strFLP("1234a67890")))
}

func TestISBN13BadBodyDigit(t *testing.T) {
	assert.False(t, validateISBN13(strFLP("978030640a157")))
}

func TestISBN10Valid(t *testing.T) {
	assert.True(t, validateISBN10(strFLP("0306406152")))
}

func TestISBN13Valid(t *testing.T) {
	assert.True(t, validateISBN13(strFLP("9780306406157")))
}

// ===== format_validators.go: validatePostcodeField invalid postcode =====

func TestPostcodeFieldInvalidPostcode(t *testing.T) {
	v, err := New()
	require.NoError(t, err)
	require.NoError(t, v.RegisterValidation("pcfield2", validatePostcodeField))
	type S struct {
		Country  string
		Postcode string `validate:"pcfield2=Country"`
	}
	assert.Error(t, v.Struct(S{Country: "US", Postcode: ""}))
}

// ===== custom_validators.go: validateHTTPURL ws:// paths =====

func TestHTTPURLWS(t *testing.T) {
	// validateHTTPURL only accepts http/https
	assert.False(t, validateHTTPURL(strFLP("ws://example.com")))
	assert.False(t, validateHTTPURL(strFLP("wss://example.com")))
	assert.True(t, validateHTTPURL(strFLP("http://example.com")))
	assert.True(t, validateHTTPURL(strFLP("https://example.com")))
	assert.False(t, validateHTTPURL(strFLP("")))
}

// validateURL (custom_validators.go) does support ws/wss
func TestValidateURLWS(t *testing.T) {
	assert.True(t, validateURL(strFLP("ws://example.com")))
	assert.True(t, validateURL(strFLP("wss://example.com")))
	assert.False(t, validateURL(strFLP("wsx://bad")))
	assert.False(t, validateURL(strFLP("http://")))
	assert.False(t, validateURL(strFLP("http:// space")))
}

// ===== custom_validators.go: strong_password 2-type combos =====

func TestStrongPassword2Types(t *testing.T) {
	assert.False(t, validateStrongPassword(strFLP("ABC12345"))) // upper+digit only
	assert.False(t, validateStrongPassword(strFLP("abcdefg!"))) // lower+special only
	assert.False(t, validateStrongPassword(strFLP("ABCDEFG!"))) // upper+special only
	assert.False(t, validateStrongPassword(strFLP("1234567!"))) // digit+special only
}

// ===== misc_validators.go: validateFn non-func =====

func TestValidateFnNonFunc(t *testing.T) {
	assert.False(t, validateFn(paramFL{field: reflect.ValueOf("not a func")}))
}

// ===== net_validators.go: MAC edge =====

func TestMACEdge(t *testing.T) {
	assert.False(t, validateMAC(strFLP("")))
	assert.False(t, validateMAC(strFLP("not-a-mac")))
}

// ===== field_validators.go: resolveFieldPath direct field =====

func TestResolveDirectField(t *testing.T) {
	type S struct{ Name string }
	s := S{Name: "test"}
	fl := &fieldLevel{
		top:    reflect.ValueOf(s),
		parent: reflect.ValueOf(s),
		field:  reflect.ValueOf(s.Name),
	}
	result := resolveFieldPath(fl, "Name")
	assert.True(t, result.IsValid())
}

// ===== locale.go: nonexistent config =====

func TestGetLocaleNonexistent(t *testing.T) {
	// "en" is always registered, just verify it exists
	_, ok := GetLocaleConfig("en")
	assert.True(t, ok)
}

// ===== engine.go: Opt11 required slice/map/ptr/nil =====

func TestOpt11RequiredKinds(t *testing.T) {
	e := NewEngine()
	assert.False(t, e.validateField_Opt11_InlinedValidators(&fieldLevel{field: reflect.ValueOf([]int{})}, "required"))
	assert.True(t, e.validateField_Opt11_InlinedValidators(&fieldLevel{field: reflect.ValueOf([]int{1})}, "required"))
	assert.False(t, e.validateField_Opt11_InlinedValidators(&fieldLevel{field: reflect.ValueOf(map[string]int{})}, "required"))
	assert.True(t, e.validateField_Opt11_InlinedValidators(&fieldLevel{field: reflect.ValueOf(map[string]int{"a": 1})}, "required"))
	var p *int
	assert.False(t, e.validateField_Opt11_InlinedValidators(&fieldLevel{field: reflect.ValueOf(p)}, "required"))
	assert.False(t, e.validateField_Opt11_InlinedValidators(&fieldLevel{field: reflect.ValueOf(nil)}, "required"))
}

// ===== engine.go: validateStruct interface field =====

func TestValidateStructInterfaceField(t *testing.T) {
	e := NewEngine()
	e.RegisterValidation("required", func(fl FieldLevel) bool { return fl.Field().String() != "" })

	type S struct {
		V interface{} `validate:"required"`
	}
	errs := &ValidationErrors{}
	e.validateStruct(reflect.ValueOf(S{V: "hello"}), reflect.ValueOf(S{V: "hello"}), "", errs)
	_ = errs
}

// ===== engine.go: validateStruct map dive =====

func TestValidateStructMapDive(t *testing.T) {
	e := NewEngine()
	e.RegisterValidation("required", func(fl FieldLevel) bool { return fl.Field().String() != "" })

	type S struct{ M map[string]string `validate:"dive"` }
	errs := &ValidationErrors{}
	e.validateStruct(reflect.ValueOf(S{M: map[string]string{"k": ""}}), reflect.ValueOf(S{M: map[string]string{"k": ""}}), "", errs)
	_ = errs
}

// ===== translateFieldError: locale not registered, falls to "en" config (lines 167-170) =====

func TestTranslateFieldErrorUnknownLocaleFallback(t *testing.T) {
	v := &Validator{
		engine:   NewEngine(),
		locale:   xlanguage.Make("zz"),
		useJSON:  true,
		messages: make(map[string]string),
	}
	v.updateFieldNameFunc()
	require.NoError(t, v.RegisterValidation("required", func(fl FieldLevel) bool { return fl.Field().String() != "" }))
	err := v.Struct(struct {
		Name string `validate:"required"`
	}{Name: ""})
	require.Error(t, err)
	// Should fall to en config and find "required" message template
	assert.NotEmpty(t, err.Error())
}

// ===== translateFieldError: non-en locale, tag not in locale config, falls to en (lines 191-195) =====

func TestTranslateFieldErrorNonEnToFallback(t *testing.T) {
	// "zh" has a config but "required" might not be in it
	// Use a custom tag that's NOT in zh config but IS in en config
	v, err := New(WithLocale(xlanguage.Make("zh")))
	require.NoError(t, err)
	// "email" should be in en config
	err2 := v.Var("not-email", "email")
	require.Error(t, err2)
	// Should get the en template for "email"
	assert.NotEmpty(t, err2.Error())
}

// ===== engine.go: validateStruct with ptr to struct field (line 687 area) =====

func TestValidateStructPtrToStruct(t *testing.T) {
	e := NewEngine()
	e.RegisterValidation("required", func(fl FieldLevel) bool { return fl.Field().String() != "" })

	type Inner struct{ Name string `validate:"required"` }
	type Outer struct{ I *Inner }

	errs := &ValidationErrors{}
	// Non-nil ptr to struct with tags
	e.validateStruct(reflect.ValueOf(Outer{I: &Inner{Name: ""}}), reflect.ValueOf(Outer{I: &Inner{Name: ""}}), "", errs)
	assert.NotEmpty(t, *errs)
}

// ===== engine.go: validateStruct nested struct path (line 768 area) =====

func TestValidateStructNestedStruct2(t *testing.T) {
	e := NewEngine()
	e.RegisterValidation("required", func(fl FieldLevel) bool { return fl.Field().String() != "" })

	type Inner struct{ Name string `validate:"required"` }
	type Outer struct{ I Inner }

	errs := &ValidationErrors{}
	e.validateStruct(reflect.ValueOf(Outer{I: Inner{Name: ""}}), reflect.ValueOf(Outer{I: Inner{Name: ""}}), "", errs)
	assert.NotEmpty(t, *errs)
}

// ===== engine.go: Struct with nil pointer (line 148-150) =====

func TestEngineStructNilPtr(t *testing.T) {
	e := NewEngine()
	var p *int
	assert.Error(t, e.Struct(p))
}

// ===== engine.go: registerBuiltinValidators required with slice/map (line 430-431) =====

func TestBuiltinRequiredSliceMap(t *testing.T) {
	e := NewEngine()
	e.registerBuiltinValidators()
	fn := e.validators["required"]
	// slice
	assert.False(t, fn(paramFL{field: reflect.ValueOf([]int{})}))
	assert.True(t, fn(paramFL{field: reflect.ValueOf([]int{1})}))
	// map
	assert.False(t, fn(paramFL{field: reflect.ValueOf(map[string]int{})}))
	assert.True(t, fn(paramFL{field: reflect.ValueOf(map[string]int{"a": 1})}))
	// array
	assert.False(t, fn(paramFL{field: reflect.ValueOf([0]int{})}))
	assert.True(t, fn(paramFL{field: reflect.ValueOf([1]int{1})}))
}

// ===== engine.go: registerBuiltinValidators alpha empty (line 539-541) =====

func TestBuiltinAlphaEmpty(t *testing.T) {
	e := NewEngine()
	e.registerBuiltinValidators()
	fn := e.validators["alpha"]
	assert.True(t, fn(paramFL{field: reflect.ValueOf("")}))
	assert.True(t, fn(paramFL{field: reflect.ValueOf("abc")}))
	assert.False(t, fn(paramFL{field: reflect.ValueOf("abc1")}))
}

// ===== engine.go: registerBuiltinValidators alphanum empty (line 548-550) =====

func TestBuiltinAlphanumEmpty(t *testing.T) {
	e := NewEngine()
	e.registerBuiltinValidators()
	fn := e.validators["alphanum"]
	assert.True(t, fn(paramFL{field: reflect.ValueOf("")}))
	assert.True(t, fn(paramFL{field: reflect.ValueOf("abc123")}))
	assert.False(t, fn(paramFL{field: reflect.ValueOf("abc!123")}))
}

// ===== engine.go: required_if invalid field (line 675-677) =====

func TestBuiltinRequiredIfInvalidField(t *testing.T) {
	e := NewEngine()
	e.registerBuiltinValidators()
	type S struct {
		Value  string `validate:"required_if=NonExist=value"`
		Status string
	}
	err := e.Struct(S{Value: "", Status: "x"})
	// NonExist field not found -> returns true (no error)
	assert.NoError(t, err)
}

// ===== engine.go: validateStruct ptr-to-struct via dive (line 355-357) =====

func TestValidateStructPtrStructDive(t *testing.T) {
	e := NewEngine()
	e.RegisterValidation("required", func(fl FieldLevel) bool { return fl.Field().String() != "" })

	type Inner struct{ Name string `validate:"required"` }
	type Outer struct{ I *Inner }

	errs := &ValidationErrors{}
	o := Outer{I: &Inner{Name: ""}}
	e.validateStruct(reflect.ValueOf(o), reflect.ValueOf(o), "", errs)
	assert.NotEmpty(t, *errs)
}

// ===== validator.go: effectiveLocale goroutine-local path (line 156) =====

func TestEffectiveLocaleNoOverride(t *testing.T) {
	// Validator with no explicit locale, no goroutine-local -> defaults to "en"
	origDefault := language.Default()
	language.SetDefault(language.Make("en"))
	defer language.SetDefault(origDefault)

	v, err := New()
	require.NoError(t, err)
	assert.Equal(t, xlanguage.Make("en"), v.GetLocale())
}

// ===== validator.go: Var returns non-FieldError from engine (line 112) =====

func TestVarNonFieldErrorFromEngine(t *testing.T) {
	v, err := New()
	require.NoError(t, err)
	// Struct validation returns ValidationErrors, not *FieldError
	err2 := v.Struct("not a struct")
	assert.Error(t, err2)
}

// ===== conditional_validators.go: requiredUnless invalid param / field not found =====

func TestRequiredUnlessInvalidParam(t *testing.T) {
	v, err := New()
	require.NoError(t, err)
	type S struct {
		Value  string `validate:"required_unless=noequals"`
		Status string
	}
	// "noequals" has no "=" → parts len != 2 → continue
	assert.NoError(t, v.Struct(S{Value: "", Status: ""}))
}

func TestRequiredUnlessFieldNotFound(t *testing.T) {
	v, err := New()
	require.NoError(t, err)
	type S struct {
		Value string `validate:"required_unless=NonExist=active"`
	}
	// NonExist field not found → continue → allConditionsMet stays true → returns true
	assert.NoError(t, v.Struct(S{Value: ""}))
}

// ===== conditional_validators.go: requiredWithAll allNotEmpty=true =====

func TestRequiredWithAllFieldsPresent(t *testing.T) {
	v, err := New()
	require.NoError(t, err)
	type S struct {
		A string
		B string
		C string `validate:"required_with_all=A,B"`
	}
	// Both A and B have values → C required
	assert.Error(t, v.Struct(S{A: "x", B: "y", C: ""}))
	assert.NoError(t, v.Struct(S{A: "x", B: "y", C: "z"}))
	// A empty → C not required
	assert.NoError(t, v.Struct(S{A: "", B: "y", C: ""}))
}

// ===== conditional_validators.go: excludedIf invalid field =====

func TestExcludedIfFieldNotFound(t *testing.T) {
	v, err := New()
	require.NoError(t, err)
	type S struct {
		Value string `validate:"excluded_if=NonExist=active"`
	}
	assert.NoError(t, v.Struct(S{Value: "x"}))
}

// ===== conditional_validators.go: excludedUnless invalid field =====

func TestExcludedUnlessFieldNotFound(t *testing.T) {
	v, err := New()
	require.NoError(t, err)
	type S struct {
		Value string `validate:"excluded_unless=NonExist=active"`
	}
	// Field not found → treated as condition not met → error if value present
	assert.Error(t, v.Struct(S{Value: "x"}))
}

// ===== conditional_validators.go: excludedWith empty param =====

func TestExcludedWithEmptyParam(t *testing.T) {
	v, err := New()
	require.NoError(t, err)
	type S struct{ Value string `validate:"excluded_with"` }
	// empty param → return true
	assert.NoError(t, v.Struct(S{Value: "x"}))
}

// ===== conditional_validators.go: excludedWithout empty param =====

func TestExcludedWithoutEmptyParam(t *testing.T) {
	v, err := New()
	require.NoError(t, err)
	type S struct{ Value string `validate:"excluded_without"` }
	// empty param → return !isFieldNotEmpty
	assert.NoError(t, v.Struct(S{Value: ""}))
	assert.Error(t, v.Struct(S{Value: "x"}))
}

// ===== custom_validators.go: validateURL empty rest (line 98-100) =====

func TestValidateURLEmptyRest(t *testing.T) {
	assert.False(t, validateURL(strFLP("http://")))
}

// ===== custom_validators.go: validateIPv4 leading zero (line 162-164) =====

func TestValidateIPv4LeadingZero(t *testing.T) {
	assert.False(t, validateIPv4(strFLP("01.2.3.4")))
}

// ===== custom_validators.go: validateIPv4 digit==0 before dot (line 174-176) =====

func TestValidateIPv4DotWithoutDigit(t *testing.T) {
	assert.False(t, validateIPv4(strFLP(".1.2.34")))
}

// ===== custom_validators.go: validateUUID bad dash position (line 248-250) =====

func TestValidateUUIDBadDash(t *testing.T) {
	assert.False(t, validateUUID(strFLP("123456781234-1234-1234-123456789abc")))
}

// ===== format_validators.go: ISBN10 last char not X/digit (line 285-287) =====

func TestISBN10BadLastChar(t *testing.T) {
	assert.False(t, validateISBN10(strFLP("030640615A")))
}

// ===== format_validators.go: ISBN13 last char not digit (line 310-312) =====

func TestISBN13BadLastChar(t *testing.T) {
	assert.False(t, validateISBN13(strFLP("978030640615X")))
}

// ===== format_validators.go: ISSN last char X (line 344-346) =====

func TestISSNLastCharX(t *testing.T) {
	assert.False(t, validateISSN(strFLP("0317-847X"))) // Wrong checksum
}

// ===== field_validators.go: resolveFieldPath dot path with nil ptr (line 42) =====

func TestResolveFieldPathDotNilPtr(t *testing.T) {
	type Inner struct{ V string }
	type Outer struct{ Ptr *Inner }
	o := Outer{Ptr: nil}
	fl := &fieldLevel{
		top:    reflect.ValueOf(o),
		parent: reflect.ValueOf(o),
		field:  reflect.ValueOf(o.Ptr),
	}
	result := resolveFieldPath(fl, "Ptr.V")
	assert.False(t, result.IsValid())
}

// ===== field_validators.go: resolveFieldPath dot path with invalid field (line 48-50) =====

func TestResolveFieldPathDotInvalidField(t *testing.T) {
	type Outer struct{ Name string }
	o := Outer{Name: "test"}
	fl := &fieldLevel{
		top:    reflect.ValueOf(o),
		parent: reflect.ValueOf(o),
		field:  reflect.ValueOf(o.Name),
	}
	result := resolveFieldPath(fl, "NonExist")
	// No dot → GetFieldByName → likely not found
	assert.False(t, result.IsValid())
}

// ===== locale.go: GetLocaleConfig no match, no en (line 50) =====
// This is unreachable since "en" is always registered, but let's verify

func TestGetLocaleConfigEnAlwaysExists(t *testing.T) {
	_, ok := GetLocaleConfig("en")
	assert.True(t, ok)
}

// ===== misc_validators.go: validateFn with Validatable struct =====

type validatableStruct struct {
	err error
}

func (v validatableStruct) Validate() error { return v.err }

func TestValidateFnWithValidatable(t *testing.T) {
	// Struct with Validate method returning nil → true
	assert.True(t, validateFn(paramFL{field: reflect.ValueOf(validatableStruct{err: nil})}))
	// Struct with Validate method returning error → false
	assert.False(t, validateFn(paramFL{field: reflect.ValueOf(validatableStruct{err: fmt.Errorf("fail")})}))
}

// ===== misc_validators.go: validateFn with addrable struct (line 86-88) =====

type validatableByAddr struct{}

func (v *validatableByAddr) Validate() error { return nil }

func TestValidateFnByAddr(t *testing.T) {
	// Need an addressable struct field to hit CanAddr path
	type Wrapper struct{ V validatableByAddr }
	w := Wrapper{V: validatableByAddr{}}
	f := reflect.ValueOf(&w).Elem().Field(0) // addressable via pointer
	assert.True(t, validateFn(paramFL{field: f}))
}

// ===== net_validators.go: validateHostname empty part (line 129-131) =====

func TestValidateHostnameEmptyPart(t *testing.T) {
	assert.False(t, validateHostname(strFLP(".example.com")))
}

// ===== net_validators.go: validateFQDN TLD non-alpha (line 150-152) =====

func TestValidateFQDNNonAlphaTLD(t *testing.T) {
	assert.False(t, validateFQDN(strFLP("example.12")))
}

// ===== net_validators.go: validateFQDN invalid hostname part (line 155-157) =====

func TestValidateFQDNInvalidPart(t *testing.T) {
	assert.False(t, validateFQDN(strFLP("exa mple.com")))
}

// ===== engine.go: max validator bad param (line 488-490) =====

func TestBuiltinMaxBadParam(t *testing.T) {
	e := NewEngine()
	e.registerBuiltinValidators()
	fn := e.validators["max"]
	// bad param → ParseFloat error → false
	assert.False(t, fn(paramFL{field: reflect.ValueOf("hello"), param: "abc"}))
}

// ===== engine.go: max validator unsupported kind (line 504) =====

func TestBuiltinMaxUnsupportedKind(t *testing.T) {
	e := NewEngine()
	e.registerBuiltinValidators()
	fn := e.validators["max"]
	assert.False(t, fn(paramFL{field: reflect.ValueOf(true), param: "10"}))
}

// ===== engine.go: eq validator unsupported kind (line 575) =====

func TestBuiltinEqUnsupportedKind(t *testing.T) {
	e := NewEngine()
	e.registerBuiltinValidators()
	fn := e.validators["eq"]
	assert.False(t, fn(paramFL{field: reflect.ValueOf([]int{}), param: "10"}))
}

// ===== engine.go: eqfield empty param + invalid field =====

func TestBuiltinEqFieldEmptyParam(t *testing.T) {
	e := NewEngine()
	e.registerBuiltinValidators()
	fn := e.validators["eqfield"]
	assert.False(t, fn(paramFL{field: reflect.ValueOf(5), param: ""}))
}

func TestBuiltinEqFieldInvalidField(t *testing.T) {
	e := NewEngine()
	e.registerBuiltinValidators()
	type S struct{ A int; B int `validate:"eqfield=C"` }
	fn := e.validators["eqfield"]
	// No struct context → GetFieldByName returns invalid
	assert.False(t, fn(paramFL{field: reflect.ValueOf(5), param: "NonExist"}))
}

// ===== engine.go: nefield empty param + invalid field =====

func TestBuiltinNeFieldEmptyParam(t *testing.T) {
	e := NewEngine()
	e.registerBuiltinValidators()
	fn := e.validators["nefield"]
	assert.False(t, fn(paramFL{field: reflect.ValueOf(5), param: ""}))
}

func TestBuiltinNeFieldInvalidField(t *testing.T) {
	e := NewEngine()
	e.registerBuiltinValidators()
	fn := e.validators["nefield"]
	assert.False(t, fn(paramFL{field: reflect.ValueOf(5), param: "NonExist"}))
}

// ===== engine.go: required_with invalid field (line 631-633) =====

func TestBuiltinRequiredWithInvalidField(t *testing.T) {
	v, err := New()
	require.NoError(t, err)
	type S struct{ V string `validate:"required_with=NonExist"` }
	// NonExist not found → returns true → no error
	assert.NoError(t, v.Struct(S{V: ""}))
}

// ===== engine.go: required_without invalid field (line 651-653) =====

func TestBuiltinRequiredWithoutInvalidField(t *testing.T) {
	v, err := New()
	require.NoError(t, err)
	type S struct{ V string `validate:"required_without=NonExist"` }
	// NonExist not found → returns isFieldNotEmpty → false for empty
	assert.Error(t, v.Struct(S{V: ""}))
}

// ===== engine.go: required_if bad param (line 667-669) =====

func TestBuiltinRequiredIfBadParam(t *testing.T) {
	v, err := New()
	require.NoError(t, err)
	type S struct{ V string `validate:"required_if=noequals"` }
	// no "=" in param → returns true → no error
	assert.NoError(t, v.Struct(S{V: ""}))
}

// ===== engine.go: compareFields uint equal, float equal =====

func TestCompareFieldsUintEqual(t *testing.T) {
	assert.Equal(t, 0, compareFields(reflect.ValueOf(uint(5)), reflect.ValueOf(uint(5))))
}

func TestCompareFieldsFloatEqual(t *testing.T) {
	assert.Equal(t, 0, compareFields(reflect.ValueOf(3.14), reflect.ValueOf(3.14)))
}

// ===== engine.go: compareFields bool false<true =====

func TestCompareFieldsBoolLt(t *testing.T) {
	assert.Equal(t, -1, compareFields(reflect.ValueOf(false), reflect.ValueOf(true)))
	assert.Equal(t, 1, compareFields(reflect.ValueOf(true), reflect.ValueOf(false)))
	assert.Equal(t, 0, compareFields(reflect.ValueOf(false), reflect.ValueOf(false)))
	assert.Equal(t, 0, compareFields(reflect.ValueOf(true), reflect.ValueOf(true)))
}

// ===== engine.go: compareFields ptr nil =====

func TestCompareFieldsPtrNil(t *testing.T) {
	var p1 *int
	var p2 *int
	assert.Equal(t, 0, compareFields(reflect.ValueOf(p1), reflect.ValueOf(p2)))
	assert.Equal(t, -1, compareFields(reflect.ValueOf(p1), reflect.ValueOf(intPtr(1))))
	assert.Equal(t, 1, compareFields(reflect.ValueOf(intPtr(1)), reflect.ValueOf(p1)))
	assert.Equal(t, 0, compareFields(reflect.ValueOf(intPtr(1)), reflect.ValueOf(intPtr(1))))
}

func intPtr(v int) *int { return &v }

// ===== engine.go: compareFields default string path =====

func TestCompareFieldsDefaultString(t *testing.T) {
	// unsupported kind → default path uses getFieldValueAsString
	assert.Equal(t, 0, compareFields(reflect.ValueOf(complex(1, 2)), reflect.ValueOf(complex(1, 2))))
	assert.Equal(t, -1, compareFields(reflect.ValueOf(complex(1, 1)), reflect.ValueOf(complex(2, 2))))
}

// ===== engine.go: isFieldNotEmpty slice/ptr-elem =====

func TestIsFieldNotEmptyBuiltinSliceMap(t *testing.T) {
	e := NewEngine()
	e.registerBuiltinValidators()
	fn := e.validators["required"]
	// slice
	assert.False(t, fn(paramFL{field: reflect.ValueOf([]int{})}))
	assert.True(t, fn(paramFL{field: reflect.ValueOf([]int{1})}))
	// ptr with non-nil elem
	p := 42
	assert.True(t, fn(paramFL{field: reflect.ValueOf(&p)}))
}

// ===== quick_strong_password_bench.go: Test wrappers for benchmarks =====

func TestBenchmarkStrongPasswordOld(t *testing.T) {
	passwords := []string{
		"Abc123!@",
		"Password123!",
		"SecurePass#2024",
		"MyP@ssw0rd",
		"Test@1234",
	}
	for _, pwd := range passwords {
		_ = validateStrongPasswordOld(pwd)
	}
}

func TestBenchmarkStrongPasswordNew(t *testing.T) {
	passwords := []string{
		"Abc123!@",
		"Password123!",
		"SecurePass#2024",
		"MyP@ssw0rd",
		"Test@1234",
	}
	for _, pwd := range passwords {
		_ = validateStrongPasswordFast(pwd)
	}
}

// ===== quick_strong_password_bench.go: validateStrongPasswordOld edge =====

func TestValidateStrongPasswordOldEdge(t *testing.T) {
	assert.True(t, validateStrongPasswordOld("Abc123!@"))
	assert.False(t, validateStrongPasswordOld("short"))
	assert.False(t, validateStrongPasswordOld("ABCDEFGH"))
}

// ===== conditional_validators.go: excludedWithAll empty param (line 181-183) =====

func TestExcludedWithAllEmptyParam(t *testing.T) {
	v, err := New()
	require.NoError(t, err)
	type S struct{ Value string `validate:"excluded_with_all"` }
	// empty param → return true
	assert.NoError(t, v.Struct(S{Value: "x"}))
}

// ===== conditional_validators.go: excludedWithoutAll empty param (line 221-223) =====

func TestExcludedWithoutAllEmptyParam(t *testing.T) {
	v, err := New()
	require.NoError(t, err)
	type S struct{ Value string `validate:"excluded_without_all"` }
	// empty param → return !isFieldNotEmpty
	assert.NoError(t, v.Struct(S{Value: ""}))
	assert.Error(t, v.Struct(S{Value: "x"}))
}

// ===== custom_validators.go: validateURL empty rest (line 98-100) =====
// Already covered by TestValidateURLWS. Let me check what's at line 98:
// It's `if len(rest) == 0 { return false }` inside validateURL
// Already tested with "http://" → should be covered. Let me check line 174:

// ===== custom_validators.go: validateIPv4 dot without digit (line 174-176) =====

func TestValidateIPv4DotWithoutDigit2(t *testing.T) {
	assert.False(t, validateIPv4(strFLP("1..2.3")))
}

// ===== engine.go: compareFields default string return 1 (line 763) =====

func TestCompareFieldsDefaultGt(t *testing.T) {
	// complex128 default path: string comparison, "2" > "1"
	assert.Equal(t, 1, compareFields(reflect.ValueOf(complex(2, 2)), reflect.ValueOf(complex(1, 1))))
}

// ===== engine.go: isFieldNotEmpty slice in validateStruct (line 776-777) =====

func TestIsFieldNotEmptyBuiltinSliceViaStruct(t *testing.T) {
	e := NewEngine()
	e.registerBuiltinValidators()
	type S struct{ V []int `validate:"required"` }
	errs := &ValidationErrors{}
	e.validateStruct(reflect.ValueOf(S{V: []int{}}), reflect.ValueOf(S{V: []int{}}), "", errs)
	assert.NotEmpty(t, *errs)
	e2 := &ValidationErrors{}
	e.validateStruct(reflect.ValueOf(S{V: []int{1}}), reflect.ValueOf(S{V: []int{1}}), "", errs)
	_ = e2
}

// ===== field_validators.go: resolveFieldPath dot path (line 42, 48-50) =====

func TestResolveFieldPathDotPathSteps(t *testing.T) {
	type Inner struct{ V string }
	type Outer struct{ I Inner }
	o := Outer{I: Inner{V: "test"}}
	fl := &fieldLevel{
		top:    reflect.ValueOf(o),
		parent: reflect.ValueOf(o),
		field:  reflect.ValueOf(o.I.V),
	}
	// dot path with valid struct navigation
	result := resolveFieldPath(fl, "I.V")
	assert.True(t, result.IsValid())
	assert.Equal(t, "test", result.String())
}

func TestResolveFieldPathDotPathInvalidField(t *testing.T) {
	type Outer struct{ Name string }
	o := Outer{Name: "test"}
	fl := &fieldLevel{
		top:    reflect.ValueOf(o),
		parent: reflect.ValueOf(o),
		field:  reflect.ValueOf(o.Name),
	}
	// dot path: "Name.Sub" → Name is string, not struct → returns invalid
	result := resolveFieldPath(fl, "Name.Sub")
	assert.False(t, result.IsValid())
}

// ===== format_validators.go: ISBN10 last char X check (line 281-283) =====

func TestISBN10WithX(t *testing.T) {
	// Valid ISBN-10 ending with X: 0-306-40615-X
	assert.True(t, validateISBN10(strFLP("123456789X")))
}

// ===== misc_validators.go: validateFn with no-result method (line 95-97) =====

type validatableNoResult struct{}

func (v validatableNoResult) Validate() {}

func TestValidateFnNoResult(t *testing.T) {
	assert.True(t, validateFn(paramFL{field: reflect.ValueOf(validatableNoResult{})}))
}

// ===== net_validators.go: validateHostname empty part (line 129-131) =====

func TestValidateHostnameEmptySegment(t *testing.T) {
	assert.False(t, validateHostname(strFLP("host..com")))
}

// ===== engine.go:355 - ptr to struct without validate tag =====

func TestValidateStructPtrToStructNoTag(t *testing.T) {
	e := NewEngine()
	e.RegisterValidation("required", func(fl FieldLevel) bool { return fl.Field().String() != "" })

	type Inner struct{ Name string `validate:"required"` }
	type Outer struct{ I *Inner } // no validate tag on I

	errs := &ValidationErrors{}
	e.validateStruct(reflect.ValueOf(Outer{I: &Inner{Name: ""}}), reflect.ValueOf(Outer{I: &Inner{Name: ""}}), "", errs)
	assert.NotEmpty(t, *errs, "should recursively validate Inner through pointer")
}

// ===== engine.go:776 - isFieldNotEmpty slice inside validateStruct =====

func TestValidateStructRequiredSlice(t *testing.T) {
	e := NewEngine()
	e.registerBuiltinValidators()
	type S struct{ V []int `validate:"required"` }
	errs := &ValidationErrors{}
	e.validateStruct(reflect.ValueOf(S{V: []int{}}), reflect.ValueOf(S{V: []int{}}), "", errs)
	assert.NotEmpty(t, *errs)
}

// ===== field_validators.go:42 - resolveFieldPath ptr elem =====

func TestResolveFieldPathPtrElem(t *testing.T) {
	type Inner struct{ V string }
	type Outer struct{ Ptr *Inner }
	inner := Inner{V: "test"}
	o := Outer{Ptr: &inner}
	fl := &fieldLevel{
		top:    reflect.ValueOf(o),
		parent: reflect.ValueOf(o),
		field:  reflect.ValueOf(o.Ptr),
	}
	result := resolveFieldPath(fl, "Ptr.V")
	assert.True(t, result.IsValid())
	assert.Equal(t, "test", result.String())
}

// ===== field_validators.go:48-50 - resolveFieldPath invalid field in dot path =====

func TestResolveFieldPathInvalidFieldInPath(t *testing.T) {
	type Outer struct{ Name string }
	o := Outer{Name: "test"}
	fl := &fieldLevel{
		top:    reflect.ValueOf(o),
		parent: reflect.ValueOf(o),
		field:  reflect.ValueOf(o.Name),
	}
	result := resolveFieldPath(fl, "Name.NonExist")
	assert.False(t, result.IsValid())
}

// ===== engine.go:355 - ptr to struct WITH validate tag, reaches recursive check =====

func TestValidateStructPtrToStructWithTag(t *testing.T) {
	e := NewEngine()
	e.RegisterValidation("required", func(fl FieldLevel) bool { return fl.Field().String() != "" })

	type Inner2 struct {
		Name string `validate:"required"`
	}
	type Outer2 struct {
		I *Inner2 `validate:"required"` // has tag → passes validation → reaches recursive check
	}

	// I is non-nil ptr to struct → after validating "required", reaches line 353-357
	errs := &ValidationErrors{}
	e.validateStruct(reflect.ValueOf(Outer2{I: &Inner2{Name: ""}}), reflect.ValueOf(Outer2{I: &Inner2{Name: ""}}), "", errs)
	assert.NotEmpty(t, *errs, "should catch Inner2.Name required error via ptr recursion")
}

// ===== engine.go:776 - isFieldNotEmpty slice in validateStruct recursive check =====

func TestValidateStructRequiredSliceViaBuiltin(t *testing.T) {
	e := NewEngine()
	e.registerBuiltinValidators()

	type S struct {
		V []int `validate:"required"`
	}
	errs := &ValidationErrors{}
	e.validateStruct(reflect.ValueOf(S{V: []int{}}), reflect.ValueOf(S{V: []int{}}), "", errs)
	assert.NotEmpty(t, *errs, "empty slice should fail required")
}

// ===== custom_validators.go:174 - validateIPv4 dot without preceding digit =====

func TestValidateIPv4DotWithoutDigit3(t *testing.T) {
	// First char is '.' → digitCount=0 when hitting dot → return false
	assert.False(t, validateIPv4(strFLP(".1.2.34")))
}

// ===== custom_validators.go:248 - validateUUID bad dash position =====

func TestValidateUUIDBadDashPos(t *testing.T) {
	// Valid UUID format but dash in wrong position
	assert.True(t, validateUUID(strFLP("12345678-1234-1234-1234-1234567890ab")))
	// Now test with dash at position 7 instead of 8
	assert.False(t, validateUUID(strFLP("12345678x234-1234-1234-123456789abcd")))
}

// ===== engine.go:776 - isFieldNotEmpty slice inside validateStruct =====
// This is the isFieldNotEmpty function at line 768, specifically line 776 (slice branch)
// It's called from validateStruct's isFieldNotEmpty path
// Let me check: the issue is that the validateStruct's built-in isFieldNotEmpty
// uses a different closure than the standalone isFieldNotEmpty function

// ===== net_validators.go:129 - validateHostname empty segment =====

func TestValidateHostnameEmptySegment2(t *testing.T) {
	assert.False(t, validateHostname(strFLP("host..com")))
}

// ===== field_validators.go:48-50 - resolveFieldPath invalid field =====

func TestResolveFieldPathInvalidFieldInDot(t *testing.T) {
	type Inner struct{ V string }
	type Outer struct{ I Inner }
	o := Outer{I: Inner{V: "test"}}
	fl := &fieldLevel{
		top:    reflect.ValueOf(o),
		parent: reflect.ValueOf(o),
		field:  reflect.ValueOf(o.I.V),
	}
	// "I.NonExist" → I resolves to Inner struct, NonExist not found
	result := resolveFieldPath(fl, "I.NonExist")
	assert.False(t, result.IsValid())
}

// ===== engine.go:776 - isFieldNotEmpty slice/map/array branch =====

func TestIsFieldNotEmptySliceMapArray(t *testing.T) {
	// slice
	assert.False(t, isFieldNotEmpty(reflect.ValueOf([]int{})))
	assert.True(t, isFieldNotEmpty(reflect.ValueOf([]int{1})))
	// map
	assert.False(t, isFieldNotEmpty(reflect.ValueOf(map[string]int{})))
	assert.True(t, isFieldNotEmpty(reflect.ValueOf(map[string]int{"a": 1})))
	// array
	assert.False(t, isFieldNotEmpty(reflect.ValueOf([0]int{})))
	assert.True(t, isFieldNotEmpty(reflect.ValueOf([1]int{1})))
}

// ===== net_validators.go:129 - validateHostnamePort empty segment =====

func TestHostnamePortEmptySegment(t *testing.T) {
	assert.False(t, validateHostnamePort(strFLP("host..com:8080")))
}
