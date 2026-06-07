package fake_test

import (
	"net"
	"regexp"
	"strconv"
	"strings"
	"testing"
	"time"

	"github.com/lazygophers/utils/country"
	"github.com/lazygophers/utils/fake"
)

func TestUUIDv4_Format(t *testing.T) {
	re := regexp.MustCompile(`^[0-9a-f]{8}-[0-9a-f]{4}-4[0-9a-f]{3}-[89ab][0-9a-f]{3}-[0-9a-f]{12}$`)
	f := fake.New(country.UnitedStates, fake.WithSeed(1))
	for i := 0; i < 100; i++ {
		u := f.UUIDv4()
		if !re.MatchString(u) {
			t.Fatalf("UUIDv4 %q not RFC 4122 v4", u)
		}
	}
}

func TestUUIDv7_Format(t *testing.T) {
	re := regexp.MustCompile(`^[0-9a-f]{8}-[0-9a-f]{4}-7[0-9a-f]{3}-[89ab][0-9a-f]{3}-[0-9a-f]{12}$`)
	f := fake.New(country.UnitedStates, fake.WithSeed(1))
	for i := 0; i < 20; i++ {
		u := f.UUIDv7()
		if !re.MatchString(u) {
			t.Fatalf("UUIDv7 %q not v7 shape", u)
		}
		// First 48 bits encode milliseconds; decode and ensure within recent window.
		hexms := strings.ReplaceAll(u[:13], "-", "")
		v, err := strconv.ParseInt(hexms, 16, 64)
		if err != nil {
			t.Fatalf("parse ms: %v", err)
		}
		now := time.Now().UnixMilli()
		if v <= 0 || v > now+1000 {
			t.Fatalf("UUIDv7 ms %d not plausible (now=%d)", v, now)
		}
	}
}

func TestIPv4_Format(t *testing.T) {
	f := fake.New(country.UnitedStates, fake.WithSeed(2))
	for i := 0; i < 50; i++ {
		s := f.IPv4()
		parts := strings.Split(s, ".")
		if len(parts) != 4 {
			t.Fatalf("IPv4 %q has %d segments", s, len(parts))
		}
		for _, p := range parts {
			n, err := strconv.Atoi(p)
			if err != nil {
				t.Fatalf("bad octet %q in %q: %v", p, s, err)
			}
			if n < 0 || n > 255 {
				t.Fatalf("octet %d out of range in %q", n, s)
			}
		}
		if net.ParseIP(s) == nil {
			t.Fatalf("net.ParseIP rejected %q", s)
		}
	}
}

func TestIPv6_Format(t *testing.T) {
	f := fake.New(country.UnitedStates, fake.WithSeed(3))
	for i := 0; i < 30; i++ {
		s := f.IPv6()
		parts := strings.Split(s, ":")
		if len(parts) != 8 {
			t.Fatalf("IPv6 %q has %d segments", s, len(parts))
		}
		for _, p := range parts {
			if len(p) != 4 {
				t.Fatalf("group %q not 4 hex digits in %q", p, s)
			}
		}
		if net.ParseIP(s) == nil {
			t.Fatalf("net.ParseIP rejected %q", s)
		}
	}
}

func TestMac_Format(t *testing.T) {
	re := regexp.MustCompile(`^([0-9a-f]{2}:){5}[0-9a-f]{2}$`)
	f := fake.New(country.UnitedStates, fake.WithSeed(4))
	for i := 0; i < 30; i++ {
		m := f.Mac()
		if !re.MatchString(m) {
			t.Fatalf("MAC %q malformed", m)
		}
		_, err := net.ParseMAC(m)
		if err != nil {
			t.Fatalf("net.ParseMAC: %v", err)
		}
	}
}

type hashCase struct {
	name string
	fn   func(*fake.Faker) string
	want int
}

func TestHashes_Length(t *testing.T) {
	cases := []hashCase{
		{name: "Md5", fn: func(f *fake.Faker) string { return f.Md5Hex() }, want: 32},
		{name: "Sha1", fn: func(f *fake.Faker) string { return f.Sha1Hex() }, want: 40},
		{name: "Sha256", fn: func(f *fake.Faker) string { return f.Sha256Hex() }, want: 64},
	}
	re := regexp.MustCompile(`^[0-9a-f]+$`)
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			f := fake.New(country.UnitedStates, fake.WithSeed(5))
			for i := 0; i < 5; i++ {
				h := tc.fn(f)
				if len(h) != tc.want {
					t.Fatalf("%s length = %d, want %d", tc.name, len(h), tc.want)
				}
				if !re.MatchString(h) {
					t.Fatalf("%s not lowercase hex: %q", tc.name, h)
				}
			}
		})
	}
}
