package country

import (
	xlanguage "golang.org/x/text/language"

	"github.com/lazygophers/utils/currency"
)

// TurksAndCaicosIslands — Turks and Caicos Islands — British Overseas Territory.
var dataTurksAndCaicosIslands = &Country{
	alpha2:       "TC",
	alpha3:       "TCA",
	numeric:      796,
	callingCodes: []string{"+1-649"},
	timezones:    []string{"America/Grand_Turk"},
	tlds:         []string{".tc"},
	languages:    []xlanguage.Tag{xlanguage.English},
	currency:     currency.Usd,
	continent:    "NA",
	region:       "Americas",
	subregion:    "Caribbean",
	flagEmoji:    "🇹🇨",
	names:        make(map[xlanguage.Tag]string),
	official:     make(map[xlanguage.Tag]string),
	capital:      make(map[xlanguage.Tag]string),
}

func init() { register(dataTurksAndCaicosIslands) }
