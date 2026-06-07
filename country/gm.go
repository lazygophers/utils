//go:build country_africa || country_all || country_gm || country_western_africa

package country

import (
	xlanguage "golang.org/x/text/language"

	"github.com/lazygophers/utils/currency"
)

// Gambia — Republic of The Gambia.
var dataGambia = &Country{
	alpha2:       "GM",
	alpha3:       "GMB",
	numeric:      270,
	callingCodes: []string{"+220"},
	timezones:    []string{"Africa/Banjul"},
	tlds:         []string{".gm"},
	officialLanguage:  xlanguage.English,
	spokenLanguages:   []xlanguage.Tag{xlanguage.English},
	currency:     currency.GMD,
	region:       RegionWesternAfrica,
	flagEmoji:    "🇬🇲",
	names:        make(map[xlanguage.Tag]string),
	official:     make(map[xlanguage.Tag]string),
	capital:      make(map[xlanguage.Tag]string),
}

func init() { register(dataGambia) }

var Gambia = dataGambia
