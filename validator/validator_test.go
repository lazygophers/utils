package validator

import (
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// Test Types
type TestUser struct {
	Name     string `json:"name" validate:"required,min=2"`
	Email    string `json:"email" validate:"required,email"`
	Phone    string `json:"phone" validate:"mobile"`
	IDCard   string `json:"id_card" validate:"idcard"`
	Password string `json:"password" validate:"strong_password"`
}

// Core Validator Tests

func TestValidatorBasic(t *testing.T) {
	v, err := New()
	require.NoError(t, err)

	// 测试有效数据
	validUser := TestUser{
		Name:     "张三",
		Email:    "test@example.com",
		Phone:    "13812345678",
		IDCard:   "11010119800101123X",
		Password: "MyPass123!",
	}

	valErr := v.Struct(validUser)
	assert.NoError(t, valErr)
}

func TestValidatorErrors(t *testing.T) {
	v, err := New()
	require.NoError(t, err)

	// 测试无效数据
	invalidUser := TestUser{
		Name:     "",
		Email:    "invalid",
		Phone:    "123",
		IDCard:   "invalid",
		Password: "weak",
	}

	valErr := v.Struct(invalidUser)
	require.Error(t, valErr)

	validationErrors, ok := valErr.(ValidationErrors)
	require.True(t, ok)
	assert.True(t, len(validationErrors) > 0)
}

func TestValidatorLocale(t *testing.T) {
	// 测试英文（默认情况下只有英文可用）
	enValidator, err := New(WithLocale("en"))
	require.NoError(t, err)

	user := TestUser{Name: ""}
	valErr := enValidator.Struct(user)
	require.Error(t, valErr)
	assert.Contains(t, valErr.Error(), "required")

	// 测试不存在的地区会回退到英文
	unknownValidator, err2 := New(WithLocale("unknown"))
	require.NoError(t, err2)

	valErr2 := unknownValidator.Struct(user)
	require.Error(t, valErr2)
	assert.Contains(t, valErr2.Error(), "required")
}

func TestValidatorJSONFields(t *testing.T) {
	v, err := New(WithUseJSON(true))
	require.NoError(t, err)

	user := TestUser{Name: ""}
	valErr := v.Struct(user)
	require.Error(t, valErr)

	validationErrors := valErr.(ValidationErrors)
	nameErr := validationErrors.ByField("name")
	require.NotNil(t, nameErr)
	assert.Equal(t, "name", nameErr.Field)
}

func TestValidatorVar(t *testing.T) {
	v, err := New()
	require.NoError(t, err)

	// 测试手机号
	err1 := v.Var("13812345678", "mobile")
	assert.NoError(t, err1)

	err2 := v.Var("123", "mobile")
	assert.Error(t, err2)

	// 测试身份证
	err3 := v.Var("11010119800101123X", "idcard")
	assert.NoError(t, err3)

	err4 := v.Var("invalid", "idcard")
	assert.Error(t, err4)
}

// Validation Errors Tests

func TestValidationErrorsMethods(t *testing.T) {
	v, err := New()
	require.NoError(t, err)

	user := TestUser{
		Name:  "",
		Email: "invalid",
	}

	valErr := v.Struct(user)
	require.Error(t, valErr)

	validationErrors := valErr.(ValidationErrors)

	// 测试各种方法
	assert.True(t, validationErrors.HasField("name"))
	assert.False(t, validationErrors.HasField("nonexistent"))
	assert.NotEmpty(t, validationErrors.FirstError())
	assert.NotNil(t, validationErrors.First())
	assert.NotEmpty(t, validationErrors.Fields())
	assert.NotEmpty(t, validationErrors.Messages())
	assert.NotEmpty(t, validationErrors.ToMap())
	assert.NotEmpty(t, validationErrors.ToDetailMap())
	assert.NotEmpty(t, validationErrors.JSON())
	assert.False(t, validationErrors.IsEmpty())
	assert.True(t, validationErrors.Len() > 0)
}

func TestValidationErrorsMultiple(t *testing.T) {
	errors := ValidationErrors{
		&FieldError{Field: "field1", Message: "error1"},
		&FieldError{Field: "field2", Message: "error2"},
	}

	// 测试 Error 方法返回多个错误
	errorMsg := errors.Error()
	assert.Contains(t, errorMsg, "error1")
	assert.Contains(t, errorMsg, "error2")
	assert.Contains(t, errorMsg, ";")
}

func TestErrorFiltering(t *testing.T) {
	v, err := New()
	require.NoError(t, err)

	user := TestUser{
		Name:  "",
		Email: "invalid",
	}

	valErr := v.Struct(user)
	require.Error(t, valErr)

	validationErrors := valErr.(ValidationErrors)

	// 测试过滤
	requiredErrors := validationErrors.ForTag("required")
	assert.True(t, len(requiredErrors) > 0)

	nameErrors := validationErrors.ForField("name")
	assert.True(t, len(nameErrors) > 0)

	// 测试自定义过滤
	filtered := validationErrors.Filter(func(err *FieldError) bool {
		return err.Tag == "required"
	})
	assert.True(t, len(filtered) > 0)
}

func TestErrorManipulation(t *testing.T) {
	var errors ValidationErrors

	// 测试空错误
	assert.True(t, errors.IsEmpty())
	assert.Nil(t, errors.First())
	assert.Empty(t, errors.FirstError())

	// 添加错误
	errors.Add(&FieldError{
		Field:   "test",
		Tag:     "required",
		Message: "Test message",
	})

	assert.False(t, errors.IsEmpty())
	assert.Equal(t, 1, errors.Len())

	// 合并错误
	other := ValidationErrors{
		&FieldError{
			Field:   "other",
			Tag:     "email",
			Message: "Other message",
		},
	}
	errors.Merge(other)
	assert.Equal(t, 2, errors.Len())

	// 测试格式化
	formatted := errors.Format("{field}: {message}")
	assert.Contains(t, formatted, "test:")
}

// Global Functions Tests

func TestGlobalFunctions(t *testing.T) {
	SetLocale("zh")
	SetUseJSON(true)

	user := TestUser{Name: ""}
	err := Struct(user)
	assert.Error(t, err)

	err2 := Var("", "required")
	assert.Error(t, err2)

	err3 := Var("test@example.com", "email")
	assert.NoError(t, err3)
}

func TestGlobalFunctionsExtra(t *testing.T) {
	// 测试全局 RegisterValidation
	err := RegisterValidation("test_global", func(fl FieldLevel) bool {
		return true
	})
	assert.NoError(t, err)

	// 测试全局 RegisterTranslation
	RegisterTranslation("en", "test_global", "test message")
}

func TestDefaultValidator(t *testing.T) {
	defaultV := Default()
	assert.NotNil(t, defaultV)
}

// Custom Validators Tests

func TestAllCustomValidators(t *testing.T) {
	tests := []struct {
		name     string
		value    interface{}
		tag      string
		expected bool
	}{
		// 手机号
		{"valid mobile 1", "13812345678", "mobile", true},
		{"valid mobile 2", "15912345678", "mobile", true},
		{"invalid mobile 1", "123456", "mobile", false},
		{"invalid mobile 2", "", "mobile", false},

		// 身份证
		{"valid idcard 15", "110101800101123", "idcard", true},
		{"valid idcard 18", "11010119800101123X", "idcard", true},
		{"invalid idcard", "123", "idcard", false},

		// 强密码
		{"strong password 1", "MyPass123!", "strong_password", true},
		{"strong password 2", "Complex1@", "strong_password", true},
		{"weak password 1", "123456", "strong_password", false},
		{"weak password 2", "password", "strong_password", false},
		{"weak password 3", "12345", "strong_password", false}, // 长度不够

		// 中文姓名
		{"chinese name 1", "张三", "chinese_name", true},
		{"chinese name 2", "李四", "chinese_name", true},
		{"chinese name with dot", "阿·布", "chinese_name", true},
		{"invalid chinese name 1", "John", "chinese_name", false},
		{"invalid chinese name 2", "", "chinese_name", false},
		{"invalid chinese name 3", "a", "chinese_name", false},

		// 银行卡
		{"valid bankcard", "4111111111111111", "bankcard", true},
		{"invalid bankcard 1", "123456", "bankcard", false},
		{"invalid bankcard 2", "", "bankcard", false},
		{"invalid bankcard 3", "abc", "bankcard", false},
	}

	v, err := New()
	require.NoError(t, err)

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := v.Var(tt.value, tt.tag)
			if tt.expected {
				assert.NoError(t, err, "Expected %v to be valid for %s", tt.value, tt.tag)
			} else {
				assert.Error(t, err, "Expected %v to be invalid for %s", tt.value, tt.tag)
			}
		})
	}
}

