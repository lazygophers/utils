//go:build country_all || country_antarctic || country_gs

package country

import (
	xlanguage "golang.org/x/text/language"

	"github.com/lazygophers/utils/currency"
)

// SouthGeorgiaAndSouthSandwich — South Georgia and the South Sandwich Islands — British Overseas Territory.
var dataSouthGeorgiaAndSouthSandwich = &Country{
	alpha2:       "GS",
	alpha3:       "SGS",
	numeric:      239,
	callingCodes: []string{"+500"},
	timezones:    []string{"Atlantic/South_Georgia"},
	tlds:         []string{".gs"},
	officialLanguage:  xlanguage.English,
	spokenLanguages:   []xlanguage.Tag{xlanguage.English},
	currency:     currency.GBP,
	region:       RegionAntarctic,
	flagEmoji:    "🇬🇸",
	names:        make(map[xlanguage.Tag]string),
	official:     make(map[xlanguage.Tag]string),
	capital:      make(map[xlanguage.Tag]string),
}

func init() { register(dataSouthGeorgiaAndSouthSandwich) }

var SouthGeorgiaAndSouthSandwich = dataSouthGeorgiaAndSouthSandwich
