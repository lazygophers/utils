package fake

import "math/rand/v2"

// intN returns a non-negative pseudo-random int in [0, n). Panics when n
// is non-positive, matching the contract of [rand.IntN].
func (f *Faker) intN(n int) int {
	if f.rng == nil {
		return rand.IntN(n)
	}
	f.mu.Lock()
	v := f.rng.IntN(n)
	f.mu.Unlock()
	return v
}

// uint64 returns a uniformly distributed random uint64.
func (f *Faker) uint64() uint64 {
	if f.rng == nil {
		return rand.Uint64()
	}
	f.mu.Lock()
	v := f.rng.Uint64()
	f.mu.Unlock()
	return v
}

// float64 returns a pseudo-random float64 in [0.0, 1.0).
func (f *Faker) float64() float64 {
	if f.rng == nil {
		return rand.Float64()
	}
	f.mu.Lock()
	v := f.rng.Float64()
	f.mu.Unlock()
	return v
}

// pickString returns a uniformly chosen element from s. Returns the empty
// string when s is empty.
func (f *Faker) pickString(s []string) string {
	if len(s) == 0 {
		return ""
	}
	return s[f.intN(len(s))]
}

// pick returns a uniformly chosen element from s. Returns the zero value
// when s is empty.
func pick[T any](rng *rand.Rand, s []T) T {
	var zero T
	if len(s) == 0 {
		return zero
	}
	if rng == nil {
		return s[rand.IntN(len(s))]
	}
	return s[rng.IntN(len(s))]
}

// shuffle reorders s in place using the Fisher–Yates algorithm.
func shuffle[T any](rng *rand.Rand, s []T) {
	if rng == nil {
		rand.Shuffle(len(s), func(i, j int) { s[i], s[j] = s[j], s[i] })
		return
	}
	rng.Shuffle(len(s), func(i, j int) { s[i], s[j] = s[j], s[i] })
}
