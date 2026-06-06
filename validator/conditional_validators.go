package validator

import (
	"strings"
)

// ConditionalValidators 返回所有条件验证器注册表
func ConditionalValidators() map[string]ValidatorFunc {
	return map[string]ValidatorFunc{
		"required_unless":      validateRequiredUnless,
		"required_with_all":    validateRequiredWithAll,
		"required_without_all": validateRequiredWithoutAll,
		"excluded_if":          validateExcludedIf,
		"excluded_unless":      validateExcludedUnless,
		"excluded_with":        validateExcludedWith,
		"excluded_with_all":    validateExcludedWithAll,
		"excluded_without":     validateExcludedWithout,
		"excluded_without_all": validateExcludedWithoutAll,
	}
}

// splitFieldList 将逗号分隔的参数拆分为字段名列表
func splitFieldList(param string) []string {
	if param == "" {
		return nil
	}
	return strings.Split(param, ",")
}

// validateRequiredUnless 当指定字段不等于某值时必填
// 格式: required_unless=Field1=Value1,Field2=Value2
// 除非所有条件都满足，否则当前字段必填
func validateRequiredUnless(fl FieldLevel) bool {
	currentField := fl.Field()
	param := fl.Param()
	if param == "" {
		return isFieldNotEmpty(currentField)
	}

	conditions := splitFieldList(param)
	allConditionsMet := true

	for _, cond := range conditions {
		parts := strings.SplitN(cond, "=", 2)
		if len(parts) != 2 {
			continue
		}
		targetField := fl.GetFieldByName(parts[0])
		if !targetField.IsValid() {
			continue
		}
		if getFieldValueAsString(targetField) != parts[1] {
			allConditionsMet = false
			break
		}
	}

	if allConditionsMet {
		return true // 所有条件满足，不需要必填
	}
	return isFieldNotEmpty(currentField)
}

// validateRequiredWithAll 当所有指定字段都有值时必填
func validateRequiredWithAll(fl FieldLevel) bool {
	currentField := fl.Field()
	param := fl.Param()
	if param == "" {
		return isFieldNotEmpty(currentField)
	}

	fields := splitFieldList(param)
	allNotEmpty := true
	for _, name := range fields {
		target := fl.GetFieldByName(name)
		if !target.IsValid() || !isFieldNotEmpty(target) {
			allNotEmpty = false
			break
		}
	}

	if allNotEmpty {
		return isFieldNotEmpty(currentField)
	}
	return true
}

// validateRequiredWithoutAll 当所有指定字段都没有值时必填
func validateRequiredWithoutAll(fl FieldLevel) bool {
	currentField := fl.Field()
	param := fl.Param()
	if param == "" {
		return isFieldNotEmpty(currentField)
	}

	fields := splitFieldList(param)
	allEmpty := true
	for _, name := range fields {
		target := fl.GetFieldByName(name)
		if target.IsValid() && isFieldNotEmpty(target) {
			allEmpty = false
			break
		}
	}

	if allEmpty {
		return isFieldNotEmpty(currentField)
	}
	return true
}

// validateExcludedIf 当条件满足时当前字段必须为空
func validateExcludedIf(fl FieldLevel) bool {
	currentField := fl.Field()
	param := fl.Param()
	if param == "" {
		return true
	}

	parts := strings.SplitN(param, "=", 2)
	if len(parts) != 2 {
		return true
	}

	target := fl.GetFieldByName(parts[0])
	if !target.IsValid() {
		return true
	}

	if getFieldValueAsString(target) == parts[1] {
		return !isFieldNotEmpty(currentField)
	}
	return true
}

// validateExcludedUnless 除非条件满足，否则当前字段必须为空
func validateExcludedUnless(fl FieldLevel) bool {
	currentField := fl.Field()
	param := fl.Param()
	if param == "" {
		return !isFieldNotEmpty(currentField)
	}

	parts := strings.SplitN(param, "=", 2)
	if len(parts) != 2 {
		return !isFieldNotEmpty(currentField)
	}

	target := fl.GetFieldByName(parts[0])
	if !target.IsValid() {
		return !isFieldNotEmpty(currentField)
	}

	if getFieldValueAsString(target) == parts[1] {
		return true
	}
	return !isFieldNotEmpty(currentField)
}

// validateExcludedWith 当任一指定字段有值时当前字段必须为空
func validateExcludedWith(fl FieldLevel) bool {
	currentField := fl.Field()
	param := fl.Param()
	if param == "" {
		return true
	}

	for _, name := range splitFieldList(param) {
		target := fl.GetFieldByName(name)
		if target.IsValid() && isFieldNotEmpty(target) {
			return !isFieldNotEmpty(currentField)
		}
	}
	return true
}

// validateExcludedWithAll 当所有指定字段都有值时当前字段必须为空
func validateExcludedWithAll(fl FieldLevel) bool {
	currentField := fl.Field()
	param := fl.Param()
	if param == "" {
		return true
	}

	allNotEmpty := true
	for _, name := range splitFieldList(param) {
		target := fl.GetFieldByName(name)
		if !target.IsValid() || !isFieldNotEmpty(target) {
			allNotEmpty = false
			break
		}
	}

	if allNotEmpty {
		return !isFieldNotEmpty(currentField)
	}
	return true
}

// validateExcludedWithout 当任一指定字段无值时当前字段必须为空
func validateExcludedWithout(fl FieldLevel) bool {
	currentField := fl.Field()
	param := fl.Param()
	if param == "" {
		return !isFieldNotEmpty(currentField)
	}

	for _, name := range splitFieldList(param) {
		target := fl.GetFieldByName(name)
		if !target.IsValid() || !isFieldNotEmpty(target) {
			return !isFieldNotEmpty(currentField)
		}
	}
	return true
}

// validateExcludedWithoutAll 当所有指定字段都无值时当前字段必须为空
func validateExcludedWithoutAll(fl FieldLevel) bool {
	currentField := fl.Field()
	param := fl.Param()
	if param == "" {
		return !isFieldNotEmpty(currentField)
	}

	allEmpty := true
	for _, name := range splitFieldList(param) {
		target := fl.GetFieldByName(name)
		if target.IsValid() && isFieldNotEmpty(target) {
			allEmpty = false
			break
		}
	}

	if allEmpty {
		return !isFieldNotEmpty(currentField)
	}
	return true
}
