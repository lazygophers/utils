//go:build country_africa || country_all || country_cd || country_middle_africa

package country

import (
	xlanguage "golang.org/x/text/language"

	"github.com/lazygophers/utils/currency"
)

// DrCongo — Democratic Republic of the Congo.
var dataDrCongo = &Country{
	alpha2:       "CD",
	alpha3:       "COD",
	numeric:      180,
	callingCodes: []string{"+243"},
	timezones:    []string{
		"Africa/Kinshasa",
		"Africa/Lubumbashi",
	},
	tlds:         []string{".cd"},
	officialLanguage:  xlanguage.French,
	spokenLanguages:   []xlanguage.Tag{xlanguage.French},
	currency:     currency.CDF,
	region:       RegionMiddleAfrica,
	flagEmoji:    "🇨🇩",
	names:        make(map[xlanguage.Tag]string),
	official:     make(map[xlanguage.Tag]string),
	capital:      make(map[xlanguage.Tag]string),
}

func init() { register(dataDrCongo) }

var DrCongo = dataDrCongo
