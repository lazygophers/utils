package validator

import (
	"fmt"
	"reflect"
	"strings"
	"sync"

	"github.com/go-playground/validator/v10"
)

var (
	defaultValidator *Validator
	once             sync.Once
)

// Validator 自定义验证器，兼容 github.com/go-playground/validator/v10
type Validator struct {
	engine   *validator.Validate
	locale   string
	useJSON  bool
	mu       sync.RWMutex
	messages map[string]string
}

// New 创建新的验证器实例
func New(opts ...Option) (*Validator, error) {
	v := &Validator{
		engine:   validator.New(),
		locale:   "en",
		useJSON:  true,
		messages: make(map[string]string),
	}

	// 应用选项
	for _, opt := range opts {
		opt(v)
	}

	// 设置字段名称函数
	v.engine.RegisterTagNameFunc(v.getFieldName)

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
				engine:   validator.New(),
				locale:   "en",
				useJSON:  true,
				messages: make(map[string]string),
			}
			defaultValidator.engine.RegisterTagNameFunc(defaultValidator.getFieldName)
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
}

// Struct 验证结构体
func (v *Validator) Struct(s interface{}) error {
	err := v.engine.Struct(s)
	if err != nil {
		if validationErrors, ok := err.(validator.ValidationErrors); ok {
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
		if validationErrors, ok := err.(validator.ValidationErrors); ok {
			return v.translateValidationErrors(validationErrors)
		}
		return err
	}
	return nil
}

// RegisterValidation 注册自定义验证规则
func (v *Validator) RegisterValidation(tag string, fn validator.Func, callValidationEvenIfNull ...bool) error {
	return v.engine.RegisterValidation(tag, fn, callValidationEvenIfNull...)
}

// RegisterTranslation 注册翻译
func (v *Validator) RegisterTranslation(locale, tag, translation string) {
	v.mu.Lock()
	defer v.mu.Unlock()
	key := fmt.Sprintf("%s.%s", locale, tag)
	v.messages[key] = translation
}

// getFieldName 获取字段名称，优先使用 JSON tag
func (v *Validator) getFieldName(fld reflect.StructField) string {
	v.mu.RLock()
	useJSON := v.useJSON
	v.mu.RUnlock()

	if useJSON {
		// 优先使用 json tag
		if jsonTag := fld.Tag.Get("json"); jsonTag != "" && jsonTag != "-" {
			// 处理 json:"name,omitempty" 格式
			if parts := strings.Split(jsonTag, ","); len(parts) > 0 && parts[0] != "" {
				return parts[0]
			}
		}
	}

	// 回退到字段名
	return fld.Name
}

// translateValidationErrors 翻译验证错误
func (v *Validator) translateValidationErrors(validationErrors validator.ValidationErrors) error {
	var errs ValidationErrors

	for _, err := range validationErrors {
		fieldError := &FieldError{
			Field:       err.Field(),
			Tag:         err.Tag(),
			Value:       err.Value(),
			Param:       err.Param(),
			ActualTag:   err.ActualTag(),
			Namespace:   err.Namespace(),
			StructField: err.StructNamespace(),
			Message:     v.translateError(err),
		}
		errs = append(errs, fieldError)
	}

	return errs
}

// translateError 翻译单个错误
func (v *Validator) translateError(err validator.FieldError) string {
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
			return fmt.Sprintf("%s failed validation for tag '%s'", err.Field(), err.Tag())
		}
	}

	// 构建翻译键
	key := fmt.Sprintf("%s.%s", locale, err.Tag())

	v.mu.RLock()
	if msg, exists := v.messages[key]; exists {
		v.mu.RUnlock()
		return v.formatMessage(msg, err)
	}
	v.mu.RUnlock()

	// 使用默认消息模板
	if template, exists := localeConfig.Messages[err.Tag()]; exists {
		return v.formatMessage(template, err)
	}

	// 最后回退到英文默认消息
	if locale != "en" {
		if englishConfig, ok := GetLocaleConfig("en"); ok {
			if template, exists := englishConfig.Messages[err.Tag()]; exists {
				return v.formatMessage(template, err)
			}
		}
	}

	// 如果没有找到翻译，返回默认格式
	return fmt.Sprintf("%s failed validation for tag '%s'", err.Field(), err.Tag())
}

// formatMessage 格式化错误消息
func (v *Validator) formatMessage(template string, err validator.FieldError) string {
	message := template

	// 替换占位符
	message = strings.ReplaceAll(message, "{field}", err.Field())
	message = strings.ReplaceAll(message, "{tag}", err.Tag())
	message = strings.ReplaceAll(message, "{param}", err.Param())

	// 处理值的显示
	if err.Value() != nil {
		message = strings.ReplaceAll(message, "{value}", fmt.Sprintf("%v", err.Value()))
	} else {
		message = strings.ReplaceAll(message, "{value}", "")
	}

	return message
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
func RegisterValidation(tag string, fn validator.Func, callValidationEvenIfNull ...bool) error {
	return Default().RegisterValidation(tag, fn, callValidationEvenIfNull...)
}

// RegisterTranslation 在默认验证器上注册翻译
func RegisterTranslation(locale, tag, translation string) {
	Default().RegisterTranslation(locale, tag, translation)
}
