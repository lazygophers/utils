//go:build country_all || country_americas || country_ec || country_south_america

package country

import (
	xlanguage "golang.org/x/text/language"

	"github.com/lazygophers/utils/currency"
)

// Ecuador — Republic of Ecuador.
var dataEcuador = &Country{
	alpha2:       "EC",
	alpha3:       "ECU",
	numeric:      218,
	callingCodes: []string{"+593"},
	timezones:    []string{
		"America/Guayaquil",
		"Pacific/Galapagos",
	},
	tlds:         []string{".ec"},
	officialLanguage:  xlanguage.Spanish,
	spokenLanguages:   []xlanguage.Tag{xlanguage.Spanish},
	currency:     currency.USD,
	region:       RegionSouthAmerica,
	flagEmoji:    "🇪🇨",
	names:        make(map[xlanguage.Tag]string),
	official:     make(map[xlanguage.Tag]string),
	capital:      make(map[xlanguage.Tag]string),
}

func init() { register(dataEcuador) }

var Ecuador = dataEcuador
