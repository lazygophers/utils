package validator

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewWithInvalidLocale(t *testing.T) {
	v, err := New(WithLocale("invalid-locale"))
	require.NoError(t, err)
	assert.NotNil(t, v)
	assert.Equal(t, "invalid-locale", v.GetLocale())
}

func TestNewWithCustomTranslations(t *testing.T) {
	translations := map[string]string{
		"en.required":        "Custom required message",
		"zh.required":        "自定义必填消息",
		"en.email":           "Custom email message",
		"en.mobile":          "Custom mobile message",
		"en.idcard":          "Custom idcard message",
		"en.bankcard":        "Custom bankcard message",
		"en.chinese_name":    "Custom chinese_name message",
		"en.strong_password": "Custom strong_password message",
	}

	v, err := New(WithTranslations(translations))
	require.NoError(t, err)
	assert.NotNil(t, v)
}

func TestNewWithCustomValidator(t *testing.T) {
	v, err := New(
		WithCustomValidator("custom", func(value interface{}) bool {
			if str, ok := value.(string); ok {
				return str == "valid"
			}
			return false
		}),
	)
	require.NoError(t, err)

	err1 := v.Var("valid", "custom")
	assert.NoError(t, err1)

	err2 := v.Var("invalid", "custom")
	assert.Error(t, err2)
}

func TestNewWithConfig(t *testing.T) {
	config := Config{
		Locale:  "en",
		UseJSON: true,
		Translations: map[string]string{
			"en.required": "Config required message",
		},
	}

	v, err := New(WithConfig(config))
	require.NoError(t, err)
	assert.Equal(t, "en", v.GetLocale())
}

func TestStructWithNilValue(t *testing.T) {
	v, err := New()
	require.NoError(t, err)

	err = v.Struct(nil)
	assert.Error(t, err)
}

func TestStructWithNonStruct(t *testing.T) {
	v, err := New()
	require.NoError(t, err)

	err = v.Struct("not a struct")
	assert.Error(t, err)
}

func TestStructWithValidData(t *testing.T) {
	v, err := New()
	require.NoError(t, err)

	type TestStruct struct {
		Name  string `validate:"required"`
		Email string `validate:"required,email"`
	}

	validData := TestStruct{
		Name:  "John Doe",
		Email: "john@example.com",
	}

	err = v.Struct(validData)
	assert.NoError(t, err)
}

func TestStructWithInvalidData(t *testing.T) {
	v, err := New()
	require.NoError(t, err)

	type TestStruct struct {
		Name  string `validate:"required"`
		Email string `validate:"required,email"`
	}

	invalidData := TestStruct{
		Name:  "",
		Email: "invalid-email",
	}

	err = v.Struct(invalidData)
	assert.Error(t, err)

	validationErrors, ok := err.(ValidationErrors)
	require.True(t, ok)
	assert.True(t, len(validationErrors) > 0)
}

func TestVarWithNilValue(t *testing.T) {
	v, err := New()
	require.NoError(t, err)

	err = v.Var(nil, "required")
	assert.Error(t, err)
}

func TestVarWithValidValue(t *testing.T) {
	v, err := New()
	require.NoError(t, err)

	err = v.Var("test@example.com", "email")
	assert.NoError(t, err)
}

func TestVarWithInvalidValue(t *testing.T) {
	v, err := New()
	require.NoError(t, err)

	err = v.Var("invalid-email", "email")
	assert.Error(t, err)

	fieldError, ok := err.(*FieldError)
	require.True(t, ok)
	assert.NotEmpty(t, fieldError.Message)
}

func TestRegisterValidationWithDuplicate(t *testing.T) {
	v, err := New()
	require.NoError(t, err)

	err1 := v.RegisterValidation("custom", func(fl FieldLevel) bool {
		return true
	})
	assert.NoError(t, err1)

	err2 := v.RegisterValidation("custom", func(fl FieldLevel) bool {
		return false
	})
	assert.NoError(t, err2)
}

