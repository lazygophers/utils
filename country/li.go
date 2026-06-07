//go:build country_all || country_europe || country_li || country_western_europe

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
	officialLanguage:  xlanguage.German,
	spokenLanguages:   []xlanguage.Tag{xlanguage.German},
	currency:     currency.CHF,
	region:       RegionWesternEurope,
	flagEmoji:    "🇱🇮",
	names:        make(map[xlanguage.Tag]string),
	official:     make(map[xlanguage.Tag]string),
	capital:      make(map[xlanguage.Tag]string),
}

func init() { register(dataLiechtenstein) }

var Liechtenstein = dataLiechtenstein
