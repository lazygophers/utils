package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// Test structs for validation
type ValidateTestStruct struct {
	Name     string `validate:"required,min=2,max=50"`
	Email    string `validate:"required,email"`
	Age      int    `validate:"required,min=18,max=120"`
	Website  string `validate:"omitempty,url"`
	Password string `validate:"required,min=8"`
}

type NestedValidateStruct struct {
	ID       int                 `validate:"required,gt=0"`
	User     ValidateTestStruct  `validate:"required"`
	Tags     []string           `validate:"required,min=1"`
	Settings map[string]string  `validate:"required"`
}

type OptionalFieldsStruct struct {
	RequiredField string `validate:"required"`
	OptionalField string `validate:"omitempty,min=5"`
	OptionalEmail string `validate:"omitempty,email"`
}

type NumericValidationStruct struct {
	PositiveInt   int     `validate:"gt=0"`
	NonNegativeInt int    `validate:"gte=0"`
	LimitedFloat  float64 `validate:"min=0.0,max=100.0"`
	RangeInt      int     `validate:"min=10,max=100"`
}

type StringValidationStruct struct {
	AlphaOnly     string `validate:"alpha"`
	AlphaNumOnly  string `validate:"alphanum"`
	NumericOnly   string `validate:"numeric"`
	FixedLength   string `validate:"len=10"`
	MinLength     string `validate:"min=5"`
	MaxLength     string `validate:"max=20"`
}

