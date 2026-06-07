package fake_test

import (
	"regexp"
	"strings"
	"testing"

	"github.com/lazygophers/utils/country"
	"github.com/lazygophers/utils/fake"
)

func TestHexColor(t *testing.T) {
	re := regexp.MustCompile(`^#[0-9a-f]{6}$`)
	f := fake.New(country.UnitedStates, fake.WithSeed(1))
	for i := 0; i < 30; i++ {
		c := f.HexColor()
		if !re.MatchString(c) {
			t.Fatalf("HexColor %q malformed", c)
		}
	}
}

func TestRgbColor(t *testing.T) {
	re := regexp.MustCompile(`^rgb\((\d{1,3}), (\d{1,3}), (\d{1,3})\)$`)
	f := fake.New(country.UnitedStates, fake.WithSeed(2))
	for i := 0; i < 30; i++ {
		c := f.RgbColor()
		m := re.FindStringSubmatch(c)
		if m == nil {
			t.Fatalf("RgbColor %q malformed", c)
		}
	}
}

func TestHslColor(t *testing.T) {
	re := regexp.MustCompile(`^hsl\((\d{1,3}), (\d{1,3})%, (\d{1,3})%\)$`)
	f := fake.New(country.UnitedStates, fake.WithSeed(3))
	for i := 0; i < 30; i++ {
		c := f.HslColor()
		if !re.MatchString(c) {
			t.Fatalf("HslColor %q malformed", c)
		}
	}
}

func TestFileName(t *testing.T) {
	f := fake.New(country.UnitedStates, fake.WithSeed(4))
	for i := 0; i < 30; i++ {
		name := f.FileName()
		if !strings.Contains(name, ".") {
			t.Fatalf("FileName missing dot: %q", name)
		}
		parts := strings.Split(name, ".")
		if len(parts) < 2 || parts[len(parts)-1] == "" {
			t.Fatalf("FileName malformed: %q", name)
		}
	}
}

func TestFileExt(t *testing.T) {
	f := fake.New(country.UnitedStates, fake.WithSeed(5))
	re := regexp.MustCompile(`^[a-z0-9]+$`)
	for i := 0; i < 30; i++ {
		ext := f.FileExt()
		if !re.MatchString(ext) {
			t.Fatalf("FileExt malformed: %q", ext)
		}
	}
}

func TestMimeType(t *testing.T) {
	f := fake.New(country.UnitedStates, fake.WithSeed(6))
	for i := 0; i < 30; i++ {
		m := f.MimeType()
		if !strings.Contains(m, "/") {
			t.Fatalf("MimeType missing /: %q", m)
		}
	}
}
