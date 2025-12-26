package validator

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestFieldLevelMethods(t *testing.T) {
	t.Run("Top", func(t *testing.T) {
		type TestStruct struct {
			Name string
		}
		s := TestStruct{Name: "test"}
		fl := &fieldLevel{
			top: reflect.ValueOf(s),
		}
		result := fl.Top()
		assert.True(t, result.IsValid())
	})

	t.Run("Parent", func(t *testing.T) {
		type ParentStruct struct {
			Child string
		}
		p := ParentStruct{Child: "test"}
		fl := &fieldLevel{
			parent: reflect.ValueOf(p),
		}
		result := fl.Parent()
		assert.True(t, result.IsValid())
	})

	t.Run("Field", func(t *testing.T) {
		fl := &fieldLevel{
			field: reflect.ValueOf("test"),
		}
		result := fl.Field()
		assert.True(t, result.IsValid())
		assert.Equal(t, "test", result.String())
	})

	t.Run("FieldName", func(t *testing.T) {
		fl := &fieldLevel{
			fieldName: "test_field",
		}
		result := fl.FieldName()
		assert.Equal(t, "test_field", result)
	})

	t.Run("StructFieldName", func(t *testing.T) {
		fl := &fieldLevel{
			structFieldName: "TestField",
		}
		result := fl.StructFieldName()
		assert.Equal(t, "TestField", result)
	})

	t.Run("Param", func(t *testing.T) {
		fl := &fieldLevel{
			param: "10",
		}
		result := fl.Param()
		assert.Equal(t, "10", result)
	})

	t.Run("GetTag", func(t *testing.T) {
		type TestStruct struct {
			Field string `json:"field_name" validate:"required"`
		}
		s := TestStruct{Field: "test"}
		rt := reflect.TypeOf(s)
		fieldType := rt.Field(0)
		fl := &fieldLevel{
			structField: fieldType,
		}

		jsonTag := fl.GetTag("json")
		assert.Equal(t, "field_name", jsonTag)

		validateTag := fl.GetTag("validate")
		assert.Equal(t, "required", validateTag)

		emptyTag := fl.GetTag("nonexistent")
		assert.Equal(t, "", emptyTag)
	})
}

func TestEngineSetTagName(t *testing.T) {
	e := NewEngine()
	e.SetTagName("custom_tag")
	assert.Equal(t, "custom_tag", e.tagName)
}

func TestEngineSetFieldNameFunc(t *testing.T) {
	e := NewEngine()

	customFunc := func(field reflect.StructField) string {
		return "custom_" + field.Name
	}
	e.SetFieldNameFunc(customFunc)
	assert.NotNil(t, e.fieldNameFunc)

	e.SetFieldNameFunc(nil)
	assert.NotNil(t, e.fieldNameFunc)
}

func TestDefaultFieldNameFunc(t *testing.T) {
	t.Run("with_json_tag", func(t *testing.T) {
		type TestStruct struct {
			FieldName string `json:"field_name"`
		}
		rt := reflect.TypeOf(TestStruct{})
		field := rt.Field(0)
		result := defaultFieldNameFunc(field)
		assert.Equal(t, "field_name", result)
	})

	t.Run("with_json_tag_omitempty", func(t *testing.T) {
		type TestStruct struct {
			FieldName string `json:"field_name,omitempty"`
		}
		rt := reflect.TypeOf(TestStruct{})
		field := rt.Field(0)
		result := defaultFieldNameFunc(field)
		assert.Equal(t, "field_name", result)
	})

	t.Run("with_json_tag_dash", func(t *testing.T) {
		type TestStruct struct {
			FieldName string `json:"-"`
		}
		rt := reflect.TypeOf(TestStruct{})
		field := rt.Field(0)
		result := defaultFieldNameFunc(field)
		assert.Equal(t, "FieldName", result)
	})

	t.Run("without_json_tag", func(t *testing.T) {
		type TestStruct struct {
			FieldName string
		}
		rt := reflect.TypeOf(TestStruct{})
		field := rt.Field(0)
		result := defaultFieldNameFunc(field)
		assert.Equal(t, "FieldName", result)
	})

	t.Run("with_empty_json_tag", func(t *testing.T) {
		type TestStruct struct {
			FieldName string `json:""`
		}
		rt := reflect.TypeOf(TestStruct{})
		field := rt.Field(0)
		result := defaultFieldNameFunc(field)
		assert.Equal(t, "FieldName", result)
	})
}

func TestStructFieldNameFunc(t *testing.T) {
	type TestStruct struct {
		FieldName string
	}
	rt := reflect.TypeOf(TestStruct{})
	field := rt.Field(0)
	result := structFieldNameFunc(field)
	assert.Equal(t, "FieldName", result)
}

func TestEngineWithCustomTagName(t *testing.T) {
	e := NewEngine()
	e.SetTagName("check")

	type TestStruct struct {
		Email string `check:"email"`
	}
	s := TestStruct{Email: "invalid-email"}
	err := e.Struct(s)
	assert.Error(t, err)
}

func TestEngineWithCustomFieldNameFunc(t *testing.T) {
	e := NewEngine()

	e.SetFieldNameFunc(func(field reflect.StructField) string {
		return "custom_" + field.Name
	})

	type TestStruct struct {
		Email string `validate:"email"`
	}
	s := TestStruct{Email: "invalid-email"}
	err := e.Struct(s)
	assert.Error(t, err)

	valErr, ok := err.(ValidationErrors)
	require.True(t, ok)
	assert.True(t, valErr.HasField("custom_Email"))
}

func TestEngineVarWithComplexTags(t *testing.T) {
	e := NewEngine()

	t.Run("multiple_tags_with_params", func(t *testing.T) {
		err := e.Var("test@example.com", "required,email")
		assert.NoError(t, err)
	})

	t.Run("tag_with_spaces", func(t *testing.T) {
		err := e.Var("test", "required , alpha")
		assert.NoError(t, err)
	})

	t.Run("empty_tag", func(t *testing.T) {
		err := e.Var("test", "")
		assert.NoError(t, err)
	})

	t.Run("tag_only_commas", func(t *testing.T) {
		err := e.Var("test", ",,")
		assert.NoError(t, err)
	})
}
