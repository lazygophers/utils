package country

import (
	xlanguage "golang.org/x/text/language"

	"github.com/lazygophers/utils/currency"
)

// Ireland — Republic of Ireland.
var dataIreland = &Country{
	alpha2:       "IE",
	alpha3:       "IRL",
	numeric:      372,
	callingCodes: []string{"+353"},
	timezones:    []string{"Europe/Dublin"},
	tlds:         []string{".ie"},
	languages:    []xlanguage.Tag{xlanguage.English, xlanguage.MustParse("ga")},
	currency:     currency.Eur,
	continent:    "EU",
	region:       "Europe",
	subregion:    "Northern Europe",
	flagEmoji:    "🇮🇪",
	names:        make(map[xlanguage.Tag]string),
	official:     make(map[xlanguage.Tag]string),
	capital:      make(map[xlanguage.Tag]string),
}

func init() { register(dataIreland) }
