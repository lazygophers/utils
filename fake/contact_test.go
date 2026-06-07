package fake_test

import (
	"regexp"
	"strings"
	"testing"

	"github.com/lazygophers/utils/country"
	"github.com/lazygophers/utils/fake"
)

type phoneCase struct {
	name    string
	c       *country.Country
	pattern string
}

func TestPhone_Formats(t *testing.T) {
	cases := []phoneCase{
		{name: "CN", c: country.China, pattern: `^\+86 1[3-9]\d-\d{4}-\d{4}$`},
		{name: "US", c: country.UnitedStates, pattern: `^\+1 \(\d{3}\) \d{3}-\d{4}$`},
		{name: "JP", c: country.Japan, pattern: `^\+81 \d{2}-\d{4}-\d{4}$`},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			f := fake.New(tc.c, fake.WithSeed(42))
			re := regexp.MustCompile(tc.pattern)
			for i := 0; i < 30; i++ {
				p := f.Phone()
				if !re.MatchString(p) {
					t.Fatalf("Phone %q doesn't match %s", p, tc.pattern)
				}
			}
		})
	}
}

func TestPhone_Default(t *testing.T) {
	// AD has calling codes (+376), no PhonePrefixes — falls through to default branch.
	c := country.Get("AD")
	if c == nil {
		t.Skip("AD country requires country_all build tag")
	}
	f := fake.New(c, fake.WithSeed(2))
	p := f.Phone()
	if !strings.HasPrefix(p, "+376 ") {
		t.Fatalf("AD phone expected to start with calling code: %q", p)
	}
}

func TestTel_Formats(t *testing.T) {
	cases := []phoneCase{
		{name: "CN", c: country.China, pattern: `^\+86 0\d{2}-\d{4}-\d{4}$`},
		{name: "US", c: country.UnitedStates, pattern: `^\+1 \(\d{3}\) \d{3}-\d{4}$`},
		{name: "JP", c: country.Japan, pattern: `^\+81 \d{1,2}-\d{4}-\d{4}$`},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			f := fake.New(tc.c, fake.WithSeed(11))
			re := regexp.MustCompile(tc.pattern)
			for i := 0; i < 20; i++ {
				p := f.Tel()
				if !re.MatchString(p) {
					t.Fatalf("Tel %q doesn't match %s", p, tc.pattern)
				}
			}
		})
	}
}

func TestTel_Default(t *testing.T) {
	c := country.Get("AD")
	if c == nil {
		t.Skip("AD country requires country_all build tag")
	}
	f := fake.New(c, fake.WithSeed(3))
	p := f.Tel()
	if !strings.HasPrefix(p, "+376 ") {
		t.Fatalf("AD tel expected calling code prefix: %q", p)
	}
}

func TestCallingCode(t *testing.T) {
	tests := []phoneCase{
		{name: "CN", c: country.China, pattern: `^\+86$`},
		{name: "US", c: country.UnitedStates, pattern: `^\+1$`},
		{name: "JP", c: country.Japan, pattern: `^\+81$`},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			f := fake.New(tc.c)
			got := f.CallingCode()
			re := regexp.MustCompile(tc.pattern)
			if !re.MatchString(got) {
				t.Fatalf("CallingCode %q doesn't match %s", got, tc.pattern)
			}
		})
	}
}

func TestEmail_Format(t *testing.T) {
	// RFC 5322 basic shape: local@domain with local being [a-z0-9.] and domain being a known set.
	re := regexp.MustCompile(`^[a-z0-9]+(\.[a-z0-9]+)*@[a-z0-9.]+$`)
	f := fake.New(country.UnitedStates, fake.WithSeed(13))
	for i := 0; i < 30; i++ {
		e := f.Email()
		if !re.MatchString(e) {
			t.Fatalf("Email %q doesn't match basic shape", e)
		}
		if !strings.Contains(e, "@") {
			t.Fatalf("Email %q missing @", e)
		}
	}
}
