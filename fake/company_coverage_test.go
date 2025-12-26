package fake

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// 测试CompanySuffix函数
func TestCompanySuffix(t *testing.T) {
	// 测试不同语言的情况
	languages := []Language{
		LanguageEnglish,
		LanguageChineseSimplified,
		LanguageChineseTraditional,
		LanguageFrench,
		LanguageRussian,
		LanguageSpanish,
		LanguagePortuguese,
	}

	for _, lang := range languages {
		t.Run("lang_"+string(lang), func(t *testing.T) {
			f := New(WithLanguage(lang))
			suffix := f.CompanySuffix()
			assert.NotEmpty(t, suffix)
		})
	}

	// 测试全局便捷函数
	t.Run("global_CompanySuffix", func(t *testing.T) {
		suffix := CompanySuffix()
		assert.NotEmpty(t, suffix)
	})
}

// 测试Industry函数
func TestIndustry(t *testing.T) {
	// 测试不同语言的情况
	languages := []Language{
		LanguageEnglish,
		LanguageChineseSimplified,
		LanguageChineseTraditional,
		LanguageFrench,
	}

	for _, lang := range languages {
		t.Run("lang_"+string(lang), func(t *testing.T) {
			f := New(WithLanguage(lang))
			industry := f.Industry()
			assert.NotEmpty(t, industry)
		})
	}

	// 测试全局便捷函数
	t.Run("global_Industry", func(t *testing.T) {
		industry := Industry()
		assert.NotEmpty(t, industry)
	})
}

// 测试JobTitle函数
func TestJobTitle(t *testing.T) {
	// 测试不同语言的情况
	languages := []Language{
		LanguageEnglish,
		LanguageChineseSimplified,
		LanguageChineseTraditional,
	}

	for _, lang := range languages {
		t.Run("lang_"+string(lang), func(t *testing.T) {
			f := New(WithLanguage(lang))
			// 多次调用，覆盖不同的分支
			for i := 0; i < 20; i++ {
				jobTitle := f.JobTitle()
				assert.NotEmpty(t, jobTitle)
			}
		})
	}

	// 测试全局便捷函数
	t.Run("global_JobTitle", func(t *testing.T) {
		jobTitle := JobTitle()
		assert.NotEmpty(t, jobTitle)
	})
}

// 测试Department函数
func TestDepartment(t *testing.T) {
	// 测试不同语言的情况
	languages := []Language{
		LanguageEnglish,
		LanguageChineseSimplified,
		LanguageChineseTraditional,
		LanguageFrench,
	}

	for _, lang := range languages {
		t.Run("lang_"+string(lang), func(t *testing.T) {
			f := New(WithLanguage(lang))
			department := f.Department()
			assert.NotEmpty(t, department)
		})
	}

	// 测试全局便捷函数
	t.Run("global_Department", func(t *testing.T) {
		department := Department()
		assert.NotEmpty(t, department)
	})
}

// 测试CompanyInfo函数
func TestCompanyInfo(t *testing.T) {
	// 测试不同语言和国家的情况
	tests := []struct {
		language Language
		country  Country
	}{
		{LanguageEnglish, CountryUS},
		{LanguageChineseSimplified, CountryChina},
		{LanguageChineseTraditional, CountryChina},
		{LanguageFrench, CountryFrance},
	}

	for _, tt := range tests {
		t.Run("lang_"+string(tt.language)+"_country_"+string(tt.country), func(t *testing.T) {
			f := New(WithLanguage(tt.language), WithCountry(tt.country))
			companyInfo := f.CompanyInfo()
			assert.NotEmpty(t, companyInfo.Name)
			assert.NotEmpty(t, companyInfo.Industry)
			assert.NotEmpty(t, companyInfo.Description)
			assert.NotEmpty(t, companyInfo.Website)
			assert.NotEmpty(t, companyInfo.Email)
			assert.NotEmpty(t, companyInfo.Phone)
			assert.NotEmpty(t, companyInfo.Address)
			assert.True(t, companyInfo.Founded >= 1900 && companyInfo.Founded <= 2024)
			assert.True(t, companyInfo.Employees >= 1 && companyInfo.Employees <= 100000)
		})
	}

	// 测试全局便捷函数
	t.Run("global_CompanyInfo", func(t *testing.T) {
		companyInfo := CompanyInfo()
		assert.NotEmpty(t, companyInfo.Name)
		assert.NotEmpty(t, companyInfo.Industry)
	})

	// 测试批量生成
	t.Run("BatchCompanyInfos", func(t *testing.T) {
		f := New()
		companyInfos := f.BatchCompanyInfos(5)
		assert.Len(t, companyInfos, 5)
		for _, info := range companyInfos {
			assert.NotEmpty(t, info.Name)
		}
	})
}

// 测试BS函数
func TestBS(t *testing.T) {
	// 测试不同语言的情况
	languages := []Language{
		LanguageEnglish,
		LanguageChineseSimplified,
		LanguageChineseTraditional,
	}

	for _, lang := range languages {
		t.Run("lang_"+string(lang), func(t *testing.T) {
			f := New(WithLanguage(lang))
			bs := f.BS()
			assert.NotEmpty(t, bs)
		})
	}

	// 测试全局便捷函数
	t.Run("global_BS", func(t *testing.T) {
		bs := BS()
		assert.NotEmpty(t, bs)
	})
}

// 测试Catchphrase函数
func TestCatchphrase(t *testing.T) {
	// 测试不同语言的情况
	languages := []Language{
		LanguageEnglish,
		LanguageChineseSimplified,
		LanguageChineseTraditional,
		LanguageFrench,
	}

	for _, lang := range languages {
		t.Run("lang_"+string(lang), func(t *testing.T) {
			f := New(WithLanguage(lang))
			catchphrase := f.Catchphrase()
			assert.NotEmpty(t, catchphrase)
		})
	}

	// 测试全局便捷函数
	t.Run("global_Catchphrase", func(t *testing.T) {
		catchphrase := Catchphrase()
		assert.NotEmpty(t, catchphrase)
	})
}

// 测试JobTitle函数的各种情况
func TestJobTitleVariations(t *testing.T) {
	f := New()
	// 多次调用，覆盖不同的分支情况
	for i := 0; i < 50; i++ {
		jobTitle := f.JobTitle()
		assert.NotEmpty(t, jobTitle)
	}

	// 测试批量生成
	jobTitles := f.BatchJobTitles(10)
	assert.Len(t, jobTitles, 10)
	for _, title := range jobTitles {
		assert.NotEmpty(t, title)
	}
}

// 测试批量生成CompanyNames
func TestBatchCompanyNames(t *testing.T) {
	f := New()
	companyNames := f.BatchCompanyNames(10)
	assert.Len(t, companyNames, 10)
	for _, name := range companyNames {
		assert.NotEmpty(t, name)
	}
}
