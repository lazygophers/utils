package country

import (
	xlanguage "golang.org/x/text/language"

	"github.com/lazygophers/utils/currency"
)

// Tuvalu — Tuvalu.
var dataTuvalu = &Country{
	alpha2:       "TV",
	alpha3:       "TUV",
	numeric:      798,
	callingCodes: []string{"+688"},
	timezones:    []string{"Pacific/Funafuti"},
	tlds:         []string{".tv"},
	languages:    []xlanguage.Tag{xlanguage.English, xlanguage.MustParse("tvl")},
	currency:     currency.Aud,
	continent:    "OC",
	region:       "Oceania",
	subregion:    "Polynesia",
	flagEmoji:    "🇹🇻",
	names:        make(map[xlanguage.Tag]string),
	official:     make(map[xlanguage.Tag]string),
	capital:      make(map[xlanguage.Tag]string),
}

func init() { register(dataTuvalu) }
