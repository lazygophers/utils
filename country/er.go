//go:build country_africa || country_all || country_eastern_africa || country_er

package country

import (
	xlanguage "golang.org/x/text/language"

	"github.com/lazygophers/utils/currency"
)

// Eritrea — State of Eritrea.
var dataEritrea = &Country{
	alpha2:       "ER",
	alpha3:       "ERI",
	numeric:      232,
	callingCodes: []string{"+291"},
	timezones:    []string{"Africa/Asmara"},
	tlds:         []string{".er"},
	languages:    []xlanguage.Tag{xlanguage.MustParse("ti"), xlanguage.Arabic, xlanguage.English},
	currency:     currency.Ern,
	continent:    "AF",
	region:       "Africa",
	subregion:    "Eastern Africa",
	flagEmoji:    "🇪🇷",
	names:        make(map[xlanguage.Tag]string),
	official:     make(map[xlanguage.Tag]string),
	capital:      make(map[xlanguage.Tag]string),
}

func init() { register(dataEritrea) }

var Eritrea = dataEritrea