func TestRegisterTranslationWithDifferentLocales(t *testing.T) {
	v, err := New()
	require.NoError(t, err)

	v.RegisterTranslation("en", "test", "Test message in English")
	v.RegisterTranslation("zh", "test", "测试消息")
	v.RegisterTranslation("fr", "test", "Message de test")
	v.RegisterTranslation("es", "test", "Mensaje de prueba")
}

func TestTranslateFieldErrorWithCustomTranslation(t *testing.T) {
	v, err := New()
	require.NoError(t, err)

	v.RegisterTranslation("en", "required", "Custom required message for {field}")

	type TestStruct struct {
		Name string `validate:"required"`
	}

	testStruct := TestStruct{Name: ""}
	valErr := v.Struct(testStruct)
	require.Error(t, valErr)

	validationErrors := valErr.(ValidationErrors)
	assert.Contains(t, validationErrors.Error(), "Custom required message")
}

func TestTranslateFieldErrorWithDefaultTranslation(t *testing.T) {
	v, err := New(WithLocale("en"))
	require.NoError(t, err)

	type TestStruct struct {
		Name string `validate:"required"`
	}

	testStruct := TestStruct{Name: ""}
	valErr := v.Struct(testStruct)
	require.Error(t, valErr)

	validationErrors := valErr.(ValidationErrors)
	assert.NotEmpty(t, validationErrors.Error())
}

func TestTranslateFieldErrorWithUnknownLocale(t *testing.T) {
	v, err := New(WithLocale("unknown-locale"))
	require.NoError(t, err)

	type TestStruct struct {
		Name string `validate:"required"`
	}

	testStruct := TestStruct{Name: ""}
	valErr := v.Struct(testStruct)
	require.Error(t, valErr)

	validationErrors := valErr.(ValidationErrors)
	assert.NotEmpty(t, validationErrors.Error())
}

func TestFormatMessageWithAllPlaceholders(t *testing.T) {
	v, err := New()
	require.NoError(t, err)

	fieldErr := &FieldError{
		Field: "email",
		Tag:   "email",
		Param: "test",
		Value: "test@example.com",
	}

	template := "{field} failed {tag} validation with param {param} and value {value}"
	result := v.formatMessage(template, fieldErr)

	assert.Contains(t, result, "email")
	assert.Contains(t, result, "email")
	assert.Contains(t, result, "test")
	assert.Contains(t, result, "test@example.com")
}

func TestFormatMessageWithNilValue(t *testing.T) {
	v, err := New()
	require.NoError(t, err)

	fieldErr := &FieldError{
		Field: "field",
		Tag:   "required",
		Param: "",
		Value: nil,
	}

	template := "{field} is required with value {value}"
	result := v.formatMessage(template, fieldErr)

	assert.Contains(t, result, "field")
	assert.Contains(t, result, "is required")
}

func TestFormatMessageWithEmptyPlaceholders(t *testing.T) {
	v, err := New()
	require.NoError(t, err)

	fieldErr := &FieldError{
		Field: "field",
		Tag:   "tag",
		Param: "",
		Value: nil,
	}

	template := "{field} {tag}"
	result := v.formatMessage(template, fieldErr)

	assert.Equal(t, "field tag", result)
}

func TestSetLocale(t *testing.T) {
	v, err := New()
	require.NoError(t, err)

	v.SetLocale("zh")
	assert.Equal(t, "zh", v.GetLocale())

	v.SetLocale("en")
	assert.Equal(t, "en", v.GetLocale())
}

func TestSetUseJSON(t *testing.T) {
	v, err := New()
	require.NoError(t, err)

	v.SetUseJSON(true)
	v.SetUseJSON(false)
}

func TestGetLocale(t *testing.T) {
	v, err := New(WithLocale("zh"))
	require.NoError(t, err)

	assert.Equal(t, "zh", v.GetLocale())
}

func TestRegisterDefaultValidatorsSuccess(t *testing.T) {
	v, err := New()
	require.NoError(t, err)

	assert.NotNil(t, v)
}

