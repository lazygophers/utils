package country

import (
	xlanguage "golang.org/x/text/language"

	"github.com/lazygophers/utils/currency"
)

// Germany — Federal Republic of Germany.
var dataGermany = &Country{
	alpha2:       "DE",
	alpha3:       "DEU",
	numeric:      276,
	callingCodes: []string{"+49"},
	timezones:    []string{"Europe/Berlin"},
	tlds:         []string{".de"},
	languages:    []xlanguage.Tag{xlanguage.German},
	currency:     currency.Eur,
	continent:    "EU",
	region:       "Europe",
	subregion:    "Western Europe",
	flagEmoji:    "🇩🇪",
	names:        make(map[xlanguage.Tag]string),
	official:     make(map[xlanguage.Tag]string),
	capital:      make(map[xlanguage.Tag]string),
}

func init() { register(dataGermany) }
