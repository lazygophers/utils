//go:build country_all || country_americas || country_caribbean || country_jm

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
	officialLanguage:  xlanguage.English,
	spokenLanguages:   []xlanguage.Tag{xlanguage.English},
	currency:     currency.JMD,
	region:       RegionCaribbean,
	flagEmoji:    "🇯🇲",
	names:        make(map[xlanguage.Tag]string),
	official:     make(map[xlanguage.Tag]string),
	capital:      make(map[xlanguage.Tag]string),
}

func init() { register(dataJamaica) }

var Jamaica = dataJamaica
