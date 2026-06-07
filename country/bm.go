//go:build country_all || country_americas || country_bm || country_northern_america

package country

import (
	xlanguage "golang.org/x/text/language"

	"github.com/lazygophers/utils/currency"
)

// Bermuda — British Overseas Territory of Bermuda.
var dataBermuda = &Country{
	alpha2:       "BM",
	alpha3:       "BMU",
	numeric:      60,
	callingCodes: []string{"+1-441"},
	timezones:    []string{"Atlantic/Bermuda"},
	tlds:         []string{".bm"},
	officialLanguage:  xlanguage.English,
	spokenLanguages:   []xlanguage.Tag{xlanguage.English},
	currency:     currency.BMD,
	region:       RegionNorthernAmerica,
	flagEmoji:    "🇧🇲",
	names:        make(map[xlanguage.Tag]string),
	official:     make(map[xlanguage.Tag]string),
	capital:      make(map[xlanguage.Tag]string),
}

func init() { register(dataBermuda) }

var Bermuda = dataBermuda
