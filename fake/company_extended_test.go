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
