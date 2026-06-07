package country

import (
	xlanguage "golang.org/x/text/language"

	"github.com/lazygophers/utils/currency"
)

// NorthernMarianaIslands — Commonwealth of the Northern Mariana Islands.
var dataNorthernMarianaIslands = &Country{
	alpha2:       "MP",
	alpha3:       "MNP",
	numeric:      580,
	callingCodes: []string{"+1-670"},
	timezones:    []string{"Pacific/Saipan"},
	tlds:         []string{".mp"},
	languages:    []xlanguage.Tag{xlanguage.English},
	currency:     currency.Usd,
	continent:    "OC",
	region:       "Oceania",
	subregion:    "Micronesia",
	flagEmoji:    "🇲🇵",
	names:        make(map[xlanguage.Tag]string),
	official:     make(map[xlanguage.Tag]string),
	capital:      make(map[xlanguage.Tag]string),
}

func init() { register(dataNorthernMarianaIslands) }
