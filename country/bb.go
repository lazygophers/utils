//go:build country_all || country_americas || country_bb || country_caribbean

package country

import (
	xlanguage "golang.org/x/text/language"

	"github.com/lazygophers/utils/currency"
)

// Barbados — Barbados.
var dataBarbados = &Country{
	alpha2:       "BB",
	alpha3:       "BRB",
	numeric:      52,
	callingCodes: []string{"+1-246"},
	timezones:    []string{"America/Barbados"},
	tlds:         []string{".bb"},
	officialLanguage:  xlanguage.English,
	spokenLanguages:   []xlanguage.Tag{xlanguage.English},
	currency:     currency.BBD,
	region:       RegionCaribbean,
	flagEmoji:    "🇧🇧",
	names:        make(map[xlanguage.Tag]string),
	official:     make(map[xlanguage.Tag]string),
	capital:      make(map[xlanguage.Tag]string),
}

func init() { register(dataBarbados) }

var Barbados = dataBarbados