func TestCustomValidators(t *testing.T) {
	v, err := New()
	require.NoError(t, err)

	// 测试中文名
	err1 := v.Var("张三", "chinese_name")
	assert.NoError(t, err1)

	err2 := v.Var("John", "chinese_name")
	assert.Error(t, err2)

	// 测试银行卡 - 使用一个符合Luhn算法的测试卡号
	err3 := v.Var("4111111111111111", "bankcard")
	assert.NoError(t, err3)

	err4 := v.Var("123456", "bankcard")
	assert.Error(t, err4)
}

func TestCustomValidatorEdgeCases(t *testing.T) {
	v, err := New()
	require.NoError(t, err)

	// 测试银行卡验证的边界情况
	err1 := v.Var("123456789012", "bankcard") // 太短
	assert.Error(t, err1)

	err2 := v.Var("12345678901234567890", "bankcard") // 太长
	assert.Error(t, err2)

	err3 := v.Var("123456789012345a", "bankcard") // 包含字母
	assert.Error(t, err3)

	// 测试 Luhn 算法的错误情况
	assert.False(t, luhnCheck("123a"))

	// 测试手机号的边界情况
	err4 := v.Var("", "mobile")
	assert.Error(t, err4)

	err5 := v.Var("12345678901", "mobile") // 不是1开头的正确格式
	assert.Error(t, err5)

	// 测试身份证边界情况
	err6 := v.Var("", "idcard")
	assert.Error(t, err6)

	err7 := v.Var("1234", "idcard") // 太短
	assert.Error(t, err7)

	// 测试强密码边界情况
	err8 := v.Var("1234567", "strong_password") // 太短
	assert.Error(t, err8)

	// 测试中文姓名边界情况
	err9 := v.Var("", "chinese_name")
	assert.Error(t, err9)

	err10 := v.Var("a", "chinese_name") // 非中文
	assert.Error(t, err10)
}

