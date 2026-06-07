//go:build country_all || country_americas || country_bz || country_central_america

package country

import (
	xlanguage "golang.org/x/text/language"

	"github.com/lazygophers/utils/currency"
)

// Belize — Belize.
var dataBelize = &Country{
	alpha2:       "BZ",
	alpha3:       "BLZ",
	numeric:      84,
	callingCodes: []string{"+501"},
	timezones:    []string{"America/Belize"},
	tlds:         []string{".bz"},
	officialLanguage:  xlanguage.English,
	spokenLanguages:   []xlanguage.Tag{xlanguage.English},
	currency:     currency.BZD,
	region:       RegionCentralAmerica,
	flagEmoji:    "🇧🇿",
	names:        make(map[xlanguage.Tag]string),
	official:     make(map[xlanguage.Tag]string),
	capital:      make(map[xlanguage.Tag]string),
}

func init() { register(dataBelize) }

var Belize = dataBelize
