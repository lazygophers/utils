package validator

import (
	"reflect"
	"strings"
)

// FieldValidators 返回所有跨字段验证器注册表
func FieldValidators() map[string]ValidatorFunc {
	return map[string]ValidatorFunc{
		"gtfield":       validateGTField,
		"gtefield":      validateGTEField,
		"ltfield":       validateLTField,
		"ltefield":      validateLTEField,
		"necsfield":     validateNECSField,
		"eqcsfield":     validateEQCSField,
		"gtcsfield":     validateGTCSField,
		"gtecsfield":    validateGTECSField,
		"ltcsfield":     validateLTCSField,
		"ltecsfield":    validateLTECSField,
		"fieldcontains": validateFieldContains,
		"fieldexcludes": validateFieldExcludes,
	}
}

// resolveFieldPath 支持点分路径的跨字段查找
// "Field" → 从当前结构体查找
// "Struct.Field" → 从顶层结构体沿路径查找
func resolveFieldPath(fl FieldLevel, path string) reflect.Value {
	if !strings.Contains(path, ".") {
		return fl.GetFieldByName(path)
	}

	// 点分路径：从顶层结构体逐级导航
	parts := strings.Split(path, ".")
	v := fl.Top()
	for _, p := range parts {
		if v.Kind() == reflect.Ptr {
			if v.IsNil() {
				return reflect.Value{}
			}
			v = v.Elem()
		}
		if v.Kind() != reflect.Struct {
			return reflect.Value{}
		}
		v = v.FieldByName(p)
		if !v.IsValid() {
			return reflect.Value{}
		}
	}
	return v
}

func validateGTField(fl FieldLevel) bool {
	target := fl.GetFieldByName(fl.Param())
	if !target.IsValid() {
		return false
	}
	return compareFields(fl.Field(), target) > 0
}

func validateGTEField(fl FieldLevel) bool {
	target := fl.GetFieldByName(fl.Param())
	if !target.IsValid() {
		return false
	}
	return compareFields(fl.Field(), target) >= 0
}

func validateLTField(fl FieldLevel) bool {
	target := fl.GetFieldByName(fl.Param())
	if !target.IsValid() {
		return false
	}
	return compareFields(fl.Field(), target) < 0
}

func validateLTEField(fl FieldLevel) bool {
	target := fl.GetFieldByName(fl.Param())
	if !target.IsValid() {
		return false
	}
	return compareFields(fl.Field(), target) <= 0
}

func validateEQCSField(fl FieldLevel) bool {
	target := resolveFieldPath(fl, fl.Param())
	if !target.IsValid() {
		return false
	}
	return compareFields(fl.Field(), target) == 0
}

func validateNECSField(fl FieldLevel) bool {
	target := resolveFieldPath(fl, fl.Param())
	if !target.IsValid() {
		return false
	}
	return compareFields(fl.Field(), target) != 0
}

func validateGTCSField(fl FieldLevel) bool {
	target := resolveFieldPath(fl, fl.Param())
	if !target.IsValid() {
		return false
	}
	return compareFields(fl.Field(), target) > 0
}

func validateGTECSField(fl FieldLevel) bool {
	target := resolveFieldPath(fl, fl.Param())
	if !target.IsValid() {
		return false
	}
	return compareFields(fl.Field(), target) >= 0
}

func validateLTCSField(fl FieldLevel) bool {
	target := resolveFieldPath(fl, fl.Param())
	if !target.IsValid() {
		return false
	}
	return compareFields(fl.Field(), target) < 0
}

func validateLTECSField(fl FieldLevel) bool {
	target := resolveFieldPath(fl, fl.Param())
	if !target.IsValid() {
		return false
	}
	return compareFields(fl.Field(), target) <= 0
}

func validateFieldContains(fl FieldLevel) bool {
	param := fl.Param()
	if param == "" {
		return false
	}
	value := fl.Field().String()
	// param 可以是多个字符，检查是否包含其中任一
	return strings.ContainsAny(value, param)
}

func validateFieldExcludes(fl FieldLevel) bool {
	param := fl.Param()
	if param == "" {
		return true
	}
	value := fl.Field().String()
	return !strings.ContainsAny(value, param)
}
