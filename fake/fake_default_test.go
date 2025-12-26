package fake

import (
	"context"
	"testing"
)

func TestSetDefaultLanguage(t *testing.T) {
	originalLang := GetDefaultLanguage()
	defer SetDefaultLanguage(originalLang)

	SetDefaultLanguage(LanguageChineseSimplified)
	if GetDefaultLanguage() != LanguageChineseSimplified {
		t.Errorf("Expected %s, got %s", LanguageChineseSimplified, GetDefaultLanguage())
	}
}

func TestSetDefaultCountry(t *testing.T) {
	originalCountry := GetDefaultCountry()
	defer SetDefaultCountry(originalCountry)

	SetDefaultCountry(CountryChina)
	if GetDefaultCountry() != CountryChina {
		t.Errorf("Expected %s, got %s", CountryChina, GetDefaultCountry())
	}
}

func TestSetDefaultGender(t *testing.T) {
	originalGender := GetDefaultGender()
	defer SetDefaultGender(originalGender)

	SetDefaultGender(GenderFemale)
	if GetDefaultGender() != GenderFemale {
		t.Errorf("Expected %s, got %s", GenderFemale, GetDefaultGender())
	}
}

func TestSetDefaultSeed(t *testing.T) {
	originalSeed := GetDefaultSeed()
	defer SetDefaultSeed(originalSeed)

	testSeed := int64(12345)
	SetDefaultSeed(testSeed)
	if GetDefaultSeed() != testSeed {
		t.Errorf("Expected %d, got %d", testSeed, GetDefaultSeed())
	}
}

func TestGetDefaultLanguage(t *testing.T) {
	lang := GetDefaultLanguage()
	if lang == "" {
		t.Error("GetDefaultLanguage should not return empty string")
	}
}

func TestGetDefaultCountry(t *testing.T) {
	country := GetDefaultCountry()
	if country == "" {
		t.Error("GetDefaultCountry should not return empty string")
	}
}

func TestGetDefaultGender(t *testing.T) {
	gender := GetDefaultGender()
	if gender == "" {
		t.Error("GetDefaultGender should not return empty string")
	}
}

func TestGetDefaultSeed(t *testing.T) {
	seed := GetDefaultSeed()
	if seed == 0 {
		t.Error("GetDefaultSeed should not return 0")
	}
}

func TestResetDefaultFaker(t *testing.T) {
	originalLang := GetDefaultLanguage()
	originalCountry := GetDefaultCountry()

	SetDefaultLanguage(LanguageChineseSimplified)
	SetDefaultCountry(CountryChina)

	ResetDefaultFaker()

	newLang := GetDefaultLanguage()
	newCountry := GetDefaultCountry()

	if newLang == LanguageChineseSimplified {
		t.Error("ResetDefaultFaker should reset language to default")
	}
	if newCountry == CountryChina {
		t.Error("ResetDefaultFaker should reset country to default")
	}

	SetDefaultLanguage(originalLang)
	SetDefaultCountry(originalCountry)
}

func TestSetDefaultFaker(t *testing.T) {
	originalFaker := getDefaultFaker()
	defer SetDefaultFaker(originalFaker)

	customFaker := New(WithLanguage(LanguageChineseSimplified), WithCountry(CountryChina))
	SetDefaultFaker(customFaker)

	if GetDefaultLanguage() != LanguageChineseSimplified {
		t.Errorf("Expected %s, got %s", LanguageChineseSimplified, GetDefaultLanguage())
	}
	if GetDefaultCountry() != CountryChina {
		t.Errorf("Expected %s, got %s", CountryChina, GetDefaultCountry())
	}
}

func TestSetDefaults(t *testing.T) {
	originalLang := GetDefaultLanguage()
	originalCountry := GetDefaultCountry()
	originalGender := GetDefaultGender()
	defer func() {
		SetDefaults(WithLanguage(originalLang), WithCountry(originalCountry), WithGender(originalGender))
	}()

	SetDefaults(
		WithLanguage(LanguageChineseSimplified),
		WithCountry(CountryChina),
		WithGender(GenderFemale),
	)

	if GetDefaultLanguage() != LanguageChineseSimplified {
		t.Errorf("Expected %s, got %s", LanguageChineseSimplified, GetDefaultLanguage())
	}
	if GetDefaultCountry() != CountryChina {
		t.Errorf("Expected %s, got %s", CountryChina, GetDefaultCountry())
	}
	if GetDefaultGender() != GenderFemale {
		t.Errorf("Expected %s, got %s", GenderFemale, GetDefaultGender())
	}
}

func TestClearDefaultCache(t *testing.T) {
	faker := getDefaultFaker()
	faker.setCachedData("test_key", "test_value")

	ClearDefaultCache()

	if _, exists := faker.getCachedData("test_key"); exists {
		t.Error("ClearDefaultCache should clear all cached data")
	}
}

func TestGetCachedData(t *testing.T) {
	faker := New()

	data, exists := faker.getCachedData("non_existent_key")
	if exists {
		t.Error("getCachedData should return false for non-existent key")
	}
	if data != nil {
		t.Error("getCachedData should return nil for non-existent key")
	}

	faker.setCachedData("test_key", "test_value")
	data, exists = faker.getCachedData("test_key")
	if !exists {
		t.Error("getCachedData should return true for existing key")
	}
	if data != "test_value" {
		t.Errorf("Expected 'test_value', got %v", data)
	}
}

