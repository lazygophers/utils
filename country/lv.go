//go:build country_all || country_europe || country_lv || country_northern_europe

package country

import (
	xlanguage "golang.org/x/text/language"

	"github.com/lazygophers/utils/currency"
)

// Latvia — Republic of Latvia.
var dataLatvia = &Country{
	alpha2:       "LV",
	alpha3:       "LVA",
	numeric:      428,
	callingCodes: []string{"+371"},
	timezones:    []string{"Europe/Riga"},
	tlds:         []string{".lv"},
	officialLanguage:  xlanguage.Latvian,
	spokenLanguages:   []xlanguage.Tag{xlanguage.Latvian, xlanguage.Russian},
	currency:     currency.EUR,
	region:       RegionNorthernEurope,
	flagEmoji:    "🇱🇻",
	names:        make(map[xlanguage.Tag]string),
	official:     make(map[xlanguage.Tag]string),
	capital:      make(map[xlanguage.Tag]string),
}

func init() { register(dataLatvia) }

var Latvia = dataLatvia
