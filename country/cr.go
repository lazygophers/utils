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
	languages:    []xlanguage.Tag{xlanguage.Spanish},
	currency:     currency.Crc,
	continent:    "NA",
	region:       "Americas",
	subregion:    "Central America",
	flagEmoji:    "🇨🇷",
	names:        make(map[xlanguage.Tag]string),
	official:     make(map[xlanguage.Tag]string),
	capital:      make(map[xlanguage.Tag]string),
}

func init() { register(dataCostaRica) }
