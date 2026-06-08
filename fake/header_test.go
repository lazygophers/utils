package fake_test

import (
	"regexp"
	"strconv"
	"strings"
	"testing"

	"github.com/lazygophers/utils/country"
	"github.com/lazygophers/utils/fake"
)

// knownAcceptVariants mirrors the three browser-flavoured Accept strings
// emitted by Faker.Accept. Kept here in the test so accidental drift in the
// production list surfaces as a hard failure.
var knownAcceptVariants = []string{
	"text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.7",
	"text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,*/*;q=0.8",
	"text/html,application/xhtml+xml,application/xml;q=0.9,*/*;q=0.8",
}

// knownAcceptEncodings mirrors the four Accept-Encoding strings emitted by
// Faker.AcceptEncoding.
var knownAcceptEncodings = []string{
	"gzip, deflate, br, zstd",
	"gzip, deflate, br",
	"gzip, deflate",
	"gzip",
}

// TestAcceptVariant asserts Accept() always returns one of the three known
// browser-flavoured variants.
func TestAcceptVariant(t *testing.T) {
	f := fake.New(country.UnitedStates, fake.WithSeed(101))
	known := make(map[string]bool, len(knownAcceptVariants))
	for _, v := range knownAcceptVariants {
		known[v] = true
	}
	for i := 0; i < 1000; i++ {
		got := f.Accept()
		if !known[got] {
			t.Fatalf("Accept() returned unknown variant: %q", got)
		}
	}
}

type acceptLangCase struct {
	name        string
	c           *country.Country
	mustContain []string
}

// TestAcceptLanguageLocale asserts the language header includes the locale's
// primary tag, a q-weighted base form, and "en" as a tail fallback.
func TestAcceptLanguageLocale(t *testing.T) {
	cases := []acceptLangCase{
		{name: "CN", c: country.China, mustContain: []string{"zh-CN", "zh;q=", "en"}},
		{name: "US", c: country.UnitedStates, mustContain: []string{"en-US"}},
		{name: "JP", c: country.Japan, mustContain: []string{"ja-JP", "ja;q="}},
	}
	qRE := regexp.MustCompile(`;q=([0-9]\.[0-9])`)
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			f := fake.New(tc.c, fake.WithSeed(42))
			got := f.AcceptLanguage()
			for _, sub := range tc.mustContain {
				if !strings.Contains(got, sub) {
					t.Errorf("AcceptLanguage(%s)=%q, missing %q", tc.name, got, sub)
				}
			}
			// All q-values must parse as floats in (0, 1).
			for _, m := range qRE.FindAllStringSubmatch(got, -1) {
				q, err := strconv.ParseFloat(m[1], 64)
				if err != nil {
					t.Errorf("AcceptLanguage(%s) bad q-value %q in %q", tc.name, m[1], got)
				}
				if q <= 0 || q >= 1 {
					t.Errorf("AcceptLanguage(%s) q-value %v out of (0,1) range in %q", tc.name, q, got)
				}
			}
		})
	}
}

// TestAcceptEncodingVariant samples AcceptEncoding repeatedly and asserts
// every result is one of the four known variants. Also confirms the modern
// zstd-aware variant dominates (> 40%).
func TestAcceptEncodingVariant(t *testing.T) {
	f := fake.New(country.UnitedStates, fake.WithSeed(202))
	known := make(map[string]bool, len(knownAcceptEncodings))
	for _, v := range knownAcceptEncodings {
		known[v] = true
	}
	zstdCount := 0
	const samples = 1000
	for i := 0; i < samples; i++ {
		got := f.AcceptEncoding()
		if !known[got] {
			t.Fatalf("AcceptEncoding() returned unknown variant: %q", got)
		}
		if got == "gzip, deflate, br, zstd" {
			zstdCount++
		}
	}
	ratio := float64(zstdCount) / float64(samples)
	if ratio < 0.4 {
		t.Errorf("zstd-aware variant share %.2f%% under 40%% floor", ratio*100)
	}
}

type refererCase struct {
	name string
	c    *country.Country
	any  []string // referer host pool must include at least one of these
}

// TestRefererRegional asserts the Referer URL pool reflects the locale's
// regional search / social landscape.
func TestRefererRegional(t *testing.T) {
	cases := []refererCase{
		{
			name: "CN",
			c:    country.China,
			any:  []string{"baidu.com", "sogou.com", "weibo.com", "zhihu.com", "bilibili.com", "douban.com", "juejin.cn", "cn.bing.com"},
		},
		{
			name: "JP",
			c:    country.Japan,
			any:  []string{"yahoo.co.jp", "google.co.jp", "amazon.co.jp"},
		},
		{
			name: "US",
			c:    country.UnitedStates,
			any:  []string{"google.com", "bing.com", "github.com", "duckduckgo.com", "reddit.com", "stackoverflow.com"},
		},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			f := fake.New(tc.c, fake.WithSeed(303))
			hits := make(map[string]bool, len(tc.any))
			for i := 0; i < 500; i++ {
				ref := f.Referer()
				if ref == "" {
					t.Fatalf("Referer() returned empty string")
				}
				for _, host := range tc.any {
					if strings.Contains(ref, host) {
						hits[host] = true
					}
				}
			}
			if len(hits) == 0 {
				t.Errorf("Referer(%s) never hit any expected host in 500 samples", tc.name)
			}
		})
	}
}

// TestHeaderKeys asserts Header() returns the documented six-key envelope
// with every value non-empty.
func TestHeaderKeys(t *testing.T) {
	f := fake.New(country.UnitedStates, fake.WithSeed(404))
	wantKeys := []string{
		"User-Agent",
		"Accept",
		"Accept-Language",
		"Accept-Encoding",
		"Connection",
		"Referer",
	}
	for i := 0; i < 50; i++ {
		h := f.Header()
		for _, k := range wantKeys {
			v, ok := h[k]
			if !ok {
				t.Fatalf("Header() missing key %q (round %d): %v", k, i, h)
			}
			if v == "" {
				t.Fatalf("Header()[%q] empty (round %d)", k, i)
			}
		}
		if h["Connection"] != "keep-alive" {
			t.Errorf("Header()[Connection]=%q, want keep-alive", h["Connection"])
		}
	}
}
