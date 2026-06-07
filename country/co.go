//go:build country_all || country_americas || country_co || country_south_america

package country

import (
	xlanguage "golang.org/x/text/language"

	"github.com/lazygophers/utils/currency"
)

// Colombia — Republic of Colombia.
var dataColombia = &Country{
	alpha2:       "CO",
	alpha3:       "COL",
	numeric:      170,
	callingCodes: []string{"+57"},
	timezones:    []string{"America/Bogota"},
	tlds:         []string{".co"},
	languages:    []xlanguage.Tag{xlanguage.Spanish},
	currency:     currency.Cop,
	continent:    "SA",
	region:       "Americas",
	subregion:    "South America",
	flagEmoji:    "🇨🇴",
	names:        make(map[xlanguage.Tag]string),
	official:     make(map[xlanguage.Tag]string),
	capital:      make(map[xlanguage.Tag]string),
}

func init() { register(dataColombia) }

var Colombia = dataColombia
