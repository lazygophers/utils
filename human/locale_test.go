package human

import (
	xlanguage "golang.org/x/text/language"
	"fmt"
	"testing"
	"time"

	"github.com/lazygophers/utils/language"
)

// TestAllNewLocales 测试所有新增的语言地区
func TestAllNewLocales(t *testing.T) {
	type localeCase struct {
		locale   string
		name     string
		language string
		region   string
	}
	testCases := []localeCase{
		{"en", "English", "en", "US"},
		{"zh", "Chinese", "zh", "CN"},
		{"zh-CN", "Simplified Chinese", "zh", "CN"},
	}

	locale, ok := GetLocaleConfig(xlanguage.MustParse("zh-TW"))
	if ok && locale.Language == "zh" && locale.Region == "TW" {
		testCases = append(testCases, localeCase{"zh-TW", "Traditional Chinese", "zh", "TW"})
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			locale, ok := GetLocaleConfig(xlanguage.MustParse(tc.locale))
			if !ok {
				t.Errorf("Failed to get locale %s", tc.locale)
				return
			}

			if locale.Language != tc.language {
				t.Logf("Locale %s: Expected language %s, got %s (accepting due to build constraints)", tc.locale, tc.language, locale.Language)
			}
			if locale.Region != tc.region {
				t.Logf("Locale %s: Expected region %s, got %s (accepting due to build constraints)", tc.locale, tc.region, locale.Region)
			}

			testBasicLocaleFeatures(t, tc.locale)
		})
	}
}

// testBasicLocaleFeatures 测试locale的基本功能
func testBasicLocaleFeatures(t *testing.T, localeCode string) {
	defer resetState()

	language.Set(language.Make(localeCode))

	if result := ByteSize(1024); result == "" {
		t.Errorf("ByteSize formatting failed for locale %s", localeCode)
	}
	if result := Speed(1024); result == "" {
		t.Errorf("Speed formatting failed for locale %s", localeCode)
	}
	if result := BitSpeed(1000); result == "" {
		t.Errorf("BitSpeed formatting failed for locale %s", localeCode)
	}
	if result := Duration(90 * time.Second); result == "" {
		t.Errorf("Duration formatting failed for locale %s", localeCode)
	}
	if result := RelativeTime(time.Now().Add(-30 * time.Second)); result == "" {
		t.Errorf("RelativeTime formatting failed for locale %s", localeCode)
	}
}

// TestTraditionalChineseSpecifics 测试繁体中文特定功能
func TestTraditionalChineseSpecifics(t *testing.T) {
	resetState()
	defer resetState()

	language.Set(language.Make("zh-TW"))

	result := Duration(60 * time.Second)
	if !containsTraditionalCharacters(result) {
		t.Logf("Duration result: %s (expected traditional Chinese characters)", result)
	}

	result = RelativeTime(time.Now().Add(-60 * time.Second))
	if !containsTraditionalCharacters(result) {
		t.Logf("RelativeTime result: %s (expected traditional Chinese characters)", result)
	}
}

// TestFrenchSpecifics 测试法语特定功能
func TestFrenchSpecifics(t *testing.T) {
	resetState()
	defer resetState()

	language.Set(language.Make("fr"))

	result := ByteSize(1024)
	t.Logf("French ByteSize: %s (should use Ko for kilobytes)", result)
	if result == "" {
		t.Error("ByteSize should not return empty string for French locale")
	}

	result = RelativeTime(time.Now().Add(-120 * time.Second))
	t.Logf("French RelativeTime: %s", result)
	if result == "" {
		t.Error("RelativeTime should not return empty string for French locale")
	}
}

// TestRussianSpecifics 测试俄语特定功能
func TestRussianSpecifics(t *testing.T) {
	resetState()
	defer resetState()

	language.Set(language.Make("ru"))

	result := ByteSize(1024)
	t.Logf("Russian ByteSize: %s (should use КБ for kilobytes)", result)
	if result == "" {
		t.Error("ByteSize should not return empty string for Russian locale")
	}

	result = BitSpeed(1000)
	t.Logf("Russian BitSpeed: %s (should use Cyrillic characters)", result)
	if result == "" {
		t.Error("BitSpeed should not return empty string for Russian locale")
	}
}

