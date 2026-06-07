//go:build country_all || country_europe || country_fo || country_northern_europe

package country

import (
	xlanguage "golang.org/x/text/language"

	"github.com/lazygophers/utils/currency"
)

// FaroeIslands — Faroe Islands — autonomous region of Denmark.
var dataFaroeIslands = &Country{
	alpha2:       "FO",
	alpha3:       "FRO",
	numeric:      234,
	callingCodes: []string{"+298"},
	timezones:    []string{"Atlantic/Faroe"},
	tlds:         []string{".fo"},
	officialLanguage:  xlanguage.MustParse("fo"),
	spokenLanguages:   []xlanguage.Tag{xlanguage.MustParse("fo"), xlanguage.Danish},
	currency:     currency.DKK,
	region:       RegionNorthernEurope,
	flagEmoji:    "🇫🇴",
	names:        make(map[xlanguage.Tag]string),
	official:     make(map[xlanguage.Tag]string),
	capital:      make(map[xlanguage.Tag]string),
}

func init() { register(dataFaroeIslands) }

var FaroeIslands = dataFaroeIslands
