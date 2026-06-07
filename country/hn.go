package country

import (
	xlanguage "golang.org/x/text/language"

	"github.com/lazygophers/utils/currency"
)

// Honduras — Republic of Honduras.
var dataHonduras = &Country{
	alpha2:       "HN",
	alpha3:       "HND",
	numeric:      340,
	callingCodes: []string{"+504"},
	timezones:    []string{"America/Tegucigalpa"},
	tlds:         []string{".hn"},
	languages:    []xlanguage.Tag{xlanguage.Spanish},
	currency:     currency.Hnl,
	continent:    "NA",
	region:       "Americas",
	subregion:    "Central America",
	flagEmoji:    "🇭🇳",
	names:        make(map[xlanguage.Tag]string),
	official:     make(map[xlanguage.Tag]string),
	capital:      make(map[xlanguage.Tag]string),
}

func init() { register(dataHonduras) }
