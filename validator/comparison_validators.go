package validator

import (
	"reflect"
	"strconv"
	"strings"
)

// ComparisonValidators 返回所有比较运算验证器注册表
func ComparisonValidators() map[string]ValidatorFunc {
	return map[string]ValidatorFunc{
		"gt":             validateGT,
		"gte":            validateGTE,
		"lt":             validateLT,
		"lte":            validateLTE,
		"eq_ignore_case": validateEqIgnoreCase,
		"ne_ignore_case": validateNeIgnoreCase,
	}
}

func validateGT(fl FieldLevel) bool {
	return compareThreshold(fl, func(v, t float64) bool { return v > t })
}

func validateGTE(fl FieldLevel) bool {
	return compareThreshold(fl, func(v, t float64) bool { return v >= t })
}

func validateLT(fl FieldLevel) bool {
	return compareThreshold(fl, func(v, t float64) bool { return v < t })
}

func validateLTE(fl FieldLevel) bool {
	return compareThreshold(fl, func(v, t float64) bool { return v <= t })
}

func compareThreshold(fl FieldLevel, cmp func(float64, float64) bool) bool {
	field := fl.Field()
	param := fl.Param()

	threshold, err := strconv.ParseFloat(param, 64)
	if err != nil {
		return false
	}

	switch field.Kind() {
	case reflect.String:
		return cmp(float64(len(field.String())), threshold)
	case reflect.Slice, reflect.Map, reflect.Array:
		return cmp(float64(field.Len()), threshold)
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return cmp(float64(field.Int()), threshold)
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return cmp(float64(field.Uint()), threshold)
	case reflect.Float32, reflect.Float64:
		return cmp(field.Float(), threshold)
	}
	return false
}

func validateEqIgnoreCase(fl FieldLevel) bool {
	field := fl.Field()
	param := fl.Param()
	if param == "" {
		return false
	}
	return strings.EqualFold(field.String(), param)
}

func validateNeIgnoreCase(fl FieldLevel) bool {
	field := fl.Field()
	param := fl.Param()
	if param == "" {
		return false
	}
	return !strings.EqualFold(field.String(), param)
}
