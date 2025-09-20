package fake

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// TestInternalDataLoaderFunctions tests internal data loader functions
func TestInternalDataLoaderFunctions(t *testing.T) {
	t.Run("Data loader internal functions", func(t *testing.T) {
		// Force the usage of internal data loader functions
		// These are private but get called through public APIs
		faker := New(WithLanguage(LanguageEnglish))

		// Generate data to trigger internal data loading
		street := faker.Street()
		assert.NotEmpty(t, street)

		// Generate more data to trigger different data paths
		city := faker.City()
		assert.NotEmpty(t, city)

		// Test with different language to trigger different data paths
		fakerZH := New(WithLanguage(LanguageChineseSimplified))
		streetZH := fakerZH.Street()
		assert.NotEmpty(t, streetZH)
	})
}

// TestCachingMechanisms tests caching functionality
func TestCachingMechanisms(t *testing.T) {
	t.Run("Cache mechanisms", func(t *testing.T) {
		faker := New()

		// Generate some data to populate cache
		for i := 0; i < 10; i++ {
			_ = faker.FirstName()
			_ = faker.LastName()
			_ = faker.Email()
		}

		// Get stats to trigger cache hit counting
		stats := faker.Stats()
		assert.NotNil(t, stats)
		assert.Contains(t, stats, "call_count")
		assert.Contains(t, stats, "cache_hits")
		assert.Contains(t, stats, "generated_data")

		// Clear cache to test cache clearing
		faker.ClearCache()

		// Get stats again
		stats2 := faker.Stats()
		assert.NotNil(t, stats2)
	})
}

// TestValidationFunctionsComprehensive tests validation functions comprehensively
func TestValidationFunctionsComprehensive(t *testing.T) {
	t.Run("Language and country validation", func(t *testing.T) {
		// Test with various language configurations to trigger validation
		languages := []Language{
			LanguageEnglish,
			LanguageChineseSimplified,
			LanguageChineseTraditional,
			LanguageFrench,
			LanguageRussian,
			LanguagePortuguese,
			LanguageSpanish,
		}

		for _, lang := range languages {
			faker := New(WithLanguage(lang))
			name := faker.FirstName()
			assert.NotEmpty(t, name)
		}

		// Test with various country configurations
		countries := []Country{
			CountryUS,
			CountryChina,
			CountryUK,
			CountryFrance,
			CountryGermany,
			CountryJapan,
		}

		for _, country := range countries {
			faker := New(WithCountry(country))
			zipCode := faker.ZipCode()
			assert.NotEmpty(t, zipCode)
		}
	})
}

// TestFormatFunctionsComprehensive tests format helper functions comprehensively
func TestFormatFunctionsComprehensive(t *testing.T) {
	t.Run("Format helper functions", func(t *testing.T) {
		faker := New()

		// Generate data that uses format functions internally
		// Phone numbers use formatNumber and formatWithParams
		for i := 0; i < 5; i++ {
			phone := faker.PhoneNumber()
			assert.NotEmpty(t, phone)
		}

		// Email addresses may use formatting
		for i := 0; i < 5; i++ {
			email := faker.Email()
			assert.NotEmpty(t, email)
			assert.Contains(t, email, "@")
		}

		// Credit card numbers use formatting
		for i := 0; i < 5; i++ {
			ccNumber := faker.CreditCardNumber()
			assert.NotEmpty(t, ccNumber)
		}
	})
}

// TestBatchFunctionsComprehensive tests batch generation functions comprehensively
func TestBatchFunctionsComprehensive(t *testing.T) {
	t.Run("Batch generation functions", func(t *testing.T) {
		faker := New()

		// Test batch generation for addresses
		streets := faker.BatchStreets(5)
		assert.Len(t, streets, 5)
		for _, street := range streets {
			assert.NotEmpty(t, street)
		}

		cities := faker.BatchCities(3)
		assert.Len(t, cities, 3)
		for _, city := range cities {
			assert.NotEmpty(t, city)
		}

		states := faker.BatchStates(4)
		assert.Len(t, states, 4)

		zipCodes := faker.BatchZipCodes(6)
		assert.Len(t, zipCodes, 6)
		for _, zip := range zipCodes {
			assert.NotEmpty(t, zip)
		}

		addresses := faker.BatchFullAddresses(2)
		assert.Len(t, addresses, 2)
		for _, addr := range addresses {
			assert.NotNil(t, addr)
			assert.NotEmpty(t, addr.Street)
			assert.NotEmpty(t, addr.City)
		}
	})
}