func TestSetCachedData(t *testing.T) {
	faker := New()

	faker.setCachedData("test_key", "test_value")
	data, exists := faker.getCachedData("test_key")
	if !exists {
		t.Error("setCachedData should store data")
	}
	if data != "test_value" {
		t.Errorf("Expected 'test_value', got %v", data)
	}

	faker.setCachedData("test_key", "new_value")
	data, _ = faker.getCachedData("test_key")
	if data != "new_value" {
		t.Errorf("Expected 'new_value', got %v", data)
	}
}

func TestValidateLanguage(t *testing.T) {
	faker := New()

	validLanguages := []Language{
		LanguageEnglish,
		LanguageChineseSimplified,
		LanguageChineseTraditional,
		LanguageFrench,
		LanguageRussian,
		LanguagePortuguese,
		LanguageSpanish,
	}

	for _, lang := range validLanguages {
		if !faker.validateLanguage(lang) {
			t.Errorf("validateLanguage should return true for %s", lang)
		}
	}

	if faker.validateLanguage(Language("invalid")) {
		t.Error("validateLanguage should return false for invalid language")
	}
}

func TestValidateCountry(t *testing.T) {
	faker := New()

	validCountries := []Country{
		CountryChina, CountryUS, CountryUK, CountryFrance, CountryGermany,
		CountryJapan, CountryKorea, CountryRussia, CountryBrazil, CountrySpain,
		CountryPortugal, CountryItaly, CountryCanada, CountryAustralia, CountryIndia,
	}

	for _, country := range validCountries {
		if !faker.validateCountry(country) {
			t.Errorf("validateCountry should return true for %s", country)
		}
	}

	if faker.validateCountry(Country("invalid")) {
		t.Error("validateCountry should return false for invalid country")
	}
}

func TestGetSupportedLanguages(t *testing.T) {
	languages := GetSupportedLanguages()
	if len(languages) == 0 {
		t.Error("GetSupportedLanguages should return non-empty list")
	}

	expectedLanguages := []Language{
		LanguageEnglish,
		LanguageChineseSimplified,
		LanguageChineseTraditional,
		LanguageFrench,
		LanguageRussian,
		LanguagePortuguese,
		LanguageSpanish,
	}

	if len(languages) != len(expectedLanguages) {
		t.Errorf("Expected %d languages, got %d", len(expectedLanguages), len(languages))
	}
}

func TestGetSupportedCountries(t *testing.T) {
	countries := GetSupportedCountries()
	if len(countries) == 0 {
		t.Error("GetSupportedCountries should return non-empty list")
	}

	expectedCountries := []Country{
		CountryChina, CountryUS, CountryUK, CountryFrance, CountryGermany,
		CountryJapan, CountryKorea, CountryRussia, CountryBrazil, CountrySpain,
		CountryPortugal, CountryItaly, CountryCanada, CountryAustralia, CountryIndia,
	}

	if len(countries) != len(expectedCountries) {
		t.Errorf("Expected %d countries, got %d", len(expectedCountries), len(countries))
	}
}

func TestFormatWithParams(t *testing.T) {
	format := "Hello {{name}}, your age is {{age}}"
	params := map[string]string{
		"name": "John",
		"age":  "30",
	}

	result := formatWithParams(format, params)
	expected := "Hello John, your age is 30"

	if result != expected {
		t.Errorf("Expected '%s', got '%s'", expected, result)
	}

	result = formatWithParams(format, map[string]string{})
	if result != format {
		t.Error("formatWithParams should not modify format when params is empty")
	}
}

func TestFormatNumber(t *testing.T) {
	result := formatNumber("#", 123)
	expected := "123"

	if result != expected {
		t.Errorf("Expected '%s', got '%s'", expected, result)
	}

	result = formatNumber("Number: #", 456)
	expected = "Number: 456"

	if result != expected {
		t.Errorf("Expected '%s', got '%s'", expected, result)
	}
}

func TestBatchGenerate(t *testing.T) {
	faker := New()

	count := 5
	generator := func() string {
		return faker.FirstName()
	}

	results := faker.batchGenerate(count, generator)

	if len(results) != count {
		t.Errorf("Expected %d results, got %d", count, len(results))
	}

	for _, result := range results {
		if result == "" {
			t.Error("batchGenerate should not return empty strings")
		}
	}
}

func TestContextWithLanguage(t *testing.T) {
	ctx := context.Background()
	ctx = ContextWithLanguage(ctx, LanguageChineseSimplified)

	lang, ok := ctx.Value(contextKeyLanguage).(Language)
	if !ok {
		t.Error("ContextWithLanguage should store language in context")
	}
	if lang != LanguageChineseSimplified {
		t.Errorf("Expected %s, got %s", LanguageChineseSimplified, lang)
	}
}

func TestContextWithCountry(t *testing.T) {
	ctx := context.Background()
	ctx = ContextWithCountry(ctx, CountryChina)

	country, ok := ctx.Value(contextKeyCountry).(Country)
	if !ok {
		t.Error("ContextWithCountry should store country in context")
	}
	if country != CountryChina {
		t.Errorf("Expected %s, got %s", CountryChina, country)
	}
}

func TestContextWithGender(t *testing.T) {
	ctx := context.Background()
	ctx = ContextWithGender(ctx, GenderFemale)

	gender, ok := ctx.Value(contextKeyGender).(Gender)
	if !ok {
		t.Error("ContextWithGender should store gender in context")
	}
	if gender != GenderFemale {
		t.Errorf("Expected %s, got %s", GenderFemale, gender)
	}
}

func TestDefaultSettings(t *testing.T) {
	lang := GetDefaultLanguage()
	country := GetDefaultCountry()
	gender := GetDefaultGender()
	seed := GetDefaultSeed()

	if lang == "" || country == "" || gender == "" || seed == 0 {
		t.Error("Default settings should be initialized")
	}
}
