package fake

import (
	"sync"
	"time"

	"github.com/lazygophers/utils/country"
	languagepkg "github.com/lazygophers/utils/language"
)

// countryByLangBase maps a base ISO 639-1 language code to the ISO 3166-1
// alpha-2 country code chosen to represent that language in the global
// helpers. The mapping is intentionally coarse — only the dominant country
// for each language is selected — because the global API exists to give
// callers a sensible zero-config default, not to model multilingual regions.
var countryByLangBase = map[string]string{
	"zh": "CN",
	"en": "US",
	"ja": "JP",
	"ko": "KR",
	"ru": "RU",
	"de": "DE",
	"fr": "FR",
	"es": "ES",
	"ar": "SA",
	"pt": "BR",
	"it": "IT",
	"nl": "NL",
	"pl": "PL",
	"tr": "TR",
	"vi": "VN",
	"th": "TH",
	"id": "ID",
	"hi": "IN",
	"uk": "UA",
	"el": "GR",
	"sv": "SE",
	"no": "NO",
	"da": "DK",
	"fi": "FI",
	"he": "IL",
}

// defaultFakers caches one [Faker] per country (keyed by alpha-2 code) so
// global helpers do not allocate on every call. Entries are lazily created
// the first time a given country is observed.
var defaultFakers sync.Map // map[string]*Faker

// inferCountry resolves the goroutine-local language to a [*country.Country]
// using [countryByLangBase]. Falls back to the United States when the base
// language has no mapping or the resolved country is not registered.
func inferCountry() *country.Country {
	tag := languagepkg.Get()
	base := tag.Base()
	code, ok := countryByLangBase[base]
	if ok {
		c := country.Get(code)
		if c != nil {
			return c
		}
	}
	return country.Get("US")
}

// defaultFaker returns the cached [Faker] for the given country, creating
// one on demand. The returned Faker is shared across goroutines; it relies
// on the package-global math/rand/v2 source which is itself goroutine-safe.
func defaultFaker(c *country.Country) *Faker {
	key := c.Alpha2()
	v, ok := defaultFakers.Load(key)
	if ok {
		return v.(*Faker)
	}
	f := New(c)
	actual, _ := defaultFakers.LoadOrStore(key, f)
	return actual.(*Faker)
}

// WithCountry returns a [Faker] bound to the given country, drawing from the
// same shared pool used by the package-level helpers. It does not touch the
// goroutine-local language state, so callers can invoke it inline to obtain
// data from a different country without affecting subsequent global calls
// (e.g. “fake.WithCountry(country.Japan).Name()“).
func WithCountry(c *country.Country) *Faker {
	return defaultFaker(c)
}

// Name returns [Faker.Name] from the goroutine's default faker.
func Name() string { return defaultFaker(inferCountry()).Name() }

// FirstName returns [Faker.FirstName] from the goroutine's default faker.
func FirstName() string { return defaultFaker(inferCountry()).FirstName() }

// LastName returns [Faker.LastName] from the goroutine's default faker.
func LastName() string { return defaultFaker(inferCountry()).LastName() }

// Username returns [Faker.Username] from the goroutine's default faker.
func Username() string { return defaultFaker(inferCountry()).Username() }

// IdCard returns [Faker.IdCard] from the goroutine's default faker.
func IdCard() string { return defaultFaker(inferCountry()).IdCard() }

// PassportNo returns [Faker.PassportNo] from the goroutine's default faker.
func PassportNo() string { return defaultFaker(inferCountry()).PassportNo() }

// CallingCode returns [Faker.CallingCode] from the goroutine's default faker.
func CallingCode() string { return defaultFaker(inferCountry()).CallingCode() }

// Email returns [Faker.Email] from the goroutine's default faker.
func Email() string { return defaultFaker(inferCountry()).Email() }

// Phone returns [Faker.Phone] from the goroutine's default faker.
func Phone() string { return defaultFaker(inferCountry()).Phone() }

// Tel returns [Faker.Tel] from the goroutine's default faker.
func Tel() string { return defaultFaker(inferCountry()).Tel() }

// City returns [Faker.City] from the goroutine's default faker.
func City() string { return defaultFaker(inferCountry()).City() }

