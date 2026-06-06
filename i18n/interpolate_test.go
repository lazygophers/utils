package i18n

import (
	"strings"
	"testing"
)

func TestInterpolate_NoPlaceholder(t *testing.T) {
	got := interpolate("hello world", nil)
	if got != "hello world" {
		t.Fatalf("expect identity, got %q", got)
	}
}

func TestInterpolate_NoBraceFastPath(t *testing.T) {
	in := "no brace here"
	got := interpolate(in, []any{"name", "Alice"})
	if got != in {
		t.Fatalf("fast path failed: %q", got)
	}
}

func TestInterpolate_NamedBasic(t *testing.T) {
	got := interpolate("hello {name}", []any{"name", "Alice"})
	if got != "hello Alice" {
		t.Fatalf("got %q", got)
	}
}

func TestInterpolate_NamedMultiple(t *testing.T) {
	got := interpolate("{greet}, {name}!", []any{"greet", "Hi", "name", "Bob"})
	if got != "Hi, Bob!" {
		t.Fatalf("got %q", got)
	}
}

func TestInterpolate_NamedMissing(t *testing.T) {
	got := interpolate("hello {name}", []any{"other", "X"})
	if got != "hello {name}" {
		t.Fatalf("missing should preserve, got %q", got)
	}
}

func TestInterpolate_NamedNonStringValue(t *testing.T) {
	got := interpolate("count={n}", []any{"n", 42})
	if got != "count=42" {
		t.Fatalf("got %q", got)
	}
}

func TestInterpolate_PositionalBasic(t *testing.T) {
	// 三元素奇数 → 位置模式
	got := interpolate("{0} loves {1} (#{2})", []any{"Alice", "Go", 3})
	if got != "Alice loves Go (#3)" {
		t.Fatalf("got %q", got)
	}
}

func TestInterpolate_PositionalOutOfRange(t *testing.T) {
	// 三元素奇数 → 位置模式
	got := interpolate("{0} {1} {2} {3}", []any{"A", "B", "C"})
	if got != "A B C {3}" {
		t.Fatalf("out-of-range should preserve, got %q", got)
	}
}

func TestInterpolate_PositionalNonNumericPreserved(t *testing.T) {
	// 首元素非 string → 位置模式，{foo} 不是数字 → 保留原样
	got := interpolate("{foo} {0}", []any{123, "x"})
	if got != "{foo} 123" {
		t.Fatalf("got %q", got)
	}
}

func TestInterpolate_NamedOddLengthFallsToPositional(t *testing.T) {
	// 长度奇数 → 走位置模式，"hello" 是 args[0]
	got := interpolate("{0}", []any{"hello", 1, 2})
	if got != "hello" {
		t.Fatalf("got %q", got)
	}
}

func TestInterpolate_FirstNotStringIsPositional(t *testing.T) {
	got := interpolate("{0} {1}", []any{1, 2})
	if got != "1 2" {
		t.Fatalf("got %q", got)
	}
}

func TestInterpolate_UnclosedBrace(t *testing.T) {
	got := interpolate("hello {name", []any{"name", "Alice"})
	if got != "hello {name" {
		t.Fatalf("unclosed should preserve, got %q", got)
	}
}

func TestInterpolate_PositionalUnclosedBrace(t *testing.T) {
	// 奇数 args → 位置模式；末尾不闭合占位符须保留
	got := interpolate("a {0 b", []any{1, 2, 3})
	if got != "a {0 b" {
		t.Fatalf("got %q", got)
	}
}

func TestInterpolate_EmptyArgsWithPlaceholder(t *testing.T) {
	got := interpolate("hello {name}", nil)
	if got != "hello {name}" {
		t.Fatalf("got %q", got)
	}
}

func TestInterpolate_EmptyTemplate(t *testing.T) {
	got := interpolate("", []any{"name", "x"})
	if got != "" {
		t.Fatalf("got %q", got)
	}
}

func TestInterpolate_AdjacentPlaceholders(t *testing.T) {
	got := interpolate("{a}{b}", []any{"a", "X", "b", "Y"})
	if got != "XY" {
		t.Fatalf("got %q", got)
	}
}

func TestInterpolate_PositionalMixedTypes(t *testing.T) {
	got := interpolate("[{0}] [{1}] [{2}] [{3}] [{4}] [{5}]", []any{
		1, int64(2), uint(3), true, 3.14, []byte("bb"),
	})
	if !strings.HasPrefix(got, "[1] [2] [3] [true] [3.14] [bb]") {
		t.Fatalf("got %q", got)
	}
}

func TestInterpolate_NamedKeyNotString(t *testing.T) {
	// 长度偶数但 args[0] 非 string → 位置模式
	got := interpolate("{0}", []any{42, "ignored"})
	if got != "42" {
		t.Fatalf("got %q", got)
	}
}

type stringerVal struct{ s string }

func (s stringerVal) String() string { return s.s }

func TestInterpolate_StringerValue(t *testing.T) {
	got := interpolate("v={v}", []any{"v", stringerVal{s: "hello"}})
	if got != "v=hello" {
		t.Fatalf("got %q", got)
	}
}

func TestInterpolate_NumericTypeBranches(t *testing.T) {
	cases := []struct {
		v    any
		want string
	}{
		{int32(7), "7"},
		{uint64(8), "8"},
		{uint32(9), "9"},
		{float32(1.5), "1.5"},
		{false, "false"},
		{struct{ X int }{X: 1}, "{1}"}, // fmt.Fprint fallback
	}
	for _, c := range cases {
		got := interpolate("{0}", []any{c.v})
		if got != c.want {
			t.Fatalf("v=%v got %q want %q", c.v, got, c.want)
		}
	}
}

func TestInterpolate_NonStringKeySkipped(t *testing.T) {
	// 中间出现非 string key 应被跳过，仍找到后面同名 key
	// 长度偶数且首元素为 string → 进入 named；查 "name"
	got := interpolate("{name}", []any{"x", 1, "name", "Alice"})
	if got != "Alice" {
		t.Fatalf("got %q", got)
	}
}

func TestInterpolate_PositionalNegativeIndex(t *testing.T) {
	// strconv.Atoi("-1") 成功但 idx < 0 → 保留原样
	got := interpolate("{-1}", []any{"a"})
	if got != "{-1}" {
		t.Fatalf("got %q", got)
	}
}

func TestInterpolate_ZeroAllocFastPath(t *testing.T) {
	tpl := "this template has no braces at all"
	args := []any{"name", "Alice"}
	allocs := testing.AllocsPerRun(100, func() {
		_ = interpolate(tpl, args)
	})
	if allocs != 0 {
		t.Fatalf("expected 0 allocs on no-brace fast path, got %v", allocs)
	}
}

func TestInterpolate_LookupNamedNonStringKey(t *testing.T) {
	// 第一个 pair 的 key 非 string → 跳过；走 named 因首元素是 string
	got := interpolate("{n}", []any{"k", 1, 99, "v", "n", "ok"})
	// 长度 6 偶数，首元素 "k" string → named；查 "n"，命中第三对
	if got != "ok" {
		t.Fatalf("got %q", got)
	}
}
