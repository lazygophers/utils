package country

import (
	xlanguage "golang.org/x/text/language"

	"github.com/lazygophers/utils/currency"
)

// Paraguay — Republic of Paraguay.
var dataParaguay = &Country{
	alpha2:       "PY",
	alpha3:       "PRY",
	numeric:      600,
	callingCodes: []string{"+595"},
	timezones:    []string{"America/Asuncion"},
	tlds:         []string{".py"},
	languages:    []xlanguage.Tag{xlanguage.Spanish},
	currency:     currency.Pyg,
	continent:    "SA",
	region:       "Americas",
	subregion:    "South America",
	flagEmoji:    "🇵🇾",
	names:        make(map[xlanguage.Tag]string),
	official:     make(map[xlanguage.Tag]string),
	capital:      make(map[xlanguage.Tag]string),
}

func init() { register(dataParaguay) }