// Province returns [Faker.Province] from the goroutine's default faker.
func Province() string { return defaultFaker(inferCountry()).Province() }

// District returns [Faker.District] from the goroutine's default faker.
func District() string { return defaultFaker(inferCountry()).District() }

// Street returns [Faker.Street] from the goroutine's default faker.
func Street() string { return defaultFaker(inferCountry()).Street() }

// StreetAddress returns [Faker.StreetAddress] from the goroutine's default faker.
func StreetAddress() string { return defaultFaker(inferCountry()).StreetAddress() }

// ZipCode returns [Faker.ZipCode] from the goroutine's default faker.
func ZipCode() string { return defaultFaker(inferCountry()).ZipCode() }

// Latitude returns [Faker.Latitude] from the goroutine's default faker.
func Latitude() float64 { return defaultFaker(inferCountry()).Latitude() }

// Longitude returns [Faker.Longitude] from the goroutine's default faker.
func Longitude() float64 { return defaultFaker(inferCountry()).Longitude() }

// FullAddress returns [Faker.FullAddress] from the goroutine's default faker.
func FullAddress() string { return defaultFaker(inferCountry()).FullAddress() }

// UUIDv4 returns [Faker.UUIDv4] from the goroutine's default faker.
func UUIDv4() string { return defaultFaker(inferCountry()).UUIDv4() }

// UUIDv7 returns [Faker.UUIDv7] from the goroutine's default faker.
func UUIDv7() string { return defaultFaker(inferCountry()).UUIDv7() }

// IPv4 returns [Faker.IPv4] from the goroutine's default faker.
func IPv4() string { return defaultFaker(inferCountry()).IPv4() }

// IPv6 returns [Faker.IPv6] from the goroutine's default faker.
func IPv6() string { return defaultFaker(inferCountry()).IPv6() }

// Mac returns [Faker.Mac] from the goroutine's default faker.
func Mac() string { return defaultFaker(inferCountry()).Mac() }

// Md5Hex returns [Faker.Md5Hex] from the goroutine's default faker.
func Md5Hex() string { return defaultFaker(inferCountry()).Md5Hex() }

// Sha1Hex returns [Faker.Sha1Hex] from the goroutine's default faker.
func Sha1Hex() string { return defaultFaker(inferCountry()).Sha1Hex() }

// Sha256Hex returns [Faker.Sha256Hex] from the goroutine's default faker.
func Sha256Hex() string { return defaultFaker(inferCountry()).Sha256Hex() }

// Word returns [Faker.Word] from the goroutine's default faker.
func Word() string { return defaultFaker(inferCountry()).Word() }

// Words returns [Faker.Words] from the goroutine's default faker.
func Words(n int) []string { return defaultFaker(inferCountry()).Words(n) }

// Sentence returns [Faker.Sentence] from the goroutine's default faker.
func Sentence(wordCount int) string { return defaultFaker(inferCountry()).Sentence(wordCount) }

// Paragraph returns [Faker.Paragraph] from the goroutine's default faker.
func Paragraph(sentenceCount int) string {
	return defaultFaker(inferCountry()).Paragraph(sentenceCount)
}

// ChineseWord returns [Faker.ChineseWord] from the goroutine's default faker.
func ChineseWord() string { return defaultFaker(inferCountry()).ChineseWord() }

// ChineseWords returns [Faker.ChineseWords] from the goroutine's default faker.
func ChineseWords(n int) []string { return defaultFaker(inferCountry()).ChineseWords(n) }

// ChineseSentence returns [Faker.ChineseSentence] from the goroutine's default faker.
func ChineseSentence(charCount int) string {
	return defaultFaker(inferCountry()).ChineseSentence(charCount)
}

// ChineseParagraph returns [Faker.ChineseParagraph] from the goroutine's default faker.
func ChineseParagraph(sentenceCount int) string {
	return defaultFaker(inferCountry()).ChineseParagraph(sentenceCount)
}

// Date returns [Faker.Date] from the goroutine's default faker.
func Date(min, max time.Time) time.Time { return defaultFaker(inferCountry()).Date(min, max) }

