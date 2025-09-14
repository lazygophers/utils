package fake

import (
	"strings"
	"testing"
)

// TestBasicFunctionality 测试基本功能
func TestBasicFunctionality(t *testing.T) {
	faker := New()
	
	// 测试名字生成
	name := faker.Name()
	if name == "" {
		t.Error("Name should not be empty")
	}
	
	// 测试邮箱生成
	email := faker.Email()
	if email == "" || !strings.Contains(email, "@") {
		t.Errorf("Invalid email: %s", email)
	}
	
	// 测试电话生成
	phone := faker.PhoneNumber()
	if phone == "" {
		t.Error("Phone should not be empty")
	}
	
	// 测试地址生成
	address := faker.FullAddress()
	if address == nil || address.FullAddress == "" {
		t.Error("Address should not be empty")
	}
}

// TestMultiLanguage 测试多语言支持
func TestMultiLanguage(t *testing.T) {
	// 英语
	enFaker := New(WithLanguage(LanguageEnglish))
	enName := enFaker.Name()
	if enName == "" {
		t.Error("English name should not be empty")
	}
	
	// 中文
	cnFaker := New(WithLanguage(LanguageChineseSimplified))
	cnName := cnFaker.Name()
	if cnName == "" {
		t.Error("Chinese name should not be empty")
	}
	
	t.Logf("EN Name: %s", enName)
	t.Logf("CN Name: %s", cnName)
}

// TestGlobalFunctions 测试全局函数
func TestGlobalFunctions(t *testing.T) {
	name := Name()
	if name == "" {
		t.Error("Global Name() should not be empty")
	}
	
	email := Email()
	if email == "" || !strings.Contains(email, "@") {
		t.Errorf("Global Email() invalid: %s", email)
	}
	
	userAgent := RandomUserAgent()
	if userAgent == "" {
		t.Error("Global RandomUserAgent() should not be empty")
	}
}