package fake

import (
	"context"
	"fmt"
	"strings"
	"sync"
	"time"
)

// Language 语言类型
type Language string

const (
	LanguageEnglish            Language = "en"
	LanguageChineseSimplified  Language = "zh-CN"
	LanguageChineseTraditional Language = "zh-TW"
	LanguageFrench             Language = "fr"
	LanguageRussian            Language = "ru"
	LanguagePortuguese         Language = "pt"
	LanguageSpanish            Language = "es"
)

// Gender 性别类型
type Gender string

const (
	GenderMale   Gender = "male"
	GenderFemale Gender = "female"
)

// Country 国家类型
type Country string

const (
	CountryChina     Country = "CN"
	CountryUS        Country = "US"
	CountryUK        Country = "UK"
	CountryFrance    Country = "FR"
	CountryGermany   Country = "DE"
	CountryJapan     Country = "JP"
	CountryKorea     Country = "KR"
	CountryRussia    Country = "RU"
	CountryBrazil    Country = "BR"
	CountrySpain     Country = "ES"
	CountryPortugal  Country = "PT"
	CountryItaly     Country = "IT"
	CountryCanada    Country = "CA"
	CountryAustralia Country = "AU"
	CountryIndia     Country = "IN"
)

// Faker 假数据生成器结构体
type Faker struct {
	language Language
	country  Country
	gender   Gender
	seed     int64

	// 数据缓存
	dataMu sync.RWMutex
	data   map[string]interface{}

	// 高性能优化数据
	fastData *FastData
}

// FakerOption 配置选项
type FakerOption func(*Faker)

// WithLanguage 设置语言
func WithLanguage(lang Language) FakerOption {
	return func(f *Faker) {
		f.language = lang
	}
}

// WithCountry 设置国家
func WithCountry(country Country) FakerOption {
	return func(f *Faker) {
		f.country = country
	}
}

// WithGender 设置性别
func WithGender(gender Gender) FakerOption {
	return func(f *Faker) {
		f.gender = gender
	}
}

// WithSeed 设置随机种子
func WithSeed(seed int64) FakerOption {
	return func(f *Faker) {
		f.seed = seed
	}
}

// New 创建新的假数据生成器
func New(opts ...FakerOption) *Faker {
	f := &Faker{
		language: LanguageEnglish,
		country:  CountryUS,
		gender:   GenderMale,
		seed:     time.Now().UnixNano(),
		data:     make(map[string]interface{}),
	}

	for _, opt := range opts {
		opt(f)
	}

	return f
}

// 全局默认实例
var (
	defaultFaker *Faker
	defaultOnce  sync.Once
)

// getDefaultFaker 获取默认的假数据生成器实例
func getDefaultFaker() *Faker {
	defaultOnce.Do(func() {
		defaultFaker = New()
	})
	return defaultFaker
}

// SetDefaultLanguage 设置默认语言
func SetDefaultLanguage(lang Language) {
	getDefaultFaker().language = lang
}

// SetDefaultCountry 设置默认国家
func SetDefaultCountry(country Country) {
	getDefaultFaker().country = country
}

// SetDefaultGender 设置默认性别
func SetDefaultGender(gender Gender) {
	getDefaultFaker().gender = gender
}

// SetDefaultSeed 设置默认随机种子
func SetDefaultSeed(seed int64) {
	getDefaultFaker().seed = seed
}

// GetDefaultLanguage 获取默认语言
func GetDefaultLanguage() Language {
	return getDefaultFaker().language
}

// GetDefaultCountry 获取默认国家
func GetDefaultCountry() Country {
	return getDefaultFaker().country
}

// GetDefaultGender 获取默认性别
func GetDefaultGender() Gender {
	return getDefaultFaker().gender
}

// GetDefaultSeed 获取默认随机种子
func GetDefaultSeed() int64 {
	return getDefaultFaker().seed
}

// ResetDefaultFaker 重置默认Faker实例
func ResetDefaultFaker() {
	defaultOnce = sync.Once{}
	defaultFaker = nil
}

// SetDefaultFaker 设置自定义的默认Faker实例
func SetDefaultFaker(faker *Faker) {
	defaultFaker = faker
}

// SetDefaults 批量设置默认选项
func SetDefaults(opts ...FakerOption) {
	f := getDefaultFaker()
	for _, opt := range opts {
		opt(f)
	}
}

// ClearDefaultCache 清空默认实例的数据缓存
func ClearDefaultCache() {
	getDefaultFaker().ClearCache()
}

// ClearCache 清空数据缓存
func (f *Faker) ClearCache() {
	f.dataMu.Lock()
	defer f.dataMu.Unlock()

	for k := range f.data {
		delete(f.data, k)
	}
}

