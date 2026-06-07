package country

import (
	xlanguage "golang.org/x/text/language"

	"github.com/lazygophers/utils/currency"
)

// Peru — Republic of Peru.
var dataPeru = &Country{
	alpha2:       "PE",
	alpha3:       "PER",
	numeric:      604,
	callingCodes: []string{"+51"},
	timezones:    []string{"America/Lima"},
	tlds:         []string{".pe"},
	languages:    []xlanguage.Tag{xlanguage.Spanish},
	currency:     currency.Pen,
	continent:    "SA",
	region:       "Americas",
	subregion:    "South America",
	flagEmoji:    "🇵🇪",
	names:        make(map[xlanguage.Tag]string),
	official:     make(map[xlanguage.Tag]string),
	capital:      make(map[xlanguage.Tag]string),
}

func init() { register(dataPeru) }
