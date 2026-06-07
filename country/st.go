//go:build country_africa || country_all || country_middle_africa || country_st

package country

import (
	xlanguage "golang.org/x/text/language"

	"github.com/lazygophers/utils/currency"
)

// SaoTomeAndPrincipe — Democratic Republic of São Tomé and Príncipe.
var dataSaoTomeAndPrincipe = &Country{
	alpha2:       "ST",
	alpha3:       "STP",
	numeric:      678,
	callingCodes: []string{"+239"},
	timezones:    []string{"Africa/Sao_Tome"},
	tlds:         []string{".st"},
	languages:    []xlanguage.Tag{xlanguage.Portuguese},
	currency:     currency.Stn,
	continent:    "AF",
	region:       "Africa",
	subregion:    "Middle Africa",
	flagEmoji:    "🇸🇹",
	names:        make(map[xlanguage.Tag]string),
	official:     make(map[xlanguage.Tag]string),
	capital:      make(map[xlanguage.Tag]string),
}

func init() { register(dataSaoTomeAndPrincipe) }

var SaoTomeAndPrincipe = dataSaoTomeAndPrincipe
