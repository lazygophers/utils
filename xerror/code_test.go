package xerror

import (
	"errors"
	"strings"
	"sync"
	"testing"

	"github.com/lazygophers/utils/language"
	xlanguage "golang.org/x/text/language"
)

// stubLocalizer 是测试用 Localizer，最小实现：按 xlanguage.Tag + key 存翻译，
// 支持命名 args ("name", value) 的 {name} 替换。
type stubLocalizer struct {
	mu      sync.RWMutex
	buckets map[xlanguage.Tag]map[string]string
}

func newStubLocalizer() *stubLocalizer {
	return &stubLocalizer{buckets: map[xlanguage.Tag]map[string]string{}}
}

func (s *stubLocalizer) LocalizeWithLang(tag xlanguage.Tag, key string, args ...any) string {
	s.mu.RLock()
	value, ok := s.buckets[tag][key]
	s.mu.RUnlock()
	if !ok {
		return key
	}
	for i := 0; i+1 < len(args); i += 2 {
		name, isStr := args[i].(string)
		if !isStr {
			continue
		}
		var sub string
		switch x := args[i+1].(type) {
		case string:
			sub = x
		case int:
			sub = itoa(int64(x))
		case int64:
			sub = itoa(x)
		case bool:
			if x {
				sub = "true"
			} else {
				sub = "false"
			}
		}
		value = strings.ReplaceAll(value, "{"+name+"}", sub)
	}
	return value
}

func (s *stubLocalizer) Register(tag xlanguage.Tag, key, value string) {
	s.mu.Lock()
	bucket, ok := s.buckets[tag]
	if !ok {
		bucket = map[string]string{}
		s.buckets[tag] = bucket
	}
	bucket[key] = value
	s.mu.Unlock()
}

func (s *stubLocalizer) RegisterBatch(tag xlanguage.Tag, data map[string]any) {
	for k, v := range data {
		if str, ok := v.(string); ok {
			s.Register(tag, k, str)
		}
	}
}

func itoa(n int64) string {
	if n == 0 {
		return "0"
	}
	neg := n < 0
	if neg {
		n = -n
	}
	buf := [20]byte{}
	i := len(buf)
	for n > 0 {
		i--
		buf[i] = byte('0' + n%10)
		n /= 10
	}
	if neg {
		i--
		buf[i] = '-'
	}
	return string(buf[i:])
}

func TestCode(t *testing.T) {
	got := Code(NewWithMsg(2001, "boom"))
	if got != 2001 {
		t.Fatalf("Code(*Error) = %d, want 2001", got)
	}
	got = Code(errors.New("plain"))
	if got != CodeSystem {
		t.Fatalf("Code(plain) = %d, want CodeSystem(-1)", got)
	}
	got = Code(nil)
	if got != CodeSuccess {
		t.Fatalf("Code(nil) = %d, want CodeSuccess(0)", got)
	}
}

func TestNewWithoutLocalizer(t *testing.T) {
	SetLocalizer(nil)
	err := NewWithMsg(3001, "fallback msg")
	if err.Error() != "fallback msg" {
		t.Fatalf("Error()=%q, want fallback msg", err.Error())
	}
}

func TestNewWithLocalizerHit(t *testing.T) {
	stub := newStubLocalizer()
	stub.Register(xlanguage.Make("en"), "error.3001", "translated")
	SetLocalizer(stub)
	defer SetLocalizer(nil)
	language.Set(language.Make("en"))
	defer language.Del()

	err := NewWithMsg(3001, "fallback")
	if err.Error() != "translated" {
		t.Fatalf("Error()=%q, want translated", err.Error())
	}
}

func TestNewWithLocalizerMissFallback(t *testing.T) {
	stub := newStubLocalizer()
	SetLocalizer(stub)
	defer SetLocalizer(nil)

	err := NewWithMsg(999, "fallback")
	if err.Error() != "fallback" {
		t.Fatalf("Error()=%q, want fallback", err.Error())
	}
}

func TestNewArgsInjected(t *testing.T) {
	stub := newStubLocalizer()
	stub.Register(xlanguage.Make("en"), "error.4001", "user %s denied")
	SetLocalizer(stub)
	defer SetLocalizer(nil)
	language.Set(language.Make("en"))
	defer language.Del()

	err := NewWithMsg(4001, "fallback %s", "alice")
	if err.Error() != "user alice denied" {
		t.Fatalf("Error()=%q, want injected", err.Error())
	}
}

