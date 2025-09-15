package validator

import (
	"sync"
	"testing"

	"github.com/go-playground/validator/v10"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// TestAdditionalCoverage 测试额外的覆盖率
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

// TestDefaultValidatorFallback 测试默认验证器回退
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

// TestNonExistentLocale 测试不存在的地区
func TestNonExistentLocale(t *testing.T) {
	v, err := New(WithLocale("nonexistent"))
	require.NoError(t, err)

	user := TestUser{Name: ""}
	valErr := v.Struct(user)
	require.Error(t, valErr)
	// 应该回退到英文错误消息
	assert.Contains(t, valErr.Error(), "required")
}

// TestFieldNameWithoutJSON 测试不使用 JSON 字段名
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

// TestGlobalFunctionsExtra 测试全局函数覆盖率
func TestGlobalFunctionsExtra(t *testing.T) {
	// 测试全局 RegisterValidation
	err := RegisterValidation("test_global", func(fl validator.FieldLevel) bool {
		return true
	})
	assert.NoError(t, err)

	// 测试全局 RegisterTranslation
	RegisterTranslation("en", "test_global", "test message")
}

// TestValidatorFields 测试验证器字段覆盖
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

// TestValidationErrorsMultiple 测试多个错误的情况
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

// TestCustomValidatorEdgeCases 测试自定义验证器的边界情况
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

// TestLocaleConfigEdgeCases 测试地区配置边界情况
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

// TestMessageFormattingEdgeCases 测试消息格式化边界情况
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