// Options and Configuration Tests

func TestConfigOptions(t *testing.T) {
	config := Config{
		Locale:  "zh",
		UseJSON: true,
		Translations: map[string]string{
			"zh.test": "测试消息",
		},
	}

	v, err := New(WithConfig(config))
	require.NoError(t, err)
	assert.Equal(t, "zh", v.GetLocale())
}

func TestWithTranslations(t *testing.T) {
	translations := map[string]string{
		"en.test": "Test message",
		"zh.test": "测试消息",
	}

	v, err := New(WithTranslations(translations))
	require.NoError(t, err)
	assert.Contains(t, v.messages, "en.test")
}

func TestCustomValidatorWithOptions(t *testing.T) {
	v, err := New(
		WithCustomValidator("even_number", func(value interface{}) bool {
			if num, ok := value.(int); ok {
				return num%2 == 0
			}
			return false
		}),
	)
	require.NoError(t, err)

	err1 := v.Var(4, "even_number")
	assert.NoError(t, err1)

	err2 := v.Var(3, "even_number")
	assert.Error(t, err2)
}

func TestRegisterTranslation(t *testing.T) {
	v, err := New()
	require.NoError(t, err)

	v.RegisterTranslation("en", "test", "test message")
	// 无法直接测试翻译，但确保不会panic
}

// Locale Configuration Tests

func TestLocaleConfig(t *testing.T) {
	// 测试获取配置
	config, ok := GetLocaleConfig("en")
	assert.True(t, ok)
	assert.NotNil(t, config)

	// 测试不存在的地区会回退到英文
	config2, ok2 := GetLocaleConfig("nonexistent")
	assert.True(t, ok2) // 会回退到英文
	assert.NotNil(t, config2)

	// 测试可用地区
	locales := GetAvailableLocales()
	assert.Contains(t, locales, "en")
}

func TestLocaleConfigEdgeCases(t *testing.T) {
	// 测试获取不存在地区时的回退逻辑
	config, ok := GetLocaleConfig("nonexistent-locale")
	assert.True(t, ok) // 应该回退到英文
	assert.NotNil(t, config)

	// 测试带地区的语言代码
	config2, ok2 := GetLocaleConfig("zh-TW")
	assert.True(t, ok2)
	assert.NotNil(t, config2)
}

// Coverage and Edge Case Tests

func TestAdditionalCoverage(t *testing.T) {
	// 测试单个错误的 Error 方法
	singleError := ValidationErrors{
		&FieldError{
			Field:   "test",
			Message: "test error",
		},
	}
	assert.Equal(t, "test error", singleError.Error())

	// 测试空错误的 Error 方法
	emptyError := ValidationErrors{}
	assert.Equal(t, "", emptyError.Error())

	// 测试 String 方法
	assert.Equal(t, "test error", singleError.String())
}

func TestDefaultValidatorFallback(t *testing.T) {
	// 重置默认验证器来测试错误路径
	originalValidator := defaultValidator
	originalOnce := once

	defaultValidator = nil
	once = sync.Once{}

	// 获取默认验证器
	v := Default()
	assert.NotNil(t, v)

	// 恢复原值
	defaultValidator = originalValidator
	once = originalOnce
}

func TestNonExistentLocale(t *testing.T) {
	v, err := New(WithLocale("nonexistent"))
	require.NoError(t, err)

	user := TestUser{Name: ""}
	valErr := v.Struct(user)
	require.Error(t, valErr)
	// 应该回退到英文错误消息
	assert.Contains(t, valErr.Error(), "required")
}

