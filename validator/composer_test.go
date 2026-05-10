package validator

import (
	"testing"
)

// TestAndComposer 测试 And 组合验证器
func TestAndComposer(t *testing.T) {
	v, _ := New()

	// 注册组合验证器：密码长度 8-20 且包含特殊字符
	v.RegisterValidation("strong_pwd", And(
		Length(8, 20),
		ContainsSpecial(),
	))

	type User struct {
		Password string `validate:"strong_pwd"`
	}

	// 测试有效密码
	valid := User{Password: "abc123!@"}
	err := v.Struct(valid)
	if err != nil {
		t.Errorf("Expected valid, got error: %v", err)
	}

	// 测试太短
	tooShort := User{Password: "ab1!"}
	err = v.Struct(tooShort)
	if err == nil {
		t.Error("Expected error for too short password")
	}

	// 测试缺少特殊字符
	noSpecial := User{Password: "abc12345"}
	err = v.Struct(noSpecial)
	if err == nil {
		t.Error("Expected error for password without special char")
	}
}

// TestOrComposer 测试 Or 组合验证器
func TestOrComposer(t *testing.T) {
	v, _ := New()

	// 注册组合验证器：可以是手机号或邮箱
	v.RegisterValidation("phone_or_email", Or(
		Pattern(`^1[3-9]\d{9}$`),
		Email(),
	))

	type Contact struct {
		Contact string `validate:"phone_or_email"`
	}

	// 测试手机号
	phone := Contact{Contact: "13812345678"}
	err := v.Struct(phone)
	if err != nil {
		t.Errorf("Expected valid phone, got error: %v", err)
	}

	// 测试邮箱
	email := Contact{Contact: "test@example.com"}
	err = v.Struct(email)
	if err != nil {
		t.Errorf("Expected valid email, got error: %v", err)
	}

	// 测试无效
	invalid := Contact{Contact: "invalid"}
	err = v.Struct(invalid)
	if err == nil {
		t.Error("Expected error for invalid contact")
	}
}

// TestNotComposer 测试 Not 组合验证器
func TestNotComposer(t *testing.T) {
	v, _ := New()

	// 注册组合验证器：不能是 admin
	v.RegisterValidation("not_admin", Not(
		In("admin", "root", "system"),
	))

	type User struct {
		Username string `validate:"not_admin"`
	}

	// 测试有效用户
	valid := User{Username: "user123"}
	err := v.Struct(valid)
	if err != nil {
		t.Errorf("Expected valid, got error: %v", err)
	}

	// 测试无效用户名
	invalid := User{Username: "admin"}
	err = v.Struct(invalid)
	if err == nil {
		t.Error("Expected error for admin username")
	}
}

// TestRangeComposer 测试范围验证器
func TestRangeComposer(t *testing.T) {
	v, _ := New()

	type Product struct {
		Price float64 `validate:"range=0,10000"`
	}

	// 注册范围验证器
	v.RegisterValidation("range", Range(0, 10000))

	// 测试有效价格
	valid := Product{Price: 99.99}
	err := v.Struct(valid)
	if err != nil {
		t.Errorf("Expected valid, got error: %v", err)
	}

	// 测试超出范围
	tooHigh := Product{Price: 15000}
	err = v.Struct(tooHigh)
	if err == nil {
		t.Error("Expected error for price too high")
	}
}

// TestLengthComposer 测试长度验证器
func TestLengthComposer(t *testing.T) {
	v, _ := New()

	type User struct {
		Name string `validate:"length=2,10"`
	}

	// 注册长度验证器
	v.RegisterValidation("length", Length(2, 10))

	// 测试有效长度
	valid := User{Name: "Alice"}
	err := v.Struct(valid)
	if err != nil {
		t.Errorf("Expected valid, got error: %v", err)
	}

	// 测试太短
	tooShort := User{Name: "A"}
	err = v.Struct(tooShort)
	if err == nil {
		t.Error("Expected error for name too short")
	}

	// 测试太长
	tooLong := User{Name: "ThisNameIsWayTooLong"}
	err = v.Struct(tooLong)
	if err == nil {
		t.Error("Expected error for name too long")
	}
}

