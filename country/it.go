//go:build country_all || country_europe || country_it || country_southern_europe

package country

import (
	xlanguage "golang.org/x/text/language"

	"github.com/lazygophers/utils/currency"
)

// Italy — Italian Republic.
var dataItaly = &Country{
	alpha2:       "IT",
	alpha3:       "ITA",
	numeric:      380,
	callingCodes: []string{"+39"},
	timezones:    []string{"Europe/Rome"},
	tlds:         []string{".it"},
	officialLanguage:  xlanguage.Italian,
	spokenLanguages:   []xlanguage.Tag{xlanguage.Italian},
	currency:     currency.EUR,
	region:       RegionSouthernEurope,
	flagEmoji:    "🇮🇹",
	names:        make(map[xlanguage.Tag]string),
	official:     make(map[xlanguage.Tag]string),
	capital:      make(map[xlanguage.Tag]string),
}

func init() { register(dataItaly) }

var Italy = dataItaly
