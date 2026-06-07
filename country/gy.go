//go:build country_all || country_americas || country_gy || country_south_america

package country

import (
	xlanguage "golang.org/x/text/language"

	"github.com/lazygophers/utils/currency"
)

// Guyana — Co-operative Republic of Guyana.
var dataGuyana = &Country{
	alpha2:       "GY",
	alpha3:       "GUY",
	numeric:      328,
	callingCodes: []string{"+592"},
	timezones:    []string{"America/Guyana"},
	tlds:         []string{".gy"},
	officialLanguage:  xlanguage.English,
	spokenLanguages:   []xlanguage.Tag{xlanguage.English},
	currency:     currency.GYD,
	region:       RegionSouthAmerica,
	flagEmoji:    "🇬🇾",
	names:        make(map[xlanguage.Tag]string),
	official:     make(map[xlanguage.Tag]string),
	capital:      make(map[xlanguage.Tag]string),
}

func init() { register(dataGuyana) }

var Guyana = dataGuyana
