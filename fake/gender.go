package fake

import "math/rand/v2"

// Gender enumerates biological gender selection modes used when generating
// gender-dependent fake data (first names, id cards, etc.).
type Gender uint8

const (
	// GenderRandom defers selection until generation time and resolves to
	// either [GenderMale] or [GenderFemale] with equal probability.
	GenderRandom Gender = iota
	// GenderMale forces male-labelled data pools.
	GenderMale
	// GenderFemale forces female-labelled data pools.
	GenderFemale
)

// Resolve returns a concrete gender. When g is [GenderRandom] it draws one of
// [GenderMale] / [GenderFemale] using the supplied rng. When rng is nil it
// falls back to the runtime-wide source from math/rand/v2.
func (g Gender) Resolve(rng *rand.Rand) Gender {
	if g != GenderRandom {
		return g
	}
	var bit uint32
	if rng != nil {
		bit = rng.Uint32()
	} else {
		bit = rand.Uint32()
	}
	if bit&1 == 0 {
		return GenderMale
	}
	return GenderFemale
}

// String returns the lower-case label of the gender (random / male / female).
func (g Gender) String() string {
	switch g {
	case GenderMale:
		return "male"
	case GenderFemale:
		return "female"
	default:
		return "random"
	}
}
