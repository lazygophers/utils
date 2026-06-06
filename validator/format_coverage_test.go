package validator

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// Helper to create a FieldLevel from a string value
func strFL(s string) FieldLevel {
	return paramFL{field: reflect.ValueOf(s)}
}

func strFLP(s string) paramFL {
	return paramFL{field: reflect.ValueOf(s)}
}

// ===== Hash validators =====

func TestValidateMD4(t *testing.T) {
	assert.True(t, validateMD4(strFL("d41d8cd98f00b204e9800998ecf8427e")))
	assert.False(t, validateMD4(strFL("invalid")))
	assert.False(t, validateMD4(strFL("")))
}

func TestValidateRIPEMD128(t *testing.T) {
	assert.True(t, validateRIPEMD128(strFL("d41d8cd98f00b204e9800998ecf8427e")))
	assert.False(t, validateRIPEMD128(strFL("invalid")))
}

func TestValidateRIPEMD160(t *testing.T) {
	assert.True(t, validateRIPEMD160(strFL("da39a3ee5e6b4b0d3255bfef95601890afd80709")))
	assert.False(t, validateRIPEMD160(strFL("invalid")))
}

func TestValidateTiger128(t *testing.T) {
	assert.True(t, validateTiger128(strFL("d41d8cd98f00b204e9800998ecf8427e")))
	assert.False(t, validateTiger128(strFL("invalid")))
}

func TestValidateTiger160(t *testing.T) {
	assert.True(t, validateTiger160(strFL("d41d8cd98f00b204e9800998ecf8427ed41d8cd9")))
	assert.False(t, validateTiger160(strFL("invalid")))
}

func TestValidateTiger192(t *testing.T) {
	assert.True(t, validateTiger192(strFL("d41d8cd98f00b204e9800998ecf8427ed41d8cd98f00b204")))
	assert.False(t, validateTiger192(strFL("invalid")))
}

func TestIsHexStringEdgeCases(t *testing.T) {
	assert.False(t, isHexString("", 32))
	assert.False(t, isHexString("gggggggggggggggggggggggggggggggg", 32))
	assert.True(t, isHexString("abcdef0123456789abcdef0123456789", 32))
}

// ===== UUID variants =====

func TestValidateUUIDRFC4122(t *testing.T) {
	assert.True(t, validateUUIDRFC4122(strFL("12345678-1234-4333-8234-123456789abc")))
	assert.False(t, validateUUIDRFC4122(strFL("12345678-1234-4333-1234-123456789abc"))) // variant not RFC4122
	assert.False(t, validateUUIDRFC4122(strFL("invalid")))
}

func TestValidateUUIDVersion(t *testing.T) {
	assert.True(t, validateUUIDVersion(strFL("12345678-1234-4333-8234-123456789abc"), '4'))
	assert.False(t, validateUUIDVersion(strFL("12345678-1234-3333-8234-123456789abc"), '4'))
}

func TestValidateUUIDVersionRFC4122(t *testing.T) {
	assert.True(t, validateUUIDVersionRFC4122(strFL("12345678-1234-4333-8234-123456789abc"), '4'))
	assert.False(t, validateUUIDVersionRFC4122(strFL("12345678-1234-4333-1234-123456789abc"), '4')) // variant not RFC4122
}

func TestValidateUUID3(t *testing.T) {
	assert.True(t, validateUUID3(strFL("12345678-1234-3333-8234-123456789abc")))
	assert.False(t, validateUUID3(strFL("12345678-1234-4333-8234-123456789abc")))
}

func TestValidateUUID3RFC4122(t *testing.T) {
	assert.True(t, validateUUID3RFC4122(strFL("12345678-1234-3333-8234-123456789abc")))
	assert.False(t, validateUUID3RFC4122(strFL("12345678-1234-3333-1234-123456789abc")))
}

func TestValidateUUID4RFC4122(t *testing.T) {
	assert.True(t, validateUUID4RFC4122(strFL("12345678-1234-4333-8234-123456789abc")))
	assert.False(t, validateUUID4RFC4122(strFL("12345678-1234-4333-1234-123456789abc")))
}

func TestValidateUUID5RFC4122(t *testing.T) {
	assert.True(t, validateUUID5RFC4122(strFL("12345678-1234-5333-8234-123456789abc")))
	assert.False(t, validateUUID5RFC4122(strFL("12345678-1234-5333-1234-123456789abc")))
}

// ===== Base64 edge cases =====

func TestValidateBase64URLEdge(t *testing.T) {
	// Invalid base64url
	assert.False(t, validateBase64URL(strFL("!!!invalid")))
	assert.False(t, validateBase64URL(strFL("")))
}

func TestValidateBase64RawURLEdge(t *testing.T) {
	assert.False(t, validateBase64RawURL(strFL("!!!invalid")))
	assert.False(t, validateBase64RawURL(strFL("")))
}

// ===== Color validators edge cases =====

