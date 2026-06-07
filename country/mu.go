//go:build country_africa || country_all || country_eastern_africa || country_mu

package country

import (
	xlanguage "golang.org/x/text/language"

	"github.com/lazygophers/utils/currency"
)

// Mauritius — Republic of Mauritius.
var dataMauritius = &Country{
	alpha2:       "MU",
	alpha3:       "MUS",
	numeric:      480,
	callingCodes: []string{"+230"},
	timezones:    []string{"Indian/Mauritius"},
	tlds:         []string{".mu"},
	officialLanguage:  xlanguage.English,
	spokenLanguages:   []xlanguage.Tag{xlanguage.English, xlanguage.French},
	currency:     currency.MUR,
	region:       RegionEasternAfrica,
	flagEmoji:    "🇲🇺",
	names:        make(map[xlanguage.Tag]string),
	official:     make(map[xlanguage.Tag]string),
	capital:      make(map[xlanguage.Tag]string),
}

func init() { register(dataMauritius) }

var Mauritius = dataMauritius
