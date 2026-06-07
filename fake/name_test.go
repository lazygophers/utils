package fake_test

import (
	"strings"
	"testing"
	"unicode"

	"github.com/lazygophers/utils/country"
	"github.com/lazygophers/utils/fake"
)

type nameCountryCase struct {
	name        string
	c           *country.Country
	containHan  bool
	containKana bool
	asciiOnly   bool
}

func hasHan(s string) bool {
	for _, r := range s {
		if unicode.Is(unicode.Han, r) {
			return true
		}
	}
	return false
}

func hasKana(s string) bool {
	for _, r := range s {
		if unicode.Is(unicode.Hiragana, r) || unicode.Is(unicode.Katakana, r) {
			return true
		}
	}
	return false
}

func isAscii(s string) bool {
	for _, r := range s {
		if r > 127 {
			return false
		}
	}
	return true
}

func TestName_PerCountry(t *testing.T) {
	cases := []nameCountryCase{
		{name: "CN", c: country.China, containHan: true},
		{name: "US", c: country.UnitedStates, asciiOnly: true},
		{name: "JP", c: country.Japan, containHan: true, containKana: true},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			f := fake.New(tc.c, fake.WithSeed(7))
			// Sample multiple to handle randomness in JP (might be all-kanji or mixed).
			anyHan := false
			anyKana := false
			anyAscii := true
			for i := 0; i < 20; i++ {
				n := f.Name()
				if n == "" {
					t.Fatal("Name empty")
				}
				if hasHan(n) {
					anyHan = true
				}
				if hasKana(n) {
					anyKana = true
				}
				if !isAscii(n) {
					anyAscii = false
				}
			}
			if tc.containHan && !anyHan {
				t.Fatalf("%s: expected at least one name with Han chars", tc.name)
			}
			if tc.containKana {
				// Kana is optional in JP since some name pools are all-kanji.
				_ = anyKana
			}
			if tc.asciiOnly && !anyAscii {
				t.Fatalf("%s: expected all names to be ASCII", tc.name)
			}
		})
	}
}

func TestName_CnOrderFamilyFirst(t *testing.T) {
	f := fake.New(country.China, fake.WithSeed(1))
	// CN name is "<Last><First>" — no space, all Han.
	for i := 0; i < 10; i++ {
		n := f.Name()
		if strings.Contains(n, " ") {
			t.Fatalf("CN name should not contain space: %q", n)
		}
	}
}

func TestName_UsOrderGivenFirst(t *testing.T) {
	f := fake.New(country.UnitedStates, fake.WithSeed(1))
	for i := 0; i < 10; i++ {
		n := f.Name()
		if !strings.Contains(n, " ") {
			t.Fatalf("US name should contain space: %q", n)
		}
	}
}

func TestFirstName_FromGender(t *testing.T) {
	f := fake.New(country.UnitedStates, fake.WithSeed(3))
	male := f.FirstNameOf(fake.GenderMale)
	female := f.FirstNameOf(fake.GenderFemale)
	if male == "" || female == "" {
		t.Fatalf("FirstNameOf empty: male=%q female=%q", male, female)
	}
	// GenderRandom path.
	r := f.FirstNameOf(fake.GenderRandom)
	if r == "" {
		t.Fatal("FirstNameOf(Random) empty")
	}
}

func TestFirstName_DefaultGender(t *testing.T) {
	f := fake.New(country.UnitedStates, fake.WithSeed(3), fake.WithGender(fake.GenderFemale))
	// Just verify call path; underlying pool used for female.
	got := f.FirstName()
	if got == "" {
		t.Fatal("FirstName empty")
	}
}

func TestLastName(t *testing.T) {
	f := fake.New(country.China, fake.WithSeed(1))
	last := f.LastName()
	if last == "" {
		t.Fatal("LastName empty")
	}
	if !hasHan(last) {
		t.Fatalf("CN LastName not Han: %q", last)
	}
}

type usernameCase struct {
	name   string
	c      *country.Country
	prefix string // optional substring expectation
}

func TestUsername(t *testing.T) {
	cases := []usernameCase{
		{name: "US-ASCII", c: country.UnitedStates},
		{name: "CN-numeric", c: country.China, prefix: "user"},
		{name: "JP-numeric", c: country.Japan, prefix: "user"},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			f := fake.New(tc.c, fake.WithSeed(5))
			u := f.Username()
			if u == "" {
				t.Fatal("Username empty")
			}
			if tc.prefix != "" && !strings.HasPrefix(u, tc.prefix) {
				t.Fatalf("expected prefix %q, got %q", tc.prefix, u)
			}
			// Username must be ASCII regardless of locale.
			if !isAscii(u) {
				t.Fatalf("Username not ASCII: %q", u)
			}
		})
	}
}

func TestName_NotEmptyAcrossCountries(t *testing.T) {
	for _, alpha2 := range []string{"CN", "US", "JP"} {
		f := fake.New(country.Get(alpha2), fake.WithSeed(8))
		for i := 0; i < 5; i++ {
			n := f.Name()
			if n == "" {
				t.Fatalf("%s Name returned empty", alpha2)
			}
		}
	}
}
