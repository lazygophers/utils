//go:build country_all || country_europe || country_rs || country_southern_europe

package country

import (
	xlanguage "golang.org/x/text/language"

	"github.com/lazygophers/utils/currency"
)

// Serbia — Republic of Serbia.
var dataSerbia = &Country{
	alpha2:       "RS",
	alpha3:       "SRB",
	numeric:      688,
	callingCodes: []string{"+381"},
	timezones:    []string{"Europe/Belgrade"},
	tlds:         []string{
		".rs",
		".срб",
	},
	officialLanguage:  xlanguage.Serbian,
	spokenLanguages:   []xlanguage.Tag{xlanguage.Serbian},
	currency:     currency.RSD,
	region:       RegionSouthernEurope,
	flagEmoji:    "🇷🇸",
	names:        make(map[xlanguage.Tag]string),
	official:     make(map[xlanguage.Tag]string),
	capital:      make(map[xlanguage.Tag]string),
}

func init() { register(dataSerbia) }

var Serbia = dataSerbia
