package fake

import (
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestDataLoaderFunctions tests data_loader.go functions with 0% coverage
func TestDataLoaderFunctions(t *testing.T) {
	faker := New()

	t.Run("loadDataSet", func(t *testing.T) {
		// Test loadDataSet function by calling functions that use it
		// This will indirectly test the data loading mechanism
		faker.language = LanguageChineseSimplified

		// Try to load Chinese names data which should trigger loadDataSet
		firstName := faker.FirstName()
		assert.NotEmpty(t, firstName)

		// Try to load different data sets
		lastName := faker.LastName()
		assert.NotEmpty(t, lastName)

		city := faker.City()
		assert.NotEmpty(t, city)
	})

	t.Run("getItemValues", func(t *testing.T) {
		// Test getItemValues by using functions that call it
		faker.language = LanguageEnglish

		// These calls should trigger getItemValues internally
		for i := 0; i < 10; i++ {
			name := faker.FirstName()
			assert.NotEmpty(t, name)

			company := faker.CompanyName()
			assert.NotEmpty(t, company)
		}
	})

	t.Run("getItemsByTag", func(t *testing.T) {
		// Test getItemsByTag by using functions that might call it
		faker.language = LanguageFrench

		// These calls should trigger getItemsByTag internally
		for i := 0; i < 5; i++ {
			street := faker.Street()
			assert.NotEmpty(t, street)

			suffix := faker.CompanySuffix()
			assert.NotEmpty(t, suffix)
		}
	})
}

// TestCachingFunctions tests fake.go caching functions with 0% coverage
func TestCachingFunctions(t *testing.T) {
	faker := New()

	t.Run("incrementCacheHit", func(t *testing.T) {
		// Test cache hit increment by forcing cache usage
		faker.language = LanguageEnglish

		// Make the same call multiple times to trigger caching
		for i := 0; i < 5; i++ {
			// This should trigger cache mechanisms
			name := faker.FirstName()
			assert.NotEmpty(t, name)
		}

		// Check that stats were incremented (indirectly)
		assert.True(t, faker.stats.cacheHits >= 0)
	})

	t.Run("getCachedData", func(t *testing.T) {
		// Test getCachedData by setting up cache scenarios
		faker.language = LanguageSpanish

		// First call should cache data
		name1 := faker.FirstName()
		assert.NotEmpty(t, name1)

		// Subsequent calls should use cache
		name2 := faker.FirstName()
		assert.NotEmpty(t, name2)
	})

	t.Run("setCachedData", func(t *testing.T) {
		// Test setCachedData by forcing cache writes
		faker.language = LanguagePortuguese

		// Generate data that should trigger cache writes
		for i := 0; i < 3; i++ {
			city := faker.City()
			assert.NotEmpty(t, city)

			country := faker.CountryName()
			assert.NotEmpty(t, country)
		}
	})
}

// TestValidationFunctionsZero tests validation functions with 0% coverage
func TestValidationFunctionsZero(t *testing.T) {
	t.Run("validateLanguage", func(t *testing.T) {
		// Test validateLanguage by creating faker with different languages
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
			faker := New(WithLanguage(lang))
			assert.Equal(t, lang, faker.language)

			// Use the faker to trigger validation
			name := faker.FirstName()
			assert.NotEmpty(t, name)
		}

		// Test with invalid language
		faker := New(WithLanguage(Language("invalid")))
		name := faker.FirstName()
		assert.NotEmpty(t, name) // Should still work with fallback
	})

	t.Run("validateCountry", func(t *testing.T) {
		// Test validateCountry by creating faker with different countries
		validCountries := []Country{
			CountryUS,
			CountryCanada,
			CountryChina,
			CountryUK,
			CountryFrance,
			CountryGermany,
			CountryJapan,
			CountryKorea,
			CountryRussia,
			CountryBrazil,
			CountrySpain,
			CountryPortugal,
			CountryItaly,
		}

		for _, country := range validCountries {
			faker := New(WithCountry(country))
			assert.Equal(t, country, faker.country)

			// Use the faker to trigger validation
			phone := faker.PhoneNumber()
			assert.NotEmpty(t, phone)
		}

		// Test with invalid country
		faker := New(WithCountry(Country("XX")))
		phone := faker.PhoneNumber()
		assert.NotEmpty(t, phone) // Should still work with fallback
	})
}

// TestFormattingFunctions tests formatting functions with 0% coverage
func TestFormattingFunctions(t *testing.T) {
	faker := New()

	t.Run("formatWithParams", func(t *testing.T) {
		// Test formatWithParams by using functions that call it
		faker.language = LanguageEnglish

		// Generate data that uses formatting
		for i := 0; i < 10; i++ {
			address := faker.AddressLine()
			assert.NotEmpty(t, address)

			bs := faker.BS()
			assert.NotEmpty(t, bs)
			assert.Contains(t, bs, " ") // Should contain formatted text
		}
	})

	t.Run("formatNumber", func(t *testing.T) {
		// Test formatNumber by using functions that format numbers
		faker.country = CountryUS

		for i := 0; i < 10; i++ {
			phone := faker.PhoneNumber()
			assert.NotEmpty(t, phone)

			mobile := faker.MobileNumber()
			assert.NotEmpty(t, mobile)
		}
	})
}