// Time returns [Faker.Time] from the goroutine's default faker.
func Time() time.Time { return defaultFaker(inferCountry()).Time() }

// Birthday returns [Faker.Birthday] from the goroutine's default faker.
func Birthday(minAge, maxAge int) time.Time {
	return defaultFaker(inferCountry()).Birthday(minAge, maxAge)
}

// IntRange returns [Faker.IntRange] from the goroutine's default faker.
func IntRange(min, max int) int { return defaultFaker(inferCountry()).IntRange(min, max) }

// Int64Range returns [Faker.Int64Range] from the goroutine's default faker.
func Int64Range(min, max int64) int64 { return defaultFaker(inferCountry()).Int64Range(min, max) }

// Float64Range returns [Faker.Float64Range] from the goroutine's default faker.
func Float64Range(min, max float64) float64 {
	return defaultFaker(inferCountry()).Float64Range(min, max)
}

// Bool returns [Faker.Bool] from the goroutine's default faker.
func Bool() bool { return defaultFaker(inferCountry()).Bool() }

// HexColor returns [Faker.HexColor] from the goroutine's default faker.
func HexColor() string { return defaultFaker(inferCountry()).HexColor() }

// RgbColor returns [Faker.RgbColor] from the goroutine's default faker.
func RgbColor() string { return defaultFaker(inferCountry()).RgbColor() }

// HslColor returns [Faker.HslColor] from the goroutine's default faker.
func HslColor() string { return defaultFaker(inferCountry()).HslColor() }

// FileName returns [Faker.FileName] from the goroutine's default faker.
func FileName() string { return defaultFaker(inferCountry()).FileName() }

// FileExt returns [Faker.FileExt] from the goroutine's default faker.
func FileExt() string { return defaultFaker(inferCountry()).FileExt() }

// MimeType returns [Faker.MimeType] from the goroutine's default faker.
func MimeType() string { return defaultFaker(inferCountry()).MimeType() }

// =================================
// User-Agent
// =================================

// UserAgent returns [Faker.UserAgent] from the goroutine's default faker.
func UserAgent() string { return defaultFaker(inferCountry()).UserAgent() }

// BrowserUA returns [Faker.BrowserUA] from the goroutine's default faker.
func BrowserUA() string { return defaultFaker(inferCountry()).BrowserUA() }

// BrowserUAOf returns [Faker.BrowserUAOf] from the goroutine's default faker.
func BrowserUAOf(os OS, br Browser) string {
	return defaultFaker(inferCountry()).BrowserUAOf(os, br)
}

// DesktopUA returns [Faker.DesktopUA] from the goroutine's default faker.
func DesktopUA() string { return defaultFaker(inferCountry()).DesktopUA() }

// MobileUA returns [Faker.MobileUA] from the goroutine's default faker.
func MobileUA() string { return defaultFaker(inferCountry()).MobileUA() }

// AppUA returns [Faker.AppUA] from the goroutine's default faker.
func AppUA() string { return defaultFaker(inferCountry()).AppUA() }

// CLIUA returns [Faker.CLIUA] from the goroutine's default faker.
func CLIUA() string { return defaultFaker(inferCountry()).CLIUA() }

// ProxyUA returns [Faker.ProxyUA] from the goroutine's default faker.
func ProxyUA() string { return defaultFaker(inferCountry()).ProxyUA() }

// =================================
// HTTP Header
// =================================

// Accept returns [Faker.Accept] from the goroutine's default faker.
func Accept() string { return defaultFaker(inferCountry()).Accept() }

// AcceptLanguage returns [Faker.AcceptLanguage] from the goroutine's default faker.
func AcceptLanguage() string { return defaultFaker(inferCountry()).AcceptLanguage() }

// AcceptEncoding returns [Faker.AcceptEncoding] from the goroutine's default faker.
func AcceptEncoding() string { return defaultFaker(inferCountry()).AcceptEncoding() }

// Referer returns [Faker.Referer] from the goroutine's default faker.
func Referer() string { return defaultFaker(inferCountry()).Referer() }

// Header returns [Faker.Header] from the goroutine's default faker.
func Header() map[string]string { return defaultFaker(inferCountry()).Header() }
