//go:build country_all || country_americas || country_caribbean || country_lc

package country

import (
	xlanguage "golang.org/x/text/language"

	"github.com/lazygophers/utils/currency"
)

// SaintLucia — Saint Lucia.
var dataSaintLucia = &Country{
	alpha2:       "LC",
	alpha3:       "LCA",
	numeric:      662,
	callingCodes: []string{"+1-758"},
	timezones:    []string{"America/St_Lucia"},
	tlds:         []string{".lc"},
	officialLanguage:  xlanguage.English,
	spokenLanguages:   []xlanguage.Tag{xlanguage.English},
	currency:     currency.XCD,
	region:       RegionCaribbean,
	flagEmoji:    "🇱🇨",
	names:        make(map[xlanguage.Tag]string),
	official:     make(map[xlanguage.Tag]string),
	capital:      make(map[xlanguage.Tag]string),
}

func init() { register(dataSaintLucia) }

var SaintLucia = dataSaintLucia
