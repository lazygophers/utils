package fake

import (
	"fmt"
	"math/rand/v2"
	"time"

	xlanguage "golang.org/x/text/language"

	"github.com/lazygophers/utils/country"
)

// CityEntry describes a single city sample inside a [Locale] city pool.
type CityEntry struct {
	Name     string
	Province string
	Lat      float64
	Lng      float64
}

// IdCardGenFunc generates a national identification number for the given
// gender and birth date using the supplied random source.
type IdCardGenFunc func(rng *rand.Rand, gender Gender, birth time.Time) string

// Locale carries all country-scoped fake data and helpers consumed by
// [Faker] generators. One Locale is registered per [country.Country].
type Locale struct {
	// Country is the owning ISO 3166-1 entry; never nil for registered locales.
	Country *country.Country
	// OfficialLangs lists language tags treated as native for this country.
	OfficialLangs []xlanguage.Tag
	// PhonePrefixes are mobile number prefixes (without the calling code).
	PhonePrefixes []string
	// LandlinePrefix lists fixed-line area / trunk prefixes.
	LandlinePrefix []string
	// ZipFormat is a free-form template describing postal code shape.
	ZipFormat string
	// IdCardGen, when non-nil, generates a national ID; nil means unsupported.
	IdCardGen IdCardGenFunc
	// Streets maps a language tag to a pool of street name templates.
	Streets map[xlanguage.Tag][]string
	// Cities maps a language tag to a pool of city entries.
	Cities map[xlanguage.Tag][]CityEntry
	// FirstNames maps language tag → gender → first name pool.
	FirstNames map[xlanguage.Tag]map[Gender][]string
	// LastNames maps a language tag to a pool of last names.
	LastNames map[xlanguage.Tag][]string
	// Domain is the country code top-level domain (without the leading dot).
	Domain string
	// BrowserBias overrides global browser market share at the country level.
	// Empty map means fall back to the package-wide share table. Keys are the
	// [Browser] enum defined in useragents.go.
	BrowserBias map[Browser]int
	// AppBias overrides the global app user-agent pool at the country level.
	// Each entry is a template key string (e.g. "wechat-android",
	// "wechat-ios", "qq-android", "alipay", "douyin", "uc", "quark",
	// "weibo"). Empty slice means use the global pool with uniform weights.
	AppBias []string
}

// registry stores all registered locales keyed by the country alpha-2 code
// (uppercase). It is populated by per-country init functions.
var registry = map[string]*Locale{}

// register installs the locale for the country into the registry. It panics
// if the same alpha-2 code is registered twice, which is always a programmer
// error during package initialisation.
func register(l *Locale) {
	if l == nil || l.Country == nil {
		panic("fake: register called with nil locale or country")
	}
	key := l.Country.Alpha2()
	_, dup := registry[key]
	if dup {
		panic(fmt.Sprintf("fake: duplicate locale registration for %s", key))
	}
	registry[key] = l
}

// lookupLocale resolves the locale for the country. Returns nil when no
// locale is registered for the country (callers handle fallback).
func lookupLocale(c *country.Country) *Locale {
	if c == nil {
		return nil
	}
	return registry[c.Alpha2()]
}
