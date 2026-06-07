package currency_test

import (
	"sync"
	"testing"

	xlanguage "golang.org/x/text/language"

	"github.com/lazygophers/utils/currency"
	"github.com/lazygophers/utils/language"
)

type getCase struct {
	name string
	in   string
	want *currency.Currency
}

func TestGet(t *testing.T) {
	cases := []getCase{
		{"upper", "CNY", currency.Cny},
		{"lower", "cny", currency.Cny},
		{"empty", "", nil},
		{"too-short", "CN", nil},
		{"too-long", "CNYY", nil},
		{"unknown", "ZZZ", nil},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			got := currency.Get(c.in)
			if got != c.want {
				t.Fatalf("Get(%q) = %v, want %v", c.in, got, c.want)
			}
		})
	}
}

func TestGetByNumeric(t *testing.T) {
	if got := currency.GetByNumeric(156); got != currency.Cny {
		t.Fatalf("GetByNumeric(156) = %v, want Cny", got)
	}
	if got := currency.GetByNumeric(840); got != currency.Usd {
		t.Fatalf("GetByNumeric(840) = %v, want Usd", got)
	}
	if got := currency.GetByNumeric(0); got != nil {
		t.Errorf("GetByNumeric(0) should be nil")
	}
	if got := currency.GetByNumeric(99999); got != nil {
		t.Errorf("GetByNumeric(99999) should be nil")
	}
}

func TestListLength(t *testing.T) {
	if got := len(currency.List()); got < 11 {
		t.Fatalf("List length = %d, want >= 11 (default set)", got)
	}
}

func TestFieldAccessors(t *testing.T) {
	c := currency.Cny
	if c.Code() != "CNY" {
		t.Errorf("Code: %q", c.Code())
	}
	if c.Symbol() != "¥" {
		t.Errorf("Symbol: %q", c.Symbol())
	}
	if c.Numeric() != 156 {
		t.Errorf("Numeric: %d", c.Numeric())
	}
	if c.String() != "CNY" {
		t.Errorf("String: %q", c.String())
	}
	if c.Decimals() != 2 {
		t.Errorf("Cny.Decimals: %d", c.Decimals())
	}
	if !c.Reserve() {
		t.Errorf("Cny.Reserve should be true")
	}
	if len(c.Banknotes()) == 0 {
		t.Errorf("Cny.Banknotes empty")
	}
	if len(c.Coins()) == 0 {
		t.Errorf("Cny.Coins empty")
	}

	jpy := currency.Jpy
	if jpy.Decimals() != 0 {
		t.Errorf("Jpy.Decimals: %d", jpy.Decimals())
	}
	if !jpy.Reserve() {
		t.Errorf("Jpy.Reserve should be true")
	}
}

func TestConstantsMatchLookup(t *testing.T) {
	if currency.Cny != currency.Get("CNY") {
		t.Error("Cny != Get(CNY)")
	}
	if currency.Usd != currency.Get("USD") {
		t.Error("Usd != Get(USD)")
	}
}

func TestNoDuplicateCodes(t *testing.T) {
	codes := make(map[string]bool, 154)
	nums := make(map[int]bool, 154)
	for _, c := range currency.List() {
		if codes[c.Code()] {
			t.Errorf("duplicate code: %s", c.Code())
		}
		codes[c.Code()] = true
		if nums[c.Numeric()] {
			t.Errorf("duplicate numeric: %d", c.Numeric())
		}
		nums[c.Numeric()] = true
	}
}

func TestNameInDirectHit(t *testing.T) {
	cny := currency.Cny
	if got := cny.NameIn(xlanguage.English); got != "Yuan Renminbi" {
		t.Errorf("en: %q", got)
	}
}

func TestNameInBaseFallback(t *testing.T) {
	cny := currency.Cny
	zhCN := xlanguage.MustParse("zh-CN")
	// zh-CN base = zh. Whether zh registered depends on cny_zh.go contents.
	got := cny.NameIn(zhCN)
	// Should never be empty — base hit, English fallback, or code.
	if got == "" {
		t.Error("empty fallback result")
	}
}

func TestNameInUnknownLangFallsBackToEnglish(t *testing.T) {
	cny := currency.Cny
	zu := xlanguage.MustParse("zu")
	if got := cny.NameIn(zu); got != "Yuan Renminbi" {
		t.Errorf("zu fallback: %q", got)
	}
}

func TestNameInFallsBackToCode(t *testing.T) {
	// Create a fresh currency with no names; should fall back to code.
	type sentinelTag struct{}
	c := currency.New("ZZX", "?", 999900)
	if got := c.NameIn(xlanguage.English); got != "ZZX" {
		t.Errorf("expected code fallback, got %q", got)
	}
	if got := c.Name(); got != "ZZX" {
		t.Errorf("Name() expected code fallback, got %q", got)
	}
	_ = sentinelTag{}
}

func TestGoroutineLocalName(t *testing.T) {
	cny := currency.Cny

	language.Del()
	if got := cny.Name(); got != "Yuan Renminbi" {
		t.Errorf("default Name: %q", got)
	}

	// Switch to a language without a registered name — should fall back to en.
	language.Set(language.Make("zu"))
	if got := cny.Name(); got != "Yuan Renminbi" {
		t.Errorf("zu Name fallback: %q", got)
	}
	language.Del()
}

func TestRegisterNameAndLookup(t *testing.T) {
	c := currency.New("ZZY", "?", 999901)
	c.RegisterName(xlanguage.MustParse("fr"), "monnaie test")
	if got := c.NameIn(xlanguage.MustParse("fr")); got != "monnaie test" {
		t.Errorf("fr: %q", got)
	}
	// fr-CA should fall back to fr via base.
	if got := c.NameIn(xlanguage.MustParse("fr-CA")); got != "monnaie test" {
		t.Errorf("fr-CA base fallback: %q", got)
	}
	// Other language — no English registered → code.
	if got := c.NameIn(xlanguage.MustParse("de")); got != "ZZY" {
		t.Errorf("de fallback to code: %q", got)
	}
}

func TestConcurrentGoroutineLocalIsolation(t *testing.T) {
	cny := currency.Cny
	var wg sync.WaitGroup
	errs := make(chan string, 32)
	for i := 0; i < 16; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			language.Set(language.Make("en"))
			if cny.Name() != "Yuan Renminbi" {
				errs <- "wrong en name"
			}
			language.Del()
		}()
	}
	wg.Wait()
	close(errs)
	for e := range errs {
		t.Error(e)
	}
}