// TestChainedComposition 测试链式组合
func TestChainedComposition(t *testing.T) {
	v, _ := New()

	// 复杂组合：用户名必须 3-20 字符，字母开头，只含字母数字下划线
	v.RegisterValidation("username", And(
		Length(3, 20),
		Pattern(`^[a-zA-Z][a-zA-Z0-9_]*$`),
		Not(In("admin", "root", "system")),
	))

	type User struct {
		Username string `validate:"username"`
	}

	// 测试有效用户名
	validCases := []string{"user123", "Alice", "bob_the_builder"}
	for _, username := range validCases {
		user := User{Username: username}
		err := v.Struct(user)
		if err != nil {
			t.Errorf("Expected valid for %s, got error: %v", username, err)
		}
	}

	// 测试无效用户名
	invalidCases := []struct {
		username string
		reason    string
	}{
		{"ab", "too short"},
		{"1invalid", "must start with letter"},
		{"admin", "reserved name"},
		{"user@name", "invalid chars"},
	}

	for _, tc := range invalidCases {
		user := User{Username: tc.username}
		err := v.Struct(user)
		if err == nil {
			t.Errorf("Expected error for %s (%s)", tc.username, tc.reason)
		}
	}
}

// TestRequiredWithMinLength 测试必填+最小长度组合
func TestRequiredWithMinLength(t *testing.T) {
	v, _ := New()

	type Form struct {
		Password string `validate:"required_min"`
	}

	// 注册组合验证器
	v.RegisterValidation("required_min", And(
		Required(),
		MinLength(8),
	))

	// 测试有效密码
	valid := Form{Password: "password123"}
	err := v.Struct(valid)
	if err != nil {
		t.Errorf("Expected valid, got error: %v", err)
	}

	// 测试空密码
	empty := Form{Password: ""}
	err = v.Struct(empty)
	if err == nil {
		t.Error("Expected error for empty password")
	}

	// 测试太短
	tooShort := Form{Password: "short"}
	err = v.Struct(tooShort)
	if err == nil {
		t.Error("Expected error for short password")
	}
}

// TestInNotInComposer 测试 In/NotIn 验证器
func TestInNotInComposer(t *testing.T) {
	v, _ := New()

	type Product struct {
		Category string `validate:"valid_category"`
		Size     string `validate:"not_reserved_size"`
	}

	// 注册验证器
	v.RegisterValidation("valid_category", In("electronics", "clothing", "food"))
	v.RegisterValidation("not_reserved_size", NotIn("xs", "xl"))

	// 测试有效产品
	valid := Product{Category: "electronics", Size: "m"}
	err := v.Struct(valid)
	if err != nil {
		t.Errorf("Expected valid, got error: %v", err)
	}

	// 测试无效分类
	invalidCat := Product{Category: "invalid", Size: "m"}
	err = v.Struct(invalidCat)
	if err == nil {
		t.Error("Expected error for invalid category")
	}

	// 测试保留尺寸
	reserved := Product{Category: "electronics", Size: "xs"}
	err = v.Struct(reserved)
	if err == nil {
		t.Error("Expected error for reserved size")
	}
}

// TestEmailURLComposer 测试 Email/URL 构造函数
func TestEmailURLComposer(t *testing.T) {
	v, _ := New()

	type Contact struct {
		EmailAddr string `validate:"custom_email"`
		Website   string `validate:"custom_url"`
	}

	// 注册验证器
	v.RegisterValidation("custom_email", Email())
	v.RegisterValidation("custom_url", URL())

	// 测试有效数据
	valid := Contact{EmailAddr: "test@example.com", Website: "https://example.com"}
	err := v.Struct(valid)
	if err != nil {
		t.Errorf("Expected valid, got error: %v", err)
	}

	// 测试无效邮箱
	invalidEmail := Contact{EmailAddr: "invalid", Website: "https://example.com"}
	err = v.Struct(invalidEmail)
	if err == nil {
		t.Error("Expected error for invalid email")
	}

	// 测试无效 URL
	invalidURL := Contact{EmailAddr: "test@example.com", Website: "invalid"}
	err = v.Struct(invalidURL)
	if err == nil {
		t.Error("Expected error for invalid URL")
	}
}

// TestPatternComposer 测试正则表达式构造函数
func TestPatternComposer(t *testing.T) {
	v, _ := New()

	type User struct {
		ZipCode string `validate:"zip_code"`
	}

	// 注册验证器
	v.RegisterValidation("zip_code", Pattern(`^\d{5}$`))

	// 测试有效邮编
	valid := User{ZipCode: "12345"}
	err := v.Struct(valid)
	if err != nil {
		t.Errorf("Expected valid, got error: %v", err)
	}

	// 测试无效邮编
	invalid := User{ZipCode: "abc"}
	err = v.Struct(invalid)
	if err == nil {
		t.Error("Expected error for invalid zip code")
	}
}
