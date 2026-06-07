package validator

import (
	"sync"

	"golang.org/x/text/language"
)

// LocaleConfig 语言地区配置
type LocaleConfig struct {
	Language language.Tag      // 语言标签（golang.org/x/text/language；用 language.Make("en") 等构造）
	Region   string            // 地区代码，如 "US", "CN"
	Messages map[string]string // 验证错误消息模板
}

// 全局地区管理器，使用 sync.Map 实现无锁读
var localeConfigs sync.Map // map[string]*LocaleConfig

// RegisterLocaleConfig 注册语言地区配置
func RegisterLocaleConfig(name string, config *LocaleConfig) {
	localeConfigs.Store(name, config)
}

// GetLocaleConfig 获取语言地区配置
// 查找链：完整匹配 > 语言前缀 > 英文兜底
func GetLocaleConfig(name string) (*LocaleConfig, bool) {
	// 尝试完整匹配
	if v, ok := localeConfigs.Load(name); ok {
		return v.(*LocaleConfig), true
	}

	// 尝试语言匹配（忽略地区）
	if len(name) > 2 && name[2] == '-' {
		if v, ok := localeConfigs.Load(name[:2]); ok {
			return v.(*LocaleConfig), true
		}
	}

	// 默认英文
	if v, ok := localeConfigs.Load("en"); ok {
		return v.(*LocaleConfig), true
	}

	return nil, false
}

// GetAvailableLocales 获取所有可用的语言地区
func GetAvailableLocales() []string {
	var locales []string
	localeConfigs.Range(func(key, _ interface{}) bool {
		locales = append(locales, key.(string))
		return true
	})
	return locales
}

// 注册默认的英文语言配置
func init() {
	RegisterLocaleConfig("en", &LocaleConfig{
		Language: language.Make("en"),
		Region:   "US",
		Messages: map[string]string{
			// 内置验证规则
			"required":             "{field} is required",
			"email":                "{field} must be a valid email address",
			"url":                  "{field} must be a valid URL",
			"min":                  "{field} must be at least {param}",
			"max":                  "{field} must be at most {param}",
			"len":                  "{field} must be exactly {param} characters long",
			"oneof":                "{field} must be one of [{param}]",
			"unique":               "{field} must contain unique values",
			"alpha":                "{field} must contain only letters",
			"alphanum":             "{field} must contain only letters and numbers",
			"numeric":              "{field} must be a valid number",
			"number":               "{field} must be a valid number",
			"hexadecimal":          "{field} must be a valid hexadecimal",
			"hexcolor":             "{field} must be a valid hex color",
			"rgb":                  "{field} must be a valid RGB color",
			"rgba":                 "{field} must be a valid RGBA color",
			"hsl":                  "{field} must be a valid HSL color",
			"hsla":                 "{field} must be a valid HSLA color",
			"e164":                 "{field} must be a valid E164 phone number",
			"json":                 "{field} must be valid JSON",
			"jwt":                  "{field} must be a valid JWT",
			"uuid":                 "{field} must be a valid UUID",
			"uuid3":                "{field} must be a valid version 3 UUID",
			"uuid4":                "{field} must be a valid version 4 UUID",
			"uuid5":                "{field} must be a valid version 5 UUID",
			"ascii":                "{field} must contain only ASCII characters",
			"printascii":           "{field} must contain only printable ASCII characters",
			"multibyte":            "{field} must contain multibyte characters",
			"datauri":              "{field} must be a valid data URI",
			"latitude":             "{field} must be a valid latitude",
			"longitude":            "{field} must be a valid longitude",
			"ssn":                  "{field} must be a valid SSN",
			"ipv4":                 "{field} must be a valid IPv4 address",
			"ipv6":                 "{field} must be a valid IPv6 address",
			"ip":                   "{field} must be a valid IP address",
			"cidr":                 "{field} must be a valid CIDR notation",
			"cidrv4":               "{field} must be a valid CIDR notation for IPv4",
			"cidrv6":               "{field} must be a valid CIDR notation for IPv6",
			"tcp_addr":             "{field} must be a valid TCP address",
			"tcp4_addr":            "{field} must be a valid TCP4 address",
			"tcp6_addr":            "{field} must be a valid TCP6 address",
			"udp_addr":             "{field} must be a valid UDP address",
			"udp4_addr":            "{field} must be a valid UDP4 address",
			"udp6_addr":            "{field} must be a valid UDP6 address",
			"ip_addr":              "{field} must be a valid IP address",
			"ip4_addr":             "{field} must be a valid IPv4 address",
			"ip6_addr":             "{field} must be a valid IPv6 address",
			"unix_addr":            "{field} must be a valid Unix address",
			"mac":                  "{field} must be a valid MAC address",
			"hostname":             "{field} must be a valid hostname",
			"hostname_rfc1123":     "{field} must be a valid RFC1123 hostname",
			"hostname_port":        "{field} must be a valid hostname:port",
			"fqdn":                 "{field} must be a valid FQDN",
			"uri":                  "{field} must be a valid URI",
			"url_encoded":          "{field} must be URL encoded",
			"dir":                  "{field} must be a valid directory path",
			"file":                 "{field} must be a valid file path",
			"base64":               "{field} must be valid base64",
			"base64url":            "{field} must be valid base64url",
			"contains":             "{field} must contain the substring '{param}'",
			"containsany":          "{field} must contain at least one of the characters '{param}'",
			"containsrune":         "{field} must contain the rune '{param}'",
			"excludes":             "{field} must not contain the substring '{param}'",
			"excludesall":          "{field} must not contain any of the characters '{param}'",
			"excludesrune":         "{field} must not contain the rune '{param}'",
			"startswith":           "{field} must start with '{param}'",
			"endswith":             "{field} must end with '{param}'",
			"gt":                   "{field} must be greater than {param}",
			"gte":                  "{field} must be greater than or equal to {param}",
			"lt":                   "{field} must be less than {param}",
			"lte":                  "{field} must be less than or equal to {param}",
			"ne":                   "{field} must not be equal to {param}",
			"eq":                   "{field} must be equal to {param}",
			"eqfield":              "{field} must be equal to {param}",
			"nefield":              "{field} must not be equal to {param}",
			"gtfield":              "{field} must be greater than {param}",
			"gtefield":             "{field} must be greater than or equal to {param}",
			"ltfield":              "{field} must be less than {param}",
			"ltefield":             "{field} must be less than or equal to {param}",
			"required_if":          "{field} is required when {param}",
			"required_unless":      "{field} is required unless {param}",
			"required_with":        "{field} is required when {param} is present",
			"required_with_all":    "{field} is required when all of {param} are present",
			"required_without":     "{field} is required when {param} is not present",
			"required_without_all": "{field} is required when none of {param} are present",
			"excluded_if":          "{field} is excluded when {param}",
			"excluded_unless":      "{field} is excluded unless {param}",
			"excluded_with":        "{field} is excluded when {param} is present",
			"excluded_with_all":    "{field} is excluded when all of {param} are present",
			"excluded_without":     "{field} is excluded when {param} is not present",
			"excluded_without_all": "{field} is excluded when none of {param} are present",

			// 自定义验证规则
			"strong_password": "{field} must be a strong password (at least 8 characters with uppercase, lowercase, numbers, and special characters)",
		},
	})
}
