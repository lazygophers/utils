package fake

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestAddressFunctions tests address-related functions for better coverage
func TestAddressFunctions(t *testing.T) {
	// Test global convenience functions
	t.Run("Global convenience functions", func(t *testing.T) {
		street := Street()
		assert.NotEmpty(t, street)

		state := State()
		assert.NotEmpty(t, state)

		zip := ZipCode()
		assert.NotEmpty(t, zip)

		country := CountryName()
		assert.NotEmpty(t, country)

		addressLine := AddressLine()
		assert.NotEmpty(t, addressLine)

		fullAddr := FullAddress()
		assert.NotEmpty(t, fullAddr)
	})

	// Test faker instance methods with different languages
	t.Run("Faker with different languages", func(t *testing.T) {
		// Test with English
		fakerEN := New(WithLanguage(LanguageEnglish))
		streetEN := fakerEN.Street()
		assert.NotEmpty(t, streetEN)

		stateEN := fakerEN.State()
		assert.NotEmpty(t, stateEN)

		zipEN := fakerEN.ZipCode()
		assert.NotEmpty(t, zipEN)

		// Test with Chinese
		fakerZH := New(WithLanguage(LanguageChineseSimplified))
		streetZH := fakerZH.Street()
		assert.NotEmpty(t, streetZH)

		// Test with French
		fakerFR := New(WithLanguage(LanguageFrench))
		streetFR := fakerFR.Street()
		assert.NotEmpty(t, streetFR)
	})

	// Test edge cases and error handling
	t.Run("Edge cases", func(t *testing.T) {
		// Test multiple calls for consistency
		streets := make(map[string]bool)
		for i := 0; i < 10; i++ {
			street := Street()
			assert.NotEmpty(t, street)
			streets[street] = true
		}
		// Should have some variety
		assert.GreaterOrEqual(t, len(streets), 2)
	})
}

// TestCompanyFunctions tests company-related functions
func TestCompanyFunctions(t *testing.T) {
	t.Run("Global company functions", func(t *testing.T) {
		suffix := CompanySuffix()
		assert.NotEmpty(t, suffix)

		bs := BS()
		assert.NotEmpty(t, bs)

		catchphrase := Catchphrase()
		assert.NotEmpty(t, catchphrase)

		companyName := CompanyName()
		assert.NotEmpty(t, companyName)

		jobTitle := JobTitle()
		assert.NotEmpty(t, jobTitle)
	})

	t.Run("Company with different languages", func(t *testing.T) {
		// Test with English
		fakerEN := New(WithLanguage(LanguageEnglish))
		suffixEN := fakerEN.CompanySuffix()
		assert.NotEmpty(t, suffixEN)

		bsEN := fakerEN.BS()
		assert.NotEmpty(t, bsEN)

		// Test with Chinese
		fakerZH := New(WithLanguage(LanguageChineseSimplified))
		companyZH := fakerZH.CompanyName()
		assert.NotEmpty(t, companyZH)

		jobTitleZH := fakerZH.JobTitle()
		assert.NotEmpty(t, jobTitleZH)
	})

	t.Run("Company name generation consistency", func(t *testing.T) {
		companies := make(map[string]bool)
		for i := 0; i < 20; i++ {
			company := CompanyName()
			assert.NotEmpty(t, company)
			companies[company] = true
		}
		// Should have reasonable variety
		assert.GreaterOrEqual(t, len(companies), 3)
	})
}

// TestContactFunctions tests contact-related functions
func TestContactFunctions(t *testing.T) {
	t.Run("Global contact functions", func(t *testing.T) {
		phone := PhoneNumber()
		assert.NotEmpty(t, phone)

		mobile := MobileNumber()
		assert.NotEmpty(t, mobile)

		email := Email()
		assert.NotEmpty(t, email)
		assert.Contains(t, email, "@")
	})

	t.Run("Contact with different languages", func(t *testing.T) {
		// Test with English
		fakerEN := New(WithLanguage(LanguageEnglish))
		phoneEN := fakerEN.PhoneNumber()
		assert.NotEmpty(t, phoneEN)

		// Test with Chinese
		fakerZH := New(WithLanguage(LanguageChineseSimplified))
		phoneZH := fakerZH.PhoneNumber()
		assert.NotEmpty(t, phoneZH)

		mobileZH := fakerZH.MobileNumber()
		assert.NotEmpty(t, mobileZH)
	})

	t.Run("Email generation patterns", func(t *testing.T) {
		emails := make(map[string]bool)
		for i := 0; i < 10; i++ {
			email := Email()
			assert.Contains(t, email, "@")
			assert.Contains(t, email, ".")
			emails[email] = true
		}
		// Should generate different emails
		assert.GreaterOrEqual(t, len(emails), 5)
	})
}

