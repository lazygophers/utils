//go:build country_all || country_americas || country_bo || country_south_america

package country

import (
	xlanguage "golang.org/x/text/language"

	"github.com/lazygophers/utils/currency"
)

// Bolivia — Plurinational State of Bolivia.
var dataBolivia = &Country{
	alpha2:       "BO",
	alpha3:       "BOL",
	numeric:      68,
	callingCodes: []string{"+591"},
	timezones:    []string{"America/La_Paz"},
	tlds:         []string{".bo"},
	languages:    []xlanguage.Tag{xlanguage.Spanish},
	currency:     currency.Bob,
	continent:    "SA",
	region:       "Americas",
	subregion:    "South America",
	flagEmoji:    "🇧🇴",
	names:        make(map[xlanguage.Tag]string),
	official:     make(map[xlanguage.Tag]string),
	capital:      make(map[xlanguage.Tag]string),
}

func init() { register(dataBolivia) }

var Bolivia = dataBolivia
