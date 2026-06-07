package country

import (
	xlanguage "golang.org/x/text/language"

	"github.com/lazygophers/utils/currency"
)

// Guyana — Co-operative Republic of Guyana.
var dataGuyana = &Country{
	alpha2:       "GY",
	alpha3:       "GUY",
	numeric:      328,
	callingCodes: []string{"+592"},
	timezones:    []string{"America/Guyana"},
	tlds:         []string{".gy"},
	languages:    []xlanguage.Tag{xlanguage.English},
	currency:     currency.Gyd,
	continent:    "SA",
	region:       "Americas",
	subregion:    "South America",
	flagEmoji:    "🇬🇾",
	names:        make(map[xlanguage.Tag]string),
	official:     make(map[xlanguage.Tag]string),
	capital:      make(map[xlanguage.Tag]string),
}

func init() { register(dataGuyana) }
