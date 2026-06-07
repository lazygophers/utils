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
	officialLanguage:  xlanguage.English,
	spokenLanguages:   []xlanguage.Tag{xlanguage.English},
	currency:     currency.TTD,
	region:       RegionCaribbean,
	flagEmoji:    "🇹🇹",
	names:        make(map[xlanguage.Tag]string),
	official:     make(map[xlanguage.Tag]string),
	capital:      make(map[xlanguage.Tag]string),
}

func init() { register(dataTrinidadAndTobago) }

var TrinidadAndTobago = dataTrinidadAndTobago
