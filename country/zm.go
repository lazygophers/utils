//go:build country_africa || country_all || country_eastern_africa || country_zm

package country

import (
	xlanguage "golang.org/x/text/language"

	"github.com/lazygophers/utils/currency"
)

// Zambia — Republic of Zambia.
var dataZambia = &Country{
	alpha2:       "ZM",
	alpha3:       "ZMB",
	numeric:      894,
	callingCodes: []string{"+260"},
	timezones:    []string{"Africa/Lusaka"},
	tlds:         []string{".zm"},
	officialLanguage:  xlanguage.English,
	spokenLanguages:   []xlanguage.Tag{xlanguage.English},
	currency:     currency.ZMW,
	region:       RegionEasternAfrica,
	flagEmoji:    "🇿🇲",
	names:        make(map[xlanguage.Tag]string),
	official:     make(map[xlanguage.Tag]string),
	capital:      make(map[xlanguage.Tag]string),
}

func init() { register(dataZambia) }

var Zambia = dataZambia