// TestDataFunctions tests data loading functions
func TestDataFunctions(t *testing.T) {
	t.Run("GetSupportedLanguages", func(t *testing.T) {
		languages := GetSupportedLanguages()
		assert.NotEmpty(t, languages)
		assert.Contains(t, languages, LanguageEnglish)
	})

	t.Run("GetSupportedCountries", func(t *testing.T) {
		countries := GetSupportedCountries()
		assert.NotEmpty(t, countries)
		assert.Contains(t, countries, CountryUS)
	})
}

// TestDataLoaderEdgeCases tests edge cases in data loading
func TestDataLoaderEdgeCases(t *testing.T) {
	t.Run("Default settings functions", func(t *testing.T) {
		// Test default language functions
		originalLang := GetDefaultLanguage()
		assert.NotEmpty(t, originalLang)

		SetDefaultLanguage(LanguageChineseSimplified)
		assert.Equal(t, LanguageChineseSimplified, GetDefaultLanguage())

		// Reset to original
		SetDefaultLanguage(originalLang)

		// Test default country functions
		originalCountry := GetDefaultCountry()
		assert.NotEmpty(t, originalCountry)

		SetDefaultCountry(CountryChina)
		assert.Equal(t, CountryChina, GetDefaultCountry())

		// Reset to original
		SetDefaultCountry(originalCountry)

		// Test default gender functions
		originalGender := GetDefaultGender()
		SetDefaultGender(GenderFemale)
		assert.Equal(t, GenderFemale, GetDefaultGender())

		// Reset to original
		SetDefaultGender(originalGender)
	})

	t.Run("Stats and cache functions", func(t *testing.T) {
		// Clear cache and get stats
		ClearDefaultCache()
		stats := GetDefaultStats()
		assert.NotNil(t, stats)

		// Generate some data to populate stats
		_ = FirstName()
		_ = LastName()

		newStats := GetDefaultStats()
		assert.NotNil(t, newStats)
	})
}

// TestPersonFunctions tests person-related functions for coverage
func TestPersonFunctions(t *testing.T) {
	t.Run("Global person functions", func(t *testing.T) {
		firstName := FirstName()
		assert.NotEmpty(t, firstName)

		lastName := LastName()
		assert.NotEmpty(t, lastName)

		fullName := FullName()
		assert.NotEmpty(t, fullName)
	})

	t.Run("Person with different languages", func(t *testing.T) {
		// Test with English
		fakerEN := New(WithLanguage(LanguageEnglish))
		firstNameEN := fakerEN.FirstName()
		assert.NotEmpty(t, firstNameEN)

		// Test with Chinese
		fakerZH := New(WithLanguage(LanguageChineseSimplified))
		firstNameZH := fakerZH.FirstName()
		assert.NotEmpty(t, firstNameZH)
	})
}

// TestInternetFunctions tests internet-related functions
func TestInternetFunctions(t *testing.T) {
	t.Run("Internet data generation", func(t *testing.T) {
		// Test URL generation
		url := URL()
		assert.NotEmpty(t, url)
		assert.Contains(t, url, "http")

		// Test IPv4 generation
		ipv4 := IPv4()
		assert.NotEmpty(t, ipv4)

		// Test IPv6 generation
		ipv6 := IPv6()
		assert.NotEmpty(t, ipv6)

		// Test MAC address generation
		mac := MAC()
		assert.NotEmpty(t, mac)
	})
}

// TestTextFunctions tests text generation functions
func TestTextFunctions(t *testing.T) {
	t.Run("Text generation", func(t *testing.T) {
		// Test word generation
		word := Word()
		assert.NotEmpty(t, word)

		// Test sentence generation
		sentence := Sentence()
		assert.NotEmpty(t, sentence)

		// Test paragraph generation
		paragraph := Paragraph()
		assert.NotEmpty(t, paragraph)

		// Test title generation
		title := Title()
		assert.NotEmpty(t, title)

		// Test lorem generation
		lorem := Lorem()
		assert.NotEmpty(t, lorem)
	})
}

