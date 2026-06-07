package country

import (
	xlanguage "golang.org/x/text/language"

	"github.com/lazygophers/utils/currency"
)

// Venezuela — Bolivarian Republic of Venezuela.
var dataVenezuela = &Country{
	alpha2:       "VE",
	alpha3:       "VEN",
	numeric:      862,
	callingCodes: []string{"+58"},
	timezones:    []string{"America/Caracas"},
	tlds:         []string{".ve"},
	languages:    []xlanguage.Tag{xlanguage.Spanish},
	currency:     currency.Ves,
	continent:    "SA",
	region:       "Americas",
	subregion:    "South America",
	flagEmoji:    "🇻🇪",
	names:        make(map[xlanguage.Tag]string),
	official:     make(map[xlanguage.Tag]string),
	capital:      make(map[xlanguage.Tag]string),
}

func init() { register(dataVenezuela) }
