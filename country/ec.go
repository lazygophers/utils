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
	languages:    []xlanguage.Tag{xlanguage.Spanish},
	currency:     currency.Usd,
	continent:    "SA",
	region:       "Americas",
	subregion:    "South America",
	flagEmoji:    "🇪🇨",
	names:        make(map[xlanguage.Tag]string),
	official:     make(map[xlanguage.Tag]string),
	capital:      make(map[xlanguage.Tag]string),
}

func init() { register(dataEcuador) }
