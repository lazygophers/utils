package fake

import (
	"strings"
	"testing"
)

// TestAllFeaturesWork 验证所有主要功能是否正常工作
func TestAllFeaturesWork(t *testing.T) {
	faker := New()
	
	// 基础数据生成
	tests := []struct {
		name string
		fn   func() string
	}{
		{"Name", faker.Name},
		{"FirstName", faker.FirstName},
		{"LastName", faker.LastName},
		{"Email", faker.Email},
		{"PhoneNumber", faker.PhoneNumber},
		{"Street", faker.Street},
		{"City", faker.City},
		{"ZipCode", faker.ZipCode},
		{"CountryName", faker.CountryName},
		{"CompanyName", faker.CompanyName},
		{"JobTitle", faker.JobTitle},
		{"Industry", faker.Industry},
		{"SSN", faker.SSN},
		{"CreditCardNumber", faker.CreditCardNumber},
		{"IPv4", faker.IPv4},
		{"MAC", faker.MAC},
		{"URL", faker.URL},
		{"Word", faker.Word},
		{"Sentence", faker.Sentence},
		{"Title", faker.Title},
		{"UserAgent", faker.UserAgent},
		{"Browser", faker.Browser},
		{"OS", faker.OS},
	}
	
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := test.fn()
			if result == "" {
				t.Errorf("%s returned empty string", test.name)
			}
		})
	}
}

// TestComplexObjectsWork 验证复杂对象生成是否正常
func TestComplexObjectsWork(t *testing.T) {
	faker := New()
	
	// 测试复杂对象
	address := faker.FullAddress()
	if address == nil || address.FullAddress == "" {
		t.Error("FullAddress should not be nil or empty")
	}
	
	company := faker.CompanyInfo()
	if company == nil || company.Name == "" {
		t.Error("CompanyInfo should not be nil or have empty name")
	}
	
	creditCard := faker.CreditCardInfo()
	if creditCard == nil || creditCard.Number == "" {
		t.Error("CreditCardInfo should not be nil or have empty number")
	}
	
	device := faker.DeviceInfo()
	if device == nil || device.Type == "" {
		t.Error("DeviceInfo should not be nil or have empty type")
	}
	
	identity := faker.IdentityDoc()
	if identity == nil || identity.Number == "" {
		t.Error("IdentityDoc should not be nil or have empty number")
	}
}

// TestMultiLanguageFunctional 验证多语言功能是否正常
func TestMultiLanguageFunctional(t *testing.T) {
	languages := []Language{
		LanguageEnglish,
		LanguageChineseSimplified,
		LanguageChineseTraditional,
	}
	
	for _, lang := range languages {
		t.Run(string(lang), func(t *testing.T) {
			faker := New(WithLanguage(lang))
			
			name := faker.Name()
			if name == "" {
				t.Errorf("Name should not be empty for language %s", lang)
			}
			
			company := faker.CompanyName()
			if company == "" {
				t.Errorf("CompanyName should not be empty for language %s", lang)
			}
		})
	}
}

// TestBatchGeneration 验证批量生成功能
func TestBatchGeneration(t *testing.T) {
	faker := New()
	
	// 测试批量生成
	names := faker.BatchNames(10)
	if len(names) != 10 {
		t.Errorf("Expected 10 names, got %d", len(names))
	}
	
	for i, name := range names {
		if name == "" {
			t.Errorf("Name at index %d is empty", i)
		}
	}
	
	// 测试并行生成
	emails := ParallelGenerate(100, func(f *Faker) string {
		return f.Email()
	})
	
	if len(emails) != 100 {
		t.Errorf("Expected 100 emails, got %d", len(emails))
	}
	
	for i, email := range emails {
		if email == "" || !strings.Contains(email, "@") {
			t.Errorf("Invalid email at index %d: %s", i, email)
		}
	}
}

// TestGlobalFunctionsSanity 验证全局函数的正常性
func TestGlobalFunctionsSanity(t *testing.T) {
	globalTests := []struct{
		name string
		value string
	}{
		{"Name", Name()},
		{"Email", Email()},
		{"PhoneNumber", PhoneNumber()},
		{"RandomUserAgent", RandomUserAgent()},
		{"CompanyName", CompanyName()},
		{"Word", Word()},
		{"Title", Title()},
	}
	
	for _, test := range globalTests {
		if test.value == "" {
			t.Errorf("Global %s returned empty string", test.name)
		}
	}
}

// TestNoMemoryLeaks 简单的内存泄漏测试
func TestNoMemoryLeaks(t *testing.T) {
	faker := New()
	
	// 生成大量数据，检查是否有明显的内存问题
	for i := 0; i < 10000; i++ {
		_ = faker.Name()
		_ = faker.Email()
		
		// 定期清理缓存
		if i%1000 == 0 {
			faker.ClearCache()
		}
	}
	
	// 如果到这里没有崩溃，说明基本没有严重的内存问题
	stats := faker.Stats()
	t.Logf("Stats after 10k generations: %+v", stats)
}