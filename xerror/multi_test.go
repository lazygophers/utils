package xerror

import (
	"errors"
	"strings"
	"sync"
	"testing"
)

var (
	errFoo = errors.New("foo")
	errBar = errors.New("bar")
	errBaz = errors.New("baz")
)

func TestJoinAllNil(t *testing.T) {
	if err := Join(); err != nil {
		t.Fatalf("Join() = %v, want nil", err)
	}
	if err := Join(nil, nil, nil); err != nil {
		t.Fatalf("Join(nil...) = %v, want nil", err)
	}
}

func TestJoinSingle(t *testing.T) {
	if err := Join(nil, errFoo, nil); err != errFoo {
		t.Fatalf("Join single = %v, want errFoo (原 error)", err)
	}
}

func TestJoinMultiple(t *testing.T) {
	err := Join(errFoo, errBar, errBaz)
	if got := err.Error(); got != "foo\nbar\nbaz" {
		t.Fatalf("Error() = %q, want %q", got, "foo\nbar\nbaz")
	}
}

func TestMultiErrorSingleString(t *testing.T) {
	m := &multiError{errs: []error{errFoo}}
	if got := m.Error(); got != "foo" {
		t.Fatalf("Error() = %q, want foo", got)
	}
}

func TestJoinIs(t *testing.T) {
	err := Join(errFoo, errBar)
	if !errors.Is(err, errFoo) {
		t.Fatal("errors.Is(err, errFoo) = false, want true")
	}
	if !errors.Is(err, errBar) {
		t.Fatal("errors.Is(err, errBar) = false, want true")
	}
	if errors.Is(err, errBaz) {
		t.Fatal("errors.Is(err, errBaz) = true, want false")
	}
}

type customError struct{ tag string }

func (e *customError) Error() string { return "custom:" + e.tag }

func TestJoinAs(t *testing.T) {
	c := &customError{tag: "x"}
	err := Join(errFoo, c)
	var target *customError
	if !errors.As(err, &target) {
		t.Fatal("errors.As = false, want true")
	}
	if target.tag != "x" {
		t.Fatalf("target.tag = %q, want x", target.tag)
	}
}

func TestAppendToMulti(t *testing.T) {
	base := Join(errFoo, errBar)
	got := Append(base, errBaz)
	if !errors.Is(got, errBaz) {
		t.Fatal("Append 未追加 errBaz")
	}
	if strings.Count(got.Error(), "\n") != 2 {
		t.Fatalf("Error() = %q, want 3 行", got.Error())
	}
}

func TestAppendToMultiSameInstance(t *testing.T) {
	base := Join(errFoo, errBar)
	got := Append(base, errBaz)
	if got != base {
		t.Fatal("Append 到 multiError 应原地扩展返回同实例")
	}
}

func TestAppendToPlain(t *testing.T) {
	got := Append(errFoo, errBar)
	if !errors.Is(got, errFoo) || !errors.Is(got, errBar) {
		t.Fatalf("Append plain = %v, 应同时命中 foo/bar", got)
	}
}

func TestAppendNilFiltered(t *testing.T) {
	base := Join(errFoo, errBar)
	got := Append(base, nil, errBaz, nil)
	if got.(*multiError).Len() != 3 {
		t.Fatalf("len = %d, want 3", got.(*multiError).Len())
	}
}

// Len 辅助便于测试断言 multiError 长度。
func (m *multiError) Len() int { return len(m.errs) }

func TestCollectorBasic(t *testing.T) {
	var c Collector
	if c.ErrorOrNil() != nil {
		t.Fatal("空 Collector ErrorOrNil 应为 nil")
	}
	c.Add(nil)
	if c.Len() != 0 {
		t.Fatalf("Add(nil) 后 Len = %d, want 0", c.Len())
	}
	c.Add(errFoo)
	c.Add(errBar)
	if c.Len() != 2 {
		t.Fatalf("Len = %d, want 2", c.Len())
	}
	err := c.ErrorOrNil()
	if !errors.Is(err, errFoo) || !errors.Is(err, errBar) {
		t.Fatalf("ErrorOrNil = %v, 应命中 foo/bar", err)
	}
}

func TestCollectorSingle(t *testing.T) {
	var c Collector
	c.Add(errFoo)
	if err := c.ErrorOrNil(); err != errFoo {
		t.Fatalf("单错 ErrorOrNil = %v, want 原 errFoo", err)
	}
}

func TestCollectorConcurrent(t *testing.T) {
	var c Collector
	const n = 100
	var wg sync.WaitGroup
	wg.Add(n)
	for i := 0; i < n; i++ {
		go func() {
			defer wg.Done()
			c.Add(errFoo)
		}()
	}
	wg.Wait()
	if c.Len() != n {
		t.Fatalf("并发 Add 后 Len = %d, want %d", c.Len(), n)
	}
}
