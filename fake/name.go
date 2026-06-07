package fake

import (
	"strconv"
	"strings"

	xlanguage "golang.org/x/text/language"
)

// lastNameFirstAlpha2 lists ISO 3166-1 alpha-2 codes whose conventional
// personal name order places the family name before the given name.
var lastNameFirstAlpha2 = map[string]bool{
	"CN": true,
	"TW": true,
	"HK": true,
	"MO": true,
	"JP": true,
	"KR": true,
	"KP": true,
	"VN": true,
	"HU": true,
	"MN": true,
}

// asciiLangs marks language base codes whose first names are reliable as
// ASCII-derived usernames. Anything else falls back to a numeric handle.
var asciiLangs = map[string]bool{
	"en": true,
	"es": true,
	"pt": true,
	"fr": true,
	"de": true,
	"it": true,
	"nl": true,
	"id": true,
	"ms": true,
	"sv": true,
	"no": true,
	"da": true,
	"fi": true,
	"pl": true,
	"cs": true,
	"sk": true,
	"hr": true,
	"sl": true,
	"ro": true,
	"hu": true,
	"tr": true,
	"vi": true,
}

// LastName returns a single locale-appropriate family name. The pool is
// resolved with a language fallback chain so that even countries lacking
// dedicated data still yield a non-empty result.
func (f *Faker) LastName() string {
	pool := f.resolveLastNamePool()
	return f.pickString(pool)
}

// FirstName returns a single locale-appropriate given name. When the faker
// was constructed with [GenderRandom] the gender is drawn uniformly at
// generation time.
func (f *Faker) FirstName() string {
	return f.FirstNameOf(f.gender)
}

// FirstNameOf returns a single locale-appropriate given name biased to the
// supplied gender. Pass [GenderRandom] to draw a gender uniformly.
func (f *Faker) FirstNameOf(gender Gender) string {
	g := gender
	if g == GenderRandom {
		if f.intN(2) == 0 {
			g = GenderMale
		} else {
			g = GenderFemale
		}
	}
	pool := f.resolveFirstNamePool(g)
	return f.pickString(pool)
}

// Name returns a full personal name composed of a first and last name in the
// order conventional for the faker's country (family name first for CJK and
// a handful of others; given name first elsewhere).
func (f *Faker) Name() string {
	first := f.FirstName()
	last := f.LastName()
	if f.lastNameFirst() {
		return last + first
	}
	if first == "" {
		return last
	}
	if last == "" {
		return first
	}
	return first + " " + last
}

// Username returns a synthetic account handle derived from the first name
// when the active language uses an ASCII-friendly script; otherwise it
// returns a numeric handle (“user“ + 6 digits) to avoid mangled output.
func (f *Faker) Username() string {
	base, _ := f.lang.Base()
	if asciiLangs[base.String()] {
		first := f.FirstName()
		clean := sanitizeUsername(first)
		if clean == "" {
			return numericUsername(f)
		}
		suffix := f.intN(1000)
		return clean + "_" + padDigits(suffix, 3)
	}
	return numericUsername(f)
}

// resolveFirstNamePool walks the fallback chain to find a non-empty pool of
// first names for the requested gender:
//
//  1. f.locale at f.lang
//  2. f.locale at its first official language
//  3. the US locale at English
//
// Returns nil when even the US locale has no data.
func (f *Faker) resolveFirstNamePool(g Gender) []string {
	pool := firstNamesFor(f.locale, f.lang, g)
	if len(pool) > 0 {
		return pool
	}
	if len(f.locale.OfficialLangs) > 0 {
		pool = firstNamesFor(f.locale, f.locale.OfficialLangs[0], g)
		if len(pool) > 0 {
			return pool
		}
	}
	return firstNamesFor(registry["US"], xlanguage.English, g)
}

// resolveLastNamePool walks the same fallback chain as
// [Faker.resolveFirstNamePool] for family names.
func (f *Faker) resolveLastNamePool() []string {
	pool := lastNamesFor(f.locale, f.lang)
	if len(pool) > 0 {
		return pool
	}
	if len(f.locale.OfficialLangs) > 0 {
		pool = lastNamesFor(f.locale, f.locale.OfficialLangs[0])
		if len(pool) > 0 {
			return pool
		}
	}
	return lastNamesFor(registry["US"], xlanguage.English)
}

// lastNameFirst reports whether the faker's country places the family name
// before the given name in its conventional written order.
func (f *Faker) lastNameFirst() bool {
	return lastNameFirstAlpha2[f.country.Alpha2()]
}

func firstNamesFor(l *Locale, tag xlanguage.Tag, g Gender) []string {
	return l.FirstNames[tag][g]
}

func lastNamesFor(l *Locale, tag xlanguage.Tag) []string {
	return l.LastNames[tag]
}

func numericUsername(f *Faker) string {
	return "user" + padDigits(f.intN(1_000_000), 6)
}

func padDigits(n, width int) string {
	s := strconv.Itoa(n)
	if len(s) >= width {
		return s
	}
	return strings.Repeat("0", width-len(s)) + s
}

// sanitizeUsername lower-cases the input and keeps only [a-z0-9]. Returns
// empty when no usable characters survive.
func sanitizeUsername(s string) string {
	if s == "" {
		return ""
	}
	lower := strings.ToLower(s)
	var b strings.Builder
	b.Grow(len(lower))
	for _, r := range lower {
		if r >= 'a' && r <= 'z' {
			b.WriteRune(r)
			continue
		}
		if r >= '0' && r <= '9' {
			b.WriteRune(r)
		}
	}
	return b.String()
}
