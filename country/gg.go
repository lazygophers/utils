//go:build country_all || country_europe || country_gg || country_northern_europe

package country

import (
	xlanguage "golang.org/x/text/language"

	"github.com/lazygophers/utils/currency"
)

// Guernsey — Bailiwick of Guernsey — British Crown Dependency.
var dataGuernsey = &Country{
	alpha2:       "GG",
	alpha3:       "GGY",
	numeric:      831,
	callingCodes: []string{"+44-1481"},
	timezones:    []string{"Europe/Guernsey"},
	tlds:         []string{".gg"},
	officialLanguage:  xlanguage.English,
	spokenLanguages:   []xlanguage.Tag{xlanguage.English},
	currency:     currency.GBP,
	region:       RegionNorthernEurope,
	flagEmoji:    "🇬🇬",
	names:        make(map[xlanguage.Tag]string),
	official:     make(map[xlanguage.Tag]string),
	capital:      make(map[xlanguage.Tag]string),
}

func init() { register(dataGuernsey) }

var Guernsey = dataGuernsey
