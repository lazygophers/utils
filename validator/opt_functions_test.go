package validator

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestOpt2InlineMap(t *testing.T) {
	e := NewEngine()
	e.validators["test_tag"] = func(fl FieldLevel) bool { return true }
	fl := &fieldLevel{field: reflect.ValueOf("x")}
	assert.True(t, e.validateField_Opt2_InlineMap(fl, "test_tag"))
	assert.True(t, e.validateField_Opt2_InlineMap(fl, "nonexistent"))
}

func TestOpt3SingleLookup(t *testing.T) {
	e := NewEngine()
	e.validators["test_tag"] = func(fl FieldLevel) bool { return false }
	fl := &fieldLevel{field: reflect.ValueOf("x")}
	assert.False(t, e.validateField_Opt3_SingleLookup(fl, "test_tag"))
	assert.True(t, e.validateField_Opt3_SingleLookup(fl, "nonexistent"))
}

func TestOpt5HotPathSwitch(t *testing.T) {
	e := NewEngine()
	e.validators["required"] = func(fl FieldLevel) bool { return fl.Field().String() != "" }
	e.validators["email"] = func(fl FieldLevel) bool { return true }
	e.validators["min"] = func(fl FieldLevel) bool { return true }
	e.validators["max"] = func(fl FieldLevel) bool { return true }
	e.validators["len"] = func(fl FieldLevel) bool { return true }
	e.validators["alpha"] = func(fl FieldLevel) bool { return true }
	e.validators["alphanum"] = func(fl FieldLevel) bool { return true }
	e.validators["url"] = func(fl FieldLevel) bool { return true }
	e.validators["custom"] = func(fl FieldLevel) bool { return false }

	fl := &fieldLevel{field: reflect.ValueOf("x")}
	// hot path tags
	assert.True(t, e.validateField_Opt5_HotPathSwitch(fl, "required"))
	assert.True(t, e.validateField_Opt5_HotPathSwitch(fl, "email"))
	assert.True(t, e.validateField_Opt5_HotPathSwitch(fl, "min"))
	assert.True(t, e.validateField_Opt5_HotPathSwitch(fl, "max"))
	assert.True(t, e.validateField_Opt5_HotPathSwitch(fl, "len"))
	assert.True(t, e.validateField_Opt5_HotPathSwitch(fl, "alpha"))
	assert.True(t, e.validateField_Opt5_HotPathSwitch(fl, "alphanum"))
	assert.True(t, e.validateField_Opt5_HotPathSwitch(fl, "url"))
	// default
	assert.False(t, e.validateField_Opt5_HotPathSwitch(fl, "custom"))
	assert.True(t, e.validateField_Opt5_HotPathSwitch(fl, "nonexistent"))
}

func TestOpt6FullSwitch(t *testing.T) {
	e := NewEngine()
	for _, tag := range []string{"required", "email", "min", "max", "len", "alpha", "alphanum",
		"url", "numeric", "eq", "ne", "eqfield", "nefield", "required_if", "required_with", "required_without"} {
		e.validators[tag] = func(fl FieldLevel) bool { return true }
	}
	fl := &fieldLevel{field: reflect.ValueOf("x")}
	for _, tag := range []string{"required", "email", "min", "max", "len", "alpha", "alphanum",
		"url", "numeric", "eq", "ne", "eqfield", "nefield", "required_if", "required_with", "required_without"} {
		assert.True(t, e.validateField_Opt6_FullSwitch(fl, tag), tag)
	}
	assert.True(t, e.validateField_Opt6_FullSwitch(fl, "nonexistent"))
}

func TestOpt11InlinedValidators(t *testing.T) {
	e := NewEngine()
	e.validators["custom"] = func(fl FieldLevel) bool { return false }

	// required: string non-empty
	fl := &fieldLevel{field: reflect.ValueOf("x")}
	assert.True(t, e.validateField_Opt11_InlinedValidators(fl, "required"))
	// required: string empty
	fl2 := &fieldLevel{field: reflect.ValueOf("")}
	assert.False(t, e.validateField_Opt11_InlinedValidators(fl2, "required"))
	// required: slice non-empty
	fl3 := &fieldLevel{field: reflect.ValueOf([]int{1})}
	assert.True(t, e.validateField_Opt11_InlinedValidators(fl3, "required"))
	// required: ptr non-nil
	p := "x"
	fl4 := &fieldLevel{field: reflect.ValueOf(&p)}
	assert.True(t, e.validateField_Opt11_InlinedValidators(fl4, "required"))
	// required: int non-zero
	fl5 := &fieldLevel{field: reflect.ValueOf(42)}
	assert.True(t, e.validateField_Opt11_InlinedValidators(fl5, "required"))

	// email empty → true
	assert.True(t, e.validateField_Opt11_InlinedValidators(&fieldLevel{field: reflect.ValueOf("")}, "email"))
	// alpha empty → true
	assert.True(t, e.validateField_Opt11_InlinedValidators(&fieldLevel{field: reflect.ValueOf("")}, "alpha"))
	// alphanum empty → true
	assert.True(t, e.validateField_Opt11_InlinedValidators(&fieldLevel{field: reflect.ValueOf("")}, "alphanum"))
	// url empty → true
	assert.True(t, e.validateField_Opt11_InlinedValidators(&fieldLevel{field: reflect.ValueOf("")}, "url"))

	// default: registered
	assert.False(t, e.validateField_Opt11_InlinedValidators(&fieldLevel{field: reflect.ValueOf("x")}, "custom"))
	// default: not found
	assert.True(t, e.validateField_Opt11_InlinedValidators(&fieldLevel{field: reflect.ValueOf("x")}, "unknown"))
}

func TestOpt13GotoOptimized(t *testing.T) {
	e := NewEngine()
	e.validators["test"] = func(fl FieldLevel) bool { return true }
	fl := &fieldLevel{field: reflect.ValueOf("x")}
	assert.True(t, e.validateField_Opt13_GotoOptimized(fl, "test"))
	assert.True(t, e.validateField_Opt13_GotoOptimized(fl, "nonexistent"))
}
