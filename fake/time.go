package fake

import "time"

// Date returns a uniformly random time in the inclusive interval [min, max].
// When min is after max the bounds are swapped so callers do not need to
// pre-order them.
func (f *Faker) Date(min, max time.Time) time.Time {
	if min.After(max) {
		min, max = max, min
	}
	lo := min.UnixNano()
	hi := max.UnixNano()
	if lo == hi {
		return min
	}
	span := hi - lo
	delta := int64(f.float64() * float64(span))
	return time.Unix(0, lo+delta).In(min.Location())
}

// Time returns a random instant within the last ten years up to now.
func (f *Faker) Time() time.Time {
	now := time.Now()
	return f.Date(now.AddDate(-10, 0, 0), now)
}

// Birthday returns a random date such that the resulting age falls in the
// inclusive year range [minAge, maxAge]. When minAge > maxAge the bounds are
// swapped. Negative ages are clamped to zero.
func (f *Faker) Birthday(minAge, maxAge int) time.Time {
	if minAge > maxAge {
		minAge, maxAge = maxAge, minAge
	}
	if minAge < 0 {
		minAge = 0
	}
	if maxAge < 0 {
		maxAge = 0
	}
	now := time.Now()
	oldest := now.AddDate(-maxAge-1, 0, 1) // one day after the (maxAge+1)-th birthday
	youngest := now.AddDate(-minAge, 0, 0)
	return f.Date(oldest, youngest)
}

// IntRange returns a uniformly random int in the inclusive range [min, max].
// When min > max the bounds are swapped.
func (f *Faker) IntRange(min, max int) int {
	if min > max {
		min, max = max, min
	}
	if min == max {
		return min
	}
	return min + f.intN(max-min+1)
}

// Int64Range returns a uniformly random int64 in the inclusive range
// [min, max]. When min > max the bounds are swapped.
func (f *Faker) Int64Range(min, max int64) int64 {
	if min > max {
		min, max = max, min
	}
	if min == max {
		return min
	}
	span := uint64(max - min + 1)
	return min + int64(f.uint64()%span)
}

// Float64Range returns a uniformly random float64 in the half-open interval
// [min, max). When min > max the bounds are swapped. When min == max the
// shared value is returned.
func (f *Faker) Float64Range(min, max float64) float64 {
	if min > max {
		min, max = max, min
	}
	if min == max {
		return min
	}
	return min + f.float64()*(max-min)
}

// Bool returns true or false with equal probability.
func (f *Faker) Bool() bool {
	return f.uint64()&1 == 1
}

// Pick returns a uniformly chosen element from s. The zero value of T is
// returned when s is empty.
func Pick[T any](f *Faker, s []T) T {
	var zero T
	if len(s) == 0 {
		return zero
	}
	return s[f.intN(len(s))]
}

// Sample returns n distinct elements drawn without replacement from s. When
// n is greater than or equal to len(s) a shuffled copy of the full slice is
// returned. When n <= 0 or s is empty an empty slice is returned.
func Sample[T any](f *Faker, s []T, n int) []T {
	if n <= 0 || len(s) == 0 {
		return []T{}
	}
	if n >= len(s) {
		n = len(s)
	}
	pool := make([]T, len(s))
	copy(pool, s)
	// Partial Fisher–Yates: swap the first n positions with random tails.
	for i := 0; i < n; i++ {
		j := i + f.intN(len(pool)-i)
		pool[i], pool[j] = pool[j], pool[i]
	}
	return pool[:n]
}

// Shuffle reorders s in place using the Fisher–Yates algorithm. The slice is
// left untouched when it has fewer than two elements.
func Shuffle[T any](f *Faker, s []T) {
	for i := len(s) - 1; i > 0; i-- {
		j := f.intN(i + 1)
		s[i], s[j] = s[j], s[i]
	}
}