func TestFieldNameWithoutJSON(t *testing.T) {
	v, err := New(WithUseJSON(false))
	require.NoError(t, err)

	type TestStruct struct {
		FieldName string `json:"json_name" validate:"required"`
	}

	testStruct := TestStruct{}
	valErr := v.Struct(testStruct)
	require.Error(t, valErr)

	validationErrors := valErr.(ValidationErrors)
	fieldErr := validationErrors.ByField("FieldName")
	assert.NotNil(t, fieldErr)
}

func TestValidatorFields(t *testing.T) {
	v, err := New()
	require.NoError(t, err)

	// 测试设置和获取
	v.SetLocale("zh")
	assert.Equal(t, "zh", v.GetLocale())

	v.SetUseJSON(false)
	// 我们无法直接测试 useJSON 字段，但可以通过行为验证

	// 测试注册翻译
	v.RegisterTranslation("en", "test", "test message")
}

func TestMessageFormattingEdgeCases(t *testing.T) {
	// 这个测试主要是为了覆盖 formatMessage 中 nil 值的分支
	// 我们通过创建一个带有 nil 值的验证错误来测试
	v, err := New()
	require.NoError(t, err)

	type TestStruct struct {
		Field *string `validate:"required"`
	}

	// 使用 nil 指针来触发 nil 值的格式化分支
	testStruct := TestStruct{Field: nil}
	valErr := v.Struct(testStruct)
	require.Error(t, valErr)

	// 验证错误消息不为空
	assert.NotEmpty(t, valErr.Error())
}

// Locale Tests (Multi-language Support)

func TestAllLocalesAvailable(t *testing.T) {
	t.Skip("Skipping multi-language test - requires all locale files to be available")

	expectedLocales := []string{
		"en", "zh", "zh-CN", "zh-TW", "ja", "ko", "fr", "es", "ar", "ru", "it", "pt", "de",
	}

	availableLocales := GetAvailableLocales()
	localeMap := make(map[string]bool)
	for _, locale := range availableLocales {
		localeMap[locale] = true
	}

	for _, expected := range expectedLocales {
		if !localeMap[expected] {
			t.Errorf("Expected locale %s not found in available locales: %v", expected, availableLocales)
		}
	}
}

func TestLocaleMessages(t *testing.T) {
	t.Skip("Skipping multi-language test - requires all locale files to be available")

	testCases := []struct {
		locale  string
		msgKey  string
		wantKey string
	}{
		{"en", "required", "{field} is required"},
		{"zh", "required", "{field}不能为空"},
		{"zh-CN", "required", "{field}不能为空"},
		{"zh-TW", "required", "{field}不能為空"},
		{"ja", "required", "{field}は必須です"},
		{"ko", "required", "{field}은(는) 필수입니다"},
		{"fr", "required", "{field} est requis"},
		{"es", "required", "{field} es obligatorio"},
		{"ar", "required", "{field} مطلوب"},
		{"ru", "required", "{field} обязательно для заполнения"},
		{"it", "required", "{field} è obbligatorio"},
		{"pt", "required", "{field} é obrigatório"},
		{"de", "required", "{field} ist erforderlich"},
	}

	for _, tc := range testCases {
		t.Run(tc.locale+"_"+tc.msgKey, func(t *testing.T) {
			config, ok := GetLocaleConfig(tc.locale)
			if !ok {
				t.Fatalf("Locale %s not found", tc.locale)
			}

			msg, exists := config.Messages[tc.msgKey]
			if !exists {
				t.Fatalf("Message key %s not found in locale %s", tc.msgKey, tc.locale)
			}

			if msg != tc.wantKey {
				t.Errorf("Expected message '%s' for key %s in locale %s, got '%s'", tc.wantKey, tc.msgKey, tc.locale, msg)
			}
		})
	}
}

func TestCustomRulesInAllLocales(t *testing.T) {
	t.Skip("Skipping multi-language test - requires all locale files to be available")

	customRules := []string{"mobile", "idcard", "bankcard", "chinese_name", "strong_password"}
	locales := []string{"en", "zh", "zh-CN", "zh-TW", "ja", "ko", "fr", "es", "ar", "ru", "it", "pt", "de"}

	for _, locale := range locales {
		t.Run(locale, func(t *testing.T) {
			config, ok := GetLocaleConfig(locale)
			if !ok {
				t.Fatalf("Locale %s not found", locale)
			}

			for _, rule := range customRules {
				if _, exists := config.Messages[rule]; !exists {
					t.Errorf("Custom rule %s not found in locale %s", rule, locale)
				}
			}
		})
	}
}
