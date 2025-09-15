package validator

import (
	"fmt"
	"reflect"
	"strings"
	"sync"
)

var (
	defaultValidator *Validator
	once             sync.Once
)

// Validator 自定义验证器
type Validator struct {
	engine   *Engine
	locale   string
	useJSON  bool
	mu       sync.RWMutex
	messages map[string]string
}

// New 创建新的验证器实例
func New(opts ...Option) (*Validator, error) {
	v := &Validator{
		engine:   NewEngine(),
		locale:   "en",
		useJSON:  true,
		messages: make(map[string]string),
	}

	// 应用选项
	for _, opt := range opts {
		opt(v)
	}

	// 设置字段名称解析函数
	v.updateFieldNameFunc()
	
	// 注册默认验证规则
	if err := v.registerDefaultValidators(); err != nil {
		return nil, fmt.Errorf("failed to register default validators: %w", err)
	}

	return v, nil
}

// Default 获取默认验证器实例
func Default() *Validator {
	once.Do(func() {
		v, err := New()
		if err != nil {
			// 如果创建默认验证器失败，创建一个基础版本
			defaultValidator = &Validator{
				engine:   NewEngine(),
				locale:   "en",
				useJSON:  true,
				messages: make(map[string]string),
			}
		} else {
			defaultValidator = v
		}
	})
	return defaultValidator
}

// SetLocale 设置语言地区
func (v *Validator) SetLocale(locale string) {
	v.mu.Lock()
	defer v.mu.Unlock()
	v.locale = locale
}

// GetLocale 获取当前语言地区
func (v *Validator) GetLocale() string {
	v.mu.RLock()
	defer v.mu.RUnlock()
	return v.locale
}

// SetUseJSON 设置是否优先使用 JSON 字段名
func (v *Validator) SetUseJSON(useJSON bool) {
	v.mu.Lock()
	defer v.mu.Unlock()
	v.useJSON = useJSON
	v.updateFieldNameFunc()
}

// Struct 验证结构体
func (v *Validator) Struct(s interface{}) error {
	err := v.engine.Struct(s)
	if err != nil {
		if validationErrors, ok := err.(ValidationErrors); ok {
			return v.translateValidationErrors(validationErrors)
		}
		return err
	}
	return nil
}

// Var 验证单个变量
func (v *Validator) Var(field interface{}, tag string) error {
	err := v.engine.Var(field, tag)
	if err != nil {
		if fieldError, ok := err.(*FieldError); ok {
			fieldError.Message = v.translateFieldError(fieldError)
			return fieldError
		}
		return err
	}
	return nil
}

// RegisterValidation 注册自定义验证规则
func (v *Validator) RegisterValidation(tag string, fn ValidatorFunc) error {
	return v.engine.RegisterValidation(tag, fn)
}

// RegisterTranslation 注册翻译
func (v *Validator) RegisterTranslation(locale, tag, translation string) {
	v.mu.Lock()
	defer v.mu.Unlock()
	key := fmt.Sprintf("%s.%s", locale, tag)
	v.messages[key] = translation
}

// translateValidationErrors 翻译验证错误
func (v *Validator) translateValidationErrors(validationErrors ValidationErrors) error {
	for _, err := range validationErrors {
		err.Message = v.translateFieldError(err)
	}
	return validationErrors
}

// translateFieldError 翻译字段错误
func (v *Validator) translateFieldError(err *FieldError) string {
	v.mu.RLock()
	locale := v.locale
	v.mu.RUnlock()

	// 获取本地化配置
	localeConfig, ok := GetLocaleConfig(locale)
	if !ok {
		if enConfig, enOk := GetLocaleConfig("en"); enOk {
			localeConfig = enConfig
		} else {
			// 如果连英文配置都没有，返回默认格式
			return fmt.Sprintf("%s failed validation for tag '%s'", err.Field, err.Tag)
		}
	}

	// 构建翻译键
	key := fmt.Sprintf("%s.%s", locale, err.Tag)

	v.mu.RLock()
	if msg, exists := v.messages[key]; exists {
		v.mu.RUnlock()
		return v.formatMessage(msg, err)
	}
	v.mu.RUnlock()

	// 使用默认消息模板
	if template, exists := localeConfig.Messages[err.Tag]; exists {
		return v.formatMessage(template, err)
	}

	// 最后回退到英文默认消息
	if locale != "en" {
		if englishConfig, ok := GetLocaleConfig("en"); ok {
			if template, exists := englishConfig.Messages[err.Tag]; exists {
				return v.formatMessage(template, err)
			}
		}
	}

	// 如果没有找到翻译，返回默认格式
	return fmt.Sprintf("%s failed validation for tag '%s'", err.Field, err.Tag)
}

