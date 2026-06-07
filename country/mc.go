//go:build country_all || country_europe || country_mc || country_western_europe

package country

import (
	xlanguage "golang.org/x/text/language"

	"github.com/lazygophers/utils/currency"
)

// Monaco — Principality of Monaco.
var dataMonaco = &Country{
	alpha2:       "MC",
	alpha3:       "MCO",
	numeric:      492,
	callingCodes: []string{"+377"},
	timezones:    []string{"Europe/Monaco"},
	tlds:         []string{".mc"},
	officialLanguage:  xlanguage.French,
	spokenLanguages:   []xlanguage.Tag{xlanguage.French},
	currency:     currency.EUR,
	region:       RegionWesternEurope,
	flagEmoji:    "🇲🇨",
	names:        make(map[xlanguage.Tag]string),
	official:     make(map[xlanguage.Tag]string),
	capital:      make(map[xlanguage.Tag]string),
}

func init() { register(dataMonaco) }

var Monaco = dataMonaco
