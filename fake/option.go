package fake

import (
	"math/rand/v2"

	xlanguage "golang.org/x/text/language"
)

// Option mutates a [Faker] during construction. Options are applied in the
// order supplied to [New]; later options override earlier ones.
type Option func(*Faker)

// WithSeed installs a deterministic random source seeded with the given
// value. Two fakers built with the same seed produce identical output for
// the same call sequence.
func WithSeed(seed int64) Option {
	return func(f *Faker) {
		seedU := uint64(seed)
		f.rng = rand.New(rand.NewPCG(seedU, seedU^0x9e3779b97f4a7c15))
	}
}

// WithRand installs an externally managed random source. The faker takes a
// mutex around every read from the source, so callers may freely share one
// *rand.Rand across multiple fakers.
func WithRand(r *rand.Rand) Option {
	return func(f *Faker) { f.rng = r }
}

// WithGender pins the default gender. Generators that do not accept a
// per-call gender argument fall back to this value.
func WithGender(g Gender) Option {
	return func(f *Faker) { f.gender = g }
}

// WithLang overrides the active language tag. When unset, the country's
// official language is used.
func WithLang(tag xlanguage.Tag) Option {
	return func(f *Faker) { f.lang = tag }
}
