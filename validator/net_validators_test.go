package validator

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValidateCIDRv4(t *testing.T) {
	v, err := New()
	assert.NoError(t, err)

	type S struct{ IP string `validate:"cidrv4"` }
	assert.NoError(t, v.Struct(S{IP: "192.168.1.0/24"}))
	assert.Error(t, v.Struct(S{IP: "::1/128"}))
	assert.Error(t, v.Struct(S{IP: "invalid"}))
}

func TestValidateCIDRv6(t *testing.T) {
	v, err := New()
	assert.NoError(t, err)

	type S struct{ IP string `validate:"cidrv6"` }
	assert.NoError(t, v.Struct(S{IP: "::1/128"}))
	assert.NoError(t, v.Struct(S{IP: "2001:db8::/32"}))
	assert.Error(t, v.Struct(S{IP: "192.168.1.0/24"}))
	assert.Error(t, v.Struct(S{IP: "invalid"}))
}

func TestValidateHostname(t *testing.T) {
	v, err := New()
	assert.NoError(t, err)

	type S struct{ H string `validate:"hostname"` }
	assert.NoError(t, v.Struct(S{H: "example.com"}))
	assert.Error(t, v.Struct(S{H: ""}))
	assert.Error(t, v.Struct(S{H: "-invalid.com"}))
	assert.NoError(t, v.Struct(S{H: "a.b"}))
}

func TestValidateHostnameRFC1123(t *testing.T) {
	v, err := New()
	assert.NoError(t, err)

	type S struct{ H string `validate:"hostname_rfc1123"` }
	assert.NoError(t, v.Struct(S{H: "1host.example.com"}))
	assert.NoError(t, v.Struct(S{H: "example.com"}))
	assert.Error(t, v.Struct(S{H: ""}))
	assert.Error(t, v.Struct(S{H: "-invalid.com"}))
}

func TestValidateHostnamePort(t *testing.T) {
	v, err := New()
	assert.NoError(t, err)

	type S struct{ H string `validate:"hostname_port"` }
	assert.NoError(t, v.Struct(S{H: "example.com:8080"}))
	assert.NoError(t, v.Struct(S{H: "192.168.1.1:443"}))
	assert.Error(t, v.Struct(S{H: "example.com:"}))
	assert.Error(t, v.Struct(S{H: ":8080"}))
	assert.Error(t, v.Struct(S{H: "example.com:abc"}))
	assert.Error(t, v.Struct(S{H: "noport"}))
}

func TestValidateFQDN(t *testing.T) {
	v, err := New()
	assert.NoError(t, err)

	type S struct{ H string `validate:"fqdn"` }
	assert.NoError(t, v.Struct(S{H: "www.example.com"}))
	assert.NoError(t, v.Struct(S{H: "example.org"}))
	assert.Error(t, v.Struct(S{H: "localhost"}))
	assert.Error(t, v.Struct(S{H: ""}))
	assert.Error(t, v.Struct(S{H: "example.1"}))
}

func TestValidateTCP4Addr(t *testing.T) {
	v, err := New()
	assert.NoError(t, err)

	type S struct{ A string `validate:"tcp4_addr"` }
	assert.NoError(t, v.Struct(S{A: "192.168.1.1:80"}))
	assert.Error(t, v.Struct(S{A: "invalid"}))
}

func TestValidateTCP6Addr(t *testing.T) {
	v, err := New()
	assert.NoError(t, err)

	type S struct{ A string `validate:"tcp6_addr"` }
	assert.NoError(t, v.Struct(S{A: "[::1]:80"}))
	assert.Error(t, v.Struct(S{A: "invalid"}))
}

func TestValidateUDPAddr(t *testing.T) {
	v, err := New()
	assert.NoError(t, err)

	type S struct{ A string `validate:"udp_addr"` }
	assert.NoError(t, v.Struct(S{A: "192.168.1.1:53"}))
	assert.Error(t, v.Struct(S{A: "invalid"}))
}

func TestValidateUDP4Addr(t *testing.T) {
	v, err := New()
	assert.NoError(t, err)

	type S struct{ A string `validate:"udp4_addr"` }
	assert.NoError(t, v.Struct(S{A: "192.168.1.1:53"}))
	assert.Error(t, v.Struct(S{A: "invalid"}))
}

func TestValidateUDP6Addr(t *testing.T) {
	v, err := New()
	assert.NoError(t, err)

	type S struct{ A string `validate:"udp6_addr"` }
	assert.NoError(t, v.Struct(S{A: "[::1]:53"}))
	assert.Error(t, v.Struct(S{A: "invalid"}))
}

func TestValidateHTTPURL(t *testing.T) {
	v, err := New()
	assert.NoError(t, err)

	type S struct{ U string `validate:"http_url"` }
	assert.NoError(t, v.Struct(S{U: "http://example.com"}))
	assert.NoError(t, v.Struct(S{U: "HTTPS://example.com/path"}))
	assert.Error(t, v.Struct(S{U: "ftp://example.com"}))
	assert.Error(t, v.Struct(S{U: "not-a-url"}))
}

func TestValidateURLEncoded(t *testing.T) {
	v, err := New()
	assert.NoError(t, err)

	type S struct{ U string `validate:"url_encoded"` }
	assert.NoError(t, v.Struct(S{U: "hello%20world"}))
	assert.NoError(t, v.Struct(S{U: "abc-123_def.txt"}))
	assert.Error(t, v.Struct(S{U: "hello world"}))
	assert.Error(t, v.Struct(S{U: ""}))
}

func TestValidateDataURI(t *testing.T) {
	v, err := New()
	assert.NoError(t, err)

	type S struct{ U string `validate:"datauri"` }
	assert.NoError(t, v.Struct(S{U: "data:text/plain;base64,SGVsbG8="}))
	assert.NoError(t, v.Struct(S{U: "data:,Hello"}))
	assert.Error(t, v.Struct(S{U: "not-a-data-uri"}))
	assert.Error(t, v.Struct(S{U: "data:text/plain"})) // no comma
}
