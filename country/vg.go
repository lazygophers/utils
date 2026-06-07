//go:build country_all || country_americas || country_caribbean || country_vg

package country

import (
	xlanguage "golang.org/x/text/language"

	"github.com/lazygophers/utils/currency"
)

// BritishVirginIslands — British Virgin Islands — British Overseas Territory.
var dataBritishVirginIslands = &Country{
	alpha2:       "VG",
	alpha3:       "VGB",
	numeric:      92,
	callingCodes: []string{"+1-284"},
	timezones:    []string{"America/Tortola"},
	tlds:         []string{".vg"},
	languages:    []xlanguage.Tag{xlanguage.English},
	currency:     currency.Usd,
	continent:    "NA",
	region:       "Americas",
	subregion:    "Caribbean",
	flagEmoji:    "🇻🇬",
	names:        make(map[xlanguage.Tag]string),
	official:     make(map[xlanguage.Tag]string),
	capital:      make(map[xlanguage.Tag]string),
}

func init() { register(dataBritishVirginIslands) }

var BritishVirginIslands = dataBritishVirginIslands
