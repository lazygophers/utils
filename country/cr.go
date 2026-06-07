//go:build country_all || country_americas || country_central_america || country_cr

package country

import (
	xlanguage "golang.org/x/text/language"

	"github.com/lazygophers/utils/currency"
)

// CostaRica — Republic of Costa Rica.
var dataCostaRica = &Country{
	alpha2:       "CR",
	alpha3:       "CRI",
	numeric:      188,
	callingCodes: []string{"+506"},
	timezones:    []string{"America/Costa_Rica"},
	tlds:         []string{".cr"},
	officialLanguage:  xlanguage.Spanish,
	spokenLanguages:   []xlanguage.Tag{xlanguage.Spanish},
	currency:     currency.CRC,
	region:       RegionCentralAmerica,
	flagEmoji:    "🇨🇷",
	names:        make(map[xlanguage.Tag]string),
	official:     make(map[xlanguage.Tag]string),
	capital:      make(map[xlanguage.Tag]string),
}

func init() { register(dataCostaRica) }

var CostaRica = dataCostaRica
