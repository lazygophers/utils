package fake

import (
	"strconv"
	"strings"

	xlanguage "golang.org/x/text/language"
)

// districtSuffixZh is the Simplified Chinese suffix appended to derived
// district names so the output reads as a real administrative division
// rather than a bare street fragment.
const districtSuffixZh = "区"

// districtSuffixJa is the Japanese suffix used for derived ward names.
const districtSuffixJa = "区"

// districtSuffixEn is the English suffix used when no native script form
// is available (fallback for ASCII pools).
const districtSuffixEn = " District"

// cjkAlpha2 marks countries whose addresses are written largest-to-smallest
// (province → city → district → street → house number) without separators.
var cjkAlpha2 = map[string]bool{
	"CN": true,
	"TW": true,
	"HK": true,
	"MO": true,
	"JP": true,
	"KR": true,
	"KP": true,
}

// City returns the name of a uniformly chosen city from the resolved city
// pool. Returns the empty string when no city data is available even after
// fallback.
func (f *Faker) City() string {
	return f.CityEntry().Name
}

// CityEntry returns a uniformly chosen [CityEntry] from the resolved city
// pool. The zero value is returned when no city data is registered for
// either the active language, the country's official languages, or the
// English fallback.
func (f *Faker) CityEntry() CityEntry {
	pool := f.resolveCityPool()
	if len(pool) == 0 {
		return CityEntry{}
	}
	return pool[f.intN(len(pool))]
}

// Province returns the administrative parent of a uniformly chosen city
// (e.g. the Chinese province, US state, or Japanese prefecture). Returns
// the empty string when no city data is available.
func (f *Faker) Province() string {
	return f.CityEntry().Province
}

// District returns a synthetic district / ward name composed of a street
// fragment and a locale-appropriate suffix (“区“ for CJK locales,
// “ District“ for ASCII fallbacks). Returns the empty string when no
// street data is available.
func (f *Faker) District() string {
	street := f.pickString(f.resolveStreetPool())
	if street == "" {
		return ""
	}
	suffix := f.districtSuffix()
	return street + suffix
}

// Street returns a uniformly chosen street name from the resolved street
// pool. Returns the empty string when no street data is available.
func (f *Faker) Street() string {
	return f.pickString(f.resolveStreetPool())
}

// StreetAddress combines a random house number with a street name. CJK
// locales render the number after the street with the “号“ suffix
// (e.g. “中山路 123号“); other locales place the number before the
// street (e.g. “742 Evergreen Terrace“).
func (f *Faker) StreetAddress() string {
	street := f.Street()
	if street == "" {
		return ""
	}
	num := f.intN(9999) + 1
	if f.isCJK() {
		return street + " " + strconv.Itoa(num) + "号"
	}
	return strconv.Itoa(num) + " " + street
}

// ZipCode returns a postal code shaped by the locale's [Locale.ZipFormat]
// template. The literal byte “#“ is replaced by a uniformly chosen
// decimal digit; all other runes are preserved. When the locale carries
// no template the result is a 5-digit numeric code.
func (f *Faker) ZipCode() string {
	template := f.locale.ZipFormat
	if template == "" {
		template = "#####"
	}
	return f.formatZip(template)
}

// Latitude returns a uniformly distributed latitude in the inclusive
// range [-90.0, 90.0].
func (f *Faker) Latitude() float64 {
	return f.float64()*180 - 90
}

// Longitude returns a uniformly distributed longitude in the inclusive
// range [-180.0, 180.0].
func (f *Faker) Longitude() float64 {
	return f.float64()*360 - 180
}

