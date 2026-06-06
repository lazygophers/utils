package validator

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestComparisonValidators(t *testing.T) {
	v, err := New()
	assert.NoError(t, err)

	t.Run("gt", func(t *testing.T) {
		assert.NoError(t, v.Var("hello", "gt=4"))
		assert.Error(t, v.Var("hi", "gt=4"))
		assert.NoError(t, v.Var(10, "gt=5"))
		assert.Error(t, v.Var(5, "gt=5"))
		assert.NoError(t, v.Var(3.14, "gt=3.0"))
	})
	t.Run("gte", func(t *testing.T) {
		assert.NoError(t, v.Var("hello", "gte=5"))
		assert.NoError(t, v.Var(5, "gte=5"))
		assert.Error(t, v.Var(4, "gte=5"))
	})
	t.Run("lt", func(t *testing.T) {
		assert.NoError(t, v.Var("hi", "lt=5"))
		assert.Error(t, v.Var("hello", "lt=5"))
		assert.NoError(t, v.Var(3, "lt=5"))
	})
	t.Run("lte", func(t *testing.T) {
		assert.NoError(t, v.Var("hi", "lte=5"))
		assert.NoError(t, v.Var(5, "lte=5"))
		assert.Error(t, v.Var(6, "lte=5"))
	})
	t.Run("eq_ignore_case", func(t *testing.T) {
		assert.NoError(t, v.Var("Hello", "eq_ignore_case=hello"))
		assert.Error(t, v.Var("Hello", "eq_ignore_case=world"))
	})
	t.Run("ne_ignore_case", func(t *testing.T) {
		assert.NoError(t, v.Var("Hello", "ne_ignore_case=world"))
		assert.Error(t, v.Var("Hello", "ne_ignore_case=hello"))
	})
}

func TestNetValidators(t *testing.T) {
	v, err := New()
	assert.NoError(t, err)

	t.Run("ip", func(t *testing.T) {
		assert.NoError(t, v.Var("192.168.1.1", "ip"))
		assert.NoError(t, v.Var("::1", "ip"))
		assert.Error(t, v.Var("not-ip", "ip"))
	})
	t.Run("ipv6", func(t *testing.T) {
		assert.NoError(t, v.Var("::1", "ipv6"))
		assert.NoError(t, v.Var("fe80::1", "ipv6"))
		assert.Error(t, v.Var("192.168.1.1", "ipv6"))
	})
	t.Run("ip_addr", func(t *testing.T) {
		assert.NoError(t, v.Var("10.0.0.1", "ip_addr"))
		assert.Error(t, v.Var("not-ip", "ip_addr"))
	})
	t.Run("ip4_addr", func(t *testing.T) {
		assert.NoError(t, v.Var("10.0.0.1", "ip4_addr"))
		assert.Error(t, v.Var("::1", "ip4_addr"))
	})
	t.Run("ip6_addr", func(t *testing.T) {
		assert.NoError(t, v.Var("::1", "ip6_addr"))
		assert.Error(t, v.Var("10.0.0.1", "ip6_addr"))
	})
	t.Run("cidr", func(t *testing.T) {
		assert.NoError(t, v.Var("192.168.1.0/24", "cidr"))
		assert.NoError(t, v.Var("::1/128", "cidr"))
		assert.Error(t, v.Var("not-cidr", "cidr"))
	})
	t.Run("cidrv4", func(t *testing.T) {
		assert.NoError(t, v.Var("10.0.0.0/8", "cidrv4"))
		assert.Error(t, v.Var("::1/128", "cidrv4"))
	})
	t.Run("cidrv6", func(t *testing.T) {
		assert.NoError(t, v.Var("::1/128", "cidrv6"))
		assert.Error(t, v.Var("10.0.0.0/8", "cidrv6"))
	})
	t.Run("hostname", func(t *testing.T) {
		assert.NoError(t, v.Var("example.com", "hostname"))
		assert.Error(t, v.Var("-invalid.com", "hostname"))
	})
	t.Run("hostname_rfc1123", func(t *testing.T) {
		assert.NoError(t, v.Var("123.example.com", "hostname_rfc1123"))
	})
	t.Run("hostname_port", func(t *testing.T) {
		assert.NoError(t, v.Var("example.com:8080", "hostname_port"))
		assert.Error(t, v.Var("example.com", "hostname_port"))
	})
	t.Run("fqdn", func(t *testing.T) {
		assert.NoError(t, v.Var("www.example.com", "fqdn"))
		assert.Error(t, v.Var("localhost", "fqdn"))
	})
	t.Run("tcp_addr", func(t *testing.T) {
		assert.NoError(t, v.Var("192.168.1.1:8080", "tcp_addr"))
		assert.Error(t, v.Var("invalid:addr", "tcp_addr"))
	})
	t.Run("udp_addr", func(t *testing.T) {
		assert.NoError(t, v.Var("192.168.1.1:53", "udp_addr"))
	})
	t.Run("unix_addr", func(t *testing.T) {
		assert.NoError(t, v.Var("/tmp/socket", "unix_addr"))
	})
	t.Run("uri", func(t *testing.T) {
		assert.NoError(t, v.Var("https://example.com/path", "uri"))
		assert.Error(t, v.Var("not a uri", "uri"))
	})
	t.Run("http_url", func(t *testing.T) {
		assert.NoError(t, v.Var("https://example.com", "http_url"))
		assert.NoError(t, v.Var("http://example.com", "http_url"))
		assert.Error(t, v.Var("ftp://example.com", "http_url"))
	})
	t.Run("url_encoded", func(t *testing.T) {
		assert.NoError(t, v.Var("hello%20world", "url_encoded"))
		assert.Error(t, v.Var("hello world", "url_encoded"))
	})
	t.Run("datauri", func(t *testing.T) {
		assert.NoError(t, v.Var("data:text/plain;base64,SGVsbG8=", "datauri"))
		assert.Error(t, v.Var("not-data", "datauri"))
	})
	t.Run("urn_rfc2141", func(t *testing.T) {
		assert.NoError(t, v.Var("urn:isbn:0451450523", "urn_rfc2141"))
		assert.Error(t, v.Var("not-a-urn", "urn_rfc2141"))
	})
}

