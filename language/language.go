package language

import (
	"sort"
	"strconv"
	"strings"

	xlanguage "golang.org/x/text/language"
)

// Tag represents a parsed language tag with extended metadata.
// It wraps golang.org/x/text/language.Tag and can be converted back via [Tag.Tag].
type Tag struct {
	underlying xlanguage.Tag
	weight     float64 // q value from Accept-Language header, 0 if not from header
}

// Make creates a Tag from a BCP 47 string.
func Make(s string) *Tag {
	return &Tag{underlying: xlanguage.Make(s)}
}

// Parse creates a Tag from a BCP 47 string. Returns an error if parsing fails.
func Parse(s string) (*Tag, error) {
	t, err := xlanguage.Parse(s)
	return &Tag{underlying: t}, err
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
	p := t.underlying.Parent()
	return &Tag{underlying: p}
}

// FallbackChain returns the full inheritance chain from the tag to root (und).
// The first element is always the tag itself.
// Example: zh-CN → [zh-CN, zh, und]
func (t *Tag) FallbackChain() []*Tag {
	var chain []*Tag
	cur := t
	for {
		chain = append(chain, cur)
		p := cur.underlying.Parent()
		if p == cur.underlying {
			break
		}
		cur = &Tag{underlying: p}
	}
	return chain
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

	type weighted struct {
		tag xlanguage.Tag
		q   float64
	}

	var tags []weighted
	for _, part := range strings.Split(header, ",") {
		part = strings.TrimSpace(part)
		if part == "" {
			continue
		}

		seg := strings.SplitN(part, ";", 2)
		raw := strings.TrimSpace(seg[0])

		q := 1.0
		skip := false
		if len(seg) == 2 {
			for _, param := range strings.Split(seg[1], ";") {
				param = strings.TrimSpace(param)
				if strings.HasPrefix(param, "q=") {
					if v, err := strconv.ParseFloat(param[2:], 64); err == nil {
						if v <= 0 {
							skip = true
						} else if v <= 1 {
							q = v
						}
					}
					break
				}
			}
		}
		if skip {
			continue
		}

		tag, err := xlanguage.Parse(raw)
		if err != nil {
			continue
		}
		tags = append(tags, weighted{tag: tag, q: q})
	}

	sort.SliceStable(tags, func(i, j int) bool {
		return tags[i].q > tags[j].q
	})

	result := make([]*Tag, len(tags))
	for i, t := range tags {
		result[i] = &Tag{underlying: t.tag, weight: t.q}
	}
	return result
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

	xtags := make([]xlanguage.Tag, len(supported))
	for i, s := range supported {
		xtags[i] = s.underlying
	}

	matcher := xlanguage.NewMatcher(xtags)
	_, idx, _ := matcher.Match(toXTags(parsed)...)
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
