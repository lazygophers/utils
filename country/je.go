//go:build country_all || country_europe || country_je || country_northern_europe

package country

import (
	xlanguage "golang.org/x/text/language"

	"github.com/lazygophers/utils/currency"
)

// Jersey — Bailiwick of Jersey — British Crown Dependency.
var dataJersey = &Country{
	alpha2:       "JE",
	alpha3:       "JEY",
	numeric:      832,
	callingCodes: []string{"+44-1534"},
	timezones:    []string{"Europe/Jersey"},
	tlds:         []string{".je"},
	officialLanguage:  xlanguage.English,
	spokenLanguages:   []xlanguage.Tag{xlanguage.English},
	currency:     currency.GBP,
	region:       RegionNorthernEurope,
	flagEmoji:    "🇯🇪",
	names:        make(map[xlanguage.Tag]string),
	official:     make(map[xlanguage.Tag]string),
	capital:      make(map[xlanguage.Tag]string),
}

func init() { register(dataJersey) }

var Jersey = dataJersey