func TestRegisterDefaultValidatorsCoverage(t *testing.T) {
	v, err := New()
	require.NoError(t, err)

	tests := []struct {
		name  string
		tag   string
		value interface{}
		valid bool
	}{
		{"mobile valid", "mobile", "13812345678", true},
		{"mobile invalid", "mobile", "123", false},
		{"idcard valid", "idcard", "11010119800101123X", true},
		{"idcard invalid", "idcard", "123", false},
		{"bankcard valid", "bankcard", "4111111111111111", true},
		{"bankcard invalid", "bankcard", "123", false},
		{"chinese_name valid", "chinese_name", "张三", true},
		{"chinese_name invalid", "chinese_name", "John", false},
		{"strong_password valid", "strong_password", "MyPass123!", true},
		{"strong_password invalid", "strong_password", "123", false},
		{"email valid", "email", "test@example.com", true},
		{"email invalid", "email", "invalid", false},
		{"url valid", "url", "http://example.com", true},
		{"url invalid", "url", "not a url", false},
		{"ipv4 valid", "ipv4", "192.168.1.1", true},
		{"ipv4 invalid", "ipv4", "not an ip", false},
		{"mac valid", "mac", "00:11:22:33:44:55", true},
		{"mac invalid", "mac", "not a mac", false},
		{"json valid", "json", `{"key":"value"}`, true},
		{"json invalid", "json", "not json", false},
		{"uuid valid", "uuid", "550e8400-e29b-41d4-a716-446655440000", true},
		{"uuid invalid", "uuid", "not a uuid", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := v.Var(tt.value, tt.tag)
			if tt.valid {
				assert.NoError(t, err)
			} else {
				assert.Error(t, err)
			}
		})
	}
}

func TestGlobalSetLocale(t *testing.T) {
	SetLocale("zh")
	SetLocale("en")
}

func TestGlobalSetUseJSON(t *testing.T) {
	SetUseJSON(true)
	SetUseJSON(false)
}

func TestGlobalStruct(t *testing.T) {
	type TestStruct struct {
		Name string `validate:"required"`
	}

	validStruct := TestStruct{Name: "John"}
	err := Struct(validStruct)
	assert.NoError(t, err)

	invalidStruct := TestStruct{Name: ""}
	err = Struct(invalidStruct)
	assert.Error(t, err)
}

func TestGlobalVar(t *testing.T) {
	err := Var("test@example.com", "email")
	assert.NoError(t, err)

	err = Var("invalid", "email")
	assert.Error(t, err)
}

func TestGlobalRegisterValidation(t *testing.T) {
	err := RegisterValidation("global_test", func(fl FieldLevel) bool {
		return true
	})
	assert.NoError(t, err)
}

func TestGlobalRegisterTranslation(t *testing.T) {
	RegisterTranslation("en", "global_test", "Global test message")
	RegisterTranslation("zh", "global_test", "全局测试消息")
}

func TestValidatorWithMultipleErrors(t *testing.T) {
	v, err := New()
	require.NoError(t, err)

	type TestStruct struct {
		Name  string `validate:"required,min=3"`
		Email string `validate:"required,email"`
		Age   int    `validate:"required,min=18"`
	}

	testStruct := TestStruct{
		Name:  "",
		Email: "invalid",
		Age:   10,
	}

	valErr := v.Struct(testStruct)
	require.Error(t, valErr)

	validationErrors := valErr.(ValidationErrors)
	assert.True(t, len(validationErrors) >= 3)
}

func TestValidatorWithNestedStruct(t *testing.T) {
	v, err := New()
	require.NoError(t, err)

	type Address struct {
		City    string `validate:"required"`
		Country string `validate:"required"`
	}

	type Person struct {
		Name    string  `validate:"required"`
		Address Address `validate:"required,dive"`
	}

	person := Person{
		Name: "",
		Address: Address{
			City:    "",
			Country: "",
		},
	}

	valErr := v.Struct(person)
	require.Error(t, valErr)
}

func TestValidatorWithSlice(t *testing.T) {
	v, err := New()
	require.NoError(t, err)

	type TestStruct struct {
		Emails []string `validate:"required,dive,email"`
	}

	testStruct := TestStruct{
		Emails: []string{"valid@example.com", "invalid", "another@example.com"},
	}

	valErr := v.Struct(testStruct)
	require.Error(t, valErr)
}