func TestValidate(t *testing.T) {
	tests := []struct {
		name        string
		input       interface{}
		expectError bool
		description string
	}{
		// Valid cases
		{
			name: "valid_struct",
			input: ValidateTestStruct{
				Name:     "John Doe",
				Email:    "john@example.com",
				Age:      30,
				Website:  "https://example.com",
				Password: "password123",
			},
			expectError: false,
			description: "All fields are valid",
		},
		{
			name: "valid_struct_minimal",
			input: ValidateTestStruct{
				Name:     "Jo",
				Email:    "j@e.co",
				Age:      18,
				Password: "12345678",
			},
			expectError: false,
			description: "Minimal valid values",
		},
		{
			name: "valid_struct_maximal",
			input: ValidateTestStruct{
				Name:     "Very Long Name That Is Still Within Limits Here",
				Email:    "very.long.email.address@example-domain.com",
				Age:      120,
				Website:  "https://very-long-website-url.example.com/path",
				Password: "very-long-password-123456789",
			},
			expectError: false,
			description: "Maximum valid values",
		},
		{
			name: "valid_nested_struct",
			input: NestedValidateStruct{
				ID: 1,
				User: ValidateTestStruct{
					Name:     "Jane Doe",
					Email:    "jane@example.com",
					Age:      25,
					Password: "password123",
				},
				Tags:     []string{"tag1", "tag2"},
				Settings: map[string]string{"key": "value"},
			},
			expectError: false,
			description: "Valid nested structure",
		},
		{
			name: "valid_optional_fields_with_values",
			input: OptionalFieldsStruct{
				RequiredField: "required",
				OptionalField: "optional value here",
				OptionalEmail: "test@example.com",
			},
			expectError: false,
			description: "Optional fields with valid values",
		},
		{
			name: "valid_optional_fields_empty",
			input: OptionalFieldsStruct{
				RequiredField: "required",
			},
			expectError: false,
			description: "Optional fields are empty (omitempty)",
		},
		{
			name: "valid_numeric_struct",
			input: NumericValidationStruct{
				PositiveInt:    5,
				NonNegativeInt: 0,
				LimitedFloat:   50.5,
				RangeInt:       55,
			},
			expectError: false,
			description: "All numeric validations pass",
		},
		{
			name: "valid_string_struct",
			input: StringValidationStruct{
				AlphaOnly:    "AbCdEfGh",
				AlphaNumOnly: "Test123",
				NumericOnly:  "123456",
				FixedLength:  "1234567890",
				MinLength:    "12345",
				MaxLength:    "short",
			},
			expectError: false,
			description: "All string validations pass",
		},

		// Invalid cases - required field violations
		{
			name: "invalid_empty_name",
			input: ValidateTestStruct{
				Name:     "",
				Email:    "john@example.com",
				Age:      30,
				Password: "password123",
			},
			expectError: true,
			description: "Name is required but empty",
		},
		{
			name: "invalid_empty_email",
			input: ValidateTestStruct{
				Name:     "John Doe",
				Email:    "",
				Age:      30,
				Password: "password123",
			},
			expectError: true,
			description: "Email is required but empty",
		},
		{
			name: "invalid_zero_age",
			input: ValidateTestStruct{
				Name:     "John Doe",
				Email:    "john@example.com",
				Age:      0,
				Password: "password123",
			},
			expectError: true,
			description: "Age is required but zero",
		},

		// Invalid cases - format violations
		{
			name: "invalid_email_format",
			input: ValidateTestStruct{
				Name:     "John Doe",
				Email:    "invalid-email",
				Age:      30,
				Password: "password123",
			},
			expectError: true,
			description: "Invalid email format",
		},
		{
			name: "invalid_website_format",
			input: ValidateTestStruct{
				Name:     "John Doe",
				Email:    "john@example.com",
				Age:      30,
				Website:  "not-a-url",
				Password: "password123",
			},
			expectError: true,
			description: "Invalid website URL format",
		},

		// Invalid cases - length/range violations
		{
			name: "invalid_name_too_short",
			input: ValidateTestStruct{
				Name:     "J",
				Email:    "john@example.com",
				Age:      30,
				Password: "password123",
			},
			expectError: true,
			description: "Name too short (min=2)",
		},
		{
			name: "invalid_name_too_long",
			input: ValidateTestStruct{
				Name:     "This name is way too long and exceeds the maximum allowed length of 50 characters",
				Email:    "john@example.com",
				Age:      30,
				Password: "password123",
			},
			expectError: true,
			description: "Name too long (max=50)",
		},
		{
			name: "invalid_age_too_young",
			input: ValidateTestStruct{
				Name:     "John Doe",
				Email:    "john@example.com",
				Age:      17,
				Password: "password123",
			},
			expectError: true,
			description: "Age too young (min=18)",
		},
		{
			name: "invalid_age_too_old",
			input: ValidateTestStruct{
				Name:     "John Doe",
				Email:    "john@example.com",
				Age:      121,
				Password: "password123",
			},
			expectError: true,
			description: "Age too old (max=120)",
		},
		{
			name: "invalid_password_too_short",
			input: ValidateTestStruct{
				Name:     "John Doe",
				Email:    "john@example.com",
				Age:      30,
				Password: "short",
			},
			expectError: true,
			description: "Password too short (min=8)",
		},

		// Invalid cases - nested struct violations
		{
			name: "invalid_nested_struct_id",
			input: NestedValidateStruct{
				ID: 0,
				User: ValidateTestStruct{
					Name:     "Jane Doe",
					Email:    "jane@example.com",
					Age:      25,
					Password: "password123",
				},
				Tags:     []string{"tag1"},
				Settings: map[string]string{"key": "value"},
			},
			expectError: true,
			description: "ID must be greater than 0",
		},
		{
			name: "invalid_nested_struct_user",
			input: NestedValidateStruct{
				ID: 1,
				User: ValidateTestStruct{
					Name:     "",
					Email:    "jane@example.com",
					Age:      25,
					Password: "password123",
				},
				Tags:     []string{"tag1"},
				Settings: map[string]string{"key": "value"},
			},
			expectError: true,
			description: "Nested user has invalid name",
		},
		{
			name: "invalid_nested_struct_empty_tags",
			input: NestedValidateStruct{
				ID: 1,
				User: ValidateTestStruct{
					Name:     "Jane Doe",
					Email:    "jane@example.com",
					Age:      25,
					Password: "password123",
				},
				Tags:     []string{},
				Settings: map[string]string{"key": "value"},
			},
			expectError: true,
			description: "Tags slice is empty but required with min=1",
		},

		// Invalid cases - optional field violations
		{
			name: "invalid_optional_field_too_short",
			input: OptionalFieldsStruct{
				RequiredField: "required",
				OptionalField: "abcd",
			},
			expectError: true,
			description: "Optional field provided but too short",
		},
		{
			name: "invalid_optional_email_format",
			input: OptionalFieldsStruct{
				RequiredField: "required",
				OptionalEmail: "invalid-email",
			},
			expectError: true,
			description: "Optional email provided but invalid format",
		},

		// Invalid cases - numeric violations
		{
			name: "invalid_numeric_not_positive",
			input: NumericValidationStruct{
				PositiveInt:    0,
				NonNegativeInt: 0,
				LimitedFloat:   50.5,
				RangeInt:       55,
			},
			expectError: true,
			description: "PositiveInt must be greater than 0",
		},
		{
			name: "invalid_numeric_negative",
			input: NumericValidationStruct{
				PositiveInt:    5,
				NonNegativeInt: -1,
				LimitedFloat:   50.5,
				RangeInt:       55,
			},
			expectError: true,
			description: "NonNegativeInt cannot be negative",
		},
		{
			name: "invalid_float_out_of_range",
			input: NumericValidationStruct{
				PositiveInt:    5,
				NonNegativeInt: 0,
				LimitedFloat:   101.0,
				RangeInt:       55,
			},
			expectError: true,
			description: "LimitedFloat exceeds maximum",
		},
		{
			name: "invalid_int_below_range",
			input: NumericValidationStruct{
				PositiveInt:    5,
				NonNegativeInt: 0,
				LimitedFloat:   50.5,
				RangeInt:       5,
			},
			expectError: true,
			description: "RangeInt below minimum",
		},

		// Invalid cases - string format violations
		{
			name: "invalid_string_non_alpha",
			input: StringValidationStruct{
				AlphaOnly:    "Test123",
				AlphaNumOnly: "Test123",
				NumericOnly:  "123456",
				FixedLength:  "1234567890",
				MinLength:    "12345",
				MaxLength:    "short",
			},
			expectError: true,
			description: "AlphaOnly contains numbers",
		},
		{
			name: "invalid_string_non_alphanum",
			input: StringValidationStruct{
				AlphaOnly:    "TestOnly",
				AlphaNumOnly: "Test123!",
				NumericOnly:  "123456",
				FixedLength:  "1234567890",
				MinLength:    "12345",
				MaxLength:    "short",
			},
			expectError: true,
			description: "AlphaNumOnly contains special characters",
		},
		{
			name: "invalid_string_non_numeric",
			input: StringValidationStruct{
				AlphaOnly:    "TestOnly",
				AlphaNumOnly: "Test123",
				NumericOnly:  "12345a",
				FixedLength:  "1234567890",
				MinLength:    "12345",
				MaxLength:    "short",
			},
			expectError: true,
			description: "NumericOnly contains letters",
		},
		{
			name: "invalid_string_wrong_length",
			input: StringValidationStruct{
				AlphaOnly:    "TestOnly",
				AlphaNumOnly: "Test123",
				NumericOnly:  "123456",
				FixedLength:  "123456789",
				MinLength:    "12345",
				MaxLength:    "short",
			},
			expectError: true,
			description: "FixedLength is not exactly 10 characters",
		},
		{
			name: "invalid_string_too_short",
			input: StringValidationStruct{
				AlphaOnly:    "TestOnly",
				AlphaNumOnly: "Test123",
				NumericOnly:  "123456",
				FixedLength:  "1234567890",
				MinLength:    "1234",
				MaxLength:    "short",
			},
			expectError: true,
			description: "MinLength is less than 5 characters",
		},
		{
			name: "invalid_string_too_long",
			input: StringValidationStruct{
				AlphaOnly:    "TestOnly",
				AlphaNumOnly: "Test123",
				NumericOnly:  "123456",
				FixedLength:  "1234567890",
				MinLength:    "12345",
				MaxLength:    "this string is too long",
			},
			expectError: true,
			description: "MaxLength exceeds 20 characters",
		},

		// Edge cases
		{
			name: "nil_input",
			input: nil,
			expectError: true,
			description: "Nil input should cause error",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := Validate(tt.input)
			if tt.expectError {
				assert.Error(t, err, tt.description)
			} else {
				assert.NoError(t, err, tt.description)
			}
		})
	}
}

