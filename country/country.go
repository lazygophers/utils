// Package country provides ISO 3166-1 country and dependent-territory data
// (alpha-2, alpha-3, numeric) together with multi-language names, primary
// IANA time zones, ITU-T calling codes, the main ISO 4217 currency, country
// code top-level domains, official languages, capital city, continent /
// region / sub-region classification, and flag emoji.
//
// Coverage includes all 249 ISO 3166-1 entries. Dependent territories with
// their own calling code, currency, or top-level domain (e.g. Hong Kong,
// Macao, Taiwan, Puerto Rico) are exposed as independent entries rather
// than merged into their parent state.
//
// Two access shapes are offered:
//
//   - Lookup by code/name: [Get], [GetByAlpha3], [GetByNumeric], [GetByName].
//   - Strongly-typed package-level constants: [China], [UnitedStates], ...
//
// All public APIs that accept a language tag use the standard library type
// golang.org/x/text/language.Tag. Goroutine-local language resolution is
// provided through github.com/lazygophers/utils/language.
package country

import (
	"sync"

	xlanguage "golang.org/x/text/language"

	"github.com/lazygophers/utils/currency"
)

// Country is an immutable ISO 3166-1 country / territory entry.
//
// All fields are unexported; accessors return defensive copies for slice
// fields. Localised names are mutated only during package init via
// [Country.RegisterName], [Country.RegisterOfficialName], and
// [Country.RegisterCapital]; readers take an RLock at runtime.
type Country struct {
	alpha2       string
	alpha3       string
	numeric      int
	callingCodes []string
	timezones    []string
	tlds         []string
	languages    []xlanguage.Tag
	currency     *currency.Currency
	continent    string
	region       string
	subregion    string
	flagEmoji    string

	namesMu  sync.RWMutex
	names    map[xlanguage.Tag]string
	official map[xlanguage.Tag]string
	capital  map[xlanguage.Tag]string
}

// Alpha2 returns the ISO 3166-1 alpha-2 code (e.g. "CN").
func (c *Country) Alpha2() string { return c.alpha2 }

// Alpha3 returns the ISO 3166-1 alpha-3 code (e.g. "CHN").
func (c *Country) Alpha3() string { return c.alpha3 }

// Numeric returns the ISO 3166-1 numeric code (e.g. 156).
func (c *Country) Numeric() int { return c.numeric }

// CallingCodes returns a copy of the ITU-T E.164 calling codes (with "+" prefix).
func (c *Country) CallingCodes() []string {
	out := make([]string, len(c.callingCodes))
	copy(out, c.callingCodes)
	return out
}

// Timezones returns a copy of the primary IANA time zones for this country.
func (c *Country) Timezones() []string {
	out := make([]string, len(c.timezones))
	copy(out, c.timezones)
	return out
}

// Tlds returns a copy of the country-code top-level domains (e.g. [".cn"]).
func (c *Country) Tlds() []string {
	out := make([]string, len(c.tlds))
	copy(out, c.tlds)
	return out
}

// Languages returns a copy of the official language tags.
func (c *Country) Languages() []xlanguage.Tag {
	out := make([]xlanguage.Tag, len(c.languages))
	copy(out, c.languages)
	return out
}

// Currency returns the main ISO 4217 currency for this country.
func (c *Country) Currency() *currency.Currency { return c.currency }

// Continent returns the two-letter continent code ("AS", "EU", "AF", "NA",
// "SA", "OC", "AN").
func (c *Country) Continent() string { return c.continent }

// Region returns the UN M.49 region (e.g. "Asia").
func (c *Country) Region() string { return c.region }

// Subregion returns the UN M.49 sub-region (e.g. "Eastern Asia").
func (c *Country) Subregion() string { return c.subregion }

// FlagEmoji returns the Unicode flag emoji (regional indicator pair).
func (c *Country) FlagEmoji() string { return c.flagEmoji }

// String returns the Alpha2 code, satisfying fmt.Stringer.
func (c *Country) String() string { return c.alpha2 }
