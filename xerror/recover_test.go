package xerror

import (
	"errors"
	"strings"
	"testing"
)

func TestTryNoPanic(t *testing.T) {
	called := false
	err := Try(func() { called = true })
	if err != nil {
		t.Fatalf("Try 正常执行应返回 nil，got %v", err)
	}
	if !called {
		t.Fatal("fn 未被调用")
	}
}

func TestTryPanicError(t *testing.T) {
	base := errors.New("boom")
	err := Try(func() { panic(base) })
	if err == nil {
		t.Fatal("Try 捕获 panic 应返回非 nil error")
	}
	var xe *Error
	if !errors.As(err, &xe) {
		t.Fatalf("应为 *Error，got %T", err)
	}
	if xe.StackTrace() == nil {
		t.Fatal("应携带堆栈")
	}
	if !errors.Is(err, base) {
		t.Fatal("应保留原始 error 链")
	}
}

func TestTryPanicNonError(t *testing.T) {
	err := Try(func() { panic("oops") })
	if err == nil {
		t.Fatal("Try 捕获 panic 应返回非 nil error")
	}
	var xe *Error
	if !errors.As(err, &xe) {
		t.Fatalf("应为 *Error，got %T", err)
	}
	if xe.StackTrace() == nil {
		t.Fatal("应携带堆栈")
	}
	if !strings.Contains(err.Error(), "oops") {
		t.Fatalf("消息应含 panic 值，got %q", err.Error())
	}
}

func TestTryPanicInt(t *testing.T) {
	err := Try(func() { panic(42) })
	if err == nil {
		t.Fatal("应返回 error")
	}
	if !strings.Contains(err.Error(), "42") {
		t.Fatalf("消息应含 panic 值，got %q", err.Error())
	}
}

func TestTryEPassthrough(t *testing.T) {
	base := errors.New("normal failure")
	err := TryE(func() error { return base })
	if err != base {
		t.Fatalf("TryE 应透传 fn 返回的 error，got %v", err)
	}
}

func TestTryENoError(t *testing.T) {
	err := TryE(func() error { return nil })
	if err != nil {
		t.Fatalf("应返回 nil，got %v", err)
	}
}

func TestTryEPanic(t *testing.T) {
	err := TryE(func() error { panic("panic in TryE") })
	if err == nil {
		t.Fatal("panic 应被转为 error")
	}
	var xe *Error
	if !errors.As(err, &xe) {
		t.Fatalf("应为 *Error，got %T", err)
	}
	if !strings.Contains(err.Error(), "panic in TryE") {
		t.Fatalf("消息应含 panic 值，got %q", err.Error())
	}
}

func TestRecoverWriteback(t *testing.T) {
	err := func() (err error) {
		defer Recover(&err)
		panic(errors.New("recovered"))
	}()
	if err == nil {
		t.Fatal("Recover 应回写 panic")
	}
	var xe *Error
	if !errors.As(err, &xe) {
		t.Fatalf("应为 *Error，got %T", err)
	}
	if xe.StackTrace() == nil {
		t.Fatal("应携带堆栈")
	}
}

func TestRecoverNoPanic(t *testing.T) {
	err := func() (err error) {
		defer Recover(&err)
		return nil
	}()
	if err != nil {
		t.Fatalf("无 panic 时不应改动 errp，got %v", err)
	}
}

func TestRecoverNonError(t *testing.T) {
	err := func() (err error) {
		defer Recover(&err)
		panic(3.14)
	}()
	if err == nil {
		t.Fatal("应回写 error")
	}
	if !strings.Contains(err.Error(), "3.14") {
		t.Fatalf("消息应含 panic 值，got %q", err.Error())
	}
}
