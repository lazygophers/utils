//go:build lang_all

package validator

import (
	"testing"
)

func TestAllLocalesAvailable(t *testing.T) {
	expectedLocales := []string{
		"en", "zh", "zh-CN", "zh-TW", "ja", "ko", "fr", "es", "ar", "ru", "it", "pt", "de",
	}

	availableLocales := GetAvailableLocales()
	localeMap := make(map[string]bool)
	for _, locale := range availableLocales {
		localeMap[locale] = true
	}

	for _, expected := range expectedLocales {
		if !localeMap[expected] {
			t.Errorf("Expected locale %s not found in available locales: %v", expected, availableLocales)
		}
	}
}

func TestLocaleMessages(t *testing.T) {
	testCases := []struct {
		locale  string
		msgKey  string
		wantKey string
	}{
		{"en", "required", "{field} is required"},
		{"zh", "required", "{field}不能为空"},
		{"zh-CN", "required", "{field}不能为空"},
		{"zh-TW", "required", "{field}不能為空"},
		{"ja", "required", "{field}は必須です"},
		{"ko", "required", "{field}은(는) 필수입니다"},
		{"fr", "required", "{field} est requis"},
		{"es", "required", "{field} es obligatorio"},
		{"ar", "required", "{field} مطلوب"},
		{"ru", "required", "{field} обязательно для заполнения"},
		{"it", "required", "{field} è obbligatorio"},
		{"pt", "required", "{field} é obrigatório"},
		{"de", "required", "{field} ist erforderlich"},
	}

	for _, tc := range testCases {
		t.Run(tc.locale+"_"+tc.msgKey, func(t *testing.T) {
			config, ok := GetLocaleConfig(tc.locale)
			if !ok {
				t.Fatalf("Locale %s not found", tc.locale)
			}

			msg, exists := config.Messages[tc.msgKey]
			if !exists {
				t.Fatalf("Message key %s not found in locale %s", tc.msgKey, tc.locale)
			}

			if msg != tc.wantKey {
				t.Errorf("Expected message '%s' for key %s in locale %s, got '%s'", tc.wantKey, tc.msgKey, tc.locale, msg)
			}
		})
	}
}

func TestCustomRulesInAllLocales(t *testing.T) {
	customRules := []string{"mobile", "idcard", "bankcard", "chinese_name", "strong_password"}
	locales := []string{"en", "zh", "zh-CN", "zh-TW", "ja", "ko", "fr", "es", "ar", "ru", "it", "pt", "de"}

	for _, locale := range locales {
		t.Run(locale, func(t *testing.T) {
			config, ok := GetLocaleConfig(locale)
			if !ok {
				t.Fatalf("Locale %s not found", locale)
			}

			for _, rule := range customRules {
				if _, exists := config.Messages[rule]; !exists {
					t.Errorf("Custom rule %s not found in locale %s", rule, locale)
				}
			}
		})
	}
}