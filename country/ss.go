//go:build country_africa || country_all || country_eastern_africa || country_ss

package country

import (
	xlanguage "golang.org/x/text/language"

	"github.com/lazygophers/utils/currency"
)

// SouthSudan — Republic of South Sudan.
var dataSouthSudan = &Country{
	alpha2:       "SS",
	alpha3:       "SSD",
	numeric:      728,
	callingCodes: []string{"+211"},
	timezones:    []string{"Africa/Juba"},
	tlds:         []string{".ss"},
	languages:    []xlanguage.Tag{xlanguage.English},
	currency:     currency.Ssp,
	continent:    "AF",
	region:       "Africa",
	subregion:    "Eastern Africa",
	flagEmoji:    "🇸🇸",
	names:        make(map[xlanguage.Tag]string),
	official:     make(map[xlanguage.Tag]string),
	capital:      make(map[xlanguage.Tag]string),
}

func init() { register(dataSouthSudan) }

var SouthSudan = dataSouthSudan
