//go:build country_all || country_europe || country_nl || country_western_europe

package country

import (
	xlanguage "golang.org/x/text/language"

	"github.com/lazygophers/utils/currency"
)

// Netherlands — Kingdom of the Netherlands.
var dataNetherlands = &Country{
	alpha2:       "NL",
	alpha3:       "NLD",
	numeric:      528,
	callingCodes: []string{"+31"},
	timezones:    []string{"Europe/Amsterdam"},
	tlds:         []string{".nl"},
	officialLanguage:  xlanguage.Dutch,
	spokenLanguages:   []xlanguage.Tag{xlanguage.Dutch},
	currency:     currency.EUR,
	region:       RegionWesternEurope,
	flagEmoji:    "🇳🇱",
	names:        make(map[xlanguage.Tag]string),
	official:     make(map[xlanguage.Tag]string),
	capital:      make(map[xlanguage.Tag]string),
}

func init() { register(dataNetherlands) }

var Netherlands = dataNetherlands