// TestIdentityFunctions tests identity-related functions
func TestIdentityFunctions(t *testing.T) {
	t.Run("Identity document generation", func(t *testing.T) {
		// Test SSN generation
		ssn := SSN()
		assert.NotEmpty(t, ssn)

		// Test Chinese ID generation
		chineseID := ChineseIDNumber()
		assert.NotEmpty(t, chineseID)

		// Test passport generation
		passport := Passport()
		assert.NotEmpty(t, passport)

		// Test drivers license generation
		license := DriversLicense()
		assert.NotEmpty(t, license)

		// Test credit card generation
		ccNumber := CreditCardNumber()
		assert.NotEmpty(t, ccNumber)

		// Test CVV generation
		cvv := CVV()
		assert.NotEmpty(t, cvv)

		// Test bank account generation
		bankAccount := BankAccount()
		assert.NotEmpty(t, bankAccount)

		// Test IBAN generation
		iban := IBAN()
		assert.NotEmpty(t, iban)
	})

	t.Run("Identity structures", func(t *testing.T) {
		// Test identity document structure
		identityDoc := IdentityDoc()
		assert.NotNil(t, identityDoc)

		// Test credit card info structure
		ccInfo := CreditCardInfo()
		assert.NotNil(t, ccInfo)
	})
}

// TestNameVariations tests name generation variations
func TestNameVariations(t *testing.T) {
	t.Run("Name variations", func(t *testing.T) {
		// Test Name function (alias for FullName)
		name := Name()
		assert.NotEmpty(t, name)

		// Test formatted name
		formattedName := FormattedName()
		assert.NotEmpty(t, formattedName)

		// Test name suffix
		suffix := NameSuffix()
		if suffix != "" {
			assert.NotEmpty(t, suffix)
		}
	})
}

// TestAdvancedContactFunctions tests additional contact functions
func TestAdvancedContactFunctions(t *testing.T) {
	t.Run("Advanced contact functions", func(t *testing.T) {
		// Test company email
		companyEmail := CompanyEmail()
		assert.NotEmpty(t, companyEmail)
		assert.Contains(t, companyEmail, "@")

		// Test safe email
		safeEmail := SafeEmail()
		assert.NotEmpty(t, safeEmail)
		assert.Contains(t, safeEmail, "@")
	})
}

// TestDataLoaderInternalFunctions tests internal data loader functions
func TestDataLoaderInternalFunctions(t *testing.T) {
	t.Run("Internal data loader functions", func(t *testing.T) {
		// Test data manager functions through public APIs
		faker := New(WithLanguage(LanguageEnglish))

		// Force loading of data which will test internal functions
		street := faker.Street()
		assert.NotEmpty(t, street)

		// Test cache functionality
		stats := faker.Stats()
		assert.NotNil(t, stats)

		faker.ClearCache()
		stats2 := faker.Stats()
		assert.NotNil(t, stats2)
	})
}

// TestValidationFunctions tests validation functions
func TestValidationFunctions(t *testing.T) {
	t.Run("Language validation", func(t *testing.T) {
		// Test with supported language
		faker2 := New(WithLanguage(LanguageChineseSimplified))
		name := faker2.FirstName()
		assert.NotEmpty(t, name)

		// Test with unsupported language (should fall back)
		faker3 := New(WithLanguage("invalid"))
		name2 := faker3.FirstName()
		assert.NotEmpty(t, name2)
	})

	t.Run("Country validation", func(t *testing.T) {
		// Test with supported country
		faker := New(WithCountry(CountryChina))
		zipCode := faker.ZipCode()
		assert.NotEmpty(t, zipCode)

		// Test with unsupported country
		faker2 := New(WithCountry("XX"))
		zipCode2 := faker2.ZipCode()
		assert.NotEmpty(t, zipCode2)
	})
}

// TestFormatFunctions tests format helper functions
func TestFormatFunctions(t *testing.T) {
	t.Run("Format functions", func(t *testing.T) {
		// Test by generating data that uses format functions internally
		faker := New()

		// Phone numbers use formatting
		phone := faker.PhoneNumber()
		assert.NotEmpty(t, phone)

		// Zip codes use number formatting
		zipCode := faker.ZipCode()
		assert.NotEmpty(t, zipCode)
	})
}

// TestUtilityFunctions tests utility functions
func TestUtilityFunctions(t *testing.T) {
	t.Run("Utility functions", func(t *testing.T) {
		faker := New()

		// Test clone functionality
		cloned := faker.Clone()
		assert.NotNil(t, cloned)

		name1 := faker.FirstName()
		name2 := cloned.FirstName()
		assert.NotEmpty(t, name1)
		assert.NotEmpty(t, name2)
	})
}