// TestValidate_NonStructTypes tests validation with non-struct types
func TestValidate_NonStructTypes(t *testing.T) {
	tests := []struct {
		name        string
		input       interface{}
		expectError bool
	}{
		{
			name:        "string_input",
			input:       "test string",
			expectError: true,
		},
		{
			name:        "int_input",
			input:       42,
			expectError: true,
		},
		{
			name:        "slice_input",
			input:       []string{"a", "b", "c"},
			expectError: true,
		},
		{
			name:        "map_input",
			input:       map[string]int{"key": 1},
			expectError: true,
		},
		{
			name:        "bool_input",
			input:       true,
			expectError: true,
		},
		{
			name:        "float_input",
			input:       3.14,
			expectError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := Validate(tt.input)
			if tt.expectError {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

// TestValidate_StructPointers tests validation with struct pointers
func TestValidate_StructPointers(t *testing.T) {
	t.Run("valid_struct_pointer", func(t *testing.T) {
		input := &ValidateTestStruct{
			Name:     "John Doe",
			Email:    "john@example.com",
			Age:      30,
			Password: "password123",
		}
		err := Validate(input)
		assert.NoError(t, err)
	})

	t.Run("invalid_struct_pointer", func(t *testing.T) {
		input := &ValidateTestStruct{
			Name:     "",
			Email:    "john@example.com",
			Age:      30,
			Password: "password123",
		}
		err := Validate(input)
		assert.Error(t, err)
	})

	t.Run("nil_pointer", func(t *testing.T) {
		var input *ValidateTestStruct
		err := Validate(input)
		assert.Error(t, err)
	})
}

// TestValidate_ComplexValidations tests more complex validation scenarios
func TestValidate_ComplexValidations(t *testing.T) {
	type ComplexStruct struct {
		Email     string   `validate:"required,email"`
		URLs      []string `validate:"required,dive,url"`
		Metadata  map[string]interface{} `validate:"required"`
		Count     int      `validate:"required,gte=1,lte=100"`
	}

	t.Run("valid_complex", func(t *testing.T) {
		input := ComplexStruct{
			Email: "test@example.com",
			URLs:  []string{"https://example.com", "https://test.com"},
			Metadata: map[string]interface{}{
				"key1": "value1",
				"key2": 42,
			},
			Count: 50,
		}
		err := Validate(input)
		assert.NoError(t, err)
	})

	t.Run("invalid_url_in_slice", func(t *testing.T) {
		input := ComplexStruct{
			Email: "test@example.com",
			URLs:  []string{"https://example.com", "not-a-url"},
			Metadata: map[string]interface{}{
				"key1": "value1",
			},
			Count: 50,
		}
		err := Validate(input)
		assert.Error(t, err)
	})
}

// Benchmark test for validation performance
func BenchmarkValidate(b *testing.B) {
	input := ValidateTestStruct{
		Name:     "John Doe",
		Email:    "john@example.com",
		Age:      30,
		Website:  "https://example.com",
		Password: "password123",
	}

	for i := 0; i < b.N; i++ {
		_ = Validate(input)
	}
}

// TestValidate_ErrorTypes tests that errors returned are proper validation errors
func TestValidate_ErrorTypes(t *testing.T) {
	input := ValidateTestStruct{
		Name:     "",
		Email:    "invalid-email",
		Age:      17,
		Password: "short",
	}

	err := Validate(input)
	require.Error(t, err)

	// The error should contain information about multiple validation failures
	errMsg := err.Error()
	assert.Contains(t, errMsg, "Name")
	assert.Contains(t, errMsg, "Email")
	assert.Contains(t, errMsg, "Age")
	assert.Contains(t, errMsg, "Password")
}