//go:build country_africa || country_all || country_ma || country_northern_africa

package country

import (
	xlanguage "golang.org/x/text/language"

	"github.com/lazygophers/utils/currency"
)

// Morocco — Kingdom of Morocco.
var dataMorocco = &Country{
	alpha2:       "MA",
	alpha3:       "MAR",
	numeric:      504,
	callingCodes: []string{"+212"},
	timezones:    []string{"Africa/Casablanca"},
	tlds:         []string{
		".ma",
		".المغرب",
	},
	languages:    []xlanguage.Tag{xlanguage.Arabic},
	currency:     currency.Mad,
	continent:    "AF",
	region:       "Africa",
	subregion:    "Northern Africa",
	flagEmoji:    "🇲🇦",
	names:        make(map[xlanguage.Tag]string),
	official:     make(map[xlanguage.Tag]string),
	capital:      make(map[xlanguage.Tag]string),
}

func init() { register(dataMorocco) }

var Morocco = dataMorocco