func TestFSValidators(t *testing.T) {
	v, err := New()
	assert.NoError(t, err)

	// 创建临时文件和目录用于测试
	tmpDir := t.TempDir()
	tmpFile, _ := os.CreateTemp(tmpDir, "testfile")
	tmpFile.Close()

	t.Run("dir", func(t *testing.T) {
		assert.NoError(t, v.Var(tmpDir, "dir"))
		assert.Error(t, v.Var(tmpFile.Name(), "dir"))
		assert.Error(t, v.Var("/nonexistent/path", "dir"))
	})
	t.Run("dirpath", func(t *testing.T) {
		assert.NoError(t, v.Var("/usr/local/bin", "dirpath"))
		assert.NoError(t, v.Var("relative/path", "dirpath"))
	})
	t.Run("file", func(t *testing.T) {
		assert.NoError(t, v.Var(tmpFile.Name(), "file"))
		assert.Error(t, v.Var(tmpDir, "file"))
	})
	t.Run("filepath", func(t *testing.T) {
		assert.NoError(t, v.Var("/tmp/test.txt", "filepath"))
	})
	t.Run("image", func(t *testing.T) {
		assert.NoError(t, v.Var("photo.jpg", "image"))
		assert.NoError(t, v.Var("icon.png", "image"))
		assert.NoError(t, v.Var("anim.gif", "image"))
		assert.Error(t, v.Var("doc.pdf", "image"))
	})
}

func TestMiscValidators(t *testing.T) {
	v, err := New()
	assert.NoError(t, err)

	t.Run("oneof", func(t *testing.T) {
		assert.NoError(t, v.Var("red", "oneof=red,green,blue"))
		assert.Error(t, v.Var("yellow", "oneof=red,green,blue"))
	})
	t.Run("unique", func(t *testing.T) {
		assert.NoError(t, v.Var([]string{"a", "b", "c"}, "unique"))
		assert.Error(t, v.Var([]string{"a", "b", "a"}, "unique"))
	})
	t.Run("isdefault", func(t *testing.T) {
		assert.NoError(t, v.Var("", "isdefault"))
		assert.NoError(t, v.Var(0, "isdefault"))
		assert.Error(t, v.Var("hello", "isdefault"))
	})
}

