package i18n

import (
	"strconv"
	"sync"
	"testing"
)

func TestMapStoreGetSet(t *testing.T) {
	s := newMapStore()

	if _, ok := s.Get("en", "hello"); ok {
		t.Fatal("expected miss before Set")
	}

	s.Set("en", "hello", "Hello")
	s.Set("zh", "hello", "你好")

	if v, ok := s.Get("en", "hello"); !ok || v != "Hello" {
		t.Fatalf("en.hello = %q,%v", v, ok)
	}
	if v, ok := s.Get("zh", "hello"); !ok || v != "你好" {
		t.Fatalf("zh.hello = %q,%v", v, ok)
	}
	if _, ok := s.Get("en", "missing"); ok {
		t.Fatal("expected miss for unknown key")
	}
}

func TestMapStoreOverwrite(t *testing.T) {
	s := newMapStore()
	s.Set("en", "k", "v1")
	s.Set("en", "k", "v2")
	if v, _ := s.Get("en", "k"); v != "v2" {
		t.Fatalf("overwrite failed, got %q", v)
	}
}

func TestMapStoreConcurrent(t *testing.T) {
	s := newMapStore()
	const n = 1000
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		defer wg.Done()
		for i := 0; i < n; i++ {
			s.Set("en", strconv.Itoa(i), strconv.Itoa(i))
		}
	}()
	go func() {
		defer wg.Done()
		for i := 0; i < n; i++ {
			_, _ = s.Get("en", strconv.Itoa(i))
		}
	}()
	wg.Wait()
}

func TestCompositeKey(t *testing.T) {
	if got := compositeKey("en", "x"); got != "en.x" {
		t.Fatalf("compositeKey = %q", got)
	}
}

// TestStoreEquivalence 双实现等价性：同序列 Set/Get 结果必须一致。
func TestStoreEquivalence(t *testing.T) {
	m := newMapStore()
	c := newChunkStore(0)

	// 显式遍历以保证插入顺序对两实现一致。
	pairs := []struct {
		locale, key, text string
	}{
		{"en", "hello", "Hello"},
		{"zh", "hello", "你好"},
		{"en", "bye", "Bye"},
		{"zh", "bye", "再见"},
		{"en", "n.items", "items"},
		{"zh-CN", "greet", "您好"},
	}
	for _, p := range pairs {
		m.Set(p.locale, p.key, p.text)
		c.Set(p.locale, p.key, p.text)
	}

	for _, p := range pairs {
		mv, mok := m.Get(p.locale, p.key)
		cv, cok := c.Get(p.locale, p.key)
		if mok != cok || mv != cv {
			t.Fatalf("mismatch for %s.%s: map=(%q,%v) chunk=(%q,%v)", p.locale, p.key, mv, mok, cv, cok)
		}
	}

	// miss 等价
	for _, p := range []struct{ l, k string }{{"en", "unknown"}, {"fr", "hello"}} {
		_, mok := m.Get(p.l, p.k)
		_, cok := c.Get(p.l, p.k)
		if mok != cok {
			t.Fatalf("miss mismatch for %s.%s: map=%v chunk=%v", p.l, p.k, mok, cok)
		}
	}
}
