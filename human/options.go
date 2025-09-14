package human

// Option 函数选项类型
type Option func(*Config)

// Config 配置结构
type Config struct {
	Precision  int
	Locale     string
	Compact    bool
	TimeFormat string
}

// DefaultConfig 返回默认配置
func DefaultConfig() Config {
	return Config{
		Precision:  1,
		Locale:     "en",
		Compact:    false,
		TimeFormat: "",
	}
}

// WithPrecision 设置精度
func WithPrecision(precision int) Option {
	return func(c *Config) {
		c.Precision = precision
	}
}

// WithLocale 设置语言地区
func WithLocale(locale string) Option {
	return func(c *Config) {
		c.Locale = locale
	}
}

// WithCompact 启用紧凑模式
func WithCompact() Option {
	return func(c *Config) {
		c.Compact = true
	}
}

// WithClockFormat 启用时钟格式 (1:20)
func WithClockFormat() Option {
	return func(c *Config) {
		c.TimeFormat = "clock"
	}
}

// applyOptions 应用选项到配置
func applyOptions(opts ...Option) Config {
	config := DefaultConfig()
	for _, opt := range opts {
		opt(&config)
	}
	return config
}