func TestValidatorWithMap(t *testing.T) {
	v, err := New()
	require.NoError(t, err)

	type TestStruct struct {
		Data map[string]string `validate:"required"`
	}

	testStruct := TestStruct{
		Data: map[string]string{"key": "value"},
	}

	valErr := v.Struct(testStruct)
	assert.NoError(t, valErr)
}

func TestValidatorWithPointer(t *testing.T) {
	v, err := New()
	require.NoError(t, err)

	type TestStruct struct {
		Name *string `validate:"required"`
	}

	name := "John"
	testStruct := TestStruct{Name: &name}
	valErr := v.Struct(testStruct)
	assert.NoError(t, valErr)

	testStruct2 := TestStruct{Name: nil}
	valErr2 := v.Struct(testStruct2)
	assert.Error(t, valErr2)
}

func TestValidatorWithInterface(t *testing.T) {
	v, err := New()
	require.NoError(t, err)

	type TestStruct struct {
		Data interface{} `validate:"required"`
	}

	testStruct := TestStruct{Data: "value"}
	valErr := v.Struct(testStruct)
	assert.NoError(t, valErr)

	testStruct2 := TestStruct{Data: nil}
	valErr2 := v.Struct(testStruct2)
	assert.Error(t, valErr2)
}

func TestValidatorWithCustomTag(t *testing.T) {
	v, err := New()
	require.NoError(t, err)

	err = v.RegisterValidation("custom_tag", func(fl FieldLevel) bool {
		value := fl.Field().String()
		return value == "valid"
	})
	require.NoError(t, err)

	type TestStruct struct {
		Field string `validate:"custom_tag"`
	}

	validStruct := TestStruct{Field: "valid"}
	valErr := v.Struct(validStruct)
	assert.NoError(t, valErr)

	invalidStruct := TestStruct{Field: "invalid"}
	valErr2 := v.Struct(invalidStruct)
	assert.Error(t, valErr2)
}

func TestValidatorWithMultipleTags(t *testing.T) {
	v, err := New()
	require.NoError(t, err)

	type TestStruct struct {
		Field string `validate:"required,min=3,max=20,email"`
	}

	validStruct := TestStruct{Field: "test@example.com"}
	valErr := v.Struct(validStruct)
	assert.NoError(t, valErr)

	invalidStruct := TestStruct{Field: "ab"}
	valErr2 := v.Struct(invalidStruct)
	assert.Error(t, valErr2)
}

func TestValidatorWithRequiredTag(t *testing.T) {
	v, err := New()
	require.NoError(t, err)

	type TestStruct struct {
		Field string `validate:"required"`
	}

	invalidStruct := TestStruct{Field: ""}
	valErr := v.Struct(invalidStruct)
	require.Error(t, valErr)

	validationErrors := valErr.(ValidationErrors)
	assert.True(t, validationErrors.HasField("Field"))
}

func TestValidatorWithMinTag(t *testing.T) {
	v, err := New()
	require.NoError(t, err)

	type TestStruct struct {
		Age int `validate:"min=18"`
	}

	validStruct := TestStruct{Age: 20}
	valErr := v.Struct(validStruct)
	assert.NoError(t, valErr)

	invalidStruct := TestStruct{Age: 15}
	valErr2 := v.Struct(invalidStruct)
	assert.Error(t, valErr2)
}

func TestValidatorWithMaxTag(t *testing.T) {
	v, err := New()
	require.NoError(t, err)

	type TestStruct struct {
		Age int `validate:"max=100"`
	}

	validStruct := TestStruct{Age: 50}
	valErr := v.Struct(validStruct)
	assert.NoError(t, valErr)

	invalidStruct := TestStruct{Age: 150}
	valErr2 := v.Struct(invalidStruct)
	assert.Error(t, valErr2)
}

