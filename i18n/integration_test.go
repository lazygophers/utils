package i18n

import (
	"testing"
)

// TestLanguageCompleteness 测试所有支持语言的完整性
func TestLanguageCompleteness(t *testing.T) {
	expectedLanguages := map[string]string{
		English:            "English",
		ChineseSimplified:  "简体中文",
		ChineseTraditional: "繁體中文",
		Japanese:           "日本語",
		Korean:             "한국어",
		French:             "Français",
		Spanish:            "Español",
		Arabic:             "العربية",
		Russian:            "Русский",
		Italian:            "Italiano",
		Portuguese:         "Português",
		German:             "Deutsch",
	}

	for langCode, _ := range expectedLanguages {
		t.Run(langCode, func(t *testing.T) {
			locale, ok := GetLocale(langCode)
			if !ok {
				t.Errorf("Language %s should be available", langCode)
				return
			}

			if locale.Language != langCode {
				t.Errorf("Expected language %s, got %s", langCode, locale.Language)
			}

			if locale.Name == "" {
				t.Errorf("Language %s should have a name", langCode)
			}

			if locale.EnglishName == "" {
				t.Errorf("Language %s should have an English name", langCode)
			}

			// 检查基本格式配置
			if locale.Formats == nil {
				t.Errorf("Language %s should have formats configured", langCode)
				return
			}

			if locale.Formats.Units == nil {
				t.Errorf("Language %s should have units configured", langCode)
			}

			// 检查消息是否有基本内容
			if len(locale.Messages) == 0 && langCode != English {
				t.Errorf("Language %s should have some messages", langCode)
			}
		})
	}
}

// TestTranslationFallback 测试翻译回退机制
func TestTranslationFallback(t *testing.T) {
	testKey := "nonexistent_test_key"
	
	// 测试不存在的键名
	result := Translate(ChineseSimplified, testKey)
	if result != testKey {
		t.Errorf("Expected %s, got %s", testKey, result)
	}

	// 测试不存在的语言
	result = Translate("nonexistent_lang", testKey)
	if result != testKey {
		t.Errorf("Expected %s, got %s", testKey, result)
	}
}

// TestAllLanguageMessages 测试所有语言的通用消息
func TestAllLanguageMessages(t *testing.T) {
	commonKeys := []string{
		"error", "warning", "info", "success", "failed",
		"loading", "saving", "done", "cancel", "confirm",
		"yes", "no", "ok",
	}

	for _, langCode := range SupportedLanguages {
		if langCode == English {
			continue // 英语是默认语言，可能没有显式的消息定义
		}

		t.Run(langCode, func(t *testing.T) {
			locale, ok := GetLocale(langCode)
			if !ok {
				t.Skipf("Language %s not available", langCode)
				return
			}

			for _, key := range commonKeys {
				if message, exists := locale.Messages[key]; exists {
					if message == "" {
						t.Errorf("Language %s has empty message for key %s", langCode, key)
					}
					if message == key {
						t.Errorf("Language %s has untranslated message for key %s", langCode, key)
					}
				}
			}
		})
	}
}

// TestConcurrentAccess 测试并发访问
func TestConcurrentAccess(t *testing.T) {
	const goroutines = 100
	const iterations = 100

	done := make(chan bool, goroutines)

	for i := 0; i < goroutines; i++ {
		go func(id int) {
			defer func() { done <- true }()

			for j := 0; j < iterations; j++ {
				// 测试并发获取语言配置
				_, ok := GetLocale(ChineseSimplified)
				if !ok {
					t.Errorf("Goroutine %d: Failed to get Chinese locale", id)
					return
				}

				// 测试并发翻译
				Translate(ChineseSimplified, "error")

				// 测试并发获取可用语言列表
				GetAvailableLocales()
			}
		}(i)
	}

	// 等待所有goroutine完成
	for i := 0; i < goroutines; i++ {
		<-done
	}
}

// TestLanguageNormalization 测试语言代码标准化
func TestLanguageNormalization(t *testing.T) {
	testCases := []struct {
		input    string
		expected string
	}{
		{"zh", ChineseSimplified},
		{"zh-cn", ChineseSimplified},
		{"zh-CN", ChineseSimplified},
		{"chinese", ChineseSimplified},
		{"ZH-TW", ChineseTraditional},
		{"ja", Japanese},
		{"JA", Japanese},
		{"ko", Korean},
		{"fr", French},
		{"es", Spanish},
		{"ar", Arabic},
		{"ru", Russian},
		{"it", Italian},
		{"pt", Portuguese},
		{"de", German},
		{"en", English},
		{"EN", English},
		{"unknown", English},
		{"", English},
	}

	for _, tc := range testCases {
		t.Run(tc.input, func(t *testing.T) {
			result := NormalizeLanguage(tc.input)
			if result != tc.expected {
				t.Errorf("NormalizeLanguage(%q): expected %q, got %q", tc.input, tc.expected, result)
			}
		})
	}
}

// TestFormatConfiguration 测试格式化配置
func TestFormatConfiguration(t *testing.T) {
	for _, langCode := range SupportedLanguages {
		t.Run(langCode, func(t *testing.T) {
			locale, ok := GetLocale(langCode)
			if !ok {
				t.Skipf("Language %s not available", langCode)
				return
			}

			if locale.Formats == nil {
				t.Errorf("Language %s missing formats", langCode)
				return
			}

			formats := locale.Formats

			// 检查日期时间格式
			if formats.DateFormat == "" {
				t.Errorf("Language %s missing date format", langCode)
			}

			if formats.TimeFormat == "" {
				t.Errorf("Language %s missing time format", langCode)
			}

			if formats.DateTimeFormat == "" {
				t.Errorf("Language %s missing datetime format", langCode)
			}

			// 检查数字格式
			if formats.DecimalSeparator == "" {
				t.Errorf("Language %s missing decimal separator", langCode)
			}

			// 检查单位配置
			if formats.Units == nil {
				t.Errorf("Language %s missing units", langCode)
				return
			}

			units := formats.Units

			if len(units.ByteUnits) == 0 {
				t.Errorf("Language %s missing byte units", langCode)
			}

			if len(units.SpeedUnits) == 0 {
				t.Errorf("Language %s missing speed units", langCode)
			}

			if len(units.BitSpeedUnits) == 0 {
				t.Errorf("Language %s missing bit speed units", langCode)
			}

			if len(units.TimeUnits) == 0 {
				t.Errorf("Language %s missing time units", langCode)
			}
		})
	}
}

// BenchmarkTranslateIntegration 翻译性能基准测试（集成测试）
func BenchmarkTranslateIntegration(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Translate(ChineseSimplified, "error")
	}
}

// BenchmarkGetLocaleIntegration 获取语言配置性能基准测试（集成测试）
func BenchmarkGetLocaleIntegration(b *testing.B) {
	for i := 0; i < b.N; i++ {
		GetLocale(ChineseSimplified)
	}
}

// BenchmarkNormalizeLanguage 语言代码标准化性能基准测试
func BenchmarkNormalizeLanguage(b *testing.B) {
	for i := 0; i < b.N; i++ {
		NormalizeLanguage("zh-cn")
	}
}