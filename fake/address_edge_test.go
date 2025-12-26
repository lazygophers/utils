package fake

import (
	"testing"
)

// 测试Street函数的各种情况
func TestStreetEdgeCases(t *testing.T) {
	// 创建不同语言的Faker实例
	fakers := []*Faker{
		New(WithLanguage(LanguageEnglish), WithCountry(CountryUS)),
		New(WithLanguage(LanguageChineseSimplified), WithCountry(CountryChina)),
		New(WithLanguage(LanguageChineseTraditional), WithCountry(CountryChina)),
	}

	for _, f := range fakers {
		// 测试Street函数
		street := f.Street()
		if street == "" {
			t.Error("Street() returned empty string")
		}
	}
}

// 测试CountryName函数的各种情况
func TestCountryNameEdgeCases(t *testing.T) {
	// 测试所有支持的语言
	languages := []Language{
		LanguageEnglish,
		LanguageChineseSimplified,
		LanguageChineseTraditional,
		LanguageSpanish,
		LanguageFrench,
		LanguageRussian,
		LanguagePortuguese,
	}

	for _, lang := range languages {
		f := New(WithLanguage(lang), WithCountry(CountryUS))
		country := f.CountryName()
		if country == "" {
			t.Errorf("CountryName() returned empty string for language %s", lang)
		}
	}
}

// 测试FullAddress函数的各种情况
func TestFullAddressEdgeCases(t *testing.T) {
	// 测试不同国家和语言组合
	tests := []struct {
		language Language
		country  Country
	}{}

	// 为所有语言和国家组合生成测试用例
	languages := []Language{
		LanguageEnglish,
		LanguageChineseSimplified,
		LanguageChineseTraditional,
	}

	countries := []Country{
		CountryUS,
		CountryUK, // UK没有州的概念
		CountryChina,
		CountryCanada,
	}

	for _, lang := range languages {
		for _, country := range countries {
			tests = append(tests, struct {
				language Language
				country  Country
			}{lang, country})
		}
	}

	for _, tt := range tests {
		f := New(WithLanguage(tt.language), WithCountry(tt.country))
		address := f.FullAddress()
		if address == nil {
			t.Errorf("FullAddress() returned nil for language %s, country %s", tt.language, tt.country)
			continue
		}

		// 验证地址字段
		if address.Street == "" {
			t.Errorf("Address.Street is empty for language %s, country %s", tt.language, tt.country)
		}
		if address.City == "" {
			t.Errorf("Address.City is empty for language %s, country %s", tt.language, tt.country)
		}
		if address.ZipCode == "" {
			t.Errorf("Address.ZipCode is empty for language %s, country %s", tt.language, tt.country)
		}
		if address.Country == "" {
			t.Errorf("Address.Country is empty for language %s, country %s", tt.language, tt.country)
		}
		if address.FullAddress == "" {
			t.Errorf("Address.FullAddress is empty for language %s, country %s", tt.language, tt.country)
		}
	}
}

// 测试AddressLine函数的各种情况
func TestAddressLineEdgeCases(t *testing.T) {
	// 测试不同国家和语言组合
	languages := []Language{
		LanguageEnglish,
		LanguageChineseSimplified,
		LanguageChineseTraditional,
	}

	countries := []Country{
		CountryUS,
		CountryUK, // UK没有州的概念
		CountryChina,
		CountryCanada,
	}

	for _, lang := range languages {
		for _, country := range countries {
			f := New(WithLanguage(lang), WithCountry(country))
			addressLine := f.AddressLine()
			if addressLine == "" {
				t.Errorf("AddressLine() returned empty string for language %s, country %s", lang, country)
			}
		}
	}
}

// 测试ZipCode函数的各种情况
func TestZipCodeEdgeCases(t *testing.T) {
	// 测试不同国家的邮政编码格式
	countries := []Country{
		CountryUS,
		CountryChina,
		CountryUK,
		CountryCanada,
	}

	for _, country := range countries {
		f := New(WithLanguage(LanguageEnglish), WithCountry(country))
		zipCode := f.ZipCode()
		if zipCode == "" {
			t.Errorf("ZipCode() returned empty string for country %s", country)
		}
	}
}