func TestValidatorWithLenTag(t *testing.T) {
	v, err := New()
	require.NoError(t, err)

	type TestStruct struct {
		Name string `validate:"len=5"`
	}

	validStruct := TestStruct{Name: "Hello"}
	valErr := v.Struct(validStruct)
	assert.NoError(t, valErr)

	invalidStruct := TestStruct{Name: "Hi"}
	valErr2 := v.Struct(invalidStruct)
	assert.Error(t, valErr2)
}

func TestValidatorWithEqTag(t *testing.T) {
	v, err := New()
	require.NoError(t, err)

	type TestStruct struct {
		Age int `validate:"eq=18"`
	}

	validStruct := TestStruct{Age: 18}
	valErr := v.Struct(validStruct)
	assert.NoError(t, valErr)

	invalidStruct := TestStruct{Age: 20}
	valErr2 := v.Struct(invalidStruct)
	assert.Error(t, valErr2)
}

func TestValidatorWithNeTag(t *testing.T) {
	v, err := New()
	require.NoError(t, err)

	type TestStruct struct {
		Age int `validate:"ne=18"`
	}

	validStruct := TestStruct{Age: 20}
	valErr := v.Struct(validStruct)
	assert.NoError(t, valErr)

	invalidStruct := TestStruct{Age: 18}
	valErr2 := v.Struct(invalidStruct)
	assert.Error(t, valErr2)
}

func TestValidatorWithAlphaTag(t *testing.T) {
	v, err := New()
	require.NoError(t, err)

	type TestStruct struct {
		Name string `validate:"alpha"`
	}

	validStruct := TestStruct{Name: "Hello"}
	valErr := v.Struct(validStruct)
	assert.NoError(t, valErr)

	invalidStruct := TestStruct{Name: "Hello123"}
	valErr2 := v.Struct(invalidStruct)
	assert.Error(t, valErr2)
}

func TestValidatorWithAlphanumTag(t *testing.T) {
	v, err := New()
	require.NoError(t, err)

	type TestStruct struct {
		Code string `validate:"alphanum"`
	}

	validStruct := TestStruct{Code: "Hello123"}
	valErr := v.Struct(validStruct)
	assert.NoError(t, valErr)

	invalidStruct := TestStruct{Code: "Hello-123"}
	valErr2 := v.Struct(invalidStruct)
	assert.Error(t, valErr2)
}

func TestValidatorWithNumericTag(t *testing.T) {
	v, err := New()
	require.NoError(t, err)

	type TestStruct struct {
		Number string `validate:"numeric"`
	}

	validStruct := TestStruct{Number: "123"}
	valErr := v.Struct(validStruct)
	assert.NoError(t, valErr)

	invalidStruct := TestStruct{Number: "abc"}
	valErr2 := v.Struct(invalidStruct)
	assert.Error(t, valErr2)
}

func TestValidatorWithEmailTag(t *testing.T) {
	v, err := New()
	require.NoError(t, err)

	type TestStruct struct {
		Email string `validate:"email"`
	}

	validStruct := TestStruct{Email: "test@example.com"}
	valErr := v.Struct(validStruct)
	assert.NoError(t, valErr)

	invalidStruct := TestStruct{Email: "invalid"}
	valErr2 := v.Struct(invalidStruct)
	assert.Error(t, valErr2)
}

func TestValidatorWithUrlTag(t *testing.T) {
	v, err := New()
	require.NoError(t, err)

	type TestStruct struct {
		URL string `validate:"url"`
	}

	validStruct := TestStruct{URL: "http://example.com"}
	valErr := v.Struct(validStruct)
	assert.NoError(t, valErr)

	invalidStruct := TestStruct{URL: "not a url"}
	valErr2 := v.Struct(invalidStruct)
	assert.Error(t, valErr2)
}

func TestValidatorWithUuidTag(t *testing.T) {
	v, err := New()
	require.NoError(t, err)

	type TestStruct struct {
		ID string `validate:"uuid"`
	}

	validStruct := TestStruct{ID: "550e8400-e29b-41d4-a716-446655440000"}
	valErr := v.Struct(validStruct)
	assert.NoError(t, valErr)

	invalidStruct := TestStruct{ID: "not a uuid"}
	valErr2 := v.Struct(invalidStruct)
	assert.Error(t, valErr2)
}

