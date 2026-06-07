//go:build country_africa || country_all || country_ly || country_northern_africa

package country

import (
	xlanguage "golang.org/x/text/language"

	"github.com/lazygophers/utils/currency"
)

// Libya — State of Libya.
var dataLibya = &Country{
	alpha2:       "LY",
	alpha3:       "LBY",
	numeric:      434,
	callingCodes: []string{"+218"},
	timezones:    []string{"Africa/Tripoli"},
	tlds:         []string{".ly"},
	officialLanguage:  xlanguage.Arabic,
	spokenLanguages:   []xlanguage.Tag{xlanguage.Arabic},
	currency:     currency.LYD,
	region:       RegionNorthernAfrica,
	flagEmoji:    "🇱🇾",
	names:        make(map[xlanguage.Tag]string),
	official:     make(map[xlanguage.Tag]string),
	capital:      make(map[xlanguage.Tag]string),
}

func init() { register(dataLibya) }

var Libya = dataLibya
