package fake

import (
	"fmt"
	"testing"
)

// Example_basicUsage 基本使用示例
func Example_basicUsage() {
	// 基本使用
	name := Name()
	email := Email()
	phone := PhoneNumber()
	
	// 验证输出不为空
	if name == "" || email == "" || phone == "" {
		fmt.Println("Error: Empty output")
	} else {
		fmt.Println("Basic usage example works correctly")
	}
	
	// Output:
	// Basic usage example works correctly
}

// Example_customFaker 自定义 Faker 示例
func Example_customFaker() {
	// 创建中文 Faker
	faker := New(
		WithLanguage(LanguageChineseSimplified),
		WithCountry(CountryChina),
		WithGender(GenderMale),
	)
	
	name := faker.Name()
	phone := faker.PhoneNumber()
	address := faker.FullAddress()
	
	// 验证输出不为空且符合预期格式
	if name == "" || phone == "" || address == nil || address.FullAddress == "" {
		fmt.Println("Error: Empty output")
	} else {
		fmt.Println("Custom faker example works correctly")
	}
	
	// Output:
	// Custom faker example works correctly
}

// TestCompleteExample 完整功能演示测试
func TestCompleteExample(t *testing.T) {
	faker := New()
	
	t.Log("=== 个人信息 ===")
	t.Logf("姓名: %s", faker.Name())
	t.Logf("邮箱: %s", faker.Email())
	t.Logf("电话: %s", faker.PhoneNumber())
	
	t.Log("\n=== 地址信息 ===")
	address := faker.FullAddress()
	t.Logf("完整地址: %s", address.FullAddress)
	t.Logf("街道: %s", address.Street)
	t.Logf("城市: %s", address.City)
	t.Logf("邮编: %s", address.ZipCode)
	
	t.Log("\n=== 公司信息 ===")
	company := faker.CompanyInfo()
	t.Logf("公司名称: %s", company.Name)
	t.Logf("行业: %s", company.Industry)
	t.Logf("网站: %s", company.Website)
	t.Logf("员工数量: %d", company.Employees)
	
	t.Log("\n=== 金融信息 ===")
	creditCard := faker.CreditCardInfo()
	t.Logf("信用卡: %s (%s)", creditCard.Number, creditCard.Brand)
	t.Logf("CVV: %s", creditCard.CVV)
	t.Logf("过期日期: %d/%d", creditCard.ExpiryMonth, creditCard.ExpiryYear)
	
	t.Log("\n=== 设备信息 ===")
	device := faker.DeviceInfo()
	t.Logf("设备: %s %s", device.Manufacturer, device.Model)
	t.Logf("操作系统: %s %s", device.OS, device.OSVersion)
	t.Logf("浏览器: %s %s", device.Browser, device.Version)
	t.Logf("屏幕分辨率: %dx%d", device.ScreenWidth, device.ScreenHeight)
	
	t.Log("\n=== 身份证件 ===")
	identity := faker.IdentityDoc()
	t.Logf("证件类型: %s", identity.Type)
	t.Logf("证件号码: %s", identity.Number)
	t.Logf("发证日期: %s", identity.IssuedDate)
	t.Logf("过期日期: %s", identity.ExpiryDate)
	
	t.Log("\n=== 文本内容 ===")
	t.Logf("标题: %s", faker.Title())
	t.Logf("句子: %s", faker.Sentence())
	t.Logf("段落: %s", faker.Text(100))
	t.Logf("标签: %v", faker.HashTags(3))
	
	t.Log("\n=== 网络信息 ===")
	t.Logf("网址: %s", faker.URL())
	t.Logf("IP地址: %s", faker.IPv4())
	t.Logf("MAC地址: %s", faker.MAC())
	t.Logf("用户代理: %s", faker.UserAgent())
}

// TestMultiLanguageExample 多语言演示测试
func TestMultiLanguageExample(t *testing.T) {
	languages := []struct{
		Lang Language
		Name string
	}{
		{LanguageEnglish, "英语"},
		{LanguageChineseSimplified, "简体中文"},
		{LanguageChineseTraditional, "繁体中文"},
	}
	
	for _, lang := range languages {
		t.Logf("\n=== %s 示例 ===", lang.Name)
		faker := New(WithLanguage(lang.Lang))
		
		t.Logf("姓名: %s", faker.Name())
		t.Logf("公司: %s", faker.CompanyName())
		t.Logf("地址: %s", faker.AddressLine())
		t.Logf("国家: %s", faker.CountryName())
	}
}

// TestPerformanceExample 性能演示测试
func TestPerformanceExample(t *testing.T) {
	faker := New()
	
	t.Log("=== 单个生成测试 ===")
	for i := 0; i < 5; i++ {
		t.Logf("用户 %d: %s - %s", i+1, faker.Name(), faker.Email())
	}
	
	t.Log("\n=== 批量生成测试 ===")
	names := faker.BatchNames(10)
	t.Logf("批量生成10个姓名: %v", names)
	
	emails := ParallelGenerate(5, func(f *Faker) string {
		return f.Email()
	})
	t.Logf("并行生成5个邮箱: %v", emails)
	
	t.Log("\n=== 性能统计 ===")
	stats := faker.Stats()
	for key, value := range stats {
		t.Logf("%s: %d", key, value)
	}
}