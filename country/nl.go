package country

import (
	xlanguage "golang.org/x/text/language"

	"github.com/lazygophers/utils/currency"
)

// Netherlands — Kingdom of the Netherlands.
var dataNetherlands = &Country{
	alpha2:       "NL",
	alpha3:       "NLD",
	numeric:      528,
	callingCodes: []string{"+31"},
	timezones:    []string{"Europe/Amsterdam"},
	tlds:         []string{".nl"},
	languages:    []xlanguage.Tag{xlanguage.Dutch},
	currency:     currency.Eur,
	continent:    "EU",
	region:       "Europe",
	subregion:    "Western Europe",
	flagEmoji:    "🇳🇱",
	names:        make(map[xlanguage.Tag]string),
	official:     make(map[xlanguage.Tag]string),
	capital:      make(map[xlanguage.Tag]string),
}

func init() { register(dataNetherlands) }
