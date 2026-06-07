//go:build country_all || country_asia || country_south_eastern_asia || country_th

package country

import (
	xlanguage "golang.org/x/text/language"

	"github.com/lazygophers/utils/currency"
)

// Thailand — Kingdom of Thailand.
var dataThailand = &Country{
	alpha2:       "TH",
	alpha3:       "THA",
	numeric:      764,
	callingCodes: []string{"+66"},
	timezones:    []string{"Asia/Bangkok"},
	tlds:         []string{
		".th",
		".ไทย",
	},
	officialLanguage:  xlanguage.Thai,
	spokenLanguages:   []xlanguage.Tag{xlanguage.Thai},
	currency:     currency.THB,
	region:       RegionSouthEasternAsia,
	flagEmoji:    "🇹🇭",
	names:        make(map[xlanguage.Tag]string),
	official:     make(map[xlanguage.Tag]string),
	capital:      make(map[xlanguage.Tag]string),
}

func init() { register(dataThailand) }

var Thailand = dataThailand
