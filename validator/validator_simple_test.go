package validator

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type TestUser struct {
	Name     string `json:"name" validate:"required,min=2"`
	Email    string `json:"email" validate:"required,email"`
	Phone    string `json:"phone" validate:"mobile"`
	IDCard   string `json:"id_card" validate:"idcard"`
	Password string `json:"password" validate:"strong_password"`
}

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

func TestDefaultValidator(t *testing.T) {
	defaultV := Default()
	assert.NotNil(t, defaultV)
}

func TestRegisterTranslation(t *testing.T) {
	v, err := New()
	require.NoError(t, err)

	v.RegisterTranslation("en", "test", "test message")
	// 无法直接测试翻译，但确保不会panic
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

// 确保所有自定义验证函数都被测试
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
