// Package i18n 提供全局的国际化支持
// 统一管理整个项目的多语言功能
package i18n

import (
	"fmt"
	"strings"
	"sync"
)

// 支持的语言列表
const (
	English            = "en"      // 英语（默认）
	ChineseSimplified  = "zh-CN"   // 简体中文
	ChineseTraditional = "zh-TW"   // 繁体中文
	Japanese           = "ja"      // 日语
	Korean             = "ko"      // 韩语
	French             = "fr"      // 法语
	Spanish            = "es"      // 西班牙语
	Arabic             = "ar"      // 阿拉伯语
	Russian            = "ru"      // 俄语
	Italian            = "it"      // 意大利语
	Portuguese         = "pt"      // 葡萄牙语
	German             = "de"      // 德语
)

// SupportedLanguages 所有支持的语言
var SupportedLanguages = []string{
	English, ChineseSimplified, ChineseTraditional,
	Japanese, Korean, French, Spanish, Arabic,
	Russian, Italian, Portuguese, German,
}

// Locale 国际化配置
type Locale struct {
	Language   string            // 语言代码 (ISO 639-1)
	Region     string            // 地区代码 (ISO 3166-1 alpha-2)
	Name       string            // 语言本地化名称
	EnglishName string           // 英语名称
	Messages   map[string]string // 消息映射
	Formats    *Formats          // 格式化配置
}

// Formats 格式化配置
type Formats struct {
	DateFormat      string   // 日期格式
	TimeFormat      string   // 时间格式
	DateTimeFormat  string   // 日期时间格式
	NumberFormat    string   // 数字格式
	CurrencyFormat  string   // 货币格式
	DecimalSeparator string  // 小数分隔符
	ThousandSeparator string // 千位分隔符
	Units           *Units   // 单位配置
}

// Units 单位配置
type Units struct {
	// 字节单位
	ByteUnits     []string // ["B", "KB", "MB", "GB", "TB", "PB"]
	SpeedUnits    []string // ["B/s", "KB/s", "MB/s", "GB/s", "TB/s", "PB/s"]
	BitSpeedUnits []string // ["bps", "Kbps", "Mbps", "Gbps", "Tbps", "Pbps"]
	
	// 时间单位
	TimeUnits     map[string]string // 时间单位映射
	
	// 距离单位
	DistanceUnits []string // ["mm", "cm", "m", "km"]
	
	// 重量单位
	WeightUnits   []string // ["g", "kg", "t"]
}

// Manager 国际化管理器
type Manager struct {
	mu             sync.RWMutex
	defaultLocale  string
	locales        map[string]*Locale
	fallbackChain  []string // 回退链
}

// 全局管理器
var globalManager = &Manager{
	defaultLocale: English,
	locales:       make(map[string]*Locale),
	fallbackChain: []string{English}, // 默认回退到英语
}

// SetDefaultLocale 设置默认语言
func SetDefaultLocale(locale string) {
	globalManager.SetDefaultLocale(locale)
}

// GetDefaultLocale 获取默认语言
func GetDefaultLocale() string {
	return globalManager.GetDefaultLocale()
}

// RegisterLocale 注册语言配置
func RegisterLocale(language string, locale *Locale) {
	globalManager.RegisterLocale(language, locale)
}

// GetLocale 获取语言配置
func GetLocale(language string) (*Locale, bool) {
	return globalManager.GetLocale(language)
}

// GetAvailableLocales 获取所有可用语言
func GetAvailableLocales() []string {
	return globalManager.GetAvailableLocales()
}

// Translate 翻译消息
func Translate(language, key string, args ...interface{}) string {
	return globalManager.Translate(language, key, args...)
}

// TranslateDefault 使用默认语言翻译消息
func TranslateDefault(key string, args ...interface{}) string {
	return globalManager.TranslateDefault(key, args...)
}

// SetDefaultLocale 设置默认语言
func (m *Manager) SetDefaultLocale(locale string) {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.defaultLocale = locale
}

// GetDefaultLocale 获取默认语言
func (m *Manager) GetDefaultLocale() string {
	m.mu.RLock()
	defer m.mu.RUnlock()
	return m.defaultLocale
}

// RegisterLocale 注册语言配置
func (m *Manager) RegisterLocale(language string, locale *Locale) {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.locales[language] = locale
	
	// 如果注册的是简体中文，同时注册为"zh"
	if language == ChineseSimplified {
		m.locales["zh"] = locale
	}
}

// GetLocale 获取语言配置
func (m *Manager) GetLocale(language string) (*Locale, bool) {
	m.mu.RLock()
	defer m.mu.RUnlock()
	
	// 直接匹配
	if locale, ok := m.locales[language]; ok {
		return locale, true
	}
	
	// 尝试语言代码匹配（忽略地区）
	langCode := strings.Split(language, "-")[0]
	if locale, ok := m.locales[langCode]; ok {
		return locale, true
	}
	
	// 尝试回退链
	for _, fallback := range m.fallbackChain {
		if locale, ok := m.locales[fallback]; ok {
			return locale, true
		}
	}
	
	return nil, false
}

// GetAvailableLocales 获取所有可用语言
func (m *Manager) GetAvailableLocales() []string {
	m.mu.RLock()
	defer m.mu.RUnlock()
	
	var locales []string
	for lang := range m.locales {
		locales = append(locales, lang)
	}
	return locales
}

// Translate 翻译消息
func (m *Manager) Translate(language, key string, args ...interface{}) string {
	locale, ok := m.GetLocale(language)
	if !ok {
		// 如果找不到语言配置，返回键名
		if len(args) > 0 {
			return fmt.Sprintf(key, args...)
		}
		return key
	}
	
	if message, exists := locale.Messages[key]; exists {
		if len(args) > 0 {
			return fmt.Sprintf(message, args...)
		}
		return message
	}
	
	// 如果没有找到翻译，返回键名
	if len(args) > 0 {
		return fmt.Sprintf(key, args...)
	}
	return key
}

// TranslateDefault 使用默认语言翻译消息
func (m *Manager) TranslateDefault(key string, args ...interface{}) string {
	return m.Translate(m.GetDefaultLocale(), key, args...)
}

// IsSupported 检查语言是否被支持
func IsSupported(language string) bool {
	for _, supported := range SupportedLanguages {
		if supported == language {
			return true
		}
	}
	// 检查是否为语言代码简写
	langCode := strings.Split(language, "-")[0]
	for _, supported := range SupportedLanguages {
		if strings.Split(supported, "-")[0] == langCode {
			return true
		}
	}
	return false
}

// NormalizeLanguage 标准化语言代码
func NormalizeLanguage(language string) string {
	switch strings.ToLower(language) {
	case "zh", "zh-cn", "chinese", "simplified chinese":
		return ChineseSimplified
	case "zh-tw", "traditional chinese":
		return ChineseTraditional
	case "ja", "japanese":
		return Japanese
	case "ko", "korean":
		return Korean
	case "fr", "french":
		return French
	case "es", "spanish":
		return Spanish
	case "ar", "arabic":
		return Arabic
	case "ru", "russian":
		return Russian
	case "it", "italian":
		return Italian
	case "pt", "portuguese":
		return Portuguese
	case "de", "german":
		return German
	case "en", "english":
		return English
	default:
		return English
	}
}