package fake

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestCompanyFunctionsExtended 测试公司相关函数的扩展功能
func TestCompanyFunctionsExtended(t *testing.T) {
	faker := New()

	// 测试CompanyName函数
	t.Run("test_company_name", func(t *testing.T) {
		name := faker.CompanyName()
		assert.NotEmpty(t, name)
	})

	// 测试generateCompanyName函数，覆盖所有分支
	t.Run("test_generate_company_name", func(t *testing.T) {
		// 调用多次，确保覆盖所有模式
		for i := 0; i < 20; i++ {
			name := faker.generateCompanyName()
			assert.NotEmpty(t, name)
		}
	})

	// 测试BS函数，覆盖所有分支
	t.Run("test_bs", func(t *testing.T) {
		// 调用多次，确保覆盖所有模式
		for i := 0; i < 20; i++ {
			bs := faker.BS()
			assert.NotEmpty(t, bs)
		}
	})

	// 测试Catchphrase函数，覆盖所有分支
	t.Run("test_catchphrase", func(t *testing.T) {
		// 调用多次，确保覆盖所有模式
		for i := 0; i < 20; i++ {
			catchphrase := faker.Catchphrase()
			assert.NotEmpty(t, catchphrase)
		}
	})

	// 测试CompanyInfo函数
	t.Run("test_company_info", func(t *testing.T) {
		company := faker.CompanyInfo()
		assert.NotNil(t, company)
		assert.NotEmpty(t, company.Name)
		assert.NotEmpty(t, company.Industry)
		assert.NotEmpty(t, company.Description)
		assert.NotEmpty(t, company.Website)
		assert.NotEmpty(t, company.Email)
		assert.NotEmpty(t, company.Phone)
		assert.NotEmpty(t, company.Address)
		assert.Greater(t, company.Founded, 0)
		assert.Greater(t, company.Employees, 0)
	})

	// 测试公司后缀
	t.Run("test_company_suffix", func(t *testing.T) {
		suffix := faker.CompanySuffix()
		assert.NotEmpty(t, suffix)
	})

	// 测试Industry函数
	t.Run("test_industry", func(t *testing.T) {
		industry := faker.Industry()
		assert.NotEmpty(t, industry)
	})

	// 测试JobTitle函数
	t.Run("test_job_title", func(t *testing.T) {
		jobTitle := faker.JobTitle()
		assert.NotEmpty(t, jobTitle)
	})

	// 测试Department函数
	t.Run("test_department", func(t *testing.T) {
		department := faker.Department()
		assert.NotEmpty(t, department)
	})

	// 测试批量生成函数
	t.Run("test_batch_company_names", func(t *testing.T) {
		names := faker.BatchCompanyNames(5)
		assert.Len(t, names, 5)
		for _, name := range names {
			assert.NotEmpty(t, name)
		}
	})

	t.Run("test_batch_job_titles", func(t *testing.T) {
		jobTitles := faker.BatchJobTitles(5)
		assert.Len(t, jobTitles, 5)
		for _, title := range jobTitles {
			assert.NotEmpty(t, title)
		}
	})

	t.Run("test_batch_company_infos", func(t *testing.T) {
		companies := faker.BatchCompanyInfos(5)
		assert.Len(t, companies, 5)
		for _, company := range companies {
			assert.NotNil(t, company)
			assert.NotEmpty(t, company.Name)
		}
	})
}

// TestGlobalCompanyFunctions 测试全局公司相关函数
func TestGlobalCompanyFunctions(t *testing.T) {
	// 测试CompanyName全局函数
	t.Run("test_global_company_name", func(t *testing.T) {
		name := CompanyName()
		assert.NotEmpty(t, name)
	})

	// 测试CompanySuffix全局函数
	t.Run("test_global_company_suffix", func(t *testing.T) {
		suffix := CompanySuffix()
		assert.NotEmpty(t, suffix)
	})

	// 测试Industry全局函数
	t.Run("test_global_industry", func(t *testing.T) {
		industry := Industry()
		assert.NotEmpty(t, industry)
	})

	// 测试JobTitle全局函数
	t.Run("test_global_job_title", func(t *testing.T) {
		jobTitle := JobTitle()
		assert.NotEmpty(t, jobTitle)
	})

	// 测试Department全局函数
	t.Run("test_global_department", func(t *testing.T) {
		department := Department()
		assert.NotEmpty(t, department)
	})

	// 测试CompanyInfo全局函数
	t.Run("test_global_company_info", func(t *testing.T) {
		company := CompanyInfo()
		assert.NotNil(t, company)
		assert.NotEmpty(t, company.Name)
	})

	// 测试BS全局函数
	t.Run("test_global_bs", func(t *testing.T) {
		bs := BS()
		assert.NotEmpty(t, bs)
	})

	// 测试Catchphrase全局函数
	t.Run("test_global_catchphrase", func(t *testing.T) {
		catchphrase := Catchphrase()
		assert.NotEmpty(t, catchphrase)
	})
}

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
	// langCountryCase 描述语言与国家组合。
	type langCountryCase struct {
		language Language
		country  Country
	}
	// 测试不同语言和国家的情况
	tests := []langCountryCase{
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
