package validator

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestStrongPasswordEdgeCases(t *testing.T) {
	v, err := New()
	require.NoError(t, err)

	type P struct{ Pass string `validate:"strong_password"` }

	// Only 1 type
	assert.Error(t, v.Struct(P{Pass: "ABCDEFGH"}))    // only upper
	assert.Error(t, v.Struct(P{Pass: "abcdefgh"}))    // only lower
	assert.Error(t, v.Struct(P{Pass: "12345678"}))    // only numbers
	// 2 types
	assert.Error(t, v.Struct(P{Pass: "Abcdefgh"}))    // upper+lower = 2
	assert.Error(t, v.Struct(P{Pass: "abcdefg1"}))    // lower+number = 2
	// 3 types → pass
	assert.NoError(t, v.Struct(P{Pass: "Abcdefg1"}))  // upper+lower+number = 3
	assert.NoError(t, v.Struct(P{Pass: "Abcdefg!"}))  // upper+lower+special = 3
	// Too short
	assert.Error(t, v.Struct(P{Pass: "Ab1!"}))        // too short
}

func TestValidateURLEdgeCases(t *testing.T) {
	v, err := New()
	require.NoError(t, err)

	type U struct{ URL string `validate:"url"` }
	assert.Error(t, v.Struct(U{URL: ""}))
	assert.Error(t, v.Struct(U{URL: "://missing-scheme"}))
	assert.NoError(t, v.Struct(U{URL: "https://example.com"}))
}

func TestValidateIPv4EdgeCases(t *testing.T) {
	v, err := New()
	require.NoError(t, err)

	type IP struct{ Addr string `validate:"ipv4"` }
	assert.Error(t, v.Struct(IP{Addr: ""}))
	assert.Error(t, v.Struct(IP{Addr: "256.0.0.1"}))
	assert.Error(t, v.Struct(IP{Addr: "1.2.3.4.5"}))
	assert.NoError(t, v.Struct(IP{Addr: "192.168.1.1"}))
}

func TestValidateUUIDEdgeCases(t *testing.T) {
	v, err := New()
	require.NoError(t, err)

	type ID struct{ Uid string `validate:"uuid"` }
	assert.Error(t, v.Struct(ID{Uid: ""}))
	assert.Error(t, v.Struct(ID{Uid: "not-a-uuid"}))
	assert.NoError(t, v.Struct(ID{Uid: "12345678-1234-1234-1234-123456789abc"}))
}
