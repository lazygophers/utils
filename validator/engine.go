package validator

import (
	"fmt"
	"reflect"
	"regexp"
	"strconv"
	"strings"
)

// Engine 自定义验证引擎
type Engine struct {
	validators    map[string]ValidatorFunc
	tagName       string
	fieldNameFunc func(reflect.StructField) string
}

// ValidatorFunc 验证函数类型
type ValidatorFunc func(fl FieldLevel) bool

// FieldLevel 字段级别接口，提供验证时的上下文信息
type FieldLevel interface {
	// Top 获取顶级结构体
	Top() reflect.Value
	// Parent 获取父级结构体
	Parent() reflect.Value
	// Field 获取当前字段值
	Field() reflect.Value
	// FieldName 获取字段名
	FieldName() string
	// StructFieldName 获取结构体字段名
	StructFieldName() string
	// Param 获取验证标签参数
	Param() string
	// GetTag 获取指定的标签值
	GetTag(key string) string
}

// fieldLevel 字段级别实现
type fieldLevel struct {
	top             reflect.Value
	parent          reflect.Value
	field           reflect.Value
	fieldName       string
	structFieldName string
	param           string
	structField     reflect.StructField
}

func (fl *fieldLevel) Top() reflect.Value {
	return fl.top
}

func (fl *fieldLevel) Parent() reflect.Value {
	return fl.parent
}

func (fl *fieldLevel) Field() reflect.Value {
	return fl.field
}

func (fl *fieldLevel) FieldName() string {
	return fl.fieldName
}

func (fl *fieldLevel) StructFieldName() string {
	return fl.structFieldName
}

func (fl *fieldLevel) Param() string {
	return fl.param
}

func (fl *fieldLevel) GetTag(key string) string {
	return fl.structField.Tag.Get(key)
}

// NewEngine 创建新的验证引擎
func NewEngine() *Engine {
	e := &Engine{
		validators:    make(map[string]ValidatorFunc),
		tagName:       "validate",
		fieldNameFunc: defaultFieldNameFunc,
	}

	// 注册内置验证器
	e.registerBuiltinValidators()

	return e
}

// SetFieldNameFunc 设置字段名称解析函数
func (e *Engine) SetFieldNameFunc(fn func(reflect.StructField) string) {
	if fn != nil {
		e.fieldNameFunc = fn
	}
}

// RegisterValidation 注册验证规则
func (e *Engine) RegisterValidation(tag string, fn ValidatorFunc) error {
	if tag == "" {
		return fmt.Errorf("validation tag cannot be empty")
	}
	if fn == nil {
		return fmt.Errorf("validation function cannot be nil")
	}

	e.validators[tag] = fn
	return nil
}

// SetTagName 设置验证标签名称
func (e *Engine) SetTagName(name string) {
	e.tagName = name
}

// Struct 验证结构体
func (e *Engine) Struct(s interface{}) error {
	rv := reflect.ValueOf(s)
	if rv.Kind() == reflect.Ptr {
		rv = rv.Elem()
	}

	if rv.Kind() != reflect.Struct {
		return fmt.Errorf("expected struct, got %s", rv.Kind())
	}

	var errors ValidationErrors
	e.validateStruct(rv, rv, "", &errors)

	if len(errors) > 0 {
		return errors
	}
	return nil
}

// Var 验证单个变量
func (e *Engine) Var(field interface{}, tag string) error {
	rv := reflect.ValueOf(field)

	// 解析验证标签
	rules := e.parseTag(tag)
	if len(rules) == 0 {
		return nil
	}

	fl := &fieldLevel{
		top:             rv,
		parent:          rv,
		field:           rv,
		fieldName:       "var",
		structFieldName: "var",
	}

	for _, rule := range rules {
		fl.param = rule.param
		if !e.validateField(fl, rule.tag) {
			return &FieldError{
				Field:   "var",
				Tag:     rule.tag,
				Value:   field,
				Param:   rule.param,
				Message: fmt.Sprintf("validation failed for tag '%s'", rule.tag),
			}
		}
	}

	return nil
}

