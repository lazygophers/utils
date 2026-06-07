package fake_test

import (
	"testing"
	"time"

	"github.com/lazygophers/utils/country"
	"github.com/lazygophers/utils/fake"
)

func TestDate_InRange(t *testing.T) {
	f := fake.New(country.UnitedStates, fake.WithSeed(1))
	min := time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC)
	max := time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC)
	for i := 0; i < 50; i++ {
		d := f.Date(min, max)
		if d.Before(min) || d.After(max) {
			t.Fatalf("Date %v outside [%v, %v]", d, min, max)
		}
	}
}

func TestDate_SwapBounds(t *testing.T) {
	f := fake.New(country.UnitedStates, fake.WithSeed(2))
	min := time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC)
	max := time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC)
	d := f.Date(max, min) // reversed
	if d.Before(min) || d.After(max) {
		t.Fatalf("swap bounds failed: %v", d)
	}
}

func TestDate_Equal(t *testing.T) {
	f := fake.New(country.UnitedStates, fake.WithSeed(3))
	when := time.Date(2022, 6, 15, 12, 0, 0, 0, time.UTC)
	got := f.Date(when, when)
	if !got.Equal(when) {
		t.Fatalf("equal bounds: %v vs %v", got, when)
	}
}

func TestTime_RecentTenYears(t *testing.T) {
	f := fake.New(country.UnitedStates, fake.WithSeed(4))
	now := time.Now()
	for i := 0; i < 20; i++ {
		ts := f.Time()
		if ts.After(now.Add(time.Minute)) {
			t.Fatalf("Time in future: %v", ts)
		}
		if ts.Before(now.AddDate(-10, 0, -1)) {
			t.Fatalf("Time too far in past: %v", ts)
		}
	}
}

func TestIntRange(t *testing.T) {
	f := fake.New(country.UnitedStates, fake.WithSeed(5))
	for i := 0; i < 100; i++ {
		v := f.IntRange(10, 20)
		if v < 10 || v > 20 {
			t.Fatalf("IntRange %d out of [10,20]", v)
		}
	}
	if f.IntRange(5, 5) != 5 {
		t.Fatal("IntRange(5,5) != 5")
	}
	// Swap bounds.
	v := f.IntRange(20, 10)
	if v < 10 || v > 20 {
		t.Fatalf("swap IntRange: %d", v)
	}
}

func TestInt64Range(t *testing.T) {
	f := fake.New(country.UnitedStates, fake.WithSeed(6))
	for i := 0; i < 50; i++ {
		v := f.Int64Range(-100, 100)
		if v < -100 || v > 100 {
			t.Fatalf("Int64Range %d out", v)
		}
	}
	if f.Int64Range(7, 7) != 7 {
		t.Fatal("Int64Range(7,7) != 7")
	}
	v := f.Int64Range(50, -50)
	if v < -50 || v > 50 {
		t.Fatalf("swap: %d", v)
	}
}

func TestFloat64Range(t *testing.T) {
	f := fake.New(country.UnitedStates, fake.WithSeed(7))
	for i := 0; i < 50; i++ {
		v := f.Float64Range(1.0, 2.0)
		if v < 1.0 || v >= 2.0 {
			t.Fatalf("Float64Range %f out", v)
		}
	}
	if f.Float64Range(3.14, 3.14) != 3.14 {
		t.Fatal("Float64Range equal bounds")
	}
	v := f.Float64Range(2.0, 1.0)
	if v < 1.0 || v >= 2.0 {
		t.Fatalf("swap: %f", v)
	}
}

func TestBool_BothValues(t *testing.T) {
	f := fake.New(country.UnitedStates, fake.WithSeed(8))
	trueCount, falseCount := 0, 0
	for i := 0; i < 200; i++ {
		if f.Bool() {
			trueCount++
		} else {
			falseCount++
		}
	}
	if trueCount == 0 || falseCount == 0 {
		t.Fatalf("Bool not balanced: t=%d f=%d", trueCount, falseCount)
	}
}

func TestPick(t *testing.T) {
	f := fake.New(country.UnitedStates, fake.WithSeed(9))
	s := []int{10, 20, 30, 40}
	for i := 0; i < 20; i++ {
		v := fake.Pick(f, s)
		found := false
		for _, x := range s {
			if x == v {
				found = true
				break
			}
		}
		if !found {
			t.Fatalf("Pick returned non-member %d", v)
		}
	}
	// Empty slice → zero value.
	zero := fake.Pick(f, []int{})
	if zero != 0 {
		t.Fatalf("Pick(empty) = %d, want 0", zero)
	}
}

func TestSample(t *testing.T) {
	f := fake.New(country.UnitedStates, fake.WithSeed(10))
	src := []string{"a", "b", "c", "d", "e"}
	got := fake.Sample(f, src, 3)
	if len(got) != 3 {
		t.Fatalf("Sample(3) len = %d", len(got))
	}
	seen := map[string]bool{}
	for _, v := range got {
		if seen[v] {
			t.Fatalf("Sample produced duplicate %q", v)
		}
		seen[v] = true
	}
	// n >= len returns shuffled full slice.
	full := fake.Sample(f, src, 10)
	if len(full) != len(src) {
		t.Fatalf("Sample(>=len) returned %d", len(full))
	}
	// n <= 0 and empty src.
	if len(fake.Sample(f, src, 0)) != 0 {
		t.Fatal("Sample(0) not empty")
	}
	if len(fake.Sample(f, []string{}, 3)) != 0 {
		t.Fatal("Sample(empty) not empty")
	}
}

func TestShuffle(t *testing.T) {
	f := fake.New(country.UnitedStates, fake.WithSeed(11))
	s := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	orig := make([]int, len(s))
	copy(orig, s)
	fake.Shuffle(f, s)
	// Same elements.
	sum := 0
	for _, v := range s {
		sum += v
	}
	if sum != 55 {
		t.Fatalf("Shuffle lost elements: sum=%d", sum)
	}
	// Highly unlikely all positions match; allow some.
	matches := 0
	for i := range s {
		if s[i] == orig[i] {
			matches++
		}
	}
	if matches == len(s) {
		t.Fatal("Shuffle made no changes")
	}
	// Empty / single slice no-op.
	empty := []int{}
	fake.Shuffle(f, empty)
	single := []int{42}
	fake.Shuffle(f, single)
	if single[0] != 42 {
		t.Fatal("Shuffle on single element changed it")
	}
}
