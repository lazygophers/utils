//go:build country_all || country_europe || country_lt || country_northern_europe

package country

import (
	xlanguage "golang.org/x/text/language"

	"github.com/lazygophers/utils/currency"
)

// Lithuania — Republic of Lithuania.
var dataLithuania = &Country{
	alpha2:       "LT",
	alpha3:       "LTU",
	numeric:      440,
	callingCodes: []string{"+370"},
	timezones:    []string{"Europe/Vilnius"},
	tlds:         []string{".lt"},
	languages:    []xlanguage.Tag{xlanguage.Lithuanian},
	currency:     currency.Eur,
	continent:    "EU",
	region:       "Europe",
	subregion:    "Northern Europe",
	flagEmoji:    "🇱🇹",
	names:        make(map[xlanguage.Tag]string),
	official:     make(map[xlanguage.Tag]string),
	capital:      make(map[xlanguage.Tag]string),
}

func init() { register(dataLithuania) }

var Lithuania = dataLithuania
