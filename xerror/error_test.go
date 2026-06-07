package xerror

import (
	"fmt"
	"errors"
	"testing"
)

func TestErrorMessage(t *testing.T) {
	e := NewWithMsg(0, "hello")
	if e.Error() != "hello" {
		t.Fatalf("Error() = %q", e.Error())
	}

	empty := &Error{cause: errors.New("from cause")}
	if empty.Error() != "from cause" {
		t.Fatalf("fallback Error() = %q", empty.Error())
	}

	bare := &Error{}
	if bare.Error() != "" {
		t.Fatalf("bare Error() = %q", bare.Error())
	}
}

func TestErrorUnwrap(t *testing.T) {
	root := errors.New("root")
	e := &Error{msg: "wrap", cause: root}
	if e.Unwrap() != root {
		t.Fatal("Unwrap mismatch")
	}
	if (&Error{}).Unwrap() != nil {
		t.Fatal("Unwrap should be nil without cause")
	}
}

func TestErrorCode(t *testing.T) {
	if NewWithMsg(42, "x").Code() != 42 {
		t.Fatal("Code mismatch")
	}
	if NewWithMsg(0, "x").Code() != 0 {
		t.Fatal("zero code mismatch")
	}
}

func TestErrorMsg(t *testing.T) {
	if NewWithMsg(1, "hello").Msg() != "hello" {
		t.Fatal("Msg should return raw msg")
	}
	bare := &Error{cause: errors.New("c")}
	if bare.Msg() != "" {
		t.Fatal("Msg should not fall back to cause")
	}
}

func TestErrorWrapMethod(t *testing.T) {
	e := NewWithMsg(1, "a")
	root := errors.New("root")
	ret := e.Wrap(root)
	if ret != e {
		t.Fatal("Wrap should return self")
	}
	if e.Unwrap() != root {
		t.Fatal("Wrap should set cause")
	}
	// 覆盖前次 cause
	root2 := errors.New("root2")
	e.Wrap(root2)
	if e.Unwrap() != root2 {
		t.Fatal("Wrap should overwrite cause")
	}
}

func TestWrapsMultiple(t *testing.T) {
	a := errors.New("a")
	b := errors.New("b")
	err := Wraps(a, nil, b)
	if err == nil {
		t.Fatal("expected non-nil")
	}
	if !errors.Is(err, a) || !errors.Is(err, b) {
		t.Fatal("Wraps should chain both")
	}
	if Wraps(nil, nil) != nil {
		t.Fatal("Wraps(nil...) should return nil")
	}
}


func TestCauseStdlibCompat(t *testing.T) {
	// 兼容 stdlib %w 包装
	root := errors.New("root")
	wrapped := fmt.Errorf("layer: %w", root)
	xWrap := Wrap(wrapped, "outer")
	if Cause(xWrap) != root {
		t.Fatalf("Cause should unwind through fmt.Errorf %%w, got %v", Cause(xWrap))
	}
}

func TestWrapStdlibIsAs(t *testing.T) {
	root := errors.New("root")
	wrapped := Wrap(root, "outer")
	if !errors.Is(wrapped, root) {
		t.Fatal("errors.Is should穿透 Wrap")
	}
	var target *Error
	if !errors.As(wrapped, &target) {
		t.Fatal("errors.As should hit *Error")
	}
}

func TestWrapsStdlibIs(t *testing.T) {
	a := errors.New("a")
	b := errors.New("b")
	wrapped := Wraps(a, b)
	if !errors.Is(wrapped, a) || !errors.Is(wrapped, b) {
		t.Fatal("errors.Is should穿透 Wraps 任一子 error")
	}
}

func TestWrapNilReturnsNilInterface(t *testing.T) {
	// 防止 *Error nil 接口非 nil 的经典坑
	if Wrap(nil, "x") != nil {
		t.Fatal("Wrap(nil) must return untyped nil")
	}
	if Wraps(nil, nil, nil) != nil {
		t.Fatal("Wraps all nil must return untyped nil")
	}
}
