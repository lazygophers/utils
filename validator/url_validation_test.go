package validator

import (
	"reflect"
	"testing"
)

func TestValidateURL(t *testing.T) {
	validURLs := []string{
		"http://example.com",
		"https://example.com",
		"https://www.example.com",
		"http://sub.domain.example.com",
		"https://example.com/path",
		"https://example.com/path/to/resource",
		"https://example.com/path?query=value",
		"https://example.com/path?query=value&other=123",
		"https://example.com/path#fragment",
		"ftp://ftp.example.com",
		"ws://websocket.example.com",
		"wss://secure.websocket.example.com",
		"http://localhost",
		"http://localhost:8080",
		"https://192.168.1.1",
		"http://example.com:8080/path?query=value#fragment",
		"https://example.co.uk",
		"http://example.io",
		"https://api.example.com/v1/users",
		"http://example.com/path/with/segments",
	}

	invalidURLs := []string{
		"",
		"not a url",
		"example.com",
		"http://",
		"https://",
		"htt://example.com",
		"http:/example.com",
		"//example.com",
		" http://example.com",
		"http://example.com ",
		"\nhttp://example.com",
		"http:// example.com",
		"http://example .com",
		"http://example.com/ path",
		"http://example.com\tpath",
		"mailto:test@example.com",
		"javascript:void(0)",
	}

	// 测试有效 URL
	for _, u := range validURLs {
		fl := &mockFieldLevel{field: reflect.ValueOf(u)}
		if !validateURL(fl) {
			t.Errorf("有效 URL 被拒绝: %q", u)
		}
	}

	// 测试无效 URL
	for _, u := range invalidURLs {
		fl := &mockFieldLevel{field: reflect.ValueOf(u)}
		if validateURL(fl) {
			t.Errorf("无效 URL 被接受: %q", u)
		}
	}
}
