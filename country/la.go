package country

import (
	xlanguage "golang.org/x/text/language"

	"github.com/lazygophers/utils/currency"
)

// Laos — Lao People's Democratic Republic.
var dataLaos = &Country{
	alpha2:       "LA",
	alpha3:       "LAO",
	numeric:      418,
	callingCodes: []string{"+856"},
	timezones:    []string{"Asia/Vientiane"},
	tlds:         []string{".la"},
	languages:    []xlanguage.Tag{xlanguage.MustParse("lo")},
	currency:     currency.Lak,
	continent:    "AS",
	region:       "Asia",
	subregion:    "South-eastern Asia",
	flagEmoji:    "🇱🇦",
	names:        make(map[xlanguage.Tag]string),
	official:     make(map[xlanguage.Tag]string),
	capital:      make(map[xlanguage.Tag]string),
}

func init() { register(dataLaos) }
