//go:build country_all || country_eastern_europe || country_europe || country_pl

package country

import (
	xlanguage "golang.org/x/text/language"

	"github.com/lazygophers/utils/currency"
)

// Poland — Republic of Poland.
var dataPoland = &Country{
	alpha2:       "PL",
	alpha3:       "POL",
	numeric:      616,
	callingCodes: []string{"+48"},
	timezones:    []string{"Europe/Warsaw"},
	tlds:         []string{".pl"},
	languages:    []xlanguage.Tag{xlanguage.Polish},
	currency:     currency.Pln,
	continent:    "EU",
	region:       "Europe",
	subregion:    "Eastern Europe",
	flagEmoji:    "🇵🇱",
	names:        make(map[xlanguage.Tag]string),
	official:     make(map[xlanguage.Tag]string),
	capital:      make(map[xlanguage.Tag]string),
}

func init() { register(dataPoland) }

var Poland = dataPoland
