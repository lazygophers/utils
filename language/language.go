package language

import (
	"sort"
	"strconv"
	"strings"
	"sync"

	xlanguage "golang.org/x/text/language"
)

// Tag represents a parsed language tag with extended metadata.
// It wraps golang.org/x/text/language.Tag and can be converted back via [Tag.Tag].
type Tag struct {
	underlying xlanguage.Tag
	weight     float64 // q value from Accept-Language header, 0 if not from header
}

// tagCache interns weight-less Tag pointers keyed by canonical BCP 47 string.
// Parent/FallbackChain hit this to skip per-call allocations.
// ParseAcceptLanguage bypasses the cache because each entry carries a per-request weight.
var tagCache sync.Map // map[string]*Tag

// makeCache short-circuits Make/Parse on the raw input string before paying for xlanguage.Make.
// Different input strings can canonicalize to the same Tag — both keys still point to the same pointer.
var makeCache sync.Map // map[string]*Tag

// intern returns the canonical *Tag for the given underlying tag.
func intern(t xlanguage.Tag) *Tag {
	key := t.String()
	v, ok := tagCache.Load(key)
	if ok {
		return v.(*Tag)
	}
	actual, _ := tagCache.LoadOrStore(key, &Tag{underlying: t})
	return actual.(*Tag)
}

// Make creates a Tag from a BCP 47 string. Repeated calls return the same pointer.
func Make(s string) *Tag {
	v, ok := makeCache.Load(s)
	if ok {
		return v.(*Tag)
	}
	tag := intern(xlanguage.Make(s))
	makeCache.Store(s, tag)
	return tag
}

// Parse creates a Tag from a BCP 47 string. Returns an error if parsing fails.
// Successful parses are interned and cached by input string.
func Parse(s string) (*Tag, error) {
	v, ok := makeCache.Load(s)
	if ok {
		return v.(*Tag), nil
	}
	t, err := xlanguage.Parse(s)
	if err != nil {
		return intern(t), err
	}
	tag := intern(t)
	makeCache.Store(s, tag)
	return tag, nil
}

// Tag returns the underlying golang.org/x/text/language.Tag for use with standard library APIs.
func (t *Tag) Tag() xlanguage.Tag {
	return t.underlying
}

// String returns the canonical BCP 47 representation.
func (t *Tag) String() string {
	return t.underlying.String()
}

// Weight returns the quality value (q) from Accept-Language header parsing.
// Defaults to 0 for tags not parsed from HTTP headers, 1.0 for header tags without explicit q.
func (t *Tag) Weight() float64 {
	return t.weight
}

// Base returns the base language code (e.g., "zh" from "zh-CN").
func (t *Tag) Base() string {
	b, _ := t.underlying.Base()
	return b.String()
}

// Region returns the region code (e.g., "CN" from "zh-CN").
func (t *Tag) Region() string {
	r, _ := t.underlying.Region()
	return r.String()
}

// Script returns the script code (e.g., "Hant" from "zh-Hant").
func (t *Tag) Script() string {
	s, _ := t.underlying.Script()
	return s.String()
}

// Parent returns the parent tag in the BCP 47 inheritance chain.
// For example: zh-CN → zh → und. Returns self if already root.
func (t *Tag) Parent() *Tag {
	return intern(t.underlying.Parent())
}

// FallbackChain returns the full inheritance chain from the tag to root (und).
// The first element is always the tag itself.
// Example: zh-CN → [zh-CN, zh, und]
func (t *Tag) FallbackChain() []*Tag {
	chain := make([]*Tag, 0, 4)
	chain = append(chain, t)
	cur := t.underlying
	for {
		p := cur.Parent()
		if p == cur {
			return chain
		}
		chain = append(chain, intern(p))
		cur = p
	}
}

// Match reports whether this tag or any of its parents equals the target.
//
//	Make("zh-CN").Match(Make("zh"))   // true
//	Make("en-US").Match(Make("en"))   // true
//	Make("zh-CN").Match(Make("en"))   // false
func (t *Tag) Match(target *Tag) bool {
	cur := t.underlying
	for {
		if cur == target.underlying {
			return true
		}
		p := cur.Parent()
		if p == cur {
			return false
		}
		cur = p
	}
}

