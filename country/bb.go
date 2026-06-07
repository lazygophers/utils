//go:build country_all || country_americas || country_bb || country_caribbean

package country

import (
	xlanguage "golang.org/x/text/language"

	"github.com/lazygophers/utils/currency"
)

// Barbados — Barbados.
var dataBarbados = &Country{
	alpha2:       "BB",
	alpha3:       "BRB",
	numeric:      52,
	callingCodes: []string{"+1-246"},
	timezones:    []string{"America/Barbados"},
	tlds:         []string{".bb"},
	languages:    []xlanguage.Tag{xlanguage.English},
	currency:     currency.Bbd,
	continent:    "NA",
	region:       "Americas",
	subregion:    "Caribbean",
	flagEmoji:    "🇧🇧",
	names:        make(map[xlanguage.Tag]string),
	official:     make(map[xlanguage.Tag]string),
	capital:      make(map[xlanguage.Tag]string),
}

func init() { register(dataBarbados) }

var Barbados = dataBarbados
