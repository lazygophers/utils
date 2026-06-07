//go:build country_africa || country_all || country_eastern_africa || country_sc

package country

import (
	xlanguage "golang.org/x/text/language"

	"github.com/lazygophers/utils/currency"
)

// Seychelles — Republic of Seychelles.
var dataSeychelles = &Country{
	alpha2:       "SC",
	alpha3:       "SYC",
	numeric:      690,
	callingCodes: []string{"+248"},
	timezones:    []string{"Indian/Mahe"},
	tlds:         []string{".sc"},
	officialLanguage:  xlanguage.English,
	spokenLanguages:   []xlanguage.Tag{xlanguage.English, xlanguage.French},
	currency:     currency.SCR,
	region:       RegionEasternAfrica,
	flagEmoji:    "🇸🇨",
	names:        make(map[xlanguage.Tag]string),
	official:     make(map[xlanguage.Tag]string),
	capital:      make(map[xlanguage.Tag]string),
}

func init() { register(dataSeychelles) }

var Seychelles = dataSeychelles
