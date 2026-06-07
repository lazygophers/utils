//go:build country_africa || country_all || country_cm || country_middle_africa

package country

import (
	xlanguage "golang.org/x/text/language"

	"github.com/lazygophers/utils/currency"
)

// Cameroon — Republic of Cameroon.
var dataCameroon = &Country{
	alpha2:       "CM",
	alpha3:       "CMR",
	numeric:      120,
	callingCodes: []string{"+237"},
	timezones:    []string{"Africa/Douala"},
	tlds:         []string{".cm"},
	officialLanguage:  xlanguage.French,
	spokenLanguages:   []xlanguage.Tag{xlanguage.French, xlanguage.English},
	currency:     currency.XAF,
	region:       RegionMiddleAfrica,
	flagEmoji:    "🇨🇲",
	names:        make(map[xlanguage.Tag]string),
	official:     make(map[xlanguage.Tag]string),
	capital:      make(map[xlanguage.Tag]string),
}

func init() { register(dataCameroon) }

var Cameroon = dataCameroon
