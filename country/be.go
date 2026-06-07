package country

import (
	xlanguage "golang.org/x/text/language"

	"github.com/lazygophers/utils/currency"
)

// Belgium — Kingdom of Belgium.
var dataBelgium = &Country{
	alpha2:       "BE",
	alpha3:       "BEL",
	numeric:      56,
	callingCodes: []string{"+32"},
	timezones:    []string{"Europe/Brussels"},
	tlds:         []string{".be"},
	languages:    []xlanguage.Tag{xlanguage.Dutch, xlanguage.French, xlanguage.German},
	currency:     currency.Eur,
	continent:    "EU",
	region:       "Europe",
	subregion:    "Western Europe",
	flagEmoji:    "🇧🇪",
	names:        make(map[xlanguage.Tag]string),
	official:     make(map[xlanguage.Tag]string),
	capital:      make(map[xlanguage.Tag]string),
}

func init() { register(dataBelgium) }
