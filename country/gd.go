//go:build country_all || country_americas || country_caribbean || country_gd

package country

import (
	xlanguage "golang.org/x/text/language"

	"github.com/lazygophers/utils/currency"
)

// Grenada — Grenada.
var dataGrenada = &Country{
	alpha2:       "GD",
	alpha3:       "GRD",
	numeric:      308,
	callingCodes: []string{"+1-473"},
	timezones:    []string{"America/Grenada"},
	tlds:         []string{".gd"},
	officialLanguage:  xlanguage.English,
	spokenLanguages:   []xlanguage.Tag{xlanguage.English},
	currency:     currency.XCD,
	region:       RegionCaribbean,
	flagEmoji:    "🇬🇩",
	names:        make(map[xlanguage.Tag]string),
	official:     make(map[xlanguage.Tag]string),
	capital:      make(map[xlanguage.Tag]string),
}

func init() { register(dataGrenada) }

var Grenada = dataGrenada
