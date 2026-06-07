package country

import (
	xlanguage "golang.org/x/text/language"

	"github.com/lazygophers/utils/currency"
)

// UsVirginIslands — United States Virgin Islands.
var dataUsVirginIslands = &Country{
	alpha2:       "VI",
	alpha3:       "VIR",
	numeric:      850,
	callingCodes: []string{"+1-340"},
	timezones:    []string{"America/St_Thomas"},
	tlds:         []string{".vi"},
	languages:    []xlanguage.Tag{xlanguage.English},
	currency:     currency.Usd,
	continent:    "NA",
	region:       "Americas",
	subregion:    "Caribbean",
	flagEmoji:    "🇻🇮",
	names:        make(map[xlanguage.Tag]string),
	official:     make(map[xlanguage.Tag]string),
	capital:      make(map[xlanguage.Tag]string),
}

func init() { register(dataUsVirginIslands) }
