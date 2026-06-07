//go:build country_all || country_europe || country_im || country_northern_europe

package country

import (
	xlanguage "golang.org/x/text/language"

	"github.com/lazygophers/utils/currency"
)

// IsleOfMan — Isle of Man — British Crown Dependency.
var dataIsleOfMan = &Country{
	alpha2:       "IM",
	alpha3:       "IMN",
	numeric:      833,
	callingCodes: []string{"+44-1624"},
	timezones:    []string{"Europe/Isle_of_Man"},
	tlds:         []string{".im"},
	officialLanguage:  xlanguage.English,
	spokenLanguages:   []xlanguage.Tag{xlanguage.English},
	currency:     currency.GBP,
	region:       RegionNorthernEurope,
	flagEmoji:    "🇮🇲",
	names:        make(map[xlanguage.Tag]string),
	official:     make(map[xlanguage.Tag]string),
	capital:      make(map[xlanguage.Tag]string),
}

func init() { register(dataIsleOfMan) }

var IsleOfMan = dataIsleOfMan
