//go:build country_all || country_americas || country_south_america || country_sr

package country

import (
	xlanguage "golang.org/x/text/language"

	"github.com/lazygophers/utils/currency"
)

// Suriname — Republic of Suriname.
var dataSuriname = &Country{
	alpha2:       "SR",
	alpha3:       "SUR",
	numeric:      740,
	callingCodes: []string{"+597"},
	timezones:    []string{"America/Paramaribo"},
	tlds:         []string{".sr"},
	officialLanguage:  xlanguage.Dutch,
	spokenLanguages:   []xlanguage.Tag{xlanguage.Dutch},
	currency:     currency.SRD,
	region:       RegionSouthAmerica,
	flagEmoji:    "🇸🇷",
	names:        make(map[xlanguage.Tag]string),
	official:     make(map[xlanguage.Tag]string),
	capital:      make(map[xlanguage.Tag]string),
}

func init() { register(dataSuriname) }

var Suriname = dataSuriname
