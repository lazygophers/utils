package validator

import (
	"reflect"
	"strings"
)

// MiscValidators 返回所有杂项验证器注册表
func MiscValidators() map[string]ValidatorFunc {
	return map[string]ValidatorFunc{
		"oneof":      validateOneOf,
		"unique":     validateUnique,
		"isdefault":  validateIsDefault,
		"validateFn": validateFn,
	}
}

func validateOneOf(fl FieldLevel) bool {
	field := fl.Field()
	param := fl.Param()
	if param == "" {
		return false
	}

	value := getFieldValueAsString(field)
	for _, item := range strings.Split(param, ",") {
		if value == item {
			return true
		}
	}
	return false
}

func validateUnique(fl FieldLevel) bool {
	field := fl.Field()

	switch field.Kind() {
	case reflect.Slice, reflect.Array:
		seen := make(map[string]bool, field.Len())
		for i := 0; i < field.Len(); i++ {
			v := getFieldValueAsString(field.Index(i))
			if seen[v] {
				return false
			}
			seen[v] = true
		}
		return true
	case reflect.Map:
		seen := make(map[string]bool, field.Len())
		iter := field.MapRange()
		for iter.Next() {
			v := getFieldValueAsString(iter.Value())
			if seen[v] {
				return false
			}
			seen[v] = true
		}
		return true
	}
	return false
}

func validateIsDefault(fl FieldLevel) bool {
	field := fl.Field()
	if !field.IsValid() {
		return true
	}
	return field.IsZero()
}

func validateFn(fl FieldLevel) bool {
	field := fl.Field()

	// 支持 *T 和 T，只要实现了 Validate() error
	val := field
	if field.Kind() == reflect.Ptr {
		if field.IsNil() {
			return true
		}
		val = field.Elem()
	}

	method := val.MethodByName("Validate")
	if !method.IsValid() {
		// 也尝试从指针上调
		if field.CanAddr() {
			method = field.Addr().MethodByName("Validate")
		}
		if !method.IsValid() {
			return false
		}
	}

	results := method.Call(nil)
	if len(results) == 0 {
		return true
	}
	err, ok := results[0].Interface().(error)
	return !ok || err == nil
}
