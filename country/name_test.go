package country_test

import (
	"sync"
	"testing"

	xlanguage "golang.org/x/text/language"

	"github.com/lazygophers/utils/country"
	"github.com/lazygophers/utils/language"
)

func TestNameInDirectHit(t *testing.T) {
	cn := country.China
	if got := cn.NameIn(xlanguage.English); got != "China" {
		t.Errorf("en: %q", got)
	}
	if got := cn.NameIn(xlanguage.Chinese); got != "中国" {
		t.Errorf("zh: %q", got)
	}
	if got := cn.OfficialNameIn(xlanguage.English); got != "People's Republic of China" {
		t.Errorf("official en: %q", got)
	}
	if got := cn.OfficialNameIn(xlanguage.Chinese); got != "中华人民共和国" {
		t.Errorf("official zh: %q", got)
	}
	if got := cn.CapitalIn(xlanguage.English); got != "Beijing" {
		t.Errorf("capital en: %q", got)
	}
	if got := cn.CapitalIn(xlanguage.Chinese); got != "北京" {
		t.Errorf("capital zh: %q", got)
	}
}

func TestNameInBaseFallback(t *testing.T) {
	cn := country.China
	zhCN := xlanguage.MustParse("zh-CN")
	if got := cn.NameIn(zhCN); got != "中国" {
		t.Errorf("zh-CN base fallback name: %q", got)
	}
	if got := cn.OfficialNameIn(zhCN); got != "中华人民共和国" {
		t.Errorf("zh-CN base fallback official: %q", got)
	}
	if got := cn.CapitalIn(zhCN); got != "北京" {
		t.Errorf("zh-CN base fallback capital: %q", got)
	}
}

func TestNameInEnglishFallback(t *testing.T) {
	cn := country.China
	// A language with no registered name should fall back to English.
	zu := xlanguage.MustParse("zu") // Zulu — unlikely registered
	if got := cn.NameIn(zu); got != "China" {
		t.Errorf("zu fallback name: %q", got)
	}
	if got := cn.OfficialNameIn(zu); got != "People's Republic of China" {
		t.Errorf("zu fallback official: %q", got)
	}
	if got := cn.CapitalIn(zu); got != "Beijing" {
		t.Errorf("zu fallback capital: %q", got)
	}
}

// Note: Antarctica-based fallback tests removed — Antarctica is not in the
// default build set. The "no-capital" branch is covered by build-tag-gated
// integration in country_all builds.

func TestGoroutineLocalName(t *testing.T) {
	cn := country.China

	// Default (English).
	language.Del()
	if got := cn.Name(); got != "China" {
		t.Errorf("default Name: %q", got)
	}
	if got := cn.OfficialName(); got != "People's Republic of China" {
		t.Errorf("default OfficialName: %q", got)
	}
	if got := cn.Capital(); got != "Beijing" {
		t.Errorf("default Capital: %q", got)
	}

	// Switch to zh.
	language.Set(language.Make("zh"))
	if got := cn.Name(); got != "中国" {
		t.Errorf("zh Name: %q", got)
	}
	if got := cn.OfficialName(); got != "中华人民共和国" {
		t.Errorf("zh OfficialName: %q", got)
	}
	if got := cn.Capital(); got != "北京" {
		t.Errorf("zh Capital: %q", got)
	}

	// Reset.
	language.Del()
	if got := cn.Name(); got != "China" {
		t.Errorf("after Del Name: %q", got)
	}
}

func TestConcurrentGoroutineLocalIsolation(t *testing.T) {
	cn := country.China
	var wg sync.WaitGroup
	errs := make(chan string, 32)

	for i := 0; i < 16; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			if i%2 == 0 {
				language.Set(language.Make("zh"))
				if cn.Name() != "中国" {
					errs <- "zh worker saw non-zh"
				}
			} else {
				language.Set(language.Make("en"))
				if cn.Name() != "China" {
					errs <- "en worker saw non-en"
				}
			}
			language.Del()
		}(i)
	}
	wg.Wait()
	close(errs)
	for e := range errs {
		t.Error(e)
	}
}