func TestValidatorWithJsonTag(t *testing.T) {
	v, err := New()
	require.NoError(t, err)

	type TestStruct struct {
		Data string `validate:"json"`
	}

	validStruct := TestStruct{Data: `{"key":"value"}`}
	valErr := v.Struct(validStruct)
	assert.NoError(t, valErr)

	invalidStruct := TestStruct{Data: "not json"}
	valErr2 := v.Struct(invalidStruct)
	assert.Error(t, valErr2)
}

func TestValidatorWithIpv4Tag(t *testing.T) {
	v, err := New()
	require.NoError(t, err)

	type TestStruct struct {
		IP string `validate:"ipv4"`
	}

	validStruct := TestStruct{IP: "192.168.1.1"}
	valErr := v.Struct(validStruct)
	assert.NoError(t, valErr)

	invalidStruct := TestStruct{IP: "not an ip"}
	valErr2 := v.Struct(invalidStruct)
	assert.Error(t, valErr2)
}

func TestValidatorWithMacTag(t *testing.T) {
	v, err := New()
	require.NoError(t, err)

	type TestStruct struct {
		MAC string `validate:"mac"`
	}

	validStruct := TestStruct{MAC: "00:11:22:33:44:55"}
	valErr := v.Struct(validStruct)
	assert.NoError(t, valErr)

	invalidStruct := TestStruct{MAC: "not a mac"}
	valErr2 := v.Struct(invalidStruct)
	assert.Error(t, valErr2)
}

func TestValidatorWithMobileTag(t *testing.T) {
	v, err := New()
	require.NoError(t, err)

	type TestStruct struct {
		Phone string `validate:"mobile"`
	}

	validStruct := TestStruct{Phone: "13812345678"}
	valErr := v.Struct(validStruct)
	assert.NoError(t, valErr)

	invalidStruct := TestStruct{Phone: "123"}
	valErr2 := v.Struct(invalidStruct)
	assert.Error(t, valErr2)
}

func TestValidatorWithIdcardTag(t *testing.T) {
	v, err := New()
	require.NoError(t, err)

	type TestStruct struct {
		IDCard string `validate:"idcard"`
	}

	validStruct := TestStruct{IDCard: "11010119800101123X"}
	valErr := v.Struct(validStruct)
	assert.NoError(t, valErr)

	invalidStruct := TestStruct{IDCard: "123"}
	valErr2 := v.Struct(invalidStruct)
	assert.Error(t, valErr2)
}

func TestValidatorWithBankcardTag(t *testing.T) {
	v, err := New()
	require.NoError(t, err)

	type TestStruct struct {
		BankCard string `validate:"bankcard"`
	}

	validStruct := TestStruct{BankCard: "4111111111111111"}
	valErr := v.Struct(validStruct)
	assert.NoError(t, valErr)

	invalidStruct := TestStruct{BankCard: "123"}
	valErr2 := v.Struct(invalidStruct)
	assert.Error(t, valErr2)
}

func TestValidatorWithChineseNameTag(t *testing.T) {
	v, err := New()
	require.NoError(t, err)

	type TestStruct struct {
		Name string `validate:"chinese_name"`
	}

	validStruct := TestStruct{Name: "张三"}
	valErr := v.Struct(validStruct)
	assert.NoError(t, valErr)

	invalidStruct := TestStruct{Name: "John"}
	valErr2 := v.Struct(invalidStruct)
	assert.Error(t, valErr2)
}

func TestValidatorWithStrongPasswordTag(t *testing.T) {
	v, err := New()
	require.NoError(t, err)

	type TestStruct struct {
		Password string `validate:"strong_password"`
	}

	validStruct := TestStruct{Password: "MyPass123!"}
	valErr := v.Struct(validStruct)
	assert.NoError(t, valErr)

	invalidStruct := TestStruct{Password: "123"}
	valErr2 := v.Struct(invalidStruct)
	assert.Error(t, valErr2)
}
