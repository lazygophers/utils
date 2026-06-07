package country

import (
	xlanguage "golang.org/x/text/language"

	"github.com/lazygophers/utils/currency"
)

// Liechtenstein — Principality of Liechtenstein.
var dataLiechtenstein = &Country{
	alpha2:       "LI",
	alpha3:       "LIE",
	numeric:      438,
	callingCodes: []string{"+423"},
	timezones:    []string{"Europe/Vaduz"},
	tlds:         []string{".li"},
	languages:    []xlanguage.Tag{xlanguage.German},
	currency:     currency.Chf,
	continent:    "EU",
	region:       "Europe",
	subregion:    "Western Europe",
	flagEmoji:    "🇱🇮",
	names:        make(map[xlanguage.Tag]string),
	official:     make(map[xlanguage.Tag]string),
	capital:      make(map[xlanguage.Tag]string),
}

func init() { register(dataLiechtenstein) }
