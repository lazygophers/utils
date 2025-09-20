package fake

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestPhoneNumberAllCountries tests PhoneNumber function for all supported countries
func TestPhoneNumberAllCountries(t *testing.T) {
	countries := []Country{
		CountryUS,
		CountryCanada,
		CountryChina,
		CountryUK,
		CountryFrance,
		CountryGermany,
		CountryJapan,
		CountryKorea,
		Country("UnknownCountry"), // Test default case
	}

	for _, country := range countries {
		t.Run("Country_"+string(country), func(t *testing.T) {
			faker := New(WithCountry(country))

			// Generate multiple phone numbers to test different paths
			for i := 0; i < 10; i++ {
				phone := faker.PhoneNumber()
				assert.NotEmpty(t, phone)
				// Phone numbers should contain digits
				assert.Regexp(t, `\d`, phone)
			}
		})
	}
}

// TestGenerateCompanyNameAllPatterns tests all company name generation patterns
func TestGenerateCompanyNameAllPatterns(t *testing.T) {
	faker := New()

	// Generate many company names to hit all patterns
	patterns := make(map[string]bool)

	for i := 0; i < 200; i++ { // Generate enough to hit all patterns
		company := faker.CompanyName()
		assert.NotEmpty(t, company)

		// Identify pattern type
		if strings.Contains(company, " & ") {
			patterns["%s & %s"] = true
		} else if strings.Contains(company, "-") {
			patterns["%s-%s"] = true
		} else if strings.Count(company, " ") == 2 {
			patterns["%s %s %s"] = true
		} else if strings.Count(company, " ") == 1 {
			patterns["%s %s"] = true
		} else {
			patterns["%s%s"] = true
		}
	}

	// Should have hit multiple patterns
	assert.GreaterOrEqual(t, len(patterns), 3, "Should generate companies with different patterns")
}

// TestCatchphraseGeneration tests the Catchphrase function thoroughly
func TestCatchphraseGeneration(t *testing.T) {

	// Test multiple languages if supported
	languages := []Language{
		LanguageEnglish,
		LanguageChineseSimplified,
		LanguageFrench,
		LanguageRussian,
		LanguageSpanish,
		LanguagePortuguese,
	}

	for _, lang := range languages {
		t.Run("Language_"+string(lang), func(t *testing.T) {
			faker := New(WithLanguage(lang))

			// Generate multiple catchphrases to test different patterns
			for i := 0; i < 20; i++ {
				catchphrase := faker.Catchphrase()
				assert.NotEmpty(t, catchphrase)
				// Catchphrases should have some structure
				assert.True(t, len(catchphrase) > 5, "Catchphrase should be meaningful length")
			}
		})
	}
}

// TestStreetGeneration tests Street function with different languages
func TestStreetGeneration(t *testing.T) {
	languages := []Language{
		LanguageEnglish,
		LanguageChineseSimplified,
		LanguageFrench,
		LanguageRussian,
		LanguageSpanish,
		LanguagePortuguese,
	}

	for _, lang := range languages {
		t.Run("Street_"+string(lang), func(t *testing.T) {
			faker := New(WithLanguage(lang))

			// Generate multiple streets to test different patterns
			for i := 0; i < 15; i++ {
				street := faker.Street()
				assert.NotEmpty(t, street)
			}
		})
	}
}

// TestCityGeneration tests City function with different languages
func TestCityGeneration(t *testing.T) {
	languages := []Language{
		LanguageEnglish,
		LanguageChineseSimplified,
		LanguageFrench,
		LanguageRussian,
		LanguageSpanish,
		LanguageRussian,
		LanguagePortuguese,
	}

	for _, lang := range languages {
		t.Run("City_"+string(lang), func(t *testing.T) {
			faker := New(WithLanguage(lang))

			// Generate multiple cities to test different data sources
			cities := make(map[string]bool)
			for i := 0; i < 20; i++ {
				city := faker.City()
				assert.NotEmpty(t, city)
				cities[city] = true
			}

			// Should generate some variety
			assert.GreaterOrEqual(t, len(cities), 2, "Should generate different cities")
		})
	}
}