// 数据缓存辅助方法
func (f *Faker) getCachedData(key string) (interface{}, bool) {
	f.dataMu.RLock()
	defer f.dataMu.RUnlock()

	data, exists := f.data[key]
	return data, exists
}

func (f *Faker) setCachedData(key string, data interface{}) {
	f.dataMu.Lock()
	defer f.dataMu.Unlock()

	f.data[key] = data
}

// validateLanguage 验证语言是否支持
func (f *Faker) validateLanguage(lang Language) bool {
	supportedLanguages := []Language{
		LanguageEnglish,
		LanguageChineseSimplified,
		LanguageChineseTraditional,
		LanguageFrench,
		LanguageRussian,
		LanguagePortuguese,
		LanguageSpanish,
	}

	for _, supported := range supportedLanguages {
		if lang == supported {
			return true
		}
	}
	return false
}

// validateCountry 验证国家是否支持
func (f *Faker) validateCountry(country Country) bool {
	supportedCountries := []Country{
		CountryChina, CountryUS, CountryUK, CountryFrance, CountryGermany,
		CountryJapan, CountryKorea, CountryRussia, CountryBrazil, CountrySpain,
		CountryPortugal, CountryItaly, CountryCanada, CountryAustralia, CountryIndia,
	}

	for _, supported := range supportedCountries {
		if country == supported {
			return true
		}
	}
	return false
}

// GetSupportedLanguages 获取支持的语言列表
func GetSupportedLanguages() []Language {
	return []Language{
		LanguageEnglish,
		LanguageChineseSimplified,
		LanguageChineseTraditional,
		LanguageFrench,
		LanguageRussian,
		LanguagePortuguese,
		LanguageSpanish,
	}
}

// GetSupportedCountries 获取支持的国家列表
func GetSupportedCountries() []Country {
	return []Country{
		CountryChina, CountryUS, CountryUK, CountryFrance, CountryGermany,
		CountryJapan, CountryKorea, CountryRussia, CountryBrazil, CountrySpain,
		CountryPortugal, CountryItaly, CountryCanada, CountryAustralia, CountryIndia,
	}
}

// 上下文支持
type contextKey string

const (
	contextKeyLanguage contextKey = "fake_language"
	contextKeyCountry  contextKey = "fake_country"
	contextKeyGender   contextKey = "fake_gender"
)

// WithContext 创建带上下文的假数据生成器
func WithContext(ctx context.Context, opts ...FakerOption) *Faker {
	f := New(opts...)

	// 从上下文中读取配置
	if lang, ok := ctx.Value(contextKeyLanguage).(Language); ok {
		f.language = lang
	}
	if country, ok := ctx.Value(contextKeyCountry).(Country); ok {
		f.country = country
	}
	if gender, ok := ctx.Value(contextKeyGender).(Gender); ok {
		f.gender = gender
	}

	return f
}

// ContextWithLanguage 为上下文添加语言信息
func ContextWithLanguage(ctx context.Context, lang Language) context.Context {
	return context.WithValue(ctx, contextKeyLanguage, lang)
}

// ContextWithCountry 为上下文添加国家信息
func ContextWithCountry(ctx context.Context, country Country) context.Context {
	return context.WithValue(ctx, contextKeyCountry, country)
}

// ContextWithGender 为上下文添加性别信息
func ContextWithGender(ctx context.Context, gender Gender) context.Context {
	return context.WithValue(ctx, contextKeyGender, gender)
}

// 字符串格式化辅助函数
func formatWithParams(format string, params map[string]string) string {
	result := format
	for key, value := range params {
		placeholder := fmt.Sprintf("{{%s}}", key)
		result = strings.ReplaceAll(result, placeholder, value)
	}
	return result
}

// 数字格式化辅助函数
func formatNumber(format string, number int) string {
	formatStr := strings.ReplaceAll(format, "#", "%d")
	return fmt.Sprintf(formatStr, number)
}

// 批量生成辅助函数
func (f *Faker) batchGenerate(count int, generator func() string) []string {
	results := make([]string, count)
	for i := 0; i < count; i++ {
		results[i] = generator()
	}
	return results
}

// Clone 克隆假数据生成器（用于并发安全）
func (f *Faker) Clone() *Faker {
	return &Faker{
		language: f.language,
		country:  f.country,
		gender:   f.gender,
		seed:     time.Now().UnixNano(), // 使用新的种子
		data:     make(map[string]interface{}),
	}
}

// UserAgent 生成用户代理字符串
func (f *Faker) UserAgent() string {
	// 使用新的智能生成器替代静态列表
	return f.GenerateRandomUserAgent()
}

// RandomUserAgent 返回随机的用户代理字符串
func RandomUserAgent() string {
	return getDefaultFaker().UserAgent()
}
