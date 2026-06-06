package xerror

import (
	"errors"
	"sync"
	"testing"

	"github.com/lazygophers/utils/language"
)

func TestCode(t *testing.T) {
	if got := Code(New(2001, "boom")); got != 2001 {
		t.Fatalf("Code(*Error) = %d, want 2001", got)
	}
	if got := Code(errors.New("plain")); got != 0 {
		t.Fatalf("Code(plain) = %d, want 0", got)
	}
	if got := Code(New(0, "no code")); got != 0 {
		t.Fatalf("Code(zero) = %d, want 0", got)
	}
}

func TestMessageRegisterAndLocalize(t *testing.T) {
	RegisterMessage(language.Make("en"), 3001, "english msg")
	RegisterMessage(language.Make("zh"), 3001, "中文消息")

	defer language.Del()

	language.Set(language.Make("en"))
	if got := New(3001, "fallback").LocalizedError(); got != "english msg" {
		t.Fatalf("en localized = %q, want %q", got, "english msg")
	}

	language.Set(language.Make("zh"))
	if got := New(3001, "fallback").LocalizedError(); got != "中文消息" {
		t.Fatalf("zh localized = %q, want %q", got, "中文消息")
	}
}

func TestLocalizeFallback(t *testing.T) {
	defer language.Del()
	language.Set(language.Make("en"))
	// 未注册的 code 回退到默认 msg。
	if got := New(999999, "default text").LocalizedError(); got != "default text" {
		t.Fatalf("fallback = %q, want %q", got, "default text")
	}
}

func TestLocalizePreRegisteredEnZh(t *testing.T) {
	defer language.Del()
	language.Set(language.Make("en"))
	if got := New(1001, "x").LocalizedError(); got != "internal server error" {
		t.Fatalf("en 1001 = %q", got)
	}
	language.Set(language.Make("zh"))
	if got := New(1001, "x").LocalizedError(); got != "服务器内部错误" {
		t.Fatalf("zh 1001 = %q", got)
	}
}

func TestMessageConcurrent(t *testing.T) {
	var wg sync.WaitGroup
	for i := 0; i < 50; i++ {
		wg.Add(2)
		go func(n int) {
			defer wg.Done()
			RegisterMessage(language.Make("en"), int64(5000+n), "msg")
		}(i)
		go func(n int) {
			defer wg.Done()
			_ = New(int64(5000+n), "fallback").LocalizedError()
		}(i)
	}
	wg.Wait()
}
