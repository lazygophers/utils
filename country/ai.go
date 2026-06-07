//go:build country_ai || country_all || country_americas || country_caribbean

package country

import (
	xlanguage "golang.org/x/text/language"

	"github.com/lazygophers/utils/currency"
)

// Anguilla — British Overseas Territory of Anguilla.
var dataAnguilla = &Country{
	alpha2:       "AI",
	alpha3:       "AIA",
	numeric:      660,
	callingCodes: []string{"+1-264"},
	timezones:    []string{"America/Anguilla"},
	tlds:         []string{".ai"},
	languages:    []xlanguage.Tag{xlanguage.English},
	currency:     currency.Xcd,
	continent:    "NA",
	region:       "Americas",
	subregion:    "Caribbean",
	flagEmoji:    "🇦🇮",
	names:        make(map[xlanguage.Tag]string),
	official:     make(map[xlanguage.Tag]string),
	capital:      make(map[xlanguage.Tag]string),
}

func init() { register(dataAnguilla) }

var Anguilla = dataAnguilla
