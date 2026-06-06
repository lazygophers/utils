package xerror

import (
	"errors"
	"fmt"
	"testing"
)

func TestErrorMessage(t *testing.T) {
	e := New(0, "hello")
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

func TestErrorIs(t *testing.T) {
	a := New(100, "a")
	b := New(100, "b different msg")
	c := New(200, "c")

	if !errors.Is(a, b) {
		t.Fatal("same code should be Is-equal")
	}
	if errors.Is(a, c) {
		t.Fatal("different code should not match")
	}

	zero := New(0, "no code")
	if errors.Is(zero, New(0, "other")) {
		t.Fatal("code 0 should never Is-match")
	}
	if a.Is(errors.New("plain")) {
		t.Fatal("non-*Error target should not match")
	}
}

func TestErrorCode(t *testing.T) {
	if New(42, "x").Code() != 42 {
		t.Fatal("Code mismatch")
	}
	if New(0, "x").Code() != 0 {
		t.Fatal("zero code mismatch")
	}
}

func TestErrorWithMetadata(t *testing.T) {
	e := New(1, "m")
	if e.meta != nil {
		t.Fatal("meta should be nil before WithMetadata")
	}
	ret := e.WithMetadata("k1", "v1").WithMetadata("k2", "v2")
	if ret != e {
		t.Fatal("WithMetadata should return self")
	}
	if e.meta["k1"] != "v1" || e.meta["k2"] != "v2" {
		t.Fatalf("meta = %v", e.meta)
	}
}

func TestErrorFormat(t *testing.T) {
	e := New(1, "msg")
	if got := fmt.Sprintf("%v", e); got != "msg" {
		t.Fatalf("%%v = %q", got)
	}
	if got := fmt.Sprintf("%s", e); got != "msg" {
		t.Fatalf("%%s = %q", got)
	}
	if got := fmt.Sprintf("%q", e); got != `"msg"` {
		t.Fatalf("%%q = %q", got)
	}
}