// TestIdentityFunctionsCoverage tests more identity functions
func TestIdentityFunctionsCoverage(t *testing.T) {
	t.Run("Identity functions comprehensive", func(t *testing.T) {
		faker := New()

		// Test various identity document functions
		ssn := faker.SSN()
		assert.NotEmpty(t, ssn)

		chineseID := faker.ChineseIDNumber()
		assert.NotEmpty(t, chineseID)

		passport := faker.Passport()
		assert.NotEmpty(t, passport)

		license := faker.DriversLicense()
		assert.NotEmpty(t, license)

		// Test credit card related functions
		ccInfo := faker.CreditCardInfo()
		assert.NotNil(t, ccInfo)
		assert.NotEmpty(t, ccInfo.Number)
		assert.NotEmpty(t, ccInfo.CVV)

		cvv := faker.CVV()
		assert.NotEmpty(t, cvv)

		bankAccount := faker.BankAccount()
		assert.NotEmpty(t, bankAccount)

		iban := faker.IBAN()
		assert.NotEmpty(t, iban)

		// Test identity document structure
		identityDoc := faker.IdentityDoc()
		assert.NotNil(t, identityDoc)
	})
}

// TestContextFunctions tests context-related functions
func TestContextFunctions(t *testing.T) {
	t.Run("Context integration", func(t *testing.T) {
		// Test context with language
		ctx := context.Background()
		ctx = ContextWithLanguage(ctx, LanguageChineseSimplified)
		ctx = ContextWithCountry(ctx, CountryChina)
		ctx = ContextWithGender(ctx, GenderFemale)

		faker := WithContext(ctx)

		name := faker.FirstName()
		assert.NotEmpty(t, name)

		address := faker.FullAddress()
		assert.NotNil(t, address)
		assert.NotEmpty(t, address.Country)
	})
}

// TestDefaultFallbacks tests default fallback functions
func TestDefaultFallbacks(t *testing.T) {
	t.Run("Default fallback mechanisms", func(t *testing.T) {
		// Test scenarios that might trigger default fallbacks

		// Test with an unsupported language (should fall back)
		faker := New(WithLanguage("unsupported"))
		name := faker.FirstName()
		assert.NotEmpty(t, name)

		// Test with unsupported country (should fall back)
		faker2 := New(WithCountry("XX"))
		zipCode := faker2.ZipCode()
		assert.NotEmpty(t, zipCode)

		// Test scenarios that might trigger getDefaultFirstName and getDefaultLastName
		// by generating many names to potentially exhaust data sources
		for i := 0; i < 100; i++ {
			firstName := faker.FirstName()
			lastName := faker.LastName()
			assert.NotEmpty(t, firstName)
			assert.NotEmpty(t, lastName)
		}
	})
}

// TestGlobalDefaultFunctions tests global default configuration functions
func TestGlobalDefaultFunctions(t *testing.T) {
	t.Run("Global default configuration", func(t *testing.T) {
		// Save original defaults
		originalLang := GetDefaultLanguage()
		originalCountry := GetDefaultCountry()
		originalGender := GetDefaultGender()
		originalSeed := GetDefaultSeed()

		// Test setting defaults
		SetDefaultLanguage(LanguageChineseSimplified)
		assert.Equal(t, LanguageChineseSimplified, GetDefaultLanguage())

		SetDefaultCountry(CountryChina)
		assert.Equal(t, CountryChina, GetDefaultCountry())

		SetDefaultGender(GenderFemale)
		assert.Equal(t, GenderFemale, GetDefaultGender())

		SetDefaultSeed(12345)
		assert.Equal(t, int64(12345), GetDefaultSeed())

		// Test batch setting
		SetDefaults(
			WithLanguage(LanguageFrench),
			WithCountry(CountryFrance),
			WithGender(GenderMale),
		)

		assert.Equal(t, LanguageFrench, GetDefaultLanguage())
		assert.Equal(t, CountryFrance, GetDefaultCountry())
		assert.Equal(t, GenderMale, GetDefaultGender())

		// Test with custom faker
		customFaker := New(WithLanguage(LanguageSpanish))
		SetDefaultFaker(customFaker)

		// Reset to original
		SetDefaultLanguage(originalLang)
		SetDefaultCountry(originalCountry)
		SetDefaultGender(originalGender)
		SetDefaultSeed(originalSeed)

		// Test reset functionality
		ResetDefaultFaker()
	})
}

