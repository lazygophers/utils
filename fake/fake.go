// Package fake generates locale-aware fake data (names, identity numbers,
// contact info, addresses, network identifiers) keyed by ISO 3166-1 country
// plus locale-agnostic helpers (UUID, IP, lorem text, time, color, file).
//
// Two API shapes are supported:
//
//   - Explicit instance: [New] returns a [Faker] bound to a country.
//   - Global helpers (added in later patches) infer the country from
//     goroutine-local language state.
//
// Public APIs that surface language tags use [golang.org/x/text/language.Tag]
// values, not internal wrapper types.
package fake

import (
	"math/rand/v2"
	"sync"

	xlanguage "golang.org/x/text/language"

	"github.com/lazygophers/utils/country"
)

// Faker is a country-bound generator. Methods are safe for concurrent use:
// when constructed with [WithSeed] or [WithRand] the embedded *rand.Rand is
// protected by an internal mutex; otherwise generators draw from the
// package-global math/rand/v2 source which is itself goroutine-safe.
type Faker struct {
	country *country.Country
	locale  *Locale
	rng     *rand.Rand
	mu      sync.Mutex
	gender  Gender
	lang    xlanguage.Tag
}

// New constructs a [Faker] bound to the given country. The country must not
// be nil — callers obtain it from package-level constants such as
// [country.China] or from [country.Get]. Options may override the random
// source, default gender, or active language tag.
//
// When no locale is registered for the country, the United States locale is
// used as a fallback so that all generators stay productive.
func New(c *country.Country, opts ...Option) *Faker {
	if c == nil {
		panic("fake: New called with nil country")
	}
	f := &Faker{
		country: c,
		locale:  lookupLocale(c),
		gender:  GenderRandom,
		lang:    c.OfficialLanguage(),
	}
	for _, opt := range opts {
		if opt != nil {
			opt(f)
		}
	}
	if f.locale == nil {
		f.locale = registry["US"]
	}
	return f
}

// Country returns the country the faker was constructed with.
func (f *Faker) Country() *country.Country { return f.country }

// Locale returns the resolved locale (possibly the fallback locale). It may
// be nil when no locale is registered for either the requested country or
// the fallback.
func (f *Faker) Locale() *Locale { return f.locale }

// Lang returns the active language tag used to pick localised data pools.
func (f *Faker) Lang() xlanguage.Tag { return f.lang }

// DefaultGender returns the gender mode used when callers do not supply one.
func (f *Faker) DefaultGender() Gender { return f.gender }
