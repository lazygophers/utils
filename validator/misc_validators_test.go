package validator

import "errors"
import (
	"reflect"
	"testing"
)

func TestValidateOneOf(t *testing.T) {
	fl := paramFL{field: reflect.ValueOf("apple"), param: "apple,banana,cherry"}
	if !validateOneOf(fl) {
		t.Error("apple should be in list")
	}
	fl2 := paramFL{field: reflect.ValueOf("grape"), param: "apple,banana,cherry"}
	if validateOneOf(fl2) {
		t.Error("grape should not be in list")
	}
	fl3 := paramFL{field: reflect.ValueOf("apple"), param: ""}
	if validateOneOf(fl3) {
		t.Error("empty param should fail")
	}
}

func TestValidateUnique(t *testing.T) {
	// Slice unique
	fl := paramFL{field: reflect.ValueOf([]string{"a", "b", "c"}), param: ""}
	if !validateUnique(fl) {
		t.Error("unique slice should pass")
	}
	fl2 := paramFL{field: reflect.ValueOf([]string{"a", "b", "a"}), param: ""}
	if validateUnique(fl2) {
		t.Error("duplicate slice should fail")
	}
	// Map unique
	fl3 := paramFL{field: reflect.ValueOf(map[string]int{"a": 1, "b": 2}), param: ""}
	if !validateUnique(fl3) {
		t.Error("unique map should pass")
	}
	// Map with duplicate values
	fl4 := paramFL{field: reflect.ValueOf(map[string]int{"a": 1, "b": 1}), param: ""}
	if validateUnique(fl4) {
		t.Error("map with duplicate values should fail")
	}
	// Non-slice/map
	fl5 := paramFL{field: reflect.ValueOf("hello"), param: ""}
	if validateUnique(fl5) {
		t.Error("string should fail unique check")
	}
}

func TestValidateIsDefault(t *testing.T) {
	fl := paramFL{field: reflect.ValueOf(0), param: ""}
	if !validateIsDefault(fl) {
		t.Error("zero int should be default")
	}
	fl2 := paramFL{field: reflect.ValueOf(42), param: ""}
	if validateIsDefault(fl2) {
		t.Error("42 should not be default")
	}
}

type validatable struct{ err error }

func (v validatable) Validate() error { return v.err }

type unvalidatable struct{}

func TestValidateFn(t *testing.T) {
	// Value with Validate returning nil
	fl := paramFL{field: reflect.ValueOf(validatable{err: nil}), param: ""}
	if !validateFn(fl) {
		t.Error("validatable with nil error should pass")
	}
	// Value with Validate returning error
	fl2 := paramFL{field: reflect.ValueOf(validatable{err: assertAnError}), param: ""}
	if validateFn(fl2) {
		t.Error("validatable with error should fail")
	}
	// Pointer to validatable
	v := validatable{err: nil}
	fl3 := paramFL{field: reflect.ValueOf(&v), param: ""}
	if !validateFn(fl3) {
		t.Error("pointer to validatable should pass")
	}
	// Nil pointer
	var pv *validatable
	fl4 := paramFL{field: reflect.ValueOf(pv), param: ""}
	if !validateFn(fl4) {
		t.Error("nil pointer should pass (treated as valid)")
	}
	// Type without Validate
	fl5 := paramFL{field: reflect.ValueOf(42), param: ""}
	if validateFn(fl5) {
		t.Error("int should fail validateFn")
	}
}

var assertAnError = errors.New("test error")
