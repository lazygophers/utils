//go:build country_all || country_europe || country_lu || country_western_europe

package country

import (
	xlanguage "golang.org/x/text/language"

	"github.com/lazygophers/utils/currency"
)

// Luxembourg — Grand Duchy of Luxembourg.
var dataLuxembourg = &Country{
	alpha2:       "LU",
	alpha3:       "LUX",
	numeric:      442,
	callingCodes: []string{"+352"},
	timezones:    []string{"Europe/Luxembourg"},
	tlds:         []string{".lu"},
	officialLanguage:  xlanguage.MustParse("lb"),
	spokenLanguages:   []xlanguage.Tag{xlanguage.MustParse("lb"), xlanguage.French, xlanguage.German},
	currency:     currency.EUR,
	region:       RegionWesternEurope,
	flagEmoji:    "🇱🇺",
	names:        make(map[xlanguage.Tag]string),
	official:     make(map[xlanguage.Tag]string),
	capital:      make(map[xlanguage.Tag]string),
}

func init() { register(dataLuxembourg) }

var Luxembourg = dataLuxembourg
