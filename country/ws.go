package country

import (
	xlanguage "golang.org/x/text/language"

	"github.com/lazygophers/utils/currency"
)

// Samoa — Independent State of Samoa.
var dataSamoa = &Country{
	alpha2:       "WS",
	alpha3:       "WSM",
	numeric:      882,
	callingCodes: []string{"+685"},
	timezones:    []string{"Pacific/Apia"},
	tlds:         []string{".ws"},
	languages:    []xlanguage.Tag{xlanguage.MustParse("sm"), xlanguage.English},
	currency:     currency.Wst,
	continent:    "OC",
	region:       "Oceania",
	subregion:    "Polynesia",
	flagEmoji:    "🇼🇸",
	names:        make(map[xlanguage.Tag]string),
	official:     make(map[xlanguage.Tag]string),
	capital:      make(map[xlanguage.Tag]string),
}

func init() { register(dataSamoa) }
