package i18n

import (
	"testing"
)

func TestSetAndGetDefaultLocale(t *testing.T) {
	original := GetDefaultLocale()
	defer SetDefaultLocale(original) // 恢复原始值

	// 测试设置和获取默认语言
	SetDefaultLocale(ChineseSimplified)
	if got := GetDefaultLocale(); got != ChineseSimplified {
		t.Errorf("Expected %s, got %s", ChineseSimplified, got)
	}
}

func TestRegisterAndGetLocale(t *testing.T) {
	// 测试注册自定义语言配置
	testLocale := &Locale{
		Language:     "test",
		Region:       "TEST",
		Name:         "Test Language",
		EnglishName:  "Test Language",
		Messages:     map[string]string{"hello": "Hello World"},
		Formats:      &Formats{},
	}

	RegisterLocale("test", testLocale)

	// 测试获取已注册的语言配置
	locale, ok := GetLocale("test")
	if !ok {
		t.Fatal("Failed to get registered locale")
	}

	if locale.Language != "test" {
		t.Errorf("Expected language 'test', got '%s'", locale.Language)
	}

	if locale.Messages["hello"] != "Hello World" {
		t.Errorf("Expected 'Hello World', got '%s'", locale.Messages["hello"])
	}
}

func TestGetLocaleWithFallback(t *testing.T) {
	// 测试不存在的语言会回退到英语
	locale, ok := GetLocale("nonexistent")
	if !ok {
		t.Fatal("Should fallback to English for nonexistent locale")
	}

	if locale.Language != English {
		t.Errorf("Expected fallback to English, got %s", locale.Language)
	}
}

func TestGetAvailableLocales(t *testing.T) {
	locales := GetAvailableLocales()
	if len(locales) == 0 {
		t.Error("Should have at least one locale available")
	}

	// 英语应该总是可用的
	found := false
	for _, locale := range locales {
		if locale == English {
			found = true
			break
		}
	}
	if !found {
		t.Error("English locale should always be available")
	}
}

func TestTranslate(t *testing.T) {
	// 测试翻译功能
	result := Translate(English, "nonexistent_key")
	if result != "nonexistent_key" {
		t.Errorf("Expected 'nonexistent_key', got '%s'", result)
	}

	// 测试带参数的翻译
	result = Translate(English, "test_%s", "value")
	if result != "test_value" {
		t.Errorf("Expected 'test_value', got '%s'", result)
	}
}

func TestTranslateDefault(t *testing.T) {
	original := GetDefaultLocale()
	defer SetDefaultLocale(original)

	SetDefaultLocale(English)
	result := TranslateDefault("test_key")
	if result != "test_key" {
		t.Errorf("Expected 'test_key', got '%s'", result)
	}
}

func TestIsSupported(t *testing.T) {
	// 测试支持的语言
	if !IsSupported(English) {
		t.Error("English should be supported")
	}

	if !IsSupported(ChineseSimplified) {
		t.Error("Chinese Simplified should be supported")
	}

	if IsSupported("unsupported") {
		t.Error("'unsupported' should not be supported")
	}

	// 测试语言代码简写
	if !IsSupported("zh") {
		t.Error("'zh' should be supported")
	}
}

func TestNormalizeLanguage(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"zh", ChineseSimplified},
		{"zh-cn", ChineseSimplified},
		{"chinese", ChineseSimplified},
		{"zh-tw", ChineseTraditional},
		{"ja", Japanese},
		{"ko", Korean},
		{"fr", French},
		{"es", Spanish},
		{"ar", Arabic},
		{"ru", Russian},
		{"it", Italian},
		{"pt", Portuguese},
		{"de", German},
		{"en", English},
		{"unknown", English}, // 默认回退到英语
	}

	for _, test := range tests {
		result := NormalizeLanguage(test.input)
		if result != test.expected {
			t.Errorf("NormalizeLanguage(%s): expected %s, got %s", test.input, test.expected, result)
		}
	}
}

// 测试语言配置的完整性
func TestLanguageConfigCompleteness(t *testing.T) {
	testLanguages := []string{
		English, ChineseSimplified, ChineseTraditional,
		Japanese, Korean, French, Spanish, Arabic,
		Russian, Italian, Portuguese, German,
	}

	for _, lang := range testLanguages {
		locale, ok := GetLocale(lang)
		if !ok {
			t.Errorf("Language %s should be available", lang)
			continue
		}

		if locale.Language == "" {
			t.Errorf("Language %s should have language field set", lang)
		}

		if locale.Name == "" {
			t.Errorf("Language %s should have name field set", lang)
		}

		if locale.EnglishName == "" {
			t.Errorf("Language %s should have english name field set", lang)
		}

		if locale.Formats == nil {
			t.Errorf("Language %s should have formats configured", lang)
		}

		if locale.Formats.Units == nil {
			t.Errorf("Language %s should have units configured", lang)
		}
	}
}

// 基准测试
func BenchmarkTranslate(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Translate(English, "nonexistent_key")
	}
}

func BenchmarkGetLocale(b *testing.B) {
	for i := 0; i < b.N; i++ {
		GetLocale(English)
	}
}