// validateStruct 验证结构体内部实现
func (e *Engine) validateStruct(top, current reflect.Value, namespace string, errors *ValidationErrors) {
	rt := current.Type()

	for i := 0; i < current.NumField(); i++ {
		field := current.Field(i)
		fieldType := rt.Field(i)

		// 跳过非导出字段
		if !fieldType.IsExported() {
			continue
		}

		fieldName := fieldType.Name
		if namespace != "" {
			fieldName = namespace + "." + fieldName
		}

		// 获取验证标签
		tag := fieldType.Tag.Get(e.tagName)
		if tag == "" || tag == "-" {
			// 如果没有验证标签，但是字段是结构体，递归验证
			if field.Kind() == reflect.Struct {
				e.validateStruct(top, field, fieldName, errors)
			} else if field.Kind() == reflect.Ptr && !field.IsNil() && field.Elem().Kind() == reflect.Struct {
				e.validateStruct(top, field.Elem(), fieldName, errors)
			}
			continue
		}

		// 解析验证规则
		rules := e.parseTag(tag)

		// 获取字段显示名称
		displayName := e.fieldNameFunc(fieldType)

		fl := &fieldLevel{
			top:             top,
			parent:          current,
			field:           field,
			fieldName:       displayName,
			structFieldName: fieldType.Name,
			structField:     fieldType,
		}

		for _, rule := range rules {
			fl.param = rule.param
			if !e.validateField(fl, rule.tag) {
				fieldError := &FieldError{
					Field:       displayName,
					Tag:         rule.tag,
					Value:       field.Interface(),
					Param:       rule.param,
					ActualTag:   rule.tag,
					Namespace:   fieldName,
					StructField: fieldType.Name,
					Message:     fmt.Sprintf("validation failed for tag '%s'", rule.tag),
				}
				*errors = append(*errors, fieldError)
			}
		}

		// 递归验证嵌套结构体
		if field.Kind() == reflect.Struct {
			e.validateStruct(top, field, fieldName, errors)
		} else if field.Kind() == reflect.Ptr && !field.IsNil() && field.Elem().Kind() == reflect.Struct {
			e.validateStruct(top, field.Elem(), fieldName, errors)
		}
	}
}

// validateField 验证单个字段
func (e *Engine) validateField(fl FieldLevel, tag string) bool {
	validator, exists := e.validators[tag]
	if !exists {
		// 如果验证器不存在，默认返回true（忽略未知的验证标签）
		return true
	}

	return validator(fl)
}

// validationRule 验证规则
type validationRule struct {
	tag   string
	param string
}

// parseTag 解析验证标签
func (e *Engine) parseTag(tag string) []validationRule {
	var rules []validationRule

	// 按逗号分割验证规则
	parts := strings.Split(tag, ",")

	for _, part := range parts {
		part = strings.TrimSpace(part)
		if part == "" {
			continue
		}

		// 解析参数
		if idx := strings.Index(part, "="); idx != -1 {
			ruleName := strings.TrimSpace(part[:idx])
			param := strings.TrimSpace(part[idx+1:])
			rules = append(rules, validationRule{tag: ruleName, param: param})
		} else {
			rules = append(rules, validationRule{tag: part, param: ""})
		}
	}

	return rules
}

// defaultFieldNameFunc 默认字段名称解析函数（优先使用JSON标签）
func defaultFieldNameFunc(field reflect.StructField) string {
	// 优先使用 json tag
	if jsonTag := field.Tag.Get("json"); jsonTag != "" && jsonTag != "-" {
		// 处理 json:"name,omitempty" 格式
		if parts := strings.Split(jsonTag, ","); len(parts) > 0 && parts[0] != "" {
			return parts[0]
		}
	}

	// 回退到字段名
	return field.Name
}

// structFieldNameFunc 结构体字段名称解析函数（不使用JSON标签）
func structFieldNameFunc(field reflect.StructField) string {
	return field.Name
}

