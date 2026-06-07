package fake_test

import (
	"regexp"
	"strings"
	"testing"

	"github.com/lazygophers/utils/country"
	"github.com/lazygophers/utils/fake"
)

func TestLatitudeLongitude_Range(t *testing.T) {
	f := fake.New(country.UnitedStates, fake.WithSeed(1))
	for i := 0; i < 100; i++ {
		lat := f.Latitude()
		if lat < -90 || lat > 90 {
			t.Fatalf("lat %f out of range", lat)
		}
		lng := f.Longitude()
		if lng < -180 || lng > 180 {
			t.Fatalf("lng %f out of range", lng)
		}
	}
}

type zipCase struct {
	name    string
	c       *country.Country
	pattern string
}

func TestZipCode_Formats(t *testing.T) {
	cases := []zipCase{
		{name: "CN", c: country.China, pattern: `^\d{6}$`},
		{name: "US", c: country.UnitedStates, pattern: `^\d{5}$`},
		{name: "JP", c: country.Japan, pattern: `^\d{3}-\d{4}$`},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			f := fake.New(tc.c, fake.WithSeed(2))
			re := regexp.MustCompile(tc.pattern)
			for i := 0; i < 20; i++ {
				z := f.ZipCode()
				if !re.MatchString(z) {
					t.Fatalf("ZipCode %q doesn't match %s", z, tc.pattern)
				}
			}
		})
	}
}

func TestZipCode_DefaultTemplate(t *testing.T) {
	// AD has empty ZipFormat → defaults to 5-digit numeric.
	c := country.Get("AD")
	if c == nil {
		t.Skip("AD country requires country_all build tag")
	}
	f := fake.New(c, fake.WithSeed(3))
	z := f.ZipCode()
	re := regexp.MustCompile(`^\d{5}$`)
	if !re.MatchString(z) {
		t.Fatalf("AD ZipCode default: %q", z)
	}
}

func TestCity_NonEmpty(t *testing.T) {
	for _, alpha2 := range []string{"CN", "US", "JP"} {
		c := country.Get(alpha2)
		f := fake.New(c, fake.WithSeed(4))
		city := f.City()
		if city == "" {
			t.Fatalf("%s City empty", alpha2)
		}
		// CityEntry must also return non-empty Name.
		ce := f.CityEntry()
		if ce.Name == "" {
			t.Fatalf("%s CityEntry empty", alpha2)
		}
	}
}

func TestProvince_NonEmpty(t *testing.T) {
	f := fake.New(country.China, fake.WithSeed(5))
	hits := 0
	for i := 0; i < 10; i++ {
		if f.Province() != "" {
			hits++
		}
	}
	if hits == 0 {
		t.Fatal("Province never returned non-empty")
	}
}

func TestStreet_NonEmpty(t *testing.T) {
	for _, alpha2 := range []string{"CN", "US", "JP"} {
		f := fake.New(country.Get(alpha2), fake.WithSeed(6))
		if f.Street() == "" {
			t.Fatalf("%s Street empty", alpha2)
		}
	}
}

func TestDistrict(t *testing.T) {
	fCn := fake.New(country.China, fake.WithSeed(7))
	d := fCn.District()
	if !strings.HasSuffix(d, "区") {
		t.Fatalf("CN District should end with 区: %q", d)
	}

	fUs := fake.New(country.UnitedStates, fake.WithSeed(7))
	du := fUs.District()
	if !strings.HasSuffix(du, " District") {
		t.Fatalf("US District should end with ' District': %q", du)
	}

	fJp := fake.New(country.Japan, fake.WithSeed(7))
	dj := fJp.District()
	if !strings.HasSuffix(dj, "区") {
		t.Fatalf("JP District should end with 区: %q", dj)
	}
}

func TestStreetAddress(t *testing.T) {
	// CJK: street + " <n>号"
	fCn := fake.New(country.China, fake.WithSeed(8))
	addr := fCn.StreetAddress()
	if !strings.HasSuffix(addr, "号") {
		t.Fatalf("CN StreetAddress should end with 号: %q", addr)
	}
	// US: "<n> Street"
	fUs := fake.New(country.UnitedStates, fake.WithSeed(8))
	addr2 := fUs.StreetAddress()
	parts := strings.SplitN(addr2, " ", 2)
	if len(parts) != 2 {
		t.Fatalf("US StreetAddress: %q", addr2)
	}
	re := regexp.MustCompile(`^\d+$`)
	if !re.MatchString(parts[0]) {
		t.Fatalf("US StreetAddress first token should be number: %q", addr2)
	}
}

type fullAddrCase struct {
	name        string
	c           *country.Country
	containsZip bool
}

func TestFullAddress(t *testing.T) {
	cases := []fullAddrCase{
		{name: "CN", c: country.China},
		{name: "US", c: country.UnitedStates, containsZip: true},
		{name: "JP", c: country.Japan},
		{name: "AD", c: country.Get("AD")},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			if tc.c == nil {
				t.Skip("country not available in current build")
			}
			f := fake.New(tc.c, fake.WithSeed(9))
			addr := f.FullAddress()
			if addr == "" {
				t.Fatal("FullAddress empty")
			}
			if tc.containsZip {
				re := regexp.MustCompile(`\d{5}`)
				if !re.MatchString(addr) {
					t.Fatalf("US FullAddress missing zip: %q", addr)
				}
			}
		})
	}
}