// TestPoolFunctionsComprehensive tests pool functionality comprehensively
func TestPoolFunctionsComprehensive(t *testing.T) {
	t.Run("Pool functionality", func(t *testing.T) {
		// Test pool operations if they exist
		faker := GetFaker()
		assert.NotNil(t, faker)

		name := faker.FirstName()
		assert.NotEmpty(t, name)

		PutFaker(faker)

		// Test with pooled faker
		WithPooledFaker(func(f *Faker) {
			result := f.LastName()
			assert.NotEmpty(t, result)
		})

		// Test parallel generation
		results := ParallelGenerate(5, func(f *Faker) string {
			return f.FirstName()
		})
		assert.Len(t, results, 5)
		for _, result := range results {
			assert.NotEmpty(t, result)
		}

		// Test batch generation
		batchResults := BatchGenerate(3, func() string {
			return FirstName()
		})
		assert.Len(t, batchResults, 3)

		// Test concurrent generation
		concurrentResults := ConcurrentGenerate(4, func(f *Faker) string {
			return f.Email()
		})
		assert.Len(t, concurrentResults, 4)

		// Test pool stats
		stats := GetPoolStats()
		assert.NotNil(t, stats)

		// Warmup pools
		WarmupPools()
	})
}

// TestCloneAndUtilities tests clone and utility functions
func TestCloneAndUtilities(t *testing.T) {
	t.Run("Clone and utility functions", func(t *testing.T) {
		original := New(WithLanguage(LanguageEnglish), WithCountry(CountryUS))

		// Test clone
		cloned := original.Clone()
		assert.NotNil(t, cloned)

		// Both should work independently
		originalName := original.FirstName()
		clonedName := cloned.FirstName()

		assert.NotEmpty(t, originalName)
		assert.NotEmpty(t, clonedName)

		// Test stats from both
		originalStats := original.Stats()
		clonedStats := cloned.Stats()

		assert.NotNil(t, originalStats)
		assert.NotNil(t, clonedStats)
	})
}

// TestComprehensiveLanguageSupport tests all supported languages
func TestComprehensiveLanguageSupport(t *testing.T) {
	t.Run("All supported languages", func(t *testing.T) {
		languages := GetSupportedLanguages()
		require.NotEmpty(t, languages)

		for _, lang := range languages {
			t.Run(string(lang), func(t *testing.T) {
				faker := New(WithLanguage(lang))

				// Test basic functions with each language
				firstName := faker.FirstName()
				lastName := faker.LastName()
				email := faker.Email()

				assert.NotEmpty(t, firstName)
				assert.NotEmpty(t, lastName)
				assert.NotEmpty(t, email)
				assert.Contains(t, email, "@")
			})
		}
	})
}

// TestComprehensiveCountrySupport tests all supported countries
func TestComprehensiveCountrySupport(t *testing.T) {
	t.Run("All supported countries", func(t *testing.T) {
		countries := GetSupportedCountries()
		require.NotEmpty(t, countries)

		for _, country := range countries {
			t.Run(string(country), func(t *testing.T) {
				faker := New(WithCountry(country))

				// Test functions that vary by country
				zipCode := faker.ZipCode()
				countryName := faker.CountryName()

				assert.NotEmpty(t, zipCode)
				assert.NotEmpty(t, countryName)
			})
		}
	})
}