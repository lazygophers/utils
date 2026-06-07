//go:build country_africa || country_all || country_eastern_africa || country_tz

package country

import (
	xlanguage "golang.org/x/text/language"

	"github.com/lazygophers/utils/currency"
)

// Tanzania — United Republic of Tanzania.
var dataTanzania = &Country{
	alpha2:       "TZ",
	alpha3:       "TZA",
	numeric:      834,
	callingCodes: []string{"+255"},
	timezones:    []string{"Africa/Dar_es_Salaam"},
	tlds:         []string{".tz"},
	languages:    []xlanguage.Tag{xlanguage.MustParse("sw"), xlanguage.English},
	currency:     currency.Tzs,
	continent:    "AF",
	region:       "Africa",
	subregion:    "Eastern Africa",
	flagEmoji:    "🇹🇿",
	names:        make(map[xlanguage.Tag]string),
	official:     make(map[xlanguage.Tag]string),
	capital:      make(map[xlanguage.Tag]string),
}

func init() { register(dataTanzania) }

var Tanzania = dataTanzania
