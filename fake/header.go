package fake

import (
	"fmt"
	"strings"

	"golang.org/x/text/language"
)

// acceptVariants are the three browser-flavoured Accept header strings sampled
// by [Faker.Accept]. The three flavours correspond to Chromium-based browsers
// (Chrome / Edge / Opera / Samsung), Firefox, and Safari respectively.
var acceptVariants = []string{
	"text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.7",
	"text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,*/*;q=0.8",
	"text/html,application/xhtml+xml,application/xml;q=0.9,*/*;q=0.8",
}

// acceptWeights weight-samples acceptVariants in proportion to the 2026
// global browser-engine market share (Chromium ~80 / Firefox ~3 / Safari ~17).
var acceptWeights = []int{80, 3, 17}

// acceptEncodingVariants are the four Accept-Encoding strings sampled by
// [Faker.AcceptEncoding], ordered from richest to leanest.
var acceptEncodingVariants = []string{
	"gzip, deflate, br, zstd",
	"gzip, deflate, br",
	"gzip, deflate",
	"gzip",
}

// acceptEncodingWeights weights the four Accept-Encoding variants. Modern
// browsers (zstd-aware) dominate; legacy single-codec headers tail off.
var acceptEncodingWeights = []int{60, 25, 10, 5}

// refererDomains is the default Referer host pool sampled when the Faker's
// country has no locale-specific override.
var refererDomains = []string{
	"https://www.google.com",
	"https://www.bing.com",
	"https://duckduckgo.com",
	"https://search.brave.com",
	"https://github.com",
	"https://stackoverflow.com",
	"https://www.reddit.com",
	"https://www.youtube.com",
	"https://twitter.com",
	"https://www.linkedin.com",
}

// refererDomainsCN is the CN-locale Referer host pool reflecting the regional
// search / social landscape (Baidu / Sogou / Weibo / Zhihu / Bilibili / ...).
var refererDomainsCN = []string{
	"https://www.baidu.com",
	"https://www.sogou.com",
	"https://cn.bing.com",
	"https://weibo.com",
	"https://www.zhihu.com",
	"https://juejin.cn",
	"https://www.bilibili.com",
	"https://www.douban.com",
}

// refererDomainsJP is the JP-locale Referer host pool covering Yahoo Japan
// and the localised Google / Amazon storefronts.
var refererDomainsJP = []string{
	"https://www.yahoo.co.jp",
	"https://www.google.co.jp",
	"https://www.amazon.co.jp",
}

// Accept returns a browser-flavoured Accept header value sampled by global
// browser-engine market share. The three flavours encode the Chromium /
// Firefox / Safari conventions observed in 2025-2026 traffic captures.
func (f *Faker) Accept() string {
	total := 0
	for _, w := range acceptWeights {
		total += w
	}
	r := f.intN(total)
	for i, w := range acceptWeights {
		if r < w {
			return acceptVariants[i]
		}
		r -= w
	}
	return acceptVariants[0]
}

// AcceptLanguage returns an Accept-Language header value built from the
// Faker's active language tag and the locale's OfficialLangs list. The
// header always leads with the active tag (q implicit), then the active
// tag's base ISO 639 code, then locale official languages, and finally
// appends "en" so that English-only services keep working. Q-values step
// down from 0.9 in 0.1 increments and are floored at 0.1.
func (f *Faker) AcceptLanguage() string {
	primary := f.lang.String()
	if primary == "" || primary == "und" {
		primary = "en-US"
	}
	// Promote base-only tags (e.g. "zh") to BCP-47 region form ("zh-CN") by
	// pairing with the Faker's country. Skip when the tag already carries a
	// region or script subtag.
	if !strings.ContainsAny(primary, "-_") {
		region := f.country.Alpha2()
		if region != "" {
			primary = primary + "-" + region
		}
	}
	type langEntry struct {
		tag string
		q   string
	}
	entries := []langEntry{{tag: primary}}
	seen := map[string]bool{strings.ToLower(primary): true}

	q := 0.9
	addTag := func(s string) {
		key := strings.ToLower(s)
		if s == "" || s == "und" || seen[key] {
			return
		}
		seen[key] = true
		entries = append(entries, langEntry{tag: s, q: fmt.Sprintf("%.1f", q)})
		q -= 0.1
		if q < 0.1 {
			q = 0.1
		}
	}

	base, conf := f.lang.Base()
	if conf != language.No {
		addTag(base.String())
	}
	for _, tag := range f.locale.OfficialLangs {
		addTag(tag.String())
		if tagBase, bc := tag.Base(); bc != language.No {
			addTag(tagBase.String())
		}
	}
	if !seen["en"] && !seen["en-us"] {
		addTag("en")
	}

	var b strings.Builder
	b.WriteString(entries[0].tag)
	for i := 1; i < len(entries); i++ {
		b.WriteByte(',')
		b.WriteString(entries[i].tag)
		b.WriteString(";q=")
		b.WriteString(entries[i].q)
	}
	return b.String()
}

// AcceptEncoding returns an Accept-Encoding header value sampled from a
// weighted pool of modern (zstd) and legacy variants.
func (f *Faker) AcceptEncoding() string {
	total := 0
	for _, w := range acceptEncodingWeights {
		total += w
	}
	r := f.intN(total)
	for i, w := range acceptEncodingWeights {
		if r < w {
			return acceptEncodingVariants[i]
		}
		r -= w
	}
	return acceptEncodingVariants[0]
}

// Referer returns a synthetic Referer URL. The host pool is selected by the
// Faker's country (CN / JP have curated regional pools; others use the
// global default). With even probability the URL is either the bare domain
// or a search-style path appending a random lorem word.
func (f *Faker) Referer() string {
	var pool []string
	switch f.country.Alpha2() {
	case "CN":
		pool = refererDomainsCN
	case "JP":
		pool = refererDomainsJP
	default:
		pool = refererDomains
	}
	domain := f.pickString(pool)
	if f.intN(2) == 0 {
		return domain
	}
	return fmt.Sprintf("%s/search?q=%s", domain, f.Word())
}

// Header returns a populated HTTP request header map covering the most
// commonly inspected fields: User-Agent, Accept, Accept-Language,
// Accept-Encoding, Connection (keep-alive) and Referer. Each call samples
// fresh values; callers that need a stable per-session header should cache
// the returned map.
func (f *Faker) Header() map[string]string {
	return map[string]string{
		"User-Agent":      f.UserAgent(),
		"Accept":          f.Accept(),
		"Accept-Language": f.AcceptLanguage(),
		"Accept-Encoding": f.AcceptEncoding(),
		"Connection":      "keep-alive",
		"Referer":         f.Referer(),
	}
}
