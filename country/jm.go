package country

import (
	xlanguage "golang.org/x/text/language"

	"github.com/lazygophers/utils/currency"
)

// Jamaica — Jamaica.
var dataJamaica = &Country{
	alpha2:       "JM",
	alpha3:       "JAM",
	numeric:      388,
	callingCodes: []string{"+1-876"},
	timezones:    []string{"America/Jamaica"},
	tlds:         []string{".jm"},
	languages:    []xlanguage.Tag{xlanguage.English},
	currency:     currency.Jmd,
	continent:    "NA",
	region:       "Americas",
	subregion:    "Caribbean",
	flagEmoji:    "🇯🇲",
	names:        make(map[xlanguage.Tag]string),
	official:     make(map[xlanguage.Tag]string),
	capital:      make(map[xlanguage.Tag]string),
}

func init() { register(dataJamaica) }
