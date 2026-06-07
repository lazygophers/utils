//go:build country_all || country_americas || country_ar || country_south_america

package country

import (
	xlanguage "golang.org/x/text/language"

	"github.com/lazygophers/utils/currency"
)

// Argentina — Argentine Republic.
var dataArgentina = &Country{
	alpha2:       "AR",
	alpha3:       "ARG",
	numeric:      32,
	callingCodes: []string{"+54"},
	timezones:    []string{"America/Argentina/Buenos_Aires"},
	tlds:         []string{".ar"},
	languages:    []xlanguage.Tag{xlanguage.Spanish},
	currency:     currency.Ars,
	continent:    "SA",
	region:       "Americas",
	subregion:    "South America",
	flagEmoji:    "🇦🇷",
	names:        make(map[xlanguage.Tag]string),
	official:     make(map[xlanguage.Tag]string),
	capital:      make(map[xlanguage.Tag]string),
}

func init() { register(dataArgentina) }

var Argentina = dataArgentina
