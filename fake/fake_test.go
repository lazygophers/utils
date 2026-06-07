package fake_test

import (
	"math/rand/v2"
	"sync"
	"testing"

	xlanguage "golang.org/x/text/language"

	"github.com/lazygophers/utils/country"
	"github.com/lazygophers/utils/fake"
)

func TestNew_BindsCountry(t *testing.T) {
	f := fake.New(country.China)
	if f.Country() != country.China {
		t.Fatalf("Country mismatch: %v", f.Country())
	}
	if f.Locale() == nil {
		t.Fatal("Locale should be non-nil for CN")
	}
}

func TestNew_NilCountryPanics(t *testing.T) {
	defer func() {
		r := recover()
		if r == nil {
			t.Fatal("expected panic")
		}
	}()
	_ = fake.New(nil)
}

func TestNew_RegisteredCountriesHaveLocales(t *testing.T) {
	for _, alpha2 := range []string{"CN", "US", "JP"} {
		c := country.Get(alpha2)
		f := fake.New(c)
		if f.Locale() == nil {
			t.Fatalf("%s locale should not be nil", alpha2)
		}
	}
}

func TestOptions(t *testing.T) {
	r := rand.New(rand.NewPCG(7, 11))
	f := fake.New(country.China,
		fake.WithRand(r),
		fake.WithGender(fake.GenderMale),
		fake.WithLang(xlanguage.Chinese),
	)
	if f.DefaultGender() != fake.GenderMale {
		t.Fatalf("DefaultGender = %v", f.DefaultGender())
	}
	if f.Lang() != xlanguage.Chinese {
		t.Fatalf("Lang = %v", f.Lang())
	}
}

func TestWithLangOverride(t *testing.T) {
	f := fake.New(country.China, fake.WithLang(xlanguage.English))
	if f.Lang() != xlanguage.English {
		t.Fatalf("Lang override failed: %v", f.Lang())
	}
}

type reproCase struct {
	name string
	fn   func(*fake.Faker) string
}

func TestWithSeedRepro(t *testing.T) {
	cases := []reproCase{
		{name: "Name", fn: func(f *fake.Faker) string { return f.Name() }},
		{name: "Email", fn: func(f *fake.Faker) string { return f.Email() }},
		{name: "Phone", fn: func(f *fake.Faker) string { return f.Phone() }},
		{name: "IdCard", fn: func(f *fake.Faker) string { return f.IdCard() }},
		{name: "UUIDv4", fn: func(f *fake.Faker) string { return f.UUIDv4() }},
		{name: "City", fn: func(f *fake.Faker) string { return f.City() }},
		{name: "Username", fn: func(f *fake.Faker) string { return f.Username() }},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			a := fake.New(country.China, fake.WithSeed(42))
			b := fake.New(country.China, fake.WithSeed(42))
			for i := 0; i < 10; i++ {
				x := tc.fn(a)
				y := tc.fn(b)
				if x != y {
					t.Fatalf("iter %d mismatch: %q vs %q", i, x, y)
				}
			}
		})
	}
}

func TestWithSeed_DifferentSeedsDiffer(t *testing.T) {
	a := fake.New(country.UnitedStates, fake.WithSeed(1))
	b := fake.New(country.UnitedStates, fake.WithSeed(2))
	same := 0
	for i := 0; i < 20; i++ {
		if a.UUIDv4() == b.UUIDv4() {
			same++
		}
	}
	if same > 2 {
		t.Fatalf("seeds 1 and 2 produced %d identical UUIDs", same)
	}
}

func TestFaker_ConcurrentSafe(t *testing.T) {
	f := fake.New(country.UnitedStates, fake.WithSeed(99))
	const goroutines = 50
	const each = 30
	var wg sync.WaitGroup
	wg.Add(goroutines)
	for g := 0; g < goroutines; g++ {
		go func() {
			defer wg.Done()
			for i := 0; i < each; i++ {
				_ = f.Name()
				_ = f.UUIDv4()
				_ = f.IPv4()
			}
		}()
	}
	wg.Wait()
}

func TestDefaultFaker_ConcurrentSafe(t *testing.T) {
	const goroutines = 50
	const each = 20
	var wg sync.WaitGroup
	wg.Add(goroutines)
	for g := 0; g < goroutines; g++ {
		go func() {
			defer wg.Done()
			for i := 0; i < each; i++ {
				_ = fake.Name()
				_ = fake.UUIDv4()
			}
		}()
	}
	wg.Wait()
}

func TestGender_String(t *testing.T) {
	if fake.GenderMale.String() != "male" {
		t.Fatalf("male: %s", fake.GenderMale.String())
	}
	if fake.GenderFemale.String() != "female" {
		t.Fatalf("female: %s", fake.GenderFemale.String())
	}
	if fake.GenderRandom.String() != "random" {
		t.Fatalf("random: %s", fake.GenderRandom.String())
	}
}

func TestGender_Resolve(t *testing.T) {
	r := rand.New(rand.NewPCG(1, 2))
	got := fake.GenderRandom.Resolve(r)
	if got != fake.GenderMale && got != fake.GenderFemale {
		t.Fatalf("Resolve returned %v", got)
	}
	// Resolve on already-fixed gender returns same.
	if fake.GenderMale.Resolve(r) != fake.GenderMale {
		t.Fatal("male.Resolve changed")
	}
	if fake.GenderFemale.Resolve(r) != fake.GenderFemale {
		t.Fatal("female.Resolve changed")
	}
	// Resolve with nil rng goes through global path.
	got2 := fake.GenderRandom.Resolve(nil)
	if got2 != fake.GenderMale && got2 != fake.GenderFemale {
		t.Fatalf("Resolve(nil) returned %v", got2)
	}
}

func TestNilOption(t *testing.T) {
	// New must tolerate a nil Option silently.
	f := fake.New(country.UnitedStates, nil, fake.WithSeed(1))
	if f.Country().Alpha2() != "US" {
		t.Fatal("Country lost after nil option")
	}
}