func TestNewEmptyMsgDirectSprint(t *testing.T) {
	SetLocalizer(nil)
	err := New(0, "code=", 42)
	if err.Error() != "code=42" {
		t.Fatalf("Error()=%q, want fmt.Sprint result", err.Error())
	}
}

func TestNewEmptyMsgNoArgs(t *testing.T) {
	SetLocalizer(nil)
	err := New(0)
	if err.Error() != "" {
		t.Fatalf("Error()=%q, want empty", err.Error())
	}
}

func TestNewBoundToConstructionLang(t *testing.T) {
	stub := newStubLocalizer()
	stub.Register(xlanguage.Make("en"), "error.5001", "english")
	stub.Register(xlanguage.Make("zh"), "error.5001", "中文")
	SetLocalizer(stub)
	defer SetLocalizer(nil)

	language.Set(language.Make("en"))
	err := NewWithMsg(5001, "fallback")
	language.Set(language.Make("zh"))
	if err.Error() != "english" {
		t.Fatalf("Error()=%q, want english (bound at construction)", err.Error())
	}
	language.Del()
}

func TestSetKeyPrefix(t *testing.T) {
	original := KeyPrefix()
	defer SetKeyPrefix(original)

	stub := newStubLocalizer()
	stub.Register(xlanguage.Make("en"), "biz.6001", "biz msg")
	SetLocalizer(stub)
	defer SetLocalizer(nil)
	language.Set(language.Make("en"))
	defer language.Del()

	SetKeyPrefix("biz.")
	err := NewWithMsg(6001, "fallback")
	if err.Error() != "biz msg" {
		t.Fatalf("Error()=%q, want biz msg via prefix", err.Error())
	}
}

func TestRegisterPackageDelegatesToLocalizer(t *testing.T) {
	stub := newStubLocalizer()
	SetLocalizer(stub)
	defer SetLocalizer(nil)
	language.Set(language.Make("en"))
	defer language.Del()

	// 包级 Register / RegisterMessage 转发到当前 Localizer
	Register(xlanguage.Make("en"), "error.7000", "via Register")
	RegisterMessage(xlanguage.Make("en"), 7001, "via RegisterMessage")
	RegisterBatch(xlanguage.Make("en"), map[string]any{"error.7002": "via RegisterBatch"})

	if got := NewWithMsg(7000, "f").Error(); got != "via Register" {
		t.Errorf("7000=%q", got)
	}
	if got := NewWithMsg(7001, "f").Error(); got != "via RegisterMessage" {
		t.Errorf("7001=%q", got)
	}
	if got := NewWithMsg(7002, "f").Error(); got != "via RegisterBatch" {
		t.Errorf("7002=%q", got)
	}
}

func TestRegisterNoopWhenNoLocalizer(t *testing.T) {
	SetLocalizer(nil)
	// 不应 panic
	Register(xlanguage.Make("en"), "error.0", "x")
	RegisterMessage(xlanguage.Make("en"), 0, "x")
	RegisterBatch(xlanguage.Make("en"), map[string]any{"error.0": "x"})
}

func TestWrapDefaultsToUnknown(t *testing.T) {
	SetLocalizer(nil)
	root := errors.New("root")
	wrapped := Wrap(root, "context")
	if Code(wrapped) != CodeSystem {
		t.Fatalf("Code(Wrap)=%d, want CodeSystem", Code(wrapped))
	}
}

func TestLocalizerConcurrent(t *testing.T) {
	stub := newStubLocalizer()
	SetLocalizer(stub)
	defer SetLocalizer(nil)
	language.Set(language.Make("en"))
	defer language.Del()

	var wg sync.WaitGroup
	for i := 0; i < 50; i++ {
		wg.Add(2)
		go func(n int) {
			defer wg.Done()
			RegisterMessage(xlanguage.Make("en"), 10000+n, "msg")
		}(i)
		go func(n int) {
			defer wg.Done()
			_ = New(10000+n, "fallback").Error()
		}(i)
	}
	wg.Wait()
}
