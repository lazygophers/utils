//go:build country_all || country_europe || country_fi || country_northern_europe

package country

import (
	xlanguage "golang.org/x/text/language"

	"github.com/lazygophers/utils/currency"
)

// Finland — Republic of Finland.
var dataFinland = &Country{
	alpha2:       "FI",
	alpha3:       "FIN",
	numeric:      246,
	callingCodes: []string{"+358"},
	timezones:    []string{"Europe/Helsinki"},
	tlds:         []string{".fi"},
	officialLanguage:  xlanguage.Finnish,
	spokenLanguages:   []xlanguage.Tag{xlanguage.Finnish, xlanguage.Swedish},
	currency:     currency.EUR,
	region:       RegionNorthernEurope,
	flagEmoji:    "🇫🇮",
	names:        make(map[xlanguage.Tag]string),
	official:     make(map[xlanguage.Tag]string),
	capital:      make(map[xlanguage.Tag]string),
}

func init() { register(dataFinland) }

var Finland = dataFinland
