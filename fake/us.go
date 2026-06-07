package fake

import (
	"fmt"
	"math/rand/v2"
	"time"

	xlanguage "golang.org/x/text/language"

	"github.com/lazygophers/utils/country"
)

// localeUS is the registered [Locale] for the United States. Language-specific
// pools (street, city, name) are populated by sibling files such as us_en.go.
var localeUS = &Locale{
	Country:       country.UnitedStates,
	OfficialLangs: []xlanguage.Tag{xlanguage.English},
	PhonePrefixes: []string{
		"201", "202", "203", "205", "206", "207", "208", "209", "210", "212",
		"213", "214", "215", "216", "217", "218", "301", "302", "303", "304",
		"305", "404", "405", "407", "408", "409", "410", "415", "417", "424",
		"425", "504", "505", "510", "512", "515", "516", "517", "518", "601",
		"602", "603", "605", "606", "607", "608", "609", "610", "612", "614",
		"615", "617", "618", "619", "630", "631", "636", "641", "646", "650",
		"651", "657", "660", "661", "662", "678", "701", "702", "703", "704",
		"706", "707", "708", "712", "713", "714", "715", "716", "717", "718",
		"719", "720", "724", "727", "731", "732", "734", "737", "740", "743",
		"747", "754", "757", "760", "762", "763", "765", "770", "771", "772",
		"773", "774", "775", "779", "781", "785", "786", "801", "802", "803",
		"804", "805", "806", "808", "810", "812", "813", "814", "815", "816",
		"817", "818", "828", "830", "831", "832", "843", "845", "847", "848",
		"850", "856", "857", "858", "859", "860", "862", "863", "864", "865",
		"870", "872", "878", "901", "903", "904", "906", "907", "908", "909",
		"910", "912", "913", "914", "915", "916", "917", "918", "919", "920",
		"925", "928", "929", "931", "934", "936", "937", "940", "941", "947",
		"949", "951", "952", "954", "956", "970", "971", "972", "973", "978",
		"979", "980", "984", "985", "989",
	},
	LandlinePrefix: nil,
	ZipFormat:      "#####",
	IdCardGen:      genSsnUS,
	Streets:        map[xlanguage.Tag][]string{},
	Cities:         map[xlanguage.Tag][]CityEntry{},
	FirstNames:     map[xlanguage.Tag]map[Gender][]string{},
	LastNames:      map[xlanguage.Tag][]string{},
	Domain:         "us",
	UserAgents:     defaultUserAgents,
}

func init() {
	register(localeUS)
}

// genSsnUS returns a syntactically valid US Social Security Number in the
// canonical AAA-GG-SSSS shape. The generator is shape-only and intentionally
// excludes reserved area numbers (000, 666, 900-999) so callers cannot mistake
// the output for a real-world identifier. gender and birth carry no semantic
// weight in the SSN scheme and are ignored.
func genSsnUS(rng *rand.Rand, _ Gender, _ time.Time) string {
	r := rng
	if r == nil {
		r = nil // use math/rand/v2 package-level helpers below
	}

	area := pickSsnArea(r)
	group := pickSsnGroup(r)
	serial := pickSsnSerial(r)

	return fmt.Sprintf("%03d-%02d-%04d", area, group, serial)
}

// pickSsnArea draws an area number from the SSA-assignable ranges 001-665 or
// 667-899, skipping the reserved 666 and 900-999 blocks.
func pickSsnArea(r *rand.Rand) int {
	// 665 values in [1,665] + 233 values in [667,899] = 898 valid options.
	const lowMax = 665
	const total = 898
	n := randIntN(r, total)
	if n < lowMax {
		return n + 1
	}
	return 667 + (n - lowMax)
}

// pickSsnGroup returns a group number in the inclusive range [1,99].
func pickSsnGroup(r *rand.Rand) int {
	return randIntN(r, 99) + 1
}

// pickSsnSerial returns a serial number in the inclusive range [1,9999].
func pickSsnSerial(r *rand.Rand) int {
	return randIntN(r, 9999) + 1
}

// randIntN returns a non-negative pseudo-random int in [0,n) using the
// supplied source, falling back to the rand/v2 global source when nil.
func randIntN(r *rand.Rand, n int) int {
	if r == nil {
		return rand.IntN(n)
	}
	return r.IntN(n)
}
