//go:build country_all || country_ch || country_europe || country_western_europe

package country

import (
	xlanguage "golang.org/x/text/language"

	"github.com/lazygophers/utils/currency"
)

// Switzerland — Swiss Confederation.
var dataSwitzerland = &Country{
	alpha2:       "CH",
	alpha3:       "CHE",
	numeric:      756,
	callingCodes: []string{"+41"},
	timezones:    []string{"Europe/Zurich"},
	tlds:         []string{".ch"},
	officialLanguage:  xlanguage.German,
	spokenLanguages:   []xlanguage.Tag{xlanguage.German, xlanguage.French, xlanguage.Italian, xlanguage.MustParse("rm")},
	currency:     currency.CHF,
	region:       RegionWesternEurope,
	flagEmoji:    "🇨🇭",
	names:        make(map[xlanguage.Tag]string),
	official:     make(map[xlanguage.Tag]string),
	capital:      make(map[xlanguage.Tag]string),
}

func init() { register(dataSwitzerland) }

var Switzerland = dataSwitzerland
