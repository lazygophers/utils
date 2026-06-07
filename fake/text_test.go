package fake_test

import (
	"strings"
	"testing"
	"unicode/utf8"

	"github.com/lazygophers/utils/country"
	"github.com/lazygophers/utils/fake"
)

func TestWord_NonEmpty(t *testing.T) {
	f := fake.New(country.UnitedStates, fake.WithSeed(1))
	for i := 0; i < 20; i++ {
		w := f.Word()
		if w == "" {
			t.Fatal("Word empty")
		}
	}
}

func TestWords(t *testing.T) {
	f := fake.New(country.UnitedStates, fake.WithSeed(2))
	got := f.Words(5)
	if len(got) != 5 {
		t.Fatalf("Words(5) returned %d", len(got))
	}
	for _, w := range got {
		if w == "" {
			t.Fatal("empty word")
		}
	}
	empty := f.Words(0)
	if len(empty) != 0 {
		t.Fatalf("Words(0) returned %d", len(empty))
	}
	if f.Words(-1) == nil || len(f.Words(-1)) != 0 {
		t.Fatal("Words(-1) should return empty slice")
	}
}

func TestSentence(t *testing.T) {
	f := fake.New(country.UnitedStates, fake.WithSeed(3))
	// Explicit word count.
	s := f.Sentence(5)
	if !strings.HasSuffix(s, ".") {
		t.Fatalf("Sentence missing period: %q", s)
	}
	words := strings.Fields(strings.TrimSuffix(s, "."))
	if len(words) != 5 {
		t.Fatalf("expected 5 words, got %d: %q", len(words), s)
	}
	// First letter must be upper.
	if s[0] < 'A' || s[0] > 'Z' {
		t.Fatalf("Sentence not capitalized: %q", s)
	}
	// Default count.
	s2 := f.Sentence(0)
	if !strings.HasSuffix(s2, ".") {
		t.Fatalf("default sentence missing period: %q", s2)
	}
}

func TestParagraph(t *testing.T) {
	f := fake.New(country.UnitedStates, fake.WithSeed(4))
	p := f.Paragraph(3)
	count := strings.Count(p, ".")
	if count != 3 {
		t.Fatalf("Paragraph expected 3 periods, got %d: %q", count, p)
	}
	p2 := f.Paragraph(0)
	if !strings.Contains(p2, ".") {
		t.Fatalf("default Paragraph missing periods: %q", p2)
	}
}

func TestChineseWord(t *testing.T) {
	f := fake.New(country.China, fake.WithSeed(5))
	w := f.ChineseWord()
	if utf8.RuneCountInString(w) < 1 {
		t.Fatalf("ChineseWord empty: %q", w)
	}
}

func TestChineseWords(t *testing.T) {
	f := fake.New(country.China, fake.WithSeed(6))
	got := f.ChineseWords(4)
	if len(got) != 4 {
		t.Fatalf("ChineseWords(4) returned %d", len(got))
	}
	if len(f.ChineseWords(0)) != 0 {
		t.Fatal("ChineseWords(0) should return empty")
	}
	if len(f.ChineseWords(-3)) != 0 {
		t.Fatal("ChineseWords(-3) should return empty")
	}
}

func TestChineseSentence(t *testing.T) {
	f := fake.New(country.China, fake.WithSeed(7))
	s := f.ChineseSentence(10)
	if !strings.HasSuffix(s, "。") {
		t.Fatalf("ChineseSentence missing 。: %q", s)
	}
	chars := utf8.RuneCountInString(strings.TrimSuffix(s, "。"))
	if chars < 10 {
		t.Fatalf("ChineseSentence too short: %d chars", chars)
	}
	// Default branch.
	s2 := f.ChineseSentence(0)
	if !strings.HasSuffix(s2, "。") {
		t.Fatalf("default ChineseSentence missing 。: %q", s2)
	}
}

func TestChineseParagraph(t *testing.T) {
	f := fake.New(country.China, fake.WithSeed(8))
	p := f.ChineseParagraph(3)
	count := strings.Count(p, "。")
	if count != 3 {
		t.Fatalf("ChineseParagraph expected 3 periods, got %d: %q", count, p)
	}
	p2 := f.ChineseParagraph(0)
	if !strings.Contains(p2, "。") {
		t.Fatalf("default ChineseParagraph missing 。: %q", p2)
	}
}