// IsRTL reports whether the language is written right-to-left.
// Covers the 10 well-known RTL languages; not derived from CLDR data.
func (t *Tag) IsRTL() bool {
	b, _ := t.underlying.Base()
	switch b.String() {
	case "ar", "he", "fa", "ur", "ps", "sd", "ug", "ku", "dv", "yi":
		return true
	}
	return false
}

// ParseAcceptLanguage parses an HTTP Accept-Language header and returns
// language tags sorted by quality (q value) in descending order.
//
// Header format (RFC 7231 §5.3.5):
//
//	Accept-Language: da, en-gb;q=0.8, en;q=0.7
//
// Tags with q=0 are excluded. Tags that fail to parse are silently skipped.
func ParseAcceptLanguage(header string) []*Tag {
	if header == "" {
		return nil
	}

	tags := make([]*Tag, 0, strings.Count(header, ",")+1)
	for _, part := range strings.Split(header, ",") {
		part = strings.TrimSpace(part)
		if part == "" {
			continue
		}

		seg := strings.SplitN(part, ";", 2)
		raw := strings.TrimSpace(seg[0])

		q, ok := 1.0, true
		if len(seg) == 2 {
			q, ok = parseAcceptLanguageQ(seg[1])
		}
		if !ok {
			continue
		}

		tag, err := xlanguage.Parse(raw)
		if err != nil {
			continue
		}
		tags = append(tags, &Tag{underlying: tag, weight: q})
	}

	sort.SliceStable(tags, func(i, j int) bool {
		return tags[i].weight > tags[j].weight
	})
	return tags
}

// parseAcceptLanguageQ parses one ";q=X" param. Returns (q, true) for valid
// q in (0, 1]; (1.0, true) when param is not a q directive (treated as default);
// (0, false) when q is explicitly 0 or invalid such that caller must skip.
func parseAcceptLanguageQ(param string) (float64, bool) {
	param = strings.TrimSpace(param)
	if !strings.HasPrefix(param, "q=") {
		return 1.0, true
	}
	v, err := strconv.ParseFloat(param[2:], 64)
	if err != nil {
		return 1.0, true
	}
	if v <= 0 {
		return 0, false
	}
	if v > 1 {
		return 1.0, true
	}
	return v, true
}

// matcherCache holds prebuilt xlanguage.Matcher instances keyed by the
// canonical fingerprint of the supported tag list. NewMatcher is expensive
// (µs-scale) so caching pays back after the first call per supported set.
var matcherCache sync.Map // map[string]xlanguage.Matcher

// matcherFor returns a Matcher for the given supported list, building once per fingerprint.
func matcherFor(supported []*Tag) xlanguage.Matcher {
	var sb strings.Builder
	sb.Grow(len(supported) * 8)
	for _, t := range supported {
		sb.WriteString(t.underlying.String())
		sb.WriteByte('|')
	}
	key := sb.String()
	v, ok := matcherCache.Load(key)
	if ok {
		return v.(xlanguage.Matcher)
	}
	m := xlanguage.NewMatcher(toXTags(supported))
	actual, _ := matcherCache.LoadOrStore(key, m)
	return actual.(xlanguage.Matcher)
}

// Detect picks the best matching language from the Accept-Language header
// against the list of supported tags. Returns the matched Tag and its index.
// If no match is found, returns the first supported tag and index 0.
func Detect(header string, supported []*Tag) (*Tag, int) {
	if len(supported) == 0 {
		return nil, -1
	}

	parsed := ParseAcceptLanguage(header)
	if len(parsed) == 0 {
		return supported[0], 0
	}

	_, idx, _ := matcherFor(supported).Match(toXTags(parsed)...)
	return supported[idx], idx
}

// DetectFromStrings is a convenience wrapper that accepts string tags.
func DetectFromStrings(header string, supported []string) (*Tag, int) {
	tags := make([]*Tag, len(supported))
	for i, s := range supported {
		tags[i] = Make(s)
	}
	return Detect(header, tags)
}

func toXTags(tags []*Tag) []xlanguage.Tag {
	out := make([]xlanguage.Tag, len(tags))
	for i, t := range tags {
		out[i] = t.underlying
	}
	return out
}