func TestValidateHexColorEdge(t *testing.T) {
	assert.False(t, validateHexColor(strFL("")))       // empty
	assert.False(t, validateHexColor(strFL("#12")))     // too short
	assert.False(t, validateHexColor(strFL("#GGGGGG"))) // invalid chars
}

func TestValidateHSLAEdge(t *testing.T) {
	assert.False(t, validateHSLA(strFL("")))            // empty
	assert.False(t, validateHSLA(strFL("hsla(0,0%,0%)"))) // missing alpha
	assert.True(t, validateHSLA(strFL("hsla(0,0%,0%,1)")))
}

func TestValidateCMYKEdge(t *testing.T) {
	assert.False(t, validateCMYK(strFL("")))
	assert.True(t, validateCMYK(strFL("cmyk(0%,0%,0%,0%)")))
}

// ===== ISBN/ISSN edge cases =====

func TestValidateISBN10Edge(t *testing.T) {
	assert.False(t, validateISBN10(strFL("")))        // empty
	assert.True(t, validateISBN10(strFL("0306406152"))) // valid
	assert.True(t, validateISBN10(strFL("0-306-40615-2")))
}

func TestValidateISBN13Edge(t *testing.T) {
	assert.False(t, validateISBN13(strFL("")))
	assert.True(t, validateISBN13(strFL("9780306406157")))
}

func TestValidateISBNEdge(t *testing.T) {
	assert.False(t, validateISBN(strFL("")))
	assert.True(t, validateISBN(strFL("0306406152")))
	assert.True(t, validateISBN(strFL("9780306406157")))
}

func TestValidateISSNEdge(t *testing.T) {
	assert.False(t, validateISSN(strFL("")))
	assert.True(t, validateISSN(strFL("0317-8471")))
}

func TestLuhnChecksumEdge(t *testing.T) {
	assert.False(t, validateLuhnChecksum(strFL("0")))    // too few digits
	assert.True(t, validateLuhnChecksum(strFL("79927398713")))
}

// ===== Other format validators =====

func TestValidatePostcodeField(t *testing.T) {
	v, err := New()
	require.NoError(t, err)
	err = v.RegisterValidation("postcodefield", validatePostcodeField)
	require.NoError(t, err)

	type S struct{ P string `validate:"postcodefield"` }
	assert.Error(t, v.Struct(S{P: ""}))
}

func TestValidateJWTEdge(t *testing.T) {
	assert.False(t, validateJWT(strFL("")))        // empty
	assert.False(t, validateJWT(strFL("not.jwt")))  // not 3 parts
}

func TestValidateSpiceDbEdge(t *testing.T) {
	assert.False(t, validateSpiceDb(strFL("")))
}

func TestValidateDatetimeEdge(t *testing.T) {
	assert.False(t, validateDatetime(strFL("")))
}

func TestValidateBICISO93622014(t *testing.T) {
	assert.True(t, validateBICISO93622014(strFL("DEUTDEFF")))
	assert.False(t, validateBICISO93622014(strFL("")))
	assert.False(t, validateBICISO93622014(strFL("AB")))
}

func TestValidateBCP47Strict(t *testing.T) {
	assert.True(t, validateBCP47Strict(strFL("en")))
	assert.True(t, validateBCP47Strict(strFL("zh-CN")))
	assert.False(t, validateBCP47Strict(strFL("")))
}

func TestValidateHexadecimalEdge(t *testing.T) {
	assert.False(t, validateHexadecimal(strFL("")))
	assert.False(t, validateHexadecimal(strFL("gg")))
	assert.True(t, validateHexadecimal(strFL("0123456789abcdef")))
}

// ===== Format Variants (dead code coverage) =====

func makeFieldError() *FieldError {
	return &FieldError{Field: "name", Tag: "required", Param: "5", Value: "test"}
}

func TestFormatMessageBuilder(t *testing.T) {
	err := makeFieldError()
	result := formatMessageBuilder("{field} must be {param}", err)
	assert.Contains(t, result, "name")
}

func TestFormatValueFast(t *testing.T) {
	assert.Equal(t, "42", formatValueFast(42))
	assert.Equal(t, "hello", formatValueFast("hello"))
	assert.Equal(t, "<nil>", formatValueFast(nil))
}

func TestFormatMessageSingleReplace(t *testing.T) {
	err := makeFieldError()
	result := formatMessageSingleReplace("{field}", err)
	assert.Equal(t, "name", result)
}

func TestFormatMessageInlineCheck(t *testing.T) {
	err := makeFieldError()
	result := formatMessageInlineCheck("{field} {tag} {param}", err)
	assert.Contains(t, result, "name")
}

func TestFormatMessageBytesBuffer(t *testing.T) {
	err := makeFieldError()
	result := formatMessageBytesBuffer("{field} {tag}", err)
	assert.Contains(t, result, "name")
}

func TestFormatMessageFastPath(t *testing.T) {
	err := makeFieldError()
	// No placeholders
	result := formatMessageFastPath("plain text", err)
	assert.Equal(t, "plain text", result)
	// With placeholders
	result = formatMessageFastPath("{field} {tag}", err)
	assert.Contains(t, result, "name")
}
