package validator

import (
	"reflect"
	"testing"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// TestBuiltinValidators tests the registerBuiltinValidators function
func TestBuiltinValidators(t *testing.T) {
	v, err := New()
	require.NoError(t, err)

	// Test required validator for different types
	t.Run("required_validator_string", func(t *testing.T) {
		type TestStruct struct {
			Value string `validate:"required"`
		}

		// Test valid string
		err := v.Struct(TestStruct{Value: "test"})
		assert.NoError(t, err)

		// Test empty string
		err = v.Struct(TestStruct{Value: ""})
		assert.Error(t, err)
	})

	t.Run("required_validator_slice", func(t *testing.T) {
		type TestStruct struct {
			Value []int `validate:"required"`
		}

		// Test non-empty slice
		err := v.Struct(TestStruct{Value: []int{1, 2, 3}})
		assert.NoError(t, err)

		// Test empty slice
		err = v.Struct(TestStruct{Value: []int{}})
		assert.Error(t, err)

		// Test nil slice
		err = v.Struct(TestStruct{Value: nil})
		assert.Error(t, err)
	})

	t.Run("required_validator_map", func(t *testing.T) {
		type TestStruct struct {
			Value map[string]int `validate:"required"`
		}

		// Test non-empty map
		err := v.Struct(TestStruct{Value: map[string]int{"key": 1}})
		assert.NoError(t, err)

		// Test empty map
		err = v.Struct(TestStruct{Value: map[string]int{}})
		assert.Error(t, err)

		// Test nil map
		err = v.Struct(TestStruct{Value: nil})
		assert.Error(t, err)
	})

	t.Run("required_validator_array", func(t *testing.T) {
		type TestStruct struct {
			Value [3]int `validate:"required"`
		}

		// Test array (arrays are always considered "present")
		err := v.Struct(TestStruct{Value: [3]int{1, 2, 3}})
		assert.NoError(t, err)

		// Test zero array (arrays are always considered present, so no error expected)
		err = v.Struct(TestStruct{Value: [3]int{}})
		assert.NoError(t, err)
	})

	t.Run("required_validator_pointer", func(t *testing.T) {
		type TestStruct struct {
			Value *string `validate:"required"`
		}

		// Test non-nil pointer
		str := "test"
		err := v.Struct(TestStruct{Value: &str})
		assert.NoError(t, err)

		// Test nil pointer
		err = v.Struct(TestStruct{Value: nil})
		assert.Error(t, err)
	})

	t.Run("required_validator_interface", func(t *testing.T) {
		type TestStruct struct {
			Value interface{} `validate:"required"`
		}

		// Test non-nil interface
		err := v.Struct(TestStruct{Value: "test"})
		assert.NoError(t, err)

		// Test nil interface
		err = v.Struct(TestStruct{Value: nil})
		assert.Error(t, err)
	})

	t.Run("required_validator_other_types", func(t *testing.T) {
		type TestStruct struct {
			IntValue    int     `validate:"required"`
			FloatValue  float64 `validate:"required"`
			BoolValue   bool    `validate:"required"`
		}

		// Test non-zero values
		err := v.Struct(TestStruct{IntValue: 1, FloatValue: 1.5, BoolValue: true})
		assert.NoError(t, err)

		// Test zero values
		err = v.Struct(TestStruct{IntValue: 0, FloatValue: 0.0, BoolValue: false})
		assert.Error(t, err)
	})

	t.Run("email_validator", func(t *testing.T) {
		type TestStruct struct {
			Email string `validate:"email"`
		}

		// Test valid emails
		validEmails := []string{
			"test@example.com",
			"user.name@domain.co.uk",
			"user+tag@example.org",
		}

		for _, email := range validEmails {
			err := v.Struct(TestStruct{Email: email})
			assert.NoError(t, err, "Expected %s to be valid", email)
		}

		// Test invalid emails
		invalidEmails := []string{
			"invalid",
			"@example.com",
			"test@",
			"test@.com",
		}

		for _, email := range invalidEmails {
			err := v.Struct(TestStruct{Email: email})
			assert.Error(t, err, "Expected %s to be invalid", email)
		}

		// Test empty email (in this implementation, empty email is considered invalid)
		err := v.Struct(TestStruct{Email: ""})
		assert.Error(t, err)
	})
}

// TestRegisterDefaultValidators tests the registerDefaultValidators function
func TestRegisterDefaultValidators(t *testing.T) {
	t.Run("mobile_validator", func(t *testing.T) {
		v, err := New()
		require.NoError(t, err)

		type TestStruct struct {
			Mobile string `validate:"mobile"`
		}

		// Test valid mobile numbers
		validMobiles := []string{
			"13812345678",
			"15987654321",
			"18612345678",
		}

		for _, mobile := range validMobiles {
			err := v.Struct(TestStruct{Mobile: mobile})
			assert.NoError(t, err, "Expected %s to be valid mobile", mobile)
		}

		// Test invalid mobile numbers
		invalidMobiles := []string{
			"123",
			"12345678901",
			"invalid",
		}

		for _, mobile := range invalidMobiles {
			err := v.Struct(TestStruct{Mobile: mobile})
			assert.Error(t, err, "Expected %s to be invalid mobile", mobile)
		}
	})

	t.Run("idcard_validator", func(t *testing.T) {
		v, err := New()
		require.NoError(t, err)

		type TestStruct struct {
			IDCard string `validate:"idcard"`
		}

		// Test valid ID cards
		validIDCards := []string{
			"11010119800101123X",
			"11010119800101124X",
		}

		for _, idcard := range validIDCards {
			err := v.Struct(TestStruct{IDCard: idcard})
			assert.NoError(t, err, "Expected %s to be valid idcard", idcard)
		}

		// Test invalid ID cards
		invalidIDCards := []string{
			"123",
			"11010119800101123Y", // Wrong checksum
			"invalid",
		}

		for _, idcard := range invalidIDCards {
			err := v.Struct(TestStruct{IDCard: idcard})
			assert.Error(t, err, "Expected %s to be invalid idcard", idcard)
		}
	})

	t.Run("bankcard_validator", func(t *testing.T) {
		v, err := New()
		require.NoError(t, err)

		type TestStruct struct {
			BankCard string `validate:"bankcard"`
		}

		// Test valid bank cards (using Luhn algorithm)
		validBankCards := []string{
			"4111111111111111", // Visa test card
			"5555555555554444", // Mastercard test card
		}

		for _, bankcard := range validBankCards {
			err := v.Struct(TestStruct{BankCard: bankcard})
			assert.NoError(t, err, "Expected %s to be valid bankcard", bankcard)
		}

		// Test invalid bank cards
		invalidBankCards := []string{
			"123",
			"4111111111111112", // Wrong checksum
			"invalid",
		}

		for _, bankcard := range invalidBankCards {
			err := v.Struct(TestStruct{BankCard: bankcard})
			assert.Error(t, err, "Expected %s to be invalid bankcard", bankcard)
		}
	})

	t.Run("chinese_name_validator", func(t *testing.T) {
		v, err := New()
		require.NoError(t, err)

		type TestStruct struct {
			Name string `validate:"chinese_name"`
		}

		// Test valid Chinese names
		validNames := []string{
			"张三",
			"李四",
			"王五",
		}

		for _, name := range validNames {
			err := v.Struct(TestStruct{Name: name})
			assert.NoError(t, err, "Expected %s to be valid chinese name", name)
		}

		// Test invalid Chinese names
		invalidNames := []string{
			"John",
			"123",
			"",
		}

		for _, name := range invalidNames {
			err := v.Struct(TestStruct{Name: name})
			assert.Error(t, err, "Expected %s to be invalid chinese name", name)
		}
	})

	t.Run("strong_password_validator", func(t *testing.T) {
		v, err := New()
		require.NoError(t, err)

		type TestStruct struct {
			Password string `validate:"strong_password"`
		}

		// Test valid strong passwords
		validPasswords := []string{
			"MyPass123!",
			"StrongP@ss1",
			"Complex#123",
		}

		for _, password := range validPasswords {
			err := v.Struct(TestStruct{Password: password})
			assert.NoError(t, err, "Expected %s to be valid strong password", password)
		}

		// Test invalid passwords
		invalidPasswords := []string{
			"weak",
			"12345678",
			"onlylowercase",
			"ONLYUPPERCASE",
		}

		for _, password := range invalidPasswords {
			err := v.Struct(TestStruct{Password: password})
			assert.Error(t, err, "Expected %s to be invalid strong password", password)
		}
	})
}

// TestTranslateFieldError tests the translateFieldError function with various error types
func TestTranslateFieldError(t *testing.T) {
	t.Run("different_locales", func(t *testing.T) {
		// Test English locale
		vEn, err := New(WithLocale("en"))
		require.NoError(t, err)

		type TestStruct struct {
			Name  string `validate:"required"`
			Email string `validate:"email"`
		}

		invalidData := TestStruct{Name: "", Email: "invalid"}
		err = vEn.Struct(invalidData)
		require.Error(t, err)

		// Should contain English error messages
		errorMsg := err.Error()
		assert.NotEmpty(t, errorMsg)

		// Test Chinese locale
		vZh, err := New(WithLocale("zh"))
		require.NoError(t, err)

		err = vZh.Struct(invalidData)
		require.Error(t, err)

		// Should contain Chinese error messages
		errorMsgZh := err.Error()
		assert.NotEmpty(t, errorMsgZh)

		// Test Traditional Chinese locale
		vZhTw, err := New(WithLocale("zh_tw"))
		require.NoError(t, err)

		err = vZhTw.Struct(invalidData)
		require.Error(t, err)

		errorMsgZhTw := err.Error()
		assert.NotEmpty(t, errorMsgZhTw)

		// Test Japanese locale
		vJa, err := New(WithLocale("ja"))
		require.NoError(t, err)

		err = vJa.Struct(invalidData)
		require.Error(t, err)

		errorMsgJa := err.Error()
		assert.NotEmpty(t, errorMsgJa)

		// Test Korean locale
		vKo, err := New(WithLocale("ko"))
		require.NoError(t, err)

		err = vKo.Struct(invalidData)
		require.Error(t, err)

		errorMsgKo := err.Error()
		assert.NotEmpty(t, errorMsgKo)

		// Test French locale
		vFr, err := New(WithLocale("fr"))
		require.NoError(t, err)

		err = vFr.Struct(invalidData)
		require.Error(t, err)

		errorMsgFr := err.Error()
		assert.NotEmpty(t, errorMsgFr)
	})

	t.Run("error_translation_coverage", func(t *testing.T) {
		v, err := New(WithLocale("en"))
		require.NoError(t, err)

		type TestStruct struct {
			Required     string  `validate:"required"`
			Email        string  `validate:"email"`
			Min          string  `validate:"min=5"`
			Max          string  `validate:"max=10"`
			Mobile       string  `validate:"mobile"`
			IDCard       string  `validate:"idcard"`
			BankCard     string  `validate:"bankcard"`
			ChineseName  string  `validate:"chinese_name"`
			StrongPass   string  `validate:"strong_password"`
		}

		// Create data that violates all validators
		invalidData := TestStruct{
			Required:    "",
			Email:       "invalid-email",
			Min:         "abc",    // Too short
			Max:         "this is too long",
			Mobile:      "invalid-mobile",
			IDCard:      "invalid-idcard",
			BankCard:    "invalid-bankcard",
			ChineseName: "Invalid Name",
			StrongPass:  "weak",
		}

		err = v.Struct(invalidData)
		require.Error(t, err)

		// The error should contain translations for all field errors
		errorMsg := err.Error()
		assert.NotEmpty(t, errorMsg)
		t.Logf("Error message: %s", errorMsg)
	})

	t.Run("var_validation_errors", func(t *testing.T) {
		v, err := New(WithLocale("en"))
		require.NoError(t, err)

		// Test various variable validations to trigger translateFieldError
		err = v.Var("", "required")
		assert.Error(t, err)
		assert.NotEmpty(t, err.Error())

		err = v.Var("invalid-email", "email")
		assert.Error(t, err)
		assert.NotEmpty(t, err.Error())

		err = v.Var("abc", "min=5")
		assert.Error(t, err)
		assert.NotEmpty(t, err.Error())

		err = v.Var("this is way too long", "max=5")
		assert.Error(t, err)
		assert.NotEmpty(t, err.Error())
	})
}

// TestEngineCustomValidation tests the Engine's built-in validators with edge cases
func TestEngineCustomValidation(t *testing.T) {
	t.Run("field_level_interface", func(t *testing.T) {
		v, err := New()
		require.NoError(t, err)

		// Register a custom validator to test FieldLevel interface
		err = v.RegisterValidation("custom_test", func(fl FieldLevel) bool {
			field := fl.Field()
			if field.Kind() == reflect.String {
				return field.String() == "expected"
			}
			return false
		})
		require.NoError(t, err)

		type TestStruct struct {
			Value string `validate:"custom_test"`
		}

		// Test valid value
		err = v.Struct(TestStruct{Value: "expected"})
		assert.NoError(t, err)

		// Test invalid value
		err = v.Struct(TestStruct{Value: "unexpected"})
		assert.Error(t, err)
	})
}

// TestValidatorConcurrency tests validator under concurrent usage
func TestValidatorConcurrency(t *testing.T) {
	v, err := New()
	require.NoError(t, err)

	type TestStruct struct {
		Name  string `validate:"required"`
		Email string `validate:"email"`
	}

	validData := TestStruct{
		Name:  "Test User",
		Email: "test@example.com",
	}

	invalidData := TestStruct{
		Name:  "",
		Email: "invalid",
	}

	// Run concurrent validations
	const numGoroutines = 10
	done := make(chan bool, numGoroutines)

	for i := 0; i < numGoroutines; i++ {
		go func(id int) {
			defer func() { done <- true }()

			// Test valid data
			err := v.Struct(validData)
			assert.NoError(t, err)

			// Test invalid data
			err = v.Struct(invalidData)
			assert.Error(t, err)

			// Test variable validation
			err = v.Var("test@example.com", "email")
			assert.NoError(t, err)

			err = v.Var("invalid", "email")
			assert.Error(t, err)
		}(i)
	}

	// Wait for all goroutines to complete
	for i := 0; i < numGoroutines; i++ {
		<-done
	}
}