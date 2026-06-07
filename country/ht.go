package country

import (
	xlanguage "golang.org/x/text/language"

	"github.com/lazygophers/utils/currency"
)

// Haiti — Republic of Haiti.
var dataHaiti = &Country{
	alpha2:       "HT",
	alpha3:       "HTI",
	numeric:      332,
	callingCodes: []string{"+509"},
	timezones:    []string{"America/Port-au-Prince"},
	tlds:         []string{".ht"},
	languages:    []xlanguage.Tag{xlanguage.French},
	currency:     currency.Htg,
	continent:    "NA",
	region:       "Americas",
	subregion:    "Caribbean",
	flagEmoji:    "🇭🇹",
	names:        make(map[xlanguage.Tag]string),
	official:     make(map[xlanguage.Tag]string),
	capital:      make(map[xlanguage.Tag]string),
}

func init() { register(dataHaiti) }