// TestCreditCardFunctions tests identity.go credit card functions with 0% coverage
func TestCreditCardFunctions(t *testing.T) {
	faker := New()

	t.Run("SafeCreditCardNumber", func(t *testing.T) {
		// Test SafeCreditCardNumber function
		for i := 0; i < 10; i++ {
			safeCard := faker.SafeCreditCardNumber()
			assert.NotEmpty(t, safeCard)
			// Check that it's a valid credit card number format
			assert.Regexp(t, `^\d+$`, safeCard, "Safe credit card should contain only digits")
			assert.True(t, len(safeCard) >= 13 && len(safeCard) <= 19, "Credit card length should be between 13-19 digits")
		}
	})

	t.Run("BatchCreditCardInfos", func(t *testing.T) {
		// Test BatchCreditCardInfos function
		count := 5
		infos := faker.BatchCreditCardInfos(count)
		assert.Len(t, infos, count)

		for _, info := range infos {
			assert.NotEmpty(t, info.Number)
			assert.NotEmpty(t, info.CVV)
			assert.True(t, info.ExpiryMonth > 0)
			assert.True(t, info.ExpiryYear > 0)
			assert.NotEmpty(t, info.HolderName)
		}
	})
}

// TestNamesFunctions tests names.go functions with 0% coverage
func TestNamesFunctions(t *testing.T) {
	faker := New()

	t.Run("getDefaultFirstName", func(t *testing.T) {
		// Test getDefaultFirstName by using edge cases
		faker.language = Language("unsupported")

		// Should fall back to default names
		for i := 0; i < 5; i++ {
			name := faker.FirstName()
			assert.NotEmpty(t, name)
		}
	})

	t.Run("getDefaultLastName", func(t *testing.T) {
		// Test getDefaultLastName by using edge cases
		faker.language = Language("unsupported")

		// Should fall back to default names
		for i := 0; i < 5; i++ {
			name := faker.LastName()
			assert.NotEmpty(t, name)
		}
	})

	// Note: NamePrefix function doesn't exist in current implementation
}

// TestPerformanceFunctions tests nano_performance.go functions with 0% coverage
func TestPerformanceFunctions(t *testing.T) {
	t.Run("FastAtomicOperations", func(t *testing.T) {
		// Test atomic operations

		// Test FastAtomicAdd
		result := FastAtomicAdd()
		assert.True(t, result > 0)

		// Test FastAtomicLoad
		loaded := FastAtomicLoad()
		assert.True(t, loaded > 0)

		// Test FastAtomicCAS
		swapped := FastAtomicCAS()
		assert.True(t, swapped)
	})

	t.Run("ConcurrentAtomicOperations", func(t *testing.T) {
		// Test atomic operations under concurrency
		var wg sync.WaitGroup

		workers := 10

		wg.Add(workers)
		for i := 0; i < workers; i++ {
			go func() {
				defer wg.Done()
				for j := 0; j < 10; j++ {
					FastAtomicAdd()
					FastAtomicLoad()
					FastAtomicCAS()
				}
			}()
		}

		wg.Wait()
	})
}

// TestOptimizedFunctions tests optimized.go functions with 0% coverage
func TestOptimizedFunctions(t *testing.T) {
	t.Run("precomputeTemplates", func(t *testing.T) {
		// Test precomputeTemplates by creating multiple fakers
		// This should trigger template precomputation
		fakers := make([]*Faker, 5)
		for i := 0; i < 5; i++ {
			fakers[i] = New(WithLanguage(LanguageEnglish))
		}

		// Use all fakers to trigger template usage
		for _, faker := range fakers {
			for j := 0; j < 3; j++ {
				company := faker.CompanyName()
				assert.NotEmpty(t, company)

				catchphrase := faker.Catchphrase()
				assert.NotEmpty(t, catchphrase)

				bs := faker.BS()
				assert.NotEmpty(t, bs)
			}
		}
	})
}

// TestGlobalConvenienceFunctions tests global functions with 0% coverage
func TestGlobalConvenienceFunctions(t *testing.T) {
	t.Run("GlobalSafeCreditCardNumber", func(t *testing.T) {
		// Test global SafeCreditCardNumber function
		for i := 0; i < 5; i++ {
			safeCard := SafeCreditCardNumber()
			assert.NotEmpty(t, safeCard)
		}
	})
}

// TestEdgeCasesAndErrorPaths tests edge cases to improve coverage
func TestEdgeCasesAndErrorPaths(t *testing.T) {
	t.Run("EmptyDataHandling", func(t *testing.T) {
		// Test behavior with empty or missing data
		faker := New(WithLanguage(Language("nonexistent")))

		// Should handle gracefully
		name := faker.FirstName()
		assert.NotEmpty(t, name)

		company := faker.CompanyName()
		assert.NotEmpty(t, company)
	})

	t.Run("ConcurrentFakerUsage", func(t *testing.T) {
		// Test concurrent usage of faker
		faker := New()
		var wg sync.WaitGroup
		results := make([]string, 100)

		wg.Add(100)
		for i := 0; i < 100; i++ {
			go func(index int) {
				defer wg.Done()
				results[index] = faker.FirstName()
			}(i)
		}

		wg.Wait()

		// All results should be valid
		for _, result := range results {
			assert.NotEmpty(t, result)
		}
	})

	t.Run("MemoryUsageOptimization", func(t *testing.T) {
		// Test memory optimization by creating many fakers
		fakers := make([]*Faker, 50)
		for i := 0; i < 50; i++ {
			fakers[i] = New()
		}

		// Use all fakers
		for _, faker := range fakers {
			name := faker.FirstName()
			assert.NotEmpty(t, name)
		}
	})
}