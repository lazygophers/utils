//go:build country_all || country_micronesia || country_nr || country_oceania

package country

import (
	xlanguage "golang.org/x/text/language"

	"github.com/lazygophers/utils/currency"
)

// Nauru — Republic of Nauru.
var dataNauru = &Country{
	alpha2:       "NR",
	alpha3:       "NRU",
	numeric:      520,
	callingCodes: []string{"+674"},
	timezones:    []string{"Pacific/Nauru"},
	tlds:         []string{".nr"},
	languages:    []xlanguage.Tag{xlanguage.MustParse("na"), xlanguage.English},
	currency:     currency.Aud,
	continent:    "OC",
	region:       "Oceania",
	subregion:    "Micronesia",
	flagEmoji:    "🇳🇷",
	names:        make(map[xlanguage.Tag]string),
	official:     make(map[xlanguage.Tag]string),
	capital:      make(map[xlanguage.Tag]string),
}

func init() { register(dataNauru) }

var Nauru = dataNauru
