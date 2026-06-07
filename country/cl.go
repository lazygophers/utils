//go:build country_all || country_americas || country_cl || country_south_america

package country

import (
	xlanguage "golang.org/x/text/language"

	"github.com/lazygophers/utils/currency"
)

// Chile — Republic of Chile.
var dataChile = &Country{
	alpha2:       "CL",
	alpha3:       "CHL",
	numeric:      152,
	callingCodes: []string{"+56"},
	timezones:    []string{
		"America/Santiago",
		"Pacific/Easter",
	},
	tlds:         []string{".cl"},
	officialLanguage:  xlanguage.Spanish,
	spokenLanguages:   []xlanguage.Tag{xlanguage.Spanish},
	currency:     currency.CLP,
	region:       RegionSouthAmerica,
	flagEmoji:    "🇨🇱",
	names:        make(map[xlanguage.Tag]string),
	official:     make(map[xlanguage.Tag]string),
	capital:      make(map[xlanguage.Tag]string),
}

func init() { register(dataChile) }

var Chile = dataChile