// formatMessage 格式化错误消息
func (v *Validator) formatMessage(template string, err *FieldError) string {
	message := template

	// 替换占位符
	message = strings.ReplaceAll(message, "{field}", err.Field)
	message = strings.ReplaceAll(message, "{tag}", err.Tag)
	message = strings.ReplaceAll(message, "{param}", err.Param)

	// 处理值的显示
	if err.Value != nil {
		message = strings.ReplaceAll(message, "{value}", fmt.Sprintf("%v", err.Value))
	} else {
		message = strings.ReplaceAll(message, "{value}", "")
	}

	return message
}

// updateFieldNameFunc 更新字段名称解析函数
func (v *Validator) updateFieldNameFunc() {
	if v.useJSON {
		v.engine.SetFieldNameFunc(v.jsonFieldNameFunc)
	} else {
		v.engine.SetFieldNameFunc(v.structFieldNameFunc)
	}
}

// jsonFieldNameFunc JSON字段名称解析函数（优先使用JSON标签）
func (v *Validator) jsonFieldNameFunc(field reflect.StructField) string {
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
func (v *Validator) structFieldNameFunc(field reflect.StructField) string {
	return field.Name
}

// registerDefaultValidators 注册默认验证器
func (v *Validator) registerDefaultValidators() error {
	// 注册自定义手机号验证
	if err := v.RegisterValidation("mobile", validateMobile); err != nil {
		return fmt.Errorf("failed to register mobile validator: %w", err)
	}

	// 注册自定义身份证验证
	if err := v.RegisterValidation("idcard", validateIDCard); err != nil {
		return fmt.Errorf("failed to register idcard validator: %w", err)
	}

	// 注册自定义银行卡验证
	if err := v.RegisterValidation("bankcard", validateBankCard); err != nil {
		return fmt.Errorf("failed to register bankcard validator: %w", err)
	}

	// 注册中文名称验证
	if err := v.RegisterValidation("chinese_name", validateChineseName); err != nil {
		return fmt.Errorf("failed to register chinese_name validator: %w", err)
	}

	// 注册强密码验证
	if err := v.RegisterValidation("strong_password", validateStrongPassword); err != nil {
		return fmt.Errorf("failed to register strong_password validator: %w", err)
	}

	// 覆盖内置的email验证
	if err := v.RegisterValidation("email", validateEmail); err != nil {
		return fmt.Errorf("failed to register email validator: %w", err)
	}

	// 覆盖内置的url验证
	if err := v.RegisterValidation("url", validateURL); err != nil {
		return fmt.Errorf("failed to register url validator: %w", err)
	}

	// 注册其他增强验证器
	if err := v.RegisterValidation("ipv4", validateIPv4); err != nil {
		return fmt.Errorf("failed to register ipv4 validator: %w", err)
	}

	if err := v.RegisterValidation("mac", validateMAC); err != nil {
		return fmt.Errorf("failed to register mac validator: %w", err)
	}

	if err := v.RegisterValidation("json", validateJSON); err != nil {
		return fmt.Errorf("failed to register json validator: %w", err)
	}

	if err := v.RegisterValidation("uuid", validateUUID); err != nil {
		return fmt.Errorf("failed to register uuid validator: %w", err)
	}

	return nil
}

// 全局便捷函数

// SetLocale 设置默认验证器的语言地区
func SetLocale(locale string) {
	Default().SetLocale(locale)
}

// SetUseJSON 设置默认验证器是否优先使用 JSON 字段名
func SetUseJSON(useJSON bool) {
	Default().SetUseJSON(useJSON)
}

// Struct 使用默认验证器验证结构体
func Struct(s interface{}) error {
	return Default().Struct(s)
}

// Var 使用默认验证器验证单个变量
func Var(field interface{}, tag string) error {
	return Default().Var(field, tag)
}

// RegisterValidation 在默认验证器上注册自定义验证规则
func RegisterValidation(tag string, fn ValidatorFunc) error {
	return Default().RegisterValidation(tag, fn)
}

// RegisterTranslation 在默认验证器上注册翻译
func RegisterTranslation(locale, tag, translation string) {
	Default().RegisterTranslation(locale, tag, translation)
}