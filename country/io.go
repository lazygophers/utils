//go:build country_all || country_asia || country_eastern_africa || country_io

package country

import (
	xlanguage "golang.org/x/text/language"

	"github.com/lazygophers/utils/currency"
)

// BritishIndianOceanTerritory — British Indian Ocean Territory.
var dataBritishIndianOceanTerritory = &Country{
	alpha2:       "IO",
	alpha3:       "IOT",
	numeric:      86,
	callingCodes: []string{"+246"},
	timezones:    []string{"Indian/Chagos"},
	tlds:         []string{".io"},
	languages:    []xlanguage.Tag{xlanguage.English},
	currency:     currency.Usd,
	continent:    "AS",
	region:       "Asia",
	subregion:    "Eastern Africa",
	flagEmoji:    "🇮🇴",
	names:        make(map[xlanguage.Tag]string),
	official:     make(map[xlanguage.Tag]string),
	capital:      make(map[xlanguage.Tag]string),
}

func init() { register(dataBritishIndianOceanTerritory) }

var BritishIndianOceanTerritory = dataBritishIndianOceanTerritory
