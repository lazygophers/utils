package i18n

import (
	"sync"
	"testing"

	"github.com/lazygophers/utils/language"
)

func TestBundleNewDefault(t *testing.T) {
	b := New()
	if b.store == nil {
		t.Fatal("store nil")
	}
	if _, ok := b.store.(*mapStore); !ok {
		t.Fatalf("default store should be mapStore, got %T", b.store)
	}
	if b.defaultLocale.String() != "en" {
		t.Fatalf("default locale = %s, want en", b.defaultLocale.String())
	}
}

func TestBundleNewLowMemory(t *testing.T) {
	b := New(WithLowMemory())
	if _, ok := b.store.(*chunkStore); !ok {
		t.Fatalf("lowmem store should be chunkStore, got %T", b.store)
	}
	b2 := New(WithLowMemory(), WithMaxBytes(1<<20))
	if _, ok := b2.store.(*chunkStore); !ok {
		t.Fatalf("lowmem+maxbytes store should be chunkStore, got %T", b2.store)
	}
	// WithMaxBytes without WithLowMemory still mapStore
	b3 := New(WithMaxBytes(1 << 20))
	if _, ok := b3.store.(*mapStore); !ok {
		t.Fatalf("maxbytes only should remain mapStore, got %T", b3.store)
	}
}

func TestBundleRegisterAndT(t *testing.T) {
	b := New()
	en := language.Make("en")
	zh := language.Make("zh")
	b.Register(en, "hello", "Hello")
	b.Register(zh, "hello", "你好")

	language.Set(en)
	defer language.Del()
	if got := b.T("hello"); got != "Hello" {
		t.Fatalf("en T = %q, want Hello", got)
	}
	language.Set(zh)
	if got := b.T("hello"); got != "你好" {
		t.Fatalf("zh T = %q, want 你好", got)
	}
}

func TestBundleRegisterMap(t *testing.T) {
	b := New()
	zh := language.Make("zh")
	b.RegisterMap(zh, map[string]string{
		"a": "甲",
		"b": "乙",
	})
	if got := b.TLocale(zh, "a"); got != "甲" {
		t.Fatalf("a = %q", got)
	}
	if got := b.TLocale(zh, "b"); got != "乙" {
		t.Fatalf("b = %q", got)
	}
}

func TestBundleFallbackChain(t *testing.T) {
	b := New()
	zh := language.Make("zh")
	b.Register(zh, "k", "中文")
	// zh-CN not registered, should fallback to zh
	zhCN := language.Make("zh-CN")
	if got := b.TLocale(zhCN, "k"); got != "中文" {
		t.Fatalf("fallback zh-CN -> zh = %q, want 中文", got)
	}
}

func TestBundleFallbackToDefaultEn(t *testing.T) {
	b := New()
	en := language.Make("en")
	b.Register(en, "only_en", "fallback")
	// Query under fr; fr chain doesn't include en, must fallback to default en.
	fr := language.Make("fr")
	if got := b.TLocale(fr, "only_en"); got != "fallback" {
		t.Fatalf("fallback to default en = %q", got)
	}
}

func TestBundleMissReturnsKey(t *testing.T) {
	b := New()
	en := language.Make("en")
	if got := b.TLocale(en, "absent"); got != "absent" {
		t.Fatalf("miss should return key, got %q", got)
	}
}

func TestBundleInterpolate(t *testing.T) {
	b := New()
	en := language.Make("en")
	b.Register(en, "greet", "Hello {name}")
	b.Register(en, "pos", "Hello {0}")

	if got := b.TLocale(en, "greet", "name", "Alice"); got != "Hello Alice" {
		t.Fatalf("named = %q", got)
	}
	if got := b.TLocale(en, "pos", "Bob"); got != "Hello Bob" {
		t.Fatalf("positional = %q", got)
	}
	// no args
	if got := b.TLocale(en, "greet"); got != "Hello {name}" {
		t.Fatalf("no args = %q", got)
	}
}

func TestBundleLowMemoryEquivalence(t *testing.T) {
	bMap := New()
	bChunk := New(WithLowMemory(), WithMaxBytes(64*1024))
	en := language.Make("en")
	zh := language.Make("zh")
	pairs := map[string]string{
		"hello":   "Hello",
		"bye":     "Bye",
		"welcome": "Welcome {name}",
	}
	bMap.RegisterMap(en, pairs)
	bChunk.RegisterMap(en, pairs)
	bMap.Register(zh, "hello", "你好")
	bChunk.Register(zh, "hello", "你好")

	for k := range pairs {
		if a, b := bMap.TLocale(en, k), bChunk.TLocale(en, k); a != b {
			t.Fatalf("mismatch key %s: map=%q chunk=%q", k, a, b)
		}
	}
	if a, b := bMap.TLocale(en, "welcome", "name", "X"), bChunk.TLocale(en, "welcome", "name", "X"); a != b {
		t.Fatalf("interpolate mismatch: %q vs %q", a, b)
	}
	if a, b := bMap.TLocale(zh, "hello"), bChunk.TLocale(zh, "hello"); a != b {
		t.Fatalf("zh mismatch: %q vs %q", a, b)
	}
}

func TestRegistryPackageLevel(t *testing.T) {
	en := language.Make("en")
	zh := language.Make("zh")
	Register(en, "pkg_hi", "Hi")
	RegisterMap(zh, map[string]string{"pkg_hi": "嗨"})

	language.Set(en)
	defer language.Del()
	if got := T("pkg_hi"); got != "Hi" {
		t.Fatalf("pkg T en = %q", got)
	}
	if got := TLocale(zh, "pkg_hi"); got != "嗨" {
		t.Fatalf("pkg TLocale zh = %q", got)
	}
	if got := T("pkg_missing"); got != "pkg_missing" {
		t.Fatalf("pkg miss = %q", got)
	}
}

func TestRegistryConcurrent(t *testing.T) {
	b := New()
	en := language.Make("en")
	b.Register(en, "k", "v")

	var wg sync.WaitGroup
	for i := 0; i < 32; i++ {
		wg.Add(2)
		go func(i int) {
			defer wg.Done()
			key := "key" + string(rune('a'+i%26))
			b.Register(en, key, "text")
		}(i)
		go func() {
			defer wg.Done()
			_ = b.TLocale(en, "k")
		}()
	}
	wg.Wait()
}

func TestBundleFallbackChunkStore(t *testing.T) {
	b := New(WithLowMemory())
	zh := language.Make("zh")
	b.Register(zh, "k", "中文")
	zhCN := language.Make("zh-CN")
	if got := b.TLocale(zhCN, "k"); got != "中文" {
		t.Fatalf("chunk fallback = %q", got)
	}
}