// registerBuiltinValidators 注册内置验证器
func (e *Engine) registerBuiltinValidators() {
	// 必填验证
	e.RegisterValidation("required", func(fl FieldLevel) bool {
		field := fl.Field()
		switch field.Kind() {
		case reflect.String:
			return field.String() != ""
		case reflect.Slice, reflect.Map, reflect.Array:
			return field.Len() > 0
		case reflect.Ptr, reflect.Interface:
			return !field.IsNil()
		default:
			return field.IsValid() && !field.IsZero()
		}
	})

	// 邮箱验证
	e.RegisterValidation("email", func(fl FieldLevel) bool {
		email := fl.Field().String()
		if email == "" {
			return true // 空值不验证，用required控制
		}
		matched, _ := regexp.MatchString(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`, email)
		return matched
	})

	// URL验证
	e.RegisterValidation("url", func(fl FieldLevel) bool {
		url := fl.Field().String()
		if url == "" {
			return true
		}
		matched, _ := regexp.MatchString(`^(https?|ftp)://[^\s/$.?#].[^\s]*$`, url)
		return matched
	})

	// 最小值验证
	e.RegisterValidation("min", func(fl FieldLevel) bool {
		field := fl.Field()
		param := fl.Param()

		min, err := strconv.ParseFloat(param, 64)
		if err != nil {
			return false
		}

		switch field.Kind() {
		case reflect.String:
			return float64(len(field.String())) >= min
		case reflect.Slice, reflect.Map, reflect.Array:
			return float64(field.Len()) >= min
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			return float64(field.Int()) >= min
		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
			return float64(field.Uint()) >= min
		case reflect.Float32, reflect.Float64:
			return field.Float() >= min
		}
		return false
	})

	// 最大值验证
	e.RegisterValidation("max", func(fl FieldLevel) bool {
		field := fl.Field()
		param := fl.Param()

		max, err := strconv.ParseFloat(param, 64)
		if err != nil {
			return false
		}

		switch field.Kind() {
		case reflect.String:
			return float64(len(field.String())) <= max
		case reflect.Slice, reflect.Map, reflect.Array:
			return float64(field.Len()) <= max
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			return float64(field.Int()) <= max
		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
			return float64(field.Uint()) <= max
		case reflect.Float32, reflect.Float64:
			return field.Float() <= max
		}
		return false
	})

	// 长度验证
	e.RegisterValidation("len", func(fl FieldLevel) bool {
		field := fl.Field()
		param := fl.Param()

		length, err := strconv.Atoi(param)
		if err != nil {
			return false
		}

		switch field.Kind() {
		case reflect.String:
			return len(field.String()) == length
		case reflect.Slice, reflect.Map, reflect.Array:
			return field.Len() == length
		}
		return false
	})

	// 数字验证
	e.RegisterValidation("numeric", func(fl FieldLevel) bool {
		value := fl.Field().String()
		if value == "" {
			return true
		}
		_, err := strconv.ParseFloat(value, 64)
		return err == nil
	})

	// 字母验证
	e.RegisterValidation("alpha", func(fl FieldLevel) bool {
		value := fl.Field().String()
		if value == "" {
			return true
		}
		matched, _ := regexp.MatchString(`^[a-zA-Z]+$`, value)
		return matched
	})

	// 字母数字验证
	e.RegisterValidation("alphanum", func(fl FieldLevel) bool {
		value := fl.Field().String()
		if value == "" {
			return true
		}
		matched, _ := regexp.MatchString(`^[a-zA-Z0-9]+$`, value)
		return matched
	})

	// 等于验证
	e.RegisterValidation("eq", func(fl FieldLevel) bool {
		field := fl.Field()
		param := fl.Param()

		switch field.Kind() {
		case reflect.String:
			return field.String() == param
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			if val, err := strconv.ParseInt(param, 10, 64); err == nil {
				return field.Int() == val
			}
		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
			if val, err := strconv.ParseUint(param, 10, 64); err == nil {
				return field.Uint() == val
			}
		case reflect.Float32, reflect.Float64:
			if val, err := strconv.ParseFloat(param, 64); err == nil {
				return field.Float() == val
			}
		}
		return false
	})

	// 不等于验证
	e.RegisterValidation("ne", func(fl FieldLevel) bool {
		// 复用eq验证，然后取反
		eqValidator := e.validators["eq"]
		return !eqValidator(fl)
	})
}
