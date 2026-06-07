//go:build country_ad || country_all || country_europe || country_southern_europe

package country

import (
	xlanguage "golang.org/x/text/language"

	"github.com/lazygophers/utils/currency"
)

// Andorra — Principality of Andorra.
var dataAndorra = &Country{
	alpha2:       "AD",
	alpha3:       "AND",
	numeric:      20,
	callingCodes: []string{"+376"},
	timezones:    []string{"Europe/Andorra"},
	tlds:         []string{".ad"},
	languages:    []xlanguage.Tag{xlanguage.Catalan},
	currency:     currency.Eur,
	continent:    "EU",
	region:       "Europe",
	subregion:    "Southern Europe",
	flagEmoji:    "🇦🇩",
	names:        make(map[xlanguage.Tag]string),
	official:     make(map[xlanguage.Tag]string),
	capital:      make(map[xlanguage.Tag]string),
}

func init() { register(dataAndorra) }

var Andorra = dataAndorra
