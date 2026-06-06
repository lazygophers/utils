package validator

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestStringValidatorsEdgeCases(t *testing.T) {
	v, err := New()
	require.NoError(t, err)

	// alphaspace with empty
	type AS struct{ S string `validate:"alphaspace"` }
	assert.Error(t, v.Struct(AS{S: ""}))
	assert.NoError(t, v.Struct(AS{S: "hello world"}))
	assert.Error(t, v.Struct(AS{S: "hello123"}))

	// alphanumspace with empty
	type ANS struct{ S string `validate:"alphanumspace"` }
	assert.Error(t, v.Struct(ANS{S: ""}))
	assert.NoError(t, v.Struct(ANS{S: "hello 123"}))
	assert.Error(t, v.Struct(ANS{S: "hello@world"}))

	// alphaunicode with empty
	type AU struct{ S string `validate:"alphaunicode"` }
	assert.Error(t, v.Struct(AU{S: ""}))
	assert.NoError(t, v.Struct(AU{S: "你好世界"}))
	assert.Error(t, v.Struct(AU{S: "hello123"}))

	// alphanumunicode with empty
	type ANU struct{ S string `validate:"alphanumunicode"` }
	assert.Error(t, v.Struct(ANU{S: ""}))
	assert.NoError(t, v.Struct(ANU{S: "hello123"}))
	assert.Error(t, v.Struct(ANU{S: "hello@world"}))

	// ascii with empty
	type ASC struct{ S string `validate:"ascii"` }
	assert.Error(t, v.Struct(ASC{S: ""}))
	assert.NoError(t, v.Struct(ASC{S: "hello"}))
	assert.Error(t, v.Struct(ASC{S: "你好"}))

	// printascii with empty
	type PAS struct{ S string `validate:"printascii"` }
	assert.Error(t, v.Struct(PAS{S: ""}))
	assert.NoError(t, v.Struct(PAS{S: "hello"}))
	assert.Error(t, v.Struct(PAS{S: string([]byte{0x1F})}))

	// number with empty
	type NUM struct{ S string `validate:"number"` }
	assert.Error(t, v.Struct(NUM{S: ""}))
	assert.NoError(t, v.Struct(NUM{S: "12345"}))
	assert.Error(t, v.Struct(NUM{S: "12a34"}))
}

func TestStringContainsRuneEdge(t *testing.T) {
	v, err := New()
	require.NoError(t, err)

	// containsrune with empty param → always false
	type CR struct{ S string `validate:"containsrune="` }
	assert.Error(t, v.Struct(CR{S: "hello"}))
}

func TestStringExcludesRuneEdge(t *testing.T) {
	v, err := New()
	require.NoError(t, err)

	// excludesrune with empty param → always true
	type ER struct{ S string `validate:"excludesrune="` }
	assert.NoError(t, v.Struct(ER{S: "hello"}))
}

func TestStringMultibyteEdge(t *testing.T) {
	v, err := New()
	require.NoError(t, err)

	type MB struct{ S string `validate:"multibyte"` }
	assert.NoError(t, v.Struct(MB{S: "你好"}))
	assert.Error(t, v.Struct(MB{S: "hello"}))
}