// TestStateGeneration tests State function with different languages
func TestStateGeneration(t *testing.T) {
	languages := []Language{
		LanguageEnglish,
		LanguageChineseSimplified,
		LanguageFrench,
		LanguageRussian,
		LanguageSpanish,
		LanguagePortuguese,
	}

	for _, lang := range languages {
		t.Run("State_"+string(lang), func(t *testing.T) {
			faker := New(WithLanguage(lang))

			// Generate multiple states to test different data sources
			for i := 0; i < 15; i++ {
				state := faker.State()
				assert.NotEmpty(t, state)
			}
		})
	}
}

// TestCountryNameGeneration tests CountryName function with different languages
func TestCountryNameGeneration(t *testing.T) {
	languages := []Language{
		LanguageEnglish,
		LanguageChineseSimplified,
		LanguageFrench,
		LanguageRussian,
		LanguageSpanish,
		LanguagePortuguese,
	}

	for _, lang := range languages {
		t.Run("CountryName_"+string(lang), func(t *testing.T) {
			faker := New(WithLanguage(lang))

			// Generate multiple country names to test different data sources
			countries := make(map[string]bool)
			for i := 0; i < 20; i++ {
				country := faker.CountryName()
				assert.NotEmpty(t, country)
				countries[country] = true
			}

			// Should generate some variety
			assert.GreaterOrEqual(t, len(countries), 3, "Should generate different countries")
		})
	}
}

// TestAddressLineGeneration tests AddressLine function with different languages
func TestAddressLineGeneration(t *testing.T) {
	languages := []Language{
		LanguageEnglish,
		LanguageChineseSimplified,
		LanguageFrench,
		LanguageRussian,
		LanguageSpanish,
		LanguagePortuguese,
	}

	for _, lang := range languages {
		t.Run("AddressLine_"+string(lang), func(t *testing.T) {
			faker := New(WithLanguage(lang))

			// Generate multiple address lines to test different formats
			for i := 0; i < 15; i++ {
				addressLine := faker.AddressLine()
				assert.NotEmpty(t, addressLine)
				// Address lines should contain some structure
				assert.True(t, len(addressLine) > 3, "Address line should be meaningful")
			}
		})
	}
}

// TestCompanySuffixGeneration tests CompanySuffix function
func TestCompanySuffixGeneration(t *testing.T) {
	languages := []Language{
		LanguageEnglish,
		LanguageChineseSimplified,
		LanguageFrench,
		LanguageRussian,
		LanguageSpanish,
		LanguagePortuguese,
	}

	for _, lang := range languages {
		t.Run("CompanySuffix_"+string(lang), func(t *testing.T) {
			faker := New(WithLanguage(lang))

			// Generate multiple suffixes to test different options
			suffixes := make(map[string]bool)
			for i := 0; i < 20; i++ {
				suffix := faker.CompanySuffix()
				assert.NotEmpty(t, suffix)
				suffixes[suffix] = true
			}

			// Should generate some variety
			assert.GreaterOrEqual(t, len(suffixes), 2, "Should generate different suffixes")
		})
	}
}

// TestJobTitleGeneration tests JobTitle function
func TestJobTitleGeneration(t *testing.T) {
	languages := []Language{
		LanguageEnglish,
		LanguageChineseSimplified,
		LanguageFrench,
		LanguageRussian,
		LanguageSpanish,
		LanguagePortuguese,
	}

	for _, lang := range languages {
		t.Run("JobTitle_"+string(lang), func(t *testing.T) {
			faker := New(WithLanguage(lang))

			// Generate multiple job titles to test different combinations
			titles := make(map[string]bool)
			for i := 0; i < 25; i++ {
				title := faker.JobTitle()
				assert.NotEmpty(t, title)
				titles[title] = true
			}

			// Should generate variety
			assert.GreaterOrEqual(t, len(titles), 5, "Should generate different job titles")
		})
	}
}

// TestBSGeneration tests BS (Business Speak) function
func TestBSGeneration(t *testing.T) {
	faker := New()

	// Generate multiple BS phrases to test different combinations
	phrases := make(map[string]bool)
	for i := 0; i < 30; i++ {
		bs := faker.BS()
		assert.NotEmpty(t, bs)
		phrases[bs] = true

		// BS should have meaningful length and structure
		assert.True(t, len(bs) > 10, "BS phrase should be substantial")
		assert.True(t, strings.Contains(bs, " "), "BS should contain spaces")
	}

	// Should generate variety
	assert.GreaterOrEqual(t, len(phrases), 10, "Should generate different BS phrases")
}

