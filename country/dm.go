//go:build country_all || country_americas || country_caribbean || country_dm

package country

import (
	xlanguage "golang.org/x/text/language"

	"github.com/lazygophers/utils/currency"
)

// Dominica — Commonwealth of Dominica.
var dataDominica = &Country{
	alpha2:       "DM",
	alpha3:       "DMA",
	numeric:      212,
	callingCodes: []string{"+1-767"},
	timezones:    []string{"America/Dominica"},
	tlds:         []string{".dm"},
	languages:    []xlanguage.Tag{xlanguage.English},
	currency:     currency.Xcd,
	continent:    "NA",
	region:       "Americas",
	subregion:    "Caribbean",
	flagEmoji:    "🇩🇲",
	names:        make(map[xlanguage.Tag]string),
	official:     make(map[xlanguage.Tag]string),
	capital:      make(map[xlanguage.Tag]string),
}

func init() { register(dataDominica) }

var Dominica = dataDominica