// FullAddress composes a complete postal address using the country's
// conventional ordering:
//
//   - CN: “<Province><City><District><Street> <num>号“
//   - JP: “<Province><City><District><Street><num>“
//   - US: “<num> <Street>, <City>, <Province> <Zip>“
//   - other: “<num> <Street>, <City>, <Province>“
//
// Empty fields are skipped so that countries lacking some data still
// produce a useful string.
func (f *Faker) FullAddress() string {
	city := f.CityEntry()
	street := f.Street()
	num := f.intN(9999) + 1
	alpha2 := f.country.Alpha2()

	switch alpha2 {
	case "CN":
		return joinNonEmpty("", city.Province, city.Name, f.districtFromStreet(street, districtSuffixZh), street+" "+strconv.Itoa(num)+"号")
	case "JP":
		return joinNonEmpty("", city.Province, city.Name, f.districtFromStreet(street, districtSuffixJa), street, strconv.Itoa(num))
	case "US":
		head := strconv.Itoa(num)
		if street != "" {
			head += " " + street
		}
		zip := f.ZipCode()
		return joinNonEmpty(", ", head, city.Name, city.Province+" "+zip)
	}

	head := strconv.Itoa(num)
	if street != "" {
		head += " " + street
	}
	return joinNonEmpty(", ", head, city.Name, city.Province)
}

// resolveCityPool walks the language fallback chain looking for a
// non-empty city pool:
//
//  1. the locale at the active language tag
//  2. the locale at each registered official language
//  3. the US locale at English
//
// Returns nil when every layer is empty.
func (f *Faker) resolveCityPool() []CityEntry {
	pool := f.locale.Cities[f.lang]
	if len(pool) > 0 {
		return pool
	}
	for _, tag := range f.locale.OfficialLangs {
		pool = f.locale.Cities[tag]
		if len(pool) > 0 {
			return pool
		}
	}
	us := registry["US"]
	if us == nil {
		return nil
	}
	return us.Cities[xlanguage.English]
}

// resolveStreetPool walks the same fallback chain as
// [Faker.resolveCityPool] for street name templates.
func (f *Faker) resolveStreetPool() []string {
	pool := f.locale.Streets[f.lang]
	if len(pool) > 0 {
		return pool
	}
	for _, tag := range f.locale.OfficialLangs {
		pool = f.locale.Streets[tag]
		if len(pool) > 0 {
			return pool
		}
	}
	us := registry["US"]
	if us == nil {
		return nil
	}
	return us.Streets[xlanguage.English]
}

// formatZip replaces every “#“ rune in template with a random decimal
// digit while preserving all other runes verbatim.
func (f *Faker) formatZip(template string) string {
	var b strings.Builder
	b.Grow(len(template))
	for _, r := range template {
		if r == '#' {
			b.WriteByte(byte('0' + f.intN(10)))
			continue
		}
		b.WriteRune(r)
	}
	return b.String()
}

// districtSuffix returns the suffix appended to street fragments when
// composing a synthetic district name. CJK locales use a native single
// character; other locales fall back to the English literal.
func (f *Faker) districtSuffix() string {
	base, _ := f.lang.Base()
	switch base.String() {
	case "zh":
		return districtSuffixZh
	case "ja":
		return districtSuffixJa
	}
	if f.isCJK() {
		return districtSuffixZh
	}
	return districtSuffixEn
}

// districtFromStreet derives a district label from a street fragment by
// appending the supplied suffix. Returns the empty string when street is
// empty so callers can skip the segment cleanly.
func (f *Faker) districtFromStreet(street, suffix string) string {
	if street == "" {
		return ""
	}
	return f.pickString(f.resolveStreetPool()) + suffix
}

// isCJK reports whether the faker's country uses the CJK address ordering
// (largest administrative unit first, no comma separators).
func (f *Faker) isCJK() bool {
	return cjkAlpha2[f.country.Alpha2()]
}

// joinNonEmpty concatenates parts with sep, skipping any empty segments
// so locales missing some fields still produce a clean string.
func joinNonEmpty(sep string, parts ...string) string {
	out := make([]string, 0, len(parts))
	for _, p := range parts {
		trimmed := strings.TrimSpace(p)
		if trimmed == "" {
			continue
		}
		out = append(out, p)
	}
	return strings.Join(out, sep)
}