// TestMobileNumberGeneration tests MobileNumber function
func TestMobileNumberGeneration(t *testing.T) {
	countries := []Country{
		CountryUS,
		CountryCanada,
		CountryChina,
		CountryUK,
		CountryFrance,
		CountryGermany,
		CountryJapan,
		CountryKorea,
	}

	for _, country := range countries {
		t.Run("Mobile_"+string(country), func(t *testing.T) {
			faker := New(WithCountry(country))

			// Generate multiple mobile numbers
			for i := 0; i < 10; i++ {
				mobile := faker.MobileNumber()
				assert.NotEmpty(t, mobile)
				// Mobile numbers should contain digits
				assert.Regexp(t, `\d`, mobile)
			}
		})
	}
}

// TestFullAddressGeneration tests FullAddress function with different languages
func TestFullAddressGeneration(t *testing.T) {
	languages := []Language{
		LanguageEnglish,
		LanguageChineseSimplified,
		LanguageFrench,
		LanguageRussian,
		LanguageSpanish,
		LanguagePortuguese,
	}

	for _, lang := range languages {
		t.Run("FullAddress_"+string(lang), func(t *testing.T) {
			faker := New(WithLanguage(lang))

			// Generate multiple full addresses to test different formats
			for i := 0; i < 10; i++ {
				address := faker.FullAddress()
				assert.NotNil(t, address)
				// Full address should be substantial
				assert.True(t, len(address.FullAddress) > 10, "Full address should be comprehensive")
				assert.NotEmpty(t, address.Street)
				assert.NotEmpty(t, address.City)
				assert.NotEmpty(t, address.Country)
			}
		})
	}
}

// TestVariousGenerationMethods tests other generation methods to improve coverage
func TestVariousGenerationMethods(t *testing.T) {
	faker := New()

	t.Run("ComprehensiveGeneration", func(t *testing.T) {
		// Test multiple generations to hit different code paths
		for i := 0; i < 50; i++ {
			// Test phone numbers
			phone := faker.PhoneNumber()
			assert.NotEmpty(t, phone)

			// Test company names
			company := faker.CompanyName()
			assert.NotEmpty(t, company)

			// Test catchphrases
			catchphrase := faker.Catchphrase()
			assert.NotEmpty(t, catchphrase)

			// Test addresses
			street := faker.Street()
			assert.NotEmpty(t, street)

			city := faker.City()
			assert.NotEmpty(t, city)

			state := faker.State()
			assert.NotEmpty(t, state)

			country := faker.CountryName()
			assert.NotEmpty(t, country)
		}
	})
}

// TestEdgeCasesAndPatterns tests edge cases to improve coverage
func TestEdgeCasesAndPatterns(t *testing.T) {
	t.Run("DifferentConfigurations", func(t *testing.T) {
		// Test different faker configurations
		configs := []struct {
			name     string
			language Language
			country  Country
		}{
			{"English_US", LanguageEnglish, CountryUS},
			{"Chinese_China", LanguageChineseSimplified, CountryChina},
			{"French_France", LanguageFrench, CountryFrance},
			{"Spanish_Spain", LanguageSpanish, CountrySpain},
			{"Portuguese_Brazil", LanguagePortuguese, CountryBrazil},
			{"Russian_Russia", LanguageRussian, CountryRussia},
			{"English_UK", LanguageEnglish, CountryUK},
			{"English_Canada", LanguageEnglish, CountryCanada},
		}

		for _, config := range configs {
			t.Run(config.name, func(t *testing.T) {
				faker := New(WithLanguage(config.language), WithCountry(config.country))

				// Test all main functions with this configuration
				phone := faker.PhoneNumber()
				assert.NotEmpty(t, phone)

				mobile := faker.MobileNumber()
				assert.NotEmpty(t, mobile)

				company := faker.CompanyName()
				assert.NotEmpty(t, company)

				street := faker.Street()
				assert.NotEmpty(t, street)

				city := faker.City()
				assert.NotEmpty(t, city)

				if config.language == LanguageEnglish {
					// Test catchphrase for English
					catchphrase := faker.Catchphrase()
					assert.NotEmpty(t, catchphrase)

					// Test BS for English
					bs := faker.BS()
					assert.NotEmpty(t, bs)
				}
			})
		}
	})
}