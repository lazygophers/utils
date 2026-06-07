//go:build country_africa || country_all || country_dj || country_eastern_africa

package country

import (
	xlanguage "golang.org/x/text/language"

	"github.com/lazygophers/utils/currency"
)

// Djibouti — Republic of Djibouti.
var dataDjibouti = &Country{
	alpha2:       "DJ",
	alpha3:       "DJI",
	numeric:      262,
	callingCodes: []string{"+253"},
	timezones:    []string{"Africa/Djibouti"},
	tlds:         []string{".dj"},
	officialLanguage:  xlanguage.French,
	spokenLanguages:   []xlanguage.Tag{xlanguage.French, xlanguage.Arabic},
	currency:     currency.DJF,
	region:       RegionEasternAfrica,
	flagEmoji:    "🇩🇯",
	names:        make(map[xlanguage.Tag]string),
	official:     make(map[xlanguage.Tag]string),
	capital:      make(map[xlanguage.Tag]string),
}

func init() { register(dataDjibouti) }

var Djibouti = dataDjibouti
