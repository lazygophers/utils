package fake

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

// 测试各种国家的电话号码生成
func TestPhoneNumber(t *testing.T) {
	// 测试美国电话号码
	f := New(WithCountry(CountryUS))
	phone := f.PhoneNumber()
	assert.NotEmpty(t, phone)

	// 测试加拿大电话号码
	f = New(WithCountry(CountryCanada))
	phone = f.PhoneNumber()
	assert.NotEmpty(t, phone)

	// 测试中国电话号码
	f = New(WithCountry(CountryChina))
	phone = f.PhoneNumber()
	assert.NotEmpty(t, phone)

	// 测试英国电话号码
	f = New(WithCountry(CountryUK))
	phone = f.PhoneNumber()
	assert.NotEmpty(t, phone)

	// 测试法国电话号码
	f = New(WithCountry(CountryFrance))
	phone = f.PhoneNumber()
	assert.NotEmpty(t, phone)

	// 测试德国电话号码
	f = New(WithCountry(CountryGermany))
	phone = f.PhoneNumber()
	assert.NotEmpty(t, phone)

	// 测试日本电话号码
	f = New(WithCountry(CountryJapan))
	phone = f.PhoneNumber()
	assert.NotEmpty(t, phone)

	// 测试韩国电话号码
	f = New(WithCountry(CountryKorea))
	phone = f.PhoneNumber()
	assert.NotEmpty(t, phone)

	// 测试默认情况
	f = New()
	phone = f.PhoneNumber()
	assert.NotEmpty(t, phone)
}

// 测试手机号码生成
func TestMobileNumber(t *testing.T) {
	f := New()
	mobile := f.MobileNumber()
	assert.NotEmpty(t, mobile)

	// 测试中国手机号码
	f = New(WithCountry(CountryChina))
	mobile = f.MobileNumber()
	assert.NotEmpty(t, mobile)

	// 测试美国手机号码
	f = New(WithCountry(CountryUS))
	mobile = f.MobileNumber()
	assert.NotEmpty(t, mobile)
}

// 测试邮箱地址生成
func TestEmail(t *testing.T) {
	f := New()
	email := f.Email()
	assert.NotEmpty(t, email)
	assert.Contains(t, email, "@")
}

// 测试企业邮箱地址生成
func TestCompanyEmail(t *testing.T) {
	f := New()
	email := f.CompanyEmail()
	assert.NotEmpty(t, email)
	assert.Contains(t, email, "@")
}

// 测试安全邮箱地址生成
func TestSafeEmail(t *testing.T) {
	f := New()
	email := f.SafeEmail()
	assert.NotEmpty(t, email)
	assert.Contains(t, email, "@")
	// 安全邮箱应该使用example.com等域名
	assert.True(t, strings.Contains(email, "example.") || strings.Contains(email, "test.") || strings.Contains(email, "sample."))
}

// 测试URL生成
func TestURL(t *testing.T) {
	f := New()
	url := f.URL()
	assert.NotEmpty(t, url)
	assert.Contains(t, url, "http")
}

// 测试IPv4地址生成
func TestIPv4(t *testing.T) {
	f := New()
	ipv4 := f.IPv4()
	assert.NotEmpty(t, ipv4)
}

// 测试IPv6地址生成
func TestIPv6(t *testing.T) {
	f := New()
	ipv6 := f.IPv6()
	assert.NotEmpty(t, ipv6)
}

// 测试MAC地址生成
func TestMAC(t *testing.T) {
	f := New()
	mac := f.MAC()
	assert.NotEmpty(t, mac)
}

// 测试批量生成电话号码
func TestBatchPhoneNumbers(t *testing.T) {
	f := New()
	phones := f.BatchPhoneNumbers(5)
	assert.Len(t, phones, 5)
	for _, phone := range phones {
		assert.NotEmpty(t, phone)
	}
}

// 测试批量生成邮箱地址
func TestBatchEmails(t *testing.T) {
	f := New()
	emails := f.BatchEmails(5)
	assert.Len(t, emails, 5)
	for _, email := range emails {
		assert.NotEmpty(t, email)
		assert.Contains(t, email, "@")
	}
}

// 测试批量生成URL
func TestBatchURLs(t *testing.T) {
	f := New()
	urls := f.BatchURLs(5)
	assert.Len(t, urls, 5)
	for _, url := range urls {
		assert.NotEmpty(t, url)
		assert.Contains(t, url, "http")
	}
}

// 测试全局函数
func TestGlobalContactFunctions(t *testing.T) {
	// 测试全局PhoneNumber函数
	phone := PhoneNumber()
	assert.NotEmpty(t, phone)

	// 测试全局MobileNumber函数
	mobile := MobileNumber()
	assert.NotEmpty(t, mobile)

	// 测试全局Email函数
	email := Email()
	assert.NotEmpty(t, email)
	assert.Contains(t, email, "@")

	// 测试全局CompanyEmail函数
	companyEmail := CompanyEmail()
	assert.NotEmpty(t, companyEmail)
	assert.Contains(t, companyEmail, "@")

	// 测试全局SafeEmail函数
	safeEmail := SafeEmail()
	assert.NotEmpty(t, safeEmail)
	assert.Contains(t, safeEmail, "@")

	// 测试全局URL函数
	url := URL()
	assert.NotEmpty(t, url)
	assert.Contains(t, url, "http")

	// 测试全局IPv4函数
	ipv4 := IPv4()
	assert.NotEmpty(t, ipv4)

	// 测试全局IPv6函数
	ipv6 := IPv6()
	assert.NotEmpty(t, ipv6)

	// 测试全局MAC函数
	mac := MAC()
	assert.NotEmpty(t, mac)
}

// 测试toASCII函数
func TestToASCII(t *testing.T) {
	f := New()

	// 测试ASCII字符串（应该保持不变）
	asciiStr := "test"
	result := f.toASCII(asciiStr)
	assert.Equal(t, asciiStr, result)

	// 测试包含中文的字符串
	chineseStr := "李小明"
	result = f.toASCII(chineseStr)
	assert.NotEmpty(t, result)

	// 测试空字符串情况
	emptyStr := ""
	result = f.toASCII(emptyStr)
	assert.NotEmpty(t, result)
}
