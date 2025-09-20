package validator

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestFieldLevelMethods tests the FieldLevel interface methods
func TestFieldLevelMethods(t *testing.T) {
	t.Run("FieldLevel interface coverage", func(t *testing.T) {
		engine := NewEngine()

		// Register a custom validator that uses all FieldLevel methods
		engine.RegisterValidation("test_field_level", func(fl FieldLevel) bool {
			// Test Top() method
			top := fl.Top()
			if !top.IsValid() {
				return false
			}

			// Test Parent() method
			parent := fl.Parent()
			if !parent.IsValid() {
				return false
			}

			// Test FieldName() method
			fieldName := fl.FieldName()
			if fieldName == "" {
				return false
			}

			// Test StructFieldName() method
			structFieldName := fl.StructFieldName()
			if structFieldName == "" {
				return false
			}

			// Test GetTag() method
			jsonTag := fl.GetTag("json")
			// This should return something (even empty string is valid)
			_ = jsonTag

			return true
		})

		// Test struct with the custom validator
		type TestStruct struct {
			Name string `json:"name" validate:"test_field_level"`
		}

		testObj := TestStruct{Name: "test"}
		err := engine.Struct(testObj)
		assert.NoError(t, err)
	})

	t.Run("FieldLevel with nested struct", func(t *testing.T) {
		engine := NewEngine()

		// Register a validator that tests nested structure access
		engine.RegisterValidation("test_nested", func(fl FieldLevel) bool {
			// Test accessing parent and top in nested context
			top := fl.Top()
			parent := fl.Parent()
			field := fl.Field()

			// Verify these are all valid reflect.Value instances
			return top.IsValid() && parent.IsValid() && field.IsValid()
		})

		type NestedStruct struct {
			Value string `validate:"test_nested"`
		}

		type ParentStruct struct {
			Nested NestedStruct
		}

		testObj := ParentStruct{
			Nested: NestedStruct{Value: "test"},
		}

		err := engine.Struct(testObj)
		assert.NoError(t, err)
	})

	t.Run("FieldLevel parameter access", func(t *testing.T) {
		engine := NewEngine()

		// Register a validator that uses parameter
		engine.RegisterValidation("test_param", func(fl FieldLevel) bool {
			param := fl.Param()
			// Test that we can access the parameter
			return param != ""
		})

		type TestStruct struct {
			Name string `validate:"test_param=somevalue"`
		}

		testObj := TestStruct{Name: "test"}
		err := engine.Struct(testObj)
		assert.NoError(t, err)
	})

	t.Run("FieldLevel tag access", func(t *testing.T) {
		engine := NewEngine()

		// Register a validator that reads various tags
		engine.RegisterValidation("test_tags", func(fl FieldLevel) bool {
			// Test GetTag with different tag names
			jsonTag := fl.GetTag("json")
			validateTag := fl.GetTag("validate")
			customTag := fl.GetTag("custom")

			// All of these should be accessible (even if empty)
			_ = jsonTag
			_ = validateTag
			_ = customTag

			return true
		})

		type TestStruct struct {
			Name string `json:"name" validate:"test_tags" custom:"value"`
		}

		testObj := TestStruct{Name: "test"}
		err := engine.Struct(testObj)
		assert.NoError(t, err)
	})
}

// TestCustomValidatorEngine tests the custom validation engine more thoroughly
func TestCustomValidatorEngine(t *testing.T) {
	t.Run("Custom validator registration and usage", func(t *testing.T) {
		engine := NewEngine()

		// Test that we can register and use custom validators
		engine.RegisterValidation("always_true", func(fl FieldLevel) bool {
			return true
		})

		engine.RegisterValidation("always_false", func(fl FieldLevel) bool {
			return false
		})

		type TestStruct struct {
			ValidField   string `validate:"always_true"`
			InvalidField string `validate:"always_false"`
		}

		testObj := TestStruct{
			ValidField:   "valid",
			InvalidField: "invalid",
		}

		err := engine.Struct(testObj)
		assert.Error(t, err) // Should fail because of always_false validator

		// Test with only valid field
		type ValidStruct struct {
			ValidField string `validate:"always_true"`
		}

		validObj := ValidStruct{ValidField: "valid"}
		err2 := engine.Struct(validObj)
		assert.NoError(t, err2)
	})

	t.Run("Engine with field name function", func(t *testing.T) {
		engine := NewEngine()

		// Set a custom field name function
		engine.SetFieldNameFunc(func(field reflect.StructField) string {
			return "custom_" + field.Name
		})

		// Register a validator that uses the field name
		engine.RegisterValidation("check_field_name", func(fl FieldLevel) bool {
			fieldName := fl.FieldName()
			return fieldName != ""
		})

		type TestStruct struct {
			TestField string `validate:"check_field_name"`
		}

		testObj := TestStruct{TestField: "test"}
		err := engine.Struct(testObj)
		assert.NoError(t, err)
	})
}