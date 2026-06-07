//go:build country_all || country_europe || country_ie || country_northern_europe

package country

import (
	xlanguage "golang.org/x/text/language"

	"github.com/lazygophers/utils/currency"
)

// Ireland — Republic of Ireland.
var dataIreland = &Country{
	alpha2:       "IE",
	alpha3:       "IRL",
	numeric:      372,
	callingCodes: []string{"+353"},
	timezones:    []string{"Europe/Dublin"},
	tlds:         []string{".ie"},
	officialLanguage:  xlanguage.English,
	spokenLanguages:   []xlanguage.Tag{xlanguage.English, xlanguage.MustParse("ga")},
	currency:     currency.EUR,
	region:       RegionNorthernEurope,
	flagEmoji:    "🇮🇪",
	names:        make(map[xlanguage.Tag]string),
	official:     make(map[xlanguage.Tag]string),
	capital:      make(map[xlanguage.Tag]string),
}

func init() { register(dataIreland) }

var Ireland = dataIreland
