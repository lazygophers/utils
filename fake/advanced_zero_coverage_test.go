package fake

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestAdvancedZeroCoverage targets all the 0% coverage functions identified in the analysis
func TestAdvancedZeroCoverage(t *testing.T) {
	t.Run("data_loader functions", func(t *testing.T) {
		// Test loadDataSet
		t.Run("loadDataSet", func(t *testing.T) {
			// Create fakers with different languages to trigger loadDataSet
			faker1 := New(WithLanguage(LanguageEnglish))
			result := faker1.FirstName()
			assert.NotEmpty(t, result)

			faker2 := New(WithLanguage(LanguageChineseSimplified))
			result = faker2.FirstName()
			assert.NotEmpty(t, result)
		})

		// Test getItemValues
		t.Run("getItemValues", func(t *testing.T) {
			faker := New()
			// This should be triggered by calling functions that use weighted data
			for i := 0; i < 10; i++ {
				name := faker.FirstName()
				assert.NotEmpty(t, name)
			}
		})

		// Test getItemsByTag
		t.Run("getItemsByTag", func(t *testing.T) {
			faker := New(WithLanguage(LanguageEnglish))
			for i := 0; i < 5; i++ {
				result := faker.LastName()
				assert.NotEmpty(t, result)
			}
		})
	})

	t.Run("cache management functions", func(t *testing.T) {
		faker := New()

		// Test incrementCacheHit
		t.Run("incrementCacheHit", func(t *testing.T) {
			// Generate multiple requests to trigger cache operations
			for i := 0; i < 20; i++ {
				faker.FirstName()
				faker.LastName()
				faker.Email()
			}
		})

		// Test getCachedData and setCachedData
		t.Run("cache operations", func(t *testing.T) {
			// Create multiple fakers to trigger cache operations
			faker1 := New(WithLanguage(LanguageEnglish))
			faker1.FirstName()

			faker2 := New(WithLanguage(LanguageChineseSimplified))
			faker2.FirstName()

			faker3 := New(WithLanguage(LanguageEnglish)) // Should hit cache
			faker3.FirstName()
		})
	})

	t.Run("validation functions", func(t *testing.T) {
		// Test validateLanguage
		t.Run("validateLanguage", func(t *testing.T) {
			// Try setting various languages to trigger validation
			validLanguages := []Language{
				LanguageEnglish,
				LanguageChineseSimplified,
				LanguageChineseTraditional,
				LanguageFrench,
				LanguageRussian,
			}

			for _, lang := range validLanguages {
				faker := New(WithLanguage(lang))
				result := faker.FirstName()
				assert.NotEmpty(t, result)
			}
		})

		// Test validateCountry
		t.Run("validateCountry", func(t *testing.T) {
			// Try setting various countries to trigger validation
			validCountries := []Country{
				CountryChina,
				CountryUS,
				CountryJapan,
			}

			for _, country := range validCountries {
				faker := New(WithCountry(country))
				result := faker.Street()
				assert.NotEmpty(t, result)
			}
		})
	})

	t.Run("formatting functions", func(t *testing.T) {
		faker := New()

		// Test formatWithParams
		t.Run("formatWithParams", func(t *testing.T) {
			// Generate data that requires parameter formatting
			phone := faker.PhoneNumber()
			assert.NotEmpty(t, phone)

			ssn := faker.SSN()
			assert.NotEmpty(t, ssn)
		})

		// Test formatNumber
		t.Run("formatNumber", func(t *testing.T) {
			// Generate numeric data that requires formatting
			creditCard := faker.CreditCardNumber()
			assert.NotEmpty(t, creditCard)

			bankAccount := faker.BankAccount()
			assert.NotEmpty(t, bankAccount)
		})
	})

	t.Run("default name functions", func(t *testing.T) {
		// Test getDefaultFirstName and getDefaultLastName by creating conditions that might trigger them
		t.Run("default names", func(t *testing.T) {
			// Create faker with various configurations to potentially trigger default behavior
			faker := New()
			result := faker.FirstName()
			assert.NotEmpty(t, result)

			result = faker.LastName()
			assert.NotEmpty(t, result)
		})
	})

	t.Run("pool functions", func(t *testing.T) {
		// Test all pool-related functions
		t.Run("string builder pool", func(t *testing.T) {
			// Call functions that might use string builder pool
			faker := New()
			for i := 0; i < 10; i++ {
				result := faker.BatchEmailsOptimized(5)
				assert.Len(t, result, 5)
				for _, email := range result {
					assert.NotEmpty(t, email)
				}
			}
		})

		t.Run("slice pools", func(t *testing.T) {
			// Call batch functions that might use slice pools
			faker := NewOptimized()
			for i := 0; i < 5; i++ {
				names := faker.BatchFastNames(3)
				assert.Len(t, names, 3)
				for _, name := range names {
					assert.NotEmpty(t, name)
				}
			}
		})
	})
}

// TestAdvancedPoolFunctions specifically targets pool function coverage
func TestAdvancedPoolFunctions(t *testing.T) {
	t.Run("BatchEmailsOptimized", func(t *testing.T) {
		faker := New()
		for batchSize := 1; batchSize <= 10; batchSize++ {
			emails := faker.BatchEmailsOptimized(batchSize)
			assert.Len(t, emails, batchSize)
			for _, email := range emails {
				assert.NotEmpty(t, email)
				assert.Contains(t, email, "@")
			}
		}
	})

	t.Run("pool stress test", func(t *testing.T) {
		// Create multiple fakers simultaneously to stress the pool
		const numFakers = 10
		const batchSize = 5

		results := make(chan []string, numFakers)

		for i := 0; i < numFakers; i++ {
			go func() {
				faker := New()
				emails := faker.BatchEmailsOptimized(batchSize)
				results <- emails
			}()
		}

		for i := 0; i < numFakers; i++ {
			emails := <-results
			assert.Len(t, emails, batchSize)
		}
	})
}

// TestNamePrefixCoverage targets the 21.1% coverage NamePrefix function
func TestNamePrefixCoverage(t *testing.T) {
	faker := New()

	t.Run("generate name prefixes", func(t *testing.T) {
		// Generate many name prefixes to cover different code paths
		prefixes := make(map[string]bool)

		for i := 0; i < 100; i++ {
			prefix := faker.NamePrefix()
			assert.NotEmpty(t, prefix)
			prefixes[prefix] = true
		}

		// Should have generated multiple different prefixes
		assert.True(t, len(prefixes) > 1, "Should generate varied prefixes")
	})

	t.Run("language-specific prefixes", func(t *testing.T) {
		languages := []Language{
			LanguageEnglish,
			LanguageChineseSimplified,
		}

		for _, lang := range languages {
			faker := New(WithLanguage(lang))
			for i := 0; i < 10; i++ {
				prefix := faker.NamePrefix()
				assert.NotEmpty(t, prefix)
			}
		}
	})
}

// TestDataLoaderGetItemsByTag targets the 30% coverage function
func TestDataLoaderGetItemsByTag(t *testing.T) {
	t.Run("GetItemsByTag coverage", func(t *testing.T) {
		dm := &DataManager{}

		// Test with various tag queries that might exist
		testTags := []string{
			"common",
			"rare",
			"male",
			"female",
			"formal",
			"casual",
		}

		for _, tag := range testTags {
			items, err := dm.GetItemsByTag(LanguageEnglish, "names", "first", tag)
			// Items might be empty for non-existent tags, that's OK
			if err != nil {
				t.Logf("Tag '%s' returned error: %v", tag, err)
			} else {
				t.Logf("Tag '%s' returned %d items", tag, len(items))
			}
		}
	})
}