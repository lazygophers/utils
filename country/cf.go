//go:build country_africa || country_all || country_cf || country_middle_africa

package country

import (
	xlanguage "golang.org/x/text/language"

	"github.com/lazygophers/utils/currency"
)

// CentralAfricanRepublic — Central African Republic.
var dataCentralAfricanRepublic = &Country{
	alpha2:       "CF",
	alpha3:       "CAF",
	numeric:      140,
	callingCodes: []string{"+236"},
	timezones:    []string{"Africa/Bangui"},
	tlds:         []string{".cf"},
	officialLanguage:  xlanguage.French,
	spokenLanguages:   []xlanguage.Tag{xlanguage.French, xlanguage.MustParse("sg")},
	currency:     currency.XAF,
	region:       RegionMiddleAfrica,
	flagEmoji:    "🇨🇫",
	names:        make(map[xlanguage.Tag]string),
	official:     make(map[xlanguage.Tag]string),
	capital:      make(map[xlanguage.Tag]string),
}

func init() { register(dataCentralAfricanRepublic) }

var CentralAfricanRepublic = dataCentralAfricanRepublic
