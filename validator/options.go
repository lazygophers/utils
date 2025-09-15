package validator

// Option 配置选项
type Option func(*Validator)

// WithLocale 设置语言地区
func WithLocale(locale string) Option {
	return func(v *Validator) {
		v.locale = locale
	}
}

// WithUseJSON 设置是否优先使用 JSON 字段名
func WithUseJSON(useJSON bool) Option {
	return func(v *Validator) {
		v.useJSON = useJSON
	}
}

// WithTranslations 设置翻译映射
func WithTranslations(translations map[string]string) Option {
	return func(v *Validator) {
		for key, value := range translations {
			v.messages[key] = value
		}
	}
}

// WithCustomValidator 添加自定义验证器
func WithCustomValidator(tag string, fn func(interface{}) bool) Option {
	return func(v *Validator) {
		_ = v.RegisterValidation(tag, func(fl FieldLevel) bool {
			return fn(fl.Field().Interface())
		})
	}
}

// Config 验证器配置（用于批量设置）
type Config struct {
	Locale       string            // 语言地区
	UseJSON      bool              // 是否优先使用 JSON 字段名
	Translations map[string]string // 自定义翻译
}

// WithConfig 使用配置对象设置验证器
func WithConfig(config Config) Option {
	return func(v *Validator) {
		if config.Locale != "" {
			v.locale = config.Locale
		}
		v.useJSON = config.UseJSON

		if config.Translations != nil {
			for key, value := range config.Translations {
				v.messages[key] = value
			}
		}
	}
}
