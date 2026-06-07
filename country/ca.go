//go:build country_all || country_americas || country_ca || country_northern_america

package country

import (
	xlanguage "golang.org/x/text/language"

	"github.com/lazygophers/utils/currency"
)

// Canada — Canada.
var dataCanada = &Country{
	alpha2:       "CA",
	alpha3:       "CAN",
	numeric:      124,
	callingCodes: []string{"+1"},
	timezones:    []string{
		"America/Toronto",
		"America/Montreal",
		"America/Vancouver",
		"America/Edmonton",
		"America/Winnipeg",
		"America/Halifax",
		"America/St_Johns",
	},
	tlds:         []string{".ca"},
	languages:    []xlanguage.Tag{xlanguage.English, xlanguage.French},
	currency:     currency.Cad,
	continent:    "NA",
	region:       "Americas",
	subregion:    "Northern America",
	flagEmoji:    "🇨🇦",
	names:        make(map[xlanguage.Tag]string),
	official:     make(map[xlanguage.Tag]string),
	capital:      make(map[xlanguage.Tag]string),
}

func init() { register(dataCanada) }

var Canada = dataCanada
