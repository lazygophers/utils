package validator

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRequiredWithoutAll(t *testing.T) {
	v, err := New()
	assert.NoError(t, err)

	type S struct {
		A string `validate:"required_without_all=B"`
		B string
	}
	// B empty → A required
	assert.Error(t, v.Struct(S{A: "", B: ""}))
	assert.NoError(t, v.Struct(S{A: "x", B: ""}))
	assert.NoError(t, v.Struct(S{A: "", B: "y"}))

	// Empty param → acts like required
	type S2 struct {
		A string `validate:"required_without_all="`
	}
	assert.Error(t, v.Struct(S2{A: ""}))
	assert.NoError(t, v.Struct(S2{A: "x"}))
}

func TestExcludedIf(t *testing.T) {
	v, err := New()
	assert.NoError(t, err)

	type S struct {
		Status string
		Reason string `validate:"excluded_if=Status=active"`
	}
	assert.NoError(t, v.Struct(S{Status: "active", Reason: ""}))
	assert.Error(t, v.Struct(S{Status: "active", Reason: "because"}))
	assert.NoError(t, v.Struct(S{Status: "inactive", Reason: "because"}))
}

func TestExcludedUnless(t *testing.T) {
	v, err := New()
	assert.NoError(t, err)

	type S struct {
		Role    string
		Details string `validate:"excluded_unless=Role=admin"`
	}
	assert.NoError(t, v.Struct(S{Role: "admin", Details: "info"}))
	assert.Error(t, v.Struct(S{Role: "user", Details: "info"}))
	assert.NoError(t, v.Struct(S{Role: "user", Details: ""}))
}

func TestExcludedWithAll(t *testing.T) {
	v, err := New()
	assert.NoError(t, err)

	// Single field: excluded_with_all=A (comma in param is parsed as tag separator)
	type S struct {
		A string
		C string `validate:"excluded_with_all=A"`
	}
	assert.NoError(t, v.Struct(S{A: "x", C: ""}))
	assert.Error(t, v.Struct(S{A: "x", C: "z"}))
	assert.NoError(t, v.Struct(S{A: "", C: "z"}))
}

func TestExcludedWithout(t *testing.T) {
	v, err := New()
	assert.NoError(t, err)

	type S struct {
		A string
		B string `validate:"excluded_without=A"`
	}
	assert.NoError(t, v.Struct(S{A: "", B: ""}))
	assert.Error(t, v.Struct(S{A: "", B: "x"}))
	assert.NoError(t, v.Struct(S{A: "y", B: "x"}))
}

func TestExcludedWithoutAll(t *testing.T) {
	v, err := New()
	assert.NoError(t, err)

	type S struct {
		A string
		C string `validate:"excluded_without_all=A"`
	}
	assert.NoError(t, v.Struct(S{A: "", C: ""}))
	assert.Error(t, v.Struct(S{A: "", C: "z"}))
	assert.NoError(t, v.Struct(S{A: "x", C: "z"}))
}

func TestSplitFieldListEmpty(t *testing.T) {
	result := splitFieldList("")
	assert.Nil(t, result)
}
