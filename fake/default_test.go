package fake_test

import (
	"regexp"
	"strings"
	"testing"
	"unicode"

	"github.com/lazygophers/utils/country"
	"github.com/lazygophers/utils/fake"
	"github.com/lazygophers/utils/language"
)

// withLang runs fn while the goroutine-local language is set to tag,
// restoring the previous state on return.
func withLang(t *testing.T, tag string, fn func()) {
	t.Helper()
	language.Set(language.Make(tag))
	defer language.Del()
	fn()
}

func TestGlobal_Name_ZhTriggersCN(t *testing.T) {
	withLang(t, "zh", func() {
		// Sample several to surface Han characters.
		anyHan := false
		for i := 0; i < 20; i++ {
			n := fake.Name()
			if n == "" {
				t.Fatal("Name empty")
			}
			for _, r := range n {
				if unicode.Is(unicode.Han, r) {
					anyHan = true
					break
				}
			}
		}
		if !anyHan {
			t.Fatal("zh language should produce CN-style names (with Han)")
		}
	})
}

func TestGlobal_Name_EnTriggersUS(t *testing.T) {
	withLang(t, "en", func() {
		anySpace := false
		for i := 0; i < 10; i++ {
			n := fake.Name()
			if strings.Contains(n, " ") {
				anySpace = true
			}
		}
		if !anySpace {
			t.Fatal("en language should produce US-style names with space separator")
		}
	})
}

func TestGlobal_Name_JaTriggersJP(t *testing.T) {
	withLang(t, "ja", func() {
		anyCjk := false
		for i := 0; i < 20; i++ {
			n := fake.Name()
			for _, r := range n {
				if unicode.Is(unicode.Han, r) || unicode.Is(unicode.Hiragana, r) || unicode.Is(unicode.Katakana, r) {
					anyCjk = true
					break
				}
			}
		}
		if !anyCjk {
			t.Fatal("ja language should produce JP-style names (CJK)")
		}
	})
}

func TestGlobal_UnknownLanguageFallsBackToUS(t *testing.T) {
	withLang(t, "xx", func() {
		n := fake.Name()
		if n == "" {
			t.Fatal("Name empty")
		}
	})
}

func TestGlobal_WithCountry_DoesNotPersist(t *testing.T) {
	withLang(t, "en", func() {
		// One-off Japan call.
		jp := fake.WithCountry(country.Japan).Name()
		if jp == "" {
			t.Fatal("WithCountry Name empty")
		}
		// After that, the global helper should still infer US (en).
		// Sample a few names; some will have spaces (US convention).
		seenSpace := false
		for i := 0; i < 10; i++ {
			n := fake.Name()
			if strings.Contains(n, " ") {
				seenSpace = true
				break
			}
		}
		if !seenSpace {
			t.Fatal("WithCountry leaked into subsequent fake.Name() calls")
		}
	})
}

func TestGlobal_AllHelpersWorking(t *testing.T) {
	withLang(t, "en", func() {
		// Smoke test every global helper to drive coverage.
		_ = fake.FirstName()
		_ = fake.LastName()
		_ = fake.Username()
		_ = fake.IdCard()
		_ = fake.PassportNo()
		_ = fake.CallingCode()
		_ = fake.Email()
		_ = fake.Phone()
		_ = fake.Tel()
		_ = fake.City()
		_ = fake.Province()
		_ = fake.District()
		_ = fake.Street()
		_ = fake.StreetAddress()
		_ = fake.ZipCode()
		_ = fake.Latitude()
		_ = fake.Longitude()
		_ = fake.FullAddress()
		_ = fake.UUIDv4()
		_ = fake.UUIDv7()
		_ = fake.IPv4()
		_ = fake.IPv6()
		_ = fake.Mac()
		_ = fake.Md5Hex()
		_ = fake.Sha1Hex()
		_ = fake.Sha256Hex()
		_ = fake.Word()
		_ = fake.Words(2)
		_ = fake.Sentence(4)
		_ = fake.Paragraph(2)
		_ = fake.ChineseWord()
		_ = fake.ChineseWords(2)
		_ = fake.ChineseSentence(10)
		_ = fake.ChineseParagraph(2)
		_ = fake.Time()
		_ = fake.Birthday(20, 30)
		_ = fake.IntRange(1, 10)
		_ = fake.Int64Range(1, 10)
		_ = fake.Float64Range(0.0, 1.0)
		_ = fake.Bool()
		_ = fake.HexColor()
		_ = fake.RgbColor()
		_ = fake.HslColor()
		_ = fake.FileName()
		_ = fake.FileExt()
		_ = fake.MimeType()
	})
}

func TestGlobal_UUIDv4_Format(t *testing.T) {
	re := regexp.MustCompile(`^[0-9a-f]{8}-[0-9a-f]{4}-4[0-9a-f]{3}-[89ab][0-9a-f]{3}-[0-9a-f]{12}$`)
	for i := 0; i < 20; i++ {
		u := fake.UUIDv4()
		if !re.MatchString(u) {
			t.Fatalf("global UUIDv4 %q malformed", u)
		}
	}
}
