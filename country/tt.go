//go:build country_all || country_americas || country_caribbean || country_tt

package country

import (
	xlanguage "golang.org/x/text/language"

	"github.com/lazygophers/utils/currency"
)

// TrinidadAndTobago — Republic of Trinidad and Tobago.
var dataTrinidadAndTobago = &Country{
	alpha2:       "TT",
	alpha3:       "TTO",
	numeric:      780,
	callingCodes: []string{"+1-868"},
	timezones:    []string{"America/Port_of_Spain"},
	tlds:         []string{".tt"},
	languages:    []xlanguage.Tag{xlanguage.English},
	currency:     currency.Ttd,
	continent:    "NA",
	region:       "Americas",
	subregion:    "Caribbean",
	flagEmoji:    "🇹🇹",
	names:        make(map[xlanguage.Tag]string),
	official:     make(map[xlanguage.Tag]string),
	capital:      make(map[xlanguage.Tag]string),
}

func init() { register(dataTrinidadAndTobago) }

var TrinidadAndTobago = dataTrinidadAndTobago