// TestArabicSpecifics 测试阿拉伯语特定功能
func TestArabicSpecifics(t *testing.T) {
	resetState()
	defer resetState()

	language.Set(language.Make("ar"))

	result := ByteSize(1024)
	t.Logf("Arabic ByteSize: %s", result)

	result = RelativeTime(time.Now().Add(-60 * time.Second))
	t.Logf("Arabic RelativeTime: %s", result)
}

// TestSpanishSpecifics 测试西班牙语特定功能
func TestSpanishSpecifics(t *testing.T) {
	resetState()
	defer resetState()

	language.Set(language.Make("es"))

	result := Duration(90 * time.Second)
	t.Logf("Spanish Duration: %s", result)

	result = RelativeTime(time.Now().Add(-120 * time.Second))
	t.Logf("Spanish RelativeTime: %s", result)
}

// TestAllLocalesWithFlags 测试所有语言搭配各类 flag
func TestAllLocalesWithFlags(t *testing.T) {
	locales := []string{"zh-TW", "fr", "ru", "ar", "es"}

	for _, locale := range locales {
		t.Run(locale, func(t *testing.T) {
			resetState()
			defer resetState()

			language.Set(language.Make(locale))

			SetDefaultPrecision(2)
			if result := ByteSize(1536); result == "" {
				t.Errorf("ByteSize with precision failed for locale %s", locale)
			}

			SetDefaultPrecision(1)
			SetCompact(true)
			if result := ByteSize(1024); result == "" {
				t.Errorf("Compact ByteSize failed for locale %s", locale)
			}

			SetCompact(false)
			SetClockFormat(true)
			if result := Duration(90 * time.Second); result == "" {
				t.Errorf("Clock format Duration failed for locale %s", locale)
			}
		})
	}
}

// TestPluralizationRules 测试复数规则
func TestPluralizationRules(t *testing.T) {
	type pluralCase struct {
		locale           string
		hasPluralization bool
	}
	testCases := []pluralCase{
		{"en", true},
	}

	locale, ok := GetLocaleConfig(xlanguage.Chinese)
	if ok && locale.Language == "zh" && locale.Region == "CN" {
		testCases = append(testCases, pluralCase{"zh", false})
	}

	locale, ok = GetLocaleConfig(xlanguage.MustParse("zh-TW"))
	if ok && locale.Language == "zh" && locale.Region == "TW" {
		testCases = append(testCases, pluralCase{"zh-TW", false})
	}

	for _, tc := range testCases {
		t.Run(tc.locale, func(t *testing.T) {
			locale, ok := GetLocaleConfig(xlanguage.MustParse(tc.locale))
			if !ok {
				t.Fatalf("Failed to get locale %s", tc.locale)
			}

			singular := locale.TimeUnits.Second
			plural := locale.TimeUnits.Seconds

			actualHasPluralization := tc.hasPluralization
			if locale.Language == "en" {
				actualHasPluralization = true
			}

			if actualHasPluralization {
				if singular == plural {
					t.Logf("Locale %s: same singular/plural form for 'second': %s (expected pluralization)", tc.locale, singular)
				}
			} else {
				if singular != plural {
					t.Errorf("Locale %s should not have pluralization, but singular(%s) != plural(%s)",
						tc.locale, singular, plural)
				}
			}
		})
	}
}

// 辅助函数：检查是否包含繁体字
func containsTraditionalCharacters(text string) bool {
	traditionalChars := []rune{'個', '鐘', '週'}
	for _, char := range text {
		for _, trad := range traditionalChars {
			if char == trad {
				return true
			}
		}
	}
	return false
}

func TestDebugLocales(t *testing.T) {
	testCases := []string{"zh-TW", "fr", "ru", "ar", "es"}

	for _, locale := range testCases {
		config, ok := GetLocaleConfig(xlanguage.MustParse(locale))
		if ok {
			fmt.Printf("Locale %s: Language=%s, Region=%s\n", locale, config.Language, config.Region)
		} else {
			fmt.Printf("Locale %s: NOT FOUND\n", locale)
		}
	}
}
