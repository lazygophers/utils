package fake

import (
	"strings"
	"time"
)

// passportPrefixByAlpha2 maps an ISO 3166-1 alpha-2 country code to the
// letter prefix observed on real passports issued by that country. Codes
// not present here fall back to the generic "P" prefix used by many issuers.
var passportPrefixByAlpha2 = map[string]string{
	"CN": "E",
	"US": "A",
	"JP": "TR",
	"GB": "P",
	"DE": "C",
	"FR": "P",
	"KR": "M",
	"RU": "P",
}

// IdCard returns a national identification number for the faker's country.
// When the locale provides an [IdCardGenFunc] (e.g. China's GB 11643 ID,
// the US Social Security Number, Japan's My Number) the locale-specific
// generator is used. Otherwise a generic 12-digit numeric string is
// returned so callers always observe a non-empty value.
//
// The gender and birthday driving the generation are picked at random
// within reasonable adult ranges. Use [Faker.IdCardOf] to supply them
// explicitly.
func (f *Faker) IdCard() string {
	g := f.gender
	if g == GenderRandom {
		if f.intN(2) == 0 {
			g = GenderMale
		} else {
			g = GenderFemale
		}
	}
	birth := f.Birthday(18, 80)
	return f.IdCardOf(g, birth)
}

// IdCardOf returns a national identification number for the given gender
// and birth date. When the locale exposes no dedicated generator a generic
// 12-digit numeric string is returned.
func (f *Faker) IdCardOf(gender Gender, birth time.Time) string {
	if f.locale.IdCardGen != nil {
		return f.locale.IdCardGen(f.rng, gender, birth)
	}
	return f.genericNumericId(12)
}

// PassportNo returns a synthetic passport number shaped like the ones
// issued by the faker's country: a one or two letter prefix followed by
// eight decimal digits. Countries without a documented prefix fall back
// to "P".
func (f *Faker) PassportNo() string {
	prefix := f.passportPrefix()
	return prefix + f.genericNumericId(8)
}

// passportPrefix returns the passport letter prefix for the faker's
// country, defaulting to "P" when no mapping exists.
func (f *Faker) passportPrefix() string {
	p, ok := passportPrefixByAlpha2[f.country.Alpha2()]
	if ok {
		return p
	}
	return "P"
}

// genericNumericId returns a string of n random decimal digits.
func (f *Faker) genericNumericId(n int) string {
	var b strings.Builder
	b.Grow(n)
	for i := 0; i < n; i++ {
		b.WriteByte(byte('0' + f.intN(10)))
	}
	return b.String()
}