func TestConditionalValidators(t *testing.T) {
	type TestStruct struct {
		Name    string `validate:"required_unless=Role=admin"`
		Role    string
		Email   string `validate:"required_with_all=FirstName,LastName"`
		FirstName string
		LastName  string
	}
	v, err := New()
	assert.NoError(t, err)

	t.Run("required_unless pass", func(t *testing.T) {
		s := TestStruct{Name: "test", Role: "user"}
		assert.NoError(t, v.Struct(s))
	})
	t.Run("required_unless fail", func(t *testing.T) {
		s := TestStruct{Name: "", Role: "user"}
		assert.Error(t, v.Struct(s))
	})
	t.Run("required_unless skip when admin", func(t *testing.T) {
		s := TestStruct{Name: "", Role: "admin"}
		assert.NoError(t, v.Struct(s))
	})

	t.Run("excluded_with", func(t *testing.T) {
		type Excl struct {
			A string `validate:"excluded_with=B"`
			B string
		}
		assert.NoError(t, v.Struct(Excl{A: "", B: "val"}))    // A 空，OK
		assert.Error(t, v.Struct(Excl{A: "x", B: "val"}))     // B 有值时 A 必须空
		assert.NoError(t, v.Struct(Excl{A: "x", B: ""}))       // B 无值，A 可任意
	})
}

func TestFieldValidators(t *testing.T) {
	type CompareStruct struct {
		Min int `validate:"gtefield=Max"`
		Max int
	}
	v, err := New()
	assert.NoError(t, err)

	t.Run("gtefield pass", func(t *testing.T) {
		s := CompareStruct{Min: 10, Max: 5}
		assert.NoError(t, v.Struct(s))
	})
	t.Run("gtefield fail", func(t *testing.T) {
		s := CompareStruct{Min: 3, Max: 5}
		assert.Error(t, v.Struct(s))
	})

	t.Run("gtfield", func(t *testing.T) {
		type S struct{ A, B int `validate:"gtfield=B"` }
		type S2 struct {
			A int `validate:"gtfield=B"`
			B int
		}
		assert.NoError(t, v.Struct(S2{A: 10, B: 5}))
		assert.Error(t, v.Struct(S2{A: 5, B: 5}))
	})

	t.Run("fieldcontains", func(t *testing.T) {
		assert.NoError(t, v.Var("hello@world", "fieldcontains=@"))
		assert.Error(t, v.Var("hello world", "fieldcontains=@"))
	})
	t.Run("fieldexcludes", func(t *testing.T) {
		assert.NoError(t, v.Var("hello world", "fieldexcludes=@"))
		assert.Error(t, v.Var("hello@world", "fieldexcludes=@"))
	})

	t.Run("eqcsfield", func(t *testing.T) {
		type Inner struct{ Val int }
		type Outer struct {
			A int `validate:"eqcsfield=Inner.Val"`
			Inner
		}
		assert.NoError(t, v.Struct(Outer{A: 42, Inner: Inner{Val: 42}}))
		assert.Error(t, v.Struct(Outer{A: 1, Inner: Inner{Val: 42}}))
	})
}

func TestAliasValidators(t *testing.T) {
	v, err := New()
	assert.NoError(t, err)

	t.Run("iscolor", func(t *testing.T) {
		assert.NoError(t, v.Var("#ff0000", "iscolor"))
		assert.NoError(t, v.Var("rgb(255,0,0)", "iscolor"))
		assert.Error(t, v.Var("not-a-color", "iscolor"))
	})
	t.Run("country_code", func(t *testing.T) {
		assert.NoError(t, v.Var("US", "country_code"))
		assert.NoError(t, v.Var("USA", "country_code"))
		assert.Error(t, v.Var("X", "country_code"))
	})
}
