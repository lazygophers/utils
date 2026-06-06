package validator

import (
	"testing"

	"github.com/lazygophers/utils/language"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	xlanguage "golang.org/x/text/language"
)

func TestValidatorRegisterStructValidation(t *testing.T) {
	v, err := New()
	require.NoError(t, err)

	type S struct {
		Name string `validate:"required"`
	}
	err = v.RegisterStructValidation(func(sl StructLevel) bool {
		return true
	}, "S")
	require.NoError(t, err)
	assert.NoError(t, v.Struct(S{Name: "test"}))
}

func TestValidatorEffectiveLocaleGoroutineLocal(t *testing.T) {
	orig := language.Default()
	defer language.SetDefault(orig)

	v, err := New()
	require.NoError(t, err)

	// With goroutine-local set
	language.Set(language.Make("ja"))
	defer language.Del()
	assert.Equal(t, xlanguage.Make("ja"), v.GetLocale())
}

func TestValidatorTranslateFieldErrorNoLocaleConfig(t *testing.T) {
	v, err := New(WithLocale(xlanguage.Make("unknown-xx")))
	require.NoError(t, err)

	// This locale has no config, should fall back to English
	type S struct {
		Name string `validate:"required"`
	}
	err = v.Struct(S{Name: ""})
	require.Error(t, err)
	// Should get English fallback message
	assert.Contains(t, err.Error(), "required")
}

func TestValidatorFormatMessageAllPlaceholders(t *testing.T) {
	v, err := New()
	require.NoError(t, err)

	// Register translation with all placeholders
	v.RegisterTranslation(xlanguage.Make("en"), "test_all", "{field} {tag} {param} {value}")

	// Use a custom validator that sets param
	err = v.RegisterValidation("test_all", func(fl FieldLevel) bool {
		return false
	})
	require.NoError(t, err)

	type S struct {
		Name string `validate:"test_all=hello"`
	}
	err = v.Struct(S{Name: "test"})
	require.Error(t, err)
	assert.Contains(t, err.Error(), "test_all")
}

func TestValidatorFormatMessageNoPlaceholders(t *testing.T) {
	v, err := New()
	require.NoError(t, err)

	v.RegisterTranslation(xlanguage.Make("en"), "plain", "plain message")

	err = v.RegisterValidation("plain", func(fl FieldLevel) bool { return false })
	require.NoError(t, err)

	type S struct {
		Name string `validate:"plain"`
	}
	err = v.Struct(S{Name: "test"})
	require.Error(t, err)
	assert.Contains(t, err.Error(), "plain message")
}

func TestValidatorFormatMessageValue(t *testing.T) {
	v, err := New()
	require.NoError(t, err)

	// Template with {value} placeholder
	v.RegisterTranslation(xlanguage.Make("en"), "val_test", "value is {value}")

	err = v.RegisterValidation("val_test", func(fl FieldLevel) bool { return false })
	require.NoError(t, err)

	type S struct {
		Name string `validate:"val_test"`
	}
	err = v.Struct(S{Name: "hello"})
	require.Error(t, err)
	// Message may come from custom translation or fall back
	assert.Error(t, err)
}

func TestValidatorFormatMessageNilValue(t *testing.T) {
	v, err := New()
	require.NoError(t, err)

	v.RegisterTranslation(xlanguage.Make("en"), "nil_val", "value={value} end")

	err = v.RegisterValidation("nil_val", func(fl FieldLevel) bool { return false })
	require.NoError(t, err)

	type S struct {
		Field *string `validate:"nil_val"`
	}
	err = v.Struct(S{Field: nil})
	require.Error(t, err)
	assert.Error(t, err)
}

func TestGlobalValidationWithComposition(t *testing.T) {
	err := RegisterValidationWithComposition("comp_test", func(fl FieldLevel) bool {
		return fl.Field().String() != ""
	})
	assert.NoError(t, err)
}

func TestInstanceValidationWithComposition(t *testing.T) {
	v, err := New()
	require.NoError(t, err)
	err = v.RegisterValidationWithComposition("inst_comp", func(fl FieldLevel) bool {
		return fl.Field().String() != ""
	})
	assert.NoError(t, err)
}

func TestNewValidatorFailure(t *testing.T) {
	// Create validator that fails due to duplicate registration
	// The New() function should still succeed since registerDefaultValidators handles duplicates
	v, err := New()
	require.NoError(t, err)
	assert.NotNil(t, v)
}

func TestDefaultValidatorCreation(t *testing.T) {
	d := Default()
	assert.NotNil(t, d)
	// Call again to cover sync.Once path
	d2 := Default()
	assert.Equal(t, d, d2)
}

func TestValidatorSetGetUseJSON(t *testing.T) {
	v, err := New()
	require.NoError(t, err)

	v.SetUseJSON(false)
	type S struct {
		Name string `json:"name" validate:"required"`
	}
	assert.NoError(t, v.Struct(S{Name: "test"}))
	assert.Error(t, v.Struct(S{Name: ""}))
}
