package xerror

import (
	"errors"
	"fmt"
	"strings"
	"testing"
)

func TestStackNew(t *testing.T) {
	err := New(1001, "boom")
	if err.Code() != 1001 {
		t.Fatalf("code = %d, want 1001", err.Code())
	}
	if err.Error() != "boom" {
		t.Fatalf("msg = %q, want boom", err.Error())
	}
	frames := err.StackTrace()
	if len(frames) == 0 {
		t.Fatal("StackTrace empty")
	}
	if !strings.Contains(frames[0].Function, "TestStackNew") {
		t.Fatalf("top frame = %q, want caller TestStackNew", frames[0].Function)
	}
}

func TestStackNewf(t *testing.T) {
	err := Newf(2002, "value=%d", 42)
	if err.Error() != "value=42" {
		t.Fatalf("msg = %q", err.Error())
	}
	if err.Code() != 2002 {
		t.Fatalf("code = %d", err.Code())
	}
	if !strings.Contains(err.StackTrace()[0].Function, "TestStackNewf") {
		t.Fatal("top frame not caller")
	}
}

func TestStackWrap(t *testing.T) {
	if Wrap(nil, "x") != nil {
		t.Fatal("Wrap(nil) should return nil")
	}
	root := errors.New("root")
	wrapped := Wrap(root, "layer")
	if wrapped.Error() != "layer" {
		t.Fatalf("msg = %q", wrapped.Error())
	}
	if !errors.Is(wrapped, root) {
		t.Fatal("errors.Is should find root")
	}
	if Cause(wrapped) != root {
		t.Fatal("Cause should resolve to root")
	}
}

func TestStackWrapNoDuplicate(t *testing.T) {
	inner := New(0, "inner")
	innerStack := inner.stack
	wrapped := Wrap(inner, "outer").(*Error)
	if wrapped.stack != innerStack {
		t.Fatal("Wrap should reuse existing *Error stack")
	}
}

func TestStackWrapf(t *testing.T) {
	if Wrapf(nil, "x%d", 1) != nil {
		t.Fatal("Wrapf(nil) should return nil")
	}
	root := errors.New("root")
	wrapped := Wrapf(root, "ctx=%s", "abc")
	if wrapped.Error() != "ctx=abc" {
		t.Fatalf("msg = %q", wrapped.Error())
	}
	if Cause(wrapped) != root {
		t.Fatal("Cause mismatch")
	}
}

func TestStackWithStack(t *testing.T) {
	if WithStack(nil) != nil {
		t.Fatal("WithStack(nil) should return nil")
	}
	root := errors.New("root")
	ws := WithStack(root).(*Error)
	if ws.Error() != "root" {
		t.Fatalf("msg = %q", ws.Error())
	}
	if ws.stack == nil {
		t.Fatal("WithStack should attach stack")
	}
	if !errors.Is(ws, root) {
		t.Fatal("errors.Is should find root")
	}

	withFresh := New(1, "has stack")
	if WithStack(withFresh) != withFresh {
		t.Fatal("WithStack should return same *Error when stack present")
	}
}

func TestStackCause(t *testing.T) {
	if Cause(nil) != nil {
		t.Fatal("Cause(nil) should return nil")
	}
	plain := errors.New("plain")
	if Cause(plain) != plain {
		t.Fatal("Cause(plain) should return plain")
	}
	l1 := errors.New("base")
	l2 := Wrap(l1, "mid")
	l3 := Wrap(l2, "top")
	if Cause(l3) != l1 {
		t.Fatal("Cause should resolve deepest")
	}
}

func TestStackFormatPlus(t *testing.T) {
	root := errors.New("root cause")
	err := Wrap(root, "wrapped")
	out := fmt.Sprintf("%+v", err)
	if !strings.Contains(out, "wrapped") || !strings.Contains(out, "root cause") {
		t.Fatalf("%%+v missing parts: %q", out)
	}
	if !strings.Contains(out, "TestStackFormatPlus") {
		t.Fatalf("%%+v missing caller frame: %q", out)
	}
}
